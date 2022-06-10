/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ApiClientCreateResponse struct for ApiClientCreateResponse
type ApiClientCreateResponse struct {
	// The opaque token to use to authenticate for other API calls. The token must be included in all HTTP API calls in a request header named \"Authorization\", and prefixed with \"apk \" to denote that it is a proprietary API key format. For instance, if the token is \"abc123\", request must contain the following HTTP request header: \"Authorization: apk abc123\". 
	Token *string `json:"token,omitempty"`
	ApiClientEntityId *int64 `json:"api_client_entity_id,omitempty"`
}

// NewApiClientCreateResponse instantiates a new ApiClientCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClientCreateResponse() *ApiClientCreateResponse {
	this := ApiClientCreateResponse{}
	return &this
}

// NewApiClientCreateResponseWithDefaults instantiates a new ApiClientCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClientCreateResponseWithDefaults() *ApiClientCreateResponse {
	this := ApiClientCreateResponse{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ApiClientCreateResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClientCreateResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ApiClientCreateResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ApiClientCreateResponse) SetToken(v string) {
	o.Token = &v
}

// GetApiClientEntityId returns the ApiClientEntityId field value if set, zero value otherwise.
func (o *ApiClientCreateResponse) GetApiClientEntityId() int64 {
	if o == nil || o.ApiClientEntityId == nil {
		var ret int64
		return ret
	}
	return *o.ApiClientEntityId
}

// GetApiClientEntityIdOk returns a tuple with the ApiClientEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClientCreateResponse) GetApiClientEntityIdOk() (*int64, bool) {
	if o == nil || o.ApiClientEntityId == nil {
		return nil, false
	}
	return o.ApiClientEntityId, true
}

// HasApiClientEntityId returns a boolean if a field has been set.
func (o *ApiClientCreateResponse) HasApiClientEntityId() bool {
	if o != nil && o.ApiClientEntityId != nil {
		return true
	}

	return false
}

// SetApiClientEntityId gets a reference to the given int64 and assigns it to the ApiClientEntityId field.
func (o *ApiClientCreateResponse) SetApiClientEntityId(v int64) {
	o.ApiClientEntityId = &v
}

func (o ApiClientCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.ApiClientEntityId != nil {
		toSerialize["api_client_entity_id"] = o.ApiClientEntityId
	}
	return json.Marshal(toSerialize)
}

type NullableApiClientCreateResponse struct {
	value *ApiClientCreateResponse
	isSet bool
}

func (v NullableApiClientCreateResponse) Get() *ApiClientCreateResponse {
	return v.value
}

func (v *NullableApiClientCreateResponse) Set(val *ApiClientCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClientCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClientCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClientCreateResponse(val *ApiClientCreateResponse) *NullableApiClientCreateResponse {
	return &NullableApiClientCreateResponse{value: val, isSet: true}
}

func (v NullableApiClientCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClientCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


