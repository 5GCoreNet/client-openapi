# NwdafCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** |  | 
**AnalyticsIds** | Pointer to **[]string** |  | [optional] 
**SnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**ServingNfTypeList** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**ServingNfSetIdList** | Pointer to **[]string** |  | [optional] 
**MlAnalyticsList** | Pointer to [**[]MlAnalyticsInfo**](MlAnalyticsInfo.md) |  | [optional] 

## Methods

### NewNwdafCond

`func NewNwdafCond(conditionType string, ) *NwdafCond`

NewNwdafCond instantiates a new NwdafCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNwdafCondWithDefaults

`func NewNwdafCondWithDefaults() *NwdafCond`

NewNwdafCondWithDefaults instantiates a new NwdafCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *NwdafCond) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *NwdafCond) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *NwdafCond) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetAnalyticsIds

`func (o *NwdafCond) GetAnalyticsIds() []string`

GetAnalyticsIds returns the AnalyticsIds field if non-nil, zero value otherwise.

### GetAnalyticsIdsOk

`func (o *NwdafCond) GetAnalyticsIdsOk() (*[]string, bool)`

GetAnalyticsIdsOk returns a tuple with the AnalyticsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsIds

`func (o *NwdafCond) SetAnalyticsIds(v []string)`

SetAnalyticsIds sets AnalyticsIds field to given value.

### HasAnalyticsIds

`func (o *NwdafCond) HasAnalyticsIds() bool`

HasAnalyticsIds returns a boolean if a field has been set.

### GetSnssaiList

`func (o *NwdafCond) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *NwdafCond) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *NwdafCond) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.

### HasSnssaiList

`func (o *NwdafCond) HasSnssaiList() bool`

HasSnssaiList returns a boolean if a field has been set.

### GetTaiList

`func (o *NwdafCond) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NwdafCond) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NwdafCond) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NwdafCond) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NwdafCond) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NwdafCond) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NwdafCond) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NwdafCond) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetServingNfTypeList

`func (o *NwdafCond) GetServingNfTypeList() []NFType`

GetServingNfTypeList returns the ServingNfTypeList field if non-nil, zero value otherwise.

### GetServingNfTypeListOk

`func (o *NwdafCond) GetServingNfTypeListOk() (*[]NFType, bool)`

GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfTypeList

`func (o *NwdafCond) SetServingNfTypeList(v []NFType)`

SetServingNfTypeList sets ServingNfTypeList field to given value.

### HasServingNfTypeList

`func (o *NwdafCond) HasServingNfTypeList() bool`

HasServingNfTypeList returns a boolean if a field has been set.

### GetServingNfSetIdList

`func (o *NwdafCond) GetServingNfSetIdList() []string`

GetServingNfSetIdList returns the ServingNfSetIdList field if non-nil, zero value otherwise.

### GetServingNfSetIdListOk

`func (o *NwdafCond) GetServingNfSetIdListOk() (*[]string, bool)`

GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfSetIdList

`func (o *NwdafCond) SetServingNfSetIdList(v []string)`

SetServingNfSetIdList sets ServingNfSetIdList field to given value.

### HasServingNfSetIdList

`func (o *NwdafCond) HasServingNfSetIdList() bool`

HasServingNfSetIdList returns a boolean if a field has been set.

### GetMlAnalyticsList

`func (o *NwdafCond) GetMlAnalyticsList() []MlAnalyticsInfo`

GetMlAnalyticsList returns the MlAnalyticsList field if non-nil, zero value otherwise.

### GetMlAnalyticsListOk

`func (o *NwdafCond) GetMlAnalyticsListOk() (*[]MlAnalyticsInfo, bool)`

GetMlAnalyticsListOk returns a tuple with the MlAnalyticsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlAnalyticsList

`func (o *NwdafCond) SetMlAnalyticsList(v []MlAnalyticsInfo)`

SetMlAnalyticsList sets MlAnalyticsList field to given value.

### HasMlAnalyticsList

`func (o *NwdafCond) HasMlAnalyticsList() bool`

HasMlAnalyticsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


