# MonitoringReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValUeIds** | Pointer to [**[]ValTargetUe**](ValTargetUe.md) | List of VAL UEs whose QoS monitoring data is requested. | [optional] 
**ValGroupId** | Pointer to **string** | The VAL Group Id which QoS monitoring data is requested. | [optional] 
**ValStreamIds** | Pointer to **[]string** | List of VAL streams for which QoS monitoring data is requested. | [optional] 
**MeasData** | [**MeasurementData**](MeasurementData.md) |  | 
**FailureRep** | Pointer to [**[]FailureReport**](FailureReport.md) | The failure report indicating the VAL UE(s) or VAL Stream ID(s) whose measurement data is not obtained successfully.  | [optional] 
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewMonitoringReport

`func NewMonitoringReport(measData MeasurementData, timestamp time.Time, ) *MonitoringReport`

NewMonitoringReport instantiates a new MonitoringReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringReportWithDefaults

`func NewMonitoringReportWithDefaults() *MonitoringReport`

NewMonitoringReportWithDefaults instantiates a new MonitoringReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValUeIds

`func (o *MonitoringReport) GetValUeIds() []ValTargetUe`

GetValUeIds returns the ValUeIds field if non-nil, zero value otherwise.

### GetValUeIdsOk

`func (o *MonitoringReport) GetValUeIdsOk() (*[]ValTargetUe, bool)`

GetValUeIdsOk returns a tuple with the ValUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValUeIds

`func (o *MonitoringReport) SetValUeIds(v []ValTargetUe)`

SetValUeIds sets ValUeIds field to given value.

### HasValUeIds

`func (o *MonitoringReport) HasValUeIds() bool`

HasValUeIds returns a boolean if a field has been set.

### GetValGroupId

`func (o *MonitoringReport) GetValGroupId() string`

GetValGroupId returns the ValGroupId field if non-nil, zero value otherwise.

### GetValGroupIdOk

`func (o *MonitoringReport) GetValGroupIdOk() (*string, bool)`

GetValGroupIdOk returns a tuple with the ValGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGroupId

`func (o *MonitoringReport) SetValGroupId(v string)`

SetValGroupId sets ValGroupId field to given value.

### HasValGroupId

`func (o *MonitoringReport) HasValGroupId() bool`

HasValGroupId returns a boolean if a field has been set.

### GetValStreamIds

`func (o *MonitoringReport) GetValStreamIds() []string`

GetValStreamIds returns the ValStreamIds field if non-nil, zero value otherwise.

### GetValStreamIdsOk

`func (o *MonitoringReport) GetValStreamIdsOk() (*[]string, bool)`

GetValStreamIdsOk returns a tuple with the ValStreamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValStreamIds

`func (o *MonitoringReport) SetValStreamIds(v []string)`

SetValStreamIds sets ValStreamIds field to given value.

### HasValStreamIds

`func (o *MonitoringReport) HasValStreamIds() bool`

HasValStreamIds returns a boolean if a field has been set.

### GetMeasData

`func (o *MonitoringReport) GetMeasData() MeasurementData`

GetMeasData returns the MeasData field if non-nil, zero value otherwise.

### GetMeasDataOk

`func (o *MonitoringReport) GetMeasDataOk() (*MeasurementData, bool)`

GetMeasDataOk returns a tuple with the MeasData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasData

`func (o *MonitoringReport) SetMeasData(v MeasurementData)`

SetMeasData sets MeasData field to given value.


### GetFailureRep

`func (o *MonitoringReport) GetFailureRep() []FailureReport`

GetFailureRep returns the FailureRep field if non-nil, zero value otherwise.

### GetFailureRepOk

`func (o *MonitoringReport) GetFailureRepOk() (*[]FailureReport, bool)`

GetFailureRepOk returns a tuple with the FailureRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRep

`func (o *MonitoringReport) SetFailureRep(v []FailureReport)`

SetFailureRep sets FailureRep field to given value.

### HasFailureRep

`func (o *MonitoringReport) HasFailureRep() bool`

HasFailureRep returns a boolean if a field has been set.

### GetTimestamp

`func (o *MonitoringReport) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MonitoringReport) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MonitoringReport) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


