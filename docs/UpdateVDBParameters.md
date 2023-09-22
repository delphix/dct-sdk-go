# UpdateVDBParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The unique name of the VDB within a group. | [optional] 
**DbUsername** | Pointer to **string** | The username of the database user (Oracle, ASE Only). | [optional] 
**DbPassword** | Pointer to **string** | The password of the database user (Oracle, ASE Only). | [optional] 
**ValidateDbCredentials** | Pointer to **bool** | Whether db_username and db_password must be validated, if present, against the VDB. This must be set to false when credentials validation is not possible, for instance if the VDB is known to be disabled. | [optional] [default to true]
**AutoRestart** | Pointer to **bool** | Whether to enable VDB restart. | [optional] 
**EnvironmentUserId** | Pointer to **string** | The environment user ID to use to connect to the target environment. | [optional] 
**TemplateId** | Pointer to **string** | The ID of the target VDB Template (Oracle Only). | [optional] 
**ListenerIds** | Pointer to **[]string** | The listener IDs for this provision operation (Oracle Only). | [optional] 
**NewDbid** | Pointer to **bool** | Whether to enable new DBID for Oracle | [optional] 
**CdcOnProvision** | Pointer to **bool** | Whether to enable CDC on provision for MSSql | [optional] 
**PreScript** | Pointer to **string** | Pre script for MSSql. | [optional] 
**PostScript** | Pointer to **string** | Post script for MSSql. | [optional] 
**Hooks** | Pointer to [**VirtualDatasetHooks**](VirtualDatasetHooks.md) |  | [optional] 
**CustomEnvVars** | Pointer to **map[string]string** | Environment variable to be set when the engine administers a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. Custom environment variables can only be updated while the VDB is disabled. | [optional] 
**CustomEnvFiles** | Pointer to **[]string** | Environment files to be sourced when the Engine administers a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. Custom environment variables can only be updated while the VDB is disabled. | [optional] 
**OracleRacCustomEnvFiles** | Pointer to [**[]OracleRacCustomEnvFile**](OracleRacCustomEnvFile.md) | Environment files to be sourced when the Engine administers an Oracle RAC VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. Custom environment variables can only be updated while the VDB is disabled. | [optional] 
**OracleRacCustomEnvVars** | Pointer to [**[]OracleRacCustomEnvVar**](OracleRacCustomEnvVar.md) | Environment variable to be set when the engine administers an Oracle RAC VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. Custom environment variables can only be updated while the VDB is disabled. | [optional] 
**ParentTdeKeystorePath** | Pointer to **string** | Path to a copy of the parent&#39;s Oracle transparent data encryption keystore on the target host. Required to provision from snapshots containing encrypted database files. (Oracle Multitenant Only) | [optional] 
**ParentTdeKeystorePassword** | Pointer to **string** | The password of the keystore specified in parentTdeKeystorePath. (Oracle Multitenant Only) | [optional] 
**TdeKeyIdentifier** | Pointer to **string** | ID of the key created by Delphix. (Oracle Multitenant Only) | [optional] 
**TargetVcdbTdeKeystorePath** | Pointer to **string** | Path to the keystore of the target vCDB. (Oracle Multitenant Only) | [optional] 
**CdbTdeKeystorePassword** | Pointer to **string** | The password for the Transparent Data Encryption keystore associated with the CDB. (Oracle Multitenant Only) | [optional] 
**AppdataSourceParams** | Pointer to **map[string]interface{}** | The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. | [optional] 
**AdditionalMountPoints** | Pointer to [**[]AdditionalMountPoint**](AdditionalMountPoint.md) | Specifies additional locations on which to mount a subdirectory of an AppData container. Can only be updated while the VDB is disabled. | [optional] 
**AppdataConfigParams** | Pointer to **map[string]interface{}** | The parameters specified by the source config schema in the toolkit | [optional] 
**ConfigParams** | Pointer to **map[string]interface{}** | Database configuration parameter overrides. | [optional] 
**MountPoint** | Pointer to **string** | Mount point for the VDB (AppData only), can only be updated while the VDB is disabled. | [optional] 

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

### GetDbUsername

`func (o *UpdateVDBParameters) GetDbUsername() string`

