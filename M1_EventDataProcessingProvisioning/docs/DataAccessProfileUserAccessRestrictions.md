# DataAccessProfileUserAccessRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupIds** | **[]string** |  | 
**UserIds** | [**[]DataAccessProfileUserAccessRestrictionsUserIdsInner**](DataAccessProfileUserAccessRestrictionsUserIdsInner.md) |  | 
**AggregationFunctions** | [**[]DataAggregationFunctionType**](DataAggregationFunctionType.md) |  | 

## Methods

### NewDataAccessProfileUserAccessRestrictions

`func NewDataAccessProfileUserAccessRestrictions(groupIds []string, userIds []DataAccessProfileUserAccessRestrictionsUserIdsInner, aggregationFunctions []DataAggregationFunctionType, ) *DataAccessProfileUserAccessRestrictions`

NewDataAccessProfileUserAccessRestrictions instantiates a new DataAccessProfileUserAccessRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAccessProfileUserAccessRestrictionsWithDefaults

`func NewDataAccessProfileUserAccessRestrictionsWithDefaults() *DataAccessProfileUserAccessRestrictions`

NewDataAccessProfileUserAccessRestrictionsWithDefaults instantiates a new DataAccessProfileUserAccessRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupIds

`func (o *DataAccessProfileUserAccessRestrictions) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *DataAccessProfileUserAccessRestrictions) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *DataAccessProfileUserAccessRestrictions) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.


### GetUserIds

`func (o *DataAccessProfileUserAccessRestrictions) GetUserIds() []DataAccessProfileUserAccessRestrictionsUserIdsInner`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *DataAccessProfileUserAccessRestrictions) GetUserIdsOk() (*[]DataAccessProfileUserAccessRestrictionsUserIdsInner, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *DataAccessProfileUserAccessRestrictions) SetUserIds(v []DataAccessProfileUserAccessRestrictionsUserIdsInner)`

SetUserIds sets UserIds field to given value.


### GetAggregationFunctions

`func (o *DataAccessProfileUserAccessRestrictions) GetAggregationFunctions() []DataAggregationFunctionType`

GetAggregationFunctions returns the AggregationFunctions field if non-nil, zero value otherwise.

### GetAggregationFunctionsOk

`func (o *DataAccessProfileUserAccessRestrictions) GetAggregationFunctionsOk() (*[]DataAggregationFunctionType, bool)`

GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationFunctions

`func (o *DataAccessProfileUserAccessRestrictions) SetAggregationFunctions(v []DataAggregationFunctionType)`

SetAggregationFunctions sets AggregationFunctions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


