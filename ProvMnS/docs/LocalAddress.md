# LocalAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressWithVlan** | Pointer to [**AddressWithVlan**](AddressWithVlan.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 

## Methods

### NewLocalAddress

`func NewLocalAddress() *LocalAddress`

NewLocalAddress instantiates a new LocalAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalAddressWithDefaults

`func NewLocalAddressWithDefaults() *LocalAddress`

NewLocalAddressWithDefaults instantiates a new LocalAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressWithVlan

`func (o *LocalAddress) GetAddressWithVlan() AddressWithVlan`

GetAddressWithVlan returns the AddressWithVlan field if non-nil, zero value otherwise.

### GetAddressWithVlanOk

`func (o *LocalAddress) GetAddressWithVlanOk() (*AddressWithVlan, bool)`

GetAddressWithVlanOk returns a tuple with the AddressWithVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressWithVlan

`func (o *LocalAddress) SetAddressWithVlan(v AddressWithVlan)`

SetAddressWithVlan sets AddressWithVlan field to given value.

### HasAddressWithVlan

`func (o *LocalAddress) HasAddressWithVlan() bool`

HasAddressWithVlan returns a boolean if a field has been set.

### GetPort

`func (o *LocalAddress) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LocalAddress) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LocalAddress) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LocalAddress) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


