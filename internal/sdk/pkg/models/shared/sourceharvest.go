// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceHarvestSchemasAuthType string

const (
	SourceHarvestSchemasAuthTypeToken SourceHarvestSchemasAuthType = "Token"
)

func (e SourceHarvestSchemasAuthType) ToPointer() *SourceHarvestSchemasAuthType {
	return &e
}

func (e *SourceHarvestSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceHarvestSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestSchemasAuthType: %v", v)
	}
}

type SourceHarvestAuthenticateWithPersonalAccessToken struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// Log into Harvest and then create new <a href="https://id.getharvest.com/developers"> personal access token</a>.
	APIToken string                        `json:"api_token"`
	authType *SourceHarvestSchemasAuthType `const:"Token" json:"auth_type,omitempty"`
}

func (s SourceHarvestAuthenticateWithPersonalAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHarvestAuthenticateWithPersonalAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceHarvestAuthenticateWithPersonalAccessToken) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceHarvestAuthenticateWithPersonalAccessToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceHarvestAuthenticateWithPersonalAccessToken) GetAuthType() *SourceHarvestSchemasAuthType {
	return SourceHarvestSchemasAuthTypeToken.ToPointer()
}

type SourceHarvestAuthType string

const (
	SourceHarvestAuthTypeClient SourceHarvestAuthType = "Client"
)

func (e SourceHarvestAuthType) ToPointer() *SourceHarvestAuthType {
	return &e
}

func (e *SourceHarvestAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceHarvestAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestAuthType: %v", v)
	}
}

type SourceHarvestAuthenticateViaHarvestOAuth struct {
	AdditionalProperties interface{}            `additionalProperties:"true" json:"-"`
	authType             *SourceHarvestAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Harvest developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Harvest developer application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceHarvestAuthenticateViaHarvestOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHarvestAuthenticateViaHarvestOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceHarvestAuthenticateViaHarvestOAuth) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceHarvestAuthenticateViaHarvestOAuth) GetAuthType() *SourceHarvestAuthType {
	return SourceHarvestAuthTypeClient.ToPointer()
}

func (o *SourceHarvestAuthenticateViaHarvestOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceHarvestAuthenticateViaHarvestOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceHarvestAuthenticateViaHarvestOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceHarvestAuthenticationMechanismType string

const (
	SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticateViaHarvestOAuth         SourceHarvestAuthenticationMechanismType = "source-harvest_Authenticate via Harvest (OAuth)"
	SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticateWithPersonalAccessToken SourceHarvestAuthenticationMechanismType = "source-harvest_Authenticate with Personal Access Token"
)

// SourceHarvestAuthenticationMechanism - Choose how to authenticate to Harvest.
type SourceHarvestAuthenticationMechanism struct {
	SourceHarvestAuthenticateViaHarvestOAuth         *SourceHarvestAuthenticateViaHarvestOAuth
	SourceHarvestAuthenticateWithPersonalAccessToken *SourceHarvestAuthenticateWithPersonalAccessToken

	Type SourceHarvestAuthenticationMechanismType
}

func CreateSourceHarvestAuthenticationMechanismSourceHarvestAuthenticateViaHarvestOAuth(sourceHarvestAuthenticateViaHarvestOAuth SourceHarvestAuthenticateViaHarvestOAuth) SourceHarvestAuthenticationMechanism {
	typ := SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticateViaHarvestOAuth

	return SourceHarvestAuthenticationMechanism{
		SourceHarvestAuthenticateViaHarvestOAuth: &sourceHarvestAuthenticateViaHarvestOAuth,
		Type:                                     typ,
	}
}

func CreateSourceHarvestAuthenticationMechanismSourceHarvestAuthenticateWithPersonalAccessToken(sourceHarvestAuthenticateWithPersonalAccessToken SourceHarvestAuthenticateWithPersonalAccessToken) SourceHarvestAuthenticationMechanism {
	typ := SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticateWithPersonalAccessToken

	return SourceHarvestAuthenticationMechanism{
		SourceHarvestAuthenticateWithPersonalAccessToken: &sourceHarvestAuthenticateWithPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceHarvestAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	sourceHarvestAuthenticateWithPersonalAccessToken := new(SourceHarvestAuthenticateWithPersonalAccessToken)
	if err := utils.UnmarshalJSON(data, &sourceHarvestAuthenticateWithPersonalAccessToken, "", true, true); err == nil {
		u.SourceHarvestAuthenticateWithPersonalAccessToken = sourceHarvestAuthenticateWithPersonalAccessToken
		u.Type = SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticateWithPersonalAccessToken
		return nil
	}

	sourceHarvestAuthenticateViaHarvestOAuth := new(SourceHarvestAuthenticateViaHarvestOAuth)
	if err := utils.UnmarshalJSON(data, &sourceHarvestAuthenticateViaHarvestOAuth, "", true, true); err == nil {
		u.SourceHarvestAuthenticateViaHarvestOAuth = sourceHarvestAuthenticateViaHarvestOAuth
		u.Type = SourceHarvestAuthenticationMechanismTypeSourceHarvestAuthenticateViaHarvestOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceHarvestAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceHarvestAuthenticateViaHarvestOAuth != nil {
		return utils.MarshalJSON(u.SourceHarvestAuthenticateViaHarvestOAuth, "", true)
	}

	if u.SourceHarvestAuthenticateWithPersonalAccessToken != nil {
		return utils.MarshalJSON(u.SourceHarvestAuthenticateWithPersonalAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Harvest string

const (
	HarvestHarvest Harvest = "harvest"
)

func (e Harvest) ToPointer() *Harvest {
	return &e
}

func (e *Harvest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "harvest":
		*e = Harvest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Harvest: %v", v)
	}
}

type SourceHarvest struct {
	// Harvest account ID. Required for all Harvest requests in pair with Personal Access Token
	AccountID string `json:"account_id"`
	// Choose how to authenticate to Harvest.
	Credentials *SourceHarvestAuthenticationMechanism `json:"credentials,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
	ReplicationEndDate *time.Time `json:"replication_end_date,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	ReplicationStartDate time.Time `json:"replication_start_date"`
	sourceType           Harvest   `const:"harvest" json:"sourceType"`
}

func (s SourceHarvest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHarvest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceHarvest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *SourceHarvest) GetCredentials() *SourceHarvestAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceHarvest) GetReplicationEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ReplicationEndDate
}

func (o *SourceHarvest) GetReplicationStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ReplicationStartDate
}

func (o *SourceHarvest) GetSourceType() Harvest {
	return HarvestHarvest
}
