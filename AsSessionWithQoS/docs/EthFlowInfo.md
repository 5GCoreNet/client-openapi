# EthFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowId** | **int32** | Indicates the Ethernet flow identifier. | 
**EthFlowDescriptions** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Indicates the packet filters of the Ethernet flow. It shall contain UL and/or DL Ethernet flow description.  | [optional] 

## Methods

### NewEthFlowInfo

`func NewEthFlowInfo(flowId int32, ) *EthFlowInfo`

NewEthFlowInfo instantiates a new EthFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEthFlowInfoWithDefaults

`func NewEthFlowInfoWithDefaults() *EthFlowInfo`

NewEthFlowInfoWithDefaults instantiates a new EthFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowId

`func (o *EthFlowInfo) GetFlowId() int32`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *EthFlowInfo) GetFlowIdOk() (*int32, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *EthFlowInfo) SetFlowId(v int32)`

SetFlowId sets FlowId field to given value.


### GetEthFlowDescriptions

`func (o *EthFlowInfo) GetEthFlowDescriptions() []EthFlowDescription`

GetEthFlowDescriptions returns the EthFlowDescriptions field if non-nil, zero value otherwise.

### GetEthFlowDescriptionsOk

`func (o *EthFlowInfo) GetEthFlowDescriptionsOk() (*[]EthFlowDescription, bool)`

GetEthFlowDescriptionsOk returns a tuple with the EthFlowDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowDescriptions

`func (o *EthFlowInfo) SetEthFlowDescriptions(v []EthFlowDescription)`

SetEthFlowDescriptions sets EthFlowDescriptions field to given value.

### HasEthFlowDescriptions

`func (o *EthFlowInfo) HasEthFlowDescriptions() bool`

HasEthFlowDescriptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


