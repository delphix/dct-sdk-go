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

// checks if the StorageSavingsReportSummarizedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageSavingsReportSummarizedData{}

// StorageSavingsReportSummarizedData struct for StorageSavingsReportSummarizedData
type StorageSavingsReportSummarizedData struct {
	// The total VDB count
	VdbCount *int32 `json:"vdb_count,omitempty"`
	// The total dSource count
	DsourceCount *int32 `json:"dsource_count,omitempty"`
	// Total Virtualized Space. This is the sum of storage size from dSources and their dependant VDBs.
	VirtualizedSpace *int64 `json:"virtualized_space,omitempty"`
	// Total disk space, in bytes, that it would take to store the filtered list of dSources and their descandant VDBs without Delphix, counting each of their timeflows as separate copy of the source data.
	UnvirtualizedSpace *int64 `json:"unvirtualized_space,omitempty"`
	// Total disk space, in bytes, that it would take to store the filtered list of dSources and their descandant VDBs without Delphix, counting only their current (active) timeflows.
	CurrentTimeflowsUnvirtualizedSpace *int64 `json:"current_timeflows_unvirtualized_space,omitempty"`
	// Total disk space that has been saved by using Delphix virtualization for all timeflows, in bytes.
	EstimatedSavings *int64 `json:"estimated_savings,omitempty"`
	// Total disk space that has been saved by using Delphix virtualization for all timeflows, in percentage.
	EstimatedSavingsPerc *float32 `json:"estimated_savings_perc,omitempty"`
	// Total disk space that has been saved by using Delphix virtualization for only the current (active) timeflows, in bytes.
	EstimatedCurrentTimeflowsSavings *int64 `json:"estimated_current_timeflows_savings,omitempty"`
	// Total disk space that has been saved by using Delphix virtualization for only the current (active) timeflows, in percentage.
	EstimatedCurrentTimeflowsSavingsPerc *float32 `json:"estimated_current_timeflows_savings_perc,omitempty"`
}

// NewStorageSavingsReportSummarizedData instantiates a new StorageSavingsReportSummarizedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageSavingsReportSummarizedData() *StorageSavingsReportSummarizedData {
	this := StorageSavingsReportSummarizedData{}
	return &this
}

// NewStorageSavingsReportSummarizedDataWithDefaults instantiates a new StorageSavingsReportSummarizedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageSavingsReportSummarizedDataWithDefaults() *StorageSavingsReportSummarizedData {
	this := StorageSavingsReportSummarizedData{}
	return &this
}

// GetVdbCount returns the VdbCount field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetVdbCount() int32 {
	if o == nil || IsNil(o.VdbCount) {
		var ret int32
		return ret
	}
	return *o.VdbCount
}

// GetVdbCountOk returns a tuple with the VdbCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetVdbCountOk() (*int32, bool) {
	if o == nil || IsNil(o.VdbCount) {
		return nil, false
	}
	return o.VdbCount, true
}

// HasVdbCount returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasVdbCount() bool {
	if o != nil && !IsNil(o.VdbCount) {
		return true
	}

	return false
}

// SetVdbCount gets a reference to the given int32 and assigns it to the VdbCount field.
func (o *StorageSavingsReportSummarizedData) SetVdbCount(v int32) {
	o.VdbCount = &v
}

// GetDsourceCount returns the DsourceCount field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetDsourceCount() int32 {
	if o == nil || IsNil(o.DsourceCount) {
		var ret int32
		return ret
	}
	return *o.DsourceCount
}

// GetDsourceCountOk returns a tuple with the DsourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetDsourceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DsourceCount) {
		return nil, false
	}
	return o.DsourceCount, true
}

// HasDsourceCount returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasDsourceCount() bool {
	if o != nil && !IsNil(o.DsourceCount) {
		return true
	}

	return false
}

// SetDsourceCount gets a reference to the given int32 and assigns it to the DsourceCount field.
func (o *StorageSavingsReportSummarizedData) SetDsourceCount(v int32) {
	o.DsourceCount = &v
}

// GetVirtualizedSpace returns the VirtualizedSpace field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetVirtualizedSpace() int64 {
	if o == nil || IsNil(o.VirtualizedSpace) {
		var ret int64
		return ret
	}
	return *o.VirtualizedSpace
}

// GetVirtualizedSpaceOk returns a tuple with the VirtualizedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetVirtualizedSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.VirtualizedSpace) {
		return nil, false
	}
	return o.VirtualizedSpace, true
}

// HasVirtualizedSpace returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasVirtualizedSpace() bool {
	if o != nil && !IsNil(o.VirtualizedSpace) {
		return true
	}

	return false
}

// SetVirtualizedSpace gets a reference to the given int64 and assigns it to the VirtualizedSpace field.
func (o *StorageSavingsReportSummarizedData) SetVirtualizedSpace(v int64) {
	o.VirtualizedSpace = &v
}

