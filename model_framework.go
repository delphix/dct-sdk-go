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

// checks if the Framework type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Framework{}

// Framework A masking algorithm framework.
type Framework struct {
	// The masking framework entity ID.
	Id *string `json:"id,omitempty"`
	// The name of this framework.
	Name *string `json:"name,omitempty"`
	// A description of this framework.
	Description NullableString `json:"description,omitempty"`
	// The masking type of this framework.
	MaskingType *string `json:"masking_type,omitempty"`
}

// NewFramework instantiates a new Framework object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFramework() *Framework {
	this := Framework{}
	return &this
}

// NewFrameworkWithDefaults instantiates a new Framework object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameworkWithDefaults() *Framework {
	this := Framework{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Framework) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Framework) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Framework) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Framework) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Framework) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Framework) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Framework) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Framework) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Framework) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Framework) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Framework) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Framework) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Framework) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Framework) UnsetDescription() {
	o.Description.Unset()
}

// GetMaskingType returns the MaskingType field value if set, zero value otherwise.
func (o *Framework) GetMaskingType() string {
	if o == nil || IsNil(o.MaskingType) {
		var ret string
		return ret
	}
	return *o.MaskingType
}

// GetMaskingTypeOk returns a tuple with the MaskingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Framework) GetMaskingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MaskingType) {
		return nil, false
	}
	return o.MaskingType, true
}

// HasMaskingType returns a boolean if a field has been set.
func (o *Framework) HasMaskingType() bool {
	if o != nil && !IsNil(o.MaskingType) {
		return true
	}

	return false
}

// SetMaskingType gets a reference to the given string and assigns it to the MaskingType field.
func (o *Framework) SetMaskingType(v string) {
	o.MaskingType = &v
}

func (o Framework) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Framework) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.MaskingType) {
		toSerialize["masking_type"] = o.MaskingType
	}
	return toSerialize, nil
}

type NullableFramework struct {
	value *Framework
	isSet bool
}

func (v NullableFramework) Get() *Framework {
	return v.value
}

func (v *NullableFramework) Set(val *Framework) {
	v.value = val
	v.isSet = true
}

func (v NullableFramework) IsSet() bool {
	return v.isSet
}

func (v *NullableFramework) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFramework(val *Framework) *NullableFramework {
	return &NullableFramework{value: val, isSet: true}
}

func (v NullableFramework) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFramework) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


