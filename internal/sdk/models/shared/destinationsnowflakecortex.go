// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SnowflakeCortex string

const (
	SnowflakeCortexSnowflakeCortex SnowflakeCortex = "snowflake-cortex"
)

func (e SnowflakeCortex) ToPointer() *SnowflakeCortex {
	return &e
}
func (e *SnowflakeCortex) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "snowflake-cortex":
		*e = SnowflakeCortex(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SnowflakeCortex: %v", v)
	}
}

type DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode string

const (
	DestinationSnowflakeCortexSchemasEmbeddingEmbedding5ModeOpenaiCompatible DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode = "openai_compatible"
)

func (e DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode) ToPointer() *DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai_compatible":
		*e = DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode: %v", v)
	}
}

// DestinationSnowflakeCortexOpenAICompatible - Use a service that's compatible with the OpenAI API to embed text.
type DestinationSnowflakeCortexOpenAICompatible struct {
	APIKey *string `default:"" json:"api_key"`
	// The base URL for your OpenAI-compatible service
	BaseURL string `json:"base_url"`
	// The number of dimensions the embedding model is generating
	Dimensions int64                                                     `json:"dimensions"`
	mode       *DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode `const:"openai_compatible" json:"mode"`
	// The name of the model to use for embedding
	ModelName *string `default:"text-embedding-ada-002" json:"model_name"`
}

func (d DestinationSnowflakeCortexOpenAICompatible) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexOpenAICompatible) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexOpenAICompatible) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *DestinationSnowflakeCortexOpenAICompatible) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *DestinationSnowflakeCortexOpenAICompatible) GetDimensions() int64 {
	if o == nil {
		return 0
	}
	return o.Dimensions
}

func (o *DestinationSnowflakeCortexOpenAICompatible) GetMode() *DestinationSnowflakeCortexSchemasEmbeddingEmbedding5Mode {
	return DestinationSnowflakeCortexSchemasEmbeddingEmbedding5ModeOpenaiCompatible.ToPointer()
}

func (o *DestinationSnowflakeCortexOpenAICompatible) GetModelName() *string {
	if o == nil {
		return nil
	}
	return o.ModelName
}

type DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode string

const (
	DestinationSnowflakeCortexSchemasEmbeddingEmbeddingModeAzureOpenai DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode = "azure_openai"
)

func (e DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode) ToPointer() *DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure_openai":
		*e = DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode: %v", v)
	}
}

// DestinationSnowflakeCortexAzureOpenAI - Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationSnowflakeCortexAzureOpenAI struct {
	// The base URL for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	APIBase string `json:"api_base"`
	// The deployment for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	Deployment string                                                   `json:"deployment"`
	mode       *DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode `const:"azure_openai" json:"mode"`
	// The API key for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	OpenaiKey string `json:"openai_key"`
}

func (d DestinationSnowflakeCortexAzureOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexAzureOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexAzureOpenAI) GetAPIBase() string {
	if o == nil {
		return ""
	}
	return o.APIBase
}

func (o *DestinationSnowflakeCortexAzureOpenAI) GetDeployment() string {
	if o == nil {
		return ""
	}
	return o.Deployment
}

func (o *DestinationSnowflakeCortexAzureOpenAI) GetMode() *DestinationSnowflakeCortexSchemasEmbeddingEmbeddingMode {
	return DestinationSnowflakeCortexSchemasEmbeddingEmbeddingModeAzureOpenai.ToPointer()
}

func (o *DestinationSnowflakeCortexAzureOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationSnowflakeCortexSchemasEmbeddingMode string

const (
	DestinationSnowflakeCortexSchemasEmbeddingModeFake DestinationSnowflakeCortexSchemasEmbeddingMode = "fake"
)

func (e DestinationSnowflakeCortexSchemasEmbeddingMode) ToPointer() *DestinationSnowflakeCortexSchemasEmbeddingMode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationSnowflakeCortexSchemasEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasEmbeddingMode: %v", v)
	}
}

// DestinationSnowflakeCortexFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationSnowflakeCortexFake struct {
	mode *DestinationSnowflakeCortexSchemasEmbeddingMode `const:"fake" json:"mode"`
}

