# ContextUpdateReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsServiceArea** | Pointer to [**MbsServiceArea**](MbsServiceArea.md) |  | [optional] 
**MbsServiceAreaInfoList** | Pointer to [**[]MbsServiceAreaInfo**](MbsServiceAreaInfo.md) |  | [optional] 
**N2MbsSmInfo** | Pointer to [**N2MbsSmInfo**](N2MbsSmInfo.md) |  | [optional] 
**RanIdList** | Pointer to [**[]GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**NoNgapSignallingInd** | Pointer to **bool** |  | [optional] 
**NotifyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**MaxResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**N2MbsInfoChangeInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewContextUpdateReqData

`func NewContextUpdateReqData() *ContextUpdateReqData`

NewContextUpdateReqData instantiates a new ContextUpdateReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextUpdateReqDataWithDefaults

`func NewContextUpdateReqDataWithDefaults() *ContextUpdateReqData`

NewContextUpdateReqDataWithDefaults instantiates a new ContextUpdateReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsServiceArea

`func (o *ContextUpdateReqData) GetMbsServiceArea() MbsServiceArea`

GetMbsServiceArea returns the MbsServiceArea field if non-nil, zero value otherwise.

### GetMbsServiceAreaOk

`func (o *ContextUpdateReqData) GetMbsServiceAreaOk() (*MbsServiceArea, bool)`

GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceArea

`func (o *ContextUpdateReqData) SetMbsServiceArea(v MbsServiceArea)`

SetMbsServiceArea sets MbsServiceArea field to given value.

### HasMbsServiceArea

`func (o *ContextUpdateReqData) HasMbsServiceArea() bool`

HasMbsServiceArea returns a boolean if a field has been set.

### GetMbsServiceAreaInfoList

`func (o *ContextUpdateReqData) GetMbsServiceAreaInfoList() []MbsServiceAreaInfo`

GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field if non-nil, zero value otherwise.

### GetMbsServiceAreaInfoListOk

`func (o *ContextUpdateReqData) GetMbsServiceAreaInfoListOk() (*[]MbsServiceAreaInfo, bool)`

GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceAreaInfoList

`func (o *ContextUpdateReqData) SetMbsServiceAreaInfoList(v []MbsServiceAreaInfo)`

SetMbsServiceAreaInfoList sets MbsServiceAreaInfoList field to given value.

### HasMbsServiceAreaInfoList

`func (o *ContextUpdateReqData) HasMbsServiceAreaInfoList() bool`

HasMbsServiceAreaInfoList returns a boolean if a field has been set.

### GetN2MbsSmInfo

`func (o *ContextUpdateReqData) GetN2MbsSmInfo() N2MbsSmInfo`

GetN2MbsSmInfo returns the N2MbsSmInfo field if non-nil, zero value otherwise.

### GetN2MbsSmInfoOk

`func (o *ContextUpdateReqData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool)`

GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfo

`func (o *ContextUpdateReqData) SetN2MbsSmInfo(v N2MbsSmInfo)`

SetN2MbsSmInfo sets N2MbsSmInfo field to given value.

### HasN2MbsSmInfo

`func (o *ContextUpdateReqData) HasN2MbsSmInfo() bool`

HasN2MbsSmInfo returns a boolean if a field has been set.

### GetRanIdList

`func (o *ContextUpdateReqData) GetRanIdList() []GlobalRanNodeId`

GetRanIdList returns the RanIdList field if non-nil, zero value otherwise.

### GetRanIdListOk

`func (o *ContextUpdateReqData) GetRanIdListOk() (*[]GlobalRanNodeId, bool)`

GetRanIdListOk returns a tuple with the RanIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanIdList

`func (o *ContextUpdateReqData) SetRanIdList(v []GlobalRanNodeId)`

SetRanIdList sets RanIdList field to given value.

### HasRanIdList

`func (o *ContextUpdateReqData) HasRanIdList() bool`

HasRanIdList returns a boolean if a field has been set.

### GetNoNgapSignallingInd

`func (o *ContextUpdateReqData) GetNoNgapSignallingInd() bool`

GetNoNgapSignallingInd returns the NoNgapSignallingInd field if non-nil, zero value otherwise.

### GetNoNgapSignallingIndOk

`func (o *ContextUpdateReqData) GetNoNgapSignallingIndOk() (*bool, bool)`

GetNoNgapSignallingIndOk returns a tuple with the NoNgapSignallingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoNgapSignallingInd

`func (o *ContextUpdateReqData) SetNoNgapSignallingInd(v bool)`

SetNoNgapSignallingInd sets NoNgapSignallingInd field to given value.

### HasNoNgapSignallingInd

`func (o *ContextUpdateReqData) HasNoNgapSignallingInd() bool`

HasNoNgapSignallingInd returns a boolean if a field has been set.

### GetNotifyUri

`func (o *ContextUpdateReqData) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *ContextUpdateReqData) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *ContextUpdateReqData) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.

### HasNotifyUri

`func (o *ContextUpdateReqData) HasNotifyUri() bool`

HasNotifyUri returns a boolean if a field has been set.

### GetMaxResponseTime

`func (o *ContextUpdateReqData) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *ContextUpdateReqData) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *ContextUpdateReqData) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *ContextUpdateReqData) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetN2MbsInfoChangeInd

`func (o *ContextUpdateReqData) GetN2MbsInfoChangeInd() bool`

GetN2MbsInfoChangeInd returns the N2MbsInfoChangeInd field if non-nil, zero value otherwise.

### GetN2MbsInfoChangeIndOk

`func (o *ContextUpdateReqData) GetN2MbsInfoChangeIndOk() (*bool, bool)`

GetN2MbsInfoChangeIndOk returns a tuple with the N2MbsInfoChangeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsInfoChangeInd

`func (o *ContextUpdateReqData) SetN2MbsInfoChangeInd(v bool)`

SetN2MbsInfoChangeInd sets N2MbsInfoChangeInd field to given value.

### HasN2MbsInfoChangeInd

`func (o *ContextUpdateReqData) HasN2MbsInfoChangeInd() bool`

HasN2MbsInfoChangeInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


