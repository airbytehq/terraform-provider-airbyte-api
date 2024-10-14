// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Faker string

const (
	FakerFaker Faker = "faker"
)

func (e Faker) ToPointer() *Faker {
	return &e
}
func (e *Faker) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "faker":
		*e = Faker(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Faker: %v", v)
	}
}

type SourceFaker struct {
	// Should the updated_at values for every record be new each sync?  Setting this to false will case the source to stop emitting records after COUNT records have been emitted.
	AlwaysUpdated *bool `default:"true" json:"always_updated"`
	// How many users should be generated in total. The purchases table will be scaled to match, with 10 purchases created per 10 users. This setting does not apply to the products stream.
	Count *int64 `default:"1000" json:"count"`
	// How many parallel workers should we use to generate fake data?  Choose a value equal to the number of CPUs you will allocate to this source.
	Parallelism *int64 `default:"4" json:"parallelism"`
	// How many fake records will be in each page (stream slice), before a state message is emitted?
	RecordsPerSlice *int64 `default:"1000" json:"records_per_slice"`
	// Manually control the faker random seed to return the same values on subsequent runs (leave -1 for random)
	Seed       *int64 `default:"-1" json:"seed"`
	sourceType Faker  `const:"faker" json:"sourceType"`
}

func (s SourceFaker) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFaker) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFaker) GetAlwaysUpdated() *bool {
	if o == nil {
		return nil
	}
	return o.AlwaysUpdated
}

func (o *SourceFaker) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *SourceFaker) GetParallelism() *int64 {
	if o == nil {
		return nil
	}
	return o.Parallelism
}

func (o *SourceFaker) GetRecordsPerSlice() *int64 {
	if o == nil {
		return nil
	}
	return o.RecordsPerSlice
}

func (o *SourceFaker) GetSeed() *int64 {
	if o == nil {
		return nil
	}
	return o.Seed
}

func (o *SourceFaker) GetSourceType() Faker {
	return FakerFaker
}
