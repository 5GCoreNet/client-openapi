# AppplicationQosRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pqi** | Pointer to **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.  | [optional] 
**ResourceType** | Pointer to [**QosResourceType**](QosResourceType.md) |  | [optional] 
**PriorityLevel** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**PacketDelayBudget** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**PacketErrorRate** | Pointer to **string** | String representing Packet Error Rate (see clause 5.7.3.5 and 5.7.4 of 3GPP TS 23.501, expressed as a \&quot;scalar x 10-k\&quot; where the scalar and the exponent k are each encoded as one decimal digit.  | [optional] 
**AveragingWindow** | Pointer to **int32** | Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] [default to 2000]
**MaxDataBurstVol** | Pointer to **int32** | Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.   | [optional] 

## Methods

### NewAppplicationQosRequirement

`func NewAppplicationQosRequirement() *AppplicationQosRequirement`

NewAppplicationQosRequirement instantiates a new AppplicationQosRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppplicationQosRequirementWithDefaults

`func NewAppplicationQosRequirementWithDefaults() *AppplicationQosRequirement`

NewAppplicationQosRequirementWithDefaults instantiates a new AppplicationQosRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPqi

`func (o *AppplicationQosRequirement) GetPqi() int32`

GetPqi returns the Pqi field if non-nil, zero value otherwise.

### GetPqiOk

`func (o *AppplicationQosRequirement) GetPqiOk() (*int32, bool)`

GetPqiOk returns a tuple with the Pqi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPqi

`func (o *AppplicationQosRequirement) SetPqi(v int32)`

SetPqi sets Pqi field to given value.

### HasPqi

`func (o *AppplicationQosRequirement) HasPqi() bool`

HasPqi returns a boolean if a field has been set.

### GetResourceType

`func (o *AppplicationQosRequirement) GetResourceType() QosResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AppplicationQosRequirement) GetResourceTypeOk() (*QosResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AppplicationQosRequirement) SetResourceType(v QosResourceType)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AppplicationQosRequirement) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetPriorityLevel

`func (o *AppplicationQosRequirement) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *AppplicationQosRequirement) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *AppplicationQosRequirement) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *AppplicationQosRequirement) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.

### GetPacketDelayBudget

`func (o *AppplicationQosRequirement) GetPacketDelayBudget() int32`

GetPacketDelayBudget returns the PacketDelayBudget field if non-nil, zero value otherwise.

### GetPacketDelayBudgetOk

`func (o *AppplicationQosRequirement) GetPacketDelayBudgetOk() (*int32, bool)`

GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDelayBudget

`func (o *AppplicationQosRequirement) SetPacketDelayBudget(v int32)`

SetPacketDelayBudget sets PacketDelayBudget field to given value.

### HasPacketDelayBudget

`func (o *AppplicationQosRequirement) HasPacketDelayBudget() bool`

HasPacketDelayBudget returns a boolean if a field has been set.

### GetPacketErrorRate

`func (o *AppplicationQosRequirement) GetPacketErrorRate() string`

GetPacketErrorRate returns the PacketErrorRate field if non-nil, zero value otherwise.

### GetPacketErrorRateOk

`func (o *AppplicationQosRequirement) GetPacketErrorRateOk() (*string, bool)`

GetPacketErrorRateOk returns a tuple with the PacketErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketErrorRate

`func (o *AppplicationQosRequirement) SetPacketErrorRate(v string)`

SetPacketErrorRate sets PacketErrorRate field to given value.

### HasPacketErrorRate

`func (o *AppplicationQosRequirement) HasPacketErrorRate() bool`

HasPacketErrorRate returns a boolean if a field has been set.

### GetAveragingWindow

`func (o *AppplicationQosRequirement) GetAveragingWindow() int32`

GetAveragingWindow returns the AveragingWindow field if non-nil, zero value otherwise.

### GetAveragingWindowOk

`func (o *AppplicationQosRequirement) GetAveragingWindowOk() (*int32, bool)`

GetAveragingWindowOk returns a tuple with the AveragingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragingWindow

`func (o *AppplicationQosRequirement) SetAveragingWindow(v int32)`

SetAveragingWindow sets AveragingWindow field to given value.

### HasAveragingWindow

`func (o *AppplicationQosRequirement) HasAveragingWindow() bool`

HasAveragingWindow returns a boolean if a field has been set.

### GetMaxDataBurstVol

`func (o *AppplicationQosRequirement) GetMaxDataBurstVol() int32`

GetMaxDataBurstVol returns the MaxDataBurstVol field if non-nil, zero value otherwise.

### GetMaxDataBurstVolOk

`func (o *AppplicationQosRequirement) GetMaxDataBurstVolOk() (*int32, bool)`

GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataBurstVol

`func (o *AppplicationQosRequirement) SetMaxDataBurstVol(v int32)`

SetMaxDataBurstVol sets MaxDataBurstVol field to given value.

### HasMaxDataBurstVol

`func (o *AppplicationQosRequirement) HasMaxDataBurstVol() bool`

HasMaxDataBurstVol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


