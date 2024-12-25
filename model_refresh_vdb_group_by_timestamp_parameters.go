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

// checks if the RefreshVDBGroupByTimestampParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBGroupByTimestampParameters{}

// RefreshVDBGroupByTimestampParameters Parameters to refresh a VDB Group by timestamp.
type RefreshVDBGroupByTimestampParameters struct {
	// List of the pair of VDB and timestamp to refresh from. If this is not set, all VDBs will be refreshed from latest timestamp of their parent.
	VdbTimestampMappings []VDBGroupRefreshByTimestamp `json:"vdb_timestamp_mappings,omitempty"`
}

// NewRefreshVDBGroupByTimestampParameters instantiates a new RefreshVDBGroupByTimestampParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBGroupByTimestampParameters() *RefreshVDBGroupByTimestampParameters {
	this := RefreshVDBGroupByTimestampParameters{}
	return &this
}

// NewRefreshVDBGroupByTimestampParametersWithDefaults instantiates a new RefreshVDBGroupByTimestampParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBGroupByTimestampParametersWithDefaults() *RefreshVDBGroupByTimestampParameters {
	this := RefreshVDBGroupByTimestampParameters{}
	return &this
}

// GetVdbTimestampMappings returns the VdbTimestampMappings field value if set, zero value otherwise.
func (o *RefreshVDBGroupByTimestampParameters) GetVdbTimestampMappings() []VDBGroupRefreshByTimestamp {
	if o == nil || IsNil(o.VdbTimestampMappings) {
		var ret []VDBGroupRefreshByTimestamp
		return ret
	}
	return o.VdbTimestampMappings
}

// GetVdbTimestampMappingsOk returns a tuple with the VdbTimestampMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBGroupByTimestampParameters) GetVdbTimestampMappingsOk() ([]VDBGroupRefreshByTimestamp, bool) {
	if o == nil || IsNil(o.VdbTimestampMappings) {
		return nil, false
	}
	return o.VdbTimestampMappings, true
}

// HasVdbTimestampMappings returns a boolean if a field has been set.
func (o *RefreshVDBGroupByTimestampParameters) HasVdbTimestampMappings() bool {
	if o != nil && !IsNil(o.VdbTimestampMappings) {
		return true
	}

	return false
}

// SetVdbTimestampMappings gets a reference to the given []VDBGroupRefreshByTimestamp and assigns it to the VdbTimestampMappings field.
func (o *RefreshVDBGroupByTimestampParameters) SetVdbTimestampMappings(v []VDBGroupRefreshByTimestamp) {
	o.VdbTimestampMappings = v
}

func (o RefreshVDBGroupByTimestampParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBGroupByTimestampParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbTimestampMappings) {
		toSerialize["vdb_timestamp_mappings"] = o.VdbTimestampMappings
	}
	return toSerialize, nil
}

type NullableRefreshVDBGroupByTimestampParameters struct {
	value *RefreshVDBGroupByTimestampParameters
	isSet bool
}

func (v NullableRefreshVDBGroupByTimestampParameters) Get() *RefreshVDBGroupByTimestampParameters {
	return v.value
}

func (v *NullableRefreshVDBGroupByTimestampParameters) Set(val *RefreshVDBGroupByTimestampParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBGroupByTimestampParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBGroupByTimestampParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBGroupByTimestampParameters(val *RefreshVDBGroupByTimestampParameters) *NullableRefreshVDBGroupByTimestampParameters {
	return &NullableRefreshVDBGroupByTimestampParameters{value: val, isSet: true}
}

func (v NullableRefreshVDBGroupByTimestampParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBGroupByTimestampParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


