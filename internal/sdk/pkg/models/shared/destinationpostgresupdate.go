// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode string

const (
	DestinationPostgresUpdateSchemasSSLModeSSLModes6ModeVerifyFull DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode = "verify-full"
)

func (e DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode) ToPointer() *DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode {
	return &e
}

func (e *DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode: %v", v)
	}
}

// VerifyFull - Verify-full SSL mode.
type VerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                               `json:"client_key_password,omitempty"`
	mode              *DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode `const:"verify-full" json:"mode"`
}

func (v VerifyFull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VerifyFull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *VerifyFull) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *VerifyFull) GetClientCertificate() string {
	if o == nil {
		return ""
	}
	return o.ClientCertificate
}

func (o *VerifyFull) GetClientKey() string {
	if o == nil {
		return ""
	}
	return o.ClientKey
}

func (o *VerifyFull) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *VerifyFull) GetMode() *DestinationPostgresUpdateSchemasSSLModeSSLModes6Mode {
	return DestinationPostgresUpdateSchemasSSLModeSSLModes6ModeVerifyFull.ToPointer()
}

type DestinationPostgresUpdateSchemasSSLModeSSLModesMode string

const (
	DestinationPostgresUpdateSchemasSSLModeSSLModesModeVerifyCa DestinationPostgresUpdateSchemasSSLModeSSLModesMode = "verify-ca"
)

func (e DestinationPostgresUpdateSchemasSSLModeSSLModesMode) ToPointer() *DestinationPostgresUpdateSchemasSSLModeSSLModesMode {
	return &e
}

func (e *DestinationPostgresUpdateSchemasSSLModeSSLModesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-ca":
		*e = DestinationPostgresUpdateSchemasSSLModeSSLModesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSchemasSSLModeSSLModesMode: %v", v)
	}
}

// VerifyCa - Verify-ca SSL mode.
type VerifyCa struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                              `json:"client_key_password,omitempty"`
	mode              *DestinationPostgresUpdateSchemasSSLModeSSLModesMode `const:"verify-ca" json:"mode"`
}

func (v VerifyCa) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VerifyCa) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *VerifyCa) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *VerifyCa) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *VerifyCa) GetMode() *DestinationPostgresUpdateSchemasSSLModeSSLModesMode {
	return DestinationPostgresUpdateSchemasSSLModeSSLModesModeVerifyCa.ToPointer()
}

type DestinationPostgresUpdateSchemasSslModeMode string

const (
	DestinationPostgresUpdateSchemasSslModeModeRequire DestinationPostgresUpdateSchemasSslModeMode = "require"
)

func (e DestinationPostgresUpdateSchemasSslModeMode) ToPointer() *DestinationPostgresUpdateSchemasSslModeMode {
	return &e
}

func (e *DestinationPostgresUpdateSchemasSslModeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		*e = DestinationPostgresUpdateSchemasSslModeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSchemasSslModeMode: %v", v)
	}
}

// Require SSL mode.
type Require struct {
	mode *DestinationPostgresUpdateSchemasSslModeMode `const:"require" json:"mode"`
}

func (r Require) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Require) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Require) GetMode() *DestinationPostgresUpdateSchemasSslModeMode {
	return DestinationPostgresUpdateSchemasSslModeModeRequire.ToPointer()
}

type DestinationPostgresUpdateSchemasMode string

const (
	DestinationPostgresUpdateSchemasModePrefer DestinationPostgresUpdateSchemasMode = "prefer"
)

func (e DestinationPostgresUpdateSchemasMode) ToPointer() *DestinationPostgresUpdateSchemasMode {
	return &e
}

func (e *DestinationPostgresUpdateSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prefer":
		*e = DestinationPostgresUpdateSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSchemasMode: %v", v)
	}
}

// Prefer SSL mode.
type Prefer struct {
	mode *DestinationPostgresUpdateSchemasMode `const:"prefer" json:"mode"`
}

func (p Prefer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Prefer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Prefer) GetMode() *DestinationPostgresUpdateSchemasMode {
	return DestinationPostgresUpdateSchemasModePrefer.ToPointer()
}

type DestinationPostgresUpdateMode string

const (
	DestinationPostgresUpdateModeAllow DestinationPostgresUpdateMode = "allow"
)

func (e DestinationPostgresUpdateMode) ToPointer() *DestinationPostgresUpdateMode {
	return &e
}

func (e *DestinationPostgresUpdateMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		*e = DestinationPostgresUpdateMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateMode: %v", v)
	}
}

// Allow SSL mode.
type Allow struct {
	mode *DestinationPostgresUpdateMode `const:"allow" json:"mode"`
}

func (a Allow) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Allow) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Allow) GetMode() *DestinationPostgresUpdateMode {
	return DestinationPostgresUpdateModeAllow.ToPointer()
}

type Mode string

const (
	ModeDisable Mode = "disable"
)

func (e Mode) ToPointer() *Mode {
	return &e
}

func (e *Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mode: %v", v)
	}
}

// Disable SSL.
type Disable struct {
	mode *Mode `const:"disable" json:"mode"`
}

