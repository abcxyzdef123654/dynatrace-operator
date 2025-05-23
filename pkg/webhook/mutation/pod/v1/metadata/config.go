package metadata

import (
	"github.com/Dynatrace/dynatrace-operator/pkg/logd"
)

const (
	workloadEnrichmentVolumeName = "metadata-enrichment"
	ingestEndpointVolumeName     = "metadata-enrichment-endpoint"
	enrichmentEndpointPath       = "/var/lib/dynatrace/enrichment/endpoint"
)

var (
	log = logd.Get().WithName("v1-pod-mutation-metadata-enrichment")
)