func (d DestinationSnowflakeCortexFake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexFake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexFake) GetMode() *DestinationSnowflakeCortexSchemasEmbeddingMode {
	return DestinationSnowflakeCortexSchemasEmbeddingModeFake.ToPointer()
}

type DestinationSnowflakeCortexSchemasMode string

const (
	DestinationSnowflakeCortexSchemasModeCohere DestinationSnowflakeCortexSchemasMode = "cohere"
)

func (e DestinationSnowflakeCortexSchemasMode) ToPointer() *DestinationSnowflakeCortexSchemasMode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationSnowflakeCortexSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasMode: %v", v)
	}
}

// DestinationSnowflakeCortexCohere - Use the Cohere API to embed text.
type DestinationSnowflakeCortexCohere struct {
	CohereKey string                                 `json:"cohere_key"`
	mode      *DestinationSnowflakeCortexSchemasMode `const:"cohere" json:"mode"`
}

func (d DestinationSnowflakeCortexCohere) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexCohere) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexCohere) GetCohereKey() string {
	if o == nil {
		return ""
	}
	return o.CohereKey
}

func (o *DestinationSnowflakeCortexCohere) GetMode() *DestinationSnowflakeCortexSchemasMode {
	return DestinationSnowflakeCortexSchemasModeCohere.ToPointer()
}

type DestinationSnowflakeCortexMode string

const (
	DestinationSnowflakeCortexModeOpenai DestinationSnowflakeCortexMode = "openai"
)

func (e DestinationSnowflakeCortexMode) ToPointer() *DestinationSnowflakeCortexMode {
	return &e
}
func (e *DestinationSnowflakeCortexMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationSnowflakeCortexMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexMode: %v", v)
	}
}

// DestinationSnowflakeCortexOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationSnowflakeCortexOpenAI struct {
	mode      *DestinationSnowflakeCortexMode `const:"openai" json:"mode"`
	OpenaiKey string                          `json:"openai_key"`
}

func (d DestinationSnowflakeCortexOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexOpenAI) GetMode() *DestinationSnowflakeCortexMode {
	return DestinationSnowflakeCortexModeOpenai.ToPointer()
}

func (o *DestinationSnowflakeCortexOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationSnowflakeCortexEmbeddingType string

const (
	DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexOpenAI           DestinationSnowflakeCortexEmbeddingType = "destination-snowflake-cortex_OpenAI"
	DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexCohere           DestinationSnowflakeCortexEmbeddingType = "destination-snowflake-cortex_Cohere"
	DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexFake             DestinationSnowflakeCortexEmbeddingType = "destination-snowflake-cortex_Fake"
	DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexAzureOpenAI      DestinationSnowflakeCortexEmbeddingType = "destination-snowflake-cortex_Azure OpenAI"
	DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexOpenAICompatible DestinationSnowflakeCortexEmbeddingType = "destination-snowflake-cortex_OpenAI-compatible"
)

// DestinationSnowflakeCortexEmbedding - Embedding configuration
type DestinationSnowflakeCortexEmbedding struct {
	DestinationSnowflakeCortexOpenAI           *DestinationSnowflakeCortexOpenAI
	DestinationSnowflakeCortexCohere           *DestinationSnowflakeCortexCohere
	DestinationSnowflakeCortexFake             *DestinationSnowflakeCortexFake
	DestinationSnowflakeCortexAzureOpenAI      *DestinationSnowflakeCortexAzureOpenAI
	DestinationSnowflakeCortexOpenAICompatible *DestinationSnowflakeCortexOpenAICompatible

	Type DestinationSnowflakeCortexEmbeddingType
}

func CreateDestinationSnowflakeCortexEmbeddingDestinationSnowflakeCortexOpenAI(destinationSnowflakeCortexOpenAI DestinationSnowflakeCortexOpenAI) DestinationSnowflakeCortexEmbedding {
	typ := DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexOpenAI

	return DestinationSnowflakeCortexEmbedding{
		DestinationSnowflakeCortexOpenAI: &destinationSnowflakeCortexOpenAI,
		Type:                             typ,
	}
}

func CreateDestinationSnowflakeCortexEmbeddingDestinationSnowflakeCortexCohere(destinationSnowflakeCortexCohere DestinationSnowflakeCortexCohere) DestinationSnowflakeCortexEmbedding {
	typ := DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexCohere

	return DestinationSnowflakeCortexEmbedding{
		DestinationSnowflakeCortexCohere: &destinationSnowflakeCortexCohere,
		Type:                             typ,
	}
}

func CreateDestinationSnowflakeCortexEmbeddingDestinationSnowflakeCortexFake(destinationSnowflakeCortexFake DestinationSnowflakeCortexFake) DestinationSnowflakeCortexEmbedding {
	typ := DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexFake

	return DestinationSnowflakeCortexEmbedding{
		DestinationSnowflakeCortexFake: &destinationSnowflakeCortexFake,
		Type:                           typ,
	}
}

func CreateDestinationSnowflakeCortexEmbeddingDestinationSnowflakeCortexAzureOpenAI(destinationSnowflakeCortexAzureOpenAI DestinationSnowflakeCortexAzureOpenAI) DestinationSnowflakeCortexEmbedding {
	typ := DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexAzureOpenAI

	return DestinationSnowflakeCortexEmbedding{
		DestinationSnowflakeCortexAzureOpenAI: &destinationSnowflakeCortexAzureOpenAI,
		Type:                                  typ,
	}
}

func CreateDestinationSnowflakeCortexEmbeddingDestinationSnowflakeCortexOpenAICompatible(destinationSnowflakeCortexOpenAICompatible DestinationSnowflakeCortexOpenAICompatible) DestinationSnowflakeCortexEmbedding {
	typ := DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexOpenAICompatible

	return DestinationSnowflakeCortexEmbedding{
		DestinationSnowflakeCortexOpenAICompatible: &destinationSnowflakeCortexOpenAICompatible,
		Type: typ,
	}
}

func (u *DestinationSnowflakeCortexEmbedding) UnmarshalJSON(data []byte) error {

	var destinationSnowflakeCortexFake DestinationSnowflakeCortexFake = DestinationSnowflakeCortexFake{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexFake, "", true, true); err == nil {
		u.DestinationSnowflakeCortexFake = &destinationSnowflakeCortexFake
		u.Type = DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexFake
		return nil
	}

	var destinationSnowflakeCortexOpenAI DestinationSnowflakeCortexOpenAI = DestinationSnowflakeCortexOpenAI{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexOpenAI, "", true, true); err == nil {
		u.DestinationSnowflakeCortexOpenAI = &destinationSnowflakeCortexOpenAI
		u.Type = DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexOpenAI
		return nil
	}

	var destinationSnowflakeCortexCohere DestinationSnowflakeCortexCohere = DestinationSnowflakeCortexCohere{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexCohere, "", true, true); err == nil {
		u.DestinationSnowflakeCortexCohere = &destinationSnowflakeCortexCohere
		u.Type = DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexCohere
		return nil
	}

	var destinationSnowflakeCortexAzureOpenAI DestinationSnowflakeCortexAzureOpenAI = DestinationSnowflakeCortexAzureOpenAI{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexAzureOpenAI, "", true, true); err == nil {
		u.DestinationSnowflakeCortexAzureOpenAI = &destinationSnowflakeCortexAzureOpenAI
		u.Type = DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexAzureOpenAI
		return nil
	}

	var destinationSnowflakeCortexOpenAICompatible DestinationSnowflakeCortexOpenAICompatible = DestinationSnowflakeCortexOpenAICompatible{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexOpenAICompatible, "", true, true); err == nil {
		u.DestinationSnowflakeCortexOpenAICompatible = &destinationSnowflakeCortexOpenAICompatible
		u.Type = DestinationSnowflakeCortexEmbeddingTypeDestinationSnowflakeCortexOpenAICompatible
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationSnowflakeCortexEmbedding", string(data))
}

func (u DestinationSnowflakeCortexEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationSnowflakeCortexOpenAI != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexOpenAI, "", true)
	}

	if u.DestinationSnowflakeCortexCohere != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexCohere, "", true)
	}

	if u.DestinationSnowflakeCortexFake != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexFake, "", true)
	}

	if u.DestinationSnowflakeCortexAzureOpenAI != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexAzureOpenAI, "", true)
	}

	if u.DestinationSnowflakeCortexOpenAICompatible != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexOpenAICompatible, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationSnowflakeCortexEmbedding: all fields are null")
}

