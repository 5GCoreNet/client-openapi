# FailureReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValUeIds** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of VAL UE(s) whose measurement data is not obtained successfully.  | [optional] 
**ValStreamIds** | Pointer to **[]string** | List of VAL stream ID(s) whose measurement data is not obtained successfully.  | [optional] 
**FailureReason** | Pointer to [**FailureReason**](FailureReason.md) |  | [optional] 
**MeasDataType** | [**MeasurementDataType**](MeasurementDataType.md) |  | 

## Methods

### NewFailureReport

`func NewFailureReport(measDataType MeasurementDataType, ) *FailureReport`

NewFailureReport instantiates a new FailureReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureReportWithDefaults

`func NewFailureReportWithDefaults() *FailureReport`

NewFailureReportWithDefaults instantiates a new FailureReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValUeIds

`func (o *FailureReport) GetValUeIds() []ValTargetUe`

GetValUeIds returns the ValUeIds field if non-nil, zero value otherwise.

### GetValUeIdsOk

`func (o *FailureReport) GetValUeIdsOk() (*[]ValTargetUe, bool)`

GetValUeIdsOk returns a tuple with the ValUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValUeIds

`func (o *FailureReport) SetValUeIds(v []ValTargetUe)`

SetValUeIds sets ValUeIds field to given value.

### HasValUeIds

`func (o *FailureReport) HasValUeIds() bool`

HasValUeIds returns a boolean if a field has been set.

### GetValStreamIds

`func (o *FailureReport) GetValStreamIds() []string`

GetValStreamIds returns the ValStreamIds field if non-nil, zero value otherwise.

### GetValStreamIdsOk

`func (o *FailureReport) GetValStreamIdsOk() (*[]string, bool)`

GetValStreamIdsOk returns a tuple with the ValStreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValStreamIds

`func (o *FailureReport) SetValStreamIds(v []string)`

SetValStreamIds sets ValStreamIds field to given value.

### HasValStreamIds

`func (o *FailureReport) HasValStreamIds() bool`

HasValStreamIds returns a boolean if a field has been set.

### GetFailureReason

`func (o *FailureReport) GetFailureReason() FailureReason`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *FailureReport) GetFailureReasonOk() (*FailureReason, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *FailureReport) SetFailureReason(v FailureReason)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *FailureReport) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetMeasDataType

`func (o *FailureReport) GetMeasDataType() MeasurementDataType`

GetMeasDataType returns the MeasDataType field if non-nil, zero value otherwise.

### GetMeasDataTypeOk

`func (o *FailureReport) GetMeasDataTypeOk() (*MeasurementDataType, bool)`

GetMeasDataTypeOk returns a tuple with the MeasDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasDataType

`func (o *FailureReport) SetMeasDataType(v MeasurementDataType)`

SetMeasDataType sets MeasDataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


