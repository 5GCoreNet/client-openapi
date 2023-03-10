# ReportingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumOfReports** | Pointer to **int32** | Maximum number of events to be reported for events in a given subscription | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ReportPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewReportingOptions

`func NewReportingOptions() *ReportingOptions`

NewReportingOptions instantiates a new ReportingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingOptionsWithDefaults

`func NewReportingOptionsWithDefaults() *ReportingOptions`

NewReportingOptionsWithDefaults instantiates a new ReportingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumOfReports

`func (o *ReportingOptions) GetMaxNumOfReports() int32`

GetMaxNumOfReports returns the MaxNumOfReports field if non-nil, zero value otherwise.

### GetMaxNumOfReportsOk

`func (o *ReportingOptions) GetMaxNumOfReportsOk() (*int32, bool)`

GetMaxNumOfReportsOk returns a tuple with the MaxNumOfReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfReports

`func (o *ReportingOptions) SetMaxNumOfReports(v int32)`

SetMaxNumOfReports sets MaxNumOfReports field to given value.

### HasMaxNumOfReports

`func (o *ReportingOptions) HasMaxNumOfReports() bool`

HasMaxNumOfReports returns a boolean if a field has been set.

### GetExpiry

`func (o *ReportingOptions) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ReportingOptions) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ReportingOptions) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ReportingOptions) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetReportPeriod

`func (o *ReportingOptions) GetReportPeriod() int32`

GetReportPeriod returns the ReportPeriod field if non-nil, zero value otherwise.

### GetReportPeriodOk

`func (o *ReportingOptions) GetReportPeriodOk() (*int32, bool)`

GetReportPeriodOk returns a tuple with the ReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPeriod

`func (o *ReportingOptions) SetReportPeriod(v int32)`

SetReportPeriod sets ReportPeriod field to given value.

### HasReportPeriod

`func (o *ReportingOptions) HasReportPeriod() bool`

HasReportPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


