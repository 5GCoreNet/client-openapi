# LocalMbmsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbmsEnbIpv4MulAddr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**MbmsEnbIpv6MulAddr** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**MbmsGwIpv4SsmAddr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**MbmsGwIpv6SsmAddr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**Cteid** | Pointer to **string** |  | [optional] 
**BmscIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**BmscIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**BmscPort** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewLocalMbmsInfo

`func NewLocalMbmsInfo() *LocalMbmsInfo`

NewLocalMbmsInfo instantiates a new LocalMbmsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalMbmsInfoWithDefaults

`func NewLocalMbmsInfoWithDefaults() *LocalMbmsInfo`

NewLocalMbmsInfoWithDefaults instantiates a new LocalMbmsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbmsEnbIpv4MulAddr

`func (o *LocalMbmsInfo) GetMbmsEnbIpv4MulAddr() string`

GetMbmsEnbIpv4MulAddr returns the MbmsEnbIpv4MulAddr field if non-nil, zero value otherwise.

### GetMbmsEnbIpv4MulAddrOk

`func (o *LocalMbmsInfo) GetMbmsEnbIpv4MulAddrOk() (*string, bool)`

GetMbmsEnbIpv4MulAddrOk returns a tuple with the MbmsEnbIpv4MulAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsEnbIpv4MulAddr

`func (o *LocalMbmsInfo) SetMbmsEnbIpv4MulAddr(v string)`

SetMbmsEnbIpv4MulAddr sets MbmsEnbIpv4MulAddr field to given value.

### HasMbmsEnbIpv4MulAddr

`func (o *LocalMbmsInfo) HasMbmsEnbIpv4MulAddr() bool`

HasMbmsEnbIpv4MulAddr returns a boolean if a field has been set.

### GetMbmsEnbIpv6MulAddr

`func (o *LocalMbmsInfo) GetMbmsEnbIpv6MulAddr() Ipv6Prefix`

GetMbmsEnbIpv6MulAddr returns the MbmsEnbIpv6MulAddr field if non-nil, zero value otherwise.

### GetMbmsEnbIpv6MulAddrOk

`func (o *LocalMbmsInfo) GetMbmsEnbIpv6MulAddrOk() (*Ipv6Prefix, bool)`

GetMbmsEnbIpv6MulAddrOk returns a tuple with the MbmsEnbIpv6MulAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsEnbIpv6MulAddr

`func (o *LocalMbmsInfo) SetMbmsEnbIpv6MulAddr(v Ipv6Prefix)`

SetMbmsEnbIpv6MulAddr sets MbmsEnbIpv6MulAddr field to given value.

### HasMbmsEnbIpv6MulAddr

`func (o *LocalMbmsInfo) HasMbmsEnbIpv6MulAddr() bool`

HasMbmsEnbIpv6MulAddr returns a boolean if a field has been set.

### GetMbmsGwIpv4SsmAddr

`func (o *LocalMbmsInfo) GetMbmsGwIpv4SsmAddr() string`

GetMbmsGwIpv4SsmAddr returns the MbmsGwIpv4SsmAddr field if non-nil, zero value otherwise.

### GetMbmsGwIpv4SsmAddrOk

`func (o *LocalMbmsInfo) GetMbmsGwIpv4SsmAddrOk() (*string, bool)`

GetMbmsGwIpv4SsmAddrOk returns a tuple with the MbmsGwIpv4SsmAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsGwIpv4SsmAddr

`func (o *LocalMbmsInfo) SetMbmsGwIpv4SsmAddr(v string)`

SetMbmsGwIpv4SsmAddr sets MbmsGwIpv4SsmAddr field to given value.

### HasMbmsGwIpv4SsmAddr

`func (o *LocalMbmsInfo) HasMbmsGwIpv4SsmAddr() bool`

HasMbmsGwIpv4SsmAddr returns a boolean if a field has been set.

