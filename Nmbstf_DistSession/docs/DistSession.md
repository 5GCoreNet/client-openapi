# DistSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistSessionId** | **string** |  | 
**DistSessionState** | [**DistSessionState**](DistSessionState.md) |  | 
**MbUpfTunAddr** | [**TunnelAddress**](TunnelAddress.md) |  | 
**UpTrafficFlowInfo** | Pointer to [**UpTrafficFlowInfo**](UpTrafficFlowInfo.md) |  | [optional] 
**Mbr** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MaxDelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**ObjDistributionData** | Pointer to [**ObjDistributionData**](ObjDistributionData.md) |  | [optional] 
**PktDistributionData** | Pointer to [**PktDistributionData**](PktDistributionData.md) |  | [optional] 
**FecInformation** | Pointer to [**FECConfig**](FECConfig.md) |  | [optional] 
**DscpMarking** | Pointer to **string** |  | [optional] 
**MbsSecurityContext** | Pointer to [**MbsSecurityContext**](MbsSecurityContext.md) |  | [optional] 

## Methods

### NewDistSession

`func NewDistSession(distSessionId string, distSessionState DistSessionState, mbUpfTunAddr TunnelAddress, mbr string, ) *DistSession`

NewDistSession instantiates a new DistSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistSessionWithDefaults

`func NewDistSessionWithDefaults() *DistSession`

NewDistSessionWithDefaults instantiates a new DistSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistSessionId

`func (o *DistSession) GetDistSessionId() string`

GetDistSessionId returns the DistSessionId field if non-nil, zero value otherwise.

### GetDistSessionIdOk

`func (o *DistSession) GetDistSessionIdOk() (*string, bool)`

GetDistSessionIdOk returns a tuple with the DistSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistSessionId

`func (o *DistSession) SetDistSessionId(v string)`

SetDistSessionId sets DistSessionId field to given value.


### GetDistSessionState

`func (o *DistSession) GetDistSessionState() DistSessionState`

GetDistSessionState returns the DistSessionState field if non-nil, zero value otherwise.

### GetDistSessionStateOk

`func (o *DistSession) GetDistSessionStateOk() (*DistSessionState, bool)`

GetDistSessionStateOk returns a tuple with the DistSessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistSessionState

`func (o *DistSession) SetDistSessionState(v DistSessionState)`

SetDistSessionState sets DistSessionState field to given value.


### GetMbUpfTunAddr

`func (o *DistSession) GetMbUpfTunAddr() TunnelAddress`

GetMbUpfTunAddr returns the MbUpfTunAddr field if non-nil, zero value otherwise.

### GetMbUpfTunAddrOk

`func (o *DistSession) GetMbUpfTunAddrOk() (*TunnelAddress, bool)`

GetMbUpfTunAddrOk returns a tuple with the MbUpfTunAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbUpfTunAddr

`func (o *DistSession) SetMbUpfTunAddr(v TunnelAddress)`

SetMbUpfTunAddr sets MbUpfTunAddr field to given value.


### GetUpTrafficFlowInfo

`func (o *DistSession) GetUpTrafficFlowInfo() UpTrafficFlowInfo`

GetUpTrafficFlowInfo returns the UpTrafficFlowInfo field if non-nil, zero value otherwise.

### GetUpTrafficFlowInfoOk

`func (o *DistSession) GetUpTrafficFlowInfoOk() (*UpTrafficFlowInfo, bool)`

GetUpTrafficFlowInfoOk returns a tuple with the UpTrafficFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTrafficFlowInfo

`func (o *DistSession) SetUpTrafficFlowInfo(v UpTrafficFlowInfo)`

SetUpTrafficFlowInfo sets UpTrafficFlowInfo field to given value.

### HasUpTrafficFlowInfo

`func (o *DistSession) HasUpTrafficFlowInfo() bool`

HasUpTrafficFlowInfo returns a boolean if a field has been set.

### GetMbr

`func (o *DistSession) GetMbr() string`

GetMbr returns the Mbr field if non-nil, zero value otherwise.

### GetMbrOk

`func (o *DistSession) GetMbrOk() (*string, bool)`

