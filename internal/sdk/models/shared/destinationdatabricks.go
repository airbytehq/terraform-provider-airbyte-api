// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationDatabricksSchemasAuthType string

const (
	DestinationDatabricksSchemasAuthTypeBasic DestinationDatabricksSchemasAuthType = "BASIC"
)

func (e DestinationDatabricksSchemasAuthType) ToPointer() *DestinationDatabricksSchemasAuthType {
	return &e
}
func (e *DestinationDatabricksSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BASIC":
		*e = DestinationDatabricksSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksSchemasAuthType: %v", v)
	}
}

type DestinationDatabricksPersonalAccessToken struct {
	authType            DestinationDatabricksSchemasAuthType `const:"BASIC" json:"auth_type"`
	PersonalAccessToken string                               `json:"personal_access_token"`
}

func (d DestinationDatabricksPersonalAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksPersonalAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksPersonalAccessToken) GetAuthType() DestinationDatabricksSchemasAuthType {
	return DestinationDatabricksSchemasAuthTypeBasic
}

func (o *DestinationDatabricksPersonalAccessToken) GetPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.PersonalAccessToken
}

type DestinationDatabricksAuthType string

const (
	DestinationDatabricksAuthTypeOauth DestinationDatabricksAuthType = "OAUTH"
)

func (e DestinationDatabricksAuthType) ToPointer() *DestinationDatabricksAuthType {
	return &e
}
func (e *DestinationDatabricksAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAUTH":
		*e = DestinationDatabricksAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksAuthType: %v", v)
	}
}

type DestinationDatabricksOAuth2Recommended struct {
	authType DestinationDatabricksAuthType `const:"OAUTH" json:"auth_type"`
	ClientID string                        `json:"client_id"`
	Secret   string                        `json:"secret"`
}

func (d DestinationDatabricksOAuth2Recommended) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksOAuth2Recommended) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksOAuth2Recommended) GetAuthType() DestinationDatabricksAuthType {
	return DestinationDatabricksAuthTypeOauth
}

func (o *DestinationDatabricksOAuth2Recommended) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *DestinationDatabricksOAuth2Recommended) GetSecret() string {
	if o == nil {
		return ""
	}
	return o.Secret
}

type DestinationDatabricksAuthenticationType string

const (
	DestinationDatabricksAuthenticationTypeDestinationDatabricksOAuth2Recommended   DestinationDatabricksAuthenticationType = "destination-databricks_OAuth2 (Recommended)"
	DestinationDatabricksAuthenticationTypeDestinationDatabricksPersonalAccessToken DestinationDatabricksAuthenticationType = "destination-databricks_Personal Access Token"
)

// DestinationDatabricksAuthentication - Authentication mechanism for Staging files and running queries
type DestinationDatabricksAuthentication struct {
	DestinationDatabricksOAuth2Recommended   *DestinationDatabricksOAuth2Recommended
	DestinationDatabricksPersonalAccessToken *DestinationDatabricksPersonalAccessToken

	Type DestinationDatabricksAuthenticationType
}

func CreateDestinationDatabricksAuthenticationDestinationDatabricksOAuth2Recommended(destinationDatabricksOAuth2Recommended DestinationDatabricksOAuth2Recommended) DestinationDatabricksAuthentication {
	typ := DestinationDatabricksAuthenticationTypeDestinationDatabricksOAuth2Recommended

	return DestinationDatabricksAuthentication{
		DestinationDatabricksOAuth2Recommended: &destinationDatabricksOAuth2Recommended,
		Type:                                   typ,
	}
}

func CreateDestinationDatabricksAuthenticationDestinationDatabricksPersonalAccessToken(destinationDatabricksPersonalAccessToken DestinationDatabricksPersonalAccessToken) DestinationDatabricksAuthentication {
	typ := DestinationDatabricksAuthenticationTypeDestinationDatabricksPersonalAccessToken

	return DestinationDatabricksAuthentication{
		DestinationDatabricksPersonalAccessToken: &destinationDatabricksPersonalAccessToken,
		Type:                                     typ,
	}
}