func (d Disable) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Disable) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Disable) GetMode() *Mode {
	return ModeDisable.ToPointer()
}

type SSLModesType string

const (
	SSLModesTypeDisable    SSLModesType = "disable"
	SSLModesTypeAllow      SSLModesType = "allow"
	SSLModesTypePrefer     SSLModesType = "prefer"
	SSLModesTypeRequire    SSLModesType = "require"
	SSLModesTypeVerifyCa   SSLModesType = "verify-ca"
	SSLModesTypeVerifyFull SSLModesType = "verify-full"
)

// SSLModes - SSL connection modes.
//
//	<b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
//	<b>allow</b> - Chose this mode to enable encryption only when required by the source database
//	<b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
//	<b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
//	 <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
//	 <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
//	See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
type SSLModes struct {
	Disable    *Disable
	Allow      *Allow
	Prefer     *Prefer
	Require    *Require
	VerifyCa   *VerifyCa
	VerifyFull *VerifyFull

	Type SSLModesType
}

func CreateSSLModesDisable(disable Disable) SSLModes {
	typ := SSLModesTypeDisable

	return SSLModes{
		Disable: &disable,
		Type:    typ,
	}
}

func CreateSSLModesAllow(allow Allow) SSLModes {
	typ := SSLModesTypeAllow

	return SSLModes{
		Allow: &allow,
		Type:  typ,
	}
}

func CreateSSLModesPrefer(prefer Prefer) SSLModes {
	typ := SSLModesTypePrefer

	return SSLModes{
		Prefer: &prefer,
		Type:   typ,
	}
}

func CreateSSLModesRequire(require Require) SSLModes {
	typ := SSLModesTypeRequire

	return SSLModes{
		Require: &require,
		Type:    typ,
	}
}

func CreateSSLModesVerifyCa(verifyCa VerifyCa) SSLModes {
	typ := SSLModesTypeVerifyCa

	return SSLModes{
		VerifyCa: &verifyCa,
		Type:     typ,
	}
}

func CreateSSLModesVerifyFull(verifyFull VerifyFull) SSLModes {
	typ := SSLModesTypeVerifyFull

	return SSLModes{
		VerifyFull: &verifyFull,
		Type:       typ,
	}
}

func (u *SSLModes) UnmarshalJSON(data []byte) error {

	disable := new(Disable)
	if err := utils.UnmarshalJSON(data, &disable, "", true, true); err == nil {
		u.Disable = disable
		u.Type = SSLModesTypeDisable
		return nil
	}

	allow := new(Allow)
	if err := utils.UnmarshalJSON(data, &allow, "", true, true); err == nil {
		u.Allow = allow
		u.Type = SSLModesTypeAllow
		return nil
	}

	prefer := new(Prefer)
	if err := utils.UnmarshalJSON(data, &prefer, "", true, true); err == nil {
		u.Prefer = prefer
		u.Type = SSLModesTypePrefer
		return nil
	}

	require := new(Require)
	if err := utils.UnmarshalJSON(data, &require, "", true, true); err == nil {
		u.Require = require
		u.Type = SSLModesTypeRequire
		return nil
	}

	verifyCa := new(VerifyCa)
	if err := utils.UnmarshalJSON(data, &verifyCa, "", true, true); err == nil {
		u.VerifyCa = verifyCa
		u.Type = SSLModesTypeVerifyCa
		return nil
	}

	verifyFull := new(VerifyFull)
	if err := utils.UnmarshalJSON(data, &verifyFull, "", true, true); err == nil {
		u.VerifyFull = verifyFull
		u.Type = SSLModesTypeVerifyFull
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SSLModes) MarshalJSON() ([]byte, error) {
	if u.Disable != nil {
		return utils.MarshalJSON(u.Disable, "", true)
	}

	if u.Allow != nil {
		return utils.MarshalJSON(u.Allow, "", true)
	}

	if u.Prefer != nil {
		return utils.MarshalJSON(u.Prefer, "", true)
	}

	if u.Require != nil {
		return utils.MarshalJSON(u.Require, "", true)
	}

	if u.VerifyCa != nil {
		return utils.MarshalJSON(u.VerifyCa, "", true)
	}

	if u.VerifyFull != nil {
		return utils.MarshalJSON(u.VerifyFull, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationPostgresUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationPostgresUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationPostgresUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationPostgresUpdatePasswordAuthentication) GetTunnelMethod() DestinationPostgresUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationPostgresUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationPostgresUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationPostgresUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationPostgresUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationPostgresUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationPostgresUpdateSchemasTunnelMethod string

const (
	DestinationPostgresUpdateSchemasTunnelMethodSSHKeyAuth DestinationPostgresUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationPostgresUpdateSchemasTunnelMethod) ToPointer() *DestinationPostgresUpdateSchemasTunnelMethod {
	return &e
}

func (e *DestinationPostgresUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationPostgresUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateSchemasTunnelMethod: %v", v)
	}
}

type DestinationPostgresUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationPostgresUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationPostgresUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationPostgresUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationPostgresUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationPostgresUpdateSchemasTunnelMethod {
	return DestinationPostgresUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationPostgresUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationPostgresUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationPostgresUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationPostgresUpdateTunnelMethod string

const (
	DestinationPostgresUpdateTunnelMethodNoTunnel DestinationPostgresUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationPostgresUpdateTunnelMethod) ToPointer() *DestinationPostgresUpdateTunnelMethod {
	return &e
}

func (e *DestinationPostgresUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationPostgresUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresUpdateTunnelMethod: %v", v)
	}
}

type DestinationPostgresUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationPostgresUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationPostgresUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresUpdateNoTunnel) GetTunnelMethod() DestinationPostgresUpdateTunnelMethod {
	return DestinationPostgresUpdateTunnelMethodNoTunnel
}

type DestinationPostgresUpdateSSHTunnelMethodType string

const (
	DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateNoTunnel               DestinationPostgresUpdateSSHTunnelMethodType = "destination-postgres-update_No Tunnel"
	DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHKeyAuthentication   DestinationPostgresUpdateSSHTunnelMethodType = "destination-postgres-update_SSH Key Authentication"
	DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdatePasswordAuthentication DestinationPostgresUpdateSSHTunnelMethodType = "destination-postgres-update_Password Authentication"
)

// DestinationPostgresUpdateSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresUpdateSSHTunnelMethod struct {
	DestinationPostgresUpdateNoTunnel               *DestinationPostgresUpdateNoTunnel
	DestinationPostgresUpdateSSHKeyAuthentication   *DestinationPostgresUpdateSSHKeyAuthentication
	DestinationPostgresUpdatePasswordAuthentication *DestinationPostgresUpdatePasswordAuthentication

	Type DestinationPostgresUpdateSSHTunnelMethodType
}

