// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationAstraUpdateSchemasEmbeddingEmbeddingMode string

const (
	DestinationAstraUpdateSchemasEmbeddingEmbeddingModeOpenaiCompatible DestinationAstraUpdateSchemasEmbeddingEmbeddingMode = "openai_compatible"
)

func (e DestinationAstraUpdateSchemasEmbeddingEmbeddingMode) ToPointer() *DestinationAstraUpdateSchemasEmbeddingEmbeddingMode {
	return &e
}
func (e *DestinationAstraUpdateSchemasEmbeddingEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai_compatible":
		*e = DestinationAstraUpdateSchemasEmbeddingEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasEmbeddingEmbeddingMode: %v", v)
	}
}

// OpenAICompatible - Use a service that's compatible with the OpenAI API to embed text.
type OpenAICompatible struct {
	APIKey *string `default:"" json:"api_key"`
	// The base URL for your OpenAI-compatible service
	BaseURL string `json:"base_url"`
	// The number of dimensions the embedding model is generating
	Dimensions int64                                                `json:"dimensions"`
	mode       *DestinationAstraUpdateSchemasEmbeddingEmbeddingMode `const:"openai_compatible" json:"mode"`
	// The name of the model to use for embedding
	ModelName *string `default:"text-embedding-ada-002" json:"model_name"`
}

func (o OpenAICompatible) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OpenAICompatible) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OpenAICompatible) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *OpenAICompatible) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *OpenAICompatible) GetDimensions() int64 {
	if o == nil {
		return 0
	}
	return o.Dimensions
}

func (o *OpenAICompatible) GetMode() *DestinationAstraUpdateSchemasEmbeddingEmbeddingMode {
	return DestinationAstraUpdateSchemasEmbeddingEmbeddingModeOpenaiCompatible.ToPointer()
}

func (o *OpenAICompatible) GetModelName() *string {
	if o == nil {
		return nil
	}
	return o.ModelName
}

type DestinationAstraUpdateSchemasEmbeddingMode string

const (
	DestinationAstraUpdateSchemasEmbeddingModeAzureOpenai DestinationAstraUpdateSchemasEmbeddingMode = "azure_openai"
)

func (e DestinationAstraUpdateSchemasEmbeddingMode) ToPointer() *DestinationAstraUpdateSchemasEmbeddingMode {
	return &e
}
func (e *DestinationAstraUpdateSchemasEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure_openai":
		*e = DestinationAstraUpdateSchemasEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasEmbeddingMode: %v", v)
	}
}

// AzureOpenAI - Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type AzureOpenAI struct {
	// The base URL for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	APIBase string `json:"api_base"`
	// The deployment for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	Deployment string                                      `json:"deployment"`
	mode       *DestinationAstraUpdateSchemasEmbeddingMode `const:"azure_openai" json:"mode"`
	// The API key for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	OpenaiKey string `json:"openai_key"`
}

func (a AzureOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AzureOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AzureOpenAI) GetAPIBase() string {
	if o == nil {
		return ""
	}
	return o.APIBase
}

func (o *AzureOpenAI) GetDeployment() string {
	if o == nil {
		return ""
	}
	return o.Deployment
}

func (o *AzureOpenAI) GetMode() *DestinationAstraUpdateSchemasEmbeddingMode {
	return DestinationAstraUpdateSchemasEmbeddingModeAzureOpenai.ToPointer()
}

func (o *AzureOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationAstraUpdateSchemasMode string

const (
	DestinationAstraUpdateSchemasModeFake DestinationAstraUpdateSchemasMode = "fake"
)

func (e DestinationAstraUpdateSchemasMode) ToPointer() *DestinationAstraUpdateSchemasMode {
	return &e
}
func (e *DestinationAstraUpdateSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationAstraUpdateSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasMode: %v", v)
	}
}

// Fake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type Fake struct {
	mode *DestinationAstraUpdateSchemasMode `const:"fake" json:"mode"`
}

