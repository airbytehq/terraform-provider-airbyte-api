// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Firestore string

const (
	FirestoreFirestore Firestore = "firestore"
)

func (e Firestore) ToPointer() *Firestore {
	return &e
}
func (e *Firestore) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "firestore":
		*e = Firestore(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Firestore: %v", v)
	}
}

type DestinationFirestore struct {
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/firestore">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.
	CredentialsJSON *string   `json:"credentials_json,omitempty"`
	destinationType Firestore `const:"firestore" json:"destinationType"`
	// The GCP project ID for the project containing the target BigQuery dataset.
	ProjectID string `json:"project_id"`
}

func (d DestinationFirestore) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationFirestore) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationFirestore) GetCredentialsJSON() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsJSON
}

func (o *DestinationFirestore) GetDestinationType() Firestore {
	return FirestoreFirestore
}

func (o *DestinationFirestore) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}
