# AfCoordinationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**SourceUeIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**SourceUeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**NotificationInfoList** | Pointer to [**[]NotificationInfo**](NotificationInfo.md) |  | [optional] 

## Methods

### NewAfCoordinationInfo

`func NewAfCoordinationInfo() *AfCoordinationInfo`

NewAfCoordinationInfo instantiates a new AfCoordinationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfCoordinationInfoWithDefaults

`func NewAfCoordinationInfoWithDefaults() *AfCoordinationInfo`

NewAfCoordinationInfoWithDefaults instantiates a new AfCoordinationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDnai

`func (o *AfCoordinationInfo) GetSourceDnai() string`

GetSourceDnai returns the SourceDnai field if non-nil, zero value otherwise.

### GetSourceDnaiOk

`func (o *AfCoordinationInfo) GetSourceDnaiOk() (*string, bool)`

GetSourceDnaiOk returns a tuple with the SourceDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDnai

`func (o *AfCoordinationInfo) SetSourceDnai(v string)`

SetSourceDnai sets SourceDnai field to given value.

### HasSourceDnai

`func (o *AfCoordinationInfo) HasSourceDnai() bool`

HasSourceDnai returns a boolean if a field has been set.

### GetSourceUeIpv4Addr

`func (o *AfCoordinationInfo) GetSourceUeIpv4Addr() string`

GetSourceUeIpv4Addr returns the SourceUeIpv4Addr field if non-nil, zero value otherwise.

### GetSourceUeIpv4AddrOk

`func (o *AfCoordinationInfo) GetSourceUeIpv4AddrOk() (*string, bool)`

GetSourceUeIpv4AddrOk returns a tuple with the SourceUeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUeIpv4Addr

`func (o *AfCoordinationInfo) SetSourceUeIpv4Addr(v string)`

SetSourceUeIpv4Addr sets SourceUeIpv4Addr field to given value.

### HasSourceUeIpv4Addr

`func (o *AfCoordinationInfo) HasSourceUeIpv4Addr() bool`

HasSourceUeIpv4Addr returns a boolean if a field has been set.

### GetSourceUeIpv6Prefix

`func (o *AfCoordinationInfo) GetSourceUeIpv6Prefix() Ipv6Prefix`

GetSourceUeIpv6Prefix returns the SourceUeIpv6Prefix field if non-nil, zero value otherwise.

### GetSourceUeIpv6PrefixOk

`func (o *AfCoordinationInfo) GetSourceUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetSourceUeIpv6PrefixOk returns a tuple with the SourceUeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUeIpv6Prefix

`func (o *AfCoordinationInfo) SetSourceUeIpv6Prefix(v Ipv6Prefix)`

SetSourceUeIpv6Prefix sets SourceUeIpv6Prefix field to given value.

### HasSourceUeIpv6Prefix

`func (o *AfCoordinationInfo) HasSourceUeIpv6Prefix() bool`

HasSourceUeIpv6Prefix returns a boolean if a field has been set.

### GetNotificationInfoList

`func (o *AfCoordinationInfo) GetNotificationInfoList() []NotificationInfo`

GetNotificationInfoList returns the NotificationInfoList field if non-nil, zero value otherwise.

### GetNotificationInfoListOk

`func (o *AfCoordinationInfo) GetNotificationInfoListOk() (*[]NotificationInfo, bool)`

GetNotificationInfoListOk returns a tuple with the NotificationInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationInfoList

`func (o *AfCoordinationInfo) SetNotificationInfoList(v []NotificationInfo)`

SetNotificationInfoList sets NotificationInfoList field to given value.

### HasNotificationInfoList

`func (o *AfCoordinationInfo) HasNotificationInfoList() bool`

HasNotificationInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


