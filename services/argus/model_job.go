/*
STACKIT Argus API

API endpoints for Argus on STACKIT

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package argus

type Job struct {
	BasicAuth       *BasicAuth       `json:"basicAuth,omitempty"`
	BearerToken     *string          `json:"bearerToken,omitempty"`
	HonorLabels     *bool            `json:"honorLabels,omitempty"`
	HonorTimeStamps *bool            `json:"honorTimeStamps,omitempty"`
	HttpSdConfigs   *[]HTTPServiceSD `json:"httpSdConfigs,omitempty"`
	// REQUIRED
	JobName               *string                 `json:"jobName"`
	MetricsPath           *string                 `json:"metricsPath,omitempty"`
	MetricsRelabelConfigs *[]MetricsRelabelConfig `json:"metricsRelabelConfigs,omitempty"`
	Oauth2                *OAuth2                 `json:"oauth2,omitempty"`
	Params                *map[string][]string    `json:"params,omitempty"`
	SampleLimit           *int32                  `json:"sampleLimit,omitempty"`
	Scheme                *string                 `json:"scheme,omitempty"`
	// REQUIRED
	ScrapeInterval *string `json:"scrapeInterval"`
	// REQUIRED
	ScrapeTimeout *string `json:"scrapeTimeout"`
	// REQUIRED
	StaticConfigs *[]StaticConfigs `json:"staticConfigs"`
	TlsConfig     *TLSConfig       `json:"tlsConfig,omitempty"`
}
