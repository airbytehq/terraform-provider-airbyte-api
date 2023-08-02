// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationSnowflakeResourceModel) ToCreateSDKType() *shared.DestinationSnowflakeCreateRequest {
	var credentials *shared.DestinationSnowflakeAuthorizationMethod
	var destinationSnowflakeAuthorizationMethodOAuth20 *shared.DestinationSnowflakeAuthorizationMethodOAuth20
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
		authType := new(shared.DestinationSnowflakeAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsNull() {
			*authType = shared.DestinationSnowflakeAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.ValueString())
		} else {
			authType = nil
		}
		destinationSnowflakeAuthorizationMethodOAuth20 = &shared.DestinationSnowflakeAuthorizationMethodOAuth20{
			AuthType: authType,
		}
	}
	if destinationSnowflakeAuthorizationMethodOAuth20 != nil {
		credentials = &shared.DestinationSnowflakeAuthorizationMethod{
			DestinationSnowflakeAuthorizationMethodOAuth20: destinationSnowflakeAuthorizationMethodOAuth20,
		}
	}
	var destinationSnowflakeAuthorizationMethodKeyPairAuthentication *shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20 != nil {
		authType1 := new(shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthenticationAuthType)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType1 = shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthenticationAuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		destinationSnowflakeAuthorizationMethodKeyPairAuthentication = &shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication{
			AuthType: authType1,
		}
	}
	if destinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
		credentials = &shared.DestinationSnowflakeAuthorizationMethod{
			DestinationSnowflakeAuthorizationMethodKeyPairAuthentication: destinationSnowflakeAuthorizationMethodKeyPairAuthentication,
		}
	}
	var destinationSnowflakeAuthorizationMethodUsernameAndPassword *shared.DestinationSnowflakeAuthorizationMethodUsernameAndPassword
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		authType2 := new(shared.DestinationSnowflakeAuthorizationMethodUsernameAndPasswordAuthType)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsNull() {
			*authType2 = shared.DestinationSnowflakeAuthorizationMethodUsernameAndPasswordAuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.ValueString())
		} else {
			authType2 = nil
		}
		password := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.Password.ValueString()
		destinationSnowflakeAuthorizationMethodUsernameAndPassword = &shared.DestinationSnowflakeAuthorizationMethodUsernameAndPassword{
			AuthType: authType2,
			Password: password,
		}
	}
	if destinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		credentials = &shared.DestinationSnowflakeAuthorizationMethod{
			DestinationSnowflakeAuthorizationMethodUsernameAndPassword: destinationSnowflakeAuthorizationMethodUsernameAndPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationSnowflakeSnowflake(r.Configuration.DestinationType.ValueString())
	fileBufferCount := new(int64)
	if !r.Configuration.FileBufferCount.IsUnknown() && !r.Configuration.FileBufferCount.IsNull() {
		*fileBufferCount = r.Configuration.FileBufferCount.ValueInt64()
	} else {
		fileBufferCount = nil
	}
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	var loadingMethod *shared.DestinationSnowflakeDataStagingMethod
	var destinationSnowflakeDataStagingMethodSelectAnotherOption *shared.DestinationSnowflakeDataStagingMethodSelectAnotherOption
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging != nil {
		method := shared.DestinationSnowflakeDataStagingMethodSelectAnotherOptionMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging.Method.ValueString())
		destinationSnowflakeDataStagingMethodSelectAnotherOption = &shared.DestinationSnowflakeDataStagingMethodSelectAnotherOption{
			Method: method,
		}
	}
	if destinationSnowflakeDataStagingMethodSelectAnotherOption != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodSelectAnotherOption: destinationSnowflakeDataStagingMethodSelectAnotherOption,
		}
	}
	var destinationSnowflakeDataStagingMethodRecommendedInternalStaging *shared.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging != nil {
		method1 := shared.DestinationSnowflakeDataStagingMethodRecommendedInternalStagingMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Method.ValueString())
		destinationSnowflakeDataStagingMethodRecommendedInternalStaging = &shared.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging{
			Method: method1,
		}
	}
	if destinationSnowflakeDataStagingMethodRecommendedInternalStaging != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodRecommendedInternalStaging: destinationSnowflakeDataStagingMethodRecommendedInternalStaging,
		}
	}
	var destinationSnowflakeDataStagingMethodAWSS3Staging *shared.DestinationSnowflakeDataStagingMethodAWSS3Staging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging != nil {
		method2 := shared.DestinationSnowflakeDataStagingMethodAWSS3StagingMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging.Method.ValueString())
		destinationSnowflakeDataStagingMethodAWSS3Staging = &shared.DestinationSnowflakeDataStagingMethodAWSS3Staging{
			Method: method2,
		}
	}
	if destinationSnowflakeDataStagingMethodAWSS3Staging != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodAWSS3Staging: destinationSnowflakeDataStagingMethodAWSS3Staging,
		}
	}
	var destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging *shared.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodSelectAnotherOption != nil {
		method3 := shared.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStagingMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodSelectAnotherOption.Method.ValueString())
		destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging = &shared.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging{
			Method: method3,
		}
	}
	if destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging != nil {
		loadingMethod = &shared.DestinationSnowflakeDataStagingMethod{
			DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging: destinationSnowflakeDataStagingMethodGoogleCloudStorageStaging,
		}
	}
	role := r.Configuration.Role.ValueString()
	schema := r.Configuration.Schema.ValueString()
	username := r.Configuration.Username.ValueString()
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.DestinationSnowflake{
		Credentials:     credentials,
		Database:        database,
		DestinationType: destinationType,
		FileBufferCount: fileBufferCount,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		LoadingMethod:   loadingMethod,
		Role:            role,
		Schema:          schema,
		Username:        username,
		Warehouse:       warehouse,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSnowflakeCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSnowflakeResourceModel) ToGetSDKType() *shared.DestinationSnowflakeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationSnowflakeResourceModel) ToUpdateSDKType() *shared.DestinationSnowflakePutRequest {
	var credentials *shared.DestinationSnowflakeUpdateAuthorizationMethod
	var destinationSnowflakeUpdateAuthorizationMethodOAuth20 *shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
		authType := new(shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsNull() {
			*authType = shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.ValueString())
		} else {
			authType = nil
		}
		destinationSnowflakeUpdateAuthorizationMethodOAuth20 = &shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20{
			AuthType: authType,
		}
	}
	if destinationSnowflakeUpdateAuthorizationMethodOAuth20 != nil {
		credentials = &shared.DestinationSnowflakeUpdateAuthorizationMethod{
			DestinationSnowflakeUpdateAuthorizationMethodOAuth20: destinationSnowflakeUpdateAuthorizationMethodOAuth20,
		}
	}
	var destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication *shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20 != nil {
		authType1 := new(shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthenticationAuthType)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType1 = shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthenticationAuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication = &shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication{
			AuthType: authType1,
		}
	}
	if destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication != nil {
		credentials = &shared.DestinationSnowflakeUpdateAuthorizationMethod{
			DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication: destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication,
		}
	}
	var destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword *shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword
	if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		authType2 := new(shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPasswordAuthType)
		if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsNull() {
			*authType2 = shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPasswordAuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.ValueString())
		} else {
			authType2 = nil
		}
		password := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.Password.ValueString()
		destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword = &shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword{
			AuthType: authType2,
			Password: password,
		}
	}
	if destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword != nil {
		credentials = &shared.DestinationSnowflakeUpdateAuthorizationMethod{
			DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword: destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	fileBufferCount := new(int64)
	if !r.Configuration.FileBufferCount.IsUnknown() && !r.Configuration.FileBufferCount.IsNull() {
		*fileBufferCount = r.Configuration.FileBufferCount.ValueInt64()
	} else {
		fileBufferCount = nil
	}
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	var loadingMethod *shared.DestinationSnowflakeUpdateDataStagingMethod
	var destinationSnowflakeUpdateDataStagingMethodSelectAnotherOption *shared.DestinationSnowflakeUpdateDataStagingMethodSelectAnotherOption
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging != nil {
		method := shared.DestinationSnowflakeUpdateDataStagingMethodSelectAnotherOptionMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodRecommendedInternalStaging.Method.ValueString())
		destinationSnowflakeUpdateDataStagingMethodSelectAnotherOption = &shared.DestinationSnowflakeUpdateDataStagingMethodSelectAnotherOption{
			Method: method,
		}
	}
	if destinationSnowflakeUpdateDataStagingMethodSelectAnotherOption != nil {
		loadingMethod = &shared.DestinationSnowflakeUpdateDataStagingMethod{
			DestinationSnowflakeUpdateDataStagingMethodSelectAnotherOption: destinationSnowflakeUpdateDataStagingMethodSelectAnotherOption,
		}
	}
	var destinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging *shared.DestinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging != nil {
		method1 := shared.DestinationSnowflakeUpdateDataStagingMethodRecommendedInternalStagingMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodAWSS3Staging.Method.ValueString())
		destinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging = &shared.DestinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging{
			Method: method1,
		}
	}
	if destinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging != nil {
		loadingMethod = &shared.DestinationSnowflakeUpdateDataStagingMethod{
			DestinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging: destinationSnowflakeUpdateDataStagingMethodRecommendedInternalStaging,
		}
	}
	var destinationSnowflakeUpdateDataStagingMethodAWSS3Staging *shared.DestinationSnowflakeUpdateDataStagingMethodAWSS3Staging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging != nil {
		method2 := shared.DestinationSnowflakeUpdateDataStagingMethodAWSS3StagingMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodGoogleCloudStorageStaging.Method.ValueString())
		destinationSnowflakeUpdateDataStagingMethodAWSS3Staging = &shared.DestinationSnowflakeUpdateDataStagingMethodAWSS3Staging{
			Method: method2,
		}
	}
	if destinationSnowflakeUpdateDataStagingMethodAWSS3Staging != nil {
		loadingMethod = &shared.DestinationSnowflakeUpdateDataStagingMethod{
			DestinationSnowflakeUpdateDataStagingMethodAWSS3Staging: destinationSnowflakeUpdateDataStagingMethodAWSS3Staging,
		}
	}
	var destinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging *shared.DestinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging
	if r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodSelectAnotherOption != nil {
		method3 := shared.DestinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStagingMethod(r.Configuration.LoadingMethod.DestinationSnowflakeDataStagingMethodSelectAnotherOption.Method.ValueString())
		destinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging = &shared.DestinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging{
			Method: method3,
		}
	}
	if destinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging != nil {
		loadingMethod = &shared.DestinationSnowflakeUpdateDataStagingMethod{
			DestinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging: destinationSnowflakeUpdateDataStagingMethodGoogleCloudStorageStaging,
		}
	}
	role := r.Configuration.Role.ValueString()
	schema := r.Configuration.Schema.ValueString()
	username := r.Configuration.Username.ValueString()
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.DestinationSnowflakeUpdate{
		Credentials:     credentials,
		Database:        database,
		FileBufferCount: fileBufferCount,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		LoadingMethod:   loadingMethod,
		Role:            role,
		Schema:          schema,
		Username:        username,
		Warehouse:       warehouse,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSnowflakePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSnowflakeResourceModel) ToDeleteSDKType() *shared.DestinationSnowflakeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationSnowflakeResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationSnowflakeResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
