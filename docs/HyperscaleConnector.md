# HyperscaleConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the Hyperscale Connector. | [optional] 
**HyperscaleInstanceId** | Pointer to **string** | The ID of the Hyperscale instance of this Dataset. | [optional] 
**SourceUsername** | Pointer to **string** | The username this Connector will use to connect to the source database. | [optional] 
**SourcePassword** | Pointer to **string** | The password this Connector will use to connect to the source database. | [optional] 
**SourceJdbcUrl** | Pointer to **string** | The JDBC URL used to connect to the source database. | [optional] 
**SourceConnectionProperties** | Pointer to **map[string]string** |  | [optional] 
**TargetUsername** | Pointer to **string** | The username this Connector will use to connect to the target database. | [optional] 
**TargetPassword** | Pointer to **string** | The username this Connector will use to connect to the target database. | [optional] 
**TargetJdbcUrl** | Pointer to **string** | The JDBC URL used to connect to the target database. | [optional] 
**TargetConnectionProperties** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewHyperscaleConnector

`func NewHyperscaleConnector() *HyperscaleConnector`

NewHyperscaleConnector instantiates a new HyperscaleConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperscaleConnectorWithDefaults

`func NewHyperscaleConnectorWithDefaults() *HyperscaleConnector`

NewHyperscaleConnectorWithDefaults instantiates a new HyperscaleConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HyperscaleConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HyperscaleConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HyperscaleConnector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HyperscaleConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHyperscaleInstanceId

`func (o *HyperscaleConnector) GetHyperscaleInstanceId() string`

GetHyperscaleInstanceId returns the HyperscaleInstanceId field if non-nil, zero value otherwise.

### GetHyperscaleInstanceIdOk

`func (o *HyperscaleConnector) GetHyperscaleInstanceIdOk() (*string, bool)`

GetHyperscaleInstanceIdOk returns a tuple with the HyperscaleInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperscaleInstanceId

`func (o *HyperscaleConnector) SetHyperscaleInstanceId(v string)`

SetHyperscaleInstanceId sets HyperscaleInstanceId field to given value.

### HasHyperscaleInstanceId

`func (o *HyperscaleConnector) HasHyperscaleInstanceId() bool`

HasHyperscaleInstanceId returns a boolean if a field has been set.

### GetSourceUsername

`func (o *HyperscaleConnector) GetSourceUsername() string`

GetSourceUsername returns the SourceUsername field if non-nil, zero value otherwise.

### GetSourceUsernameOk

`func (o *HyperscaleConnector) GetSourceUsernameOk() (*string, bool)`

GetSourceUsernameOk returns a tuple with the SourceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUsername

`func (o *HyperscaleConnector) SetSourceUsername(v string)`

SetSourceUsername sets SourceUsername field to given value.

### HasSourceUsername

`func (o *HyperscaleConnector) HasSourceUsername() bool`

HasSourceUsername returns a boolean if a field has been set.

### GetSourcePassword

`func (o *HyperscaleConnector) GetSourcePassword() string`

GetSourcePassword returns the SourcePassword field if non-nil, zero value otherwise.

### GetSourcePasswordOk

`func (o *HyperscaleConnector) GetSourcePasswordOk() (*string, bool)`

GetSourcePasswordOk returns a tuple with the SourcePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePassword

`func (o *HyperscaleConnector) SetSourcePassword(v string)`

SetSourcePassword sets SourcePassword field to given value.

### HasSourcePassword

`func (o *HyperscaleConnector) HasSourcePassword() bool`

HasSourcePassword returns a boolean if a field has been set.

### GetSourceJdbcUrl

`func (o *HyperscaleConnector) GetSourceJdbcUrl() string`

GetSourceJdbcUrl returns the SourceJdbcUrl field if non-nil, zero value otherwise.

### GetSourceJdbcUrlOk

`func (o *HyperscaleConnector) GetSourceJdbcUrlOk() (*string, bool)`

GetSourceJdbcUrlOk returns a tuple with the SourceJdbcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceJdbcUrl

