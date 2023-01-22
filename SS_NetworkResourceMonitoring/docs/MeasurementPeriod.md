# MeasurementPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeasStartTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**MeasDuration** | **int32** | indicating a time in seconds. | 

## Methods

### NewMeasurementPeriod

`func NewMeasurementPeriod(measStartTime time.Time, measDuration int32, ) *MeasurementPeriod`

NewMeasurementPeriod instantiates a new MeasurementPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementPeriodWithDefaults

`func NewMeasurementPeriodWithDefaults() *MeasurementPeriod`

NewMeasurementPeriodWithDefaults instantiates a new MeasurementPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasStartTime

`func (o *MeasurementPeriod) GetMeasStartTime() time.Time`

GetMeasStartTime returns the MeasStartTime field if non-nil, zero value otherwise.

### GetMeasStartTimeOk

`func (o *MeasurementPeriod) GetMeasStartTimeOk() (*time.Time, bool)`

GetMeasStartTimeOk returns a tuple with the MeasStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasStartTime

`func (o *MeasurementPeriod) SetMeasStartTime(v time.Time)`

SetMeasStartTime sets MeasStartTime field to given value.


### GetMeasDuration

`func (o *MeasurementPeriod) GetMeasDuration() int32`

GetMeasDuration returns the MeasDuration field if non-nil, zero value otherwise.

### GetMeasDurationOk

`func (o *MeasurementPeriod) GetMeasDurationOk() (*int32, bool)`

GetMeasDurationOk returns a tuple with the MeasDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasDuration

`func (o *MeasurementPeriod) SetMeasDuration(v int32)`

SetMeasDuration sets MeasDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


