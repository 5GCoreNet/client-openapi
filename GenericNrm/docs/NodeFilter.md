# NodeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaOfInterest** | Pointer to [**AreaOfInterest**](AreaOfInterest.md) |  | [optional] 
**NetworkDomain** | Pointer to **string** |  | [optional] 
**CpUpType** | Pointer to **string** |  | [optional] 
**Sst** | Pointer to **int32** |  | [optional] 

## Methods

### NewNodeFilter

`func NewNodeFilter() *NodeFilter`

NewNodeFilter instantiates a new NodeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeFilterWithDefaults

`func NewNodeFilterWithDefaults() *NodeFilter`

NewNodeFilterWithDefaults instantiates a new NodeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaOfInterest

`func (o *NodeFilter) GetAreaOfInterest() AreaOfInterest`

GetAreaOfInterest returns the AreaOfInterest field if non-nil, zero value otherwise.

### GetAreaOfInterestOk

`func (o *NodeFilter) GetAreaOfInterestOk() (*AreaOfInterest, bool)`

GetAreaOfInterestOk returns a tuple with the AreaOfInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaOfInterest

`func (o *NodeFilter) SetAreaOfInterest(v AreaOfInterest)`

SetAreaOfInterest sets AreaOfInterest field to given value.

### HasAreaOfInterest

`func (o *NodeFilter) HasAreaOfInterest() bool`

HasAreaOfInterest returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *NodeFilter) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *NodeFilter) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *NodeFilter) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *NodeFilter) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetCpUpType

`func (o *NodeFilter) GetCpUpType() string`

GetCpUpType returns the CpUpType field if non-nil, zero value otherwise.

### GetCpUpTypeOk

`func (o *NodeFilter) GetCpUpTypeOk() (*string, bool)`

GetCpUpTypeOk returns a tuple with the CpUpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpUpType

`func (o *NodeFilter) SetCpUpType(v string)`

SetCpUpType sets CpUpType field to given value.

### HasCpUpType

`func (o *NodeFilter) HasCpUpType() bool`

HasCpUpType returns a boolean if a field has been set.

### GetSst

`func (o *NodeFilter) GetSst() int32`

GetSst returns the Sst field if non-nil, zero value otherwise.

### GetSstOk

`func (o *NodeFilter) GetSstOk() (*int32, bool)`

GetSstOk returns a tuple with the Sst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSst

`func (o *NodeFilter) SetSst(v int32)`

SetSst sets Sst field to given value.

### HasSst

`func (o *NodeFilter) HasSst() bool`

HasSst returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


