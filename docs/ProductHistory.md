# ProductHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Product Version. | [optional] 
**InstalledOn** | Pointer to **time.Time** | This version installed on date. | [optional] 

## Methods

### NewProductHistory

`func NewProductHistory() *ProductHistory`

NewProductHistory instantiates a new ProductHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductHistoryWithDefaults

`func NewProductHistoryWithDefaults() *ProductHistory`

NewProductHistoryWithDefaults instantiates a new ProductHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ProductHistory) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProductHistory) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProductHistory) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProductHistory) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetInstalledOn

`func (o *ProductHistory) GetInstalledOn() time.Time`

GetInstalledOn returns the InstalledOn field if non-nil, zero value otherwise.

### GetInstalledOnOk

`func (o *ProductHistory) GetInstalledOnOk() (*time.Time, bool)`

GetInstalledOnOk returns a tuple with the InstalledOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledOn

`func (o *ProductHistory) SetInstalledOn(v time.Time)`

SetInstalledOn sets InstalledOn field to given value.

### HasInstalledOn

`func (o *ProductHistory) HasInstalledOn() bool`

HasInstalledOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


