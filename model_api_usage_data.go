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

// ApiUsageData struct for ApiUsageData
type ApiUsageData struct {
	// API called.
	ApiEndpoint string `json:"api_endpoint"`
	// HTTP method for API called.
	ApiMethod string `json:"api_method"`
	// Count of API calls over the requested timeframe.
	ApiCount int64 `json:"api_count"`
}

// NewApiUsageData instantiates a new ApiUsageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUsageData(apiEndpoint string, apiMethod string, apiCount int64) *ApiUsageData {
	this := ApiUsageData{}
	this.ApiEndpoint = apiEndpoint
	this.ApiMethod = apiMethod
	this.ApiCount = apiCount
	return &this
}

// NewApiUsageDataWithDefaults instantiates a new ApiUsageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUsageDataWithDefaults() *ApiUsageData {
	this := ApiUsageData{}
	return &this
}

// GetApiEndpoint returns the ApiEndpoint field value
func (o *ApiUsageData) GetApiEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiEndpoint
}

// GetApiEndpointOk returns a tuple with the ApiEndpoint field value
// and a boolean to check if the value has been set.
func (o *ApiUsageData) GetApiEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiEndpoint, true
}

// SetApiEndpoint sets field value
func (o *ApiUsageData) SetApiEndpoint(v string) {
	o.ApiEndpoint = v
}

// GetApiMethod returns the ApiMethod field value
func (o *ApiUsageData) GetApiMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiMethod
}

// GetApiMethodOk returns a tuple with the ApiMethod field value
// and a boolean to check if the value has been set.
func (o *ApiUsageData) GetApiMethodOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiMethod, true
}

// SetApiMethod sets field value
func (o *ApiUsageData) SetApiMethod(v string) {
	o.ApiMethod = v
}

// GetApiCount returns the ApiCount field value
func (o *ApiUsageData) GetApiCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ApiCount
}

// GetApiCountOk returns a tuple with the ApiCount field value
// and a boolean to check if the value has been set.
func (o *ApiUsageData) GetApiCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ApiCount, true
}

// SetApiCount sets field value
func (o *ApiUsageData) SetApiCount(v int64) {
	o.ApiCount = v
}

func (o ApiUsageData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_endpoint"] = o.ApiEndpoint
	}
	if true {
		toSerialize["api_method"] = o.ApiMethod
	}
	if true {
		toSerialize["api_count"] = o.ApiCount
	}
	return json.Marshal(toSerialize)
}

type NullableApiUsageData struct {
	value *ApiUsageData
	isSet bool
}

func (v NullableApiUsageData) Get() *ApiUsageData {
	return v.value
}

func (v *NullableApiUsageData) Set(val *ApiUsageData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUsageData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUsageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUsageData(val *ApiUsageData) *NullableApiUsageData {
	return &NullableApiUsageData{value: val, isSet: true}
}

func (v NullableApiUsageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUsageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


