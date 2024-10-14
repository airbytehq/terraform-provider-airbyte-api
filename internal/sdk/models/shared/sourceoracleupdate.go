// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceOracleUpdateConnectionType string

const (
	SourceOracleUpdateConnectionTypeSid SourceOracleUpdateConnectionType = "sid"
)

func (e SourceOracleUpdateConnectionType) ToPointer() *SourceOracleUpdateConnectionType {
	return &e
}
func (e *SourceOracleUpdateConnectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sid":
		*e = SourceOracleUpdateConnectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateConnectionType: %v", v)
	}
}

// SystemIDSID - Use SID (Oracle System Identifier)
type SystemIDSID struct {
	connectionType *SourceOracleUpdateConnectionType `const:"sid" json:"connection_type,omitempty"`
	Sid            string                            `json:"sid"`
}

func (s SystemIDSID) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SystemIDSID) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SystemIDSID) GetConnectionType() *SourceOracleUpdateConnectionType {
	return SourceOracleUpdateConnectionTypeSid.ToPointer()
}

func (o *SystemIDSID) GetSid() string {
	if o == nil {
		return ""
	}
	return o.Sid
}

type ConnectionType string

const (
	ConnectionTypeServiceName ConnectionType = "service_name"
)

func (e ConnectionType) ToPointer() *ConnectionType {
	return &e
}
func (e *ConnectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_name":
		*e = ConnectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionType: %v", v)
	}
}

// ServiceName - Use service name
type ServiceName struct {
	connectionType *ConnectionType `const:"service_name" json:"connection_type,omitempty"`
	ServiceName    string          `json:"service_name"`
}

func (s ServiceName) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *ServiceName) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ServiceName) GetConnectionType() *ConnectionType {
	return ConnectionTypeServiceName.ToPointer()
}

func (o *ServiceName) GetServiceName() string {
	if o == nil {
		return ""
	}
	return o.ServiceName
}

type ConnectByType string

const (
	ConnectByTypeServiceName ConnectByType = "Service name"
	ConnectByTypeSystemIDSID ConnectByType = "System ID (SID)"
)

// ConnectBy - Connect data that will be used for DB connection
type ConnectBy struct {
	ServiceName *ServiceName
	SystemIDSID *SystemIDSID

	Type ConnectByType
}

func CreateConnectByServiceName(serviceName ServiceName) ConnectBy {
	typ := ConnectByTypeServiceName

	return ConnectBy{
		ServiceName: &serviceName,
		Type:        typ,
	}
}

func CreateConnectBySystemIDSID(systemIDSID SystemIDSID) ConnectBy {
	typ := ConnectByTypeSystemIDSID

	return ConnectBy{
		SystemIDSID: &systemIDSID,
		Type:        typ,
	}
}

func (u *ConnectBy) UnmarshalJSON(data []byte) error {

	var serviceName ServiceName = ServiceName{}
	if err := utils.UnmarshalJSON(data, &serviceName, "", true, true); err == nil {
		u.ServiceName = &serviceName
		u.Type = ConnectByTypeServiceName
		return nil
	}

	var systemIDSID SystemIDSID = SystemIDSID{}
	if err := utils.UnmarshalJSON(data, &systemIDSID, "", true, true); err == nil {
		u.SystemIDSID = &systemIDSID
		u.Type = ConnectByTypeSystemIDSID
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ConnectBy", string(data))
}

func (u ConnectBy) MarshalJSON() ([]byte, error) {
	if u.ServiceName != nil {
		return utils.MarshalJSON(u.ServiceName, "", true)
	}

	if u.SystemIDSID != nil {
		return utils.MarshalJSON(u.SystemIDSID, "", true)
	}

	return nil, errors.New("could not marshal union type ConnectBy: all fields are null")
}

type SourceOracleUpdateSchemasEncryptionEncryptionMethod string

const (
	SourceOracleUpdateSchemasEncryptionEncryptionMethodEncryptedVerifyCertificate SourceOracleUpdateSchemasEncryptionEncryptionMethod = "encrypted_verify_certificate"
)

func (e SourceOracleUpdateSchemasEncryptionEncryptionMethod) ToPointer() *SourceOracleUpdateSchemasEncryptionEncryptionMethod {
	return &e
}
func (e *SourceOracleUpdateSchemasEncryptionEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = SourceOracleUpdateSchemasEncryptionEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateSchemasEncryptionEncryptionMethod: %v", v)
	}
}

// SourceOracleUpdateTLSEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type SourceOracleUpdateTLSEncryptedVerifyCertificate struct {
	encryptionMethod SourceOracleUpdateSchemasEncryptionEncryptionMethod `const:"encrypted_verify_certificate" json:"encryption_method"`
	// Privacy Enhanced Mail (PEM) files are concatenated certificate containers frequently used in certificate installations.
	SslCertificate string `json:"ssl_certificate"`
}

func (s SourceOracleUpdateTLSEncryptedVerifyCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdateTLSEncryptedVerifyCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdateTLSEncryptedVerifyCertificate) GetEncryptionMethod() SourceOracleUpdateSchemasEncryptionEncryptionMethod {
	return SourceOracleUpdateSchemasEncryptionEncryptionMethodEncryptedVerifyCertificate
}

func (o *SourceOracleUpdateTLSEncryptedVerifyCertificate) GetSslCertificate() string {
	if o == nil {
		return ""
	}
	return o.SslCertificate
}

// SourceOracleUpdateEncryptionAlgorithm - This parameter defines what encryption algorithm is used.
type SourceOracleUpdateEncryptionAlgorithm string

const (
	SourceOracleUpdateEncryptionAlgorithmAes256      SourceOracleUpdateEncryptionAlgorithm = "AES256"
	SourceOracleUpdateEncryptionAlgorithmRc456       SourceOracleUpdateEncryptionAlgorithm = "RC4_56"
	SourceOracleUpdateEncryptionAlgorithmThreeDes168 SourceOracleUpdateEncryptionAlgorithm = "3DES168"
)

func (e SourceOracleUpdateEncryptionAlgorithm) ToPointer() *SourceOracleUpdateEncryptionAlgorithm {
	return &e
}
func (e *SourceOracleUpdateEncryptionAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AES256":
		fallthrough
	case "RC4_56":
		fallthrough
	case "3DES168":
		*e = SourceOracleUpdateEncryptionAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateEncryptionAlgorithm: %v", v)
	}
}

type SourceOracleUpdateSchemasEncryptionMethod string

const (
	SourceOracleUpdateSchemasEncryptionMethodClientNne SourceOracleUpdateSchemasEncryptionMethod = "client_nne"
)

func (e SourceOracleUpdateSchemasEncryptionMethod) ToPointer() *SourceOracleUpdateSchemasEncryptionMethod {
	return &e
}
func (e *SourceOracleUpdateSchemasEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_nne":
		*e = SourceOracleUpdateSchemasEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateSchemasEncryptionMethod: %v", v)
	}
}

// SourceOracleUpdateNativeNetworkEncryptionNNE - The native network encryption gives you the ability to encrypt database connections, without the configuration overhead of TCP/IP and SSL/TLS and without the need to open and listen on different ports.
type SourceOracleUpdateNativeNetworkEncryptionNNE struct {
	// This parameter defines what encryption algorithm is used.
	EncryptionAlgorithm *SourceOracleUpdateEncryptionAlgorithm    `default:"AES256" json:"encryption_algorithm"`
	encryptionMethod    SourceOracleUpdateSchemasEncryptionMethod `const:"client_nne" json:"encryption_method"`
}

func (s SourceOracleUpdateNativeNetworkEncryptionNNE) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdateNativeNetworkEncryptionNNE) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdateNativeNetworkEncryptionNNE) GetEncryptionAlgorithm() *SourceOracleUpdateEncryptionAlgorithm {
	if o == nil {
		return nil
	}
	return o.EncryptionAlgorithm
}

func (o *SourceOracleUpdateNativeNetworkEncryptionNNE) GetEncryptionMethod() SourceOracleUpdateSchemasEncryptionMethod {
	return SourceOracleUpdateSchemasEncryptionMethodClientNne
}

type SourceOracleUpdateEncryptionMethod string

const (
	SourceOracleUpdateEncryptionMethodUnencrypted SourceOracleUpdateEncryptionMethod = "unencrypted"
)

func (e SourceOracleUpdateEncryptionMethod) ToPointer() *SourceOracleUpdateEncryptionMethod {
	return &e
}
func (e *SourceOracleUpdateEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unencrypted":
		*e = SourceOracleUpdateEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateEncryptionMethod: %v", v)
	}
}

// SourceOracleUpdateUnencrypted - Data transfer will not be encrypted.
type SourceOracleUpdateUnencrypted struct {
	encryptionMethod SourceOracleUpdateEncryptionMethod `const:"unencrypted" json:"encryption_method"`
}

