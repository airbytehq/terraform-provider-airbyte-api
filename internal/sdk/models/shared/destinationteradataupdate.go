// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode string

const (
	DestinationTeradataUpdateSchemasSSLModeSSLModes6ModeVerifyFull DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode = "verify-full"
)

func (e DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode) ToPointer() *DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode {
	return &e
}
func (e *DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode: %v", v)
	}
}

// DestinationTeradataUpdateVerifyFull - Verify-full SSL mode.
type DestinationTeradataUpdateVerifyFull struct {
	mode *DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode `const:"verify-full" json:"mode"`
	// Specifies the file name of a PEM file that contains Certificate Authority (CA) certificates for use with SSLMODE=verify-full.
	//  See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLCA"> in the docs</a>.
	SslCaCertificate string `json:"ssl_ca_certificate"`
}

func (d DestinationTeradataUpdateVerifyFull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdateVerifyFull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdateVerifyFull) GetMode() *DestinationTeradataUpdateSchemasSSLModeSSLModes6Mode {
	return DestinationTeradataUpdateSchemasSSLModeSSLModes6ModeVerifyFull.ToPointer()
}

func (o *DestinationTeradataUpdateVerifyFull) GetSslCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.SslCaCertificate
}

type DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode string

const (
	DestinationTeradataUpdateSchemasSSLModeSSLModes5ModeVerifyCa DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode = "verify-ca"
)

func (e DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode) ToPointer() *DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode {
	return &e
}
func (e *DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-ca":
		*e = DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode: %v", v)
	}
}

// DestinationTeradataUpdateVerifyCa - Verify-ca SSL mode.
type DestinationTeradataUpdateVerifyCa struct {
	mode *DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode `const:"verify-ca" json:"mode"`
	// Specifies the file name of a PEM file that contains Certificate Authority (CA) certificates for use with SSLMODE=verify-ca.
	//  See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLCA"> in the docs</a>.
	SslCaCertificate string `json:"ssl_ca_certificate"`
}

func (d DestinationTeradataUpdateVerifyCa) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdateVerifyCa) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdateVerifyCa) GetMode() *DestinationTeradataUpdateSchemasSSLModeSSLModes5Mode {
	return DestinationTeradataUpdateSchemasSSLModeSSLModes5ModeVerifyCa.ToPointer()
}

func (o *DestinationTeradataUpdateVerifyCa) GetSslCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.SslCaCertificate
}

type DestinationTeradataUpdateSchemasSSLModeSSLModesMode string

const (
	DestinationTeradataUpdateSchemasSSLModeSSLModesModeRequire DestinationTeradataUpdateSchemasSSLModeSSLModesMode = "require"
)

func (e DestinationTeradataUpdateSchemasSSLModeSSLModesMode) ToPointer() *DestinationTeradataUpdateSchemasSSLModeSSLModesMode {
	return &e
}
func (e *DestinationTeradataUpdateSchemasSSLModeSSLModesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		*e = DestinationTeradataUpdateSchemasSSLModeSSLModesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTeradataUpdateSchemasSSLModeSSLModesMode: %v", v)
	}
}

// DestinationTeradataUpdateRequire - Require SSL mode.
type DestinationTeradataUpdateRequire struct {
	mode *DestinationTeradataUpdateSchemasSSLModeSSLModesMode `const:"require" json:"mode"`
}

func (d DestinationTeradataUpdateRequire) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdateRequire) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdateRequire) GetMode() *DestinationTeradataUpdateSchemasSSLModeSSLModesMode {
	return DestinationTeradataUpdateSchemasSSLModeSSLModesModeRequire.ToPointer()
}

type DestinationTeradataUpdateSchemasSslModeMode string

const (
	DestinationTeradataUpdateSchemasSslModeModePrefer DestinationTeradataUpdateSchemasSslModeMode = "prefer"
)

