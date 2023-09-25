# UpdateVDBGroupParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the VDB group. | [optional] 
**VdbIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateVDBGroupParameters

`func NewUpdateVDBGroupParameters() *UpdateVDBGroupParameters`

NewUpdateVDBGroupParameters instantiates a new UpdateVDBGroupParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVDBGroupParametersWithDefaults

`func NewUpdateVDBGroupParametersWithDefaults() *UpdateVDBGroupParameters`

NewUpdateVDBGroupParametersWithDefaults instantiates a new UpdateVDBGroupParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateVDBGroupParameters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDBGroupParameters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDBGroupParameters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVDBGroupParameters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVdbIds

`func (o *UpdateVDBGroupParameters) GetVdbIds() []string`

GetVdbIds returns the VdbIds field if non-nil, zero value otherwise.

### GetVdbIdsOk

`func (o *UpdateVDBGroupParameters) GetVdbIdsOk() (*[]string, bool)`

GetVdbIdsOk returns a tuple with the VdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbIds

`func (o *UpdateVDBGroupParameters) SetVdbIds(v []string)`

SetVdbIds sets VdbIds field to given value.

### HasVdbIds

`func (o *UpdateVDBGroupParameters) HasVdbIds() bool`

HasVdbIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


