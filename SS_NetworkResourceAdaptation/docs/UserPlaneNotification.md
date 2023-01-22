# UserPlaneNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifId** | **string** | String providing an URI formatted according to RFC 3986. | 
**EventNotifs** | [**[]NrmEventNotification**](NrmEventNotification.md) |  | 

## Methods

### NewUserPlaneNotification

`func NewUserPlaneNotification(notifId string, eventNotifs []NrmEventNotification, ) *UserPlaneNotification`

NewUserPlaneNotification instantiates a new UserPlaneNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlaneNotificationWithDefaults

`func NewUserPlaneNotificationWithDefaults() *UserPlaneNotification`

NewUserPlaneNotificationWithDefaults instantiates a new UserPlaneNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifId

`func (o *UserPlaneNotification) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *UserPlaneNotification) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *UserPlaneNotification) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetEventNotifs

`func (o *UserPlaneNotification) GetEventNotifs() []NrmEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *UserPlaneNotification) GetEventNotifsOk() (*[]NrmEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *UserPlaneNotification) SetEventNotifs(v []NrmEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


