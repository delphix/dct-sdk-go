# AllObjectPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectPermissions** | Pointer to [**[]PermissionObject**](PermissionObject.md) |  | [optional] 

## Methods

### NewAllObjectPermissionsResponse

`func NewAllObjectPermissionsResponse() *AllObjectPermissionsResponse`

NewAllObjectPermissionsResponse instantiates a new AllObjectPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllObjectPermissionsResponseWithDefaults

`func NewAllObjectPermissionsResponseWithDefaults() *AllObjectPermissionsResponse`

NewAllObjectPermissionsResponseWithDefaults instantiates a new AllObjectPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectPermissions

`func (o *AllObjectPermissionsResponse) GetObjectPermissions() []PermissionObject`

GetObjectPermissions returns the ObjectPermissions field if non-nil, zero value otherwise.

### GetObjectPermissionsOk

`func (o *AllObjectPermissionsResponse) GetObjectPermissionsOk() (*[]PermissionObject, bool)`

GetObjectPermissionsOk returns a tuple with the ObjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPermissions

`func (o *AllObjectPermissionsResponse) SetObjectPermissions(v []PermissionObject)`

SetObjectPermissions sets ObjectPermissions field to given value.

### HasObjectPermissions

`func (o *AllObjectPermissionsResponse) HasObjectPermissions() bool`

HasObjectPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