func (f Fake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *Fake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Fake) GetMode() *DestinationAstraUpdateSchemasMode {
	return DestinationAstraUpdateSchemasModeFake.ToPointer()
}

type DestinationAstraUpdateMode string

const (
	DestinationAstraUpdateModeCohere DestinationAstraUpdateMode = "cohere"
)

func (e DestinationAstraUpdateMode) ToPointer() *DestinationAstraUpdateMode {
	return &e
}
func (e *DestinationAstraUpdateMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationAstraUpdateMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateMode: %v", v)
	}
}

// Cohere - Use the Cohere API to embed text.
type Cohere struct {
	CohereKey string                      `json:"cohere_key"`
	mode      *DestinationAstraUpdateMode `const:"cohere" json:"mode"`
}

func (c Cohere) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Cohere) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Cohere) GetCohereKey() string {
	if o == nil {
		return ""
	}
	return o.CohereKey
}

func (o *Cohere) GetMode() *DestinationAstraUpdateMode {
	return DestinationAstraUpdateModeCohere.ToPointer()
}

type DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode string

const (
	DestinationAstraUpdateSchemasEmbeddingEmbedding1ModeOpenai DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode = "openai"
)

func (e DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode) ToPointer() *DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode {
	return &e
}
func (e *DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode: %v", v)
	}
}

// OpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type OpenAI struct {
	mode      *DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode `const:"openai" json:"mode"`
	OpenaiKey string                                                `json:"openai_key"`
}

func (o OpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OpenAI) GetMode() *DestinationAstraUpdateSchemasEmbeddingEmbedding1Mode {
	return DestinationAstraUpdateSchemasEmbeddingEmbedding1ModeOpenai.ToPointer()
}