func (u *DestinationDatabricksAuthentication) UnmarshalJSON(data []byte) error {

	var destinationDatabricksPersonalAccessToken DestinationDatabricksPersonalAccessToken = DestinationDatabricksPersonalAccessToken{}
	if err := utils.UnmarshalJSON(data, &destinationDatabricksPersonalAccessToken, "", true, true); err == nil {
		u.DestinationDatabricksPersonalAccessToken = &destinationDatabricksPersonalAccessToken
		u.Type = DestinationDatabricksAuthenticationTypeDestinationDatabricksPersonalAccessToken
		return nil
	}

	var destinationDatabricksOAuth2Recommended DestinationDatabricksOAuth2Recommended = DestinationDatabricksOAuth2Recommended{}
	if err := utils.UnmarshalJSON(data, &destinationDatabricksOAuth2Recommended, "", true, true); err == nil {
		u.DestinationDatabricksOAuth2Recommended = &destinationDatabricksOAuth2Recommended
		u.Type = DestinationDatabricksAuthenticationTypeDestinationDatabricksOAuth2Recommended
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationDatabricksAuthentication", string(data))
}

func (u DestinationDatabricksAuthentication) MarshalJSON() ([]byte, error) {
	if u.DestinationDatabricksOAuth2Recommended != nil {
		return utils.MarshalJSON(u.DestinationDatabricksOAuth2Recommended, "", true)
	}

	if u.DestinationDatabricksPersonalAccessToken != nil {
		return utils.MarshalJSON(u.DestinationDatabricksPersonalAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationDatabricksAuthentication: all fields are null")
}

type Databricks string

const (
	DatabricksDatabricks Databricks = "databricks"
)

func (e Databricks) ToPointer() *Databricks {
	return &e
}
func (e *Databricks) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "databricks":
		*e = Databricks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Databricks: %v", v)
	}
}

type DestinationDatabricks struct {
	// You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.
	AcceptTerms *bool `default:"false" json:"accept_terms"`
	// Authentication mechanism for Staging files and running queries
	Authentication DestinationDatabricksAuthentication `json:"authentication"`
	// The name of the unity catalog for the database
	Database        string     `json:"database"`
	destinationType Databricks `const:"databricks" json:"destinationType"`
	// Databricks Cluster Server Hostname.
	Hostname string `json:"hostname"`
	// Databricks Cluster HTTP Path.
	HTTPPath string `json:"http_path"`
	// Databricks Cluster Port.
	Port *string `default:"443" json:"port"`
	// Default to 'true'. Switch it to 'false' for debugging purpose.
	PurgeStagingData *bool `default:"true" json:"purge_staging_data"`
	// The schema to write raw tables into (default: airbyte_internal)
	RawSchemaOverride *string `default:"airbyte_internal" json:"raw_schema_override"`
	// The default schema tables are written. If not specified otherwise, the "default" will be used.
	Schema *string `default:"default" json:"schema"`
}

func (d DestinationDatabricks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricks) GetAcceptTerms() *bool {
	if o == nil {
		return nil
	}
	return o.AcceptTerms
}

func (o *DestinationDatabricks) GetAuthentication() DestinationDatabricksAuthentication {
	if o == nil {
		return DestinationDatabricksAuthentication{}
	}
	return o.Authentication
}

func (o *DestinationDatabricks) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationDatabricks) GetDestinationType() Databricks {
	return DatabricksDatabricks
}

func (o *DestinationDatabricks) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *DestinationDatabricks) GetHTTPPath() string {
	if o == nil {
		return ""
	}
	return o.HTTPPath
}

func (o *DestinationDatabricks) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationDatabricks) GetPurgeStagingData() *bool {
	if o == nil {
		return nil
	}
	return o.PurgeStagingData
}

func (o *DestinationDatabricks) GetRawSchemaOverride() *string {
	if o == nil {
		return nil
	}
	return o.RawSchemaOverride
}

func (o *DestinationDatabricks) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
