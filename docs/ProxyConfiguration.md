# ProxyConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The host name or IP address of the proxy server. | 
**Port** | **int32** | The port number of the proxy server. | 
**Username** | Pointer to **string** | The username to use when authenticating with the proxy server. | [optional] 
**Password** | Pointer to **string** | The password to use when authenticating with the proxy server. | [optional] 
**Enabled** | **bool** | When set, these settings are enabled. True by default. | 
**TruststoreFilename** | Pointer to **string** | File name of a truststore which can be used to validate the TLS certificate of the proxy server. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt; | [optional] 
**TruststorePassword** | Pointer to **string** | Password for reading trustStore file provided in &#39;truststore_filename&#39; property | [optional] 

## Methods

### NewProxyConfiguration

`func NewProxyConfiguration(host string, port int32, enabled bool, ) *ProxyConfiguration`

NewProxyConfiguration instantiates a new ProxyConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProxyConfigurationWithDefaults

`func NewProxyConfigurationWithDefaults() *ProxyConfiguration`

NewProxyConfigurationWithDefaults instantiates a new ProxyConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ProxyConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ProxyConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ProxyConfiguration) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *ProxyConfiguration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ProxyConfiguration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ProxyConfiguration) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *ProxyConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ProxyConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ProxyConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ProxyConfiguration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ProxyConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ProxyConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ProxyConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ProxyConfiguration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEnabled

`func (o *ProxyConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProxyConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProxyConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetTruststoreFilename

`func (o *ProxyConfiguration) GetTruststoreFilename() string`

GetTruststoreFilename returns the TruststoreFilename field if non-nil, zero value otherwise.

### GetTruststoreFilenameOk

`func (o *ProxyConfiguration) GetTruststoreFilenameOk() (*string, bool)`

GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFilename

`func (o *ProxyConfiguration) SetTruststoreFilename(v string)`

SetTruststoreFilename sets TruststoreFilename field to given value.

### HasTruststoreFilename

`func (o *ProxyConfiguration) HasTruststoreFilename() bool`

HasTruststoreFilename returns a boolean if a field has been set.

### GetTruststorePassword

`func (o *ProxyConfiguration) GetTruststorePassword() string`

GetTruststorePassword returns the TruststorePassword field if non-nil, zero value otherwise.

### GetTruststorePasswordOk

`func (o *ProxyConfiguration) GetTruststorePasswordOk() (*string, bool)`

GetTruststorePasswordOk returns a tuple with the TruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststorePassword

`func (o *ProxyConfiguration) SetTruststorePassword(v string)`

SetTruststorePassword sets TruststorePassword field to given value.

### HasTruststorePassword

`func (o *ProxyConfiguration) HasTruststorePassword() bool`

HasTruststorePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


