# PccRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PccRuleId** | Pointer to **string** | Univocally identifies the PCC rule within a PDU session. | [optional] 
**FlowInfoList** | Pointer to [**[]FlowInformation**](FlowInformation.md) |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**AppDescriptor** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**ContentVersion** | Pointer to **int32** | Represents the content version of some content. | [optional] 
**Precedence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AfSigProtocol** | Pointer to [**AfSigProtocol**](AfSigProtocol.md) |  | [optional] 
**IsAppRelocatable** | Pointer to **bool** |  | [optional] 
**IsUeAddrPreserved** | Pointer to **bool** |  | [optional] 
**QosData** | Pointer to [**[][]QosData**]([]QosData.md) |  | [optional] 
**AltQosParams** | Pointer to [**[][]QosData**]([]QosData.md) |  | [optional] 
**TrafficControlData** | Pointer to [**[][]TrafficControlData**]([]TrafficControlData.md) |  | [optional] 
**ConditionData** | Pointer to [**NullableConditionData**](ConditionData.md) |  | [optional] 
**TscaiInputDl** | Pointer to [**NullableTscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**TscaiInputUl** | Pointer to [**NullableTscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 

## Methods

### NewPccRule

`func NewPccRule() *PccRule`

NewPccRule instantiates a new PccRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPccRuleWithDefaults

`func NewPccRuleWithDefaults() *PccRule`

NewPccRuleWithDefaults instantiates a new PccRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPccRuleId

`func (o *PccRule) GetPccRuleId() string`

GetPccRuleId returns the PccRuleId field if non-nil, zero value otherwise.

### GetPccRuleIdOk

`func (o *PccRule) GetPccRuleIdOk() (*string, bool)`

GetPccRuleIdOk returns a tuple with the PccRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRuleId

`func (o *PccRule) SetPccRuleId(v string)`

SetPccRuleId sets PccRuleId field to given value.

### HasPccRuleId

`func (o *PccRule) HasPccRuleId() bool`

HasPccRuleId returns a boolean if a field has been set.

### GetFlowInfoList

`func (o *PccRule) GetFlowInfoList() []FlowInformation`

GetFlowInfoList returns the FlowInfoList field if non-nil, zero value otherwise.

### GetFlowInfoListOk

`func (o *PccRule) GetFlowInfoListOk() (*[]FlowInformation, bool)`

GetFlowInfoListOk returns a tuple with the FlowInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfoList

`func (o *PccRule) SetFlowInfoList(v []FlowInformation)`

SetFlowInfoList sets FlowInfoList field to given value.

### HasFlowInfoList

`func (o *PccRule) HasFlowInfoList() bool`

HasFlowInfoList returns a boolean if a field has been set.

### GetApplicationId

`func (o *PccRule) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PccRule) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PccRule) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *PccRule) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetAppDescriptor

`func (o *PccRule) GetAppDescriptor() string`

GetAppDescriptor returns the AppDescriptor field if non-nil, zero value otherwise.

### GetAppDescriptorOk

`func (o *PccRule) GetAppDescriptorOk() (*string, bool)`

GetAppDescriptorOk returns a tuple with the AppDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescriptor

`func (o *PccRule) SetAppDescriptor(v string)`

SetAppDescriptor sets AppDescriptor field to given value.

### HasAppDescriptor

`func (o *PccRule) HasAppDescriptor() bool`

HasAppDescriptor returns a boolean if a field has been set.

### GetContentVersion

`func (o *PccRule) GetContentVersion() int32`

GetContentVersion returns the ContentVersion field if non-nil, zero value otherwise.

### GetContentVersionOk

`func (o *PccRule) GetContentVersionOk() (*int32, bool)`

GetContentVersionOk returns a tuple with the ContentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentVersion

`func (o *PccRule) SetContentVersion(v int32)`

SetContentVersion sets ContentVersion field to given value.

### HasContentVersion

`func (o *PccRule) HasContentVersion() bool`

HasContentVersion returns a boolean if a field has been set.

### GetPrecedence

`func (o *PccRule) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *PccRule) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *PccRule) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *PccRule) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetAfSigProtocol

`func (o *PccRule) GetAfSigProtocol() AfSigProtocol`

GetAfSigProtocol returns the AfSigProtocol field if non-nil, zero value otherwise.

### GetAfSigProtocolOk

`func (o *PccRule) GetAfSigProtocolOk() (*AfSigProtocol, bool)`

GetAfSigProtocolOk returns a tuple with the AfSigProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfSigProtocol

`func (o *PccRule) SetAfSigProtocol(v AfSigProtocol)`

SetAfSigProtocol sets AfSigProtocol field to given value.

### HasAfSigProtocol

`func (o *PccRule) HasAfSigProtocol() bool`

HasAfSigProtocol returns a boolean if a field has been set.

### GetIsAppRelocatable

`func (o *PccRule) GetIsAppRelocatable() bool`

GetIsAppRelocatable returns the IsAppRelocatable field if non-nil, zero value otherwise.

### GetIsAppRelocatableOk

`func (o *PccRule) GetIsAppRelocatableOk() (*bool, bool)`

GetIsAppRelocatableOk returns a tuple with the IsAppRelocatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppRelocatable

`func (o *PccRule) SetIsAppRelocatable(v bool)`

SetIsAppRelocatable sets IsAppRelocatable field to given value.

### HasIsAppRelocatable

`func (o *PccRule) HasIsAppRelocatable() bool`

