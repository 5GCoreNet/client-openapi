# ObservedRedundantTransExp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgPktDropRateUl** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 
**VarPktDropRateUl** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**AvgPktDropRateDl** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 
**VarPktDropRateDl** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**AvgPktDelayUl** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**VarPktDelayUl** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 
**AvgPktDelayDl** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**VarPktDelayDl** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewObservedRedundantTransExp

`func NewObservedRedundantTransExp() *ObservedRedundantTransExp`

NewObservedRedundantTransExp instantiates a new ObservedRedundantTransExp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservedRedundantTransExpWithDefaults

`func NewObservedRedundantTransExpWithDefaults() *ObservedRedundantTransExp`

NewObservedRedundantTransExpWithDefaults instantiates a new ObservedRedundantTransExp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgPktDropRateUl

`func (o *ObservedRedundantTransExp) GetAvgPktDropRateUl() int32`

GetAvgPktDropRateUl returns the AvgPktDropRateUl field if non-nil, zero value otherwise.

### GetAvgPktDropRateUlOk

`func (o *ObservedRedundantTransExp) GetAvgPktDropRateUlOk() (*int32, bool)`

GetAvgPktDropRateUlOk returns a tuple with the AvgPktDropRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPktDropRateUl

`func (o *ObservedRedundantTransExp) SetAvgPktDropRateUl(v int32)`

SetAvgPktDropRateUl sets AvgPktDropRateUl field to given value.

### HasAvgPktDropRateUl

`func (o *ObservedRedundantTransExp) HasAvgPktDropRateUl() bool`

HasAvgPktDropRateUl returns a boolean if a field has been set.

### GetVarPktDropRateUl

`func (o *ObservedRedundantTransExp) GetVarPktDropRateUl() float32`

GetVarPktDropRateUl returns the VarPktDropRateUl field if non-nil, zero value otherwise.

### GetVarPktDropRateUlOk

`func (o *ObservedRedundantTransExp) GetVarPktDropRateUlOk() (*float32, bool)`

GetVarPktDropRateUlOk returns a tuple with the VarPktDropRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarPktDropRateUl

`func (o *ObservedRedundantTransExp) SetVarPktDropRateUl(v float32)`

SetVarPktDropRateUl sets VarPktDropRateUl field to given value.

### HasVarPktDropRateUl

`func (o *ObservedRedundantTransExp) HasVarPktDropRateUl() bool`

HasVarPktDropRateUl returns a boolean if a field has been set.

### GetAvgPktDropRateDl

`func (o *ObservedRedundantTransExp) GetAvgPktDropRateDl() int32`

GetAvgPktDropRateDl returns the AvgPktDropRateDl field if non-nil, zero value otherwise.

### GetAvgPktDropRateDlOk

`func (o *ObservedRedundantTransExp) GetAvgPktDropRateDlOk() (*int32, bool)`

GetAvgPktDropRateDlOk returns a tuple with the AvgPktDropRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPktDropRateDl

`func (o *ObservedRedundantTransExp) SetAvgPktDropRateDl(v int32)`

SetAvgPktDropRateDl sets AvgPktDropRateDl field to given value.

### HasAvgPktDropRateDl

`func (o *ObservedRedundantTransExp) HasAvgPktDropRateDl() bool`

HasAvgPktDropRateDl returns a boolean if a field has been set.

### GetVarPktDropRateDl

`func (o *ObservedRedundantTransExp) GetVarPktDropRateDl() float32`

GetVarPktDropRateDl returns the VarPktDropRateDl field if non-nil, zero value otherwise.

### GetVarPktDropRateDlOk

`func (o *ObservedRedundantTransExp) GetVarPktDropRateDlOk() (*float32, bool)`

GetVarPktDropRateDlOk returns a tuple with the VarPktDropRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarPktDropRateDl

`func (o *ObservedRedundantTransExp) SetVarPktDropRateDl(v float32)`

SetVarPktDropRateDl sets VarPktDropRateDl field to given value.

### HasVarPktDropRateDl

`func (o *ObservedRedundantTransExp) HasVarPktDropRateDl() bool`

HasVarPktDropRateDl returns a boolean if a field has been set.

### GetAvgPktDelayUl

`func (o *ObservedRedundantTransExp) GetAvgPktDelayUl() int32`

GetAvgPktDelayUl returns the AvgPktDelayUl field if non-nil, zero value otherwise.

### GetAvgPktDelayUlOk

`func (o *ObservedRedundantTransExp) GetAvgPktDelayUlOk() (*int32, bool)`

GetAvgPktDelayUlOk returns a tuple with the AvgPktDelayUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPktDelayUl

`func (o *ObservedRedundantTransExp) SetAvgPktDelayUl(v int32)`

SetAvgPktDelayUl sets AvgPktDelayUl field to given value.

### HasAvgPktDelayUl

`func (o *ObservedRedundantTransExp) HasAvgPktDelayUl() bool`

HasAvgPktDelayUl returns a boolean if a field has been set.

### GetVarPktDelayUl

`func (o *ObservedRedundantTransExp) GetVarPktDelayUl() float32`

GetVarPktDelayUl returns the VarPktDelayUl field if non-nil, zero value otherwise.

### GetVarPktDelayUlOk

`func (o *ObservedRedundantTransExp) GetVarPktDelayUlOk() (*float32, bool)`

GetVarPktDelayUlOk returns a tuple with the VarPktDelayUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarPktDelayUl

`func (o *ObservedRedundantTransExp) SetVarPktDelayUl(v float32)`

SetVarPktDelayUl sets VarPktDelayUl field to given value.

### HasVarPktDelayUl

`func (o *ObservedRedundantTransExp) HasVarPktDelayUl() bool`

HasVarPktDelayUl returns a boolean if a field has been set.

### GetAvgPktDelayDl

`func (o *ObservedRedundantTransExp) GetAvgPktDelayDl() int32`

GetAvgPktDelayDl returns the AvgPktDelayDl field if non-nil, zero value otherwise.

### GetAvgPktDelayDlOk

`func (o *ObservedRedundantTransExp) GetAvgPktDelayDlOk() (*int32, bool)`

GetAvgPktDelayDlOk returns a tuple with the AvgPktDelayDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPktDelayDl

`func (o *ObservedRedundantTransExp) SetAvgPktDelayDl(v int32)`

SetAvgPktDelayDl sets AvgPktDelayDl field to given value.

### HasAvgPktDelayDl

`func (o *ObservedRedundantTransExp) HasAvgPktDelayDl() bool`

HasAvgPktDelayDl returns a boolean if a field has been set.

### GetVarPktDelayDl

`func (o *ObservedRedundantTransExp) GetVarPktDelayDl() float32`

GetVarPktDelayDl returns the VarPktDelayDl field if non-nil, zero value otherwise.

### GetVarPktDelayDlOk

`func (o *ObservedRedundantTransExp) GetVarPktDelayDlOk() (*float32, bool)`

GetVarPktDelayDlOk returns a tuple with the VarPktDelayDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarPktDelayDl

`func (o *ObservedRedundantTransExp) SetVarPktDelayDl(v float32)`

SetVarPktDelayDl sets VarPktDelayDl field to given value.

### HasVarPktDelayDl

`func (o *ObservedRedundantTransExp) HasVarPktDelayDl() bool`

HasVarPktDelayDl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


