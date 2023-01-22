# MDAOutputIEFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MDAOutputIEName** | Pointer to **string** |  | [optional] 
**FilterValue** | Pointer to **string** |  | [optional] 
**Threshold** | Pointer to [**ThresholdInfo**](ThresholdInfo.md) |  | [optional] 
**AnalyticsPeriod** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**TimeOut** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewMDAOutputIEFilter

`func NewMDAOutputIEFilter() *MDAOutputIEFilter`

NewMDAOutputIEFilter instantiates a new MDAOutputIEFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMDAOutputIEFilterWithDefaults

`func NewMDAOutputIEFilterWithDefaults() *MDAOutputIEFilter`

NewMDAOutputIEFilterWithDefaults instantiates a new MDAOutputIEFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMDAOutputIEName

`func (o *MDAOutputIEFilter) GetMDAOutputIEName() string`

GetMDAOutputIEName returns the MDAOutputIEName field if non-nil, zero value otherwise.

### GetMDAOutputIENameOk

`func (o *MDAOutputIEFilter) GetMDAOutputIENameOk() (*string, bool)`

GetMDAOutputIENameOk returns a tuple with the MDAOutputIEName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAOutputIEName

`func (o *MDAOutputIEFilter) SetMDAOutputIEName(v string)`

SetMDAOutputIEName sets MDAOutputIEName field to given value.

### HasMDAOutputIEName

`func (o *MDAOutputIEFilter) HasMDAOutputIEName() bool`

HasMDAOutputIEName returns a boolean if a field has been set.

### GetFilterValue

`func (o *MDAOutputIEFilter) GetFilterValue() string`

GetFilterValue returns the FilterValue field if non-nil, zero value otherwise.

### GetFilterValueOk

`func (o *MDAOutputIEFilter) GetFilterValueOk() (*string, bool)`

GetFilterValueOk returns a tuple with the FilterValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterValue

`func (o *MDAOutputIEFilter) SetFilterValue(v string)`

SetFilterValue sets FilterValue field to given value.

### HasFilterValue

`func (o *MDAOutputIEFilter) HasFilterValue() bool`

HasFilterValue returns a boolean if a field has been set.

### GetThreshold

`func (o *MDAOutputIEFilter) GetThreshold() ThresholdInfo`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MDAOutputIEFilter) GetThresholdOk() (*ThresholdInfo, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MDAOutputIEFilter) SetThreshold(v ThresholdInfo)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MDAOutputIEFilter) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetAnalyticsPeriod

`func (o *MDAOutputIEFilter) GetAnalyticsPeriod() []time.Time`

GetAnalyticsPeriod returns the AnalyticsPeriod field if non-nil, zero value otherwise.

### GetAnalyticsPeriodOk

`func (o *MDAOutputIEFilter) GetAnalyticsPeriodOk() (*[]time.Time, bool)`

GetAnalyticsPeriodOk returns a tuple with the AnalyticsPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsPeriod

`func (o *MDAOutputIEFilter) SetAnalyticsPeriod(v []time.Time)`

SetAnalyticsPeriod sets AnalyticsPeriod field to given value.

### HasAnalyticsPeriod

`func (o *MDAOutputIEFilter) HasAnalyticsPeriod() bool`

HasAnalyticsPeriod returns a boolean if a field has been set.

### GetTimeOut

`func (o *MDAOutputIEFilter) GetTimeOut() time.Time`

GetTimeOut returns the TimeOut field if non-nil, zero value otherwise.

### GetTimeOutOk

`func (o *MDAOutputIEFilter) GetTimeOutOk() (*time.Time, bool)`

GetTimeOutOk returns a tuple with the TimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOut

`func (o *MDAOutputIEFilter) SetTimeOut(v time.Time)`

SetTimeOut sets TimeOut field to given value.

### HasTimeOut

`func (o *MDAOutputIEFilter) HasTimeOut() bool`

HasTimeOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


