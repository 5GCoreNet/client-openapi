# AssignEbiFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**FailedArpList** | Pointer to [**[]Arp**](Arp.md) |  | [optional] 

## Methods

### NewAssignEbiFailed

`func NewAssignEbiFailed(pduSessionId int32, ) *AssignEbiFailed`

NewAssignEbiFailed instantiates a new AssignEbiFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignEbiFailedWithDefaults

`func NewAssignEbiFailedWithDefaults() *AssignEbiFailed`

NewAssignEbiFailedWithDefaults instantiates a new AssignEbiFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *AssignEbiFailed) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *AssignEbiFailed) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *AssignEbiFailed) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetFailedArpList

`func (o *AssignEbiFailed) GetFailedArpList() []Arp`

GetFailedArpList returns the FailedArpList field if non-nil, zero value otherwise.

### GetFailedArpListOk

`func (o *AssignEbiFailed) GetFailedArpListOk() (*[]Arp, bool)`

GetFailedArpListOk returns a tuple with the FailedArpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedArpList

`func (o *AssignEbiFailed) SetFailedArpList(v []Arp)`

SetFailedArpList sets FailedArpList field to given value.

### HasFailedArpList

`func (o *AssignEbiFailed) HasFailedArpList() bool`

HasFailedArpList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