func (e DestinationTeradataUpdateSchemasSslModeMode) ToPointer() *DestinationTeradataUpdateSchemasSslModeMode {
	return &e
}
func (e *DestinationTeradataUpdateSchemasSslModeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prefer":
		*e = DestinationTeradataUpdateSchemasSslModeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTeradataUpdateSchemasSslModeMode: %v", v)
	}
}

// DestinationTeradataUpdatePrefer - Prefer SSL mode.
type DestinationTeradataUpdatePrefer struct {
	mode *DestinationTeradataUpdateSchemasSslModeMode `const:"prefer" json:"mode"`
}

func (d DestinationTeradataUpdatePrefer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdatePrefer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdatePrefer) GetMode() *DestinationTeradataUpdateSchemasSslModeMode {
	return DestinationTeradataUpdateSchemasSslModeModePrefer.ToPointer()
}

type DestinationTeradataUpdateSchemasMode string

const (
	DestinationTeradataUpdateSchemasModeAllow DestinationTeradataUpdateSchemasMode = "allow"
)

func (e DestinationTeradataUpdateSchemasMode) ToPointer() *DestinationTeradataUpdateSchemasMode {
	return &e
}
func (e *DestinationTeradataUpdateSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		*e = DestinationTeradataUpdateSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTeradataUpdateSchemasMode: %v", v)
	}
}

// DestinationTeradataUpdateAllow - Allow SSL mode.
type DestinationTeradataUpdateAllow struct {
	mode *DestinationTeradataUpdateSchemasMode `const:"allow" json:"mode"`
}

func (d DestinationTeradataUpdateAllow) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdateAllow) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdateAllow) GetMode() *DestinationTeradataUpdateSchemasMode {
	return DestinationTeradataUpdateSchemasModeAllow.ToPointer()
}

type DestinationTeradataUpdateMode string

const (
	DestinationTeradataUpdateModeDisable DestinationTeradataUpdateMode = "disable"
)

func (e DestinationTeradataUpdateMode) ToPointer() *DestinationTeradataUpdateMode {
	return &e
}
func (e *DestinationTeradataUpdateMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = DestinationTeradataUpdateMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationTeradataUpdateMode: %v", v)
	}
}

// DestinationTeradataUpdateDisable - Disable SSL.
type DestinationTeradataUpdateDisable struct {
	mode *DestinationTeradataUpdateMode `const:"disable" json:"mode"`
}

func (d DestinationTeradataUpdateDisable) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdateDisable) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdateDisable) GetMode() *DestinationTeradataUpdateMode {
	return DestinationTeradataUpdateModeDisable.ToPointer()
}

type DestinationTeradataUpdateSSLModesType string

const (
	DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateDisable    DestinationTeradataUpdateSSLModesType = "destination-teradata-update_disable"
	DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateAllow      DestinationTeradataUpdateSSLModesType = "destination-teradata-update_allow"
	DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdatePrefer     DestinationTeradataUpdateSSLModesType = "destination-teradata-update_prefer"
	DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateRequire    DestinationTeradataUpdateSSLModesType = "destination-teradata-update_require"
	DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateVerifyCa   DestinationTeradataUpdateSSLModesType = "destination-teradata-update_verify-ca"
	DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateVerifyFull DestinationTeradataUpdateSSLModesType = "destination-teradata-update_verify-full"
)

// DestinationTeradataUpdateSSLModes - SSL connection modes.
//
//	<b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
//	<b>allow</b> - Chose this mode to enable encryption only when required by the destination database
//	<b>prefer</b> - Chose this mode to allow unencrypted connection only if the destination database does not support encryption
//	<b>require</b> - Chose this mode to always require encryption. If the destination database server does not support encryption, connection will fail
//	 <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the destination database server has a valid SSL certificate
//	 <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the destination database server
//	See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLMODE"> in the docs</a>.
type DestinationTeradataUpdateSSLModes struct {
	DestinationTeradataUpdateDisable    *DestinationTeradataUpdateDisable
	DestinationTeradataUpdateAllow      *DestinationTeradataUpdateAllow
	DestinationTeradataUpdatePrefer     *DestinationTeradataUpdatePrefer
	DestinationTeradataUpdateRequire    *DestinationTeradataUpdateRequire
	DestinationTeradataUpdateVerifyCa   *DestinationTeradataUpdateVerifyCa
	DestinationTeradataUpdateVerifyFull *DestinationTeradataUpdateVerifyFull

	Type DestinationTeradataUpdateSSLModesType
}

