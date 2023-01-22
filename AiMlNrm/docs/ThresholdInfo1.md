# ThresholdInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservedMeasurement** | **string** |  | 
**ObservedValue** | [**ThresholdInfoThresholdValue**](ThresholdInfoThresholdValue.md) |  | 
**ThresholdLevel** | Pointer to [**ThresholdLevelInd**](ThresholdLevelInd.md) |  | [optional] 
**ArmTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewThresholdInfo1

`func NewThresholdInfo1(observedMeasurement string, observedValue ThresholdInfoThresholdValue, ) *ThresholdInfo1`

NewThresholdInfo1 instantiates a new ThresholdInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdInfo1WithDefaults

`func NewThresholdInfo1WithDefaults() *ThresholdInfo1`

NewThresholdInfo1WithDefaults instantiates a new ThresholdInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservedMeasurement

`func (o *ThresholdInfo1) GetObservedMeasurement() string`

GetObservedMeasurement returns the ObservedMeasurement field if non-nil, zero value otherwise.

### GetObservedMeasurementOk

`func (o *ThresholdInfo1) GetObservedMeasurementOk() (*string, bool)`

GetObservedMeasurementOk returns a tuple with the ObservedMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedMeasurement

`func (o *ThresholdInfo1) SetObservedMeasurement(v string)`

SetObservedMeasurement sets ObservedMeasurement field to given value.


### GetObservedValue

`func (o *ThresholdInfo1) GetObservedValue() ThresholdInfoThresholdValue`

GetObservedValue returns the ObservedValue field if non-nil, zero value otherwise.

### GetObservedValueOk

`func (o *ThresholdInfo1) GetObservedValueOk() (*ThresholdInfoThresholdValue, bool)`

GetObservedValueOk returns a tuple with the ObservedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedValue

`func (o *ThresholdInfo1) SetObservedValue(v ThresholdInfoThresholdValue)`

SetObservedValue sets ObservedValue field to given value.


### GetThresholdLevel

`func (o *ThresholdInfo1) GetThresholdLevel() ThresholdLevelInd`

GetThresholdLevel returns the ThresholdLevel field if non-nil, zero value otherwise.

### GetThresholdLevelOk

`func (o *ThresholdInfo1) GetThresholdLevelOk() (*ThresholdLevelInd, bool)`

GetThresholdLevelOk returns a tuple with the ThresholdLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdLevel

`func (o *ThresholdInfo1) SetThresholdLevel(v ThresholdLevelInd)`

SetThresholdLevel sets ThresholdLevel field to given value.

### HasThresholdLevel

`func (o *ThresholdInfo1) HasThresholdLevel() bool`

HasThresholdLevel returns a boolean if a field has been set.

### GetArmTime

`func (o *ThresholdInfo1) GetArmTime() time.Time`

GetArmTime returns the ArmTime field if non-nil, zero value otherwise.

### GetArmTimeOk

`func (o *ThresholdInfo1) GetArmTimeOk() (*time.Time, bool)`

GetArmTimeOk returns a tuple with the ArmTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArmTime

`func (o *ThresholdInfo1) SetArmTime(v time.Time)`

SetArmTime sets ArmTime field to given value.

### HasArmTime

`func (o *ThresholdInfo1) HasArmTime() bool`

HasArmTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


