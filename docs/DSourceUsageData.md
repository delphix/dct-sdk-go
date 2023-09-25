# DSourceUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineName** | Pointer to **string** | The name of the engine the dSource belongs to. | [optional] 
**Name** | Pointer to **string** | The name of the dSource | [optional] 
**UnvirtualizedSpace** | Pointer to **int64** | The the disk space that would be required if not using Delphix virtualizion, in bytes. | [optional] 
**ActualSpace** | Pointer to **int64** | The actual space used by this dSource, in bytes. This includes the disk space used by the current copy of the data as well as the snapshots and log files retained to enable provisioning from historical versions. | [optional] 
**DependantVdbs** | Pointer to **int32** | The number of VDBs that are dependant on this dSource. This includes all VDB descendants that have this dSource as an ancestor. | [optional] 

## Methods

### NewDSourceUsageData

`func NewDSourceUsageData() *DSourceUsageData`

NewDSourceUsageData instantiates a new DSourceUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSourceUsageDataWithDefaults

`func NewDSourceUsageDataWithDefaults() *DSourceUsageData`

NewDSourceUsageDataWithDefaults instantiates a new DSourceUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineName

`func (o *DSourceUsageData) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *DSourceUsageData) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *DSourceUsageData) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *DSourceUsageData) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetName

`func (o *DSourceUsageData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DSourceUsageData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DSourceUsageData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DSourceUsageData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUnvirtualizedSpace

`func (o *DSourceUsageData) GetUnvirtualizedSpace() int64`

GetUnvirtualizedSpace returns the UnvirtualizedSpace field if non-nil, zero value otherwise.

### GetUnvirtualizedSpaceOk

`func (o *DSourceUsageData) GetUnvirtualizedSpaceOk() (*int64, bool)`

GetUnvirtualizedSpaceOk returns a tuple with the UnvirtualizedSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnvirtualizedSpace

`func (o *DSourceUsageData) SetUnvirtualizedSpace(v int64)`

SetUnvirtualizedSpace sets UnvirtualizedSpace field to given value.

### HasUnvirtualizedSpace

`func (o *DSourceUsageData) HasUnvirtualizedSpace() bool`

HasUnvirtualizedSpace returns a boolean if a field has been set.

### GetActualSpace

`func (o *DSourceUsageData) GetActualSpace() int64`

GetActualSpace returns the ActualSpace field if non-nil, zero value otherwise.

### GetActualSpaceOk

`func (o *DSourceUsageData) GetActualSpaceOk() (*int64, bool)`

GetActualSpaceOk returns a tuple with the ActualSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpace

`func (o *DSourceUsageData) SetActualSpace(v int64)`

SetActualSpace sets ActualSpace field to given value.

### HasActualSpace

`func (o *DSourceUsageData) HasActualSpace() bool`

HasActualSpace returns a boolean if a field has been set.

### GetDependantVdbs

`func (o *DSourceUsageData) GetDependantVdbs() int32`

GetDependantVdbs returns the DependantVdbs field if non-nil, zero value otherwise.

### GetDependantVdbsOk

`func (o *DSourceUsageData) GetDependantVdbsOk() (*int32, bool)`

GetDependantVdbsOk returns a tuple with the DependantVdbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependantVdbs

`func (o *DSourceUsageData) SetDependantVdbs(v int32)`

SetDependantVdbs sets DependantVdbs field to given value.

### HasDependantVdbs

`func (o *DSourceUsageData) HasDependantVdbs() bool`

HasDependantVdbs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


