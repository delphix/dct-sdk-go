# PasswordPoliciesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | True if password policies are enforced/enabled. | [optional] 
**MinLength** | Pointer to **int32** | Minimum length for password. | [optional] 
**ReuseDisallowLimit** | Pointer to **int32** | The password can not be the same as any of the previous n passwords. | [optional] 
**Digit** | Pointer to **bool** | Mandate at least one digit in password. | [optional] 
**UppercaseLetter** | Pointer to **bool** | Mandate at least one uppercase letter in password. | [optional] 
**LowercaseLetter** | Pointer to **bool** | Mandate at least one lower letter in password. | [optional] 
**SpecialCharacter** | Pointer to **bool** | Mandate at least one special character in password. | [optional] 
**DisallowUsernameAsPassword** | Pointer to **bool** | Disallows password containing case-insensitive user name or reversed user name. | [optional] 
**MaximumPasswordAttempts** | Pointer to **int32** | The number of allowed attempts for incorrect password, after which the account gets locked. | [optional] [default to 0]

## Methods

### NewPasswordPoliciesParams

`func NewPasswordPoliciesParams() *PasswordPoliciesParams`

NewPasswordPoliciesParams instantiates a new PasswordPoliciesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPoliciesParamsWithDefaults

`func NewPasswordPoliciesParamsWithDefaults() *PasswordPoliciesParams`

NewPasswordPoliciesParamsWithDefaults instantiates a new PasswordPoliciesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PasswordPoliciesParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PasswordPoliciesParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PasswordPoliciesParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PasswordPoliciesParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMinLength

`func (o *PasswordPoliciesParams) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PasswordPoliciesParams) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PasswordPoliciesParams) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PasswordPoliciesParams) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetReuseDisallowLimit

`func (o *PasswordPoliciesParams) GetReuseDisallowLimit() int32`

GetReuseDisallowLimit returns the ReuseDisallowLimit field if non-nil, zero value otherwise.

### GetReuseDisallowLimitOk

`func (o *PasswordPoliciesParams) GetReuseDisallowLimitOk() (*int32, bool)`

GetReuseDisallowLimitOk returns a tuple with the ReuseDisallowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseDisallowLimit

`func (o *PasswordPoliciesParams) SetReuseDisallowLimit(v int32)`

SetReuseDisallowLimit sets ReuseDisallowLimit field to given value.

### HasReuseDisallowLimit

`func (o *PasswordPoliciesParams) HasReuseDisallowLimit() bool`

HasReuseDisallowLimit returns a boolean if a field has been set.

### GetDigit

`func (o *PasswordPoliciesParams) GetDigit() bool`

GetDigit returns the Digit field if non-nil, zero value otherwise.

### GetDigitOk

`func (o *PasswordPoliciesParams) GetDigitOk() (*bool, bool)`

GetDigitOk returns a tuple with the Digit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigit

`func (o *PasswordPoliciesParams) SetDigit(v bool)`

SetDigit sets Digit field to given value.

### HasDigit

`func (o *PasswordPoliciesParams) HasDigit() bool`

HasDigit returns a boolean if a field has been set.

### GetUppercaseLetter

`func (o *PasswordPoliciesParams) GetUppercaseLetter() bool`

GetUppercaseLetter returns the UppercaseLetter field if non-nil, zero value otherwise.

### GetUppercaseLetterOk

`func (o *PasswordPoliciesParams) GetUppercaseLetterOk() (*bool, bool)`

GetUppercaseLetterOk returns a tuple with the UppercaseLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUppercaseLetter

`func (o *PasswordPoliciesParams) SetUppercaseLetter(v bool)`

SetUppercaseLetter sets UppercaseLetter field to given value.

### HasUppercaseLetter

`func (o *PasswordPoliciesParams) HasUppercaseLetter() bool`

HasUppercaseLetter returns a boolean if a field has been set.

### GetLowercaseLetter

`func (o *PasswordPoliciesParams) GetLowercaseLetter() bool`

GetLowercaseLetter returns the LowercaseLetter field if non-nil, zero value otherwise.

### GetLowercaseLetterOk

`func (o *PasswordPoliciesParams) GetLowercaseLetterOk() (*bool, bool)`

GetLowercaseLetterOk returns a tuple with the LowercaseLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowercaseLetter

`func (o *PasswordPoliciesParams) SetLowercaseLetter(v bool)`

SetLowercaseLetter sets LowercaseLetter field to given value.

### HasLowercaseLetter

`func (o *PasswordPoliciesParams) HasLowercaseLetter() bool`

HasLowercaseLetter returns a boolean if a field has been set.

### GetSpecialCharacter

`func (o *PasswordPoliciesParams) GetSpecialCharacter() bool`

GetSpecialCharacter returns the SpecialCharacter field if non-nil, zero value otherwise.

### GetSpecialCharacterOk

`func (o *PasswordPoliciesParams) GetSpecialCharacterOk() (*bool, bool)`

GetSpecialCharacterOk returns a tuple with the SpecialCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialCharacter

`func (o *PasswordPoliciesParams) SetSpecialCharacter(v bool)`

SetSpecialCharacter sets SpecialCharacter field to given value.

### HasSpecialCharacter

`func (o *PasswordPoliciesParams) HasSpecialCharacter() bool`

HasSpecialCharacter returns a boolean if a field has been set.

### GetDisallowUsernameAsPassword

`func (o *PasswordPoliciesParams) GetDisallowUsernameAsPassword() bool`

GetDisallowUsernameAsPassword returns the DisallowUsernameAsPassword field if non-nil, zero value otherwise.

### GetDisallowUsernameAsPasswordOk

`func (o *PasswordPoliciesParams) GetDisallowUsernameAsPasswordOk() (*bool, bool)`

GetDisallowUsernameAsPasswordOk returns a tuple with the DisallowUsernameAsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowUsernameAsPassword

`func (o *PasswordPoliciesParams) SetDisallowUsernameAsPassword(v bool)`

SetDisallowUsernameAsPassword sets DisallowUsernameAsPassword field to given value.

### HasDisallowUsernameAsPassword

`func (o *PasswordPoliciesParams) HasDisallowUsernameAsPassword() bool`

HasDisallowUsernameAsPassword returns a boolean if a field has been set.

### GetMaximumPasswordAttempts

`func (o *PasswordPoliciesParams) GetMaximumPasswordAttempts() int32`

GetMaximumPasswordAttempts returns the MaximumPasswordAttempts field if non-nil, zero value otherwise.

### GetMaximumPasswordAttemptsOk

`func (o *PasswordPoliciesParams) GetMaximumPasswordAttemptsOk() (*int32, bool)`

GetMaximumPasswordAttemptsOk returns a tuple with the MaximumPasswordAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPasswordAttempts

`func (o *PasswordPoliciesParams) SetMaximumPasswordAttempts(v int32)`

SetMaximumPasswordAttempts sets MaximumPasswordAttempts field to given value.

### HasMaximumPasswordAttempts

`func (o *PasswordPoliciesParams) HasMaximumPasswordAttempts() bool`

HasMaximumPasswordAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


