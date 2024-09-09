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
)

// checks if the ListHyperscaleDatasetTablesOrFilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHyperscaleDatasetTablesOrFilesResponse{}

// ListHyperscaleDatasetTablesOrFilesResponse struct for ListHyperscaleDatasetTablesOrFilesResponse
type ListHyperscaleDatasetTablesOrFilesResponse struct {
	Items []HyperscaleDatasetTableOrFile `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListHyperscaleDatasetTablesOrFilesResponse instantiates a new ListHyperscaleDatasetTablesOrFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHyperscaleDatasetTablesOrFilesResponse() *ListHyperscaleDatasetTablesOrFilesResponse {
	this := ListHyperscaleDatasetTablesOrFilesResponse{}
	return &this
}

// NewListHyperscaleDatasetTablesOrFilesResponseWithDefaults instantiates a new ListHyperscaleDatasetTablesOrFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHyperscaleDatasetTablesOrFilesResponseWithDefaults() *ListHyperscaleDatasetTablesOrFilesResponse {
	this := ListHyperscaleDatasetTablesOrFilesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) GetItems() []HyperscaleDatasetTableOrFile {
	if o == nil || IsNil(o.Items) {
		var ret []HyperscaleDatasetTableOrFile
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) GetItemsOk() ([]HyperscaleDatasetTableOrFile, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []HyperscaleDatasetTableOrFile and assigns it to the Items field.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) SetItems(v []HyperscaleDatasetTableOrFile) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListHyperscaleDatasetTablesOrFilesResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListHyperscaleDatasetTablesOrFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHyperscaleDatasetTablesOrFilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListHyperscaleDatasetTablesOrFilesResponse struct {
	value *ListHyperscaleDatasetTablesOrFilesResponse
	isSet bool
}

func (v NullableListHyperscaleDatasetTablesOrFilesResponse) Get() *ListHyperscaleDatasetTablesOrFilesResponse {
	return v.value
}

func (v *NullableListHyperscaleDatasetTablesOrFilesResponse) Set(val *ListHyperscaleDatasetTablesOrFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListHyperscaleDatasetTablesOrFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListHyperscaleDatasetTablesOrFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHyperscaleDatasetTablesOrFilesResponse(val *ListHyperscaleDatasetTablesOrFilesResponse) *NullableListHyperscaleDatasetTablesOrFilesResponse {
	return &NullableListHyperscaleDatasetTablesOrFilesResponse{value: val, isSet: true}
}

func (v NullableListHyperscaleDatasetTablesOrFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHyperscaleDatasetTablesOrFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


