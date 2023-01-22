# NumberAverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **float32** | string with format &#39;float&#39; as defined in OpenAPI. | 
**Variance** | **float32** | string with format &#39;float&#39; as defined in OpenAPI. | 
**Skewness** | Pointer to **float32** | string with format &#39;float&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewNumberAverage

`func NewNumberAverage(number float32, variance float32, ) *NumberAverage`

NewNumberAverage instantiates a new NumberAverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberAverageWithDefaults

`func NewNumberAverageWithDefaults() *NumberAverage`

NewNumberAverageWithDefaults instantiates a new NumberAverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *NumberAverage) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NumberAverage) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NumberAverage) SetNumber(v float32)`

SetNumber sets Number field to given value.


### GetVariance

`func (o *NumberAverage) GetVariance() float32`

GetVariance returns the Variance field if non-nil, zero value otherwise.

### GetVarianceOk

`func (o *NumberAverage) GetVarianceOk() (*float32, bool)`

GetVarianceOk returns a tuple with the Variance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariance

`func (o *NumberAverage) SetVariance(v float32)`

SetVariance sets Variance field to given value.


### GetSkewness

`func (o *NumberAverage) GetSkewness() float32`

GetSkewness returns the Skewness field if non-nil, zero value otherwise.

### GetSkewnessOk

`func (o *NumberAverage) GetSkewnessOk() (*float32, bool)`

GetSkewnessOk returns a tuple with the Skewness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkewness

`func (o *NumberAverage) SetSkewness(v float32)`

SetSkewness sets Skewness field to given value.

### HasSkewness

`func (o *NumberAverage) HasSkewness() bool`

HasSkewness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


