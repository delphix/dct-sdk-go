# SourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Command** | **string** |  | 
**Shell** | Pointer to **string** |  | [optional] [default to "bash"]
**CredentialsEnvVars** | Pointer to [**[]CredentialsEnvVariable**](CredentialsEnvVariable.md) | List of environment variables that will contain credentials for this operation. | [optional] 

## Methods

### NewSourceOperation

`func NewSourceOperation(name string, command string, ) *SourceOperation`

NewSourceOperation instantiates a new SourceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceOperationWithDefaults

`func NewSourceOperationWithDefaults() *SourceOperation`

NewSourceOperationWithDefaults instantiates a new SourceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceOperation) SetName(v string)`

SetName sets Name field to given value.


### GetCommand

`func (o *SourceOperation) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SourceOperation) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SourceOperation) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetShell

`func (o *SourceOperation) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *SourceOperation) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *SourceOperation) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *SourceOperation) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetCredentialsEnvVars

`func (o *SourceOperation) GetCredentialsEnvVars() []CredentialsEnvVariable`

GetCredentialsEnvVars returns the CredentialsEnvVars field if non-nil, zero value otherwise.

### GetCredentialsEnvVarsOk

`func (o *SourceOperation) GetCredentialsEnvVarsOk() (*[]CredentialsEnvVariable, bool)`

GetCredentialsEnvVarsOk returns a tuple with the CredentialsEnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsEnvVars

`func (o *SourceOperation) SetCredentialsEnvVars(v []CredentialsEnvVariable)`

SetCredentialsEnvVars sets CredentialsEnvVars field to given value.

### HasCredentialsEnvVars

`func (o *SourceOperation) HasCredentialsEnvVars() bool`

HasCredentialsEnvVars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


