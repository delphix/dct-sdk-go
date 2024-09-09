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

// checks if the ProvisionVDBByLocationDefaultsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionVDBByLocationDefaultsRequest{}

// ProvisionVDBByLocationDefaultsRequest struct for ProvisionVDBByLocationDefaultsRequest
type ProvisionVDBByLocationDefaultsRequest struct {
	// The ID of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source.
	SourceDataId *string `json:"source_data_id,omitempty"`
	// The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.
	EngineId *string `json:"engine_id,omitempty"`
	// The location to get the defaults from.
	Location *string `json:"location,omitempty"`
	// ID of the timeflow to provision from.
	TimeflowId *string `json:"timeflow_id,omitempty"`
}

// NewProvisionVDBByLocationDefaultsRequest instantiates a new ProvisionVDBByLocationDefaultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionVDBByLocationDefaultsRequest() *ProvisionVDBByLocationDefaultsRequest {
	this := ProvisionVDBByLocationDefaultsRequest{}
	return &this
}

// NewProvisionVDBByLocationDefaultsRequestWithDefaults instantiates a new ProvisionVDBByLocationDefaultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionVDBByLocationDefaultsRequestWithDefaults() *ProvisionVDBByLocationDefaultsRequest {
	this := ProvisionVDBByLocationDefaultsRequest{}
	return &this
}

// GetSourceDataId returns the SourceDataId field value if set, zero value otherwise.
func (o *ProvisionVDBByLocationDefaultsRequest) GetSourceDataId() string {
	if o == nil || IsNil(o.SourceDataId) {
		var ret string
		return ret
	}
	return *o.SourceDataId
}

// GetSourceDataIdOk returns a tuple with the SourceDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) GetSourceDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDataId) {
		return nil, false
	}
	return o.SourceDataId, true
}

// HasSourceDataId returns a boolean if a field has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) HasSourceDataId() bool {
	if o != nil && !IsNil(o.SourceDataId) {
		return true
	}

	return false
}

// SetSourceDataId gets a reference to the given string and assigns it to the SourceDataId field.
func (o *ProvisionVDBByLocationDefaultsRequest) SetSourceDataId(v string) {
	o.SourceDataId = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *ProvisionVDBByLocationDefaultsRequest) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *ProvisionVDBByLocationDefaultsRequest) SetEngineId(v string) {
	o.EngineId = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ProvisionVDBByLocationDefaultsRequest) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ProvisionVDBByLocationDefaultsRequest) SetLocation(v string) {
	o.Location = &v
}

// GetTimeflowId returns the TimeflowId field value if set, zero value otherwise.
func (o *ProvisionVDBByLocationDefaultsRequest) GetTimeflowId() string {
	if o == nil || IsNil(o.TimeflowId) {
		var ret string
		return ret
	}
	return *o.TimeflowId
}

// GetTimeflowIdOk returns a tuple with the TimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) GetTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeflowId) {
		return nil, false
	}
	return o.TimeflowId, true
}

// HasTimeflowId returns a boolean if a field has been set.
func (o *ProvisionVDBByLocationDefaultsRequest) HasTimeflowId() bool {
	if o != nil && !IsNil(o.TimeflowId) {
		return true
	}

	return false
}

// SetTimeflowId gets a reference to the given string and assigns it to the TimeflowId field.
func (o *ProvisionVDBByLocationDefaultsRequest) SetTimeflowId(v string) {
	o.TimeflowId = &v
}

func (o ProvisionVDBByLocationDefaultsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionVDBByLocationDefaultsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceDataId) {
		toSerialize["source_data_id"] = o.SourceDataId
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.TimeflowId) {
		toSerialize["timeflow_id"] = o.TimeflowId
	}
	return toSerialize, nil
}

type NullableProvisionVDBByLocationDefaultsRequest struct {
	value *ProvisionVDBByLocationDefaultsRequest
	isSet bool
}

func (v NullableProvisionVDBByLocationDefaultsRequest) Get() *ProvisionVDBByLocationDefaultsRequest {
	return v.value
}

func (v *NullableProvisionVDBByLocationDefaultsRequest) Set(val *ProvisionVDBByLocationDefaultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionVDBByLocationDefaultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionVDBByLocationDefaultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionVDBByLocationDefaultsRequest(val *ProvisionVDBByLocationDefaultsRequest) *NullableProvisionVDBByLocationDefaultsRequest {
	return &NullableProvisionVDBByLocationDefaultsRequest{value: val, isSet: true}
}

func (v NullableProvisionVDBByLocationDefaultsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionVDBByLocationDefaultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


