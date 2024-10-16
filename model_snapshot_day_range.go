/*
Delphix DCT API

Delphix DCT API

API version: 3.17.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
	"time"
)

// checks if the SnapshotDayRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotDayRange{}

// SnapshotDayRange Count of TimeFlow snapshots aggregated by day.
type SnapshotDayRange struct {
	// Number of TimeFlow snapshots on that day.
	Count *int64 `json:"count,omitempty"`
	// Date for which TimeFlow snapshots have been aggregated.
	Date *string `json:"date,omitempty"`
	// Start of day of this range in the time zone used for computation.
	StartOfDay *time.Time `json:"start_of_day,omitempty"`
	// End of day of this range in the time zone used for computation.
	EndOfDay *time.Time `json:"end_of_day,omitempty"`
}

// NewSnapshotDayRange instantiates a new SnapshotDayRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotDayRange() *SnapshotDayRange {
	this := SnapshotDayRange{}
	return &this
}

// NewSnapshotDayRangeWithDefaults instantiates a new SnapshotDayRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotDayRangeWithDefaults() *SnapshotDayRange {
	this := SnapshotDayRange{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *SnapshotDayRange) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDayRange) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *SnapshotDayRange) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *SnapshotDayRange) SetCount(v int64) {
	o.Count = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SnapshotDayRange) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDayRange) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SnapshotDayRange) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SnapshotDayRange) SetDate(v string) {
	o.Date = &v
}

// GetStartOfDay returns the StartOfDay field value if set, zero value otherwise.
func (o *SnapshotDayRange) GetStartOfDay() time.Time {
	if o == nil || IsNil(o.StartOfDay) {
		var ret time.Time
		return ret
	}
	return *o.StartOfDay
}

// GetStartOfDayOk returns a tuple with the StartOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDayRange) GetStartOfDayOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartOfDay) {
		return nil, false
	}
	return o.StartOfDay, true
}

// HasStartOfDay returns a boolean if a field has been set.
func (o *SnapshotDayRange) HasStartOfDay() bool {
	if o != nil && !IsNil(o.StartOfDay) {
		return true
	}

	return false
}

// SetStartOfDay gets a reference to the given time.Time and assigns it to the StartOfDay field.
func (o *SnapshotDayRange) SetStartOfDay(v time.Time) {
	o.StartOfDay = &v
}

// GetEndOfDay returns the EndOfDay field value if set, zero value otherwise.
func (o *SnapshotDayRange) GetEndOfDay() time.Time {
	if o == nil || IsNil(o.EndOfDay) {
		var ret time.Time
		return ret
	}
	return *o.EndOfDay
}

// GetEndOfDayOk returns a tuple with the EndOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotDayRange) GetEndOfDayOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndOfDay) {
		return nil, false
	}
	return o.EndOfDay, true
}

// HasEndOfDay returns a boolean if a field has been set.
func (o *SnapshotDayRange) HasEndOfDay() bool {
	if o != nil && !IsNil(o.EndOfDay) {
		return true
	}

	return false
}

// SetEndOfDay gets a reference to the given time.Time and assigns it to the EndOfDay field.
func (o *SnapshotDayRange) SetEndOfDay(v time.Time) {
	o.EndOfDay = &v
}

func (o SnapshotDayRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotDayRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.StartOfDay) {
		toSerialize["start_of_day"] = o.StartOfDay
	}
	if !IsNil(o.EndOfDay) {
		toSerialize["end_of_day"] = o.EndOfDay
	}
	return toSerialize, nil
}

type NullableSnapshotDayRange struct {
	value *SnapshotDayRange
	isSet bool
}

func (v NullableSnapshotDayRange) Get() *SnapshotDayRange {
	return v.value
}

func (v *NullableSnapshotDayRange) Set(val *SnapshotDayRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotDayRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotDayRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotDayRange(val *SnapshotDayRange) *NullableSnapshotDayRange {
	return &NullableSnapshotDayRange{value: val, isSet: true}
}

func (v NullableSnapshotDayRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotDayRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


