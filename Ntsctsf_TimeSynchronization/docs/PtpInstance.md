# PtpInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceType** | [**InstanceType**](InstanceType.md) |  | 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**PtpProfile** | **string** |  | 
**PortConfigs** | Pointer to [**[]ConfigForPort**](ConfigForPort.md) |  | [optional] 

## Methods

### NewPtpInstance

`func NewPtpInstance(instanceType InstanceType, protocol Protocol, ptpProfile string, ) *PtpInstance`

NewPtpInstance instantiates a new PtpInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPtpInstanceWithDefaults

`func NewPtpInstanceWithDefaults() *PtpInstance`

NewPtpInstanceWithDefaults instantiates a new PtpInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceType

`func (o *PtpInstance) GetInstanceType() InstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *PtpInstance) GetInstanceTypeOk() (*InstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *PtpInstance) SetInstanceType(v InstanceType)`

SetInstanceType sets InstanceType field to given value.


### GetProtocol

`func (o *PtpInstance) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PtpInstance) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PtpInstance) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetPtpProfile

`func (o *PtpInstance) GetPtpProfile() string`

GetPtpProfile returns the PtpProfile field if non-nil, zero value otherwise.

### GetPtpProfileOk

`func (o *PtpInstance) GetPtpProfileOk() (*string, bool)`

GetPtpProfileOk returns a tuple with the PtpProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpProfile

`func (o *PtpInstance) SetPtpProfile(v string)`

SetPtpProfile sets PtpProfile field to given value.


### GetPortConfigs

`func (o *PtpInstance) GetPortConfigs() []ConfigForPort`

GetPortConfigs returns the PortConfigs field if non-nil, zero value otherwise.

### GetPortConfigsOk

`func (o *PtpInstance) GetPortConfigsOk() (*[]ConfigForPort, bool)`

GetPortConfigsOk returns a tuple with the PortConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfigs

`func (o *PtpInstance) SetPortConfigs(v []ConfigForPort)`

SetPortConfigs sets PortConfigs field to given value.

### HasPortConfigs

`func (o *PtpInstance) HasPortConfigs() bool`

HasPortConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


