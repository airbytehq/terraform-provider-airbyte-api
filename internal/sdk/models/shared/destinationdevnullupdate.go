// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationDevNullUpdateSchemasTestDestinationTestDestinationType string

const (
	DestinationDevNullUpdateSchemasTestDestinationTestDestinationTypeFailing DestinationDevNullUpdateSchemasTestDestinationTestDestinationType = "FAILING"
)

func (e DestinationDevNullUpdateSchemasTestDestinationTestDestinationType) ToPointer() *DestinationDevNullUpdateSchemasTestDestinationTestDestinationType {
	return &e
}
func (e *DestinationDevNullUpdateSchemasTestDestinationTestDestinationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FAILING":
		*e = DestinationDevNullUpdateSchemasTestDestinationTestDestinationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullUpdateSchemasTestDestinationTestDestinationType: %v", v)
	}
}

type Failing struct {
	// Number of messages after which to fail.
	NumMessages         int64                                                              `json:"num_messages"`
	testDestinationType *DestinationDevNullUpdateSchemasTestDestinationTestDestinationType `const:"FAILING" json:"test_destination_type"`
}

func (f Failing) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *Failing) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Failing) GetNumMessages() int64 {
	if o == nil {
		return 0
	}
	return o.NumMessages
}

func (o *Failing) GetTestDestinationType() *DestinationDevNullUpdateSchemasTestDestinationTestDestinationType {
	return DestinationDevNullUpdateSchemasTestDestinationTestDestinationTypeFailing.ToPointer()
}

type DestinationDevNullUpdateSchemasTestDestinationType string

const (
	DestinationDevNullUpdateSchemasTestDestinationTypeThrottled DestinationDevNullUpdateSchemasTestDestinationType = "THROTTLED"
)

func (e DestinationDevNullUpdateSchemasTestDestinationType) ToPointer() *DestinationDevNullUpdateSchemasTestDestinationType {
	return &e
}
func (e *DestinationDevNullUpdateSchemasTestDestinationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "THROTTLED":
		*e = DestinationDevNullUpdateSchemasTestDestinationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullUpdateSchemasTestDestinationType: %v", v)
	}
}

type Throttled struct {
	// Number of milli-second to pause in between records.
	MillisPerRecord     int64                                               `json:"millis_per_record"`
	testDestinationType *DestinationDevNullUpdateSchemasTestDestinationType `const:"THROTTLED" json:"test_destination_type"`
}

func (t Throttled) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *Throttled) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Throttled) GetMillisPerRecord() int64 {
	if o == nil {
		return 0
	}
	return o.MillisPerRecord
}

func (o *Throttled) GetTestDestinationType() *DestinationDevNullUpdateSchemasTestDestinationType {
	return DestinationDevNullUpdateSchemasTestDestinationTypeThrottled.ToPointer()
}

type DestinationDevNullUpdateTestDestinationType string

const (
	DestinationDevNullUpdateTestDestinationTypeSilent DestinationDevNullUpdateTestDestinationType = "SILENT"
)

func (e DestinationDevNullUpdateTestDestinationType) ToPointer() *DestinationDevNullUpdateTestDestinationType {
	return &e
}
func (e *DestinationDevNullUpdateTestDestinationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SILENT":
		*e = DestinationDevNullUpdateTestDestinationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullUpdateTestDestinationType: %v", v)
	}
}

type Silent struct {
	testDestinationType *DestinationDevNullUpdateTestDestinationType `const:"SILENT" json:"test_destination_type"`
}

func (s Silent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Silent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Silent) GetTestDestinationType() *DestinationDevNullUpdateTestDestinationType {
	return DestinationDevNullUpdateTestDestinationTypeSilent.ToPointer()
}

type DestinationDevNullUpdateSchemasLoggingType string

const (
	DestinationDevNullUpdateSchemasLoggingTypeRandomSampling DestinationDevNullUpdateSchemasLoggingType = "RandomSampling"
)

func (e DestinationDevNullUpdateSchemasLoggingType) ToPointer() *DestinationDevNullUpdateSchemasLoggingType {
	return &e
}
func (e *DestinationDevNullUpdateSchemasLoggingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RandomSampling":
		*e = DestinationDevNullUpdateSchemasLoggingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullUpdateSchemasLoggingType: %v", v)
	}
}

// RandomSampling - For each stream, randomly log a percentage of the entries with a maximum cap.
type RandomSampling struct {
	LoggingType *DestinationDevNullUpdateSchemasLoggingType `default:"RandomSampling" json:"logging_type"`
	// Max number of entries to log. This destination is for testing only. So it won't make sense to log infinitely. The maximum is 1,000 entries.
	MaxEntryCount *float64 `default:"100" json:"max_entry_count"`
	// A positive floating number smaller than 1.
	SamplingRatio *float64 `default:"0.001" json:"sampling_ratio"`
	// When the seed is unspecified, the current time millis will be used as the seed.
	Seed *float64 `json:"seed,omitempty"`
}

