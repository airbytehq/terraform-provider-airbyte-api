// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationS3GlueUpdateSchemasCompressionType string

const (
	DestinationS3GlueUpdateSchemasCompressionTypeGzip DestinationS3GlueUpdateSchemasCompressionType = "GZIP"
)

func (e DestinationS3GlueUpdateSchemasCompressionType) ToPointer() *DestinationS3GlueUpdateSchemasCompressionType {
	return &e
}
func (e *DestinationS3GlueUpdateSchemasCompressionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GZIP":
		*e = DestinationS3GlueUpdateSchemasCompressionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueUpdateSchemasCompressionType: %v", v)
	}
}

type DestinationS3GlueUpdateGZIP struct {
	CompressionType *DestinationS3GlueUpdateSchemasCompressionType `default:"GZIP" json:"compression_type"`
}

func (d DestinationS3GlueUpdateGZIP) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationS3GlueUpdateGZIP) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationS3GlueUpdateGZIP) GetCompressionType() *DestinationS3GlueUpdateSchemasCompressionType {
	if o == nil {
		return nil
	}
	return o.CompressionType
}

type DestinationS3GlueUpdateCompressionType string

const (
	DestinationS3GlueUpdateCompressionTypeNoCompression DestinationS3GlueUpdateCompressionType = "No Compression"
)

func (e DestinationS3GlueUpdateCompressionType) ToPointer() *DestinationS3GlueUpdateCompressionType {
	return &e
}
func (e *DestinationS3GlueUpdateCompressionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No Compression":
		*e = DestinationS3GlueUpdateCompressionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueUpdateCompressionType: %v", v)
	}
}

type DestinationS3GlueUpdateNoCompression struct {
	CompressionType *DestinationS3GlueUpdateCompressionType `default:"No Compression" json:"compression_type"`
}

func (d DestinationS3GlueUpdateNoCompression) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationS3GlueUpdateNoCompression) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationS3GlueUpdateNoCompression) GetCompressionType() *DestinationS3GlueUpdateCompressionType {
	if o == nil {
		return nil
	}
	return o.CompressionType
}

type DestinationS3GlueUpdateCompressionUnionType string

const (
	DestinationS3GlueUpdateCompressionUnionTypeDestinationS3GlueUpdateNoCompression DestinationS3GlueUpdateCompressionUnionType = "destination-s3-glue-update_No Compression"
	DestinationS3GlueUpdateCompressionUnionTypeDestinationS3GlueUpdateGZIP          DestinationS3GlueUpdateCompressionUnionType = "destination-s3-glue-update_GZIP"
)

// DestinationS3GlueUpdateCompression - Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").
type DestinationS3GlueUpdateCompression struct {
	DestinationS3GlueUpdateNoCompression *DestinationS3GlueUpdateNoCompression
	DestinationS3GlueUpdateGZIP          *DestinationS3GlueUpdateGZIP

	Type DestinationS3GlueUpdateCompressionUnionType
}

func CreateDestinationS3GlueUpdateCompressionDestinationS3GlueUpdateNoCompression(destinationS3GlueUpdateNoCompression DestinationS3GlueUpdateNoCompression) DestinationS3GlueUpdateCompression {
	typ := DestinationS3GlueUpdateCompressionUnionTypeDestinationS3GlueUpdateNoCompression

	return DestinationS3GlueUpdateCompression{
		DestinationS3GlueUpdateNoCompression: &destinationS3GlueUpdateNoCompression,
		Type:                                 typ,
	}
}

func CreateDestinationS3GlueUpdateCompressionDestinationS3GlueUpdateGZIP(destinationS3GlueUpdateGZIP DestinationS3GlueUpdateGZIP) DestinationS3GlueUpdateCompression {
	typ := DestinationS3GlueUpdateCompressionUnionTypeDestinationS3GlueUpdateGZIP

	return DestinationS3GlueUpdateCompression{
		DestinationS3GlueUpdateGZIP: &destinationS3GlueUpdateGZIP,
		Type:                        typ,
	}
}

