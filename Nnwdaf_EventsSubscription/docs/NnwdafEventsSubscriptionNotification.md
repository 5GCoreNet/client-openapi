# NnwdafEventsSubscriptionNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotifications** | Pointer to [**[]EventNotification**](EventNotification.md) | Notifications about Individual Events | [optional] 
**SubscriptionId** | **string** | String identifying a subscription to the Nnwdaf_EventsSubscription Service | 
**NotifCorrId** | Pointer to **string** | Notification correlation identifier. | [optional] 
**OldSubscriptionId** | Pointer to **string** | Subscription ID which was allocated by the source NWDAF. This parameter shall be present if the notification is for informing the assignment of a new Subscription Id by the target NWDAF.  | [optional] 

## Methods

### NewNnwdafEventsSubscriptionNotification

`func NewNnwdafEventsSubscriptionNotification(subscriptionId string, ) *NnwdafEventsSubscriptionNotification`

NewNnwdafEventsSubscriptionNotification instantiates a new NnwdafEventsSubscriptionNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNnwdafEventsSubscriptionNotificationWithDefaults

`func NewNnwdafEventsSubscriptionNotificationWithDefaults() *NnwdafEventsSubscriptionNotification`

NewNnwdafEventsSubscriptionNotificationWithDefaults instantiates a new NnwdafEventsSubscriptionNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotifications

`func (o *NnwdafEventsSubscriptionNotification) GetEventNotifications() []EventNotification`

GetEventNotifications returns the EventNotifications field if non-nil, zero value otherwise.

### GetEventNotificationsOk

`func (o *NnwdafEventsSubscriptionNotification) GetEventNotificationsOk() (*[]EventNotification, bool)`

GetEventNotificationsOk returns a tuple with the EventNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifications

`func (o *NnwdafEventsSubscriptionNotification) SetEventNotifications(v []EventNotification)`

SetEventNotifications sets EventNotifications field to given value.

### HasEventNotifications

`func (o *NnwdafEventsSubscriptionNotification) HasEventNotifications() bool`

HasEventNotifications returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *NnwdafEventsSubscriptionNotification) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetNotifCorrId

`func (o *NnwdafEventsSubscriptionNotification) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *NnwdafEventsSubscriptionNotification) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *NnwdafEventsSubscriptionNotification) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.

### HasNotifCorrId

`func (o *NnwdafEventsSubscriptionNotification) HasNotifCorrId() bool`

HasNotifCorrId returns a boolean if a field has been set.

### GetOldSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) GetOldSubscriptionId() string`

GetOldSubscriptionId returns the OldSubscriptionId field if non-nil, zero value otherwise.

### GetOldSubscriptionIdOk

`func (o *NnwdafEventsSubscriptionNotification) GetOldSubscriptionIdOk() (*string, bool)`

GetOldSubscriptionIdOk returns a tuple with the OldSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) SetOldSubscriptionId(v string)`

SetOldSubscriptionId sets OldSubscriptionId field to given value.

### HasOldSubscriptionId

`func (o *NnwdafEventsSubscriptionNotification) HasOldSubscriptionId() bool`

HasOldSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


