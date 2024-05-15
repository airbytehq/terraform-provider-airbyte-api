// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

// SourceSftpSchemasAuthMethod - Connect through ssh key
type SourceSftpSchemasAuthMethod string

const (
	SourceSftpSchemasAuthMethodSSHKeyAuth SourceSftpSchemasAuthMethod = "SSH_KEY_AUTH"
)

func (e SourceSftpSchemasAuthMethod) ToPointer() *SourceSftpSchemasAuthMethod {
	return &e
}
func (e *SourceSftpSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceSftpSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpSchemasAuthMethod: %v", v)
	}
}

type SourceSftpSSHKeyAuthentication struct {
	// Connect through ssh key
	authMethod SourceSftpSchemasAuthMethod `const:"SSH_KEY_AUTH" json:"auth_method"`
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	AuthSSHKey string `json:"auth_ssh_key"`
}

func (s SourceSftpSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSftpSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSftpSSHKeyAuthentication) GetAuthMethod() SourceSftpSchemasAuthMethod {
	return SourceSftpSchemasAuthMethodSSHKeyAuth
}

func (o *SourceSftpSSHKeyAuthentication) GetAuthSSHKey() string {
	if o == nil {
		return ""
	}
	return o.AuthSSHKey
}

// SourceSftpAuthMethod - Connect through password authentication
type SourceSftpAuthMethod string

const (
	SourceSftpAuthMethodSSHPasswordAuth SourceSftpAuthMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceSftpAuthMethod) ToPointer() *SourceSftpAuthMethod {
	return &e
}
func (e *SourceSftpAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceSftpAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSftpAuthMethod: %v", v)
	}
}

type SourceSftpPasswordAuthentication struct {
	// Connect through password authentication
	authMethod SourceSftpAuthMethod `const:"SSH_PASSWORD_AUTH" json:"auth_method"`
	// OS-level password for logging into the jump server host
	AuthUserPassword string `json:"auth_user_password"`
}

func (s SourceSftpPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSftpPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSftpPasswordAuthentication) GetAuthMethod() SourceSftpAuthMethod {
	return SourceSftpAuthMethodSSHPasswordAuth
}

func (o *SourceSftpPasswordAuthentication) GetAuthUserPassword() string {
	if o == nil {
		return ""
	}
	return o.AuthUserPassword
}

type SourceSftpAuthenticationType string

const (
	SourceSftpAuthenticationTypeSourceSftpPasswordAuthentication SourceSftpAuthenticationType = "source-sftp_Password Authentication"
	SourceSftpAuthenticationTypeSourceSftpSSHKeyAuthentication   SourceSftpAuthenticationType = "source-sftp_SSH Key Authentication"
)

// SourceSftpAuthentication - The server authentication method
type SourceSftpAuthentication struct {
	SourceSftpPasswordAuthentication *SourceSftpPasswordAuthentication
	SourceSftpSSHKeyAuthentication   *SourceSftpSSHKeyAuthentication

	Type SourceSftpAuthenticationType
}

func CreateSourceSftpAuthenticationSourceSftpPasswordAuthentication(sourceSftpPasswordAuthentication SourceSftpPasswordAuthentication) SourceSftpAuthentication {
	typ := SourceSftpAuthenticationTypeSourceSftpPasswordAuthentication

	return SourceSftpAuthentication{
		SourceSftpPasswordAuthentication: &sourceSftpPasswordAuthentication,
		Type:                             typ,
	}
}

func CreateSourceSftpAuthenticationSourceSftpSSHKeyAuthentication(sourceSftpSSHKeyAuthentication SourceSftpSSHKeyAuthentication) SourceSftpAuthentication {
	typ := SourceSftpAuthenticationTypeSourceSftpSSHKeyAuthentication

	return SourceSftpAuthentication{
		SourceSftpSSHKeyAuthentication: &sourceSftpSSHKeyAuthentication,
		Type:                           typ,
	}
}

func (u *SourceSftpAuthentication) UnmarshalJSON(data []byte) error {

	var sourceSftpPasswordAuthentication SourceSftpPasswordAuthentication = SourceSftpPasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &sourceSftpPasswordAuthentication, "", true, true); err == nil {
		u.SourceSftpPasswordAuthentication = &sourceSftpPasswordAuthentication
		u.Type = SourceSftpAuthenticationTypeSourceSftpPasswordAuthentication
		return nil
	}

	var sourceSftpSSHKeyAuthentication SourceSftpSSHKeyAuthentication = SourceSftpSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &sourceSftpSSHKeyAuthentication, "", true, true); err == nil {
		u.SourceSftpSSHKeyAuthentication = &sourceSftpSSHKeyAuthentication
		u.Type = SourceSftpAuthenticationTypeSourceSftpSSHKeyAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSftpAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceSftpPasswordAuthentication != nil {
		return utils.MarshalJSON(u.SourceSftpPasswordAuthentication, "", true)
	}

	if u.SourceSftpSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceSftpSSHKeyAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Sftp string

const (
	SftpSftp Sftp = "sftp"
)

func (e Sftp) ToPointer() *Sftp {
	return &e
}
func (e *Sftp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sftp":
		*e = Sftp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sftp: %v", v)
	}
}

type SourceSftp struct {
	// The server authentication method
	Credentials *SourceSftpAuthentication `json:"credentials,omitempty"`
	// The regular expression to specify files for sync in a chosen Folder Path
	FilePattern *string `default:"" json:"file_pattern"`
	// Coma separated file types. Currently only 'csv' and 'json' types are supported.
	FileTypes *string `default:"csv,json" json:"file_types"`
	// The directory to search files for sync
	FolderPath *string `default:"" json:"folder_path"`
	// The server host address
	Host string `json:"host"`
	// The server port
	Port       *int64 `default:"22" json:"port"`
	sourceType Sftp   `const:"sftp" json:"sourceType"`
	// The server user
	User string `json:"user"`
}

func (s SourceSftp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSftp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSftp) GetCredentials() *SourceSftpAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceSftp) GetFilePattern() *string {
	if o == nil {
		return nil
	}
	return o.FilePattern
}

func (o *SourceSftp) GetFileTypes() *string {
	if o == nil {
		return nil
	}
	return o.FileTypes
}

func (o *SourceSftp) GetFolderPath() *string {
	if o == nil {
		return nil
	}
	return o.FolderPath
}

func (o *SourceSftp) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceSftp) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceSftp) GetSourceType() Sftp {
	return SftpSftp
}

func (o *SourceSftp) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}
