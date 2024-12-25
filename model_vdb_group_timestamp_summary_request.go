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
	"time"
)

// checks if the VDBGroupTimestampSummaryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VDBGroupTimestampSummaryRequest{}

// VDBGroupTimestampSummaryRequest struct for VDBGroupTimestampSummaryRequest
type VDBGroupTimestampSummaryRequest struct {
	// The timestamp to get the summary for.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewVDBGroupTimestampSummaryRequest instantiates a new VDBGroupTimestampSummaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDBGroupTimestampSummaryRequest() *VDBGroupTimestampSummaryRequest {
	this := VDBGroupTimestampSummaryRequest{}
	return &this
}

// NewVDBGroupTimestampSummaryRequestWithDefaults instantiates a new VDBGroupTimestampSummaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDBGroupTimestampSummaryRequestWithDefaults() *VDBGroupTimestampSummaryRequest {
	this := VDBGroupTimestampSummaryRequest{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *VDBGroupTimestampSummaryRequest) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBGroupTimestampSummaryRequest) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *VDBGroupTimestampSummaryRequest) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *VDBGroupTimestampSummaryRequest) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o VDBGroupTimestampSummaryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VDBGroupTimestampSummaryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableVDBGroupTimestampSummaryRequest struct {
	value *VDBGroupTimestampSummaryRequest
	isSet bool
}

func (v NullableVDBGroupTimestampSummaryRequest) Get() *VDBGroupTimestampSummaryRequest {
	return v.value
}

func (v *NullableVDBGroupTimestampSummaryRequest) Set(val *VDBGroupTimestampSummaryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVDBGroupTimestampSummaryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVDBGroupTimestampSummaryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDBGroupTimestampSummaryRequest(val *VDBGroupTimestampSummaryRequest) *NullableVDBGroupTimestampSummaryRequest {
	return &NullableVDBGroupTimestampSummaryRequest{value: val, isSet: true}
}

func (v NullableVDBGroupTimestampSummaryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDBGroupTimestampSummaryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


