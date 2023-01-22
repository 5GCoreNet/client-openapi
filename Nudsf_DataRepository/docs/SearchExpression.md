# SearchExpression

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cond** | [**ConditionOperator**](ConditionOperator.md) |  | 
**Units** | [**[]SearchExpression**](SearchExpression.md) |  | 
**SchemaId** | Pointer to **string** | Represents the Identifier of a Meta schema. | [optional] 
**Op** | [**ComparisonOperator**](ComparisonOperator.md) |  | 
**Tag** | **string** |  | 
**Value** | **string** |  | 
**RecordIdList** | **[]string** |  | 

## Methods

### NewSearchExpression

`func NewSearchExpression(cond ConditionOperator, units []SearchExpression, op ComparisonOperator, tag string, value string, recordIdList []string, ) *SearchExpression`

NewSearchExpression instantiates a new SearchExpression object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchExpressionWithDefaults

`func NewSearchExpressionWithDefaults() *SearchExpression`

NewSearchExpressionWithDefaults instantiates a new SearchExpression object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCond

`func (o *SearchExpression) GetCond() ConditionOperator`

GetCond returns the Cond field if non-nil, zero value otherwise.

### GetCondOk

`func (o *SearchExpression) GetCondOk() (*ConditionOperator, bool)`

GetCondOk returns a tuple with the Cond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCond

`func (o *SearchExpression) SetCond(v ConditionOperator)`

SetCond sets Cond field to given value.


### GetUnits

`func (o *SearchExpression) GetUnits() []SearchExpression`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *SearchExpression) GetUnitsOk() (*[]SearchExpression, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *SearchExpression) SetUnits(v []SearchExpression)`

SetUnits sets Units field to given value.


### GetSchemaId

`func (o *SearchExpression) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SearchExpression) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SearchExpression) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *SearchExpression) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetOp

`func (o *SearchExpression) GetOp() ComparisonOperator`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *SearchExpression) GetOpOk() (*ComparisonOperator, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *SearchExpression) SetOp(v ComparisonOperator)`

SetOp sets Op field to given value.


### GetTag

`func (o *SearchExpression) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SearchExpression) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SearchExpression) SetTag(v string)`

SetTag sets Tag field to given value.


### GetValue

`func (o *SearchExpression) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchExpression) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchExpression) SetValue(v string)`

SetValue sets Value field to given value.


### GetRecordIdList

`func (o *SearchExpression) GetRecordIdList() []string`

GetRecordIdList returns the RecordIdList field if non-nil, zero value otherwise.

### GetRecordIdListOk

`func (o *SearchExpression) GetRecordIdListOk() (*[]string, bool)`

GetRecordIdListOk returns a tuple with the RecordIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordIdList

`func (o *SearchExpression) SetRecordIdList(v []string)`

SetRecordIdList sets RecordIdList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


