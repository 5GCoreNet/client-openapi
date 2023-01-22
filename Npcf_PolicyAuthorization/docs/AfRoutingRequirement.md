# AfRoutingRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppReloc** | Pointer to **bool** |  | [optional] 
**RouteToLocs** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) |  | [optional] 
**SpVal** | Pointer to [**SpatialValidity**](SpatialValidity.md) |  | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 
**UpPathChgSub** | Pointer to [**NullableUpPathChgEvent**](UpPathChgEvent.md) |  | [optional] 
**AddrPreserInd** | Pointer to **bool** |  | [optional] 
**SimConnInd** | Pointer to **bool** | Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA. | [optional] 
**SimConnTerm** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**EasIpReplaceInfos** | Pointer to [**[]EasIpReplacementInfo**](EasIpReplacementInfo.md) | Contains EAS IP replacement information. | [optional] 
**EasRedisInd** | Pointer to **bool** | Indicates the EAS rediscovery is required. | [optional] 
**MaxAllowedUpLat** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewAfRoutingRequirement

`func NewAfRoutingRequirement() *AfRoutingRequirement`

NewAfRoutingRequirement instantiates a new AfRoutingRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfRoutingRequirementWithDefaults

`func NewAfRoutingRequirementWithDefaults() *AfRoutingRequirement`

NewAfRoutingRequirementWithDefaults instantiates a new AfRoutingRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppReloc

`func (o *AfRoutingRequirement) GetAppReloc() bool`

GetAppReloc returns the AppReloc field if non-nil, zero value otherwise.

### GetAppRelocOk

`func (o *AfRoutingRequirement) GetAppRelocOk() (*bool, bool)`

GetAppRelocOk returns a tuple with the AppReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloc

`func (o *AfRoutingRequirement) SetAppReloc(v bool)`

SetAppReloc sets AppReloc field to given value.

### HasAppReloc

`func (o *AfRoutingRequirement) HasAppReloc() bool`

HasAppReloc returns a boolean if a field has been set.

### GetRouteToLocs

`func (o *AfRoutingRequirement) GetRouteToLocs() []RouteToLocation`

GetRouteToLocs returns the RouteToLocs field if non-nil, zero value otherwise.

### GetRouteToLocsOk

`func (o *AfRoutingRequirement) GetRouteToLocsOk() (*[]RouteToLocation, bool)`

GetRouteToLocsOk returns a tuple with the RouteToLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToLocs

`func (o *AfRoutingRequirement) SetRouteToLocs(v []RouteToLocation)`

SetRouteToLocs sets RouteToLocs field to given value.

### HasRouteToLocs

`func (o *AfRoutingRequirement) HasRouteToLocs() bool`

HasRouteToLocs returns a boolean if a field has been set.

### GetSpVal

`func (o *AfRoutingRequirement) GetSpVal() SpatialValidity`

GetSpVal returns the SpVal field if non-nil, zero value otherwise.

### GetSpValOk

`func (o *AfRoutingRequirement) GetSpValOk() (*SpatialValidity, bool)`

GetSpValOk returns a tuple with the SpVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVal

`func (o *AfRoutingRequirement) SetSpVal(v SpatialValidity)`

SetSpVal sets SpVal field to given value.

### HasSpVal

`func (o *AfRoutingRequirement) HasSpVal() bool`

HasSpVal returns a boolean if a field has been set.

### GetTempVals

`func (o *AfRoutingRequirement) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *AfRoutingRequirement) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *AfRoutingRequirement) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *AfRoutingRequirement) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.

### GetUpPathChgSub

`func (o *AfRoutingRequirement) GetUpPathChgSub() UpPathChgEvent`

GetUpPathChgSub returns the UpPathChgSub field if non-nil, zero value otherwise.

### GetUpPathChgSubOk

`func (o *AfRoutingRequirement) GetUpPathChgSubOk() (*UpPathChgEvent, bool)`

GetUpPathChgSubOk returns a tuple with the UpPathChgSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgSub

`func (o *AfRoutingRequirement) SetUpPathChgSub(v UpPathChgEvent)`

SetUpPathChgSub sets UpPathChgSub field to given value.

### HasUpPathChgSub

`func (o *AfRoutingRequirement) HasUpPathChgSub() bool`

HasUpPathChgSub returns a boolean if a field has been set.

### SetUpPathChgSubNil

`func (o *AfRoutingRequirement) SetUpPathChgSubNil(b bool)`

 SetUpPathChgSubNil sets the value for UpPathChgSub to be an explicit nil

### UnsetUpPathChgSub
`func (o *AfRoutingRequirement) UnsetUpPathChgSub()`

UnsetUpPathChgSub ensures that no value is present for UpPathChgSub, not even an explicit nil
### GetAddrPreserInd

`func (o *AfRoutingRequirement) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *AfRoutingRequirement) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *AfRoutingRequirement) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *AfRoutingRequirement) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### GetSimConnInd

