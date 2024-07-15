// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

// DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}
func (e *DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

type DestinationRedshiftUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationRedshiftUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelMethod() DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationRedshiftUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationRedshiftUpdateSchemasTunnelMethod string

const (
	DestinationRedshiftUpdateSchemasTunnelMethodSSHKeyAuth DestinationRedshiftUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationRedshiftUpdateSchemasTunnelMethod) ToPointer() *DestinationRedshiftUpdateSchemasTunnelMethod {
	return &e
}
func (e *DestinationRedshiftUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationRedshiftUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateSchemasTunnelMethod: %v", v)
	}
}

type DestinationRedshiftUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationRedshiftUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationRedshiftUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationRedshiftUpdateSchemasTunnelMethod {
	return DestinationRedshiftUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationRedshiftUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationRedshiftUpdateTunnelMethod string

const (
	DestinationRedshiftUpdateTunnelMethodNoTunnel DestinationRedshiftUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationRedshiftUpdateTunnelMethod) ToPointer() *DestinationRedshiftUpdateTunnelMethod {
	return &e
}
func (e *DestinationRedshiftUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationRedshiftUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateTunnelMethod: %v", v)
	}
}

type DestinationRedshiftUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationRedshiftUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationRedshiftUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdateNoTunnel) GetTunnelMethod() DestinationRedshiftUpdateTunnelMethod {
	return DestinationRedshiftUpdateTunnelMethodNoTunnel
}

type DestinationRedshiftUpdateSSHTunnelMethodType string

const (
	DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateNoTunnel               DestinationRedshiftUpdateSSHTunnelMethodType = "destination-redshift-update_No Tunnel"
	DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateSSHKeyAuthentication   DestinationRedshiftUpdateSSHTunnelMethodType = "destination-redshift-update_SSH Key Authentication"
	DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdatePasswordAuthentication DestinationRedshiftUpdateSSHTunnelMethodType = "destination-redshift-update_Password Authentication"
)

// DestinationRedshiftUpdateSSHTunnelMethod - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedshiftUpdateSSHTunnelMethod struct {
	DestinationRedshiftUpdateNoTunnel               *DestinationRedshiftUpdateNoTunnel
	DestinationRedshiftUpdateSSHKeyAuthentication   *DestinationRedshiftUpdateSSHKeyAuthentication
	DestinationRedshiftUpdatePasswordAuthentication *DestinationRedshiftUpdatePasswordAuthentication

	Type DestinationRedshiftUpdateSSHTunnelMethodType
}

func CreateDestinationRedshiftUpdateSSHTunnelMethodDestinationRedshiftUpdateNoTunnel(destinationRedshiftUpdateNoTunnel DestinationRedshiftUpdateNoTunnel) DestinationRedshiftUpdateSSHTunnelMethod {
	typ := DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateNoTunnel

	return DestinationRedshiftUpdateSSHTunnelMethod{
		DestinationRedshiftUpdateNoTunnel: &destinationRedshiftUpdateNoTunnel,
		Type:                              typ,
	}
}

