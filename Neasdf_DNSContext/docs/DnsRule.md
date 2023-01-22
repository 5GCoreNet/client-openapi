# DnsRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsRuleId** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Precedence** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**DnsQueryMdtList** | Pointer to [**map[string]DnsQueryMdt**](DnsQueryMdt.md) | map of DNS query message detection templates where a valid JSON string serves as key | [optional] 
**BaseDnsQueryMdtList** | Pointer to [**[]BaselineDnsQueryMdtInfo**](BaselineDnsQueryMdtInfo.md) |  | [optional] 
**DnsRspMdtList** | Pointer to [**map[string]DnsRspMdt**](DnsRspMdt.md) | map of DNS response message detection templates where a valid JSON string serves as key | [optional] 
**BaseDnsRspMdtList** | Pointer to [**[]BaselineDnsRspMdtInfo**](BaselineDnsRspMdtInfo.md) |  | [optional] 
**DnsMsgId** | Pointer to **string** |  | [optional] 
**ActionList** | [**map[string]Action**](Action.md) | map of actions where a valid JSON string serves as key | 

## Methods

### NewDnsRule

`func NewDnsRule(actionList map[string]Action, ) *DnsRule`

NewDnsRule instantiates a new DnsRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRuleWithDefaults

`func NewDnsRuleWithDefaults() *DnsRule`

NewDnsRuleWithDefaults instantiates a new DnsRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsRuleId

`func (o *DnsRule) GetDnsRuleId() string`

GetDnsRuleId returns the DnsRuleId field if non-nil, zero value otherwise.

### GetDnsRuleIdOk

`func (o *DnsRule) GetDnsRuleIdOk() (*string, bool)`

GetDnsRuleIdOk returns a tuple with the DnsRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRuleId

`func (o *DnsRule) SetDnsRuleId(v string)`

SetDnsRuleId sets DnsRuleId field to given value.

### HasDnsRuleId

`func (o *DnsRule) HasDnsRuleId() bool`

HasDnsRuleId returns a boolean if a field has been set.

### GetLabel

`func (o *DnsRule) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DnsRule) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DnsRule) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DnsRule) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrecedence

`func (o *DnsRule) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *DnsRule) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *DnsRule) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *DnsRule) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetDnsQueryMdtList

`func (o *DnsRule) GetDnsQueryMdtList() map[string]DnsQueryMdt`

GetDnsQueryMdtList returns the DnsQueryMdtList field if non-nil, zero value otherwise.

### GetDnsQueryMdtListOk

`func (o *DnsRule) GetDnsQueryMdtListOk() (*map[string]DnsQueryMdt, bool)`

GetDnsQueryMdtListOk returns a tuple with the DnsQueryMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQueryMdtList

`func (o *DnsRule) SetDnsQueryMdtList(v map[string]DnsQueryMdt)`

SetDnsQueryMdtList sets DnsQueryMdtList field to given value.

### HasDnsQueryMdtList

`func (o *DnsRule) HasDnsQueryMdtList() bool`

HasDnsQueryMdtList returns a boolean if a field has been set.

### GetBaseDnsQueryMdtList

`func (o *DnsRule) GetBaseDnsQueryMdtList() []BaselineDnsQueryMdtInfo`

GetBaseDnsQueryMdtList returns the BaseDnsQueryMdtList field if non-nil, zero value otherwise.

### GetBaseDnsQueryMdtListOk

`func (o *DnsRule) GetBaseDnsQueryMdtListOk() (*[]BaselineDnsQueryMdtInfo, bool)`

GetBaseDnsQueryMdtListOk returns a tuple with the BaseDnsQueryMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDnsQueryMdtList

`func (o *DnsRule) SetBaseDnsQueryMdtList(v []BaselineDnsQueryMdtInfo)`

SetBaseDnsQueryMdtList sets BaseDnsQueryMdtList field to given value.

### HasBaseDnsQueryMdtList

`func (o *DnsRule) HasBaseDnsQueryMdtList() bool`

HasBaseDnsQueryMdtList returns a boolean if a field has been set.

### GetDnsRspMdtList

`func (o *DnsRule) GetDnsRspMdtList() map[string]DnsRspMdt`

GetDnsRspMdtList returns the DnsRspMdtList field if non-nil, zero value otherwise.

### GetDnsRspMdtListOk

`func (o *DnsRule) GetDnsRspMdtListOk() (*map[string]DnsRspMdt, bool)`

GetDnsRspMdtListOk returns a tuple with the DnsRspMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRspMdtList

`func (o *DnsRule) SetDnsRspMdtList(v map[string]DnsRspMdt)`

SetDnsRspMdtList sets DnsRspMdtList field to given value.

### HasDnsRspMdtList

`func (o *DnsRule) HasDnsRspMdtList() bool`

HasDnsRspMdtList returns a boolean if a field has been set.

### GetBaseDnsRspMdtList

`func (o *DnsRule) GetBaseDnsRspMdtList() []BaselineDnsRspMdtInfo`

GetBaseDnsRspMdtList returns the BaseDnsRspMdtList field if non-nil, zero value otherwise.

### GetBaseDnsRspMdtListOk

`func (o *DnsRule) GetBaseDnsRspMdtListOk() (*[]BaselineDnsRspMdtInfo, bool)`

GetBaseDnsRspMdtListOk returns a tuple with the BaseDnsRspMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDnsRspMdtList

`func (o *DnsRule) SetBaseDnsRspMdtList(v []BaselineDnsRspMdtInfo)`

SetBaseDnsRspMdtList sets BaseDnsRspMdtList field to given value.

### HasBaseDnsRspMdtList

`func (o *DnsRule) HasBaseDnsRspMdtList() bool`

HasBaseDnsRspMdtList returns a boolean if a field has been set.

### GetDnsMsgId

`func (o *DnsRule) GetDnsMsgId() string`

GetDnsMsgId returns the DnsMsgId field if non-nil, zero value otherwise.

### GetDnsMsgIdOk

`func (o *DnsRule) GetDnsMsgIdOk() (*string, bool)`

GetDnsMsgIdOk returns a tuple with the DnsMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsMsgId

`func (o *DnsRule) SetDnsMsgId(v string)`

SetDnsMsgId sets DnsMsgId field to given value.

### HasDnsMsgId

`func (o *DnsRule) HasDnsMsgId() bool`

HasDnsMsgId returns a boolean if a field has been set.

### GetActionList

`func (o *DnsRule) GetActionList() map[string]Action`

GetActionList returns the ActionList field if non-nil, zero value otherwise.

### GetActionListOk

`func (o *DnsRule) GetActionListOk() (*map[string]Action, bool)`

GetActionListOk returns a tuple with the ActionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionList

`func (o *DnsRule) SetActionList(v map[string]Action)`

SetActionList sets ActionList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


