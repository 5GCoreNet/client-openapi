# NotifyAlarmListRebuilt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**Reason** | **string** |  | 
**AlarmListAlignmentRequirement** | Pointer to [**AlarmListAlignmentRequirement**](AlarmListAlignmentRequirement.md) |  | [optional] 

## Methods

### NewNotifyAlarmListRebuilt

`func NewNotifyAlarmListRebuilt(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, reason string, ) *NotifyAlarmListRebuilt`

NewNotifyAlarmListRebuilt instantiates a new NotifyAlarmListRebuilt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyAlarmListRebuiltWithDefaults

`func NewNotifyAlarmListRebuiltWithDefaults() *NotifyAlarmListRebuilt`

NewNotifyAlarmListRebuiltWithDefaults instantiates a new NotifyAlarmListRebuilt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyAlarmListRebuilt) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyAlarmListRebuilt) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyAlarmListRebuilt) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyAlarmListRebuilt) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyAlarmListRebuilt) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyAlarmListRebuilt) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyAlarmListRebuilt) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyAlarmListRebuilt) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyAlarmListRebuilt) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyAlarmListRebuilt) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyAlarmListRebuilt) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyAlarmListRebuilt) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyAlarmListRebuilt) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyAlarmListRebuilt) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyAlarmListRebuilt) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetReason

`func (o *NotifyAlarmListRebuilt) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotifyAlarmListRebuilt) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotifyAlarmListRebuilt) SetReason(v string)`

SetReason sets Reason field to given value.


### GetAlarmListAlignmentRequirement

`func (o *NotifyAlarmListRebuilt) GetAlarmListAlignmentRequirement() AlarmListAlignmentRequirement`

GetAlarmListAlignmentRequirement returns the AlarmListAlignmentRequirement field if non-nil, zero value otherwise.

### GetAlarmListAlignmentRequirementOk

`func (o *NotifyAlarmListRebuilt) GetAlarmListAlignmentRequirementOk() (*AlarmListAlignmentRequirement, bool)`

GetAlarmListAlignmentRequirementOk returns a tuple with the AlarmListAlignmentRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmListAlignmentRequirement

`func (o *NotifyAlarmListRebuilt) SetAlarmListAlignmentRequirement(v AlarmListAlignmentRequirement)`

SetAlarmListAlignmentRequirement sets AlarmListAlignmentRequirement field to given value.

### HasAlarmListAlignmentRequirement

`func (o *NotifyAlarmListRebuilt) HasAlarmListAlignmentRequirement() bool`

HasAlarmListAlignmentRequirement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


