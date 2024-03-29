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

// checks if the ProductRegistrationOnlinePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductRegistrationOnlinePayload{}

// ProductRegistrationOnlinePayload Product registration object for a manually generated payload.
type ProductRegistrationOnlinePayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewProductRegistrationOnlinePayload instantiates a new ProductRegistrationOnlinePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductRegistrationOnlinePayload(username string, password string) *ProductRegistrationOnlinePayload {
	this := ProductRegistrationOnlinePayload{}
	this.Username = username
	this.Password = password
	return &this
}

// NewProductRegistrationOnlinePayloadWithDefaults instantiates a new ProductRegistrationOnlinePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductRegistrationOnlinePayloadWithDefaults() *ProductRegistrationOnlinePayload {
	this := ProductRegistrationOnlinePayload{}
	return &this
}

// GetUsername returns the Username field value
func (o *ProductRegistrationOnlinePayload) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ProductRegistrationOnlinePayload) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ProductRegistrationOnlinePayload) SetUsername(v string) {
	o.Username = v
}

// GetPassword returns the Password field value
func (o *ProductRegistrationOnlinePayload) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ProductRegistrationOnlinePayload) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ProductRegistrationOnlinePayload) SetPassword(v string) {
	o.Password = v
}

func (o ProductRegistrationOnlinePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductRegistrationOnlinePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableProductRegistrationOnlinePayload struct {
	value *ProductRegistrationOnlinePayload
	isSet bool
}

func (v NullableProductRegistrationOnlinePayload) Get() *ProductRegistrationOnlinePayload {
	return v.value
}

func (v *NullableProductRegistrationOnlinePayload) Set(val *ProductRegistrationOnlinePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProductRegistrationOnlinePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProductRegistrationOnlinePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductRegistrationOnlinePayload(val *ProductRegistrationOnlinePayload) *NullableProductRegistrationOnlinePayload {
	return &NullableProductRegistrationOnlinePayload{value: val, isSet: true}
}

func (v NullableProductRegistrationOnlinePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductRegistrationOnlinePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