func (o *OpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type EmbeddingType string

const (
	EmbeddingTypeOpenAI           EmbeddingType = "OpenAI"
	EmbeddingTypeCohere           EmbeddingType = "Cohere"
	EmbeddingTypeFake             EmbeddingType = "Fake"
	EmbeddingTypeAzureOpenAI      EmbeddingType = "Azure OpenAI"
	EmbeddingTypeOpenAICompatible EmbeddingType = "OpenAI-compatible"
)

// Embedding configuration
type Embedding struct {
	OpenAI           *OpenAI
	Cohere           *Cohere
	Fake             *Fake
	AzureOpenAI      *AzureOpenAI
	OpenAICompatible *OpenAICompatible

	Type EmbeddingType
}

func CreateEmbeddingOpenAI(openAI OpenAI) Embedding {
	typ := EmbeddingTypeOpenAI

	return Embedding{
		OpenAI: &openAI,
		Type:   typ,
	}
}

func CreateEmbeddingCohere(cohere Cohere) Embedding {
	typ := EmbeddingTypeCohere

	return Embedding{
		Cohere: &cohere,
		Type:   typ,
	}
}

func CreateEmbeddingFake(fake Fake) Embedding {
	typ := EmbeddingTypeFake

	return Embedding{
		Fake: &fake,
		Type: typ,
	}
}

func CreateEmbeddingAzureOpenAI(azureOpenAI AzureOpenAI) Embedding {
	typ := EmbeddingTypeAzureOpenAI

	return Embedding{
		AzureOpenAI: &azureOpenAI,
		Type:        typ,
	}
}

func CreateEmbeddingOpenAICompatible(openAICompatible OpenAICompatible) Embedding {
	typ := EmbeddingTypeOpenAICompatible

	return Embedding{
		OpenAICompatible: &openAICompatible,
		Type:             typ,
	}
}

func (u *Embedding) UnmarshalJSON(data []byte) error {

	var fake Fake = Fake{}
	if err := utils.UnmarshalJSON(data, &fake, "", true, true); err == nil {
		u.Fake = &fake
		u.Type = EmbeddingTypeFake
		return nil
	}

	var openAI OpenAI = OpenAI{}
	if err := utils.UnmarshalJSON(data, &openAI, "", true, true); err == nil {
		u.OpenAI = &openAI
		u.Type = EmbeddingTypeOpenAI
		return nil
	}

	var cohere Cohere = Cohere{}
	if err := utils.UnmarshalJSON(data, &cohere, "", true, true); err == nil {
		u.Cohere = &cohere
		u.Type = EmbeddingTypeCohere
		return nil
	}

	var azureOpenAI AzureOpenAI = AzureOpenAI{}
	if err := utils.UnmarshalJSON(data, &azureOpenAI, "", true, true); err == nil {
		u.AzureOpenAI = &azureOpenAI
		u.Type = EmbeddingTypeAzureOpenAI
		return nil
	}

	var openAICompatible OpenAICompatible = OpenAICompatible{}
	if err := utils.UnmarshalJSON(data, &openAICompatible, "", true, true); err == nil {
		u.OpenAICompatible = &openAICompatible
		u.Type = EmbeddingTypeOpenAICompatible
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Embedding", string(data))
}

func (u Embedding) MarshalJSON() ([]byte, error) {
	if u.OpenAI != nil {
		return utils.MarshalJSON(u.OpenAI, "", true)
	}

	if u.Cohere != nil {
		return utils.MarshalJSON(u.Cohere, "", true)
	}

	if u.Fake != nil {
		return utils.MarshalJSON(u.Fake, "", true)
	}

	if u.AzureOpenAI != nil {
		return utils.MarshalJSON(u.AzureOpenAI, "", true)
	}

	if u.OpenAICompatible != nil {
		return utils.MarshalJSON(u.OpenAICompatible, "", true)
	}

	return nil, errors.New("could not marshal union type Embedding: all fields are null")
}

// Indexing - Astra DB gives developers the APIs, real-time data and ecosystem integrations to put accurate RAG and Gen AI apps with fewer hallucinations in production.
type Indexing struct {
	// The application token authorizes a user to connect to a specific Astra DB database. It is created when the user clicks the Generate Token button on the Overview tab of the Database page in the Astra UI.
	AstraDbAppToken string `json:"astra_db_app_token"`
	// The endpoint specifies which Astra DB database queries are sent to. It can be copied from the Database Details section of the Overview tab of the Database page in the Astra UI.
	AstraDbEndpoint string `json:"astra_db_endpoint"`
	// Keyspaces (or Namespaces) serve as containers for organizing data within a database. You can create a new keyspace uisng the Data Explorer tab in the Astra UI. The keyspace default_keyspace is created for you when you create a Vector Database in Astra DB.
	AstraDbKeyspace string `json:"astra_db_keyspace"`
	// Collections hold data. They are analagous to tables in traditional Cassandra terminology. This tool will create the collection with the provided name automatically if it does not already exist. Alternatively, you can create one thorugh the Data Explorer tab in the Astra UI.
	Collection string `json:"collection"`
}

func (o *Indexing) GetAstraDbAppToken() string {
	if o == nil {
		return ""
	}
	return o.AstraDbAppToken
}

func (o *Indexing) GetAstraDbEndpoint() string {
	if o == nil {
		return ""
	}
	return o.AstraDbEndpoint
}

func (o *Indexing) GetAstraDbKeyspace() string {
	if o == nil {
		return ""
	}
	return o.AstraDbKeyspace
}

func (o *Indexing) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

type FieldNameMappingConfigModel struct {
	// The field name in the source
	FromField string `json:"from_field"`
	// The field name to use in the destination
	ToField string `json:"to_field"`
}

func (o *FieldNameMappingConfigModel) GetFromField() string {
	if o == nil {
		return ""
	}
	return o.FromField
}

func (o *FieldNameMappingConfigModel) GetToField() string {
	if o == nil {
		return ""
	}
	return o.ToField
}

// DestinationAstraUpdateLanguage - Split code in suitable places based on the programming language
type DestinationAstraUpdateLanguage string

const (
	DestinationAstraUpdateLanguageCpp      DestinationAstraUpdateLanguage = "cpp"
	DestinationAstraUpdateLanguageGo       DestinationAstraUpdateLanguage = "go"
	DestinationAstraUpdateLanguageJava     DestinationAstraUpdateLanguage = "java"
	DestinationAstraUpdateLanguageJs       DestinationAstraUpdateLanguage = "js"
	DestinationAstraUpdateLanguagePhp      DestinationAstraUpdateLanguage = "php"
	DestinationAstraUpdateLanguageProto    DestinationAstraUpdateLanguage = "proto"
	DestinationAstraUpdateLanguagePython   DestinationAstraUpdateLanguage = "python"
	DestinationAstraUpdateLanguageRst      DestinationAstraUpdateLanguage = "rst"
	DestinationAstraUpdateLanguageRuby     DestinationAstraUpdateLanguage = "ruby"
	DestinationAstraUpdateLanguageRust     DestinationAstraUpdateLanguage = "rust"
	DestinationAstraUpdateLanguageScala    DestinationAstraUpdateLanguage = "scala"
	DestinationAstraUpdateLanguageSwift    DestinationAstraUpdateLanguage = "swift"
	DestinationAstraUpdateLanguageMarkdown DestinationAstraUpdateLanguage = "markdown"
	DestinationAstraUpdateLanguageLatex    DestinationAstraUpdateLanguage = "latex"
	DestinationAstraUpdateLanguageHTML     DestinationAstraUpdateLanguage = "html"
	DestinationAstraUpdateLanguageSol      DestinationAstraUpdateLanguage = "sol"
)

func (e DestinationAstraUpdateLanguage) ToPointer() *DestinationAstraUpdateLanguage {
	return &e
}
func (e *DestinationAstraUpdateLanguage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cpp":
		fallthrough
	case "go":
		fallthrough
	case "java":
		fallthrough
	case "js":
		fallthrough
	case "php":
		fallthrough
	case "proto":
		fallthrough
	case "python":
		fallthrough
	case "rst":
		fallthrough
	case "ruby":
		fallthrough
	case "rust":
		fallthrough
	case "scala":
		fallthrough
	case "swift":
		fallthrough
	case "markdown":
		fallthrough
	case "latex":
		fallthrough
	case "html":
		fallthrough
	case "sol":
		*e = DestinationAstraUpdateLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateLanguage: %v", v)
	}
}

type DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode string

const (
	DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterModeCode DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode = "code"
)

func (e DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode) ToPointer() *DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode {
	return &e
}
func (e *DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "code":
		*e = DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode: %v", v)
	}
}

