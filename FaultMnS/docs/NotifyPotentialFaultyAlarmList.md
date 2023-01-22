# NotifyPotentialFaultyAlarmList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewNotifyPotentialFaultyAlarmList

`func NewNotifyPotentialFaultyAlarmList(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, reason string, ) *NotifyPotentialFaultyAlarmList`

NewNotifyPotentialFaultyAlarmList instantiates a new NotifyPotentialFaultyAlarmList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyPotentialFaultyAlarmListWithDefaults

`func NewNotifyPotentialFaultyAlarmListWithDefaults() *NotifyPotentialFaultyAlarmList`

NewNotifyPotentialFaultyAlarmListWithDefaults instantiates a new NotifyPotentialFaultyAlarmList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyPotentialFaultyAlarmList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyPotentialFaultyAlarmList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyPotentialFaultyAlarmList) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyPotentialFaultyAlarmList) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyPotentialFaultyAlarmList) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyPotentialFaultyAlarmList) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyPotentialFaultyAlarmList) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyPotentialFaultyAlarmList) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyPotentialFaultyAlarmList) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyPotentialFaultyAlarmList) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyPotentialFaultyAlarmList) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyPotentialFaultyAlarmList) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyPotentialFaultyAlarmList) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyPotentialFaultyAlarmList) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyPotentialFaultyAlarmList) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetReason

`func (o *NotifyPotentialFaultyAlarmList) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotifyPotentialFaultyAlarmList) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotifyPotentialFaultyAlarmList) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


