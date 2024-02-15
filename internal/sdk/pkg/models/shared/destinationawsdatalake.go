// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// DestinationAwsDatalakeSchemasCredentialsTitle - Name of the credentials
type DestinationAwsDatalakeSchemasCredentialsTitle string

const (
	DestinationAwsDatalakeSchemasCredentialsTitleIamUser DestinationAwsDatalakeSchemasCredentialsTitle = "IAM User"
)

func (e DestinationAwsDatalakeSchemasCredentialsTitle) ToPointer() *DestinationAwsDatalakeSchemasCredentialsTitle {
	return &e
}

func (e *DestinationAwsDatalakeSchemasCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IAM User":
		*e = DestinationAwsDatalakeSchemasCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeSchemasCredentialsTitle: %v", v)
	}
}

type DestinationAwsDatalakeIAMUser struct {
	// AWS User Access Key Id
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// Secret Access Key
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// Name of the credentials
	credentialsTitle *DestinationAwsDatalakeSchemasCredentialsTitle `const:"IAM User" json:"credentials_title"`
}

func (d DestinationAwsDatalakeIAMUser) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAwsDatalakeIAMUser) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAwsDatalakeIAMUser) GetAwsAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AwsAccessKeyID
}

func (o *DestinationAwsDatalakeIAMUser) GetAwsSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.AwsSecretAccessKey
}

func (o *DestinationAwsDatalakeIAMUser) GetCredentialsTitle() *DestinationAwsDatalakeSchemasCredentialsTitle {
	return DestinationAwsDatalakeSchemasCredentialsTitleIamUser.ToPointer()
}

// DestinationAwsDatalakeCredentialsTitle - Name of the credentials
type DestinationAwsDatalakeCredentialsTitle string

const (
	DestinationAwsDatalakeCredentialsTitleIamRole DestinationAwsDatalakeCredentialsTitle = "IAM Role"
)

func (e DestinationAwsDatalakeCredentialsTitle) ToPointer() *DestinationAwsDatalakeCredentialsTitle {
	return &e
}

func (e *DestinationAwsDatalakeCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IAM Role":
		*e = DestinationAwsDatalakeCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeCredentialsTitle: %v", v)
	}
}

type DestinationAwsDatalakeIAMRole struct {
	// Name of the credentials
	credentialsTitle *DestinationAwsDatalakeCredentialsTitle `const:"IAM Role" json:"credentials_title"`
	// Will assume this role to write data to s3
	RoleArn string `json:"role_arn"`
}

func (d DestinationAwsDatalakeIAMRole) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAwsDatalakeIAMRole) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAwsDatalakeIAMRole) GetCredentialsTitle() *DestinationAwsDatalakeCredentialsTitle {
	return DestinationAwsDatalakeCredentialsTitleIamRole.ToPointer()
}

func (o *DestinationAwsDatalakeIAMRole) GetRoleArn() string {
	if o == nil {
		return ""
	}
	return o.RoleArn
}

type DestinationAwsDatalakeAuthenticationModeType string

const (
	DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeIAMRole DestinationAwsDatalakeAuthenticationModeType = "destination-aws-datalake_IAM Role"
	DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeIAMUser DestinationAwsDatalakeAuthenticationModeType = "destination-aws-datalake_IAM User"
)

// DestinationAwsDatalakeAuthenticationMode - Choose How to Authenticate to AWS.
type DestinationAwsDatalakeAuthenticationMode struct {
	DestinationAwsDatalakeIAMRole *DestinationAwsDatalakeIAMRole
	DestinationAwsDatalakeIAMUser *DestinationAwsDatalakeIAMUser

	Type DestinationAwsDatalakeAuthenticationModeType
}

func CreateDestinationAwsDatalakeAuthenticationModeDestinationAwsDatalakeIAMRole(destinationAwsDatalakeIAMRole DestinationAwsDatalakeIAMRole) DestinationAwsDatalakeAuthenticationMode {
	typ := DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeIAMRole

	return DestinationAwsDatalakeAuthenticationMode{
		DestinationAwsDatalakeIAMRole: &destinationAwsDatalakeIAMRole,
		Type:                          typ,
	}
}

