// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSftpBulkResourceModel) ToSharedSourceSftpBulkCreateRequest() *shared.SourceSftpBulkCreateRequest {
	fileMostRecent := new(bool)
	if !r.Configuration.FileMostRecent.IsUnknown() && !r.Configuration.FileMostRecent.IsNull() {
		*fileMostRecent = r.Configuration.FileMostRecent.ValueBool()
	} else {
		fileMostRecent = nil
	}
	filePattern := new(string)
	if !r.Configuration.FilePattern.IsUnknown() && !r.Configuration.FilePattern.IsNull() {
		*filePattern = r.Configuration.FilePattern.ValueString()
	} else {
		filePattern = nil
	}
	fileType := new(shared.SourceSftpBulkFileType)
	if !r.Configuration.FileType.IsUnknown() && !r.Configuration.FileType.IsNull() {
		*fileType = shared.SourceSftpBulkFileType(r.Configuration.FileType.ValueString())
	} else {
		fileType = nil
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	privateKey := new(string)
	if !r.Configuration.PrivateKey.IsUnknown() && !r.Configuration.PrivateKey.IsNull() {
		*privateKey = r.Configuration.PrivateKey.ValueString()
	} else {
		privateKey = nil
	}
	separator := new(string)
	if !r.Configuration.Separator.IsUnknown() && !r.Configuration.Separator.IsNull() {
		*separator = r.Configuration.Separator.ValueString()
	} else {
		separator = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	streamName := r.Configuration.StreamName.ValueString()
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceSftpBulk{
		FileMostRecent: fileMostRecent,
		FilePattern:    filePattern,
		FileType:       fileType,
		FolderPath:     folderPath,
		Host:           host,
		Password:       password,
		Port:           port,
		PrivateKey:     privateKey,
		Separator:      separator,
		StartDate:      startDate,
		StreamName:     streamName,
		Username:       username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSftpBulkCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSftpBulkResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSftpBulkResourceModel) ToSharedSourceSftpBulkPutRequest() *shared.SourceSftpBulkPutRequest {
	fileMostRecent := new(bool)
	if !r.Configuration.FileMostRecent.IsUnknown() && !r.Configuration.FileMostRecent.IsNull() {
		*fileMostRecent = r.Configuration.FileMostRecent.ValueBool()
	} else {
		fileMostRecent = nil
	}
	filePattern := new(string)
	if !r.Configuration.FilePattern.IsUnknown() && !r.Configuration.FilePattern.IsNull() {
		*filePattern = r.Configuration.FilePattern.ValueString()
	} else {
		filePattern = nil
	}
	fileType := new(shared.FileType)
	if !r.Configuration.FileType.IsUnknown() && !r.Configuration.FileType.IsNull() {
		*fileType = shared.FileType(r.Configuration.FileType.ValueString())
	} else {
		fileType = nil
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	privateKey := new(string)
	if !r.Configuration.PrivateKey.IsUnknown() && !r.Configuration.PrivateKey.IsNull() {
		*privateKey = r.Configuration.PrivateKey.ValueString()
	} else {
		privateKey = nil
	}
	separator := new(string)
	if !r.Configuration.Separator.IsUnknown() && !r.Configuration.Separator.IsNull() {
		*separator = r.Configuration.Separator.ValueString()
	} else {
		separator = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	streamName := r.Configuration.StreamName.ValueString()
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceSftpBulkUpdate{
		FileMostRecent: fileMostRecent,
		FilePattern:    filePattern,
		FileType:       fileType,
		FolderPath:     folderPath,
		Host:           host,
		Password:       password,
		Port:           port,
		PrivateKey:     privateKey,
		Separator:      separator,
		StartDate:      startDate,
		StreamName:     streamName,
		Username:       username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSftpBulkPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
