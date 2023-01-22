# ACServiceKPIs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnBand** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ReqRate** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RespTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**Avail** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**ReqComp** | Pointer to **string** | The compute resources required by the AC. | [optional] 
**ReqGrapComp** | Pointer to **string** | The graphical compute resources required by the AC. | [optional] 
**ReqMem** | Pointer to **string** | The memory resources required by the AC. | [optional] 
**ReqStrg** | Pointer to **string** | The storage resources required by the AC. | [optional] 

## Methods

### NewACServiceKPIs

`func NewACServiceKPIs() *ACServiceKPIs`

NewACServiceKPIs instantiates a new ACServiceKPIs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACServiceKPIsWithDefaults

`func NewACServiceKPIsWithDefaults() *ACServiceKPIs`

NewACServiceKPIsWithDefaults instantiates a new ACServiceKPIs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnBand

`func (o *ACServiceKPIs) GetConnBand() string`

GetConnBand returns the ConnBand field if non-nil, zero value otherwise.

### GetConnBandOk

`func (o *ACServiceKPIs) GetConnBandOk() (*string, bool)`

GetConnBandOk returns a tuple with the ConnBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnBand

`func (o *ACServiceKPIs) SetConnBand(v string)`

SetConnBand sets ConnBand field to given value.

### HasConnBand

`func (o *ACServiceKPIs) HasConnBand() bool`

HasConnBand returns a boolean if a field has been set.

### GetReqRate

`func (o *ACServiceKPIs) GetReqRate() int32`

GetReqRate returns the ReqRate field if non-nil, zero value otherwise.

### GetReqRateOk

`func (o *ACServiceKPIs) GetReqRateOk() (*int32, bool)`

GetReqRateOk returns a tuple with the ReqRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqRate

`func (o *ACServiceKPIs) SetReqRate(v int32)`

SetReqRate sets ReqRate field to given value.

### HasReqRate

`func (o *ACServiceKPIs) HasReqRate() bool`

HasReqRate returns a boolean if a field has been set.

### GetRespTime

`func (o *ACServiceKPIs) GetRespTime() int32`

GetRespTime returns the RespTime field if non-nil, zero value otherwise.

### GetRespTimeOk

`func (o *ACServiceKPIs) GetRespTimeOk() (*int32, bool)`

GetRespTimeOk returns a tuple with the RespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespTime

`func (o *ACServiceKPIs) SetRespTime(v int32)`

SetRespTime sets RespTime field to given value.

### HasRespTime

`func (o *ACServiceKPIs) HasRespTime() bool`

HasRespTime returns a boolean if a field has been set.

### GetAvail

`func (o *ACServiceKPIs) GetAvail() int32`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *ACServiceKPIs) GetAvailOk() (*int32, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *ACServiceKPIs) SetAvail(v int32)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *ACServiceKPIs) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetReqComp

`func (o *ACServiceKPIs) GetReqComp() string`

GetReqComp returns the ReqComp field if non-nil, zero value otherwise.

### GetReqCompOk

`func (o *ACServiceKPIs) GetReqCompOk() (*string, bool)`

GetReqCompOk returns a tuple with the ReqComp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqComp

`func (o *ACServiceKPIs) SetReqComp(v string)`

SetReqComp sets ReqComp field to given value.

### HasReqComp

`func (o *ACServiceKPIs) HasReqComp() bool`

HasReqComp returns a boolean if a field has been set.

### GetReqGrapComp

`func (o *ACServiceKPIs) GetReqGrapComp() string`

GetReqGrapComp returns the ReqGrapComp field if non-nil, zero value otherwise.

### GetReqGrapCompOk

`func (o *ACServiceKPIs) GetReqGrapCompOk() (*string, bool)`

GetReqGrapCompOk returns a tuple with the ReqGrapComp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqGrapComp

`func (o *ACServiceKPIs) SetReqGrapComp(v string)`

SetReqGrapComp sets ReqGrapComp field to given value.

### HasReqGrapComp

`func (o *ACServiceKPIs) HasReqGrapComp() bool`

HasReqGrapComp returns a boolean if a field has been set.

### GetReqMem

`func (o *ACServiceKPIs) GetReqMem() string`

GetReqMem returns the ReqMem field if non-nil, zero value otherwise.

### GetReqMemOk

`func (o *ACServiceKPIs) GetReqMemOk() (*string, bool)`

GetReqMemOk returns a tuple with the ReqMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqMem

`func (o *ACServiceKPIs) SetReqMem(v string)`

SetReqMem sets ReqMem field to given value.

### HasReqMem

`func (o *ACServiceKPIs) HasReqMem() bool`

HasReqMem returns a boolean if a field has been set.

### GetReqStrg

`func (o *ACServiceKPIs) GetReqStrg() string`

GetReqStrg returns the ReqStrg field if non-nil, zero value otherwise.

### GetReqStrgOk

`func (o *ACServiceKPIs) GetReqStrgOk() (*string, bool)`

GetReqStrgOk returns a tuple with the ReqStrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqStrg

`func (o *ACServiceKPIs) SetReqStrg(v string)`

SetReqStrg sets ReqStrg field to given value.

### HasReqStrg

`func (o *ACServiceKPIs) HasReqStrg() bool`

HasReqStrg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


