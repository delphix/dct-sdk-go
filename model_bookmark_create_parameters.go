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
	"time"
	"bytes"
	"fmt"
)

// checks if the BookmarkCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkCreateParameters{}

// BookmarkCreateParameters A Data Control Tower object that references points in time for one or more datasets.
type BookmarkCreateParameters struct {
	// The user-defined name of this bookmark.
	Name string `json:"name"`
	// The IDs of the VDBs to create the Bookmark on. This parameter is mutually exclusive with snapshot_ids and timeflow_ids.
	VdbIds []string `json:"vdb_ids,omitempty"`
	// The ID of the VDB group to create the Bookmark on. This parameter is mutually exclusive with vdb_ids.
	VdbGroupId *string `json:"vdb_group_id,omitempty"`
	// The IDs of the snapshots that will be part of the Bookmark. This parameter is mutually exclusive with vdb_ids, timestamp, timestamp_in_database_timezone, location and timeflow_ids. 
	SnapshotIds []string `json:"snapshot_ids,omitempty"`
	// The array of timeflow Id. Only allowed to set when timestamp, timestamp_in_database_timezone or location is provided.
	TimeflowIds []string `json:"timeflow_ids,omitempty"`
	// The point in time from which to execute the operation. Mutually exclusive with snapshot_ids, timestamp_in_database_timezone and location.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with snapshot_ids, timestamp and location.
	TimestampInDatabaseTimezone *string `json:"timestamp_in_database_timezone,omitempty" validate:"regexp=[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(.[0-9]{0,3})?"`
	// The location to create bookmark from. Mutually exclusive with snapshot_ids, timestamp, and timestamp_in_database_timezone.
	Location *string `json:"location,omitempty"`
	// The retention policy for this bookmark, in days. A value of -1 indicates the bookmark should be kept forever. Deprecated in favor of expiration and retain_forever.
	// Deprecated
	Retention *int64 `json:"retention,omitempty"`
	// The expiration for this bookmark. Mutually exclusive with retention and retain_forever.
	Expiration *string `json:"expiration,omitempty"`
	// Indicates that the bookmark should be retained forever.
	RetainForever *bool `json:"retain_forever,omitempty"`
	// The tags to be created for this Bookmark.
	Tags []Tag `json:"tags,omitempty"`
	// Type of the bookmark, either PUBLIC or PRIVATE.
	BookmarkType *string `json:"bookmark_type,omitempty"`
	// Whether the account creating this bookmark must be configured as owner of the bookmark.
	MakeCurrentAccountOwner *bool `json:"make_current_account_owner,omitempty"`
	// This field has been deprecated in favour of new field 'inherit_parent_tags'.
	// Deprecated
	InheritParentVdbTags *bool `json:"inherit_parent_vdb_tags,omitempty"`
	// Whether this bookmark should inherit tags from the parent dataset.
	InheritParentTags *bool `json:"inherit_parent_tags,omitempty"`
}

type _BookmarkCreateParameters BookmarkCreateParameters

// NewBookmarkCreateParameters instantiates a new BookmarkCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkCreateParameters(name string) *BookmarkCreateParameters {
	this := BookmarkCreateParameters{}
	this.Name = name
	var bookmarkType string = "PRIVATE"
	this.BookmarkType = &bookmarkType
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	var inheritParentVdbTags bool = false
	this.InheritParentVdbTags = &inheritParentVdbTags
	var inheritParentTags bool = false
	this.InheritParentTags = &inheritParentTags
	return &this
}

// NewBookmarkCreateParametersWithDefaults instantiates a new BookmarkCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkCreateParametersWithDefaults() *BookmarkCreateParameters {
	this := BookmarkCreateParameters{}
	var bookmarkType string = "PRIVATE"
	this.BookmarkType = &bookmarkType
	var makeCurrentAccountOwner bool = true
	this.MakeCurrentAccountOwner = &makeCurrentAccountOwner
	var inheritParentVdbTags bool = false
	this.InheritParentVdbTags = &inheritParentVdbTags
	var inheritParentTags bool = false
	this.InheritParentTags = &inheritParentTags
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

// GetVdbGroupId returns the VdbGroupId field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetVdbGroupId() string {
	if o == nil || IsNil(o.VdbGroupId) {
		var ret string
		return ret
	}
	return *o.VdbGroupId
}

// GetVdbGroupIdOk returns a tuple with the VdbGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetVdbGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.VdbGroupId) {
		return nil, false
	}
	return o.VdbGroupId, true
}

// HasVdbGroupId returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasVdbGroupId() bool {
	if o != nil && !IsNil(o.VdbGroupId) {
		return true
	}

	return false
}

// SetVdbGroupId gets a reference to the given string and assigns it to the VdbGroupId field.
func (o *BookmarkCreateParameters) SetVdbGroupId(v string) {
	o.VdbGroupId = &v
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

// GetTimeflowIds returns the TimeflowIds field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetTimeflowIds() []string {
	if o == nil || IsNil(o.TimeflowIds) {
		var ret []string
		return ret
	}
	return o.TimeflowIds
}

// GetTimeflowIdsOk returns a tuple with the TimeflowIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetTimeflowIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TimeflowIds) {
		return nil, false
	}
	return o.TimeflowIds, true
}

// HasTimeflowIds returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasTimeflowIds() bool {
	if o != nil && !IsNil(o.TimeflowIds) {
		return true
	}

	return false
}

