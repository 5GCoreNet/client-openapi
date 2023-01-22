# MbsContextInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**AnyUeInd** | Pointer to **bool** |  | [optional] [default to false]
**LlSsm** | Pointer to [**Ssm**](Ssm.md) |  | [optional] 
**CTeid** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**MbsServiceArea** | Pointer to [**MbsServiceArea**](MbsServiceArea.md) |  | [optional] 
**MbsServiceAreaInfoList** | Pointer to [**map[string]MbsServiceAreaInfo**](MbsServiceAreaInfo.md) | A map (list of key-value pairs) where the key identifies an areaSessionId  | [optional] 

## Methods

### NewMbsContextInfo

`func NewMbsContextInfo() *MbsContextInfo`

NewMbsContextInfo instantiates a new MbsContextInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsContextInfoWithDefaults

`func NewMbsContextInfoWithDefaults() *MbsContextInfo`

NewMbsContextInfoWithDefaults instantiates a new MbsContextInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *MbsContextInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MbsContextInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MbsContextInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MbsContextInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *MbsContextInfo) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *MbsContextInfo) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *MbsContextInfo) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *MbsContextInfo) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetLlSsm

`func (o *MbsContextInfo) GetLlSsm() Ssm`

GetLlSsm returns the LlSsm field if non-nil, zero value otherwise.

### GetLlSsmOk

`func (o *MbsContextInfo) GetLlSsmOk() (*Ssm, bool)`

GetLlSsmOk returns a tuple with the LlSsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlSsm

`func (o *MbsContextInfo) SetLlSsm(v Ssm)`

SetLlSsm sets LlSsm field to given value.

### HasLlSsm

`func (o *MbsContextInfo) HasLlSsm() bool`

HasLlSsm returns a boolean if a field has been set.

### GetCTeid

`func (o *MbsContextInfo) GetCTeid() int32`

GetCTeid returns the CTeid field if non-nil, zero value otherwise.

### GetCTeidOk

`func (o *MbsContextInfo) GetCTeidOk() (*int32, bool)`

GetCTeidOk returns a tuple with the CTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTeid

`func (o *MbsContextInfo) SetCTeid(v int32)`

SetCTeid sets CTeid field to given value.

### HasCTeid

`func (o *MbsContextInfo) HasCTeid() bool`

HasCTeid returns a boolean if a field has been set.

### GetMbsServiceArea

`func (o *MbsContextInfo) GetMbsServiceArea() MbsServiceArea`

GetMbsServiceArea returns the MbsServiceArea field if non-nil, zero value otherwise.

### GetMbsServiceAreaOk

`func (o *MbsContextInfo) GetMbsServiceAreaOk() (*MbsServiceArea, bool)`

GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceArea

`func (o *MbsContextInfo) SetMbsServiceArea(v MbsServiceArea)`

SetMbsServiceArea sets MbsServiceArea field to given value.

### HasMbsServiceArea

`func (o *MbsContextInfo) HasMbsServiceArea() bool`

HasMbsServiceArea returns a boolean if a field has been set.

### GetMbsServiceAreaInfoList

`func (o *MbsContextInfo) GetMbsServiceAreaInfoList() map[string]MbsServiceAreaInfo`

GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field if non-nil, zero value otherwise.

### GetMbsServiceAreaInfoListOk

`func (o *MbsContextInfo) GetMbsServiceAreaInfoListOk() (*map[string]MbsServiceAreaInfo, bool)`

GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceAreaInfoList

`func (o *MbsContextInfo) SetMbsServiceAreaInfoList(v map[string]MbsServiceAreaInfo)`

SetMbsServiceAreaInfoList sets MbsServiceAreaInfoList field to given value.

### HasMbsServiceAreaInfoList

`func (o *MbsContextInfo) HasMbsServiceAreaInfoList() bool`

HasMbsServiceAreaInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


