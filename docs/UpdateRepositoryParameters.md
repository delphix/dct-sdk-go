# UpdateRepositoryParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseType** | **string** | The database type of this repository. | 
**AllowProvisioning** | Pointer to **bool** | Flag indicating whether the repository should be used for provisioning. | [optional] 
**IsStaging** | Pointer to **bool** | Flag indicating whether this repository can be used by the Delphix Engine for internal processing. | [optional] 
**Version** | Pointer to **string** | Version of the repository. | [optional] 

## Methods

### NewUpdateRepositoryParameters

`func NewUpdateRepositoryParameters(databaseType string, ) *UpdateRepositoryParameters`

NewUpdateRepositoryParameters instantiates a new UpdateRepositoryParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRepositoryParametersWithDefaults

`func NewUpdateRepositoryParametersWithDefaults() *UpdateRepositoryParameters`

NewUpdateRepositoryParametersWithDefaults instantiates a new UpdateRepositoryParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseType

`func (o *UpdateRepositoryParameters) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *UpdateRepositoryParameters) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *UpdateRepositoryParameters) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.


### GetAllowProvisioning

`func (o *UpdateRepositoryParameters) GetAllowProvisioning() bool`

GetAllowProvisioning returns the AllowProvisioning field if non-nil, zero value otherwise.

### GetAllowProvisioningOk

`func (o *UpdateRepositoryParameters) GetAllowProvisioningOk() (*bool, bool)`

GetAllowProvisioningOk returns a tuple with the AllowProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProvisioning

`func (o *UpdateRepositoryParameters) SetAllowProvisioning(v bool)`

SetAllowProvisioning sets AllowProvisioning field to given value.

### HasAllowProvisioning

`func (o *UpdateRepositoryParameters) HasAllowProvisioning() bool`

HasAllowProvisioning returns a boolean if a field has been set.

### GetIsStaging

`func (o *UpdateRepositoryParameters) GetIsStaging() bool`

GetIsStaging returns the IsStaging field if non-nil, zero value otherwise.

### GetIsStagingOk

`func (o *UpdateRepositoryParameters) GetIsStagingOk() (*bool, bool)`

GetIsStagingOk returns a tuple with the IsStaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaging

`func (o *UpdateRepositoryParameters) SetIsStaging(v bool)`

SetIsStaging sets IsStaging field to given value.

### HasIsStaging

`func (o *UpdateRepositoryParameters) HasIsStaging() bool`

HasIsStaging returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateRepositoryParameters) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateRepositoryParameters) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateRepositoryParameters) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateRepositoryParameters) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


