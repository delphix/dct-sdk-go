# VDBInventoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineName** | Pointer to **string** | The name of the engine the VDB belongs to. | [optional] 
**Name** | Pointer to **string** | The name of the VDB. | [optional] 
**Type** | Pointer to **string** | The database type of the VDB. | [optional] 
**Version** | Pointer to **string** | The database version of the VDB. | [optional] 
**ParentName** | Pointer to **string** | The name of the VDB&#39;s parent dataset. | [optional] 
**ParentId** | Pointer to **string** | A reference to the parent dataset of the VDB. | [optional] 
**CreationDate** | Pointer to **time.Time** | The date the VDB was created. | [optional] 
**LastRefreshedDate** | Pointer to **time.Time** | The date the VDB was last refreshed. | [optional] 
**ParentTimeflowLocation** | Pointer to **string** | The location for the VDB&#39;s parent timeflow. | [optional] 
**ParentTimeflowTimestamp** | Pointer to **time.Time** | The timestamp for the VDB&#39;s parent timeflow. | [optional] 
**ParentTimeflowTimezone** | Pointer to **string** | The timezone for the VDB&#39;s parent timeflow. | [optional] 
**Enabled** | Pointer to **bool** | Whether the VDB is enabled | [optional] 
**Status** | Pointer to **string** | The runtime status of the VDB. &#39;Unknown&#39; if all attempts to connect to the dataset failed. | [optional] 
**StorageSize** | Pointer to **int64** | The actual space used by the VDB, in bytes. | [optional] 

## Methods

### NewVDBInventoryData

`func NewVDBInventoryData() *VDBInventoryData`

NewVDBInventoryData instantiates a new VDBInventoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDBInventoryDataWithDefaults

`func NewVDBInventoryDataWithDefaults() *VDBInventoryData`

NewVDBInventoryDataWithDefaults instantiates a new VDBInventoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineName

`func (o *VDBInventoryData) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *VDBInventoryData) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *VDBInventoryData) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *VDBInventoryData) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetName

`func (o *VDBInventoryData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VDBInventoryData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VDBInventoryData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VDBInventoryData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VDBInventoryData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VDBInventoryData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VDBInventoryData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VDBInventoryData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *VDBInventoryData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VDBInventoryData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VDBInventoryData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VDBInventoryData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetParentName

`func (o *VDBInventoryData) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *VDBInventoryData) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *VDBInventoryData) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *VDBInventoryData) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetParentId

`func (o *VDBInventoryData) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *VDBInventoryData) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *VDBInventoryData) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *VDBInventoryData) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetCreationDate

`func (o *VDBInventoryData) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *VDBInventoryData) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *VDBInventoryData) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *VDBInventoryData) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetLastRefreshedDate

`func (o *VDBInventoryData) GetLastRefreshedDate() time.Time`

GetLastRefreshedDate returns the LastRefreshedDate field if non-nil, zero value otherwise.

### GetLastRefreshedDateOk

`func (o *VDBInventoryData) GetLastRefreshedDateOk() (*time.Time, bool)`

GetLastRefreshedDateOk returns a tuple with the LastRefreshedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshedDate

`func (o *VDBInventoryData) SetLastRefreshedDate(v time.Time)`

SetLastRefreshedDate sets LastRefreshedDate field to given value.

### HasLastRefreshedDate

`func (o *VDBInventoryData) HasLastRefreshedDate() bool`

HasLastRefreshedDate returns a boolean if a field has been set.

### GetParentTimeflowLocation

`func (o *VDBInventoryData) GetParentTimeflowLocation() string`

GetParentTimeflowLocation returns the ParentTimeflowLocation field if non-nil, zero value otherwise.

### GetParentTimeflowLocationOk

`func (o *VDBInventoryData) GetParentTimeflowLocationOk() (*string, bool)`

GetParentTimeflowLocationOk returns a tuple with the ParentTimeflowLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeflowLocation

`func (o *VDBInventoryData) SetParentTimeflowLocation(v string)`

SetParentTimeflowLocation sets ParentTimeflowLocation field to given value.

### HasParentTimeflowLocation

`func (o *VDBInventoryData) HasParentTimeflowLocation() bool`

HasParentTimeflowLocation returns a boolean if a field has been set.

### GetParentTimeflowTimestamp

`func (o *VDBInventoryData) GetParentTimeflowTimestamp() time.Time`

GetParentTimeflowTimestamp returns the ParentTimeflowTimestamp field if non-nil, zero value otherwise.

### GetParentTimeflowTimestampOk

`func (o *VDBInventoryData) GetParentTimeflowTimestampOk() (*time.Time, bool)`

GetParentTimeflowTimestampOk returns a tuple with the ParentTimeflowTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeflowTimestamp

`func (o *VDBInventoryData) SetParentTimeflowTimestamp(v time.Time)`

SetParentTimeflowTimestamp sets ParentTimeflowTimestamp field to given value.

### HasParentTimeflowTimestamp

`func (o *VDBInventoryData) HasParentTimeflowTimestamp() bool`

HasParentTimeflowTimestamp returns a boolean if a field has been set.

### GetParentTimeflowTimezone

`func (o *VDBInventoryData) GetParentTimeflowTimezone() string`

GetParentTimeflowTimezone returns the ParentTimeflowTimezone field if non-nil, zero value otherwise.

### GetParentTimeflowTimezoneOk

`func (o *VDBInventoryData) GetParentTimeflowTimezoneOk() (*string, bool)`

GetParentTimeflowTimezoneOk returns a tuple with the ParentTimeflowTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTimeflowTimezone

`func (o *VDBInventoryData) SetParentTimeflowTimezone(v string)`

SetParentTimeflowTimezone sets ParentTimeflowTimezone field to given value.

### HasParentTimeflowTimezone

`func (o *VDBInventoryData) HasParentTimeflowTimezone() bool`

HasParentTimeflowTimezone returns a boolean if a field has been set.

### GetEnabled

`func (o *VDBInventoryData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VDBInventoryData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VDBInventoryData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VDBInventoryData) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *VDBInventoryData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VDBInventoryData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VDBInventoryData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VDBInventoryData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageSize

`func (o *VDBInventoryData) GetStorageSize() int64`

GetStorageSize returns the StorageSize field if non-nil, zero value otherwise.

### GetStorageSizeOk

`func (o *VDBInventoryData) GetStorageSizeOk() (*int64, bool)`

GetStorageSizeOk returns a tuple with the StorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSize

`func (o *VDBInventoryData) SetStorageSize(v int64)`

SetStorageSize sets StorageSize field to given value.

### HasStorageSize

`func (o *VDBInventoryData) HasStorageSize() bool`

HasStorageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


