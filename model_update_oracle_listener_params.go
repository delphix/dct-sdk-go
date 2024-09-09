/*
Delphix DCT API

Delphix DCT API

API version: 3.16.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the UpdateOracleListenerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOracleListenerParams{}

// UpdateOracleListenerParams struct for UpdateOracleListenerParams
type UpdateOracleListenerParams struct {
	// The name of the Oracle listener.
	Name *string `json:"name,omitempty"`
	// The protocol addresses of the Oracle listener.
	ProtocolAddresses []string `json:"protocol_addresses,omitempty"`
}

// NewUpdateOracleListenerParams instantiates a new UpdateOracleListenerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOracleListenerParams() *UpdateOracleListenerParams {
	this := UpdateOracleListenerParams{}
	return &this
}

// NewUpdateOracleListenerParamsWithDefaults instantiates a new UpdateOracleListenerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOracleListenerParamsWithDefaults() *UpdateOracleListenerParams {
	this := UpdateOracleListenerParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateOracleListenerParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOracleListenerParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateOracleListenerParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateOracleListenerParams) SetName(v string) {
	o.Name = &v
}

// GetProtocolAddresses returns the ProtocolAddresses field value if set, zero value otherwise.
func (o *UpdateOracleListenerParams) GetProtocolAddresses() []string {
	if o == nil || IsNil(o.ProtocolAddresses) {
		var ret []string
		return ret
	}
	return o.ProtocolAddresses
}

// GetProtocolAddressesOk returns a tuple with the ProtocolAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOracleListenerParams) GetProtocolAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProtocolAddresses) {
		return nil, false
	}
	return o.ProtocolAddresses, true
}

// HasProtocolAddresses returns a boolean if a field has been set.
func (o *UpdateOracleListenerParams) HasProtocolAddresses() bool {
	if o != nil && !IsNil(o.ProtocolAddresses) {
		return true
	}

	return false
}

// SetProtocolAddresses gets a reference to the given []string and assigns it to the ProtocolAddresses field.
func (o *UpdateOracleListenerParams) SetProtocolAddresses(v []string) {
	o.ProtocolAddresses = v
}

func (o UpdateOracleListenerParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOracleListenerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ProtocolAddresses) {
		toSerialize["protocol_addresses"] = o.ProtocolAddresses
	}
	return toSerialize, nil
}

type NullableUpdateOracleListenerParams struct {
	value *UpdateOracleListenerParams
	isSet bool
}

func (v NullableUpdateOracleListenerParams) Get() *UpdateOracleListenerParams {
	return v.value
}

func (v *NullableUpdateOracleListenerParams) Set(val *UpdateOracleListenerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOracleListenerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOracleListenerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOracleListenerParams(val *UpdateOracleListenerParams) *NullableUpdateOracleListenerParams {
	return &NullableUpdateOracleListenerParams{value: val, isSet: true}
}

func (v NullableUpdateOracleListenerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOracleListenerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


