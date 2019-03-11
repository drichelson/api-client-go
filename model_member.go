/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.14
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Member struct {
	Links *Links `json:"_links,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	Role *Role `json:"role,omitempty"`
	Email string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName string `json:"lastName,omitempty"`
	PendingInvite bool `json:"_pendingInvite,omitempty"`
	IsBeta bool `json:"isBeta,omitempty"`
	CustomRoles []string `json:"customRoles,omitempty"`
}
