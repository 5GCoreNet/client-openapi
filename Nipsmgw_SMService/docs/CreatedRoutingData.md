# CreatedRoutingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpsmgwIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**IpsmgwIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**IpsmgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewCreatedRoutingData

`func NewCreatedRoutingData() *CreatedRoutingData`

NewCreatedRoutingData instantiates a new CreatedRoutingData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedRoutingDataWithDefaults

`func NewCreatedRoutingDataWithDefaults() *CreatedRoutingData`

NewCreatedRoutingDataWithDefaults instantiates a new CreatedRoutingData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpsmgwIpv4

`func (o *CreatedRoutingData) GetIpsmgwIpv4() string`

GetIpsmgwIpv4 returns the IpsmgwIpv4 field if non-nil, zero value otherwise.

### GetIpsmgwIpv4Ok

`func (o *CreatedRoutingData) GetIpsmgwIpv4Ok() (*string, bool)`

GetIpsmgwIpv4Ok returns a tuple with the IpsmgwIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsmgwIpv4

`func (o *CreatedRoutingData) SetIpsmgwIpv4(v string)`

SetIpsmgwIpv4 sets IpsmgwIpv4 field to given value.

### HasIpsmgwIpv4

`func (o *CreatedRoutingData) HasIpsmgwIpv4() bool`

HasIpsmgwIpv4 returns a boolean if a field has been set.

### GetIpsmgwIpv6

`func (o *CreatedRoutingData) GetIpsmgwIpv6() Ipv6Addr`

GetIpsmgwIpv6 returns the IpsmgwIpv6 field if non-nil, zero value otherwise.

### GetIpsmgwIpv6Ok

`func (o *CreatedRoutingData) GetIpsmgwIpv6Ok() (*Ipv6Addr, bool)`

GetIpsmgwIpv6Ok returns a tuple with the IpsmgwIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsmgwIpv6

`func (o *CreatedRoutingData) SetIpsmgwIpv6(v Ipv6Addr)`

SetIpsmgwIpv6 sets IpsmgwIpv6 field to given value.

### HasIpsmgwIpv6

`func (o *CreatedRoutingData) HasIpsmgwIpv6() bool`

HasIpsmgwIpv6 returns a boolean if a field has been set.

### GetIpsmgwFqdn

`func (o *CreatedRoutingData) GetIpsmgwFqdn() string`

GetIpsmgwFqdn returns the IpsmgwFqdn field if non-nil, zero value otherwise.

### GetIpsmgwFqdnOk

`func (o *CreatedRoutingData) GetIpsmgwFqdnOk() (*string, bool)`

GetIpsmgwFqdnOk returns a tuple with the IpsmgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsmgwFqdn

`func (o *CreatedRoutingData) SetIpsmgwFqdn(v string)`

SetIpsmgwFqdn sets IpsmgwFqdn field to given value.

### HasIpsmgwFqdn

`func (o *CreatedRoutingData) HasIpsmgwFqdn() bool`

HasIpsmgwFqdn returns a boolean if a field has been set.

### GetCorrelationId

`func (o *CreatedRoutingData) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *CreatedRoutingData) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *CreatedRoutingData) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *CreatedRoutingData) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *CreatedRoutingData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *CreatedRoutingData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *CreatedRoutingData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *CreatedRoutingData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


