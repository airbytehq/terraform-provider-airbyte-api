// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGcsResourceModel) ToSharedSourceGcsCreateRequest() *shared.SourceGcsCreateRequest {
	bucket := r.Configuration.Bucket.ValueString()
	serviceAccount := r.Configuration.ServiceAccount.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceGcsFileBasedStreamConfig = []shared.SourceGcsFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceGcsFormat
		var sourceGcsAvroFormat *shared.SourceGcsAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceGcsAvroFormat = &shared.SourceGcsAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceGcsAvroFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsAvroFormat: sourceGcsAvroFormat,
			}
		}
		var sourceGcsCSVFormat *shared.SourceGcsCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = []string{}
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceGcsCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceGcsFromCSV *shared.SourceGcsFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceGcsFromCSV = &shared.SourceGcsFromCSV{}
				}
				if sourceGcsFromCSV != nil {
					headerDefinition = &shared.SourceGcsCSVHeaderDefinition{
						SourceGcsFromCSV: sourceGcsFromCSV,
					}
				}
				var sourceGcsAutogenerated *shared.SourceGcsAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceGcsAutogenerated = &shared.SourceGcsAutogenerated{}
				}
				if sourceGcsAutogenerated != nil {
					headerDefinition = &shared.SourceGcsCSVHeaderDefinition{
						SourceGcsAutogenerated: sourceGcsAutogenerated,
					}
				}
				var sourceGcsUserProvided *shared.SourceGcsUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceGcsUserProvided = &shared.SourceGcsUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceGcsUserProvided != nil {
					headerDefinition = &shared.SourceGcsCSVHeaderDefinition{
						SourceGcsUserProvided: sourceGcsUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
			}
			inferenceType := new(shared.SourceGcsInferenceType)
			if !streamsItem.Format.CSVFormat.InferenceType.IsUnknown() && !streamsItem.Format.CSVFormat.InferenceType.IsNull() {
				*inferenceType = shared.SourceGcsInferenceType(streamsItem.Format.CSVFormat.InferenceType.ValueString())
			} else {
				inferenceType = nil
			}
			var nullValues []string = []string{}
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = []string{}
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceGcsCSVFormat = &shared.SourceGcsCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				InferenceType:                inferenceType,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceGcsCSVFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsCSVFormat: sourceGcsCSVFormat,
			}
		}
		var sourceGcsJsonlFormat *shared.SourceGcsJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceGcsJsonlFormat = &shared.SourceGcsJsonlFormat{}
		}
		if sourceGcsJsonlFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsJsonlFormat: sourceGcsJsonlFormat,
			}
		}
		var sourceGcsParquetFormat *shared.SourceGcsParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceGcsParquetFormat = &shared.SourceGcsParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceGcsParquetFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsParquetFormat: sourceGcsParquetFormat,
			}
		}
		var sourceGcsUnstructuredDocumentFormat *shared.SourceGcsUnstructuredDocumentFormat
		if streamsItem.Format.UnstructuredDocumentFormat != nil {
			var processing *shared.SourceGcsProcessing
			if streamsItem.Format.UnstructuredDocumentFormat.Processing != nil {
				var sourceGcsLocal *shared.SourceGcsLocal
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.Local != nil {
					sourceGcsLocal = &shared.SourceGcsLocal{}
				}
				if sourceGcsLocal != nil {
					processing = &shared.SourceGcsProcessing{
						SourceGcsLocal: sourceGcsLocal,
					}
				}
				var sourceGcsViaAPI *shared.SourceGcsViaAPI
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI != nil {
					apiKey := new(string)
					if !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIKey.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIKey.IsNull() {
						*apiKey = streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIKey.ValueString()
					} else {
						apiKey = nil
					}
					apiURL := new(string)
					if !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIURL.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIURL.IsNull() {
						*apiURL = streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIURL.ValueString()
					} else {
						apiURL = nil
					}
					var parameters []shared.SourceGcsAPIParameterConfigModel = []shared.SourceGcsAPIParameterConfigModel{}
					for _, parametersItem := range streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.Parameters {
						name := parametersItem.Name.ValueString()
						value := parametersItem.Value.ValueString()
						parameters = append(parameters, shared.SourceGcsAPIParameterConfigModel{
							Name:  name,
							Value: value,
						})
					}
					sourceGcsViaAPI = &shared.SourceGcsViaAPI{
						APIKey:     apiKey,
						APIURL:     apiURL,
						Parameters: parameters,
					}
				}
				if sourceGcsViaAPI != nil {
					processing = &shared.SourceGcsProcessing{
						SourceGcsViaAPI: sourceGcsViaAPI,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceGcsParsingStrategy)
			if !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsNull() {
				*strategy = shared.SourceGcsParsingStrategy(streamsItem.Format.UnstructuredDocumentFormat.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceGcsUnstructuredDocumentFormat = &shared.SourceGcsUnstructuredDocumentFormat{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceGcsUnstructuredDocumentFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsUnstructuredDocumentFormat: sourceGcsUnstructuredDocumentFormat,
			}
		}
		var sourceGcsExcelFormat *shared.SourceGcsExcelFormat
		if streamsItem.Format.ExcelFormat != nil {
			sourceGcsExcelFormat = &shared.SourceGcsExcelFormat{}
		}
		if sourceGcsExcelFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsExcelFormat: sourceGcsExcelFormat,
			}
		}
		var globs []string = []string{}
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		legacyPrefix := new(string)
		if !streamsItem.LegacyPrefix.IsUnknown() && !streamsItem.LegacyPrefix.IsNull() {
			*legacyPrefix = streamsItem.LegacyPrefix.ValueString()
		} else {
			legacyPrefix = nil
		}
		name1 := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		recentNFilesToReadForSchemaDiscovery := new(int64)
		if !streamsItem.RecentNFilesToReadForSchemaDiscovery.IsUnknown() && !streamsItem.RecentNFilesToReadForSchemaDiscovery.IsNull() {
			*recentNFilesToReadForSchemaDiscovery = streamsItem.RecentNFilesToReadForSchemaDiscovery.ValueInt64()
		} else {
			recentNFilesToReadForSchemaDiscovery = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceGcsValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceGcsValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceGcsFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull:            daysToSyncIfHistoryIsFull,
			Format:                               format,
			Globs:                                globs,
			InputSchema:                          inputSchema,
			LegacyPrefix:                         legacyPrefix,
			Name:                                 name1,
			PrimaryKey:                           primaryKey,
			RecentNFilesToReadForSchemaDiscovery: recentNFilesToReadForSchemaDiscovery,
			Schemaless:                           schemaless,
			ValidationPolicy:                     validationPolicy,
		})
	}
	configuration := shared.SourceGcs{
		Bucket:         bucket,
		ServiceAccount: serviceAccount,
		StartDate:      startDate,
		Streams:        streams,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name2 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGcsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name2,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGcsResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGcsResourceModel) ToSharedSourceGcsPutRequest() *shared.SourceGcsPutRequest {
	bucket := r.Configuration.Bucket.ValueString()
	serviceAccount := r.Configuration.ServiceAccount.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceGcsUpdateFileBasedStreamConfig = []shared.SourceGcsUpdateFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceGcsUpdateFormat
		var sourceGcsUpdateAvroFormat *shared.SourceGcsUpdateAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceGcsUpdateAvroFormat = &shared.SourceGcsUpdateAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceGcsUpdateAvroFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				SourceGcsUpdateAvroFormat: sourceGcsUpdateAvroFormat,
			}
		}
		var sourceGcsUpdateCSVFormat *shared.SourceGcsUpdateCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = []string{}
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceGcsUpdateCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceGcsUpdateFromCSV *shared.SourceGcsUpdateFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceGcsUpdateFromCSV = &shared.SourceGcsUpdateFromCSV{}
				}
				if sourceGcsUpdateFromCSV != nil {
					headerDefinition = &shared.SourceGcsUpdateCSVHeaderDefinition{
						SourceGcsUpdateFromCSV: sourceGcsUpdateFromCSV,
					}
				}
				var sourceGcsUpdateAutogenerated *shared.SourceGcsUpdateAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceGcsUpdateAutogenerated = &shared.SourceGcsUpdateAutogenerated{}
				}
				if sourceGcsUpdateAutogenerated != nil {
					headerDefinition = &shared.SourceGcsUpdateCSVHeaderDefinition{
						SourceGcsUpdateAutogenerated: sourceGcsUpdateAutogenerated,
					}
				}
				var sourceGcsUpdateUserProvided *shared.SourceGcsUpdateUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceGcsUpdateUserProvided = &shared.SourceGcsUpdateUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceGcsUpdateUserProvided != nil {
					headerDefinition = &shared.SourceGcsUpdateCSVHeaderDefinition{
						SourceGcsUpdateUserProvided: sourceGcsUpdateUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
			}
			inferenceType := new(shared.SourceGcsUpdateInferenceType)
			if !streamsItem.Format.CSVFormat.InferenceType.IsUnknown() && !streamsItem.Format.CSVFormat.InferenceType.IsNull() {
				*inferenceType = shared.SourceGcsUpdateInferenceType(streamsItem.Format.CSVFormat.InferenceType.ValueString())
			} else {
				inferenceType = nil
			}
			var nullValues []string = []string{}
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = []string{}
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceGcsUpdateCSVFormat = &shared.SourceGcsUpdateCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				InferenceType:                inferenceType,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceGcsUpdateCSVFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				SourceGcsUpdateCSVFormat: sourceGcsUpdateCSVFormat,
			}
		}
		var sourceGcsUpdateJsonlFormat *shared.SourceGcsUpdateJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceGcsUpdateJsonlFormat = &shared.SourceGcsUpdateJsonlFormat{}
		}
		if sourceGcsUpdateJsonlFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				SourceGcsUpdateJsonlFormat: sourceGcsUpdateJsonlFormat,
			}
		}
		var sourceGcsUpdateParquetFormat *shared.SourceGcsUpdateParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceGcsUpdateParquetFormat = &shared.SourceGcsUpdateParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceGcsUpdateParquetFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				SourceGcsUpdateParquetFormat: sourceGcsUpdateParquetFormat,
			}
		}
		var unstructuredDocumentFormat *shared.UnstructuredDocumentFormat
		if streamsItem.Format.UnstructuredDocumentFormat != nil {
			var processing *shared.SourceGcsUpdateProcessing
			if streamsItem.Format.UnstructuredDocumentFormat.Processing != nil {
				var sourceGcsUpdateLocal *shared.SourceGcsUpdateLocal
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.Local != nil {
					sourceGcsUpdateLocal = &shared.SourceGcsUpdateLocal{}
				}
				if sourceGcsUpdateLocal != nil {
					processing = &shared.SourceGcsUpdateProcessing{
						SourceGcsUpdateLocal: sourceGcsUpdateLocal,
					}
				}
				var viaAPI *shared.ViaAPI
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI != nil {
					apiKey := new(string)
					if !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIKey.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIKey.IsNull() {
						*apiKey = streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIKey.ValueString()
					} else {
						apiKey = nil
					}
					apiURL := new(string)
					if !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIURL.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIURL.IsNull() {
						*apiURL = streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.APIURL.ValueString()
					} else {
						apiURL = nil
					}
					var parameters []shared.APIParameterConfigModel = []shared.APIParameterConfigModel{}
					for _, parametersItem := range streamsItem.Format.UnstructuredDocumentFormat.Processing.ViaAPI.Parameters {
						name := parametersItem.Name.ValueString()
						value := parametersItem.Value.ValueString()
						parameters = append(parameters, shared.APIParameterConfigModel{
							Name:  name,
							Value: value,
						})
					}
					viaAPI = &shared.ViaAPI{
						APIKey:     apiKey,
						APIURL:     apiURL,
						Parameters: parameters,
					}
				}
				if viaAPI != nil {
					processing = &shared.SourceGcsUpdateProcessing{
						ViaAPI: viaAPI,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceGcsUpdateParsingStrategy)
			if !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsNull() {
				*strategy = shared.SourceGcsUpdateParsingStrategy(streamsItem.Format.UnstructuredDocumentFormat.Strategy.ValueString())
			} else {
				strategy = nil
			}
			unstructuredDocumentFormat = &shared.UnstructuredDocumentFormat{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if unstructuredDocumentFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				UnstructuredDocumentFormat: unstructuredDocumentFormat,
			}
		}
		var excelFormat *shared.ExcelFormat
		if streamsItem.Format.ExcelFormat != nil {
			excelFormat = &shared.ExcelFormat{}
		}
		if excelFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				ExcelFormat: excelFormat,
			}
		}
		var globs []string = []string{}
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		legacyPrefix := new(string)
		if !streamsItem.LegacyPrefix.IsUnknown() && !streamsItem.LegacyPrefix.IsNull() {
			*legacyPrefix = streamsItem.LegacyPrefix.ValueString()
		} else {
			legacyPrefix = nil
		}
		name1 := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		recentNFilesToReadForSchemaDiscovery := new(int64)
		if !streamsItem.RecentNFilesToReadForSchemaDiscovery.IsUnknown() && !streamsItem.RecentNFilesToReadForSchemaDiscovery.IsNull() {
			*recentNFilesToReadForSchemaDiscovery = streamsItem.RecentNFilesToReadForSchemaDiscovery.ValueInt64()
		} else {
			recentNFilesToReadForSchemaDiscovery = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceGcsUpdateValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceGcsUpdateValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceGcsUpdateFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull:            daysToSyncIfHistoryIsFull,
			Format:                               format,
			Globs:                                globs,
			InputSchema:                          inputSchema,
			LegacyPrefix:                         legacyPrefix,
			Name:                                 name1,
			PrimaryKey:                           primaryKey,
			RecentNFilesToReadForSchemaDiscovery: recentNFilesToReadForSchemaDiscovery,
			Schemaless:                           schemaless,
			ValidationPolicy:                     validationPolicy,
		})
	}
	configuration := shared.SourceGcsUpdate{
		Bucket:         bucket,
		ServiceAccount: serviceAccount,
		StartDate:      startDate,
		Streams:        streams,
	}
	name2 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGcsPutRequest{
		Configuration: configuration,
		Name:          name2,
		WorkspaceID:   workspaceID,
	}
	return &out
}