type DestinationSnowflakeCortexCredentials struct {
	// Enter the password you want to use to access the database
	Password string `json:"password"`
}

func (o *DestinationSnowflakeCortexCredentials) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

// DestinationSnowflakeCortexSnowflakeConnection - Snowflake can be used to store vector data and retrieve embeddings.
type DestinationSnowflakeCortexSnowflakeConnection struct {
	Credentials DestinationSnowflakeCortexCredentials `json:"credentials"`
	// Enter the name of the database that you want to sync data into
	Database string `json:"database"`
	// Enter the name of the default schema
	DefaultSchema string `json:"default_schema"`
	// Enter the account name you want to use to access the database. This is usually the identifier before .snowflakecomputing.com
	Host string `json:"host"`
	// Enter the role that you want to use to access Snowflake
	Role string `json:"role"`
	// Enter the name of the user you want to use to access the database
	Username string `json:"username"`
	// Enter the name of the warehouse that you want to use as a compute cluster
	Warehouse string `json:"warehouse"`
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetCredentials() DestinationSnowflakeCortexCredentials {
	if o == nil {
		return DestinationSnowflakeCortexCredentials{}
	}
	return o.Credentials
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetDefaultSchema() string {
	if o == nil {
		return ""
	}
	return o.DefaultSchema
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *DestinationSnowflakeCortexSnowflakeConnection) GetWarehouse() string {
	if o == nil {
		return ""
	}
	return o.Warehouse
}

type DestinationSnowflakeCortexFieldNameMappingConfigModel struct {
	// The field name in the source
	FromField string `json:"from_field"`
	// The field name to use in the destination
	ToField string `json:"to_field"`
}

func (o *DestinationSnowflakeCortexFieldNameMappingConfigModel) GetFromField() string {
	if o == nil {
		return ""
	}
	return o.FromField
}

func (o *DestinationSnowflakeCortexFieldNameMappingConfigModel) GetToField() string {
	if o == nil {
		return ""
	}
	return o.ToField
}

// DestinationSnowflakeCortexLanguage - Split code in suitable places based on the programming language
type DestinationSnowflakeCortexLanguage string

const (
	DestinationSnowflakeCortexLanguageCpp      DestinationSnowflakeCortexLanguage = "cpp"
	DestinationSnowflakeCortexLanguageGo       DestinationSnowflakeCortexLanguage = "go"
	DestinationSnowflakeCortexLanguageJava     DestinationSnowflakeCortexLanguage = "java"
	DestinationSnowflakeCortexLanguageJs       DestinationSnowflakeCortexLanguage = "js"
	DestinationSnowflakeCortexLanguagePhp      DestinationSnowflakeCortexLanguage = "php"
	DestinationSnowflakeCortexLanguageProto    DestinationSnowflakeCortexLanguage = "proto"
	DestinationSnowflakeCortexLanguagePython   DestinationSnowflakeCortexLanguage = "python"
	DestinationSnowflakeCortexLanguageRst      DestinationSnowflakeCortexLanguage = "rst"
	DestinationSnowflakeCortexLanguageRuby     DestinationSnowflakeCortexLanguage = "ruby"
	DestinationSnowflakeCortexLanguageRust     DestinationSnowflakeCortexLanguage = "rust"
	DestinationSnowflakeCortexLanguageScala    DestinationSnowflakeCortexLanguage = "scala"
	DestinationSnowflakeCortexLanguageSwift    DestinationSnowflakeCortexLanguage = "swift"
	DestinationSnowflakeCortexLanguageMarkdown DestinationSnowflakeCortexLanguage = "markdown"
	DestinationSnowflakeCortexLanguageLatex    DestinationSnowflakeCortexLanguage = "latex"
	DestinationSnowflakeCortexLanguageHTML     DestinationSnowflakeCortexLanguage = "html"
	DestinationSnowflakeCortexLanguageSol      DestinationSnowflakeCortexLanguage = "sol"
)

func (e DestinationSnowflakeCortexLanguage) ToPointer() *DestinationSnowflakeCortexLanguage {
	return &e
}
func (e *DestinationSnowflakeCortexLanguage) UnmarshalJSON(data []byte) error {
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
		*e = DestinationSnowflakeCortexLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexLanguage: %v", v)
	}
}

type DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode string

const (
	DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterModeCode DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode = "code"
)

func (e DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode) ToPointer() *DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "code":
		*e = DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode: %v", v)
	}
}

