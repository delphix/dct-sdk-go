# VirtualizationStorageSummaryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | **string** | A reference to the engine. | 
**EngineName** | **string** | The engine name. | 
**EngineHostname** | **string** | The engine hostname. | 
**TotalCapacity** | Pointer to **int64** | The total amount of storage allocated for engine objects and system metadata, in bytes. | [optional] 
**FreeStorage** | Pointer to **int64** | The amount of available storage, in bytes. | [optional] 
**UsedStorage** | Pointer to **int64** | The amount of storage used by engine objects and system metadata, in bytes. | [optional] 
**UsedPercentage** | Pointer to **float32** | The percentage of storage used. | [optional] 
**DsourceCount** | Pointer to **int64** | The number of dSources on the engine. | [optional] 
**VdbCount** | Pointer to **int64** | The number of VDBs on the engine. | [optional] 
**TotalObjectCount** | Pointer to **int64** | The total number of dSources and VDBs on the engine. | [optional] 

## Methods

### NewVirtualizationStorageSummaryData

`func NewVirtualizationStorageSummaryData(engineId string, engineName string, engineHostname string, ) *VirtualizationStorageSummaryData`

NewVirtualizationStorageSummaryData instantiates a new VirtualizationStorageSummaryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationStorageSummaryDataWithDefaults

`func NewVirtualizationStorageSummaryDataWithDefaults() *VirtualizationStorageSummaryData`

NewVirtualizationStorageSummaryDataWithDefaults instantiates a new VirtualizationStorageSummaryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *VirtualizationStorageSummaryData) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *VirtualizationStorageSummaryData) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *VirtualizationStorageSummaryData) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetEngineName

`func (o *VirtualizationStorageSummaryData) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *VirtualizationStorageSummaryData) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *VirtualizationStorageSummaryData) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.


### GetEngineHostname

`func (o *VirtualizationStorageSummaryData) GetEngineHostname() string`

GetEngineHostname returns the EngineHostname field if non-nil, zero value otherwise.

### GetEngineHostnameOk

`func (o *VirtualizationStorageSummaryData) GetEngineHostnameOk() (*string, bool)`

GetEngineHostnameOk returns a tuple with the EngineHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHostname

`func (o *VirtualizationStorageSummaryData) SetEngineHostname(v string)`

SetEngineHostname sets EngineHostname field to given value.


### GetTotalCapacity

`func (o *VirtualizationStorageSummaryData) GetTotalCapacity() int64`

GetTotalCapacity returns the TotalCapacity field if non-nil, zero value otherwise.

### GetTotalCapacityOk

`func (o *VirtualizationStorageSummaryData) GetTotalCapacityOk() (*int64, bool)`

GetTotalCapacityOk returns a tuple with the TotalCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCapacity

`func (o *VirtualizationStorageSummaryData) SetTotalCapacity(v int64)`

SetTotalCapacity sets TotalCapacity field to given value.

### HasTotalCapacity

`func (o *VirtualizationStorageSummaryData) HasTotalCapacity() bool`

HasTotalCapacity returns a boolean if a field has been set.

### GetFreeStorage

`func (o *VirtualizationStorageSummaryData) GetFreeStorage() int64`

GetFreeStorage returns the FreeStorage field if non-nil, zero value otherwise.

### GetFreeStorageOk

`func (o *VirtualizationStorageSummaryData) GetFreeStorageOk() (*int64, bool)`

GetFreeStorageOk returns a tuple with the FreeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStorage

`func (o *VirtualizationStorageSummaryData) SetFreeStorage(v int64)`

SetFreeStorage sets FreeStorage field to given value.

### HasFreeStorage

`func (o *VirtualizationStorageSummaryData) HasFreeStorage() bool`

HasFreeStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *VirtualizationStorageSummaryData) GetUsedStorage() int64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *VirtualizationStorageSummaryData) GetUsedStorageOk() (*int64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *VirtualizationStorageSummaryData) SetUsedStorage(v int64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *VirtualizationStorageSummaryData) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.

### GetUsedPercentage

`func (o *VirtualizationStorageSummaryData) GetUsedPercentage() float32`

GetUsedPercentage returns the UsedPercentage field if non-nil, zero value otherwise.

### GetUsedPercentageOk

`func (o *VirtualizationStorageSummaryData) GetUsedPercentageOk() (*float32, bool)`

GetUsedPercentageOk returns a tuple with the UsedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPercentage

`func (o *VirtualizationStorageSummaryData) SetUsedPercentage(v float32)`

SetUsedPercentage sets UsedPercentage field to given value.

### HasUsedPercentage

`func (o *VirtualizationStorageSummaryData) HasUsedPercentage() bool`

HasUsedPercentage returns a boolean if a field has been set.

### GetDsourceCount

`func (o *VirtualizationStorageSummaryData) GetDsourceCount() int64`

GetDsourceCount returns the DsourceCount field if non-nil, zero value otherwise.

### GetDsourceCountOk

`func (o *VirtualizationStorageSummaryData) GetDsourceCountOk() (*int64, bool)`

GetDsourceCountOk returns a tuple with the DsourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsourceCount

`func (o *VirtualizationStorageSummaryData) SetDsourceCount(v int64)`

SetDsourceCount sets DsourceCount field to given value.

### HasDsourceCount

`func (o *VirtualizationStorageSummaryData) HasDsourceCount() bool`

HasDsourceCount returns a boolean if a field has been set.

### GetVdbCount

`func (o *VirtualizationStorageSummaryData) GetVdbCount() int64`

GetVdbCount returns the VdbCount field if non-nil, zero value otherwise.

### GetVdbCountOk

`func (o *VirtualizationStorageSummaryData) GetVdbCountOk() (*int64, bool)`

GetVdbCountOk returns a tuple with the VdbCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbCount

`func (o *VirtualizationStorageSummaryData) SetVdbCount(v int64)`

SetVdbCount sets VdbCount field to given value.

### HasVdbCount

`func (o *VirtualizationStorageSummaryData) HasVdbCount() bool`

HasVdbCount returns a boolean if a field has been set.

### GetTotalObjectCount

`func (o *VirtualizationStorageSummaryData) GetTotalObjectCount() int64`

GetTotalObjectCount returns the TotalObjectCount field if non-nil, zero value otherwise.

### GetTotalObjectCountOk

`func (o *VirtualizationStorageSummaryData) GetTotalObjectCountOk() (*int64, bool)`

GetTotalObjectCountOk returns a tuple with the TotalObjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjectCount

`func (o *VirtualizationStorageSummaryData) SetTotalObjectCount(v int64)`

SetTotalObjectCount sets TotalObjectCount field to given value.

### HasTotalObjectCount

`func (o *VirtualizationStorageSummaryData) HasTotalObjectCount() bool`

HasTotalObjectCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


