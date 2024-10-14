// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

type SourceS3S3 string

const (
	SourceS3S3S3 SourceS3S3 = "s3"
)

func (e SourceS3S3) ToPointer() *SourceS3S3 {
	return &e
}
func (e *SourceS3S3) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "s3":
		*e = SourceS3S3(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3S3: %v", v)
	}
}

type SourceS3SchemasStreamsFormatFormatFiletype string

const (
	SourceS3SchemasStreamsFormatFormatFiletypeUnstructured SourceS3SchemasStreamsFormatFormatFiletype = "unstructured"
)

func (e SourceS3SchemasStreamsFormatFormatFiletype) ToPointer() *SourceS3SchemasStreamsFormatFormatFiletype {
	return &e
}
func (e *SourceS3SchemasStreamsFormatFormatFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unstructured":
		*e = SourceS3SchemasStreamsFormatFormatFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3SchemasStreamsFormatFormatFiletype: %v", v)
	}
}

type SourceS3Mode string

const (
	SourceS3ModeLocal SourceS3Mode = "local"
)

func (e SourceS3Mode) ToPointer() *SourceS3Mode {
	return &e
}
func (e *SourceS3Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "local":
		*e = SourceS3Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3Mode: %v", v)
	}
}

// SourceS3Local - Process files locally, supporting `fast` and `ocr` modes. This is the default option.
type SourceS3Local struct {
	mode *SourceS3Mode `const:"local" json:"mode"`
}

func (s SourceS3Local) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3Local) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3Local) GetMode() *SourceS3Mode {
	return SourceS3ModeLocal.ToPointer()
}

type SourceS3ProcessingType string

const (
	SourceS3ProcessingTypeSourceS3Local SourceS3ProcessingType = "source-s3_Local"
)

// SourceS3Processing - Processing configuration
type SourceS3Processing struct {
	SourceS3Local *SourceS3Local

	Type SourceS3ProcessingType
}

func CreateSourceS3ProcessingSourceS3Local(sourceS3Local SourceS3Local) SourceS3Processing {
	typ := SourceS3ProcessingTypeSourceS3Local

	return SourceS3Processing{
		SourceS3Local: &sourceS3Local,
		Type:          typ,
	}
}

func (u *SourceS3Processing) UnmarshalJSON(data []byte) error {

	var sourceS3Local SourceS3Local = SourceS3Local{}
	if err := utils.UnmarshalJSON(data, &sourceS3Local, "", true, true); err == nil {
		u.SourceS3Local = &sourceS3Local
		u.Type = SourceS3ProcessingTypeSourceS3Local
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceS3Processing", string(data))
}

func (u SourceS3Processing) MarshalJSON() ([]byte, error) {
	if u.SourceS3Local != nil {
		return utils.MarshalJSON(u.SourceS3Local, "", true)
	}

	return nil, errors.New("could not marshal union type SourceS3Processing: all fields are null")
}

// SourceS3ParsingStrategy - The strategy used to parse documents. `fast` extracts text directly from the document which doesn't work for all files. `ocr_only` is more reliable, but slower. `hi_res` is the most reliable, but requires an API key and a hosted instance of unstructured and can't be used with local mode. See the unstructured.io documentation for more details: https://unstructured-io.github.io/unstructured/core/partition.html#partition-pdf
type SourceS3ParsingStrategy string

const (
	SourceS3ParsingStrategyAuto    SourceS3ParsingStrategy = "auto"
	SourceS3ParsingStrategyFast    SourceS3ParsingStrategy = "fast"
	SourceS3ParsingStrategyOcrOnly SourceS3ParsingStrategy = "ocr_only"
	SourceS3ParsingStrategyHiRes   SourceS3ParsingStrategy = "hi_res"
)

func (e SourceS3ParsingStrategy) ToPointer() *SourceS3ParsingStrategy {
	return &e
}
func (e *SourceS3ParsingStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auto":
		fallthrough
	case "fast":
		fallthrough
	case "ocr_only":
		fallthrough
	case "hi_res":
		*e = SourceS3ParsingStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3ParsingStrategy: %v", v)
	}
}

