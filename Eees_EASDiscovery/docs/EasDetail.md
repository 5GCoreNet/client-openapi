# EasDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | **string** | Application identifier of the EAS. | 
**ExpectedSvcKPIs** | Pointer to [**ACServiceKPIs**](ACServiceKPIs.md) |  | [optional] 
**MinimumReqSvcKPIs** | Pointer to [**ACServiceKPIs**](ACServiceKPIs.md) |  | [optional] 

## Methods

### NewEasDetail

`func NewEasDetail(easId string, ) *EasDetail`

NewEasDetail instantiates a new EasDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDetailWithDefaults

`func NewEasDetailWithDefaults() *EasDetail`

NewEasDetailWithDefaults instantiates a new EasDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *EasDetail) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *EasDetail) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *EasDetail) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetExpectedSvcKPIs

`func (o *EasDetail) GetExpectedSvcKPIs() ACServiceKPIs`

GetExpectedSvcKPIs returns the ExpectedSvcKPIs field if non-nil, zero value otherwise.

### GetExpectedSvcKPIsOk

`func (o *EasDetail) GetExpectedSvcKPIsOk() (*ACServiceKPIs, bool)`

GetExpectedSvcKPIsOk returns a tuple with the ExpectedSvcKPIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedSvcKPIs

`func (o *EasDetail) SetExpectedSvcKPIs(v ACServiceKPIs)`

SetExpectedSvcKPIs sets ExpectedSvcKPIs field to given value.

### HasExpectedSvcKPIs

`func (o *EasDetail) HasExpectedSvcKPIs() bool`

HasExpectedSvcKPIs returns a boolean if a field has been set.

### GetMinimumReqSvcKPIs

`func (o *EasDetail) GetMinimumReqSvcKPIs() ACServiceKPIs`

GetMinimumReqSvcKPIs returns the MinimumReqSvcKPIs field if non-nil, zero value otherwise.

### GetMinimumReqSvcKPIsOk

`func (o *EasDetail) GetMinimumReqSvcKPIsOk() (*ACServiceKPIs, bool)`

GetMinimumReqSvcKPIsOk returns a tuple with the MinimumReqSvcKPIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumReqSvcKPIs

`func (o *EasDetail) SetMinimumReqSvcKPIs(v ACServiceKPIs)`

SetMinimumReqSvcKPIs sets MinimumReqSvcKPIs field to given value.

### HasMinimumReqSvcKPIs

`func (o *EasDetail) HasMinimumReqSvcKPIs() bool`

HasMinimumReqSvcKPIs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


