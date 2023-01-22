# EventsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifCorreId** | **string** |  | 
**Events** | [**[]EventNotification**](EventNotification.md) |  | 

## Methods

### NewEventsNotification

`func NewEventsNotification(notifCorreId string, events []EventNotification, ) *EventsNotification`

NewEventsNotification instantiates a new EventsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsNotificationWithDefaults

`func NewEventsNotificationWithDefaults() *EventsNotification`

NewEventsNotificationWithDefaults instantiates a new EventsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifCorreId

`func (o *EventsNotification) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *EventsNotification) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *EventsNotification) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.


### GetEvents

`func (o *EventsNotification) GetEvents() []EventNotification`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsNotification) GetEventsOk() (*[]EventNotification, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsNotification) SetEvents(v []EventNotification)`

SetEvents sets Events field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


