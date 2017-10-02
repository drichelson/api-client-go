/* 
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@launchdarkly.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type FeatureFlag struct {

	Key string `json:"key,omitempty"`

	Name string `json:"name,omitempty"`

	Kind string `json:"kind,omitempty"`

	CreationDate float32 `json:"creationDate,omitempty"`

	IncludeInSnippet bool `json:"includeInSnippet,omitempty"`

	Temporary bool `json:"temporary,omitempty"`

	MaintainerId string `json:"maintainerId,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Variations []Variation `json:"variations,omitempty"`

	Links Links `json:"_links,omitempty"`

	Maintainer Member `json:"_maintainer,omitempty"`

	Environments map[string]FeatureFlagConfig `json:"environments,omitempty"`
}