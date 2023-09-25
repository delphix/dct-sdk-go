# Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Connector entity ID. | [optional] 
**Name** | Pointer to **string** | The Connector name. | [optional] 
**EngineId** | Pointer to **string** | A reference to the Engine that this Connector belongs to. | [optional] 
**Type** | Pointer to **string** | The type of Connector. One of Database, File, or Mainframe. | [optional] 
**Hostname** | Pointer to **string** | The network hostname or IP address of the database server. | [optional] 
**Port** | Pointer to **int32** | The TCP port of the server. | [optional] 
**Username** | Pointer to **string** | The username this Connector will use to connect to the database. | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewConnector

`func NewConnector() *Connector`

NewConnector instantiates a new Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWithDefaults

`func NewConnectorWithDefaults() *Connector`

NewConnectorWithDefaults instantiates a new Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Connector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Connector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Connector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Connector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEngineId

`func (o *Connector) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *Connector) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *Connector) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *Connector) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetType

`func (o *Connector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Connector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Connector) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Connector) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHostname

`func (o *Connector) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Connector) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Connector) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Connector) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPort

`func (o *Connector) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Connector) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Connector) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Connector) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *Connector) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Connector) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Connector) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Connector) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetTags

`func (o *Connector) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Connector) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Connector) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Connector) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


