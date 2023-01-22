# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrType** | [**AddressType**](AddressType.md) |  | 
**Addr** | **string** |  | 

## Methods

### NewAddress

`func NewAddress(addrType AddressType, addr string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrType

`func (o *Address) GetAddrType() AddressType`

GetAddrType returns the AddrType field if non-nil, zero value otherwise.

### GetAddrTypeOk

`func (o *Address) GetAddrTypeOk() (*AddressType, bool)`

GetAddrTypeOk returns a tuple with the AddrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrType

`func (o *Address) SetAddrType(v AddressType)`

SetAddrType sets AddrType field to given value.


### GetAddr

`func (o *Address) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *Address) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *Address) SetAddr(v string)`

SetAddr sets Addr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