// SourceS3UnstructuredDocumentFormat - Extract text from document formats (.pdf, .docx, .md, .pptx) and emit as one record per file.
type SourceS3UnstructuredDocumentFormat struct {
	filetype *SourceS3SchemasStreamsFormatFormatFiletype `const:"unstructured" json:"filetype"`
	// Processing configuration
	Processing *SourceS3Processing `json:"processing,omitempty"`
	// If true, skip files that cannot be parsed and pass the error message along as the _ab_source_file_parse_error field. If false, fail the sync.
	SkipUnprocessableFiles *bool `default:"true" json:"skip_unprocessable_files"`
	// The strategy used to parse documents. `fast` extracts text directly from the document which doesn't work for all files. `ocr_only` is more reliable, but slower. `hi_res` is the most reliable, but requires an API key and a hosted instance of unstructured and can't be used with local mode. See the unstructured.io documentation for more details: https://unstructured-io.github.io/unstructured/core/partition.html#partition-pdf
	Strategy *SourceS3ParsingStrategy `default:"auto" json:"strategy"`
}

func (s SourceS3UnstructuredDocumentFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3UnstructuredDocumentFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3UnstructuredDocumentFormat) GetFiletype() *SourceS3SchemasStreamsFormatFormatFiletype {
	return SourceS3SchemasStreamsFormatFormatFiletypeUnstructured.ToPointer()
}

func (o *SourceS3UnstructuredDocumentFormat) GetProcessing() *SourceS3Processing {
	if o == nil {
		return nil
	}
	return o.Processing
}

func (o *SourceS3UnstructuredDocumentFormat) GetSkipUnprocessableFiles() *bool {
	if o == nil {
		return nil
	}
	return o.SkipUnprocessableFiles
}

func (o *SourceS3UnstructuredDocumentFormat) GetStrategy() *SourceS3ParsingStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

type SourceS3SchemasStreamsFormatFiletype string

const (
	SourceS3SchemasStreamsFormatFiletypeParquet SourceS3SchemasStreamsFormatFiletype = "parquet"
)

func (e SourceS3SchemasStreamsFormatFiletype) ToPointer() *SourceS3SchemasStreamsFormatFiletype {
	return &e
}
func (e *SourceS3SchemasStreamsFormatFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "parquet":
		*e = SourceS3SchemasStreamsFormatFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3SchemasStreamsFormatFiletype: %v", v)
	}
}

type SourceS3ParquetFormat struct {
	// Whether to convert decimal fields to floats. There is a loss of precision when converting decimals to floats, so this is not recommended.
	DecimalAsFloat *bool                                 `default:"false" json:"decimal_as_float"`
	filetype       *SourceS3SchemasStreamsFormatFiletype `const:"parquet" json:"filetype"`
}

func (s SourceS3ParquetFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3ParquetFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3ParquetFormat) GetDecimalAsFloat() *bool {
	if o == nil {
		return nil
	}
	return o.DecimalAsFloat
}

func (o *SourceS3ParquetFormat) GetFiletype() *SourceS3SchemasStreamsFormatFiletype {
	return SourceS3SchemasStreamsFormatFiletypeParquet.ToPointer()
}

type SourceS3SchemasStreamsFiletype string

const (
	SourceS3SchemasStreamsFiletypeJsonl SourceS3SchemasStreamsFiletype = "jsonl"
)

func (e SourceS3SchemasStreamsFiletype) ToPointer() *SourceS3SchemasStreamsFiletype {
	return &e
}
func (e *SourceS3SchemasStreamsFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "jsonl":
		*e = SourceS3SchemasStreamsFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3SchemasStreamsFiletype: %v", v)
	}
}

type SourceS3JsonlFormat struct {
	filetype *SourceS3SchemasStreamsFiletype `const:"jsonl" json:"filetype"`
}

func (s SourceS3JsonlFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3JsonlFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3JsonlFormat) GetFiletype() *SourceS3SchemasStreamsFiletype {
	return SourceS3SchemasStreamsFiletypeJsonl.ToPointer()
}

type SourceS3SchemasFiletype string

const (
	SourceS3SchemasFiletypeCsv SourceS3SchemasFiletype = "csv"
)