### GetMbmsGwIpv6SsmAddr

`func (o *LocalMbmsInfo) GetMbmsGwIpv6SsmAddr() Ipv6Addr`

GetMbmsGwIpv6SsmAddr returns the MbmsGwIpv6SsmAddr field if non-nil, zero value otherwise.

### GetMbmsGwIpv6SsmAddrOk

`func (o *LocalMbmsInfo) GetMbmsGwIpv6SsmAddrOk() (*Ipv6Addr, bool)`

GetMbmsGwIpv6SsmAddrOk returns a tuple with the MbmsGwIpv6SsmAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsGwIpv6SsmAddr

`func (o *LocalMbmsInfo) SetMbmsGwIpv6SsmAddr(v Ipv6Addr)`

SetMbmsGwIpv6SsmAddr sets MbmsGwIpv6SsmAddr field to given value.

### HasMbmsGwIpv6SsmAddr

`func (o *LocalMbmsInfo) HasMbmsGwIpv6SsmAddr() bool`

HasMbmsGwIpv6SsmAddr returns a boolean if a field has been set.

### GetCteid

`func (o *LocalMbmsInfo) GetCteid() string`

GetCteid returns the Cteid field if non-nil, zero value otherwise.

### GetCteidOk

`func (o *LocalMbmsInfo) GetCteidOk() (*string, bool)`

GetCteidOk returns a tuple with the Cteid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCteid

`func (o *LocalMbmsInfo) SetCteid(v string)`

SetCteid sets Cteid field to given value.

### HasCteid

`func (o *LocalMbmsInfo) HasCteid() bool`

HasCteid returns a boolean if a field has been set.

### GetBmscIpv4Addr

`func (o *LocalMbmsInfo) GetBmscIpv4Addr() string`

GetBmscIpv4Addr returns the BmscIpv4Addr field if non-nil, zero value otherwise.

### GetBmscIpv4AddrOk

`func (o *LocalMbmsInfo) GetBmscIpv4AddrOk() (*string, bool)`

GetBmscIpv4AddrOk returns a tuple with the BmscIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmscIpv4Addr

`func (o *LocalMbmsInfo) SetBmscIpv4Addr(v string)`

SetBmscIpv4Addr sets BmscIpv4Addr field to given value.

### HasBmscIpv4Addr

`func (o *LocalMbmsInfo) HasBmscIpv4Addr() bool`

HasBmscIpv4Addr returns a boolean if a field has been set.

### GetBmscIpv6Addr

`func (o *LocalMbmsInfo) GetBmscIpv6Addr() Ipv6Addr`

GetBmscIpv6Addr returns the BmscIpv6Addr field if non-nil, zero value otherwise.

### GetBmscIpv6AddrOk

`func (o *LocalMbmsInfo) GetBmscIpv6AddrOk() (*Ipv6Addr, bool)`

GetBmscIpv6AddrOk returns a tuple with the BmscIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmscIpv6Addr

`func (o *LocalMbmsInfo) SetBmscIpv6Addr(v Ipv6Addr)`

SetBmscIpv6Addr sets BmscIpv6Addr field to given value.

### HasBmscIpv6Addr

`func (o *LocalMbmsInfo) HasBmscIpv6Addr() bool`

HasBmscIpv6Addr returns a boolean if a field has been set.

### GetBmscPort

`func (o *LocalMbmsInfo) GetBmscPort() int32`

GetBmscPort returns the BmscPort field if non-nil, zero value otherwise.

### GetBmscPortOk

`func (o *LocalMbmsInfo) GetBmscPortOk() (*int32, bool)`

GetBmscPortOk returns a tuple with the BmscPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmscPort

`func (o *LocalMbmsInfo) SetBmscPort(v int32)`

SetBmscPort sets BmscPort field to given value.

### HasBmscPort

`func (o *LocalMbmsInfo) HasBmscPort() bool`

HasBmscPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


