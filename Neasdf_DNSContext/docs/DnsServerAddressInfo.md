# DnsServerAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServerAddressList** | Pointer to [**[]IpAddr**](IpAddr.md) |  | [optional] 
**BaseDnsAitId** | Pointer to [**BaselineDnsAitId**](BaselineDnsAitId.md) |  | [optional] 

## Methods

### NewDnsServerAddressInfo

`func NewDnsServerAddressInfo() *DnsServerAddressInfo`

NewDnsServerAddressInfo instantiates a new DnsServerAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsServerAddressInfoWithDefaults

`func NewDnsServerAddressInfoWithDefaults() *DnsServerAddressInfo`

NewDnsServerAddressInfoWithDefaults instantiates a new DnsServerAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServerAddressList

`func (o *DnsServerAddressInfo) GetDnsServerAddressList() []IpAddr`

GetDnsServerAddressList returns the DnsServerAddressList field if non-nil, zero value otherwise.

### GetDnsServerAddressListOk

`func (o *DnsServerAddressInfo) GetDnsServerAddressListOk() (*[]IpAddr, bool)`

GetDnsServerAddressListOk returns a tuple with the DnsServerAddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerAddressList

`func (o *DnsServerAddressInfo) SetDnsServerAddressList(v []IpAddr)`

SetDnsServerAddressList sets DnsServerAddressList field to given value.

### HasDnsServerAddressList

`func (o *DnsServerAddressInfo) HasDnsServerAddressList() bool`

HasDnsServerAddressList returns a boolean if a field has been set.

### GetBaseDnsAitId

`func (o *DnsServerAddressInfo) GetBaseDnsAitId() BaselineDnsAitId`

GetBaseDnsAitId returns the BaseDnsAitId field if non-nil, zero value otherwise.

### GetBaseDnsAitIdOk

`func (o *DnsServerAddressInfo) GetBaseDnsAitIdOk() (*BaselineDnsAitId, bool)`

GetBaseDnsAitIdOk returns a tuple with the BaseDnsAitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDnsAitId

`func (o *DnsServerAddressInfo) SetBaseDnsAitId(v BaselineDnsAitId)`

SetBaseDnsAitId sets BaseDnsAitId field to given value.

### HasBaseDnsAitId

`func (o *DnsServerAddressInfo) HasBaseDnsAitId() bool`

HasBaseDnsAitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