func (e SourceS3SchemasFiletype) ToPointer() *SourceS3SchemasFiletype {
	return &e
}
func (e *SourceS3SchemasFiletype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		*e = SourceS3SchemasFiletype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3SchemasFiletype: %v", v)
	}
}

type SourceS3SchemasStreamsHeaderDefinitionType string

const (
	SourceS3SchemasStreamsHeaderDefinitionTypeUserProvided SourceS3SchemasStreamsHeaderDefinitionType = "User Provided"
)

func (e SourceS3SchemasStreamsHeaderDefinitionType) ToPointer() *SourceS3SchemasStreamsHeaderDefinitionType {
	return &e
}
func (e *SourceS3SchemasStreamsHeaderDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "User Provided":
		*e = SourceS3SchemasStreamsHeaderDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3SchemasStreamsHeaderDefinitionType: %v", v)
	}
}

type SourceS3UserProvided struct {
	// The column names that will be used while emitting the CSV records
	ColumnNames          []string                                    `json:"column_names"`
	headerDefinitionType *SourceS3SchemasStreamsHeaderDefinitionType `const:"User Provided" json:"header_definition_type"`
}

func (s SourceS3UserProvided) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3UserProvided) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3UserProvided) GetColumnNames() []string {
	if o == nil {
		return []string{}
	}
	return o.ColumnNames
}

func (o *SourceS3UserProvided) GetHeaderDefinitionType() *SourceS3SchemasStreamsHeaderDefinitionType {
	return SourceS3SchemasStreamsHeaderDefinitionTypeUserProvided.ToPointer()
}

type SourceS3SchemasHeaderDefinitionType string

const (
	SourceS3SchemasHeaderDefinitionTypeAutogenerated SourceS3SchemasHeaderDefinitionType = "Autogenerated"
)

func (e SourceS3SchemasHeaderDefinitionType) ToPointer() *SourceS3SchemasHeaderDefinitionType {
	return &e
}
func (e *SourceS3SchemasHeaderDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Autogenerated":
		*e = SourceS3SchemasHeaderDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3SchemasHeaderDefinitionType: %v", v)
	}
}

type SourceS3Autogenerated struct {
	headerDefinitionType *SourceS3SchemasHeaderDefinitionType `const:"Autogenerated" json:"header_definition_type"`
}

func (s SourceS3Autogenerated) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3Autogenerated) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3Autogenerated) GetHeaderDefinitionType() *SourceS3SchemasHeaderDefinitionType {
	return SourceS3SchemasHeaderDefinitionTypeAutogenerated.ToPointer()
}

type SourceS3HeaderDefinitionType string

const (
	SourceS3HeaderDefinitionTypeFromCsv SourceS3HeaderDefinitionType = "From CSV"
)

func (e SourceS3HeaderDefinitionType) ToPointer() *SourceS3HeaderDefinitionType {
	return &e
}
func (e *SourceS3HeaderDefinitionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "From CSV":
		*e = SourceS3HeaderDefinitionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3HeaderDefinitionType: %v", v)
	}
}

type SourceS3FromCSV struct {
	headerDefinitionType *SourceS3HeaderDefinitionType `const:"From CSV" json:"header_definition_type"`
}

func (s SourceS3FromCSV) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3FromCSV) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3FromCSV) GetHeaderDefinitionType() *SourceS3HeaderDefinitionType {
	return SourceS3HeaderDefinitionTypeFromCsv.ToPointer()
}

type SourceS3CSVHeaderDefinitionType string

const (
	SourceS3CSVHeaderDefinitionTypeSourceS3FromCSV       SourceS3CSVHeaderDefinitionType = "source-s3_From CSV"
	SourceS3CSVHeaderDefinitionTypeSourceS3Autogenerated SourceS3CSVHeaderDefinitionType = "source-s3_Autogenerated"
	SourceS3CSVHeaderDefinitionTypeSourceS3UserProvided  SourceS3CSVHeaderDefinitionType = "source-s3_User Provided"
)

// SourceS3CSVHeaderDefinition - How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.
type SourceS3CSVHeaderDefinition struct {
	SourceS3FromCSV       *SourceS3FromCSV
	SourceS3Autogenerated *SourceS3Autogenerated
	SourceS3UserProvided  *SourceS3UserProvided

	Type SourceS3CSVHeaderDefinitionType
}

