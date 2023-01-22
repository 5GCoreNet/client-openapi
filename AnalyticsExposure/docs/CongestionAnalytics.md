# CongestionAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CngType** | [**CongestionType**](CongestionType.md) |  | 
**TmWdw** | [**TimeWindow**](TimeWindow.md) |  | 
**Nsi** | [**ThresholdLevel**](ThresholdLevel.md) |  | 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TopAppListUl** | Pointer to [**[]TopApplication**](TopApplication.md) |  | [optional] 
**TopAppListDl** | Pointer to [**[]TopApplication**](TopApplication.md) |  | [optional] 

## Methods

### NewCongestionAnalytics

`func NewCongestionAnalytics(cngType CongestionType, tmWdw TimeWindow, nsi ThresholdLevel, ) *CongestionAnalytics`

NewCongestionAnalytics instantiates a new CongestionAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCongestionAnalyticsWithDefaults

`func NewCongestionAnalyticsWithDefaults() *CongestionAnalytics`

NewCongestionAnalyticsWithDefaults instantiates a new CongestionAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCngType

`func (o *CongestionAnalytics) GetCngType() CongestionType`

GetCngType returns the CngType field if non-nil, zero value otherwise.

### GetCngTypeOk

`func (o *CongestionAnalytics) GetCngTypeOk() (*CongestionType, bool)`

GetCngTypeOk returns a tuple with the CngType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCngType

`func (o *CongestionAnalytics) SetCngType(v CongestionType)`

SetCngType sets CngType field to given value.


### GetTmWdw

`func (o *CongestionAnalytics) GetTmWdw() TimeWindow`

GetTmWdw returns the TmWdw field if non-nil, zero value otherwise.

### GetTmWdwOk

`func (o *CongestionAnalytics) GetTmWdwOk() (*TimeWindow, bool)`

GetTmWdwOk returns a tuple with the TmWdw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmWdw

`func (o *CongestionAnalytics) SetTmWdw(v TimeWindow)`

SetTmWdw sets TmWdw field to given value.


### GetNsi

`func (o *CongestionAnalytics) GetNsi() ThresholdLevel`

GetNsi returns the Nsi field if non-nil, zero value otherwise.

### GetNsiOk

`func (o *CongestionAnalytics) GetNsiOk() (*ThresholdLevel, bool)`

GetNsiOk returns a tuple with the Nsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsi

`func (o *CongestionAnalytics) SetNsi(v ThresholdLevel)`

SetNsi sets Nsi field to given value.


### GetConfidence

`func (o *CongestionAnalytics) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CongestionAnalytics) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CongestionAnalytics) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *CongestionAnalytics) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetTopAppListUl

`func (o *CongestionAnalytics) GetTopAppListUl() []TopApplication`

GetTopAppListUl returns the TopAppListUl field if non-nil, zero value otherwise.

### GetTopAppListUlOk

`func (o *CongestionAnalytics) GetTopAppListUlOk() (*[]TopApplication, bool)`

GetTopAppListUlOk returns a tuple with the TopAppListUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAppListUl

`func (o *CongestionAnalytics) SetTopAppListUl(v []TopApplication)`

SetTopAppListUl sets TopAppListUl field to given value.

### HasTopAppListUl

`func (o *CongestionAnalytics) HasTopAppListUl() bool`

HasTopAppListUl returns a boolean if a field has been set.

### GetTopAppListDl

`func (o *CongestionAnalytics) GetTopAppListDl() []TopApplication`

GetTopAppListDl returns the TopAppListDl field if non-nil, zero value otherwise.

### GetTopAppListDlOk

`func (o *CongestionAnalytics) GetTopAppListDlOk() (*[]TopApplication, bool)`

GetTopAppListDlOk returns a tuple with the TopAppListDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAppListDl

`func (o *CongestionAnalytics) SetTopAppListDl(v []TopApplication)`

SetTopAppListDl sets TopAppListDl field to given value.

### HasTopAppListDl

`func (o *CongestionAnalytics) HasTopAppListDl() bool`

HasTopAppListDl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


