// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceMongodbV2UpdateSchemasClusterType string

const (
	SourceMongodbV2UpdateSchemasClusterTypeSelfManagedReplicaSet SourceMongodbV2UpdateSchemasClusterType = "SELF_MANAGED_REPLICA_SET"
)

func (e SourceMongodbV2UpdateSchemasClusterType) ToPointer() *SourceMongodbV2UpdateSchemasClusterType {
	return &e
}

func (e *SourceMongodbV2UpdateSchemasClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SELF_MANAGED_REPLICA_SET":
		*e = SourceMongodbV2UpdateSchemasClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2UpdateSchemasClusterType: %v", v)
	}
}

// SelfManagedReplicaSet - MongoDB self-hosted cluster configured as a replica set
type SelfManagedReplicaSet struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The authentication source where the user information is stored.
	AuthSource  *string                                 `default:"admin" json:"auth_source"`
	clusterType SourceMongodbV2UpdateSchemasClusterType `const:"SELF_MANAGED_REPLICA_SET" json:"cluster_type"`
	// The connection string of the cluster that you want to replicate.  https://www.mongodb.com/docs/manual/reference/connection-string/#find-your-self-hosted-deployment-s-connection-string for more information.
	ConnectionString string `json:"connection_string"`
	// The name of the MongoDB database that contains the collection(s) to replicate.
	Database string `json:"database"`
	// The password associated with this username.
	Password *string `json:"password,omitempty"`
	// When enabled, syncs will validate and structure records against the stream's schema.
	SchemaEnforced *bool `default:"true" json:"schema_enforced"`
	// The username which is used to access the database.
	Username *string `json:"username,omitempty"`
}

func (s SelfManagedReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SelfManagedReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SelfManagedReplicaSet) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SelfManagedReplicaSet) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *SelfManagedReplicaSet) GetClusterType() SourceMongodbV2UpdateSchemasClusterType {
	return SourceMongodbV2UpdateSchemasClusterTypeSelfManagedReplicaSet
}

func (o *SelfManagedReplicaSet) GetConnectionString() string {
	if o == nil {
		return ""
	}
	return o.ConnectionString
}

func (o *SelfManagedReplicaSet) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SelfManagedReplicaSet) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SelfManagedReplicaSet) GetSchemaEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaEnforced
}

func (o *SelfManagedReplicaSet) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type SourceMongodbV2UpdateClusterType string

const (
	SourceMongodbV2UpdateClusterTypeAtlasReplicaSet SourceMongodbV2UpdateClusterType = "ATLAS_REPLICA_SET"
)

func (e SourceMongodbV2UpdateClusterType) ToPointer() *SourceMongodbV2UpdateClusterType {
	return &e
}

func (e *SourceMongodbV2UpdateClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ATLAS_REPLICA_SET":
		*e = SourceMongodbV2UpdateClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2UpdateClusterType: %v", v)
	}
}

