# NotifyClearedAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**AlarmId** | **string** |  | 
**AlarmType** | [**AlarmType**](AlarmType.md) |  | 
**ProbableCause** | [**ProbableCause**](ProbableCause.md) |  | 
**PerceivedSeverity** | [**PerceivedSeverity**](PerceivedSeverity.md) |  | 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**ClearUserId** | Pointer to **string** |  | [optional] 
**ClearSystemId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyClearedAlarm

`func NewNotifyClearedAlarm(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, ) *NotifyClearedAlarm`

NewNotifyClearedAlarm instantiates a new NotifyClearedAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyClearedAlarmWithDefaults

`func NewNotifyClearedAlarmWithDefaults() *NotifyClearedAlarm`

NewNotifyClearedAlarmWithDefaults instantiates a new NotifyClearedAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyClearedAlarm) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyClearedAlarm) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyClearedAlarm) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyClearedAlarm) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyClearedAlarm) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyClearedAlarm) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyClearedAlarm) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyClearedAlarm) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyClearedAlarm) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyClearedAlarm) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyClearedAlarm) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyClearedAlarm) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyClearedAlarm) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyClearedAlarm) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyClearedAlarm) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetAlarmId

`func (o *NotifyClearedAlarm) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyClearedAlarm) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyClearedAlarm) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyClearedAlarm) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyClearedAlarm) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyClearedAlarm) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyClearedAlarm) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyClearedAlarm) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyClearedAlarm) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetPerceivedSeverity

`func (o *NotifyClearedAlarm) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyClearedAlarm) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyClearedAlarm) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.


### GetCorrelatedNotifications

`func (o *NotifyClearedAlarm) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyClearedAlarm) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyClearedAlarm) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyClearedAlarm) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetClearUserId

`func (o *NotifyClearedAlarm) GetClearUserId() string`

GetClearUserId returns the ClearUserId field if non-nil, zero value otherwise.

### GetClearUserIdOk

`func (o *NotifyClearedAlarm) GetClearUserIdOk() (*string, bool)`

GetClearUserIdOk returns a tuple with the ClearUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearUserId

`func (o *NotifyClearedAlarm) SetClearUserId(v string)`

SetClearUserId sets ClearUserId field to given value.

### HasClearUserId

`func (o *NotifyClearedAlarm) HasClearUserId() bool`

HasClearUserId returns a boolean if a field has been set.

### GetClearSystemId

`func (o *NotifyClearedAlarm) GetClearSystemId() string`

GetClearSystemId returns the ClearSystemId field if non-nil, zero value otherwise.

### GetClearSystemIdOk

`func (o *NotifyClearedAlarm) GetClearSystemIdOk() (*string, bool)`

GetClearSystemIdOk returns a tuple with the ClearSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearSystemId

`func (o *NotifyClearedAlarm) SetClearSystemId(v string)`

SetClearSystemId sets ClearSystemId field to given value.

### HasClearSystemId

`func (o *NotifyClearedAlarm) HasClearSystemId() bool`

HasClearSystemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


