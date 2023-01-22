# PushInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | **string** | Public Identity of the UE | 
**UeIdType** | [**UeIdType**](UeIdType.md) |  | 
**UiccAppLabel** | **string** | Character string representing an UICC Application Label | 
**NafId** | [**NafId**](NafId.md) |  | 
**PtId** | **string** | Character string representing a P-TID | 
**UiccOrMe** | [**UiccOrMe**](UiccOrMe.md) |  | 
**RequestedLifeTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**PrivateIdRequest** | Pointer to **bool** |  | [optional] 
**GbaUAware** | Pointer to **bool** |  | [optional] 
**GsIds** | Pointer to **[]int32** |  | [optional] 
**Auts** | Pointer to **string** | AUTS value in UMTS AKA | [optional] 
**Rand** | Pointer to **string** | RAND in UMTS AKA | [optional] 
**SecurityFeaturesRequest** | Pointer to [**[]SecFeature**](SecFeature.md) |  | [optional] 

## Methods

### NewPushInfoRequest

`func NewPushInfoRequest(ueId string, ueIdType UeIdType, uiccAppLabel string, nafId NafId, ptId string, uiccOrMe UiccOrMe, requestedLifeTime time.Time, ) *PushInfoRequest`

NewPushInfoRequest instantiates a new PushInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushInfoRequestWithDefaults

`func NewPushInfoRequestWithDefaults() *PushInfoRequest`

NewPushInfoRequestWithDefaults instantiates a new PushInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *PushInfoRequest) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *PushInfoRequest) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *PushInfoRequest) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetUeIdType

`func (o *PushInfoRequest) GetUeIdType() UeIdType`

GetUeIdType returns the UeIdType field if non-nil, zero value otherwise.

### GetUeIdTypeOk

`func (o *PushInfoRequest) GetUeIdTypeOk() (*UeIdType, bool)`

GetUeIdTypeOk returns a tuple with the UeIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIdType

`func (o *PushInfoRequest) SetUeIdType(v UeIdType)`

SetUeIdType sets UeIdType field to given value.


### GetUiccAppLabel

`func (o *PushInfoRequest) GetUiccAppLabel() string`

GetUiccAppLabel returns the UiccAppLabel field if non-nil, zero value otherwise.

### GetUiccAppLabelOk

`func (o *PushInfoRequest) GetUiccAppLabelOk() (*string, bool)`

GetUiccAppLabelOk returns a tuple with the UiccAppLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiccAppLabel

`func (o *PushInfoRequest) SetUiccAppLabel(v string)`

SetUiccAppLabel sets UiccAppLabel field to given value.


### GetNafId

`func (o *PushInfoRequest) GetNafId() NafId`

GetNafId returns the NafId field if non-nil, zero value otherwise.

### GetNafIdOk

`func (o *PushInfoRequest) GetNafIdOk() (*NafId, bool)`

GetNafIdOk returns a tuple with the NafId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNafId

`func (o *PushInfoRequest) SetNafId(v NafId)`

SetNafId sets NafId field to given value.


### GetPtId

`func (o *PushInfoRequest) GetPtId() string`

GetPtId returns the PtId field if non-nil, zero value otherwise.

### GetPtIdOk

`func (o *PushInfoRequest) GetPtIdOk() (*string, bool)`

GetPtIdOk returns a tuple with the PtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtId

`func (o *PushInfoRequest) SetPtId(v string)`

SetPtId sets PtId field to given value.


### GetUiccOrMe

`func (o *PushInfoRequest) GetUiccOrMe() UiccOrMe`

GetUiccOrMe returns the UiccOrMe field if non-nil, zero value otherwise.

### GetUiccOrMeOk

`func (o *PushInfoRequest) GetUiccOrMeOk() (*UiccOrMe, bool)`

GetUiccOrMeOk returns a tuple with the UiccOrMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiccOrMe

`func (o *PushInfoRequest) SetUiccOrMe(v UiccOrMe)`

SetUiccOrMe sets UiccOrMe field to given value.


### GetRequestedLifeTime

`func (o *PushInfoRequest) GetRequestedLifeTime() time.Time`

GetRequestedLifeTime returns the RequestedLifeTime field if non-nil, zero value otherwise.

### GetRequestedLifeTimeOk

