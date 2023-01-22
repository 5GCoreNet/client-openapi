# ThresholdHysteresis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**High** | [**ThresholdHysteresisHigh**](ThresholdHysteresisHigh.md) |  | 
**Low** | Pointer to **float32** |  | [optional] 

## Methods

### NewThresholdHysteresis

`func NewThresholdHysteresis(high ThresholdHysteresisHigh, ) *ThresholdHysteresis`

NewThresholdHysteresis instantiates a new ThresholdHysteresis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdHysteresisWithDefaults

`func NewThresholdHysteresisWithDefaults() *ThresholdHysteresis`

NewThresholdHysteresisWithDefaults instantiates a new ThresholdHysteresis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigh

`func (o *ThresholdHysteresis) GetHigh() ThresholdHysteresisHigh`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ThresholdHysteresis) GetHighOk() (*ThresholdHysteresisHigh, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ThresholdHysteresis) SetHigh(v ThresholdHysteresisHigh)`

SetHigh sets High field to given value.


### GetLow

`func (o *ThresholdHysteresis) GetLow() float32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *ThresholdHysteresis) GetLowOk() (*float32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *ThresholdHysteresis) SetLow(v float32)`

SetLow sets Low field to given value.

### HasLow

`func (o *ThresholdHysteresis) HasLow() bool`

HasLow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


