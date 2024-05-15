// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DestinationResponse - Provides details of a single destination.
type DestinationResponse struct {
	// The values required to configure the destination.
	Configuration   any    `json:"configuration"`
	DestinationID   string `json:"destinationId"`
	DestinationType string `json:"destinationType"`
	Name            string `json:"name"`
	WorkspaceID     string `json:"workspaceId"`
}

func (o *DestinationResponse) GetConfiguration() any {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *DestinationResponse) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *DestinationResponse) GetDestinationType() string {
	if o == nil {
		return ""
	}
	return o.DestinationType
}

func (o *DestinationResponse) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationResponse) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}