`func (o *AfRoutingRequirement) GetSimConnInd() bool`

GetSimConnInd returns the SimConnInd field if non-nil, zero value otherwise.

### GetSimConnIndOk

`func (o *AfRoutingRequirement) GetSimConnIndOk() (*bool, bool)`

GetSimConnIndOk returns a tuple with the SimConnInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnInd

`func (o *AfRoutingRequirement) SetSimConnInd(v bool)`

SetSimConnInd sets SimConnInd field to given value.

### HasSimConnInd

`func (o *AfRoutingRequirement) HasSimConnInd() bool`

HasSimConnInd returns a boolean if a field has been set.

### GetSimConnTerm

`func (o *AfRoutingRequirement) GetSimConnTerm() int32`

GetSimConnTerm returns the SimConnTerm field if non-nil, zero value otherwise.

### GetSimConnTermOk

`func (o *AfRoutingRequirement) GetSimConnTermOk() (*int32, bool)`

GetSimConnTermOk returns a tuple with the SimConnTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnTerm

`func (o *AfRoutingRequirement) SetSimConnTerm(v int32)`

SetSimConnTerm sets SimConnTerm field to given value.

### HasSimConnTerm

`func (o *AfRoutingRequirement) HasSimConnTerm() bool`

HasSimConnTerm returns a boolean if a field has been set.

### GetEasIpReplaceInfos

`func (o *AfRoutingRequirement) GetEasIpReplaceInfos() []EasIpReplacementInfo`

GetEasIpReplaceInfos returns the EasIpReplaceInfos field if non-nil, zero value otherwise.

### GetEasIpReplaceInfosOk

`func (o *AfRoutingRequirement) GetEasIpReplaceInfosOk() (*[]EasIpReplacementInfo, bool)`

GetEasIpReplaceInfosOk returns a tuple with the EasIpReplaceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpReplaceInfos

`func (o *AfRoutingRequirement) SetEasIpReplaceInfos(v []EasIpReplacementInfo)`

SetEasIpReplaceInfos sets EasIpReplaceInfos field to given value.

### HasEasIpReplaceInfos

`func (o *AfRoutingRequirement) HasEasIpReplaceInfos() bool`

HasEasIpReplaceInfos returns a boolean if a field has been set.

### GetEasRedisInd

`func (o *AfRoutingRequirement) GetEasRedisInd() bool`

GetEasRedisInd returns the EasRedisInd field if non-nil, zero value otherwise.

### GetEasRedisIndOk

`func (o *AfRoutingRequirement) GetEasRedisIndOk() (*bool, bool)`

GetEasRedisIndOk returns a tuple with the EasRedisInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasRedisInd

`func (o *AfRoutingRequirement) SetEasRedisInd(v bool)`

SetEasRedisInd sets EasRedisInd field to given value.

### HasEasRedisInd

`func (o *AfRoutingRequirement) HasEasRedisInd() bool`

HasEasRedisInd returns a boolean if a field has been set.

### GetMaxAllowedUpLat

`func (o *AfRoutingRequirement) GetMaxAllowedUpLat() int32`

GetMaxAllowedUpLat returns the MaxAllowedUpLat field if non-nil, zero value otherwise.

### GetMaxAllowedUpLatOk

`func (o *AfRoutingRequirement) GetMaxAllowedUpLatOk() (*int32, bool)`

GetMaxAllowedUpLatOk returns a tuple with the MaxAllowedUpLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedUpLat

`func (o *AfRoutingRequirement) SetMaxAllowedUpLat(v int32)`

SetMaxAllowedUpLat sets MaxAllowedUpLat field to given value.

### HasMaxAllowedUpLat

`func (o *AfRoutingRequirement) HasMaxAllowedUpLat() bool`

HasMaxAllowedUpLat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


