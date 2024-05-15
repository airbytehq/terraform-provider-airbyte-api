// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"time"
)

// SourceChargebeeProductCatalog - Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section. If left blank, the product catalog version will be set to 2.0.
type SourceChargebeeProductCatalog string

const (
	SourceChargebeeProductCatalogOne0 SourceChargebeeProductCatalog = "1.0"
	SourceChargebeeProductCatalogTwo0 SourceChargebeeProductCatalog = "2.0"
)

func (e SourceChargebeeProductCatalog) ToPointer() *SourceChargebeeProductCatalog {
	return &e
}
func (e *SourceChargebeeProductCatalog) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "1.0":
		fallthrough
	case "2.0":
		*e = SourceChargebeeProductCatalog(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceChargebeeProductCatalog: %v", v)
	}
}

type Chargebee string

const (
	ChargebeeChargebee Chargebee = "chargebee"
)

func (e Chargebee) ToPointer() *Chargebee {
	return &e
}
func (e *Chargebee) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chargebee":
		*e = Chargebee(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Chargebee: %v", v)
	}
}

type SourceChargebee struct {
	// Product Catalog version of your Chargebee site. Instructions on how to find your version you may find <a href="https://apidocs.chargebee.com/docs/api?prod_cat_ver=2">here</a> under `API Version` section. If left blank, the product catalog version will be set to 2.0.
	ProductCatalog *SourceChargebeeProductCatalog `default:"2.0" json:"product_catalog"`
	// The site prefix for your Chargebee instance.
	Site string `json:"site"`
	// Chargebee API Key. See the <a href="https://docs.airbyte.com/integrations/sources/chargebee">docs</a> for more information on how to obtain this key.
	SiteAPIKey string    `json:"site_api_key"`
	sourceType Chargebee `const:"chargebee" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00.000Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceChargebee) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceChargebee) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceChargebee) GetProductCatalog() *SourceChargebeeProductCatalog {
	if o == nil {
		return nil
	}
	return o.ProductCatalog
}

func (o *SourceChargebee) GetSite() string {
	if o == nil {
		return ""
	}
	return o.Site
}

func (o *SourceChargebee) GetSiteAPIKey() string {
	if o == nil {
		return ""
	}
	return o.SiteAPIKey
}

func (o *SourceChargebee) GetSourceType() Chargebee {
	return ChargebeeChargebee
}

func (o *SourceChargebee) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}