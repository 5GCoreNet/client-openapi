# NotifyMoiChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**MoiChanges** | [**[]MoiChange**](MoiChange.md) |  | 

## Methods

### NewNotifyMoiChanges

`func NewNotifyMoiChanges(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, moiChanges []MoiChange, ) *NotifyMoiChanges`

NewNotifyMoiChanges instantiates a new NotifyMoiChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMoiChangesWithDefaults

`func NewNotifyMoiChangesWithDefaults() *NotifyMoiChanges`

NewNotifyMoiChangesWithDefaults instantiates a new NotifyMoiChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyMoiChanges) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyMoiChanges) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyMoiChanges) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyMoiChanges) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyMoiChanges) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyMoiChanges) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyMoiChanges) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyMoiChanges) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyMoiChanges) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyMoiChanges) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyMoiChanges) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyMoiChanges) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyMoiChanges) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyMoiChanges) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyMoiChanges) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetMoiChanges

`func (o *NotifyMoiChanges) GetMoiChanges() []MoiChange`

GetMoiChanges returns the MoiChanges field if non-nil, zero value otherwise.

### GetMoiChangesOk

`func (o *NotifyMoiChanges) GetMoiChangesOk() (*[]MoiChange, bool)`

GetMoiChangesOk returns a tuple with the MoiChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoiChanges

`func (o *NotifyMoiChanges) SetMoiChanges(v []MoiChange)`

SetMoiChanges sets MoiChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


