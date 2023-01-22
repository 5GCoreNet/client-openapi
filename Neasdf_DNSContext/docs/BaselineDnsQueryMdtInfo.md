# BaselineDnsQueryMdtInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**SourceIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**BaseDnsMdtList** | [**[]BaselineDnsMdtId**](BaselineDnsMdtId.md) |  | 

## Methods

### NewBaselineDnsQueryMdtInfo

`func NewBaselineDnsQueryMdtInfo(baseDnsMdtList []BaselineDnsMdtId, ) *BaselineDnsQueryMdtInfo`

NewBaselineDnsQueryMdtInfo instantiates a new BaselineDnsQueryMdtInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaselineDnsQueryMdtInfoWithDefaults

`func NewBaselineDnsQueryMdtInfoWithDefaults() *BaselineDnsQueryMdtInfo`

NewBaselineDnsQueryMdtInfoWithDefaults instantiates a new BaselineDnsQueryMdtInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceIpv4Addr

`func (o *BaselineDnsQueryMdtInfo) GetSourceIpv4Addr() string`

GetSourceIpv4Addr returns the SourceIpv4Addr field if non-nil, zero value otherwise.

### GetSourceIpv4AddrOk

`func (o *BaselineDnsQueryMdtInfo) GetSourceIpv4AddrOk() (*string, bool)`

GetSourceIpv4AddrOk returns a tuple with the SourceIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpv4Addr

`func (o *BaselineDnsQueryMdtInfo) SetSourceIpv4Addr(v string)`

SetSourceIpv4Addr sets SourceIpv4Addr field to given value.

### HasSourceIpv4Addr

`func (o *BaselineDnsQueryMdtInfo) HasSourceIpv4Addr() bool`

HasSourceIpv4Addr returns a boolean if a field has been set.

### GetSourceIpv6Prefix

`func (o *BaselineDnsQueryMdtInfo) GetSourceIpv6Prefix() Ipv6Prefix`

GetSourceIpv6Prefix returns the SourceIpv6Prefix field if non-nil, zero value otherwise.

### GetSourceIpv6PrefixOk

`func (o *BaselineDnsQueryMdtInfo) GetSourceIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSourceIpv6PrefixOk returns a tuple with the SourceIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpv6Prefix

`func (o *BaselineDnsQueryMdtInfo) SetSourceIpv6Prefix(v Ipv6Prefix)`

SetSourceIpv6Prefix sets SourceIpv6Prefix field to given value.

### HasSourceIpv6Prefix

`func (o *BaselineDnsQueryMdtInfo) HasSourceIpv6Prefix() bool`

HasSourceIpv6Prefix returns a boolean if a field has been set.

### GetBaseDnsMdtList

`func (o *BaselineDnsQueryMdtInfo) GetBaseDnsMdtList() []BaselineDnsMdtId`

GetBaseDnsMdtList returns the BaseDnsMdtList field if non-nil, zero value otherwise.

### GetBaseDnsMdtListOk

`func (o *BaselineDnsQueryMdtInfo) GetBaseDnsMdtListOk() (*[]BaselineDnsMdtId, bool)`

GetBaseDnsMdtListOk returns a tuple with the BaseDnsMdtList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDnsMdtList

`func (o *BaselineDnsQueryMdtInfo) SetBaseDnsMdtList(v []BaselineDnsMdtId)`

SetBaseDnsMdtList sets BaseDnsMdtList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


