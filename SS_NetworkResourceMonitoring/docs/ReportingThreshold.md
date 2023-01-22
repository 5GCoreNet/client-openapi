# ReportingThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeasThrValues** | [**MeasurementData**](MeasurementData.md) |  | 
**ThrDirection** | [**MatchingDirection**](MatchingDirection.md) |  | 

## Methods

### NewReportingThreshold

`func NewReportingThreshold(measThrValues MeasurementData, thrDirection MatchingDirection, ) *ReportingThreshold`

NewReportingThreshold instantiates a new ReportingThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingThresholdWithDefaults

`func NewReportingThresholdWithDefaults() *ReportingThreshold`

NewReportingThresholdWithDefaults instantiates a new ReportingThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasThrValues

`func (o *ReportingThreshold) GetMeasThrValues() MeasurementData`

GetMeasThrValues returns the MeasThrValues field if non-nil, zero value otherwise.

### GetMeasThrValuesOk

`func (o *ReportingThreshold) GetMeasThrValuesOk() (*MeasurementData, bool)`

GetMeasThrValuesOk returns a tuple with the MeasThrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasThrValues

`func (o *ReportingThreshold) SetMeasThrValues(v MeasurementData)`

SetMeasThrValues sets MeasThrValues field to given value.


### GetThrDirection

`func (o *ReportingThreshold) GetThrDirection() MatchingDirection`

GetThrDirection returns the ThrDirection field if non-nil, zero value otherwise.

### GetThrDirectionOk

`func (o *ReportingThreshold) GetThrDirectionOk() (*MatchingDirection, bool)`

GetThrDirectionOk returns a tuple with the ThrDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrDirection

`func (o *ReportingThreshold) SetThrDirection(v MatchingDirection)`

SetThrDirection sets ThrDirection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


