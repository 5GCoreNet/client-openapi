# ForwardingParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EcsOptionInfo** | Pointer to [**EcsOptionInfo**](EcsOptionInfo.md) |  | [optional] 
**DnsServerAddressInfo** | Pointer to [**DnsServerAddressInfo**](DnsServerAddressInfo.md) |  | [optional] 

## Methods

### NewForwardingParameters

`func NewForwardingParameters() *ForwardingParameters`

NewForwardingParameters instantiates a new ForwardingParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardingParametersWithDefaults

`func NewForwardingParametersWithDefaults() *ForwardingParameters`

NewForwardingParametersWithDefaults instantiates a new ForwardingParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEcsOptionInfo

`func (o *ForwardingParameters) GetEcsOptionInfo() EcsOptionInfo`

GetEcsOptionInfo returns the EcsOptionInfo field if non-nil, zero value otherwise.

### GetEcsOptionInfoOk

`func (o *ForwardingParameters) GetEcsOptionInfoOk() (*EcsOptionInfo, bool)`

GetEcsOptionInfoOk returns a tuple with the EcsOptionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcsOptionInfo

`func (o *ForwardingParameters) SetEcsOptionInfo(v EcsOptionInfo)`

SetEcsOptionInfo sets EcsOptionInfo field to given value.

### HasEcsOptionInfo

`func (o *ForwardingParameters) HasEcsOptionInfo() bool`

HasEcsOptionInfo returns a boolean if a field has been set.

### GetDnsServerAddressInfo

`func (o *ForwardingParameters) GetDnsServerAddressInfo() DnsServerAddressInfo`

GetDnsServerAddressInfo returns the DnsServerAddressInfo field if non-nil, zero value otherwise.

### GetDnsServerAddressInfoOk

`func (o *ForwardingParameters) GetDnsServerAddressInfoOk() (*DnsServerAddressInfo, bool)`

GetDnsServerAddressInfoOk returns a tuple with the DnsServerAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerAddressInfo

`func (o *ForwardingParameters) SetDnsServerAddressInfo(v DnsServerAddressInfo)`

SetDnsServerAddressInfo sets DnsServerAddressInfo field to given value.

### HasDnsServerAddressInfo

`func (o *ForwardingParameters) HasDnsServerAddressInfo() bool`

HasDnsServerAddressInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


