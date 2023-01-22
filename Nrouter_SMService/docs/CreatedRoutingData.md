# CreatedRoutingData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouterIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**RouterIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**RouterFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
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

### GetRouterIpv4

`func (o *CreatedRoutingData) GetRouterIpv4() string`

GetRouterIpv4 returns the RouterIpv4 field if non-nil, zero value otherwise.

### GetRouterIpv4Ok

`func (o *CreatedRoutingData) GetRouterIpv4Ok() (*string, bool)`

GetRouterIpv4Ok returns a tuple with the RouterIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIpv4

`func (o *CreatedRoutingData) SetRouterIpv4(v string)`

SetRouterIpv4 sets RouterIpv4 field to given value.

### HasRouterIpv4

`func (o *CreatedRoutingData) HasRouterIpv4() bool`

HasRouterIpv4 returns a boolean if a field has been set.

### GetRouterIpv6

`func (o *CreatedRoutingData) GetRouterIpv6() Ipv6Addr`

GetRouterIpv6 returns the RouterIpv6 field if non-nil, zero value otherwise.

### GetRouterIpv6Ok

`func (o *CreatedRoutingData) GetRouterIpv6Ok() (*Ipv6Addr, bool)`

GetRouterIpv6Ok returns a tuple with the RouterIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterIpv6

`func (o *CreatedRoutingData) SetRouterIpv6(v Ipv6Addr)`

SetRouterIpv6 sets RouterIpv6 field to given value.

### HasRouterIpv6

`func (o *CreatedRoutingData) HasRouterIpv6() bool`

HasRouterIpv6 returns a boolean if a field has been set.

### GetRouterFqdn

`func (o *CreatedRoutingData) GetRouterFqdn() string`

GetRouterFqdn returns the RouterFqdn field if non-nil, zero value otherwise.

### GetRouterFqdnOk

`func (o *CreatedRoutingData) GetRouterFqdnOk() (*string, bool)`

GetRouterFqdnOk returns a tuple with the RouterFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterFqdn

`func (o *CreatedRoutingData) SetRouterFqdn(v string)`

SetRouterFqdn sets RouterFqdn field to given value.

### HasRouterFqdn

`func (o *CreatedRoutingData) HasRouterFqdn() bool`

HasRouterFqdn returns a boolean if a field has been set.

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


