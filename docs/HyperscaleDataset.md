# HyperscaleDataset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Hyperscale Dataset. | [optional] 
**HyperscaleInstanceId** | Pointer to **string** | The ID of the Hyperscale instance of this Dataset. | [optional] 
**MountPointId** | Pointer to **string** | The Id of the Hyperscale Mount Point used for this Dataset. | [optional] 
**ConnectorId** | Pointer to **string** | Id the Hyperscale Connector used to read sensitive data and write masked data. | [optional] 

## Methods

### NewHyperscaleDataset

`func NewHyperscaleDataset() *HyperscaleDataset`

NewHyperscaleDataset instantiates a new HyperscaleDataset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleDatasetWithDefaults

`func NewHyperscaleDatasetWithDefaults() *HyperscaleDataset`

NewHyperscaleDatasetWithDefaults instantiates a new HyperscaleDataset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HyperscaleDataset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperscaleDataset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperscaleDataset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HyperscaleDataset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHyperscaleInstanceId

`func (o *HyperscaleDataset) GetHyperscaleInstanceId() string`

GetHyperscaleInstanceId returns the HyperscaleInstanceId field if non-nil, zero value otherwise.

### GetHyperscaleInstanceIdOk

`func (o *HyperscaleDataset) GetHyperscaleInstanceIdOk() (*string, bool)`

GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperscaleInstanceId

`func (o *HyperscaleDataset) SetHyperscaleInstanceId(v string)`

SetHyperscaleInstanceId sets HyperscaleInstanceId field to given value.

### HasHyperscaleInstanceId

`func (o *HyperscaleDataset) HasHyperscaleInstanceId() bool`

HasHyperscaleInstanceId returns a boolean if a field has been set.

### GetMountPointId

`func (o *HyperscaleDataset) GetMountPointId() string`

GetMountPointId returns the MountPointId field if non-nil, zero value otherwise.

### GetMountPointIdOk

`func (o *HyperscaleDataset) GetMountPointIdOk() (*string, bool)`

GetMountPointIdOk returns a tuple with the MountPointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPointId

`func (o *HyperscaleDataset) SetMountPointId(v string)`

SetMountPointId sets MountPointId field to given value.

### HasMountPointId

`func (o *HyperscaleDataset) HasMountPointId() bool`

HasMountPointId returns a boolean if a field has been set.

### GetConnectorId

`func (o *HyperscaleDataset) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *HyperscaleDataset) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *HyperscaleDataset) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *HyperscaleDataset) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