// GetUnvirtualizedSpace returns the UnvirtualizedSpace field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetUnvirtualizedSpace() int64 {
	if o == nil || IsNil(o.UnvirtualizedSpace) {
		var ret int64
		return ret
	}
	return *o.UnvirtualizedSpace
}

// GetUnvirtualizedSpaceOk returns a tuple with the UnvirtualizedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetUnvirtualizedSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.UnvirtualizedSpace) {
		return nil, false
	}
	return o.UnvirtualizedSpace, true
}

// HasUnvirtualizedSpace returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasUnvirtualizedSpace() bool {
	if o != nil && !IsNil(o.UnvirtualizedSpace) {
		return true
	}

	return false
}

// SetUnvirtualizedSpace gets a reference to the given int64 and assigns it to the UnvirtualizedSpace field.
func (o *StorageSavingsReportSummarizedData) SetUnvirtualizedSpace(v int64) {
	o.UnvirtualizedSpace = &v
}

// GetCurrentTimeflowsUnvirtualizedSpace returns the CurrentTimeflowsUnvirtualizedSpace field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetCurrentTimeflowsUnvirtualizedSpace() int64 {
	if o == nil || IsNil(o.CurrentTimeflowsUnvirtualizedSpace) {
		var ret int64
		return ret
	}
	return *o.CurrentTimeflowsUnvirtualizedSpace
}

// GetCurrentTimeflowsUnvirtualizedSpaceOk returns a tuple with the CurrentTimeflowsUnvirtualizedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetCurrentTimeflowsUnvirtualizedSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentTimeflowsUnvirtualizedSpace) {
		return nil, false
	}
	return o.CurrentTimeflowsUnvirtualizedSpace, true
}

// HasCurrentTimeflowsUnvirtualizedSpace returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasCurrentTimeflowsUnvirtualizedSpace() bool {
	if o != nil && !IsNil(o.CurrentTimeflowsUnvirtualizedSpace) {
		return true
	}

	return false
}

// SetCurrentTimeflowsUnvirtualizedSpace gets a reference to the given int64 and assigns it to the CurrentTimeflowsUnvirtualizedSpace field.
func (o *StorageSavingsReportSummarizedData) SetCurrentTimeflowsUnvirtualizedSpace(v int64) {
	o.CurrentTimeflowsUnvirtualizedSpace = &v
}

// GetEstimatedSavings returns the EstimatedSavings field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetEstimatedSavings() int64 {
	if o == nil || IsNil(o.EstimatedSavings) {
		var ret int64
		return ret
	}
	return *o.EstimatedSavings
}

// GetEstimatedSavingsOk returns a tuple with the EstimatedSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetEstimatedSavingsOk() (*int64, bool) {
	if o == nil || IsNil(o.EstimatedSavings) {
		return nil, false
	}
	return o.EstimatedSavings, true
}

// HasEstimatedSavings returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasEstimatedSavings() bool {
	if o != nil && !IsNil(o.EstimatedSavings) {
		return true
	}

	return false
}

// SetEstimatedSavings gets a reference to the given int64 and assigns it to the EstimatedSavings field.
func (o *StorageSavingsReportSummarizedData) SetEstimatedSavings(v int64) {
	o.EstimatedSavings = &v
}

// GetEstimatedSavingsPerc returns the EstimatedSavingsPerc field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetEstimatedSavingsPerc() float32 {
	if o == nil || IsNil(o.EstimatedSavingsPerc) {
		var ret float32
		return ret
	}
	return *o.EstimatedSavingsPerc
}

// GetEstimatedSavingsPercOk returns a tuple with the EstimatedSavingsPerc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetEstimatedSavingsPercOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedSavingsPerc) {
		return nil, false
	}
	return o.EstimatedSavingsPerc, true
}

// HasEstimatedSavingsPerc returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasEstimatedSavingsPerc() bool {
	if o != nil && !IsNil(o.EstimatedSavingsPerc) {
		return true
	}

	return false
}

// SetEstimatedSavingsPerc gets a reference to the given float32 and assigns it to the EstimatedSavingsPerc field.
func (o *StorageSavingsReportSummarizedData) SetEstimatedSavingsPerc(v float32) {
	o.EstimatedSavingsPerc = &v
}

// GetEstimatedCurrentTimeflowsSavings returns the EstimatedCurrentTimeflowsSavings field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetEstimatedCurrentTimeflowsSavings() int64 {
	if o == nil || IsNil(o.EstimatedCurrentTimeflowsSavings) {
		var ret int64
		return ret
	}
	return *o.EstimatedCurrentTimeflowsSavings
}

