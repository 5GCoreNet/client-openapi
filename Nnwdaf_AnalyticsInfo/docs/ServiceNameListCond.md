# ServiceNameListCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** |  | 
**ServiceNameList** | [**[]ServiceName**](ServiceName.md) |  | 

## Methods

### NewServiceNameListCond

`func NewServiceNameListCond(conditionType string, serviceNameList []ServiceName, ) *ServiceNameListCond`

NewServiceNameListCond instantiates a new ServiceNameListCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNameListCondWithDefaults

`func NewServiceNameListCondWithDefaults() *ServiceNameListCond`

NewServiceNameListCondWithDefaults instantiates a new ServiceNameListCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *ServiceNameListCond) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ServiceNameListCond) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ServiceNameListCond) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetServiceNameList

`func (o *ServiceNameListCond) GetServiceNameList() []ServiceName`

GetServiceNameList returns the ServiceNameList field if non-nil, zero value otherwise.

### GetServiceNameListOk

`func (o *ServiceNameListCond) GetServiceNameListOk() (*[]ServiceName, bool)`

GetServiceNameListOk returns a tuple with the ServiceNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNameList

`func (o *ServiceNameListCond) SetServiceNameList(v []ServiceName)`

SetServiceNameList sets ServiceNameList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


