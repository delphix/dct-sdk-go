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
)

// ListSnapshotsResponse struct for ListSnapshotsResponse
type ListSnapshotsResponse struct {
	Items []Snapshot `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListSnapshotsResponse instantiates a new ListSnapshotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSnapshotsResponse() *ListSnapshotsResponse {
	this := ListSnapshotsResponse{}
	return &this
}

// NewListSnapshotsResponseWithDefaults instantiates a new ListSnapshotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSnapshotsResponseWithDefaults() *ListSnapshotsResponse {
	this := ListSnapshotsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListSnapshotsResponse) GetItems() []Snapshot {
	if o == nil || o.Items == nil {
		var ret []Snapshot
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSnapshotsResponse) GetItemsOk() ([]Snapshot, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListSnapshotsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Snapshot and assigns it to the Items field.
func (o *ListSnapshotsResponse) SetItems(v []Snapshot) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListSnapshotsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || o.ResponseMetadata == nil {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListSnapshotsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || o.ResponseMetadata == nil {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListSnapshotsResponse) HasResponseMetadata() bool {
	if o != nil && o.ResponseMetadata != nil {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListSnapshotsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListSnapshotsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.ResponseMetadata != nil {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return json.Marshal(toSerialize)
}

type NullableListSnapshotsResponse struct {
	value *ListSnapshotsResponse
	isSet bool
}

func (v NullableListSnapshotsResponse) Get() *ListSnapshotsResponse {
	return v.value
}

func (v *NullableListSnapshotsResponse) Set(val *ListSnapshotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListSnapshotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListSnapshotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSnapshotsResponse(val *ListSnapshotsResponse) *NullableListSnapshotsResponse {
	return &NullableListSnapshotsResponse{value: val, isSet: true}
}

func (v NullableListSnapshotsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSnapshotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


