# C2LinkQualityThrlds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NrRsrpThrldLow** | Pointer to **int32** |  | [optional] 
**NrRsrpThrldHigh** | Pointer to **int32** |  | [optional] 
**NrRsrqThrldLow** | Pointer to **int32** |  | [optional] 
**NrRsrqThrldHigh** | Pointer to **int32** |  | [optional] 
**PacketLossThrldLow** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 
**PacketLossThrldHigh** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 

## Methods

### NewC2LinkQualityThrlds

`func NewC2LinkQualityThrlds() *C2LinkQualityThrlds`

NewC2LinkQualityThrlds instantiates a new C2LinkQualityThrlds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC2LinkQualityThrldsWithDefaults

`func NewC2LinkQualityThrldsWithDefaults() *C2LinkQualityThrlds`

NewC2LinkQualityThrldsWithDefaults instantiates a new C2LinkQualityThrlds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNrRsrpThrldLow

`func (o *C2LinkQualityThrlds) GetNrRsrpThrldLow() int32`

GetNrRsrpThrldLow returns the NrRsrpThrldLow field if non-nil, zero value otherwise.

### GetNrRsrpThrldLowOk

`func (o *C2LinkQualityThrlds) GetNrRsrpThrldLowOk() (*int32, bool)`

GetNrRsrpThrldLowOk returns a tuple with the NrRsrpThrldLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrRsrpThrldLow

`func (o *C2LinkQualityThrlds) SetNrRsrpThrldLow(v int32)`

SetNrRsrpThrldLow sets NrRsrpThrldLow field to given value.

### HasNrRsrpThrldLow

`func (o *C2LinkQualityThrlds) HasNrRsrpThrldLow() bool`

HasNrRsrpThrldLow returns a boolean if a field has been set.

### GetNrRsrpThrldHigh

`func (o *C2LinkQualityThrlds) GetNrRsrpThrldHigh() int32`

GetNrRsrpThrldHigh returns the NrRsrpThrldHigh field if non-nil, zero value otherwise.

### GetNrRsrpThrldHighOk

`func (o *C2LinkQualityThrlds) GetNrRsrpThrldHighOk() (*int32, bool)`

GetNrRsrpThrldHighOk returns a tuple with the NrRsrpThrldHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrRsrpThrldHigh

`func (o *C2LinkQualityThrlds) SetNrRsrpThrldHigh(v int32)`

SetNrRsrpThrldHigh sets NrRsrpThrldHigh field to given value.

### HasNrRsrpThrldHigh

`func (o *C2LinkQualityThrlds) HasNrRsrpThrldHigh() bool`

HasNrRsrpThrldHigh returns a boolean if a field has been set.

### GetNrRsrqThrldLow

`func (o *C2LinkQualityThrlds) GetNrRsrqThrldLow() int32`

GetNrRsrqThrldLow returns the NrRsrqThrldLow field if non-nil, zero value otherwise.

### GetNrRsrqThrldLowOk

`func (o *C2LinkQualityThrlds) GetNrRsrqThrldLowOk() (*int32, bool)`

GetNrRsrqThrldLowOk returns a tuple with the NrRsrqThrldLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrRsrqThrldLow

`func (o *C2LinkQualityThrlds) SetNrRsrqThrldLow(v int32)`

SetNrRsrqThrldLow sets NrRsrqThrldLow field to given value.

### HasNrRsrqThrldLow

`func (o *C2LinkQualityThrlds) HasNrRsrqThrldLow() bool`

HasNrRsrqThrldLow returns a boolean if a field has been set.

### GetNrRsrqThrldHigh

`func (o *C2LinkQualityThrlds) GetNrRsrqThrldHigh() int32`

GetNrRsrqThrldHigh returns the NrRsrqThrldHigh field if non-nil, zero value otherwise.

### GetNrRsrqThrldHighOk

`func (o *C2LinkQualityThrlds) GetNrRsrqThrldHighOk() (*int32, bool)`

GetNrRsrqThrldHighOk returns a tuple with the NrRsrqThrldHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrRsrqThrldHigh

`func (o *C2LinkQualityThrlds) SetNrRsrqThrldHigh(v int32)`

SetNrRsrqThrldHigh sets NrRsrqThrldHigh field to given value.

### HasNrRsrqThrldHigh

`func (o *C2LinkQualityThrlds) HasNrRsrqThrldHigh() bool`

HasNrRsrqThrldHigh returns a boolean if a field has been set.

### GetPacketLossThrldLow

`func (o *C2LinkQualityThrlds) GetPacketLossThrldLow() int32`

GetPacketLossThrldLow returns the PacketLossThrldLow field if non-nil, zero value otherwise.

### GetPacketLossThrldLowOk

`func (o *C2LinkQualityThrlds) GetPacketLossThrldLowOk() (*int32, bool)`

GetPacketLossThrldLowOk returns a tuple with the PacketLossThrldLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLossThrldLow

`func (o *C2LinkQualityThrlds) SetPacketLossThrldLow(v int32)`

SetPacketLossThrldLow sets PacketLossThrldLow field to given value.

### HasPacketLossThrldLow

`func (o *C2LinkQualityThrlds) HasPacketLossThrldLow() bool`

HasPacketLossThrldLow returns a boolean if a field has been set.

### GetPacketLossThrldHigh

`func (o *C2LinkQualityThrlds) GetPacketLossThrldHigh() int32`

GetPacketLossThrldHigh returns the PacketLossThrldHigh field if non-nil, zero value otherwise.

### GetPacketLossThrldHighOk

`func (o *C2LinkQualityThrlds) GetPacketLossThrldHighOk() (*int32, bool)`

GetPacketLossThrldHighOk returns a tuple with the PacketLossThrldHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLossThrldHigh

`func (o *C2LinkQualityThrlds) SetPacketLossThrldHigh(v int32)`

SetPacketLossThrldHigh sets PacketLossThrldHigh field to given value.

### HasPacketLossThrldHigh

`func (o *C2LinkQualityThrlds) HasPacketLossThrldHigh() bool`

HasPacketLossThrldHigh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


