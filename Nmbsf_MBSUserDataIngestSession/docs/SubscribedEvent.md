# SubscribedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusEvent** | [**Event**](Event.md) |  | 
**MbsDistSessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewSubscribedEvent

`func NewSubscribedEvent(statusEvent Event, ) *SubscribedEvent`

NewSubscribedEvent instantiates a new SubscribedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribedEventWithDefaults

`func NewSubscribedEventWithDefaults() *SubscribedEvent`

NewSubscribedEventWithDefaults instantiates a new SubscribedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusEvent

`func (o *SubscribedEvent) GetStatusEvent() Event`

GetStatusEvent returns the StatusEvent field if non-nil, zero value otherwise.

### GetStatusEventOk

`func (o *SubscribedEvent) GetStatusEventOk() (*Event, bool)`

GetStatusEventOk returns a tuple with the StatusEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEvent

`func (o *SubscribedEvent) SetStatusEvent(v Event)`

SetStatusEvent sets StatusEvent field to given value.


### GetMbsDistSessionId

`func (o *SubscribedEvent) GetMbsDistSessionId() string`

GetMbsDistSessionId returns the MbsDistSessionId field if non-nil, zero value otherwise.

### GetMbsDistSessionIdOk

`func (o *SubscribedEvent) GetMbsDistSessionIdOk() (*string, bool)`

GetMbsDistSessionIdOk returns a tuple with the MbsDistSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDistSessionId

`func (o *SubscribedEvent) SetMbsDistSessionId(v string)`

SetMbsDistSessionId sets MbsDistSessionId field to given value.

### HasMbsDistSessionId

`func (o *SubscribedEvent) HasMbsDistSessionId() bool`

HasMbsDistSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


