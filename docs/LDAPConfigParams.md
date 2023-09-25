# LDAPConfigParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | When set, these settings are enabled. True by default. | [optional] [default to true]
**AutoCreateUsers** | Pointer to **bool** | When set, the system will automatically create new Accounts for those who have logged in using LDAP. This must be true if LDAP user is not already registered in system. True by default. | [optional] [default to true]
**Hostname** | Pointer to **string** | The hostname of the LDAP server. | [optional] 
**Port** | Pointer to **int32** | The port of the LDAP server. Default port is 389 for non-SSL and 636 for SSL. | [optional] 
**Domains** | Pointer to [**[]Domain**](Domain.md) | DCT will try to authenticate using each Domain given in this list. | [optional] 
**EnableSsl** | Pointer to **bool** | True if LDAP should be used over SSL. | [optional] [default to true]
**TruststoreFilename** | Pointer to **string** | File name of a truststore which can be used to validate the TLS certificate of the LDAP server. The truststore must be available at /etc/config/certs/&lt;truststore_filename&gt; | [optional] 
**TruststorePassword** | Pointer to **string** | Password for reading trustStore file provided in &#39;truststore_filename&#39; property | [optional] 
**InsecureSsl** | Pointer to **bool** | Allow connections to the LDAP server over LDAPS without validating the TLS certificate. Even though the connection to the server might be performed over LDAPS, setting this property eliminates the protection against a man-in-the-middle attach for connections to this server. Instead, consider creating a truststore with a Certificate Authority to validate the server&#39;s certificate, and set the truststore_filename property.  | [optional] [default to false]
**UnsafeSslHostnameCheck** | Pointer to **bool** | Ignore validation of the name associated to the TLS certificate when connecting to the LDAP server over LDAPS. Setting this value must only be done if the TLS certificate of the server does not match the hostname, and the TLS configuration of the server cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this server. This is ignored if insecure_ssl is set.  | [optional] [default to false]

## Methods

### NewLDAPConfigParams

`func NewLDAPConfigParams() *LDAPConfigParams`

NewLDAPConfigParams instantiates a new LDAPConfigParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLDAPConfigParamsWithDefaults

`func NewLDAPConfigParamsWithDefaults() *LDAPConfigParams`

NewLDAPConfigParamsWithDefaults instantiates a new LDAPConfigParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LDAPConfigParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LDAPConfigParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LDAPConfigParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LDAPConfigParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateUsers

`func (o *LDAPConfigParams) GetAutoCreateUsers() bool`

GetAutoCreateUsers returns the AutoCreateUsers field if non-nil, zero value otherwise.

### GetAutoCreateUsersOk

`func (o *LDAPConfigParams) GetAutoCreateUsersOk() (*bool, bool)`

GetAutoCreateUsersOk returns a tuple with the AutoCreateUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateUsers

`func (o *LDAPConfigParams) SetAutoCreateUsers(v bool)`

SetAutoCreateUsers sets AutoCreateUsers field to given value.

### HasAutoCreateUsers

`func (o *LDAPConfigParams) HasAutoCreateUsers() bool`

HasAutoCreateUsers returns a boolean if a field has been set.

### GetHostname

`func (o *LDAPConfigParams) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *LDAPConfigParams) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *LDAPConfigParams) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *LDAPConfigParams) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *LDAPConfigParams) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LDAPConfigParams) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LDAPConfigParams) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LDAPConfigParams) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDomains

`func (o *LDAPConfigParams) GetDomains() []Domain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *LDAPConfigParams) GetDomainsOk() (*[]Domain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *LDAPConfigParams) SetDomains(v []Domain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *LDAPConfigParams) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetEnableSsl

`func (o *LDAPConfigParams) GetEnableSsl() bool`

GetEnableSsl returns the EnableSsl field if non-nil, zero value otherwise.

### GetEnableSslOk

`func (o *LDAPConfigParams) GetEnableSslOk() (*bool, bool)`

GetEnableSslOk returns a tuple with the EnableSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSsl

`func (o *LDAPConfigParams) SetEnableSsl(v bool)`

SetEnableSsl sets EnableSsl field to given value.

### HasEnableSsl

`func (o *LDAPConfigParams) HasEnableSsl() bool`

HasEnableSsl returns a boolean if a field has been set.

### GetTruststoreFilename

`func (o *LDAPConfigParams) GetTruststoreFilename() string`

GetTruststoreFilename returns the TruststoreFilename field if non-nil, zero value otherwise.

### GetTruststoreFilenameOk

`func (o *LDAPConfigParams) GetTruststoreFilenameOk() (*string, bool)`

GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststoreFilename

`func (o *LDAPConfigParams) SetTruststoreFilename(v string)`

SetTruststoreFilename sets TruststoreFilename field to given value.

### HasTruststoreFilename

`func (o *LDAPConfigParams) HasTruststoreFilename() bool`

HasTruststoreFilename returns a boolean if a field has been set.

### GetTruststorePassword

`func (o *LDAPConfigParams) GetTruststorePassword() string`

GetTruststorePassword returns the TruststorePassword field if non-nil, zero value otherwise.

### GetTruststorePasswordOk

`func (o *LDAPConfigParams) GetTruststorePasswordOk() (*string, bool)`

GetTruststorePasswordOk returns a tuple with the TruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruststorePassword

`func (o *LDAPConfigParams) SetTruststorePassword(v string)`

SetTruststorePassword sets TruststorePassword field to given value.

### HasTruststorePassword

`func (o *LDAPConfigParams) HasTruststorePassword() bool`

HasTruststorePassword returns a boolean if a field has been set.

### GetInsecureSsl

`func (o *LDAPConfigParams) GetInsecureSsl() bool`

GetInsecureSsl returns the InsecureSsl field if non-nil, zero value otherwise.

### GetInsecureSslOk

`func (o *LDAPConfigParams) GetInsecureSslOk() (*bool, bool)`

GetInsecureSslOk returns a tuple with the InsecureSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSsl

`func (o *LDAPConfigParams) SetInsecureSsl(v bool)`

SetInsecureSsl sets InsecureSsl field to given value.

### HasInsecureSsl

`func (o *LDAPConfigParams) HasInsecureSsl() bool`

HasInsecureSsl returns a boolean if a field has been set.

### GetUnsafeSslHostnameCheck

`func (o *LDAPConfigParams) GetUnsafeSslHostnameCheck() bool`

GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field if non-nil, zero value otherwise.

### GetUnsafeSslHostnameCheckOk

`func (o *LDAPConfigParams) GetUnsafeSslHostnameCheckOk() (*bool, bool)`

GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsafeSslHostnameCheck

`func (o *LDAPConfigParams) SetUnsafeSslHostnameCheck(v bool)`

SetUnsafeSslHostnameCheck sets UnsafeSslHostnameCheck field to given value.

### HasUnsafeSslHostnameCheck

`func (o *LDAPConfigParams) HasUnsafeSslHostnameCheck() bool`

HasUnsafeSslHostnameCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


