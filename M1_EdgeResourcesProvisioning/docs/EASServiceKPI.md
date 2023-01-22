# EASServiceKPI

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxReqRate** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MaxRespTime** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Avail** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AvlComp** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AvlGraComp** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AvlMem** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**AvlStrg** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**ConnBand** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewEASServiceKPI

`func NewEASServiceKPI() *EASServiceKPI`

NewEASServiceKPI instantiates a new EASServiceKPI object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEASServiceKPIWithDefaults

`func NewEASServiceKPIWithDefaults() *EASServiceKPI`

NewEASServiceKPIWithDefaults instantiates a new EASServiceKPI object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxReqRate

`func (o *EASServiceKPI) GetMaxReqRate() int32`

GetMaxReqRate returns the MaxReqRate field if non-nil, zero value otherwise.

### GetMaxReqRateOk

`func (o *EASServiceKPI) GetMaxReqRateOk() (*int32, bool)`

GetMaxReqRateOk returns a tuple with the MaxReqRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReqRate

`func (o *EASServiceKPI) SetMaxReqRate(v int32)`

SetMaxReqRate sets MaxReqRate field to given value.

### HasMaxReqRate

`func (o *EASServiceKPI) HasMaxReqRate() bool`

HasMaxReqRate returns a boolean if a field has been set.

### GetMaxRespTime

`func (o *EASServiceKPI) GetMaxRespTime() int32`

GetMaxRespTime returns the MaxRespTime field if non-nil, zero value otherwise.

### GetMaxRespTimeOk

`func (o *EASServiceKPI) GetMaxRespTimeOk() (*int32, bool)`

GetMaxRespTimeOk returns a tuple with the MaxRespTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRespTime

`func (o *EASServiceKPI) SetMaxRespTime(v int32)`

SetMaxRespTime sets MaxRespTime field to given value.

### HasMaxRespTime

`func (o *EASServiceKPI) HasMaxRespTime() bool`

HasMaxRespTime returns a boolean if a field has been set.

### GetAvail

`func (o *EASServiceKPI) GetAvail() int32`

GetAvail returns the Avail field if non-nil, zero value otherwise.

### GetAvailOk

`func (o *EASServiceKPI) GetAvailOk() (*int32, bool)`

GetAvailOk returns a tuple with the Avail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvail

`func (o *EASServiceKPI) SetAvail(v int32)`

SetAvail sets Avail field to given value.

### HasAvail

`func (o *EASServiceKPI) HasAvail() bool`

HasAvail returns a boolean if a field has been set.

### GetAvlComp

`func (o *EASServiceKPI) GetAvlComp() int32`

GetAvlComp returns the AvlComp field if non-nil, zero value otherwise.

### GetAvlCompOk

`func (o *EASServiceKPI) GetAvlCompOk() (*int32, bool)`

GetAvlCompOk returns a tuple with the AvlComp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvlComp

`func (o *EASServiceKPI) SetAvlComp(v int32)`

SetAvlComp sets AvlComp field to given value.

### HasAvlComp

`func (o *EASServiceKPI) HasAvlComp() bool`

HasAvlComp returns a boolean if a field has been set.

### GetAvlGraComp

`func (o *EASServiceKPI) GetAvlGraComp() int32`

GetAvlGraComp returns the AvlGraComp field if non-nil, zero value otherwise.

### GetAvlGraCompOk

`func (o *EASServiceKPI) GetAvlGraCompOk() (*int32, bool)`

GetAvlGraCompOk returns a tuple with the AvlGraComp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvlGraComp

`func (o *EASServiceKPI) SetAvlGraComp(v int32)`

SetAvlGraComp sets AvlGraComp field to given value.

### HasAvlGraComp

`func (o *EASServiceKPI) HasAvlGraComp() bool`

HasAvlGraComp returns a boolean if a field has been set.

### GetAvlMem

`func (o *EASServiceKPI) GetAvlMem() int32`

GetAvlMem returns the AvlMem field if non-nil, zero value otherwise.

### GetAvlMemOk

`func (o *EASServiceKPI) GetAvlMemOk() (*int32, bool)`

GetAvlMemOk returns a tuple with the AvlMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvlMem

`func (o *EASServiceKPI) SetAvlMem(v int32)`

SetAvlMem sets AvlMem field to given value.

### HasAvlMem

`func (o *EASServiceKPI) HasAvlMem() bool`

HasAvlMem returns a boolean if a field has been set.

### GetAvlStrg

`func (o *EASServiceKPI) GetAvlStrg() int32`

GetAvlStrg returns the AvlStrg field if non-nil, zero value otherwise.

### GetAvlStrgOk

`func (o *EASServiceKPI) GetAvlStrgOk() (*int32, bool)`

GetAvlStrgOk returns a tuple with the AvlStrg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvlStrg

`func (o *EASServiceKPI) SetAvlStrg(v int32)`

SetAvlStrg sets AvlStrg field to given value.

### HasAvlStrg

`func (o *EASServiceKPI) HasAvlStrg() bool`

HasAvlStrg returns a boolean if a field has been set.

### GetConnBand

`func (o *EASServiceKPI) GetConnBand() string`

GetConnBand returns the ConnBand field if non-nil, zero value otherwise.

### GetConnBandOk

`func (o *EASServiceKPI) GetConnBandOk() (*string, bool)`

GetConnBandOk returns a tuple with the ConnBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnBand

`func (o *EASServiceKPI) SetConnBand(v string)`

SetConnBand sets ConnBand field to given value.

### HasConnBand

`func (o *EASServiceKPI) HasConnBand() bool`

HasConnBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


