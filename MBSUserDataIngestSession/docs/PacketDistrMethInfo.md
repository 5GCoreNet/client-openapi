# PacketDistrMethInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatingMode** | Pointer to [**PktDistributionOperatingMode**](PktDistributionOperatingMode.md) |  | [optional] 
**PckIngMethod** | Pointer to [**PktIngestMethod**](PktIngestMethod.md) |  | [optional] 
**IngEndpointAddrs** | [**MbStfIngestAddr**](MbStfIngestAddr.md) |  | 

## Methods

### NewPacketDistrMethInfo

`func NewPacketDistrMethInfo(ingEndpointAddrs MbStfIngestAddr, ) *PacketDistrMethInfo`

NewPacketDistrMethInfo instantiates a new PacketDistrMethInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPacketDistrMethInfoWithDefaults

`func NewPacketDistrMethInfoWithDefaults() *PacketDistrMethInfo`

NewPacketDistrMethInfoWithDefaults instantiates a new PacketDistrMethInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatingMode

`func (o *PacketDistrMethInfo) GetOperatingMode() PktDistributionOperatingMode`

GetOperatingMode returns the OperatingMode field if non-nil, zero value otherwise.

### GetOperatingModeOk

`func (o *PacketDistrMethInfo) GetOperatingModeOk() (*PktDistributionOperatingMode, bool)`

GetOperatingModeOk returns a tuple with the OperatingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingMode

`func (o *PacketDistrMethInfo) SetOperatingMode(v PktDistributionOperatingMode)`

SetOperatingMode sets OperatingMode field to given value.

### HasOperatingMode

`func (o *PacketDistrMethInfo) HasOperatingMode() bool`

HasOperatingMode returns a boolean if a field has been set.

### GetPckIngMethod

`func (o *PacketDistrMethInfo) GetPckIngMethod() PktIngestMethod`

GetPckIngMethod returns the PckIngMethod field if non-nil, zero value otherwise.

### GetPckIngMethodOk

`func (o *PacketDistrMethInfo) GetPckIngMethodOk() (*PktIngestMethod, bool)`

GetPckIngMethodOk returns a tuple with the PckIngMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPckIngMethod

`func (o *PacketDistrMethInfo) SetPckIngMethod(v PktIngestMethod)`

SetPckIngMethod sets PckIngMethod field to given value.

### HasPckIngMethod

`func (o *PacketDistrMethInfo) HasPckIngMethod() bool`

HasPckIngMethod returns a boolean if a field has been set.

### GetIngEndpointAddrs

`func (o *PacketDistrMethInfo) GetIngEndpointAddrs() MbStfIngestAddr`

GetIngEndpointAddrs returns the IngEndpointAddrs field if non-nil, zero value otherwise.

### GetIngEndpointAddrsOk

`func (o *PacketDistrMethInfo) GetIngEndpointAddrsOk() (*MbStfIngestAddr, bool)`

GetIngEndpointAddrsOk returns a tuple with the IngEndpointAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngEndpointAddrs

`func (o *PacketDistrMethInfo) SetIngEndpointAddrs(v MbStfIngestAddr)`

SetIngEndpointAddrs sets IngEndpointAddrs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


