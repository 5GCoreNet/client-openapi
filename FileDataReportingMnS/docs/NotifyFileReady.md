# NotifyFileReady

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**NotificationId** | **int32** |  | 
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**EventTime** | **time.Time** |  | 
**SystemDN** | **string** |  | 
**FileInfoList** | Pointer to [**[]FileInfo**](FileInfo.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 

## Methods

### NewNotifyFileReady

`func NewNotifyFileReady(href string, notificationId int32, notificationType NotificationType, eventTime time.Time, systemDN string, ) *NotifyFileReady`

NewNotifyFileReady instantiates a new NotifyFileReady object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyFileReadyWithDefaults

`func NewNotifyFileReadyWithDefaults() *NotifyFileReady`

NewNotifyFileReadyWithDefaults instantiates a new NotifyFileReady object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *NotifyFileReady) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *NotifyFileReady) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *NotifyFileReady) SetHref(v string)`

SetHref sets Href field to given value.


### GetNotificationId

`func (o *NotifyFileReady) GetNotificationId() int32`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *NotifyFileReady) GetNotificationIdOk() (*int32, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *NotifyFileReady) SetNotificationId(v int32)`

SetNotificationId sets NotificationId field to given value.


### GetNotificationType

`func (o *NotifyFileReady) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *NotifyFileReady) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *NotifyFileReady) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetEventTime

`func (o *NotifyFileReady) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *NotifyFileReady) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *NotifyFileReady) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSystemDN

`func (o *NotifyFileReady) GetSystemDN() string`

GetSystemDN returns the SystemDN field if non-nil, zero value otherwise.

### GetSystemDNOk

`func (o *NotifyFileReady) GetSystemDNOk() (*string, bool)`

GetSystemDNOk returns a tuple with the SystemDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDN

`func (o *NotifyFileReady) SetSystemDN(v string)`

SetSystemDN sets SystemDN field to given value.


### GetFileInfoList

`func (o *NotifyFileReady) GetFileInfoList() []FileInfo`

GetFileInfoList returns the FileInfoList field if non-nil, zero value otherwise.

### GetFileInfoListOk

`func (o *NotifyFileReady) GetFileInfoListOk() (*[]FileInfo, bool)`

GetFileInfoListOk returns a tuple with the FileInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileInfoList

`func (o *NotifyFileReady) SetFileInfoList(v []FileInfo)`

SetFileInfoList sets FileInfoList field to given value.

### HasFileInfoList

`func (o *NotifyFileReady) HasFileInfoList() bool`

HasFileInfoList returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyFileReady) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyFileReady) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyFileReady) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyFileReady) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


