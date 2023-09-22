# VirtualizationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 
**EngineId** | **string** |  | 
**PolicyType** | **string** |  | 
**TimezoneId** | **string** |  | 
**Schedules** | [**[]VirtualizationSchedule**](VirtualizationSchedule.md) |  | 

## Methods

### NewVirtualizationPolicy

`func NewVirtualizationPolicy(id string, name string, engineId string, policyType string, timezoneId string, schedules []VirtualizationSchedule, ) *VirtualizationPolicy`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


