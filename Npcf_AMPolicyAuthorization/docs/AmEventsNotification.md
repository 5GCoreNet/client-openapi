# AmEventsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAmContextId** | Pointer to **string** | Contains the AM Policy Events Subscription resource identifier related to the event notification. | [optional] 
**RepEvents** | [**[]AmEventNotification**](AmEventNotification.md) |  | 

## Methods

### NewAmEventsNotification

`func NewAmEventsNotification(repEvents []AmEventNotification, ) *AmEventsNotification`

NewAmEventsNotification instantiates a new AmEventsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmEventsNotificationWithDefaults

`func NewAmEventsNotificationWithDefaults() *AmEventsNotification`

NewAmEventsNotificationWithDefaults instantiates a new AmEventsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAmContextId

`func (o *AmEventsNotification) GetAppAmContextId() string`

GetAppAmContextId returns the AppAmContextId field if non-nil, zero value otherwise.

### GetAppAmContextIdOk

`func (o *AmEventsNotification) GetAppAmContextIdOk() (*string, bool)`

GetAppAmContextIdOk returns a tuple with the AppAmContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAmContextId

`func (o *AmEventsNotification) SetAppAmContextId(v string)`

SetAppAmContextId sets AppAmContextId field to given value.

### HasAppAmContextId

`func (o *AmEventsNotification) HasAppAmContextId() bool`

HasAppAmContextId returns a boolean if a field has been set.

### GetRepEvents

`func (o *AmEventsNotification) GetRepEvents() []AmEventNotification`

GetRepEvents returns the RepEvents field if non-nil, zero value otherwise.

### GetRepEventsOk

`func (o *AmEventsNotification) GetRepEventsOk() (*[]AmEventNotification, bool)`

GetRepEventsOk returns a tuple with the RepEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepEvents

`func (o *AmEventsNotification) SetRepEvents(v []AmEventNotification)`

SetRepEvents sets RepEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