// GetEstimatedCurrentTimeflowsSavingsOk returns a tuple with the EstimatedCurrentTimeflowsSavings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetEstimatedCurrentTimeflowsSavingsOk() (*int64, bool) {
	if o == nil || IsNil(o.EstimatedCurrentTimeflowsSavings) {
		return nil, false
	}
	return o.EstimatedCurrentTimeflowsSavings, true
}

// HasEstimatedCurrentTimeflowsSavings returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasEstimatedCurrentTimeflowsSavings() bool {
	if o != nil && !IsNil(o.EstimatedCurrentTimeflowsSavings) {
		return true
	}

	return false
}

// SetEstimatedCurrentTimeflowsSavings gets a reference to the given int64 and assigns it to the EstimatedCurrentTimeflowsSavings field.
func (o *StorageSavingsReportSummarizedData) SetEstimatedCurrentTimeflowsSavings(v int64) {
	o.EstimatedCurrentTimeflowsSavings = &v
}

// GetEstimatedCurrentTimeflowsSavingsPerc returns the EstimatedCurrentTimeflowsSavingsPerc field value if set, zero value otherwise.
func (o *StorageSavingsReportSummarizedData) GetEstimatedCurrentTimeflowsSavingsPerc() float32 {
	if o == nil || IsNil(o.EstimatedCurrentTimeflowsSavingsPerc) {
		var ret float32
		return ret
	}
	return *o.EstimatedCurrentTimeflowsSavingsPerc
}

// GetEstimatedCurrentTimeflowsSavingsPercOk returns a tuple with the EstimatedCurrentTimeflowsSavingsPerc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageSavingsReportSummarizedData) GetEstimatedCurrentTimeflowsSavingsPercOk() (*float32, bool) {
	if o == nil || IsNil(o.EstimatedCurrentTimeflowsSavingsPerc) {
		return nil, false
	}
	return o.EstimatedCurrentTimeflowsSavingsPerc, true
}

// HasEstimatedCurrentTimeflowsSavingsPerc returns a boolean if a field has been set.
func (o *StorageSavingsReportSummarizedData) HasEstimatedCurrentTimeflowsSavingsPerc() bool {
	if o != nil && !IsNil(o.EstimatedCurrentTimeflowsSavingsPerc) {
		return true
	}

	return false
}

// SetEstimatedCurrentTimeflowsSavingsPerc gets a reference to the given float32 and assigns it to the EstimatedCurrentTimeflowsSavingsPerc field.
func (o *StorageSavingsReportSummarizedData) SetEstimatedCurrentTimeflowsSavingsPerc(v float32) {
	o.EstimatedCurrentTimeflowsSavingsPerc = &v
}

func (o StorageSavingsReportSummarizedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageSavingsReportSummarizedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbCount) {
		toSerialize["vdb_count"] = o.VdbCount
	}
	if !IsNil(o.DsourceCount) {
		toSerialize["dsource_count"] = o.DsourceCount
	}
	if !IsNil(o.VirtualizedSpace) {
		toSerialize["virtualized_space"] = o.VirtualizedSpace
	}
	if !IsNil(o.UnvirtualizedSpace) {
		toSerialize["unvirtualized_space"] = o.UnvirtualizedSpace
	}
	if !IsNil(o.CurrentTimeflowsUnvirtualizedSpace) {
		toSerialize["current_timeflows_unvirtualized_space"] = o.CurrentTimeflowsUnvirtualizedSpace
	}
	if !IsNil(o.EstimatedSavings) {
		toSerialize["estimated_savings"] = o.EstimatedSavings
	}
	if !IsNil(o.EstimatedSavingsPerc) {
		toSerialize["estimated_savings_perc"] = o.EstimatedSavingsPerc
	}
	if !IsNil(o.EstimatedCurrentTimeflowsSavings) {
		toSerialize["estimated_current_timeflows_savings"] = o.EstimatedCurrentTimeflowsSavings
	}
	if !IsNil(o.EstimatedCurrentTimeflowsSavingsPerc) {
		toSerialize["estimated_current_timeflows_savings_perc"] = o.EstimatedCurrentTimeflowsSavingsPerc
	}
	return toSerialize, nil
}

type NullableStorageSavingsReportSummarizedData struct {
	value *StorageSavingsReportSummarizedData
	isSet bool
}

func (v NullableStorageSavingsReportSummarizedData) Get() *StorageSavingsReportSummarizedData {
	return v.value
}

func (v *NullableStorageSavingsReportSummarizedData) Set(val *StorageSavingsReportSummarizedData) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageSavingsReportSummarizedData) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageSavingsReportSummarizedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageSavingsReportSummarizedData(val *StorageSavingsReportSummarizedData) *NullableStorageSavingsReportSummarizedData {
	return &NullableStorageSavingsReportSummarizedData{value: val, isSet: true}
}

func (v NullableStorageSavingsReportSummarizedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageSavingsReportSummarizedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


