# EnvironmentUpdateParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the environment. | [optional] 
**Hostname** | Pointer to **string** | host address of the machine. | [optional] 
**StagingEnvironment** | Pointer to **string** | Id of the connector environment which is used to connect to this source environment. | [optional] 
**ConnectorPort** | Pointer to **int32** | Specify port on which Delphix connector will run. | [optional] 
**NfsAddresses** | Pointer to **[]string** | array of ip addresses or hostnames | [optional] 
**SshPort** | Pointer to **int64** | ssh port of the host. | [optional] 
**ToolkitPath** | Pointer to **string** | The path for the toolkit that resides on the host. | [optional] 
**JavaHome** | Pointer to **string** | The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used. | [optional] 
**DspKeystorePath** | Pointer to **string** | DSP keystore path. | [optional] 
**DspKeystorePassword** | Pointer to **string** | DSP keystore password. | [optional] 
**DspKeystoreAlias** | Pointer to **string** | DSP keystore alias. | [optional] 
**DspTruststorePath** | Pointer to **string** | DSP truststore path. | [optional] 
**DspTruststorePassword** | Pointer to **string** | DSP truststore password. | [optional] 
**Description** | Pointer to **string** | The environment description. | [optional] 

## Methods

### NewEnvironmentUpdateParameters

`func NewEnvironmentUpdateParameters() *EnvironmentUpdateParameters`

NewEnvironmentUpdateParameters instantiates a new EnvironmentUpdateParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentUpdateParametersWithDefaults

`func NewEnvironmentUpdateParametersWithDefaults() *EnvironmentUpdateParameters`

NewEnvironmentUpdateParametersWithDefaults instantiates a new EnvironmentUpdateParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentUpdateParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentUpdateParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentUpdateParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentUpdateParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *EnvironmentUpdateParameters) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EnvironmentUpdateParameters) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EnvironmentUpdateParameters) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *EnvironmentUpdateParameters) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetStagingEnvironment

`func (o *EnvironmentUpdateParameters) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *EnvironmentUpdateParameters) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *EnvironmentUpdateParameters) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *EnvironmentUpdateParameters) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetConnectorPort

`func (o *EnvironmentUpdateParameters) GetConnectorPort() int32`

GetConnectorPort returns the ConnectorPort field if non-nil, zero value otherwise.

### GetConnectorPortOk

`func (o *EnvironmentUpdateParameters) GetConnectorPortOk() (*int32, bool)`

GetConnectorPortOk returns a tuple with the ConnectorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPort

`func (o *EnvironmentUpdateParameters) SetConnectorPort(v int32)`

SetConnectorPort sets ConnectorPort field to given value.

### HasConnectorPort

`func (o *EnvironmentUpdateParameters) HasConnectorPort() bool`

HasConnectorPort returns a boolean if a field has been set.

### GetNfsAddresses

`func (o *EnvironmentUpdateParameters) GetNfsAddresses() []string`

GetNfsAddresses returns the NfsAddresses field if non-nil, zero value otherwise.

### GetNfsAddressesOk

`func (o *EnvironmentUpdateParameters) GetNfsAddressesOk() (*[]string, bool)`

GetNfsAddressesOk returns a tuple with the NfsAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsAddresses

`func (o *EnvironmentUpdateParameters) SetNfsAddresses(v []string)`

SetNfsAddresses sets NfsAddresses field to given value.

### HasNfsAddresses

`func (o *EnvironmentUpdateParameters) HasNfsAddresses() bool`

HasNfsAddresses returns a boolean if a field has been set.

### GetSshPort

`func (o *EnvironmentUpdateParameters) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *EnvironmentUpdateParameters) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *EnvironmentUpdateParameters) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *EnvironmentUpdateParameters) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetToolkitPath

`func (o *EnvironmentUpdateParameters) GetToolkitPath() string`

GetToolkitPath returns the ToolkitPath field if non-nil, zero value otherwise.

### GetToolkitPathOk

`func (o *EnvironmentUpdateParameters) GetToolkitPathOk() (*string, bool)`

GetToolkitPathOk returns a tuple with the ToolkitPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolkitPath

`func (o *EnvironmentUpdateParameters) SetToolkitPath(v string)`

SetToolkitPath sets ToolkitPath field to given value.

### HasToolkitPath

`func (o *EnvironmentUpdateParameters) HasToolkitPath() bool`

HasToolkitPath returns a boolean if a field has been set.

### GetJavaHome

`func (o *EnvironmentUpdateParameters) GetJavaHome() string`

