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
	"bytes"
	"fmt"
)

// checks if the TokenInfoRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenInfoRequest{}

// TokenInfoRequest struct for TokenInfoRequest
type TokenInfoRequest struct {
	// API Key or JWT token for fetching information
	Token string `json:"token"`
}

type _TokenInfoRequest TokenInfoRequest

// NewTokenInfoRequest instantiates a new TokenInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenInfoRequest(token string) *TokenInfoRequest {
	this := TokenInfoRequest{}
	this.Token = token
	return &this
}

// NewTokenInfoRequestWithDefaults instantiates a new TokenInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenInfoRequestWithDefaults() *TokenInfoRequest {
	this := TokenInfoRequest{}
	return &this
}

// GetToken returns the Token field value
func (o *TokenInfoRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *TokenInfoRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *TokenInfoRequest) SetToken(v string) {
	o.Token = v
}

func (o TokenInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *TokenInfoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"token",
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

	varTokenInfoRequest := _TokenInfoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenInfoRequest)

	if err != nil {
		return err
	}

	*o = TokenInfoRequest(varTokenInfoRequest)

	return err
}

type NullableTokenInfoRequest struct {
	value *TokenInfoRequest
	isSet bool
}

func (v NullableTokenInfoRequest) Get() *TokenInfoRequest {
	return v.value
}

func (v *NullableTokenInfoRequest) Set(val *TokenInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenInfoRequest(val *TokenInfoRequest) *NullableTokenInfoRequest {
	return &NullableTokenInfoRequest{value: val, isSet: true}
}

func (v NullableTokenInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