func CreateDestinationAwsDatalakeAuthenticationModeDestinationAwsDatalakeIAMUser(destinationAwsDatalakeIAMUser DestinationAwsDatalakeIAMUser) DestinationAwsDatalakeAuthenticationMode {
	typ := DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeIAMUser

	return DestinationAwsDatalakeAuthenticationMode{
		DestinationAwsDatalakeIAMUser: &destinationAwsDatalakeIAMUser,
		Type:                          typ,
	}
}

func (u *DestinationAwsDatalakeAuthenticationMode) UnmarshalJSON(data []byte) error {

	destinationAwsDatalakeIAMRole := new(DestinationAwsDatalakeIAMRole)
	if err := utils.UnmarshalJSON(data, &destinationAwsDatalakeIAMRole, "", true, true); err == nil {
		u.DestinationAwsDatalakeIAMRole = destinationAwsDatalakeIAMRole
		u.Type = DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeIAMRole
		return nil
	}

	destinationAwsDatalakeIAMUser := new(DestinationAwsDatalakeIAMUser)
	if err := utils.UnmarshalJSON(data, &destinationAwsDatalakeIAMUser, "", true, true); err == nil {
		u.DestinationAwsDatalakeIAMUser = destinationAwsDatalakeIAMUser
		u.Type = DestinationAwsDatalakeAuthenticationModeTypeDestinationAwsDatalakeIAMUser
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAwsDatalakeAuthenticationMode) MarshalJSON() ([]byte, error) {
	if u.DestinationAwsDatalakeIAMRole != nil {
		return utils.MarshalJSON(u.DestinationAwsDatalakeIAMRole, "", true)
	}

	if u.DestinationAwsDatalakeIAMUser != nil {
		return utils.MarshalJSON(u.DestinationAwsDatalakeIAMUser, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type AwsDatalake string

const (
	AwsDatalakeAwsDatalake AwsDatalake = "aws-datalake"
)

func (e AwsDatalake) ToPointer() *AwsDatalake {
	return &e
}

func (e *AwsDatalake) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aws-datalake":
		*e = AwsDatalake(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AwsDatalake: %v", v)
	}
}

// DestinationAwsDatalakeSchemasCompressionCodecOptional - The compression algorithm used to compress data.
type DestinationAwsDatalakeSchemasCompressionCodecOptional string

const (
	DestinationAwsDatalakeSchemasCompressionCodecOptionalUncompressed DestinationAwsDatalakeSchemasCompressionCodecOptional = "UNCOMPRESSED"
	DestinationAwsDatalakeSchemasCompressionCodecOptionalSnappy       DestinationAwsDatalakeSchemasCompressionCodecOptional = "SNAPPY"
	DestinationAwsDatalakeSchemasCompressionCodecOptionalGzip         DestinationAwsDatalakeSchemasCompressionCodecOptional = "GZIP"
	DestinationAwsDatalakeSchemasCompressionCodecOptionalZstd         DestinationAwsDatalakeSchemasCompressionCodecOptional = "ZSTD"
)

func (e DestinationAwsDatalakeSchemasCompressionCodecOptional) ToPointer() *DestinationAwsDatalakeSchemasCompressionCodecOptional {
	return &e
}

func (e *DestinationAwsDatalakeSchemasCompressionCodecOptional) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCOMPRESSED":
		fallthrough
	case "SNAPPY":
		fallthrough
	case "GZIP":
		fallthrough
	case "ZSTD":
		*e = DestinationAwsDatalakeSchemasCompressionCodecOptional(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeSchemasCompressionCodecOptional: %v", v)
	}
}

type DestinationAwsDatalakeSchemasFormatTypeWildcard string

const (
	DestinationAwsDatalakeSchemasFormatTypeWildcardParquet DestinationAwsDatalakeSchemasFormatTypeWildcard = "Parquet"
)

func (e DestinationAwsDatalakeSchemasFormatTypeWildcard) ToPointer() *DestinationAwsDatalakeSchemasFormatTypeWildcard {
	return &e
}

func (e *DestinationAwsDatalakeSchemasFormatTypeWildcard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Parquet":
		*e = DestinationAwsDatalakeSchemasFormatTypeWildcard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeSchemasFormatTypeWildcard: %v", v)
	}
}

type DestinationAwsDatalakeParquetColumnarStorage struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeSchemasCompressionCodecOptional `default:"SNAPPY" json:"compression_codec"`
	FormatType       *DestinationAwsDatalakeSchemasFormatTypeWildcard       `default:"Parquet" json:"format_type"`
}

