# MbsQoSReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qi** | **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.  | 
**GuarBitRate** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxBitRate** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**AverWindow** | Pointer to **int32** | Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] [default to 2000]
**ReqMbsArp** | Pointer to [**Arp**](Arp.md) |  | [optional] 

## Methods

### NewMbsQoSReq

`func NewMbsQoSReq(var5qi int32, ) *MbsQoSReq`

NewMbsQoSReq instantiates a new MbsQoSReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsQoSReqWithDefaults

`func NewMbsQoSReqWithDefaults() *MbsQoSReq`

NewMbsQoSReqWithDefaults instantiates a new MbsQoSReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qi

`func (o *MbsQoSReq) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *MbsQoSReq) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *MbsQoSReq) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.


### GetGuarBitRate

`func (o *MbsQoSReq) GetGuarBitRate() string`

GetGuarBitRate returns the GuarBitRate field if non-nil, zero value otherwise.

### GetGuarBitRateOk

`func (o *MbsQoSReq) GetGuarBitRateOk() (*string, bool)`

GetGuarBitRateOk returns a tuple with the GuarBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuarBitRate

`func (o *MbsQoSReq) SetGuarBitRate(v string)`

SetGuarBitRate sets GuarBitRate field to given value.

### HasGuarBitRate

`func (o *MbsQoSReq) HasGuarBitRate() bool`

HasGuarBitRate returns a boolean if a field has been set.

### GetMaxBitRate

`func (o *MbsQoSReq) GetMaxBitRate() string`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *MbsQoSReq) GetMaxBitRateOk() (*string, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *MbsQoSReq) SetMaxBitRate(v string)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *MbsQoSReq) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetAverWindow

`func (o *MbsQoSReq) GetAverWindow() int32`

GetAverWindow returns the AverWindow field if non-nil, zero value otherwise.

### GetAverWindowOk

`func (o *MbsQoSReq) GetAverWindowOk() (*int32, bool)`

GetAverWindowOk returns a tuple with the AverWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverWindow

`func (o *MbsQoSReq) SetAverWindow(v int32)`

SetAverWindow sets AverWindow field to given value.

### HasAverWindow

`func (o *MbsQoSReq) HasAverWindow() bool`

HasAverWindow returns a boolean if a field has been set.

### GetReqMbsArp

`func (o *MbsQoSReq) GetReqMbsArp() Arp`

GetReqMbsArp returns the ReqMbsArp field if non-nil, zero value otherwise.

### GetReqMbsArpOk

`func (o *MbsQoSReq) GetReqMbsArpOk() (*Arp, bool)`

GetReqMbsArpOk returns a tuple with the ReqMbsArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMbsArp

`func (o *MbsQoSReq) SetReqMbsArp(v Arp)`

SetReqMbsArp sets ReqMbsArp field to given value.

### HasReqMbsArp

`func (o *MbsQoSReq) HasReqMbsArp() bool`

HasReqMbsArp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


