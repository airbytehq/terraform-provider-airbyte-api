// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceFileSecureResourceModel) ToCreateSDKType() *shared.SourceFileSecureCreateRequest {
	datasetName := r.Configuration.DatasetName.ValueString()
	format := new(shared.SourceFileSecureFileFormat)
	if !r.Configuration.Format.IsUnknown() && !r.Configuration.Format.IsNull() {
		*format = shared.SourceFileSecureFileFormat(r.Configuration.Format.ValueString())
	} else {
		format = nil
	}
	var provider shared.SourceFileSecureStorageProvider
	var sourceFileSecureHTTPSPublicWeb *shared.SourceFileSecureHTTPSPublicWeb
	if r.Configuration.Provider.HTTPSPublicWeb != nil {
		userAgent := new(bool)
		if !r.Configuration.Provider.HTTPSPublicWeb.UserAgent.IsUnknown() && !r.Configuration.Provider.HTTPSPublicWeb.UserAgent.IsNull() {
			*userAgent = r.Configuration.Provider.HTTPSPublicWeb.UserAgent.ValueBool()
		} else {
			userAgent = nil
		}
		sourceFileSecureHTTPSPublicWeb = &shared.SourceFileSecureHTTPSPublicWeb{
			UserAgent: userAgent,
		}
	}
	if sourceFileSecureHTTPSPublicWeb != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureHTTPSPublicWeb: sourceFileSecureHTTPSPublicWeb,
		}
	}
	var sourceFileSecureGCSGoogleCloudStorage *shared.SourceFileSecureGCSGoogleCloudStorage
	if r.Configuration.Provider.GCSGoogleCloudStorage != nil {
		serviceAccountJSON := new(string)
		if !r.Configuration.Provider.GCSGoogleCloudStorage.ServiceAccountJSON.IsUnknown() && !r.Configuration.Provider.GCSGoogleCloudStorage.ServiceAccountJSON.IsNull() {
			*serviceAccountJSON = r.Configuration.Provider.GCSGoogleCloudStorage.ServiceAccountJSON.ValueString()
		} else {
			serviceAccountJSON = nil
		}
		sourceFileSecureGCSGoogleCloudStorage = &shared.SourceFileSecureGCSGoogleCloudStorage{
			ServiceAccountJSON: serviceAccountJSON,
		}
	}
	if sourceFileSecureGCSGoogleCloudStorage != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureGCSGoogleCloudStorage: sourceFileSecureGCSGoogleCloudStorage,
		}
	}
	var sourceFileSecureS3AmazonWebServices *shared.SourceFileSecureS3AmazonWebServices
	if r.Configuration.Provider.S3AmazonWebServices != nil {
		awsAccessKeyID := new(string)
		if !r.Configuration.Provider.S3AmazonWebServices.AwsAccessKeyID.IsUnknown() && !r.Configuration.Provider.S3AmazonWebServices.AwsAccessKeyID.IsNull() {
			*awsAccessKeyID = r.Configuration.Provider.S3AmazonWebServices.AwsAccessKeyID.ValueString()
		} else {
			awsAccessKeyID = nil
		}
		awsSecretAccessKey := new(string)
		if !r.Configuration.Provider.S3AmazonWebServices.AwsSecretAccessKey.IsUnknown() && !r.Configuration.Provider.S3AmazonWebServices.AwsSecretAccessKey.IsNull() {
			*awsSecretAccessKey = r.Configuration.Provider.S3AmazonWebServices.AwsSecretAccessKey.ValueString()
		} else {
			awsSecretAccessKey = nil
		}
		sourceFileSecureS3AmazonWebServices = &shared.SourceFileSecureS3AmazonWebServices{
			AwsAccessKeyID:     awsAccessKeyID,
			AwsSecretAccessKey: awsSecretAccessKey,
		}
	}
	if sourceFileSecureS3AmazonWebServices != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureS3AmazonWebServices: sourceFileSecureS3AmazonWebServices,
		}
	}
	var sourceFileSecureAzBlobAzureBlobStorage *shared.SourceFileSecureAzBlobAzureBlobStorage
	if r.Configuration.Provider.AzBlobAzureBlobStorage != nil {
		sasToken := new(string)
		if !r.Configuration.Provider.AzBlobAzureBlobStorage.SasToken.IsUnknown() && !r.Configuration.Provider.AzBlobAzureBlobStorage.SasToken.IsNull() {
			*sasToken = r.Configuration.Provider.AzBlobAzureBlobStorage.SasToken.ValueString()
		} else {
			sasToken = nil
		}
		sharedKey := new(string)
		if !r.Configuration.Provider.AzBlobAzureBlobStorage.SharedKey.IsUnknown() && !r.Configuration.Provider.AzBlobAzureBlobStorage.SharedKey.IsNull() {
			*sharedKey = r.Configuration.Provider.AzBlobAzureBlobStorage.SharedKey.ValueString()
		} else {
			sharedKey = nil
		}
		storageAccount := r.Configuration.Provider.AzBlobAzureBlobStorage.StorageAccount.ValueString()
		sourceFileSecureAzBlobAzureBlobStorage = &shared.SourceFileSecureAzBlobAzureBlobStorage{
			SasToken:       sasToken,
			SharedKey:      sharedKey,
			StorageAccount: storageAccount,
		}
	}
	if sourceFileSecureAzBlobAzureBlobStorage != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureAzBlobAzureBlobStorage: sourceFileSecureAzBlobAzureBlobStorage,
		}
	}
	var sourceFileSecureSSHSecureShell *shared.SourceFileSecureSSHSecureShell
	if r.Configuration.Provider.SSHSecureShell != nil {
		host := r.Configuration.Provider.SSHSecureShell.Host.ValueString()
		password := new(string)
		if !r.Configuration.Provider.SSHSecureShell.Password.IsUnknown() && !r.Configuration.Provider.SSHSecureShell.Password.IsNull() {
			*password = r.Configuration.Provider.SSHSecureShell.Password.ValueString()
		} else {
			password = nil
		}
		port := new(string)
		if !r.Configuration.Provider.SSHSecureShell.Port.IsUnknown() && !r.Configuration.Provider.SSHSecureShell.Port.IsNull() {
			*port = r.Configuration.Provider.SSHSecureShell.Port.ValueString()
		} else {
			port = nil
		}
		user := r.Configuration.Provider.SSHSecureShell.User.ValueString()
		sourceFileSecureSSHSecureShell = &shared.SourceFileSecureSSHSecureShell{
			Host:     host,
			Password: password,
			Port:     port,
			User:     user,
		}
	}
	if sourceFileSecureSSHSecureShell != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureSSHSecureShell: sourceFileSecureSSHSecureShell,
		}
	}
	var sourceFileSecureSCPSecureCopyProtocol *shared.SourceFileSecureSCPSecureCopyProtocol
	if r.Configuration.Provider.SCPSecureCopyProtocol != nil {
		host1 := r.Configuration.Provider.SCPSecureCopyProtocol.Host.ValueString()
		password1 := new(string)
		if !r.Configuration.Provider.SCPSecureCopyProtocol.Password.IsUnknown() && !r.Configuration.Provider.SCPSecureCopyProtocol.Password.IsNull() {
			*password1 = r.Configuration.Provider.SCPSecureCopyProtocol.Password.ValueString()
		} else {
			password1 = nil
		}
		port1 := new(string)
		if !r.Configuration.Provider.SCPSecureCopyProtocol.Port.IsUnknown() && !r.Configuration.Provider.SCPSecureCopyProtocol.Port.IsNull() {
			*port1 = r.Configuration.Provider.SCPSecureCopyProtocol.Port.ValueString()
		} else {
			port1 = nil
		}
		user1 := r.Configuration.Provider.SCPSecureCopyProtocol.User.ValueString()
		sourceFileSecureSCPSecureCopyProtocol = &shared.SourceFileSecureSCPSecureCopyProtocol{
			Host:     host1,
			Password: password1,
			Port:     port1,
			User:     user1,
		}
	}
	if sourceFileSecureSCPSecureCopyProtocol != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureSCPSecureCopyProtocol: sourceFileSecureSCPSecureCopyProtocol,
		}
	}
	var sourceFileSecureSFTPSecureFileTransferProtocol *shared.SourceFileSecureSFTPSecureFileTransferProtocol
	if r.Configuration.Provider.SFTPSecureFileTransferProtocol != nil {
		host2 := r.Configuration.Provider.SFTPSecureFileTransferProtocol.Host.ValueString()
		password2 := new(string)
		if !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Password.IsUnknown() && !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Password.IsNull() {
			*password2 = r.Configuration.Provider.SFTPSecureFileTransferProtocol.Password.ValueString()
		} else {
			password2 = nil
		}
		port2 := new(string)
		if !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Port.IsUnknown() && !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Port.IsNull() {
			*port2 = r.Configuration.Provider.SFTPSecureFileTransferProtocol.Port.ValueString()
		} else {
			port2 = nil
		}
		user2 := r.Configuration.Provider.SFTPSecureFileTransferProtocol.User.ValueString()
		sourceFileSecureSFTPSecureFileTransferProtocol = &shared.SourceFileSecureSFTPSecureFileTransferProtocol{
			Host:     host2,
			Password: password2,
			Port:     port2,
			User:     user2,
		}
	}
	if sourceFileSecureSFTPSecureFileTransferProtocol != nil {
		provider = shared.SourceFileSecureStorageProvider{
			SourceFileSecureSFTPSecureFileTransferProtocol: sourceFileSecureSFTPSecureFileTransferProtocol,
		}
	}
	readerOptions := new(string)
	if !r.Configuration.ReaderOptions.IsUnknown() && !r.Configuration.ReaderOptions.IsNull() {
		*readerOptions = r.Configuration.ReaderOptions.ValueString()
	} else {
		readerOptions = nil
	}
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceFileSecure{
		DatasetName:   datasetName,
		Format:        format,
		Provider:      provider,
		ReaderOptions: readerOptions,
		URL:           url,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFileSecureCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFileSecureResourceModel) ToGetSDKType() *shared.SourceFileSecureCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFileSecureResourceModel) ToUpdateSDKType() *shared.SourceFileSecurePutRequest {
	datasetName := r.Configuration.DatasetName.ValueString()
	format := new(shared.FileFormat)
	if !r.Configuration.Format.IsUnknown() && !r.Configuration.Format.IsNull() {
		*format = shared.FileFormat(r.Configuration.Format.ValueString())
	} else {
		format = nil
	}
	var provider shared.StorageProvider
	var httpsPublicWeb *shared.HTTPSPublicWeb
	if r.Configuration.Provider.HTTPSPublicWeb != nil {
		userAgent := new(bool)
		if !r.Configuration.Provider.HTTPSPublicWeb.UserAgent.IsUnknown() && !r.Configuration.Provider.HTTPSPublicWeb.UserAgent.IsNull() {
			*userAgent = r.Configuration.Provider.HTTPSPublicWeb.UserAgent.ValueBool()
		} else {
			userAgent = nil
		}
		httpsPublicWeb = &shared.HTTPSPublicWeb{
			UserAgent: userAgent,
		}
	}
	if httpsPublicWeb != nil {
		provider = shared.StorageProvider{
			HTTPSPublicWeb: httpsPublicWeb,
		}
	}
	var gcsGoogleCloudStorage *shared.GCSGoogleCloudStorage
	if r.Configuration.Provider.GCSGoogleCloudStorage != nil {
		serviceAccountJSON := new(string)
		if !r.Configuration.Provider.GCSGoogleCloudStorage.ServiceAccountJSON.IsUnknown() && !r.Configuration.Provider.GCSGoogleCloudStorage.ServiceAccountJSON.IsNull() {
			*serviceAccountJSON = r.Configuration.Provider.GCSGoogleCloudStorage.ServiceAccountJSON.ValueString()
		} else {
			serviceAccountJSON = nil
		}
		gcsGoogleCloudStorage = &shared.GCSGoogleCloudStorage{
			ServiceAccountJSON: serviceAccountJSON,
		}
	}
	if gcsGoogleCloudStorage != nil {
		provider = shared.StorageProvider{
			GCSGoogleCloudStorage: gcsGoogleCloudStorage,
		}
	}
	var sourceFileSecureUpdateS3AmazonWebServices *shared.SourceFileSecureUpdateS3AmazonWebServices
	if r.Configuration.Provider.S3AmazonWebServices != nil {
		awsAccessKeyID := new(string)
		if !r.Configuration.Provider.S3AmazonWebServices.AwsAccessKeyID.IsUnknown() && !r.Configuration.Provider.S3AmazonWebServices.AwsAccessKeyID.IsNull() {
			*awsAccessKeyID = r.Configuration.Provider.S3AmazonWebServices.AwsAccessKeyID.ValueString()
		} else {
			awsAccessKeyID = nil
		}
		awsSecretAccessKey := new(string)
		if !r.Configuration.Provider.S3AmazonWebServices.AwsSecretAccessKey.IsUnknown() && !r.Configuration.Provider.S3AmazonWebServices.AwsSecretAccessKey.IsNull() {
			*awsSecretAccessKey = r.Configuration.Provider.S3AmazonWebServices.AwsSecretAccessKey.ValueString()
		} else {
			awsSecretAccessKey = nil
		}
		sourceFileSecureUpdateS3AmazonWebServices = &shared.SourceFileSecureUpdateS3AmazonWebServices{
			AwsAccessKeyID:     awsAccessKeyID,
			AwsSecretAccessKey: awsSecretAccessKey,
		}
	}
	if sourceFileSecureUpdateS3AmazonWebServices != nil {
		provider = shared.StorageProvider{
			SourceFileSecureUpdateS3AmazonWebServices: sourceFileSecureUpdateS3AmazonWebServices,
		}
	}
	var azBlobAzureBlobStorage *shared.AzBlobAzureBlobStorage
	if r.Configuration.Provider.AzBlobAzureBlobStorage != nil {
		sasToken := new(string)
		if !r.Configuration.Provider.AzBlobAzureBlobStorage.SasToken.IsUnknown() && !r.Configuration.Provider.AzBlobAzureBlobStorage.SasToken.IsNull() {
			*sasToken = r.Configuration.Provider.AzBlobAzureBlobStorage.SasToken.ValueString()
		} else {
			sasToken = nil
		}
		sharedKey := new(string)
		if !r.Configuration.Provider.AzBlobAzureBlobStorage.SharedKey.IsUnknown() && !r.Configuration.Provider.AzBlobAzureBlobStorage.SharedKey.IsNull() {
			*sharedKey = r.Configuration.Provider.AzBlobAzureBlobStorage.SharedKey.ValueString()
		} else {
			sharedKey = nil
		}
		storageAccount := r.Configuration.Provider.AzBlobAzureBlobStorage.StorageAccount.ValueString()
		azBlobAzureBlobStorage = &shared.AzBlobAzureBlobStorage{
			SasToken:       sasToken,
			SharedKey:      sharedKey,
			StorageAccount: storageAccount,
		}
	}
	if azBlobAzureBlobStorage != nil {
		provider = shared.StorageProvider{
			AzBlobAzureBlobStorage: azBlobAzureBlobStorage,
		}
	}
	var sshSecureShell *shared.SSHSecureShell
	if r.Configuration.Provider.SSHSecureShell != nil {
		host := r.Configuration.Provider.SSHSecureShell.Host.ValueString()
		password := new(string)
		if !r.Configuration.Provider.SSHSecureShell.Password.IsUnknown() && !r.Configuration.Provider.SSHSecureShell.Password.IsNull() {
			*password = r.Configuration.Provider.SSHSecureShell.Password.ValueString()
		} else {
			password = nil
		}
		port := new(string)
		if !r.Configuration.Provider.SSHSecureShell.Port.IsUnknown() && !r.Configuration.Provider.SSHSecureShell.Port.IsNull() {
			*port = r.Configuration.Provider.SSHSecureShell.Port.ValueString()
		} else {
			port = nil
		}
		user := r.Configuration.Provider.SSHSecureShell.User.ValueString()
		sshSecureShell = &shared.SSHSecureShell{
			Host:     host,
			Password: password,
			Port:     port,
			User:     user,
		}
	}
	if sshSecureShell != nil {
		provider = shared.StorageProvider{
			SSHSecureShell: sshSecureShell,
		}
	}
	var scpSecureCopyProtocol *shared.SCPSecureCopyProtocol
	if r.Configuration.Provider.SCPSecureCopyProtocol != nil {
		host1 := r.Configuration.Provider.SCPSecureCopyProtocol.Host.ValueString()
		password1 := new(string)
		if !r.Configuration.Provider.SCPSecureCopyProtocol.Password.IsUnknown() && !r.Configuration.Provider.SCPSecureCopyProtocol.Password.IsNull() {
			*password1 = r.Configuration.Provider.SCPSecureCopyProtocol.Password.ValueString()
		} else {
			password1 = nil
		}
		port1 := new(string)
		if !r.Configuration.Provider.SCPSecureCopyProtocol.Port.IsUnknown() && !r.Configuration.Provider.SCPSecureCopyProtocol.Port.IsNull() {
			*port1 = r.Configuration.Provider.SCPSecureCopyProtocol.Port.ValueString()
		} else {
			port1 = nil
		}
		user1 := r.Configuration.Provider.SCPSecureCopyProtocol.User.ValueString()
		scpSecureCopyProtocol = &shared.SCPSecureCopyProtocol{
			Host:     host1,
			Password: password1,
			Port:     port1,
			User:     user1,
		}
	}
	if scpSecureCopyProtocol != nil {
		provider = shared.StorageProvider{
			SCPSecureCopyProtocol: scpSecureCopyProtocol,
		}
	}
	var sftpSecureFileTransferProtocol *shared.SFTPSecureFileTransferProtocol
	if r.Configuration.Provider.SFTPSecureFileTransferProtocol != nil {
		host2 := r.Configuration.Provider.SFTPSecureFileTransferProtocol.Host.ValueString()
		password2 := new(string)
		if !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Password.IsUnknown() && !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Password.IsNull() {
			*password2 = r.Configuration.Provider.SFTPSecureFileTransferProtocol.Password.ValueString()
		} else {
			password2 = nil
		}
		port2 := new(string)
		if !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Port.IsUnknown() && !r.Configuration.Provider.SFTPSecureFileTransferProtocol.Port.IsNull() {
			*port2 = r.Configuration.Provider.SFTPSecureFileTransferProtocol.Port.ValueString()
		} else {
			port2 = nil
		}
		user2 := r.Configuration.Provider.SFTPSecureFileTransferProtocol.User.ValueString()
		sftpSecureFileTransferProtocol = &shared.SFTPSecureFileTransferProtocol{
			Host:     host2,
			Password: password2,
			Port:     port2,
			User:     user2,
		}
	}
	if sftpSecureFileTransferProtocol != nil {
		provider = shared.StorageProvider{
			SFTPSecureFileTransferProtocol: sftpSecureFileTransferProtocol,
		}
	}
	readerOptions := new(string)
	if !r.Configuration.ReaderOptions.IsUnknown() && !r.Configuration.ReaderOptions.IsNull() {
		*readerOptions = r.Configuration.ReaderOptions.ValueString()
	} else {
		readerOptions = nil
	}
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceFileSecureUpdate{
		DatasetName:   datasetName,
		Format:        format,
		Provider:      provider,
		ReaderOptions: readerOptions,
		URL:           url,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFileSecurePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFileSecureResourceModel) ToDeleteSDKType() *shared.SourceFileSecureCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFileSecureResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFileSecureResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
