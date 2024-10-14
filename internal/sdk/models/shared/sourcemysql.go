// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceMysqlSchemasMethod string

const (
	SourceMysqlSchemasMethodStandard SourceMysqlSchemasMethod = "STANDARD"
)

func (e SourceMysqlSchemasMethod) ToPointer() *SourceMysqlSchemasMethod {
	return &e
}
func (e *SourceMysqlSchemasMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		*e = SourceMysqlSchemasMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlSchemasMethod: %v", v)
	}
}

// SourceMysqlScanChangesWithUserDefinedCursor - Incrementally detects new inserts and updates using the <a href="https://docs.airbyte.com/understanding-airbyte/connections/incremental-append/#user-defined-cursor">cursor column</a> chosen when configuring a connection (e.g. created_at, updated_at).
type SourceMysqlScanChangesWithUserDefinedCursor struct {
	method SourceMysqlSchemasMethod `const:"STANDARD" json:"method"`
}

func (s SourceMysqlScanChangesWithUserDefinedCursor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlScanChangesWithUserDefinedCursor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlScanChangesWithUserDefinedCursor) GetMethod() SourceMysqlSchemasMethod {
	return SourceMysqlSchemasMethodStandard
}

// SourceMysqlInvalidCDCPositionBehaviorAdvanced - Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss.
type SourceMysqlInvalidCDCPositionBehaviorAdvanced string

const (
	SourceMysqlInvalidCDCPositionBehaviorAdvancedFailSync   SourceMysqlInvalidCDCPositionBehaviorAdvanced = "Fail sync"
	SourceMysqlInvalidCDCPositionBehaviorAdvancedReSyncData SourceMysqlInvalidCDCPositionBehaviorAdvanced = "Re-sync data"
)

func (e SourceMysqlInvalidCDCPositionBehaviorAdvanced) ToPointer() *SourceMysqlInvalidCDCPositionBehaviorAdvanced {
	return &e
}
func (e *SourceMysqlInvalidCDCPositionBehaviorAdvanced) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Fail sync":
		fallthrough
	case "Re-sync data":
		*e = SourceMysqlInvalidCDCPositionBehaviorAdvanced(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlInvalidCDCPositionBehaviorAdvanced: %v", v)
	}
}

type SourceMysqlMethod string

const (
	SourceMysqlMethodCdc SourceMysqlMethod = "CDC"
)

func (e SourceMysqlMethod) ToPointer() *SourceMysqlMethod {
	return &e
}
func (e *SourceMysqlMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CDC":
		*e = SourceMysqlMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlMethod: %v", v)
	}
}

