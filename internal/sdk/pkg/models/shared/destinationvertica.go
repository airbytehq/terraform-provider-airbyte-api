// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Vertica string

const (
	VerticaVertica Vertica = "vertica"
)

func (e Vertica) ToPointer() *Vertica {
	return &e
}

func (e *Vertica) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vertica":
		*e = Vertica(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Vertica: %v", v)
	}
}

// DestinationVerticaSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationVerticaSchemasTunnelMethodTunnelMethod string

const (
	DestinationVerticaSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationVerticaSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationVerticaSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationVerticaSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationVerticaSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationVerticaSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// DestinationVerticaPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationVerticaSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationVerticaPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationVerticaPasswordAuthentication) GetTunnelMethod() DestinationVerticaSchemasTunnelMethodTunnelMethod {
	return DestinationVerticaSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationVerticaPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationVerticaPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationVerticaPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationVerticaSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationVerticaSchemasTunnelMethod string

const (
	DestinationVerticaSchemasTunnelMethodSSHKeyAuth DestinationVerticaSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationVerticaSchemasTunnelMethod) ToPointer() *DestinationVerticaSchemasTunnelMethod {
	return &e
}

func (e *DestinationVerticaSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationVerticaSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaSchemasTunnelMethod: %v", v)
	}
}

// DestinationVerticaSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationVerticaSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationVerticaSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationVerticaSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationVerticaSSHKeyAuthentication) GetTunnelMethod() DestinationVerticaSchemasTunnelMethod {
	return DestinationVerticaSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationVerticaSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationVerticaSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationVerticaTunnelMethod - No ssh tunnel needed to connect to database
type DestinationVerticaTunnelMethod string

const (
	DestinationVerticaTunnelMethodNoTunnel DestinationVerticaTunnelMethod = "NO_TUNNEL"
)

func (e DestinationVerticaTunnelMethod) ToPointer() *DestinationVerticaTunnelMethod {
	return &e
}

func (e *DestinationVerticaTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationVerticaTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaTunnelMethod: %v", v)
	}
}

// DestinationVerticaNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationVerticaTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationVerticaNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaNoTunnel) GetTunnelMethod() DestinationVerticaTunnelMethod {
	return DestinationVerticaTunnelMethodNoTunnel
}

type DestinationVerticaSSHTunnelMethodType string

const (
	DestinationVerticaSSHTunnelMethodTypeDestinationVerticaNoTunnel               DestinationVerticaSSHTunnelMethodType = "destination-vertica_No Tunnel"
	DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHKeyAuthentication   DestinationVerticaSSHTunnelMethodType = "destination-vertica_SSH Key Authentication"
	DestinationVerticaSSHTunnelMethodTypeDestinationVerticaPasswordAuthentication DestinationVerticaSSHTunnelMethodType = "destination-vertica_Password Authentication"
)

type DestinationVerticaSSHTunnelMethod struct {
	DestinationVerticaNoTunnel               *DestinationVerticaNoTunnel
	DestinationVerticaSSHKeyAuthentication   *DestinationVerticaSSHKeyAuthentication
	DestinationVerticaPasswordAuthentication *DestinationVerticaPasswordAuthentication

	Type DestinationVerticaSSHTunnelMethodType
}

func CreateDestinationVerticaSSHTunnelMethodDestinationVerticaNoTunnel(destinationVerticaNoTunnel DestinationVerticaNoTunnel) DestinationVerticaSSHTunnelMethod {
	typ := DestinationVerticaSSHTunnelMethodTypeDestinationVerticaNoTunnel

	return DestinationVerticaSSHTunnelMethod{
		DestinationVerticaNoTunnel: &destinationVerticaNoTunnel,
		Type:                       typ,
	}
}

func CreateDestinationVerticaSSHTunnelMethodDestinationVerticaSSHKeyAuthentication(destinationVerticaSSHKeyAuthentication DestinationVerticaSSHKeyAuthentication) DestinationVerticaSSHTunnelMethod {
	typ := DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHKeyAuthentication

	return DestinationVerticaSSHTunnelMethod{
		DestinationVerticaSSHKeyAuthentication: &destinationVerticaSSHKeyAuthentication,
		Type:                                   typ,
	}
}

func CreateDestinationVerticaSSHTunnelMethodDestinationVerticaPasswordAuthentication(destinationVerticaPasswordAuthentication DestinationVerticaPasswordAuthentication) DestinationVerticaSSHTunnelMethod {
	typ := DestinationVerticaSSHTunnelMethodTypeDestinationVerticaPasswordAuthentication

	return DestinationVerticaSSHTunnelMethod{
		DestinationVerticaPasswordAuthentication: &destinationVerticaPasswordAuthentication,
		Type:                                     typ,
	}
}

func (u *DestinationVerticaSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationVerticaNoTunnel := new(DestinationVerticaNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationVerticaNoTunnel, "", true, true); err == nil {
		u.DestinationVerticaNoTunnel = destinationVerticaNoTunnel
		u.Type = DestinationVerticaSSHTunnelMethodTypeDestinationVerticaNoTunnel
		return nil
	}

	destinationVerticaSSHKeyAuthentication := new(DestinationVerticaSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationVerticaSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationVerticaSSHKeyAuthentication = destinationVerticaSSHKeyAuthentication
		u.Type = DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHKeyAuthentication
		return nil
	}

	destinationVerticaPasswordAuthentication := new(DestinationVerticaPasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationVerticaPasswordAuthentication, "", true, true); err == nil {
		u.DestinationVerticaPasswordAuthentication = destinationVerticaPasswordAuthentication
		u.Type = DestinationVerticaSSHTunnelMethodTypeDestinationVerticaPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationVerticaSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationVerticaNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationVerticaNoTunnel, "", true)
	}

	if u.DestinationVerticaSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationVerticaSSHKeyAuthentication, "", true)
	}

	if u.DestinationVerticaPasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationVerticaPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationVertica struct {
	// Name of the database.
	Database        string  `json:"database"`
	destinationType Vertica `const:"vertica" json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"5433" json:"port"`
	// Schema for vertica destination
	Schema string `json:"schema"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationVerticaSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationVertica) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVertica) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVertica) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationVertica) GetDestinationType() Vertica {
	return VerticaVertica
}

func (o *DestinationVertica) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationVertica) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationVertica) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationVertica) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationVertica) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *DestinationVertica) GetTunnelMethod() *DestinationVerticaSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationVertica) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
