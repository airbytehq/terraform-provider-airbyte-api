// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type ReadChangesUsingWriteAheadLogCDC struct {
	AdditionalProperties             types.String `tfsdk:"additional_properties"`
	HeartbeatActionQuery             types.String `tfsdk:"heartbeat_action_query"`
	InitialLoadTimeoutHours          types.Int64  `tfsdk:"initial_load_timeout_hours"`
	InitialWaitingSeconds            types.Int64  `tfsdk:"initial_waiting_seconds"`
	InvalidCdcCursorPositionBehavior types.String `tfsdk:"invalid_cdc_cursor_position_behavior"`
	LsnCommitBehaviour               types.String `tfsdk:"lsn_commit_behaviour"`
	Plugin                           types.String `tfsdk:"plugin"`
	Publication                      types.String `tfsdk:"publication"`
	QueueSize                        types.Int64  `tfsdk:"queue_size"`
	ReplicationSlot                  types.String `tfsdk:"replication_slot"`
}