`func (o *PushInfoRequest) GetRequestedLifeTimeOk() (*time.Time, bool)`

GetRequestedLifeTimeOk returns a tuple with the RequestedLifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedLifeTime

`func (o *PushInfoRequest) SetRequestedLifeTime(v time.Time)`

SetRequestedLifeTime sets RequestedLifeTime field to given value.


### GetPrivateIdRequest

`func (o *PushInfoRequest) GetPrivateIdRequest() bool`

GetPrivateIdRequest returns the PrivateIdRequest field if non-nil, zero value otherwise.

### GetPrivateIdRequestOk

`func (o *PushInfoRequest) GetPrivateIdRequestOk() (*bool, bool)`

GetPrivateIdRequestOk returns a tuple with the PrivateIdRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIdRequest

`func (o *PushInfoRequest) SetPrivateIdRequest(v bool)`

SetPrivateIdRequest sets PrivateIdRequest field to given value.

### HasPrivateIdRequest

`func (o *PushInfoRequest) HasPrivateIdRequest() bool`

HasPrivateIdRequest returns a boolean if a field has been set.

### GetGbaUAware

`func (o *PushInfoRequest) GetGbaUAware() bool`

GetGbaUAware returns the GbaUAware field if non-nil, zero value otherwise.

### GetGbaUAwareOk

`func (o *PushInfoRequest) GetGbaUAwareOk() (*bool, bool)`

GetGbaUAwareOk returns a tuple with the GbaUAware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbaUAware

`func (o *PushInfoRequest) SetGbaUAware(v bool)`

SetGbaUAware sets GbaUAware field to given value.

### HasGbaUAware

`func (o *PushInfoRequest) HasGbaUAware() bool`

HasGbaUAware returns a boolean if a field has been set.

### GetGsIds

`func (o *PushInfoRequest) GetGsIds() []int32`

GetGsIds returns the GsIds field if non-nil, zero value otherwise.

### GetGsIdsOk

`func (o *PushInfoRequest) GetGsIdsOk() (*[]int32, bool)`

GetGsIdsOk returns a tuple with the GsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGsIds

`func (o *PushInfoRequest) SetGsIds(v []int32)`

SetGsIds sets GsIds field to given value.

### HasGsIds

`func (o *PushInfoRequest) HasGsIds() bool`

HasGsIds returns a boolean if a field has been set.

### GetAuts

`func (o *PushInfoRequest) GetAuts() string`

GetAuts returns the Auts field if non-nil, zero value otherwise.

### GetAutsOk

`func (o *PushInfoRequest) GetAutsOk() (*string, bool)`

GetAutsOk returns a tuple with the Auts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuts

`func (o *PushInfoRequest) SetAuts(v string)`

SetAuts sets Auts field to given value.

### HasAuts

`func (o *PushInfoRequest) HasAuts() bool`

HasAuts returns a boolean if a field has been set.

### GetRand

`func (o *PushInfoRequest) GetRand() string`

GetRand returns the Rand field if non-nil, zero value otherwise.

### GetRandOk

`func (o *PushInfoRequest) GetRandOk() (*string, bool)`

GetRandOk returns a tuple with the Rand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRand

`func (o *PushInfoRequest) SetRand(v string)`

SetRand sets Rand field to given value.

### HasRand

`func (o *PushInfoRequest) HasRand() bool`

HasRand returns a boolean if a field has been set.

### GetSecurityFeaturesRequest

`func (o *PushInfoRequest) GetSecurityFeaturesRequest() []SecFeature`

GetSecurityFeaturesRequest returns the SecurityFeaturesRequest field if non-nil, zero value otherwise.

### GetSecurityFeaturesRequestOk

`func (o *PushInfoRequest) GetSecurityFeaturesRequestOk() (*[]SecFeature, bool)`

GetSecurityFeaturesRequestOk returns a tuple with the SecurityFeaturesRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityFeaturesRequest

`func (o *PushInfoRequest) SetSecurityFeaturesRequest(v []SecFeature)`

SetSecurityFeaturesRequest sets SecurityFeaturesRequest field to given value.

### HasSecurityFeaturesRequest

`func (o *PushInfoRequest) HasSecurityFeaturesRequest() bool`

HasSecurityFeaturesRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


