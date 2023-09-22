# HyperscaleMountPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Hyperscale Mount Point. | [optional] 
**HyperscaleInstanceId** | Pointer to **string** | The ID of the Hyperscale instance of this Mount Point. | [optional] 
**Name** | Pointer to **string** | Name of the mount, unique for a hyperscale instance. This name will be used as a directory name by the Hyperscale instance. | [optional] 
**Hostname** | Pointer to **string** | The host name of the server. | [optional] 
**MountPath** | Pointer to **string** | The path to the directory on the filesystem to mount. | [optional] 
**MountType** | Pointer to **string** | The type of filesystem. Enum having values- CIFS, NFS3, NFS4. | [optional] 
**Options** | Pointer to **string** | The options for mount. The endpoint will return all default options and user specified options. | [optional] 

## Methods

### NewHyperscaleMountPoint

`func NewHyperscaleMountPoint() *HyperscaleMountPoint`

NewHyperscaleMountPoint instantiates a new HyperscaleMountPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleMountPointWithDefaults

`func NewHyperscaleMountPointWithDefaults() *HyperscaleMountPoint`

NewHyperscaleMountPointWithDefaults instantiates a new HyperscaleMountPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HyperscaleMountPoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperscaleMountPoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperscaleMountPoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HyperscaleMountPoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHyperscaleInstanceId

`func (o *HyperscaleMountPoint) GetHyperscaleInstanceId() string`

GetHyperscaleInstanceId returns the HyperscaleInstanceId field if non-nil, zero value otherwise.

### GetHyperscaleInstanceIdOk

`func (o *HyperscaleMountPoint) GetHyperscaleInstanceIdOk() (*string, bool)`

GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperscaleInstanceId

`func (o *HyperscaleMountPoint) SetHyperscaleInstanceId(v string)`

SetHyperscaleInstanceId sets HyperscaleInstanceId field to given value.

### HasHyperscaleInstanceId

`func (o *HyperscaleMountPoint) HasHyperscaleInstanceId() bool`

HasHyperscaleInstanceId returns a boolean if a field has been set.

### GetName

`func (o *HyperscaleMountPoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperscaleMountPoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperscaleMountPoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperscaleMountPoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHostname

`func (o *HyperscaleMountPoint) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HyperscaleMountPoint) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HyperscaleMountPoint) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HyperscaleMountPoint) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetMountPath

`func (o *HyperscaleMountPoint) GetMountPath() string`

GetMountPath returns the MountPath field if non-nil, zero value otherwise.

### GetMountPathOk

`func (o *HyperscaleMountPoint) GetMountPathOk() (*string, bool)`

GetMountPathOk returns a tuple with the MountPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPath

`func (o *HyperscaleMountPoint) SetMountPath(v string)`

SetMountPath sets MountPath field to given value.

### HasMountPath

`func (o *HyperscaleMountPoint) HasMountPath() bool`

HasMountPath returns a boolean if a field has been set.

### GetMountType

`func (o *HyperscaleMountPoint) GetMountType() string`

GetMountType returns the MountType field if non-nil, zero value otherwise.

### GetMountTypeOk

`func (o *HyperscaleMountPoint) GetMountTypeOk() (*string, bool)`

GetMountTypeOk returns a tuple with the MountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountType

`func (o *HyperscaleMountPoint) SetMountType(v string)`

SetMountType sets MountType field to given value.

### HasMountType

`func (o *HyperscaleMountPoint) HasMountType() bool`

HasMountType returns a boolean if a field has been set.

### GetOptions

`func (o *HyperscaleMountPoint) GetOptions() string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *HyperscaleMountPoint) GetOptionsOk() (*string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *HyperscaleMountPoint) SetOptions(v string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *HyperscaleMountPoint) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


