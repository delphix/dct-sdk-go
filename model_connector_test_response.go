/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ConnectorTestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorTestResponse{}

// ConnectorTestResponse The result of the masking connector test.
type ConnectorTestResponse struct {
	// Connection status, SUCCEEDED or FAILED
	Status string `json:"status"`
	// A message describing the result of the masking connector test.
	Message string `json:"message"`
}

// NewConnectorTestResponse instantiates a new ConnectorTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorTestResponse(status string, message string) *ConnectorTestResponse {
	this := ConnectorTestResponse{}
	this.Status = status
	this.Message = message
	return &this
}

// NewConnectorTestResponseWithDefaults instantiates a new ConnectorTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorTestResponseWithDefaults() *ConnectorTestResponse {
	this := ConnectorTestResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *ConnectorTestResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ConnectorTestResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ConnectorTestResponse) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *ConnectorTestResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ConnectorTestResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ConnectorTestResponse) SetMessage(v string) {
	o.Message = v
}

func (o ConnectorTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorTestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

type NullableConnectorTestResponse struct {
	value *ConnectorTestResponse
	isSet bool
}

func (v NullableConnectorTestResponse) Get() *ConnectorTestResponse {
	return v.value
}

func (v *NullableConnectorTestResponse) Set(val *ConnectorTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorTestResponse(val *ConnectorTestResponse) *NullableConnectorTestResponse {
	return &NullableConnectorTestResponse{value: val, isSet: true}
}

func (v NullableConnectorTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