func CreateDestinationTeradataUpdateSSLModesDestinationTeradataUpdateDisable(destinationTeradataUpdateDisable DestinationTeradataUpdateDisable) DestinationTeradataUpdateSSLModes {
	typ := DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateDisable

	return DestinationTeradataUpdateSSLModes{
		DestinationTeradataUpdateDisable: &destinationTeradataUpdateDisable,
		Type:                             typ,
	}
}

func CreateDestinationTeradataUpdateSSLModesDestinationTeradataUpdateAllow(destinationTeradataUpdateAllow DestinationTeradataUpdateAllow) DestinationTeradataUpdateSSLModes {
	typ := DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateAllow

	return DestinationTeradataUpdateSSLModes{
		DestinationTeradataUpdateAllow: &destinationTeradataUpdateAllow,
		Type:                           typ,
	}
}

func CreateDestinationTeradataUpdateSSLModesDestinationTeradataUpdatePrefer(destinationTeradataUpdatePrefer DestinationTeradataUpdatePrefer) DestinationTeradataUpdateSSLModes {
	typ := DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdatePrefer

	return DestinationTeradataUpdateSSLModes{
		DestinationTeradataUpdatePrefer: &destinationTeradataUpdatePrefer,
		Type:                            typ,
	}
}

func CreateDestinationTeradataUpdateSSLModesDestinationTeradataUpdateRequire(destinationTeradataUpdateRequire DestinationTeradataUpdateRequire) DestinationTeradataUpdateSSLModes {
	typ := DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateRequire

	return DestinationTeradataUpdateSSLModes{
		DestinationTeradataUpdateRequire: &destinationTeradataUpdateRequire,
		Type:                             typ,
	}
}

func CreateDestinationTeradataUpdateSSLModesDestinationTeradataUpdateVerifyCa(destinationTeradataUpdateVerifyCa DestinationTeradataUpdateVerifyCa) DestinationTeradataUpdateSSLModes {
	typ := DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateVerifyCa

	return DestinationTeradataUpdateSSLModes{
		DestinationTeradataUpdateVerifyCa: &destinationTeradataUpdateVerifyCa,
		Type:                              typ,
	}
}

func CreateDestinationTeradataUpdateSSLModesDestinationTeradataUpdateVerifyFull(destinationTeradataUpdateVerifyFull DestinationTeradataUpdateVerifyFull) DestinationTeradataUpdateSSLModes {
	typ := DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateVerifyFull

	return DestinationTeradataUpdateSSLModes{
		DestinationTeradataUpdateVerifyFull: &destinationTeradataUpdateVerifyFull,
		Type:                                typ,
	}
}

