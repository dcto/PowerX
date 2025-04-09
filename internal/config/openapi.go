package config

import (
	"PowerX/internal/config/openapiPlatform"
	"PowerX/internal/config/openapiProvider"
)

type OpenAPI struct {
	Platforms struct {
		BrainX openapiPlatform.BrainX
	}
	Providers struct {
		BrainX openapiProvider.BrainX
	}
}
