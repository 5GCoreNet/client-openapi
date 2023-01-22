# NwdafInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIds** | Pointer to [**[]EventId**](EventId.md) |  | [optional] 
**NwdafEvents** | Pointer to [**[]NwdafEvent**](NwdafEvent.md) |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**NwdafCapability** | Pointer to [**NwdafCapability**](NwdafCapability.md) |  | [optional] 
**AnalyticsDelay** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ServingNfSetIdList** | Pointer to **[]string** |  | [optional] 
**ServingNfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**MlAnalyticsList** | Pointer to [**[]MlAnalyticsInfo**](MlAnalyticsInfo.md) |  | [optional] 

## Methods

### NewNwdafInfo

`func NewNwdafInfo() *NwdafInfo`

NewNwdafInfo instantiates a new NwdafInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafInfoWithDefaults

`func NewNwdafInfoWithDefaults() *NwdafInfo`

NewNwdafInfoWithDefaults instantiates a new NwdafInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIds

`func (o *NwdafInfo) GetEventIds() []EventId`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *NwdafInfo) GetEventIdsOk() (*[]EventId, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *NwdafInfo) SetEventIds(v []EventId)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *NwdafInfo) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetNwdafEvents

`func (o *NwdafInfo) GetNwdafEvents() []NwdafEvent`

GetNwdafEvents returns the NwdafEvents field if non-nil, zero value otherwise.

### GetNwdafEventsOk

`func (o *NwdafInfo) GetNwdafEventsOk() (*[]NwdafEvent, bool)`

GetNwdafEventsOk returns a tuple with the NwdafEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvents

`func (o *NwdafInfo) SetNwdafEvents(v []NwdafEvent)`

SetNwdafEvents sets NwdafEvents field to given value.

### HasNwdafEvents

`func (o *NwdafInfo) HasNwdafEvents() bool`

HasNwdafEvents returns a boolean if a field has been set.

### GetTaiList

`func (o *NwdafInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NwdafInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NwdafInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NwdafInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NwdafInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NwdafInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NwdafInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NwdafInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetNwdafCapability

`func (o *NwdafInfo) GetNwdafCapability() NwdafCapability`

GetNwdafCapability returns the NwdafCapability field if non-nil, zero value otherwise.

### GetNwdafCapabilityOk

`func (o *NwdafInfo) GetNwdafCapabilityOk() (*NwdafCapability, bool)`

GetNwdafCapabilityOk returns a tuple with the NwdafCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafCapability

`func (o *NwdafInfo) SetNwdafCapability(v NwdafCapability)`

SetNwdafCapability sets NwdafCapability field to given value.

### HasNwdafCapability

`func (o *NwdafInfo) HasNwdafCapability() bool`

HasNwdafCapability returns a boolean if a field has been set.

### GetAnalyticsDelay

`func (o *NwdafInfo) GetAnalyticsDelay() int32`

GetAnalyticsDelay returns the AnalyticsDelay field if non-nil, zero value otherwise.

### GetAnalyticsDelayOk

`func (o *NwdafInfo) GetAnalyticsDelayOk() (*int32, bool)`

GetAnalyticsDelayOk returns a tuple with the AnalyticsDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsDelay

`func (o *NwdafInfo) SetAnalyticsDelay(v int32)`

SetAnalyticsDelay sets AnalyticsDelay field to given value.

### HasAnalyticsDelay

`func (o *NwdafInfo) HasAnalyticsDelay() bool`

HasAnalyticsDelay returns a boolean if a field has been set.

### GetServingNfSetIdList

`func (o *NwdafInfo) GetServingNfSetIdList() []string`

GetServingNfSetIdList returns the ServingNfSetIdList field if non-nil, zero value otherwise.

### GetServingNfSetIdListOk

`func (o *NwdafInfo) GetServingNfSetIdListOk() (*[]string, bool)`

GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfSetIdList

`func (o *NwdafInfo) SetServingNfSetIdList(v []string)`

SetServingNfSetIdList sets ServingNfSetIdList field to given value.

### HasServingNfSetIdList

`func (o *NwdafInfo) HasServingNfSetIdList() bool`

HasServingNfSetIdList returns a boolean if a field has been set.

### GetServingNfTypeList

`func (o *NwdafInfo) GetServingNfTypeList() []NFType`

GetServingNfTypeList returns the ServingNfTypeList field if non-nil, zero value otherwise.

### GetServingNfTypeListOk

`func (o *NwdafInfo) GetServingNfTypeListOk() (*[]NFType, bool)`

GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfTypeList

`func (o *NwdafInfo) SetServingNfTypeList(v []NFType)`

SetServingNfTypeList sets ServingNfTypeList field to given value.

### HasServingNfTypeList

`func (o *NwdafInfo) HasServingNfTypeList() bool`

HasServingNfTypeList returns a boolean if a field has been set.

### GetMlAnalyticsList

`func (o *NwdafInfo) GetMlAnalyticsList() []MlAnalyticsInfo`

GetMlAnalyticsList returns the MlAnalyticsList field if non-nil, zero value otherwise.

### GetMlAnalyticsListOk

`func (o *NwdafInfo) GetMlAnalyticsListOk() (*[]MlAnalyticsInfo, bool)`

GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlAnalyticsList

`func (o *NwdafInfo) SetMlAnalyticsList(v []MlAnalyticsInfo)`

SetMlAnalyticsList sets MlAnalyticsList field to given value.

### HasMlAnalyticsList

`func (o *NwdafInfo) HasMlAnalyticsList() bool`

HasMlAnalyticsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


