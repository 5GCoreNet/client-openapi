# PLMNContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextAttribute** | Pointer to **string** |  | [optional] 
**ContextCondition** | Pointer to **string** |  | [optional] 
**ContextValueRange** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewPLMNContext

`func NewPLMNContext() *PLMNContext`

NewPLMNContext instantiates a new PLMNContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPLMNContextWithDefaults

`func NewPLMNContextWithDefaults() *PLMNContext`

NewPLMNContextWithDefaults instantiates a new PLMNContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextAttribute

`func (o *PLMNContext) GetContextAttribute() string`

GetContextAttribute returns the ContextAttribute field if non-nil, zero value otherwise.

### GetContextAttributeOk

`func (o *PLMNContext) GetContextAttributeOk() (*string, bool)`

GetContextAttributeOk returns a tuple with the ContextAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAttribute

`func (o *PLMNContext) SetContextAttribute(v string)`

SetContextAttribute sets ContextAttribute field to given value.

### HasContextAttribute

`func (o *PLMNContext) HasContextAttribute() bool`

HasContextAttribute returns a boolean if a field has been set.

### GetContextCondition

`func (o *PLMNContext) GetContextCondition() string`

GetContextCondition returns the ContextCondition field if non-nil, zero value otherwise.

### GetContextConditionOk

`func (o *PLMNContext) GetContextConditionOk() (*string, bool)`

GetContextConditionOk returns a tuple with the ContextCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextCondition

`func (o *PLMNContext) SetContextCondition(v string)`

SetContextCondition sets ContextCondition field to given value.

### HasContextCondition

`func (o *PLMNContext) HasContextCondition() bool`

HasContextCondition returns a boolean if a field has been set.

### GetContextValueRange

`func (o *PLMNContext) GetContextValueRange() []PlmnId`

GetContextValueRange returns the ContextValueRange field if non-nil, zero value otherwise.

### GetContextValueRangeOk

`func (o *PLMNContext) GetContextValueRangeOk() (*[]PlmnId, bool)`

GetContextValueRangeOk returns a tuple with the ContextValueRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextValueRange

`func (o *PLMNContext) SetContextValueRange(v []PlmnId)`

SetContextValueRange sets ContextValueRange field to given value.

### HasContextValueRange

`func (o *PLMNContext) HasContextValueRange() bool`

HasContextValueRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


