# PolicyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Triggers** | Pointer to [**[]RequestTrigger**](RequestTrigger.md) | Request Triggers that the PCF subscribes. | [optional] 
**ServAreaRes** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**WlServAreaRes** | Pointer to [**WirelineServiceAreaRestriction**](WirelineServiceAreaRestriction.md) |  | [optional] 
**Rfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**TargetRfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**SmfSelInfo** | Pointer to [**NullableSmfSelectionData**](SmfSelectionData.md) |  | [optional] 
**UeAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**UeSliceMbrs** | Pointer to [**[]UeSliceMbr**](UeSliceMbr.md) | One or more UE-Slice-MBR(s) for S-NSSAI(s) of serving PLMN the allowed NSSAI as part of the AMF Access and Mobility Policy as determined by the PCF.  | [optional] 
**Pras** | Pointer to [**map[string]PresenceInfoRm**](PresenceInfoRm.md) | Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map.  | [optional] 
**PcfUeInfo** | Pointer to [**NullablePcfUeCallbackInfo**](PcfUeCallbackInfo.md) |  | [optional] 
**MatchPdus** | Pointer to [**[]PduSessionInfo**](PduSessionInfo.md) |  | [optional] 
**AsTimeDisParam** | Pointer to [**NullableAsTimeDistributionParam**](AsTimeDistributionParam.md) |  | [optional] 

## Methods

### NewPolicyUpdate

`func NewPolicyUpdate(resourceUri string, ) *PolicyUpdate`

NewPolicyUpdate instantiates a new PolicyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUpdateWithDefaults

`func NewPolicyUpdateWithDefaults() *PolicyUpdate`

NewPolicyUpdateWithDefaults instantiates a new PolicyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *PolicyUpdate) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *PolicyUpdate) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *PolicyUpdate) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetTriggers

