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

// checks if the TokenInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenInfoResponse{}

// TokenInfoResponse struct for TokenInfoResponse
type TokenInfoResponse struct {
	// Flag to identify if the token is active.
	Active *bool `json:"active,omitempty"`
	// Type of the token.
	TokenType *string `json:"token_type,omitempty"`
	// Numeric ID of the account.
	AccountId *int64 `json:"account_id,omitempty"`
	// First name for the Account.
	FirstName *string `json:"first_name,omitempty"`
	// Last name for the Account.
	LastName *string `json:"last_name,omitempty"`
	// Email for the Account.
	Email *string `json:"email,omitempty"`
	// The username or logical name for the Account.
	Username *string `json:"username,omitempty"`
	// The LDAP Principal for the Account.
	LdapPrincipal *string `json:"ldap_principal,omitempty"`
	// Seconds duration after which the token will expire.
	Exp *int64 `json:"exp,omitempty"`
}

// NewTokenInfoResponse instantiates a new TokenInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenInfoResponse() *TokenInfoResponse {
	this := TokenInfoResponse{}
	return &this
}

// NewTokenInfoResponseWithDefaults instantiates a new TokenInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenInfoResponseWithDefaults() *TokenInfoResponse {
	this := TokenInfoResponse{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TokenInfoResponse) SetActive(v bool) {
	o.Active = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *TokenInfoResponse) SetTokenType(v string) {
	o.TokenType = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *TokenInfoResponse) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *TokenInfoResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *TokenInfoResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *TokenInfoResponse) SetEmail(v string) {
	o.Email = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *TokenInfoResponse) SetUsername(v string) {
	o.Username = &v
}

// GetLdapPrincipal returns the LdapPrincipal field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetLdapPrincipal() string {
	if o == nil || IsNil(o.LdapPrincipal) {
		var ret string
		return ret
	}
	return *o.LdapPrincipal
}

// GetLdapPrincipalOk returns a tuple with the LdapPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetLdapPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.LdapPrincipal) {
		return nil, false
	}
	return o.LdapPrincipal, true
}

// HasLdapPrincipal returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasLdapPrincipal() bool {
	if o != nil && !IsNil(o.LdapPrincipal) {
		return true
	}

	return false
}

// SetLdapPrincipal gets a reference to the given string and assigns it to the LdapPrincipal field.
func (o *TokenInfoResponse) SetLdapPrincipal(v string) {
	o.LdapPrincipal = &v
}

// GetExp returns the Exp field value if set, zero value otherwise.
func (o *TokenInfoResponse) GetExp() int64 {
	if o == nil || IsNil(o.Exp) {
		var ret int64
		return ret
	}
	return *o.Exp
}

// GetExpOk returns a tuple with the Exp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenInfoResponse) GetExpOk() (*int64, bool) {
	if o == nil || IsNil(o.Exp) {
		return nil, false
	}
	return o.Exp, true
}

// HasExp returns a boolean if a field has been set.
func (o *TokenInfoResponse) HasExp() bool {
	if o != nil && !IsNil(o.Exp) {
		return true
	}

	return false
}

// SetExp gets a reference to the given int64 and assigns it to the Exp field.
func (o *TokenInfoResponse) SetExp(v int64) {
	o.Exp = &v
}

func (o TokenInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.LdapPrincipal) {
		toSerialize["ldap_principal"] = o.LdapPrincipal
	}
	if !IsNil(o.Exp) {
		toSerialize["exp"] = o.Exp
	}
	return toSerialize, nil
}

type NullableTokenInfoResponse struct {
	value *TokenInfoResponse
	isSet bool
}

func (v NullableTokenInfoResponse) Get() *TokenInfoResponse {
	return v.value
}

func (v *NullableTokenInfoResponse) Set(val *TokenInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenInfoResponse(val *TokenInfoResponse) *NullableTokenInfoResponse {
	return &NullableTokenInfoResponse{value: val, isSet: true}
}

func (v NullableTokenInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


