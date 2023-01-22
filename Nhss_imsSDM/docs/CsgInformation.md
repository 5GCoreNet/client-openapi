# CsgInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CsgId** | **string** |  | 
**AccessMode** | Pointer to **string** |  | [optional] 
**CMi** | Pointer to **bool** |  | [optional] 

## Methods

### NewCsgInformation

`func NewCsgInformation(csgId string, ) *CsgInformation`

NewCsgInformation instantiates a new CsgInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsgInformationWithDefaults

`func NewCsgInformationWithDefaults() *CsgInformation`

NewCsgInformationWithDefaults instantiates a new CsgInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCsgId

`func (o *CsgInformation) GetCsgId() string`

GetCsgId returns the CsgId field if non-nil, zero value otherwise.

### GetCsgIdOk

`func (o *CsgInformation) GetCsgIdOk() (*string, bool)`

GetCsgIdOk returns a tuple with the CsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsgId

`func (o *CsgInformation) SetCsgId(v string)`

SetCsgId sets CsgId field to given value.


### GetAccessMode

`func (o *CsgInformation) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CsgInformation) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CsgInformation) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *CsgInformation) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetCMi

`func (o *CsgInformation) GetCMi() bool`

GetCMi returns the CMi field if non-nil, zero value otherwise.

### GetCMiOk

`func (o *CsgInformation) GetCMiOk() (*bool, bool)`

GetCMiOk returns a tuple with the CMi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCMi

`func (o *CsgInformation) SetCMi(v bool)`

SetCMi sets CMi field to given value.

### HasCMi

`func (o *CsgInformation) HasCMi() bool`

HasCMi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


