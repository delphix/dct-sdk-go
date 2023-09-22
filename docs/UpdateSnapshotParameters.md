# UpdateSnapshotParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **string** | The expiration for this snapshot. Mutually exclusive with retain_forever. | [optional] 
**RetainForever** | Pointer to **bool** | Indicates that the snapshot should be retained forever. | [optional] 

## Methods

### NewUpdateSnapshotParameters

`func NewUpdateSnapshotParameters() *UpdateSnapshotParameters`

NewUpdateSnapshotParameters instantiates a new UpdateSnapshotParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotParametersWithDefaults

`func NewUpdateSnapshotParametersWithDefaults() *UpdateSnapshotParameters`

NewUpdateSnapshotParametersWithDefaults instantiates a new UpdateSnapshotParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *UpdateSnapshotParameters) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *UpdateSnapshotParameters) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *UpdateSnapshotParameters) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *UpdateSnapshotParameters) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRetainForever

`func (o *UpdateSnapshotParameters) GetRetainForever() bool`

GetRetainForever returns the RetainForever field if non-nil, zero value otherwise.

### GetRetainForeverOk

`func (o *UpdateSnapshotParameters) GetRetainForeverOk() (*bool, bool)`

GetRetainForeverOk returns a tuple with the RetainForever field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainForever

`func (o *UpdateSnapshotParameters) SetRetainForever(v bool)`

SetRetainForever sets RetainForever field to given value.

### HasRetainForever

`func (o *UpdateSnapshotParameters) HasRetainForever() bool`

HasRetainForever returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


