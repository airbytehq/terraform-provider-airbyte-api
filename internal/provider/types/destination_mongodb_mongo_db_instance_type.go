// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type DestinationMongodbMongoDbInstanceType struct {
	MongoDBAtlas              *MongoDBAtlas              `tfsdk:"mongo_db_atlas" tfPlanOnly:"true"`
	ReplicaSet                *ReplicaSet                `tfsdk:"replica_set" tfPlanOnly:"true"`
	StandaloneMongoDbInstance *StandaloneMongoDbInstance `tfsdk:"standalone_mongo_db_instance" tfPlanOnly:"true"`
}
