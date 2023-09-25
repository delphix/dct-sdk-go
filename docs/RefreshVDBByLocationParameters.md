# RefreshVDBByLocationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **string** | The database specific identifier for tracking transactions (SCN, LSN, etc). | [optional] 
**DatasetId** | Pointer to **string** | ID of the dataset to refresh to, mutually exclusive with timeflow_id. | [optional] 
**TimeflowId** | Pointer to **string** | ID of the timeflow to refresh to, mutually exclusive with dataset_id. | [optional] 

## Methods

### NewRefreshVDBByLocationParameters

`func NewRefreshVDBByLocationParameters() *RefreshVDBByLocationParameters`

NewRefreshVDBByLocationParameters instantiates a new RefreshVDBByLocationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshVDBByLocationParametersWithDefaults

`func NewRefreshVDBByLocationParametersWithDefaults() *RefreshVDBByLocationParameters`

NewRefreshVDBByLocationParametersWithDefaults instantiates a new RefreshVDBByLocationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *RefreshVDBByLocationParameters) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *RefreshVDBByLocationParameters) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *RefreshVDBByLocationParameters) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *RefreshVDBByLocationParameters) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDatasetId

`func (o *RefreshVDBByLocationParameters) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *RefreshVDBByLocationParameters) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *RefreshVDBByLocationParameters) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *RefreshVDBByLocationParameters) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetTimeflowId

`func (o *RefreshVDBByLocationParameters) GetTimeflowId() string`

GetTimeflowId returns the TimeflowId field if non-nil, zero value otherwise.

### GetTimeflowIdOk

`func (o *RefreshVDBByLocationParameters) GetTimeflowIdOk() (*string, bool)`

GetTimeflowIdOk returns a tuple with the TimeflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeflowId

`func (o *RefreshVDBByLocationParameters) SetTimeflowId(v string)`

SetTimeflowId sets TimeflowId field to given value.

### HasTimeflowId

`func (o *RefreshVDBByLocationParameters) HasTimeflowId() bool`

HasTimeflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


