// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

type SourceMysqlUpdateMethod struct {
	ReadChangesUsingBinaryLogCDC     *ReadChangesUsingBinaryLogCDC `tfsdk:"read_changes_using_binary_log_cdc" tfPlanOnly:"true"`
	ScanChangesWithUserDefinedCursor *Fake                         `tfsdk:"scan_changes_with_user_defined_cursor" tfPlanOnly:"true"`
}
