# NotifyMoiDeletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**SourceIndicator** | Pointer to [**SourceIndicator**](SourceIndicator.md) |  | [optional] 
**AttributeList** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 

## Methods

### NewNotifyMoiDeletion

`func NewNotifyMoiDeletion(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, ) *NotifyMoiDeletion`

NewNotifyMoiDeletion instantiates a new NotifyMoiDeletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMoiDeletionWithDefaults

`func NewNotifyMoiDeletionWithDefaults() *NotifyMoiDeletion`

NewNotifyMoiDeletionWithDefaults instantiates a new NotifyMoiDeletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyMoiDeletion) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyMoiDeletion) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyMoiDeletion) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyMoiDeletion) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyMoiDeletion) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyMoiDeletion) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyMoiDeletion) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyMoiDeletion) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyMoiDeletion) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyMoiDeletion) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyMoiDeletion) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyMoiDeletion) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyMoiDeletion) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyMoiDeletion) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyMoiDeletion) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetCorrelatedNotifications

`func (o *NotifyMoiDeletion) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyMoiDeletion) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyMoiDeletion) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyMoiDeletion) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyMoiDeletion) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyMoiDeletion) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyMoiDeletion) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyMoiDeletion) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetSourceIndicator

`func (o *NotifyMoiDeletion) GetSourceIndicator() SourceIndicator`

GetSourceIndicator returns the SourceIndicator field if non-nil, zero value otherwise.

### GetSourceIndicatorOk

`func (o *NotifyMoiDeletion) GetSourceIndicatorOk() (*SourceIndicator, bool)`

GetSourceIndicatorOk returns a tuple with the SourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndicator

`func (o *NotifyMoiDeletion) SetSourceIndicator(v SourceIndicator)`

SetSourceIndicator sets SourceIndicator field to given value.

### HasSourceIndicator

`func (o *NotifyMoiDeletion) HasSourceIndicator() bool`

HasSourceIndicator returns a boolean if a field has been set.

### GetAttributeList

`func (o *NotifyMoiDeletion) GetAttributeList() map[string]interface{}`

GetAttributeList returns the AttributeList field if non-nil, zero value otherwise.

### GetAttributeListOk

`func (o *NotifyMoiDeletion) GetAttributeListOk() (*map[string]interface{}, bool)`

GetAttributeListOk returns a tuple with the AttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeList

`func (o *NotifyMoiDeletion) SetAttributeList(v map[string]interface{})`

SetAttributeList sets AttributeList field to given value.

### HasAttributeList

`func (o *NotifyMoiDeletion) HasAttributeList() bool`

HasAttributeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


