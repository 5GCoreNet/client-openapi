# RegistrationInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiProvPubKey** | **string** | Public Key of API Provider domain function. | 
**ApiProvCert** | Pointer to **string** | API provider domain function&#39;s client certificate | [optional] 

## Methods

### NewRegistrationInformation

`func NewRegistrationInformation(apiProvPubKey string, ) *RegistrationInformation`

NewRegistrationInformation instantiates a new RegistrationInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationInformationWithDefaults

`func NewRegistrationInformationWithDefaults() *RegistrationInformation`

NewRegistrationInformationWithDefaults instantiates a new RegistrationInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiProvPubKey

`func (o *RegistrationInformation) GetApiProvPubKey() string`

GetApiProvPubKey returns the ApiProvPubKey field if non-nil, zero value otherwise.

### GetApiProvPubKeyOk

`func (o *RegistrationInformation) GetApiProvPubKeyOk() (*string, bool)`

GetApiProvPubKeyOk returns a tuple with the ApiProvPubKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvPubKey

`func (o *RegistrationInformation) SetApiProvPubKey(v string)`

SetApiProvPubKey sets ApiProvPubKey field to given value.


### GetApiProvCert

`func (o *RegistrationInformation) GetApiProvCert() string`

GetApiProvCert returns the ApiProvCert field if non-nil, zero value otherwise.

### GetApiProvCertOk

`func (o *RegistrationInformation) GetApiProvCertOk() (*string, bool)`

GetApiProvCertOk returns a tuple with the ApiProvCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProvCert

`func (o *RegistrationInformation) SetApiProvCert(v string)`

SetApiProvCert sets ApiProvCert field to given value.

### HasApiProvCert

`func (o *RegistrationInformation) HasApiProvCert() bool`

HasApiProvCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


