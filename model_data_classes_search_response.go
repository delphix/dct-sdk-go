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

// checks if the DataClassesSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataClassesSearchResponse{}

// DataClassesSearchResponse struct for DataClassesSearchResponse
type DataClassesSearchResponse struct {
	Items []DataClass `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewDataClassesSearchResponse instantiates a new DataClassesSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataClassesSearchResponse() *DataClassesSearchResponse {
	this := DataClassesSearchResponse{}
	return &this
}

// NewDataClassesSearchResponseWithDefaults instantiates a new DataClassesSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataClassesSearchResponseWithDefaults() *DataClassesSearchResponse {
	this := DataClassesSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DataClassesSearchResponse) GetItems() []DataClass {
	if o == nil || IsNil(o.Items) {
		var ret []DataClass
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataClassesSearchResponse) GetItemsOk() ([]DataClass, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DataClassesSearchResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DataClass and assigns it to the Items field.
func (o *DataClassesSearchResponse) SetItems(v []DataClass) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *DataClassesSearchResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataClassesSearchResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *DataClassesSearchResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *DataClassesSearchResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o DataClassesSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataClassesSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableDataClassesSearchResponse struct {
	value *DataClassesSearchResponse
	isSet bool
}

func (v NullableDataClassesSearchResponse) Get() *DataClassesSearchResponse {
	return v.value
}

func (v *NullableDataClassesSearchResponse) Set(val *DataClassesSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataClassesSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataClassesSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataClassesSearchResponse(val *DataClassesSearchResponse) *NullableDataClassesSearchResponse {
	return &NullableDataClassesSearchResponse{value: val, isSet: true}
}

func (v NullableDataClassesSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataClassesSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


