# HyperscaleTaskError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** | the name of the table for which the error occurred, including the schema. | [optional] 
**Error** | Pointer to **string** | A textual description of the error. | [optional] 

## Methods

### NewHyperscaleTaskError

`func NewHyperscaleTaskError() *HyperscaleTaskError`

NewHyperscaleTaskError instantiates a new HyperscaleTaskError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleTaskErrorWithDefaults

`func NewHyperscaleTaskErrorWithDefaults() *HyperscaleTaskError`

NewHyperscaleTaskErrorWithDefaults instantiates a new HyperscaleTaskError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *HyperscaleTaskError) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *HyperscaleTaskError) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *HyperscaleTaskError) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *HyperscaleTaskError) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetError

`func (o *HyperscaleTaskError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HyperscaleTaskError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HyperscaleTaskError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *HyperscaleTaskError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


