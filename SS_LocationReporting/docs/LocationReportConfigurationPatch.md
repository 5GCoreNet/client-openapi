# LocationReportConfigurationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValTgtUe** | Pointer to [**ValTargetUe**](ValTargetUe.md) |  | [optional] 
**MonDur** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**Accuracy** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 

## Methods

### NewLocationReportConfigurationPatch

`func NewLocationReportConfigurationPatch() *LocationReportConfigurationPatch`

NewLocationReportConfigurationPatch instantiates a new LocationReportConfigurationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationReportConfigurationPatchWithDefaults

`func NewLocationReportConfigurationPatchWithDefaults() *LocationReportConfigurationPatch`

NewLocationReportConfigurationPatchWithDefaults instantiates a new LocationReportConfigurationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValTgtUe

`func (o *LocationReportConfigurationPatch) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *LocationReportConfigurationPatch) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *LocationReportConfigurationPatch) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.

### HasValTgtUe

`func (o *LocationReportConfigurationPatch) HasValTgtUe() bool`

HasValTgtUe returns a boolean if a field has been set.

### GetMonDur

`func (o *LocationReportConfigurationPatch) GetMonDur() time.Time`

GetMonDur returns the MonDur field if non-nil, zero value otherwise.

### GetMonDurOk

`func (o *LocationReportConfigurationPatch) GetMonDurOk() (*time.Time, bool)`

GetMonDurOk returns a tuple with the MonDur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonDur

`func (o *LocationReportConfigurationPatch) SetMonDur(v time.Time)`

SetMonDur sets MonDur field to given value.

### HasMonDur

`func (o *LocationReportConfigurationPatch) HasMonDur() bool`

HasMonDur returns a boolean if a field has been set.

### GetRepPeriod

`func (o *LocationReportConfigurationPatch) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *LocationReportConfigurationPatch) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *LocationReportConfigurationPatch) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *LocationReportConfigurationPatch) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetAccuracy

`func (o *LocationReportConfigurationPatch) GetAccuracy() Accuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationReportConfigurationPatch) GetAccuracyOk() (*Accuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationReportConfigurationPatch) SetAccuracy(v Accuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationReportConfigurationPatch) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