func CreateDestinationRedshiftUpdateSSHTunnelMethodDestinationRedshiftUpdateSSHKeyAuthentication(destinationRedshiftUpdateSSHKeyAuthentication DestinationRedshiftUpdateSSHKeyAuthentication) DestinationRedshiftUpdateSSHTunnelMethod {
	typ := DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateSSHKeyAuthentication

	return DestinationRedshiftUpdateSSHTunnelMethod{
		DestinationRedshiftUpdateSSHKeyAuthentication: &destinationRedshiftUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationRedshiftUpdateSSHTunnelMethodDestinationRedshiftUpdatePasswordAuthentication(destinationRedshiftUpdatePasswordAuthentication DestinationRedshiftUpdatePasswordAuthentication) DestinationRedshiftUpdateSSHTunnelMethod {
	typ := DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdatePasswordAuthentication

	return DestinationRedshiftUpdateSSHTunnelMethod{
		DestinationRedshiftUpdatePasswordAuthentication: &destinationRedshiftUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationRedshiftUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	var destinationRedshiftUpdateNoTunnel DestinationRedshiftUpdateNoTunnel = DestinationRedshiftUpdateNoTunnel{}
	if err := utils.UnmarshalJSON(data, &destinationRedshiftUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationRedshiftUpdateNoTunnel = &destinationRedshiftUpdateNoTunnel
		u.Type = DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateNoTunnel
		return nil
	}

	var destinationRedshiftUpdateSSHKeyAuthentication DestinationRedshiftUpdateSSHKeyAuthentication = DestinationRedshiftUpdateSSHKeyAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationRedshiftUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationRedshiftUpdateSSHKeyAuthentication = &destinationRedshiftUpdateSSHKeyAuthentication
		u.Type = DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateSSHKeyAuthentication
		return nil
	}

	var destinationRedshiftUpdatePasswordAuthentication DestinationRedshiftUpdatePasswordAuthentication = DestinationRedshiftUpdatePasswordAuthentication{}
	if err := utils.UnmarshalJSON(data, &destinationRedshiftUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationRedshiftUpdatePasswordAuthentication = &destinationRedshiftUpdatePasswordAuthentication
		u.Type = DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdatePasswordAuthentication
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationRedshiftUpdateSSHTunnelMethod", string(data))
}

func (u DestinationRedshiftUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationRedshiftUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationRedshiftUpdateNoTunnel, "", true)
	}

	if u.DestinationRedshiftUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationRedshiftUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationRedshiftUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationRedshiftUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationRedshiftUpdateSSHTunnelMethod: all fields are null")
}

type DestinationRedshiftUpdateMethod string

const (
	DestinationRedshiftUpdateMethodS3Staging DestinationRedshiftUpdateMethod = "S3 Staging"
)

func (e DestinationRedshiftUpdateMethod) ToPointer() *DestinationRedshiftUpdateMethod {
	return &e
}
func (e *DestinationRedshiftUpdateMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3 Staging":
		*e = DestinationRedshiftUpdateMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateMethod: %v", v)
	}
}

// DestinationRedshiftUpdateS3BucketRegion - The region of the S3 staging bucket.
type DestinationRedshiftUpdateS3BucketRegion string