// SetTimeflowIds gets a reference to the given []string and assigns it to the TimeflowIds field.
func (o *BookmarkCreateParameters) SetTimeflowIds(v []string) {
	o.TimeflowIds = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *BookmarkCreateParameters) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetTimestampInDatabaseTimezone() string {
	if o == nil || IsNil(o.TimestampInDatabaseTimezone) {
		var ret string
		return ret
	}
	return *o.TimestampInDatabaseTimezone
}

// GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimestampInDatabaseTimezone) {
		return nil, false
	}
	return o.TimestampInDatabaseTimezone, true
}

// HasTimestampInDatabaseTimezone returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasTimestampInDatabaseTimezone() bool {
	if o != nil && !IsNil(o.TimestampInDatabaseTimezone) {
		return true
	}

	return false
}

// SetTimestampInDatabaseTimezone gets a reference to the given string and assigns it to the TimestampInDatabaseTimezone field.
func (o *BookmarkCreateParameters) SetTimestampInDatabaseTimezone(v string) {
	o.TimestampInDatabaseTimezone = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *BookmarkCreateParameters) SetLocation(v string) {
	o.Location = &v
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

// GetBookmarkType returns the BookmarkType field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetBookmarkType() string {
	if o == nil || IsNil(o.BookmarkType) {
		var ret string
		return ret
	}
	return *o.BookmarkType
}

// GetBookmarkTypeOk returns a tuple with the BookmarkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetBookmarkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BookmarkType) {
		return nil, false
	}
	return o.BookmarkType, true
}

// HasBookmarkType returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasBookmarkType() bool {
	if o != nil && !IsNil(o.BookmarkType) {
		return true
	}

	return false
}

// SetBookmarkType gets a reference to the given string and assigns it to the BookmarkType field.
func (o *BookmarkCreateParameters) SetBookmarkType(v string) {
	o.BookmarkType = &v
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

// GetInheritParentVdbTags returns the InheritParentVdbTags field value if set, zero value otherwise.
// Deprecated
func (o *BookmarkCreateParameters) GetInheritParentVdbTags() bool {
	if o == nil || IsNil(o.InheritParentVdbTags) {
		var ret bool
		return ret
	}
	return *o.InheritParentVdbTags
}

// GetInheritParentVdbTagsOk returns a tuple with the InheritParentVdbTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *BookmarkCreateParameters) GetInheritParentVdbTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritParentVdbTags) {
		return nil, false
	}
	return o.InheritParentVdbTags, true
}

// HasInheritParentVdbTags returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasInheritParentVdbTags() bool {
	if o != nil && !IsNil(o.InheritParentVdbTags) {
		return true
	}

	return false
}

// SetInheritParentVdbTags gets a reference to the given bool and assigns it to the InheritParentVdbTags field.
// Deprecated
func (o *BookmarkCreateParameters) SetInheritParentVdbTags(v bool) {
	o.InheritParentVdbTags = &v
}

// GetInheritParentTags returns the InheritParentTags field value if set, zero value otherwise.
func (o *BookmarkCreateParameters) GetInheritParentTags() bool {
	if o == nil || IsNil(o.InheritParentTags) {
		var ret bool
		return ret
	}
	return *o.InheritParentTags
}

// GetInheritParentTagsOk returns a tuple with the InheritParentTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BookmarkCreateParameters) GetInheritParentTagsOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritParentTags) {
		return nil, false
	}
	return o.InheritParentTags, true
}

// HasInheritParentTags returns a boolean if a field has been set.
func (o *BookmarkCreateParameters) HasInheritParentTags() bool {
	if o != nil && !IsNil(o.InheritParentTags) {
		return true
	}

	return false
}

// SetInheritParentTags gets a reference to the given bool and assigns it to the InheritParentTags field.
func (o *BookmarkCreateParameters) SetInheritParentTags(v bool) {
	o.InheritParentTags = &v
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
	if !IsNil(o.VdbGroupId) {
		toSerialize["vdb_group_id"] = o.VdbGroupId
	}
	if !IsNil(o.SnapshotIds) {
		toSerialize["snapshot_ids"] = o.SnapshotIds
	}
	if !IsNil(o.TimeflowIds) {
		toSerialize["timeflow_ids"] = o.TimeflowIds
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TimestampInDatabaseTimezone) {
		toSerialize["timestamp_in_database_timezone"] = o.TimestampInDatabaseTimezone
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
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
	if !IsNil(o.BookmarkType) {
		toSerialize["bookmark_type"] = o.BookmarkType
	}
	if !IsNil(o.MakeCurrentAccountOwner) {
		toSerialize["make_current_account_owner"] = o.MakeCurrentAccountOwner
	}
	if !IsNil(o.InheritParentVdbTags) {
		toSerialize["inherit_parent_vdb_tags"] = o.InheritParentVdbTags
	}
	if !IsNil(o.InheritParentTags) {
		toSerialize["inherit_parent_tags"] = o.InheritParentTags
	}
	return toSerialize, nil
}

func (o *BookmarkCreateParameters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBookmarkCreateParameters := _BookmarkCreateParameters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBookmarkCreateParameters)

	if err != nil {
		return err
	}

	*o = BookmarkCreateParameters(varBookmarkCreateParameters)

	return err
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


