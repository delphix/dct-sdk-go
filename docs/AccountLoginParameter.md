# AccountLoginParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username of the account that needs to login | 
**Password** | **string** | Password of the account that needs to login. | 

## Methods

### NewAccountLoginParameter

`func NewAccountLoginParameter(username string, password string, ) *AccountLoginParameter`

NewAccountLoginParameter instantiates a new AccountLoginParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountLoginParameterWithDefaults

`func NewAccountLoginParameterWithDefaults() *AccountLoginParameter`

NewAccountLoginParameterWithDefaults instantiates a new AccountLoginParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AccountLoginParameter) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountLoginParameter) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountLoginParameter) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AccountLoginParameter) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccountLoginParameter) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccountLoginParameter) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


