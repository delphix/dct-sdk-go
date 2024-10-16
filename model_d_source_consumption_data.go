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
	"time"
)

// checks if the DSourceConsumptionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DSourceConsumptionData{}

// DSourceConsumptionData struct for DSourceConsumptionData
type DSourceConsumptionData struct {
	// The name of the dSource
	Name *string `json:"name,omitempty"`
	// The status of the dSource
	Status *string `json:"status,omitempty"`
	// The type of the dSource
	DatabaseType *string `json:"database_type,omitempty"`
	// The id of the engine the dSource belongs to.
	EngineId *string `json:"engine_id,omitempty"`
	// The name of the engine the dSource belongs to
	EngineName *string `json:"engine_name,omitempty"`
	// The most recent date where consumption data was captured for this dSource.
	LastConsumptionDate *time.Time `json:"last_consumption_date,omitempty"`
	// The ingested size of the dSource
	IngestedSize *int64 `json:"ingested_size,omitempty"`
}

// NewDSourceConsumptionData instantiates a new DSourceConsumptionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSourceConsumptionData() *DSourceConsumptionData {
	this := DSourceConsumptionData{}
	return &this
}

// NewDSourceConsumptionDataWithDefaults instantiates a new DSourceConsumptionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSourceConsumptionDataWithDefaults() *DSourceConsumptionData {
	this := DSourceConsumptionData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DSourceConsumptionData) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DSourceConsumptionData) SetStatus(v string) {
	o.Status = &v
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetDatabaseType() string {
	if o == nil || IsNil(o.DatabaseType) {
		var ret string
		return ret
	}
	return *o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetDatabaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseType) {
		return nil, false
	}
	return o.DatabaseType, true
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasDatabaseType() bool {
	if o != nil && !IsNil(o.DatabaseType) {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given string and assigns it to the DatabaseType field.
func (o *DSourceConsumptionData) SetDatabaseType(v string) {
	o.DatabaseType = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *DSourceConsumptionData) SetEngineId(v string) {
	o.EngineId = &v
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *DSourceConsumptionData) SetEngineName(v string) {
	o.EngineName = &v
}

// GetLastConsumptionDate returns the LastConsumptionDate field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetLastConsumptionDate() time.Time {
	if o == nil || IsNil(o.LastConsumptionDate) {
		var ret time.Time
		return ret
	}
	return *o.LastConsumptionDate
}

// GetLastConsumptionDateOk returns a tuple with the LastConsumptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetLastConsumptionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastConsumptionDate) {
		return nil, false
	}
	return o.LastConsumptionDate, true
}

// HasLastConsumptionDate returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasLastConsumptionDate() bool {
	if o != nil && !IsNil(o.LastConsumptionDate) {
		return true
	}

	return false
}

// SetLastConsumptionDate gets a reference to the given time.Time and assigns it to the LastConsumptionDate field.
func (o *DSourceConsumptionData) SetLastConsumptionDate(v time.Time) {
	o.LastConsumptionDate = &v
}

// GetIngestedSize returns the IngestedSize field value if set, zero value otherwise.
func (o *DSourceConsumptionData) GetIngestedSize() int64 {
	if o == nil || IsNil(o.IngestedSize) {
		var ret int64
		return ret
	}
	return *o.IngestedSize
}

// GetIngestedSizeOk returns a tuple with the IngestedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSourceConsumptionData) GetIngestedSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.IngestedSize) {
		return nil, false
	}
	return o.IngestedSize, true
}

// HasIngestedSize returns a boolean if a field has been set.
func (o *DSourceConsumptionData) HasIngestedSize() bool {
	if o != nil && !IsNil(o.IngestedSize) {
		return true
	}

	return false
}

// SetIngestedSize gets a reference to the given int64 and assigns it to the IngestedSize field.
func (o *DSourceConsumptionData) SetIngestedSize(v int64) {
	o.IngestedSize = &v
}

func (o DSourceConsumptionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSourceConsumptionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.DatabaseType) {
		toSerialize["database_type"] = o.DatabaseType
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.LastConsumptionDate) {
		toSerialize["last_consumption_date"] = o.LastConsumptionDate
	}
	if !IsNil(o.IngestedSize) {
		toSerialize["ingested_size"] = o.IngestedSize
	}
	return toSerialize, nil
}

type NullableDSourceConsumptionData struct {
	value *DSourceConsumptionData
	isSet bool
}

func (v NullableDSourceConsumptionData) Get() *DSourceConsumptionData {
	return v.value
}

func (v *NullableDSourceConsumptionData) Set(val *DSourceConsumptionData) {
	v.value = val
	v.isSet = true
}

func (v NullableDSourceConsumptionData) IsSet() bool {
	return v.isSet
}

func (v *NullableDSourceConsumptionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSourceConsumptionData(val *DSourceConsumptionData) *NullableDSourceConsumptionData {
	return &NullableDSourceConsumptionData{value: val, isSet: true}
}

func (v NullableDSourceConsumptionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSourceConsumptionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


