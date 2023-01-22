# DnsContextCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**UeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**SNssai** | [**Snssai**](Snssai.md) |  | 
**DnsRules** | [**map[string]DnsRule**](DnsRule.md) | map of DNS message handling rules where a valid JSON string serves as key | 
**NotifyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewDnsContextCreateData

`func NewDnsContextCreateData(dnn string, sNssai Snssai, dnsRules map[string]DnsRule, ) *DnsContextCreateData`

NewDnsContextCreateData instantiates a new DnsContextCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsContextCreateDataWithDefaults

`func NewDnsContextCreateDataWithDefaults() *DnsContextCreateData`

NewDnsContextCreateDataWithDefaults instantiates a new DnsContextCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeIpv4Addr

`func (o *DnsContextCreateData) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *DnsContextCreateData) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *DnsContextCreateData) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *DnsContextCreateData) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetUeIpv6Prefix

`func (o *DnsContextCreateData) GetUeIpv6Prefix() Ipv6Prefix`

GetUeIpv6Prefix returns the UeIpv6Prefix field if non-nil, zero value otherwise.

### GetUeIpv6PrefixOk

`func (o *DnsContextCreateData) GetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetUeIpv6PrefixOk returns a tuple with the UeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Prefix

`func (o *DnsContextCreateData) SetUeIpv6Prefix(v Ipv6Prefix)`

SetUeIpv6Prefix sets UeIpv6Prefix field to given value.

### HasUeIpv6Prefix

`func (o *DnsContextCreateData) HasUeIpv6Prefix() bool`

HasUeIpv6Prefix returns a boolean if a field has been set.

### GetDnn

`func (o *DnsContextCreateData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *DnsContextCreateData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *DnsContextCreateData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSNssai

`func (o *DnsContextCreateData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *DnsContextCreateData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *DnsContextCreateData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetDnsRules

`func (o *DnsContextCreateData) GetDnsRules() map[string]DnsRule`

GetDnsRules returns the DnsRules field if non-nil, zero value otherwise.

### GetDnsRulesOk

`func (o *DnsContextCreateData) GetDnsRulesOk() (*map[string]DnsRule, bool)`

GetDnsRulesOk returns a tuple with the DnsRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRules

`func (o *DnsContextCreateData) SetDnsRules(v map[string]DnsRule)`

SetDnsRules sets DnsRules field to given value.


### GetNotifyUri

`func (o *DnsContextCreateData) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *DnsContextCreateData) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *DnsContextCreateData) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.

### HasNotifyUri

`func (o *DnsContextCreateData) HasNotifyUri() bool`

HasNotifyUri returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *DnsContextCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DnsContextCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DnsContextCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *DnsContextCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