func (d DestinationAwsDatalakeParquetColumnarStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAwsDatalakeParquetColumnarStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAwsDatalakeParquetColumnarStorage) GetCompressionCodec() *DestinationAwsDatalakeSchemasCompressionCodecOptional {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *DestinationAwsDatalakeParquetColumnarStorage) GetFormatType() *DestinationAwsDatalakeSchemasFormatTypeWildcard {
	if o == nil {
		return nil
	}
	return o.FormatType
}

// DestinationAwsDatalakeCompressionCodecOptional - The compression algorithm used to compress data.
type DestinationAwsDatalakeCompressionCodecOptional string

const (
	DestinationAwsDatalakeCompressionCodecOptionalUncompressed DestinationAwsDatalakeCompressionCodecOptional = "UNCOMPRESSED"
	DestinationAwsDatalakeCompressionCodecOptionalGzip         DestinationAwsDatalakeCompressionCodecOptional = "GZIP"
)

func (e DestinationAwsDatalakeCompressionCodecOptional) ToPointer() *DestinationAwsDatalakeCompressionCodecOptional {
	return &e
}

func (e *DestinationAwsDatalakeCompressionCodecOptional) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCOMPRESSED":
		fallthrough
	case "GZIP":
		*e = DestinationAwsDatalakeCompressionCodecOptional(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeCompressionCodecOptional: %v", v)
	}
}

type DestinationAwsDatalakeFormatTypeWildcard string

const (
	DestinationAwsDatalakeFormatTypeWildcardJsonl DestinationAwsDatalakeFormatTypeWildcard = "JSONL"
)

func (e DestinationAwsDatalakeFormatTypeWildcard) ToPointer() *DestinationAwsDatalakeFormatTypeWildcard {
	return &e
}

func (e *DestinationAwsDatalakeFormatTypeWildcard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationAwsDatalakeFormatTypeWildcard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeFormatTypeWildcard: %v", v)
	}
}

type DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeCompressionCodecOptional `default:"UNCOMPRESSED" json:"compression_codec"`
	FormatType       *DestinationAwsDatalakeFormatTypeWildcard       `default:"JSONL" json:"format_type"`
}

func (d DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON) GetCompressionCodec() *DestinationAwsDatalakeCompressionCodecOptional {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON) GetFormatType() *DestinationAwsDatalakeFormatTypeWildcard {
	if o == nil {
		return nil
	}
	return o.FormatType
}

type DestinationAwsDatalakeOutputFormatWildcardType string

const (
	DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeJSONLinesNewlineDelimitedJSON DestinationAwsDatalakeOutputFormatWildcardType = "destination-aws-datalake_JSON Lines: Newline-delimited JSON"
	DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeParquetColumnarStorage        DestinationAwsDatalakeOutputFormatWildcardType = "destination-aws-datalake_Parquet: Columnar Storage"
)

// DestinationAwsDatalakeOutputFormatWildcard - Format of the data output.
type DestinationAwsDatalakeOutputFormatWildcard struct {
	DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON *DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON
	DestinationAwsDatalakeParquetColumnarStorage        *DestinationAwsDatalakeParquetColumnarStorage

	Type DestinationAwsDatalakeOutputFormatWildcardType
}

