# ExpectationTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetName** | Pointer to **string** |  | [optional] 
**TargetCondition** | Pointer to [**Condition**](Condition.md) |  | [optional] 
**TargetValueRange** | Pointer to **float32** |  | [optional] 
**TargetContexts** | Pointer to [**[]TargetContext**](TargetContext.md) |  | [optional] 

## Methods

### NewExpectationTarget

`func NewExpectationTarget() *ExpectationTarget`

NewExpectationTarget instantiates a new ExpectationTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpectationTargetWithDefaults

`func NewExpectationTargetWithDefaults() *ExpectationTarget`

NewExpectationTargetWithDefaults instantiates a new ExpectationTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetName

`func (o *ExpectationTarget) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ExpectationTarget) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ExpectationTarget) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ExpectationTarget) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetCondition

`func (o *ExpectationTarget) GetTargetCondition() Condition`

GetTargetCondition returns the TargetCondition field if non-nil, zero value otherwise.

### GetTargetConditionOk

`func (o *ExpectationTarget) GetTargetConditionOk() (*Condition, bool)`

GetTargetConditionOk returns a tuple with the TargetCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCondition

`func (o *ExpectationTarget) SetTargetCondition(v Condition)`

SetTargetCondition sets TargetCondition field to given value.

### HasTargetCondition

`func (o *ExpectationTarget) HasTargetCondition() bool`

HasTargetCondition returns a boolean if a field has been set.

### GetTargetValueRange

`func (o *ExpectationTarget) GetTargetValueRange() float32`

GetTargetValueRange returns the TargetValueRange field if non-nil, zero value otherwise.

### GetTargetValueRangeOk

`func (o *ExpectationTarget) GetTargetValueRangeOk() (*float32, bool)`

GetTargetValueRangeOk returns a tuple with the TargetValueRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValueRange

`func (o *ExpectationTarget) SetTargetValueRange(v float32)`

SetTargetValueRange sets TargetValueRange field to given value.

### HasTargetValueRange

`func (o *ExpectationTarget) HasTargetValueRange() bool`

HasTargetValueRange returns a boolean if a field has been set.

### GetTargetContexts

`func (o *ExpectationTarget) GetTargetContexts() []TargetContext`

GetTargetContexts returns the TargetContexts field if non-nil, zero value otherwise.

### GetTargetContextsOk

`func (o *ExpectationTarget) GetTargetContextsOk() (*[]TargetContext, bool)`

GetTargetContextsOk returns a tuple with the TargetContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetContexts

`func (o *ExpectationTarget) SetTargetContexts(v []TargetContext)`

SetTargetContexts sets TargetContexts field to given value.

### HasTargetContexts

`func (o *ExpectationTarget) HasTargetContexts() bool`

HasTargetContexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


