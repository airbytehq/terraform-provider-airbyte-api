// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceDynamodbUpdateSchemasAuthType string

const (
	SourceDynamodbUpdateSchemasAuthTypeRole SourceDynamodbUpdateSchemasAuthType = "Role"
)

func (e SourceDynamodbUpdateSchemasAuthType) ToPointer() *SourceDynamodbUpdateSchemasAuthType {
	return &e
}
func (e *SourceDynamodbUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Role":
		*e = SourceDynamodbUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDynamodbUpdateSchemasAuthType: %v", v)
	}
}

type RoleBasedAuthentication struct {
	AdditionalProperties any                                  `additionalProperties:"true" json:"-"`
	authType             *SourceDynamodbUpdateSchemasAuthType `const:"Role" json:"auth_type,omitempty"`
}

func (r RoleBasedAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RoleBasedAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *RoleBasedAuthentication) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *RoleBasedAuthentication) GetAuthType() *SourceDynamodbUpdateSchemasAuthType {
	return SourceDynamodbUpdateSchemasAuthTypeRole.ToPointer()
}

type SourceDynamodbUpdateAuthType string

const (
	SourceDynamodbUpdateAuthTypeUser SourceDynamodbUpdateAuthType = "User"
)

func (e SourceDynamodbUpdateAuthType) ToPointer() *SourceDynamodbUpdateAuthType {
	return &e
}
func (e *SourceDynamodbUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "User":
		*e = SourceDynamodbUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDynamodbUpdateAuthType: %v", v)
	}
}

type AuthenticateViaAccessKeys struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The access key id to access Dynamodb. Airbyte requires read permissions to the database
	AccessKeyID string                        `json:"access_key_id"`
	authType    *SourceDynamodbUpdateAuthType `const:"User" json:"auth_type,omitempty"`
	// The corresponding secret to the access key id.
	SecretAccessKey string `json:"secret_access_key"`
}

func (a AuthenticateViaAccessKeys) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaAccessKeys) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaAccessKeys) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AuthenticateViaAccessKeys) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *AuthenticateViaAccessKeys) GetAuthType() *SourceDynamodbUpdateAuthType {
	return SourceDynamodbUpdateAuthTypeUser.ToPointer()
}

func (o *AuthenticateViaAccessKeys) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

type CredentialsType string

const (
	CredentialsTypeAuthenticateViaAccessKeys CredentialsType = "Authenticate via Access Keys"
	CredentialsTypeRoleBasedAuthentication   CredentialsType = "Role Based Authentication"
)

// Credentials for the service
type Credentials struct {
	AuthenticateViaAccessKeys *AuthenticateViaAccessKeys
	RoleBasedAuthentication   *RoleBasedAuthentication

	Type CredentialsType
}

func CreateCredentialsAuthenticateViaAccessKeys(authenticateViaAccessKeys AuthenticateViaAccessKeys) Credentials {
	typ := CredentialsTypeAuthenticateViaAccessKeys

	return Credentials{
		AuthenticateViaAccessKeys: &authenticateViaAccessKeys,
		Type:                      typ,
	}
}

func CreateCredentialsRoleBasedAuthentication(roleBasedAuthentication RoleBasedAuthentication) Credentials {
	typ := CredentialsTypeRoleBasedAuthentication

	return Credentials{
		RoleBasedAuthentication: &roleBasedAuthentication,
		Type:                    typ,
	}
}

