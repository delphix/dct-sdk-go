# MaskingExecutionMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The MaskingJob entity ID. | [optional] 
**MaskingJobName** | Pointer to **string** | The name of the MaskingJob. | [optional] 
**MaskingJobType** | Pointer to **string** | The type of the Masking job. | [optional] 
**ConnectorType** | Pointer to **string** | The type of data being masked by this Job. If the Masking Job is masking a database this is the type of the database, otherwise \&quot;FILE\&quot; or \&quot;MAINFRAME_DATASET\&quot;. | [optional] 
**RulesetName** | Pointer to **string** | Name of the ruleset for the Masking job. | [optional] 
**RowsMasked** | Pointer to **int64** | The number of rows masked. This is not applicable for JSON file type. | [optional] 
**RowsTotal** | Pointer to **int64** | The total number of rows. This is not applicable for JSON file type. | [optional] 
**BytesMasked** | Pointer to **int64** | The number of bytes masked. This is only applicable for JSON file type. | [optional] 
**BytesTotal** | Pointer to **int64** | The total number of bytes. This is only applicable for JSON file type. | [optional] 
**Duration** | Pointer to **int64** | The time taken by the execution in ms. Only available for successful executions. | [optional] 
**TablesFilesCount** | Pointer to **int64** | The number of tables or files in the ruleset associated to the Masking job. | [optional] 
**MaskedTablesFilesCount** | Pointer to **int64** | The number of tables or files in the ruleset associated to the Masking job for which at least one column or field is masked. | [optional] 
**ColumnsFieldsCount** | Pointer to **int64** | The number of columns or fields across all tables or files in the ruleset associated to the Masking job. | [optional] 
**MaskedColumnsFieldsCount** | Pointer to **int64** | The number of columns or fields across all tables or files in the ruleset associated to the Masking job which are masked. | [optional] 

## Methods

### NewMaskingExecutionMetrics

`func NewMaskingExecutionMetrics() *MaskingExecutionMetrics`

NewMaskingExecutionMetrics instantiates a new MaskingExecutionMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingExecutionMetricsWithDefaults

`func NewMaskingExecutionMetricsWithDefaults() *MaskingExecutionMetrics`

NewMaskingExecutionMetricsWithDefaults instantiates a new MaskingExecutionMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MaskingExecutionMetrics) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MaskingExecutionMetrics) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MaskingExecutionMetrics) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MaskingExecutionMetrics) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaskingJobName

`func (o *MaskingExecutionMetrics) GetMaskingJobName() string`

GetMaskingJobName returns the MaskingJobName field if non-nil, zero value otherwise.

### GetMaskingJobNameOk

`func (o *MaskingExecutionMetrics) GetMaskingJobNameOk() (*string, bool)`

GetMaskingJobNameOk returns a tuple with the MaskingJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingJobName

`func (o *MaskingExecutionMetrics) SetMaskingJobName(v string)`

SetMaskingJobName sets MaskingJobName field to given value.

### HasMaskingJobName

`func (o *MaskingExecutionMetrics) HasMaskingJobName() bool`

HasMaskingJobName returns a boolean if a field has been set.

### GetMaskingJobType

`func (o *MaskingExecutionMetrics) GetMaskingJobType() string`

GetMaskingJobType returns the MaskingJobType field if non-nil, zero value otherwise.

### GetMaskingJobTypeOk

`func (o *MaskingExecutionMetrics) GetMaskingJobTypeOk() (*string, bool)`

GetMaskingJobTypeOk returns a tuple with the MaskingJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingJobType

`func (o *MaskingExecutionMetrics) SetMaskingJobType(v string)`

SetMaskingJobType sets MaskingJobType field to given value.

### HasMaskingJobType

`func (o *MaskingExecutionMetrics) HasMaskingJobType() bool`

HasMaskingJobType returns a boolean if a field has been set.

### GetConnectorType