// ByProgrammingLanguage - Split the text by suitable delimiters based on the programming language. This is useful for splitting code into chunks.
type ByProgrammingLanguage struct {
	// Split code in suitable places based on the programming language
	Language DestinationAstraUpdateLanguage                                       `json:"language"`
	mode     *DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode `const:"code" json:"mode"`
}

func (b ByProgrammingLanguage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *ByProgrammingLanguage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ByProgrammingLanguage) GetLanguage() DestinationAstraUpdateLanguage {
	if o == nil {
		return DestinationAstraUpdateLanguage("")
	}
	return o.Language
}

func (o *ByProgrammingLanguage) GetMode() *DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterMode {
	return DestinationAstraUpdateSchemasProcessingTextSplitterTextSplitterModeCode.ToPointer()
}

type DestinationAstraUpdateSchemasProcessingTextSplitterMode string

const (
	DestinationAstraUpdateSchemasProcessingTextSplitterModeMarkdown DestinationAstraUpdateSchemasProcessingTextSplitterMode = "markdown"
)

func (e DestinationAstraUpdateSchemasProcessingTextSplitterMode) ToPointer() *DestinationAstraUpdateSchemasProcessingTextSplitterMode {
	return &e
}
func (e *DestinationAstraUpdateSchemasProcessingTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "markdown":
		*e = DestinationAstraUpdateSchemasProcessingTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasProcessingTextSplitterMode: %v", v)
	}
}

