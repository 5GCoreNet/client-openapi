# NrfInfoServedBsfInfoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnnList** | Pointer to **[]string** |  | [optional] 
**IpDomainList** | Pointer to **[]string** |  | [optional] 
**Ipv4AddressRanges** | Pointer to [**[]Ipv4AddressRange**](Ipv4AddressRange.md) |  | [optional] 
**Ipv6PrefixRanges** | Pointer to [**[]Ipv6PrefixRange**](Ipv6PrefixRange.md) |  | [optional] 
**RxDiamHost** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**RxDiamRealm** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 

## Methods

### NewNrfInfoServedBsfInfoValue

`func NewNrfInfoServedBsfInfoValue() *NrfInfoServedBsfInfoValue`

NewNrfInfoServedBsfInfoValue instantiates a new NrfInfoServedBsfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedBsfInfoValueWithDefaults

`func NewNrfInfoServedBsfInfoValueWithDefaults() *NrfInfoServedBsfInfoValue`

NewNrfInfoServedBsfInfoValueWithDefaults instantiates a new NrfInfoServedBsfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnnList

`func (o *NrfInfoServedBsfInfoValue) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *NrfInfoServedBsfInfoValue) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *NrfInfoServedBsfInfoValue) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *NrfInfoServedBsfInfoValue) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetIpDomainList

`func (o *NrfInfoServedBsfInfoValue) GetIpDomainList() []string`

GetIpDomainList returns the IpDomainList field if non-nil, zero value otherwise.

### GetIpDomainListOk

`func (o *NrfInfoServedBsfInfoValue) GetIpDomainListOk() (*[]string, bool)`

GetIpDomainListOk returns a tuple with the IpDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomainList

`func (o *NrfInfoServedBsfInfoValue) SetIpDomainList(v []string)`

SetIpDomainList sets IpDomainList field to given value.

### HasIpDomainList

`func (o *NrfInfoServedBsfInfoValue) HasIpDomainList() bool`

HasIpDomainList returns a boolean if a field has been set.

### GetIpv4AddressRanges

`func (o *NrfInfoServedBsfInfoValue) GetIpv4AddressRanges() []Ipv4AddressRange`

GetIpv4AddressRanges returns the Ipv4AddressRanges field if non-nil, zero value otherwise.

### GetIpv4AddressRangesOk

`func (o *NrfInfoServedBsfInfoValue) GetIpv4AddressRangesOk() (*[]Ipv4AddressRange, bool)`

GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressRanges

`func (o *NrfInfoServedBsfInfoValue) SetIpv4AddressRanges(v []Ipv4AddressRange)`

SetIpv4AddressRanges sets Ipv4AddressRanges field to given value.

### HasIpv4AddressRanges

`func (o *NrfInfoServedBsfInfoValue) HasIpv4AddressRanges() bool`

HasIpv4AddressRanges returns a boolean if a field has been set.

### GetIpv6PrefixRanges

`func (o *NrfInfoServedBsfInfoValue) GetIpv6PrefixRanges() []Ipv6PrefixRange`

GetIpv6PrefixRanges returns the Ipv6PrefixRanges field if non-nil, zero value otherwise.

### GetIpv6PrefixRangesOk

`func (o *NrfInfoServedBsfInfoValue) GetIpv6PrefixRangesOk() (*[]Ipv6PrefixRange, bool)`

GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixRanges

`func (o *NrfInfoServedBsfInfoValue) SetIpv6PrefixRanges(v []Ipv6PrefixRange)`

SetIpv6PrefixRanges sets Ipv6PrefixRanges field to given value.

### HasIpv6PrefixRanges

`func (o *NrfInfoServedBsfInfoValue) HasIpv6PrefixRanges() bool`

HasIpv6PrefixRanges returns a boolean if a field has been set.

### GetRxDiamHost

`func (o *NrfInfoServedBsfInfoValue) GetRxDiamHost() string`

GetRxDiamHost returns the RxDiamHost field if non-nil, zero value otherwise.

### GetRxDiamHostOk

`func (o *NrfInfoServedBsfInfoValue) GetRxDiamHostOk() (*string, bool)`

GetRxDiamHostOk returns a tuple with the RxDiamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamHost

`func (o *NrfInfoServedBsfInfoValue) SetRxDiamHost(v string)`

SetRxDiamHost sets RxDiamHost field to given value.

### HasRxDiamHost

`func (o *NrfInfoServedBsfInfoValue) HasRxDiamHost() bool`

HasRxDiamHost returns a boolean if a field has been set.

### GetRxDiamRealm

`func (o *NrfInfoServedBsfInfoValue) GetRxDiamRealm() string`

GetRxDiamRealm returns the RxDiamRealm field if non-nil, zero value otherwise.

### GetRxDiamRealmOk

`func (o *NrfInfoServedBsfInfoValue) GetRxDiamRealmOk() (*string, bool)`

GetRxDiamRealmOk returns a tuple with the RxDiamRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxDiamRealm

`func (o *NrfInfoServedBsfInfoValue) SetRxDiamRealm(v string)`

SetRxDiamRealm sets RxDiamRealm field to given value.

### HasRxDiamRealm

`func (o *NrfInfoServedBsfInfoValue) HasRxDiamRealm() bool`

HasRxDiamRealm returns a boolean if a field has been set.

### GetGroupId

`func (o *NrfInfoServedBsfInfoValue) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NrfInfoServedBsfInfoValue) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NrfInfoServedBsfInfoValue) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NrfInfoServedBsfInfoValue) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *NrfInfoServedBsfInfoValue) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *NrfInfoServedBsfInfoValue) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *NrfInfoServedBsfInfoValue) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *NrfInfoServedBsfInfoValue) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NrfInfoServedBsfInfoValue) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NrfInfoServedBsfInfoValue) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NrfInfoServedBsfInfoValue) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NrfInfoServedBsfInfoValue) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


