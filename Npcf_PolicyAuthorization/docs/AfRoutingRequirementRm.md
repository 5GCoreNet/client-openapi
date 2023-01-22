# AfRoutingRequirementRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppReloc** | Pointer to **bool** |  | [optional] 
**RouteToLocs** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) |  | [optional] 
**SpVal** | Pointer to [**NullableSpatialValidityRm**](SpatialValidityRm.md) |  | [optional] 
**TempVals** | Pointer to [**[]TemporalValidity**](TemporalValidity.md) |  | [optional] 
**UpPathChgSub** | Pointer to [**NullableUpPathChgEvent**](UpPathChgEvent.md) |  | [optional] 
**AddrPreserInd** | Pointer to **NullableBool** |  | [optional] 
**SimConnInd** | Pointer to **NullableBool** | Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA. | [optional] 
**SimConnTerm** | Pointer to **NullableInt32** | indicating a time in seconds with OpenAPI defined &#39;nullable: true&#39; property. | [optional] 
**EasIpReplaceInfos** | Pointer to [**[]EasIpReplacementInfo**](EasIpReplacementInfo.md) | Contains EAS IP replacement information. | [optional] 
**EasRedisInd** | Pointer to **bool** | Indicates the EAS rediscovery is required. | [optional] 
**MaxAllowedUpLat** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 

## Methods

### NewAfRoutingRequirementRm

`func NewAfRoutingRequirementRm() *AfRoutingRequirementRm`

NewAfRoutingRequirementRm instantiates a new AfRoutingRequirementRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfRoutingRequirementRmWithDefaults

`func NewAfRoutingRequirementRmWithDefaults() *AfRoutingRequirementRm`

NewAfRoutingRequirementRmWithDefaults instantiates a new AfRoutingRequirementRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppReloc

`func (o *AfRoutingRequirementRm) GetAppReloc() bool`

GetAppReloc returns the AppReloc field if non-nil, zero value otherwise.

### GetAppRelocOk

`func (o *AfRoutingRequirementRm) GetAppRelocOk() (*bool, bool)`

GetAppRelocOk returns a tuple with the AppReloc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppReloc

`func (o *AfRoutingRequirementRm) SetAppReloc(v bool)`

SetAppReloc sets AppReloc field to given value.

### HasAppReloc

`func (o *AfRoutingRequirementRm) HasAppReloc() bool`

HasAppReloc returns a boolean if a field has been set.

### GetRouteToLocs

`func (o *AfRoutingRequirementRm) GetRouteToLocs() []RouteToLocation`

GetRouteToLocs returns the RouteToLocs field if non-nil, zero value otherwise.

### GetRouteToLocsOk

`func (o *AfRoutingRequirementRm) GetRouteToLocsOk() (*[]RouteToLocation, bool)`

GetRouteToLocsOk returns a tuple with the RouteToLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToLocs

`func (o *AfRoutingRequirementRm) SetRouteToLocs(v []RouteToLocation)`

SetRouteToLocs sets RouteToLocs field to given value.

### HasRouteToLocs

`func (o *AfRoutingRequirementRm) HasRouteToLocs() bool`

HasRouteToLocs returns a boolean if a field has been set.

### SetRouteToLocsNil

`func (o *AfRoutingRequirementRm) SetRouteToLocsNil(b bool)`

 SetRouteToLocsNil sets the value for RouteToLocs to be an explicit nil

### UnsetRouteToLocs
`func (o *AfRoutingRequirementRm) UnsetRouteToLocs()`

UnsetRouteToLocs ensures that no value is present for RouteToLocs, not even an explicit nil
### GetSpVal

`func (o *AfRoutingRequirementRm) GetSpVal() SpatialValidityRm`

GetSpVal returns the SpVal field if non-nil, zero value otherwise.

### GetSpValOk

`func (o *AfRoutingRequirementRm) GetSpValOk() (*SpatialValidityRm, bool)`

GetSpValOk returns a tuple with the SpVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVal

`func (o *AfRoutingRequirementRm) SetSpVal(v SpatialValidityRm)`

SetSpVal sets SpVal field to given value.

### HasSpVal

`func (o *AfRoutingRequirementRm) HasSpVal() bool`

HasSpVal returns a boolean if a field has been set.

### SetSpValNil

`func (o *AfRoutingRequirementRm) SetSpValNil(b bool)`

 SetSpValNil sets the value for SpVal to be an explicit nil

### UnsetSpVal
`func (o *AfRoutingRequirementRm) UnsetSpVal()`

UnsetSpVal ensures that no value is present for SpVal, not even an explicit nil
### GetTempVals