`func (o *PolicyUpdate) GetTriggers() []RequestTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PolicyUpdate) GetTriggersOk() (*[]RequestTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PolicyUpdate) SetTriggers(v []RequestTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PolicyUpdate) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *PolicyUpdate) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *PolicyUpdate) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetServAreaRes

`func (o *PolicyUpdate) GetServAreaRes() ServiceAreaRestriction`

GetServAreaRes returns the ServAreaRes field if non-nil, zero value otherwise.

### GetServAreaResOk

`func (o *PolicyUpdate) GetServAreaResOk() (*ServiceAreaRestriction, bool)`

GetServAreaResOk returns a tuple with the ServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAreaRes

`func (o *PolicyUpdate) SetServAreaRes(v ServiceAreaRestriction)`

SetServAreaRes sets ServAreaRes field to given value.

### HasServAreaRes

`func (o *PolicyUpdate) HasServAreaRes() bool`

HasServAreaRes returns a boolean if a field has been set.

### GetWlServAreaRes

`func (o *PolicyUpdate) GetWlServAreaRes() WirelineServiceAreaRestriction`

GetWlServAreaRes returns the WlServAreaRes field if non-nil, zero value otherwise.

### GetWlServAreaResOk

`func (o *PolicyUpdate) GetWlServAreaResOk() (*WirelineServiceAreaRestriction, bool)`

GetWlServAreaResOk returns a tuple with the WlServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlServAreaRes

`func (o *PolicyUpdate) SetWlServAreaRes(v WirelineServiceAreaRestriction)`

SetWlServAreaRes sets WlServAreaRes field to given value.

### HasWlServAreaRes

`func (o *PolicyUpdate) HasWlServAreaRes() bool`

HasWlServAreaRes returns a boolean if a field has been set.

### GetRfsp

`func (o *PolicyUpdate) GetRfsp() int32`

GetRfsp returns the Rfsp field if non-nil, zero value otherwise.

### GetRfspOk

`func (o *PolicyUpdate) GetRfspOk() (*int32, bool)`

GetRfspOk returns a tuple with the Rfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfsp

`func (o *PolicyUpdate) SetRfsp(v int32)`

SetRfsp sets Rfsp field to given value.

### HasRfsp

`func (o *PolicyUpdate) HasRfsp() bool`

HasRfsp returns a boolean if a field has been set.

### GetTargetRfsp

`func (o *PolicyUpdate) GetTargetRfsp() int32`

GetTargetRfsp returns the TargetRfsp field if non-nil, zero value otherwise.

### GetTargetRfspOk

`func (o *PolicyUpdate) GetTargetRfspOk() (*int32, bool)`

GetTargetRfspOk returns a tuple with the TargetRfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRfsp

`func (o *PolicyUpdate) SetTargetRfsp(v int32)`

SetTargetRfsp sets TargetRfsp field to given value.

### HasTargetRfsp

`func (o *PolicyUpdate) HasTargetRfsp() bool`

HasTargetRfsp returns a boolean if a field has been set.

### GetSmfSelInfo

`func (o *PolicyUpdate) GetSmfSelInfo() SmfSelectionData`

GetSmfSelInfo returns the SmfSelInfo field if non-nil, zero value otherwise.

### GetSmfSelInfoOk

`func (o *PolicyUpdate) GetSmfSelInfoOk() (*SmfSelectionData, bool)`

GetSmfSelInfoOk returns a tuple with the SmfSelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelInfo

`func (o *PolicyUpdate) SetSmfSelInfo(v SmfSelectionData)`

SetSmfSelInfo sets SmfSelInfo field to given value.

### HasSmfSelInfo

`func (o *PolicyUpdate) HasSmfSelInfo() bool`

HasSmfSelInfo returns a boolean if a field has been set.

### SetSmfSelInfoNil

`func (o *PolicyUpdate) SetSmfSelInfoNil(b bool)`

 SetSmfSelInfoNil sets the value for SmfSelInfo to be an explicit nil

### UnsetSmfSelInfo
`func (o *PolicyUpdate) UnsetSmfSelInfo()`

UnsetSmfSelInfo ensures that no value is present for SmfSelInfo, not even an explicit nil
### GetUeAmbr

`func (o *PolicyUpdate) GetUeAmbr() Ambr`

GetUeAmbr returns the UeAmbr field if non-nil, zero value otherwise.

### GetUeAmbrOk

`func (o *PolicyUpdate) GetUeAmbrOk() (*Ambr, bool)`

GetUeAmbrOk returns a tuple with the UeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAmbr

`func (o *PolicyUpdate) SetUeAmbr(v Ambr)`

SetUeAmbr sets UeAmbr field to given value.

### HasUeAmbr

`func (o *PolicyUpdate) HasUeAmbr() bool`

HasUeAmbr returns a boolean if a field has been set.

### GetUeSliceMbrs

`func (o *PolicyUpdate) GetUeSliceMbrs() []UeSliceMbr`

GetUeSliceMbrs returns the UeSliceMbrs field if non-nil, zero value otherwise.

### GetUeSliceMbrsOk

`func (o *PolicyUpdate) GetUeSliceMbrsOk() (*[]UeSliceMbr, bool)`

GetUeSliceMbrsOk returns a tuple with the UeSliceMbrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSliceMbrs

`func (o *PolicyUpdate) SetUeSliceMbrs(v []UeSliceMbr)`

SetUeSliceMbrs sets UeSliceMbrs field to given value.

### HasUeSliceMbrs

`func (o *PolicyUpdate) HasUeSliceMbrs() bool`

HasUeSliceMbrs returns a boolean if a field has been set.

### GetPras

`func (o *PolicyUpdate) GetPras() map[string]PresenceInfoRm`

GetPras returns the Pras field if non-nil, zero value otherwise.

### GetPrasOk

`func (o *PolicyUpdate) GetPrasOk() (*map[string]PresenceInfoRm, bool)`

GetPrasOk returns a tuple with the Pras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPras

`func (o *PolicyUpdate) SetPras(v map[string]PresenceInfoRm)`

SetPras sets Pras field to given value.

### HasPras

`func (o *PolicyUpdate) HasPras() bool`

HasPras returns a boolean if a field has been set.

### SetPrasNil

`func (o *PolicyUpdate) SetPrasNil(b bool)`

 SetPrasNil sets the value for Pras to be an explicit nil

### UnsetPras
`func (o *PolicyUpdate) UnsetPras()`

UnsetPras ensures that no value is present for Pras, not even an explicit nil
### GetPcfUeInfo

`func (o *PolicyUpdate) GetPcfUeInfo() PcfUeCallbackInfo`

GetPcfUeInfo returns the PcfUeInfo field if non-nil, zero value otherwise.

### GetPcfUeInfoOk

`func (o *PolicyUpdate) GetPcfUeInfoOk() (*PcfUeCallbackInfo, bool)`

GetPcfUeInfoOk returns a tuple with the PcfUeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeInfo

`func (o *PolicyUpdate) SetPcfUeInfo(v PcfUeCallbackInfo)`

SetPcfUeInfo sets PcfUeInfo field to given value.

### HasPcfUeInfo

`func (o *PolicyUpdate) HasPcfUeInfo() bool`

HasPcfUeInfo returns a boolean if a field has been set.

### SetPcfUeInfoNil

`func (o *PolicyUpdate) SetPcfUeInfoNil(b bool)`

 SetPcfUeInfoNil sets the value for PcfUeInfo to be an explicit nil

### UnsetPcfUeInfo
`func (o *PolicyUpdate) UnsetPcfUeInfo()`

UnsetPcfUeInfo ensures that no value is present for PcfUeInfo, not even an explicit nil
### GetMatchPdus

`func (o *PolicyUpdate) GetMatchPdus() []PduSessionInfo`

GetMatchPdus returns the MatchPdus field if non-nil, zero value otherwise.

### GetMatchPdusOk

`func (o *PolicyUpdate) GetMatchPdusOk() (*[]PduSessionInfo, bool)`

GetMatchPdusOk returns a tuple with the MatchPdus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPdus

`func (o *PolicyUpdate) SetMatchPdus(v []PduSessionInfo)`

SetMatchPdus sets MatchPdus field to given value.

### HasMatchPdus

`func (o *PolicyUpdate) HasMatchPdus() bool`

HasMatchPdus returns a boolean if a field has been set.

### SetMatchPdusNil

`func (o *PolicyUpdate) SetMatchPdusNil(b bool)`

 SetMatchPdusNil sets the value for MatchPdus to be an explicit nil

### UnsetMatchPdus
`func (o *PolicyUpdate) UnsetMatchPdus()`

UnsetMatchPdus ensures that no value is present for MatchPdus, not even an explicit nil
### GetAsTimeDisParam

`func (o *PolicyUpdate) GetAsTimeDisParam() AsTimeDistributionParam`

GetAsTimeDisParam returns the AsTimeDisParam field if non-nil, zero value otherwise.

### GetAsTimeDisParamOk

`func (o *PolicyUpdate) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool)`

GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeDisParam

`func (o *PolicyUpdate) SetAsTimeDisParam(v AsTimeDistributionParam)`

SetAsTimeDisParam sets AsTimeDisParam field to given value.

### HasAsTimeDisParam

`func (o *PolicyUpdate) HasAsTimeDisParam() bool`

HasAsTimeDisParam returns a boolean if a field has been set.

### SetAsTimeDisParamNil

`func (o *PolicyUpdate) SetAsTimeDisParamNil(b bool)`

 SetAsTimeDisParamNil sets the value for AsTimeDisParam to be an explicit nil

### UnsetAsTimeDisParam
`func (o *PolicyUpdate) UnsetAsTimeDisParam()`

UnsetAsTimeDisParam ensures that no value is present for AsTimeDisParam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