`func (o *MaskingExecutionMetrics) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *MaskingExecutionMetrics) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *MaskingExecutionMetrics) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *MaskingExecutionMetrics) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetRulesetName

`func (o *MaskingExecutionMetrics) GetRulesetName() string`

GetRulesetName returns the RulesetName field if non-nil, zero value otherwise.

### GetRulesetNameOk

`func (o *MaskingExecutionMetrics) GetRulesetNameOk() (*string, bool)`

GetRulesetNameOk returns a tuple with the RulesetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesetName

`func (o *MaskingExecutionMetrics) SetRulesetName(v string)`

SetRulesetName sets RulesetName field to given value.

### HasRulesetName

`func (o *MaskingExecutionMetrics) HasRulesetName() bool`

HasRulesetName returns a boolean if a field has been set.

### GetRowsMasked

`func (o *MaskingExecutionMetrics) GetRowsMasked() int64`

GetRowsMasked returns the RowsMasked field if non-nil, zero value otherwise.

### GetRowsMaskedOk

`func (o *MaskingExecutionMetrics) GetRowsMaskedOk() (*int64, bool)`

GetRowsMaskedOk returns a tuple with the RowsMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsMasked

`func (o *MaskingExecutionMetrics) SetRowsMasked(v int64)`

SetRowsMasked sets RowsMasked field to given value.

### HasRowsMasked

`func (o *MaskingExecutionMetrics) HasRowsMasked() bool`

HasRowsMasked returns a boolean if a field has been set.

### GetRowsTotal

`func (o *MaskingExecutionMetrics) GetRowsTotal() int64`

GetRowsTotal returns the RowsTotal field if non-nil, zero value otherwise.

### GetRowsTotalOk

`func (o *MaskingExecutionMetrics) GetRowsTotalOk() (*int64, bool)`

GetRowsTotalOk returns a tuple with the RowsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowsTotal

`func (o *MaskingExecutionMetrics) SetRowsTotal(v int64)`

SetRowsTotal sets RowsTotal field to given value.

### HasRowsTotal

`func (o *MaskingExecutionMetrics) HasRowsTotal() bool`

HasRowsTotal returns a boolean if a field has been set.

### GetBytesMasked

`func (o *MaskingExecutionMetrics) GetBytesMasked() int64`

GetBytesMasked returns the BytesMasked field if non-nil, zero value otherwise.

### GetBytesMaskedOk

`func (o *MaskingExecutionMetrics) GetBytesMaskedOk() (*int64, bool)`

GetBytesMaskedOk returns a tuple with the BytesMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesMasked

`func (o *MaskingExecutionMetrics) SetBytesMasked(v int64)`

SetBytesMasked sets BytesMasked field to given value.

### HasBytesMasked

`func (o *MaskingExecutionMetrics) HasBytesMasked() bool`

HasBytesMasked returns a boolean if a field has been set.

### GetBytesTotal

`func (o *MaskingExecutionMetrics) GetBytesTotal() int64`

GetBytesTotal returns the BytesTotal field if non-nil, zero value otherwise.

### GetBytesTotalOk

`func (o *MaskingExecutionMetrics) GetBytesTotalOk() (*int64, bool)`

GetBytesTotalOk returns a tuple with the BytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesTotal

`func (o *MaskingExecutionMetrics) SetBytesTotal(v int64)`

SetBytesTotal sets BytesTotal field to given value.

### HasBytesTotal

`func (o *MaskingExecutionMetrics) HasBytesTotal() bool`

HasBytesTotal returns a boolean if a field has been set.

### GetDuration

`func (o *MaskingExecutionMetrics) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MaskingExecutionMetrics) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MaskingExecutionMetrics) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MaskingExecutionMetrics) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTablesFilesCount

`func (o *MaskingExecutionMetrics) GetTablesFilesCount() int64`

GetTablesFilesCount returns the TablesFilesCount field if non-nil, zero value otherwise.

### GetTablesFilesCountOk

`func (o *MaskingExecutionMetrics) GetTablesFilesCountOk() (*int64, bool)`

GetTablesFilesCountOk returns a tuple with the TablesFilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablesFilesCount

`func (o *MaskingExecutionMetrics) SetTablesFilesCount(v int64)`

SetTablesFilesCount sets TablesFilesCount field to given value.

### HasTablesFilesCount

`func (o *MaskingExecutionMetrics) HasTablesFilesCount() bool`

HasTablesFilesCount returns a boolean if a field has been set.

### GetMaskedTablesFilesCount

`func (o *MaskingExecutionMetrics) GetMaskedTablesFilesCount() int64`

GetMaskedTablesFilesCount returns the MaskedTablesFilesCount field if non-nil, zero value otherwise.

### GetMaskedTablesFilesCountOk

`func (o *MaskingExecutionMetrics) GetMaskedTablesFilesCountOk() (*int64, bool)`

GetMaskedTablesFilesCountOk returns a tuple with the MaskedTablesFilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedTablesFilesCount

`func (o *MaskingExecutionMetrics) SetMaskedTablesFilesCount(v int64)`

SetMaskedTablesFilesCount sets MaskedTablesFilesCount field to given value.

### HasMaskedTablesFilesCount

`func (o *MaskingExecutionMetrics) HasMaskedTablesFilesCount() bool`

HasMaskedTablesFilesCount returns a boolean if a field has been set.

### GetColumnsFieldsCount

`func (o *MaskingExecutionMetrics) GetColumnsFieldsCount() int64`

GetColumnsFieldsCount returns the ColumnsFieldsCount field if non-nil, zero value otherwise.

### GetColumnsFieldsCountOk

`func (o *MaskingExecutionMetrics) GetColumnsFieldsCountOk() (*int64, bool)`

GetColumnsFieldsCountOk returns a tuple with the ColumnsFieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnsFieldsCount

`func (o *MaskingExecutionMetrics) SetColumnsFieldsCount(v int64)`

SetColumnsFieldsCount sets ColumnsFieldsCount field to given value.

### HasColumnsFieldsCount

`func (o *MaskingExecutionMetrics) HasColumnsFieldsCount() bool`

HasColumnsFieldsCount returns a boolean if a field has been set.

### GetMaskedColumnsFieldsCount

`func (o *MaskingExecutionMetrics) GetMaskedColumnsFieldsCount() int64`

GetMaskedColumnsFieldsCount returns the MaskedColumnsFieldsCount field if non-nil, zero value otherwise.

### GetMaskedColumnsFieldsCountOk

`func (o *MaskingExecutionMetrics) GetMaskedColumnsFieldsCountOk() (*int64, bool)`

GetMaskedColumnsFieldsCountOk returns a tuple with the MaskedColumnsFieldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedColumnsFieldsCount

`func (o *MaskingExecutionMetrics) SetMaskedColumnsFieldsCount(v int64)`

SetMaskedColumnsFieldsCount sets MaskedColumnsFieldsCount field to given value.

### HasMaskedColumnsFieldsCount

`func (o *MaskingExecutionMetrics) HasMaskedColumnsFieldsCount() bool`

HasMaskedColumnsFieldsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


