# IpPacketFilterSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcIp** | Pointer to **string** |  | [optional] 
**DstIp** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **int32** |  | [optional] 
**SrcPort** | Pointer to **int32** |  | [optional] 
**DstPort** | Pointer to **int32** |  | [optional] 
**ToSTc** | Pointer to **string** |  | [optional] 
**FlowLabel** | Pointer to **int32** |  | [optional] 
**Spi** | Pointer to **int32** |  | [optional] 
**Direction** | **string** |  | 

## Methods

### NewIpPacketFilterSet

`func NewIpPacketFilterSet(direction string, ) *IpPacketFilterSet`

NewIpPacketFilterSet instantiates a new IpPacketFilterSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpPacketFilterSetWithDefaults

`func NewIpPacketFilterSetWithDefaults() *IpPacketFilterSet`

NewIpPacketFilterSetWithDefaults instantiates a new IpPacketFilterSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrcIp

`func (o *IpPacketFilterSet) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *IpPacketFilterSet) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *IpPacketFilterSet) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *IpPacketFilterSet) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetDstIp

`func (o *IpPacketFilterSet) GetDstIp() string`

GetDstIp returns the DstIp field if non-nil, zero value otherwise.

### GetDstIpOk

`func (o *IpPacketFilterSet) GetDstIpOk() (*string, bool)`

GetDstIpOk returns a tuple with the DstIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstIp

`func (o *IpPacketFilterSet) SetDstIp(v string)`

SetDstIp sets DstIp field to given value.

### HasDstIp

`func (o *IpPacketFilterSet) HasDstIp() bool`

HasDstIp returns a boolean if a field has been set.

### GetProtocol

`func (o *IpPacketFilterSet) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpPacketFilterSet) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpPacketFilterSet) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpPacketFilterSet) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPort

`func (o *IpPacketFilterSet) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *IpPacketFilterSet) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *IpPacketFilterSet) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *IpPacketFilterSet) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetDstPort

`func (o *IpPacketFilterSet) GetDstPort() int32`

GetDstPort returns the DstPort field if non-nil, zero value otherwise.

### GetDstPortOk

`func (o *IpPacketFilterSet) GetDstPortOk() (*int32, bool)`

GetDstPortOk returns a tuple with the DstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPort

`func (o *IpPacketFilterSet) SetDstPort(v int32)`

SetDstPort sets DstPort field to given value.

### HasDstPort

`func (o *IpPacketFilterSet) HasDstPort() bool`

HasDstPort returns a boolean if a field has been set.

### GetToSTc

`func (o *IpPacketFilterSet) GetToSTc() string`

GetToSTc returns the ToSTc field if non-nil, zero value otherwise.

### GetToSTcOk

`func (o *IpPacketFilterSet) GetToSTcOk() (*string, bool)`

GetToSTcOk returns a tuple with the ToSTc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSTc

`func (o *IpPacketFilterSet) SetToSTc(v string)`

SetToSTc sets ToSTc field to given value.

### HasToSTc

`func (o *IpPacketFilterSet) HasToSTc() bool`

HasToSTc returns a boolean if a field has been set.

### GetFlowLabel

`func (o *IpPacketFilterSet) GetFlowLabel() int32`

GetFlowLabel returns the FlowLabel field if non-nil, zero value otherwise.

### GetFlowLabelOk

`func (o *IpPacketFilterSet) GetFlowLabelOk() (*int32, bool)`

GetFlowLabelOk returns a tuple with the FlowLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowLabel

`func (o *IpPacketFilterSet) SetFlowLabel(v int32)`

SetFlowLabel sets FlowLabel field to given value.

### HasFlowLabel

`func (o *IpPacketFilterSet) HasFlowLabel() bool`

HasFlowLabel returns a boolean if a field has been set.

### GetSpi

`func (o *IpPacketFilterSet) GetSpi() int32`

GetSpi returns the Spi field if non-nil, zero value otherwise.

### GetSpiOk

`func (o *IpPacketFilterSet) GetSpiOk() (*int32, bool)`

GetSpiOk returns a tuple with the Spi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpi

`func (o *IpPacketFilterSet) SetSpi(v int32)`

SetSpi sets Spi field to given value.

### HasSpi

`func (o *IpPacketFilterSet) HasSpi() bool`

HasSpi returns a boolean if a field has been set.

### GetDirection

`func (o *IpPacketFilterSet) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IpPacketFilterSet) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IpPacketFilterSet) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