`func (o *AfRoutingRequirementRm) GetTempVals() []TemporalValidity`

GetTempVals returns the TempVals field if non-nil, zero value otherwise.

### GetTempValsOk

`func (o *AfRoutingRequirementRm) GetTempValsOk() (*[]TemporalValidity, bool)`

GetTempValsOk returns a tuple with the TempVals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempVals

`func (o *AfRoutingRequirementRm) SetTempVals(v []TemporalValidity)`

SetTempVals sets TempVals field to given value.

### HasTempVals

`func (o *AfRoutingRequirementRm) HasTempVals() bool`

HasTempVals returns a boolean if a field has been set.

### SetTempValsNil

`func (o *AfRoutingRequirementRm) SetTempValsNil(b bool)`

 SetTempValsNil sets the value for TempVals to be an explicit nil

### UnsetTempVals
`func (o *AfRoutingRequirementRm) UnsetTempVals()`

UnsetTempVals ensures that no value is present for TempVals, not even an explicit nil
### GetUpPathChgSub

`func (o *AfRoutingRequirementRm) GetUpPathChgSub() UpPathChgEvent`

GetUpPathChgSub returns the UpPathChgSub field if non-nil, zero value otherwise.

### GetUpPathChgSubOk

`func (o *AfRoutingRequirementRm) GetUpPathChgSubOk() (*UpPathChgEvent, bool)`

GetUpPathChgSubOk returns a tuple with the UpPathChgSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPathChgSub

`func (o *AfRoutingRequirementRm) SetUpPathChgSub(v UpPathChgEvent)`

SetUpPathChgSub sets UpPathChgSub field to given value.

### HasUpPathChgSub

`func (o *AfRoutingRequirementRm) HasUpPathChgSub() bool`

HasUpPathChgSub returns a boolean if a field has been set.

### SetUpPathChgSubNil

`func (o *AfRoutingRequirementRm) SetUpPathChgSubNil(b bool)`

 SetUpPathChgSubNil sets the value for UpPathChgSub to be an explicit nil

### UnsetUpPathChgSub
`func (o *AfRoutingRequirementRm) UnsetUpPathChgSub()`

UnsetUpPathChgSub ensures that no value is present for UpPathChgSub, not even an explicit nil
### GetAddrPreserInd

`func (o *AfRoutingRequirementRm) GetAddrPreserInd() bool`

GetAddrPreserInd returns the AddrPreserInd field if non-nil, zero value otherwise.

### GetAddrPreserIndOk

`func (o *AfRoutingRequirementRm) GetAddrPreserIndOk() (*bool, bool)`

GetAddrPreserIndOk returns a tuple with the AddrPreserInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrPreserInd

`func (o *AfRoutingRequirementRm) SetAddrPreserInd(v bool)`

SetAddrPreserInd sets AddrPreserInd field to given value.

### HasAddrPreserInd

`func (o *AfRoutingRequirementRm) HasAddrPreserInd() bool`

HasAddrPreserInd returns a boolean if a field has been set.

### SetAddrPreserIndNil

`func (o *AfRoutingRequirementRm) SetAddrPreserIndNil(b bool)`

 SetAddrPreserIndNil sets the value for AddrPreserInd to be an explicit nil

### UnsetAddrPreserInd
`func (o *AfRoutingRequirementRm) UnsetAddrPreserInd()`

UnsetAddrPreserInd ensures that no value is present for AddrPreserInd, not even an explicit nil
### GetSimConnInd

`func (o *AfRoutingRequirementRm) GetSimConnInd() bool`

GetSimConnInd returns the SimConnInd field if non-nil, zero value otherwise.

### GetSimConnIndOk

`func (o *AfRoutingRequirementRm) GetSimConnIndOk() (*bool, bool)`

GetSimConnIndOk returns a tuple with the SimConnInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnInd

`func (o *AfRoutingRequirementRm) SetSimConnInd(v bool)`

SetSimConnInd sets SimConnInd field to given value.

### HasSimConnInd

`func (o *AfRoutingRequirementRm) HasSimConnInd() bool`

HasSimConnInd returns a boolean if a field has been set.

### SetSimConnIndNil

`func (o *AfRoutingRequirementRm) SetSimConnIndNil(b bool)`

 SetSimConnIndNil sets the value for SimConnInd to be an explicit nil

### UnsetSimConnInd
`func (o *AfRoutingRequirementRm) UnsetSimConnInd()`

UnsetSimConnInd ensures that no value is present for SimConnInd, not even an explicit nil
### GetSimConnTerm

