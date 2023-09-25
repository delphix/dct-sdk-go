/*
Delphix DCT API

Delphix DCT API

API version: 3.5.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the ValidateJavaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateJavaResponse{}

// ValidateJavaResponse The result of the validating java path on remote host.
type ValidateJavaResponse struct {
	// A message describing the result of the connectivity check.
	Message string `json:"message"`
	// A status describing the status of the connectivity check.
	Status *string `json:"status,omitempty"`
}

// NewValidateJavaResponse instantiates a new ValidateJavaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateJavaResponse(message string) *ValidateJavaResponse {
	this := ValidateJavaResponse{}
	this.Message = message
	return &this
}

// NewValidateJavaResponseWithDefaults instantiates a new ValidateJavaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateJavaResponseWithDefaults() *ValidateJavaResponse {
	this := ValidateJavaResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *ValidateJavaResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ValidateJavaResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ValidateJavaResponse) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ValidateJavaResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateJavaResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ValidateJavaResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ValidateJavaResponse) SetStatus(v string) {
	o.Status = &v
}

func (o ValidateJavaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateJavaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableValidateJavaResponse struct {
	value *ValidateJavaResponse
	isSet bool
}

func (v NullableValidateJavaResponse) Get() *ValidateJavaResponse {
	return v.value
}

func (v *NullableValidateJavaResponse) Set(val *ValidateJavaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateJavaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateJavaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateJavaResponse(val *ValidateJavaResponse) *NullableValidateJavaResponse {
	return &NullableValidateJavaResponse{value: val, isSet: true}
}

func (v NullableValidateJavaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateJavaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

