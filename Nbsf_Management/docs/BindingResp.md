# BindingResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PcfSmFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfSmIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_SMPolicyControl service. | [optional] 

## Methods

### NewBindingResp

`func NewBindingResp() *BindingResp`

NewBindingResp instantiates a new BindingResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindingRespWithDefaults

`func NewBindingRespWithDefaults() *BindingResp`

NewBindingRespWithDefaults instantiates a new BindingResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcfSmFqdn

`func (o *BindingResp) GetPcfSmFqdn() string`

GetPcfSmFqdn returns the PcfSmFqdn field if non-nil, zero value otherwise.

### GetPcfSmFqdnOk

`func (o *BindingResp) GetPcfSmFqdnOk() (*string, bool)`

GetPcfSmFqdnOk returns a tuple with the PcfSmFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSmFqdn

`func (o *BindingResp) SetPcfSmFqdn(v string)`

SetPcfSmFqdn sets PcfSmFqdn field to given value.

### HasPcfSmFqdn

`func (o *BindingResp) HasPcfSmFqdn() bool`

HasPcfSmFqdn returns a boolean if a field has been set.

### GetPcfSmIpEndPoints

`func (o *BindingResp) GetPcfSmIpEndPoints() []IpEndPoint`

GetPcfSmIpEndPoints returns the PcfSmIpEndPoints field if non-nil, zero value otherwise.

### GetPcfSmIpEndPointsOk

`func (o *BindingResp) GetPcfSmIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfSmIpEndPointsOk returns a tuple with the PcfSmIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSmIpEndPoints

`func (o *BindingResp) SetPcfSmIpEndPoints(v []IpEndPoint)`

SetPcfSmIpEndPoints sets PcfSmIpEndPoints field to given value.

### HasPcfSmIpEndPoints

`func (o *BindingResp) HasPcfSmIpEndPoints() bool`

HasPcfSmIpEndPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


