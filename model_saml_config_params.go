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

// checks if the SAMLConfigParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLConfigParams{}

// SAMLConfigParams Parameters to read or update SAML Config
type SAMLConfigParams struct {
	// When set, SAML settings are enabled. False by default.
	Enabled *bool `json:"enabled,omitempty"`
	// When set, the system will automatically create new Accounts for those who have logged in using SAML. This must be true if SAML user is not already registered in system. True by default.
	AutoCreateUsers *bool `json:"auto_create_users,omitempty"`
	// IdP metadata for this service provider. This is a required property for successful SAML authentication.
	Metadata *string `json:"metadata,omitempty"`
	// Unique identifier of this instance as a SAML/SSO service provider.
	EntityId *string `json:"entity_id,omitempty"`
	// Maximum time difference allowed between a SAML response and the DCT's current time, in seconds. If not set, it defaults to 120 seconds.
	ResponseSkew *int32 `json:"response_skew,omitempty"`
	// Group mapped attribute on SAML to create account tags in DCT.
	GroupAttr *string `json:"group_attr,omitempty"`
	// First name attribute mapped on SAML used for mapping on DCT account.
	FirstNameAttr *string `json:"first_name_attr,omitempty"`
	// Last name attribute mapped on SAML used for mapping on DCT account.
	LastNameAttr *string `json:"last_name_attr,omitempty"`
}

// NewSAMLConfigParams instantiates a new SAMLConfigParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLConfigParams() *SAMLConfigParams {
	this := SAMLConfigParams{}
	var enabled bool = false
	this.Enabled = &enabled
	var autoCreateUsers bool = true
	this.AutoCreateUsers = &autoCreateUsers
	var responseSkew int32 = 120
	this.ResponseSkew = &responseSkew
	var groupAttr string = "groups"
	this.GroupAttr = &groupAttr
	var firstNameAttr string = "firstName"
	this.FirstNameAttr = &firstNameAttr
	var lastNameAttr string = "lastName"
	this.LastNameAttr = &lastNameAttr
	return &this
}

// NewSAMLConfigParamsWithDefaults instantiates a new SAMLConfigParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLConfigParamsWithDefaults() *SAMLConfigParams {
	this := SAMLConfigParams{}
	var enabled bool = false
	this.Enabled = &enabled
	var autoCreateUsers bool = true
	this.AutoCreateUsers = &autoCreateUsers
	var responseSkew int32 = 120
	this.ResponseSkew = &responseSkew
	var groupAttr string = "groups"
	this.GroupAttr = &groupAttr
	var firstNameAttr string = "firstName"
	this.FirstNameAttr = &firstNameAttr
	var lastNameAttr string = "lastName"
	this.LastNameAttr = &lastNameAttr
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SAMLConfigParams) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAutoCreateUsers returns the AutoCreateUsers field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetAutoCreateUsers() bool {
	if o == nil || IsNil(o.AutoCreateUsers) {
		var ret bool
		return ret
	}
	return *o.AutoCreateUsers
}

// GetAutoCreateUsersOk returns a tuple with the AutoCreateUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetAutoCreateUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCreateUsers) {
		return nil, false
	}
	return o.AutoCreateUsers, true
}

// HasAutoCreateUsers returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasAutoCreateUsers() bool {
	if o != nil && !IsNil(o.AutoCreateUsers) {
		return true
	}

	return false
}

// SetAutoCreateUsers gets a reference to the given bool and assigns it to the AutoCreateUsers field.
func (o *SAMLConfigParams) SetAutoCreateUsers(v bool) {
	o.AutoCreateUsers = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetMetadata() string {
	if o == nil || IsNil(o.Metadata) {
		var ret string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given string and assigns it to the Metadata field.
func (o *SAMLConfigParams) SetMetadata(v string) {
	o.Metadata = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SAMLConfigParams) SetEntityId(v string) {
	o.EntityId = &v
}

// GetResponseSkew returns the ResponseSkew field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetResponseSkew() int32 {
	if o == nil || IsNil(o.ResponseSkew) {
		var ret int32
		return ret
	}
	return *o.ResponseSkew
}

// GetResponseSkewOk returns a tuple with the ResponseSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetResponseSkewOk() (*int32, bool) {
	if o == nil || IsNil(o.ResponseSkew) {
		return nil, false
	}
	return o.ResponseSkew, true
}

// HasResponseSkew returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasResponseSkew() bool {
	if o != nil && !IsNil(o.ResponseSkew) {
		return true
	}

	return false
}

