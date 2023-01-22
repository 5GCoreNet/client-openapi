# NmfafDataAnaNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnaNotifications** | Pointer to [**[]NnwdafEventsSubscriptionNotification**](NnwdafEventsSubscriptionNotification.md) | List of analytics subscription notifications. | [optional] 
**DataNotif** | Pointer to [**DataNotification**](DataNotification.md) |  | [optional] 

## Methods

### NewNmfafDataAnaNotification

`func NewNmfafDataAnaNotification() *NmfafDataAnaNotification`

NewNmfafDataAnaNotification instantiates a new NmfafDataAnaNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNmfafDataAnaNotificationWithDefaults

`func NewNmfafDataAnaNotificationWithDefaults() *NmfafDataAnaNotification`

NewNmfafDataAnaNotificationWithDefaults instantiates a new NmfafDataAnaNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaNotifications

`func (o *NmfafDataAnaNotification) GetAnaNotifications() []NnwdafEventsSubscriptionNotification`

GetAnaNotifications returns the AnaNotifications field if non-nil, zero value otherwise.

### GetAnaNotificationsOk

`func (o *NmfafDataAnaNotification) GetAnaNotificationsOk() (*[]NnwdafEventsSubscriptionNotification, bool)`

GetAnaNotificationsOk returns a tuple with the AnaNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifications

`func (o *NmfafDataAnaNotification) SetAnaNotifications(v []NnwdafEventsSubscriptionNotification)`

SetAnaNotifications sets AnaNotifications field to given value.

### HasAnaNotifications

`func (o *NmfafDataAnaNotification) HasAnaNotifications() bool`

HasAnaNotifications returns a boolean if a field has been set.

### GetDataNotif

`func (o *NmfafDataAnaNotification) GetDataNotif() DataNotification`

GetDataNotif returns the DataNotif field if non-nil, zero value otherwise.

### GetDataNotifOk

`func (o *NmfafDataAnaNotification) GetDataNotifOk() (*DataNotification, bool)`

GetDataNotifOk returns a tuple with the DataNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNotif

`func (o *NmfafDataAnaNotification) SetDataNotif(v DataNotification)`

SetDataNotif sets DataNotif field to given value.

### HasDataNotif

`func (o *NmfafDataAnaNotification) HasDataNotif() bool`

HasDataNotif returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


