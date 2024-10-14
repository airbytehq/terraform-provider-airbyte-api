// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

// DestinationGoogleSheetsAuthenticationViaGoogleOAuth - Google API Credentials for connecting to Google Sheets and Google Drive APIs
type DestinationGoogleSheetsAuthenticationViaGoogleOAuth struct {
	// The Client ID of your Google Sheets developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Google Sheets developer application.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining new access token.
	RefreshToken string `json:"refresh_token"`
}

func (o *DestinationGoogleSheetsAuthenticationViaGoogleOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *DestinationGoogleSheetsAuthenticationViaGoogleOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *DestinationGoogleSheetsAuthenticationViaGoogleOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type GoogleSheets string

const (
	GoogleSheetsGoogleSheets GoogleSheets = "google-sheets"
)

func (e GoogleSheets) ToPointer() *GoogleSheets {
	return &e
}
func (e *GoogleSheets) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-sheets":
		*e = GoogleSheets(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GoogleSheets: %v", v)
	}
}

type DestinationGoogleSheets struct {
	// Google API Credentials for connecting to Google Sheets and Google Drive APIs
	Credentials     DestinationGoogleSheetsAuthenticationViaGoogleOAuth `json:"credentials"`
	destinationType GoogleSheets                                        `const:"google-sheets" json:"destinationType"`
	// The link to your spreadsheet. See <a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'>this guide</a> for more details.
	SpreadsheetID string `json:"spreadsheet_id"`
}

func (d DestinationGoogleSheets) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationGoogleSheets) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationGoogleSheets) GetCredentials() DestinationGoogleSheetsAuthenticationViaGoogleOAuth {
	if o == nil {
		return DestinationGoogleSheetsAuthenticationViaGoogleOAuth{}
	}
	return o.Credentials
}

func (o *DestinationGoogleSheets) GetDestinationType() GoogleSheets {
	return GoogleSheetsGoogleSheets
}

func (o *DestinationGoogleSheets) GetSpreadsheetID() string {
	if o == nil {
		return ""
	}
	return o.SpreadsheetID
}
