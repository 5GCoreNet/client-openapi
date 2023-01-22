# AmEventsSubscData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Events** | Pointer to [**[]AmEventData**](AmEventData.md) |  | [optional] 

## Methods

### NewAmEventsSubscData

`func NewAmEventsSubscData(eventNotifUri string, ) *AmEventsSubscData`

NewAmEventsSubscData instantiates a new AmEventsSubscData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmEventsSubscDataWithDefaults

`func NewAmEventsSubscDataWithDefaults() *AmEventsSubscData`

NewAmEventsSubscDataWithDefaults instantiates a new AmEventsSubscData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventNotifUri

`func (o *AmEventsSubscData) GetEventNotifUri() string`

GetEventNotifUri returns the EventNotifUri field if non-nil, zero value otherwise.

### GetEventNotifUriOk

`func (o *AmEventsSubscData) GetEventNotifUriOk() (*string, bool)`

GetEventNotifUriOk returns a tuple with the EventNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifUri

`func (o *AmEventsSubscData) SetEventNotifUri(v string)`

SetEventNotifUri sets EventNotifUri field to given value.


### GetEvents

`func (o *AmEventsSubscData) GetEvents() []AmEventData`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AmEventsSubscData) GetEventsOk() (*[]AmEventData, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AmEventsSubscData) SetEvents(v []AmEventData)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AmEventsSubscData) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


