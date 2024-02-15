// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationQdrantResourceModel) ToSharedDestinationQdrantCreateRequest() *shared.DestinationQdrantCreateRequest {
	var embedding shared.DestinationQdrantEmbedding
	var destinationQdrantOpenAI *shared.DestinationQdrantOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		openaiKey := r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()
		destinationQdrantOpenAI = &shared.DestinationQdrantOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationQdrantOpenAI != nil {
		embedding = shared.DestinationQdrantEmbedding{
			DestinationQdrantOpenAI: destinationQdrantOpenAI,
		}
	}
	var destinationQdrantCohere *shared.DestinationQdrantCohere
	if r.Configuration.Embedding.Cohere != nil {
		cohereKey := r.Configuration.Embedding.Cohere.CohereKey.ValueString()
		destinationQdrantCohere = &shared.DestinationQdrantCohere{
			CohereKey: cohereKey,
		}
	}
	if destinationQdrantCohere != nil {
		embedding = shared.DestinationQdrantEmbedding{
			DestinationQdrantCohere: destinationQdrantCohere,
		}
	}
	var destinationQdrantFake *shared.DestinationQdrantFake
	if r.Configuration.Embedding.Fake != nil {
		destinationQdrantFake = &shared.DestinationQdrantFake{}
	}
	if destinationQdrantFake != nil {
		embedding = shared.DestinationQdrantEmbedding{
			DestinationQdrantFake: destinationQdrantFake,
		}
	}
	var destinationQdrantAzureOpenAI *shared.DestinationQdrantAzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		apiBase := r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()
		deployment := r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()
		openaiKey1 := r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()
		destinationQdrantAzureOpenAI = &shared.DestinationQdrantAzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if destinationQdrantAzureOpenAI != nil {
		embedding = shared.DestinationQdrantEmbedding{
			DestinationQdrantAzureOpenAI: destinationQdrantAzureOpenAI,
		}
	}
	var destinationQdrantOpenAICompatible *shared.DestinationQdrantOpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		baseURL := r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()
		dimensions := r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()
		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		destinationQdrantOpenAICompatible = &shared.DestinationQdrantOpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions,
			ModelName:  modelName,
		}
	}
	if destinationQdrantOpenAICompatible != nil {
		embedding = shared.DestinationQdrantEmbedding{
			DestinationQdrantOpenAICompatible: destinationQdrantOpenAICompatible,
		}
	}
	var authMethod *shared.DestinationQdrantAuthenticationMethod
	if r.Configuration.Indexing.AuthMethod != nil {
		var destinationQdrantAPIKeyAuth *shared.DestinationQdrantAPIKeyAuth
		if r.Configuration.Indexing.AuthMethod.APIKeyAuth != nil {
			apiKey1 := r.Configuration.Indexing.AuthMethod.APIKeyAuth.APIKey.ValueString()
			destinationQdrantAPIKeyAuth = &shared.DestinationQdrantAPIKeyAuth{
				APIKey: apiKey1,
			}
		}
		if destinationQdrantAPIKeyAuth != nil {
			authMethod = &shared.DestinationQdrantAuthenticationMethod{
				DestinationQdrantAPIKeyAuth: destinationQdrantAPIKeyAuth,
			}
		}
		var destinationQdrantNoAuth *shared.DestinationQdrantNoAuth
		if r.Configuration.Indexing.AuthMethod.NoAuth != nil {
			destinationQdrantNoAuth = &shared.DestinationQdrantNoAuth{}
		}
		if destinationQdrantNoAuth != nil {
			authMethod = &shared.DestinationQdrantAuthenticationMethod{
				DestinationQdrantNoAuth: destinationQdrantNoAuth,
			}
		}
	}
	collection := r.Configuration.Indexing.Collection.ValueString()
	distanceMetric := new(shared.DestinationQdrantDistanceMetric)
	if !r.Configuration.Indexing.DistanceMetric.IsUnknown() && !r.Configuration.Indexing.DistanceMetric.IsNull() {
		*distanceMetric = shared.DestinationQdrantDistanceMetric(r.Configuration.Indexing.DistanceMetric.ValueString())
	} else {
		distanceMetric = nil
	}
	preferGrpc := new(bool)
	if !r.Configuration.Indexing.PreferGrpc.IsUnknown() && !r.Configuration.Indexing.PreferGrpc.IsNull() {
		*preferGrpc = r.Configuration.Indexing.PreferGrpc.ValueBool()
	} else {
		preferGrpc = nil
	}
	textField := new(string)
	if !r.Configuration.Indexing.TextField.IsUnknown() && !r.Configuration.Indexing.TextField.IsNull() {
		*textField = r.Configuration.Indexing.TextField.ValueString()
	} else {
		textField = nil
	}
	url := r.Configuration.Indexing.URL.ValueString()
	indexing := shared.DestinationQdrantIndexing{
		AuthMethod:     authMethod,
		Collection:     collection,
		DistanceMetric: distanceMetric,
		PreferGrpc:     preferGrpc,
		TextField:      textField,
		URL:            url,
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
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var fieldNameMappings []shared.DestinationQdrantFieldNameMappingConfigModel = nil
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		fromField := fieldNameMappingsItem.FromField.ValueString()
		toField := fieldNameMappingsItem.ToField.ValueString()
		fieldNameMappings = append(fieldNameMappings, shared.DestinationQdrantFieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = nil
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = nil
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.DestinationQdrantTextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var destinationQdrantBySeparator *shared.DestinationQdrantBySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = nil
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			destinationQdrantBySeparator = &shared.DestinationQdrantBySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if destinationQdrantBySeparator != nil {
			textSplitter = &shared.DestinationQdrantTextSplitter{
				DestinationQdrantBySeparator: destinationQdrantBySeparator,
			}
		}
		var destinationQdrantByMarkdownHeader *shared.DestinationQdrantByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			destinationQdrantByMarkdownHeader = &shared.DestinationQdrantByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if destinationQdrantByMarkdownHeader != nil {
			textSplitter = &shared.DestinationQdrantTextSplitter{
				DestinationQdrantByMarkdownHeader: destinationQdrantByMarkdownHeader,
			}
		}
		var destinationQdrantByProgrammingLanguage *shared.DestinationQdrantByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationQdrantLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			destinationQdrantByProgrammingLanguage = &shared.DestinationQdrantByProgrammingLanguage{
				Language: language,
			}
		}
		if destinationQdrantByProgrammingLanguage != nil {
			textSplitter = &shared.DestinationQdrantTextSplitter{
				DestinationQdrantByProgrammingLanguage: destinationQdrantByProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationQdrantProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationQdrant{
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
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationQdrantCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationQdrantResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationQdrantResourceModel) ToSharedDestinationQdrantPutRequest() *shared.DestinationQdrantPutRequest {
	var embedding shared.DestinationQdrantUpdateEmbedding
	var destinationQdrantUpdateOpenAI *shared.DestinationQdrantUpdateOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		openaiKey := r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()
		destinationQdrantUpdateOpenAI = &shared.DestinationQdrantUpdateOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationQdrantUpdateOpenAI != nil {
		embedding = shared.DestinationQdrantUpdateEmbedding{
			DestinationQdrantUpdateOpenAI: destinationQdrantUpdateOpenAI,
		}
	}
	var destinationQdrantUpdateCohere *shared.DestinationQdrantUpdateCohere
	if r.Configuration.Embedding.Cohere != nil {
		cohereKey := r.Configuration.Embedding.Cohere.CohereKey.ValueString()
		destinationQdrantUpdateCohere = &shared.DestinationQdrantUpdateCohere{
			CohereKey: cohereKey,
		}
	}
	if destinationQdrantUpdateCohere != nil {
		embedding = shared.DestinationQdrantUpdateEmbedding{
			DestinationQdrantUpdateCohere: destinationQdrantUpdateCohere,
		}
	}
	var destinationQdrantUpdateFake *shared.DestinationQdrantUpdateFake
	if r.Configuration.Embedding.Fake != nil {
		destinationQdrantUpdateFake = &shared.DestinationQdrantUpdateFake{}
	}
	if destinationQdrantUpdateFake != nil {
		embedding = shared.DestinationQdrantUpdateEmbedding{
			DestinationQdrantUpdateFake: destinationQdrantUpdateFake,
		}
	}
	var destinationQdrantUpdateAzureOpenAI *shared.DestinationQdrantUpdateAzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		apiBase := r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()
		deployment := r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()
		openaiKey1 := r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()
		destinationQdrantUpdateAzureOpenAI = &shared.DestinationQdrantUpdateAzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if destinationQdrantUpdateAzureOpenAI != nil {
		embedding = shared.DestinationQdrantUpdateEmbedding{
			DestinationQdrantUpdateAzureOpenAI: destinationQdrantUpdateAzureOpenAI,
		}
	}
	var destinationQdrantUpdateOpenAICompatible *shared.DestinationQdrantUpdateOpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		baseURL := r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()
		dimensions := r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()
		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		destinationQdrantUpdateOpenAICompatible = &shared.DestinationQdrantUpdateOpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions,
			ModelName:  modelName,
		}
	}
	if destinationQdrantUpdateOpenAICompatible != nil {
		embedding = shared.DestinationQdrantUpdateEmbedding{
			DestinationQdrantUpdateOpenAICompatible: destinationQdrantUpdateOpenAICompatible,
		}
	}
	var authMethod *shared.DestinationQdrantUpdateAuthenticationMethod
	if r.Configuration.Indexing.AuthMethod != nil {
		var apiKeyAuth *shared.APIKeyAuth
		if r.Configuration.Indexing.AuthMethod.APIKeyAuth != nil {
			apiKey1 := r.Configuration.Indexing.AuthMethod.APIKeyAuth.APIKey.ValueString()
			apiKeyAuth = &shared.APIKeyAuth{
				APIKey: apiKey1,
			}
		}
		if apiKeyAuth != nil {
			authMethod = &shared.DestinationQdrantUpdateAuthenticationMethod{
				APIKeyAuth: apiKeyAuth,
			}
		}
		var destinationQdrantUpdateNoAuth *shared.DestinationQdrantUpdateNoAuth
		if r.Configuration.Indexing.AuthMethod.NoAuth != nil {
			destinationQdrantUpdateNoAuth = &shared.DestinationQdrantUpdateNoAuth{}
		}
		if destinationQdrantUpdateNoAuth != nil {
			authMethod = &shared.DestinationQdrantUpdateAuthenticationMethod{
				DestinationQdrantUpdateNoAuth: destinationQdrantUpdateNoAuth,
			}
		}
	}
	collection := r.Configuration.Indexing.Collection.ValueString()
	distanceMetric := new(shared.DistanceMetric)
	if !r.Configuration.Indexing.DistanceMetric.IsUnknown() && !r.Configuration.Indexing.DistanceMetric.IsNull() {
		*distanceMetric = shared.DistanceMetric(r.Configuration.Indexing.DistanceMetric.ValueString())
	} else {
		distanceMetric = nil
	}
	preferGrpc := new(bool)
	if !r.Configuration.Indexing.PreferGrpc.IsUnknown() && !r.Configuration.Indexing.PreferGrpc.IsNull() {
		*preferGrpc = r.Configuration.Indexing.PreferGrpc.ValueBool()
	} else {
		preferGrpc = nil
	}
	textField := new(string)
	if !r.Configuration.Indexing.TextField.IsUnknown() && !r.Configuration.Indexing.TextField.IsNull() {
		*textField = r.Configuration.Indexing.TextField.ValueString()
	} else {
		textField = nil
	}
	url := r.Configuration.Indexing.URL.ValueString()
	indexing := shared.DestinationQdrantUpdateIndexing{
		AuthMethod:     authMethod,
		Collection:     collection,
		DistanceMetric: distanceMetric,
		PreferGrpc:     preferGrpc,
		TextField:      textField,
		URL:            url,
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
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var fieldNameMappings []shared.DestinationQdrantUpdateFieldNameMappingConfigModel = nil
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		fromField := fieldNameMappingsItem.FromField.ValueString()
		toField := fieldNameMappingsItem.ToField.ValueString()
		fieldNameMappings = append(fieldNameMappings, shared.DestinationQdrantUpdateFieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = nil
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = nil
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.DestinationQdrantUpdateTextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var destinationQdrantUpdateBySeparator *shared.DestinationQdrantUpdateBySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = nil
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			destinationQdrantUpdateBySeparator = &shared.DestinationQdrantUpdateBySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if destinationQdrantUpdateBySeparator != nil {
			textSplitter = &shared.DestinationQdrantUpdateTextSplitter{
				DestinationQdrantUpdateBySeparator: destinationQdrantUpdateBySeparator,
			}
		}
		var destinationQdrantUpdateByMarkdownHeader *shared.DestinationQdrantUpdateByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			destinationQdrantUpdateByMarkdownHeader = &shared.DestinationQdrantUpdateByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if destinationQdrantUpdateByMarkdownHeader != nil {
			textSplitter = &shared.DestinationQdrantUpdateTextSplitter{
				DestinationQdrantUpdateByMarkdownHeader: destinationQdrantUpdateByMarkdownHeader,
			}
		}
		var destinationQdrantUpdateByProgrammingLanguage *shared.DestinationQdrantUpdateByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationQdrantUpdateLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			destinationQdrantUpdateByProgrammingLanguage = &shared.DestinationQdrantUpdateByProgrammingLanguage{
				Language: language,
			}
		}
		if destinationQdrantUpdateByProgrammingLanguage != nil {
			textSplitter = &shared.DestinationQdrantUpdateTextSplitter{
				DestinationQdrantUpdateByProgrammingLanguage: destinationQdrantUpdateByProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationQdrantUpdateProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationQdrantUpdate{
		Embedding:   embedding,
		Indexing:    indexing,
		OmitRawText: omitRawText,
		Processing:  processing,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationQdrantPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
