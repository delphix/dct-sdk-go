# RefreshVDBByTimestampParametersAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** | ID of the dataset to refresh to, mutually exclusive with timeflow_id. | [optional] 
**TimeflowId** | Pointer to **string** | ID of the timeflow to refresh to, mutually exclusive with dataset_id. | [optional] 

## Methods

### NewRefreshVDBByTimestampParametersAllOf

`func NewRefreshVDBByTimestampParametersAllOf() *RefreshVDBByTimestampParametersAllOf`

NewRefreshVDBByTimestampParametersAllOf instantiates a new RefreshVDBByTimestampParametersAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshVDBByTimestampParametersAllOfWithDefaults

`func NewRefreshVDBByTimestampParametersAllOfWithDefaults() *RefreshVDBByTimestampParametersAllOf`

NewRefreshVDBByTimestampParametersAllOfWithDefaults instantiates a new RefreshVDBByTimestampParametersAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetId

`func (o *RefreshVDBByTimestampParametersAllOf) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *RefreshVDBByTimestampParametersAllOf) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *RefreshVDBByTimestampParametersAllOf) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *RefreshVDBByTimestampParametersAllOf) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetTimeflowId

`func (o *RefreshVDBByTimestampParametersAllOf) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *RefreshVDBByTimestampParametersAllOf) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *RefreshVDBByTimestampParametersAllOf) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *RefreshVDBByTimestampParametersAllOf) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


