# ExtSsm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ssm** | [**Ssm**](Ssm.md) |  | 
**PortNumber** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 

## Methods

### NewExtSsm

`func NewExtSsm(ssm Ssm, portNumber int32, ) *ExtSsm`

NewExtSsm instantiates a new ExtSsm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtSsmWithDefaults

`func NewExtSsmWithDefaults() *ExtSsm`

NewExtSsmWithDefaults instantiates a new ExtSsm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsm

`func (o *ExtSsm) GetSsm() Ssm`

GetSsm returns the Ssm field if non-nil, zero value otherwise.

### GetSsmOk

`func (o *ExtSsm) GetSsmOk() (*Ssm, bool)`

GetSsmOk returns a tuple with the Ssm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsm

`func (o *ExtSsm) SetSsm(v Ssm)`

SetSsm sets Ssm field to given value.


### GetPortNumber

`func (o *ExtSsm) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *ExtSsm) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *ExtSsm) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


