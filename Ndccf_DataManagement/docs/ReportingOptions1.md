# ReportingOptions1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportMode** | Pointer to [**EventReportMode**](EventReportMode.md) |  | [optional] 
**MaxNumOfReports** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SamplingRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 
**GuardTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ReportPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NotifFlag** | Pointer to [**NotificationFlag**](NotificationFlag.md) |  | [optional] 

## Methods

### NewReportingOptions1

`func NewReportingOptions1() *ReportingOptions1`

NewReportingOptions1 instantiates a new ReportingOptions1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingOptions1WithDefaults

`func NewReportingOptions1WithDefaults() *ReportingOptions1`

NewReportingOptions1WithDefaults instantiates a new ReportingOptions1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportMode

`func (o *ReportingOptions1) GetReportMode() EventReportMode`

GetReportMode returns the ReportMode field if non-nil, zero value otherwise.

### GetReportModeOk

`func (o *ReportingOptions1) GetReportModeOk() (*EventReportMode, bool)`

GetReportModeOk returns a tuple with the ReportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportMode

`func (o *ReportingOptions1) SetReportMode(v EventReportMode)`

SetReportMode sets ReportMode field to given value.

### HasReportMode

`func (o *ReportingOptions1) HasReportMode() bool`

HasReportMode returns a boolean if a field has been set.

### GetMaxNumOfReports

`func (o *ReportingOptions1) GetMaxNumOfReports() int32`

GetMaxNumOfReports returns the MaxNumOfReports field if non-nil, zero value otherwise.

### GetMaxNumOfReportsOk

`func (o *ReportingOptions1) GetMaxNumOfReportsOk() (*int32, bool)`

GetMaxNumOfReportsOk returns a tuple with the MaxNumOfReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumOfReports

`func (o *ReportingOptions1) SetMaxNumOfReports(v int32)`

SetMaxNumOfReports sets MaxNumOfReports field to given value.

### HasMaxNumOfReports

`func (o *ReportingOptions1) HasMaxNumOfReports() bool`

HasMaxNumOfReports returns a boolean if a field has been set.

### GetExpiry

`func (o *ReportingOptions1) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ReportingOptions1) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ReportingOptions1) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ReportingOptions1) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSamplingRatio

`func (o *ReportingOptions1) GetSamplingRatio() int32`

GetSamplingRatio returns the SamplingRatio field if non-nil, zero value otherwise.

### GetSamplingRatioOk

`func (o *ReportingOptions1) GetSamplingRatioOk() (*int32, bool)`

GetSamplingRatioOk returns a tuple with the SamplingRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingRatio

`func (o *ReportingOptions1) SetSamplingRatio(v int32)`

SetSamplingRatio sets SamplingRatio field to given value.

### HasSamplingRatio

`func (o *ReportingOptions1) HasSamplingRatio() bool`

HasSamplingRatio returns a boolean if a field has been set.

### GetGuardTime

`func (o *ReportingOptions1) GetGuardTime() int32`

GetGuardTime returns the GuardTime field if non-nil, zero value otherwise.

### GetGuardTimeOk

`func (o *ReportingOptions1) GetGuardTimeOk() (*int32, bool)`

GetGuardTimeOk returns a tuple with the GuardTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardTime

`func (o *ReportingOptions1) SetGuardTime(v int32)`

SetGuardTime sets GuardTime field to given value.

### HasGuardTime

`func (o *ReportingOptions1) HasGuardTime() bool`

HasGuardTime returns a boolean if a field has been set.

### GetReportPeriod

`func (o *ReportingOptions1) GetReportPeriod() int32`

GetReportPeriod returns the ReportPeriod field if non-nil, zero value otherwise.

### GetReportPeriodOk

`func (o *ReportingOptions1) GetReportPeriodOk() (*int32, bool)`

GetReportPeriodOk returns a tuple with the ReportPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPeriod

`func (o *ReportingOptions1) SetReportPeriod(v int32)`

SetReportPeriod sets ReportPeriod field to given value.

### HasReportPeriod

`func (o *ReportingOptions1) HasReportPeriod() bool`

HasReportPeriod returns a boolean if a field has been set.

### GetNotifFlag

`func (o *ReportingOptions1) GetNotifFlag() NotificationFlag`

GetNotifFlag returns the NotifFlag field if non-nil, zero value otherwise.

### GetNotifFlagOk

`func (o *ReportingOptions1) GetNotifFlagOk() (*NotificationFlag, bool)`

GetNotifFlagOk returns a tuple with the NotifFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifFlag

`func (o *ReportingOptions1) SetNotifFlag(v NotificationFlag)`

SetNotifFlag sets NotifFlag field to given value.

### HasNotifFlag

`func (o *ReportingOptions1) HasNotifFlag() bool`

HasNotifFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


