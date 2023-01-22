# ObjectContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextAttribute** | Pointer to **string** |  | [optional] 
**ContextCondition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**ContextValueRange** | Pointer to **[]float32** |  | [optional] 

## Methods

### NewObjectContext

`func NewObjectContext() *ObjectContext`

NewObjectContext instantiates a new ObjectContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectContextWithDefaults

`func NewObjectContextWithDefaults() *ObjectContext`

NewObjectContextWithDefaults instantiates a new ObjectContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextAttribute

`func (o *ObjectContext) GetContextAttribute() string`

GetContextAttribute returns the ContextAttribute field if non-nil, zero value otherwise.

### GetContextAttributeOk

`func (o *ObjectContext) GetContextAttributeOk() (*string, bool)`

GetContextAttributeOk returns a tuple with the ContextAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAttribute

`func (o *ObjectContext) SetContextAttribute(v string)`

SetContextAttribute sets ContextAttribute field to given value.

### HasContextAttribute

`func (o *ObjectContext) HasContextAttribute() bool`

HasContextAttribute returns a boolean if a field has been set.

### GetContextCondition

`func (o *ObjectContext) GetContextCondition() Condition`

GetContextCondition returns the ContextCondition field if non-nil, zero value otherwise.

### GetContextConditionOk

`func (o *ObjectContext) GetContextConditionOk() (*Condition, bool)`

GetContextConditionOk returns a tuple with the ContextCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextCondition

`func (o *ObjectContext) SetContextCondition(v Condition)`

SetContextCondition sets ContextCondition field to given value.

### HasContextCondition

`func (o *ObjectContext) HasContextCondition() bool`

HasContextCondition returns a boolean if a field has been set.

### GetContextValueRange

`func (o *ObjectContext) GetContextValueRange() []float32`

GetContextValueRange returns the ContextValueRange field if non-nil, zero value otherwise.

### GetContextValueRangeOk

`func (o *ObjectContext) GetContextValueRangeOk() (*[]float32, bool)`

GetContextValueRangeOk returns a tuple with the ContextValueRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextValueRange

`func (o *ObjectContext) SetContextValueRange(v []float32)`

SetContextValueRange sets ContextValueRange field to given value.

### HasContextValueRange

`func (o *ObjectContext) HasContextValueRange() bool`

HasContextValueRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


