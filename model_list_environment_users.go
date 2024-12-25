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

// checks if the ListEnvironmentUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEnvironmentUsers{}

// ListEnvironmentUsers struct for ListEnvironmentUsers
type ListEnvironmentUsers struct {
	// List of users
	Users []EnvironmentUser `json:"users,omitempty"`
}

// NewListEnvironmentUsers instantiates a new ListEnvironmentUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEnvironmentUsers() *ListEnvironmentUsers {
	this := ListEnvironmentUsers{}
	return &this
}

// NewListEnvironmentUsersWithDefaults instantiates a new ListEnvironmentUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEnvironmentUsersWithDefaults() *ListEnvironmentUsers {
	this := ListEnvironmentUsers{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ListEnvironmentUsers) GetUsers() []EnvironmentUser {
	if o == nil || IsNil(o.Users) {
		var ret []EnvironmentUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEnvironmentUsers) GetUsersOk() ([]EnvironmentUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ListEnvironmentUsers) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []EnvironmentUser and assigns it to the Users field.
func (o *ListEnvironmentUsers) SetUsers(v []EnvironmentUser) {
	o.Users = v
}

func (o ListEnvironmentUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEnvironmentUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableListEnvironmentUsers struct {
	value *ListEnvironmentUsers
	isSet bool
}

func (v NullableListEnvironmentUsers) Get() *ListEnvironmentUsers {
	return v.value
}

func (v *NullableListEnvironmentUsers) Set(val *ListEnvironmentUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnvironmentUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnvironmentUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnvironmentUsers(val *ListEnvironmentUsers) *NullableListEnvironmentUsers {
	return &NullableListEnvironmentUsers{value: val, isSet: true}
}

func (v NullableListEnvironmentUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnvironmentUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