func (u *DestinationTeradataUpdateSSLModes) UnmarshalJSON(data []byte) error {

	var destinationTeradataUpdateDisable DestinationTeradataUpdateDisable = DestinationTeradataUpdateDisable{}
	if err := utils.UnmarshalJSON(data, &destinationTeradataUpdateDisable, "", true, true); err == nil {
		u.DestinationTeradataUpdateDisable = &destinationTeradataUpdateDisable
		u.Type = DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateDisable
		return nil
	}

	var destinationTeradataUpdateAllow DestinationTeradataUpdateAllow = DestinationTeradataUpdateAllow{}
	if err := utils.UnmarshalJSON(data, &destinationTeradataUpdateAllow, "", true, true); err == nil {
		u.DestinationTeradataUpdateAllow = &destinationTeradataUpdateAllow
		u.Type = DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateAllow
		return nil
	}

	var destinationTeradataUpdatePrefer DestinationTeradataUpdatePrefer = DestinationTeradataUpdatePrefer{}
	if err := utils.UnmarshalJSON(data, &destinationTeradataUpdatePrefer, "", true, true); err == nil {
		u.DestinationTeradataUpdatePrefer = &destinationTeradataUpdatePrefer
		u.Type = DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdatePrefer
		return nil
	}

	var destinationTeradataUpdateRequire DestinationTeradataUpdateRequire = DestinationTeradataUpdateRequire{}
	if err := utils.UnmarshalJSON(data, &destinationTeradataUpdateRequire, "", true, true); err == nil {
		u.DestinationTeradataUpdateRequire = &destinationTeradataUpdateRequire
		u.Type = DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateRequire
		return nil
	}

	var destinationTeradataUpdateVerifyCa DestinationTeradataUpdateVerifyCa = DestinationTeradataUpdateVerifyCa{}
	if err := utils.UnmarshalJSON(data, &destinationTeradataUpdateVerifyCa, "", true, true); err == nil {
		u.DestinationTeradataUpdateVerifyCa = &destinationTeradataUpdateVerifyCa
		u.Type = DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateVerifyCa
		return nil
	}

	var destinationTeradataUpdateVerifyFull DestinationTeradataUpdateVerifyFull = DestinationTeradataUpdateVerifyFull{}
	if err := utils.UnmarshalJSON(data, &destinationTeradataUpdateVerifyFull, "", true, true); err == nil {
		u.DestinationTeradataUpdateVerifyFull = &destinationTeradataUpdateVerifyFull
		u.Type = DestinationTeradataUpdateSSLModesTypeDestinationTeradataUpdateVerifyFull
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationTeradataUpdateSSLModes", string(data))
}

func (u DestinationTeradataUpdateSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationTeradataUpdateDisable != nil {
		return utils.MarshalJSON(u.DestinationTeradataUpdateDisable, "", true)
	}

	if u.DestinationTeradataUpdateAllow != nil {
		return utils.MarshalJSON(u.DestinationTeradataUpdateAllow, "", true)
	}

	if u.DestinationTeradataUpdatePrefer != nil {
		return utils.MarshalJSON(u.DestinationTeradataUpdatePrefer, "", true)
	}

	if u.DestinationTeradataUpdateRequire != nil {
		return utils.MarshalJSON(u.DestinationTeradataUpdateRequire, "", true)
	}

	if u.DestinationTeradataUpdateVerifyCa != nil {
		return utils.MarshalJSON(u.DestinationTeradataUpdateVerifyCa, "", true)
	}

	if u.DestinationTeradataUpdateVerifyFull != nil {
		return utils.MarshalJSON(u.DestinationTeradataUpdateVerifyFull, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationTeradataUpdateSSLModes: all fields are null")
}

type DestinationTeradataUpdate struct {
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema *string `default:"airbyte_td" json:"schema"`
	// Encrypt data using SSL. When activating SSL, please select one of the connection modes.
	Ssl *bool `default:"false" json:"ssl"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the destination database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the destination database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the destination database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the destination database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the destination database server
	//  See more information - <a href="https://teradata-docs.s3.amazonaws.com/doc/connectivity/jdbc/reference/current/jdbcug_chapter_2.html#URL_SSLMODE"> in the docs</a>.
	SslMode *DestinationTeradataUpdateSSLModes `json:"ssl_mode,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationTeradataUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationTeradataUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationTeradataUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationTeradataUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationTeradataUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationTeradataUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationTeradataUpdate) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *DestinationTeradataUpdate) GetSslMode() *DestinationTeradataUpdateSSLModes {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *DestinationTeradataUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
