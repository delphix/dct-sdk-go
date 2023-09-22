# PasswordVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The password vault ID. | [optional] 
**Name** | Pointer to **string** | The name of this password vault. | [optional] 
**NamespaceId** | Pointer to **string** | The namespace id of this password vault. | [optional] 
**NamespaceName** | Pointer to **string** | The namespace name of this password vault. | [optional] 
**IsReplica** | Pointer to **bool** | Is this a replicated object. | [optional] 
**EngineId** | Pointer to **string** | Id of the Engine that this password vault belongs to. | [optional] 
**EngineName** | Pointer to **string** | Name of the Engine that this password vault belongs to. | [optional] 
**Type** | Pointer to **string** | The type of this password vault. | [optional] 
**Host** | Pointer to **string** | Host name or IP address of this password vault server. | [optional] 
**Port** | Pointer to **int64** | Port of this password vault server. | [optional] 
**Enabled** | Pointer to **NullableBool** | The vault is enabled or not. | [optional] 
**Namespace** | Pointer to **string** | The namespace of this password vault. | [optional] 

## Methods

### NewPasswordVault

`func NewPasswordVault() *PasswordVault`

NewPasswordVault instantiates a new PasswordVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordVaultWithDefaults

`func NewPasswordVaultWithDefaults() *PasswordVault`

NewPasswordVaultWithDefaults instantiates a new PasswordVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PasswordVault) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordVault) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordVault) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PasswordVault) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PasswordVault) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordVault) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordVault) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PasswordVault) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceId

`func (o *PasswordVault) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *PasswordVault) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *PasswordVault) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *PasswordVault) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetNamespaceName

`func (o *PasswordVault) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *PasswordVault) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *PasswordVault) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.

### HasNamespaceName

`func (o *PasswordVault) HasNamespaceName() bool`

HasNamespaceName returns a boolean if a field has been set.

### GetIsReplica

`func (o *PasswordVault) GetIsReplica() bool`

GetIsReplica returns the IsReplica field if non-nil, zero value otherwise.

### GetIsReplicaOk

`func (o *PasswordVault) GetIsReplicaOk() (*bool, bool)`

GetIsReplicaOk returns a tuple with the IsReplica field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReplica

`func (o *PasswordVault) SetIsReplica(v bool)`

SetIsReplica sets IsReplica field to given value.

### HasIsReplica

`func (o *PasswordVault) HasIsReplica() bool`

HasIsReplica returns a boolean if a field has been set.

### GetEngineId

`func (o *PasswordVault) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *PasswordVault) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *PasswordVault) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *PasswordVault) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetEngineName

`func (o *PasswordVault) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *PasswordVault) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *PasswordVault) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *PasswordVault) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetType

`func (o *PasswordVault) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PasswordVault) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PasswordVault) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PasswordVault) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHost

`func (o *PasswordVault) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PasswordVault) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PasswordVault) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PasswordVault) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PasswordVault) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PasswordVault) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PasswordVault) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *PasswordVault) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEnabled

`func (o *PasswordVault) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PasswordVault) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PasswordVault) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PasswordVault) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *PasswordVault) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *PasswordVault) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetNamespace

`func (o *PasswordVault) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PasswordVault) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PasswordVault) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PasswordVault) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


