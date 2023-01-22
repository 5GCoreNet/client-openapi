# EasServerAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | [**IpAddr**](IpAddr.md) |  | 
**Port** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 

## Methods

### NewEasServerAddress

`func NewEasServerAddress(ip IpAddr, port int32, ) *EasServerAddress`

NewEasServerAddress instantiates a new EasServerAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasServerAddressWithDefaults

`func NewEasServerAddressWithDefaults() *EasServerAddress`

NewEasServerAddressWithDefaults instantiates a new EasServerAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *EasServerAddress) GetIp() IpAddr`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *EasServerAddress) GetIpOk() (*IpAddr, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *EasServerAddress) SetIp(v IpAddr)`

SetIp sets Ip field to given value.


### GetPort

`func (o *EasServerAddress) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EasServerAddress) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EasServerAddress) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


