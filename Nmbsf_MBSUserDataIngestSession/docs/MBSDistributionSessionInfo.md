# MBSDistributionSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsDistSessionId** | Pointer to **string** |  | [optional] 
**MbsSessionId** | Pointer to [**MbsSessionId**](MbsSessionId.md) |  | [optional] 
**MbsServInfo** | Pointer to [**MbsServiceInfo**](MbsServiceInfo.md) |  | [optional] 
**MaxContBitRate** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MaxContDelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**DistrMethod** | [**DistributionMethod**](DistributionMethod.md) |  | 
**FecConfig** | Pointer to [**FECConfig**](FECConfig.md) |  | [optional] 
**ObjDistrInfo** | Pointer to [**ObjectDistrMethInfo**](ObjectDistrMethInfo.md) |  | [optional] 
**PckDistrInfo** | Pointer to [**PacketDistrMethInfo**](PacketDistrMethInfo.md) |  | [optional] 
**TrafficMarkingInfo** | Pointer to **string** |  | [optional] 
**MbsDistSessState** | Pointer to [**DistSessionState**](DistSessionState.md) |  | [optional] 
**TgtServAreas** | Pointer to [**MbsServiceArea**](MbsServiceArea.md) |  | [optional] 
**ExtTgtServAreas** | Pointer to [**ExternalMbsServiceArea**](ExternalMbsServiceArea.md) |  | [optional] 
**MbsFSAId** | Pointer to **string** | MBS Frequency Selection Area Identifier | [optional] 
**LocationDependent** | Pointer to **bool** | Represents an indication that this MBS Distribution Session belongs to a location- dependent MBS. This attribute shall be set to \&quot;true\&quot; to indicate that the MBS  Distribution Session belongs to a location-dependent MBS; or set to \&quot;false\&quot; to  indicate that the MBS Distribution Session does not belong to a location-dependent MBS. The default value is \&quot;false\&quot;, if omitted.  | [optional] [default to false]
**MultiplexedServFlag** | Pointer to **bool** | Represents an indication that this MBS Distribution Session belongs to a multiplex, i.e.  forms part of a set of MBS Distribution Sessions under the same parent MBS User Data  Ingest Session with identical or empty sets of target service areas and multiplexed onto  the same MBS Session at the MB-SMF.  | [optional] [default to false]
**RestrictedFlag** | Pointer to **bool** | Represents an indication that this MBS Distribution Session is not open to any UE, i.e.  restricted to a set of UEs according to their MBS related subscription information. This attribute may be included only if the parent MBS User Service is of Multicast service type. This attribute shall be set to \&quot;true\&quot; to indicate that this MBS Distribution Session is restricted to a set of UE(s); or set to \&quot;false\&quot; to indicate that this MBS Distribution Session is open to any UE. The default value is \&quot;false\&quot;, if omitted.  | [optional] [default to false]

## Methods

### NewMBSDistributionSessionInfo

`func NewMBSDistributionSessionInfo(maxContBitRate string, distrMethod DistributionMethod, ) *MBSDistributionSessionInfo`

NewMBSDistributionSessionInfo instantiates a new MBSDistributionSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMBSDistributionSessionInfoWithDefaults

`func NewMBSDistributionSessionInfoWithDefaults() *MBSDistributionSessionInfo`

NewMBSDistributionSessionInfoWithDefaults instantiates a new MBSDistributionSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsDistSessionId

`func (o *MBSDistributionSessionInfo) GetMbsDistSessionId() string`

GetMbsDistSessionId returns the MbsDistSessionId field if non-nil, zero value otherwise.

### GetMbsDistSessionIdOk

`func (o *MBSDistributionSessionInfo) GetMbsDistSessionIdOk() (*string, bool)`

GetMbsDistSessionIdOk returns a tuple with the MbsDistSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDistSessionId

`func (o *MBSDistributionSessionInfo) SetMbsDistSessionId(v string)`

SetMbsDistSessionId sets MbsDistSessionId field to given value.

