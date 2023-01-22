# PcfAddressingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PcfFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_PolicyAuthorization service. | [optional] 
**BindingInfo** | Pointer to **string** | contains the binding indications of the PCF. | [optional] 

## Methods

### NewPcfAddressingInfo

`func NewPcfAddressingInfo() *PcfAddressingInfo`

NewPcfAddressingInfo instantiates a new PcfAddressingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfAddressingInfoWithDefaults

`func NewPcfAddressingInfoWithDefaults() *PcfAddressingInfo`

NewPcfAddressingInfoWithDefaults instantiates a new PcfAddressingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcfFqdn

`func (o *PcfAddressingInfo) GetPcfFqdn() string`

GetPcfFqdn returns the PcfFqdn field if non-nil, zero value otherwise.

### GetPcfFqdnOk

`func (o *PcfAddressingInfo) GetPcfFqdnOk() (*string, bool)`

GetPcfFqdnOk returns a tuple with the PcfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFqdn

`func (o *PcfAddressingInfo) SetPcfFqdn(v string)`

SetPcfFqdn sets PcfFqdn field to given value.

### HasPcfFqdn

`func (o *PcfAddressingInfo) HasPcfFqdn() bool`

HasPcfFqdn returns a boolean if a field has been set.

### GetPcfIpEndPoints

`func (o *PcfAddressingInfo) GetPcfIpEndPoints() []IpEndPoint`

GetPcfIpEndPoints returns the PcfIpEndPoints field if non-nil, zero value otherwise.

### GetPcfIpEndPointsOk

`func (o *PcfAddressingInfo) GetPcfIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfIpEndPoints

`func (o *PcfAddressingInfo) SetPcfIpEndPoints(v []IpEndPoint)`

SetPcfIpEndPoints sets PcfIpEndPoints field to given value.

### HasPcfIpEndPoints

`func (o *PcfAddressingInfo) HasPcfIpEndPoints() bool`

HasPcfIpEndPoints returns a boolean if a field has been set.

### GetBindingInfo

`func (o *PcfAddressingInfo) GetBindingInfo() string`

GetBindingInfo returns the BindingInfo field if non-nil, zero value otherwise.

### GetBindingInfoOk

`func (o *PcfAddressingInfo) GetBindingInfoOk() (*string, bool)`

GetBindingInfoOk returns a tuple with the BindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingInfo

`func (o *PcfAddressingInfo) SetBindingInfo(v string)`

SetBindingInfo sets BindingInfo field to given value.

### HasBindingInfo

`func (o *PcfAddressingInfo) HasBindingInfo() bool`

HasBindingInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


