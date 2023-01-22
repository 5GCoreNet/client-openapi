# MbsQosDec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsQosId** | **string** |  | 
**Var5qi** | Pointer to **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.  | [optional] 
**PriorityLevel** | Pointer to **int32** | Unsigned integer indicating the 5QI Priority Level (see clauses 5.7.3.3 and 5.7.4 of 3GPP TS 23.501, within the range 1 to 127.Values are ordered in decreasing order of priority,  i.e. with 1 as the highest priority and 127 as the lowest priority.   | [optional] 
**MbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**GbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**AverWindow** | Pointer to **int32** | Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] [default to 2000]
**MbsMaxDataBurstVol** | Pointer to **int32** | Represents MBS Maximum Data Burst Volume, expressed in Bytes. | [optional] 

## Methods

### NewMbsQosDec

`func NewMbsQosDec(mbsQosId string, ) *MbsQosDec`

NewMbsQosDec instantiates a new MbsQosDec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsQosDecWithDefaults

`func NewMbsQosDecWithDefaults() *MbsQosDec`

NewMbsQosDecWithDefaults instantiates a new MbsQosDec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsQosId

`func (o *MbsQosDec) GetMbsQosId() string`

GetMbsQosId returns the MbsQosId field if non-nil, zero value otherwise.

### GetMbsQosIdOk

`func (o *MbsQosDec) GetMbsQosIdOk() (*string, bool)`

GetMbsQosIdOk returns a tuple with the MbsQosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsQosId

`func (o *MbsQosDec) SetMbsQosId(v string)`

SetMbsQosId sets MbsQosId field to given value.


### GetVar5qi

`func (o *MbsQosDec) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *MbsQosDec) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *MbsQosDec) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.

### HasVar5qi

`func (o *MbsQosDec) HasVar5qi() bool`

HasVar5qi returns a boolean if a field has been set.

### GetPriorityLevel

`func (o *MbsQosDec) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *MbsQosDec) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *MbsQosDec) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *MbsQosDec) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.

### GetMbrDl

`func (o *MbsQosDec) GetMbrDl() string`

GetMbrDl returns the MbrDl field if non-nil, zero value otherwise.

### GetMbrDlOk

`func (o *MbsQosDec) GetMbrDlOk() (*string, bool)`

GetMbrDlOk returns a tuple with the MbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbrDl

`func (o *MbsQosDec) SetMbrDl(v string)`

SetMbrDl sets MbrDl field to given value.

### HasMbrDl

`func (o *MbsQosDec) HasMbrDl() bool`

HasMbrDl returns a boolean if a field has been set.

### GetGbrDl

`func (o *MbsQosDec) GetGbrDl() string`

GetGbrDl returns the GbrDl field if non-nil, zero value otherwise.

### GetGbrDlOk

`func (o *MbsQosDec) GetGbrDlOk() (*string, bool)`

GetGbrDlOk returns a tuple with the GbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrDl

`func (o *MbsQosDec) SetGbrDl(v string)`

SetGbrDl sets GbrDl field to given value.

### HasGbrDl

`func (o *MbsQosDec) HasGbrDl() bool`

HasGbrDl returns a boolean if a field has been set.

### GetArp

`func (o *MbsQosDec) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *MbsQosDec) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *MbsQosDec) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *MbsQosDec) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetAverWindow

`func (o *MbsQosDec) GetAverWindow() int32`

GetAverWindow returns the AverWindow field if non-nil, zero value otherwise.

### GetAverWindowOk

`func (o *MbsQosDec) GetAverWindowOk() (*int32, bool)`

GetAverWindowOk returns a tuple with the AverWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverWindow

`func (o *MbsQosDec) SetAverWindow(v int32)`

SetAverWindow sets AverWindow field to given value.

### HasAverWindow

`func (o *MbsQosDec) HasAverWindow() bool`

HasAverWindow returns a boolean if a field has been set.

### GetMbsMaxDataBurstVol

`func (o *MbsQosDec) GetMbsMaxDataBurstVol() int32`

GetMbsMaxDataBurstVol returns the MbsMaxDataBurstVol field if non-nil, zero value otherwise.

### GetMbsMaxDataBurstVolOk

`func (o *MbsQosDec) GetMbsMaxDataBurstVolOk() (*int32, bool)`

GetMbsMaxDataBurstVolOk returns a tuple with the MbsMaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsMaxDataBurstVol

`func (o *MbsQosDec) SetMbsMaxDataBurstVol(v int32)`

SetMbsMaxDataBurstVol sets MbsMaxDataBurstVol field to given value.

### HasMbsMaxDataBurstVol

`func (o *MbsQosDec) HasMbsMaxDataBurstVol() bool`

HasMbsMaxDataBurstVol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


