# AdditionalMountPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedPath** | Pointer to **string** | Relative path within the container of the directory that should be mounted. | [optional] 
**MountPath** | Pointer to **string** | Absolute path on the target environment were the filesystem should be mounted | [optional] 
**EnvironmentId** | Pointer to **string** | The entity ID of the environment on which the file system will be mounted. | [optional] 

## Methods

### NewAdditionalMountPoint

`func NewAdditionalMountPoint() *AdditionalMountPoint`

NewAdditionalMountPoint instantiates a new AdditionalMountPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalMountPointWithDefaults

`func NewAdditionalMountPointWithDefaults() *AdditionalMountPoint`

NewAdditionalMountPointWithDefaults instantiates a new AdditionalMountPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedPath

`func (o *AdditionalMountPoint) GetSharedPath() string`

GetSharedPath returns the SharedPath field if non-nil, zero value otherwise.

### GetSharedPathOk

`func (o *AdditionalMountPoint) GetSharedPathOk() (*string, bool)`

GetSharedPathOk returns a tuple with the SharedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPath

`func (o *AdditionalMountPoint) SetSharedPath(v string)`

SetSharedPath sets SharedPath field to given value.

### HasSharedPath

`func (o *AdditionalMountPoint) HasSharedPath() bool`

HasSharedPath returns a boolean if a field has been set.

### GetMountPath

`func (o *AdditionalMountPoint) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *AdditionalMountPoint) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *AdditionalMountPoint) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *AdditionalMountPoint) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *AdditionalMountPoint) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *AdditionalMountPoint) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *AdditionalMountPoint) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *AdditionalMountPoint) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