// SourceMysqlReadChangesUsingBinaryLogCDC - <i>Recommended</i> - Incrementally reads new inserts, updates, and deletes using the MySQL <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">binary log</a>. This must be enabled on your database.
type SourceMysqlReadChangesUsingBinaryLogCDC struct {
	// The amount of time an initial load is allowed to continue for before catching up on CDC logs.
	InitialLoadTimeoutHours *int64 `default:"8" json:"initial_load_timeout_hours"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.
	InitialWaitingSeconds *int64 `default:"300" json:"initial_waiting_seconds"`
	// Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss.
	InvalidCdcCursorPositionBehavior *SourceMysqlInvalidCDCPositionBehaviorAdvanced `default:"Fail sync" json:"invalid_cdc_cursor_position_behavior"`
	method                           SourceMysqlMethod                              `const:"CDC" json:"method"`
	// Enter the configured MySQL server timezone. This should only be done if the configured timezone in your MySQL instance does not conform to IANNA standard.
	ServerTimeZone *string `json:"server_time_zone,omitempty"`
}

func (s SourceMysqlReadChangesUsingBinaryLogCDC) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlReadChangesUsingBinaryLogCDC) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlReadChangesUsingBinaryLogCDC) GetInitialLoadTimeoutHours() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialLoadTimeoutHours
}

func (o *SourceMysqlReadChangesUsingBinaryLogCDC) GetInitialWaitingSeconds() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialWaitingSeconds
}

func (o *SourceMysqlReadChangesUsingBinaryLogCDC) GetInvalidCdcCursorPositionBehavior() *SourceMysqlInvalidCDCPositionBehaviorAdvanced {
	if o == nil {
		return nil
	}
	return o.InvalidCdcCursorPositionBehavior
}

func (o *SourceMysqlReadChangesUsingBinaryLogCDC) GetMethod() SourceMysqlMethod {
	return SourceMysqlMethodCdc
}

func (o *SourceMysqlReadChangesUsingBinaryLogCDC) GetServerTimeZone() *string {
	if o == nil {
		return nil
	}
	return o.ServerTimeZone
}

type SourceMysqlUpdateMethodType string

const (
	SourceMysqlUpdateMethodTypeSourceMysqlReadChangesUsingBinaryLogCDC     SourceMysqlUpdateMethodType = "source-mysql_Read Changes using Binary Log (CDC)"
	SourceMysqlUpdateMethodTypeSourceMysqlScanChangesWithUserDefinedCursor SourceMysqlUpdateMethodType = "source-mysql_Scan Changes with User Defined Cursor"
)

// SourceMysqlUpdateMethod - Configures how data is extracted from the database.
type SourceMysqlUpdateMethod struct {
	SourceMysqlReadChangesUsingBinaryLogCDC     *SourceMysqlReadChangesUsingBinaryLogCDC
	SourceMysqlScanChangesWithUserDefinedCursor *SourceMysqlScanChangesWithUserDefinedCursor

	Type SourceMysqlUpdateMethodType
}

func CreateSourceMysqlUpdateMethodSourceMysqlReadChangesUsingBinaryLogCDC(sourceMysqlReadChangesUsingBinaryLogCDC SourceMysqlReadChangesUsingBinaryLogCDC) SourceMysqlUpdateMethod {
	typ := SourceMysqlUpdateMethodTypeSourceMysqlReadChangesUsingBinaryLogCDC

	return SourceMysqlUpdateMethod{
		SourceMysqlReadChangesUsingBinaryLogCDC: &sourceMysqlReadChangesUsingBinaryLogCDC,
		Type:                                    typ,
	}
}

func CreateSourceMysqlUpdateMethodSourceMysqlScanChangesWithUserDefinedCursor(sourceMysqlScanChangesWithUserDefinedCursor SourceMysqlScanChangesWithUserDefinedCursor) SourceMysqlUpdateMethod {
	typ := SourceMysqlUpdateMethodTypeSourceMysqlScanChangesWithUserDefinedCursor

	return SourceMysqlUpdateMethod{
		SourceMysqlScanChangesWithUserDefinedCursor: &sourceMysqlScanChangesWithUserDefinedCursor,
		Type: typ,
	}
}

func (u *SourceMysqlUpdateMethod) UnmarshalJSON(data []byte) error {

	var sourceMysqlScanChangesWithUserDefinedCursor SourceMysqlScanChangesWithUserDefinedCursor = SourceMysqlScanChangesWithUserDefinedCursor{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlScanChangesWithUserDefinedCursor, "", true, true); err == nil {
		u.SourceMysqlScanChangesWithUserDefinedCursor = &sourceMysqlScanChangesWithUserDefinedCursor
		u.Type = SourceMysqlUpdateMethodTypeSourceMysqlScanChangesWithUserDefinedCursor
		return nil
	}

	var sourceMysqlReadChangesUsingBinaryLogCDC SourceMysqlReadChangesUsingBinaryLogCDC = SourceMysqlReadChangesUsingBinaryLogCDC{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlReadChangesUsingBinaryLogCDC, "", true, true); err == nil {
		u.SourceMysqlReadChangesUsingBinaryLogCDC = &sourceMysqlReadChangesUsingBinaryLogCDC
		u.Type = SourceMysqlUpdateMethodTypeSourceMysqlReadChangesUsingBinaryLogCDC
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceMysqlUpdateMethod", string(data))
}

func (u SourceMysqlUpdateMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMysqlReadChangesUsingBinaryLogCDC != nil {
		return utils.MarshalJSON(u.SourceMysqlReadChangesUsingBinaryLogCDC, "", true)
	}

	if u.SourceMysqlScanChangesWithUserDefinedCursor != nil {
		return utils.MarshalJSON(u.SourceMysqlScanChangesWithUserDefinedCursor, "", true)
	}

	return nil, errors.New("could not marshal union type SourceMysqlUpdateMethod: all fields are null")
}

type SourceMysqlMysql string

const (
	SourceMysqlMysqlMysql SourceMysqlMysql = "mysql"
)

func (e SourceMysqlMysql) ToPointer() *SourceMysqlMysql {
	return &e
}
func (e *SourceMysqlMysql) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mysql":
		*e = SourceMysqlMysql(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlMysql: %v", v)
	}
}

type SourceMysqlSchemasSSLModeSSLModesMode string

const (
	SourceMysqlSchemasSSLModeSSLModesModeVerifyIdentity SourceMysqlSchemasSSLModeSSLModesMode = "verify_identity"
)

func (e SourceMysqlSchemasSSLModeSSLModesMode) ToPointer() *SourceMysqlSchemasSSLModeSSLModesMode {
	return &e
}
func (e *SourceMysqlSchemasSSLModeSSLModesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify_identity":
		*e = SourceMysqlSchemasSSLModeSSLModesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlSchemasSSLModeSSLModesMode: %v", v)
	}
}

// SourceMysqlVerifyIdentity - Always connect with SSL. Verify both CA and Hostname.
type SourceMysqlVerifyIdentity struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate (this is not a required field, but if you want to use it, you will need to add the <b>Client key</b> as well)
	ClientCertificate *string `json:"client_certificate,omitempty"`
	// Client key (this is not a required field, but if you want to use it, you will need to add the <b>Client certificate</b> as well)
	ClientKey *string `json:"client_key,omitempty"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                               `json:"client_key_password,omitempty"`
	mode              SourceMysqlSchemasSSLModeSSLModesMode `const:"verify_identity" json:"mode"`
}

