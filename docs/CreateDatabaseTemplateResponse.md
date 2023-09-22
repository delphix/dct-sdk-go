# CreateDatabaseTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseTemplate** | Pointer to [**DatabaseTemplate**](DatabaseTemplate.md) |  | [optional] 
**Job** | Pointer to [**Job**](Job.md) |  | [optional] 

## Methods

### NewCreateDatabaseTemplateResponse

`func NewCreateDatabaseTemplateResponse() *CreateDatabaseTemplateResponse`

NewCreateDatabaseTemplateResponse instantiates a new CreateDatabaseTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatabaseTemplateResponseWithDefaults

`func NewCreateDatabaseTemplateResponseWithDefaults() *CreateDatabaseTemplateResponse`

NewCreateDatabaseTemplateResponseWithDefaults instantiates a new CreateDatabaseTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseTemplate

`func (o *CreateDatabaseTemplateResponse) GetDatabaseTemplate() DatabaseTemplate`

GetDatabaseTemplate returns the DatabaseTemplate field if non-nil, zero value otherwise.

### GetDatabaseTemplateOk

`func (o *CreateDatabaseTemplateResponse) GetDatabaseTemplateOk() (*DatabaseTemplate, bool)`

GetDatabaseTemplateOk returns a tuple with the DatabaseTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseTemplate

`func (o *CreateDatabaseTemplateResponse) SetDatabaseTemplate(v DatabaseTemplate)`

SetDatabaseTemplate sets DatabaseTemplate field to given value.

### HasDatabaseTemplate

`func (o *CreateDatabaseTemplateResponse) HasDatabaseTemplate() bool`

HasDatabaseTemplate returns a boolean if a field has been set.

### GetJob

`func (o *CreateDatabaseTemplateResponse) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *CreateDatabaseTemplateResponse) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *CreateDatabaseTemplateResponse) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *CreateDatabaseTemplateResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