GetMbrOk returns a tuple with the Mbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbr

`func (o *DistSession) SetMbr(v string)`

SetMbr sets Mbr field to given value.


### GetMaxDelay

`func (o *DistSession) GetMaxDelay() int32`

GetMaxDelay returns the MaxDelay field if non-nil, zero value otherwise.

### GetMaxDelayOk

`func (o *DistSession) GetMaxDelayOk() (*int32, bool)`

GetMaxDelayOk returns a tuple with the MaxDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDelay

`func (o *DistSession) SetMaxDelay(v int32)`

SetMaxDelay sets MaxDelay field to given value.

### HasMaxDelay

`func (o *DistSession) HasMaxDelay() bool`

HasMaxDelay returns a boolean if a field has been set.

### GetObjDistributionData

`func (o *DistSession) GetObjDistributionData() ObjDistributionData`

GetObjDistributionData returns the ObjDistributionData field if non-nil, zero value otherwise.

### GetObjDistributionDataOk

`func (o *DistSession) GetObjDistributionDataOk() (*ObjDistributionData, bool)`

GetObjDistributionDataOk returns a tuple with the ObjDistributionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDistributionData

`func (o *DistSession) SetObjDistributionData(v ObjDistributionData)`

SetObjDistributionData sets ObjDistributionData field to given value.

### HasObjDistributionData

`func (o *DistSession) HasObjDistributionData() bool`

HasObjDistributionData returns a boolean if a field has been set.

### GetPktDistributionData

`func (o *DistSession) GetPktDistributionData() PktDistributionData`

GetPktDistributionData returns the PktDistributionData field if non-nil, zero value otherwise.

### GetPktDistributionDataOk

`func (o *DistSession) GetPktDistributionDataOk() (*PktDistributionData, bool)`

GetPktDistributionDataOk returns a tuple with the PktDistributionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPktDistributionData

`func (o *DistSession) SetPktDistributionData(v PktDistributionData)`

SetPktDistributionData sets PktDistributionData field to given value.

### HasPktDistributionData

`func (o *DistSession) HasPktDistributionData() bool`

HasPktDistributionData returns a boolean if a field has been set.

### GetFecInformation

`func (o *DistSession) GetFecInformation() FECConfig`

GetFecInformation returns the FecInformation field if non-nil, zero value otherwise.

### GetFecInformationOk

`func (o *DistSession) GetFecInformationOk() (*FECConfig, bool)`

GetFecInformationOk returns a tuple with the FecInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecInformation

`func (o *DistSession) SetFecInformation(v FECConfig)`

SetFecInformation sets FecInformation field to given value.

### HasFecInformation

`func (o *DistSession) HasFecInformation() bool`

HasFecInformation returns a boolean if a field has been set.

### GetDscpMarking

`func (o *DistSession) GetDscpMarking() string`

GetDscpMarking returns the DscpMarking field if non-nil, zero value otherwise.

### GetDscpMarkingOk

`func (o *DistSession) GetDscpMarkingOk() (*string, bool)`

GetDscpMarkingOk returns a tuple with the DscpMarking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpMarking

`func (o *DistSession) SetDscpMarking(v string)`

SetDscpMarking sets DscpMarking field to given value.

### HasDscpMarking

`func (o *DistSession) HasDscpMarking() bool`

HasDscpMarking returns a boolean if a field has been set.

### GetMbsSecurityContext

`func (o *DistSession) GetMbsSecurityContext() MbsSecurityContext`

GetMbsSecurityContext returns the MbsSecurityContext field if non-nil, zero value otherwise.

### GetMbsSecurityContextOk

`func (o *DistSession) GetMbsSecurityContextOk() (*MbsSecurityContext, bool)`

GetMbsSecurityContextOk returns a tuple with the MbsSecurityContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSecurityContext

`func (o *DistSession) SetMbsSecurityContext(v MbsSecurityContext)`

SetMbsSecurityContext sets MbsSecurityContext field to given value.

### HasMbsSecurityContext

`func (o *DistSession) HasMbsSecurityContext() bool`

HasMbsSecurityContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