func CreateSourceS3CSVHeaderDefinitionSourceS3FromCSV(sourceS3FromCSV SourceS3FromCSV) SourceS3CSVHeaderDefinition {
	typ := SourceS3CSVHeaderDefinitionTypeSourceS3FromCSV

	return SourceS3CSVHeaderDefinition{
		SourceS3FromCSV: &sourceS3FromCSV,
		Type:            typ,
	}
}

func CreateSourceS3CSVHeaderDefinitionSourceS3Autogenerated(sourceS3Autogenerated SourceS3Autogenerated) SourceS3CSVHeaderDefinition {
	typ := SourceS3CSVHeaderDefinitionTypeSourceS3Autogenerated

	return SourceS3CSVHeaderDefinition{
		SourceS3Autogenerated: &sourceS3Autogenerated,
		Type:                  typ,
	}
}

func CreateSourceS3CSVHeaderDefinitionSourceS3UserProvided(sourceS3UserProvided SourceS3UserProvided) SourceS3CSVHeaderDefinition {
	typ := SourceS3CSVHeaderDefinitionTypeSourceS3UserProvided

	return SourceS3CSVHeaderDefinition{
		SourceS3UserProvided: &sourceS3UserProvided,
		Type:                 typ,
	}
}

func (u *SourceS3CSVHeaderDefinition) UnmarshalJSON(data []byte) error {

	var sourceS3FromCSV SourceS3FromCSV = SourceS3FromCSV{}
	if err := utils.UnmarshalJSON(data, &sourceS3FromCSV, "", true, true); err == nil {
		u.SourceS3FromCSV = &sourceS3FromCSV
		u.Type = SourceS3CSVHeaderDefinitionTypeSourceS3FromCSV
		return nil
	}

	var sourceS3Autogenerated SourceS3Autogenerated = SourceS3Autogenerated{}
	if err := utils.UnmarshalJSON(data, &sourceS3Autogenerated, "", true, true); err == nil {
		u.SourceS3Autogenerated = &sourceS3Autogenerated
		u.Type = SourceS3CSVHeaderDefinitionTypeSourceS3Autogenerated
		return nil
	}

	var sourceS3UserProvided SourceS3UserProvided = SourceS3UserProvided{}
	if err := utils.UnmarshalJSON(data, &sourceS3UserProvided, "", true, true); err == nil {
		u.SourceS3UserProvided = &sourceS3UserProvided
		u.Type = SourceS3CSVHeaderDefinitionTypeSourceS3UserProvided
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceS3CSVHeaderDefinition", string(data))
}

func (u SourceS3CSVHeaderDefinition) MarshalJSON() ([]byte, error) {
	if u.SourceS3FromCSV != nil {
		return utils.MarshalJSON(u.SourceS3FromCSV, "", true)
	}

	if u.SourceS3Autogenerated != nil {
		return utils.MarshalJSON(u.SourceS3Autogenerated, "", true)
	}

	if u.SourceS3UserProvided != nil {
		return utils.MarshalJSON(u.SourceS3UserProvided, "", true)
	}

	return nil, errors.New("could not marshal union type SourceS3CSVHeaderDefinition: all fields are null")
}

type SourceS3CSVFormat struct {
	// The character delimiting individual cells in the CSV data. This may only be a 1-character string. For tab-delimited data enter '\t'.
	Delimiter *string `default:"," json:"delimiter"`
	// Whether two quotes in a quoted CSV value denote a single quote in the data.
	DoubleQuote *bool `default:"true" json:"double_quote"`
	// The character encoding of the CSV data. Leave blank to default to <strong>UTF8</strong>. See <a href="https://docs.python.org/3/library/codecs.html#standard-encodings" target="_blank">list of python encodings</a> for allowable options.
	Encoding *string `default:"utf8" json:"encoding"`
	// The character used for escaping special characters. To disallow escaping, leave this field blank.
	EscapeChar *string `json:"escape_char,omitempty"`
	// A set of case-sensitive strings that should be interpreted as false values.
	FalseValues []string                 `json:"false_values,omitempty"`
	filetype    *SourceS3SchemasFiletype `const:"csv" json:"filetype"`
	// How headers will be defined. `User Provided` assumes the CSV does not have a header row and uses the headers provided and `Autogenerated` assumes the CSV does not have a header row and the CDK will generate headers using for `f{i}` where `i` is the index starting from 0. Else, the default behavior is to use the header from the CSV file. If a user wants to autogenerate or provide column names for a CSV having headers, they can skip rows.
	HeaderDefinition *SourceS3CSVHeaderDefinition `json:"header_definition,omitempty"`
	// Whether to ignore errors that occur when the number of fields in the CSV does not match the number of columns in the schema.
	IgnoreErrorsOnFieldsMismatch *bool `default:"false" json:"ignore_errors_on_fields_mismatch"`
	// A set of case-sensitive strings that should be interpreted as null values. For example, if the value 'NA' should be interpreted as null, enter 'NA' in this field.
	NullValues []string `json:"null_values,omitempty"`
	// The character used for quoting CSV values. To disallow quoting, make this field blank.
	QuoteChar *string `default:"\"" json:"quote_char"`
	// The number of rows to skip after the header row.
	SkipRowsAfterHeader *int64 `default:"0" json:"skip_rows_after_header"`
	// The number of rows to skip before the header row. For example, if the header row is on the 3rd row, enter 2 in this field.
	SkipRowsBeforeHeader *int64 `default:"0" json:"skip_rows_before_header"`
	// Whether strings can be interpreted as null values. If true, strings that match the null_values set will be interpreted as null. If false, strings that match the null_values set will be interpreted as the string itself.
	StringsCanBeNull *bool `default:"true" json:"strings_can_be_null"`
	// A set of case-sensitive strings that should be interpreted as true values.
	TrueValues []string `json:"true_values,omitempty"`
}

func (s SourceS3CSVFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3CSVFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3CSVFormat) GetDelimiter() *string {
	if o == nil {
		return nil
	}
	return o.Delimiter
}

func (o *SourceS3CSVFormat) GetDoubleQuote() *bool {
	if o == nil {
		return nil
	}
	return o.DoubleQuote
}

func (o *SourceS3CSVFormat) GetEncoding() *string {
	if o == nil {
		return nil
	}
	return o.Encoding
}

func (o *SourceS3CSVFormat) GetEscapeChar() *string {
	if o == nil {
		return nil
	}
	return o.EscapeChar
}

func (o *SourceS3CSVFormat) GetFalseValues() []string {
	if o == nil {
		return nil
	}
	return o.FalseValues
}

func (o *SourceS3CSVFormat) GetFiletype() *SourceS3SchemasFiletype {
	return SourceS3SchemasFiletypeCsv.ToPointer()
}

func (o *SourceS3CSVFormat) GetHeaderDefinition() *SourceS3CSVHeaderDefinition {
	if o == nil {
		return nil
	}
	return o.HeaderDefinition
}

func (o *SourceS3CSVFormat) GetIgnoreErrorsOnFieldsMismatch() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreErrorsOnFieldsMismatch
}

