# NotifyMoiCreation

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

### NewNotifyMoiCreation

`func NewNotifyMoiCreation(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, ) *NotifyMoiCreation`

NewNotifyMoiCreation instantiates a new NotifyMoiCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMoiCreationWithDefaults

`func NewNotifyMoiCreationWithDefaults() *NotifyMoiCreation`

NewNotifyMoiCreationWithDefaults instantiates a new NotifyMoiCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyMoiCreation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyMoiCreation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyMoiCreation) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyMoiCreation) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyMoiCreation) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyMoiCreation) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyMoiCreation) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyMoiCreation) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyMoiCreation) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyMoiCreation) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyMoiCreation) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyMoiCreation) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyMoiCreation) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyMoiCreation) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyMoiCreation) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetCorrelatedNotifications

`func (o *NotifyMoiCreation) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyMoiCreation) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyMoiCreation) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyMoiCreation) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyMoiCreation) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyMoiCreation) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyMoiCreation) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyMoiCreation) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetSourceIndicator

`func (o *NotifyMoiCreation) GetSourceIndicator() SourceIndicator`

GetSourceIndicator returns the SourceIndicator field if non-nil, zero value otherwise.

### GetSourceIndicatorOk

`func (o *NotifyMoiCreation) GetSourceIndicatorOk() (*SourceIndicator, bool)`

GetSourceIndicatorOk returns a tuple with the SourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndicator

`func (o *NotifyMoiCreation) SetSourceIndicator(v SourceIndicator)`

SetSourceIndicator sets SourceIndicator field to given value.

### HasSourceIndicator

`func (o *NotifyMoiCreation) HasSourceIndicator() bool`

HasSourceIndicator returns a boolean if a field has been set.

### GetAttributeList

`func (o *NotifyMoiCreation) GetAttributeList() map[string]interface{}`

GetAttributeList returns the AttributeList field if non-nil, zero value otherwise.

### GetAttributeListOk

`func (o *NotifyMoiCreation) GetAttributeListOk() (*map[string]interface{}, bool)`

GetAttributeListOk returns a tuple with the AttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeList

`func (o *NotifyMoiCreation) SetAttributeList(v map[string]interface{})`

SetAttributeList sets AttributeList field to given value.

### HasAttributeList

`func (o *NotifyMoiCreation) HasAttributeList() bool`

HasAttributeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


