# NotifyCorrelatedNotificationChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**AlarmId** | **string** |  | 
**CorrelatedNotifications** | [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotifyCorrelatedNotificationChanged

`func NewNotifyCorrelatedNotificationChanged(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, correlatedNotifications []CorrelatedNotification, ) *NotifyCorrelatedNotificationChanged`

NewNotifyCorrelatedNotificationChanged instantiates a new NotifyCorrelatedNotificationChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyCorrelatedNotificationChangedWithDefaults

`func NewNotifyCorrelatedNotificationChangedWithDefaults() *NotifyCorrelatedNotificationChanged`

NewNotifyCorrelatedNotificationChangedWithDefaults instantiates a new NotifyCorrelatedNotificationChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyCorrelatedNotificationChanged) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyCorrelatedNotificationChanged) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyCorrelatedNotificationChanged) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyCorrelatedNotificationChanged) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyCorrelatedNotificationChanged) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyCorrelatedNotificationChanged) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyCorrelatedNotificationChanged) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyCorrelatedNotificationChanged) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyCorrelatedNotificationChanged) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyCorrelatedNotificationChanged) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyCorrelatedNotificationChanged) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyCorrelatedNotificationChanged) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyCorrelatedNotificationChanged) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyCorrelatedNotificationChanged) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyCorrelatedNotificationChanged) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetAlarmId

`func (o *NotifyCorrelatedNotificationChanged) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyCorrelatedNotificationChanged) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyCorrelatedNotificationChanged) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetCorrelatedNotifications

`func (o *NotifyCorrelatedNotificationChanged) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyCorrelatedNotificationChanged) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyCorrelatedNotificationChanged) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.


### GetRootCauseIndicator

`func (o *NotifyCorrelatedNotificationChanged) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyCorrelatedNotificationChanged) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyCorrelatedNotificationChanged) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyCorrelatedNotificationChanged) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