func (o *SourceS3CSVFormat) GetNullValues() []string {
	if o == nil {
		return nil
	}
	return o.NullValues
}

func (o *SourceS3CSVFormat) GetQuoteChar() *string {
	if o == nil {
		return nil
	}
	return o.QuoteChar
}

func (o *SourceS3CSVFormat) GetSkipRowsAfterHeader() *int64 {
	if o == nil {
		return nil
	}
	return o.SkipRowsAfterHeader
}

func (o *SourceS3CSVFormat) GetSkipRowsBeforeHeader() *int64 {
	if o == nil {
		return nil
	}
	return o.SkipRowsBeforeHeader
}

func (o *SourceS3CSVFormat) GetStringsCanBeNull() *bool {
	if o == nil {
		return nil
	}
	return o.StringsCanBeNull
}

func (o *SourceS3CSVFormat) GetTrueValues() []string {
	if o == nil {
		return nil
	}
	return o.TrueValues
}

type SourceS3Filetype string

const (
	SourceS3FiletypeAvro SourceS3Filetype = "avro"
)

func (e SourceS3Filetype) ToPointer() *SourceS3Filetype {
	return &e
}
func (e *SourceS3Filetype) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "avro":
		*e = SourceS3Filetype(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3Filetype: %v", v)
	}
}

