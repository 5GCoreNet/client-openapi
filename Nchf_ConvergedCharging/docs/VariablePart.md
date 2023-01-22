# VariablePart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariablePartType** | [**VariablePartType**](VariablePartType.md) |  | 
**VariablePartValue** | **[]string** |  | 
**VariablePartOrder** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 

## Methods

### NewVariablePart

`func NewVariablePart(variablePartType VariablePartType, variablePartValue []string, ) *VariablePart`

NewVariablePart instantiates a new VariablePart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariablePartWithDefaults

`func NewVariablePartWithDefaults() *VariablePart`

NewVariablePartWithDefaults instantiates a new VariablePart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariablePartType

`func (o *VariablePart) GetVariablePartType() VariablePartType`

GetVariablePartType returns the VariablePartType field if non-nil, zero value otherwise.

### GetVariablePartTypeOk

`func (o *VariablePart) GetVariablePartTypeOk() (*VariablePartType, bool)`

GetVariablePartTypeOk returns a tuple with the VariablePartType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePartType

`func (o *VariablePart) SetVariablePartType(v VariablePartType)`

SetVariablePartType sets VariablePartType field to given value.


### GetVariablePartValue

`func (o *VariablePart) GetVariablePartValue() []string`

GetVariablePartValue returns the VariablePartValue field if non-nil, zero value otherwise.

### GetVariablePartValueOk

`func (o *VariablePart) GetVariablePartValueOk() (*[]string, bool)`

GetVariablePartValueOk returns a tuple with the VariablePartValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePartValue

`func (o *VariablePart) SetVariablePartValue(v []string)`

SetVariablePartValue sets VariablePartValue field to given value.


### GetVariablePartOrder

`func (o *VariablePart) GetVariablePartOrder() int32`

GetVariablePartOrder returns the VariablePartOrder field if non-nil, zero value otherwise.

### GetVariablePartOrderOk

`func (o *VariablePart) GetVariablePartOrderOk() (*int32, bool)`

GetVariablePartOrderOk returns a tuple with the VariablePartOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablePartOrder

`func (o *VariablePart) SetVariablePartOrder(v int32)`

SetVariablePartOrder sets VariablePartOrder field to given value.

### HasVariablePartOrder

`func (o *VariablePart) HasVariablePartOrder() bool`

HasVariablePartOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


