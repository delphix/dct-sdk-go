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

// checks if the Domain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Domain{}

// Domain struct for Domain
type Domain struct {
	// This is used to get full DN for authentication and search. Provide this value only if server is microsoft AD.
	MsadDomainName *string `json:"msad_domain_name,omitempty"`
	// The username_patterns can be used to avoid providing full-dn during login. This will also be used for search of groups,email, first name and last name.
	UsernamePattern *string `json:"username_pattern,omitempty"`
	// Search base used to search for ldap user groups. Leave this field empty if a full username_pattern is provided and server is non microsoft AD.
	SearchBase *string `json:"search_base,omitempty"`
	// Group mapped attribute on ldap side used for user group search.
	GroupAttr *string `json:"group_attr,omitempty"`
	// Email mapped attribute on ldap side used for mapping on DCT side account.
	EmailAttr *string `json:"email_attr,omitempty"`
	// First name attribute mapped on ldap side used for mapping on DCT side account.
	FirstNameAttr *string `json:"first_name_attr,omitempty"`
	// Last name attribute mapped on ldap side used for mapping on DCT side account.
	LastNameAttr *string `json:"last_name_attr,omitempty"`
	// The name of the objectClass on ldap side under which the user is mapped.This is used to search for the user details.
	ObjectClassAttr *string `json:"object_class_attr,omitempty"`
	// Search attribute mapped on ldap side using which search on ldap side will be made.
	SearchAttr *string `json:"search_attr,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetMsadDomainName returns the MsadDomainName field value if set, zero value otherwise.
func (o *Domain) GetMsadDomainName() string {
	if o == nil || IsNil(o.MsadDomainName) {
		var ret string
		return ret
	}
	return *o.MsadDomainName
}

// GetMsadDomainNameOk returns a tuple with the MsadDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetMsadDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.MsadDomainName) {
		return nil, false
	}
	return o.MsadDomainName, true
}

// HasMsadDomainName returns a boolean if a field has been set.
func (o *Domain) HasMsadDomainName() bool {
	if o != nil && !IsNil(o.MsadDomainName) {
		return true
	}

	return false
}

// SetMsadDomainName gets a reference to the given string and assigns it to the MsadDomainName field.
func (o *Domain) SetMsadDomainName(v string) {
	o.MsadDomainName = &v
}

// GetUsernamePattern returns the UsernamePattern field value if set, zero value otherwise.
func (o *Domain) GetUsernamePattern() string {
	if o == nil || IsNil(o.UsernamePattern) {
		var ret string
		return ret
	}
	return *o.UsernamePattern
}

// GetUsernamePatternOk returns a tuple with the UsernamePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetUsernamePatternOk() (*string, bool) {
	if o == nil || IsNil(o.UsernamePattern) {
		return nil, false
	}
	return o.UsernamePattern, true
}

// HasUsernamePattern returns a boolean if a field has been set.
func (o *Domain) HasUsernamePattern() bool {
	if o != nil && !IsNil(o.UsernamePattern) {
		return true
	}

	return false
}

// SetUsernamePattern gets a reference to the given string and assigns it to the UsernamePattern field.
func (o *Domain) SetUsernamePattern(v string) {
	o.UsernamePattern = &v
}

// GetSearchBase returns the SearchBase field value if set, zero value otherwise.
func (o *Domain) GetSearchBase() string {
	if o == nil || IsNil(o.SearchBase) {
		var ret string
		return ret
	}
	return *o.SearchBase
}

// GetSearchBaseOk returns a tuple with the SearchBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSearchBaseOk() (*string, bool) {
	if o == nil || IsNil(o.SearchBase) {
		return nil, false
	}
	return o.SearchBase, true
}

// HasSearchBase returns a boolean if a field has been set.
func (o *Domain) HasSearchBase() bool {
	if o != nil && !IsNil(o.SearchBase) {
		return true
	}

	return false
}

// SetSearchBase gets a reference to the given string and assigns it to the SearchBase field.
func (o *Domain) SetSearchBase(v string) {
	o.SearchBase = &v
}

// GetGroupAttr returns the GroupAttr field value if set, zero value otherwise.
func (o *Domain) GetGroupAttr() string {
	if o == nil || IsNil(o.GroupAttr) {
		var ret string
		return ret
	}
	return *o.GroupAttr
}

// GetGroupAttrOk returns a tuple with the GroupAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetGroupAttrOk() (*string, bool) {
	if o == nil || IsNil(o.GroupAttr) {
		return nil, false
	}
	return o.GroupAttr, true
}

// HasGroupAttr returns a boolean if a field has been set.
func (o *Domain) HasGroupAttr() bool {
	if o != nil && !IsNil(o.GroupAttr) {
		return true
	}

	return false
}

// SetGroupAttr gets a reference to the given string and assigns it to the GroupAttr field.
func (o *Domain) SetGroupAttr(v string) {
	o.GroupAttr = &v
}

// GetEmailAttr returns the EmailAttr field value if set, zero value otherwise.
func (o *Domain) GetEmailAttr() string {
	if o == nil || IsNil(o.EmailAttr) {
		var ret string
		return ret
	}
	return *o.EmailAttr
}

// GetEmailAttrOk returns a tuple with the EmailAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetEmailAttrOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttr) {
		return nil, false
	}
	return o.EmailAttr, true
}

// HasEmailAttr returns a boolean if a field has been set.
func (o *Domain) HasEmailAttr() bool {
	if o != nil && !IsNil(o.EmailAttr) {
		return true
	}

	return false
}

// SetEmailAttr gets a reference to the given string and assigns it to the EmailAttr field.
func (o *Domain) SetEmailAttr(v string) {
	o.EmailAttr = &v
}

// GetFirstNameAttr returns the FirstNameAttr field value if set, zero value otherwise.
func (o *Domain) GetFirstNameAttr() string {
	if o == nil || IsNil(o.FirstNameAttr) {
		var ret string
		return ret
	}
	return *o.FirstNameAttr
}

// GetFirstNameAttrOk returns a tuple with the FirstNameAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetFirstNameAttrOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNameAttr) {
		return nil, false
	}
	return o.FirstNameAttr, true
}

// HasFirstNameAttr returns a boolean if a field has been set.
func (o *Domain) HasFirstNameAttr() bool {
	if o != nil && !IsNil(o.FirstNameAttr) {
		return true
	}

	return false
}

// SetFirstNameAttr gets a reference to the given string and assigns it to the FirstNameAttr field.
func (o *Domain) SetFirstNameAttr(v string) {
	o.FirstNameAttr = &v
}

// GetLastNameAttr returns the LastNameAttr field value if set, zero value otherwise.
func (o *Domain) GetLastNameAttr() string {
	if o == nil || IsNil(o.LastNameAttr) {
		var ret string
		return ret
	}
	return *o.LastNameAttr
}

// GetLastNameAttrOk returns a tuple with the LastNameAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetLastNameAttrOk() (*string, bool) {
	if o == nil || IsNil(o.LastNameAttr) {
		return nil, false
	}
	return o.LastNameAttr, true
}

// HasLastNameAttr returns a boolean if a field has been set.
func (o *Domain) HasLastNameAttr() bool {
	if o != nil && !IsNil(o.LastNameAttr) {
		return true
	}

	return false
}

// SetLastNameAttr gets a reference to the given string and assigns it to the LastNameAttr field.
func (o *Domain) SetLastNameAttr(v string) {
	o.LastNameAttr = &v
}

// GetObjectClassAttr returns the ObjectClassAttr field value if set, zero value otherwise.
func (o *Domain) GetObjectClassAttr() string {
	if o == nil || IsNil(o.ObjectClassAttr) {
		var ret string
		return ret
	}
	return *o.ObjectClassAttr
}

// GetObjectClassAttrOk returns a tuple with the ObjectClassAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetObjectClassAttrOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectClassAttr) {
		return nil, false
	}
	return o.ObjectClassAttr, true
}

// HasObjectClassAttr returns a boolean if a field has been set.
func (o *Domain) HasObjectClassAttr() bool {
	if o != nil && !IsNil(o.ObjectClassAttr) {
		return true
	}

	return false
}

// SetObjectClassAttr gets a reference to the given string and assigns it to the ObjectClassAttr field.
func (o *Domain) SetObjectClassAttr(v string) {
	o.ObjectClassAttr = &v
}

// GetSearchAttr returns the SearchAttr field value if set, zero value otherwise.
func (o *Domain) GetSearchAttr() string {
	if o == nil || IsNil(o.SearchAttr) {
		var ret string
		return ret
	}
	return *o.SearchAttr
}

// GetSearchAttrOk returns a tuple with the SearchAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetSearchAttrOk() (*string, bool) {
	if o == nil || IsNil(o.SearchAttr) {
		return nil, false
	}
	return o.SearchAttr, true
}

// HasSearchAttr returns a boolean if a field has been set.
func (o *Domain) HasSearchAttr() bool {
	if o != nil && !IsNil(o.SearchAttr) {
		return true
	}

	return false
}

// SetSearchAttr gets a reference to the given string and assigns it to the SearchAttr field.
func (o *Domain) SetSearchAttr(v string) {
	o.SearchAttr = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Domain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MsadDomainName) {
		toSerialize["msad_domain_name"] = o.MsadDomainName
	}
	if !IsNil(o.UsernamePattern) {
		toSerialize["username_pattern"] = o.UsernamePattern
	}
	if !IsNil(o.SearchBase) {
		toSerialize["search_base"] = o.SearchBase
	}
	if !IsNil(o.GroupAttr) {
		toSerialize["group_attr"] = o.GroupAttr
	}
	if !IsNil(o.EmailAttr) {
		toSerialize["email_attr"] = o.EmailAttr
	}
	if !IsNil(o.FirstNameAttr) {
		toSerialize["first_name_attr"] = o.FirstNameAttr
	}
	if !IsNil(o.LastNameAttr) {
		toSerialize["last_name_attr"] = o.LastNameAttr
	}
	if !IsNil(o.ObjectClassAttr) {
		toSerialize["object_class_attr"] = o.ObjectClassAttr
	}
	if !IsNil(o.SearchAttr) {
		toSerialize["search_attr"] = o.SearchAttr
	}
	return toSerialize, nil
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


