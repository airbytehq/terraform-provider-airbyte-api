// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Pinecone string

const (
	PineconePinecone Pinecone = "pinecone"
)

func (e Pinecone) ToPointer() *Pinecone {
	return &e
}

func (e *Pinecone) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinecone":
		*e = Pinecone(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pinecone: %v", v)
	}
}

type DestinationPineconeSchemasEmbeddingMode string

const (
	DestinationPineconeSchemasEmbeddingModeFake DestinationPineconeSchemasEmbeddingMode = "fake"
)

func (e DestinationPineconeSchemasEmbeddingMode) ToPointer() *DestinationPineconeSchemasEmbeddingMode {
	return &e
}

func (e *DestinationPineconeSchemasEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationPineconeSchemasEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeSchemasEmbeddingMode: %v", v)
	}
}

// DestinationPineconeFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationPineconeFake struct {
	mode *DestinationPineconeSchemasEmbeddingMode `const:"fake" json:"mode"`
}

func (d DestinationPineconeFake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeFake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeFake) GetMode() *DestinationPineconeSchemasEmbeddingMode {
	return DestinationPineconeSchemasEmbeddingModeFake.ToPointer()
}

type DestinationPineconeSchemasMode string

const (
	DestinationPineconeSchemasModeCohere DestinationPineconeSchemasMode = "cohere"
)

func (e DestinationPineconeSchemasMode) ToPointer() *DestinationPineconeSchemasMode {
	return &e
}

func (e *DestinationPineconeSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationPineconeSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeSchemasMode: %v", v)
	}
}

// DestinationPineconeCohere - Use the Cohere API to embed text.
type DestinationPineconeCohere struct {
	CohereKey string                          `json:"cohere_key"`
	mode      *DestinationPineconeSchemasMode `const:"cohere" json:"mode"`
}

func (d DestinationPineconeCohere) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeCohere) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeCohere) GetCohereKey() string {
	if o == nil {
		return ""
	}
	return o.CohereKey
}

func (o *DestinationPineconeCohere) GetMode() *DestinationPineconeSchemasMode {
	return DestinationPineconeSchemasModeCohere.ToPointer()
}

type DestinationPineconeMode string

const (
	DestinationPineconeModeOpenai DestinationPineconeMode = "openai"
)

func (e DestinationPineconeMode) ToPointer() *DestinationPineconeMode {
	return &e
}

func (e *DestinationPineconeMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationPineconeMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeMode: %v", v)
	}
}

// DestinationPineconeOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPineconeOpenAI struct {
	mode      *DestinationPineconeMode `const:"openai" json:"mode"`
	OpenaiKey string                   `json:"openai_key"`
}

func (d DestinationPineconeOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeOpenAI) GetMode() *DestinationPineconeMode {
	return DestinationPineconeModeOpenai.ToPointer()
}

