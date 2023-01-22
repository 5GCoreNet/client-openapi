# PcfForUeBindingPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PcfForUeFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfForUeIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) | IP end points of the PCF hosting the Npcf_AmPolicyAuthorization service. | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewPcfForUeBindingPatch

`func NewPcfForUeBindingPatch() *PcfForUeBindingPatch`

NewPcfForUeBindingPatch instantiates a new PcfForUeBindingPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfForUeBindingPatchWithDefaults

`func NewPcfForUeBindingPatchWithDefaults() *PcfForUeBindingPatch`

NewPcfForUeBindingPatchWithDefaults instantiates a new PcfForUeBindingPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPcfForUeFqdn

`func (o *PcfForUeBindingPatch) GetPcfForUeFqdn() string`

GetPcfForUeFqdn returns the PcfForUeFqdn field if non-nil, zero value otherwise.

### GetPcfForUeFqdnOk

`func (o *PcfForUeBindingPatch) GetPcfForUeFqdnOk() (*string, bool)`

GetPcfForUeFqdnOk returns a tuple with the PcfForUeFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfForUeFqdn

`func (o *PcfForUeBindingPatch) SetPcfForUeFqdn(v string)`

SetPcfForUeFqdn sets PcfForUeFqdn field to given value.

### HasPcfForUeFqdn

`func (o *PcfForUeBindingPatch) HasPcfForUeFqdn() bool`

HasPcfForUeFqdn returns a boolean if a field has been set.

### GetPcfForUeIpEndPoints

`func (o *PcfForUeBindingPatch) GetPcfForUeIpEndPoints() []IpEndPoint`

GetPcfForUeIpEndPoints returns the PcfForUeIpEndPoints field if non-nil, zero value otherwise.

### GetPcfForUeIpEndPointsOk

`func (o *PcfForUeBindingPatch) GetPcfForUeIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfForUeIpEndPointsOk returns a tuple with the PcfForUeIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfForUeIpEndPoints

`func (o *PcfForUeBindingPatch) SetPcfForUeIpEndPoints(v []IpEndPoint)`

SetPcfForUeIpEndPoints sets PcfForUeIpEndPoints field to given value.

### HasPcfForUeIpEndPoints

`func (o *PcfForUeBindingPatch) HasPcfForUeIpEndPoints() bool`

HasPcfForUeIpEndPoints returns a boolean if a field has been set.

### GetPcfId

`func (o *PcfForUeBindingPatch) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PcfForUeBindingPatch) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PcfForUeBindingPatch) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PcfForUeBindingPatch) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


