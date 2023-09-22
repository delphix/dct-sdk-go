# ProductInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | Current API version. | [optional] 
**ProductVersion** | Pointer to **string** | Current installed product version. | [optional] 
**ProductUpgradeHistory** | Pointer to [**[]ProductHistory**](ProductHistory.md) | Product upgrade history. | [optional] 
**SupportedApiVersions** | Pointer to **[]string** | All the supported API versions. | [optional] 

## Methods

### NewProductInfo

`func NewProductInfo() *ProductInfo`

NewProductInfo instantiates a new ProductInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductInfoWithDefaults

`func NewProductInfoWithDefaults() *ProductInfo`

NewProductInfoWithDefaults instantiates a new ProductInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *ProductInfo) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ProductInfo) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ProductInfo) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ProductInfo) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetProductVersion

`func (o *ProductInfo) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ProductInfo) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ProductInfo) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ProductInfo) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetProductUpgradeHistory

`func (o *ProductInfo) GetProductUpgradeHistory() []ProductHistory`

GetProductUpgradeHistory returns the ProductUpgradeHistory field if non-nil, zero value otherwise.

### GetProductUpgradeHistoryOk

`func (o *ProductInfo) GetProductUpgradeHistoryOk() (*[]ProductHistory, bool)`

GetProductUpgradeHistoryOk returns a tuple with the ProductUpgradeHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductUpgradeHistory

`func (o *ProductInfo) SetProductUpgradeHistory(v []ProductHistory)`

SetProductUpgradeHistory sets ProductUpgradeHistory field to given value.

### HasProductUpgradeHistory

`func (o *ProductInfo) HasProductUpgradeHistory() bool`

HasProductUpgradeHistory returns a boolean if a field has been set.

### GetSupportedApiVersions

`func (o *ProductInfo) GetSupportedApiVersions() []string`

GetSupportedApiVersions returns the SupportedApiVersions field if non-nil, zero value otherwise.

### GetSupportedApiVersionsOk

`func (o *ProductInfo) GetSupportedApiVersionsOk() (*[]string, bool)`

GetSupportedApiVersionsOk returns a tuple with the SupportedApiVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedApiVersions

`func (o *ProductInfo) SetSupportedApiVersions(v []string)`

SetSupportedApiVersions sets SupportedApiVersions field to given value.

### HasSupportedApiVersions

`func (o *ProductInfo) HasSupportedApiVersions() bool`

HasSupportedApiVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


