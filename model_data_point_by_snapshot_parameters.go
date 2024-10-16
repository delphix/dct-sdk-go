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
)

// checks if the DataPointBySnapshotParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPointBySnapshotParameters{}

// DataPointBySnapshotParameters struct for DataPointBySnapshotParameters
type DataPointBySnapshotParameters struct {
	// The ID of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot.
	SnapshotId *string `json:"snapshot_id,omitempty"`
}

// NewDataPointBySnapshotParameters instantiates a new DataPointBySnapshotParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPointBySnapshotParameters() *DataPointBySnapshotParameters {
	this := DataPointBySnapshotParameters{}
	return &this
}

// NewDataPointBySnapshotParametersWithDefaults instantiates a new DataPointBySnapshotParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointBySnapshotParametersWithDefaults() *DataPointBySnapshotParameters {
	this := DataPointBySnapshotParameters{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *DataPointBySnapshotParameters) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPointBySnapshotParameters) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *DataPointBySnapshotParameters) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *DataPointBySnapshotParameters) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

func (o DataPointBySnapshotParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPointBySnapshotParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	return toSerialize, nil
}

type NullableDataPointBySnapshotParameters struct {
	value *DataPointBySnapshotParameters
	isSet bool
}

func (v NullableDataPointBySnapshotParameters) Get() *DataPointBySnapshotParameters {
	return v.value
}

func (v *NullableDataPointBySnapshotParameters) Set(val *DataPointBySnapshotParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPointBySnapshotParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPointBySnapshotParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPointBySnapshotParameters(val *DataPointBySnapshotParameters) *NullableDataPointBySnapshotParameters {
	return &NullableDataPointBySnapshotParameters{value: val, isSet: true}
}

func (v NullableDataPointBySnapshotParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPointBySnapshotParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


