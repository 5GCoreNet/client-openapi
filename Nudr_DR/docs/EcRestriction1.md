# EcRestriction1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfInstanceId** | **string** |  | 
**ReferenceId** | **int32** |  | 
**PlmnEcInfos** | Pointer to [**[]PlmnEcInfo1**](PlmnEcInfo1.md) |  | [optional] 
**MtcProviderInformation** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 

## Methods

### NewEcRestriction1

`func NewEcRestriction1(afInstanceId string, referenceId int32, ) *EcRestriction1`

NewEcRestriction1 instantiates a new EcRestriction1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcRestriction1WithDefaults

`func NewEcRestriction1WithDefaults() *EcRestriction1`

NewEcRestriction1WithDefaults instantiates a new EcRestriction1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfInstanceId

`func (o *EcRestriction1) GetAfInstanceId() string`

GetAfInstanceId returns the AfInstanceId field if non-nil, zero value otherwise.

### GetAfInstanceIdOk

`func (o *EcRestriction1) GetAfInstanceIdOk() (*string, bool)`

GetAfInstanceIdOk returns a tuple with the AfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfInstanceId

`func (o *EcRestriction1) SetAfInstanceId(v string)`

SetAfInstanceId sets AfInstanceId field to given value.


### GetReferenceId

`func (o *EcRestriction1) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *EcRestriction1) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *EcRestriction1) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### GetPlmnEcInfos

`func (o *EcRestriction1) GetPlmnEcInfos() []PlmnEcInfo1`

GetPlmnEcInfos returns the PlmnEcInfos field if non-nil, zero value otherwise.

### GetPlmnEcInfosOk

`func (o *EcRestriction1) GetPlmnEcInfosOk() (*[]PlmnEcInfo1, bool)`

GetPlmnEcInfosOk returns a tuple with the PlmnEcInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnEcInfos

`func (o *EcRestriction1) SetPlmnEcInfos(v []PlmnEcInfo1)`

SetPlmnEcInfos sets PlmnEcInfos field to given value.

### HasPlmnEcInfos

`func (o *EcRestriction1) HasPlmnEcInfos() bool`

HasPlmnEcInfos returns a boolean if a field has been set.

### GetMtcProviderInformation

`func (o *EcRestriction1) GetMtcProviderInformation() string`

GetMtcProviderInformation returns the MtcProviderInformation field if non-nil, zero value otherwise.

### GetMtcProviderInformationOk

`func (o *EcRestriction1) GetMtcProviderInformationOk() (*string, bool)`

GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderInformation

`func (o *EcRestriction1) SetMtcProviderInformation(v string)`

SetMtcProviderInformation sets MtcProviderInformation field to given value.

### HasMtcProviderInformation

`func (o *EcRestriction1) HasMtcProviderInformation() bool`

HasMtcProviderInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


