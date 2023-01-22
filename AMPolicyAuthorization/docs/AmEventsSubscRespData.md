# AmEventsSubscRespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Events** | Pointer to [**[]AmEventData**](AmEventData.md) |  | [optional] 
**AppAmContextId** | Pointer to **string** | Contains the AM Policy Events Subscription resource identifier related to the event notification. | [optional] 
**RepEvents** | [**[]AmEventNotification**](AmEventNotification.md) |  | 

## Methods

### NewAmEventsSubscRespData

`func NewAmEventsSubscRespData(eventNotifUri string, repEvents []AmEventNotification, ) *AmEventsSubscRespData`

NewAmEventsSubscRespData instantiates a new AmEventsSubscRespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmEventsSubscRespDataWithDefaults

`func NewAmEventsSubscRespDataWithDefaults() *AmEventsSubscRespData`

NewAmEventsSubscRespDataWithDefaults instantiates a new AmEventsSubscRespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotifUri

`func (o *AmEventsSubscRespData) GetEventNotifUri() string`

GetEventNotifUri returns the EventNotifUri field if non-nil, zero value otherwise.

### GetEventNotifUriOk

`func (o *AmEventsSubscRespData) GetEventNotifUriOk() (*string, bool)`

GetEventNotifUriOk returns a tuple with the EventNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifUri

`func (o *AmEventsSubscRespData) SetEventNotifUri(v string)`

SetEventNotifUri sets EventNotifUri field to given value.


### GetEvents

`func (o *AmEventsSubscRespData) GetEvents() []AmEventData`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AmEventsSubscRespData) GetEventsOk() (*[]AmEventData, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AmEventsSubscRespData) SetEvents(v []AmEventData)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AmEventsSubscRespData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetAppAmContextId

`func (o *AmEventsSubscRespData) GetAppAmContextId() string`

GetAppAmContextId returns the AppAmContextId field if non-nil, zero value otherwise.

### GetAppAmContextIdOk

`func (o *AmEventsSubscRespData) GetAppAmContextIdOk() (*string, bool)`

GetAppAmContextIdOk returns a tuple with the AppAmContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAmContextId

`func (o *AmEventsSubscRespData) SetAppAmContextId(v string)`

SetAppAmContextId sets AppAmContextId field to given value.

### HasAppAmContextId

`func (o *AmEventsSubscRespData) HasAppAmContextId() bool`

HasAppAmContextId returns a boolean if a field has been set.

### GetRepEvents

`func (o *AmEventsSubscRespData) GetRepEvents() []AmEventNotification`

GetRepEvents returns the RepEvents field if non-nil, zero value otherwise.

### GetRepEventsOk

`func (o *AmEventsSubscRespData) GetRepEventsOk() (*[]AmEventNotification, bool)`

GetRepEventsOk returns a tuple with the RepEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepEvents

`func (o *AmEventsSubscRespData) SetRepEvents(v []AmEventNotification)`

SetRepEvents sets RepEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


