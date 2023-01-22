# NotificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** |  | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**UpBufferInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNotificationInfo

`func NewNotificationInfo(notifId string, notifUri string, ) *NotificationInfo`

NewNotificationInfo instantiates a new NotificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationInfoWithDefaults

`func NewNotificationInfoWithDefaults() *NotificationInfo`

NewNotificationInfoWithDefaults instantiates a new NotificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *NotificationInfo) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *NotificationInfo) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *NotificationInfo) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetNotifUri

`func (o *NotificationInfo) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *NotificationInfo) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *NotificationInfo) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetUpBufferInd

`func (o *NotificationInfo) GetUpBufferInd() bool`

GetUpBufferInd returns the UpBufferInd field if non-nil, zero value otherwise.

### GetUpBufferIndOk

`func (o *NotificationInfo) GetUpBufferIndOk() (*bool, bool)`

GetUpBufferIndOk returns a tuple with the UpBufferInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpBufferInd

`func (o *NotificationInfo) SetUpBufferInd(v bool)`

SetUpBufferInd sets UpBufferInd field to given value.

### HasUpBufferInd

`func (o *NotificationInfo) HasUpBufferInd() bool`

HasUpBufferInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


