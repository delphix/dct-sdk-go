/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the DataPointByLocationParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPointByLocationParameters{}

// DataPointByLocationParameters struct for DataPointByLocationParameters
type DataPointByLocationParameters struct {
	// The location to provision from.
	Location *string `json:"location,omitempty"`
	// ID of the timeflow to provision from.
	TimeflowId *string `json:"timeflow_id,omitempty"`
}

// NewDataPointByLocationParameters instantiates a new DataPointByLocationParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPointByLocationParameters() *DataPointByLocationParameters {
	this := DataPointByLocationParameters{}
	return &this
}

// NewDataPointByLocationParametersWithDefaults instantiates a new DataPointByLocationParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointByLocationParametersWithDefaults() *DataPointByLocationParameters {
	this := DataPointByLocationParameters{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *DataPointByLocationParameters) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPointByLocationParameters) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *DataPointByLocationParameters) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *DataPointByLocationParameters) SetLocation(v string) {
	o.Location = &v
}

// GetTimeflowId returns the TimeflowId field value if set, zero value otherwise.
func (o *DataPointByLocationParameters) GetTimeflowId() string {
	if o == nil || IsNil(o.TimeflowId) {
		var ret string
		return ret
	}
	return *o.TimeflowId
}

// GetTimeflowIdOk returns a tuple with the TimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPointByLocationParameters) GetTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeflowId) {
		return nil, false
	}
	return o.TimeflowId, true
}

// HasTimeflowId returns a boolean if a field has been set.
func (o *DataPointByLocationParameters) HasTimeflowId() bool {
	if o != nil && !IsNil(o.TimeflowId) {
		return true
	}

	return false
}

// SetTimeflowId gets a reference to the given string and assigns it to the TimeflowId field.
func (o *DataPointByLocationParameters) SetTimeflowId(v string) {
	o.TimeflowId = &v
}

func (o DataPointByLocationParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPointByLocationParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.TimeflowId) {
		toSerialize["timeflow_id"] = o.TimeflowId
	}
	return toSerialize, nil
}

type NullableDataPointByLocationParameters struct {
	value *DataPointByLocationParameters
	isSet bool
}

func (v NullableDataPointByLocationParameters) Get() *DataPointByLocationParameters {
	return v.value
}

func (v *NullableDataPointByLocationParameters) Set(val *DataPointByLocationParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPointByLocationParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPointByLocationParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPointByLocationParameters(val *DataPointByLocationParameters) *NullableDataPointByLocationParameters {
	return &NullableDataPointByLocationParameters{value: val, isSet: true}
}

func (v NullableDataPointByLocationParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPointByLocationParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


