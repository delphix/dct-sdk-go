/*
Delphix DCT API

Delphix DCT API

API version: 3.18.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the HistoricalStorageSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HistoricalStorageSummary{}

// HistoricalStorageSummary struct for HistoricalStorageSummary
type HistoricalStorageSummary struct {
	Engines []EngineHistoricalStorageSummary `json:"engines,omitempty"`
}

// NewHistoricalStorageSummary instantiates a new HistoricalStorageSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalStorageSummary() *HistoricalStorageSummary {
	this := HistoricalStorageSummary{}
	return &this
}

// NewHistoricalStorageSummaryWithDefaults instantiates a new HistoricalStorageSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalStorageSummaryWithDefaults() *HistoricalStorageSummary {
	this := HistoricalStorageSummary{}
	return &this
}

// GetEngines returns the Engines field value if set, zero value otherwise.
func (o *HistoricalStorageSummary) GetEngines() []EngineHistoricalStorageSummary {
	if o == nil || IsNil(o.Engines) {
		var ret []EngineHistoricalStorageSummary
		return ret
	}
	return o.Engines
}

// GetEnginesOk returns a tuple with the Engines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalStorageSummary) GetEnginesOk() ([]EngineHistoricalStorageSummary, bool) {
	if o == nil || IsNil(o.Engines) {
		return nil, false
	}
	return o.Engines, true
}

// HasEngines returns a boolean if a field has been set.
func (o *HistoricalStorageSummary) HasEngines() bool {
	if o != nil && !IsNil(o.Engines) {
		return true
	}

	return false
}

// SetEngines gets a reference to the given []EngineHistoricalStorageSummary and assigns it to the Engines field.
func (o *HistoricalStorageSummary) SetEngines(v []EngineHistoricalStorageSummary) {
	o.Engines = v
}

func (o HistoricalStorageSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HistoricalStorageSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Engines) {
		toSerialize["engines"] = o.Engines
	}
	return toSerialize, nil
}

type NullableHistoricalStorageSummary struct {
	value *HistoricalStorageSummary
	isSet bool
}

func (v NullableHistoricalStorageSummary) Get() *HistoricalStorageSummary {
	return v.value
}

func (v *NullableHistoricalStorageSummary) Set(val *HistoricalStorageSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalStorageSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalStorageSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalStorageSummary(val *HistoricalStorageSummary) *NullableHistoricalStorageSummary {
	return &NullableHistoricalStorageSummary{value: val, isSet: true}
}

func (v NullableHistoricalStorageSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalStorageSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


