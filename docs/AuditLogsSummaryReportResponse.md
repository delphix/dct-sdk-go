# AuditLogsSummaryReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AuditLogsSummary**](AuditLogsSummary.md) |  | [optional] 
**ResponseMetadata** | Pointer to [**PaginatedResponseMetadata**](PaginatedResponseMetadata.md) |  | [optional] 
**Totals** | Pointer to [**AuditLogsSummaryTotals**](AuditLogsSummaryTotals.md) |  | [optional] 

## Methods

### NewAuditLogsSummaryReportResponse

`func NewAuditLogsSummaryReportResponse() *AuditLogsSummaryReportResponse`

NewAuditLogsSummaryReportResponse instantiates a new AuditLogsSummaryReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsSummaryReportResponseWithDefaults

`func NewAuditLogsSummaryReportResponseWithDefaults() *AuditLogsSummaryReportResponse`

NewAuditLogsSummaryReportResponseWithDefaults instantiates a new AuditLogsSummaryReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AuditLogsSummaryReportResponse) GetItems() []AuditLogsSummary`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AuditLogsSummaryReportResponse) GetItemsOk() (*[]AuditLogsSummary, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AuditLogsSummaryReportResponse) SetItems(v []AuditLogsSummary)`

SetItems sets Items field to given value.

### HasItems

`func (o *AuditLogsSummaryReportResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetResponseMetadata

`func (o *AuditLogsSummaryReportResponse) GetResponseMetadata() PaginatedResponseMetadata`

GetResponseMetadata returns the ResponseMetadata field if non-nil, zero value otherwise.

### GetResponseMetadataOk

`func (o *AuditLogsSummaryReportResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool)`

GetResponseMetadataOk returns a tuple with the ResponseMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMetadata

`func (o *AuditLogsSummaryReportResponse) SetResponseMetadata(v PaginatedResponseMetadata)`

SetResponseMetadata sets ResponseMetadata field to given value.

### HasResponseMetadata

`func (o *AuditLogsSummaryReportResponse) HasResponseMetadata() bool`

HasResponseMetadata returns a boolean if a field has been set.

### GetTotals

`func (o *AuditLogsSummaryReportResponse) GetTotals() AuditLogsSummaryTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *AuditLogsSummaryReportResponse) GetTotalsOk() (*AuditLogsSummaryTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *AuditLogsSummaryReportResponse) SetTotals(v AuditLogsSummaryTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *AuditLogsSummaryReportResponse) HasTotals() bool`

HasTotals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


