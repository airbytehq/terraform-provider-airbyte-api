// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationAstraResourceModel) ToSharedDestinationAstraCreateRequest() *shared.DestinationAstraCreateRequest {
	var embedding shared.DestinationAstraEmbedding
	var destinationAstraOpenAI *shared.DestinationAstraOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		var openaiKey string
		openaiKey = r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()

		destinationAstraOpenAI = &shared.DestinationAstraOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationAstraOpenAI != nil {
		embedding = shared.DestinationAstraEmbedding{
			DestinationAstraOpenAI: destinationAstraOpenAI,
		}
	}
	var destinationAstraCohere *shared.DestinationAstraCohere
	if r.Configuration.Embedding.Cohere != nil {
		var cohereKey string
		cohereKey = r.Configuration.Embedding.Cohere.CohereKey.ValueString()

		destinationAstraCohere = &shared.DestinationAstraCohere{
			CohereKey: cohereKey,
		}
	}
	if destinationAstraCohere != nil {
		embedding = shared.DestinationAstraEmbedding{
			DestinationAstraCohere: destinationAstraCohere,
		}
	}
	var destinationAstraFake *shared.DestinationAstraFake
	if r.Configuration.Embedding.Fake != nil {
		destinationAstraFake = &shared.DestinationAstraFake{}
	}
	if destinationAstraFake != nil {
		embedding = shared.DestinationAstraEmbedding{
			DestinationAstraFake: destinationAstraFake,
		}
	}
	var destinationAstraAzureOpenAI *shared.DestinationAstraAzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		var apiBase string
		apiBase = r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()

		var deployment string
		deployment = r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()

		var openaiKey1 string
		openaiKey1 = r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()

		destinationAstraAzureOpenAI = &shared.DestinationAstraAzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if destinationAstraAzureOpenAI != nil {
		embedding = shared.DestinationAstraEmbedding{
			DestinationAstraAzureOpenAI: destinationAstraAzureOpenAI,
		}
	}
	var destinationAstraOpenAICompatible *shared.DestinationAstraOpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		var baseURL string
		baseURL = r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()

		var dimensions int64
		dimensions = r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()

		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		destinationAstraOpenAICompatible = &shared.DestinationAstraOpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions,
			ModelName:  modelName,
		}
	}
	if destinationAstraOpenAICompatible != nil {
		embedding = shared.DestinationAstraEmbedding{
			DestinationAstraOpenAICompatible: destinationAstraOpenAICompatible,
		}
	}
	var astraDbAppToken string
	astraDbAppToken = r.Configuration.Indexing.AstraDbAppToken.ValueString()

	var astraDbEndpoint string
	astraDbEndpoint = r.Configuration.Indexing.AstraDbEndpoint.ValueString()

	var astraDbKeyspace string
	astraDbKeyspace = r.Configuration.Indexing.AstraDbKeyspace.ValueString()

	var collection string
	collection = r.Configuration.Indexing.Collection.ValueString()

	indexing := shared.DestinationAstraIndexing{
		AstraDbAppToken: astraDbAppToken,
		AstraDbEndpoint: astraDbEndpoint,
		AstraDbKeyspace: astraDbKeyspace,
		Collection:      collection,
	}
	omitRawText := new(bool)
	if !r.Configuration.OmitRawText.IsUnknown() && !r.Configuration.OmitRawText.IsNull() {
		*omitRawText = r.Configuration.OmitRawText.ValueBool()
	} else {
		omitRawText = nil
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	var chunkSize int64
	chunkSize = r.Configuration.Processing.ChunkSize.ValueInt64()

	var fieldNameMappings []shared.DestinationAstraFieldNameMappingConfigModel = []shared.DestinationAstraFieldNameMappingConfigModel{}
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		var fromField string
		fromField = fieldNameMappingsItem.FromField.ValueString()

		var toField string
		toField = fieldNameMappingsItem.ToField.ValueString()

		fieldNameMappings = append(fieldNameMappings, shared.DestinationAstraFieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = []string{}
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = []string{}
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.DestinationAstraTextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var destinationAstraBySeparator *shared.DestinationAstraBySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = []string{}
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			destinationAstraBySeparator = &shared.DestinationAstraBySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if destinationAstraBySeparator != nil {
			textSplitter = &shared.DestinationAstraTextSplitter{
				DestinationAstraBySeparator: destinationAstraBySeparator,
			}
		}
		var destinationAstraByMarkdownHeader *shared.DestinationAstraByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			destinationAstraByMarkdownHeader = &shared.DestinationAstraByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if destinationAstraByMarkdownHeader != nil {
			textSplitter = &shared.DestinationAstraTextSplitter{
				DestinationAstraByMarkdownHeader: destinationAstraByMarkdownHeader,
			}
		}
		var destinationAstraByProgrammingLanguage *shared.DestinationAstraByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationAstraLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			destinationAstraByProgrammingLanguage = &shared.DestinationAstraByProgrammingLanguage{
				Language: language,
			}
		}
		if destinationAstraByProgrammingLanguage != nil {
			textSplitter = &shared.DestinationAstraTextSplitter{
				DestinationAstraByProgrammingLanguage: destinationAstraByProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationAstraProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationAstra{
		Embedding:   embedding,
		Indexing:    indexing,
		OmitRawText: omitRawText,
		Processing:  processing,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.DestinationAstraCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationAstraResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationAstraResourceModel) ToSharedDestinationAstraPutRequest() *shared.DestinationAstraPutRequest {
	var embedding shared.Embedding
	var openAI *shared.OpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		var openaiKey string
		openaiKey = r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()

		openAI = &shared.OpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if openAI != nil {
		embedding = shared.Embedding{
			OpenAI: openAI,
		}
	}
	var cohere *shared.Cohere
	if r.Configuration.Embedding.Cohere != nil {
		var cohereKey string
		cohereKey = r.Configuration.Embedding.Cohere.CohereKey.ValueString()

		cohere = &shared.Cohere{
			CohereKey: cohereKey,
		}
	}
	if cohere != nil {
		embedding = shared.Embedding{
			Cohere: cohere,
		}
	}
	var fake *shared.Fake
	if r.Configuration.Embedding.Fake != nil {
		fake = &shared.Fake{}
	}
	if fake != nil {
		embedding = shared.Embedding{
			Fake: fake,
		}
	}
	var azureOpenAI *shared.AzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		var apiBase string
		apiBase = r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()

		var deployment string
		deployment = r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()

		var openaiKey1 string
		openaiKey1 = r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()

		azureOpenAI = &shared.AzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if azureOpenAI != nil {
		embedding = shared.Embedding{
			AzureOpenAI: azureOpenAI,
		}
	}
	var openAICompatible *shared.OpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		var baseURL string
		baseURL = r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()

		var dimensions int64
		dimensions = r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()

		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		openAICompatible = &shared.OpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions,
			ModelName:  modelName,
		}
	}
	if openAICompatible != nil {
		embedding = shared.Embedding{
			OpenAICompatible: openAICompatible,
		}
	}
	var astraDbAppToken string
	astraDbAppToken = r.Configuration.Indexing.AstraDbAppToken.ValueString()

	var astraDbEndpoint string
	astraDbEndpoint = r.Configuration.Indexing.AstraDbEndpoint.ValueString()

	var astraDbKeyspace string
	astraDbKeyspace = r.Configuration.Indexing.AstraDbKeyspace.ValueString()

	var collection string
	collection = r.Configuration.Indexing.Collection.ValueString()

	indexing := shared.Indexing{
		AstraDbAppToken: astraDbAppToken,
		AstraDbEndpoint: astraDbEndpoint,
		AstraDbKeyspace: astraDbKeyspace,
		Collection:      collection,
	}
	omitRawText := new(bool)
	if !r.Configuration.OmitRawText.IsUnknown() && !r.Configuration.OmitRawText.IsNull() {
		*omitRawText = r.Configuration.OmitRawText.ValueBool()
	} else {
		omitRawText = nil
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	var chunkSize int64
	chunkSize = r.Configuration.Processing.ChunkSize.ValueInt64()

	var fieldNameMappings []shared.FieldNameMappingConfigModel = []shared.FieldNameMappingConfigModel{}
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		var fromField string
		fromField = fieldNameMappingsItem.FromField.ValueString()

		var toField string
		toField = fieldNameMappingsItem.ToField.ValueString()

		fieldNameMappings = append(fieldNameMappings, shared.FieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = []string{}
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = []string{}
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.TextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var bySeparator *shared.BySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = []string{}
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			bySeparator = &shared.BySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if bySeparator != nil {
			textSplitter = &shared.TextSplitter{
				BySeparator: bySeparator,
			}
		}
		var byMarkdownHeader *shared.ByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			byMarkdownHeader = &shared.ByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if byMarkdownHeader != nil {
			textSplitter = &shared.TextSplitter{
				ByMarkdownHeader: byMarkdownHeader,
			}
		}
		var byProgrammingLanguage *shared.ByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationAstraUpdateLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			byProgrammingLanguage = &shared.ByProgrammingLanguage{
				Language: language,
			}
		}
		if byProgrammingLanguage != nil {
			textSplitter = &shared.TextSplitter{
				ByProgrammingLanguage: byProgrammingLanguage,
			}
		}
	}
	processing := shared.ProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationAstraUpdate{
		Embedding:   embedding,
		Indexing:    indexing,
		OmitRawText: omitRawText,
		Processing:  processing,
	}
	var name string
	name = r.Name.ValueString()

	var workspaceID string
	workspaceID = r.WorkspaceID.ValueString()

	out := shared.DestinationAstraPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
