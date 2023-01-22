# NotifyCorrelatedNotificationChangedAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlarmId** | **string** |  | 
**CorrelatedNotifications** | [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | 
**RootCauseIndicator** | Pointer to **bool** |  | [optional] 

## Methods

### NewNotifyCorrelatedNotificationChangedAllOf

`func NewNotifyCorrelatedNotificationChangedAllOf(alarmId string, correlatedNotifications []CorrelatedNotification, ) *NotifyCorrelatedNotificationChangedAllOf`

NewNotifyCorrelatedNotificationChangedAllOf instantiates a new NotifyCorrelatedNotificationChangedAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyCorrelatedNotificationChangedAllOfWithDefaults

`func NewNotifyCorrelatedNotificationChangedAllOfWithDefaults() *NotifyCorrelatedNotificationChangedAllOf`

NewNotifyCorrelatedNotificationChangedAllOfWithDefaults instantiates a new NotifyCorrelatedNotificationChangedAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlarmId

`func (o *NotifyCorrelatedNotificationChangedAllOf) GetAlarmId() string`

GetAlarmId returns the AlarmId field if non-nil, zero value otherwise.

### GetAlarmIdOk

`func (o *NotifyCorrelatedNotificationChangedAllOf) GetAlarmIdOk() (*string, bool)`

GetAlarmIdOk returns a tuple with the AlarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlarmId

`func (o *NotifyCorrelatedNotificationChangedAllOf) SetAlarmId(v string)`

SetAlarmId sets AlarmId field to given value.


### GetCorrelatedNotifications

`func (o *NotifyCorrelatedNotificationChangedAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyCorrelatedNotificationChangedAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyCorrelatedNotificationChangedAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.


### GetRootCauseIndicator

`func (o *NotifyCorrelatedNotificationChangedAllOf) GetRootCauseIndicator() bool`

GetRootCauseIndicator returns the RootCauseIndicator field if non-nil, zero value otherwise.

### GetRootCauseIndicatorOk

`func (o *NotifyCorrelatedNotificationChangedAllOf) GetRootCauseIndicatorOk() (*bool, bool)`

GetRootCauseIndicatorOk returns a tuple with the RootCauseIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseIndicator

`func (o *NotifyCorrelatedNotificationChangedAllOf) SetRootCauseIndicator(v bool)`

SetRootCauseIndicator sets RootCauseIndicator field to given value.

### HasRootCauseIndicator

`func (o *NotifyCorrelatedNotificationChangedAllOf) HasRootCauseIndicator() bool`

HasRootCauseIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


