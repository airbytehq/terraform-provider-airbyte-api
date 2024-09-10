// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceMicrosoftSharepointResourceModel) ToSharedSourceMicrosoftSharepointCreateRequest() *shared.SourceMicrosoftSharepointCreateRequest {
	var credentials shared.SourceMicrosoftSharepointAuthentication
	var sourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth *shared.SourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth
	if r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientSecret.ValueString()
		refreshToken := new(string)
		if !r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.IsNull() {
			*refreshToken = r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.ValueString()
		} else {
			refreshToken = nil
		}
		tenantID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.TenantID.ValueString()
		sourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth = &shared.SourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
			TenantID:     tenantID,
		}
	}
	if sourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth != nil {
		credentials = shared.SourceMicrosoftSharepointAuthentication{
			SourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth: sourceMicrosoftSharepointAuthenticateViaMicrosoftOAuth,
		}
	}
	var sourceMicrosoftSharepointServiceKeyAuthentication *shared.SourceMicrosoftSharepointServiceKeyAuthentication
	if r.Configuration.Credentials.ServiceKeyAuthentication != nil {
		clientId1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientID.ValueString()
		clientSecret1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientSecret.ValueString()
		tenantId1 := r.Configuration.Credentials.ServiceKeyAuthentication.TenantID.ValueString()
		userPrincipalName := r.Configuration.Credentials.ServiceKeyAuthentication.UserPrincipalName.ValueString()
		sourceMicrosoftSharepointServiceKeyAuthentication = &shared.SourceMicrosoftSharepointServiceKeyAuthentication{
			ClientID:          clientId1,
			ClientSecret:      clientSecret1,
			TenantID:          tenantId1,
			UserPrincipalName: userPrincipalName,
		}
	}
	if sourceMicrosoftSharepointServiceKeyAuthentication != nil {
		credentials = shared.SourceMicrosoftSharepointAuthentication{
			SourceMicrosoftSharepointServiceKeyAuthentication: sourceMicrosoftSharepointServiceKeyAuthentication,
		}
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	searchScope := new(shared.SourceMicrosoftSharepointSearchScope)
	if !r.Configuration.SearchScope.IsUnknown() && !r.Configuration.SearchScope.IsNull() {
		*searchScope = shared.SourceMicrosoftSharepointSearchScope(r.Configuration.SearchScope.ValueString())
	} else {
		searchScope = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceMicrosoftSharepointFileBasedStreamConfig = []shared.SourceMicrosoftSharepointFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceMicrosoftSharepointFormat
		var sourceMicrosoftSharepointAvroFormat *shared.SourceMicrosoftSharepointAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceMicrosoftSharepointAvroFormat = &shared.SourceMicrosoftSharepointAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceMicrosoftSharepointAvroFormat != nil {
			format = shared.SourceMicrosoftSharepointFormat{
				SourceMicrosoftSharepointAvroFormat: sourceMicrosoftSharepointAvroFormat,
			}
		}
		var sourceMicrosoftSharepointCSVFormat *shared.SourceMicrosoftSharepointCSVFormat
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
			var headerDefinition *shared.SourceMicrosoftSharepointCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceMicrosoftSharepointFromCSV *shared.SourceMicrosoftSharepointFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceMicrosoftSharepointFromCSV = &shared.SourceMicrosoftSharepointFromCSV{}
				}
				if sourceMicrosoftSharepointFromCSV != nil {
					headerDefinition = &shared.SourceMicrosoftSharepointCSVHeaderDefinition{
						SourceMicrosoftSharepointFromCSV: sourceMicrosoftSharepointFromCSV,
					}
				}
				var sourceMicrosoftSharepointAutogenerated *shared.SourceMicrosoftSharepointAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceMicrosoftSharepointAutogenerated = &shared.SourceMicrosoftSharepointAutogenerated{}
				}
				if sourceMicrosoftSharepointAutogenerated != nil {
					headerDefinition = &shared.SourceMicrosoftSharepointCSVHeaderDefinition{
						SourceMicrosoftSharepointAutogenerated: sourceMicrosoftSharepointAutogenerated,
					}
				}
				var sourceMicrosoftSharepointUserProvided *shared.SourceMicrosoftSharepointUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceMicrosoftSharepointUserProvided = &shared.SourceMicrosoftSharepointUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceMicrosoftSharepointUserProvided != nil {
					headerDefinition = &shared.SourceMicrosoftSharepointCSVHeaderDefinition{
						SourceMicrosoftSharepointUserProvided: sourceMicrosoftSharepointUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
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
			sourceMicrosoftSharepointCSVFormat = &shared.SourceMicrosoftSharepointCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceMicrosoftSharepointCSVFormat != nil {
			format = shared.SourceMicrosoftSharepointFormat{
				SourceMicrosoftSharepointCSVFormat: sourceMicrosoftSharepointCSVFormat,
			}
		}
		var sourceMicrosoftSharepointJsonlFormat *shared.SourceMicrosoftSharepointJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceMicrosoftSharepointJsonlFormat = &shared.SourceMicrosoftSharepointJsonlFormat{}
		}
		if sourceMicrosoftSharepointJsonlFormat != nil {
			format = shared.SourceMicrosoftSharepointFormat{
				SourceMicrosoftSharepointJsonlFormat: sourceMicrosoftSharepointJsonlFormat,
			}
		}
		var sourceMicrosoftSharepointParquetFormat *shared.SourceMicrosoftSharepointParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceMicrosoftSharepointParquetFormat = &shared.SourceMicrosoftSharepointParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceMicrosoftSharepointParquetFormat != nil {
			format = shared.SourceMicrosoftSharepointFormat{
				SourceMicrosoftSharepointParquetFormat: sourceMicrosoftSharepointParquetFormat,
			}
		}
		var sourceMicrosoftSharepointUnstructuredDocumentFormat *shared.SourceMicrosoftSharepointUnstructuredDocumentFormat
		if streamsItem.Format.UnstructuredDocumentFormat != nil {
			var processing *shared.SourceMicrosoftSharepointProcessing
			if streamsItem.Format.UnstructuredDocumentFormat.Processing != nil {
				var sourceMicrosoftSharepointLocal *shared.SourceMicrosoftSharepointLocal
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.Local != nil {
					sourceMicrosoftSharepointLocal = &shared.SourceMicrosoftSharepointLocal{}
				}
				if sourceMicrosoftSharepointLocal != nil {
					processing = &shared.SourceMicrosoftSharepointProcessing{
						SourceMicrosoftSharepointLocal: sourceMicrosoftSharepointLocal,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceMicrosoftSharepointParsingStrategy)
			if !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsNull() {
				*strategy = shared.SourceMicrosoftSharepointParsingStrategy(streamsItem.Format.UnstructuredDocumentFormat.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceMicrosoftSharepointUnstructuredDocumentFormat = &shared.SourceMicrosoftSharepointUnstructuredDocumentFormat{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceMicrosoftSharepointUnstructuredDocumentFormat != nil {
			format = shared.SourceMicrosoftSharepointFormat{
				SourceMicrosoftSharepointUnstructuredDocumentFormat: sourceMicrosoftSharepointUnstructuredDocumentFormat,
			}
		}
		var sourceMicrosoftSharepointExcelFormat *shared.SourceMicrosoftSharepointExcelFormat
		if streamsItem.Format.ExcelFormat != nil {
			sourceMicrosoftSharepointExcelFormat = &shared.SourceMicrosoftSharepointExcelFormat{}
		}
		if sourceMicrosoftSharepointExcelFormat != nil {
			format = shared.SourceMicrosoftSharepointFormat{
				SourceMicrosoftSharepointExcelFormat: sourceMicrosoftSharepointExcelFormat,
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
		name := streamsItem.Name.ValueString()
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
		validationPolicy := new(shared.SourceMicrosoftSharepointValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceMicrosoftSharepointValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceMicrosoftSharepointFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull:            daysToSyncIfHistoryIsFull,
			Format:                               format,
			Globs:                                globs,
			InputSchema:                          inputSchema,
			Name:                                 name,
			PrimaryKey:                           primaryKey,
			RecentNFilesToReadForSchemaDiscovery: recentNFilesToReadForSchemaDiscovery,
			Schemaless:                           schemaless,
			ValidationPolicy:                     validationPolicy,
		})
	}
	configuration := shared.SourceMicrosoftSharepoint{
		Credentials: credentials,
		FolderPath:  folderPath,
		SearchScope: searchScope,
		StartDate:   startDate,
		Streams:     streams,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMicrosoftSharepointCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMicrosoftSharepointResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceMicrosoftSharepointResourceModel) ToSharedSourceMicrosoftSharepointPutRequest() *shared.SourceMicrosoftSharepointPutRequest {
	var credentials shared.SourceMicrosoftSharepointUpdateAuthentication
	var sourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth *shared.SourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth
	if r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientSecret.ValueString()
		refreshToken := new(string)
		if !r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.IsUnknown() && !r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.IsNull() {
			*refreshToken = r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.ValueString()
		} else {
			refreshToken = nil
		}
		tenantID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.TenantID.ValueString()
		sourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth = &shared.SourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
			TenantID:     tenantID,
		}
	}
	if sourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth != nil {
		credentials = shared.SourceMicrosoftSharepointUpdateAuthentication{
			SourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth: sourceMicrosoftSharepointUpdateAuthenticateViaMicrosoftOAuth,
		}
	}
	var sourceMicrosoftSharepointUpdateServiceKeyAuthentication *shared.SourceMicrosoftSharepointUpdateServiceKeyAuthentication
	if r.Configuration.Credentials.ServiceKeyAuthentication != nil {
		clientId1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientID.ValueString()
		clientSecret1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientSecret.ValueString()
		tenantId1 := r.Configuration.Credentials.ServiceKeyAuthentication.TenantID.ValueString()
		userPrincipalName := r.Configuration.Credentials.ServiceKeyAuthentication.UserPrincipalName.ValueString()
		sourceMicrosoftSharepointUpdateServiceKeyAuthentication = &shared.SourceMicrosoftSharepointUpdateServiceKeyAuthentication{
			ClientID:          clientId1,
			ClientSecret:      clientSecret1,
			TenantID:          tenantId1,
			UserPrincipalName: userPrincipalName,
		}
	}
	if sourceMicrosoftSharepointUpdateServiceKeyAuthentication != nil {
		credentials = shared.SourceMicrosoftSharepointUpdateAuthentication{
			SourceMicrosoftSharepointUpdateServiceKeyAuthentication: sourceMicrosoftSharepointUpdateServiceKeyAuthentication,
		}
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	searchScope := new(shared.SourceMicrosoftSharepointUpdateSearchScope)
	if !r.Configuration.SearchScope.IsUnknown() && !r.Configuration.SearchScope.IsNull() {
		*searchScope = shared.SourceMicrosoftSharepointUpdateSearchScope(r.Configuration.SearchScope.ValueString())
	} else {
		searchScope = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceMicrosoftSharepointUpdateFileBasedStreamConfig = []shared.SourceMicrosoftSharepointUpdateFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceMicrosoftSharepointUpdateFormat
		var sourceMicrosoftSharepointUpdateAvroFormat *shared.SourceMicrosoftSharepointUpdateAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceMicrosoftSharepointUpdateAvroFormat = &shared.SourceMicrosoftSharepointUpdateAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceMicrosoftSharepointUpdateAvroFormat != nil {
			format = shared.SourceMicrosoftSharepointUpdateFormat{
				SourceMicrosoftSharepointUpdateAvroFormat: sourceMicrosoftSharepointUpdateAvroFormat,
			}
		}
		var sourceMicrosoftSharepointUpdateCSVFormat *shared.SourceMicrosoftSharepointUpdateCSVFormat
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
			var headerDefinition *shared.SourceMicrosoftSharepointUpdateCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceMicrosoftSharepointUpdateFromCSV *shared.SourceMicrosoftSharepointUpdateFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceMicrosoftSharepointUpdateFromCSV = &shared.SourceMicrosoftSharepointUpdateFromCSV{}
				}
				if sourceMicrosoftSharepointUpdateFromCSV != nil {
					headerDefinition = &shared.SourceMicrosoftSharepointUpdateCSVHeaderDefinition{
						SourceMicrosoftSharepointUpdateFromCSV: sourceMicrosoftSharepointUpdateFromCSV,
					}
				}
				var sourceMicrosoftSharepointUpdateAutogenerated *shared.SourceMicrosoftSharepointUpdateAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceMicrosoftSharepointUpdateAutogenerated = &shared.SourceMicrosoftSharepointUpdateAutogenerated{}
				}
				if sourceMicrosoftSharepointUpdateAutogenerated != nil {
					headerDefinition = &shared.SourceMicrosoftSharepointUpdateCSVHeaderDefinition{
						SourceMicrosoftSharepointUpdateAutogenerated: sourceMicrosoftSharepointUpdateAutogenerated,
					}
				}
				var sourceMicrosoftSharepointUpdateUserProvided *shared.SourceMicrosoftSharepointUpdateUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceMicrosoftSharepointUpdateUserProvided = &shared.SourceMicrosoftSharepointUpdateUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceMicrosoftSharepointUpdateUserProvided != nil {
					headerDefinition = &shared.SourceMicrosoftSharepointUpdateCSVHeaderDefinition{
						SourceMicrosoftSharepointUpdateUserProvided: sourceMicrosoftSharepointUpdateUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
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
			sourceMicrosoftSharepointUpdateCSVFormat = &shared.SourceMicrosoftSharepointUpdateCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceMicrosoftSharepointUpdateCSVFormat != nil {
			format = shared.SourceMicrosoftSharepointUpdateFormat{
				SourceMicrosoftSharepointUpdateCSVFormat: sourceMicrosoftSharepointUpdateCSVFormat,
			}
		}
		var sourceMicrosoftSharepointUpdateJsonlFormat *shared.SourceMicrosoftSharepointUpdateJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceMicrosoftSharepointUpdateJsonlFormat = &shared.SourceMicrosoftSharepointUpdateJsonlFormat{}
		}
		if sourceMicrosoftSharepointUpdateJsonlFormat != nil {
			format = shared.SourceMicrosoftSharepointUpdateFormat{
				SourceMicrosoftSharepointUpdateJsonlFormat: sourceMicrosoftSharepointUpdateJsonlFormat,
			}
		}
		var sourceMicrosoftSharepointUpdateParquetFormat *shared.SourceMicrosoftSharepointUpdateParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceMicrosoftSharepointUpdateParquetFormat = &shared.SourceMicrosoftSharepointUpdateParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceMicrosoftSharepointUpdateParquetFormat != nil {
			format = shared.SourceMicrosoftSharepointUpdateFormat{
				SourceMicrosoftSharepointUpdateParquetFormat: sourceMicrosoftSharepointUpdateParquetFormat,
			}
		}
		var sourceMicrosoftSharepointUpdateUnstructuredDocumentFormat *shared.SourceMicrosoftSharepointUpdateUnstructuredDocumentFormat
		if streamsItem.Format.UnstructuredDocumentFormat != nil {
			var processing *shared.SourceMicrosoftSharepointUpdateProcessing
			if streamsItem.Format.UnstructuredDocumentFormat.Processing != nil {
				var sourceMicrosoftSharepointUpdateLocal *shared.SourceMicrosoftSharepointUpdateLocal
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.Local != nil {
					sourceMicrosoftSharepointUpdateLocal = &shared.SourceMicrosoftSharepointUpdateLocal{}
				}
				if sourceMicrosoftSharepointUpdateLocal != nil {
					processing = &shared.SourceMicrosoftSharepointUpdateProcessing{
						SourceMicrosoftSharepointUpdateLocal: sourceMicrosoftSharepointUpdateLocal,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceMicrosoftSharepointUpdateParsingStrategy)
			if !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsNull() {
				*strategy = shared.SourceMicrosoftSharepointUpdateParsingStrategy(streamsItem.Format.UnstructuredDocumentFormat.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceMicrosoftSharepointUpdateUnstructuredDocumentFormat = &shared.SourceMicrosoftSharepointUpdateUnstructuredDocumentFormat{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceMicrosoftSharepointUpdateUnstructuredDocumentFormat != nil {
			format = shared.SourceMicrosoftSharepointUpdateFormat{
				SourceMicrosoftSharepointUpdateUnstructuredDocumentFormat: sourceMicrosoftSharepointUpdateUnstructuredDocumentFormat,
			}
		}
		var sourceMicrosoftSharepointUpdateExcelFormat *shared.SourceMicrosoftSharepointUpdateExcelFormat
		if streamsItem.Format.ExcelFormat != nil {
			sourceMicrosoftSharepointUpdateExcelFormat = &shared.SourceMicrosoftSharepointUpdateExcelFormat{}
		}
		if sourceMicrosoftSharepointUpdateExcelFormat != nil {
			format = shared.SourceMicrosoftSharepointUpdateFormat{
				SourceMicrosoftSharepointUpdateExcelFormat: sourceMicrosoftSharepointUpdateExcelFormat,
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
		name := streamsItem.Name.ValueString()
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
		validationPolicy := new(shared.SourceMicrosoftSharepointUpdateValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceMicrosoftSharepointUpdateValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceMicrosoftSharepointUpdateFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull:            daysToSyncIfHistoryIsFull,
			Format:                               format,
			Globs:                                globs,
			InputSchema:                          inputSchema,
			Name:                                 name,
			PrimaryKey:                           primaryKey,
			RecentNFilesToReadForSchemaDiscovery: recentNFilesToReadForSchemaDiscovery,
			Schemaless:                           schemaless,
			ValidationPolicy:                     validationPolicy,
		})
	}
	configuration := shared.SourceMicrosoftSharepointUpdate{
		Credentials: credentials,
		FolderPath:  folderPath,
		SearchScope: searchScope,
		StartDate:   startDate,
		Streams:     streams,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMicrosoftSharepointPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}
