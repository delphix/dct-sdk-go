# LockVDBParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** | Id of the account on whose behalf this request is being made. Only accounts having LOCK_FOR_OTHER_ACCOUNT permission can lock VDBs on behalf of other accounts. If this property is not specified then the account id of the authenticated user making the request is used. | [optional] 

## Methods

### NewLockVDBParameters

`func NewLockVDBParameters() *LockVDBParameters`

NewLockVDBParameters instantiates a new LockVDBParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockVDBParametersWithDefaults

`func NewLockVDBParametersWithDefaults() *LockVDBParameters`

NewLockVDBParametersWithDefaults instantiates a new LockVDBParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LockVDBParameters) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LockVDBParameters) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LockVDBParameters) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LockVDBParameters) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


