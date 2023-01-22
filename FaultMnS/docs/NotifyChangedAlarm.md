# NotifyChangedAlarm

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

## Methods

### NewNotifyChangedAlarm

`func NewNotifyChangedAlarm(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, ) *NotifyChangedAlarm`

NewNotifyChangedAlarm instantiates a new NotifyChangedAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyChangedAlarmWithDefaults

`func NewNotifyChangedAlarmWithDefaults() *NotifyChangedAlarm`

NewNotifyChangedAlarmWithDefaults instantiates a new NotifyChangedAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyChangedAlarm) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyChangedAlarm) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyChangedAlarm) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyChangedAlarm) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyChangedAlarm) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyChangedAlarm) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyChangedAlarm) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyChangedAlarm) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyChangedAlarm) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyChangedAlarm) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyChangedAlarm) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyChangedAlarm) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyChangedAlarm) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyChangedAlarm) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyChangedAlarm) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetAlarmId

`func (o *NotifyChangedAlarm) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyChangedAlarm) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyChangedAlarm) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyChangedAlarm) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyChangedAlarm) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyChangedAlarm) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyChangedAlarm) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyChangedAlarm) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyChangedAlarm) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetPerceivedSeverity

`func (o *NotifyChangedAlarm) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyChangedAlarm) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyChangedAlarm) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