### HasMbsDistSessionId

`func (o *MBSDistributionSessionInfo) HasMbsDistSessionId() bool`

HasMbsDistSessionId returns a boolean if a field has been set.

### GetMbsSessionId

`func (o *MBSDistributionSessionInfo) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MBSDistributionSessionInfo) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MBSDistributionSessionInfo) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.

### HasMbsSessionId

`func (o *MBSDistributionSessionInfo) HasMbsSessionId() bool`

HasMbsSessionId returns a boolean if a field has been set.

### GetMbsServInfo

`func (o *MBSDistributionSessionInfo) GetMbsServInfo() MbsServiceInfo`

GetMbsServInfo returns the MbsServInfo field if non-nil, zero value otherwise.

### GetMbsServInfoOk

`func (o *MBSDistributionSessionInfo) GetMbsServInfoOk() (*MbsServiceInfo, bool)`

GetMbsServInfoOk returns a tuple with the MbsServInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServInfo

`func (o *MBSDistributionSessionInfo) SetMbsServInfo(v MbsServiceInfo)`

SetMbsServInfo sets MbsServInfo field to given value.

### HasMbsServInfo

`func (o *MBSDistributionSessionInfo) HasMbsServInfo() bool`

HasMbsServInfo returns a boolean if a field has been set.

### GetMaxContBitRate

`func (o *MBSDistributionSessionInfo) GetMaxContBitRate() string`

GetMaxContBitRate returns the MaxContBitRate field if non-nil, zero value otherwise.

### GetMaxContBitRateOk

`func (o *MBSDistributionSessionInfo) GetMaxContBitRateOk() (*string, bool)`

GetMaxContBitRateOk returns a tuple with the MaxContBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContBitRate

`func (o *MBSDistributionSessionInfo) SetMaxContBitRate(v string)`

SetMaxContBitRate sets MaxContBitRate field to given value.


### GetMaxContDelay

`func (o *MBSDistributionSessionInfo) GetMaxContDelay() int32`

GetMaxContDelay returns the MaxContDelay field if non-nil, zero value otherwise.

### GetMaxContDelayOk

`func (o *MBSDistributionSessionInfo) GetMaxContDelayOk() (*int32, bool)`

GetMaxContDelayOk returns a tuple with the MaxContDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxContDelay

`func (o *MBSDistributionSessionInfo) SetMaxContDelay(v int32)`

SetMaxContDelay sets MaxContDelay field to given value.

### HasMaxContDelay

`func (o *MBSDistributionSessionInfo) HasMaxContDelay() bool`

HasMaxContDelay returns a boolean if a field has been set.

### GetDistrMethod

`func (o *MBSDistributionSessionInfo) GetDistrMethod() DistributionMethod`

GetDistrMethod returns the DistrMethod field if non-nil, zero value otherwise.

### GetDistrMethodOk

`func (o *MBSDistributionSessionInfo) GetDistrMethodOk() (*DistributionMethod, bool)`

GetDistrMethodOk returns a tuple with the DistrMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistrMethod

`func (o *MBSDistributionSessionInfo) SetDistrMethod(v DistributionMethod)`

SetDistrMethod sets DistrMethod field to given value.


### GetFecConfig

`func (o *MBSDistributionSessionInfo) GetFecConfig() FECConfig`

GetFecConfig returns the FecConfig field if non-nil, zero value otherwise.

### GetFecConfigOk

`func (o *MBSDistributionSessionInfo) GetFecConfigOk() (*FECConfig, bool)`

GetFecConfigOk returns a tuple with the FecConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFecConfig

`func (o *MBSDistributionSessionInfo) SetFecConfig(v FECConfig)`

SetFecConfig sets FecConfig field to given value.

### HasFecConfig

`func (o *MBSDistributionSessionInfo) HasFecConfig() bool`

HasFecConfig returns a boolean if a field has been set.

### GetObjDistrInfo

`func (o *MBSDistributionSessionInfo) GetObjDistrInfo() ObjectDistrMethInfo`

GetObjDistrInfo returns the ObjDistrInfo field if non-nil, zero value otherwise.

### GetObjDistrInfoOk

`func (o *MBSDistributionSessionInfo) GetObjDistrInfoOk() (*ObjectDistrMethInfo, bool)`

GetObjDistrInfoOk returns a tuple with the ObjDistrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDistrInfo

`func (o *MBSDistributionSessionInfo) SetObjDistrInfo(v ObjectDistrMethInfo)`

SetObjDistrInfo sets ObjDistrInfo field to given value.

### HasObjDistrInfo

`func (o *MBSDistributionSessionInfo) HasObjDistrInfo() bool`

HasObjDistrInfo returns a boolean if a field has been set.

### GetPckDistrInfo

`func (o *MBSDistributionSessionInfo) GetPckDistrInfo() PacketDistrMethInfo`

GetPckDistrInfo returns the PckDistrInfo field if non-nil, zero value otherwise.

### GetPckDistrInfoOk

`func (o *MBSDistributionSessionInfo) GetPckDistrInfoOk() (*PacketDistrMethInfo, bool)`

GetPckDistrInfoOk returns a tuple with the PckDistrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPckDistrInfo

`func (o *MBSDistributionSessionInfo) SetPckDistrInfo(v PacketDistrMethInfo)`

SetPckDistrInfo sets PckDistrInfo field to given value.

### HasPckDistrInfo

`func (o *MBSDistributionSessionInfo) HasPckDistrInfo() bool`

HasPckDistrInfo returns a boolean if a field has been set.

### GetTrafficMarkingInfo

`func (o *MBSDistributionSessionInfo) GetTrafficMarkingInfo() string`

GetTrafficMarkingInfo returns the TrafficMarkingInfo field if non-nil, zero value otherwise.

### GetTrafficMarkingInfoOk

`func (o *MBSDistributionSessionInfo) GetTrafficMarkingInfoOk() (*string, bool)`

GetTrafficMarkingInfoOk returns a tuple with the TrafficMarkingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficMarkingInfo

`func (o *MBSDistributionSessionInfo) SetTrafficMarkingInfo(v string)`

SetTrafficMarkingInfo sets TrafficMarkingInfo field to given value.

### HasTrafficMarkingInfo

`func (o *MBSDistributionSessionInfo) HasTrafficMarkingInfo() bool`

HasTrafficMarkingInfo returns a boolean if a field has been set.

### GetMbsDistSessState

`func (o *MBSDistributionSessionInfo) GetMbsDistSessState() DistSessionState`

GetMbsDistSessState returns the MbsDistSessState field if non-nil, zero value otherwise.

### GetMbsDistSessStateOk

`func (o *MBSDistributionSessionInfo) GetMbsDistSessStateOk() (*DistSessionState, bool)`

GetMbsDistSessStateOk returns a tuple with the MbsDistSessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDistSessState

`func (o *MBSDistributionSessionInfo) SetMbsDistSessState(v DistSessionState)`

SetMbsDistSessState sets MbsDistSessState field to given value.

### HasMbsDistSessState

`func (o *MBSDistributionSessionInfo) HasMbsDistSessState() bool`

HasMbsDistSessState returns a boolean if a field has been set.

### GetTgtServAreas

`func (o *MBSDistributionSessionInfo) GetTgtServAreas() MbsServiceArea`

GetTgtServAreas returns the TgtServAreas field if non-nil, zero value otherwise.

### GetTgtServAreasOk

`func (o *MBSDistributionSessionInfo) GetTgtServAreasOk() (*MbsServiceArea, bool)`

GetTgtServAreasOk returns a tuple with the TgtServAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtServAreas

`func (o *MBSDistributionSessionInfo) SetTgtServAreas(v MbsServiceArea)`

SetTgtServAreas sets TgtServAreas field to given value.

### HasTgtServAreas

`func (o *MBSDistributionSessionInfo) HasTgtServAreas() bool`

HasTgtServAreas returns a boolean if a field has been set.

### GetExtTgtServAreas

`func (o *MBSDistributionSessionInfo) GetExtTgtServAreas() ExternalMbsServiceArea`

GetExtTgtServAreas returns the ExtTgtServAreas field if non-nil, zero value otherwise.

### GetExtTgtServAreasOk

`func (o *MBSDistributionSessionInfo) GetExtTgtServAreasOk() (*ExternalMbsServiceArea, bool)`

GetExtTgtServAreasOk returns a tuple with the ExtTgtServAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtTgtServAreas

`func (o *MBSDistributionSessionInfo) SetExtTgtServAreas(v ExternalMbsServiceArea)`

SetExtTgtServAreas sets ExtTgtServAreas field to given value.

### HasExtTgtServAreas

`func (o *MBSDistributionSessionInfo) HasExtTgtServAreas() bool`

HasExtTgtServAreas returns a boolean if a field has been set.

### GetMbsFSAId

`func (o *MBSDistributionSessionInfo) GetMbsFSAId() string`

GetMbsFSAId returns the MbsFSAId field if non-nil, zero value otherwise.

### GetMbsFSAIdOk

`func (o *MBSDistributionSessionInfo) GetMbsFSAIdOk() (*string, bool)`

GetMbsFSAIdOk returns a tuple with the MbsFSAId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsFSAId

`func (o *MBSDistributionSessionInfo) SetMbsFSAId(v string)`

SetMbsFSAId sets MbsFSAId field to given value.

### HasMbsFSAId

`func (o *MBSDistributionSessionInfo) HasMbsFSAId() bool`

HasMbsFSAId returns a boolean if a field has been set.

### GetLocationDependent

`func (o *MBSDistributionSessionInfo) GetLocationDependent() bool`

GetLocationDependent returns the LocationDependent field if non-nil, zero value otherwise.

### GetLocationDependentOk

`func (o *MBSDistributionSessionInfo) GetLocationDependentOk() (*bool, bool)`

GetLocationDependentOk returns a tuple with the LocationDependent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationDependent

`func (o *MBSDistributionSessionInfo) SetLocationDependent(v bool)`

SetLocationDependent sets LocationDependent field to given value.

### HasLocationDependent

`func (o *MBSDistributionSessionInfo) HasLocationDependent() bool`

HasLocationDependent returns a boolean if a field has been set.

### GetMultiplexedServFlag

`func (o *MBSDistributionSessionInfo) GetMultiplexedServFlag() bool`

GetMultiplexedServFlag returns the MultiplexedServFlag field if non-nil, zero value otherwise.

### GetMultiplexedServFlagOk

`func (o *MBSDistributionSessionInfo) GetMultiplexedServFlagOk() (*bool, bool)`

GetMultiplexedServFlagOk returns a tuple with the MultiplexedServFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplexedServFlag

`func (o *MBSDistributionSessionInfo) SetMultiplexedServFlag(v bool)`

SetMultiplexedServFlag sets MultiplexedServFlag field to given value.

### HasMultiplexedServFlag

`func (o *MBSDistributionSessionInfo) HasMultiplexedServFlag() bool`

HasMultiplexedServFlag returns a boolean if a field has been set.

### GetRestrictedFlag

`func (o *MBSDistributionSessionInfo) GetRestrictedFlag() bool`

GetRestrictedFlag returns the RestrictedFlag field if non-nil, zero value otherwise.

### GetRestrictedFlagOk

`func (o *MBSDistributionSessionInfo) GetRestrictedFlagOk() (*bool, bool)`

GetRestrictedFlagOk returns a tuple with the RestrictedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedFlag

`func (o *MBSDistributionSessionInfo) SetRestrictedFlag(v bool)`

SetRestrictedFlag sets RestrictedFlag field to given value.

### HasRestrictedFlag

`func (o *MBSDistributionSessionInfo) HasRestrictedFlag() bool`

HasRestrictedFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


