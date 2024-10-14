// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SourceRedshiftUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Host Endpoint of the Redshift Cluster (must include the cluster-id, region and end with .redshift.amazonaws.com).
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password string `json:"password"`
	// Port of the database.
	Port *int64 `default:"5439" json:"port"`
	// The list of schemas to sync from. Specify one or more explicitly or keep empty to process all schemas. Schema names are case sensitive.
	Schemas []string `json:"schemas,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (s SourceRedshiftUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRedshiftUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRedshiftUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceRedshiftUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceRedshiftUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceRedshiftUpdate) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceRedshiftUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceRedshiftUpdate) GetSchemas() []string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *SourceRedshiftUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
