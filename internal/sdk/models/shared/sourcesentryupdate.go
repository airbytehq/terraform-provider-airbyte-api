// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceSentryUpdate struct {
	// Log into Sentry and then <a href="https://sentry.io/settings/account/api/auth-tokens/">create authentication tokens</a>.For self-hosted, you can find or create authentication tokens by visiting "{instance_url_prefix}/settings/account/api/auth-tokens/"
	AuthToken string `json:"auth_token"`
	// Fields to retrieve when fetching discover events
	DiscoverFields []any `json:"discover_fields,omitempty"`
	// Host name of Sentry API server.For self-hosted, specify your host name here. Otherwise, leave it empty.
	Hostname *string `default:"sentry.io" json:"hostname"`
	// The slug of the organization the groups belong to.
	Organization string `json:"organization"`
	// The name (slug) of the Project you want to sync.
	Project string `json:"project"`
}

func (s SourceSentryUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSentryUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSentryUpdate) GetAuthToken() string {
	if o == nil {
		return ""
	}
	return o.AuthToken
}

func (o *SourceSentryUpdate) GetDiscoverFields() []any {
	if o == nil {
		return nil
	}
	return o.DiscoverFields
}

func (o *SourceSentryUpdate) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *SourceSentryUpdate) GetOrganization() string {
	if o == nil {
		return ""
	}
	return o.Organization
}

func (o *SourceSentryUpdate) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}