func (s SourceOracleUpdateUnencrypted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdateUnencrypted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdateUnencrypted) GetEncryptionMethod() SourceOracleUpdateEncryptionMethod {
	return SourceOracleUpdateEncryptionMethodUnencrypted
}

type SourceOracleUpdateEncryptionType string

const (
	SourceOracleUpdateEncryptionTypeSourceOracleUpdateUnencrypted                   SourceOracleUpdateEncryptionType = "source-oracle-update_Unencrypted"
	SourceOracleUpdateEncryptionTypeSourceOracleUpdateNativeNetworkEncryptionNNE    SourceOracleUpdateEncryptionType = "source-oracle-update_Native Network Encryption (NNE)"
	SourceOracleUpdateEncryptionTypeSourceOracleUpdateTLSEncryptedVerifyCertificate SourceOracleUpdateEncryptionType = "source-oracle-update_TLS Encrypted (verify certificate)"
)

// SourceOracleUpdateEncryption - The encryption method with is used when communicating with the database.
type SourceOracleUpdateEncryption struct {
	SourceOracleUpdateUnencrypted                   *SourceOracleUpdateUnencrypted
	SourceOracleUpdateNativeNetworkEncryptionNNE    *SourceOracleUpdateNativeNetworkEncryptionNNE
	SourceOracleUpdateTLSEncryptedVerifyCertificate *SourceOracleUpdateTLSEncryptedVerifyCertificate

	Type SourceOracleUpdateEncryptionType
}

func CreateSourceOracleUpdateEncryptionSourceOracleUpdateUnencrypted(sourceOracleUpdateUnencrypted SourceOracleUpdateUnencrypted) SourceOracleUpdateEncryption {
	typ := SourceOracleUpdateEncryptionTypeSourceOracleUpdateUnencrypted

	return SourceOracleUpdateEncryption{
		SourceOracleUpdateUnencrypted: &sourceOracleUpdateUnencrypted,
		Type:                          typ,
	}
}

func CreateSourceOracleUpdateEncryptionSourceOracleUpdateNativeNetworkEncryptionNNE(sourceOracleUpdateNativeNetworkEncryptionNNE SourceOracleUpdateNativeNetworkEncryptionNNE) SourceOracleUpdateEncryption {
	typ := SourceOracleUpdateEncryptionTypeSourceOracleUpdateNativeNetworkEncryptionNNE

	return SourceOracleUpdateEncryption{
		SourceOracleUpdateNativeNetworkEncryptionNNE: &sourceOracleUpdateNativeNetworkEncryptionNNE,
		Type: typ,
	}
}

