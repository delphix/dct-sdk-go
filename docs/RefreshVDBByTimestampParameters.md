# RefreshVDBByTimestampParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point. | [optional] 
**TimestampInDatabaseTimezone** | Pointer to **string** | The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp. | [optional] 
**TimeflowId** | Pointer to **string** | ID of the timeflow to refresh to, mutually exclusive with dataset_id. | [optional] 
**DatasetId** | Pointer to **string** | ID of the dataset to refresh to, mutually exclusive with timeflow_id. | [optional] 

## Methods

### NewRefreshVDBByTimestampParameters

`func NewRefreshVDBByTimestampParameters() *RefreshVDBByTimestampParameters`

NewRefreshVDBByTimestampParameters instantiates a new RefreshVDBByTimestampParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshVDBByTimestampParametersWithDefaults

`func NewRefreshVDBByTimestampParametersWithDefaults() *RefreshVDBByTimestampParameters`

NewRefreshVDBByTimestampParametersWithDefaults instantiates a new RefreshVDBByTimestampParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *RefreshVDBByTimestampParameters) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RefreshVDBByTimestampParameters) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RefreshVDBByTimestampParameters) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RefreshVDBByTimestampParameters) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampInDatabaseTimezone

`func (o *RefreshVDBByTimestampParameters) GetTimestampInDatabaseTimezone() string`

GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field if non-nil, zero value otherwise.

### GetTimestampInDatabaseTimezoneOk

`func (o *RefreshVDBByTimestampParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool)`

GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampInDatabaseTimezone

`func (o *RefreshVDBByTimestampParameters) SetTimestampInDatabaseTimezone(v string)`

SetTimestampInDatabaseTimezone sets TimestampInDatabaseTimezone field to given value.

### HasTimestampInDatabaseTimezone

`func (o *RefreshVDBByTimestampParameters) HasTimestampInDatabaseTimezone() bool`

HasTimestampInDatabaseTimezone returns a boolean if a field has been set.

### GetTimeflowId

`func (o *RefreshVDBByTimestampParameters) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *RefreshVDBByTimestampParameters) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *RefreshVDBByTimestampParameters) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *RefreshVDBByTimestampParameters) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.

### GetDatasetId

`func (o *RefreshVDBByTimestampParameters) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *RefreshVDBByTimestampParameters) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *RefreshVDBByTimestampParameters) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *RefreshVDBByTimestampParameters) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