func (r RandomSampling) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RandomSampling) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *RandomSampling) GetLoggingType() *DestinationDevNullUpdateSchemasLoggingType {
	if o == nil {
		return nil
	}
	return o.LoggingType
}

func (o *RandomSampling) GetMaxEntryCount() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxEntryCount
}

func (o *RandomSampling) GetSamplingRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.SamplingRatio
}

func (o *RandomSampling) GetSeed() *float64 {
	if o == nil {
		return nil
	}
	return o.Seed
}

type DestinationDevNullUpdateLoggingType string

const (
	DestinationDevNullUpdateLoggingTypeEveryNth DestinationDevNullUpdateLoggingType = "EveryNth"
)

func (e DestinationDevNullUpdateLoggingType) ToPointer() *DestinationDevNullUpdateLoggingType {
	return &e
}
func (e *DestinationDevNullUpdateLoggingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EveryNth":
		*e = DestinationDevNullUpdateLoggingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDevNullUpdateLoggingType: %v", v)
	}
}

// EveryNThEntry - For each stream, log every N-th entry with a maximum cap.
type EveryNThEntry struct {
	LoggingType *DestinationDevNullUpdateLoggingType `default:"EveryNth" json:"logging_type"`
	// Max number of entries to log. This destination is for testing only. So it won't make sense to log infinitely. The maximum is 1,000 entries.
	MaxEntryCount *float64 `default:"100" json:"max_entry_count"`
	// The N-th entry to log for each stream. N starts from 1. For example, when N = 1, every entry is logged; when N = 2, every other entry is logged; when N = 3, one out of three entries is logged.
	NthEntryToLog float64 `json:"nth_entry_to_log"`
}

func (e EveryNThEntry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EveryNThEntry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *EveryNThEntry) GetLoggingType() *DestinationDevNullUpdateLoggingType {
	if o == nil {
		return nil
	}
	return o.LoggingType
}

func (o *EveryNThEntry) GetMaxEntryCount() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxEntryCount
}

func (o *EveryNThEntry) GetNthEntryToLog() float64 {
	if o == nil {
		return 0.0
	}
	return o.NthEntryToLog
}

type LoggingType string

const (
	LoggingTypeFirstN LoggingType = "FirstN"
)

func (e LoggingType) ToPointer() *LoggingType {
	return &e
}
func (e *LoggingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FirstN":
		*e = LoggingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LoggingType: %v", v)
	}
}

// FirstNEntries - Log first N entries per stream.
type FirstNEntries struct {
	LoggingType *LoggingType `default:"FirstN" json:"logging_type"`
	// Number of entries to log. This destination is for testing only. So it won't make sense to log infinitely. The maximum is 1,000 entries.
	MaxEntryCount *float64 `default:"100" json:"max_entry_count"`
}

func (f FirstNEntries) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FirstNEntries) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *FirstNEntries) GetLoggingType() *LoggingType {
	if o == nil {
		return nil
	}
	return o.LoggingType
}

func (o *FirstNEntries) GetMaxEntryCount() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxEntryCount
}

type LoggingConfigurationType string

const (
	LoggingConfigurationTypeFirstNEntries  LoggingConfigurationType = "First N Entries"
	LoggingConfigurationTypeEveryNThEntry  LoggingConfigurationType = "Every N-th Entry"
	LoggingConfigurationTypeRandomSampling LoggingConfigurationType = "Random Sampling"
)

// LoggingConfiguration - Configurate how the messages are logged.
type LoggingConfiguration struct {
	FirstNEntries  *FirstNEntries
	EveryNThEntry  *EveryNThEntry
	RandomSampling *RandomSampling

	Type LoggingConfigurationType
}

func CreateLoggingConfigurationFirstNEntries(firstNEntries FirstNEntries) LoggingConfiguration {
	typ := LoggingConfigurationTypeFirstNEntries

	return LoggingConfiguration{
		FirstNEntries: &firstNEntries,
		Type:          typ,
	}
}

func CreateLoggingConfigurationEveryNThEntry(everyNThEntry EveryNThEntry) LoggingConfiguration {
	typ := LoggingConfigurationTypeEveryNThEntry

	return LoggingConfiguration{
		EveryNThEntry: &everyNThEntry,
		Type:          typ,
	}
}

func CreateLoggingConfigurationRandomSampling(randomSampling RandomSampling) LoggingConfiguration {
	typ := LoggingConfigurationTypeRandomSampling

	return LoggingConfiguration{
		RandomSampling: &randomSampling,
		Type:           typ,
	}
}