// ByMarkdownHeader - Split the text by Markdown headers down to the specified header level. If the chunk size fits multiple sections, they will be combined into a single chunk.
type ByMarkdownHeader struct {
	mode *DestinationAstraUpdateSchemasProcessingTextSplitterMode `const:"markdown" json:"mode"`
	// Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points
	SplitLevel *int64 `default:"1" json:"split_level"`
}

func (b ByMarkdownHeader) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *ByMarkdownHeader) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ByMarkdownHeader) GetMode() *DestinationAstraUpdateSchemasProcessingTextSplitterMode {
	return DestinationAstraUpdateSchemasProcessingTextSplitterModeMarkdown.ToPointer()
}

func (o *ByMarkdownHeader) GetSplitLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.SplitLevel
}

type DestinationAstraUpdateSchemasProcessingMode string

const (
	DestinationAstraUpdateSchemasProcessingModeSeparator DestinationAstraUpdateSchemasProcessingMode = "separator"
)

func (e DestinationAstraUpdateSchemasProcessingMode) ToPointer() *DestinationAstraUpdateSchemasProcessingMode {
	return &e
}
func (e *DestinationAstraUpdateSchemasProcessingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "separator":
		*e = DestinationAstraUpdateSchemasProcessingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAstraUpdateSchemasProcessingMode: %v", v)
	}
}

// BySeparator - Split the text by the list of separators until the chunk size is reached, using the earlier mentioned separators where possible. This is useful for splitting text fields by paragraphs, sentences, words, etc.
type BySeparator struct {
	// Whether to keep the separator in the resulting chunks
	KeepSeparator *bool                                        `default:"false" json:"keep_separator"`
	mode          *DestinationAstraUpdateSchemasProcessingMode `const:"separator" json:"mode"`
	// List of separator strings to split text fields by. The separator itself needs to be wrapped in double quotes, e.g. to split by the dot character, use ".". To split by a newline, use "\n".
	Separators []string `json:"separators,omitempty"`
}

func (b BySeparator) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BySeparator) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *BySeparator) GetKeepSeparator() *bool {
	if o == nil {
		return nil
	}
	return o.KeepSeparator
}

func (o *BySeparator) GetMode() *DestinationAstraUpdateSchemasProcessingMode {
	return DestinationAstraUpdateSchemasProcessingModeSeparator.ToPointer()
}

func (o *BySeparator) GetSeparators() []string {
	if o == nil {
		return nil
	}
	return o.Separators
}

type TextSplitterType string

const (
	TextSplitterTypeBySeparator           TextSplitterType = "By Separator"
	TextSplitterTypeByMarkdownHeader      TextSplitterType = "By Markdown header"
	TextSplitterTypeByProgrammingLanguage TextSplitterType = "By Programming Language"
)

// TextSplitter - Split text fields into chunks based on the specified method.
type TextSplitter struct {
	BySeparator           *BySeparator
	ByMarkdownHeader      *ByMarkdownHeader
	ByProgrammingLanguage *ByProgrammingLanguage

	Type TextSplitterType
}

func CreateTextSplitterBySeparator(bySeparator BySeparator) TextSplitter {
	typ := TextSplitterTypeBySeparator

	return TextSplitter{
		BySeparator: &bySeparator,
		Type:        typ,
	}
}

func CreateTextSplitterByMarkdownHeader(byMarkdownHeader ByMarkdownHeader) TextSplitter {
	typ := TextSplitterTypeByMarkdownHeader

	return TextSplitter{
		ByMarkdownHeader: &byMarkdownHeader,
		Type:             typ,
	}
}

func CreateTextSplitterByProgrammingLanguage(byProgrammingLanguage ByProgrammingLanguage) TextSplitter {
	typ := TextSplitterTypeByProgrammingLanguage

	return TextSplitter{
		ByProgrammingLanguage: &byProgrammingLanguage,
		Type:                  typ,
	}
}

