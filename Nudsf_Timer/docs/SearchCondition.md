# SearchCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cond** | [**ConditionOperator**](ConditionOperator.md) |  | 
**Units** | [**[]SearchExpression**](SearchExpression.md) |  | 
**SchemaId** | Pointer to **string** | Represents the Identifier of a Meta schema. | [optional] 

## Methods

### NewSearchCondition

`func NewSearchCondition(cond ConditionOperator, units []SearchExpression, ) *SearchCondition`

NewSearchCondition instantiates a new SearchCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchConditionWithDefaults

`func NewSearchConditionWithDefaults() *SearchCondition`

NewSearchConditionWithDefaults instantiates a new SearchCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCond

`func (o *SearchCondition) GetCond() ConditionOperator`

GetCond returns the Cond field if non-nil, zero value otherwise.

### GetCondOk

`func (o *SearchCondition) GetCondOk() (*ConditionOperator, bool)`

GetCondOk returns a tuple with the Cond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCond

`func (o *SearchCondition) SetCond(v ConditionOperator)`

SetCond sets Cond field to given value.


### GetUnits

`func (o *SearchCondition) GetUnits() []SearchExpression`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SearchCondition) GetUnitsOk() (*[]SearchExpression, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *SearchCondition) SetUnits(v []SearchExpression)`

SetUnits sets Units field to given value.


### GetSchemaId

`func (o *SearchCondition) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SearchCondition) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SearchCondition) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *SearchCondition) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


