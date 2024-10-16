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

// checks if the DSourceUsageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DSourceUsageData{}

// DSourceUsageData struct for DSourceUsageData
type DSourceUsageData struct {
	// The name of the engine the dSource belongs to.
	EngineName *string `json:"engine_name,omitempty"`
	// The name of the dSource
	Name *string `json:"name,omitempty"`
	// The the disk space that would be required if not using Delphix virtualizion, in bytes.
	UnvirtualizedSpace *int64 `json:"unvirtualized_space,omitempty"`
	// The actual space used by this dSource, in bytes. This includes the disk space used by the current copy of the data as well as the snapshots and log files retained to enable provisioning from historical versions.
	ActualSpace *int64 `json:"actual_space,omitempty"`
	// The number of VDBs that are dependant on this dSource. This includes all VDB descendants that have this dSource as an ancestor.
	DependantVdbs *int32 `json:"dependant_vdbs,omitempty"`
}

// NewDSourceUsageData instantiates a new DSourceUsageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSourceUsageData() *DSourceUsageData {
	this := DSourceUsageData{}
	return &this
}

// NewDSourceUsageDataWithDefaults instantiates a new DSourceUsageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSourceUsageDataWithDefaults() *DSourceUsageData {
	this := DSourceUsageData{}
	return &this
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *DSourceUsageData) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceUsageData) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *DSourceUsageData) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *DSourceUsageData) SetEngineName(v string) {
	o.EngineName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DSourceUsageData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceUsageData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DSourceUsageData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DSourceUsageData) SetName(v string) {
	o.Name = &v
}

// GetUnvirtualizedSpace returns the UnvirtualizedSpace field value if set, zero value otherwise.
func (o *DSourceUsageData) GetUnvirtualizedSpace() int64 {
	if o == nil || IsNil(o.UnvirtualizedSpace) {
		var ret int64
		return ret
	}
	return *o.UnvirtualizedSpace
}

// GetUnvirtualizedSpaceOk returns a tuple with the UnvirtualizedSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceUsageData) GetUnvirtualizedSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.UnvirtualizedSpace) {
		return nil, false
	}
	return o.UnvirtualizedSpace, true
}

// HasUnvirtualizedSpace returns a boolean if a field has been set.
func (o *DSourceUsageData) HasUnvirtualizedSpace() bool {
	if o != nil && !IsNil(o.UnvirtualizedSpace) {
		return true
	}

	return false
}

// SetUnvirtualizedSpace gets a reference to the given int64 and assigns it to the UnvirtualizedSpace field.
func (o *DSourceUsageData) SetUnvirtualizedSpace(v int64) {
	o.UnvirtualizedSpace = &v
}

// GetActualSpace returns the ActualSpace field value if set, zero value otherwise.
func (o *DSourceUsageData) GetActualSpace() int64 {
	if o == nil || IsNil(o.ActualSpace) {
		var ret int64
		return ret
	}
	return *o.ActualSpace
}

// GetActualSpaceOk returns a tuple with the ActualSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceUsageData) GetActualSpaceOk() (*int64, bool) {
	if o == nil || IsNil(o.ActualSpace) {
		return nil, false
	}
	return o.ActualSpace, true
}

// HasActualSpace returns a boolean if a field has been set.
func (o *DSourceUsageData) HasActualSpace() bool {
	if o != nil && !IsNil(o.ActualSpace) {
		return true
	}

	return false
}

// SetActualSpace gets a reference to the given int64 and assigns it to the ActualSpace field.
func (o *DSourceUsageData) SetActualSpace(v int64) {
	o.ActualSpace = &v
}

// GetDependantVdbs returns the DependantVdbs field value if set, zero value otherwise.
func (o *DSourceUsageData) GetDependantVdbs() int32 {
	if o == nil || IsNil(o.DependantVdbs) {
		var ret int32
		return ret
	}
	return *o.DependantVdbs
}

// GetDependantVdbsOk returns a tuple with the DependantVdbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceUsageData) GetDependantVdbsOk() (*int32, bool) {
	if o == nil || IsNil(o.DependantVdbs) {
		return nil, false
	}
	return o.DependantVdbs, true
}

// HasDependantVdbs returns a boolean if a field has been set.
func (o *DSourceUsageData) HasDependantVdbs() bool {
	if o != nil && !IsNil(o.DependantVdbs) {
		return true
	}

	return false
}

// SetDependantVdbs gets a reference to the given int32 and assigns it to the DependantVdbs field.
func (o *DSourceUsageData) SetDependantVdbs(v int32) {
	o.DependantVdbs = &v
}

func (o DSourceUsageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSourceUsageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UnvirtualizedSpace) {
		toSerialize["unvirtualized_space"] = o.UnvirtualizedSpace
	}
	if !IsNil(o.ActualSpace) {
		toSerialize["actual_space"] = o.ActualSpace
	}
	if !IsNil(o.DependantVdbs) {
		toSerialize["dependant_vdbs"] = o.DependantVdbs
	}
	return toSerialize, nil
}

type NullableDSourceUsageData struct {
	value *DSourceUsageData
	isSet bool
}

func (v NullableDSourceUsageData) Get() *DSourceUsageData {
	return v.value
}

func (v *NullableDSourceUsageData) Set(val *DSourceUsageData) {
	v.value = val
	v.isSet = true
}

func (v NullableDSourceUsageData) IsSet() bool {
	return v.isSet
}

func (v *NullableDSourceUsageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSourceUsageData(val *DSourceUsageData) *NullableDSourceUsageData {
	return &NullableDSourceUsageData{value: val, isSet: true}
}

func (v NullableDSourceUsageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSourceUsageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


