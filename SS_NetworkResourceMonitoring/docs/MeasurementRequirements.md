# MeasurementRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeasDataTypes** | [**[]MeasurementDataType**](MeasurementDataType.md) | Indicates the required the QoS measurement data types. | 
**MeasAggrGranWnd** | Pointer to **int32** | Unsigned integer indicating Averaging Window (see clause 5.7.3.6 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. | [optional] [default to 2000]
**MeasPeriod** | Pointer to [**MeasurementPeriod**](MeasurementPeriod.md) |  | [optional] 

## Methods

### NewMeasurementRequirements

`func NewMeasurementRequirements(measDataTypes []MeasurementDataType, ) *MeasurementRequirements`

NewMeasurementRequirements instantiates a new MeasurementRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementRequirementsWithDefaults

`func NewMeasurementRequirementsWithDefaults() *MeasurementRequirements`

NewMeasurementRequirementsWithDefaults instantiates a new MeasurementRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasDataTypes

`func (o *MeasurementRequirements) GetMeasDataTypes() []MeasurementDataType`

GetMeasDataTypes returns the MeasDataTypes field if non-nil, zero value otherwise.

### GetMeasDataTypesOk

`func (o *MeasurementRequirements) GetMeasDataTypesOk() (*[]MeasurementDataType, bool)`

GetMeasDataTypesOk returns a tuple with the MeasDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasDataTypes

`func (o *MeasurementRequirements) SetMeasDataTypes(v []MeasurementDataType)`

SetMeasDataTypes sets MeasDataTypes field to given value.


### GetMeasAggrGranWnd

`func (o *MeasurementRequirements) GetMeasAggrGranWnd() int32`

GetMeasAggrGranWnd returns the MeasAggrGranWnd field if non-nil, zero value otherwise.

### GetMeasAggrGranWndOk

`func (o *MeasurementRequirements) GetMeasAggrGranWndOk() (*int32, bool)`

GetMeasAggrGranWndOk returns a tuple with the MeasAggrGranWnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasAggrGranWnd

`func (o *MeasurementRequirements) SetMeasAggrGranWnd(v int32)`

SetMeasAggrGranWnd sets MeasAggrGranWnd field to given value.

### HasMeasAggrGranWnd

`func (o *MeasurementRequirements) HasMeasAggrGranWnd() bool`

HasMeasAggrGranWnd returns a boolean if a field has been set.

### GetMeasPeriod

`func (o *MeasurementRequirements) GetMeasPeriod() MeasurementPeriod`

GetMeasPeriod returns the MeasPeriod field if non-nil, zero value otherwise.

### GetMeasPeriodOk

`func (o *MeasurementRequirements) GetMeasPeriodOk() (*MeasurementPeriod, bool)`

GetMeasPeriodOk returns a tuple with the MeasPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasPeriod

`func (o *MeasurementRequirements) SetMeasPeriod(v MeasurementPeriod)`

SetMeasPeriod sets MeasPeriod field to given value.

### HasMeasPeriod

`func (o *MeasurementRequirements) HasMeasPeriod() bool`

HasMeasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


