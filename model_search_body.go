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

// SearchBody Search body.
type SearchBody struct {
	FilterExpression *string `json:"filter_expression,omitempty"`
}

// NewSearchBody instantiates a new SearchBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchBody() *SearchBody {
	this := SearchBody{}
	return &this
}

// NewSearchBodyWithDefaults instantiates a new SearchBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchBodyWithDefaults() *SearchBody {
	this := SearchBody{}
	return &this
}

// GetFilterExpression returns the FilterExpression field value if set, zero value otherwise.
func (o *SearchBody) GetFilterExpression() string {
	if o == nil || o.FilterExpression == nil {
		var ret string
		return ret
	}
	return *o.FilterExpression
}

// GetFilterExpressionOk returns a tuple with the FilterExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchBody) GetFilterExpressionOk() (*string, bool) {
	if o == nil || o.FilterExpression == nil {
		return nil, false
	}
	return o.FilterExpression, true
}

// HasFilterExpression returns a boolean if a field has been set.
func (o *SearchBody) HasFilterExpression() bool {
	if o != nil && o.FilterExpression != nil {
		return true
	}

	return false
}

// SetFilterExpression gets a reference to the given string and assigns it to the FilterExpression field.
func (o *SearchBody) SetFilterExpression(v string) {
	o.FilterExpression = &v
}

func (o SearchBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilterExpression != nil {
		toSerialize["filter_expression"] = o.FilterExpression
	}
	return json.Marshal(toSerialize)
}

type NullableSearchBody struct {
	value *SearchBody
	isSet bool
}

func (v NullableSearchBody) Get() *SearchBody {
	return v.value
}

func (v *NullableSearchBody) Set(val *SearchBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchBody(val *SearchBody) *NullableSearchBody {
	return &NullableSearchBody{value: val, isSet: true}
}

func (v NullableSearchBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


