# KerberosConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The kerberos config ID. | [optional] 
**Name** | Pointer to **string** | The name of the kerberos config object. | [optional] 
**NamespaceId** | Pointer to **string** | The namespace id of this kerberos config object. | [optional] 
**NamespaceName** | Pointer to **string** | The namespace name of this kerberos config object. | [optional] 
**IsReplica** | Pointer to **bool** | Is this a replicated object. | [optional] 
**EngineId** | Pointer to **string** | Id of the Engine that this kerberos config object belongs to. | [optional] 
**EngineName** | Pointer to **string** | Name of the Engine that this kerberos config object belongs to. | [optional] 
**Realm** | Pointer to **string** | Kerberos Realm name. | [optional] 
**Principal** | Pointer to **string** | Kerberos principal name. | [optional] 
**Enabled** | Pointer to **NullableBool** | The kerberos is enabled or not. | [optional] 
**Keytab** | Pointer to **string** | Kerberos keytab. | [optional] 
**KdcServers** | Pointer to [**[]KDCServers**](KDCServers.md) | One of more KDC servers. | [optional] 

## Methods

### NewKerberosConfig

`func NewKerberosConfig() *KerberosConfig`

NewKerberosConfig instantiates a new KerberosConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosConfigWithDefaults

`func NewKerberosConfigWithDefaults() *KerberosConfig`

NewKerberosConfigWithDefaults instantiates a new KerberosConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KerberosConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KerberosConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KerberosConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KerberosConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KerberosConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KerberosConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KerberosConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KerberosConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *KerberosConfig) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *KerberosConfig) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *KerberosConfig) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *KerberosConfig) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNamespaceName

`func (o *KerberosConfig) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *KerberosConfig) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *KerberosConfig) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *KerberosConfig) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetIsReplica

`func (o *KerberosConfig) GetIsReplica() bool`

GetIsReplica returns the IsReplica field if non-nil, zero value otherwise.

### GetIsReplicaOk

`func (o *KerberosConfig) GetIsReplicaOk() (*bool, bool)`

GetIsReplicaOk returns a tuple with the IsReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplica

`func (o *KerberosConfig) SetIsReplica(v bool)`

SetIsReplica sets IsReplica field to given value.

### HasIsReplica

`func (o *KerberosConfig) HasIsReplica() bool`

HasIsReplica returns a boolean if a field has been set.

### GetEngineId

`func (o *KerberosConfig) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *KerberosConfig) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *KerberosConfig) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *KerberosConfig) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineName

`func (o *KerberosConfig) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *KerberosConfig) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *KerberosConfig) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *KerberosConfig) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetRealm

`func (o *KerberosConfig) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *KerberosConfig) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *KerberosConfig) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *KerberosConfig) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetPrincipal

`func (o *KerberosConfig) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *KerberosConfig) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *KerberosConfig) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *KerberosConfig) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetEnabled

`func (o *KerberosConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KerberosConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KerberosConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KerberosConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *KerberosConfig) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *KerberosConfig) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetKeytab

`func (o *KerberosConfig) GetKeytab() string`

GetKeytab returns the Keytab field if non-nil, zero value otherwise.

### GetKeytabOk

`func (o *KerberosConfig) GetKeytabOk() (*string, bool)`

GetKeytabOk returns a tuple with the Keytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytab

`func (o *KerberosConfig) SetKeytab(v string)`

SetKeytab sets Keytab field to given value.

### HasKeytab

`func (o *KerberosConfig) HasKeytab() bool`

HasKeytab returns a boolean if a field has been set.

### GetKdcServers

`func (o *KerberosConfig) GetKdcServers() []KDCServers`

GetKdcServers returns the KdcServers field if non-nil, zero value otherwise.

### GetKdcServersOk

`func (o *KerberosConfig) GetKdcServersOk() (*[]KDCServers, bool)`

GetKdcServersOk returns a tuple with the KdcServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcServers

`func (o *KerberosConfig) SetKdcServers(v []KDCServers)`

SetKdcServers sets KdcServers field to given value.

### HasKdcServers

`func (o *KerberosConfig) HasKdcServers() bool`

HasKdcServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


