# MLEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MLEntityId** | Pointer to **string** |  | [optional] 
**InferenceType** | Pointer to **string** |  | [optional] 
**AIMLEntityVersion** | Pointer to **string** |  | [optional] 
**ExpectedRunTimeContext** | Pointer to [**AIMLContext**](AIMLContext.md) |  | [optional] 
**TrainingContext** | Pointer to [**AIMLContext**](AIMLContext.md) |  | [optional] 
**RunTimeContext** | Pointer to [**AIMLContext**](AIMLContext.md) |  | [optional] 

## Methods

### NewMLEntity

`func NewMLEntity() *MLEntity`

NewMLEntity instantiates a new MLEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLEntityWithDefaults

`func NewMLEntityWithDefaults() *MLEntity`

NewMLEntityWithDefaults instantiates a new MLEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMLEntityId

`func (o *MLEntity) GetMLEntityId() string`

GetMLEntityId returns the MLEntityId field if non-nil, zero value otherwise.

### GetMLEntityIdOk

`func (o *MLEntity) GetMLEntityIdOk() (*string, bool)`

GetMLEntityIdOk returns a tuple with the MLEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLEntityId

`func (o *MLEntity) SetMLEntityId(v string)`

SetMLEntityId sets MLEntityId field to given value.

### HasMLEntityId

`func (o *MLEntity) HasMLEntityId() bool`

HasMLEntityId returns a boolean if a field has been set.

### GetInferenceType

`func (o *MLEntity) GetInferenceType() string`

GetInferenceType returns the InferenceType field if non-nil, zero value otherwise.

### GetInferenceTypeOk

`func (o *MLEntity) GetInferenceTypeOk() (*string, bool)`

GetInferenceTypeOk returns a tuple with the InferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceType

`func (o *MLEntity) SetInferenceType(v string)`

SetInferenceType sets InferenceType field to given value.

### HasInferenceType

`func (o *MLEntity) HasInferenceType() bool`

HasInferenceType returns a boolean if a field has been set.

### GetAIMLEntityVersion

`func (o *MLEntity) GetAIMLEntityVersion() string`

GetAIMLEntityVersion returns the AIMLEntityVersion field if non-nil, zero value otherwise.

### GetAIMLEntityVersionOk

`func (o *MLEntity) GetAIMLEntityVersionOk() (*string, bool)`

GetAIMLEntityVersionOk returns a tuple with the AIMLEntityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAIMLEntityVersion

`func (o *MLEntity) SetAIMLEntityVersion(v string)`

SetAIMLEntityVersion sets AIMLEntityVersion field to given value.

### HasAIMLEntityVersion

`func (o *MLEntity) HasAIMLEntityVersion() bool`

HasAIMLEntityVersion returns a boolean if a field has been set.

### GetExpectedRunTimeContext

`func (o *MLEntity) GetExpectedRunTimeContext() AIMLContext`

GetExpectedRunTimeContext returns the ExpectedRunTimeContext field if non-nil, zero value otherwise.

### GetExpectedRunTimeContextOk

`func (o *MLEntity) GetExpectedRunTimeContextOk() (*AIMLContext, bool)`

GetExpectedRunTimeContextOk returns a tuple with the ExpectedRunTimeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRunTimeContext

`func (o *MLEntity) SetExpectedRunTimeContext(v AIMLContext)`

SetExpectedRunTimeContext sets ExpectedRunTimeContext field to given value.

### HasExpectedRunTimeContext

`func (o *MLEntity) HasExpectedRunTimeContext() bool`

HasExpectedRunTimeContext returns a boolean if a field has been set.

### GetTrainingContext

`func (o *MLEntity) GetTrainingContext() AIMLContext`

GetTrainingContext returns the TrainingContext field if non-nil, zero value otherwise.

### GetTrainingContextOk

`func (o *MLEntity) GetTrainingContextOk() (*AIMLContext, bool)`

GetTrainingContextOk returns a tuple with the TrainingContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingContext

`func (o *MLEntity) SetTrainingContext(v AIMLContext)`

SetTrainingContext sets TrainingContext field to given value.

### HasTrainingContext

`func (o *MLEntity) HasTrainingContext() bool`

HasTrainingContext returns a boolean if a field has been set.

### GetRunTimeContext

`func (o *MLEntity) GetRunTimeContext() AIMLContext`

GetRunTimeContext returns the RunTimeContext field if non-nil, zero value otherwise.

### GetRunTimeContextOk

`func (o *MLEntity) GetRunTimeContextOk() (*AIMLContext, bool)`

GetRunTimeContextOk returns a tuple with the RunTimeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeContext

`func (o *MLEntity) SetRunTimeContext(v AIMLContext)`

SetRunTimeContext sets RunTimeContext field to given value.

### HasRunTimeContext

`func (o *MLEntity) HasRunTimeContext() bool`

HasRunTimeContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