func (u *TextSplitter) UnmarshalJSON(data []byte) error {

	var byMarkdownHeader ByMarkdownHeader = ByMarkdownHeader{}
	if err := utils.UnmarshalJSON(data, &byMarkdownHeader, "", true, true); err == nil {
		u.ByMarkdownHeader = &byMarkdownHeader
		u.Type = TextSplitterTypeByMarkdownHeader
		return nil
	}

	var byProgrammingLanguage ByProgrammingLanguage = ByProgrammingLanguage{}
	if err := utils.UnmarshalJSON(data, &byProgrammingLanguage, "", true, true); err == nil {
		u.ByProgrammingLanguage = &byProgrammingLanguage
		u.Type = TextSplitterTypeByProgrammingLanguage
		return nil
	}

	var bySeparator BySeparator = BySeparator{}
	if err := utils.UnmarshalJSON(data, &bySeparator, "", true, true); err == nil {
		u.BySeparator = &bySeparator
		u.Type = TextSplitterTypeBySeparator
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TextSplitter", string(data))
}

func (u TextSplitter) MarshalJSON() ([]byte, error) {
	if u.BySeparator != nil {
		return utils.MarshalJSON(u.BySeparator, "", true)
	}

	if u.ByMarkdownHeader != nil {
		return utils.MarshalJSON(u.ByMarkdownHeader, "", true)
	}

	if u.ByProgrammingLanguage != nil {
		return utils.MarshalJSON(u.ByProgrammingLanguage, "", true)
	}

	return nil, errors.New("could not marshal union type TextSplitter: all fields are null")
}

type ProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.
	FieldNameMappings []FieldNameMappingConfigModel `json:"field_name_mappings,omitempty"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
	// Split text fields into chunks based on the specified method.
	TextSplitter *TextSplitter `json:"text_splitter,omitempty"`
}

func (p ProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *ProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *ProcessingConfigModel) GetFieldNameMappings() []FieldNameMappingConfigModel {
	if o == nil {
		return nil
	}
	return o.FieldNameMappings
}

func (o *ProcessingConfigModel) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *ProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

func (o *ProcessingConfigModel) GetTextSplitter() *TextSplitter {
	if o == nil {
		return nil
	}
	return o.TextSplitter
}

// DestinationAstraUpdate - The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
// as well as to provide type safety for the configuration passed to the destination.
//
// The configuration model is composed of four parts:
// * Processing configuration
// * Embedding configuration
// * Indexing configuration
// * Advanced configuration
//
// Processing, embedding and advanced configuration are provided by this base class, while the indexing configuration is provided by the destination connector in the sub class.
type DestinationAstraUpdate struct {
	// Embedding configuration
	Embedding Embedding `json:"embedding"`
	// Astra DB gives developers the APIs, real-time data and ecosystem integrations to put accurate RAG and Gen AI apps with fewer hallucinations in production.
	Indexing Indexing `json:"indexing"`
	// Do not store the text that gets embedded along with the vector and the metadata in the destination. If set to true, only the vector and the metadata will be stored - in this case raw text for LLM use cases needs to be retrieved from another source.
	OmitRawText *bool                 `default:"false" json:"omit_raw_text"`
	Processing  ProcessingConfigModel `json:"processing"`
}

func (d DestinationAstraUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAstraUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAstraUpdate) GetEmbedding() Embedding {
	if o == nil {
		return Embedding{}
	}
	return o.Embedding
}

func (o *DestinationAstraUpdate) GetIndexing() Indexing {
	if o == nil {
		return Indexing{}
	}
	return o.Indexing
}

func (o *DestinationAstraUpdate) GetOmitRawText() *bool {
	if o == nil {
		return nil
	}
	return o.OmitRawText
}

func (o *DestinationAstraUpdate) GetProcessing() ProcessingConfigModel {
	if o == nil {
		return ProcessingConfigModel{}
	}
	return o.Processing
}