func (u *Credentials) UnmarshalJSON(data []byte) error {

	var roleBasedAuthentication RoleBasedAuthentication = RoleBasedAuthentication{}
	if err := utils.UnmarshalJSON(data, &roleBasedAuthentication, "", true, true); err == nil {
		u.RoleBasedAuthentication = &roleBasedAuthentication
		u.Type = CredentialsTypeRoleBasedAuthentication
		return nil
	}

	var authenticateViaAccessKeys AuthenticateViaAccessKeys = AuthenticateViaAccessKeys{}
	if err := utils.UnmarshalJSON(data, &authenticateViaAccessKeys, "", true, true); err == nil {
		u.AuthenticateViaAccessKeys = &authenticateViaAccessKeys
		u.Type = CredentialsTypeAuthenticateViaAccessKeys
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Credentials", string(data))
}

func (u Credentials) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaAccessKeys != nil {
		return utils.MarshalJSON(u.AuthenticateViaAccessKeys, "", true)
	}

	if u.RoleBasedAuthentication != nil {
		return utils.MarshalJSON(u.RoleBasedAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type Credentials: all fields are null")
}

// SourceDynamodbUpdateDynamodbRegion - The region of the Dynamodb database
type SourceDynamodbUpdateDynamodbRegion string

const (
	SourceDynamodbUpdateDynamodbRegionUnknown      SourceDynamodbUpdateDynamodbRegion = ""
	SourceDynamodbUpdateDynamodbRegionAfSouth1     SourceDynamodbUpdateDynamodbRegion = "af-south-1"
	SourceDynamodbUpdateDynamodbRegionApEast1      SourceDynamodbUpdateDynamodbRegion = "ap-east-1"
	SourceDynamodbUpdateDynamodbRegionApNortheast1 SourceDynamodbUpdateDynamodbRegion = "ap-northeast-1"
	SourceDynamodbUpdateDynamodbRegionApNortheast2 SourceDynamodbUpdateDynamodbRegion = "ap-northeast-2"
	SourceDynamodbUpdateDynamodbRegionApNortheast3 SourceDynamodbUpdateDynamodbRegion = "ap-northeast-3"
	SourceDynamodbUpdateDynamodbRegionApSouth1     SourceDynamodbUpdateDynamodbRegion = "ap-south-1"
	SourceDynamodbUpdateDynamodbRegionApSouth2     SourceDynamodbUpdateDynamodbRegion = "ap-south-2"
	SourceDynamodbUpdateDynamodbRegionApSoutheast1 SourceDynamodbUpdateDynamodbRegion = "ap-southeast-1"
	SourceDynamodbUpdateDynamodbRegionApSoutheast2 SourceDynamodbUpdateDynamodbRegion = "ap-southeast-2"
	SourceDynamodbUpdateDynamodbRegionApSoutheast3 SourceDynamodbUpdateDynamodbRegion = "ap-southeast-3"
	SourceDynamodbUpdateDynamodbRegionApSoutheast4 SourceDynamodbUpdateDynamodbRegion = "ap-southeast-4"
	SourceDynamodbUpdateDynamodbRegionCaCentral1   SourceDynamodbUpdateDynamodbRegion = "ca-central-1"
	SourceDynamodbUpdateDynamodbRegionCaWest1      SourceDynamodbUpdateDynamodbRegion = "ca-west-1"
	SourceDynamodbUpdateDynamodbRegionCnNorth1     SourceDynamodbUpdateDynamodbRegion = "cn-north-1"
	SourceDynamodbUpdateDynamodbRegionCnNorthwest1 SourceDynamodbUpdateDynamodbRegion = "cn-northwest-1"
	SourceDynamodbUpdateDynamodbRegionEuCentral1   SourceDynamodbUpdateDynamodbRegion = "eu-central-1"
	SourceDynamodbUpdateDynamodbRegionEuCentral2   SourceDynamodbUpdateDynamodbRegion = "eu-central-2"
	SourceDynamodbUpdateDynamodbRegionEuNorth1     SourceDynamodbUpdateDynamodbRegion = "eu-north-1"
	SourceDynamodbUpdateDynamodbRegionEuSouth1     SourceDynamodbUpdateDynamodbRegion = "eu-south-1"
	SourceDynamodbUpdateDynamodbRegionEuSouth2     SourceDynamodbUpdateDynamodbRegion = "eu-south-2"
	SourceDynamodbUpdateDynamodbRegionEuWest1      SourceDynamodbUpdateDynamodbRegion = "eu-west-1"
	SourceDynamodbUpdateDynamodbRegionEuWest2      SourceDynamodbUpdateDynamodbRegion = "eu-west-2"
	SourceDynamodbUpdateDynamodbRegionEuWest3      SourceDynamodbUpdateDynamodbRegion = "eu-west-3"
	SourceDynamodbUpdateDynamodbRegionIlCentral1   SourceDynamodbUpdateDynamodbRegion = "il-central-1"
	SourceDynamodbUpdateDynamodbRegionMeCentral1   SourceDynamodbUpdateDynamodbRegion = "me-central-1"
	SourceDynamodbUpdateDynamodbRegionMeSouth1     SourceDynamodbUpdateDynamodbRegion = "me-south-1"
	SourceDynamodbUpdateDynamodbRegionSaEast1      SourceDynamodbUpdateDynamodbRegion = "sa-east-1"
	SourceDynamodbUpdateDynamodbRegionUsEast1      SourceDynamodbUpdateDynamodbRegion = "us-east-1"
	SourceDynamodbUpdateDynamodbRegionUsEast2      SourceDynamodbUpdateDynamodbRegion = "us-east-2"
	SourceDynamodbUpdateDynamodbRegionUsGovEast1   SourceDynamodbUpdateDynamodbRegion = "us-gov-east-1"
	SourceDynamodbUpdateDynamodbRegionUsGovWest1   SourceDynamodbUpdateDynamodbRegion = "us-gov-west-1"
	SourceDynamodbUpdateDynamodbRegionUsWest1      SourceDynamodbUpdateDynamodbRegion = "us-west-1"
	SourceDynamodbUpdateDynamodbRegionUsWest2      SourceDynamodbUpdateDynamodbRegion = "us-west-2"
)

func (e SourceDynamodbUpdateDynamodbRegion) ToPointer() *SourceDynamodbUpdateDynamodbRegion {
	return &e
}
func (e *SourceDynamodbUpdateDynamodbRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-south-2":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ap-southeast-3":
		fallthrough
	case "ap-southeast-4":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "ca-west-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-central-2":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-south-2":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "il-central-1":
		fallthrough
	case "me-central-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		*e = SourceDynamodbUpdateDynamodbRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDynamodbUpdateDynamodbRegion: %v", v)
	}
}

type SourceDynamodbUpdate struct {
	// Credentials for the service
	Credentials *Credentials `json:"credentials,omitempty"`
	// the URL of the Dynamodb database
	Endpoint *string `default:"" json:"endpoint"`
	// Ignore tables with missing scan/read permissions
	IgnoreMissingReadPermissionsTables *bool `default:"false" json:"ignore_missing_read_permissions_tables"`
	// The region of the Dynamodb database
	Region *SourceDynamodbUpdateDynamodbRegion `default:"" json:"region"`
	// Comma separated reserved attribute names present in your tables
	ReservedAttributeNames *string `json:"reserved_attribute_names,omitempty"`
}

func (s SourceDynamodbUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDynamodbUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDynamodbUpdate) GetCredentials() *Credentials {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceDynamodbUpdate) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *SourceDynamodbUpdate) GetIgnoreMissingReadPermissionsTables() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreMissingReadPermissionsTables
}

func (o *SourceDynamodbUpdate) GetRegion() *SourceDynamodbUpdateDynamodbRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *SourceDynamodbUpdate) GetReservedAttributeNames() *string {
	if o == nil {
		return nil
	}
	return o.ReservedAttributeNames
}