// SetResponseSkew gets a reference to the given int32 and assigns it to the ResponseSkew field.
func (o *SAMLConfigParams) SetResponseSkew(v int32) {
	o.ResponseSkew = &v
}

// GetGroupAttr returns the GroupAttr field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetGroupAttr() string {
	if o == nil || IsNil(o.GroupAttr) {
		var ret string
		return ret
	}
	return *o.GroupAttr
}

// GetGroupAttrOk returns a tuple with the GroupAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetGroupAttrOk() (*string, bool) {
	if o == nil || IsNil(o.GroupAttr) {
		return nil, false
	}
	return o.GroupAttr, true
}

// HasGroupAttr returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasGroupAttr() bool {
	if o != nil && !IsNil(o.GroupAttr) {
		return true
	}

	return false
}

// SetGroupAttr gets a reference to the given string and assigns it to the GroupAttr field.
func (o *SAMLConfigParams) SetGroupAttr(v string) {
	o.GroupAttr = &v
}

// GetFirstNameAttr returns the FirstNameAttr field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetFirstNameAttr() string {
	if o == nil || IsNil(o.FirstNameAttr) {
		var ret string
		return ret
	}
	return *o.FirstNameAttr
}

// GetFirstNameAttrOk returns a tuple with the FirstNameAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetFirstNameAttrOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNameAttr) {
		return nil, false
	}
	return o.FirstNameAttr, true
}

// HasFirstNameAttr returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasFirstNameAttr() bool {
	if o != nil && !IsNil(o.FirstNameAttr) {
		return true
	}

	return false
}

// SetFirstNameAttr gets a reference to the given string and assigns it to the FirstNameAttr field.
func (o *SAMLConfigParams) SetFirstNameAttr(v string) {
	o.FirstNameAttr = &v
}

// GetLastNameAttr returns the LastNameAttr field value if set, zero value otherwise.
func (o *SAMLConfigParams) GetLastNameAttr() string {
	if o == nil || IsNil(o.LastNameAttr) {
		var ret string
		return ret
	}
	return *o.LastNameAttr
}

// GetLastNameAttrOk returns a tuple with the LastNameAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLConfigParams) GetLastNameAttrOk() (*string, bool) {
	if o == nil || IsNil(o.LastNameAttr) {
		return nil, false
	}
	return o.LastNameAttr, true
}

// HasLastNameAttr returns a boolean if a field has been set.
func (o *SAMLConfigParams) HasLastNameAttr() bool {
	if o != nil && !IsNil(o.LastNameAttr) {
		return true
	}

	return false
}

// SetLastNameAttr gets a reference to the given string and assigns it to the LastNameAttr field.
func (o *SAMLConfigParams) SetLastNameAttr(v string) {
	o.LastNameAttr = &v
}

func (o SAMLConfigParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLConfigParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.AutoCreateUsers) {
		toSerialize["auto_create_users"] = o.AutoCreateUsers
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.EntityId) {
		toSerialize["entity_id"] = o.EntityId
	}
	if !IsNil(o.ResponseSkew) {
		toSerialize["response_skew"] = o.ResponseSkew
	}
	if !IsNil(o.GroupAttr) {
		toSerialize["group_attr"] = o.GroupAttr
	}
	if !IsNil(o.FirstNameAttr) {
		toSerialize["first_name_attr"] = o.FirstNameAttr
	}
	if !IsNil(o.LastNameAttr) {
		toSerialize["last_name_attr"] = o.LastNameAttr
	}
	return toSerialize, nil
}

type NullableSAMLConfigParams struct {
	value *SAMLConfigParams
	isSet bool
}

func (v NullableSAMLConfigParams) Get() *SAMLConfigParams {
	return v.value
}

func (v *NullableSAMLConfigParams) Set(val *SAMLConfigParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLConfigParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLConfigParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLConfigParams(val *SAMLConfigParams) *NullableSAMLConfigParams {
	return &NullableSAMLConfigParams{value: val, isSet: true}
}

func (v NullableSAMLConfigParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLConfigParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


