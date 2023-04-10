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

// checks if the BookmarkCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkCreateParameters{}

// BookmarkCreateParameters A Data Control Tower object that references points in time for one or more datasets.
type BookmarkCreateParameters struct {
	// The user-defined name of this bookmark.
	Name string `json:"name"`
	// The IDs of the VDBs to create the Bookmark on. This parameter is mutually exclusive with snapshot_ids.
	VdbIds []string `json:"vdb_ids,omitempty"`
	// The IDs of the snapshots that will be part of the Bookmark. This parameter is mutually exclusive with vdb_ids. 
	SnapshotIds []string `json:"snapshot_ids,omitempty"`
	// The retention policy for this bookmark, in days. A value of -1 indicates the bookmark should be kept forever. Deprecated in favor of expiration and retain_forever.
	// Deprecated
	Retention *int64 `json:"retention,omitempty"`
	// The expiration for this bookmark. Mutually exclusive with retention and retain_forever.
	Expiration *string `json:"expiration,omitempty"`
	// Indicates that the bookmark should be retained forever.
	RetainForever *bool `json:"retain_forever,omitempty"`
	// The tags to be created for this Bookmark.
	Tags []Tag `json:"tags,omitempty"`
	// Whether the account creating this bookmark must be configured as owner of the bookmark.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
}

// NewBookmarkCreateParameters instantiates a new BookmarkCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkCreateParameters(name string) *BookmarkCreateParameters {
	this := BookmarkCreateParameters{}
	this.Name = name
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// NewBookmarkCreateParametersWithDefaults instantiates a new BookmarkCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkCreateParametersWithDefaults() *BookmarkCreateParameters {
	this := BookmarkCreateParameters{}
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	return &this
}

// GetName returns the Name field value
func (o *BookmarkCreateParameters) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BookmarkCreateParameters) SetName(v string) {
	o.Name = v
}

// GetVdbIds returns the VdbIds field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetVdbIds() []string {
	if o == nil || IsNil(o.VdbIds) {
		var ret []string
		return ret
	}
	return o.VdbIds
}

// GetVdbIdsOk returns a tuple with the VdbIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetVdbIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VdbIds) {
		return nil, false
	}
	return o.VdbIds, true
}

// HasVdbIds returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasVdbIds() bool {
	if o != nil && !IsNil(o.VdbIds) {
		return true
	}

	return false
}

// SetVdbIds gets a reference to the given []string and assigns it to the VdbIds field.
func (o *BookmarkCreateParameters) SetVdbIds(v []string) {
	o.VdbIds = v
}

// GetSnapshotIds returns the SnapshotIds field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetSnapshotIds() []string {
	if o == nil || IsNil(o.SnapshotIds) {
		var ret []string
		return ret
	}
	return o.SnapshotIds
}

// GetSnapshotIdsOk returns a tuple with the SnapshotIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetSnapshotIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SnapshotIds) {
		return nil, false
	}
	return o.SnapshotIds, true
}

// HasSnapshotIds returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasSnapshotIds() bool {
	if o != nil && !IsNil(o.SnapshotIds) {
		return true
	}

	return false
}

// SetSnapshotIds gets a reference to the given []string and assigns it to the SnapshotIds field.
func (o *BookmarkCreateParameters) SetSnapshotIds(v []string) {
	o.SnapshotIds = v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
// Deprecated
func (o *BookmarkCreateParameters) GetRetention() int64 {
	if o == nil || IsNil(o.Retention) {
		var ret int64
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BookmarkCreateParameters) GetRetentionOk() (*int64, bool) {
	if o == nil || IsNil(o.Retention) {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasRetention() bool {
	if o != nil && !IsNil(o.Retention) {
		return true
	}

	return false
}

// SetRetention gets a reference to the given int64 and assigns it to the Retention field.
// Deprecated
func (o *BookmarkCreateParameters) SetRetention(v int64) {
	o.Retention = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetExpiration() string {
	if o == nil || IsNil(o.Expiration) {
		var ret string
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetExpirationOk() (*string, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given string and assigns it to the Expiration field.
func (o *BookmarkCreateParameters) SetExpiration(v string) {
	o.Expiration = &v
}

// GetRetainForever returns the RetainForever field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetRetainForever() bool {
	if o == nil || IsNil(o.RetainForever) {
		var ret bool
		return ret
	}
	return *o.RetainForever
}

// GetRetainForeverOk returns a tuple with the RetainForever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetRetainForeverOk() (*bool, bool) {
	if o == nil || IsNil(o.RetainForever) {
		return nil, false
	}
	return o.RetainForever, true
}

// HasRetainForever returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasRetainForever() bool {
	if o != nil && !IsNil(o.RetainForever) {
		return true
	}

	return false
}

// SetRetainForever gets a reference to the given bool and assigns it to the RetainForever field.
func (o *BookmarkCreateParameters) SetRetainForever(v bool) {
	o.RetainForever = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *BookmarkCreateParameters) SetTags(v []Tag) {
	o.Tags = v
}

// GetMakeCurrentAccountOwner returns the MakeCurrentAccountOwner field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetMakeCurrentAccountOwner() bool {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		var ret bool
		return ret
	}
	return *o.MakeCurrentAccountOwner
}

// GetMakeCurrentAccountOwnerOk returns a tuple with the MakeCurrentAccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetMakeCurrentAccountOwnerOk() (*bool, bool) {
	if o == nil || IsNil(o.MakeCurrentAccountOwner) {
		return nil, false
	}
	return o.MakeCurrentAccountOwner, true
}

// HasMakeCurrentAccountOwner returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasMakeCurrentAccountOwner() bool {
	if o != nil && !IsNil(o.MakeCurrentAccountOwner) {
		return true
	}

	return false
}

// SetMakeCurrentAccountOwner gets a reference to the given bool and assigns it to the MakeCurrentAccountOwner field.
func (o *BookmarkCreateParameters) SetMakeCurrentAccountOwner(v bool) {
	o.MakeCurrentAccountOwner = &v
}

func (o BookmarkCreateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkCreateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.VdbIds) {
		toSerialize["vdb_ids"] = o.VdbIds
	}
	if !IsNil(o.SnapshotIds) {
		toSerialize["snapshot_ids"] = o.SnapshotIds
	}
	if !IsNil(o.Retention) {
		toSerialize["retention"] = o.Retention
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	if !IsNil(o.RetainForever) {
		toSerialize["retain_forever"] = o.RetainForever
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.MakeCurrentAccountOwner) {
		toSerialize["make_current_account_owner"] = o.MakeCurrentAccountOwner
	}
	return toSerialize, nil
}

type NullableBookmarkCreateParameters struct {
	value *BookmarkCreateParameters
	isSet bool
}

func (v NullableBookmarkCreateParameters) Get() *BookmarkCreateParameters {
	return v.value
}

func (v *NullableBookmarkCreateParameters) Set(val *BookmarkCreateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkCreateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkCreateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkCreateParameters(val *BookmarkCreateParameters) *NullableBookmarkCreateParameters {
	return &NullableBookmarkCreateParameters{value: val, isSet: true}
}

func (v NullableBookmarkCreateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkCreateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


