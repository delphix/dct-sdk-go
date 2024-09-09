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

// checks if the DatabaseTableMetadataSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseTableMetadataSearchResponse{}

// DatabaseTableMetadataSearchResponse struct for DatabaseTableMetadataSearchResponse
type DatabaseTableMetadataSearchResponse struct {
	Items []DatabaseTableMetadata `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewDatabaseTableMetadataSearchResponse instantiates a new DatabaseTableMetadataSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseTableMetadataSearchResponse() *DatabaseTableMetadataSearchResponse {
	this := DatabaseTableMetadataSearchResponse{}
	return &this
}

// NewDatabaseTableMetadataSearchResponseWithDefaults instantiates a new DatabaseTableMetadataSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseTableMetadataSearchResponseWithDefaults() *DatabaseTableMetadataSearchResponse {
	this := DatabaseTableMetadataSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DatabaseTableMetadataSearchResponse) GetItems() []DatabaseTableMetadata {
	if o == nil || IsNil(o.Items) {
		var ret []DatabaseTableMetadata
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadataSearchResponse) GetItemsOk() ([]DatabaseTableMetadata, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DatabaseTableMetadataSearchResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DatabaseTableMetadata and assigns it to the Items field.
func (o *DatabaseTableMetadataSearchResponse) SetItems(v []DatabaseTableMetadata) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *DatabaseTableMetadataSearchResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseTableMetadataSearchResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *DatabaseTableMetadataSearchResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *DatabaseTableMetadataSearchResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o DatabaseTableMetadataSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseTableMetadataSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableDatabaseTableMetadataSearchResponse struct {
	value *DatabaseTableMetadataSearchResponse
	isSet bool
}

func (v NullableDatabaseTableMetadataSearchResponse) Get() *DatabaseTableMetadataSearchResponse {
	return v.value
}

func (v *NullableDatabaseTableMetadataSearchResponse) Set(val *DatabaseTableMetadataSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseTableMetadataSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseTableMetadataSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseTableMetadataSearchResponse(val *DatabaseTableMetadataSearchResponse) *NullableDatabaseTableMetadataSearchResponse {
	return &NullableDatabaseTableMetadataSearchResponse{value: val, isSet: true}
}

func (v NullableDatabaseTableMetadataSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseTableMetadataSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


