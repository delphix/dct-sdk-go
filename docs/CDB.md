# CDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The CDB object entity ID. | [optional] 
**Name** | Pointer to **NullableString** | The name of this CDB. | [optional] 
**DatabaseVersion** | Pointer to **NullableString** | The version of this CDB. | [optional] 
**EnvironmentId** | Pointer to **NullableString** | A reference to the Environment that hosts this CDB. | [optional] 
**Size** | Pointer to **NullableInt64** | The total size of the data files used by this CDB, in bytes. | [optional] 
**JdbcConnectionString** | Pointer to **NullableString** | The JDBC connection URL for this CDB. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this CDB belongs to. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewCDB

`func NewCDB() *CDB`

NewCDB instantiates a new CDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCDBWithDefaults

`func NewCDBWithDefaults() *CDB`

NewCDBWithDefaults instantiates a new CDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CDB) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CDB) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CDB) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CDB) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CDB) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CDB) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CDB) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CDB) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CDB) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CDB) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDatabaseVersion

`func (o *CDB) GetDatabaseVersion() string`

GetDatabaseVersion returns the DatabaseVersion field if non-nil, zero value otherwise.

### GetDatabaseVersionOk

`func (o *CDB) GetDatabaseVersionOk() (*string, bool)`

GetDatabaseVersionOk returns a tuple with the DatabaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseVersion

`func (o *CDB) SetDatabaseVersion(v string)`

SetDatabaseVersion sets DatabaseVersion field to given value.

### HasDatabaseVersion

`func (o *CDB) HasDatabaseVersion() bool`

HasDatabaseVersion returns a boolean if a field has been set.

### SetDatabaseVersionNil

`func (o *CDB) SetDatabaseVersionNil(b bool)`

 SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil

### UnsetDatabaseVersion
`func (o *CDB) UnsetDatabaseVersion()`

UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
### GetEnvironmentId

`func (o *CDB) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CDB) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CDB) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CDB) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### SetEnvironmentIdNil

`func (o *CDB) SetEnvironmentIdNil(b bool)`

 SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil

### UnsetEnvironmentId
`func (o *CDB) UnsetEnvironmentId()`

UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
### GetSize

`func (o *CDB) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CDB) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CDB) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *CDB) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *CDB) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *CDB) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetJdbcConnectionString

`func (o *CDB) GetJdbcConnectionString() string`

GetJdbcConnectionString returns the JdbcConnectionString field if non-nil, zero value otherwise.

### GetJdbcConnectionStringOk

`func (o *CDB) GetJdbcConnectionStringOk() (*string, bool)`

GetJdbcConnectionStringOk returns a tuple with the JdbcConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcConnectionString

`func (o *CDB) SetJdbcConnectionString(v string)`

SetJdbcConnectionString sets JdbcConnectionString field to given value.

### HasJdbcConnectionString

`func (o *CDB) HasJdbcConnectionString() bool`

HasJdbcConnectionString returns a boolean if a field has been set.

### SetJdbcConnectionStringNil

`func (o *CDB) SetJdbcConnectionStringNil(b bool)`

 SetJdbcConnectionStringNil sets the value for JdbcConnectionString to be an explicit nil

### UnsetJdbcConnectionString
`func (o *CDB) UnsetJdbcConnectionString()`

UnsetJdbcConnectionString ensures that no value is present for JdbcConnectionString, not even an explicit nil
### GetEngineId

`func (o *CDB) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *CDB) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *CDB) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *CDB) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetTags

`func (o *CDB) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CDB) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CDB) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CDB) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