func CreateDestinationPostgresUpdateSSHTunnelMethodDestinationPostgresUpdateNoTunnel(destinationPostgresUpdateNoTunnel DestinationPostgresUpdateNoTunnel) DestinationPostgresUpdateSSHTunnelMethod {
	typ := DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateNoTunnel

	return DestinationPostgresUpdateSSHTunnelMethod{
		DestinationPostgresUpdateNoTunnel: &destinationPostgresUpdateNoTunnel,
		Type:                              typ,
	}
}

func CreateDestinationPostgresUpdateSSHTunnelMethodDestinationPostgresUpdateSSHKeyAuthentication(destinationPostgresUpdateSSHKeyAuthentication DestinationPostgresUpdateSSHKeyAuthentication) DestinationPostgresUpdateSSHTunnelMethod {
	typ := DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHKeyAuthentication

	return DestinationPostgresUpdateSSHTunnelMethod{
		DestinationPostgresUpdateSSHKeyAuthentication: &destinationPostgresUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationPostgresUpdateSSHTunnelMethodDestinationPostgresUpdatePasswordAuthentication(destinationPostgresUpdatePasswordAuthentication DestinationPostgresUpdatePasswordAuthentication) DestinationPostgresUpdateSSHTunnelMethod {
	typ := DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdatePasswordAuthentication

	return DestinationPostgresUpdateSSHTunnelMethod{
		DestinationPostgresUpdatePasswordAuthentication: &destinationPostgresUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationPostgresUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationPostgresUpdateNoTunnel := new(DestinationPostgresUpdateNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationPostgresUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationPostgresUpdateNoTunnel = destinationPostgresUpdateNoTunnel
		u.Type = DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateNoTunnel
		return nil
	}

	destinationPostgresUpdateSSHKeyAuthentication := new(DestinationPostgresUpdateSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationPostgresUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationPostgresUpdateSSHKeyAuthentication = destinationPostgresUpdateSSHKeyAuthentication
		u.Type = DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdateSSHKeyAuthentication
		return nil
	}

	destinationPostgresUpdatePasswordAuthentication := new(DestinationPostgresUpdatePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationPostgresUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationPostgresUpdatePasswordAuthentication = destinationPostgresUpdatePasswordAuthentication
		u.Type = DestinationPostgresUpdateSSHTunnelMethodTypeDestinationPostgresUpdatePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationPostgresUpdateNoTunnel, "", true)
	}

	if u.DestinationPostgresUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationPostgresUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationPostgresUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationPostgresUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationPostgresUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions
	DisableTypeDedupe *bool `default:"false" json:"disable_type_dedupe"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"5432" json:"port"`
	// The schema to write raw tables into
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the source database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
	//  See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
	SslMode *SSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationPostgresUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationPostgresUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationPostgresUpdate) GetDisableTypeDedupe() *bool {
	if o == nil {
		return nil
	}
	return o.DisableTypeDedupe
}

func (o *DestinationPostgresUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationPostgresUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationPostgresUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationPostgresUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationPostgresUpdate) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationPostgresUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationPostgresUpdate) GetSslMode() *SSLModes {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *DestinationPostgresUpdate) GetTunnelMethod() *DestinationPostgresUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationPostgresUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
