# ApiUsageReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ApiUsageData**](ApiUsageData.md) |  | [optional] 
**TotalAutomationApiCount** | Pointer to **int64** | Total count of automation API calls over the requested timeframe. | [optional] 
**TotalManagementApiCount** | Pointer to **int64** | Total count of management API calls over the requested timeframe. | [optional] 

## Methods

### NewApiUsageReportResponse

`func NewApiUsageReportResponse() *ApiUsageReportResponse`

NewApiUsageReportResponse instantiates a new ApiUsageReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsageReportResponseWithDefaults

`func NewApiUsageReportResponseWithDefaults() *ApiUsageReportResponse`

NewApiUsageReportResponseWithDefaults instantiates a new ApiUsageReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ApiUsageReportResponse) GetItems() []ApiUsageData`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ApiUsageReportResponse) GetItemsOk() (*[]ApiUsageData, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ApiUsageReportResponse) SetItems(v []ApiUsageData)`

SetItems sets Items field to given value.

### HasItems

`func (o *ApiUsageReportResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalAutomationApiCount

`func (o *ApiUsageReportResponse) GetTotalAutomationApiCount() int64`

GetTotalAutomationApiCount returns the TotalAutomationApiCount field if non-nil, zero value otherwise.

### GetTotalAutomationApiCountOk

`func (o *ApiUsageReportResponse) GetTotalAutomationApiCountOk() (*int64, bool)`

GetTotalAutomationApiCountOk returns a tuple with the TotalAutomationApiCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAutomationApiCount

`func (o *ApiUsageReportResponse) SetTotalAutomationApiCount(v int64)`

SetTotalAutomationApiCount sets TotalAutomationApiCount field to given value.

### HasTotalAutomationApiCount

`func (o *ApiUsageReportResponse) HasTotalAutomationApiCount() bool`

HasTotalAutomationApiCount returns a boolean if a field has been set.

### GetTotalManagementApiCount

`func (o *ApiUsageReportResponse) GetTotalManagementApiCount() int64`

GetTotalManagementApiCount returns the TotalManagementApiCount field if non-nil, zero value otherwise.

### GetTotalManagementApiCountOk

`func (o *ApiUsageReportResponse) GetTotalManagementApiCountOk() (*int64, bool)`

GetTotalManagementApiCountOk returns a tuple with the TotalManagementApiCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalManagementApiCount

`func (o *ApiUsageReportResponse) SetTotalManagementApiCount(v int64)`

SetTotalManagementApiCount sets TotalManagementApiCount field to given value.

### HasTotalManagementApiCount

`func (o *ApiUsageReportResponse) HasTotalManagementApiCount() bool`

HasTotalManagementApiCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


