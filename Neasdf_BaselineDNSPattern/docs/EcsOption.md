# EcsOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePrefixLength** | **int32** |  | 
**ScopePrefixLength** | Pointer to **int32** |  | [optional] 
**IpAddr** | [**IpAddr**](IpAddr.md) |  | 

## Methods

### NewEcsOption

`func NewEcsOption(sourcePrefixLength int32, ipAddr IpAddr, ) *EcsOption`

NewEcsOption instantiates a new EcsOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcsOptionWithDefaults

`func NewEcsOptionWithDefaults() *EcsOption`

NewEcsOptionWithDefaults instantiates a new EcsOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePrefixLength

`func (o *EcsOption) GetSourcePrefixLength() int32`

GetSourcePrefixLength returns the SourcePrefixLength field if non-nil, zero value otherwise.

### GetSourcePrefixLengthOk

`func (o *EcsOption) GetSourcePrefixLengthOk() (*int32, bool)`

GetSourcePrefixLengthOk returns a tuple with the SourcePrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePrefixLength

`func (o *EcsOption) SetSourcePrefixLength(v int32)`

SetSourcePrefixLength sets SourcePrefixLength field to given value.


### GetScopePrefixLength

`func (o *EcsOption) GetScopePrefixLength() int32`

GetScopePrefixLength returns the ScopePrefixLength field if non-nil, zero value otherwise.

### GetScopePrefixLengthOk

`func (o *EcsOption) GetScopePrefixLengthOk() (*int32, bool)`

GetScopePrefixLengthOk returns a tuple with the ScopePrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopePrefixLength

`func (o *EcsOption) SetScopePrefixLength(v int32)`

SetScopePrefixLength sets ScopePrefixLength field to given value.

### HasScopePrefixLength

`func (o *EcsOption) HasScopePrefixLength() bool`

HasScopePrefixLength returns a boolean if a field has been set.

### GetIpAddr

`func (o *EcsOption) GetIpAddr() IpAddr`

GetIpAddr returns the IpAddr field if non-nil, zero value otherwise.

### GetIpAddrOk

`func (o *EcsOption) GetIpAddrOk() (*IpAddr, bool)`

GetIpAddrOk returns a tuple with the IpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddr

`func (o *EcsOption) SetIpAddr(v IpAddr)`

SetIpAddr sets IpAddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


