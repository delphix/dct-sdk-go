# ConfigSettingsStg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyName** | Pointer to **string** | Name of the property. | [optional] 
**Value** | Pointer to **string** | Value of the property. | [optional] 
**CommentProperty** | Pointer to **bool** | Select this option to comment out the provided property name in the configuration file. | [optional] 

## Methods

### NewConfigSettingsStg

`func NewConfigSettingsStg() *ConfigSettingsStg`

NewConfigSettingsStg instantiates a new ConfigSettingsStg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingsStgWithDefaults

`func NewConfigSettingsStgWithDefaults() *ConfigSettingsStg`

NewConfigSettingsStgWithDefaults instantiates a new ConfigSettingsStg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyName

`func (o *ConfigSettingsStg) GetPropertyName() string`

GetPropertyName returns the PropertyName field if non-nil, zero value otherwise.

### GetPropertyNameOk

`func (o *ConfigSettingsStg) GetPropertyNameOk() (*string, bool)`

GetPropertyNameOk returns a tuple with the PropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyName

`func (o *ConfigSettingsStg) SetPropertyName(v string)`

SetPropertyName sets PropertyName field to given value.

### HasPropertyName

`func (o *ConfigSettingsStg) HasPropertyName() bool`

HasPropertyName returns a boolean if a field has been set.

### GetValue

`func (o *ConfigSettingsStg) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigSettingsStg) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigSettingsStg) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigSettingsStg) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCommentProperty

`func (o *ConfigSettingsStg) GetCommentProperty() bool`

GetCommentProperty returns the CommentProperty field if non-nil, zero value otherwise.

### GetCommentPropertyOk

`func (o *ConfigSettingsStg) GetCommentPropertyOk() (*bool, bool)`

GetCommentPropertyOk returns a tuple with the CommentProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentProperty

`func (o *ConfigSettingsStg) SetCommentProperty(v bool)`

SetCommentProperty sets CommentProperty field to given value.

### HasCommentProperty

`func (o *ConfigSettingsStg) HasCommentProperty() bool`

HasCommentProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


