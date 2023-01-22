# NotifyHeartbeat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**HeartbeatNtfPeriod** | Pointer to **int32** |  | [optional] 

## Methods

### NewNotifyHeartbeat

`func NewNotifyHeartbeat(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, ) *NotifyHeartbeat`

NewNotifyHeartbeat instantiates a new NotifyHeartbeat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyHeartbeatWithDefaults

`func NewNotifyHeartbeatWithDefaults() *NotifyHeartbeat`

NewNotifyHeartbeatWithDefaults instantiates a new NotifyHeartbeat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyHeartbeat) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyHeartbeat) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyHeartbeat) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyHeartbeat) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyHeartbeat) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyHeartbeat) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyHeartbeat) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyHeartbeat) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyHeartbeat) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyHeartbeat) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyHeartbeat) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyHeartbeat) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyHeartbeat) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyHeartbeat) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyHeartbeat) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetHeartbeatNtfPeriod

`func (o *NotifyHeartbeat) GetHeartbeatNtfPeriod() int32`

GetHeartbeatNtfPeriod returns the HeartbeatNtfPeriod field if non-nil, zero value otherwise.

### GetHeartbeatNtfPeriodOk

`func (o *NotifyHeartbeat) GetHeartbeatNtfPeriodOk() (*int32, bool)`

GetHeartbeatNtfPeriodOk returns a tuple with the HeartbeatNtfPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatNtfPeriod

`func (o *NotifyHeartbeat) SetHeartbeatNtfPeriod(v int32)`

SetHeartbeatNtfPeriod sets HeartbeatNtfPeriod field to given value.

### HasHeartbeatNtfPeriod

`func (o *NotifyHeartbeat) HasHeartbeatNtfPeriod() bool`

HasHeartbeatNtfPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


