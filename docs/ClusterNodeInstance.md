# ClusterNodeInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeReference** | **string** | The cluster node id, name or addresses for this provision operation | 
**InstanceNumber** | **int32** | The instance number for this provision operation | 
**InstanceName** | **string** | The instance name for this provision operation | 

## Methods

### NewClusterNodeInstance

`func NewClusterNodeInstance(nodeReference string, instanceNumber int32, instanceName string, ) *ClusterNodeInstance`

NewClusterNodeInstance instantiates a new ClusterNodeInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeInstanceWithDefaults

`func NewClusterNodeInstanceWithDefaults() *ClusterNodeInstance`

NewClusterNodeInstanceWithDefaults instantiates a new ClusterNodeInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeReference

`func (o *ClusterNodeInstance) GetNodeReference() string`

GetNodeReference returns the NodeReference field if non-nil, zero value otherwise.

### GetNodeReferenceOk

`func (o *ClusterNodeInstance) GetNodeReferenceOk() (*string, bool)`

GetNodeReferenceOk returns a tuple with the NodeReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeReference

`func (o *ClusterNodeInstance) SetNodeReference(v string)`

SetNodeReference sets NodeReference field to given value.


### GetInstanceNumber

`func (o *ClusterNodeInstance) GetInstanceNumber() int32`

GetInstanceNumber returns the InstanceNumber field if non-nil, zero value otherwise.

### GetInstanceNumberOk

`func (o *ClusterNodeInstance) GetInstanceNumberOk() (*int32, bool)`

GetInstanceNumberOk returns a tuple with the InstanceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceNumber

`func (o *ClusterNodeInstance) SetInstanceNumber(v int32)`

SetInstanceNumber sets InstanceNumber field to given value.


### GetInstanceName

`func (o *ClusterNodeInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *ClusterNodeInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *ClusterNodeInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


