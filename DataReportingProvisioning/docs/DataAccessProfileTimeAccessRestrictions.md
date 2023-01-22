# DataAccessProfileTimeAccessRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | indicating a time in seconds. | 
**AggregationFunctions** | [**[]DataAggregationFunctionType**](DataAggregationFunctionType.md) |  | 

## Methods

### NewDataAccessProfileTimeAccessRestrictions

`func NewDataAccessProfileTimeAccessRestrictions(duration int32, aggregationFunctions []DataAggregationFunctionType, ) *DataAccessProfileTimeAccessRestrictions`

NewDataAccessProfileTimeAccessRestrictions instantiates a new DataAccessProfileTimeAccessRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAccessProfileTimeAccessRestrictionsWithDefaults

`func NewDataAccessProfileTimeAccessRestrictionsWithDefaults() *DataAccessProfileTimeAccessRestrictions`

NewDataAccessProfileTimeAccessRestrictionsWithDefaults instantiates a new DataAccessProfileTimeAccessRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DataAccessProfileTimeAccessRestrictions) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DataAccessProfileTimeAccessRestrictions) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DataAccessProfileTimeAccessRestrictions) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetAggregationFunctions

`func (o *DataAccessProfileTimeAccessRestrictions) GetAggregationFunctions() []DataAggregationFunctionType`

GetAggregationFunctions returns the AggregationFunctions field if non-nil, zero value otherwise.

### GetAggregationFunctionsOk

`func (o *DataAccessProfileTimeAccessRestrictions) GetAggregationFunctionsOk() (*[]DataAggregationFunctionType, bool)`

GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunctions

`func (o *DataAccessProfileTimeAccessRestrictions) SetAggregationFunctions(v []DataAggregationFunctionType)`

SetAggregationFunctions sets AggregationFunctions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


