# AuditLogsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** |  | [optional] 
**AccountFirstName** | Pointer to **string** |  | [optional] 
**AccountLastName** | Pointer to **string** |  | [optional] 
**VdbRefreshes** | Pointer to **int32** | The number of VDB refreshes performed by the account. | [optional] 
**MaskingJobs** | Pointer to **int32** | The number of compliance jobs executed by the account. | [optional] 

## Methods

### NewAuditLogsSummary

`func NewAuditLogsSummary() *AuditLogsSummary`

NewAuditLogsSummary instantiates a new AuditLogsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogsSummaryWithDefaults

`func NewAuditLogsSummaryWithDefaults() *AuditLogsSummary`

NewAuditLogsSummaryWithDefaults instantiates a new AuditLogsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AuditLogsSummary) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AuditLogsSummary) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AuditLogsSummary) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AuditLogsSummary) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountFirstName

`func (o *AuditLogsSummary) GetAccountFirstName() string`

GetAccountFirstName returns the AccountFirstName field if non-nil, zero value otherwise.

### GetAccountFirstNameOk

`func (o *AuditLogsSummary) GetAccountFirstNameOk() (*string, bool)`

GetAccountFirstNameOk returns a tuple with the AccountFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFirstName

`func (o *AuditLogsSummary) SetAccountFirstName(v string)`

SetAccountFirstName sets AccountFirstName field to given value.

### HasAccountFirstName

`func (o *AuditLogsSummary) HasAccountFirstName() bool`

HasAccountFirstName returns a boolean if a field has been set.

### GetAccountLastName

`func (o *AuditLogsSummary) GetAccountLastName() string`

GetAccountLastName returns the AccountLastName field if non-nil, zero value otherwise.

### GetAccountLastNameOk

`func (o *AuditLogsSummary) GetAccountLastNameOk() (*string, bool)`

GetAccountLastNameOk returns a tuple with the AccountLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLastName

`func (o *AuditLogsSummary) SetAccountLastName(v string)`

SetAccountLastName sets AccountLastName field to given value.

### HasAccountLastName

`func (o *AuditLogsSummary) HasAccountLastName() bool`

HasAccountLastName returns a boolean if a field has been set.

### GetVdbRefreshes

`func (o *AuditLogsSummary) GetVdbRefreshes() int32`

GetVdbRefreshes returns the VdbRefreshes field if non-nil, zero value otherwise.

### GetVdbRefreshesOk

`func (o *AuditLogsSummary) GetVdbRefreshesOk() (*int32, bool)`

GetVdbRefreshesOk returns a tuple with the VdbRefreshes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbRefreshes

`func (o *AuditLogsSummary) SetVdbRefreshes(v int32)`

SetVdbRefreshes sets VdbRefreshes field to given value.

### HasVdbRefreshes

`func (o *AuditLogsSummary) HasVdbRefreshes() bool`

HasVdbRefreshes returns a boolean if a field has been set.

### GetMaskingJobs

`func (o *AuditLogsSummary) GetMaskingJobs() int32`

GetMaskingJobs returns the MaskingJobs field if non-nil, zero value otherwise.

### GetMaskingJobsOk

`func (o *AuditLogsSummary) GetMaskingJobsOk() (*int32, bool)`

GetMaskingJobsOk returns a tuple with the MaskingJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingJobs

`func (o *AuditLogsSummary) SetMaskingJobs(v int32)`

SetMaskingJobs sets MaskingJobs field to given value.

### HasMaskingJobs

`func (o *AuditLogsSummary) HasMaskingJobs() bool`

HasMaskingJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


