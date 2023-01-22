# SACEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**SACEventType**](SACEventType.md) |  | 
**EventTrigger** | Pointer to [**SACEventTrigger**](SACEventTrigger.md) |  | [optional] 
**EventFilter** | [**[]Snssai**](Snssai.md) |  | 
**NotificationPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NotifThreshold** | Pointer to [**SACInfo**](SACInfo.md) |  | [optional] 
**ImmediateFlag** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSACEvent

`func NewSACEvent(eventType SACEventType, eventFilter []Snssai, ) *SACEvent`

NewSACEvent instantiates a new SACEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSACEventWithDefaults

`func NewSACEventWithDefaults() *SACEvent`

NewSACEventWithDefaults instantiates a new SACEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *SACEvent) GetEventType() SACEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SACEvent) GetEventTypeOk() (*SACEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SACEvent) SetEventType(v SACEventType)`

SetEventType sets EventType field to given value.


### GetEventTrigger

`func (o *SACEvent) GetEventTrigger() SACEventTrigger`

GetEventTrigger returns the EventTrigger field if non-nil, zero value otherwise.

### GetEventTriggerOk

`func (o *SACEvent) GetEventTriggerOk() (*SACEventTrigger, bool)`

GetEventTriggerOk returns a tuple with the EventTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTrigger

`func (o *SACEvent) SetEventTrigger(v SACEventTrigger)`

SetEventTrigger sets EventTrigger field to given value.

### HasEventTrigger

`func (o *SACEvent) HasEventTrigger() bool`

HasEventTrigger returns a boolean if a field has been set.

### GetEventFilter

`func (o *SACEvent) GetEventFilter() []Snssai`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *SACEvent) GetEventFilterOk() (*[]Snssai, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *SACEvent) SetEventFilter(v []Snssai)`

SetEventFilter sets EventFilter field to given value.


### GetNotificationPeriod

`func (o *SACEvent) GetNotificationPeriod() int32`

GetNotificationPeriod returns the NotificationPeriod field if non-nil, zero value otherwise.

### GetNotificationPeriodOk

`func (o *SACEvent) GetNotificationPeriodOk() (*int32, bool)`

GetNotificationPeriodOk returns a tuple with the NotificationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationPeriod

`func (o *SACEvent) SetNotificationPeriod(v int32)`

SetNotificationPeriod sets NotificationPeriod field to given value.

### HasNotificationPeriod

`func (o *SACEvent) HasNotificationPeriod() bool`

HasNotificationPeriod returns a boolean if a field has been set.

### GetNotifThreshold

`func (o *SACEvent) GetNotifThreshold() SACInfo`

GetNotifThreshold returns the NotifThreshold field if non-nil, zero value otherwise.

### GetNotifThresholdOk

`func (o *SACEvent) GetNotifThresholdOk() (*SACInfo, bool)`

GetNotifThresholdOk returns a tuple with the NotifThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifThreshold

`func (o *SACEvent) SetNotifThreshold(v SACInfo)`

SetNotifThreshold sets NotifThreshold field to given value.

### HasNotifThreshold

`func (o *SACEvent) HasNotifThreshold() bool`

HasNotifThreshold returns a boolean if a field has been set.

### GetImmediateFlag

`func (o *SACEvent) GetImmediateFlag() bool`

GetImmediateFlag returns the ImmediateFlag field if non-nil, zero value otherwise.

### GetImmediateFlagOk

`func (o *SACEvent) GetImmediateFlagOk() (*bool, bool)`

GetImmediateFlagOk returns a tuple with the ImmediateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateFlag

`func (o *SACEvent) SetImmediateFlag(v bool)`

SetImmediateFlag sets ImmediateFlag field to given value.

### HasImmediateFlag

`func (o *SACEvent) HasImmediateFlag() bool`

HasImmediateFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


