# VirtualizationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NamespaceId** | Pointer to **string** | The namespace id of this virtualization policy. | [optional] 
**NamespaceName** | Pointer to **string** | The namespace name of this virtualization policy.. | [optional] 
**IsReplica** | Pointer to **bool** | Is this a replicated object. | [optional] 
**EngineId** | Pointer to **string** |  | [optional] 
**PolicyType** | Pointer to **string** |  | [optional] 
**TimezoneId** | Pointer to **string** |  | [optional] 
**Schedules** | Pointer to [**[]VirtualizationSchedule**](VirtualizationSchedule.md) |  | [optional] 

## Methods

### NewVirtualizationPolicy

`func NewVirtualizationPolicy() *VirtualizationPolicy`

NewVirtualizationPolicy instantiates a new VirtualizationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationPolicyWithDefaults

`func NewVirtualizationPolicyWithDefaults() *VirtualizationPolicy`

NewVirtualizationPolicyWithDefaults instantiates a new VirtualizationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualizationPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualizationPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualizationPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualizationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *VirtualizationPolicy) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VirtualizationPolicy) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VirtualizationPolicy) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *VirtualizationPolicy) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNamespaceId

`func (o *VirtualizationPolicy) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *VirtualizationPolicy) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *VirtualizationPolicy) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *VirtualizationPolicy) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNamespaceName

`func (o *VirtualizationPolicy) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *VirtualizationPolicy) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *VirtualizationPolicy) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *VirtualizationPolicy) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetIsReplica

`func (o *VirtualizationPolicy) GetIsReplica() bool`

GetIsReplica returns the IsReplica field if non-nil, zero value otherwise.

### GetIsReplicaOk

`func (o *VirtualizationPolicy) GetIsReplicaOk() (*bool, bool)`

GetIsReplicaOk returns a tuple with the IsReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplica

`func (o *VirtualizationPolicy) SetIsReplica(v bool)`

SetIsReplica sets IsReplica field to given value.

### HasIsReplica

`func (o *VirtualizationPolicy) HasIsReplica() bool`

HasIsReplica returns a boolean if a field has been set.

### GetEngineId

`func (o *VirtualizationPolicy) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *VirtualizationPolicy) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *VirtualizationPolicy) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *VirtualizationPolicy) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetPolicyType

`func (o *VirtualizationPolicy) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *VirtualizationPolicy) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *VirtualizationPolicy) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *VirtualizationPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetTimezoneId

`func (o *VirtualizationPolicy) GetTimezoneId() string`

GetTimezoneId returns the TimezoneId field if non-nil, zero value otherwise.

### GetTimezoneIdOk

`func (o *VirtualizationPolicy) GetTimezoneIdOk() (*string, bool)`

GetTimezoneIdOk returns a tuple with the TimezoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneId

`func (o *VirtualizationPolicy) SetTimezoneId(v string)`

SetTimezoneId sets TimezoneId field to given value.

### HasTimezoneId

`func (o *VirtualizationPolicy) HasTimezoneId() bool`

HasTimezoneId returns a boolean if a field has been set.

### GetSchedules

`func (o *VirtualizationPolicy) GetSchedules() []VirtualizationSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *VirtualizationPolicy) GetSchedulesOk() (*[]VirtualizationSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *VirtualizationPolicy) SetSchedules(v []VirtualizationSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *VirtualizationPolicy) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


