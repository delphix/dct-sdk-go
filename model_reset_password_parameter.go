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

// checks if the ResetPasswordParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResetPasswordParameter{}

// ResetPasswordParameter struct for ResetPasswordParameter
type ResetPasswordParameter struct {
	// New password that needs to be set for the Account. Set this to null for unsetting the current password. Not including this property also results in unsetting of the current password.
	NewPassword *string `json:"new_password,omitempty"`
}

// NewResetPasswordParameter instantiates a new ResetPasswordParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResetPasswordParameter() *ResetPasswordParameter {
	this := ResetPasswordParameter{}
	return &this
}

// NewResetPasswordParameterWithDefaults instantiates a new ResetPasswordParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResetPasswordParameterWithDefaults() *ResetPasswordParameter {
	this := ResetPasswordParameter{}
	return &this
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *ResetPasswordParameter) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword) {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResetPasswordParameter) GetNewPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.NewPassword) {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *ResetPasswordParameter) HasNewPassword() bool {
	if o != nil && !IsNil(o.NewPassword) {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *ResetPasswordParameter) SetNewPassword(v string) {
	o.NewPassword = &v
}

func (o ResetPasswordParameter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResetPasswordParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewPassword) {
		toSerialize["new_password"] = o.NewPassword
	}
	return toSerialize, nil
}

type NullableResetPasswordParameter struct {
	value *ResetPasswordParameter
	isSet bool
}

func (v NullableResetPasswordParameter) Get() *ResetPasswordParameter {
	return v.value
}

func (v *NullableResetPasswordParameter) Set(val *ResetPasswordParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableResetPasswordParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableResetPasswordParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResetPasswordParameter(val *ResetPasswordParameter) *NullableResetPasswordParameter {
	return &NullableResetPasswordParameter{value: val, isSet: true}
}

func (v NullableResetPasswordParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResetPasswordParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


