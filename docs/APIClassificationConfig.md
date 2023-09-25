# APIClassificationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Api Classification Config Version. | [optional] 
**ApiClassification** | Pointer to [**[]ApiClassificationObject**](ApiClassificationObject.md) | The classification of each APIs, either it is automation or not. | [optional] 

## Methods

### NewAPIClassificationConfig

`func NewAPIClassificationConfig() *APIClassificationConfig`

NewAPIClassificationConfig instantiates a new APIClassificationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIClassificationConfigWithDefaults

`func NewAPIClassificationConfigWithDefaults() *APIClassificationConfig`

NewAPIClassificationConfigWithDefaults instantiates a new APIClassificationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *APIClassificationConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *APIClassificationConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *APIClassificationConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *APIClassificationConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetApiClassification

`func (o *APIClassificationConfig) GetApiClassification() []ApiClassificationObject`

GetApiClassification returns the ApiClassification field if non-nil, zero value otherwise.

### GetApiClassificationOk

`func (o *APIClassificationConfig) GetApiClassificationOk() (*[]ApiClassificationObject, bool)`

GetApiClassificationOk returns a tuple with the ApiClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClassification

`func (o *APIClassificationConfig) SetApiClassification(v []ApiClassificationObject)`

SetApiClassification sets ApiClassification field to given value.

### HasApiClassification

`func (o *APIClassificationConfig) HasApiClassification() bool`

HasApiClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


