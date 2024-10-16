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

// checks if the UpdateStagingSourceParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStagingSourceParameters{}

// UpdateStagingSourceParameters Parameters to update a Staging Source.
type UpdateStagingSourceParameters struct {
	// List of jdbc connection strings which are used to connect with the database.
	OracleServices []string `json:"oracle_services,omitempty"`
}

// NewUpdateStagingSourceParameters instantiates a new UpdateStagingSourceParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStagingSourceParameters() *UpdateStagingSourceParameters {
	this := UpdateStagingSourceParameters{}
	return &this
}

// NewUpdateStagingSourceParametersWithDefaults instantiates a new UpdateStagingSourceParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStagingSourceParametersWithDefaults() *UpdateStagingSourceParameters {
	this := UpdateStagingSourceParameters{}
	return &this
}

// GetOracleServices returns the OracleServices field value if set, zero value otherwise.
func (o *UpdateStagingSourceParameters) GetOracleServices() []string {
	if o == nil || IsNil(o.OracleServices) {
		var ret []string
		return ret
	}
	return o.OracleServices
}

// GetOracleServicesOk returns a tuple with the OracleServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStagingSourceParameters) GetOracleServicesOk() ([]string, bool) {
	if o == nil || IsNil(o.OracleServices) {
		return nil, false
	}
	return o.OracleServices, true
}

// HasOracleServices returns a boolean if a field has been set.
func (o *UpdateStagingSourceParameters) HasOracleServices() bool {
	if o != nil && !IsNil(o.OracleServices) {
		return true
	}

	return false
}

// SetOracleServices gets a reference to the given []string and assigns it to the OracleServices field.
func (o *UpdateStagingSourceParameters) SetOracleServices(v []string) {
	o.OracleServices = v
}

func (o UpdateStagingSourceParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateStagingSourceParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OracleServices) {
		toSerialize["oracle_services"] = o.OracleServices
	}
	return toSerialize, nil
}

type NullableUpdateStagingSourceParameters struct {
	value *UpdateStagingSourceParameters
	isSet bool
}

func (v NullableUpdateStagingSourceParameters) Get() *UpdateStagingSourceParameters {
	return v.value
}

func (v *NullableUpdateStagingSourceParameters) Set(val *UpdateStagingSourceParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStagingSourceParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStagingSourceParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStagingSourceParameters(val *UpdateStagingSourceParameters) *NullableUpdateStagingSourceParameters {
	return &NullableUpdateStagingSourceParameters{value: val, isSet: true}
}

func (v NullableUpdateStagingSourceParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStagingSourceParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


