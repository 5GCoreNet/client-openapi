# DccfEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwdafEvent** | Pointer to [**NwdafEvent**](NwdafEvent.md) |  | [optional] 
**SmfEvent** | Pointer to [**SmfEvent**](SmfEvent.md) |  | [optional] 
**AmfEvent** | Pointer to [**AmfEventType**](AmfEventType.md) |  | [optional] 
**NefEvent** | Pointer to [**NefEvent**](NefEvent.md) |  | [optional] 
**UdmEvent** | Pointer to [**EventType**](EventType.md) |  | [optional] 
**AfEvent** | Pointer to [**AfEvent**](AfEvent.md) |  | [optional] 
**SacEvent** | Pointer to [**SACEvent**](SACEvent.md) |  | [optional] 
**NrfEvent** | Pointer to [**NotificationEventType**](NotificationEventType.md) |  | [optional] 

## Methods

### NewDccfEvent

`func NewDccfEvent() *DccfEvent`

NewDccfEvent instantiates a new DccfEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDccfEventWithDefaults

`func NewDccfEventWithDefaults() *DccfEvent`

NewDccfEventWithDefaults instantiates a new DccfEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwdafEvent

`func (o *DccfEvent) GetNwdafEvent() NwdafEvent`

GetNwdafEvent returns the NwdafEvent field if non-nil, zero value otherwise.

### GetNwdafEventOk

`func (o *DccfEvent) GetNwdafEventOk() (*NwdafEvent, bool)`

GetNwdafEventOk returns a tuple with the NwdafEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafEvent

`func (o *DccfEvent) SetNwdafEvent(v NwdafEvent)`

SetNwdafEvent sets NwdafEvent field to given value.

### HasNwdafEvent

`func (o *DccfEvent) HasNwdafEvent() bool`

HasNwdafEvent returns a boolean if a field has been set.

### GetSmfEvent

`func (o *DccfEvent) GetSmfEvent() SmfEvent`

GetSmfEvent returns the SmfEvent field if non-nil, zero value otherwise.

### GetSmfEventOk

`func (o *DccfEvent) GetSmfEventOk() (*SmfEvent, bool)`

GetSmfEventOk returns a tuple with the SmfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfEvent

`func (o *DccfEvent) SetSmfEvent(v SmfEvent)`

SetSmfEvent sets SmfEvent field to given value.

### HasSmfEvent

`func (o *DccfEvent) HasSmfEvent() bool`

HasSmfEvent returns a boolean if a field has been set.

### GetAmfEvent

`func (o *DccfEvent) GetAmfEvent() AmfEventType`

GetAmfEvent returns the AmfEvent field if non-nil, zero value otherwise.

### GetAmfEventOk

`func (o *DccfEvent) GetAmfEventOk() (*AmfEventType, bool)`

GetAmfEventOk returns a tuple with the AmfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfEvent

`func (o *DccfEvent) SetAmfEvent(v AmfEventType)`

SetAmfEvent sets AmfEvent field to given value.

### HasAmfEvent

`func (o *DccfEvent) HasAmfEvent() bool`

HasAmfEvent returns a boolean if a field has been set.

### GetNefEvent

`func (o *DccfEvent) GetNefEvent() NefEvent`

GetNefEvent returns the NefEvent field if non-nil, zero value otherwise.

### GetNefEventOk

`func (o *DccfEvent) GetNefEventOk() (*NefEvent, bool)`

GetNefEventOk returns a tuple with the NefEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefEvent

`func (o *DccfEvent) SetNefEvent(v NefEvent)`

SetNefEvent sets NefEvent field to given value.

### HasNefEvent

`func (o *DccfEvent) HasNefEvent() bool`

HasNefEvent returns a boolean if a field has been set.

### GetUdmEvent

`func (o *DccfEvent) GetUdmEvent() EventType`

GetUdmEvent returns the UdmEvent field if non-nil, zero value otherwise.

### GetUdmEventOk

`func (o *DccfEvent) GetUdmEventOk() (*EventType, bool)`

GetUdmEventOk returns a tuple with the UdmEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmEvent

`func (o *DccfEvent) SetUdmEvent(v EventType)`

SetUdmEvent sets UdmEvent field to given value.

### HasUdmEvent

`func (o *DccfEvent) HasUdmEvent() bool`

HasUdmEvent returns a boolean if a field has been set.

### GetAfEvent

`func (o *DccfEvent) GetAfEvent() AfEvent`

GetAfEvent returns the AfEvent field if non-nil, zero value otherwise.

### GetAfEventOk

`func (o *DccfEvent) GetAfEventOk() (*AfEvent, bool)`

GetAfEventOk returns a tuple with the AfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEvent

`func (o *DccfEvent) SetAfEvent(v AfEvent)`

SetAfEvent sets AfEvent field to given value.

### HasAfEvent

`func (o *DccfEvent) HasAfEvent() bool`

HasAfEvent returns a boolean if a field has been set.

### GetSacEvent

`func (o *DccfEvent) GetSacEvent() SACEvent`

GetSacEvent returns the SacEvent field if non-nil, zero value otherwise.

### GetSacEventOk

`func (o *DccfEvent) GetSacEventOk() (*SACEvent, bool)`

GetSacEventOk returns a tuple with the SacEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSacEvent

`func (o *DccfEvent) SetSacEvent(v SACEvent)`

SetSacEvent sets SacEvent field to given value.

### HasSacEvent

`func (o *DccfEvent) HasSacEvent() bool`

HasSacEvent returns a boolean if a field has been set.

### GetNrfEvent

`func (o *DccfEvent) GetNrfEvent() NotificationEventType`

GetNrfEvent returns the NrfEvent field if non-nil, zero value otherwise.

### GetNrfEventOk

`func (o *DccfEvent) GetNrfEventOk() (*NotificationEventType, bool)`

GetNrfEventOk returns a tuple with the NrfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfEvent

`func (o *DccfEvent) SetNrfEvent(v NotificationEventType)`

SetNrfEvent sets NrfEvent field to given value.

### HasNrfEvent

`func (o *DccfEvent) HasNrfEvent() bool`

HasNrfEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