func CreateDestinationAwsDatalakeOutputFormatWildcardDestinationAwsDatalakeJSONLinesNewlineDelimitedJSON(destinationAwsDatalakeJSONLinesNewlineDelimitedJSON DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON) DestinationAwsDatalakeOutputFormatWildcard {
	typ := DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeJSONLinesNewlineDelimitedJSON

	return DestinationAwsDatalakeOutputFormatWildcard{
		DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON: &destinationAwsDatalakeJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func CreateDestinationAwsDatalakeOutputFormatWildcardDestinationAwsDatalakeParquetColumnarStorage(destinationAwsDatalakeParquetColumnarStorage DestinationAwsDatalakeParquetColumnarStorage) DestinationAwsDatalakeOutputFormatWildcard {
	typ := DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeParquetColumnarStorage

	return DestinationAwsDatalakeOutputFormatWildcard{
		DestinationAwsDatalakeParquetColumnarStorage: &destinationAwsDatalakeParquetColumnarStorage,
		Type: typ,
	}
}

func (u *DestinationAwsDatalakeOutputFormatWildcard) UnmarshalJSON(data []byte) error {

	destinationAwsDatalakeJSONLinesNewlineDelimitedJSON := new(DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON)
	if err := utils.UnmarshalJSON(data, &destinationAwsDatalakeJSONLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON = destinationAwsDatalakeJSONLinesNewlineDelimitedJSON
		u.Type = DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeJSONLinesNewlineDelimitedJSON
		return nil
	}

	destinationAwsDatalakeParquetColumnarStorage := new(DestinationAwsDatalakeParquetColumnarStorage)
	if err := utils.UnmarshalJSON(data, &destinationAwsDatalakeParquetColumnarStorage, "", true, true); err == nil {
		u.DestinationAwsDatalakeParquetColumnarStorage = destinationAwsDatalakeParquetColumnarStorage
		u.Type = DestinationAwsDatalakeOutputFormatWildcardTypeDestinationAwsDatalakeParquetColumnarStorage
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAwsDatalakeOutputFormatWildcard) MarshalJSON() ([]byte, error) {
	if u.DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.DestinationAwsDatalakeJSONLinesNewlineDelimitedJSON, "", true)
	}

	if u.DestinationAwsDatalakeParquetColumnarStorage != nil {
		return utils.MarshalJSON(u.DestinationAwsDatalakeParquetColumnarStorage, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationAwsDatalakeChooseHowToPartitionData - Partition data by cursor fields when a cursor field is a date
type DestinationAwsDatalakeChooseHowToPartitionData string

const (
	DestinationAwsDatalakeChooseHowToPartitionDataNoPartitioning DestinationAwsDatalakeChooseHowToPartitionData = "NO PARTITIONING"
	DestinationAwsDatalakeChooseHowToPartitionDataDate           DestinationAwsDatalakeChooseHowToPartitionData = "DATE"
	DestinationAwsDatalakeChooseHowToPartitionDataYear           DestinationAwsDatalakeChooseHowToPartitionData = "YEAR"
	DestinationAwsDatalakeChooseHowToPartitionDataMonth          DestinationAwsDatalakeChooseHowToPartitionData = "MONTH"
	DestinationAwsDatalakeChooseHowToPartitionDataDay            DestinationAwsDatalakeChooseHowToPartitionData = "DAY"
	DestinationAwsDatalakeChooseHowToPartitionDataYearMonth      DestinationAwsDatalakeChooseHowToPartitionData = "YEAR/MONTH"
	DestinationAwsDatalakeChooseHowToPartitionDataYearMonthDay   DestinationAwsDatalakeChooseHowToPartitionData = "YEAR/MONTH/DAY"
)

func (e DestinationAwsDatalakeChooseHowToPartitionData) ToPointer() *DestinationAwsDatalakeChooseHowToPartitionData {
	return &e
}

func (e *DestinationAwsDatalakeChooseHowToPartitionData) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO PARTITIONING":
		fallthrough
	case "DATE":
		fallthrough
	case "YEAR":
		fallthrough
	case "MONTH":
		fallthrough
	case "DAY":
		fallthrough
	case "YEAR/MONTH":
		fallthrough
	case "YEAR/MONTH/DAY":
		*e = DestinationAwsDatalakeChooseHowToPartitionData(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeChooseHowToPartitionData: %v", v)
	}
}

// DestinationAwsDatalakeS3BucketRegion - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type DestinationAwsDatalakeS3BucketRegion string

const (
	DestinationAwsDatalakeS3BucketRegionUnknown      DestinationAwsDatalakeS3BucketRegion = ""
	DestinationAwsDatalakeS3BucketRegionAfSouth1     DestinationAwsDatalakeS3BucketRegion = "af-south-1"
	DestinationAwsDatalakeS3BucketRegionApEast1      DestinationAwsDatalakeS3BucketRegion = "ap-east-1"
	DestinationAwsDatalakeS3BucketRegionApNortheast1 DestinationAwsDatalakeS3BucketRegion = "ap-northeast-1"
	DestinationAwsDatalakeS3BucketRegionApNortheast2 DestinationAwsDatalakeS3BucketRegion = "ap-northeast-2"
	DestinationAwsDatalakeS3BucketRegionApNortheast3 DestinationAwsDatalakeS3BucketRegion = "ap-northeast-3"
	DestinationAwsDatalakeS3BucketRegionApSouth1     DestinationAwsDatalakeS3BucketRegion = "ap-south-1"
	DestinationAwsDatalakeS3BucketRegionApSouth2     DestinationAwsDatalakeS3BucketRegion = "ap-south-2"
	DestinationAwsDatalakeS3BucketRegionApSoutheast1 DestinationAwsDatalakeS3BucketRegion = "ap-southeast-1"
	DestinationAwsDatalakeS3BucketRegionApSoutheast2 DestinationAwsDatalakeS3BucketRegion = "ap-southeast-2"
	DestinationAwsDatalakeS3BucketRegionApSoutheast3 DestinationAwsDatalakeS3BucketRegion = "ap-southeast-3"
	DestinationAwsDatalakeS3BucketRegionApSoutheast4 DestinationAwsDatalakeS3BucketRegion = "ap-southeast-4"
	DestinationAwsDatalakeS3BucketRegionCaCentral1   DestinationAwsDatalakeS3BucketRegion = "ca-central-1"
	DestinationAwsDatalakeS3BucketRegionCaWest1      DestinationAwsDatalakeS3BucketRegion = "ca-west-1"
	DestinationAwsDatalakeS3BucketRegionCnNorth1     DestinationAwsDatalakeS3BucketRegion = "cn-north-1"
	DestinationAwsDatalakeS3BucketRegionCnNorthwest1 DestinationAwsDatalakeS3BucketRegion = "cn-northwest-1"
	DestinationAwsDatalakeS3BucketRegionEuCentral1   DestinationAwsDatalakeS3BucketRegion = "eu-central-1"
	DestinationAwsDatalakeS3BucketRegionEuCentral2   DestinationAwsDatalakeS3BucketRegion = "eu-central-2"
	DestinationAwsDatalakeS3BucketRegionEuNorth1     DestinationAwsDatalakeS3BucketRegion = "eu-north-1"
	DestinationAwsDatalakeS3BucketRegionEuSouth1     DestinationAwsDatalakeS3BucketRegion = "eu-south-1"
	DestinationAwsDatalakeS3BucketRegionEuSouth2     DestinationAwsDatalakeS3BucketRegion = "eu-south-2"
	DestinationAwsDatalakeS3BucketRegionEuWest1      DestinationAwsDatalakeS3BucketRegion = "eu-west-1"
	DestinationAwsDatalakeS3BucketRegionEuWest2      DestinationAwsDatalakeS3BucketRegion = "eu-west-2"
	DestinationAwsDatalakeS3BucketRegionEuWest3      DestinationAwsDatalakeS3BucketRegion = "eu-west-3"
	DestinationAwsDatalakeS3BucketRegionIlCentral1   DestinationAwsDatalakeS3BucketRegion = "il-central-1"
	DestinationAwsDatalakeS3BucketRegionMeCentral1   DestinationAwsDatalakeS3BucketRegion = "me-central-1"
	DestinationAwsDatalakeS3BucketRegionMeSouth1     DestinationAwsDatalakeS3BucketRegion = "me-south-1"
	DestinationAwsDatalakeS3BucketRegionSaEast1      DestinationAwsDatalakeS3BucketRegion = "sa-east-1"
	DestinationAwsDatalakeS3BucketRegionUsEast1      DestinationAwsDatalakeS3BucketRegion = "us-east-1"
	DestinationAwsDatalakeS3BucketRegionUsEast2      DestinationAwsDatalakeS3BucketRegion = "us-east-2"
	DestinationAwsDatalakeS3BucketRegionUsGovEast1   DestinationAwsDatalakeS3BucketRegion = "us-gov-east-1"
	DestinationAwsDatalakeS3BucketRegionUsGovWest1   DestinationAwsDatalakeS3BucketRegion = "us-gov-west-1"
	DestinationAwsDatalakeS3BucketRegionUsWest1      DestinationAwsDatalakeS3BucketRegion = "us-west-1"
	DestinationAwsDatalakeS3BucketRegionUsWest2      DestinationAwsDatalakeS3BucketRegion = "us-west-2"
)

func (e DestinationAwsDatalakeS3BucketRegion) ToPointer() *DestinationAwsDatalakeS3BucketRegion {
	return &e
}

func (e *DestinationAwsDatalakeS3BucketRegion) UnmarshalJSON(data []byte) error {
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
		*e = DestinationAwsDatalakeS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeS3BucketRegion: %v", v)
	}
}

type DestinationAwsDatalake struct {
	// target aws account id
	AwsAccountID *string `json:"aws_account_id,omitempty"`
	// The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
	BucketName string `json:"bucket_name"`
	// S3 prefix
	BucketPrefix *string `json:"bucket_prefix,omitempty"`
	// Choose How to Authenticate to AWS.
	Credentials     DestinationAwsDatalakeAuthenticationMode `json:"credentials"`
	destinationType AwsDatalake                              `const:"aws-datalake" json:"destinationType"`
	// Format of the data output.
	Format *DestinationAwsDatalakeOutputFormatWildcard `json:"format,omitempty"`
	// Cast float/double as decimal(38,18). This can help achieve higher accuracy and represent numbers correctly as received from the source.
	GlueCatalogFloatAsDecimal *bool `default:"false" json:"glue_catalog_float_as_decimal"`
	// Add a default tag key to databases created by this destination
	LakeformationDatabaseDefaultTagKey *string `json:"lakeformation_database_default_tag_key,omitempty"`
	// Add default values for the `Tag Key` to databases created by this destination. Comma separate for multiple values.
	LakeformationDatabaseDefaultTagValues *string `json:"lakeformation_database_default_tag_values,omitempty"`
	// The default database this destination will use to create tables in per stream. Can be changed per connection by customizing the namespace.
	LakeformationDatabaseName string `json:"lakeformation_database_name"`
	// Whether to create tables as LF governed tables.
	LakeformationGovernedTables *bool `default:"false" json:"lakeformation_governed_tables"`
	// Partition data by cursor fields when a cursor field is a date
	Partitioning *DestinationAwsDatalakeChooseHowToPartitionData `default:"NO PARTITIONING" json:"partitioning"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	Region *DestinationAwsDatalakeS3BucketRegion `default:"" json:"region"`
}

func (d DestinationAwsDatalake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAwsDatalake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAwsDatalake) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *DestinationAwsDatalake) GetBucketName() string {
	if o == nil {
		return ""
	}
	return o.BucketName
}

func (o *DestinationAwsDatalake) GetBucketPrefix() *string {
	if o == nil {
		return nil
	}
	return o.BucketPrefix
}

func (o *DestinationAwsDatalake) GetCredentials() DestinationAwsDatalakeAuthenticationMode {
	if o == nil {
		return DestinationAwsDatalakeAuthenticationMode{}
	}
	return o.Credentials
}

func (o *DestinationAwsDatalake) GetDestinationType() AwsDatalake {
	return AwsDatalakeAwsDatalake
}

func (o *DestinationAwsDatalake) GetFormat() *DestinationAwsDatalakeOutputFormatWildcard {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *DestinationAwsDatalake) GetGlueCatalogFloatAsDecimal() *bool {
	if o == nil {
		return nil
	}
	return o.GlueCatalogFloatAsDecimal
}

func (o *DestinationAwsDatalake) GetLakeformationDatabaseDefaultTagKey() *string {
	if o == nil {
		return nil
	}
	return o.LakeformationDatabaseDefaultTagKey
}

func (o *DestinationAwsDatalake) GetLakeformationDatabaseDefaultTagValues() *string {
	if o == nil {
		return nil
	}
	return o.LakeformationDatabaseDefaultTagValues
}

func (o *DestinationAwsDatalake) GetLakeformationDatabaseName() string {
	if o == nil {
		return ""
	}
	return o.LakeformationDatabaseName
}

func (o *DestinationAwsDatalake) GetLakeformationGovernedTables() *bool {
	if o == nil {
		return nil
	}
	return o.LakeformationGovernedTables
}

func (o *DestinationAwsDatalake) GetPartitioning() *DestinationAwsDatalakeChooseHowToPartitionData {
	if o == nil {
		return nil
	}
	return o.Partitioning
}

func (o *DestinationAwsDatalake) GetRegion() *DestinationAwsDatalakeS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.Region
}
