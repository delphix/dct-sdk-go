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

// checks if the DeleteScopedObjectItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteScopedObjectItem{}

// DeleteScopedObjectItem struct for DeleteScopedObjectItem
type DeleteScopedObjectItem struct {
	// List of scoped objects to be deleted
	Objects []ScopedObjectItem `json:"objects"`
}

type _DeleteScopedObjectItem DeleteScopedObjectItem

// NewDeleteScopedObjectItem instantiates a new DeleteScopedObjectItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteScopedObjectItem(objects []ScopedObjectItem) *DeleteScopedObjectItem {
	this := DeleteScopedObjectItem{}
	this.Objects = objects
	return &this
}

// NewDeleteScopedObjectItemWithDefaults instantiates a new DeleteScopedObjectItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteScopedObjectItemWithDefaults() *DeleteScopedObjectItem {
	this := DeleteScopedObjectItem{}
	return &this
}

// GetObjects returns the Objects field value
func (o *DeleteScopedObjectItem) GetObjects() []ScopedObjectItem {
	if o == nil {
		var ret []ScopedObjectItem
		return ret
	}

	return o.Objects
}

// GetObjectsOk returns a tuple with the Objects field value
// and a boolean to check if the value has been set.
func (o *DeleteScopedObjectItem) GetObjectsOk() ([]ScopedObjectItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Objects, true
}

// SetObjects sets field value
func (o *DeleteScopedObjectItem) SetObjects(v []ScopedObjectItem) {
	o.Objects = v
}

func (o DeleteScopedObjectItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteScopedObjectItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objects"] = o.Objects
	return toSerialize, nil
}

func (o *DeleteScopedObjectItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objects",
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

	varDeleteScopedObjectItem := _DeleteScopedObjectItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteScopedObjectItem)

	if err != nil {
		return err
	}

	*o = DeleteScopedObjectItem(varDeleteScopedObjectItem)

	return err
}

type NullableDeleteScopedObjectItem struct {
	value *DeleteScopedObjectItem
	isSet bool
}

func (v NullableDeleteScopedObjectItem) Get() *DeleteScopedObjectItem {
	return v.value
}

func (v *NullableDeleteScopedObjectItem) Set(val *DeleteScopedObjectItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteScopedObjectItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteScopedObjectItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteScopedObjectItem(val *DeleteScopedObjectItem) *NullableDeleteScopedObjectItem {
	return &NullableDeleteScopedObjectItem{value: val, isSet: true}
}

func (v NullableDeleteScopedObjectItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteScopedObjectItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


