# RdsPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortUE** | **int32** | Unsigned integer with valid values between 0 and 65535. | 
**PortSCEF** | **int32** | Unsigned integer with valid values between 0 and 65535. | 

## Methods

### NewRdsPort

`func NewRdsPort(portUE int32, portSCEF int32, ) *RdsPort`

NewRdsPort instantiates a new RdsPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsPortWithDefaults

`func NewRdsPortWithDefaults() *RdsPort`

NewRdsPortWithDefaults instantiates a new RdsPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortUE

`func (o *RdsPort) GetPortUE() int32`

GetPortUE returns the PortUE field if non-nil, zero value otherwise.

### GetPortUEOk

`func (o *RdsPort) GetPortUEOk() (*int32, bool)`

GetPortUEOk returns a tuple with the PortUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUE

`func (o *RdsPort) SetPortUE(v int32)`

SetPortUE sets PortUE field to given value.


### GetPortSCEF

`func (o *RdsPort) GetPortSCEF() int32`

GetPortSCEF returns the PortSCEF field if non-nil, zero value otherwise.

### GetPortSCEFOk

`func (o *RdsPort) GetPortSCEFOk() (*int32, bool)`

GetPortSCEFOk returns a tuple with the PortSCEF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSCEF

`func (o *RdsPort) SetPortSCEF(v int32)`

SetPortSCEF sets PortSCEF field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