// DestinationSnowflakeCortexByProgrammingLanguage - Split the text by suitable delimiters based on the programming language. This is useful for splitting code into chunks.
type DestinationSnowflakeCortexByProgrammingLanguage struct {
	// Split code in suitable places based on the programming language
	Language DestinationSnowflakeCortexLanguage                                       `json:"language"`
	mode     *DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode `const:"code" json:"mode"`
}

func (d DestinationSnowflakeCortexByProgrammingLanguage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexByProgrammingLanguage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexByProgrammingLanguage) GetLanguage() DestinationSnowflakeCortexLanguage {
	if o == nil {
		return DestinationSnowflakeCortexLanguage("")
	}
	return o.Language
}

func (o *DestinationSnowflakeCortexByProgrammingLanguage) GetMode() *DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterMode {
	return DestinationSnowflakeCortexSchemasProcessingTextSplitterTextSplitterModeCode.ToPointer()
}

type DestinationSnowflakeCortexSchemasProcessingTextSplitterMode string

const (
	DestinationSnowflakeCortexSchemasProcessingTextSplitterModeMarkdown DestinationSnowflakeCortexSchemasProcessingTextSplitterMode = "markdown"
)

func (e DestinationSnowflakeCortexSchemasProcessingTextSplitterMode) ToPointer() *DestinationSnowflakeCortexSchemasProcessingTextSplitterMode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasProcessingTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "markdown":
		*e = DestinationSnowflakeCortexSchemasProcessingTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasProcessingTextSplitterMode: %v", v)
	}
}

// DestinationSnowflakeCortexByMarkdownHeader - Split the text by Markdown headers down to the specified header level. If the chunk size fits multiple sections, they will be combined into a single chunk.
type DestinationSnowflakeCortexByMarkdownHeader struct {
	mode *DestinationSnowflakeCortexSchemasProcessingTextSplitterMode `const:"markdown" json:"mode"`
	// Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points
	SplitLevel *int64 `default:"1" json:"split_level"`
}

func (d DestinationSnowflakeCortexByMarkdownHeader) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexByMarkdownHeader) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexByMarkdownHeader) GetMode() *DestinationSnowflakeCortexSchemasProcessingTextSplitterMode {
	return DestinationSnowflakeCortexSchemasProcessingTextSplitterModeMarkdown.ToPointer()
}

func (o *DestinationSnowflakeCortexByMarkdownHeader) GetSplitLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.SplitLevel
}

type DestinationSnowflakeCortexSchemasProcessingMode string

const (
	DestinationSnowflakeCortexSchemasProcessingModeSeparator DestinationSnowflakeCortexSchemasProcessingMode = "separator"
)

func (e DestinationSnowflakeCortexSchemasProcessingMode) ToPointer() *DestinationSnowflakeCortexSchemasProcessingMode {
	return &e
}
func (e *DestinationSnowflakeCortexSchemasProcessingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "separator":
		*e = DestinationSnowflakeCortexSchemasProcessingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeCortexSchemasProcessingMode: %v", v)
	}
}

// DestinationSnowflakeCortexBySeparator - Split the text by the list of separators until the chunk size is reached, using the earlier mentioned separators where possible. This is useful for splitting text fields by paragraphs, sentences, words, etc.
type DestinationSnowflakeCortexBySeparator struct {
	// Whether to keep the separator in the resulting chunks
	KeepSeparator *bool                                            `default:"false" json:"keep_separator"`
	mode          *DestinationSnowflakeCortexSchemasProcessingMode `const:"separator" json:"mode"`
	// List of separator strings to split text fields by. The separator itself needs to be wrapped in double quotes, e.g. to split by the dot character, use ".". To split by a newline, use "\n".
	Separators []string `json:"separators,omitempty"`
}

