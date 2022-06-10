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
	"time"
)

// Bookmark A Data Control Tower object that references points in time for one or more datasets.
type Bookmark struct {
	// The Bookmark object entity ID.
	Id *string `json:"id,omitempty"`
	// The user-defined name of this bookmark.
	Name string `json:"name"`
	// The date and time that this bookmark was created.
	CreationDate *time.Time `json:"creation_date,omitempty"`
	// The list of VDB IDs associated with this bookmark.
	VdbIds []string `json:"vdb_ids"`
	// The retention policy for this bookmark, in days. A value of -1 indicates the bookmark should be kept forever.
	Retention *int64 `json:"retention,omitempty"`
	// A message with details about operation progress or state of this bookmark.
	Status NullableString `json:"status,omitempty"`
}

// NewBookmark instantiates a new Bookmark object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmark(name string, vdbIds []string) *Bookmark {
	this := Bookmark{}
	this.Name = name
	this.VdbIds = vdbIds
	return &this
}

// NewBookmarkWithDefaults instantiates a new Bookmark object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkWithDefaults() *Bookmark {
	this := Bookmark{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Bookmark) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookmark) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Bookmark) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Bookmark) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Bookmark) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Bookmark) SetName(v string) {
	o.Name = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *Bookmark) GetCreationDate() time.Time {
	if o == nil || o.CreationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookmark) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *Bookmark) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *Bookmark) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetVdbIds returns the VdbIds field value
func (o *Bookmark) GetVdbIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VdbIds
}

// GetVdbIdsOk returns a tuple with the VdbIds field value
// and a boolean to check if the value has been set.
func (o *Bookmark) GetVdbIdsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VdbIds, true
}

// SetVdbIds sets field value
func (o *Bookmark) SetVdbIds(v []string) {
	o.VdbIds = v
}

// GetRetention returns the Retention field value if set, zero value otherwise.
func (o *Bookmark) GetRetention() int64 {
	if o == nil || o.Retention == nil {
		var ret int64
		return ret
	}
	return *o.Retention
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Bookmark) GetRetentionOk() (*int64, bool) {
	if o == nil || o.Retention == nil {
		return nil, false
	}
	return o.Retention, true
}

// HasRetention returns a boolean if a field has been set.
func (o *Bookmark) HasRetention() bool {
	if o != nil && o.Retention != nil {
		return true
	}

	return false
}

// SetRetention gets a reference to the given int64 and assigns it to the Retention field.
func (o *Bookmark) SetRetention(v int64) {
	o.Retention = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Bookmark) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Bookmark) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Bookmark) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Bookmark) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Bookmark) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Bookmark) UnsetStatus() {
	o.Status.Unset()
}

func (o Bookmark) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CreationDate != nil {
		toSerialize["creation_date"] = o.CreationDate
	}
	if true {
		toSerialize["vdb_ids"] = o.VdbIds
	}
	if o.Retention != nil {
		toSerialize["retention"] = o.Retention
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBookmark struct {
	value *Bookmark
	isSet bool
}

func (v NullableBookmark) Get() *Bookmark {
	return v.value
}

func (v *NullableBookmark) Set(val *Bookmark) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmark) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmark) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmark(val *Bookmark) *NullableBookmark {
	return &NullableBookmark{value: val, isSet: true}
}

func (v NullableBookmark) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmark) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