GetDbUsername returns the DbUsername field if non-nil, zero value otherwise.

### GetDbUsernameOk

`func (o *UpdateVDBParameters) GetDbUsernameOk() (*string, bool)`

GetDbUsernameOk returns a tuple with the DbUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbUsername

`func (o *UpdateVDBParameters) SetDbUsername(v string)`

SetDbUsername sets DbUsername field to given value.

### HasDbUsername

`func (o *UpdateVDBParameters) HasDbUsername() bool`

HasDbUsername returns a boolean if a field has been set.

### GetDbPassword

`func (o *UpdateVDBParameters) GetDbPassword() string`

GetDbPassword returns the DbPassword field if non-nil, zero value otherwise.

### GetDbPasswordOk

`func (o *UpdateVDBParameters) GetDbPasswordOk() (*string, bool)`

GetDbPasswordOk returns a tuple with the DbPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbPassword

`func (o *UpdateVDBParameters) SetDbPassword(v string)`

SetDbPassword sets DbPassword field to given value.

### HasDbPassword

`func (o *UpdateVDBParameters) HasDbPassword() bool`

HasDbPassword returns a boolean if a field has been set.

### GetValidateDbCredentials

`func (o *UpdateVDBParameters) GetValidateDbCredentials() bool`

GetValidateDbCredentials returns the ValidateDbCredentials field if non-nil, zero value otherwise.

### GetValidateDbCredentialsOk

`func (o *UpdateVDBParameters) GetValidateDbCredentialsOk() (*bool, bool)`

GetValidateDbCredentialsOk returns a tuple with the ValidateDbCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateDbCredentials

`func (o *UpdateVDBParameters) SetValidateDbCredentials(v bool)`

SetValidateDbCredentials sets ValidateDbCredentials field to given value.

### HasValidateDbCredentials

`func (o *UpdateVDBParameters) HasValidateDbCredentials() bool`

HasValidateDbCredentials returns a boolean if a field has been set.

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

### GetEnvironmentUserId

`func (o *UpdateVDBParameters) GetEnvironmentUserId() string`

GetEnvironmentUserId returns the EnvironmentUserId field if non-nil, zero value otherwise.

### GetEnvironmentUserIdOk

`func (o *UpdateVDBParameters) GetEnvironmentUserIdOk() (*string, bool)`

GetEnvironmentUserIdOk returns a tuple with the EnvironmentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUserId

`func (o *UpdateVDBParameters) SetEnvironmentUserId(v string)`

SetEnvironmentUserId sets EnvironmentUserId field to given value.

### HasEnvironmentUserId

`func (o *UpdateVDBParameters) HasEnvironmentUserId() bool`

HasEnvironmentUserId returns a boolean if a field has been set.

### GetTemplateId