type SourceS3AvroFormat struct {
	// Whether to convert double fields to strings. This is recommended if you have decimal numbers with a high degree of precision because there can be a loss precision when handling floating point numbers.
	DoubleAsString *bool             `default:"false" json:"double_as_string"`
	filetype       *SourceS3Filetype `const:"avro" json:"filetype"`
}

func (s SourceS3AvroFormat) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3AvroFormat) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3AvroFormat) GetDoubleAsString() *bool {
	if o == nil {
		return nil
	}
	return o.DoubleAsString
}

func (o *SourceS3AvroFormat) GetFiletype() *SourceS3Filetype {
	return SourceS3FiletypeAvro.ToPointer()
}

type SourceS3FormatType string

const (
	SourceS3FormatTypeSourceS3AvroFormat                 SourceS3FormatType = "source-s3_Avro Format"
	SourceS3FormatTypeSourceS3CSVFormat                  SourceS3FormatType = "source-s3_CSV Format"
	SourceS3FormatTypeSourceS3JsonlFormat                SourceS3FormatType = "source-s3_Jsonl Format"
	SourceS3FormatTypeSourceS3ParquetFormat              SourceS3FormatType = "source-s3_Parquet Format"
	SourceS3FormatTypeSourceS3UnstructuredDocumentFormat SourceS3FormatType = "source-s3_Unstructured Document Format"
)

// SourceS3Format - The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.
type SourceS3Format struct {
	SourceS3AvroFormat                 *SourceS3AvroFormat
	SourceS3CSVFormat                  *SourceS3CSVFormat
	SourceS3JsonlFormat                *SourceS3JsonlFormat
	SourceS3ParquetFormat              *SourceS3ParquetFormat
	SourceS3UnstructuredDocumentFormat *SourceS3UnstructuredDocumentFormat

	Type SourceS3FormatType
}

func CreateSourceS3FormatSourceS3AvroFormat(sourceS3AvroFormat SourceS3AvroFormat) SourceS3Format {
	typ := SourceS3FormatTypeSourceS3AvroFormat

	return SourceS3Format{
		SourceS3AvroFormat: &sourceS3AvroFormat,
		Type:               typ,
	}
}

func CreateSourceS3FormatSourceS3CSVFormat(sourceS3CSVFormat SourceS3CSVFormat) SourceS3Format {
	typ := SourceS3FormatTypeSourceS3CSVFormat

	return SourceS3Format{
		SourceS3CSVFormat: &sourceS3CSVFormat,
		Type:              typ,
	}
}

func CreateSourceS3FormatSourceS3JsonlFormat(sourceS3JsonlFormat SourceS3JsonlFormat) SourceS3Format {
	typ := SourceS3FormatTypeSourceS3JsonlFormat

	return SourceS3Format{
		SourceS3JsonlFormat: &sourceS3JsonlFormat,
		Type:                typ,
	}
}

func CreateSourceS3FormatSourceS3ParquetFormat(sourceS3ParquetFormat SourceS3ParquetFormat) SourceS3Format {
	typ := SourceS3FormatTypeSourceS3ParquetFormat

	return SourceS3Format{
		SourceS3ParquetFormat: &sourceS3ParquetFormat,
		Type:                  typ,
	}
}

func CreateSourceS3FormatSourceS3UnstructuredDocumentFormat(sourceS3UnstructuredDocumentFormat SourceS3UnstructuredDocumentFormat) SourceS3Format {
	typ := SourceS3FormatTypeSourceS3UnstructuredDocumentFormat

	return SourceS3Format{
		SourceS3UnstructuredDocumentFormat: &sourceS3UnstructuredDocumentFormat,
		Type:                               typ,
	}
}

