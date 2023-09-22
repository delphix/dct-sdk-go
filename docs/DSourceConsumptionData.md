# DSourceConsumptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the dSource | [optional] 
**Status** | Pointer to **string** | The status of the dSource | [optional] 
**DatabaseType** | Pointer to **string** | The type of the dSource | [optional] 
**EngineId** | Pointer to **string** | The id of the engine the dSource belongs to. | [optional] 
**EngineName** | Pointer to **string** | The name of the engine the dSource belongs to | [optional] 
**LastConsumptionDate** | Pointer to **time.Time** | The most recent date where consumption data was captured for this dSource. | [optional] 
**IngestedSize** | Pointer to **int64** | The ingested size of the dSource | [optional] 

## Methods

### NewDSourceConsumptionData

`func NewDSourceConsumptionData() *DSourceConsumptionData`

NewDSourceConsumptionData instantiates a new DSourceConsumptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDSourceConsumptionDataWithDefaults

`func NewDSourceConsumptionDataWithDefaults() *DSourceConsumptionData`

NewDSourceConsumptionDataWithDefaults instantiates a new DSourceConsumptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DSourceConsumptionData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DSourceConsumptionData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DSourceConsumptionData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DSourceConsumptionData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DSourceConsumptionData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DSourceConsumptionData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DSourceConsumptionData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DSourceConsumptionData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDatabaseType

`func (o *DSourceConsumptionData) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *DSourceConsumptionData) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *DSourceConsumptionData) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *DSourceConsumptionData) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetEngineId

`func (o *DSourceConsumptionData) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *DSourceConsumptionData) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *DSourceConsumptionData) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *DSourceConsumptionData) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineName

`func (o *DSourceConsumptionData) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *DSourceConsumptionData) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *DSourceConsumptionData) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *DSourceConsumptionData) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetLastConsumptionDate

`func (o *DSourceConsumptionData) GetLastConsumptionDate() time.Time`

GetLastConsumptionDate returns the LastConsumptionDate field if non-nil, zero value otherwise.

### GetLastConsumptionDateOk

`func (o *DSourceConsumptionData) GetLastConsumptionDateOk() (*time.Time, bool)`

GetLastConsumptionDateOk returns a tuple with the LastConsumptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumptionDate

`func (o *DSourceConsumptionData) SetLastConsumptionDate(v time.Time)`

SetLastConsumptionDate sets LastConsumptionDate field to given value.

### HasLastConsumptionDate

`func (o *DSourceConsumptionData) HasLastConsumptionDate() bool`

HasLastConsumptionDate returns a boolean if a field has been set.

### GetIngestedSize

`func (o *DSourceConsumptionData) GetIngestedSize() int64`

GetIngestedSize returns the IngestedSize field if non-nil, zero value otherwise.

### GetIngestedSizeOk

`func (o *DSourceConsumptionData) GetIngestedSizeOk() (*int64, bool)`

GetIngestedSizeOk returns a tuple with the IngestedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestedSize

`func (o *DSourceConsumptionData) SetIngestedSize(v int64)`

SetIngestedSize sets IngestedSize field to given value.

### HasIngestedSize

`func (o *DSourceConsumptionData) HasIngestedSize() bool`

HasIngestedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


