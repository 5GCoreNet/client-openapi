# EventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceTypes** | Pointer to [**[]InstanceType**](InstanceType.md) |  | [optional] 
**TransProtocols** | Pointer to [**[]Protocol**](Protocol.md) |  | [optional] 
**PtpProfiles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEventFilter

`func NewEventFilter() *EventFilter`

NewEventFilter instantiates a new EventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFilterWithDefaults

`func NewEventFilterWithDefaults() *EventFilter`

NewEventFilterWithDefaults instantiates a new EventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceTypes

`func (o *EventFilter) GetInstanceTypes() []InstanceType`

GetInstanceTypes returns the InstanceTypes field if non-nil, zero value otherwise.

### GetInstanceTypesOk

`func (o *EventFilter) GetInstanceTypesOk() (*[]InstanceType, bool)`

GetInstanceTypesOk returns a tuple with the InstanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypes

`func (o *EventFilter) SetInstanceTypes(v []InstanceType)`

SetInstanceTypes sets InstanceTypes field to given value.

### HasInstanceTypes

`func (o *EventFilter) HasInstanceTypes() bool`

HasInstanceTypes returns a boolean if a field has been set.

### GetTransProtocols

`func (o *EventFilter) GetTransProtocols() []Protocol`

GetTransProtocols returns the TransProtocols field if non-nil, zero value otherwise.

### GetTransProtocolsOk

`func (o *EventFilter) GetTransProtocolsOk() (*[]Protocol, bool)`

GetTransProtocolsOk returns a tuple with the TransProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransProtocols

`func (o *EventFilter) SetTransProtocols(v []Protocol)`

SetTransProtocols sets TransProtocols field to given value.

### HasTransProtocols

`func (o *EventFilter) HasTransProtocols() bool`

HasTransProtocols returns a boolean if a field has been set.

### GetPtpProfiles

`func (o *EventFilter) GetPtpProfiles() []string`

GetPtpProfiles returns the PtpProfiles field if non-nil, zero value otherwise.

### GetPtpProfilesOk

`func (o *EventFilter) GetPtpProfilesOk() (*[]string, bool)`

GetPtpProfilesOk returns a tuple with the PtpProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpProfiles

`func (o *EventFilter) SetPtpProfiles(v []string)`

SetPtpProfiles sets PtpProfiles field to given value.

### HasPtpProfiles

`func (o *EventFilter) HasPtpProfiles() bool`

HasPtpProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


