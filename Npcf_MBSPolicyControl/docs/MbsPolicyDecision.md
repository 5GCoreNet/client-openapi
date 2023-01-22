# MbsPolicyDecision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsPccRules** | Pointer to [**map[string]MbsPccRule**](MbsPccRule.md) | A map of MBS PCC rule(s) with each map entry containing the MbsPccRule data structure. The key of the map for each entry is the mbsPccRuleId attribute of the corresponding MbsPccRule data structure.  | [optional] 
**MbsQosDecs** | Pointer to [**map[string]MbsQosDec**](MbsQosDec.md) | A map of MBS QoS Decision(s) with each map entry containing the MbsQosDecdata data structure.The key of the map for each entry is the mbsQosId attribute of the corresponding to the MbsQosDec data structure.  | [optional] 
**MbsQosChars** | Pointer to [**map[string]MbsQosChar**](MbsQosChar.md) | A map of MBS QoS Characteristics set(s) with each map entry containing the MbsQosChar data structure. The key of the map for each entry is the 5QI attribute of the corresponding MbsQosDec data structure.  | [optional] 
**AuthMbsSessAmbr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MbsPcrts** | Pointer to [**[]MbsPcrt**](MbsPcrt.md) |  | [optional] 

## Methods

### NewMbsPolicyDecision

`func NewMbsPolicyDecision() *MbsPolicyDecision`

NewMbsPolicyDecision instantiates a new MbsPolicyDecision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsPolicyDecisionWithDefaults

`func NewMbsPolicyDecisionWithDefaults() *MbsPolicyDecision`

NewMbsPolicyDecisionWithDefaults instantiates a new MbsPolicyDecision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsPccRules

`func (o *MbsPolicyDecision) GetMbsPccRules() map[string]MbsPccRule`

GetMbsPccRules returns the MbsPccRules field if non-nil, zero value otherwise.

### GetMbsPccRulesOk

`func (o *MbsPolicyDecision) GetMbsPccRulesOk() (*map[string]MbsPccRule, bool)`

GetMbsPccRulesOk returns a tuple with the MbsPccRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsPccRules

`func (o *MbsPolicyDecision) SetMbsPccRules(v map[string]MbsPccRule)`

SetMbsPccRules sets MbsPccRules field to given value.

### HasMbsPccRules

`func (o *MbsPolicyDecision) HasMbsPccRules() bool`

HasMbsPccRules returns a boolean if a field has been set.

### SetMbsPccRulesNil

`func (o *MbsPolicyDecision) SetMbsPccRulesNil(b bool)`

 SetMbsPccRulesNil sets the value for MbsPccRules to be an explicit nil

### UnsetMbsPccRules
`func (o *MbsPolicyDecision) UnsetMbsPccRules()`

UnsetMbsPccRules ensures that no value is present for MbsPccRules, not even an explicit nil
### GetMbsQosDecs

`func (o *MbsPolicyDecision) GetMbsQosDecs() map[string]MbsQosDec`

GetMbsQosDecs returns the MbsQosDecs field if non-nil, zero value otherwise.

### GetMbsQosDecsOk

`func (o *MbsPolicyDecision) GetMbsQosDecsOk() (*map[string]MbsQosDec, bool)`

GetMbsQosDecsOk returns a tuple with the MbsQosDecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsQosDecs

`func (o *MbsPolicyDecision) SetMbsQosDecs(v map[string]MbsQosDec)`

SetMbsQosDecs sets MbsQosDecs field to given value.

### HasMbsQosDecs

`func (o *MbsPolicyDecision) HasMbsQosDecs() bool`

HasMbsQosDecs returns a boolean if a field has been set.

### GetMbsQosChars

`func (o *MbsPolicyDecision) GetMbsQosChars() map[string]MbsQosChar`

GetMbsQosChars returns the MbsQosChars field if non-nil, zero value otherwise.

### GetMbsQosCharsOk

`func (o *MbsPolicyDecision) GetMbsQosCharsOk() (*map[string]MbsQosChar, bool)`

GetMbsQosCharsOk returns a tuple with the MbsQosChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsQosChars

`func (o *MbsPolicyDecision) SetMbsQosChars(v map[string]MbsQosChar)`

SetMbsQosChars sets MbsQosChars field to given value.

### HasMbsQosChars

`func (o *MbsPolicyDecision) HasMbsQosChars() bool`

HasMbsQosChars returns a boolean if a field has been set.

### GetAuthMbsSessAmbr

`func (o *MbsPolicyDecision) GetAuthMbsSessAmbr() string`

GetAuthMbsSessAmbr returns the AuthMbsSessAmbr field if non-nil, zero value otherwise.

### GetAuthMbsSessAmbrOk

`func (o *MbsPolicyDecision) GetAuthMbsSessAmbrOk() (*string, bool)`

GetAuthMbsSessAmbrOk returns a tuple with the AuthMbsSessAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMbsSessAmbr

`func (o *MbsPolicyDecision) SetAuthMbsSessAmbr(v string)`

SetAuthMbsSessAmbr sets AuthMbsSessAmbr field to given value.

### HasAuthMbsSessAmbr

`func (o *MbsPolicyDecision) HasAuthMbsSessAmbr() bool`

HasAuthMbsSessAmbr returns a boolean if a field has been set.

### GetMbsPcrts

`func (o *MbsPolicyDecision) GetMbsPcrts() []MbsPcrt`

GetMbsPcrts returns the MbsPcrts field if non-nil, zero value otherwise.

### GetMbsPcrtsOk

`func (o *MbsPolicyDecision) GetMbsPcrtsOk() (*[]MbsPcrt, bool)`

GetMbsPcrtsOk returns a tuple with the MbsPcrts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsPcrts

`func (o *MbsPolicyDecision) SetMbsPcrts(v []MbsPcrt)`

SetMbsPcrts sets MbsPcrts field to given value.

### HasMbsPcrts

`func (o *MbsPolicyDecision) HasMbsPcrts() bool`

HasMbsPcrts returns a boolean if a field has been set.

### SetMbsPcrtsNil

`func (o *MbsPolicyDecision) SetMbsPcrtsNil(b bool)`

 SetMbsPcrtsNil sets the value for MbsPcrts to be an explicit nil

### UnsetMbsPcrts
`func (o *MbsPolicyDecision) UnsetMbsPcrts()`

UnsetMbsPcrts ensures that no value is present for MbsPcrts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


