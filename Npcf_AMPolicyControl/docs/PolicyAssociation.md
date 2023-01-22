# PolicyAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | Pointer to [**PolicyAssociationRequest**](PolicyAssociationRequest.md) |  | [optional] 
**Triggers** | Pointer to [**[]RequestTrigger**](RequestTrigger.md) | Request Triggers that the PCF subscribes. | [optional] 
**ServAreaRes** | Pointer to [**ServiceAreaRestriction**](ServiceAreaRestriction.md) |  | [optional] 
**WlServAreaRes** | Pointer to [**WirelineServiceAreaRestriction**](WirelineServiceAreaRestriction.md) |  | [optional] 
**Rfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**TargetRfsp** | Pointer to **int32** | Unsigned integer representing the \&quot;Subscriber Profile ID for RAT/Frequency Priority\&quot;  as specified in 3GPP TS 36.413.  | [optional] 
**SmfSelInfo** | Pointer to [**NullableSmfSelectionData**](SmfSelectionData.md) |  | [optional] 
**UeAmbr** | Pointer to [**Ambr**](Ambr.md) |  | [optional] 
**UeSliceMbrs** | Pointer to [**[]UeSliceMbr**](UeSliceMbr.md) | One or more UE-Slice-MBR(s) for S-NSSAI(s) of serving PLMN as part of the AMF Access and Mobility Policy as determined by the PCF.  | [optional] 
**Pras** | Pointer to [**map[string]PresenceInfo**](PresenceInfo.md) | Contains the presence reporting area(s) for which reporting was requested. The praId attribute within the PresenceInfo data type is the key of the map.  | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 
**PcfUeInfo** | Pointer to [**NullablePcfUeCallbackInfo**](PcfUeCallbackInfo.md) |  | [optional] 
**MatchPdus** | Pointer to [**[]PduSessionInfo**](PduSessionInfo.md) |  | [optional] 
**AsTimeDisParam** | Pointer to [**NullableAsTimeDistributionParam**](AsTimeDistributionParam.md) |  | [optional] 

## Methods

### NewPolicyAssociation

`func NewPolicyAssociation(suppFeat string, ) *PolicyAssociation`

NewPolicyAssociation instantiates a new PolicyAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAssociationWithDefaults

`func NewPolicyAssociationWithDefaults() *PolicyAssociation`

NewPolicyAssociationWithDefaults instantiates a new PolicyAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *PolicyAssociation) GetRequest() PolicyAssociationRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *PolicyAssociation) GetRequestOk() (*PolicyAssociationRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *PolicyAssociation) SetRequest(v PolicyAssociationRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *PolicyAssociation) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetTriggers

`func (o *PolicyAssociation) GetTriggers() []RequestTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *PolicyAssociation) GetTriggersOk() (*[]RequestTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *PolicyAssociation) SetTriggers(v []RequestTrigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *PolicyAssociation) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetServAreaRes

`func (o *PolicyAssociation) GetServAreaRes() ServiceAreaRestriction`

GetServAreaRes returns the ServAreaRes field if non-nil, zero value otherwise.

### GetServAreaResOk

`func (o *PolicyAssociation) GetServAreaResOk() (*ServiceAreaRestriction, bool)`

GetServAreaResOk returns a tuple with the ServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAreaRes

`func (o *PolicyAssociation) SetServAreaRes(v ServiceAreaRestriction)`

SetServAreaRes sets ServAreaRes field to given value.

### HasServAreaRes

`func (o *PolicyAssociation) HasServAreaRes() bool`

HasServAreaRes returns a boolean if a field has been set.

### GetWlServAreaRes

`func (o *PolicyAssociation) GetWlServAreaRes() WirelineServiceAreaRestriction`

GetWlServAreaRes returns the WlServAreaRes field if non-nil, zero value otherwise.

### GetWlServAreaResOk

`func (o *PolicyAssociation) GetWlServAreaResOk() (*WirelineServiceAreaRestriction, bool)`

GetWlServAreaResOk returns a tuple with the WlServAreaRes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlServAreaRes

`func (o *PolicyAssociation) SetWlServAreaRes(v WirelineServiceAreaRestriction)`

SetWlServAreaRes sets WlServAreaRes field to given value.

### HasWlServAreaRes

`func (o *PolicyAssociation) HasWlServAreaRes() bool`

HasWlServAreaRes returns a boolean if a field has been set.

### GetRfsp

`func (o *PolicyAssociation) GetRfsp() int32`

GetRfsp returns the Rfsp field if non-nil, zero value otherwise.

### GetRfspOk

`func (o *PolicyAssociation) GetRfspOk() (*int32, bool)`

GetRfspOk returns a tuple with the Rfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRfsp

`func (o *PolicyAssociation) SetRfsp(v int32)`

SetRfsp sets Rfsp field to given value.

### HasRfsp

`func (o *PolicyAssociation) HasRfsp() bool`

HasRfsp returns a boolean if a field has been set.

### GetTargetRfsp

`func (o *PolicyAssociation) GetTargetRfsp() int32`

GetTargetRfsp returns the TargetRfsp field if non-nil, zero value otherwise.

### GetTargetRfspOk

`func (o *PolicyAssociation) GetTargetRfspOk() (*int32, bool)`

GetTargetRfspOk returns a tuple with the TargetRfsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRfsp

`func (o *PolicyAssociation) SetTargetRfsp(v int32)`

SetTargetRfsp sets TargetRfsp field to given value.

### HasTargetRfsp

`func (o *PolicyAssociation) HasTargetRfsp() bool`

HasTargetRfsp returns a boolean if a field has been set.

### GetSmfSelInfo

`func (o *PolicyAssociation) GetSmfSelInfo() SmfSelectionData`

GetSmfSelInfo returns the SmfSelInfo field if non-nil, zero value otherwise.

### GetSmfSelInfoOk

`func (o *PolicyAssociation) GetSmfSelInfoOk() (*SmfSelectionData, bool)`

GetSmfSelInfoOk returns a tuple with the SmfSelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfSelInfo

`func (o *PolicyAssociation) SetSmfSelInfo(v SmfSelectionData)`

SetSmfSelInfo sets SmfSelInfo field to given value.

### HasSmfSelInfo

`func (o *PolicyAssociation) HasSmfSelInfo() bool`

HasSmfSelInfo returns a boolean if a field has been set.

### SetSmfSelInfoNil

`func (o *PolicyAssociation) SetSmfSelInfoNil(b bool)`

 SetSmfSelInfoNil sets the value for SmfSelInfo to be an explicit nil

### UnsetSmfSelInfo
`func (o *PolicyAssociation) UnsetSmfSelInfo()`

UnsetSmfSelInfo ensures that no value is present for SmfSelInfo, not even an explicit nil
### GetUeAmbr

`func (o *PolicyAssociation) GetUeAmbr() Ambr`

GetUeAmbr returns the UeAmbr field if non-nil, zero value otherwise.

### GetUeAmbrOk

`func (o *PolicyAssociation) GetUeAmbrOk() (*Ambr, bool)`

GetUeAmbrOk returns a tuple with the UeAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAmbr

`func (o *PolicyAssociation) SetUeAmbr(v Ambr)`

SetUeAmbr sets UeAmbr field to given value.

### HasUeAmbr

`func (o *PolicyAssociation) HasUeAmbr() bool`

HasUeAmbr returns a boolean if a field has been set.

### GetUeSliceMbrs

`func (o *PolicyAssociation) GetUeSliceMbrs() []UeSliceMbr`

GetUeSliceMbrs returns the UeSliceMbrs field if non-nil, zero value otherwise.

### GetUeSliceMbrsOk

`func (o *PolicyAssociation) GetUeSliceMbrsOk() (*[]UeSliceMbr, bool)`

GetUeSliceMbrsOk returns a tuple with the UeSliceMbrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSliceMbrs

`func (o *PolicyAssociation) SetUeSliceMbrs(v []UeSliceMbr)`

SetUeSliceMbrs sets UeSliceMbrs field to given value.

### HasUeSliceMbrs

`func (o *PolicyAssociation) HasUeSliceMbrs() bool`

HasUeSliceMbrs returns a boolean if a field has been set.

### GetPras

`func (o *PolicyAssociation) GetPras() map[string]PresenceInfo`

GetPras returns the Pras field if non-nil, zero value otherwise.

### GetPrasOk

`func (o *PolicyAssociation) GetPrasOk() (*map[string]PresenceInfo, bool)`

GetPrasOk returns a tuple with the Pras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPras

`func (o *PolicyAssociation) SetPras(v map[string]PresenceInfo)`

SetPras sets Pras field to given value.

### HasPras

`func (o *PolicyAssociation) HasPras() bool`

HasPras returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PolicyAssociation) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PolicyAssociation) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PolicyAssociation) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.


### GetPcfUeInfo

`func (o *PolicyAssociation) GetPcfUeInfo() PcfUeCallbackInfo`

GetPcfUeInfo returns the PcfUeInfo field if non-nil, zero value otherwise.

### GetPcfUeInfoOk

`func (o *PolicyAssociation) GetPcfUeInfoOk() (*PcfUeCallbackInfo, bool)`

GetPcfUeInfoOk returns a tuple with the PcfUeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeInfo

`func (o *PolicyAssociation) SetPcfUeInfo(v PcfUeCallbackInfo)`

SetPcfUeInfo sets PcfUeInfo field to given value.

### HasPcfUeInfo

`func (o *PolicyAssociation) HasPcfUeInfo() bool`

HasPcfUeInfo returns a boolean if a field has been set.

### SetPcfUeInfoNil

`func (o *PolicyAssociation) SetPcfUeInfoNil(b bool)`

 SetPcfUeInfoNil sets the value for PcfUeInfo to be an explicit nil

### UnsetPcfUeInfo
`func (o *PolicyAssociation) UnsetPcfUeInfo()`

UnsetPcfUeInfo ensures that no value is present for PcfUeInfo, not even an explicit nil
### GetMatchPdus

`func (o *PolicyAssociation) GetMatchPdus() []PduSessionInfo`

GetMatchPdus returns the MatchPdus field if non-nil, zero value otherwise.

### GetMatchPdusOk

`func (o *PolicyAssociation) GetMatchPdusOk() (*[]PduSessionInfo, bool)`

GetMatchPdusOk returns a tuple with the MatchPdus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPdus

`func (o *PolicyAssociation) SetMatchPdus(v []PduSessionInfo)`

SetMatchPdus sets MatchPdus field to given value.

### HasMatchPdus

`func (o *PolicyAssociation) HasMatchPdus() bool`

HasMatchPdus returns a boolean if a field has been set.

### SetMatchPdusNil

`func (o *PolicyAssociation) SetMatchPdusNil(b bool)`

 SetMatchPdusNil sets the value for MatchPdus to be an explicit nil

### UnsetMatchPdus
`func (o *PolicyAssociation) UnsetMatchPdus()`

UnsetMatchPdus ensures that no value is present for MatchPdus, not even an explicit nil
### GetAsTimeDisParam

`func (o *PolicyAssociation) GetAsTimeDisParam() AsTimeDistributionParam`

GetAsTimeDisParam returns the AsTimeDisParam field if non-nil, zero value otherwise.

### GetAsTimeDisParamOk

`func (o *PolicyAssociation) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool)`

GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeDisParam

`func (o *PolicyAssociation) SetAsTimeDisParam(v AsTimeDistributionParam)`

SetAsTimeDisParam sets AsTimeDisParam field to given value.

### HasAsTimeDisParam

`func (o *PolicyAssociation) HasAsTimeDisParam() bool`

HasAsTimeDisParam returns a boolean if a field has been set.

### SetAsTimeDisParamNil

`func (o *PolicyAssociation) SetAsTimeDisParamNil(b bool)`

 SetAsTimeDisParamNil sets the value for AsTimeDisParam to be an explicit nil

### UnsetAsTimeDisParam
`func (o *PolicyAssociation) UnsetAsTimeDisParam()`

UnsetAsTimeDisParam ensures that no value is present for AsTimeDisParam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