GetJavaHome returns the JavaHome field if non-nil, zero value otherwise.

### GetJavaHomeOk

`func (o *EnvironmentUpdateParameters) GetJavaHomeOk() (*string, bool)`

GetJavaHomeOk returns a tuple with the JavaHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaHome

`func (o *EnvironmentUpdateParameters) SetJavaHome(v string)`

SetJavaHome sets JavaHome field to given value.

### HasJavaHome

`func (o *EnvironmentUpdateParameters) HasJavaHome() bool`

HasJavaHome returns a boolean if a field has been set.

### GetDspKeystorePath

`func (o *EnvironmentUpdateParameters) GetDspKeystorePath() string`

GetDspKeystorePath returns the DspKeystorePath field if non-nil, zero value otherwise.

### GetDspKeystorePathOk

`func (o *EnvironmentUpdateParameters) GetDspKeystorePathOk() (*string, bool)`

GetDspKeystorePathOk returns a tuple with the DspKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePath

`func (o *EnvironmentUpdateParameters) SetDspKeystorePath(v string)`

SetDspKeystorePath sets DspKeystorePath field to given value.

### HasDspKeystorePath

`func (o *EnvironmentUpdateParameters) HasDspKeystorePath() bool`

HasDspKeystorePath returns a boolean if a field has been set.

### GetDspKeystorePassword

`func (o *EnvironmentUpdateParameters) GetDspKeystorePassword() string`

GetDspKeystorePassword returns the DspKeystorePassword field if non-nil, zero value otherwise.

### GetDspKeystorePasswordOk

`func (o *EnvironmentUpdateParameters) GetDspKeystorePasswordOk() (*string, bool)`

GetDspKeystorePasswordOk returns a tuple with the DspKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystorePassword

`func (o *EnvironmentUpdateParameters) SetDspKeystorePassword(v string)`

SetDspKeystorePassword sets DspKeystorePassword field to given value.

### HasDspKeystorePassword

`func (o *EnvironmentUpdateParameters) HasDspKeystorePassword() bool`

HasDspKeystorePassword returns a boolean if a field has been set.

### GetDspKeystoreAlias

`func (o *EnvironmentUpdateParameters) GetDspKeystoreAlias() string`

GetDspKeystoreAlias returns the DspKeystoreAlias field if non-nil, zero value otherwise.

### GetDspKeystoreAliasOk

`func (o *EnvironmentUpdateParameters) GetDspKeystoreAliasOk() (*string, bool)`

GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspKeystoreAlias

`func (o *EnvironmentUpdateParameters) SetDspKeystoreAlias(v string)`

SetDspKeystoreAlias sets DspKeystoreAlias field to given value.

### HasDspKeystoreAlias

`func (o *EnvironmentUpdateParameters) HasDspKeystoreAlias() bool`

HasDspKeystoreAlias returns a boolean if a field has been set.

### GetDspTruststorePath

`func (o *EnvironmentUpdateParameters) GetDspTruststorePath() string`

GetDspTruststorePath returns the DspTruststorePath field if non-nil, zero value otherwise.

### GetDspTruststorePathOk

`func (o *EnvironmentUpdateParameters) GetDspTruststorePathOk() (*string, bool)`

GetDspTruststorePathOk returns a tuple with the DspTruststorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePath

`func (o *EnvironmentUpdateParameters) SetDspTruststorePath(v string)`

SetDspTruststorePath sets DspTruststorePath field to given value.

### HasDspTruststorePath

`func (o *EnvironmentUpdateParameters) HasDspTruststorePath() bool`

HasDspTruststorePath returns a boolean if a field has been set.

### GetDspTruststorePassword

`func (o *EnvironmentUpdateParameters) GetDspTruststorePassword() string`

GetDspTruststorePassword returns the DspTruststorePassword field if non-nil, zero value otherwise.

### GetDspTruststorePasswordOk

`func (o *EnvironmentUpdateParameters) GetDspTruststorePasswordOk() (*string, bool)`

GetDspTruststorePasswordOk returns a tuple with the DspTruststorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDspTruststorePassword

`func (o *EnvironmentUpdateParameters) SetDspTruststorePassword(v string)`

SetDspTruststorePassword sets DspTruststorePassword field to given value.

### HasDspTruststorePassword

`func (o *EnvironmentUpdateParameters) HasDspTruststorePassword() bool`

HasDspTruststorePassword returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentUpdateParameters) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentUpdateParameters) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentUpdateParameters) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentUpdateParameters) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


