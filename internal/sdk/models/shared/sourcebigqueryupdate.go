// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type SourceBigqueryUpdate struct {
	// The contents of your Service Account Key JSON file. See the <a href="https://docs.airbyte.com/integrations/sources/bigquery#setup-the-bigquery-source-in-airbyte">docs</a> for more information on how to obtain this key.
	CredentialsJSON string `json:"credentials_json"`
	// The dataset ID to search for tables and views. If you are only loading data from one dataset, setting this option could result in much faster schema discovery.
	DatasetID *string `json:"dataset_id,omitempty"`
	// The GCP project ID for the project containing the target BigQuery dataset.
	ProjectID string `json:"project_id"`
}

func (o *SourceBigqueryUpdate) GetCredentialsJSON() string {
	if o == nil {
		return ""
	}
	return o.CredentialsJSON
}

func (o *SourceBigqueryUpdate) GetDatasetID() *string {
	if o == nil {
		return nil
	}
	return o.DatasetID
}

func (o *SourceBigqueryUpdate) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}
