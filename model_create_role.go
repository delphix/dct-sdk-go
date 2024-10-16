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

// checks if the CreateRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRole{}

// CreateRole A role Object to create a custom role.
type CreateRole struct {
	// The Role name.
	Name string `json:"name"`
	// Role description.
	Description *string `json:"description,omitempty"`
	// The list of permissions granted by this role.
	PermissionObjects []PermissionObject `json:"permission_objects"`
	// If set to true, adding or removing permission is not allowed.
	Immutable *bool `json:"immutable,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

type _CreateRole CreateRole

// NewCreateRole instantiates a new CreateRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRole(name string, permissionObjects []PermissionObject) *CreateRole {
	this := CreateRole{}
	this.Name = name
	this.PermissionObjects = permissionObjects
	var immutable bool = false
	this.Immutable = &immutable
	return &this
}

// NewCreateRoleWithDefaults instantiates a new CreateRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRoleWithDefaults() *CreateRole {
	this := CreateRole{}
	var immutable bool = false
	this.Immutable = &immutable
	return &this
}

// GetName returns the Name field value
func (o *CreateRole) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRole) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRole) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateRole) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateRole) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateRole) SetDescription(v string) {
	o.Description = &v
}

// GetPermissionObjects returns the PermissionObjects field value
func (o *CreateRole) GetPermissionObjects() []PermissionObject {
	if o == nil {
		var ret []PermissionObject
		return ret
	}

	return o.PermissionObjects
}

// GetPermissionObjectsOk returns a tuple with the PermissionObjects field value
// and a boolean to check if the value has been set.
func (o *CreateRole) GetPermissionObjectsOk() ([]PermissionObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.PermissionObjects, true
}

// SetPermissionObjects sets field value
func (o *CreateRole) SetPermissionObjects(v []PermissionObject) {
	o.PermissionObjects = v
}

// GetImmutable returns the Immutable field value if set, zero value otherwise.
func (o *CreateRole) GetImmutable() bool {
	if o == nil || IsNil(o.Immutable) {
		var ret bool
		return ret
	}
	return *o.Immutable
}

// GetImmutableOk returns a tuple with the Immutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetImmutableOk() (*bool, bool) {
	if o == nil || IsNil(o.Immutable) {
		return nil, false
	}
	return o.Immutable, true
}

// HasImmutable returns a boolean if a field has been set.
func (o *CreateRole) HasImmutable() bool {
	if o != nil && !IsNil(o.Immutable) {
		return true
	}

	return false
}

// SetImmutable gets a reference to the given bool and assigns it to the Immutable field.
func (o *CreateRole) SetImmutable(v bool) {
	o.Immutable = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreateRole) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreateRole) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *CreateRole) SetTags(v []Tag) {
	o.Tags = v
}

func (o CreateRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["permission_objects"] = o.PermissionObjects
	if !IsNil(o.Immutable) {
		toSerialize["immutable"] = o.Immutable
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *CreateRole) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"permission_objects",
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

	varCreateRole := _CreateRole{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRole)

	if err != nil {
		return err
	}

	*o = CreateRole(varCreateRole)

	return err
}

type NullableCreateRole struct {
	value *CreateRole
	isSet bool
}

func (v NullableCreateRole) Get() *CreateRole {
	return v.value
}

func (v *NullableCreateRole) Set(val *CreateRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRole(val *CreateRole) *NullableCreateRole {
	return &NullableCreateRole{value: val, isSet: true}
}

func (v NullableCreateRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


