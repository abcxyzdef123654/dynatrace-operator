package dynakube

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OneAgentMode string

type OneAgentSpec struct {
	// Has a single OneAgent per node via DaemonSet.
	// Injection is performed via the same OneAgent DaemonSet.
	// +nullable
	ClassicFullStack *HostInjectSpec `json:"classicFullStack,omitempty"`

	// Has a single OneAgent per node via DaemonSet.
	// dynatrace-webhook injects into application pods based on labeled namespaces.
	// Has a CSI driver per node via DaemonSet to provide binaries to pods.
	// +nullable
	CloudNativeFullStack *CloudNativeFullStackSpec `json:"cloudNativeFullStack,omitempty"`

	// dynatrace-webhook injects into application pods based on labeled namespaces.
	// Has an optional CSI driver per node via DaemonSet to provide binaries to pods.
	// +nullable
	ApplicationMonitoring *ApplicationMonitoringSpec `json:"applicationMonitoring,omitempty"`

	// Has a single OneAgent per node via DaemonSet.
	// Doesn't inject into application pods.
	// +nullable
	HostMonitoring *HostInjectSpec `json:"hostMonitoring,omitempty"`

	// Sets a host group for OneAgent.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Host Group",order=5,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	HostGroup string `json:"hostGroup,omitempty"`
}

type CloudNativeFullStackSpec struct {
	HostInjectSpec   `json:",inline"`
	AppInjectionSpec `json:",inline"`
}

type HostInjectSpec struct {

	// Add custom OneAgent annotations.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Annotations",order=27,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	Annotations map[string]string `json:"annotations,omitempty"`

	// Your defined labels for OneAgent pods in order to structure workloads as desired.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Labels",order=26,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	Labels map[string]string `json:"labels,omitempty"`

	// Specify the node selector that controls on which nodes OneAgent will be deployed.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Node Selector",order=17,xDescriptors="urn:alm:descriptor:com.tectonic.ui:selector:Node"
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Disables automatic restarts of OneAgent pods in case a new version is available (https://www.dynatrace.com/support/help/setup-and-configuration/setup-on-container-platforms/kubernetes/get-started-with-kubernetes-monitoring#disable-auto).
	// Enabled by default.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Automatically update Agent",order=13,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:booleanSwitch"}
	AutoUpdate *bool `json:"autoUpdate"`

	// The OneAgent version to be used.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OneAgent version",order=11,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	Version string `json:"version,omitempty"`

	// Use a custom OneAgent Docker image. Defaults to the image from the Dynatrace cluster.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Image",order=12,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	Image string `json:"image,omitempty"`

	// Set the DNS Policy for OneAgent pods. For details, see Pods DNS Policy (https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy).
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="DNS Policy",order=24,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	DNSPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// Assign a priority class to the OneAgent pods. By default, no class is set.
	// For details, see Pod Priority and Preemption (https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/).
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Priority Class name",order=23,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:io.kubernetes:PriorityClass"}
	PriorityClassName string `json:"priorityClassName,omitempty"`

	// The SecComp Profile that will be configured in order to run in secure computing mode.
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OneAgent SecComp Profile",order=17,xDescriptors="urn:alm:descriptor:com.tectonic.ui:selector:core:v1:Namespace"
	SecCompProfile string `json:"secCompProfile,omitempty"`

	// Resource settings for OneAgent container. Consumption of the OneAgent heavily depends on the workload to monitor. You can use the default settings in the CR.
	// Note: resource.requests shows the values needed to run; resource.limits shows the maximum limits for the pod.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements",order=20,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	OneAgentResources corev1.ResourceRequirements `json:"oneAgentResources,omitempty"`

	// Tolerations to include with the OneAgent DaemonSet. For details, see Taints and Tolerations (https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/).
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Tolerations",order=18,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:hidden"}
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// Set additional environment variables for the OneAgent pods.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OneAgent environment variable installer arguments",order=22,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:hidden"}
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Set additional arguments to the OneAgent installer.
	// For available options, see Linux custom installation (https://www.dynatrace.com/support/help/setup-and-configuration/dynatrace-oneagent/installation-and-operation/linux/installation/customize-oneagent-installation-on-linux).
	// For the list of limitations, see Limitations (https://www.dynatrace.com/support/help/setup-and-configuration/setup-on-container-platforms/docker/set-up-dynatrace-oneagent-as-docker-container#limitations).
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OneAgent installer arguments",order=21,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:hidden"}
	// +listType=set
	Args []string `json:"args,omitempty"`
}

type ApplicationMonitoringSpec struct {

	// Set if you want to use the CSIDriver. Don't enable it if you do not have access to Kubernetes nodes or if you lack privileges.
	// +kubebuilder:validation:Optional
	UseCSIDriver *bool `json:"useCSIDriver,omitempty"`
	// The OneAgent version to be used.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="OneAgent version",order=11,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	Version string `json:"version,omitempty"`

	AppInjectionSpec `json:",inline"`
}

type AppInjectionSpec struct {
	// Define resources requests and limits for the initContainer. For details, see Managing resources for containers
	// (https://kubernetes.io/docs/concepts/configuration/manage-resources-containers).
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Resource Requirements",order=15,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:resourceRequirements"}
	InitResources *corev1.ResourceRequirements `json:"initResources,omitempty"`

	// The OneAgent image that is used to inject into Pods.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="CodeModulesImage",order=12,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:advanced","urn:alm:descriptor:com.tectonic.ui:text"}
	CodeModulesImage string `json:"codeModulesImage,omitempty"`

	// Applicable only for applicationMonitoring or cloudNativeFullStack configuration types. The namespaces where you want Dynatrace Operator to inject.
	// For more information, see Configure monitoring for namespaces and pods (https://www.dynatrace.com/support/help/setup-and-configuration/setup-on-container-platforms/kubernetes/get-started-with-kubernetes-monitoring/dto-config-options-k8s#annotate).
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="Namespace Selector",order=17,xDescriptors="urn:alm:descriptor:com.tectonic.ui:selector:core:v1:Namespace"
	NamespaceSelector metav1.LabelSelector `json:"namespaceSelector,omitempty"`
}
