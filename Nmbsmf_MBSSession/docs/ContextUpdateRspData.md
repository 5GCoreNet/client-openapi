# ContextUpdateRspData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LlSsm** | Pointer to [**Ssm**](Ssm.md) |  | [optional] 
**CTeid** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**N2MbsSmInfo** | Pointer to [**N2MbsSmInfo**](N2MbsSmInfo.md) |  | [optional] 

## Methods

### NewContextUpdateRspData

`func NewContextUpdateRspData() *ContextUpdateRspData`

NewContextUpdateRspData instantiates a new ContextUpdateRspData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextUpdateRspDataWithDefaults

`func NewContextUpdateRspDataWithDefaults() *ContextUpdateRspData`

NewContextUpdateRspDataWithDefaults instantiates a new ContextUpdateRspData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLlSsm

`func (o *ContextUpdateRspData) GetLlSsm() Ssm`

GetLlSsm returns the LlSsm field if non-nil, zero value otherwise.

### GetLlSsmOk

`func (o *ContextUpdateRspData) GetLlSsmOk() (*Ssm, bool)`

GetLlSsmOk returns a tuple with the LlSsm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlSsm

`func (o *ContextUpdateRspData) SetLlSsm(v Ssm)`

SetLlSsm sets LlSsm field to given value.

### HasLlSsm

`func (o *ContextUpdateRspData) HasLlSsm() bool`

HasLlSsm returns a boolean if a field has been set.

### GetCTeid

`func (o *ContextUpdateRspData) GetCTeid() int32`

GetCTeid returns the CTeid field if non-nil, zero value otherwise.

### GetCTeidOk

`func (o *ContextUpdateRspData) GetCTeidOk() (*int32, bool)`

GetCTeidOk returns a tuple with the CTeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCTeid

`func (o *ContextUpdateRspData) SetCTeid(v int32)`

SetCTeid sets CTeid field to given value.

### HasCTeid

`func (o *ContextUpdateRspData) HasCTeid() bool`

HasCTeid returns a boolean if a field has been set.

### GetN2MbsSmInfo

`func (o *ContextUpdateRspData) GetN2MbsSmInfo() N2MbsSmInfo`

GetN2MbsSmInfo returns the N2MbsSmInfo field if non-nil, zero value otherwise.

### GetN2MbsSmInfoOk

`func (o *ContextUpdateRspData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool)`

GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfo

`func (o *ContextUpdateRspData) SetN2MbsSmInfo(v N2MbsSmInfo)`

SetN2MbsSmInfo sets N2MbsSmInfo field to given value.

### HasN2MbsSmInfo

`func (o *ContextUpdateRspData) HasN2MbsSmInfo() bool`

HasN2MbsSmInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


