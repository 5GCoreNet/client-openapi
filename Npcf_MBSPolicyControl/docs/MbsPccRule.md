# MbsPccRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsPccRuleId** | **string** |  | 
**MbsDlIpFlowInfo** | Pointer to **[]string** |  | [optional] 
**Precedence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RefMbsQosDec** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMbsPccRule

`func NewMbsPccRule(mbsPccRuleId string, ) *MbsPccRule`

NewMbsPccRule instantiates a new MbsPccRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsPccRuleWithDefaults

`func NewMbsPccRuleWithDefaults() *MbsPccRule`

NewMbsPccRuleWithDefaults instantiates a new MbsPccRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsPccRuleId

`func (o *MbsPccRule) GetMbsPccRuleId() string`

GetMbsPccRuleId returns the MbsPccRuleId field if non-nil, zero value otherwise.

### GetMbsPccRuleIdOk

`func (o *MbsPccRule) GetMbsPccRuleIdOk() (*string, bool)`

GetMbsPccRuleIdOk returns a tuple with the MbsPccRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsPccRuleId

`func (o *MbsPccRule) SetMbsPccRuleId(v string)`

SetMbsPccRuleId sets MbsPccRuleId field to given value.


### GetMbsDlIpFlowInfo

`func (o *MbsPccRule) GetMbsDlIpFlowInfo() []string`

GetMbsDlIpFlowInfo returns the MbsDlIpFlowInfo field if non-nil, zero value otherwise.

### GetMbsDlIpFlowInfoOk

`func (o *MbsPccRule) GetMbsDlIpFlowInfoOk() (*[]string, bool)`

GetMbsDlIpFlowInfoOk returns a tuple with the MbsDlIpFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsDlIpFlowInfo

`func (o *MbsPccRule) SetMbsDlIpFlowInfo(v []string)`

SetMbsDlIpFlowInfo sets MbsDlIpFlowInfo field to given value.

### HasMbsDlIpFlowInfo

`func (o *MbsPccRule) HasMbsDlIpFlowInfo() bool`

HasMbsDlIpFlowInfo returns a boolean if a field has been set.

### GetPrecedence

`func (o *MbsPccRule) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *MbsPccRule) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *MbsPccRule) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *MbsPccRule) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetRefMbsQosDec

`func (o *MbsPccRule) GetRefMbsQosDec() []string`

GetRefMbsQosDec returns the RefMbsQosDec field if non-nil, zero value otherwise.

### GetRefMbsQosDecOk

`func (o *MbsPccRule) GetRefMbsQosDecOk() (*[]string, bool)`

GetRefMbsQosDecOk returns a tuple with the RefMbsQosDec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefMbsQosDec

`func (o *MbsPccRule) SetRefMbsQosDec(v []string)`

SetRefMbsQosDec sets RefMbsQosDec field to given value.

### HasRefMbsQosDec

`func (o *MbsPccRule) HasRefMbsQosDec() bool`

HasRefMbsQosDec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


