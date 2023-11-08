// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationVerticaResourceModel) ToCreateSDKType() *shared.DestinationVerticaCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
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
	schema := r.Configuration.Schema.ValueString()
	var tunnelMethod *shared.DestinationVerticaSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationVerticaNoTunnel *shared.DestinationVerticaNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationVerticaNoTunnel = &shared.DestinationVerticaNoTunnel{}
		}
		if destinationVerticaNoTunnel != nil {
			tunnelMethod = &shared.DestinationVerticaSSHTunnelMethod{
				DestinationVerticaNoTunnel: destinationVerticaNoTunnel,
			}
		}
		var destinationVerticaSSHKeyAuthentication *shared.DestinationVerticaSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationVerticaSSHKeyAuthentication = &shared.DestinationVerticaSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationVerticaSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationVerticaSSHTunnelMethod{
				DestinationVerticaSSHKeyAuthentication: destinationVerticaSSHKeyAuthentication,
			}
		}
		var destinationVerticaPasswordAuthentication *shared.DestinationVerticaPasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationVerticaPasswordAuthentication = &shared.DestinationVerticaPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationVerticaPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationVerticaSSHTunnelMethod{
				DestinationVerticaPasswordAuthentication: destinationVerticaPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationVertica{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schema:        schema,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationVerticaCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationVerticaResourceModel) ToGetSDKType() *shared.DestinationVerticaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationVerticaResourceModel) ToUpdateSDKType() *shared.DestinationVerticaPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
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
	schema := r.Configuration.Schema.ValueString()
	var tunnelMethod *shared.DestinationVerticaUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationVerticaUpdateNoTunnel *shared.DestinationVerticaUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationVerticaUpdateNoTunnel = &shared.DestinationVerticaUpdateNoTunnel{}
		}
		if destinationVerticaUpdateNoTunnel != nil {
			tunnelMethod = &shared.DestinationVerticaUpdateSSHTunnelMethod{
				DestinationVerticaUpdateNoTunnel: destinationVerticaUpdateNoTunnel,
			}
		}
		var destinationVerticaUpdateSSHKeyAuthentication *shared.DestinationVerticaUpdateSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationVerticaUpdateSSHKeyAuthentication = &shared.DestinationVerticaUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationVerticaUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationVerticaUpdateSSHTunnelMethod{
				DestinationVerticaUpdateSSHKeyAuthentication: destinationVerticaUpdateSSHKeyAuthentication,
			}
		}
		var destinationVerticaUpdatePasswordAuthentication *shared.DestinationVerticaUpdatePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationVerticaUpdatePasswordAuthentication = &shared.DestinationVerticaUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationVerticaUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationVerticaUpdateSSHTunnelMethod{
				DestinationVerticaUpdatePasswordAuthentication: destinationVerticaUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationVerticaUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		Schema:        schema,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationVerticaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationVerticaResourceModel) ToDeleteSDKType() *shared.DestinationVerticaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationVerticaResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationVerticaResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