`func (o *AfRoutingRequirementRm) GetSimConnTerm() int32`

GetSimConnTerm returns the SimConnTerm field if non-nil, zero value otherwise.

### GetSimConnTermOk

`func (o *AfRoutingRequirementRm) GetSimConnTermOk() (*int32, bool)`

GetSimConnTermOk returns a tuple with the SimConnTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimConnTerm

`func (o *AfRoutingRequirementRm) SetSimConnTerm(v int32)`

SetSimConnTerm sets SimConnTerm field to given value.

### HasSimConnTerm

`func (o *AfRoutingRequirementRm) HasSimConnTerm() bool`

HasSimConnTerm returns a boolean if a field has been set.

### SetSimConnTermNil

`func (o *AfRoutingRequirementRm) SetSimConnTermNil(b bool)`

 SetSimConnTermNil sets the value for SimConnTerm to be an explicit nil

### UnsetSimConnTerm
`func (o *AfRoutingRequirementRm) UnsetSimConnTerm()`

UnsetSimConnTerm ensures that no value is present for SimConnTerm, not even an explicit nil
### GetEasIpReplaceInfos

`func (o *AfRoutingRequirementRm) GetEasIpReplaceInfos() []EasIpReplacementInfo`

GetEasIpReplaceInfos returns the EasIpReplaceInfos field if non-nil, zero value otherwise.

### GetEasIpReplaceInfosOk

`func (o *AfRoutingRequirementRm) GetEasIpReplaceInfosOk() (*[]EasIpReplacementInfo, bool)`

GetEasIpReplaceInfosOk returns a tuple with the EasIpReplaceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIpReplaceInfos

`func (o *AfRoutingRequirementRm) SetEasIpReplaceInfos(v []EasIpReplacementInfo)`

SetEasIpReplaceInfos sets EasIpReplaceInfos field to given value.

### HasEasIpReplaceInfos

`func (o *AfRoutingRequirementRm) HasEasIpReplaceInfos() bool`

HasEasIpReplaceInfos returns a boolean if a field has been set.

### SetEasIpReplaceInfosNil

`func (o *AfRoutingRequirementRm) SetEasIpReplaceInfosNil(b bool)`

 SetEasIpReplaceInfosNil sets the value for EasIpReplaceInfos to be an explicit nil

### UnsetEasIpReplaceInfos
`func (o *AfRoutingRequirementRm) UnsetEasIpReplaceInfos()`

UnsetEasIpReplaceInfos ensures that no value is present for EasIpReplaceInfos, not even an explicit nil
### GetEasRedisInd

`func (o *AfRoutingRequirementRm) GetEasRedisInd() bool`

GetEasRedisInd returns the EasRedisInd field if non-nil, zero value otherwise.

### GetEasRedisIndOk

`func (o *AfRoutingRequirementRm) GetEasRedisIndOk() (*bool, bool)`

GetEasRedisIndOk returns a tuple with the EasRedisInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasRedisInd

`func (o *AfRoutingRequirementRm) SetEasRedisInd(v bool)`

SetEasRedisInd sets EasRedisInd field to given value.

### HasEasRedisInd

`func (o *AfRoutingRequirementRm) HasEasRedisInd() bool`

HasEasRedisInd returns a boolean if a field has been set.

### GetMaxAllowedUpLat

`func (o *AfRoutingRequirementRm) GetMaxAllowedUpLat() int32`

GetMaxAllowedUpLat returns the MaxAllowedUpLat field if non-nil, zero value otherwise.

### GetMaxAllowedUpLatOk

`func (o *AfRoutingRequirementRm) GetMaxAllowedUpLatOk() (*int32, bool)`

GetMaxAllowedUpLatOk returns a tuple with the MaxAllowedUpLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedUpLat

`func (o *AfRoutingRequirementRm) SetMaxAllowedUpLat(v int32)`

SetMaxAllowedUpLat sets MaxAllowedUpLat field to given value.

### HasMaxAllowedUpLat

`func (o *AfRoutingRequirementRm) HasMaxAllowedUpLat() bool`

HasMaxAllowedUpLat returns a boolean if a field has been set.

### SetMaxAllowedUpLatNil

`func (o *AfRoutingRequirementRm) SetMaxAllowedUpLatNil(b bool)`

 SetMaxAllowedUpLatNil sets the value for MaxAllowedUpLat to be an explicit nil

### UnsetMaxAllowedUpLat
`func (o *AfRoutingRequirementRm) UnsetMaxAllowedUpLat()`

UnsetMaxAllowedUpLat ensures that no value is present for MaxAllowedUpLat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