const (
	DestinationRedshiftUpdateS3BucketRegionUnknown      DestinationRedshiftUpdateS3BucketRegion = ""
	DestinationRedshiftUpdateS3BucketRegionAfSouth1     DestinationRedshiftUpdateS3BucketRegion = "af-south-1"
	DestinationRedshiftUpdateS3BucketRegionApEast1      DestinationRedshiftUpdateS3BucketRegion = "ap-east-1"
	DestinationRedshiftUpdateS3BucketRegionApNortheast1 DestinationRedshiftUpdateS3BucketRegion = "ap-northeast-1"
	DestinationRedshiftUpdateS3BucketRegionApNortheast2 DestinationRedshiftUpdateS3BucketRegion = "ap-northeast-2"
	DestinationRedshiftUpdateS3BucketRegionApNortheast3 DestinationRedshiftUpdateS3BucketRegion = "ap-northeast-3"
	DestinationRedshiftUpdateS3BucketRegionApSouth1     DestinationRedshiftUpdateS3BucketRegion = "ap-south-1"
	DestinationRedshiftUpdateS3BucketRegionApSouth2     DestinationRedshiftUpdateS3BucketRegion = "ap-south-2"
	DestinationRedshiftUpdateS3BucketRegionApSoutheast1 DestinationRedshiftUpdateS3BucketRegion = "ap-southeast-1"
	DestinationRedshiftUpdateS3BucketRegionApSoutheast2 DestinationRedshiftUpdateS3BucketRegion = "ap-southeast-2"
	DestinationRedshiftUpdateS3BucketRegionApSoutheast3 DestinationRedshiftUpdateS3BucketRegion = "ap-southeast-3"
	DestinationRedshiftUpdateS3BucketRegionApSoutheast4 DestinationRedshiftUpdateS3BucketRegion = "ap-southeast-4"
	DestinationRedshiftUpdateS3BucketRegionCaCentral1   DestinationRedshiftUpdateS3BucketRegion = "ca-central-1"
	DestinationRedshiftUpdateS3BucketRegionCaWest1      DestinationRedshiftUpdateS3BucketRegion = "ca-west-1"
	DestinationRedshiftUpdateS3BucketRegionCnNorth1     DestinationRedshiftUpdateS3BucketRegion = "cn-north-1"
	DestinationRedshiftUpdateS3BucketRegionCnNorthwest1 DestinationRedshiftUpdateS3BucketRegion = "cn-northwest-1"
	DestinationRedshiftUpdateS3BucketRegionEuCentral1   DestinationRedshiftUpdateS3BucketRegion = "eu-central-1"
	DestinationRedshiftUpdateS3BucketRegionEuCentral2   DestinationRedshiftUpdateS3BucketRegion = "eu-central-2"
	DestinationRedshiftUpdateS3BucketRegionEuNorth1     DestinationRedshiftUpdateS3BucketRegion = "eu-north-1"
	DestinationRedshiftUpdateS3BucketRegionEuSouth1     DestinationRedshiftUpdateS3BucketRegion = "eu-south-1"
	DestinationRedshiftUpdateS3BucketRegionEuSouth2     DestinationRedshiftUpdateS3BucketRegion = "eu-south-2"
	DestinationRedshiftUpdateS3BucketRegionEuWest1      DestinationRedshiftUpdateS3BucketRegion = "eu-west-1"
	DestinationRedshiftUpdateS3BucketRegionEuWest2      DestinationRedshiftUpdateS3BucketRegion = "eu-west-2"
	DestinationRedshiftUpdateS3BucketRegionEuWest3      DestinationRedshiftUpdateS3BucketRegion = "eu-west-3"
	DestinationRedshiftUpdateS3BucketRegionIlCentral1   DestinationRedshiftUpdateS3BucketRegion = "il-central-1"
	DestinationRedshiftUpdateS3BucketRegionMeCentral1   DestinationRedshiftUpdateS3BucketRegion = "me-central-1"
	DestinationRedshiftUpdateS3BucketRegionMeSouth1     DestinationRedshiftUpdateS3BucketRegion = "me-south-1"
	DestinationRedshiftUpdateS3BucketRegionSaEast1      DestinationRedshiftUpdateS3BucketRegion = "sa-east-1"
	DestinationRedshiftUpdateS3BucketRegionUsEast1      DestinationRedshiftUpdateS3BucketRegion = "us-east-1"
	DestinationRedshiftUpdateS3BucketRegionUsEast2      DestinationRedshiftUpdateS3BucketRegion = "us-east-2"
	DestinationRedshiftUpdateS3BucketRegionUsGovEast1   DestinationRedshiftUpdateS3BucketRegion = "us-gov-east-1"
	DestinationRedshiftUpdateS3BucketRegionUsGovWest1   DestinationRedshiftUpdateS3BucketRegion = "us-gov-west-1"
	DestinationRedshiftUpdateS3BucketRegionUsWest1      DestinationRedshiftUpdateS3BucketRegion = "us-west-1"
	DestinationRedshiftUpdateS3BucketRegionUsWest2      DestinationRedshiftUpdateS3BucketRegion = "us-west-2"
)

func (e DestinationRedshiftUpdateS3BucketRegion) ToPointer() *DestinationRedshiftUpdateS3BucketRegion {
	return &e
}
func (e *DestinationRedshiftUpdateS3BucketRegion) UnmarshalJSON(data []byte) error {
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
		*e = DestinationRedshiftUpdateS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateS3BucketRegion: %v", v)
	}
}

