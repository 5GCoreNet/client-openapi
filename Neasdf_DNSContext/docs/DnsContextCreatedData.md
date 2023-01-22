# DnsContextCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasdfIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**EasdfIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewDnsContextCreatedData

`func NewDnsContextCreatedData() *DnsContextCreatedData`

NewDnsContextCreatedData instantiates a new DnsContextCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsContextCreatedDataWithDefaults

`func NewDnsContextCreatedDataWithDefaults() *DnsContextCreatedData`

NewDnsContextCreatedDataWithDefaults instantiates a new DnsContextCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasdfIpv4Addr

`func (o *DnsContextCreatedData) GetEasdfIpv4Addr() string`

GetEasdfIpv4Addr returns the EasdfIpv4Addr field if non-nil, zero value otherwise.

### GetEasdfIpv4AddrOk

`func (o *DnsContextCreatedData) GetEasdfIpv4AddrOk() (*string, bool)`

GetEasdfIpv4AddrOk returns a tuple with the EasdfIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasdfIpv4Addr

`func (o *DnsContextCreatedData) SetEasdfIpv4Addr(v string)`

SetEasdfIpv4Addr sets EasdfIpv4Addr field to given value.

### HasEasdfIpv4Addr

`func (o *DnsContextCreatedData) HasEasdfIpv4Addr() bool`

HasEasdfIpv4Addr returns a boolean if a field has been set.

### GetEasdfIpv6Addr

`func (o *DnsContextCreatedData) GetEasdfIpv6Addr() Ipv6Addr`

GetEasdfIpv6Addr returns the EasdfIpv6Addr field if non-nil, zero value otherwise.

### GetEasdfIpv6AddrOk

`func (o *DnsContextCreatedData) GetEasdfIpv6AddrOk() (*Ipv6Addr, bool)`

GetEasdfIpv6AddrOk returns a tuple with the EasdfIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasdfIpv6Addr

`func (o *DnsContextCreatedData) SetEasdfIpv6Addr(v Ipv6Addr)`

SetEasdfIpv6Addr sets EasdfIpv6Addr field to given value.

### HasEasdfIpv6Addr

`func (o *DnsContextCreatedData) HasEasdfIpv6Addr() bool`

HasEasdfIpv6Addr returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *DnsContextCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DnsContextCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DnsContextCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *DnsContextCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


