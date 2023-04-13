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

// checks if the SnapshotCompatibleRepositories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnapshotCompatibleRepositories{}

// SnapshotCompatibleRepositories Compatible repositories corresponding to the snapshot.
type SnapshotCompatibleRepositories struct {
	// Repositories corresponding to the snapshot. A Repository typically corresponds to a database installation.
	Repositories []EnvironmentRepository `json:"repositories,omitempty"`
}

// NewSnapshotCompatibleRepositories instantiates a new SnapshotCompatibleRepositories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotCompatibleRepositories() *SnapshotCompatibleRepositories {
	this := SnapshotCompatibleRepositories{}
	return &this
}

// NewSnapshotCompatibleRepositoriesWithDefaults instantiates a new SnapshotCompatibleRepositories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotCompatibleRepositoriesWithDefaults() *SnapshotCompatibleRepositories {
	this := SnapshotCompatibleRepositories{}
	return &this
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *SnapshotCompatibleRepositories) GetRepositories() []EnvironmentRepository {
	if o == nil || IsNil(o.Repositories) {
		var ret []EnvironmentRepository
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotCompatibleRepositories) GetRepositoriesOk() ([]EnvironmentRepository, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *SnapshotCompatibleRepositories) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []EnvironmentRepository and assigns it to the Repositories field.
func (o *SnapshotCompatibleRepositories) SetRepositories(v []EnvironmentRepository) {
	o.Repositories = v
}

func (o SnapshotCompatibleRepositories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotCompatibleRepositories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	return toSerialize, nil
}

type NullableSnapshotCompatibleRepositories struct {
	value *SnapshotCompatibleRepositories
	isSet bool
}

func (v NullableSnapshotCompatibleRepositories) Get() *SnapshotCompatibleRepositories {
	return v.value
}

func (v *NullableSnapshotCompatibleRepositories) Set(val *SnapshotCompatibleRepositories) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotCompatibleRepositories) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotCompatibleRepositories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotCompatibleRepositories(val *SnapshotCompatibleRepositories) *NullableSnapshotCompatibleRepositories {
	return &NullableSnapshotCompatibleRepositories{value: val, isSet: true}
}

func (v NullableSnapshotCompatibleRepositories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotCompatibleRepositories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


