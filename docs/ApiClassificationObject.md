# ApiClassificationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiMethod** | Pointer to **string** | HTTP method of the API. | [optional] 
**Path** | Pointer to **string** | context path of the API. | [optional] 
**IsAutomation** | Pointer to **bool** | Either this API is automation or not. | [optional] 
**StartDate** | Pointer to **time.Time** | The start date and time from when this api&#39;s is_automation definition has changed. | [optional] 
**EndDate** | Pointer to **time.Time** | The end date and time from when this api&#39;s is_automation definition has changed. | [optional] 

## Methods

### NewApiClassificationObject

`func NewApiClassificationObject() *ApiClassificationObject`

NewApiClassificationObject instantiates a new ApiClassificationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClassificationObjectWithDefaults

`func NewApiClassificationObjectWithDefaults() *ApiClassificationObject`

NewApiClassificationObjectWithDefaults instantiates a new ApiClassificationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiMethod

`func (o *ApiClassificationObject) GetApiMethod() string`

GetApiMethod returns the ApiMethod field if non-nil, zero value otherwise.

### GetApiMethodOk

`func (o *ApiClassificationObject) GetApiMethodOk() (*string, bool)`

GetApiMethodOk returns a tuple with the ApiMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMethod

`func (o *ApiClassificationObject) SetApiMethod(v string)`

SetApiMethod sets ApiMethod field to given value.

### HasApiMethod

`func (o *ApiClassificationObject) HasApiMethod() bool`

HasApiMethod returns a boolean if a field has been set.

### GetPath

`func (o *ApiClassificationObject) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ApiClassificationObject) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ApiClassificationObject) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ApiClassificationObject) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetIsAutomation

`func (o *ApiClassificationObject) GetIsAutomation() bool`

GetIsAutomation returns the IsAutomation field if non-nil, zero value otherwise.

### GetIsAutomationOk

`func (o *ApiClassificationObject) GetIsAutomationOk() (*bool, bool)`

GetIsAutomationOk returns a tuple with the IsAutomation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomation

`func (o *ApiClassificationObject) SetIsAutomation(v bool)`

SetIsAutomation sets IsAutomation field to given value.

### HasIsAutomation

`func (o *ApiClassificationObject) HasIsAutomation() bool`

HasIsAutomation returns a boolean if a field has been set.

### GetStartDate

`func (o *ApiClassificationObject) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ApiClassificationObject) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ApiClassificationObject) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ApiClassificationObject) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ApiClassificationObject) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ApiClassificationObject) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ApiClassificationObject) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ApiClassificationObject) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


