// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Vectara string

const (
	VectaraVectara Vectara = "vectara"
)

func (e Vectara) ToPointer() *Vectara {
	return &e
}
func (e *Vectara) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vectara":
		*e = Vectara(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Vectara: %v", v)
	}
}

// DestinationVectaraOAuth20Credentials - OAuth2.0 credentials used to authenticate admin actions (creating/deleting corpora)
type DestinationVectaraOAuth20Credentials struct {
	// OAuth2.0 client id
	ClientID string `json:"client_id"`
	// OAuth2.0 client secret
	ClientSecret string `json:"client_secret"`
}

func (o *DestinationVectaraOAuth20Credentials) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *DestinationVectaraOAuth20Credentials) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

// DestinationVectara - Configuration to connect to the Vectara instance
type DestinationVectara struct {
	// The Name of Corpus to load data into
	CorpusName string `json:"corpus_name"`
	// Your customer id as it is in the authenticaion url
	CustomerID      string  `json:"customer_id"`
	destinationType Vectara `const:"vectara" json:"destinationType"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// OAuth2.0 credentials used to authenticate admin actions (creating/deleting corpora)
	Oauth2 DestinationVectaraOAuth20Credentials `json:"oauth2"`
	// Parallelize indexing into Vectara with multiple threads
	Parallelize *bool `default:"false" json:"parallelize"`
	// List of fields in the record that should be in the section of the document. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
	// A field that will be used to populate the `title` of each document. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TitleField *string `default:"" json:"title_field"`
}

func (d DestinationVectara) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVectara) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVectara) GetCorpusName() string {
	if o == nil {
		return ""
	}
	return o.CorpusName
}

func (o *DestinationVectara) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *DestinationVectara) GetDestinationType() Vectara {
	return VectaraVectara
}

func (o *DestinationVectara) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *DestinationVectara) GetOauth2() DestinationVectaraOAuth20Credentials {
	if o == nil {
		return DestinationVectaraOAuth20Credentials{}
	}
	return o.Oauth2
}

func (o *DestinationVectara) GetParallelize() *bool {
	if o == nil {
		return nil
	}
	return o.Parallelize
}

func (o *DestinationVectara) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

func (o *DestinationVectara) GetTitleField() *string {
	if o == nil {
		return nil
	}
	return o.TitleField
}
