# PerformanceInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeasObjDn** | [**MeasObjDnType**](MeasObjDnType.md) |  | 
**PerformanceMetrics** | **[]string** | an ordered list of performance metric names (see clause 4.4.1 of 3GPP TS 28.622[11]) whose values are to be reported by the Performance Data Stream Units (see Annex C of TS 28.550 [42]) via this stream. Performance metrics include measurement and KPI | 
**JobId** | Pointer to **string** |  | [optional] 

## Methods

### NewPerformanceInfoType

`func NewPerformanceInfoType(measObjDn MeasObjDnType, performanceMetrics []string, ) *PerformanceInfoType`

NewPerformanceInfoType instantiates a new PerformanceInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceInfoTypeWithDefaults

`func NewPerformanceInfoTypeWithDefaults() *PerformanceInfoType`

NewPerformanceInfoTypeWithDefaults instantiates a new PerformanceInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeasObjDn

`func (o *PerformanceInfoType) GetMeasObjDn() MeasObjDnType`

GetMeasObjDn returns the MeasObjDn field if non-nil, zero value otherwise.

### GetMeasObjDnOk

`func (o *PerformanceInfoType) GetMeasObjDnOk() (*MeasObjDnType, bool)`

GetMeasObjDnOk returns a tuple with the MeasObjDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasObjDn

`func (o *PerformanceInfoType) SetMeasObjDn(v MeasObjDnType)`

SetMeasObjDn sets MeasObjDn field to given value.


### GetPerformanceMetrics

`func (o *PerformanceInfoType) GetPerformanceMetrics() []string`

GetPerformanceMetrics returns the PerformanceMetrics field if non-nil, zero value otherwise.

### GetPerformanceMetricsOk

`func (o *PerformanceInfoType) GetPerformanceMetricsOk() (*[]string, bool)`

GetPerformanceMetricsOk returns a tuple with the PerformanceMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceMetrics

`func (o *PerformanceInfoType) SetPerformanceMetrics(v []string)`

SetPerformanceMetrics sets PerformanceMetrics field to given value.


### GetJobId

`func (o *PerformanceInfoType) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *PerformanceInfoType) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *PerformanceInfoType) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *PerformanceInfoType) HasJobId() bool`

HasJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


