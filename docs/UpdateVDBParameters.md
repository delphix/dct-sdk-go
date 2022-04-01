# UpdateVDBParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name to set the VDB to. | [optional] 
**User** | Pointer to **string** | Source config username. | [optional] 
**Password** | Pointer to **string** | Source config password | [optional] 
**AutoRestart** | Pointer to **bool** | Whether to enable VDB restart. | [optional] 
**EnvironmentUser** | Pointer to **string** | Environment to set the VDB to. | [optional] 
**ConfigTemplate** | Pointer to **string** | VDB template reference. | [optional] 
**Listeners** | Pointer to **[]string** | List of Oracle listener references. | [optional] 
**NewDbid** | Pointer to **bool** | Whether to enable new DBID for Oracle | [optional] 
**CdcOnProvision** | Pointer to **bool** | Whether to enable CDC on provision for MSSQL | [optional] 
**PreScript** | Pointer to **string** | Pre script for MSSQL. | [optional] 
**PostScript** | Pointer to **string** | Post script for MSSQL. | [optional] 

## Methods

### NewUpdateVDBParameters

`func NewUpdateVDBParameters() *UpdateVDBParameters`

NewUpdateVDBParameters instantiates a new UpdateVDBParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVDBParametersWithDefaults

`func NewUpdateVDBParametersWithDefaults() *UpdateVDBParameters`

NewUpdateVDBParametersWithDefaults instantiates a new UpdateVDBParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVDBParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDBParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDBParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVDBParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUser

`func (o *UpdateVDBParameters) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateVDBParameters) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateVDBParameters) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateVDBParameters) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateVDBParameters) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateVDBParameters) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateVDBParameters) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateVDBParameters) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAutoRestart

`func (o *UpdateVDBParameters) GetAutoRestart() bool`

GetAutoRestart returns the AutoRestart field if non-nil, zero value otherwise.

### GetAutoRestartOk

`func (o *UpdateVDBParameters) GetAutoRestartOk() (*bool, bool)`

GetAutoRestartOk returns a tuple with the AutoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRestart

`func (o *UpdateVDBParameters) SetAutoRestart(v bool)`

SetAutoRestart sets AutoRestart field to given value.

### HasAutoRestart

`func (o *UpdateVDBParameters) HasAutoRestart() bool`

HasAutoRestart returns a boolean if a field has been set.

### GetEnvironmentUser

`func (o *UpdateVDBParameters) GetEnvironmentUser() string`

GetEnvironmentUser returns the EnvironmentUser field if non-nil, zero value otherwise.

### GetEnvironmentUserOk

`func (o *UpdateVDBParameters) GetEnvironmentUserOk() (*string, bool)`

GetEnvironmentUserOk returns a tuple with the EnvironmentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUser

`func (o *UpdateVDBParameters) SetEnvironmentUser(v string)`

SetEnvironmentUser sets EnvironmentUser field to given value.

### HasEnvironmentUser

`func (o *UpdateVDBParameters) HasEnvironmentUser() bool`

HasEnvironmentUser returns a boolean if a field has been set.

### GetConfigTemplate

`func (o *UpdateVDBParameters) GetConfigTemplate() string`

GetConfigTemplate returns the ConfigTemplate field if non-nil, zero value otherwise.

### GetConfigTemplateOk

`func (o *UpdateVDBParameters) GetConfigTemplateOk() (*string, bool)`

GetConfigTemplateOk returns a tuple with the ConfigTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplate

`func (o *UpdateVDBParameters) SetConfigTemplate(v string)`

SetConfigTemplate sets ConfigTemplate field to given value.

### HasConfigTemplate

`func (o *UpdateVDBParameters) HasConfigTemplate() bool`

HasConfigTemplate returns a boolean if a field has been set.

### GetListeners

`func (o *UpdateVDBParameters) GetListeners() []string`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *UpdateVDBParameters) GetListenersOk() (*[]string, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *UpdateVDBParameters) SetListeners(v []string)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *UpdateVDBParameters) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### GetNewDbid

`func (o *UpdateVDBParameters) GetNewDbid() bool`

GetNewDbid returns the NewDbid field if non-nil, zero value otherwise.

### GetNewDbidOk

`func (o *UpdateVDBParameters) GetNewDbidOk() (*bool, bool)`

GetNewDbidOk returns a tuple with the NewDbid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDbid

`func (o *UpdateVDBParameters) SetNewDbid(v bool)`

SetNewDbid sets NewDbid field to given value.

### HasNewDbid

`func (o *UpdateVDBParameters) HasNewDbid() bool`

HasNewDbid returns a boolean if a field has been set.

### GetCdcOnProvision

`func (o *UpdateVDBParameters) GetCdcOnProvision() bool`

GetCdcOnProvision returns the CdcOnProvision field if non-nil, zero value otherwise.

### GetCdcOnProvisionOk

`func (o *UpdateVDBParameters) GetCdcOnProvisionOk() (*bool, bool)`

GetCdcOnProvisionOk returns a tuple with the CdcOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcOnProvision

`func (o *UpdateVDBParameters) SetCdcOnProvision(v bool)`

SetCdcOnProvision sets CdcOnProvision field to given value.

### HasCdcOnProvision

`func (o *UpdateVDBParameters) HasCdcOnProvision() bool`

HasCdcOnProvision returns a boolean if a field has been set.

### GetPreScript

`func (o *UpdateVDBParameters) GetPreScript() string`

GetPreScript returns the PreScript field if non-nil, zero value otherwise.

### GetPreScriptOk

`func (o *UpdateVDBParameters) GetPreScriptOk() (*string, bool)`

GetPreScriptOk returns a tuple with the PreScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreScript

`func (o *UpdateVDBParameters) SetPreScript(v string)`

SetPreScript sets PreScript field to given value.

### HasPreScript

`func (o *UpdateVDBParameters) HasPreScript() bool`

HasPreScript returns a boolean if a field has been set.

### GetPostScript

`func (o *UpdateVDBParameters) GetPostScript() string`

GetPostScript returns the PostScript field if non-nil, zero value otherwise.

### GetPostScriptOk

`func (o *UpdateVDBParameters) GetPostScriptOk() (*string, bool)`

GetPostScriptOk returns a tuple with the PostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostScript

`func (o *UpdateVDBParameters) SetPostScript(v string)`

SetPostScript sets PostScript field to given value.

### HasPostScript

`func (o *UpdateVDBParameters) HasPostScript() bool`

HasPostScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


