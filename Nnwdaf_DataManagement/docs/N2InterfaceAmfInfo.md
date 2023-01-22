# N2InterfaceAmfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipv4EndpointAddress** | Pointer to **[]string** |  | [optional] 
**Ipv6EndpointAddress** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AmfName** | Pointer to **string** | Fully Qualified Domain Name | [optional] 

## Methods

### NewN2InterfaceAmfInfo

`func NewN2InterfaceAmfInfo() *N2InterfaceAmfInfo`

NewN2InterfaceAmfInfo instantiates a new N2InterfaceAmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN2InterfaceAmfInfoWithDefaults

`func NewN2InterfaceAmfInfoWithDefaults() *N2InterfaceAmfInfo`

NewN2InterfaceAmfInfoWithDefaults instantiates a new N2InterfaceAmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpv4EndpointAddress

`func (o *N2InterfaceAmfInfo) GetIpv4EndpointAddress() []string`

GetIpv4EndpointAddress returns the Ipv4EndpointAddress field if non-nil, zero value otherwise.

### GetIpv4EndpointAddressOk

`func (o *N2InterfaceAmfInfo) GetIpv4EndpointAddressOk() (*[]string, bool)`

GetIpv4EndpointAddressOk returns a tuple with the Ipv4EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4EndpointAddress

`func (o *N2InterfaceAmfInfo) SetIpv4EndpointAddress(v []string)`

SetIpv4EndpointAddress sets Ipv4EndpointAddress field to given value.

### HasIpv4EndpointAddress

`func (o *N2InterfaceAmfInfo) HasIpv4EndpointAddress() bool`

HasIpv4EndpointAddress returns a boolean if a field has been set.

### GetIpv6EndpointAddress

`func (o *N2InterfaceAmfInfo) GetIpv6EndpointAddress() []Ipv6Addr`

GetIpv6EndpointAddress returns the Ipv6EndpointAddress field if non-nil, zero value otherwise.

### GetIpv6EndpointAddressOk

`func (o *N2InterfaceAmfInfo) GetIpv6EndpointAddressOk() (*[]Ipv6Addr, bool)`

GetIpv6EndpointAddressOk returns a tuple with the Ipv6EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6EndpointAddress

`func (o *N2InterfaceAmfInfo) SetIpv6EndpointAddress(v []Ipv6Addr)`

SetIpv6EndpointAddress sets Ipv6EndpointAddress field to given value.

### HasIpv6EndpointAddress

`func (o *N2InterfaceAmfInfo) HasIpv6EndpointAddress() bool`

HasIpv6EndpointAddress returns a boolean if a field has been set.

### GetAmfName

`func (o *N2InterfaceAmfInfo) GetAmfName() string`

GetAmfName returns the AmfName field if non-nil, zero value otherwise.

### GetAmfNameOk

`func (o *N2InterfaceAmfInfo) GetAmfNameOk() (*string, bool)`

GetAmfNameOk returns a tuple with the AmfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfName

`func (o *N2InterfaceAmfInfo) SetAmfName(v string)`

SetAmfName sets AmfName field to given value.

### HasAmfName

`func (o *N2InterfaceAmfInfo) HasAmfName() bool`

HasAmfName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


