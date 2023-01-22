# NotifyFilePreparationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**FileInfoList** | Pointer to [**[]FileInfo**](FileInfo.md) |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyFilePreparationError

`func NewNotifyFilePreparationError(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, ) *NotifyFilePreparationError`

NewNotifyFilePreparationError instantiates a new NotifyFilePreparationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyFilePreparationErrorWithDefaults

`func NewNotifyFilePreparationErrorWithDefaults() *NotifyFilePreparationError`

NewNotifyFilePreparationErrorWithDefaults instantiates a new NotifyFilePreparationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyFilePreparationError) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyFilePreparationError) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyFilePreparationError) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyFilePreparationError) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyFilePreparationError) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyFilePreparationError) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyFilePreparationError) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyFilePreparationError) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyFilePreparationError) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyFilePreparationError) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyFilePreparationError) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyFilePreparationError) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyFilePreparationError) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyFilePreparationError) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyFilePreparationError) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetFileInfoList

`func (o *NotifyFilePreparationError) GetFileInfoList() []FileInfo`

GetFileInfoList returns the FileInfoList field if non-nil, zero value otherwise.

### GetFileInfoListOk

`func (o *NotifyFilePreparationError) GetFileInfoListOk() (*[]FileInfo, bool)`

GetFileInfoListOk returns a tuple with the FileInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfoList

`func (o *NotifyFilePreparationError) SetFileInfoList(v []FileInfo)`

SetFileInfoList sets FileInfoList field to given value.

### HasFileInfoList

`func (o *NotifyFilePreparationError) HasFileInfoList() bool`

HasFileInfoList returns a boolean if a field has been set.

### GetReason

`func (o *NotifyFilePreparationError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *NotifyFilePreparationError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *NotifyFilePreparationError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *NotifyFilePreparationError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyFilePreparationError) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyFilePreparationError) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyFilePreparationError) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyFilePreparationError) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


