# NotifyMoiAttributeValueChanges

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
**AttributeListValueChanges** | **[]map[string]interface{}** | The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values. | 

## Methods

### NewNotifyMoiAttributeValueChanges

`func NewNotifyMoiAttributeValueChanges(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, attributeListValueChanges []map[string]interface{}, ) *NotifyMoiAttributeValueChanges`

NewNotifyMoiAttributeValueChanges instantiates a new NotifyMoiAttributeValueChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMoiAttributeValueChangesWithDefaults

`func NewNotifyMoiAttributeValueChangesWithDefaults() *NotifyMoiAttributeValueChanges`

NewNotifyMoiAttributeValueChangesWithDefaults instantiates a new NotifyMoiAttributeValueChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyMoiAttributeValueChanges) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyMoiAttributeValueChanges) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyMoiAttributeValueChanges) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyMoiAttributeValueChanges) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyMoiAttributeValueChanges) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyMoiAttributeValueChanges) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyMoiAttributeValueChanges) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyMoiAttributeValueChanges) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyMoiAttributeValueChanges) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyMoiAttributeValueChanges) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyMoiAttributeValueChanges) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyMoiAttributeValueChanges) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyMoiAttributeValueChanges) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyMoiAttributeValueChanges) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyMoiAttributeValueChanges) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetCorrelatedNotifications

`func (o *NotifyMoiAttributeValueChanges) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyMoiAttributeValueChanges) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyMoiAttributeValueChanges) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyMoiAttributeValueChanges) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyMoiAttributeValueChanges) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyMoiAttributeValueChanges) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyMoiAttributeValueChanges) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyMoiAttributeValueChanges) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetSourceIndicator

`func (o *NotifyMoiAttributeValueChanges) GetSourceIndicator() SourceIndicator`

GetSourceIndicator returns the SourceIndicator field if non-nil, zero value otherwise.

### GetSourceIndicatorOk

`func (o *NotifyMoiAttributeValueChanges) GetSourceIndicatorOk() (*SourceIndicator, bool)`

GetSourceIndicatorOk returns a tuple with the SourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndicator

`func (o *NotifyMoiAttributeValueChanges) SetSourceIndicator(v SourceIndicator)`

SetSourceIndicator sets SourceIndicator field to given value.

### HasSourceIndicator

`func (o *NotifyMoiAttributeValueChanges) HasSourceIndicator() bool`

HasSourceIndicator returns a boolean if a field has been set.

### GetAttributeListValueChanges

`func (o *NotifyMoiAttributeValueChanges) GetAttributeListValueChanges() []map[string]interface{}`

GetAttributeListValueChanges returns the AttributeListValueChanges field if non-nil, zero value otherwise.

### GetAttributeListValueChangesOk

`func (o *NotifyMoiAttributeValueChanges) GetAttributeListValueChangesOk() (*[]map[string]interface{}, bool)`

GetAttributeListValueChangesOk returns a tuple with the AttributeListValueChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeListValueChanges

`func (o *NotifyMoiAttributeValueChanges) SetAttributeListValueChanges(v []map[string]interface{})`

SetAttributeListValueChanges sets AttributeListValueChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


