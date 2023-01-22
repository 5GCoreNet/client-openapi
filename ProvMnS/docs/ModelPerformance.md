# ModelPerformance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InferenceOutputName** | Pointer to **string** |  | [optional] 
**PerformanceMetric** | Pointer to **string** |  | [optional] 
**PerformanceScore** | Pointer to **float32** |  | [optional] 
**DecisionConfidenceScore** | Pointer to **float32** |  | [optional] 

## Methods

### NewModelPerformance

`func NewModelPerformance() *ModelPerformance`

NewModelPerformance instantiates a new ModelPerformance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPerformanceWithDefaults

`func NewModelPerformanceWithDefaults() *ModelPerformance`

NewModelPerformanceWithDefaults instantiates a new ModelPerformance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInferenceOutputName

`func (o *ModelPerformance) GetInferenceOutputName() string`

GetInferenceOutputName returns the InferenceOutputName field if non-nil, zero value otherwise.

### GetInferenceOutputNameOk

`func (o *ModelPerformance) GetInferenceOutputNameOk() (*string, bool)`

GetInferenceOutputNameOk returns a tuple with the InferenceOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceOutputName

`func (o *ModelPerformance) SetInferenceOutputName(v string)`

SetInferenceOutputName sets InferenceOutputName field to given value.

### HasInferenceOutputName

`func (o *ModelPerformance) HasInferenceOutputName() bool`

HasInferenceOutputName returns a boolean if a field has been set.

### GetPerformanceMetric

`func (o *ModelPerformance) GetPerformanceMetric() string`

GetPerformanceMetric returns the PerformanceMetric field if non-nil, zero value otherwise.

### GetPerformanceMetricOk

`func (o *ModelPerformance) GetPerformanceMetricOk() (*string, bool)`

GetPerformanceMetricOk returns a tuple with the PerformanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMetric

`func (o *ModelPerformance) SetPerformanceMetric(v string)`

SetPerformanceMetric sets PerformanceMetric field to given value.

### HasPerformanceMetric

`func (o *ModelPerformance) HasPerformanceMetric() bool`

HasPerformanceMetric returns a boolean if a field has been set.

### GetPerformanceScore

`func (o *ModelPerformance) GetPerformanceScore() float32`

GetPerformanceScore returns the PerformanceScore field if non-nil, zero value otherwise.

### GetPerformanceScoreOk

`func (o *ModelPerformance) GetPerformanceScoreOk() (*float32, bool)`

GetPerformanceScoreOk returns a tuple with the PerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceScore

`func (o *ModelPerformance) SetPerformanceScore(v float32)`

SetPerformanceScore sets PerformanceScore field to given value.

### HasPerformanceScore

`func (o *ModelPerformance) HasPerformanceScore() bool`

HasPerformanceScore returns a boolean if a field has been set.

### GetDecisionConfidenceScore

`func (o *ModelPerformance) GetDecisionConfidenceScore() float32`

GetDecisionConfidenceScore returns the DecisionConfidenceScore field if non-nil, zero value otherwise.

### GetDecisionConfidenceScoreOk

`func (o *ModelPerformance) GetDecisionConfidenceScoreOk() (*float32, bool)`

GetDecisionConfidenceScoreOk returns a tuple with the DecisionConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionConfidenceScore

`func (o *ModelPerformance) SetDecisionConfidenceScore(v float32)`

SetDecisionConfidenceScore sets DecisionConfidenceScore field to given value.

### HasDecisionConfidenceScore

`func (o *ModelPerformance) HasDecisionConfidenceScore() bool`

HasDecisionConfidenceScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


