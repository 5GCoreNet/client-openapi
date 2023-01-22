# EventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | [**SEALEvent**](SEALEvent.md) |  | 
**ValGroups** | Pointer to [**[]VALGroupFilter**](VALGroupFilter.md) | Each element of the array represents the VAL group identifier(s) of a VAL service that the subscriber wants to know in the interested event.  | [optional] 
**Identities** | Pointer to [**[]IdentityFilter**](IdentityFilter.md) | Each element of the array represents the VAL User / UE IDs of a VAL service that the event subscriber wants to know in the interested event.  | [optional] 
**MonFltr** | Pointer to [**[]MonitorFilter**](MonitorFilter.md) | List of event monitoring details that the subscriber wishes to mmonitor the VAL UEs, VAL group and/or VAL service.  | [optional] 
**AreaInt** | Pointer to [**[]MonitorLocationInterestFilter**](MonitorLocationInterestFilter.md) | Represents the list of VAL User / UE IDs and the area of interest information which the subscriber wishes to monitor the location deviation of VAL User / UEs.  | [optional] 
**LocAreaMon** | Pointer to [**[]MonLocAreaInterestFltr**](MonLocAreaInterestFltr.md) | Each element represents the location area monitoring details to monitor the VA UEs moving in and out of the provided location area.  | [optional] 

## Methods

### NewEventSubscription

`func NewEventSubscription(eventId SEALEvent, ) *EventSubscription`

NewEventSubscription instantiates a new EventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionWithDefaults

`func NewEventSubscriptionWithDefaults() *EventSubscription`

NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *EventSubscription) GetEventId() SEALEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventSubscription) GetEventIdOk() (*SEALEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventSubscription) SetEventId(v SEALEvent)`

SetEventId sets EventId field to given value.


### GetValGroups

`func (o *EventSubscription) GetValGroups() []VALGroupFilter`

GetValGroups returns the ValGroups field if non-nil, zero value otherwise.

### GetValGroupsOk

`func (o *EventSubscription) GetValGroupsOk() (*[]VALGroupFilter, bool)`

GetValGroupsOk returns a tuple with the ValGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGroups

`func (o *EventSubscription) SetValGroups(v []VALGroupFilter)`

SetValGroups sets ValGroups field to given value.

### HasValGroups

`func (o *EventSubscription) HasValGroups() bool`

HasValGroups returns a boolean if a field has been set.

### GetIdentities

`func (o *EventSubscription) GetIdentities() []IdentityFilter`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *EventSubscription) GetIdentitiesOk() (*[]IdentityFilter, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *EventSubscription) SetIdentities(v []IdentityFilter)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *EventSubscription) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### GetMonFltr

`func (o *EventSubscription) GetMonFltr() []MonitorFilter`

GetMonFltr returns the MonFltr field if non-nil, zero value otherwise.

### GetMonFltrOk

`func (o *EventSubscription) GetMonFltrOk() (*[]MonitorFilter, bool)`

GetMonFltrOk returns a tuple with the MonFltr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonFltr

`func (o *EventSubscription) SetMonFltr(v []MonitorFilter)`

SetMonFltr sets MonFltr field to given value.

### HasMonFltr

`func (o *EventSubscription) HasMonFltr() bool`

HasMonFltr returns a boolean if a field has been set.

### GetAreaInt

`func (o *EventSubscription) GetAreaInt() []MonitorLocationInterestFilter`

GetAreaInt returns the AreaInt field if non-nil, zero value otherwise.

### GetAreaIntOk

`func (o *EventSubscription) GetAreaIntOk() (*[]MonitorLocationInterestFilter, bool)`

GetAreaIntOk returns a tuple with the AreaInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaInt

`func (o *EventSubscription) SetAreaInt(v []MonitorLocationInterestFilter)`

SetAreaInt sets AreaInt field to given value.

### HasAreaInt

`func (o *EventSubscription) HasAreaInt() bool`

HasAreaInt returns a boolean if a field has been set.

### GetLocAreaMon

`func (o *EventSubscription) GetLocAreaMon() []MonLocAreaInterestFltr`

GetLocAreaMon returns the LocAreaMon field if non-nil, zero value otherwise.

### GetLocAreaMonOk

`func (o *EventSubscription) GetLocAreaMonOk() (*[]MonLocAreaInterestFltr, bool)`

GetLocAreaMonOk returns a tuple with the LocAreaMon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocAreaMon

`func (o *EventSubscription) SetLocAreaMon(v []MonLocAreaInterestFltr)`

SetLocAreaMon sets LocAreaMon field to given value.

### HasLocAreaMon

`func (o *EventSubscription) HasLocAreaMon() bool`

HasLocAreaMon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


