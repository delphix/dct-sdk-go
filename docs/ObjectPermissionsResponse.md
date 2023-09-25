# ObjectPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]ObjectPermissionAccount**](ObjectPermissionAccount.md) | The Accounts permitted for this object. | [optional] 

## Methods

### NewObjectPermissionsResponse

`func NewObjectPermissionsResponse() *ObjectPermissionsResponse`

NewObjectPermissionsResponse instantiates a new ObjectPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectPermissionsResponseWithDefaults

`func NewObjectPermissionsResponseWithDefaults() *ObjectPermissionsResponse`

NewObjectPermissionsResponseWithDefaults instantiates a new ObjectPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ObjectPermissionsResponse) GetAccounts() []ObjectPermissionAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ObjectPermissionsResponse) GetAccountsOk() (*[]ObjectPermissionAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ObjectPermissionsResponse) SetAccounts(v []ObjectPermissionAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ObjectPermissionsResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


