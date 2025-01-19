package config

import (
	"PowerX/internal/config/openapiplatform"
	"PowerX/internal/config/openapiprovider"
)

type OpenAPI struct {
	Platforms struct {
		BrainX openapiPlatform.BrainX
	}
	Providers struct {
		BrainX openapiProvider.BrainX
	}
}
