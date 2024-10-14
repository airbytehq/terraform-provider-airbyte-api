// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ConnectionDataSourceModel) RefreshFromSharedConnectionResponse(resp *shared.ConnectionResponse) {
	if resp != nil {
		r.Configurations.Streams = []tfTypes.StreamConfiguration{}
		if len(r.Configurations.Streams) > len(resp.Configurations.Streams) {
			r.Configurations.Streams = r.Configurations.Streams[:len(resp.Configurations.Streams)]
		}
		for streamsCount, streamsItem := range resp.Configurations.Streams {
			var streams1 tfTypes.StreamConfiguration
			streams1.CursorField = []types.String{}
			for _, v := range streamsItem.CursorField {
				streams1.CursorField = append(streams1.CursorField, types.StringValue(v))
			}
			streams1.Name = types.StringValue(streamsItem.Name)
			streams1.PrimaryKey = nil
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				var primaryKey1 []types.String
				primaryKey1 = []types.String{}
				for _, v := range primaryKeyItem {
					primaryKey1 = append(primaryKey1, types.StringValue(v))
				}
				streams1.PrimaryKey = append(streams1.PrimaryKey, primaryKey1)
			}
			streams1.SelectedFields = []tfTypes.SelectedFieldInfo{}
			for selectedFieldsCount, selectedFieldsItem := range streamsItem.SelectedFields {
				var selectedFields1 tfTypes.SelectedFieldInfo
				selectedFields1.FieldPath = []types.String{}
				for _, v := range selectedFieldsItem.FieldPath {
					selectedFields1.FieldPath = append(selectedFields1.FieldPath, types.StringValue(v))
				}
				if selectedFieldsCount+1 > len(streams1.SelectedFields) {
					streams1.SelectedFields = append(streams1.SelectedFields, selectedFields1)
				} else {
					streams1.SelectedFields[selectedFieldsCount].FieldPath = selectedFields1.FieldPath
				}
			}
			if streamsItem.SyncMode != nil {
				streams1.SyncMode = types.StringValue(string(*streamsItem.SyncMode))
			} else {
				streams1.SyncMode = types.StringNull()
			}
			if streamsCount+1 > len(r.Configurations.Streams) {
				r.Configurations.Streams = append(r.Configurations.Streams, streams1)
			} else {
				r.Configurations.Streams[streamsCount].CursorField = streams1.CursorField
				r.Configurations.Streams[streamsCount].Name = streams1.Name
				r.Configurations.Streams[streamsCount].PrimaryKey = streams1.PrimaryKey
				r.Configurations.Streams[streamsCount].SelectedFields = streams1.SelectedFields
				r.Configurations.Streams[streamsCount].SyncMode = streams1.SyncMode
			}
		}
		r.ConnectionID = types.StringValue(resp.ConnectionID)
		if resp.DataResidency != nil {
			r.DataResidency = types.StringValue(string(*resp.DataResidency))
		} else {
			r.DataResidency = types.StringNull()
		}
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.Name = types.StringValue(resp.Name)
		if resp.NamespaceDefinition != nil {
			r.NamespaceDefinition = types.StringValue(string(*resp.NamespaceDefinition))
		} else {
			r.NamespaceDefinition = types.StringNull()
		}
		r.NamespaceFormat = types.StringPointerValue(resp.NamespaceFormat)
		if resp.NonBreakingSchemaUpdatesBehavior != nil {
			r.NonBreakingSchemaUpdatesBehavior = types.StringValue(string(*resp.NonBreakingSchemaUpdatesBehavior))
		} else {
			r.NonBreakingSchemaUpdatesBehavior = types.StringNull()
		}
		r.Prefix = types.StringPointerValue(resp.Prefix)
		r.Schedule.BasicTiming = types.StringPointerValue(resp.Schedule.BasicTiming)
		r.Schedule.CronExpression = types.StringPointerValue(resp.Schedule.CronExpression)
		r.Schedule.ScheduleType = types.StringValue(string(resp.Schedule.ScheduleType))
		r.SourceID = types.StringValue(resp.SourceID)
		r.Status = types.StringValue(string(resp.Status))
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}
