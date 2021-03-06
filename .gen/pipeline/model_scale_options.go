/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type ScaleOptions struct {
	Enabled bool `json:"enabled"`
	DesiredCpu float64 `json:"desiredCpu,omitempty"`
	DesiredMem float64 `json:"desiredMem,omitempty"`
	DesiredGpu int32 `json:"desiredGpu,omitempty"`
	OnDemandPct int32 `json:"onDemandPct,omitempty"`
	Excludes []string `json:"excludes,omitempty"`
	KeepDesiredCapacity bool `json:"keepDesiredCapacity,omitempty"`
}
