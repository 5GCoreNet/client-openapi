# RATContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextAttribute** | Pointer to **string** |  | [optional] 
**ContextCondition** | Pointer to **string** |  | [optional] 
**ContextValueRange** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRATContext

`func NewRATContext() *RATContext`

NewRATContext instantiates a new RATContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRATContextWithDefaults

`func NewRATContextWithDefaults() *RATContext`

NewRATContextWithDefaults instantiates a new RATContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextAttribute

`func (o *RATContext) GetContextAttribute() string`

GetContextAttribute returns the ContextAttribute field if non-nil, zero value otherwise.

### GetContextAttributeOk

`func (o *RATContext) GetContextAttributeOk() (*string, bool)`

GetContextAttributeOk returns a tuple with the ContextAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextAttribute

`func (o *RATContext) SetContextAttribute(v string)`

SetContextAttribute sets ContextAttribute field to given value.

### HasContextAttribute

`func (o *RATContext) HasContextAttribute() bool`

HasContextAttribute returns a boolean if a field has been set.

### GetContextCondition

`func (o *RATContext) GetContextCondition() string`

GetContextCondition returns the ContextCondition field if non-nil, zero value otherwise.

### GetContextConditionOk

`func (o *RATContext) GetContextConditionOk() (*string, bool)`

GetContextConditionOk returns a tuple with the ContextCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextCondition

`func (o *RATContext) SetContextCondition(v string)`

SetContextCondition sets ContextCondition field to given value.

### HasContextCondition

`func (o *RATContext) HasContextCondition() bool`

HasContextCondition returns a boolean if a field has been set.

### GetContextValueRange

`func (o *RATContext) GetContextValueRange() []string`

GetContextValueRange returns the ContextValueRange field if non-nil, zero value otherwise.

### GetContextValueRangeOk

`func (o *RATContext) GetContextValueRangeOk() (*[]string, bool)`

GetContextValueRangeOk returns a tuple with the ContextValueRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextValueRange

`func (o *RATContext) SetContextValueRange(v []string)`

SetContextValueRange sets ContextValueRange field to given value.

### HasContextValueRange

`func (o *RATContext) HasContextValueRange() bool`

HasContextValueRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


