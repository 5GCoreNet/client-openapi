# NrfInfoServedNwdafInfoValue

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

### NewNrfInfoServedNwdafInfoValue

`func NewNrfInfoServedNwdafInfoValue() *NrfInfoServedNwdafInfoValue`

NewNrfInfoServedNwdafInfoValue instantiates a new NrfInfoServedNwdafInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedNwdafInfoValueWithDefaults

`func NewNrfInfoServedNwdafInfoValueWithDefaults() *NrfInfoServedNwdafInfoValue`

NewNrfInfoServedNwdafInfoValueWithDefaults instantiates a new NrfInfoServedNwdafInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIds

`func (o *NrfInfoServedNwdafInfoValue) GetEventIds() []EventId`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *NrfInfoServedNwdafInfoValue) GetEventIdsOk() (*[]EventId, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *NrfInfoServedNwdafInfoValue) SetEventIds(v []EventId)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *NrfInfoServedNwdafInfoValue) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetNwdafEvents

`func (o *NrfInfoServedNwdafInfoValue) GetNwdafEvents() []NwdafEvent`

GetNwdafEvents returns the NwdafEvents field if non-nil, zero value otherwise.

### GetNwdafEventsOk

`func (o *NrfInfoServedNwdafInfoValue) GetNwdafEventsOk() (*[]NwdafEvent, bool)`

GetNwdafEventsOk returns a tuple with the NwdafEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvents

`func (o *NrfInfoServedNwdafInfoValue) SetNwdafEvents(v []NwdafEvent)`

SetNwdafEvents sets NwdafEvents field to given value.

### HasNwdafEvents

`func (o *NrfInfoServedNwdafInfoValue) HasNwdafEvents() bool`

HasNwdafEvents returns a boolean if a field has been set.

### GetTaiList

`func (o *NrfInfoServedNwdafInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedNwdafInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedNwdafInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedNwdafInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedNwdafInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedNwdafInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedNwdafInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedNwdafInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetNwdafCapability

`func (o *NrfInfoServedNwdafInfoValue) GetNwdafCapability() NwdafCapability`

GetNwdafCapability returns the NwdafCapability field if non-nil, zero value otherwise.

### GetNwdafCapabilityOk

`func (o *NrfInfoServedNwdafInfoValue) GetNwdafCapabilityOk() (*NwdafCapability, bool)`

GetNwdafCapabilityOk returns a tuple with the NwdafCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafCapability

`func (o *NrfInfoServedNwdafInfoValue) SetNwdafCapability(v NwdafCapability)`

SetNwdafCapability sets NwdafCapability field to given value.

### HasNwdafCapability

`func (o *NrfInfoServedNwdafInfoValue) HasNwdafCapability() bool`

HasNwdafCapability returns a boolean if a field has been set.

### GetAnalyticsDelay

`func (o *NrfInfoServedNwdafInfoValue) GetAnalyticsDelay() int32`

GetAnalyticsDelay returns the AnalyticsDelay field if non-nil, zero value otherwise.

### GetAnalyticsDelayOk

`func (o *NrfInfoServedNwdafInfoValue) GetAnalyticsDelayOk() (*int32, bool)`

GetAnalyticsDelayOk returns a tuple with the AnalyticsDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsDelay

`func (o *NrfInfoServedNwdafInfoValue) SetAnalyticsDelay(v int32)`

SetAnalyticsDelay sets AnalyticsDelay field to given value.

### HasAnalyticsDelay

`func (o *NrfInfoServedNwdafInfoValue) HasAnalyticsDelay() bool`

HasAnalyticsDelay returns a boolean if a field has been set.

### GetServingNfSetIdList

`func (o *NrfInfoServedNwdafInfoValue) GetServingNfSetIdList() []string`

GetServingNfSetIdList returns the ServingNfSetIdList field if non-nil, zero value otherwise.

### GetServingNfSetIdListOk

`func (o *NrfInfoServedNwdafInfoValue) GetServingNfSetIdListOk() (*[]string, bool)`

GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfSetIdList

`func (o *NrfInfoServedNwdafInfoValue) SetServingNfSetIdList(v []string)`

SetServingNfSetIdList sets ServingNfSetIdList field to given value.

### HasServingNfSetIdList

`func (o *NrfInfoServedNwdafInfoValue) HasServingNfSetIdList() bool`

HasServingNfSetIdList returns a boolean if a field has been set.

### GetServingNfTypeList

`func (o *NrfInfoServedNwdafInfoValue) GetServingNfTypeList() []NFType`

GetServingNfTypeList returns the ServingNfTypeList field if non-nil, zero value otherwise.

### GetServingNfTypeListOk

`func (o *NrfInfoServedNwdafInfoValue) GetServingNfTypeListOk() (*[]NFType, bool)`

GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfTypeList

`func (o *NrfInfoServedNwdafInfoValue) SetServingNfTypeList(v []NFType)`

SetServingNfTypeList sets ServingNfTypeList field to given value.

### HasServingNfTypeList

`func (o *NrfInfoServedNwdafInfoValue) HasServingNfTypeList() bool`

HasServingNfTypeList returns a boolean if a field has been set.

### GetMlAnalyticsList

`func (o *NrfInfoServedNwdafInfoValue) GetMlAnalyticsList() []MlAnalyticsInfo`

GetMlAnalyticsList returns the MlAnalyticsList field if non-nil, zero value otherwise.

### GetMlAnalyticsListOk

`func (o *NrfInfoServedNwdafInfoValue) GetMlAnalyticsListOk() (*[]MlAnalyticsInfo, bool)`

GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlAnalyticsList

`func (o *NrfInfoServedNwdafInfoValue) SetMlAnalyticsList(v []MlAnalyticsInfo)`

SetMlAnalyticsList sets MlAnalyticsList field to given value.

### HasMlAnalyticsList

`func (o *NrfInfoServedNwdafInfoValue) HasMlAnalyticsList() bool`

HasMlAnalyticsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