func (o *DestinationPineconeOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationPineconeEmbeddingType string

const (
	DestinationPineconeEmbeddingTypeDestinationPineconeOpenAI DestinationPineconeEmbeddingType = "destination-pinecone_OpenAI"
	DestinationPineconeEmbeddingTypeDestinationPineconeCohere DestinationPineconeEmbeddingType = "destination-pinecone_Cohere"
	DestinationPineconeEmbeddingTypeDestinationPineconeFake   DestinationPineconeEmbeddingType = "destination-pinecone_Fake"
)

type DestinationPineconeEmbedding struct {
	DestinationPineconeOpenAI *DestinationPineconeOpenAI
	DestinationPineconeCohere *DestinationPineconeCohere
	DestinationPineconeFake   *DestinationPineconeFake

	Type DestinationPineconeEmbeddingType
}

func CreateDestinationPineconeEmbeddingDestinationPineconeOpenAI(destinationPineconeOpenAI DestinationPineconeOpenAI) DestinationPineconeEmbedding {
	typ := DestinationPineconeEmbeddingTypeDestinationPineconeOpenAI

	return DestinationPineconeEmbedding{
		DestinationPineconeOpenAI: &destinationPineconeOpenAI,
		Type:                      typ,
	}
}

func CreateDestinationPineconeEmbeddingDestinationPineconeCohere(destinationPineconeCohere DestinationPineconeCohere) DestinationPineconeEmbedding {
	typ := DestinationPineconeEmbeddingTypeDestinationPineconeCohere

	return DestinationPineconeEmbedding{
		DestinationPineconeCohere: &destinationPineconeCohere,
		Type:                      typ,
	}
}

func CreateDestinationPineconeEmbeddingDestinationPineconeFake(destinationPineconeFake DestinationPineconeFake) DestinationPineconeEmbedding {
	typ := DestinationPineconeEmbeddingTypeDestinationPineconeFake

	return DestinationPineconeEmbedding{
		DestinationPineconeFake: &destinationPineconeFake,
		Type:                    typ,
	}
}

func (u *DestinationPineconeEmbedding) UnmarshalJSON(data []byte) error {

	destinationPineconeFake := new(DestinationPineconeFake)
	if err := utils.UnmarshalJSON(data, &destinationPineconeFake, "", true, true); err == nil {
		u.DestinationPineconeFake = destinationPineconeFake
		u.Type = DestinationPineconeEmbeddingTypeDestinationPineconeFake
		return nil
	}

	destinationPineconeOpenAI := new(DestinationPineconeOpenAI)
	if err := utils.UnmarshalJSON(data, &destinationPineconeOpenAI, "", true, true); err == nil {
		u.DestinationPineconeOpenAI = destinationPineconeOpenAI
		u.Type = DestinationPineconeEmbeddingTypeDestinationPineconeOpenAI
		return nil
	}

	destinationPineconeCohere := new(DestinationPineconeCohere)
	if err := utils.UnmarshalJSON(data, &destinationPineconeCohere, "", true, true); err == nil {
		u.DestinationPineconeCohere = destinationPineconeCohere
		u.Type = DestinationPineconeEmbeddingTypeDestinationPineconeCohere
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPineconeEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationPineconeOpenAI != nil {
		return utils.MarshalJSON(u.DestinationPineconeOpenAI, "", true)
	}

	if u.DestinationPineconeCohere != nil {
		return utils.MarshalJSON(u.DestinationPineconeCohere, "", true)
	}

	if u.DestinationPineconeFake != nil {
		return utils.MarshalJSON(u.DestinationPineconeFake, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationPineconeIndexing - Pinecone is a popular vector store that can be used to store and retrieve embeddings.
type DestinationPineconeIndexing struct {
	// Pinecone index to use
	Index string `json:"index"`
	// Pinecone environment to use
	PineconeEnvironment string `json:"pinecone_environment"`
	PineconeKey         string `json:"pinecone_key"`
}

func (o *DestinationPineconeIndexing) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *DestinationPineconeIndexing) GetPineconeEnvironment() string {
	if o == nil {
		return ""
	}
	return o.PineconeEnvironment
}

func (o *DestinationPineconeIndexing) GetPineconeKey() string {
	if o == nil {
		return ""
	}
	return o.PineconeKey
}

type DestinationPineconeProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
}

func (d DestinationPineconeProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *DestinationPineconeProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *DestinationPineconeProcessingConfigModel) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *DestinationPineconeProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

type DestinationPinecone struct {
	destinationType Pinecone `const:"pinecone" json:"destinationType"`
	// Embedding configuration
	Embedding DestinationPineconeEmbedding `json:"embedding"`
	// Pinecone is a popular vector store that can be used to store and retrieve embeddings.
	Indexing   DestinationPineconeIndexing              `json:"indexing"`
	Processing DestinationPineconeProcessingConfigModel `json:"processing"`
}

func (d DestinationPinecone) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPinecone) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPinecone) GetDestinationType() Pinecone {
	return PineconePinecone
}

func (o *DestinationPinecone) GetEmbedding() DestinationPineconeEmbedding {
	if o == nil {
		return DestinationPineconeEmbedding{}
	}
	return o.Embedding
}

func (o *DestinationPinecone) GetIndexing() DestinationPineconeIndexing {
	if o == nil {
		return DestinationPineconeIndexing{}
	}
	return o.Indexing
}

func (o *DestinationPinecone) GetProcessing() DestinationPineconeProcessingConfigModel {
	if o == nil {
		return DestinationPineconeProcessingConfigModel{}
	}
	return o.Processing
}