`func (o *HyperscaleConnector) SetSourceJdbcUrl(v string)`

SetSourceJdbcUrl sets SourceJdbcUrl field to given value.

### HasSourceJdbcUrl

`func (o *HyperscaleConnector) HasSourceJdbcUrl() bool`

HasSourceJdbcUrl returns a boolean if a field has been set.

### GetSourceConnectionProperties

`func (o *HyperscaleConnector) GetSourceConnectionProperties() map[string]string`

GetSourceConnectionProperties returns the SourceConnectionProperties field if non-nil, zero value otherwise.

### GetSourceConnectionPropertiesOk

`func (o *HyperscaleConnector) GetSourceConnectionPropertiesOk() (*map[string]string, bool)`

GetSourceConnectionPropertiesOk returns a tuple with the SourceConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnectionProperties

`func (o *HyperscaleConnector) SetSourceConnectionProperties(v map[string]string)`

SetSourceConnectionProperties sets SourceConnectionProperties field to given value.

### HasSourceConnectionProperties

`func (o *HyperscaleConnector) HasSourceConnectionProperties() bool`

HasSourceConnectionProperties returns a boolean if a field has been set.

### GetTargetUsername

`func (o *HyperscaleConnector) GetTargetUsername() string`

GetTargetUsername returns the TargetUsername field if non-nil, zero value otherwise.

### GetTargetUsernameOk

`func (o *HyperscaleConnector) GetTargetUsernameOk() (*string, bool)`

GetTargetUsernameOk returns a tuple with the TargetUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUsername

`func (o *HyperscaleConnector) SetTargetUsername(v string)`

SetTargetUsername sets TargetUsername field to given value.

### HasTargetUsername

`func (o *HyperscaleConnector) HasTargetUsername() bool`

HasTargetUsername returns a boolean if a field has been set.

### GetTargetPassword

`func (o *HyperscaleConnector) GetTargetPassword() string`

GetTargetPassword returns the TargetPassword field if non-nil, zero value otherwise.

### GetTargetPasswordOk

`func (o *HyperscaleConnector) GetTargetPasswordOk() (*string, bool)`

GetTargetPasswordOk returns a tuple with the TargetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPassword

`func (o *HyperscaleConnector) SetTargetPassword(v string)`

SetTargetPassword sets TargetPassword field to given value.

### HasTargetPassword

`func (o *HyperscaleConnector) HasTargetPassword() bool`

HasTargetPassword returns a boolean if a field has been set.

### GetTargetJdbcUrl

`func (o *HyperscaleConnector) GetTargetJdbcUrl() string`

GetTargetJdbcUrl returns the TargetJdbcUrl field if non-nil, zero value otherwise.

### GetTargetJdbcUrlOk

`func (o *HyperscaleConnector) GetTargetJdbcUrlOk() (*string, bool)`

GetTargetJdbcUrlOk returns a tuple with the TargetJdbcUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetJdbcUrl

`func (o *HyperscaleConnector) SetTargetJdbcUrl(v string)`

SetTargetJdbcUrl sets TargetJdbcUrl field to given value.

### HasTargetJdbcUrl

`func (o *HyperscaleConnector) HasTargetJdbcUrl() bool`

HasTargetJdbcUrl returns a boolean if a field has been set.

### GetTargetConnectionProperties

`func (o *HyperscaleConnector) GetTargetConnectionProperties() map[string]string`

GetTargetConnectionProperties returns the TargetConnectionProperties field if non-nil, zero value otherwise.

### GetTargetConnectionPropertiesOk

`func (o *HyperscaleConnector) GetTargetConnectionPropertiesOk() (*map[string]string, bool)`

GetTargetConnectionPropertiesOk returns a tuple with the TargetConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetConnectionProperties

`func (o *HyperscaleConnector) SetTargetConnectionProperties(v map[string]string)`

SetTargetConnectionProperties sets TargetConnectionProperties field to given value.

### HasTargetConnectionProperties

`func (o *HyperscaleConnector) HasTargetConnectionProperties() bool`

HasTargetConnectionProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