func (u *DestinationS3GlueUpdateCompression) UnmarshalJSON(data []byte) error {

	var destinationS3GlueUpdateNoCompression DestinationS3GlueUpdateNoCompression = DestinationS3GlueUpdateNoCompression{}
	if err := utils.UnmarshalJSON(data, &destinationS3GlueUpdateNoCompression, "", true, true); err == nil {
		u.DestinationS3GlueUpdateNoCompression = &destinationS3GlueUpdateNoCompression
		u.Type = DestinationS3GlueUpdateCompressionUnionTypeDestinationS3GlueUpdateNoCompression
		return nil
	}

	var destinationS3GlueUpdateGZIP DestinationS3GlueUpdateGZIP = DestinationS3GlueUpdateGZIP{}
	if err := utils.UnmarshalJSON(data, &destinationS3GlueUpdateGZIP, "", true, true); err == nil {
		u.DestinationS3GlueUpdateGZIP = &destinationS3GlueUpdateGZIP
		u.Type = DestinationS3GlueUpdateCompressionUnionTypeDestinationS3GlueUpdateGZIP
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationS3GlueUpdateCompression", string(data))
}

func (u DestinationS3GlueUpdateCompression) MarshalJSON() ([]byte, error) {
	if u.DestinationS3GlueUpdateNoCompression != nil {
		return utils.MarshalJSON(u.DestinationS3GlueUpdateNoCompression, "", true)
	}

	if u.DestinationS3GlueUpdateGZIP != nil {
		return utils.MarshalJSON(u.DestinationS3GlueUpdateGZIP, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationS3GlueUpdateCompression: all fields are null")
}

// Flattening - Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.
type Flattening string

const (
	FlatteningNoFlattening        Flattening = "No flattening"
	FlatteningRootLevelFlattening Flattening = "Root level flattening"
)

func (e Flattening) ToPointer() *Flattening {
	return &e
}
func (e *Flattening) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No flattening":
		fallthrough
	case "Root level flattening":
		*e = Flattening(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Flattening: %v", v)
	}
}

type DestinationS3GlueUpdateFormatType string

const (
	DestinationS3GlueUpdateFormatTypeJsonl DestinationS3GlueUpdateFormatType = "JSONL"
)

func (e DestinationS3GlueUpdateFormatType) ToPointer() *DestinationS3GlueUpdateFormatType {
	return &e
}
func (e *DestinationS3GlueUpdateFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationS3GlueUpdateFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueUpdateFormatType: %v", v)
	}
}

type DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON struct {
	// Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").
	Compression *DestinationS3GlueUpdateCompression `json:"compression,omitempty"`
	// Whether the input json data should be normalized (flattened) in the output JSON Lines. Please refer to docs for details.
	Flattening *Flattening                        `default:"Root level flattening" json:"flattening"`
	FormatType *DestinationS3GlueUpdateFormatType `default:"JSONL" json:"format_type"`
}

func (d DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON) GetCompression() *DestinationS3GlueUpdateCompression {
	if o == nil {
		return nil
	}
	return o.Compression
}

func (o *DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON) GetFlattening() *Flattening {
	if o == nil {
		return nil
	}
	return o.Flattening
}

func (o *DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON) GetFormatType() *DestinationS3GlueUpdateFormatType {
	if o == nil {
		return nil
	}
	return o.FormatType
}

type DestinationS3GlueUpdateOutputFormatType string

const (
	DestinationS3GlueUpdateOutputFormatTypeDestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON DestinationS3GlueUpdateOutputFormatType = "destination-s3-glue-update_JSON Lines: Newline-delimited JSON"
)

// DestinationS3GlueUpdateOutputFormat - Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details
type DestinationS3GlueUpdateOutputFormat struct {
	DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON *DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON

	Type DestinationS3GlueUpdateOutputFormatType
}

func CreateDestinationS3GlueUpdateOutputFormatDestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON(destinationS3GlueUpdateJSONLinesNewlineDelimitedJSON DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON) DestinationS3GlueUpdateOutputFormat {
	typ := DestinationS3GlueUpdateOutputFormatTypeDestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON

	return DestinationS3GlueUpdateOutputFormat{
		DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON: &destinationS3GlueUpdateJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *DestinationS3GlueUpdateOutputFormat) UnmarshalJSON(data []byte) error {

	var destinationS3GlueUpdateJSONLinesNewlineDelimitedJSON DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON = DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON{}
	if err := utils.UnmarshalJSON(data, &destinationS3GlueUpdateJSONLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON = &destinationS3GlueUpdateJSONLinesNewlineDelimitedJSON
		u.Type = DestinationS3GlueUpdateOutputFormatTypeDestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationS3GlueUpdateOutputFormat", string(data))
}

func (u DestinationS3GlueUpdateOutputFormat) MarshalJSON() ([]byte, error) {
	if u.DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.DestinationS3GlueUpdateJSONLinesNewlineDelimitedJSON, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationS3GlueUpdateOutputFormat: all fields are null")
}

// SerializationLibrary - The library that your query engine will use for reading and writing data in your lake.
type SerializationLibrary string

const (
	SerializationLibraryOrgOpenxDataJsonserdeJSONSerDe     SerializationLibrary = "org.openx.data.jsonserde.JsonSerDe"
	SerializationLibraryOrgApacheHiveHcatalogDataJSONSerDe SerializationLibrary = "org.apache.hive.hcatalog.data.JsonSerDe"
)

func (e SerializationLibrary) ToPointer() *SerializationLibrary {
	return &e
}
func (e *SerializationLibrary) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "org.openx.data.jsonserde.JsonSerDe":
		fallthrough
	case "org.apache.hive.hcatalog.data.JsonSerDe":
		*e = SerializationLibrary(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SerializationLibrary: %v", v)
	}
}

// DestinationS3GlueUpdateS3BucketRegion - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type DestinationS3GlueUpdateS3BucketRegion string

const (
	DestinationS3GlueUpdateS3BucketRegionUnknown      DestinationS3GlueUpdateS3BucketRegion = ""
	DestinationS3GlueUpdateS3BucketRegionAfSouth1     DestinationS3GlueUpdateS3BucketRegion = "af-south-1"
	DestinationS3GlueUpdateS3BucketRegionApEast1      DestinationS3GlueUpdateS3BucketRegion = "ap-east-1"
	DestinationS3GlueUpdateS3BucketRegionApNortheast1 DestinationS3GlueUpdateS3BucketRegion = "ap-northeast-1"
	DestinationS3GlueUpdateS3BucketRegionApNortheast2 DestinationS3GlueUpdateS3BucketRegion = "ap-northeast-2"
	DestinationS3GlueUpdateS3BucketRegionApNortheast3 DestinationS3GlueUpdateS3BucketRegion = "ap-northeast-3"
	DestinationS3GlueUpdateS3BucketRegionApSouth1     DestinationS3GlueUpdateS3BucketRegion = "ap-south-1"
	DestinationS3GlueUpdateS3BucketRegionApSouth2     DestinationS3GlueUpdateS3BucketRegion = "ap-south-2"
	DestinationS3GlueUpdateS3BucketRegionApSoutheast1 DestinationS3GlueUpdateS3BucketRegion = "ap-southeast-1"
	DestinationS3GlueUpdateS3BucketRegionApSoutheast2 DestinationS3GlueUpdateS3BucketRegion = "ap-southeast-2"
	DestinationS3GlueUpdateS3BucketRegionApSoutheast3 DestinationS3GlueUpdateS3BucketRegion = "ap-southeast-3"
	DestinationS3GlueUpdateS3BucketRegionApSoutheast4 DestinationS3GlueUpdateS3BucketRegion = "ap-southeast-4"
	DestinationS3GlueUpdateS3BucketRegionCaCentral1   DestinationS3GlueUpdateS3BucketRegion = "ca-central-1"
	DestinationS3GlueUpdateS3BucketRegionCaWest1      DestinationS3GlueUpdateS3BucketRegion = "ca-west-1"
	DestinationS3GlueUpdateS3BucketRegionCnNorth1     DestinationS3GlueUpdateS3BucketRegion = "cn-north-1"
	DestinationS3GlueUpdateS3BucketRegionCnNorthwest1 DestinationS3GlueUpdateS3BucketRegion = "cn-northwest-1"
	DestinationS3GlueUpdateS3BucketRegionEuCentral1   DestinationS3GlueUpdateS3BucketRegion = "eu-central-1"
	DestinationS3GlueUpdateS3BucketRegionEuCentral2   DestinationS3GlueUpdateS3BucketRegion = "eu-central-2"
	DestinationS3GlueUpdateS3BucketRegionEuNorth1     DestinationS3GlueUpdateS3BucketRegion = "eu-north-1"
	DestinationS3GlueUpdateS3BucketRegionEuSouth1     DestinationS3GlueUpdateS3BucketRegion = "eu-south-1"
	DestinationS3GlueUpdateS3BucketRegionEuSouth2     DestinationS3GlueUpdateS3BucketRegion = "eu-south-2"
	DestinationS3GlueUpdateS3BucketRegionEuWest1      DestinationS3GlueUpdateS3BucketRegion = "eu-west-1"
	DestinationS3GlueUpdateS3BucketRegionEuWest2      DestinationS3GlueUpdateS3BucketRegion = "eu-west-2"
	DestinationS3GlueUpdateS3BucketRegionEuWest3      DestinationS3GlueUpdateS3BucketRegion = "eu-west-3"
	DestinationS3GlueUpdateS3BucketRegionIlCentral1   DestinationS3GlueUpdateS3BucketRegion = "il-central-1"
	DestinationS3GlueUpdateS3BucketRegionMeCentral1   DestinationS3GlueUpdateS3BucketRegion = "me-central-1"
	DestinationS3GlueUpdateS3BucketRegionMeSouth1     DestinationS3GlueUpdateS3BucketRegion = "me-south-1"
	DestinationS3GlueUpdateS3BucketRegionSaEast1      DestinationS3GlueUpdateS3BucketRegion = "sa-east-1"
	DestinationS3GlueUpdateS3BucketRegionUsEast1      DestinationS3GlueUpdateS3BucketRegion = "us-east-1"
	DestinationS3GlueUpdateS3BucketRegionUsEast2      DestinationS3GlueUpdateS3BucketRegion = "us-east-2"
	DestinationS3GlueUpdateS3BucketRegionUsGovEast1   DestinationS3GlueUpdateS3BucketRegion = "us-gov-east-1"
	DestinationS3GlueUpdateS3BucketRegionUsGovWest1   DestinationS3GlueUpdateS3BucketRegion = "us-gov-west-1"
	DestinationS3GlueUpdateS3BucketRegionUsWest1      DestinationS3GlueUpdateS3BucketRegion = "us-west-1"
	DestinationS3GlueUpdateS3BucketRegionUsWest2      DestinationS3GlueUpdateS3BucketRegion = "us-west-2"
)

func (e DestinationS3GlueUpdateS3BucketRegion) ToPointer() *DestinationS3GlueUpdateS3BucketRegion {
	return &e
}
func (e *DestinationS3GlueUpdateS3BucketRegion) UnmarshalJSON(data []byte) error {
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
		*e = DestinationS3GlueUpdateS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationS3GlueUpdateS3BucketRegion: %v", v)
	}
}

type DestinationS3GlueUpdate struct {
	// The access key ID to access the S3 bucket. Airbyte requires Read and Write permissions to the given bucket. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>.
	AccessKeyID *string `json:"access_key_id,omitempty"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
	// Format of the data output. See <a href="https://docs.airbyte.com/integrations/destinations/s3/#supported-output-schema">here</a> for more details
	Format DestinationS3GlueUpdateOutputFormat `json:"format"`
	// Name of the glue database for creating the tables, leave blank if no integration
	GlueDatabase string `json:"glue_database"`
	// The library that your query engine will use for reading and writing data in your lake.
	GlueSerializationLibrary *SerializationLibrary `default:"org.openx.data.jsonserde.JsonSerDe" json:"glue_serialization_library"`
	// The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
	S3BucketName string `json:"s3_bucket_name"`
	// Directory under the S3 bucket where data will be written. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=to%20format%20the-,bucket%20path,-%3A">here</a>
	S3BucketPath string `json:"s3_bucket_path"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	S3BucketRegion *DestinationS3GlueUpdateS3BucketRegion `default:"" json:"s3_bucket_region"`
	// Your S3 endpoint url. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/s3.html#:~:text=Service%20endpoints-,Amazon%20S3%20endpoints,-When%20you%20use">here</a>
	S3Endpoint *string `default:"" json:"s3_endpoint"`
	// Format string on how data will be organized inside the S3 bucket directory. Read more <a href="https://docs.airbyte.com/integrations/destinations/s3#:~:text=The%20full%20path%20of%20the%20output%20data%20with%20the%20default%20S3%20path%20format">here</a>
	S3PathFormat *string `json:"s3_path_format,omitempty"`
	// The corresponding secret to the access key ID. Read more <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">here</a>
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}

func (d DestinationS3GlueUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationS3GlueUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationS3GlueUpdate) GetAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AccessKeyID
}

