# TargetContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextAttribute** | Pointer to **string** |  | [optional] 
**ContextCondition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**ContextValueRange** | Pointer to **float32** |  | [optional] 

## Methods

### NewTargetContext

`func NewTargetContext() *TargetContext`

NewTargetContext instantiates a new TargetContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetContextWithDefaults

`func NewTargetContextWithDefaults() *TargetContext`

NewTargetContextWithDefaults instantiates a new TargetContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextAttribute

`func (o *TargetContext) GetContextAttribute() string`

GetContextAttribute returns the ContextAttribute field if non-nil, zero value otherwise.

### GetContextAttributeOk

`func (o *TargetContext) GetContextAttributeOk() (*string, bool)`

GetContextAttributeOk returns a tuple with the ContextAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAttribute

`func (o *TargetContext) SetContextAttribute(v string)`

SetContextAttribute sets ContextAttribute field to given value.

### HasContextAttribute

`func (o *TargetContext) HasContextAttribute() bool`

HasContextAttribute returns a boolean if a field has been set.

### GetContextCondition

`func (o *TargetContext) GetContextCondition() Condition`

GetContextCondition returns the ContextCondition field if non-nil, zero value otherwise.

### GetContextConditionOk

`func (o *TargetContext) GetContextConditionOk() (*Condition, bool)`

GetContextConditionOk returns a tuple with the ContextCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextCondition

`func (o *TargetContext) SetContextCondition(v Condition)`

SetContextCondition sets ContextCondition field to given value.

### HasContextCondition

`func (o *TargetContext) HasContextCondition() bool`

HasContextCondition returns a boolean if a field has been set.

### GetContextValueRange

`func (o *TargetContext) GetContextValueRange() float32`

GetContextValueRange returns the ContextValueRange field if non-nil, zero value otherwise.

### GetContextValueRangeOk

`func (o *TargetContext) GetContextValueRangeOk() (*float32, bool)`

GetContextValueRangeOk returns a tuple with the ContextValueRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextValueRange

`func (o *TargetContext) SetContextValueRange(v float32)`

SetContextValueRange sets ContextValueRange field to given value.

### HasContextValueRange

`func (o *TargetContext) HasContextValueRange() bool`

HasContextValueRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


