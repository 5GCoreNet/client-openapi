# TrafficSpecInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PcpValue** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**MaxFramInt** | **int32** | Unsigned integer identifying a period of time in units of seconds. | 
**MaxFramSize** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**MaxIntFrames** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**MaxLatency** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 

## Methods

### NewTrafficSpecInformation

`func NewTrafficSpecInformation(pcpValue int32, maxFramInt int32, maxFramSize int32, maxIntFrames int32, maxLatency int32, ) *TrafficSpecInformation`

NewTrafficSpecInformation instantiates a new TrafficSpecInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficSpecInformationWithDefaults

`func NewTrafficSpecInformationWithDefaults() *TrafficSpecInformation`

NewTrafficSpecInformationWithDefaults instantiates a new TrafficSpecInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcpValue

`func (o *TrafficSpecInformation) GetPcpValue() int32`

GetPcpValue returns the PcpValue field if non-nil, zero value otherwise.

### GetPcpValueOk

`func (o *TrafficSpecInformation) GetPcpValueOk() (*int32, bool)`

GetPcpValueOk returns a tuple with the PcpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcpValue

`func (o *TrafficSpecInformation) SetPcpValue(v int32)`

SetPcpValue sets PcpValue field to given value.


### GetMaxFramInt

`func (o *TrafficSpecInformation) GetMaxFramInt() int32`

GetMaxFramInt returns the MaxFramInt field if non-nil, zero value otherwise.

### GetMaxFramIntOk

`func (o *TrafficSpecInformation) GetMaxFramIntOk() (*int32, bool)`

GetMaxFramIntOk returns a tuple with the MaxFramInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFramInt

`func (o *TrafficSpecInformation) SetMaxFramInt(v int32)`

SetMaxFramInt sets MaxFramInt field to given value.


### GetMaxFramSize

`func (o *TrafficSpecInformation) GetMaxFramSize() int32`

GetMaxFramSize returns the MaxFramSize field if non-nil, zero value otherwise.

### GetMaxFramSizeOk

`func (o *TrafficSpecInformation) GetMaxFramSizeOk() (*int32, bool)`

GetMaxFramSizeOk returns a tuple with the MaxFramSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFramSize

`func (o *TrafficSpecInformation) SetMaxFramSize(v int32)`

SetMaxFramSize sets MaxFramSize field to given value.


### GetMaxIntFrames

`func (o *TrafficSpecInformation) GetMaxIntFrames() int32`

GetMaxIntFrames returns the MaxIntFrames field if non-nil, zero value otherwise.

### GetMaxIntFramesOk

`func (o *TrafficSpecInformation) GetMaxIntFramesOk() (*int32, bool)`

GetMaxIntFramesOk returns a tuple with the MaxIntFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntFrames

`func (o *TrafficSpecInformation) SetMaxIntFrames(v int32)`

SetMaxIntFrames sets MaxIntFrames field to given value.


### GetMaxLatency

`func (o *TrafficSpecInformation) GetMaxLatency() int32`

GetMaxLatency returns the MaxLatency field if non-nil, zero value otherwise.

### GetMaxLatencyOk

`func (o *TrafficSpecInformation) GetMaxLatencyOk() (*int32, bool)`

GetMaxLatencyOk returns a tuple with the MaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatency

`func (o *TrafficSpecInformation) SetMaxLatency(v int32)`

SetMaxLatency sets MaxLatency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


