# SSHVerificationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the verification strategy. | 
**KeyType** | Pointer to **string** | The type of SSH key. | [optional] 
**RawKey** | Pointer to **string** | Base64-encoded ssh key of the host for RAW_KEY verification. | [optional] 
**FingerprintType** | Pointer to **string** | Hash function for the fingerprint for FINGERPRINT verification. | [optional] 
**Fingerprint** | Pointer to **string** | Base-64 encoded fingerprint of the ssh key of the host for FINGERPRINT verification. | [optional] 

## Methods

### NewSSHVerificationStrategy

`func NewSSHVerificationStrategy(name string, ) *SSHVerificationStrategy`

NewSSHVerificationStrategy instantiates a new SSHVerificationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHVerificationStrategyWithDefaults

`func NewSSHVerificationStrategyWithDefaults() *SSHVerificationStrategy`

NewSSHVerificationStrategyWithDefaults instantiates a new SSHVerificationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SSHVerificationStrategy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SSHVerificationStrategy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SSHVerificationStrategy) SetName(v string)`

SetName sets Name field to given value.


### GetKeyType

`func (o *SSHVerificationStrategy) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *SSHVerificationStrategy) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *SSHVerificationStrategy) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *SSHVerificationStrategy) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetRawKey

`func (o *SSHVerificationStrategy) GetRawKey() string`

GetRawKey returns the RawKey field if non-nil, zero value otherwise.

### GetRawKeyOk

`func (o *SSHVerificationStrategy) GetRawKeyOk() (*string, bool)`

GetRawKeyOk returns a tuple with the RawKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawKey

`func (o *SSHVerificationStrategy) SetRawKey(v string)`

SetRawKey sets RawKey field to given value.

### HasRawKey

`func (o *SSHVerificationStrategy) HasRawKey() bool`

HasRawKey returns a boolean if a field has been set.

### GetFingerprintType

`func (o *SSHVerificationStrategy) GetFingerprintType() string`

GetFingerprintType returns the FingerprintType field if non-nil, zero value otherwise.

### GetFingerprintTypeOk

`func (o *SSHVerificationStrategy) GetFingerprintTypeOk() (*string, bool)`

GetFingerprintTypeOk returns a tuple with the FingerprintType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintType

`func (o *SSHVerificationStrategy) SetFingerprintType(v string)`

SetFingerprintType sets FingerprintType field to given value.

### HasFingerprintType

`func (o *SSHVerificationStrategy) HasFingerprintType() bool`

HasFingerprintType returns a boolean if a field has been set.

### GetFingerprint

`func (o *SSHVerificationStrategy) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SSHVerificationStrategy) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SSHVerificationStrategy) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SSHVerificationStrategy) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


