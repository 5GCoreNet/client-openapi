# FqdnPatternMatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regex** | Pointer to **string** |  | [optional] 
**StringMatchingRule** | Pointer to [**StringMatchingRule**](StringMatchingRule.md) |  | [optional] 

## Methods

### NewFqdnPatternMatchingRule

`func NewFqdnPatternMatchingRule() *FqdnPatternMatchingRule`

NewFqdnPatternMatchingRule instantiates a new FqdnPatternMatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFqdnPatternMatchingRuleWithDefaults

`func NewFqdnPatternMatchingRuleWithDefaults() *FqdnPatternMatchingRule`

NewFqdnPatternMatchingRuleWithDefaults instantiates a new FqdnPatternMatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegex

`func (o *FqdnPatternMatchingRule) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *FqdnPatternMatchingRule) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *FqdnPatternMatchingRule) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *FqdnPatternMatchingRule) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetStringMatchingRule

`func (o *FqdnPatternMatchingRule) GetStringMatchingRule() StringMatchingRule`

GetStringMatchingRule returns the StringMatchingRule field if non-nil, zero value otherwise.

### GetStringMatchingRuleOk

`func (o *FqdnPatternMatchingRule) GetStringMatchingRuleOk() (*StringMatchingRule, bool)`

GetStringMatchingRuleOk returns a tuple with the StringMatchingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringMatchingRule

`func (o *FqdnPatternMatchingRule) SetStringMatchingRule(v StringMatchingRule)`

SetStringMatchingRule sets StringMatchingRule field to given value.

### HasStringMatchingRule

`func (o *FqdnPatternMatchingRule) HasStringMatchingRule() bool`

HasStringMatchingRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


