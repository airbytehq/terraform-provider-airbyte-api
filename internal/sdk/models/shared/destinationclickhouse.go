// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Clickhouse string

const (
	ClickhouseClickhouse Clickhouse = "clickhouse"
)

func (e Clickhouse) ToPointer() *Clickhouse {
	return &e
}
func (e *Clickhouse) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "clickhouse":
		*e = Clickhouse(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Clickhouse: %v", v)
	}
}

// DestinationClickhouseSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationClickhouseSchemasTunnelMethodTunnelMethod string

const (
	DestinationClickhouseSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationClickhouseSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationClickhouseSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationClickhouseSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationClickhouseSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationClickhouseSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationClickhousePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationClickhouseSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationClickhousePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationClickhousePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationClickhousePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationClickhousePasswordAuthentication) GetTunnelMethod() DestinationClickhouseSchemasTunnelMethodTunnelMethod {
	return DestinationClickhouseSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationClickhousePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationClickhousePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationClickhousePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationClickhouseSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationClickhouseSchemasTunnelMethod string

const (
	DestinationClickhouseSchemasTunnelMethodSSHKeyAuth DestinationClickhouseSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationClickhouseSchemasTunnelMethod) ToPointer() *DestinationClickhouseSchemasTunnelMethod {
	return &e
}
func (e *DestinationClickhouseSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationClickhouseSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseSchemasTunnelMethod: %v", v)
	}
}

type DestinationClickhouseSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationClickhouseSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationClickhouseSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationClickhouseSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationClickhouseSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationClickhouseSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationClickhouseSSHKeyAuthentication) GetTunnelMethod() DestinationClickhouseSchemasTunnelMethod {
	return DestinationClickhouseSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationClickhouseSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationClickhouseSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationClickhouseTunnelMethod - No ssh tunnel needed to connect to database
type DestinationClickhouseTunnelMethod string

const (
	DestinationClickhouseTunnelMethodNoTunnel DestinationClickhouseTunnelMethod = "NO_TUNNEL"
)

func (e DestinationClickhouseTunnelMethod) ToPointer() *DestinationClickhouseTunnelMethod {
	return &e
}
func (e *DestinationClickhouseTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationClickhouseTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationClickhouseTunnelMethod: %v", v)
	}
}

type DestinationClickhouseNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationClickhouseTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationClickhouseNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationClickhouseNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationClickhouseNoTunnel) GetTunnelMethod() DestinationClickhouseTunnelMethod {
	return DestinationClickhouseTunnelMethodNoTunnel
}

type DestinationClickhouseSSHTunnelMethodType string

const (
	DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseNoTunnel               DestinationClickhouseSSHTunnelMethodType = "destination-clickhouse_No Tunnel"
	DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHKeyAuthentication   DestinationClickhouseSSHTunnelMethodType = "destination-clickhouse_SSH Key Authentication"
	DestinationClickhouseSSHTunnelMethodTypeDestinationClickhousePasswordAuthentication DestinationClickhouseSSHTunnelMethodType = "destination-clickhouse_Password Authentication"
)

// DestinationClickhouseSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationClickhouseSSHTunnelMethod struct {
	DestinationClickhouseNoTunnel               *DestinationClickhouseNoTunnel
	DestinationClickhouseSSHKeyAuthentication   *DestinationClickhouseSSHKeyAuthentication
	DestinationClickhousePasswordAuthentication *DestinationClickhousePasswordAuthentication

	Type DestinationClickhouseSSHTunnelMethodType
}

func CreateDestinationClickhouseSSHTunnelMethodDestinationClickhouseNoTunnel(destinationClickhouseNoTunnel DestinationClickhouseNoTunnel) DestinationClickhouseSSHTunnelMethod {
	typ := DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseNoTunnel

	return DestinationClickhouseSSHTunnelMethod{
		DestinationClickhouseNoTunnel: &destinationClickhouseNoTunnel,
		Type:                          typ,
	}
}

func CreateDestinationClickhouseSSHTunnelMethodDestinationClickhouseSSHKeyAuthentication(destinationClickhouseSSHKeyAuthentication DestinationClickhouseSSHKeyAuthentication) DestinationClickhouseSSHTunnelMethod {
	typ := DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHKeyAuthentication

	return DestinationClickhouseSSHTunnelMethod{
		DestinationClickhouseSSHKeyAuthentication: &destinationClickhouseSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationClickhouseSSHTunnelMethodDestinationClickhousePasswordAuthentication(destinationClickhousePasswordAuthentication DestinationClickhousePasswordAuthentication) DestinationClickhouseSSHTunnelMethod {
	typ := DestinationClickhouseSSHTunnelMethodTypeDestinationClickhousePasswordAuthentication

	return DestinationClickhouseSSHTunnelMethod{
		DestinationClickhousePasswordAuthentication: &destinationClickhousePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationClickhouseSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationClickhouseNoTunnel DestinationClickhouseNoTunnel = DestinationClickhouseNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationClickhouseNoTunnel, "", true, true); err == nil {
		u.DestinationClickhouseNoTunnel = &destinationClickhouseNoTunnel
		u.Type = DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseNoTunnel
		return nil
	}

	var destinationClickhouseSSHKeyAuthentication DestinationClickhouseSSHKeyAuthentication = DestinationClickhouseSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationClickhouseSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationClickhouseSSHKeyAuthentication = &destinationClickhouseSSHKeyAuthentication
		u.Type = DestinationClickhouseSSHTunnelMethodTypeDestinationClickhouseSSHKeyAuthentication
		return nil
	}

	var destinationClickhousePasswordAuthentication DestinationClickhousePasswordAuthentication = DestinationClickhousePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationClickhousePasswordAuthentication, "", true, true); err == nil {
		u.DestinationClickhousePasswordAuthentication = &destinationClickhousePasswordAuthentication
		u.Type = DestinationClickhouseSSHTunnelMethodTypeDestinationClickhousePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationClickhouseSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationClickhouseNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationClickhouseNoTunnel, "", true)
	}

	if u.DestinationClickhouseSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationClickhouseSSHKeyAuthentication, "", true)
	}

	if u.DestinationClickhousePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationClickhousePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationClickhouse struct {
	// Name of the database.
	Database        string     `json:"database"`
	destinationType Clickhouse `const:"clickhouse" json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// HTTP port of the database.
	Port *int64 `default:"8123" json:"port"`
	// The schema to write raw tables into (default: airbyte_internal)
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationClickhouseSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationClickhouse) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationClickhouse) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationClickhouse) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationClickhouse) GetDestinationType() Clickhouse {
	return ClickhouseClickhouse
}

func (o *DestinationClickhouse) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationClickhouse) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationClickhouse) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationClickhouse) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationClickhouse) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationClickhouse) GetTunnelMethod() *DestinationClickhouseSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationClickhouse) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
