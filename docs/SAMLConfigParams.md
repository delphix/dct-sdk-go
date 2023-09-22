# SAMLConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | When set, SAML settings are enabled. False by default. | [optional] [default to false]
**AutoCreateUsers** | Pointer to **bool** | When set, the system will automatically create new Accounts for those who have logged in using SAML. This must be true if SAML user is not already registered in system. True by default. | [optional] [default to true]
**Metadata** | Pointer to **string** | IdP metadata for this service provider. This is a required property for successful SAML authentication. | [optional] 
**EntityId** | Pointer to **string** | Unique identifier of this instance as a SAML/SSO service provider. | [optional] 
**ResponseSkew** | Pointer to **int32** | Maximum time difference allowed between a SAML response and the DCT&#39;s current time, in seconds. If not set, it defaults to 120 seconds. | [optional] [default to 120]
**GroupAttr** | Pointer to **string** | Group mapped attribute on SAML to create account tags in DCT. | [optional] [default to "groups"]
**FirstNameAttr** | Pointer to **string** | First name attribute mapped on SAML used for mapping on DCT account. | [optional] [default to "firstName"]
**LastNameAttr** | Pointer to **string** | Last name attribute mapped on SAML used for mapping on DCT account. | [optional] [default to "lastName"]

## Methods

### NewSAMLConfigParams

`func NewSAMLConfigParams() *SAMLConfigParams`

NewSAMLConfigParams instantiates a new SAMLConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLConfigParamsWithDefaults

`func NewSAMLConfigParamsWithDefaults() *SAMLConfigParams`

NewSAMLConfigParamsWithDefaults instantiates a new SAMLConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SAMLConfigParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SAMLConfigParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SAMLConfigParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SAMLConfigParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateUsers

`func (o *SAMLConfigParams) GetAutoCreateUsers() bool`

GetAutoCreateUsers returns the AutoCreateUsers field if non-nil, zero value otherwise.

### GetAutoCreateUsersOk

`func (o *SAMLConfigParams) GetAutoCreateUsersOk() (*bool, bool)`

GetAutoCreateUsersOk returns a tuple with the AutoCreateUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateUsers

`func (o *SAMLConfigParams) SetAutoCreateUsers(v bool)`

SetAutoCreateUsers sets AutoCreateUsers field to given value.

### HasAutoCreateUsers

`func (o *SAMLConfigParams) HasAutoCreateUsers() bool`

HasAutoCreateUsers returns a boolean if a field has been set.

### GetMetadata

`func (o *SAMLConfigParams) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SAMLConfigParams) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SAMLConfigParams) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SAMLConfigParams) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetEntityId

`func (o *SAMLConfigParams) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SAMLConfigParams) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SAMLConfigParams) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SAMLConfigParams) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetResponseSkew

`func (o *SAMLConfigParams) GetResponseSkew() int32`

GetResponseSkew returns the ResponseSkew field if non-nil, zero value otherwise.

### GetResponseSkewOk

`func (o *SAMLConfigParams) GetResponseSkewOk() (*int32, bool)`

GetResponseSkewOk returns a tuple with the ResponseSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSkew

`func (o *SAMLConfigParams) SetResponseSkew(v int32)`

SetResponseSkew sets ResponseSkew field to given value.

### HasResponseSkew

`func (o *SAMLConfigParams) HasResponseSkew() bool`

HasResponseSkew returns a boolean if a field has been set.

### GetGroupAttr

`func (o *SAMLConfigParams) GetGroupAttr() string`

GetGroupAttr returns the GroupAttr field if non-nil, zero value otherwise.

### GetGroupAttrOk

`func (o *SAMLConfigParams) GetGroupAttrOk() (*string, bool)`

GetGroupAttrOk returns a tuple with the GroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAttr

`func (o *SAMLConfigParams) SetGroupAttr(v string)`

SetGroupAttr sets GroupAttr field to given value.

### HasGroupAttr

`func (o *SAMLConfigParams) HasGroupAttr() bool`

HasGroupAttr returns a boolean if a field has been set.

### GetFirstNameAttr

`func (o *SAMLConfigParams) GetFirstNameAttr() string`

GetFirstNameAttr returns the FirstNameAttr field if non-nil, zero value otherwise.

### GetFirstNameAttrOk

`func (o *SAMLConfigParams) GetFirstNameAttrOk() (*string, bool)`

GetFirstNameAttrOk returns a tuple with the FirstNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameAttr

`func (o *SAMLConfigParams) SetFirstNameAttr(v string)`

SetFirstNameAttr sets FirstNameAttr field to given value.

### HasFirstNameAttr

`func (o *SAMLConfigParams) HasFirstNameAttr() bool`

HasFirstNameAttr returns a boolean if a field has been set.

### GetLastNameAttr

`func (o *SAMLConfigParams) GetLastNameAttr() string`

GetLastNameAttr returns the LastNameAttr field if non-nil, zero value otherwise.

### GetLastNameAttrOk

`func (o *SAMLConfigParams) GetLastNameAttrOk() (*string, bool)`

GetLastNameAttrOk returns a tuple with the LastNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameAttr

`func (o *SAMLConfigParams) SetLastNameAttr(v string)`

SetLastNameAttr sets LastNameAttr field to given value.

### HasLastNameAttr

`func (o *SAMLConfigParams) HasLastNameAttr() bool`

HasLastNameAttr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


