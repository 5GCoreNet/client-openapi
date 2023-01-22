# SAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to [**HostAddr**](HostAddr.md) |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 

## Methods

### NewSAP

`func NewSAP() *SAP`

NewSAP instantiates a new SAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAPWithDefaults

`func NewSAPWithDefaults() *SAP`

NewSAPWithDefaults instantiates a new SAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SAP) GetHost() HostAddr`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SAP) GetHostOk() (*HostAddr, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SAP) SetHost(v HostAddr)`

SetHost sets Host field to given value.

### HasHost

`func (o *SAP) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *SAP) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SAP) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SAP) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SAP) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


