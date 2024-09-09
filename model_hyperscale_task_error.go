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

// checks if the HyperscaleTaskError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HyperscaleTaskError{}

// HyperscaleTaskError struct for HyperscaleTaskError
type HyperscaleTaskError struct {
	// the name of the table for which the error occurred, including the schema.
	TableName *string `json:"table_name,omitempty"`
	// A textual description of the error.
	Error *string `json:"error,omitempty"`
}

// NewHyperscaleTaskError instantiates a new HyperscaleTaskError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperscaleTaskError() *HyperscaleTaskError {
	this := HyperscaleTaskError{}
	return &this
}

// NewHyperscaleTaskErrorWithDefaults instantiates a new HyperscaleTaskError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperscaleTaskErrorWithDefaults() *HyperscaleTaskError {
	this := HyperscaleTaskError{}
	return &this
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *HyperscaleTaskError) GetTableName() string {
	if o == nil || IsNil(o.TableName) {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleTaskError) GetTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.TableName) {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *HyperscaleTaskError) HasTableName() bool {
	if o != nil && !IsNil(o.TableName) {
		return true
	}

	return false
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *HyperscaleTaskError) SetTableName(v string) {
	o.TableName = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *HyperscaleTaskError) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperscaleTaskError) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *HyperscaleTaskError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *HyperscaleTaskError) SetError(v string) {
	o.Error = &v
}

func (o HyperscaleTaskError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HyperscaleTaskError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TableName) {
		toSerialize["table_name"] = o.TableName
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableHyperscaleTaskError struct {
	value *HyperscaleTaskError
	isSet bool
}

func (v NullableHyperscaleTaskError) Get() *HyperscaleTaskError {
	return v.value
}

func (v *NullableHyperscaleTaskError) Set(val *HyperscaleTaskError) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperscaleTaskError) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperscaleTaskError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperscaleTaskError(val *HyperscaleTaskError) *NullableHyperscaleTaskError {
	return &NullableHyperscaleTaskError{value: val, isSet: true}
}

func (v NullableHyperscaleTaskError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperscaleTaskError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


