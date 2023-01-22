# TscQosRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReqGbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ReqGbrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ReqMbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ReqMbrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxTscBurstSize** | Pointer to **int32** | Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.   | [optional] 
**Req5Gsdelay** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**Priority** | Pointer to **int32** | Represents the priority level of TSC Flows. | [optional] 
**TscaiTimeDom** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**TscaiInputDl** | Pointer to [**NullableTscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 
**TscaiInputUl** | Pointer to [**NullableTscaiInputContainer**](TscaiInputContainer.md) |  | [optional] 

## Methods

### NewTscQosRequirement

`func NewTscQosRequirement() *TscQosRequirement`

NewTscQosRequirement instantiates a new TscQosRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTscQosRequirementWithDefaults

`func NewTscQosRequirementWithDefaults() *TscQosRequirement`

NewTscQosRequirementWithDefaults instantiates a new TscQosRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReqGbrDl

`func (o *TscQosRequirement) GetReqGbrDl() string`

GetReqGbrDl returns the ReqGbrDl field if non-nil, zero value otherwise.

### GetReqGbrDlOk

`func (o *TscQosRequirement) GetReqGbrDlOk() (*string, bool)`

GetReqGbrDlOk returns a tuple with the ReqGbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqGbrDl

`func (o *TscQosRequirement) SetReqGbrDl(v string)`

SetReqGbrDl sets ReqGbrDl field to given value.

### HasReqGbrDl

`func (o *TscQosRequirement) HasReqGbrDl() bool`

HasReqGbrDl returns a boolean if a field has been set.

### GetReqGbrUl

`func (o *TscQosRequirement) GetReqGbrUl() string`

GetReqGbrUl returns the ReqGbrUl field if non-nil, zero value otherwise.

### GetReqGbrUlOk

`func (o *TscQosRequirement) GetReqGbrUlOk() (*string, bool)`

GetReqGbrUlOk returns a tuple with the ReqGbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqGbrUl

`func (o *TscQosRequirement) SetReqGbrUl(v string)`

SetReqGbrUl sets ReqGbrUl field to given value.

### HasReqGbrUl

`func (o *TscQosRequirement) HasReqGbrUl() bool`

HasReqGbrUl returns a boolean if a field has been set.

### GetReqMbrDl

`func (o *TscQosRequirement) GetReqMbrDl() string`

GetReqMbrDl returns the ReqMbrDl field if non-nil, zero value otherwise.

### GetReqMbrDlOk

`func (o *TscQosRequirement) GetReqMbrDlOk() (*string, bool)`

GetReqMbrDlOk returns a tuple with the ReqMbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMbrDl

`func (o *TscQosRequirement) SetReqMbrDl(v string)`

SetReqMbrDl sets ReqMbrDl field to given value.

### HasReqMbrDl

`func (o *TscQosRequirement) HasReqMbrDl() bool`

HasReqMbrDl returns a boolean if a field has been set.

### GetReqMbrUl

`func (o *TscQosRequirement) GetReqMbrUl() string`

GetReqMbrUl returns the ReqMbrUl field if non-nil, zero value otherwise.

### GetReqMbrUlOk

`func (o *TscQosRequirement) GetReqMbrUlOk() (*string, bool)`

GetReqMbrUlOk returns a tuple with the ReqMbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMbrUl

`func (o *TscQosRequirement) SetReqMbrUl(v string)`

SetReqMbrUl sets ReqMbrUl field to given value.

### HasReqMbrUl

`func (o *TscQosRequirement) HasReqMbrUl() bool`

HasReqMbrUl returns a boolean if a field has been set.

### GetMaxTscBurstSize

`func (o *TscQosRequirement) GetMaxTscBurstSize() int32`

GetMaxTscBurstSize returns the MaxTscBurstSize field if non-nil, zero value otherwise.

### GetMaxTscBurstSizeOk

`func (o *TscQosRequirement) GetMaxTscBurstSizeOk() (*int32, bool)`

GetMaxTscBurstSizeOk returns a tuple with the MaxTscBurstSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTscBurstSize

`func (o *TscQosRequirement) SetMaxTscBurstSize(v int32)`

SetMaxTscBurstSize sets MaxTscBurstSize field to given value.

### HasMaxTscBurstSize

`func (o *TscQosRequirement) HasMaxTscBurstSize() bool`

HasMaxTscBurstSize returns a boolean if a field has been set.

### GetReq5Gsdelay

`func (o *TscQosRequirement) GetReq5Gsdelay() int32`

GetReq5Gsdelay returns the Req5Gsdelay field if non-nil, zero value otherwise.

### GetReq5GsdelayOk

`func (o *TscQosRequirement) GetReq5GsdelayOk() (*int32, bool)`

GetReq5GsdelayOk returns a tuple with the Req5Gsdelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReq5Gsdelay

`func (o *TscQosRequirement) SetReq5Gsdelay(v int32)`

SetReq5Gsdelay sets Req5Gsdelay field to given value.

### HasReq5Gsdelay

`func (o *TscQosRequirement) HasReq5Gsdelay() bool`

HasReq5Gsdelay returns a boolean if a field has been set.

### GetPriority

`func (o *TscQosRequirement) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TscQosRequirement) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TscQosRequirement) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TscQosRequirement) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTscaiTimeDom

`func (o *TscQosRequirement) GetTscaiTimeDom() int32`

GetTscaiTimeDom returns the TscaiTimeDom field if non-nil, zero value otherwise.

### GetTscaiTimeDomOk

`func (o *TscQosRequirement) GetTscaiTimeDomOk() (*int32, bool)`

GetTscaiTimeDomOk returns a tuple with the TscaiTimeDom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiTimeDom

`func (o *TscQosRequirement) SetTscaiTimeDom(v int32)`

SetTscaiTimeDom sets TscaiTimeDom field to given value.

### HasTscaiTimeDom

`func (o *TscQosRequirement) HasTscaiTimeDom() bool`

HasTscaiTimeDom returns a boolean if a field has been set.

### GetTscaiInputDl

`func (o *TscQosRequirement) GetTscaiInputDl() TscaiInputContainer`

GetTscaiInputDl returns the TscaiInputDl field if non-nil, zero value otherwise.

### GetTscaiInputDlOk

`func (o *TscQosRequirement) GetTscaiInputDlOk() (*TscaiInputContainer, bool)`

GetTscaiInputDlOk returns a tuple with the TscaiInputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputDl

`func (o *TscQosRequirement) SetTscaiInputDl(v TscaiInputContainer)`

SetTscaiInputDl sets TscaiInputDl field to given value.

### HasTscaiInputDl

`func (o *TscQosRequirement) HasTscaiInputDl() bool`

HasTscaiInputDl returns a boolean if a field has been set.

### SetTscaiInputDlNil

`func (o *TscQosRequirement) SetTscaiInputDlNil(b bool)`

 SetTscaiInputDlNil sets the value for TscaiInputDl to be an explicit nil

### UnsetTscaiInputDl
`func (o *TscQosRequirement) UnsetTscaiInputDl()`

UnsetTscaiInputDl ensures that no value is present for TscaiInputDl, not even an explicit nil
### GetTscaiInputUl

`func (o *TscQosRequirement) GetTscaiInputUl() TscaiInputContainer`

GetTscaiInputUl returns the TscaiInputUl field if non-nil, zero value otherwise.

### GetTscaiInputUlOk

`func (o *TscQosRequirement) GetTscaiInputUlOk() (*TscaiInputContainer, bool)`

GetTscaiInputUlOk returns a tuple with the TscaiInputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTscaiInputUl

`func (o *TscQosRequirement) SetTscaiInputUl(v TscaiInputContainer)`

SetTscaiInputUl sets TscaiInputUl field to given value.

### HasTscaiInputUl

`func (o *TscQosRequirement) HasTscaiInputUl() bool`

HasTscaiInputUl returns a boolean if a field has been set.

### SetTscaiInputUlNil

`func (o *TscQosRequirement) SetTscaiInputUlNil(b bool)`

 SetTscaiInputUlNil sets the value for TscaiInputUl to be an explicit nil

### UnsetTscaiInputUl
`func (o *TscQosRequirement) UnsetTscaiInputUl()`

UnsetTscaiInputUl ensures that no value is present for TscaiInputUl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


