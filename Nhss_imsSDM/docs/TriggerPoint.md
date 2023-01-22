# TriggerPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | [**TypeOfCondition**](TypeOfCondition.md) |  | 
**SptList** | [**[]Spt**](Spt.md) |  | 

## Methods

### NewTriggerPoint

`func NewTriggerPoint(conditionType TypeOfCondition, sptList []Spt, ) *TriggerPoint`

NewTriggerPoint instantiates a new TriggerPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerPointWithDefaults

`func NewTriggerPointWithDefaults() *TriggerPoint`

NewTriggerPointWithDefaults instantiates a new TriggerPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *TriggerPoint) GetConditionType() TypeOfCondition`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *TriggerPoint) GetConditionTypeOk() (*TypeOfCondition, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *TriggerPoint) SetConditionType(v TypeOfCondition)`

SetConditionType sets ConditionType field to given value.


### GetSptList

`func (o *TriggerPoint) GetSptList() []Spt`

GetSptList returns the SptList field if non-nil, zero value otherwise.

### GetSptListOk

`func (o *TriggerPoint) GetSptListOk() (*[]Spt, bool)`

GetSptListOk returns a tuple with the SptList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSptList

`func (o *TriggerPoint) SetSptList(v []Spt)`

SetSptList sets SptList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