func (u *SourceS3Format) UnmarshalJSON(data []byte) error {

	var sourceS3JsonlFormat SourceS3JsonlFormat = SourceS3JsonlFormat{}
	if err := utils.UnmarshalJSON(data, &sourceS3JsonlFormat, "", true, true); err == nil {
		u.SourceS3JsonlFormat = &sourceS3JsonlFormat
		u.Type = SourceS3FormatTypeSourceS3JsonlFormat
		return nil
	}

	var sourceS3AvroFormat SourceS3AvroFormat = SourceS3AvroFormat{}
	if err := utils.UnmarshalJSON(data, &sourceS3AvroFormat, "", true, true); err == nil {
		u.SourceS3AvroFormat = &sourceS3AvroFormat
		u.Type = SourceS3FormatTypeSourceS3AvroFormat
		return nil
	}

	var sourceS3ParquetFormat SourceS3ParquetFormat = SourceS3ParquetFormat{}
	if err := utils.UnmarshalJSON(data, &sourceS3ParquetFormat, "", true, true); err == nil {
		u.SourceS3ParquetFormat = &sourceS3ParquetFormat
		u.Type = SourceS3FormatTypeSourceS3ParquetFormat
		return nil
	}

	var sourceS3UnstructuredDocumentFormat SourceS3UnstructuredDocumentFormat = SourceS3UnstructuredDocumentFormat{}
	if err := utils.UnmarshalJSON(data, &sourceS3UnstructuredDocumentFormat, "", true, true); err == nil {
		u.SourceS3UnstructuredDocumentFormat = &sourceS3UnstructuredDocumentFormat
		u.Type = SourceS3FormatTypeSourceS3UnstructuredDocumentFormat
		return nil
	}

	var sourceS3CSVFormat SourceS3CSVFormat = SourceS3CSVFormat{}
	if err := utils.UnmarshalJSON(data, &sourceS3CSVFormat, "", true, true); err == nil {
		u.SourceS3CSVFormat = &sourceS3CSVFormat
		u.Type = SourceS3FormatTypeSourceS3CSVFormat
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceS3Format", string(data))
}

func (u SourceS3Format) MarshalJSON() ([]byte, error) {
	if u.SourceS3AvroFormat != nil {
		return utils.MarshalJSON(u.SourceS3AvroFormat, "", true)
	}

	if u.SourceS3CSVFormat != nil {
		return utils.MarshalJSON(u.SourceS3CSVFormat, "", true)
	}

	if u.SourceS3JsonlFormat != nil {
		return utils.MarshalJSON(u.SourceS3JsonlFormat, "", true)
	}

	if u.SourceS3ParquetFormat != nil {
		return utils.MarshalJSON(u.SourceS3ParquetFormat, "", true)
	}

	if u.SourceS3UnstructuredDocumentFormat != nil {
		return utils.MarshalJSON(u.SourceS3UnstructuredDocumentFormat, "", true)
	}

	return nil, errors.New("could not marshal union type SourceS3Format: all fields are null")
}

// SourceS3ValidationPolicy - The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema.
type SourceS3ValidationPolicy string

const (
	SourceS3ValidationPolicyEmitRecord      SourceS3ValidationPolicy = "Emit Record"
	SourceS3ValidationPolicySkipRecord      SourceS3ValidationPolicy = "Skip Record"
	SourceS3ValidationPolicyWaitForDiscover SourceS3ValidationPolicy = "Wait for Discover"
)

func (e SourceS3ValidationPolicy) ToPointer() *SourceS3ValidationPolicy {
	return &e
}
func (e *SourceS3ValidationPolicy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Emit Record":
		fallthrough
	case "Skip Record":
		fallthrough
	case "Wait for Discover":
		*e = SourceS3ValidationPolicy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3ValidationPolicy: %v", v)
	}
}

type SourceS3FileBasedStreamConfig struct {
	// When the state history of the file store is full, syncs will only read files that were last modified in the provided day range.
	DaysToSyncIfHistoryIsFull *int64 `default:"3" json:"days_to_sync_if_history_is_full"`
	// The configuration options that are used to alter how to read incoming files that deviate from the standard formatting.
	Format SourceS3Format `json:"format"`
	// The pattern used to specify which files should be selected from the file system. For more information on glob pattern matching look <a href="https://en.wikipedia.org/wiki/Glob_(programming)">here</a>.
	Globs []string `json:"globs,omitempty"`
	// The schema that will be used to validate records extracted from the file. This will override the stream schema that is auto-detected from incoming files.
	InputSchema *string `json:"input_schema,omitempty"`
	// The name of the stream.
	Name string `json:"name"`
	// The number of resent files which will be used to discover the schema for this stream.
	RecentNFilesToReadForSchemaDiscovery *int64 `json:"recent_n_files_to_read_for_schema_discovery,omitempty"`
	// When enabled, syncs will not validate or structure records against the stream's schema.
	Schemaless *bool `default:"false" json:"schemaless"`
	// The name of the validation policy that dictates sync behavior when a record does not adhere to the stream schema.
	ValidationPolicy *SourceS3ValidationPolicy `default:"Emit Record" json:"validation_policy"`
}

