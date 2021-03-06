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

type HelmChartDetailsResponse struct {
	Name string `json:"name,omitempty"`
	Repo string `json:"repo,omitempty"`
	Versions []HelmChartDetailsResponseVersions `json:"versions,omitempty"`
}