// AWSS3Staging - <i>(recommended)</i> Uploads data to S3 and then uses a COPY to insert the data into Redshift. COPY is recommended for production workloads for better speed and scalability. See <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html">AWS docs</a> for more details.
type AWSS3Staging struct {
	// This ID grants access to the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket. See <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">AWS docs</a> on how to generate an access key ID and secret access key.
	AccessKeyID string `json:"access_key_id"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string                         `json:"file_name_pattern,omitempty"`
	method          DestinationRedshiftUpdateMethod `const:"S3 Staging" json:"method"`
	// Whether to delete the staging files from S3 after completing the sync. See <a href="https://docs.airbyte.com/integrations/destinations/redshift/#:~:text=the%20root%20directory.-,Purge%20Staging%20Data,-Whether%20to%20delete"> docs</a> for details.
	PurgeStagingData *bool `default:"true" json:"purge_staging_data"`
	// The name of the staging S3 bucket.
	S3BucketName string `json:"s3_bucket_name"`
	// The directory under the S3 bucket where data will be written. If not provided, then defaults to the root directory. See <a href="https://docs.aws.amazon.com/prescriptive-guidance/latest/defining-bucket-names-data-lakes/faq.html#:~:text=be%20globally%20unique.-,For%20S3%20bucket%20paths,-%2C%20you%20can%20use">path's name recommendations</a> for more details.
	S3BucketPath *string `json:"s3_bucket_path,omitempty"`
	// The region of the S3 staging bucket.
	S3BucketRegion *DestinationRedshiftUpdateS3BucketRegion `default:"" json:"s3_bucket_region"`
	// The corresponding secret to the above access key id. See <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">AWS docs</a> on how to generate an access key ID and secret access key.
	SecretAccessKey string `json:"secret_access_key"`
}

func (a AWSS3Staging) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AWSS3Staging) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AWSS3Staging) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *AWSS3Staging) GetFileNamePattern() *string {
	if o == nil {
		return nil
	}
	return o.FileNamePattern
}

func (o *AWSS3Staging) GetMethod() DestinationRedshiftUpdateMethod {
	return DestinationRedshiftUpdateMethodS3Staging
}

func (o *AWSS3Staging) GetPurgeStagingData() *bool {
	if o == nil {
		return nil
	}
	return o.PurgeStagingData
}

func (o *AWSS3Staging) GetS3BucketName() string {
	if o == nil {
		return ""
	}
	return o.S3BucketName
}

func (o *AWSS3Staging) GetS3BucketPath() *string {
	if o == nil {
		return nil
	}
	return o.S3BucketPath
}

func (o *AWSS3Staging) GetS3BucketRegion() *DestinationRedshiftUpdateS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.S3BucketRegion
}

func (o *AWSS3Staging) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

type UploadingMethodType string

const (
	UploadingMethodTypeAWSS3Staging UploadingMethodType = "AWS S3 Staging"
)

// UploadingMethod - The way data will be uploaded to Redshift.
type UploadingMethod struct {
	AWSS3Staging *AWSS3Staging

	Type UploadingMethodType
}

func CreateUploadingMethodAWSS3Staging(awsS3Staging AWSS3Staging) UploadingMethod {
	typ := UploadingMethodTypeAWSS3Staging

	return UploadingMethod{
		AWSS3Staging: &awsS3Staging,
		Type:         typ,
	}
}

func (u *UploadingMethod) UnmarshalJSON(data []byte) error {

	var awsS3Staging AWSS3Staging = AWSS3Staging{}
	if err := utils.UnmarshalJSON(data, &awsS3Staging, "", true, true); err == nil {
		u.AWSS3Staging = &awsS3Staging
		u.Type = UploadingMethodTypeAWSS3Staging
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for UploadingMethod", string(data))
}

func (u UploadingMethod) MarshalJSON() ([]byte, error) {
	if u.AWSS3Staging != nil {
		return utils.MarshalJSON(u.AWSS3Staging, "", true)
	}

	return nil, errors.New("could not marshal union type UploadingMethod: all fields are null")
}

type DestinationRedshiftUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions
	DisableTypeDedupe *bool `default:"false" json:"disable_type_dedupe"`
	// Drop tables with CASCADE. WARNING! This will delete all data in all dependent objects (views, etc.). Use with caution. This option is intended for usecases which can easily rebuild the dependent objects.
	DropCascade *bool `default:"false" json:"drop_cascade"`
	// Host Endpoint of the Redshift Cluster (must include the cluster-id, region and end with .redshift.amazonaws.com)
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password string `json:"password"`
	// Port of the database.
	Port *int64 `default:"5439" json:"port"`
	// The schema to write raw tables into (default: airbyte_internal).
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// The default schema tables are written to if the source does not specify a namespace. Unless specifically configured, the usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationRedshiftUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The way data will be uploaded to Redshift.
	UploadingMethod *UploadingMethod `json:"uploading_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationRedshiftUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationRedshiftUpdate) GetDisableTypeDedupe() *bool {
	if o == nil {
		return nil
	}
	return o.DisableTypeDedupe
}

func (o *DestinationRedshiftUpdate) GetDropCascade() *bool {
	if o == nil {
		return nil
	}
	return o.DropCascade
}

func (o *DestinationRedshiftUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationRedshiftUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationRedshiftUpdate) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *DestinationRedshiftUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationRedshiftUpdate) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationRedshiftUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationRedshiftUpdate) GetTunnelMethod() *DestinationRedshiftUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationRedshiftUpdate) GetUploadingMethod() *UploadingMethod {
	if o == nil {
		return nil
	}
	return o.UploadingMethod
}

func (o *DestinationRedshiftUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