// MongoDBAtlasReplicaSet - MongoDB Atlas-hosted cluster configured as a replica set
type MongoDBAtlasReplicaSet struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The authentication source where the user information is stored.  See https://www.mongodb.com/docs/manual/reference/connection-string/#mongodb-urioption-urioption.authSource for more details.
	AuthSource  *string                          `default:"admin" json:"auth_source"`
	clusterType SourceMongodbV2UpdateClusterType `const:"ATLAS_REPLICA_SET" json:"cluster_type"`
	// The connection string of the cluster that you want to replicate.
	ConnectionString string `json:"connection_string"`
	// The name of the MongoDB database that contains the collection(s) to replicate.
	Database string `json:"database"`
	// The password associated with this username.
	Password string `json:"password"`
	// When enabled, syncs will validate and structure records against the stream's schema.
	SchemaEnforced *bool `default:"true" json:"schema_enforced"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (m MongoDBAtlasReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MongoDBAtlasReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MongoDBAtlasReplicaSet) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *MongoDBAtlasReplicaSet) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *MongoDBAtlasReplicaSet) GetClusterType() SourceMongodbV2UpdateClusterType {
	return SourceMongodbV2UpdateClusterTypeAtlasReplicaSet
}

func (o *MongoDBAtlasReplicaSet) GetConnectionString() string {
	if o == nil {
		return ""
	}
	return o.ConnectionString
}

func (o *MongoDBAtlasReplicaSet) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *MongoDBAtlasReplicaSet) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *MongoDBAtlasReplicaSet) GetSchemaEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaEnforced
}

func (o *MongoDBAtlasReplicaSet) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type ClusterTypeType string

const (
	ClusterTypeTypeMongoDBAtlasReplicaSet ClusterTypeType = "MongoDB Atlas Replica Set"
	ClusterTypeTypeSelfManagedReplicaSet  ClusterTypeType = "Self-Managed Replica Set"
)

// ClusterType - Configures the MongoDB cluster type.
type ClusterType struct {
	MongoDBAtlasReplicaSet *MongoDBAtlasReplicaSet
	SelfManagedReplicaSet  *SelfManagedReplicaSet

	Type ClusterTypeType
}

func CreateClusterTypeMongoDBAtlasReplicaSet(mongoDBAtlasReplicaSet MongoDBAtlasReplicaSet) ClusterType {
	typ := ClusterTypeTypeMongoDBAtlasReplicaSet

	return ClusterType{
		MongoDBAtlasReplicaSet: &mongoDBAtlasReplicaSet,
		Type:                   typ,
	}
}

func CreateClusterTypeSelfManagedReplicaSet(selfManagedReplicaSet SelfManagedReplicaSet) ClusterType {
	typ := ClusterTypeTypeSelfManagedReplicaSet

	return ClusterType{
		SelfManagedReplicaSet: &selfManagedReplicaSet,
		Type:                  typ,
	}
}

func (u *ClusterType) UnmarshalJSON(data []byte) error {

	mongoDBAtlasReplicaSet := new(MongoDBAtlasReplicaSet)
	if err := utils.UnmarshalJSON(data, &mongoDBAtlasReplicaSet, "", true, true); err == nil {
		u.MongoDBAtlasReplicaSet = mongoDBAtlasReplicaSet
		u.Type = ClusterTypeTypeMongoDBAtlasReplicaSet
		return nil
	}

	selfManagedReplicaSet := new(SelfManagedReplicaSet)
	if err := utils.UnmarshalJSON(data, &selfManagedReplicaSet, "", true, true); err == nil {
		u.SelfManagedReplicaSet = selfManagedReplicaSet
		u.Type = ClusterTypeTypeSelfManagedReplicaSet
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ClusterType) MarshalJSON() ([]byte, error) {
	if u.MongoDBAtlasReplicaSet != nil {
		return utils.MarshalJSON(u.MongoDBAtlasReplicaSet, "", true)
	}

	if u.SelfManagedReplicaSet != nil {
		return utils.MarshalJSON(u.SelfManagedReplicaSet, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMongodbV2Update struct {
	// Configures the MongoDB cluster type.
	DatabaseConfig ClusterType `json:"database_config"`
	// The maximum number of documents to sample when attempting to discover the unique fields for a collection.
	DiscoverSampleSize *int64 `default:"10000" json:"discover_sample_size"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds.
	InitialWaitingSeconds *int64 `default:"300" json:"initial_waiting_seconds"`
	// The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.
	QueueSize *int64 `default:"10000" json:"queue_size"`
}

func (s SourceMongodbV2Update) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbV2Update) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbV2Update) GetDatabaseConfig() ClusterType {
	if o == nil {
		return ClusterType{}
	}
	return o.DatabaseConfig
}

func (o *SourceMongodbV2Update) GetDiscoverSampleSize() *int64 {
	if o == nil {
		return nil
	}
	return o.DiscoverSampleSize
}

func (o *SourceMongodbV2Update) GetInitialWaitingSeconds() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialWaitingSeconds
}

func (o *SourceMongodbV2Update) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}
