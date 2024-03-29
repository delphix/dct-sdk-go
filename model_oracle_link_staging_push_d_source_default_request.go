/*
Delphix DCT API

Delphix DCT API

API version: 3.9.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the OracleLinkStagingPushDSourceDefaultRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OracleLinkStagingPushDSourceDefaultRequest{}

// OracleLinkStagingPushDSourceDefaultRequest struct for OracleLinkStagingPushDSourceDefaultRequest
type OracleLinkStagingPushDSourceDefaultRequest struct {
	// The ID of the environment to be linked.
	EnvironmentId string `json:"environment_id"`
	// The container type of this database.If not provided the request would be considered for a PDB database.
	ContainerType *string `json:"container_type,omitempty"`
}

// NewOracleLinkStagingPushDSourceDefaultRequest instantiates a new OracleLinkStagingPushDSourceDefaultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleLinkStagingPushDSourceDefaultRequest(environmentId string) *OracleLinkStagingPushDSourceDefaultRequest {
	this := OracleLinkStagingPushDSourceDefaultRequest{}
	this.EnvironmentId = environmentId
	return &this
}

// NewOracleLinkStagingPushDSourceDefaultRequestWithDefaults instantiates a new OracleLinkStagingPushDSourceDefaultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleLinkStagingPushDSourceDefaultRequestWithDefaults() *OracleLinkStagingPushDSourceDefaultRequest {
	this := OracleLinkStagingPushDSourceDefaultRequest{}
	return &this
}

// GetEnvironmentId returns the EnvironmentId field value
func (o *OracleLinkStagingPushDSourceDefaultRequest) GetEnvironmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value
// and a boolean to check if the value has been set.
func (o *OracleLinkStagingPushDSourceDefaultRequest) GetEnvironmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentId, true
}

// SetEnvironmentId sets field value
func (o *OracleLinkStagingPushDSourceDefaultRequest) SetEnvironmentId(v string) {
	o.EnvironmentId = v
}

// GetContainerType returns the ContainerType field value if set, zero value otherwise.
func (o *OracleLinkStagingPushDSourceDefaultRequest) GetContainerType() string {
	if o == nil || IsNil(o.ContainerType) {
		var ret string
		return ret
	}
	return *o.ContainerType
}

// GetContainerTypeOk returns a tuple with the ContainerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleLinkStagingPushDSourceDefaultRequest) GetContainerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerType) {
		return nil, false
	}
	return o.ContainerType, true
}

// HasContainerType returns a boolean if a field has been set.
func (o *OracleLinkStagingPushDSourceDefaultRequest) HasContainerType() bool {
	if o != nil && !IsNil(o.ContainerType) {
		return true
	}

	return false
}

// SetContainerType gets a reference to the given string and assigns it to the ContainerType field.
func (o *OracleLinkStagingPushDSourceDefaultRequest) SetContainerType(v string) {
	o.ContainerType = &v
}

func (o OracleLinkStagingPushDSourceDefaultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OracleLinkStagingPushDSourceDefaultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment_id"] = o.EnvironmentId
	if !IsNil(o.ContainerType) {
		toSerialize["container_type"] = o.ContainerType
	}
	return toSerialize, nil
}

type NullableOracleLinkStagingPushDSourceDefaultRequest struct {
	value *OracleLinkStagingPushDSourceDefaultRequest
	isSet bool
}

func (v NullableOracleLinkStagingPushDSourceDefaultRequest) Get() *OracleLinkStagingPushDSourceDefaultRequest {
	return v.value
}

func (v *NullableOracleLinkStagingPushDSourceDefaultRequest) Set(val *OracleLinkStagingPushDSourceDefaultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleLinkStagingPushDSourceDefaultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleLinkStagingPushDSourceDefaultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleLinkStagingPushDSourceDefaultRequest(val *OracleLinkStagingPushDSourceDefaultRequest) *NullableOracleLinkStagingPushDSourceDefaultRequest {
	return &NullableOracleLinkStagingPushDSourceDefaultRequest{value: val, isSet: true}
}

func (v NullableOracleLinkStagingPushDSourceDefaultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleLinkStagingPushDSourceDefaultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