`func (o *UpdateVDBParameters) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *UpdateVDBParameters) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *UpdateVDBParameters) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *UpdateVDBParameters) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetListenerIds

`func (o *UpdateVDBParameters) GetListenerIds() []string`

GetListenerIds returns the ListenerIds field if non-nil, zero value otherwise.

### GetListenerIdsOk

`func (o *UpdateVDBParameters) GetListenerIdsOk() (*[]string, bool)`

GetListenerIdsOk returns a tuple with the ListenerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerIds

`func (o *UpdateVDBParameters) SetListenerIds(v []string)`

SetListenerIds sets ListenerIds field to given value.

### HasListenerIds

`func (o *UpdateVDBParameters) HasListenerIds() bool`

HasListenerIds returns a boolean if a field has been set.

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

### GetHooks

`func (o *UpdateVDBParameters) GetHooks() VirtualDatasetHooks`

GetHooks returns the Hooks field if non-nil, zero value otherwise.

### GetHooksOk

`func (o *UpdateVDBParameters) GetHooksOk() (*VirtualDatasetHooks, bool)`

GetHooksOk returns a tuple with the Hooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHooks

`func (o *UpdateVDBParameters) SetHooks(v VirtualDatasetHooks)`

SetHooks sets Hooks field to given value.

### HasHooks

`func (o *UpdateVDBParameters) HasHooks() bool`

HasHooks returns a boolean if a field has been set.

### GetCustomEnvVars

`func (o *UpdateVDBParameters) GetCustomEnvVars() map[string]string`

GetCustomEnvVars returns the CustomEnvVars field if non-nil, zero value otherwise.

### GetCustomEnvVarsOk

`func (o *UpdateVDBParameters) GetCustomEnvVarsOk() (*map[string]string, bool)`

GetCustomEnvVarsOk returns a tuple with the CustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvVars

`func (o *UpdateVDBParameters) SetCustomEnvVars(v map[string]string)`

SetCustomEnvVars sets CustomEnvVars field to given value.

### HasCustomEnvVars

`func (o *UpdateVDBParameters) HasCustomEnvVars() bool`

HasCustomEnvVars returns a boolean if a field has been set.

### GetCustomEnvFiles

`func (o *UpdateVDBParameters) GetCustomEnvFiles() []string`

GetCustomEnvFiles returns the CustomEnvFiles field if non-nil, zero value otherwise.

### GetCustomEnvFilesOk

`func (o *UpdateVDBParameters) GetCustomEnvFilesOk() (*[]string, bool)`

GetCustomEnvFilesOk returns a tuple with the CustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEnvFiles

`func (o *UpdateVDBParameters) SetCustomEnvFiles(v []string)`

SetCustomEnvFiles sets CustomEnvFiles field to given value.

### HasCustomEnvFiles

`func (o *UpdateVDBParameters) HasCustomEnvFiles() bool`

HasCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvFiles

`func (o *UpdateVDBParameters) GetOracleRacCustomEnvFiles() []OracleRacCustomEnvFile`

GetOracleRacCustomEnvFiles returns the OracleRacCustomEnvFiles field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvFilesOk

`func (o *UpdateVDBParameters) GetOracleRacCustomEnvFilesOk() (*[]OracleRacCustomEnvFile, bool)`

GetOracleRacCustomEnvFilesOk returns a tuple with the OracleRacCustomEnvFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvFiles

`func (o *UpdateVDBParameters) SetOracleRacCustomEnvFiles(v []OracleRacCustomEnvFile)`

SetOracleRacCustomEnvFiles sets OracleRacCustomEnvFiles field to given value.

### HasOracleRacCustomEnvFiles

`func (o *UpdateVDBParameters) HasOracleRacCustomEnvFiles() bool`

HasOracleRacCustomEnvFiles returns a boolean if a field has been set.

### GetOracleRacCustomEnvVars

`func (o *UpdateVDBParameters) GetOracleRacCustomEnvVars() []OracleRacCustomEnvVar`

GetOracleRacCustomEnvVars returns the OracleRacCustomEnvVars field if non-nil, zero value otherwise.

### GetOracleRacCustomEnvVarsOk

`func (o *UpdateVDBParameters) GetOracleRacCustomEnvVarsOk() (*[]OracleRacCustomEnvVar, bool)`

GetOracleRacCustomEnvVarsOk returns a tuple with the OracleRacCustomEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleRacCustomEnvVars

`func (o *UpdateVDBParameters) SetOracleRacCustomEnvVars(v []OracleRacCustomEnvVar)`

SetOracleRacCustomEnvVars sets OracleRacCustomEnvVars field to given value.

### HasOracleRacCustomEnvVars

`func (o *UpdateVDBParameters) HasOracleRacCustomEnvVars() bool`

HasOracleRacCustomEnvVars returns a boolean if a field has been set.

### GetParentTdeKeystorePath

`func (o *UpdateVDBParameters) GetParentTdeKeystorePath() string`

GetParentTdeKeystorePath returns the ParentTdeKeystorePath field if non-nil, zero value otherwise.

### GetParentTdeKeystorePathOk

`func (o *UpdateVDBParameters) GetParentTdeKeystorePathOk() (*string, bool)`

GetParentTdeKeystorePathOk returns a tuple with the ParentTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePath

`func (o *UpdateVDBParameters) SetParentTdeKeystorePath(v string)`

SetParentTdeKeystorePath sets ParentTdeKeystorePath field to given value.

### HasParentTdeKeystorePath

`func (o *UpdateVDBParameters) HasParentTdeKeystorePath() bool`

HasParentTdeKeystorePath returns a boolean if a field has been set.

### GetParentTdeKeystorePassword

`func (o *UpdateVDBParameters) GetParentTdeKeystorePassword() string`

GetParentTdeKeystorePassword returns the ParentTdeKeystorePassword field if non-nil, zero value otherwise.

### GetParentTdeKeystorePasswordOk

`func (o *UpdateVDBParameters) GetParentTdeKeystorePasswordOk() (*string, bool)`

GetParentTdeKeystorePasswordOk returns a tuple with the ParentTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTdeKeystorePassword

`func (o *UpdateVDBParameters) SetParentTdeKeystorePassword(v string)`

SetParentTdeKeystorePassword sets ParentTdeKeystorePassword field to given value.

### HasParentTdeKeystorePassword

`func (o *UpdateVDBParameters) HasParentTdeKeystorePassword() bool`

HasParentTdeKeystorePassword returns a boolean if a field has been set.

### GetTdeKeyIdentifier

`func (o *UpdateVDBParameters) GetTdeKeyIdentifier() string`

GetTdeKeyIdentifier returns the TdeKeyIdentifier field if non-nil, zero value otherwise.

### GetTdeKeyIdentifierOk

`func (o *UpdateVDBParameters) GetTdeKeyIdentifierOk() (*string, bool)`

GetTdeKeyIdentifierOk returns a tuple with the TdeKeyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdeKeyIdentifier

`func (o *UpdateVDBParameters) SetTdeKeyIdentifier(v string)`

SetTdeKeyIdentifier sets TdeKeyIdentifier field to given value.

### HasTdeKeyIdentifier

`func (o *UpdateVDBParameters) HasTdeKeyIdentifier() bool`

HasTdeKeyIdentifier returns a boolean if a field has been set.

### GetTargetVcdbTdeKeystorePath

`func (o *UpdateVDBParameters) GetTargetVcdbTdeKeystorePath() string`

GetTargetVcdbTdeKeystorePath returns the TargetVcdbTdeKeystorePath field if non-nil, zero value otherwise.

### GetTargetVcdbTdeKeystorePathOk

`func (o *UpdateVDBParameters) GetTargetVcdbTdeKeystorePathOk() (*string, bool)`

GetTargetVcdbTdeKeystorePathOk returns a tuple with the TargetVcdbTdeKeystorePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVcdbTdeKeystorePath

`func (o *UpdateVDBParameters) SetTargetVcdbTdeKeystorePath(v string)`

SetTargetVcdbTdeKeystorePath sets TargetVcdbTdeKeystorePath field to given value.

### HasTargetVcdbTdeKeystorePath

`func (o *UpdateVDBParameters) HasTargetVcdbTdeKeystorePath() bool`

HasTargetVcdbTdeKeystorePath returns a boolean if a field has been set.

### GetCdbTdeKeystorePassword

`func (o *UpdateVDBParameters) GetCdbTdeKeystorePassword() string`

GetCdbTdeKeystorePassword returns the CdbTdeKeystorePassword field if non-nil, zero value otherwise.

### GetCdbTdeKeystorePasswordOk

`func (o *UpdateVDBParameters) GetCdbTdeKeystorePasswordOk() (*string, bool)`

GetCdbTdeKeystorePasswordOk returns a tuple with the CdbTdeKeystorePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdbTdeKeystorePassword

`func (o *UpdateVDBParameters) SetCdbTdeKeystorePassword(v string)`

SetCdbTdeKeystorePassword sets CdbTdeKeystorePassword field to given value.

### HasCdbTdeKeystorePassword

`func (o *UpdateVDBParameters) HasCdbTdeKeystorePassword() bool`

HasCdbTdeKeystorePassword returns a boolean if a field has been set.

### GetAppdataSourceParams

`func (o *UpdateVDBParameters) GetAppdataSourceParams() map[string]interface{}`

GetAppdataSourceParams returns the AppdataSourceParams field if non-nil, zero value otherwise.

### GetAppdataSourceParamsOk

`func (o *UpdateVDBParameters) GetAppdataSourceParamsOk() (*map[string]interface{}, bool)`

GetAppdataSourceParamsOk returns a tuple with the AppdataSourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataSourceParams

`func (o *UpdateVDBParameters) SetAppdataSourceParams(v map[string]interface{})`

SetAppdataSourceParams sets AppdataSourceParams field to given value.

### HasAppdataSourceParams

`func (o *UpdateVDBParameters) HasAppdataSourceParams() bool`

HasAppdataSourceParams returns a boolean if a field has been set.

### GetAdditionalMountPoints

`func (o *UpdateVDBParameters) GetAdditionalMountPoints() []AdditionalMountPoint`

GetAdditionalMountPoints returns the AdditionalMountPoints field if non-nil, zero value otherwise.

### GetAdditionalMountPointsOk

`func (o *UpdateVDBParameters) GetAdditionalMountPointsOk() (*[]AdditionalMountPoint, bool)`

GetAdditionalMountPointsOk returns a tuple with the AdditionalMountPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMountPoints

`func (o *UpdateVDBParameters) SetAdditionalMountPoints(v []AdditionalMountPoint)`

SetAdditionalMountPoints sets AdditionalMountPoints field to given value.

### HasAdditionalMountPoints

`func (o *UpdateVDBParameters) HasAdditionalMountPoints() bool`

HasAdditionalMountPoints returns a boolean if a field has been set.

### SetAdditionalMountPointsNil

`func (o *UpdateVDBParameters) SetAdditionalMountPointsNil(b bool)`

 SetAdditionalMountPointsNil sets the value for AdditionalMountPoints to be an explicit nil

### UnsetAdditionalMountPoints
`func (o *UpdateVDBParameters) UnsetAdditionalMountPoints()`

UnsetAdditionalMountPoints ensures that no value is present for AdditionalMountPoints, not even an explicit nil
### GetAppdataConfigParams

`func (o *UpdateVDBParameters) GetAppdataConfigParams() map[string]interface{}`

GetAppdataConfigParams returns the AppdataConfigParams field if non-nil, zero value otherwise.

### GetAppdataConfigParamsOk

`func (o *UpdateVDBParameters) GetAppdataConfigParamsOk() (*map[string]interface{}, bool)`

GetAppdataConfigParamsOk returns a tuple with the AppdataConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppdataConfigParams

`func (o *UpdateVDBParameters) SetAppdataConfigParams(v map[string]interface{})`

SetAppdataConfigParams sets AppdataConfigParams field to given value.

### HasAppdataConfigParams

`func (o *UpdateVDBParameters) HasAppdataConfigParams() bool`

HasAppdataConfigParams returns a boolean if a field has been set.

### SetAppdataConfigParamsNil

`func (o *UpdateVDBParameters) SetAppdataConfigParamsNil(b bool)`

 SetAppdataConfigParamsNil sets the value for AppdataConfigParams to be an explicit nil

### UnsetAppdataConfigParams
`func (o *UpdateVDBParameters) UnsetAppdataConfigParams()`

UnsetAppdataConfigParams ensures that no value is present for AppdataConfigParams, not even an explicit nil
### GetConfigParams

`func (o *UpdateVDBParameters) GetConfigParams() map[string]interface{}`

GetConfigParams returns the ConfigParams field if non-nil, zero value otherwise.

### GetConfigParamsOk

`func (o *UpdateVDBParameters) GetConfigParamsOk() (*map[string]interface{}, bool)`

GetConfigParamsOk returns a tuple with the ConfigParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigParams

`func (o *UpdateVDBParameters) SetConfigParams(v map[string]interface{})`

SetConfigParams sets ConfigParams field to given value.

### HasConfigParams

`func (o *UpdateVDBParameters) HasConfigParams() bool`

HasConfigParams returns a boolean if a field has been set.

### SetConfigParamsNil

`func (o *UpdateVDBParameters) SetConfigParamsNil(b bool)`

 SetConfigParamsNil sets the value for ConfigParams to be an explicit nil

### UnsetConfigParams
`func (o *UpdateVDBParameters) UnsetConfigParams()`

UnsetConfigParams ensures that no value is present for ConfigParams, not even an explicit nil
### GetMountPoint

`func (o *UpdateVDBParameters) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *UpdateVDBParameters) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *UpdateVDBParameters) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.

### HasMountPoint

`func (o *UpdateVDBParameters) HasMountPoint() bool`

HasMountPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