func (d DestinationSnowflakeCortexBySeparator) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexBySeparator) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexBySeparator) GetKeepSeparator() *bool {
	if o == nil {
		return nil
	}
	return o.KeepSeparator
}

func (o *DestinationSnowflakeCortexBySeparator) GetMode() *DestinationSnowflakeCortexSchemasProcessingMode {
	return DestinationSnowflakeCortexSchemasProcessingModeSeparator.ToPointer()
}

func (o *DestinationSnowflakeCortexBySeparator) GetSeparators() []string {
	if o == nil {
		return nil
	}
	return o.Separators
}

type DestinationSnowflakeCortexTextSplitterType string

const (
	DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexBySeparator           DestinationSnowflakeCortexTextSplitterType = "destination-snowflake-cortex_By Separator"
	DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexByMarkdownHeader      DestinationSnowflakeCortexTextSplitterType = "destination-snowflake-cortex_By Markdown header"
	DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexByProgrammingLanguage DestinationSnowflakeCortexTextSplitterType = "destination-snowflake-cortex_By Programming Language"
)

// DestinationSnowflakeCortexTextSplitter - Split text fields into chunks based on the specified method.
type DestinationSnowflakeCortexTextSplitter struct {
	DestinationSnowflakeCortexBySeparator           *DestinationSnowflakeCortexBySeparator
	DestinationSnowflakeCortexByMarkdownHeader      *DestinationSnowflakeCortexByMarkdownHeader
	DestinationSnowflakeCortexByProgrammingLanguage *DestinationSnowflakeCortexByProgrammingLanguage

	Type DestinationSnowflakeCortexTextSplitterType
}

func CreateDestinationSnowflakeCortexTextSplitterDestinationSnowflakeCortexBySeparator(destinationSnowflakeCortexBySeparator DestinationSnowflakeCortexBySeparator) DestinationSnowflakeCortexTextSplitter {
	typ := DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexBySeparator

	return DestinationSnowflakeCortexTextSplitter{
		DestinationSnowflakeCortexBySeparator: &destinationSnowflakeCortexBySeparator,
		Type:                                  typ,
	}
}

func CreateDestinationSnowflakeCortexTextSplitterDestinationSnowflakeCortexByMarkdownHeader(destinationSnowflakeCortexByMarkdownHeader DestinationSnowflakeCortexByMarkdownHeader) DestinationSnowflakeCortexTextSplitter {
	typ := DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexByMarkdownHeader

	return DestinationSnowflakeCortexTextSplitter{
		DestinationSnowflakeCortexByMarkdownHeader: &destinationSnowflakeCortexByMarkdownHeader,
		Type: typ,
	}
}

func CreateDestinationSnowflakeCortexTextSplitterDestinationSnowflakeCortexByProgrammingLanguage(destinationSnowflakeCortexByProgrammingLanguage DestinationSnowflakeCortexByProgrammingLanguage) DestinationSnowflakeCortexTextSplitter {
	typ := DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexByProgrammingLanguage

	return DestinationSnowflakeCortexTextSplitter{
		DestinationSnowflakeCortexByProgrammingLanguage: &destinationSnowflakeCortexByProgrammingLanguage,
		Type: typ,
	}
}

func (u *DestinationSnowflakeCortexTextSplitter) UnmarshalJSON(data []byte) error {

	var destinationSnowflakeCortexByMarkdownHeader DestinationSnowflakeCortexByMarkdownHeader = DestinationSnowflakeCortexByMarkdownHeader{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexByMarkdownHeader, "", true, true); err == nil {
		u.DestinationSnowflakeCortexByMarkdownHeader = &destinationSnowflakeCortexByMarkdownHeader
		u.Type = DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexByMarkdownHeader
		return nil
	}

	var destinationSnowflakeCortexByProgrammingLanguage DestinationSnowflakeCortexByProgrammingLanguage = DestinationSnowflakeCortexByProgrammingLanguage{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexByProgrammingLanguage, "", true, true); err == nil {
		u.DestinationSnowflakeCortexByProgrammingLanguage = &destinationSnowflakeCortexByProgrammingLanguage
		u.Type = DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexByProgrammingLanguage
		return nil
	}

	var destinationSnowflakeCortexBySeparator DestinationSnowflakeCortexBySeparator = DestinationSnowflakeCortexBySeparator{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeCortexBySeparator, "", true, true); err == nil {
		u.DestinationSnowflakeCortexBySeparator = &destinationSnowflakeCortexBySeparator
		u.Type = DestinationSnowflakeCortexTextSplitterTypeDestinationSnowflakeCortexBySeparator
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationSnowflakeCortexTextSplitter", string(data))
}

func (u DestinationSnowflakeCortexTextSplitter) MarshalJSON() ([]byte, error) {
	if u.DestinationSnowflakeCortexBySeparator != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexBySeparator, "", true)
	}

	if u.DestinationSnowflakeCortexByMarkdownHeader != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexByMarkdownHeader, "", true)
	}

	if u.DestinationSnowflakeCortexByProgrammingLanguage != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeCortexByProgrammingLanguage, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationSnowflakeCortexTextSplitter: all fields are null")
}

type DestinationSnowflakeCortexProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.
	FieldNameMappings []DestinationSnowflakeCortexFieldNameMappingConfigModel `json:"field_name_mappings,omitempty"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
	// Split text fields into chunks based on the specified method.
	TextSplitter *DestinationSnowflakeCortexTextSplitter `json:"text_splitter,omitempty"`
}

func (d DestinationSnowflakeCortexProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortexProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortexProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *DestinationSnowflakeCortexProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *DestinationSnowflakeCortexProcessingConfigModel) GetFieldNameMappings() []DestinationSnowflakeCortexFieldNameMappingConfigModel {
	if o == nil {
		return nil
	}
	return o.FieldNameMappings
}

func (o *DestinationSnowflakeCortexProcessingConfigModel) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *DestinationSnowflakeCortexProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

func (o *DestinationSnowflakeCortexProcessingConfigModel) GetTextSplitter() *DestinationSnowflakeCortexTextSplitter {
	if o == nil {
		return nil
	}
	return o.TextSplitter
}

// DestinationSnowflakeCortex - The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
// as well as to provide type safety for the configuration passed to the destination.
//
// The configuration model is composed of four parts:
// * Processing configuration
// * Embedding configuration
// * Indexing configuration
// * Advanced configuration
//
// Processing, embedding and advanced configuration are provided by this base class, while the indexing configuration is provided by the destination connector in the sub class.
type DestinationSnowflakeCortex struct {
	destinationType SnowflakeCortex `const:"snowflake-cortex" json:"destinationType"`
	// Embedding configuration
	Embedding DestinationSnowflakeCortexEmbedding `json:"embedding"`
	// Snowflake can be used to store vector data and retrieve embeddings.
	Indexing DestinationSnowflakeCortexSnowflakeConnection `json:"indexing"`
	// Do not store the text that gets embedded along with the vector and the metadata in the destination. If set to true, only the vector and the metadata will be stored - in this case raw text for LLM use cases needs to be retrieved from another source.
	OmitRawText *bool                                           `default:"false" json:"omit_raw_text"`
	Processing  DestinationSnowflakeCortexProcessingConfigModel `json:"processing"`
}

func (d DestinationSnowflakeCortex) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeCortex) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeCortex) GetDestinationType() SnowflakeCortex {
	return SnowflakeCortexSnowflakeCortex
}

func (o *DestinationSnowflakeCortex) GetEmbedding() DestinationSnowflakeCortexEmbedding {
	if o == nil {
		return DestinationSnowflakeCortexEmbedding{}
	}
	return o.Embedding
}

func (o *DestinationSnowflakeCortex) GetIndexing() DestinationSnowflakeCortexSnowflakeConnection {
	if o == nil {
		return DestinationSnowflakeCortexSnowflakeConnection{}
	}
	return o.Indexing
}

func (o *DestinationSnowflakeCortex) GetOmitRawText() *bool {
	if o == nil {
		return nil
	}
	return o.OmitRawText
}

func (o *DestinationSnowflakeCortex) GetProcessing() DestinationSnowflakeCortexProcessingConfigModel {
	if o == nil {
		return DestinationSnowflakeCortexProcessingConfigModel{}
	}
	return o.Processing
}
