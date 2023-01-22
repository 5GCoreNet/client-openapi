# EventsSubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**AfEvent**](AfEvent.md) |  | 
**EventFilter** | [**EventFilter**](EventFilter.md) |  | 

## Methods

### NewEventsSubs

`func NewEventsSubs(event AfEvent, eventFilter EventFilter, ) *EventsSubs`

NewEventsSubs instantiates a new EventsSubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSubsWithDefaults

`func NewEventsSubsWithDefaults() *EventsSubs`

NewEventsSubsWithDefaults instantiates a new EventsSubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventsSubs) GetEvent() AfEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventsSubs) GetEventOk() (*AfEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventsSubs) SetEvent(v AfEvent)`

SetEvent sets Event field to given value.


### GetEventFilter

`func (o *EventsSubs) GetEventFilter() EventFilter`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *EventsSubs) GetEventFilterOk() (*EventFilter, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *EventsSubs) SetEventFilter(v EventFilter)`

SetEventFilter sets EventFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


