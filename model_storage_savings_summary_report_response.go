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

// checks if the StorageSavingsSummaryReportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageSavingsSummaryReportResponse{}

// StorageSavingsSummaryReportResponse struct for StorageSavingsSummaryReportResponse
type StorageSavingsSummaryReportResponse struct {
	Items []StorageSavingsSummaryData `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
	Totals *StorageSavingsReportSummarizedData `json:"totals,omitempty"`
}

// NewStorageSavingsSummaryReportResponse instantiates a new StorageSavingsSummaryReportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSavingsSummaryReportResponse() *StorageSavingsSummaryReportResponse {
	this := StorageSavingsSummaryReportResponse{}
	return &this
}

// NewStorageSavingsSummaryReportResponseWithDefaults instantiates a new StorageSavingsSummaryReportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSavingsSummaryReportResponseWithDefaults() *StorageSavingsSummaryReportResponse {
	this := StorageSavingsSummaryReportResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *StorageSavingsSummaryReportResponse) GetItems() []StorageSavingsSummaryData {
	if o == nil || IsNil(o.Items) {
		var ret []StorageSavingsSummaryData
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsSummaryReportResponse) GetItemsOk() ([]StorageSavingsSummaryData, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *StorageSavingsSummaryReportResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []StorageSavingsSummaryData and assigns it to the Items field.
func (o *StorageSavingsSummaryReportResponse) SetItems(v []StorageSavingsSummaryData) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *StorageSavingsSummaryReportResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsSummaryReportResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *StorageSavingsSummaryReportResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *StorageSavingsSummaryReportResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

// GetTotals returns the Totals field value if set, zero value otherwise.
func (o *StorageSavingsSummaryReportResponse) GetTotals() StorageSavingsReportSummarizedData {
	if o == nil || IsNil(o.Totals) {
		var ret StorageSavingsReportSummarizedData
		return ret
	}
	return *o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsSummaryReportResponse) GetTotalsOk() (*StorageSavingsReportSummarizedData, bool) {
	if o == nil || IsNil(o.Totals) {
		return nil, false
	}
	return o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *StorageSavingsSummaryReportResponse) HasTotals() bool {
	if o != nil && !IsNil(o.Totals) {
		return true
	}

	return false
}

// SetTotals gets a reference to the given StorageSavingsReportSummarizedData and assigns it to the Totals field.
func (o *StorageSavingsSummaryReportResponse) SetTotals(v StorageSavingsReportSummarizedData) {
	o.Totals = &v
}

func (o StorageSavingsSummaryReportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageSavingsSummaryReportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	if !IsNil(o.Totals) {
		toSerialize["totals"] = o.Totals
	}
	return toSerialize, nil
}

type NullableStorageSavingsSummaryReportResponse struct {
	value *StorageSavingsSummaryReportResponse
	isSet bool
}

func (v NullableStorageSavingsSummaryReportResponse) Get() *StorageSavingsSummaryReportResponse {
	return v.value
}

func (v *NullableStorageSavingsSummaryReportResponse) Set(val *StorageSavingsSummaryReportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSavingsSummaryReportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSavingsSummaryReportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSavingsSummaryReportResponse(val *StorageSavingsSummaryReportResponse) *NullableStorageSavingsSummaryReportResponse {
	return &NullableStorageSavingsSummaryReportResponse{value: val, isSet: true}
}

func (v NullableStorageSavingsSummaryReportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSavingsSummaryReportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


