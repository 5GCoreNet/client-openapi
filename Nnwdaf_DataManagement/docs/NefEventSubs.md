# NefEventSubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**NefEvent**](NefEvent.md) |  | 
**EventFilter** | Pointer to [**NefEventFilter**](NefEventFilter.md) |  | [optional] 

## Methods

### NewNefEventSubs

`func NewNefEventSubs(event NefEvent, ) *NefEventSubs`

NewNefEventSubs instantiates a new NefEventSubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefEventSubsWithDefaults

`func NewNefEventSubsWithDefaults() *NefEventSubs`

NewNefEventSubsWithDefaults instantiates a new NefEventSubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *NefEventSubs) GetEvent() NefEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *NefEventSubs) GetEventOk() (*NefEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *NefEventSubs) SetEvent(v NefEvent)`

SetEvent sets Event field to given value.


### GetEventFilter

`func (o *NefEventSubs) GetEventFilter() NefEventFilter`

GetEventFilter returns the EventFilter field if non-nil, zero value otherwise.

### GetEventFilterOk

`func (o *NefEventSubs) GetEventFilterOk() (*NefEventFilter, bool)`

GetEventFilterOk returns a tuple with the EventFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilter

`func (o *NefEventSubs) SetEventFilter(v NefEventFilter)`

SetEventFilter sets EventFilter field to given value.

### HasEventFilter

`func (o *NefEventSubs) HasEventFilter() bool`

HasEventFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


