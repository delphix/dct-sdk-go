# VCDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The vCDB object entity ID. | [optional] 
**Name** | Pointer to **NullableString** | The name of this vCDB. | [optional] 
**DatabaseVersion** | Pointer to **NullableString** | The version of this vCDB. | [optional] 
**EnvironmentId** | Pointer to **NullableString** | A reference to the Environment that hosts this vCDB. | [optional] 
**Size** | Pointer to **NullableInt64** | The total size of the data files used by this vCDB, in bytes. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this vCDB belongs to. | [optional] 
**Status** | Pointer to **NullableString** | The runtime status of the vCDB. | [optional] 
**ParentId** | Pointer to **string** | A reference to the parent CDB of this vCDB. | [optional] 
**CreationDate** | Pointer to **NullableTime** | The date this vCDB was created. | [optional] 
**GroupName** | Pointer to **NullableString** | The name of the group containing this vCDB. | [optional] 
**Enabled** | Pointer to **bool** | Whether the vCDB is enabled or not. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewVCDB

`func NewVCDB() *VCDB`

NewVCDB instantiates a new VCDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCDBWithDefaults

`func NewVCDBWithDefaults() *VCDB`

NewVCDBWithDefaults instantiates a new VCDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VCDB) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VCDB) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VCDB) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VCDB) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VCDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VCDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VCDB) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VCDB) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VCDB) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VCDB) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDatabaseVersion

`func (o *VCDB) GetDatabaseVersion() string`

GetDatabaseVersion returns the DatabaseVersion field if non-nil, zero value otherwise.

### GetDatabaseVersionOk

`func (o *VCDB) GetDatabaseVersionOk() (*string, bool)`

GetDatabaseVersionOk returns a tuple with the DatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVersion

`func (o *VCDB) SetDatabaseVersion(v string)`

SetDatabaseVersion sets DatabaseVersion field to given value.

### HasDatabaseVersion

`func (o *VCDB) HasDatabaseVersion() bool`

HasDatabaseVersion returns a boolean if a field has been set.

### SetDatabaseVersionNil

`func (o *VCDB) SetDatabaseVersionNil(b bool)`

 SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil

### UnsetDatabaseVersion
`func (o *VCDB) UnsetDatabaseVersion()`

UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
### GetEnvironmentId

`func (o *VCDB) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *VCDB) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *VCDB) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *VCDB) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *VCDB) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *VCDB) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetSize

`func (o *VCDB) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VCDB) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VCDB) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VCDB) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VCDB) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VCDB) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetEngineId

`func (o *VCDB) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *VCDB) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *VCDB) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *VCDB) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetStatus

`func (o *VCDB) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VCDB) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VCDB) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VCDB) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VCDB) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VCDB) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetParentId

`func (o *VCDB) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *VCDB) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *VCDB) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *VCDB) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetCreationDate

`func (o *VCDB) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VCDB) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VCDB) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VCDB) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *VCDB) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *VCDB) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetGroupName

`func (o *VCDB) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *VCDB) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *VCDB) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *VCDB) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *VCDB) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *VCDB) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetEnabled

`func (o *VCDB) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VCDB) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VCDB) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VCDB) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTags

`func (o *VCDB) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VCDB) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VCDB) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VCDB) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


