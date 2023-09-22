# ResetPasswordParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | Pointer to **string** | New password that needs to be set for the Account. Set this to null for unsetting the current password. Not including this property also results in unsetting of the current password. | [optional] 

## Methods

### NewResetPasswordParameter

`func NewResetPasswordParameter() *ResetPasswordParameter`

NewResetPasswordParameter instantiates a new ResetPasswordParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordParameterWithDefaults

`func NewResetPasswordParameterWithDefaults() *ResetPasswordParameter`

NewResetPasswordParameterWithDefaults instantiates a new ResetPasswordParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *ResetPasswordParameter) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ResetPasswordParameter) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ResetPasswordParameter) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *ResetPasswordParameter) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


