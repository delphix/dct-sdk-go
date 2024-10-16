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
)

// checks if the ExecutionComponentsSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionComponentsSearchResponse{}

// ExecutionComponentsSearchResponse struct for ExecutionComponentsSearchResponse
type ExecutionComponentsSearchResponse struct {
	Items []ExecutionComponent `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewExecutionComponentsSearchResponse instantiates a new ExecutionComponentsSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionComponentsSearchResponse() *ExecutionComponentsSearchResponse {
	this := ExecutionComponentsSearchResponse{}
	return &this
}

// NewExecutionComponentsSearchResponseWithDefaults instantiates a new ExecutionComponentsSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionComponentsSearchResponseWithDefaults() *ExecutionComponentsSearchResponse {
	this := ExecutionComponentsSearchResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ExecutionComponentsSearchResponse) GetItems() []ExecutionComponent {
	if o == nil || IsNil(o.Items) {
		var ret []ExecutionComponent
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionComponentsSearchResponse) GetItemsOk() ([]ExecutionComponent, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ExecutionComponentsSearchResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ExecutionComponent and assigns it to the Items field.
func (o *ExecutionComponentsSearchResponse) SetItems(v []ExecutionComponent) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ExecutionComponentsSearchResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionComponentsSearchResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ExecutionComponentsSearchResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ExecutionComponentsSearchResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ExecutionComponentsSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionComponentsSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableExecutionComponentsSearchResponse struct {
	value *ExecutionComponentsSearchResponse
	isSet bool
}

func (v NullableExecutionComponentsSearchResponse) Get() *ExecutionComponentsSearchResponse {
	return v.value
}

func (v *NullableExecutionComponentsSearchResponse) Set(val *ExecutionComponentsSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionComponentsSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionComponentsSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionComponentsSearchResponse(val *ExecutionComponentsSearchResponse) *NullableExecutionComponentsSearchResponse {
	return &NullableExecutionComponentsSearchResponse{value: val, isSet: true}
}

func (v NullableExecutionComponentsSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionComponentsSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


