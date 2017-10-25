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

type FeatureFlagConfig struct {

	On bool `json:"on,omitempty"`

	Archived bool `json:"archived,omitempty"`

	Salt string `json:"salt,omitempty"`

	Sel string `json:"sel,omitempty"`

	LastModified int64 `json:"lastModified,omitempty"`

	Version int32 `json:"version,omitempty"`

	Targets []Target `json:"targets,omitempty"`

	Rules []Rule `json:"rules,omitempty"`

	Fallthrough_ FeatureFlagConfigFallthrough `json:"fallthrough,omitempty"`
}