func (u *LoggingConfiguration) UnmarshalJSON(data []byte) error {

	var firstNEntries FirstNEntries = FirstNEntries{}
	if err := utils.UnmarshalJSON(data, &firstNEntries, "", true, true); err == nil {
		u.FirstNEntries = &firstNEntries
		u.Type = LoggingConfigurationTypeFirstNEntries
		return nil
	}

	var everyNThEntry EveryNThEntry = EveryNThEntry{}
	if err := utils.UnmarshalJSON(data, &everyNThEntry, "", true, true); err == nil {
		u.EveryNThEntry = &everyNThEntry
		u.Type = LoggingConfigurationTypeEveryNThEntry
		return nil
	}

	var randomSampling RandomSampling = RandomSampling{}
	if err := utils.UnmarshalJSON(data, &randomSampling, "", true, true); err == nil {
		u.RandomSampling = &randomSampling
		u.Type = LoggingConfigurationTypeRandomSampling
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for LoggingConfiguration", string(data))
}

func (u LoggingConfiguration) MarshalJSON() ([]byte, error) {
	if u.FirstNEntries != nil {
		return utils.MarshalJSON(u.FirstNEntries, "", true)
	}

	if u.EveryNThEntry != nil {
		return utils.MarshalJSON(u.EveryNThEntry, "", true)
	}

	if u.RandomSampling != nil {
		return utils.MarshalJSON(u.RandomSampling, "", true)
	}

	return nil, errors.New("could not marshal union type LoggingConfiguration: all fields are null")
}

type TestDestinationType string

const (
	TestDestinationTypeLogging TestDestinationType = "LOGGING"
)

func (e TestDestinationType) ToPointer() *TestDestinationType {
	return &e
}
func (e *TestDestinationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOGGING":
		*e = TestDestinationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TestDestinationType: %v", v)
	}
}

type Logging struct {
	// Configurate how the messages are logged.
	LoggingConfig       LoggingConfiguration `json:"logging_config"`
	testDestinationType *TestDestinationType `const:"LOGGING" json:"test_destination_type"`
}

func (l Logging) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Logging) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Logging) GetLoggingConfig() LoggingConfiguration {
	if o == nil {
		return LoggingConfiguration{}
	}
	return o.LoggingConfig
}

func (o *Logging) GetTestDestinationType() *TestDestinationType {
	return TestDestinationTypeLogging.ToPointer()
}

type TestDestinationUnionType string

const (
	TestDestinationUnionTypeLogging   TestDestinationUnionType = "Logging"
	TestDestinationUnionTypeSilent    TestDestinationUnionType = "Silent"
	TestDestinationUnionTypeThrottled TestDestinationUnionType = "Throttled"
	TestDestinationUnionTypeFailing   TestDestinationUnionType = "Failing"
)

// TestDestination - The type of destination to be used
type TestDestination struct {
	Logging   *Logging
	Silent    *Silent
	Throttled *Throttled
	Failing   *Failing

	Type TestDestinationUnionType
}

func CreateTestDestinationLogging(logging Logging) TestDestination {
	typ := TestDestinationUnionTypeLogging

	return TestDestination{
		Logging: &logging,
		Type:    typ,
	}
}

func CreateTestDestinationSilent(silent Silent) TestDestination {
	typ := TestDestinationUnionTypeSilent

	return TestDestination{
		Silent: &silent,
		Type:   typ,
	}
}

func CreateTestDestinationThrottled(throttled Throttled) TestDestination {
	typ := TestDestinationUnionTypeThrottled

	return TestDestination{
		Throttled: &throttled,
		Type:      typ,
	}
}

func CreateTestDestinationFailing(failing Failing) TestDestination {
	typ := TestDestinationUnionTypeFailing

	return TestDestination{
		Failing: &failing,
		Type:    typ,
	}
}

func (u *TestDestination) UnmarshalJSON(data []byte) error {

	var silent Silent = Silent{}
	if err := utils.UnmarshalJSON(data, &silent, "", true, true); err == nil {
		u.Silent = &silent
		u.Type = TestDestinationUnionTypeSilent
		return nil
	}

	var logging Logging = Logging{}
	if err := utils.UnmarshalJSON(data, &logging, "", true, true); err == nil {
		u.Logging = &logging
		u.Type = TestDestinationUnionTypeLogging
		return nil
	}

	var throttled Throttled = Throttled{}
	if err := utils.UnmarshalJSON(data, &throttled, "", true, true); err == nil {
		u.Throttled = &throttled
		u.Type = TestDestinationUnionTypeThrottled
		return nil
	}

	var failing Failing = Failing{}
	if err := utils.UnmarshalJSON(data, &failing, "", true, true); err == nil {
		u.Failing = &failing
		u.Type = TestDestinationUnionTypeFailing
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for TestDestination", string(data))
}

func (u TestDestination) MarshalJSON() ([]byte, error) {
	if u.Logging != nil {
		return utils.MarshalJSON(u.Logging, "", true)
	}

	if u.Silent != nil {
		return utils.MarshalJSON(u.Silent, "", true)
	}

	if u.Throttled != nil {
		return utils.MarshalJSON(u.Throttled, "", true)
	}

	if u.Failing != nil {
		return utils.MarshalJSON(u.Failing, "", true)
	}

	return nil, errors.New("could not marshal union type TestDestination: all fields are null")
}

type DestinationDevNullUpdate struct {
	// The type of destination to be used
	TestDestination TestDestination `json:"test_destination"`
}

func (o *DestinationDevNullUpdate) GetTestDestination() TestDestination {
	if o == nil {
		return TestDestination{}
	}
	return o.TestDestination
}
