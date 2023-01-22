# NotifyThresholdCrossingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservedPerfMetricName** | Pointer to **string** |  | [optional] 
**ObservedPerfMetricValue** | Pointer to [**PerfMetricValue**](PerfMetricValue.md) |  | [optional] 
**ObservedPerfMetricDirection** | Pointer to [**PerfMetricDirection**](PerfMetricDirection.md) |  | [optional] 
**ThresholdValue** | Pointer to [**PerfMetricValue**](PerfMetricValue.md) |  | [optional] 
**Hysteresis** | Pointer to [**PerfMetricValue**](PerfMetricValue.md) |  | [optional] 
**MonitorGranularityPeriod** | Pointer to **int32** |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyThresholdCrossingAllOf

`func NewNotifyThresholdCrossingAllOf() *NotifyThresholdCrossingAllOf`

NewNotifyThresholdCrossingAllOf instantiates a new NotifyThresholdCrossingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyThresholdCrossingAllOfWithDefaults

`func NewNotifyThresholdCrossingAllOfWithDefaults() *NotifyThresholdCrossingAllOf`

NewNotifyThresholdCrossingAllOfWithDefaults instantiates a new NotifyThresholdCrossingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservedPerfMetricName

`func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricName() string`

GetObservedPerfMetricName returns the ObservedPerfMetricName field if non-nil, zero value otherwise.

### GetObservedPerfMetricNameOk

`func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricNameOk() (*string, bool)`

GetObservedPerfMetricNameOk returns a tuple with the ObservedPerfMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedPerfMetricName

`func (o *NotifyThresholdCrossingAllOf) SetObservedPerfMetricName(v string)`

SetObservedPerfMetricName sets ObservedPerfMetricName field to given value.

### HasObservedPerfMetricName

`func (o *NotifyThresholdCrossingAllOf) HasObservedPerfMetricName() bool`

HasObservedPerfMetricName returns a boolean if a field has been set.

### GetObservedPerfMetricValue

`func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricValue() PerfMetricValue`

GetObservedPerfMetricValue returns the ObservedPerfMetricValue field if non-nil, zero value otherwise.

### GetObservedPerfMetricValueOk

`func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricValueOk() (*PerfMetricValue, bool)`

GetObservedPerfMetricValueOk returns a tuple with the ObservedPerfMetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedPerfMetricValue

`func (o *NotifyThresholdCrossingAllOf) SetObservedPerfMetricValue(v PerfMetricValue)`

SetObservedPerfMetricValue sets ObservedPerfMetricValue field to given value.

### HasObservedPerfMetricValue

`func (o *NotifyThresholdCrossingAllOf) HasObservedPerfMetricValue() bool`

HasObservedPerfMetricValue returns a boolean if a field has been set.

### GetObservedPerfMetricDirection

`func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricDirection() PerfMetricDirection`

GetObservedPerfMetricDirection returns the ObservedPerfMetricDirection field if non-nil, zero value otherwise.

### GetObservedPerfMetricDirectionOk

`func (o *NotifyThresholdCrossingAllOf) GetObservedPerfMetricDirectionOk() (*PerfMetricDirection, bool)`

GetObservedPerfMetricDirectionOk returns a tuple with the ObservedPerfMetricDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedPerfMetricDirection

`func (o *NotifyThresholdCrossingAllOf) SetObservedPerfMetricDirection(v PerfMetricDirection)`

SetObservedPerfMetricDirection sets ObservedPerfMetricDirection field to given value.

### HasObservedPerfMetricDirection

`func (o *NotifyThresholdCrossingAllOf) HasObservedPerfMetricDirection() bool`

HasObservedPerfMetricDirection returns a boolean if a field has been set.

### GetThresholdValue

`func (o *NotifyThresholdCrossingAllOf) GetThresholdValue() PerfMetricValue`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *NotifyThresholdCrossingAllOf) GetThresholdValueOk() (*PerfMetricValue, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *NotifyThresholdCrossingAllOf) SetThresholdValue(v PerfMetricValue)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *NotifyThresholdCrossingAllOf) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetHysteresis

`func (o *NotifyThresholdCrossingAllOf) GetHysteresis() PerfMetricValue`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *NotifyThresholdCrossingAllOf) GetHysteresisOk() (*PerfMetricValue, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *NotifyThresholdCrossingAllOf) SetHysteresis(v PerfMetricValue)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *NotifyThresholdCrossingAllOf) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.

### GetMonitorGranularityPeriod

`func (o *NotifyThresholdCrossingAllOf) GetMonitorGranularityPeriod() int32`

GetMonitorGranularityPeriod returns the MonitorGranularityPeriod field if non-nil, zero value otherwise.

### GetMonitorGranularityPeriodOk

`func (o *NotifyThresholdCrossingAllOf) GetMonitorGranularityPeriodOk() (*int32, bool)`

GetMonitorGranularityPeriodOk returns a tuple with the MonitorGranularityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGranularityPeriod

`func (o *NotifyThresholdCrossingAllOf) SetMonitorGranularityPeriod(v int32)`

SetMonitorGranularityPeriod sets MonitorGranularityPeriod field to given value.

### HasMonitorGranularityPeriod

`func (o *NotifyThresholdCrossingAllOf) HasMonitorGranularityPeriod() bool`

HasMonitorGranularityPeriod returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyThresholdCrossingAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyThresholdCrossingAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyThresholdCrossingAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyThresholdCrossingAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


