// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceMongodbV2SchemasDatabaseConfigClusterType string

const (
	SourceMongodbV2SchemasDatabaseConfigClusterTypeSelfManagedReplicaSet SourceMongodbV2SchemasDatabaseConfigClusterType = "SELF_MANAGED_REPLICA_SET"
)

func (e SourceMongodbV2SchemasDatabaseConfigClusterType) ToPointer() *SourceMongodbV2SchemasDatabaseConfigClusterType {
	return &e
}
func (e *SourceMongodbV2SchemasDatabaseConfigClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SELF_MANAGED_REPLICA_SET":
		*e = SourceMongodbV2SchemasDatabaseConfigClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2SchemasDatabaseConfigClusterType: %v", v)
	}
}

// SourceMongodbV2SelfManagedReplicaSet - MongoDB self-hosted cluster configured as a replica set
type SourceMongodbV2SelfManagedReplicaSet struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The authentication source where the user information is stored.
	AuthSource  *string                                         `default:"admin" json:"auth_source"`
	clusterType SourceMongodbV2SchemasDatabaseConfigClusterType `const:"SELF_MANAGED_REPLICA_SET" json:"cluster_type"`
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

func (s SourceMongodbV2SelfManagedReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbV2SelfManagedReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetClusterType() SourceMongodbV2SchemasDatabaseConfigClusterType {
	return SourceMongodbV2SchemasDatabaseConfigClusterTypeSelfManagedReplicaSet
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetConnectionString() string {
	if o == nil {
		return ""
	}
	return o.ConnectionString
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetSchemaEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaEnforced
}

func (o *SourceMongodbV2SelfManagedReplicaSet) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type SourceMongodbV2SchemasClusterType string

const (
	SourceMongodbV2SchemasClusterTypeAtlasReplicaSet SourceMongodbV2SchemasClusterType = "ATLAS_REPLICA_SET"
)

func (e SourceMongodbV2SchemasClusterType) ToPointer() *SourceMongodbV2SchemasClusterType {
	return &e
}
func (e *SourceMongodbV2SchemasClusterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ATLAS_REPLICA_SET":
		*e = SourceMongodbV2SchemasClusterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2SchemasClusterType: %v", v)
	}
}

// SourceMongodbV2MongoDBAtlasReplicaSet - MongoDB Atlas-hosted cluster configured as a replica set
type SourceMongodbV2MongoDBAtlasReplicaSet struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// The authentication source where the user information is stored.  See https://www.mongodb.com/docs/manual/reference/connection-string/#mongodb-urioption-urioption.authSource for more details.
	AuthSource  *string                           `default:"admin" json:"auth_source"`
	clusterType SourceMongodbV2SchemasClusterType `const:"ATLAS_REPLICA_SET" json:"cluster_type"`
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

func (s SourceMongodbV2MongoDBAtlasReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbV2MongoDBAtlasReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetClusterType() SourceMongodbV2SchemasClusterType {
	return SourceMongodbV2SchemasClusterTypeAtlasReplicaSet
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetConnectionString() string {
	if o == nil {
		return ""
	}
	return o.ConnectionString
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetSchemaEnforced() *bool {
	if o == nil {
		return nil
	}
	return o.SchemaEnforced
}

func (o *SourceMongodbV2MongoDBAtlasReplicaSet) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type SourceMongodbV2ClusterTypeType string

const (
	SourceMongodbV2ClusterTypeTypeSourceMongodbV2MongoDBAtlasReplicaSet SourceMongodbV2ClusterTypeType = "source-mongodb-v2_MongoDB Atlas Replica Set"
	SourceMongodbV2ClusterTypeTypeSourceMongodbV2SelfManagedReplicaSet  SourceMongodbV2ClusterTypeType = "source-mongodb-v2_Self-Managed Replica Set"
)

// SourceMongodbV2ClusterType - Configures the MongoDB cluster type.
type SourceMongodbV2ClusterType struct {
	SourceMongodbV2MongoDBAtlasReplicaSet *SourceMongodbV2MongoDBAtlasReplicaSet
	SourceMongodbV2SelfManagedReplicaSet  *SourceMongodbV2SelfManagedReplicaSet

	Type SourceMongodbV2ClusterTypeType
}

func CreateSourceMongodbV2ClusterTypeSourceMongodbV2MongoDBAtlasReplicaSet(sourceMongodbV2MongoDBAtlasReplicaSet SourceMongodbV2MongoDBAtlasReplicaSet) SourceMongodbV2ClusterType {
	typ := SourceMongodbV2ClusterTypeTypeSourceMongodbV2MongoDBAtlasReplicaSet

	return SourceMongodbV2ClusterType{
		SourceMongodbV2MongoDBAtlasReplicaSet: &sourceMongodbV2MongoDBAtlasReplicaSet,
		Type:                                  typ,
	}
}

func CreateSourceMongodbV2ClusterTypeSourceMongodbV2SelfManagedReplicaSet(sourceMongodbV2SelfManagedReplicaSet SourceMongodbV2SelfManagedReplicaSet) SourceMongodbV2ClusterType {
	typ := SourceMongodbV2ClusterTypeTypeSourceMongodbV2SelfManagedReplicaSet

	return SourceMongodbV2ClusterType{
		SourceMongodbV2SelfManagedReplicaSet: &sourceMongodbV2SelfManagedReplicaSet,
		Type:                                 typ,
	}
}

func (u *SourceMongodbV2ClusterType) UnmarshalJSON(data []byte) error {

	var sourceMongodbV2MongoDBAtlasReplicaSet SourceMongodbV2MongoDBAtlasReplicaSet = SourceMongodbV2MongoDBAtlasReplicaSet{}
	if err := utils.UnmarshalJSON(data, &sourceMongodbV2MongoDBAtlasReplicaSet, "", true, true); err == nil {
		u.SourceMongodbV2MongoDBAtlasReplicaSet = &sourceMongodbV2MongoDBAtlasReplicaSet
		u.Type = SourceMongodbV2ClusterTypeTypeSourceMongodbV2MongoDBAtlasReplicaSet
		return nil
	}

	var sourceMongodbV2SelfManagedReplicaSet SourceMongodbV2SelfManagedReplicaSet = SourceMongodbV2SelfManagedReplicaSet{}
	if err := utils.UnmarshalJSON(data, &sourceMongodbV2SelfManagedReplicaSet, "", true, true); err == nil {
		u.SourceMongodbV2SelfManagedReplicaSet = &sourceMongodbV2SelfManagedReplicaSet
		u.Type = SourceMongodbV2ClusterTypeTypeSourceMongodbV2SelfManagedReplicaSet
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SourceMongodbV2ClusterType", string(data))
}

func (u SourceMongodbV2ClusterType) MarshalJSON() ([]byte, error) {
	if u.SourceMongodbV2MongoDBAtlasReplicaSet != nil {
		return utils.MarshalJSON(u.SourceMongodbV2MongoDBAtlasReplicaSet, "", true)
	}

	if u.SourceMongodbV2SelfManagedReplicaSet != nil {
		return utils.MarshalJSON(u.SourceMongodbV2SelfManagedReplicaSet, "", true)
	}

	return nil, errors.New("could not marshal union type SourceMongodbV2ClusterType: all fields are null")
}

// SourceMongodbV2InvalidCDCPositionBehaviorAdvanced - Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss.
type SourceMongodbV2InvalidCDCPositionBehaviorAdvanced string

const (
	SourceMongodbV2InvalidCDCPositionBehaviorAdvancedFailSync   SourceMongodbV2InvalidCDCPositionBehaviorAdvanced = "Fail sync"
	SourceMongodbV2InvalidCDCPositionBehaviorAdvancedReSyncData SourceMongodbV2InvalidCDCPositionBehaviorAdvanced = "Re-sync data"
)

func (e SourceMongodbV2InvalidCDCPositionBehaviorAdvanced) ToPointer() *SourceMongodbV2InvalidCDCPositionBehaviorAdvanced {
	return &e
}
func (e *SourceMongodbV2InvalidCDCPositionBehaviorAdvanced) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Fail sync":
		fallthrough
	case "Re-sync data":
		*e = SourceMongodbV2InvalidCDCPositionBehaviorAdvanced(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2InvalidCDCPositionBehaviorAdvanced: %v", v)
	}
}

type MongodbV2 string

const (
	MongodbV2MongodbV2 MongodbV2 = "mongodb-v2"
)

func (e MongodbV2) ToPointer() *MongodbV2 {
	return &e
}
func (e *MongodbV2) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mongodb-v2":
		*e = MongodbV2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MongodbV2: %v", v)
	}
}

// SourceMongodbV2CaptureModeAdvanced - Determines how Airbyte looks up the value of an updated document. If 'Lookup' is chosen, the current value of the document will be read. If 'Post Image' is chosen, then the version of the document immediately after an update will be read. WARNING : Severe data loss will occur if this option is chosen and the appropriate settings are not set on your Mongo instance : https://www.mongodb.com/docs/manual/changeStreams/#change-streams-with-document-pre-and-post-images.
type SourceMongodbV2CaptureModeAdvanced string

const (
	SourceMongodbV2CaptureModeAdvancedLookup    SourceMongodbV2CaptureModeAdvanced = "Lookup"
	SourceMongodbV2CaptureModeAdvancedPostImage SourceMongodbV2CaptureModeAdvanced = "Post Image"
)

func (e SourceMongodbV2CaptureModeAdvanced) ToPointer() *SourceMongodbV2CaptureModeAdvanced {
	return &e
}
func (e *SourceMongodbV2CaptureModeAdvanced) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Lookup":
		fallthrough
	case "Post Image":
		*e = SourceMongodbV2CaptureModeAdvanced(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbV2CaptureModeAdvanced: %v", v)
	}
}

type SourceMongodbV2 struct {
	// Configures the MongoDB cluster type.
	DatabaseConfig SourceMongodbV2ClusterType `json:"database_config"`
	// The maximum number of documents to sample when attempting to discover the unique fields for a collection.
	DiscoverSampleSize *int64 `default:"10000" json:"discover_sample_size"`
	// The amount of time an initial load is allowed to continue for before catching up on CDC logs.
	InitialLoadTimeoutHours *int64 `default:"8" json:"initial_load_timeout_hours"`
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds.
	InitialWaitingSeconds *int64 `default:"300" json:"initial_waiting_seconds"`
	// Determines whether Airbyte should fail or re-sync data in case of an stale/invalid cursor value into the WAL. If 'Fail sync' is chosen, a user will have to manually reset the connection before being able to continue syncing data. If 'Re-sync data' is chosen, Airbyte will automatically trigger a refresh but could lead to higher cloud costs and data loss.
	InvalidCdcCursorPositionBehavior *SourceMongodbV2InvalidCDCPositionBehaviorAdvanced `default:"Fail sync" json:"invalid_cdc_cursor_position_behavior"`
	// The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.
	QueueSize  *int64    `default:"10000" json:"queue_size"`
	sourceType MongodbV2 `const:"mongodb-v2" json:"sourceType"`
	// Determines how Airbyte looks up the value of an updated document. If 'Lookup' is chosen, the current value of the document will be read. If 'Post Image' is chosen, then the version of the document immediately after an update will be read. WARNING : Severe data loss will occur if this option is chosen and the appropriate settings are not set on your Mongo instance : https://www.mongodb.com/docs/manual/changeStreams/#change-streams-with-document-pre-and-post-images.
	UpdateCaptureMode *SourceMongodbV2CaptureModeAdvanced `default:"Lookup" json:"update_capture_mode"`
}

func (s SourceMongodbV2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbV2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbV2) GetDatabaseConfig() SourceMongodbV2ClusterType {
	if o == nil {
		return SourceMongodbV2ClusterType{}
	}
	return o.DatabaseConfig
}

func (o *SourceMongodbV2) GetDiscoverSampleSize() *int64 {
	if o == nil {
		return nil
	}
	return o.DiscoverSampleSize
}

func (o *SourceMongodbV2) GetInitialLoadTimeoutHours() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialLoadTimeoutHours
}

func (o *SourceMongodbV2) GetInitialWaitingSeconds() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialWaitingSeconds
}

func (o *SourceMongodbV2) GetInvalidCdcCursorPositionBehavior() *SourceMongodbV2InvalidCDCPositionBehaviorAdvanced {
	if o == nil {
		return nil
	}
	return o.InvalidCdcCursorPositionBehavior
}

func (o *SourceMongodbV2) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *SourceMongodbV2) GetSourceType() MongodbV2 {
	return MongodbV2MongodbV2
}

func (o *SourceMongodbV2) GetUpdateCaptureMode() *SourceMongodbV2CaptureModeAdvanced {
	if o == nil {
		return nil
	}
	return o.UpdateCaptureMode
}
