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

// checks if the SearchEngineGlobalObjectStateReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchEngineGlobalObjectStateReportResponse{}

// SearchEngineGlobalObjectStateReportResponse struct for SearchEngineGlobalObjectStateReportResponse
type SearchEngineGlobalObjectStateReportResponse struct {
	Items []EngineGlobalObjectStateData `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchEngineGlobalObjectStateReportResponse instantiates a new SearchEngineGlobalObjectStateReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEngineGlobalObjectStateReportResponse() *SearchEngineGlobalObjectStateReportResponse {
	this := SearchEngineGlobalObjectStateReportResponse{}
	return &this
}

// NewSearchEngineGlobalObjectStateReportResponseWithDefaults instantiates a new SearchEngineGlobalObjectStateReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEngineGlobalObjectStateReportResponseWithDefaults() *SearchEngineGlobalObjectStateReportResponse {
	this := SearchEngineGlobalObjectStateReportResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchEngineGlobalObjectStateReportResponse) GetItems() []EngineGlobalObjectStateData {
	if o == nil || IsNil(o.Items) {
		var ret []EngineGlobalObjectStateData
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEngineGlobalObjectStateReportResponse) GetItemsOk() ([]EngineGlobalObjectStateData, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchEngineGlobalObjectStateReportResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []EngineGlobalObjectStateData and assigns it to the Items field.
func (o *SearchEngineGlobalObjectStateReportResponse) SetItems(v []EngineGlobalObjectStateData) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchEngineGlobalObjectStateReportResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEngineGlobalObjectStateReportResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchEngineGlobalObjectStateReportResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchEngineGlobalObjectStateReportResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchEngineGlobalObjectStateReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchEngineGlobalObjectStateReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchEngineGlobalObjectStateReportResponse struct {
	value *SearchEngineGlobalObjectStateReportResponse
	isSet bool
}

func (v NullableSearchEngineGlobalObjectStateReportResponse) Get() *SearchEngineGlobalObjectStateReportResponse {
	return v.value
}

func (v *NullableSearchEngineGlobalObjectStateReportResponse) Set(val *SearchEngineGlobalObjectStateReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEngineGlobalObjectStateReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEngineGlobalObjectStateReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEngineGlobalObjectStateReportResponse(val *SearchEngineGlobalObjectStateReportResponse) *NullableSearchEngineGlobalObjectStateReportResponse {
	return &NullableSearchEngineGlobalObjectStateReportResponse{value: val, isSet: true}
}

func (v NullableSearchEngineGlobalObjectStateReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEngineGlobalObjectStateReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


