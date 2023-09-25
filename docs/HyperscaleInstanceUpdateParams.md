# HyperscaleInstanceUpdateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name in DCT of the Hyperscale instance. | [optional] 
**Hostname** | Pointer to **string** | Hostname of the Hyperscale instance. | [optional] 
**ApiKey** | Pointer to **NullableString** | API key to connect to the Hyperscale instance. | [optional] 
**InsecureSsl** | Pointer to **bool** | Allow connections to the hyperscale instance over HTTPs without validating the TLS certificate. Even though the connection to the hyperscale instance might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the hyperscale instance&#39;s certificate, and set the truststore_filename property.  | [optional] 
**UnsafeSslHostnameCheck** | Pointer to **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the hyperscale instance over HTTPs. Setting this value must only be done if the TLS certificate of the hyperscale instance does not match the hostname, and the TLS configuration of the hyperscale instance cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set.  | [optional] 
**TruststoreFilename** | Pointer to **NullableString** | File name of a truststore which can be used to validate the TLS certificate of the hyperscale instance. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt;. Set this property to an empty string to clear the value.  | [optional] 
**TruststorePassword** | Pointer to **NullableString** | Password to read the truststore. Set this property to an empty string to clear the value.  | [optional] 

## Methods

### NewHyperscaleInstanceUpdateParams

`func NewHyperscaleInstanceUpdateParams() *HyperscaleInstanceUpdateParams`

NewHyperscaleInstanceUpdateParams instantiates a new HyperscaleInstanceUpdateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleInstanceUpdateParamsWithDefaults

`func NewHyperscaleInstanceUpdateParamsWithDefaults() *HyperscaleInstanceUpdateParams`

NewHyperscaleInstanceUpdateParamsWithDefaults instantiates a new HyperscaleInstanceUpdateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HyperscaleInstanceUpdateParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperscaleInstanceUpdateParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperscaleInstanceUpdateParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperscaleInstanceUpdateParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *HyperscaleInstanceUpdateParams) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HyperscaleInstanceUpdateParams) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HyperscaleInstanceUpdateParams) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HyperscaleInstanceUpdateParams) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetApiKey

`func (o *HyperscaleInstanceUpdateParams) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *HyperscaleInstanceUpdateParams) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *HyperscaleInstanceUpdateParams) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *HyperscaleInstanceUpdateParams) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *HyperscaleInstanceUpdateParams) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *HyperscaleInstanceUpdateParams) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetInsecureSsl

`func (o *HyperscaleInstanceUpdateParams) GetInsecureSsl() bool`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *HyperscaleInstanceUpdateParams) GetInsecureSslOk() (*bool, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *HyperscaleInstanceUpdateParams) SetInsecureSsl(v bool)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *HyperscaleInstanceUpdateParams) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetUnsafeSslHostnameCheck

`func (o *HyperscaleInstanceUpdateParams) GetUnsafeSslHostnameCheck() bool`

GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field if non-nil, zero value otherwise.

### GetUnsafeSslHostnameCheckOk

`func (o *HyperscaleInstanceUpdateParams) GetUnsafeSslHostnameCheckOk() (*bool, bool)`

GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeSslHostnameCheck

`func (o *HyperscaleInstanceUpdateParams) SetUnsafeSslHostnameCheck(v bool)`

SetUnsafeSslHostnameCheck sets UnsafeSslHostnameCheck field to given value.

### HasUnsafeSslHostnameCheck

`func (o *HyperscaleInstanceUpdateParams) HasUnsafeSslHostnameCheck() bool`

HasUnsafeSslHostnameCheck returns a boolean if a field has been set.

### GetTruststoreFilename

`func (o *HyperscaleInstanceUpdateParams) GetTruststoreFilename() string`

GetTruststoreFilename returns the TruststoreFilename field if non-nil, zero value otherwise.

### GetTruststoreFilenameOk

`func (o *HyperscaleInstanceUpdateParams) GetTruststoreFilenameOk() (*string, bool)`

GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFilename

`func (o *HyperscaleInstanceUpdateParams) SetTruststoreFilename(v string)`

SetTruststoreFilename sets TruststoreFilename field to given value.

### HasTruststoreFilename

`func (o *HyperscaleInstanceUpdateParams) HasTruststoreFilename() bool`

HasTruststoreFilename returns a boolean if a field has been set.

### SetTruststoreFilenameNil

`func (o *HyperscaleInstanceUpdateParams) SetTruststoreFilenameNil(b bool)`

 SetTruststoreFilenameNil sets the value for TruststoreFilename to be an explicit nil

### UnsetTruststoreFilename
`func (o *HyperscaleInstanceUpdateParams) UnsetTruststoreFilename()`

UnsetTruststoreFilename ensures that no value is present for TruststoreFilename, not even an explicit nil
### GetTruststorePassword

`func (o *HyperscaleInstanceUpdateParams) GetTruststorePassword() string`

GetTruststorePassword returns the TruststorePassword field if non-nil, zero value otherwise.

### GetTruststorePasswordOk

`func (o *HyperscaleInstanceUpdateParams) GetTruststorePasswordOk() (*string, bool)`

GetTruststorePasswordOk returns a tuple with the TruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststorePassword

`func (o *HyperscaleInstanceUpdateParams) SetTruststorePassword(v string)`

SetTruststorePassword sets TruststorePassword field to given value.

### HasTruststorePassword

`func (o *HyperscaleInstanceUpdateParams) HasTruststorePassword() bool`

HasTruststorePassword returns a boolean if a field has been set.

### SetTruststorePasswordNil

`func (o *HyperscaleInstanceUpdateParams) SetTruststorePasswordNil(b bool)`

 SetTruststorePasswordNil sets the value for TruststorePassword to be an explicit nil

### UnsetTruststorePassword
`func (o *HyperscaleInstanceUpdateParams) UnsetTruststorePassword()`

UnsetTruststorePassword ensures that no value is present for TruststorePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


