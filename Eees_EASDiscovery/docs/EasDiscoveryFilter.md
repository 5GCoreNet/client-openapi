# EasDiscoveryFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcChars** | Pointer to [**[]ACCharacteristics**](ACCharacteristics.md) | AC description for which an EAS is needed. | [optional] 
**EasChars** | Pointer to [**[]EasCharacteristics**](EasCharacteristics.md) | Required EAS chararcteristics. | [optional] 

## Methods

### NewEasDiscoveryFilter

`func NewEasDiscoveryFilter() *EasDiscoveryFilter`

NewEasDiscoveryFilter instantiates a new EasDiscoveryFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDiscoveryFilterWithDefaults

`func NewEasDiscoveryFilterWithDefaults() *EasDiscoveryFilter`

NewEasDiscoveryFilterWithDefaults instantiates a new EasDiscoveryFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcChars

`func (o *EasDiscoveryFilter) GetAcChars() []ACCharacteristics`

GetAcChars returns the AcChars field if non-nil, zero value otherwise.

### GetAcCharsOk

`func (o *EasDiscoveryFilter) GetAcCharsOk() (*[]ACCharacteristics, bool)`

GetAcCharsOk returns a tuple with the AcChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcChars

`func (o *EasDiscoveryFilter) SetAcChars(v []ACCharacteristics)`

SetAcChars sets AcChars field to given value.

### HasAcChars

`func (o *EasDiscoveryFilter) HasAcChars() bool`

HasAcChars returns a boolean if a field has been set.

### GetEasChars

`func (o *EasDiscoveryFilter) GetEasChars() []EasCharacteristics`

GetEasChars returns the EasChars field if non-nil, zero value otherwise.

### GetEasCharsOk

`func (o *EasDiscoveryFilter) GetEasCharsOk() (*[]EasCharacteristics, bool)`

GetEasCharsOk returns a tuple with the EasChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasChars

`func (o *EasDiscoveryFilter) SetEasChars(v []EasCharacteristics)`

SetEasChars sets EasChars field to given value.

### HasEasChars

`func (o *EasDiscoveryFilter) HasEasChars() bool`

HasEasChars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


