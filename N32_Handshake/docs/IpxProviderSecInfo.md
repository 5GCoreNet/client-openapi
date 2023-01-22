# IpxProviderSecInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpxProviderId** | **string** | Fully Qualified Domain Name | 
**RawPublicKeyList** | Pointer to **[]string** |  | [optional] 
**CertificateList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIpxProviderSecInfo

`func NewIpxProviderSecInfo(ipxProviderId string, ) *IpxProviderSecInfo`

NewIpxProviderSecInfo instantiates a new IpxProviderSecInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpxProviderSecInfoWithDefaults

`func NewIpxProviderSecInfoWithDefaults() *IpxProviderSecInfo`

NewIpxProviderSecInfoWithDefaults instantiates a new IpxProviderSecInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpxProviderId

`func (o *IpxProviderSecInfo) GetIpxProviderId() string`

GetIpxProviderId returns the IpxProviderId field if non-nil, zero value otherwise.

### GetIpxProviderIdOk

`func (o *IpxProviderSecInfo) GetIpxProviderIdOk() (*string, bool)`

GetIpxProviderIdOk returns a tuple with the IpxProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxProviderId

`func (o *IpxProviderSecInfo) SetIpxProviderId(v string)`

SetIpxProviderId sets IpxProviderId field to given value.


### GetRawPublicKeyList

`func (o *IpxProviderSecInfo) GetRawPublicKeyList() []string`

GetRawPublicKeyList returns the RawPublicKeyList field if non-nil, zero value otherwise.

### GetRawPublicKeyListOk

`func (o *IpxProviderSecInfo) GetRawPublicKeyListOk() (*[]string, bool)`

GetRawPublicKeyListOk returns a tuple with the RawPublicKeyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawPublicKeyList

`func (o *IpxProviderSecInfo) SetRawPublicKeyList(v []string)`

SetRawPublicKeyList sets RawPublicKeyList field to given value.

### HasRawPublicKeyList

`func (o *IpxProviderSecInfo) HasRawPublicKeyList() bool`

HasRawPublicKeyList returns a boolean if a field has been set.

### GetCertificateList

`func (o *IpxProviderSecInfo) GetCertificateList() []string`

GetCertificateList returns the CertificateList field if non-nil, zero value otherwise.

### GetCertificateListOk

`func (o *IpxProviderSecInfo) GetCertificateListOk() (*[]string, bool)`

GetCertificateListOk returns a tuple with the CertificateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateList

`func (o *IpxProviderSecInfo) SetCertificateList(v []string)`

SetCertificateList sets CertificateList field to given value.

### HasCertificateList

`func (o *IpxProviderSecInfo) HasCertificateList() bool`

HasCertificateList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


