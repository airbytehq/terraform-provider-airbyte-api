// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMongodbResourceModel) ToCreateSDKType() *shared.DestinationMongodbCreateRequest {
	var authType shared.DestinationMongodbAuthorizationType
	var destinationMongodbAuthorizationTypeNone *shared.DestinationMongodbAuthorizationTypeNone
	if r.Configuration.AuthType.DestinationMongodbAuthorizationTypeLoginPassword != nil {
		authorization := shared.DestinationMongodbAuthorizationTypeNoneAuthorization(r.Configuration.AuthType.DestinationMongodbAuthorizationTypeLoginPassword.Authorization.ValueString())
		destinationMongodbAuthorizationTypeNone = &shared.DestinationMongodbAuthorizationTypeNone{
			Authorization: authorization,
		}
	}
	if destinationMongodbAuthorizationTypeNone != nil {
		authType = shared.DestinationMongodbAuthorizationType{
			DestinationMongodbAuthorizationTypeNone: destinationMongodbAuthorizationTypeNone,
		}
	}
	var destinationMongodbAuthorizationTypeLoginPassword *shared.DestinationMongodbAuthorizationTypeLoginPassword
	if r.Configuration.AuthType.DestinationMongodbAuthorizationTypeNone != nil {
		authorization1 := shared.DestinationMongodbAuthorizationTypeLoginPasswordAuthorization(r.Configuration.AuthType.DestinationMongodbAuthorizationTypeNone.Authorization.ValueString())
		destinationMongodbAuthorizationTypeLoginPassword = &shared.DestinationMongodbAuthorizationTypeLoginPassword{
			Authorization: authorization1,
		}
	}
	if destinationMongodbAuthorizationTypeLoginPassword != nil {
		authType = shared.DestinationMongodbAuthorizationType{
			DestinationMongodbAuthorizationTypeLoginPassword: destinationMongodbAuthorizationTypeLoginPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationMongodbMongodb(r.Configuration.DestinationType.ValueString())
	var instanceType *shared.DestinationMongodbMongoDbInstanceType
	var destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance *shared.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance
	if r.Configuration.InstanceType.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
		instance := shared.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstanceInstance(r.Configuration.InstanceType.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas.Instance.ValueString())
		destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance = &shared.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance{
			Instance: instance,
		}
	}
	if destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		instanceType = &shared.DestinationMongodbMongoDbInstanceType{
			DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance: destinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance,
		}
	}
	var destinationMongodbMongoDbInstanceTypeReplicaSet *shared.DestinationMongodbMongoDbInstanceTypeReplicaSet
	if r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet != nil {
		instance1 := shared.DestinationMongodbMongoDbInstanceTypeReplicaSetInstance(r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.Instance.ValueString())
		replicaSet := new(string)
		if !r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.IsUnknown() && !r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.IsNull() {
			*replicaSet = r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.ValueString()
		} else {
			replicaSet = nil
		}
		serverAddresses := r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ServerAddresses.ValueString()
		destinationMongodbMongoDbInstanceTypeReplicaSet = &shared.DestinationMongodbMongoDbInstanceTypeReplicaSet{
			Instance:        instance1,
			ReplicaSet:      replicaSet,
			ServerAddresses: serverAddresses,
		}
	}
	if destinationMongodbMongoDbInstanceTypeReplicaSet != nil {
		instanceType = &shared.DestinationMongodbMongoDbInstanceType{
			DestinationMongodbMongoDbInstanceTypeReplicaSet: destinationMongodbMongoDbInstanceTypeReplicaSet,
		}
	}
	var destinationMongodbMongoDBInstanceTypeMongoDBAtlas *shared.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas
	if r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		instance2 := shared.DestinationMongodbMongoDBInstanceTypeMongoDBAtlasInstance(r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance.Instance.ValueString())
		destinationMongodbMongoDBInstanceTypeMongoDBAtlas = &shared.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas{
			Instance: instance2,
		}
	}
	if destinationMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
		instanceType = &shared.DestinationMongodbMongoDbInstanceType{
			DestinationMongodbMongoDBInstanceTypeMongoDBAtlas: destinationMongodbMongoDBInstanceTypeMongoDBAtlas,
		}
	}
	var tunnelMethod *shared.DestinationMongodbSSHTunnelMethod
	var destinationMongodbSSHTunnelMethodNoTunnel *shared.DestinationMongodbSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationMongodbSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationMongodbSSHTunnelMethodNoTunnel = &shared.DestinationMongodbSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationMongodbSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationMongodbSSHTunnelMethod{
			DestinationMongodbSSHTunnelMethodNoTunnel: destinationMongodbSSHTunnelMethodNoTunnel,
		}
	}
	var destinationMongodbSSHTunnelMethodSSHKeyAuthentication *shared.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationMongodbSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		destinationMongodbSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication{
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationMongodbSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationMongodbSSHTunnelMethod{
			DestinationMongodbSSHTunnelMethodSSHKeyAuthentication: destinationMongodbSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationMongodbSSHTunnelMethodPasswordAuthentication *shared.DestinationMongodbSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationMongodbSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationMongodbSSHTunnelMethodPasswordAuthentication = &shared.DestinationMongodbSSHTunnelMethodPasswordAuthentication{
			TunnelHost:   tunnelHost1,
			TunnelMethod: tunnelMethod3,
			TunnelPort:   tunnelPort1,
			TunnelUser:   tunnelUser1,
		}
	}
	if destinationMongodbSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationMongodbSSHTunnelMethod{
			DestinationMongodbSSHTunnelMethodPasswordAuthentication: destinationMongodbSSHTunnelMethodPasswordAuthentication,
		}
	}
	configuration := shared.DestinationMongodb{
		AuthType:        authType,
		Database:        database,
		DestinationType: destinationType,
		InstanceType:    instanceType,
		TunnelMethod:    tunnelMethod,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMongodbCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMongodbResourceModel) ToGetSDKType() *shared.DestinationMongodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMongodbResourceModel) ToUpdateSDKType() *shared.DestinationMongodbPutRequest {
	var authType shared.DestinationMongodbUpdateAuthorizationType
	var destinationMongodbUpdateAuthorizationTypeNone *shared.DestinationMongodbUpdateAuthorizationTypeNone
	if r.Configuration.AuthType.DestinationMongodbAuthorizationTypeLoginPassword != nil {
		authorization := shared.DestinationMongodbUpdateAuthorizationTypeNoneAuthorization(r.Configuration.AuthType.DestinationMongodbAuthorizationTypeLoginPassword.Authorization.ValueString())
		destinationMongodbUpdateAuthorizationTypeNone = &shared.DestinationMongodbUpdateAuthorizationTypeNone{
			Authorization: authorization,
		}
	}
	if destinationMongodbUpdateAuthorizationTypeNone != nil {
		authType = shared.DestinationMongodbUpdateAuthorizationType{
			DestinationMongodbUpdateAuthorizationTypeNone: destinationMongodbUpdateAuthorizationTypeNone,
		}
	}
	var destinationMongodbUpdateAuthorizationTypeLoginPassword *shared.DestinationMongodbUpdateAuthorizationTypeLoginPassword
	if r.Configuration.AuthType.DestinationMongodbAuthorizationTypeNone != nil {
		authorization1 := shared.DestinationMongodbUpdateAuthorizationTypeLoginPasswordAuthorization(r.Configuration.AuthType.DestinationMongodbAuthorizationTypeNone.Authorization.ValueString())
		destinationMongodbUpdateAuthorizationTypeLoginPassword = &shared.DestinationMongodbUpdateAuthorizationTypeLoginPassword{
			Authorization: authorization1,
		}
	}
	if destinationMongodbUpdateAuthorizationTypeLoginPassword != nil {
		authType = shared.DestinationMongodbUpdateAuthorizationType{
			DestinationMongodbUpdateAuthorizationTypeLoginPassword: destinationMongodbUpdateAuthorizationTypeLoginPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	var instanceType *shared.DestinationMongodbUpdateMongoDbInstanceType
	var destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance *shared.DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance
	if r.Configuration.InstanceType.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas != nil {
		instance := shared.DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstanceInstance(r.Configuration.InstanceType.DestinationMongodbMongoDBInstanceTypeMongoDBAtlas.Instance.ValueString())
		destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance = &shared.DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance{
			Instance: instance,
		}
	}
	if destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		instanceType = &shared.DestinationMongodbUpdateMongoDbInstanceType{
			DestinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance: destinationMongodbUpdateMongoDbInstanceTypeStandaloneMongoDbInstance,
		}
	}
	var destinationMongodbUpdateMongoDbInstanceTypeReplicaSet *shared.DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet
	if r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet != nil {
		instance1 := shared.DestinationMongodbUpdateMongoDbInstanceTypeReplicaSetInstance(r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.Instance.ValueString())
		replicaSet := new(string)
		if !r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.IsUnknown() && !r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.IsNull() {
			*replicaSet = r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ReplicaSet.ValueString()
		} else {
			replicaSet = nil
		}
		serverAddresses := r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeReplicaSet.ServerAddresses.ValueString()
		destinationMongodbUpdateMongoDbInstanceTypeReplicaSet = &shared.DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet{
			Instance:        instance1,
			ReplicaSet:      replicaSet,
			ServerAddresses: serverAddresses,
		}
	}
	if destinationMongodbUpdateMongoDbInstanceTypeReplicaSet != nil {
		instanceType = &shared.DestinationMongodbUpdateMongoDbInstanceType{
			DestinationMongodbUpdateMongoDbInstanceTypeReplicaSet: destinationMongodbUpdateMongoDbInstanceTypeReplicaSet,
		}
	}
	var destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas *shared.DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas
	if r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance != nil {
		instance2 := shared.DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlasInstance(r.Configuration.InstanceType.DestinationMongodbMongoDbInstanceTypeStandaloneMongoDbInstance.Instance.ValueString())
		destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas = &shared.DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas{
			Instance: instance2,
		}
	}
	if destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas != nil {
		instanceType = &shared.DestinationMongodbUpdateMongoDbInstanceType{
			DestinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas: destinationMongodbUpdateMongoDBInstanceTypeMongoDBAtlas,
		}
	}
	var tunnelMethod *shared.DestinationMongodbUpdateSSHTunnelMethod
	var destinationMongodbUpdateSSHTunnelMethodNoTunnel *shared.DestinationMongodbUpdateSSHTunnelMethodNoTunnel
	if r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodNoTunnel != nil {
		tunnelMethod1 := shared.DestinationMongodbUpdateSSHTunnelMethodNoTunnelTunnelMethod(r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodNoTunnel.TunnelMethod.ValueString())
		destinationMongodbUpdateSSHTunnelMethodNoTunnel = &shared.DestinationMongodbUpdateSSHTunnelMethodNoTunnel{
			TunnelMethod: tunnelMethod1,
		}
	}
	if destinationMongodbUpdateSSHTunnelMethodNoTunnel != nil {
		tunnelMethod = &shared.DestinationMongodbUpdateSSHTunnelMethod{
			DestinationMongodbUpdateSSHTunnelMethodNoTunnel: destinationMongodbUpdateSSHTunnelMethodNoTunnel,
		}
	}
	var destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication *shared.DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication
	if r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication != nil {
		tunnelHost := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelHost.ValueString()
		tunnelMethod2 := shared.DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelMethod.ValueString())
		tunnelPort := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelPort.ValueInt64()
		tunnelUser := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodPasswordAuthentication.TunnelUser.ValueString()
		destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication = &shared.DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication{
			TunnelHost:   tunnelHost,
			TunnelMethod: tunnelMethod2,
			TunnelPort:   tunnelPort,
			TunnelUser:   tunnelUser,
		}
	}
	if destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelMethod = &shared.DestinationMongodbUpdateSSHTunnelMethod{
			DestinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication: destinationMongodbUpdateSSHTunnelMethodSSHKeyAuthentication,
		}
	}
	var destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication *shared.DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication
	if r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication != nil {
		tunnelHost1 := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelHost.ValueString()
		tunnelMethod3 := shared.DestinationMongodbUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelMethod.ValueString())
		tunnelPort1 := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelPort.ValueInt64()
		tunnelUser1 := r.Configuration.TunnelMethod.DestinationMongodbSSHTunnelMethodSSHKeyAuthentication.TunnelUser.ValueString()
		destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication = &shared.DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication{
			TunnelHost:   tunnelHost1,
			TunnelMethod: tunnelMethod3,
			TunnelPort:   tunnelPort1,
			TunnelUser:   tunnelUser1,
		}
	}
	if destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication != nil {
		tunnelMethod = &shared.DestinationMongodbUpdateSSHTunnelMethod{
			DestinationMongodbUpdateSSHTunnelMethodPasswordAuthentication: destinationMongodbUpdateSSHTunnelMethodPasswordAuthentication,
		}
	}
	configuration := shared.DestinationMongodbUpdate{
		AuthType:     authType,
		Database:     database,
		InstanceType: instanceType,
		TunnelMethod: tunnelMethod,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMongodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMongodbResourceModel) ToDeleteSDKType() *shared.DestinationMongodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMongodbResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationMongodbResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
