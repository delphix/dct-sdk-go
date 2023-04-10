/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the PermissionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionObject{}

// PermissionObject A Permission Object which is mapping between object type and its permissions.
type PermissionObject struct {
	// Object type.
	ObjectType string `json:"object_type"`
	// List of permissions.
	Permissions []string `json:"permissions"`
}

// NewPermissionObject instantiates a new PermissionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionObject(objectType string, permissions []string) *PermissionObject {
	this := PermissionObject{}
	this.ObjectType = objectType
	this.Permissions = permissions
	return &this
}

// NewPermissionObjectWithDefaults instantiates a new PermissionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionObjectWithDefaults() *PermissionObject {
	this := PermissionObject{}
	return &this
}

// GetObjectType returns the ObjectType field value
func (o *PermissionObject) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *PermissionObject) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *PermissionObject) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPermissions returns the Permissions field value
func (o *PermissionObject) GetPermissions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *PermissionObject) GetPermissionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *PermissionObject) SetPermissions(v []string) {
	o.Permissions = v
}

func (o PermissionObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_type"] = o.ObjectType
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

type NullablePermissionObject struct {
	value *PermissionObject
	isSet bool
}

func (v NullablePermissionObject) Get() *PermissionObject {
	return v.value
}

func (v *NullablePermissionObject) Set(val *PermissionObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionObject(val *PermissionObject) *NullablePermissionObject {
	return &NullablePermissionObject{value: val, isSet: true}
}

func (v NullablePermissionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


