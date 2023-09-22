# ApiUsageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiEndpoint** | Pointer to **string** | API called. | [optional] 
**ApiMethod** | Pointer to **string** | HTTP method for API called. | [optional] 
**ApiCount** | **int64** | Count of API calls over the requested timeframe. | 
**Kind** | Pointer to **string** | Whether the API calls are of kind automation or governance | [optional] 
**ClientName** | Pointer to **string** | Name of the api client that called the API endpoint | [optional] 
**UserAgent** | Pointer to **string** | Version of the api client that called the API endpoint | [optional] 
**DctVersion** | Pointer to **string** | DCT version at the time of api call | [optional] 

## Methods

### NewApiUsageData

`func NewApiUsageData(apiCount int64, ) *ApiUsageData`

NewApiUsageData instantiates a new ApiUsageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsageDataWithDefaults

`func NewApiUsageDataWithDefaults() *ApiUsageData`

NewApiUsageDataWithDefaults instantiates a new ApiUsageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiEndpoint

`func (o *ApiUsageData) GetApiEndpoint() string`

GetApiEndpoint returns the ApiEndpoint field if non-nil, zero value otherwise.

### GetApiEndpointOk

`func (o *ApiUsageData) GetApiEndpointOk() (*string, bool)`

GetApiEndpointOk returns a tuple with the ApiEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndpoint

`func (o *ApiUsageData) SetApiEndpoint(v string)`

SetApiEndpoint sets ApiEndpoint field to given value.

### HasApiEndpoint

`func (o *ApiUsageData) HasApiEndpoint() bool`

HasApiEndpoint returns a boolean if a field has been set.

### GetApiMethod

`func (o *ApiUsageData) GetApiMethod() string`

GetApiMethod returns the ApiMethod field if non-nil, zero value otherwise.

### GetApiMethodOk

`func (o *ApiUsageData) GetApiMethodOk() (*string, bool)`

GetApiMethodOk returns a tuple with the ApiMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMethod

`func (o *ApiUsageData) SetApiMethod(v string)`

SetApiMethod sets ApiMethod field to given value.

### HasApiMethod

`func (o *ApiUsageData) HasApiMethod() bool`

HasApiMethod returns a boolean if a field has been set.

### GetApiCount

`func (o *ApiUsageData) GetApiCount() int64`

GetApiCount returns the ApiCount field if non-nil, zero value otherwise.

### GetApiCountOk

`func (o *ApiUsageData) GetApiCountOk() (*int64, bool)`

GetApiCountOk returns a tuple with the ApiCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCount

`func (o *ApiUsageData) SetApiCount(v int64)`

SetApiCount sets ApiCount field to given value.


### GetKind

`func (o *ApiUsageData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiUsageData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiUsageData) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiUsageData) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetClientName

`func (o *ApiUsageData) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ApiUsageData) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ApiUsageData) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ApiUsageData) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetUserAgent

`func (o *ApiUsageData) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ApiUsageData) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ApiUsageData) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ApiUsageData) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetDctVersion

`func (o *ApiUsageData) GetDctVersion() string`

GetDctVersion returns the DctVersion field if non-nil, zero value otherwise.

### GetDctVersionOk

`func (o *ApiUsageData) GetDctVersionOk() (*string, bool)`

GetDctVersionOk returns a tuple with the DctVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDctVersion

`func (o *ApiUsageData) SetDctVersion(v string)`

SetDctVersion sets DctVersion field to given value.

### HasDctVersion

`func (o *ApiUsageData) HasDctVersion() bool`

HasDctVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


