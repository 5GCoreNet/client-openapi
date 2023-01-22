# MbsQosChar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qi** | **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.  | 
**PriorityLevel** | **int32** | Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.   | 
**ResourceType** | [**QosResourceType**](QosResourceType.md) |  | 
**PacketDelayBudget** | **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | 
**PacketErrorRate** | **string** | String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \&quot;scalar x 10-k\&quot; where the scalar and the exponent k are each encoded as one decimal digit.  | 
**AverWindow** | Pointer to **int32** | Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] [default to 2000]
**MbsMaxDataBurstVol** | **int32** | Represents MBS Maximum Data Burst Volume, expressed in Bytes. | 

## Methods

### NewMbsQosChar

`func NewMbsQosChar(var5qi int32, priorityLevel int32, resourceType QosResourceType, packetDelayBudget int32, packetErrorRate string, mbsMaxDataBurstVol int32, ) *MbsQosChar`

NewMbsQosChar instantiates a new MbsQosChar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsQosCharWithDefaults

`func NewMbsQosCharWithDefaults() *MbsQosChar`

NewMbsQosCharWithDefaults instantiates a new MbsQosChar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qi

`func (o *MbsQosChar) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *MbsQosChar) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *MbsQosChar) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.


### GetPriorityLevel

`func (o *MbsQosChar) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *MbsQosChar) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *MbsQosChar) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.


### GetResourceType

`func (o *MbsQosChar) GetResourceType() QosResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MbsQosChar) GetResourceTypeOk() (*QosResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MbsQosChar) SetResourceType(v QosResourceType)`

SetResourceType sets ResourceType field to given value.


### GetPacketDelayBudget

`func (o *MbsQosChar) GetPacketDelayBudget() int32`

GetPacketDelayBudget returns the PacketDelayBudget field if non-nil, zero value otherwise.

### GetPacketDelayBudgetOk

`func (o *MbsQosChar) GetPacketDelayBudgetOk() (*int32, bool)`

GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDelayBudget

`func (o *MbsQosChar) SetPacketDelayBudget(v int32)`

SetPacketDelayBudget sets PacketDelayBudget field to given value.


### GetPacketErrorRate

`func (o *MbsQosChar) GetPacketErrorRate() string`

GetPacketErrorRate returns the PacketErrorRate field if non-nil, zero value otherwise.

### GetPacketErrorRateOk

`func (o *MbsQosChar) GetPacketErrorRateOk() (*string, bool)`

GetPacketErrorRateOk returns a tuple with the PacketErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketErrorRate

`func (o *MbsQosChar) SetPacketErrorRate(v string)`

SetPacketErrorRate sets PacketErrorRate field to given value.


### GetAverWindow

`func (o *MbsQosChar) GetAverWindow() int32`

GetAverWindow returns the AverWindow field if non-nil, zero value otherwise.

### GetAverWindowOk

`func (o *MbsQosChar) GetAverWindowOk() (*int32, bool)`

GetAverWindowOk returns a tuple with the AverWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverWindow

`func (o *MbsQosChar) SetAverWindow(v int32)`

SetAverWindow sets AverWindow field to given value.

### HasAverWindow

`func (o *MbsQosChar) HasAverWindow() bool`

HasAverWindow returns a boolean if a field has been set.

### GetMbsMaxDataBurstVol

`func (o *MbsQosChar) GetMbsMaxDataBurstVol() int32`

GetMbsMaxDataBurstVol returns the MbsMaxDataBurstVol field if non-nil, zero value otherwise.

### GetMbsMaxDataBurstVolOk

`func (o *MbsQosChar) GetMbsMaxDataBurstVolOk() (*int32, bool)`

GetMbsMaxDataBurstVolOk returns a tuple with the MbsMaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsMaxDataBurstVol

`func (o *MbsQosChar) SetMbsMaxDataBurstVol(v int32)`

SetMbsMaxDataBurstVol sets MbsMaxDataBurstVol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


