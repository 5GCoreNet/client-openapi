# NotifyAckStateChanged

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
**AckState** | [**AckState**](AckState.md) |  | 
**AckUserId** | **string** |  | 
**AckSystemId** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyAckStateChanged

`func NewNotifyAckStateChanged(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, alarmId string, alarmType AlarmType, probableCause ProbableCause, perceivedSeverity PerceivedSeverity, ackState AckState, ackUserId string, ) *NotifyAckStateChanged`

NewNotifyAckStateChanged instantiates a new NotifyAckStateChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyAckStateChangedWithDefaults

`func NewNotifyAckStateChangedWithDefaults() *NotifyAckStateChanged`

NewNotifyAckStateChangedWithDefaults instantiates a new NotifyAckStateChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyAckStateChanged) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyAckStateChanged) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyAckStateChanged) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyAckStateChanged) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyAckStateChanged) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyAckStateChanged) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyAckStateChanged) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyAckStateChanged) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyAckStateChanged) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyAckStateChanged) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyAckStateChanged) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyAckStateChanged) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyAckStateChanged) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyAckStateChanged) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyAckStateChanged) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetAlarmId

`func (o *NotifyAckStateChanged) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyAckStateChanged) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyAckStateChanged) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetAlarmType

`func (o *NotifyAckStateChanged) GetAlarmType() AlarmType`

GetAlarmType returns the AlarmType field if non-nil, zero value otherwise.

### GetAlarmTypeOk

`func (o *NotifyAckStateChanged) GetAlarmTypeOk() (*AlarmType, bool)`

GetAlarmTypeOk returns a tuple with the AlarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmType

`func (o *NotifyAckStateChanged) SetAlarmType(v AlarmType)`

SetAlarmType sets AlarmType field to given value.


### GetProbableCause

`func (o *NotifyAckStateChanged) GetProbableCause() ProbableCause`

GetProbableCause returns the ProbableCause field if non-nil, zero value otherwise.

### GetProbableCauseOk

`func (o *NotifyAckStateChanged) GetProbableCauseOk() (*ProbableCause, bool)`

GetProbableCauseOk returns a tuple with the ProbableCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbableCause

`func (o *NotifyAckStateChanged) SetProbableCause(v ProbableCause)`

SetProbableCause sets ProbableCause field to given value.


### GetPerceivedSeverity

`func (o *NotifyAckStateChanged) GetPerceivedSeverity() PerceivedSeverity`

GetPerceivedSeverity returns the PerceivedSeverity field if non-nil, zero value otherwise.

### GetPerceivedSeverityOk

`func (o *NotifyAckStateChanged) GetPerceivedSeverityOk() (*PerceivedSeverity, bool)`

GetPerceivedSeverityOk returns a tuple with the PerceivedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerceivedSeverity

`func (o *NotifyAckStateChanged) SetPerceivedSeverity(v PerceivedSeverity)`

SetPerceivedSeverity sets PerceivedSeverity field to given value.


### GetAckState

`func (o *NotifyAckStateChanged) GetAckState() AckState`

GetAckState returns the AckState field if non-nil, zero value otherwise.

### GetAckStateOk

`func (o *NotifyAckStateChanged) GetAckStateOk() (*AckState, bool)`

GetAckStateOk returns a tuple with the AckState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckState

`func (o *NotifyAckStateChanged) SetAckState(v AckState)`

SetAckState sets AckState field to given value.


### GetAckUserId

`func (o *NotifyAckStateChanged) GetAckUserId() string`

GetAckUserId returns the AckUserId field if non-nil, zero value otherwise.

### GetAckUserIdOk

`func (o *NotifyAckStateChanged) GetAckUserIdOk() (*string, bool)`

GetAckUserIdOk returns a tuple with the AckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUserId

`func (o *NotifyAckStateChanged) SetAckUserId(v string)`

SetAckUserId sets AckUserId field to given value.


### GetAckSystemId

`func (o *NotifyAckStateChanged) GetAckSystemId() string`

GetAckSystemId returns the AckSystemId field if non-nil, zero value otherwise.

### GetAckSystemIdOk

`func (o *NotifyAckStateChanged) GetAckSystemIdOk() (*string, bool)`

GetAckSystemIdOk returns a tuple with the AckSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSystemId

`func (o *NotifyAckStateChanged) SetAckSystemId(v string)`

SetAckSystemId sets AckSystemId field to given value.

### HasAckSystemId

`func (o *NotifyAckStateChanged) HasAckSystemId() bool`

HasAckSystemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