HasIsAppRelocatable returns a boolean if a field has been set.

### GetIsUeAddrPreserved

`func (o *PccRule) GetIsUeAddrPreserved() bool`

GetIsUeAddrPreserved returns the IsUeAddrPreserved field if non-nil, zero value otherwise.

### GetIsUeAddrPreservedOk

`func (o *PccRule) GetIsUeAddrPreservedOk() (*bool, bool)`

GetIsUeAddrPreservedOk returns a tuple with the IsUeAddrPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUeAddrPreserved

`func (o *PccRule) SetIsUeAddrPreserved(v bool)`

SetIsUeAddrPreserved sets IsUeAddrPreserved field to given value.

### HasIsUeAddrPreserved

`func (o *PccRule) HasIsUeAddrPreserved() bool`

HasIsUeAddrPreserved returns a boolean if a field has been set.

### GetQosData

`func (o *PccRule) GetQosData() [][]QosData`

GetQosData returns the QosData field if non-nil, zero value otherwise.

### GetQosDataOk

`func (o *PccRule) GetQosDataOk() (*[][]QosData, bool)`

GetQosDataOk returns a tuple with the QosData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosData

`func (o *PccRule) SetQosData(v [][]QosData)`

SetQosData sets QosData field to given value.

### HasQosData

`func (o *PccRule) HasQosData() bool`

HasQosData returns a boolean if a field has been set.

### GetAltQosParams

`func (o *PccRule) GetAltQosParams() [][]QosData`

GetAltQosParams returns the AltQosParams field if non-nil, zero value otherwise.

### GetAltQosParamsOk

`func (o *PccRule) GetAltQosParamsOk() (*[][]QosData, bool)`

GetAltQosParamsOk returns a tuple with the AltQosParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosParams

`func (o *PccRule) SetAltQosParams(v [][]QosData)`

SetAltQosParams sets AltQosParams field to given value.

### HasAltQosParams

`func (o *PccRule) HasAltQosParams() bool`

HasAltQosParams returns a boolean if a field has been set.

### GetTrafficControlData

`func (o *PccRule) GetTrafficControlData() [][]TrafficControlData`

GetTrafficControlData returns the TrafficControlData field if non-nil, zero value otherwise.

### GetTrafficControlDataOk

`func (o *PccRule) GetTrafficControlDataOk() (*[][]TrafficControlData, bool)`

GetTrafficControlDataOk returns a tuple with the TrafficControlData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficControlData

`func (o *PccRule) SetTrafficControlData(v [][]TrafficControlData)`

SetTrafficControlData sets TrafficControlData field to given value.

### HasTrafficControlData

`func (o *PccRule) HasTrafficControlData() bool`

HasTrafficControlData returns a boolean if a field has been set.

### GetConditionData

`func (o *PccRule) GetConditionData() ConditionData`

GetConditionData returns the ConditionData field if non-nil, zero value otherwise.

### GetConditionDataOk

`func (o *PccRule) GetConditionDataOk() (*ConditionData, bool)`

GetConditionDataOk returns a tuple with the ConditionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionData

`func (o *PccRule) SetConditionData(v ConditionData)`

SetConditionData sets ConditionData field to given value.

### HasConditionData

`func (o *PccRule) HasConditionData() bool`

HasConditionData returns a boolean if a field has been set.

### SetConditionDataNil

`func (o *PccRule) SetConditionDataNil(b bool)`

 SetConditionDataNil sets the value for ConditionData to be an explicit nil

### UnsetConditionData
`func (o *PccRule) UnsetConditionData()`

UnsetConditionData ensures that no value is present for ConditionData, not even an explicit nil
### GetTscaiInputDl

`func (o *PccRule) GetTscaiInputDl() TscaiInputContainer`

GetTscaiInputDl returns the TscaiInputDl field if non-nil, zero value otherwise.

### GetTscaiInputDlOk

`func (o *PccRule) GetTscaiInputDlOk() (*TscaiInputContainer, bool)`

GetTscaiInputDlOk returns a tuple with the TscaiInputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputDl

`func (o *PccRule) SetTscaiInputDl(v TscaiInputContainer)`

SetTscaiInputDl sets TscaiInputDl field to given value.

### HasTscaiInputDl

`func (o *PccRule) HasTscaiInputDl() bool`

HasTscaiInputDl returns a boolean if a field has been set.

### SetTscaiInputDlNil

`func (o *PccRule) SetTscaiInputDlNil(b bool)`

 SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil

### UnsetTscaiInputDl
`func (o *PccRule) UnsetTscaiInputDl()`

UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
### GetTscaiInputUl

`func (o *PccRule) GetTscaiInputUl() TscaiInputContainer`

GetTscaiInputUl returns the TscaiInputUl field if non-nil, zero value otherwise.

### GetTscaiInputUlOk

`func (o *PccRule) GetTscaiInputUlOk() (*TscaiInputContainer, bool)`

GetTscaiInputUlOk returns a tuple with the TscaiInputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputUl

`func (o *PccRule) SetTscaiInputUl(v TscaiInputContainer)`

SetTscaiInputUl sets TscaiInputUl field to given value.

### HasTscaiInputUl

`func (o *PccRule) HasTscaiInputUl() bool`

HasTscaiInputUl returns a boolean if a field has been set.

### SetTscaiInputUlNil

`func (o *PccRule) SetTscaiInputUlNil(b bool)`

 SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil

### UnsetTscaiInputUl
`func (o *PccRule) UnsetTscaiInputUl()`

UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


