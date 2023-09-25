# HyperscaleDatasetTableOrFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaName** | Pointer to **string** | Name of the table schema. | [optional] 
**TableName** | Pointer to **string** | Name of the table. | [optional] 
**FilterKey** | Pointer to **string** | The unique database column field to filter the source data. | [optional] 
**ColumnArrayRows** | Pointer to **int64** | The number of column array rows to be used by the sqlldr oracle utility which determines the number of rows loaded before the stream buffer is built. | [optional] 
**UnloadSplit** | Pointer to **int64** | The number of unloaded files to be generated from the source database. | [optional] 
**StreamSize** | Pointer to **int64** | Long The stream size to be used by the sqlldr oracle utility which specifies the size (in bytes) of the data stream sent from the client to the server. | [optional] 
**MaskingInventory** | Pointer to [**[]HyperscaleColumnOrField**](HyperscaleColumnOrField.md) | DataSet information for masking inventory. | [optional] 

## Methods

### NewHyperscaleDatasetTableOrFile

`func NewHyperscaleDatasetTableOrFile() *HyperscaleDatasetTableOrFile`

NewHyperscaleDatasetTableOrFile instantiates a new HyperscaleDatasetTableOrFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleDatasetTableOrFileWithDefaults

`func NewHyperscaleDatasetTableOrFileWithDefaults() *HyperscaleDatasetTableOrFile`

NewHyperscaleDatasetTableOrFileWithDefaults instantiates a new HyperscaleDatasetTableOrFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaName

`func (o *HyperscaleDatasetTableOrFile) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *HyperscaleDatasetTableOrFile) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *HyperscaleDatasetTableOrFile) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *HyperscaleDatasetTableOrFile) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### GetTableName

`func (o *HyperscaleDatasetTableOrFile) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *HyperscaleDatasetTableOrFile) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *HyperscaleDatasetTableOrFile) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *HyperscaleDatasetTableOrFile) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetFilterKey

`func (o *HyperscaleDatasetTableOrFile) GetFilterKey() string`

GetFilterKey returns the FilterKey field if non-nil, zero value otherwise.

### GetFilterKeyOk

`func (o *HyperscaleDatasetTableOrFile) GetFilterKeyOk() (*string, bool)`

GetFilterKeyOk returns a tuple with the FilterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterKey

`func (o *HyperscaleDatasetTableOrFile) SetFilterKey(v string)`

SetFilterKey sets FilterKey field to given value.

### HasFilterKey

`func (o *HyperscaleDatasetTableOrFile) HasFilterKey() bool`

HasFilterKey returns a boolean if a field has been set.

### GetColumnArrayRows

`func (o *HyperscaleDatasetTableOrFile) GetColumnArrayRows() int64`

GetColumnArrayRows returns the ColumnArrayRows field if non-nil, zero value otherwise.

### GetColumnArrayRowsOk

`func (o *HyperscaleDatasetTableOrFile) GetColumnArrayRowsOk() (*int64, bool)`

GetColumnArrayRowsOk returns a tuple with the ColumnArrayRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnArrayRows

`func (o *HyperscaleDatasetTableOrFile) SetColumnArrayRows(v int64)`

SetColumnArrayRows sets ColumnArrayRows field to given value.

### HasColumnArrayRows

`func (o *HyperscaleDatasetTableOrFile) HasColumnArrayRows() bool`

HasColumnArrayRows returns a boolean if a field has been set.

### GetUnloadSplit

`func (o *HyperscaleDatasetTableOrFile) GetUnloadSplit() int64`

GetUnloadSplit returns the UnloadSplit field if non-nil, zero value otherwise.

### GetUnloadSplitOk

`func (o *HyperscaleDatasetTableOrFile) GetUnloadSplitOk() (*int64, bool)`

GetUnloadSplitOk returns a tuple with the UnloadSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnloadSplit

`func (o *HyperscaleDatasetTableOrFile) SetUnloadSplit(v int64)`

SetUnloadSplit sets UnloadSplit field to given value.

### HasUnloadSplit

`func (o *HyperscaleDatasetTableOrFile) HasUnloadSplit() bool`

HasUnloadSplit returns a boolean if a field has been set.

### GetStreamSize

`func (o *HyperscaleDatasetTableOrFile) GetStreamSize() int64`

GetStreamSize returns the StreamSize field if non-nil, zero value otherwise.

### GetStreamSizeOk

`func (o *HyperscaleDatasetTableOrFile) GetStreamSizeOk() (*int64, bool)`

GetStreamSizeOk returns a tuple with the StreamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSize

`func (o *HyperscaleDatasetTableOrFile) SetStreamSize(v int64)`

SetStreamSize sets StreamSize field to given value.

### HasStreamSize

`func (o *HyperscaleDatasetTableOrFile) HasStreamSize() bool`

HasStreamSize returns a boolean if a field has been set.

### GetMaskingInventory

`func (o *HyperscaleDatasetTableOrFile) GetMaskingInventory() []HyperscaleColumnOrField`

GetMaskingInventory returns the MaskingInventory field if non-nil, zero value otherwise.

### GetMaskingInventoryOk

`func (o *HyperscaleDatasetTableOrFile) GetMaskingInventoryOk() (*[]HyperscaleColumnOrField, bool)`

GetMaskingInventoryOk returns a tuple with the MaskingInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingInventory

`func (o *HyperscaleDatasetTableOrFile) SetMaskingInventory(v []HyperscaleColumnOrField)`

SetMaskingInventory sets MaskingInventory field to given value.

### HasMaskingInventory

`func (o *HyperscaleDatasetTableOrFile) HasMaskingInventory() bool`

HasMaskingInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


