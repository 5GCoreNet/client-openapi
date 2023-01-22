# StringMatchingCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchingString** | Pointer to **string** |  | [optional] 
**MatchingOperator** | [**MatchingOperator**](MatchingOperator.md) |  | 

## Methods

### NewStringMatchingCondition

`func NewStringMatchingCondition(matchingOperator MatchingOperator, ) *StringMatchingCondition`

NewStringMatchingCondition instantiates a new StringMatchingCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringMatchingConditionWithDefaults

`func NewStringMatchingConditionWithDefaults() *StringMatchingCondition`

NewStringMatchingConditionWithDefaults instantiates a new StringMatchingCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchingString

`func (o *StringMatchingCondition) GetMatchingString() string`

GetMatchingString returns the MatchingString field if non-nil, zero value otherwise.

### GetMatchingStringOk

`func (o *StringMatchingCondition) GetMatchingStringOk() (*string, bool)`

GetMatchingStringOk returns a tuple with the MatchingString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingString

`func (o *StringMatchingCondition) SetMatchingString(v string)`

SetMatchingString sets MatchingString field to given value.

### HasMatchingString

`func (o *StringMatchingCondition) HasMatchingString() bool`

HasMatchingString returns a boolean if a field has been set.

### GetMatchingOperator

`func (o *StringMatchingCondition) GetMatchingOperator() MatchingOperator`

GetMatchingOperator returns the MatchingOperator field if non-nil, zero value otherwise.

### GetMatchingOperatorOk

`func (o *StringMatchingCondition) GetMatchingOperatorOk() (*MatchingOperator, bool)`

GetMatchingOperatorOk returns a tuple with the MatchingOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingOperator

`func (o *StringMatchingCondition) SetMatchingOperator(v MatchingOperator)`

SetMatchingOperator sets MatchingOperator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


