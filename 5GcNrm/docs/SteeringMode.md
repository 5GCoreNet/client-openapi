# SteeringMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SteerModeValue** | Pointer to [**SteerModeValue**](SteerModeValue.md) |  | [optional] 
**Active** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**Standby** | Pointer to [**AccessTypeRm**](AccessTypeRm.md) |  | [optional] 
**ThreeGLoad** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**PrioAcc** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewSteeringMode

`func NewSteeringMode() *SteeringMode`

NewSteeringMode instantiates a new SteeringMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSteeringModeWithDefaults

`func NewSteeringModeWithDefaults() *SteeringMode`

NewSteeringModeWithDefaults instantiates a new SteeringMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSteerModeValue

`func (o *SteeringMode) GetSteerModeValue() SteerModeValue`

GetSteerModeValue returns the SteerModeValue field if non-nil, zero value otherwise.

### GetSteerModeValueOk

`func (o *SteeringMode) GetSteerModeValueOk() (*SteerModeValue, bool)`

GetSteerModeValueOk returns a tuple with the SteerModeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteerModeValue

`func (o *SteeringMode) SetSteerModeValue(v SteerModeValue)`

SetSteerModeValue sets SteerModeValue field to given value.

### HasSteerModeValue

`func (o *SteeringMode) HasSteerModeValue() bool`

HasSteerModeValue returns a boolean if a field has been set.

### GetActive

`func (o *SteeringMode) GetActive() AccessType`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SteeringMode) GetActiveOk() (*AccessType, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SteeringMode) SetActive(v AccessType)`

SetActive sets Active field to given value.

### HasActive

`func (o *SteeringMode) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStandby

`func (o *SteeringMode) GetStandby() AccessTypeRm`

GetStandby returns the Standby field if non-nil, zero value otherwise.

### GetStandbyOk

`func (o *SteeringMode) GetStandbyOk() (*AccessTypeRm, bool)`

GetStandbyOk returns a tuple with the Standby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandby

`func (o *SteeringMode) SetStandby(v AccessTypeRm)`

SetStandby sets Standby field to given value.

### HasStandby

`func (o *SteeringMode) HasStandby() bool`

HasStandby returns a boolean if a field has been set.

### GetThreeGLoad

`func (o *SteeringMode) GetThreeGLoad() int32`

GetThreeGLoad returns the ThreeGLoad field if non-nil, zero value otherwise.

### GetThreeGLoadOk

`func (o *SteeringMode) GetThreeGLoadOk() (*int32, bool)`

GetThreeGLoadOk returns a tuple with the ThreeGLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeGLoad

`func (o *SteeringMode) SetThreeGLoad(v int32)`

SetThreeGLoad sets ThreeGLoad field to given value.

### HasThreeGLoad

`func (o *SteeringMode) HasThreeGLoad() bool`

HasThreeGLoad returns a boolean if a field has been set.

### GetPrioAcc

`func (o *SteeringMode) GetPrioAcc() AccessType`

GetPrioAcc returns the PrioAcc field if non-nil, zero value otherwise.

### GetPrioAccOk

`func (o *SteeringMode) GetPrioAccOk() (*AccessType, bool)`

GetPrioAccOk returns a tuple with the PrioAcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioAcc

`func (o *SteeringMode) SetPrioAcc(v AccessType)`

SetPrioAcc sets PrioAcc field to given value.

### HasPrioAcc

`func (o *SteeringMode) HasPrioAcc() bool`

HasPrioAcc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


