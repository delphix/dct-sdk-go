# AccessGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Access group ID. | [optional] [readonly] 
**Name** | **string** | The Access group name | 
**SingleAccount** | Pointer to **bool** | Indicates that this Access group defines the permissions of a single account, and thus account and account tags cannot be modified. Instead create a new Access group to manage permissions of multiple accounts. | [optional] 
**AccountIds** | Pointer to **[]int64** | List of accounts ids included individually (as opposed to added by tags) in the Access group. | [optional] 
**TaggedAccountIds** | Pointer to **[]int64** | List of accounts ids included by tags in the Access group. | [optional] [readonly] 
**AccountTags** | Pointer to [**[]Tag**](Tag.md) | List of account tags. Accounts matching any of these tags will be automatically added to the Access group. | [optional] 
**Scopes** | Pointer to [**[]AccessGroupScope**](AccessGroupScope.md) | The Access group scopes. | [optional] 

## Methods

### NewAccessGroup

`func NewAccessGroup(name string, ) *AccessGroup`

NewAccessGroup instantiates a new AccessGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessGroupWithDefaults

`func NewAccessGroupWithDefaults() *AccessGroup`

NewAccessGroupWithDefaults instantiates a new AccessGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccessGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccessGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSingleAccount

`func (o *AccessGroup) GetSingleAccount() bool`

GetSingleAccount returns the SingleAccount field if non-nil, zero value otherwise.

### GetSingleAccountOk

`func (o *AccessGroup) GetSingleAccountOk() (*bool, bool)`

GetSingleAccountOk returns a tuple with the SingleAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleAccount

`func (o *AccessGroup) SetSingleAccount(v bool)`

SetSingleAccount sets SingleAccount field to given value.

### HasSingleAccount

`func (o *AccessGroup) HasSingleAccount() bool`

HasSingleAccount returns a boolean if a field has been set.

### GetAccountIds

`func (o *AccessGroup) GetAccountIds() []int64`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *AccessGroup) GetAccountIdsOk() (*[]int64, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *AccessGroup) SetAccountIds(v []int64)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *AccessGroup) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetTaggedAccountIds

`func (o *AccessGroup) GetTaggedAccountIds() []int64`

GetTaggedAccountIds returns the TaggedAccountIds field if non-nil, zero value otherwise.

### GetTaggedAccountIdsOk

`func (o *AccessGroup) GetTaggedAccountIdsOk() (*[]int64, bool)`

GetTaggedAccountIdsOk returns a tuple with the TaggedAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaggedAccountIds

`func (o *AccessGroup) SetTaggedAccountIds(v []int64)`

SetTaggedAccountIds sets TaggedAccountIds field to given value.

### HasTaggedAccountIds

`func (o *AccessGroup) HasTaggedAccountIds() bool`

HasTaggedAccountIds returns a boolean if a field has been set.

### GetAccountTags

`func (o *AccessGroup) GetAccountTags() []Tag`

GetAccountTags returns the AccountTags field if non-nil, zero value otherwise.

### GetAccountTagsOk

`func (o *AccessGroup) GetAccountTagsOk() (*[]Tag, bool)`

GetAccountTagsOk returns a tuple with the AccountTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTags

`func (o *AccessGroup) SetAccountTags(v []Tag)`

SetAccountTags sets AccountTags field to given value.

### HasAccountTags

`func (o *AccessGroup) HasAccountTags() bool`

HasAccountTags returns a boolean if a field has been set.

### GetScopes

`func (o *AccessGroup) GetScopes() []AccessGroupScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessGroup) GetScopesOk() (*[]AccessGroupScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessGroup) SetScopes(v []AccessGroupScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AccessGroup) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


