# LinkDSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 
**DsourceId** | Pointer to **string** | The ID of the dSource. | [optional] 

## Methods

### NewLinkDSourceResponse

`func NewLinkDSourceResponse() *LinkDSourceResponse`

NewLinkDSourceResponse instantiates a new LinkDSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDSourceResponseWithDefaults

`func NewLinkDSourceResponseWithDefaults() *LinkDSourceResponse`

NewLinkDSourceResponseWithDefaults instantiates a new LinkDSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *LinkDSourceResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *LinkDSourceResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *LinkDSourceResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *LinkDSourceResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetDsourceId

`func (o *LinkDSourceResponse) GetDsourceId() string`

GetDsourceId returns the DsourceId field if non-nil, zero value otherwise.

### GetDsourceIdOk

`func (o *LinkDSourceResponse) GetDsourceIdOk() (*string, bool)`

GetDsourceIdOk returns a tuple with the DsourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsourceId

`func (o *LinkDSourceResponse) SetDsourceId(v string)`

SetDsourceId sets DsourceId field to given value.

### HasDsourceId

`func (o *LinkDSourceResponse) HasDsourceId() bool`

HasDsourceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


