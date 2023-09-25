# HyperscaleInstanceRegistrationParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name in DCT of the Hyperscale instance. | 
**Hostname** | **string** | Hostname of the Hyperscale instance. | 
**ApiKey** | **NullableString** | API key to connect to the Hyperscale instance. | 
**InsecureSsl** | Pointer to **bool** | Allow connections to the hyperscale instance over HTTPs without validating the TLS certificate. Even though the connection to the hyperscale instance might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the hyperscale instance&#39;s certificate, and set the truststore_filename property.  | [optional] [default to false]
**UnsafeSslHostnameCheck** | Pointer to **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the hyperscale instance over HTTPs. Setting this value must only be done if the TLS certificate of the hyperscale instance does not match the hostname, and the TLS configuration of the hyperscale instance cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set.  | [optional] [default to false]
**TruststoreFilename** | Pointer to **NullableString** | File name of a truststore which can be used to validate the TLS certificate of the hyperscale instance. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt;  | [optional] 
**TruststorePassword** | Pointer to **NullableString** | Password to read the truststore.  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | The tags to be created for this engine. | [optional] 

## Methods

### NewHyperscaleInstanceRegistrationParameter

`func NewHyperscaleInstanceRegistrationParameter(name string, hostname string, apiKey NullableString, ) *HyperscaleInstanceRegistrationParameter`

NewHyperscaleInstanceRegistrationParameter instantiates a new HyperscaleInstanceRegistrationParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleInstanceRegistrationParameterWithDefaults

`func NewHyperscaleInstanceRegistrationParameterWithDefaults() *HyperscaleInstanceRegistrationParameter`

NewHyperscaleInstanceRegistrationParameterWithDefaults instantiates a new HyperscaleInstanceRegistrationParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HyperscaleInstanceRegistrationParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperscaleInstanceRegistrationParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperscaleInstanceRegistrationParameter) SetName(v string)`

SetName sets Name field to given value.


### GetHostname

`func (o *HyperscaleInstanceRegistrationParameter) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HyperscaleInstanceRegistrationParameter) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HyperscaleInstanceRegistrationParameter) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetApiKey

`func (o *HyperscaleInstanceRegistrationParameter) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HyperscaleInstanceRegistrationParameter) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HyperscaleInstanceRegistrationParameter) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *HyperscaleInstanceRegistrationParameter) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *HyperscaleInstanceRegistrationParameter) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetInsecureSsl

`func (o *HyperscaleInstanceRegistrationParameter) GetInsecureSsl() bool`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *HyperscaleInstanceRegistrationParameter) GetInsecureSslOk() (*bool, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *HyperscaleInstanceRegistrationParameter) SetInsecureSsl(v bool)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *HyperscaleInstanceRegistrationParameter) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetUnsafeSslHostnameCheck

`func (o *HyperscaleInstanceRegistrationParameter) GetUnsafeSslHostnameCheck() bool`

GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field if non-nil, zero value otherwise.

### GetUnsafeSslHostnameCheckOk

`func (o *HyperscaleInstanceRegistrationParameter) GetUnsafeSslHostnameCheckOk() (*bool, bool)`

GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeSslHostnameCheck

`func (o *HyperscaleInstanceRegistrationParameter) SetUnsafeSslHostnameCheck(v bool)`

SetUnsafeSslHostnameCheck sets UnsafeSslHostnameCheck field to given value.

### HasUnsafeSslHostnameCheck

`func (o *HyperscaleInstanceRegistrationParameter) HasUnsafeSslHostnameCheck() bool`

HasUnsafeSslHostnameCheck returns a boolean if a field has been set.

### GetTruststoreFilename

`func (o *HyperscaleInstanceRegistrationParameter) GetTruststoreFilename() string`

GetTruststoreFilename returns the TruststoreFilename field if non-nil, zero value otherwise.

### GetTruststoreFilenameOk

`func (o *HyperscaleInstanceRegistrationParameter) GetTruststoreFilenameOk() (*string, bool)`

GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFilename

`func (o *HyperscaleInstanceRegistrationParameter) SetTruststoreFilename(v string)`

SetTruststoreFilename sets TruststoreFilename field to given value.

### HasTruststoreFilename

`func (o *HyperscaleInstanceRegistrationParameter) HasTruststoreFilename() bool`

HasTruststoreFilename returns a boolean if a field has been set.

### SetTruststoreFilenameNil

`func (o *HyperscaleInstanceRegistrationParameter) SetTruststoreFilenameNil(b bool)`

 SetTruststoreFilenameNil sets the value for TruststoreFilename to be an explicit nil

### UnsetTruststoreFilename
`func (o *HyperscaleInstanceRegistrationParameter) UnsetTruststoreFilename()`

UnsetTruststoreFilename ensures that no value is present for TruststoreFilename, not even an explicit nil
### GetTruststorePassword

`func (o *HyperscaleInstanceRegistrationParameter) GetTruststorePassword() string`

GetTruststorePassword returns the TruststorePassword field if non-nil, zero value otherwise.

### GetTruststorePasswordOk

`func (o *HyperscaleInstanceRegistrationParameter) GetTruststorePasswordOk() (*string, bool)`

GetTruststorePasswordOk returns a tuple with the TruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststorePassword

`func (o *HyperscaleInstanceRegistrationParameter) SetTruststorePassword(v string)`

SetTruststorePassword sets TruststorePassword field to given value.

### HasTruststorePassword

`func (o *HyperscaleInstanceRegistrationParameter) HasTruststorePassword() bool`

HasTruststorePassword returns a boolean if a field has been set.

### SetTruststorePasswordNil

`func (o *HyperscaleInstanceRegistrationParameter) SetTruststorePasswordNil(b bool)`

 SetTruststorePasswordNil sets the value for TruststorePassword to be an explicit nil

### UnsetTruststorePassword
`func (o *HyperscaleInstanceRegistrationParameter) UnsetTruststorePassword()`

UnsetTruststorePassword ensures that no value is present for TruststorePassword, not even an explicit nil
### GetTags

`func (o *HyperscaleInstanceRegistrationParameter) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperscaleInstanceRegistrationParameter) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperscaleInstanceRegistrationParameter) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperscaleInstanceRegistrationParameter) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