func (o *DestinationS3GlueUpdate) GetFileNamePattern() *string {
	if o == nil {
		return nil
	}
	return o.FileNamePattern
}

func (o *DestinationS3GlueUpdate) GetFormat() DestinationS3GlueUpdateOutputFormat {
	if o == nil {
		return DestinationS3GlueUpdateOutputFormat{}
	}
	return o.Format
}

func (o *DestinationS3GlueUpdate) GetGlueDatabase() string {
	if o == nil {
		return ""
	}
	return o.GlueDatabase
}

func (o *DestinationS3GlueUpdate) GetGlueSerializationLibrary() *SerializationLibrary {
	if o == nil {
		return nil
	}
	return o.GlueSerializationLibrary
}

func (o *DestinationS3GlueUpdate) GetS3BucketName() string {
	if o == nil {
		return ""
	}
	return o.S3BucketName
}

func (o *DestinationS3GlueUpdate) GetS3BucketPath() string {
	if o == nil {
		return ""
	}
	return o.S3BucketPath
}

func (o *DestinationS3GlueUpdate) GetS3BucketRegion() *DestinationS3GlueUpdateS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.S3BucketRegion
}

func (o *DestinationS3GlueUpdate) GetS3Endpoint() *string {
	if o == nil {
		return nil
	}
	return o.S3Endpoint
}

func (o *DestinationS3GlueUpdate) GetS3PathFormat() *string {
	if o == nil {
		return nil
	}
	return o.S3PathFormat
}

func (o *DestinationS3GlueUpdate) GetSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.SecretAccessKey
}
