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
	"bytes"
	"fmt"
)

// checks if the NameValuePair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameValuePair{}

// NameValuePair struct for NameValuePair
type NameValuePair struct {
	// The name of the environment variable.
	VarName string `json:"var_name"`
	// The value of the environment variable.
	VarValue string `json:"var_value"`
}

type _NameValuePair NameValuePair

// NewNameValuePair instantiates a new NameValuePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameValuePair(varName string, varValue string) *NameValuePair {
	this := NameValuePair{}
	this.VarName = varName
	this.VarValue = varValue
	return &this
}

// NewNameValuePairWithDefaults instantiates a new NameValuePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameValuePairWithDefaults() *NameValuePair {
	this := NameValuePair{}
	return &this
}

// GetVarName returns the VarName field value
func (o *NameValuePair) GetVarName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VarName
}

// GetVarNameOk returns a tuple with the VarName field value
// and a boolean to check if the value has been set.
func (o *NameValuePair) GetVarNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VarName, true
}

// SetVarName sets field value
func (o *NameValuePair) SetVarName(v string) {
	o.VarName = v
}

// GetVarValue returns the VarValue field value
func (o *NameValuePair) GetVarValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VarValue
}

// GetVarValueOk returns a tuple with the VarValue field value
// and a boolean to check if the value has been set.
func (o *NameValuePair) GetVarValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VarValue, true
}

// SetVarValue sets field value
func (o *NameValuePair) SetVarValue(v string) {
	o.VarValue = v
}

func (o NameValuePair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameValuePair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["var_name"] = o.VarName
	toSerialize["var_value"] = o.VarValue
	return toSerialize, nil
}

func (o *NameValuePair) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"var_name",
		"var_value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNameValuePair := _NameValuePair{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNameValuePair)

	if err != nil {
		return err
	}

	*o = NameValuePair(varNameValuePair)

	return err
}

type NullableNameValuePair struct {
	value *NameValuePair
	isSet bool
}

func (v NullableNameValuePair) Get() *NameValuePair {
	return v.value
}

func (v *NullableNameValuePair) Set(val *NameValuePair) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValuePair) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValuePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValuePair(val *NameValuePair) *NullableNameValuePair {
	return &NullableNameValuePair{value: val, isSet: true}
}

func (v NullableNameValuePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValuePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


