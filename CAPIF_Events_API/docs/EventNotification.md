# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | Identifier of the subscription resource to which the notification is related – CAPIF resource identifier  | 
**Events** | [**CAPIFEvent**](CAPIFEvent.md) |  | 
**EventDetail** | Pointer to [**CAPIFEventDetail**](CAPIFEventDetail.md) |  | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification(subscriptionId string, events CAPIFEvent, ) *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *EventNotification) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *EventNotification) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *EventNotification) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetEvents

`func (o *EventNotification) GetEvents() CAPIFEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventNotification) GetEventsOk() (*CAPIFEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventNotification) SetEvents(v CAPIFEvent)`

SetEvents sets Events field to given value.


### GetEventDetail

`func (o *EventNotification) GetEventDetail() CAPIFEventDetail`

GetEventDetail returns the EventDetail field if non-nil, zero value otherwise.

### GetEventDetailOk

`func (o *EventNotification) GetEventDetailOk() (*CAPIFEventDetail, bool)`

GetEventDetailOk returns a tuple with the EventDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetail

`func (o *EventNotification) SetEventDetail(v CAPIFEventDetail)`

SetEventDetail sets EventDetail field to given value.

### HasEventDetail

`func (o *EventNotification) HasEventDetail() bool`

HasEventDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


