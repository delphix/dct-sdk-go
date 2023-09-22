# AppDataDSourceLinkSourceParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkType** | Pointer to **string** | The type of link to create. Default is AppDataDirect.  * &#x60;AppDataDirect&#x60; - Represents the AppData specific parameters of a link request for a source directly replicated into the Delphix Engine. * &#x60;AppDataStaged&#x60; - Represents the AppData specific parameters of a link request for a source with a staging source.  | [optional] [default to "AppDataDirect"]
**StagingMountBase** | Pointer to **string** | The base mount point for the NFS mount on the staging environment [AppDataStaged only]. | [optional] 
**StagingEnvironment** | Pointer to **string** | The environment used as an intermediate stage to pull data into Delphix [AppDataStaged only]. | [optional] 
**StagingEnvironmentUser** | Pointer to **string** | The environment user used to access the staging environment [AppDataStaged only]. | [optional] 
**EnvironmentUser** | Pointer to **string** | The OS user to use for linking. | [optional] 
**Excludes** | Pointer to **[]string** | List of subdirectories in the source to exclude when syncing data. These paths are relative to the root of the source directory. [AppDataDirect only] | [optional] 
**FollowSymlinks** | Pointer to **[]string** | List of symlinks in the source to follow when syncing data. These paths are relative to the root of the source directory. All other symlinks are preserved. [AppDataDirect only] | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | The JSON payload conforming to the DraftV4 schema based on the type of application data being manipulated. | [optional] 
**SyncParameters** | Pointer to **map[string]interface{}** | The JSON payload conforming to the snapshot parameters definition in a LUA toolkit or platform plugin. | [optional] 

## Methods

### NewAppDataDSourceLinkSourceParametersAllOf

`func NewAppDataDSourceLinkSourceParametersAllOf() *AppDataDSourceLinkSourceParametersAllOf`

NewAppDataDSourceLinkSourceParametersAllOf instantiates a new AppDataDSourceLinkSourceParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDataDSourceLinkSourceParametersAllOfWithDefaults

`func NewAppDataDSourceLinkSourceParametersAllOfWithDefaults() *AppDataDSourceLinkSourceParametersAllOf`

NewAppDataDSourceLinkSourceParametersAllOfWithDefaults instantiates a new AppDataDSourceLinkSourceParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkType

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetLinkType() string`

GetLinkType returns the LinkType field if non-nil, zero value otherwise.

### GetLinkTypeOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetLinkTypeOk() (*string, bool)`

GetLinkTypeOk returns a tuple with the LinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkType

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetLinkType(v string)`

SetLinkType sets LinkType field to given value.

### HasLinkType

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasLinkType() bool`

HasLinkType returns a boolean if a field has been set.

### GetStagingMountBase

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetStagingMountBase() string`

GetStagingMountBase returns the StagingMountBase field if non-nil, zero value otherwise.

### GetStagingMountBaseOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetStagingMountBaseOk() (*string, bool)`

GetStagingMountBaseOk returns a tuple with the StagingMountBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingMountBase

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetStagingMountBase(v string)`

SetStagingMountBase sets StagingMountBase field to given value.

### HasStagingMountBase

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasStagingMountBase() bool`

HasStagingMountBase returns a boolean if a field has been set.

### GetStagingEnvironment

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetStagingEnvironment() string`

GetStagingEnvironment returns the StagingEnvironment field if non-nil, zero value otherwise.

### GetStagingEnvironmentOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetStagingEnvironmentOk() (*string, bool)`

GetStagingEnvironmentOk returns a tuple with the StagingEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironment

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetStagingEnvironment(v string)`

SetStagingEnvironment sets StagingEnvironment field to given value.

### HasStagingEnvironment

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasStagingEnvironment() bool`

HasStagingEnvironment returns a boolean if a field has been set.

### GetStagingEnvironmentUser

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetStagingEnvironmentUser() string`

GetStagingEnvironmentUser returns the StagingEnvironmentUser field if non-nil, zero value otherwise.

### GetStagingEnvironmentUserOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetStagingEnvironmentUserOk() (*string, bool)`

GetStagingEnvironmentUserOk returns a tuple with the StagingEnvironmentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagingEnvironmentUser

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetStagingEnvironmentUser(v string)`

SetStagingEnvironmentUser sets StagingEnvironmentUser field to given value.

### HasStagingEnvironmentUser

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasStagingEnvironmentUser() bool`

HasStagingEnvironmentUser returns a boolean if a field has been set.

### GetEnvironmentUser

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetEnvironmentUser() string`

GetEnvironmentUser returns the EnvironmentUser field if non-nil, zero value otherwise.

### GetEnvironmentUserOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetEnvironmentUserOk() (*string, bool)`

GetEnvironmentUserOk returns a tuple with the EnvironmentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUser

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetEnvironmentUser(v string)`

SetEnvironmentUser sets EnvironmentUser field to given value.

### HasEnvironmentUser

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasEnvironmentUser() bool`

HasEnvironmentUser returns a boolean if a field has been set.

### GetExcludes

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetExcludes() []string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetExcludesOk() (*[]string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetExcludes(v []string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetFollowSymlinks

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetFollowSymlinks() []string`

GetFollowSymlinks returns the FollowSymlinks field if non-nil, zero value otherwise.

### GetFollowSymlinksOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetFollowSymlinksOk() (*[]string, bool)`

GetFollowSymlinksOk returns a tuple with the FollowSymlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowSymlinks

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetFollowSymlinks(v []string)`

SetFollowSymlinks sets FollowSymlinks field to given value.

### HasFollowSymlinks

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasFollowSymlinks() bool`

HasFollowSymlinks returns a boolean if a field has been set.

### GetParameters

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetSyncParameters

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetSyncParameters() map[string]interface{}`

GetSyncParameters returns the SyncParameters field if non-nil, zero value otherwise.

### GetSyncParametersOk

`func (o *AppDataDSourceLinkSourceParametersAllOf) GetSyncParametersOk() (*map[string]interface{}, bool)`

GetSyncParametersOk returns a tuple with the SyncParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncParameters

`func (o *AppDataDSourceLinkSourceParametersAllOf) SetSyncParameters(v map[string]interface{})`

SetSyncParameters sets SyncParameters field to given value.

### HasSyncParameters

`func (o *AppDataDSourceLinkSourceParametersAllOf) HasSyncParameters() bool`

HasSyncParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


