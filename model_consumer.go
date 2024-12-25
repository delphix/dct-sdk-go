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
)

// checks if the Consumer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Consumer{}

// Consumer Consumer object.
type Consumer struct {
	// ID of the parent object.
	ParentId *string `json:"parent_id,omitempty"`
	// The name of the parent object.
	ParentName *string `json:"parent_name,omitempty"`
	// The type of the object.
	ParentType *string `json:"parent_type,omitempty"`
}

// NewConsumer instantiates a new Consumer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumer() *Consumer {
	this := Consumer{}
	return &this
}

// NewConsumerWithDefaults instantiates a new Consumer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerWithDefaults() *Consumer {
	this := Consumer{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Consumer) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Consumer) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Consumer) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Consumer) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *Consumer) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Consumer) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *Consumer) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *Consumer) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *Consumer) GetParentType() string {
	if o == nil || IsNil(o.ParentType) {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Consumer) GetParentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentType) {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *Consumer) HasParentType() bool {
	if o != nil && !IsNil(o.ParentType) {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *Consumer) SetParentType(v string) {
	o.ParentType = &v
}

func (o Consumer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Consumer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.ParentName) {
		toSerialize["parent_name"] = o.ParentName
	}
	if !IsNil(o.ParentType) {
		toSerialize["parent_type"] = o.ParentType
	}
	return toSerialize, nil
}

type NullableConsumer struct {
	value *Consumer
	isSet bool
}

func (v NullableConsumer) Get() *Consumer {
	return v.value
}

func (v *NullableConsumer) Set(val *Consumer) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumer) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumer(val *Consumer) *NullableConsumer {
	return &NullableConsumer{value: val, isSet: true}
}

func (v NullableConsumer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsumer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


