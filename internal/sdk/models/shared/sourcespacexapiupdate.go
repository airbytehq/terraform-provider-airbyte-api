// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSpacexAPIUpdate struct {
	ID      *string `json:"id,omitempty"`
	Options *string `json:"options,omitempty"`
}

func (o *SourceSpacexAPIUpdate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SourceSpacexAPIUpdate) GetOptions() *string {
	if o == nil {
		return nil
	}
	return o.Options
}