func (s SourceS3FileBasedStreamConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3FileBasedStreamConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3FileBasedStreamConfig) GetDaysToSyncIfHistoryIsFull() *int64 {
	if o == nil {
		return nil
	}
	return o.DaysToSyncIfHistoryIsFull
}

func (o *SourceS3FileBasedStreamConfig) GetFormat() SourceS3Format {
	if o == nil {
		return SourceS3Format{}
	}
	return o.Format
}

func (o *SourceS3FileBasedStreamConfig) GetGlobs() []string {
	if o == nil {
		return nil
	}
	return o.Globs
}

func (o *SourceS3FileBasedStreamConfig) GetInputSchema() *string {
	if o == nil {
		return nil
	}
	return o.InputSchema
}

func (o *SourceS3FileBasedStreamConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceS3FileBasedStreamConfig) GetRecentNFilesToReadForSchemaDiscovery() *int64 {
	if o == nil {
		return nil
	}
	return o.RecentNFilesToReadForSchemaDiscovery
}

func (o *SourceS3FileBasedStreamConfig) GetSchemaless() *bool {
	if o == nil {
		return nil
	}
	return o.Schemaless
}

func (o *SourceS3FileBasedStreamConfig) GetValidationPolicy() *SourceS3ValidationPolicy {
	if o == nil {
		return nil
	}
	return o.ValidationPolicy
}

// SourceS3 - NOTE: When this Spec is changed, legacy_config_transformer.py must also be modified to uptake the changes
// because it is responsible for converting legacy S3 v3 configs into v4 configs using the File-Based CDK.
type SourceS3 struct {
	// In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`
	// In order to access private Buckets stored on AWS S3, this connector requires credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsSecretAccessKey *string `json:"aws_secret_access_key,omitempty"`
	// Name of the S3 bucket where the file(s) exist.
	Bucket string `json:"bucket"`
	// Endpoint to an S3 compatible service. Leave empty to use AWS.
	Endpoint *string `default:"" json:"endpoint"`
	// AWS region where the S3 bucket is located. If not provided, the region will be determined automatically.
	RegionName *string `json:"region_name,omitempty"`
	// Specifies the Amazon Resource Name (ARN) of an IAM role that you want to use to perform operations requested using this profile. Set the External ID to the Airbyte workspace ID, which can be found in the URL of this page.
	RoleArn    *string    `json:"role_arn,omitempty"`
	sourceType SourceS3S3 `const:"s3" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00.000000Z. Any file modified before this date will not be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Each instance of this configuration defines a <a href="https://docs.airbyte.com/cloud/core-concepts#stream">stream</a>. Use this to define which files belong in the stream, their format, and how they should be parsed and validated. When sending data to warehouse destination such as Snowflake or BigQuery, each stream is a separate table.
	Streams []SourceS3FileBasedStreamConfig `json:"streams"`
}

func (s SourceS3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *SourceS3) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *SourceS3) GetBucket() string {
	if o == nil {
		return ""
	}
	return o.Bucket
}

func (o *SourceS3) GetEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.Endpoint
}

func (o *SourceS3) GetRegionName() *string {
	if o == nil {
		return nil
	}
	return o.RegionName
}

func (o *SourceS3) GetRoleArn() *string {
	if o == nil {
		return nil
	}
	return o.RoleArn
}

func (o *SourceS3) GetSourceType() SourceS3S3 {
	return SourceS3S3S3
}

func (o *SourceS3) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceS3) GetStreams() []SourceS3FileBasedStreamConfig {
	if o == nil {
		return []SourceS3FileBasedStreamConfig{}
	}
	return o.Streams
}
