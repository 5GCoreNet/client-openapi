# NotifyThresholdCrossing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**ObservedPerfMetricName** | Pointer to **string** |  | [optional] 
**ObservedPerfMetricValue** | Pointer to [**PerfMetricValue**](PerfMetricValue.md) |  | [optional] 
**ObservedPerfMetricDirection** | Pointer to [**PerfMetricDirection**](PerfMetricDirection.md) |  | [optional] 
**ThresholdValue** | Pointer to [**PerfMetricValue**](PerfMetricValue.md) |  | [optional] 
**Hysteresis** | Pointer to [**PerfMetricValue**](PerfMetricValue.md) |  | [optional] 
**MonitorGranularityPeriod** | Pointer to **int32** |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyThresholdCrossing

`func NewNotifyThresholdCrossing(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, ) *NotifyThresholdCrossing`

NewNotifyThresholdCrossing instantiates a new NotifyThresholdCrossing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyThresholdCrossingWithDefaults

`func NewNotifyThresholdCrossingWithDefaults() *NotifyThresholdCrossing`

NewNotifyThresholdCrossingWithDefaults instantiates a new NotifyThresholdCrossing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyThresholdCrossing) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyThresholdCrossing) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyThresholdCrossing) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyThresholdCrossing) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyThresholdCrossing) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyThresholdCrossing) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyThresholdCrossing) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyThresholdCrossing) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyThresholdCrossing) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyThresholdCrossing) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyThresholdCrossing) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyThresholdCrossing) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyThresholdCrossing) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyThresholdCrossing) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyThresholdCrossing) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetObservedPerfMetricName

`func (o *NotifyThresholdCrossing) GetObservedPerfMetricName() string`

GetObservedPerfMetricName returns the ObservedPerfMetricName field if non-nil, zero value otherwise.

### GetObservedPerfMetricNameOk

`func (o *NotifyThresholdCrossing) GetObservedPerfMetricNameOk() (*string, bool)`

GetObservedPerfMetricNameOk returns a tuple with the ObservedPerfMetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedPerfMetricName

`func (o *NotifyThresholdCrossing) SetObservedPerfMetricName(v string)`

SetObservedPerfMetricName sets ObservedPerfMetricName field to given value.

### HasObservedPerfMetricName

`func (o *NotifyThresholdCrossing) HasObservedPerfMetricName() bool`

HasObservedPerfMetricName returns a boolean if a field has been set.

### GetObservedPerfMetricValue

`func (o *NotifyThresholdCrossing) GetObservedPerfMetricValue() PerfMetricValue`

GetObservedPerfMetricValue returns the ObservedPerfMetricValue field if non-nil, zero value otherwise.

### GetObservedPerfMetricValueOk

`func (o *NotifyThresholdCrossing) GetObservedPerfMetricValueOk() (*PerfMetricValue, bool)`

GetObservedPerfMetricValueOk returns a tuple with the ObservedPerfMetricValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedPerfMetricValue

`func (o *NotifyThresholdCrossing) SetObservedPerfMetricValue(v PerfMetricValue)`

SetObservedPerfMetricValue sets ObservedPerfMetricValue field to given value.

### HasObservedPerfMetricValue

`func (o *NotifyThresholdCrossing) HasObservedPerfMetricValue() bool`

HasObservedPerfMetricValue returns a boolean if a field has been set.

### GetObservedPerfMetricDirection

`func (o *NotifyThresholdCrossing) GetObservedPerfMetricDirection() PerfMetricDirection`

GetObservedPerfMetricDirection returns the ObservedPerfMetricDirection field if non-nil, zero value otherwise.

### GetObservedPerfMetricDirectionOk

`func (o *NotifyThresholdCrossing) GetObservedPerfMetricDirectionOk() (*PerfMetricDirection, bool)`

GetObservedPerfMetricDirectionOk returns a tuple with the ObservedPerfMetricDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedPerfMetricDirection

`func (o *NotifyThresholdCrossing) SetObservedPerfMetricDirection(v PerfMetricDirection)`

SetObservedPerfMetricDirection sets ObservedPerfMetricDirection field to given value.

### HasObservedPerfMetricDirection

`func (o *NotifyThresholdCrossing) HasObservedPerfMetricDirection() bool`

HasObservedPerfMetricDirection returns a boolean if a field has been set.

### GetThresholdValue

`func (o *NotifyThresholdCrossing) GetThresholdValue() PerfMetricValue`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *NotifyThresholdCrossing) GetThresholdValueOk() (*PerfMetricValue, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *NotifyThresholdCrossing) SetThresholdValue(v PerfMetricValue)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *NotifyThresholdCrossing) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.

### GetHysteresis

`func (o *NotifyThresholdCrossing) GetHysteresis() PerfMetricValue`

GetHysteresis returns the Hysteresis field if non-nil, zero value otherwise.

### GetHysteresisOk

`func (o *NotifyThresholdCrossing) GetHysteresisOk() (*PerfMetricValue, bool)`

GetHysteresisOk returns a tuple with the Hysteresis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHysteresis

`func (o *NotifyThresholdCrossing) SetHysteresis(v PerfMetricValue)`

SetHysteresis sets Hysteresis field to given value.

### HasHysteresis

`func (o *NotifyThresholdCrossing) HasHysteresis() bool`

HasHysteresis returns a boolean if a field has been set.

### GetMonitorGranularityPeriod

`func (o *NotifyThresholdCrossing) GetMonitorGranularityPeriod() int32`

GetMonitorGranularityPeriod returns the MonitorGranularityPeriod field if non-nil, zero value otherwise.

### GetMonitorGranularityPeriodOk

`func (o *NotifyThresholdCrossing) GetMonitorGranularityPeriodOk() (*int32, bool)`

GetMonitorGranularityPeriodOk returns a tuple with the MonitorGranularityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorGranularityPeriod

`func (o *NotifyThresholdCrossing) SetMonitorGranularityPeriod(v int32)`

SetMonitorGranularityPeriod sets MonitorGranularityPeriod field to given value.

### HasMonitorGranularityPeriod

`func (o *NotifyThresholdCrossing) HasMonitorGranularityPeriod() bool`

HasMonitorGranularityPeriod returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyThresholdCrossing) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyThresholdCrossing) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyThresholdCrossing) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyThresholdCrossing) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