func (s SourceMysqlVerifyIdentity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlVerifyIdentity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlVerifyIdentity) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *SourceMysqlVerifyIdentity) GetClientCertificate() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *SourceMysqlVerifyIdentity) GetClientKey() *string {
	if o == nil {
		return nil
	}
	return o.ClientKey
}

func (o *SourceMysqlVerifyIdentity) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *SourceMysqlVerifyIdentity) GetMode() SourceMysqlSchemasSSLModeSSLModesMode {
	return SourceMysqlSchemasSSLModeSSLModesModeVerifyIdentity
}

type SourceMysqlSchemasSslModeMode string

const (
	SourceMysqlSchemasSslModeModeVerifyCa SourceMysqlSchemasSslModeMode = "verify_ca"
)

func (e SourceMysqlSchemasSslModeMode) ToPointer() *SourceMysqlSchemasSslModeMode {
	return &e
}
func (e *SourceMysqlSchemasSslModeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify_ca":
		*e = SourceMysqlSchemasSslModeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlSchemasSslModeMode: %v", v)
	}
}

// SourceMysqlVerifyCA - Always connect with SSL. Verifies CA, but allows connection even if Hostname does not match.
type SourceMysqlVerifyCA struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate (this is not a required field, but if you want to use it, you will need to add the <b>Client key</b> as well)
	ClientCertificate *string `json:"client_certificate,omitempty"`
	// Client key (this is not a required field, but if you want to use it, you will need to add the <b>Client certificate</b> as well)
	ClientKey *string `json:"client_key,omitempty"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                       `json:"client_key_password,omitempty"`
	mode              SourceMysqlSchemasSslModeMode `const:"verify_ca" json:"mode"`
}

func (s SourceMysqlVerifyCA) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlVerifyCA) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlVerifyCA) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *SourceMysqlVerifyCA) GetClientCertificate() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *SourceMysqlVerifyCA) GetClientKey() *string {
	if o == nil {
		return nil
	}
	return o.ClientKey
}

func (o *SourceMysqlVerifyCA) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *SourceMysqlVerifyCA) GetMode() SourceMysqlSchemasSslModeMode {
	return SourceMysqlSchemasSslModeModeVerifyCa
}

type SourceMysqlSchemasMode string

const (
	SourceMysqlSchemasModeRequired SourceMysqlSchemasMode = "required"
)

func (e SourceMysqlSchemasMode) ToPointer() *SourceMysqlSchemasMode {
	return &e
}
func (e *SourceMysqlSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "required":
		*e = SourceMysqlSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlSchemasMode: %v", v)
	}
}

// SourceMysqlRequired - Always connect with SSL. If the MySQL server doesn’t support SSL, the connection will not be established. Certificate Authority (CA) and Hostname are not verified.
type SourceMysqlRequired struct {
	mode SourceMysqlSchemasMode `const:"required" json:"mode"`
}

func (s SourceMysqlRequired) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlRequired) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlRequired) GetMode() SourceMysqlSchemasMode {
	return SourceMysqlSchemasModeRequired
}

type SourceMysqlMode string

const (
	SourceMysqlModePreferred SourceMysqlMode = "preferred"
)

func (e SourceMysqlMode) ToPointer() *SourceMysqlMode {
	return &e
}
func (e *SourceMysqlMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preferred":
		*e = SourceMysqlMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlMode: %v", v)
	}
}

// SourceMysqlPreferred - Automatically attempt SSL connection. If the MySQL server does not support SSL, continue with a regular connection.
type SourceMysqlPreferred struct {
	mode SourceMysqlMode `const:"preferred" json:"mode"`
}

func (s SourceMysqlPreferred) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlPreferred) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlPreferred) GetMode() SourceMysqlMode {
	return SourceMysqlModePreferred
}

type SourceMysqlSSLModesType string

const (
	SourceMysqlSSLModesTypeSourceMysqlPreferred      SourceMysqlSSLModesType = "source-mysql_preferred"
	SourceMysqlSSLModesTypeSourceMysqlRequired       SourceMysqlSSLModesType = "source-mysql_required"
	SourceMysqlSSLModesTypeSourceMysqlVerifyCA       SourceMysqlSSLModesType = "source-mysql_Verify CA"
	SourceMysqlSSLModesTypeSourceMysqlVerifyIdentity SourceMysqlSSLModesType = "source-mysql_Verify Identity"
)

// SourceMysqlSSLModes - SSL connection modes. Read more <a href="https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-using-ssl.html"> in the docs</a>.
type SourceMysqlSSLModes struct {
	SourceMysqlPreferred      *SourceMysqlPreferred
	SourceMysqlRequired       *SourceMysqlRequired
	SourceMysqlVerifyCA       *SourceMysqlVerifyCA
	SourceMysqlVerifyIdentity *SourceMysqlVerifyIdentity

	Type SourceMysqlSSLModesType
}

func CreateSourceMysqlSSLModesSourceMysqlPreferred(sourceMysqlPreferred SourceMysqlPreferred) SourceMysqlSSLModes {
	typ := SourceMysqlSSLModesTypeSourceMysqlPreferred

	return SourceMysqlSSLModes{
		SourceMysqlPreferred: &sourceMysqlPreferred,
		Type:                 typ,
	}
}

func CreateSourceMysqlSSLModesSourceMysqlRequired(sourceMysqlRequired SourceMysqlRequired) SourceMysqlSSLModes {
	typ := SourceMysqlSSLModesTypeSourceMysqlRequired

	return SourceMysqlSSLModes{
		SourceMysqlRequired: &sourceMysqlRequired,
		Type:                typ,
	}
}

func CreateSourceMysqlSSLModesSourceMysqlVerifyCA(sourceMysqlVerifyCA SourceMysqlVerifyCA) SourceMysqlSSLModes {
	typ := SourceMysqlSSLModesTypeSourceMysqlVerifyCA

	return SourceMysqlSSLModes{
		SourceMysqlVerifyCA: &sourceMysqlVerifyCA,
		Type:                typ,
	}
}

func CreateSourceMysqlSSLModesSourceMysqlVerifyIdentity(sourceMysqlVerifyIdentity SourceMysqlVerifyIdentity) SourceMysqlSSLModes {
	typ := SourceMysqlSSLModesTypeSourceMysqlVerifyIdentity

	return SourceMysqlSSLModes{
		SourceMysqlVerifyIdentity: &sourceMysqlVerifyIdentity,
		Type:                      typ,
	}
}

func (u *SourceMysqlSSLModes) UnmarshalJSON(data []byte) error {

	var sourceMysqlPreferred SourceMysqlPreferred = SourceMysqlPreferred{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlPreferred, "", true, true); err == nil {
		u.SourceMysqlPreferred = &sourceMysqlPreferred
		u.Type = SourceMysqlSSLModesTypeSourceMysqlPreferred
		return nil
	}

	var sourceMysqlRequired SourceMysqlRequired = SourceMysqlRequired{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlRequired, "", true, true); err == nil {
		u.SourceMysqlRequired = &sourceMysqlRequired
		u.Type = SourceMysqlSSLModesTypeSourceMysqlRequired
		return nil
	}

	var sourceMysqlVerifyCA SourceMysqlVerifyCA = SourceMysqlVerifyCA{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlVerifyCA, "", true, true); err == nil {
		u.SourceMysqlVerifyCA = &sourceMysqlVerifyCA
		u.Type = SourceMysqlSSLModesTypeSourceMysqlVerifyCA
		return nil
	}

	var sourceMysqlVerifyIdentity SourceMysqlVerifyIdentity = SourceMysqlVerifyIdentity{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlVerifyIdentity, "", true, true); err == nil {
		u.SourceMysqlVerifyIdentity = &sourceMysqlVerifyIdentity
		u.Type = SourceMysqlSSLModesTypeSourceMysqlVerifyIdentity
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceMysqlSSLModes", string(data))
}

func (u SourceMysqlSSLModes) MarshalJSON() ([]byte, error) {
	if u.SourceMysqlPreferred != nil {
		return utils.MarshalJSON(u.SourceMysqlPreferred, "", true)
	}

	if u.SourceMysqlRequired != nil {
		return utils.MarshalJSON(u.SourceMysqlRequired, "", true)
	}

	if u.SourceMysqlVerifyCA != nil {
		return utils.MarshalJSON(u.SourceMysqlVerifyCA, "", true)
	}

	if u.SourceMysqlVerifyIdentity != nil {
		return utils.MarshalJSON(u.SourceMysqlVerifyIdentity, "", true)
	}

	return nil, errors.New("could not marshal union type SourceMysqlSSLModes: all fields are null")
}

// SourceMysqlSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceMysqlSchemasTunnelMethodTunnelMethod string

const (
	SourceMysqlSchemasTunnelMethodTunnelMethodSSHPasswordAuth SourceMysqlSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceMysqlSchemasTunnelMethodTunnelMethod) ToPointer() *SourceMysqlSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *SourceMysqlSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceMysqlSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type SourceMysqlPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod SourceMysqlSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (s SourceMysqlPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceMysqlPasswordAuthentication) GetTunnelMethod() SourceMysqlSchemasTunnelMethodTunnelMethod {
	return SourceMysqlSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *SourceMysqlPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceMysqlPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *SourceMysqlPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// SourceMysqlSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceMysqlSchemasTunnelMethod string

const (
	SourceMysqlSchemasTunnelMethodSSHKeyAuth SourceMysqlSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceMysqlSchemasTunnelMethod) ToPointer() *SourceMysqlSchemasTunnelMethod {
	return &e
}
func (e *SourceMysqlSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceMysqlSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlSchemasTunnelMethod: %v", v)
	}
}

type SourceMysqlSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod SourceMysqlSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (s SourceMysqlSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *SourceMysqlSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceMysqlSSHKeyAuthentication) GetTunnelMethod() SourceMysqlSchemasTunnelMethod {
	return SourceMysqlSchemasTunnelMethodSSHKeyAuth
}

func (o *SourceMysqlSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceMysqlSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// SourceMysqlTunnelMethod - No ssh tunnel needed to connect to database
type SourceMysqlTunnelMethod string

const (
	SourceMysqlTunnelMethodNoTunnel SourceMysqlTunnelMethod = "NO_TUNNEL"
)

func (e SourceMysqlTunnelMethod) ToPointer() *SourceMysqlTunnelMethod {
	return &e
}
func (e *SourceMysqlTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceMysqlTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlTunnelMethod: %v", v)
	}
}

type SourceMysqlNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod SourceMysqlTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (s SourceMysqlNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlNoTunnel) GetTunnelMethod() SourceMysqlTunnelMethod {
	return SourceMysqlTunnelMethodNoTunnel
}

type SourceMysqlSSHTunnelMethodType string

const (
	SourceMysqlSSHTunnelMethodTypeSourceMysqlNoTunnel               SourceMysqlSSHTunnelMethodType = "source-mysql_No Tunnel"
	SourceMysqlSSHTunnelMethodTypeSourceMysqlSSHKeyAuthentication   SourceMysqlSSHTunnelMethodType = "source-mysql_SSH Key Authentication"
	SourceMysqlSSHTunnelMethodTypeSourceMysqlPasswordAuthentication SourceMysqlSSHTunnelMethodType = "source-mysql_Password Authentication"
)

// SourceMysqlSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMysqlSSHTunnelMethod struct {
	SourceMysqlNoTunnel               *SourceMysqlNoTunnel
	SourceMysqlSSHKeyAuthentication   *SourceMysqlSSHKeyAuthentication
	SourceMysqlPasswordAuthentication *SourceMysqlPasswordAuthentication

	Type SourceMysqlSSHTunnelMethodType
}

func CreateSourceMysqlSSHTunnelMethodSourceMysqlNoTunnel(sourceMysqlNoTunnel SourceMysqlNoTunnel) SourceMysqlSSHTunnelMethod {
	typ := SourceMysqlSSHTunnelMethodTypeSourceMysqlNoTunnel

	return SourceMysqlSSHTunnelMethod{
		SourceMysqlNoTunnel: &sourceMysqlNoTunnel,
		Type:                typ,
	}
}

func CreateSourceMysqlSSHTunnelMethodSourceMysqlSSHKeyAuthentication(sourceMysqlSSHKeyAuthentication SourceMysqlSSHKeyAuthentication) SourceMysqlSSHTunnelMethod {
	typ := SourceMysqlSSHTunnelMethodTypeSourceMysqlSSHKeyAuthentication

	return SourceMysqlSSHTunnelMethod{
		SourceMysqlSSHKeyAuthentication: &sourceMysqlSSHKeyAuthentication,
		Type:                            typ,
	}
}

func CreateSourceMysqlSSHTunnelMethodSourceMysqlPasswordAuthentication(sourceMysqlPasswordAuthentication SourceMysqlPasswordAuthentication) SourceMysqlSSHTunnelMethod {
	typ := SourceMysqlSSHTunnelMethodTypeSourceMysqlPasswordAuthentication

	return SourceMysqlSSHTunnelMethod{
		SourceMysqlPasswordAuthentication: &sourceMysqlPasswordAuthentication,
		Type:                              typ,
	}
}

func (u *SourceMysqlSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var sourceMysqlNoTunnel SourceMysqlNoTunnel = SourceMysqlNoTunnel{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlNoTunnel, "", true, true); err == nil {
		u.SourceMysqlNoTunnel = &sourceMysqlNoTunnel
		u.Type = SourceMysqlSSHTunnelMethodTypeSourceMysqlNoTunnel
		return nil
	}

	var sourceMysqlSSHKeyAuthentication SourceMysqlSSHKeyAuthentication = SourceMysqlSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlSSHKeyAuthentication, "", true, true); err == nil {
		u.SourceMysqlSSHKeyAuthentication = &sourceMysqlSSHKeyAuthentication
		u.Type = SourceMysqlSSHTunnelMethodTypeSourceMysqlSSHKeyAuthentication
		return nil
	}

	var sourceMysqlPasswordAuthentication SourceMysqlPasswordAuthentication = SourceMysqlPasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &sourceMysqlPasswordAuthentication, "", true, true); err == nil {
		u.SourceMysqlPasswordAuthentication = &sourceMysqlPasswordAuthentication
		u.Type = SourceMysqlSSHTunnelMethodTypeSourceMysqlPasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceMysqlSSHTunnelMethod", string(data))
}

func (u SourceMysqlSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMysqlNoTunnel != nil {
		return utils.MarshalJSON(u.SourceMysqlNoTunnel, "", true)
	}

	if u.SourceMysqlSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceMysqlSSHKeyAuthentication, "", true)
	}

	if u.SourceMysqlPasswordAuthentication != nil {
		return utils.MarshalJSON(u.SourceMysqlPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type SourceMysqlSSHTunnelMethod: all fields are null")
}

type SourceMysql struct {
	// The database name.
	Database string `json:"database"`
	// The host name of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3). For more information read about <a href="https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-jdbc-url-format.html">JDBC URL parameters</a>.
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// The port to connect to.
	Port *int64 `default:"3306" json:"port"`
	// Configures how data is extracted from the database.
	ReplicationMethod SourceMysqlUpdateMethod `json:"replication_method"`
	sourceType        SourceMysqlMysql        `const:"mysql" json:"sourceType"`
	// Encrypt data using SSL.
	Ssl *bool `default:"true" json:"ssl"`
	// SSL connection modes. Read more <a href="https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-using-ssl.html"> in the docs</a>.
	SslMode *SourceMysqlSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceMysqlSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (s SourceMysql) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysql) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysql) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceMysql) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceMysql) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceMysql) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceMysql) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceMysql) GetReplicationMethod() SourceMysqlUpdateMethod {
	if o == nil {
		return SourceMysqlUpdateMethod{}
	}
	return o.ReplicationMethod
}

func (o *SourceMysql) GetSourceType() SourceMysqlMysql {
	return SourceMysqlMysqlMysql
}

func (o *SourceMysql) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *SourceMysql) GetSslMode() *SourceMysqlSSLModes {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *SourceMysql) GetTunnelMethod() *SourceMysqlSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *SourceMysql) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
