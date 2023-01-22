# ScpDomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScpFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**ScpIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) |  | [optional] 
**ScpPrefix** | Pointer to **string** |  | [optional] 
**ScpPorts** | Pointer to **map[string]int32** | Port numbers for HTTP and HTTPS. The key of the map shall be \&quot;http\&quot; or \&quot;https\&quot;.  | [optional] 

## Methods

### NewScpDomainInfo

`func NewScpDomainInfo() *ScpDomainInfo`

NewScpDomainInfo instantiates a new ScpDomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScpDomainInfoWithDefaults

`func NewScpDomainInfoWithDefaults() *ScpDomainInfo`

NewScpDomainInfoWithDefaults instantiates a new ScpDomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScpFqdn

`func (o *ScpDomainInfo) GetScpFqdn() string`

GetScpFqdn returns the ScpFqdn field if non-nil, zero value otherwise.

### GetScpFqdnOk

`func (o *ScpDomainInfo) GetScpFqdnOk() (*string, bool)`

GetScpFqdnOk returns a tuple with the ScpFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpFqdn

`func (o *ScpDomainInfo) SetScpFqdn(v string)`

SetScpFqdn sets ScpFqdn field to given value.

### HasScpFqdn

`func (o *ScpDomainInfo) HasScpFqdn() bool`

HasScpFqdn returns a boolean if a field has been set.

### GetScpIpEndPoints

`func (o *ScpDomainInfo) GetScpIpEndPoints() []IpEndPoint`

GetScpIpEndPoints returns the ScpIpEndPoints field if non-nil, zero value otherwise.

### GetScpIpEndPointsOk

`func (o *ScpDomainInfo) GetScpIpEndPointsOk() (*[]IpEndPoint, bool)`

GetScpIpEndPointsOk returns a tuple with the ScpIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpIpEndPoints

`func (o *ScpDomainInfo) SetScpIpEndPoints(v []IpEndPoint)`

SetScpIpEndPoints sets ScpIpEndPoints field to given value.

### HasScpIpEndPoints

`func (o *ScpDomainInfo) HasScpIpEndPoints() bool`

HasScpIpEndPoints returns a boolean if a field has been set.

### GetScpPrefix

`func (o *ScpDomainInfo) GetScpPrefix() string`

GetScpPrefix returns the ScpPrefix field if non-nil, zero value otherwise.

### GetScpPrefixOk

`func (o *ScpDomainInfo) GetScpPrefixOk() (*string, bool)`

GetScpPrefixOk returns a tuple with the ScpPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpPrefix

`func (o *ScpDomainInfo) SetScpPrefix(v string)`

SetScpPrefix sets ScpPrefix field to given value.

### HasScpPrefix

`func (o *ScpDomainInfo) HasScpPrefix() bool`

HasScpPrefix returns a boolean if a field has been set.

### GetScpPorts

`func (o *ScpDomainInfo) GetScpPorts() map[string]int32`

GetScpPorts returns the ScpPorts field if non-nil, zero value otherwise.

### GetScpPortsOk

`func (o *ScpDomainInfo) GetScpPortsOk() (*map[string]int32, bool)`

GetScpPortsOk returns a tuple with the ScpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpPorts

`func (o *ScpDomainInfo) SetScpPorts(v map[string]int32)`

SetScpPorts sets ScpPorts field to given value.

### HasScpPorts

`func (o *ScpDomainInfo) HasScpPorts() bool`

HasScpPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