func CreateSourceOracleUpdateEncryptionSourceOracleUpdateTLSEncryptedVerifyCertificate(sourceOracleUpdateTLSEncryptedVerifyCertificate SourceOracleUpdateTLSEncryptedVerifyCertificate) SourceOracleUpdateEncryption {
	typ := SourceOracleUpdateEncryptionTypeSourceOracleUpdateTLSEncryptedVerifyCertificate

	return SourceOracleUpdateEncryption{
		SourceOracleUpdateTLSEncryptedVerifyCertificate: &sourceOracleUpdateTLSEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *SourceOracleUpdateEncryption) UnmarshalJSON(data []byte) error {

	var sourceOracleUpdateUnencrypted SourceOracleUpdateUnencrypted = SourceOracleUpdateUnencrypted{}
	if err := utils.UnmarshalJSON(data, &sourceOracleUpdateUnencrypted, "", true, true); err == nil {
		u.SourceOracleUpdateUnencrypted = &sourceOracleUpdateUnencrypted
		u.Type = SourceOracleUpdateEncryptionTypeSourceOracleUpdateUnencrypted
		return nil
	}

	var sourceOracleUpdateNativeNetworkEncryptionNNE SourceOracleUpdateNativeNetworkEncryptionNNE = SourceOracleUpdateNativeNetworkEncryptionNNE{}
	if err := utils.UnmarshalJSON(data, &sourceOracleUpdateNativeNetworkEncryptionNNE, "", true, true); err == nil {
		u.SourceOracleUpdateNativeNetworkEncryptionNNE = &sourceOracleUpdateNativeNetworkEncryptionNNE
		u.Type = SourceOracleUpdateEncryptionTypeSourceOracleUpdateNativeNetworkEncryptionNNE
		return nil
	}

	var sourceOracleUpdateTLSEncryptedVerifyCertificate SourceOracleUpdateTLSEncryptedVerifyCertificate = SourceOracleUpdateTLSEncryptedVerifyCertificate{}
	if err := utils.UnmarshalJSON(data, &sourceOracleUpdateTLSEncryptedVerifyCertificate, "", true, true); err == nil {
		u.SourceOracleUpdateTLSEncryptedVerifyCertificate = &sourceOracleUpdateTLSEncryptedVerifyCertificate
		u.Type = SourceOracleUpdateEncryptionTypeSourceOracleUpdateTLSEncryptedVerifyCertificate
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceOracleUpdateEncryption", string(data))
}

func (u SourceOracleUpdateEncryption) MarshalJSON() ([]byte, error) {
	if u.SourceOracleUpdateUnencrypted != nil {
		return utils.MarshalJSON(u.SourceOracleUpdateUnencrypted, "", true)
	}

	if u.SourceOracleUpdateNativeNetworkEncryptionNNE != nil {
		return utils.MarshalJSON(u.SourceOracleUpdateNativeNetworkEncryptionNNE, "", true)
	}

	if u.SourceOracleUpdateTLSEncryptedVerifyCertificate != nil {
		return utils.MarshalJSON(u.SourceOracleUpdateTLSEncryptedVerifyCertificate, "", true)
	}

	return nil, errors.New("could not marshal union type SourceOracleUpdateEncryption: all fields are null")
}

// SourceOracleUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceOracleUpdateSchemasTunnelMethodTunnelMethod string

const (
	SourceOracleUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth SourceOracleUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceOracleUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *SourceOracleUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *SourceOracleUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceOracleUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type SourceOracleUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod SourceOracleUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (s SourceOracleUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceOracleUpdatePasswordAuthentication) GetTunnelMethod() SourceOracleUpdateSchemasTunnelMethodTunnelMethod {
	return SourceOracleUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *SourceOracleUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceOracleUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *SourceOracleUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// SourceOracleUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceOracleUpdateSchemasTunnelMethod string

const (
	SourceOracleUpdateSchemasTunnelMethodSSHKeyAuth SourceOracleUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceOracleUpdateSchemasTunnelMethod) ToPointer() *SourceOracleUpdateSchemasTunnelMethod {
	return &e
}
func (e *SourceOracleUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceOracleUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateSchemasTunnelMethod: %v", v)
	}
}

type SourceOracleUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod SourceOracleUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (s SourceOracleUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *SourceOracleUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceOracleUpdateSSHKeyAuthentication) GetTunnelMethod() SourceOracleUpdateSchemasTunnelMethod {
	return SourceOracleUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *SourceOracleUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceOracleUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// SourceOracleUpdateTunnelMethod - No ssh tunnel needed to connect to database
type SourceOracleUpdateTunnelMethod string

const (
	SourceOracleUpdateTunnelMethodNoTunnel SourceOracleUpdateTunnelMethod = "NO_TUNNEL"
)

func (e SourceOracleUpdateTunnelMethod) ToPointer() *SourceOracleUpdateTunnelMethod {
	return &e
}
func (e *SourceOracleUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceOracleUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleUpdateTunnelMethod: %v", v)
	}
}

type SourceOracleUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod SourceOracleUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (s SourceOracleUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdateNoTunnel) GetTunnelMethod() SourceOracleUpdateTunnelMethod {
	return SourceOracleUpdateTunnelMethodNoTunnel
}

type SourceOracleUpdateSSHTunnelMethodType string

const (
	SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdateNoTunnel               SourceOracleUpdateSSHTunnelMethodType = "source-oracle-update_No Tunnel"
	SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdateSSHKeyAuthentication   SourceOracleUpdateSSHTunnelMethodType = "source-oracle-update_SSH Key Authentication"
	SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdatePasswordAuthentication SourceOracleUpdateSSHTunnelMethodType = "source-oracle-update_Password Authentication"
)

// SourceOracleUpdateSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceOracleUpdateSSHTunnelMethod struct {
	SourceOracleUpdateNoTunnel               *SourceOracleUpdateNoTunnel
	SourceOracleUpdateSSHKeyAuthentication   *SourceOracleUpdateSSHKeyAuthentication
	SourceOracleUpdatePasswordAuthentication *SourceOracleUpdatePasswordAuthentication

	Type SourceOracleUpdateSSHTunnelMethodType
}

func CreateSourceOracleUpdateSSHTunnelMethodSourceOracleUpdateNoTunnel(sourceOracleUpdateNoTunnel SourceOracleUpdateNoTunnel) SourceOracleUpdateSSHTunnelMethod {
	typ := SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdateNoTunnel

	return SourceOracleUpdateSSHTunnelMethod{
		SourceOracleUpdateNoTunnel: &sourceOracleUpdateNoTunnel,
		Type:                       typ,
	}
}

func CreateSourceOracleUpdateSSHTunnelMethodSourceOracleUpdateSSHKeyAuthentication(sourceOracleUpdateSSHKeyAuthentication SourceOracleUpdateSSHKeyAuthentication) SourceOracleUpdateSSHTunnelMethod {
	typ := SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdateSSHKeyAuthentication

	return SourceOracleUpdateSSHTunnelMethod{
		SourceOracleUpdateSSHKeyAuthentication: &sourceOracleUpdateSSHKeyAuthentication,
		Type:                                   typ,
	}
}

func CreateSourceOracleUpdateSSHTunnelMethodSourceOracleUpdatePasswordAuthentication(sourceOracleUpdatePasswordAuthentication SourceOracleUpdatePasswordAuthentication) SourceOracleUpdateSSHTunnelMethod {
	typ := SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdatePasswordAuthentication

	return SourceOracleUpdateSSHTunnelMethod{
		SourceOracleUpdatePasswordAuthentication: &sourceOracleUpdatePasswordAuthentication,
		Type:                                     typ,
	}
}

func (u *SourceOracleUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var sourceOracleUpdateNoTunnel SourceOracleUpdateNoTunnel = SourceOracleUpdateNoTunnel{}
	if err := utils.UnmarshalJSON(data, &sourceOracleUpdateNoTunnel, "", true, true); err == nil {
		u.SourceOracleUpdateNoTunnel = &sourceOracleUpdateNoTunnel
		u.Type = SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdateNoTunnel
		return nil
	}

	var sourceOracleUpdateSSHKeyAuthentication SourceOracleUpdateSSHKeyAuthentication = SourceOracleUpdateSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &sourceOracleUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.SourceOracleUpdateSSHKeyAuthentication = &sourceOracleUpdateSSHKeyAuthentication
		u.Type = SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdateSSHKeyAuthentication
		return nil
	}

	var sourceOracleUpdatePasswordAuthentication SourceOracleUpdatePasswordAuthentication = SourceOracleUpdatePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &sourceOracleUpdatePasswordAuthentication, "", true, true); err == nil {
		u.SourceOracleUpdatePasswordAuthentication = &sourceOracleUpdatePasswordAuthentication
		u.Type = SourceOracleUpdateSSHTunnelMethodTypeSourceOracleUpdatePasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceOracleUpdateSSHTunnelMethod", string(data))
}

func (u SourceOracleUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceOracleUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.SourceOracleUpdateNoTunnel, "", true)
	}

	if u.SourceOracleUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceOracleUpdateSSHKeyAuthentication, "", true)
	}

	if u.SourceOracleUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.SourceOracleUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type SourceOracleUpdateSSHTunnelMethod: all fields are null")
}

type SourceOracleUpdate struct {
	// Connect data that will be used for DB connection
	ConnectionData *ConnectBy `json:"connection_data,omitempty"`
	// The encryption method with is used when communicating with the database.
	Encryption *SourceOracleUpdateEncryption `json:"encryption,omitempty"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	// Oracle Corporations recommends the following port numbers:
	// 1521 - Default listening port for client connections to the listener.
	// 2484 - Recommended and officially registered listening port for client connections to the listener using TCP/IP with SSL
	Port *int64 `default:"1521" json:"port"`
	// The list of schemas to sync from. Defaults to user. Case sensitive.
	Schemas []string `json:"schemas,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceOracleUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (s SourceOracleUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleUpdate) GetConnectionData() *ConnectBy {
	if o == nil {
		return nil
	}
	return o.ConnectionData
}

func (o *SourceOracleUpdate) GetEncryption() *SourceOracleUpdateEncryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *SourceOracleUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceOracleUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceOracleUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceOracleUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceOracleUpdate) GetSchemas() []string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *SourceOracleUpdate) GetTunnelMethod() *SourceOracleUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *SourceOracleUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
