# QosMonitoringInformationRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReqQosMonParams** | Pointer to [**[]RequestedQosMonitoringParameter**](RequestedQosMonitoringParameter.md) |  | [optional] 
**RepFreqs** | Pointer to [**[]ReportingFrequency**](ReportingFrequency.md) |  | [optional] 
**RepThreshDl** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**RepThreshUl** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**RepThreshRp** | Pointer to **NullableInt32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**WaitTime** | Pointer to **NullableInt32** | indicating a time in seconds with OpenAPI defined &#39;nullable: true&#39; property. | [optional] 
**RepPeriod** | Pointer to **NullableInt32** | indicating a time in seconds with OpenAPI defined &#39;nullable: true&#39; property. | [optional] 

## Methods

### NewQosMonitoringInformationRm

`func NewQosMonitoringInformationRm() *QosMonitoringInformationRm`

NewQosMonitoringInformationRm instantiates a new QosMonitoringInformationRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosMonitoringInformationRmWithDefaults

`func NewQosMonitoringInformationRmWithDefaults() *QosMonitoringInformationRm`

NewQosMonitoringInformationRmWithDefaults instantiates a new QosMonitoringInformationRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReqQosMonParams

`func (o *QosMonitoringInformationRm) GetReqQosMonParams() []RequestedQosMonitoringParameter`

GetReqQosMonParams returns the ReqQosMonParams field if non-nil, zero value otherwise.

### GetReqQosMonParamsOk

`func (o *QosMonitoringInformationRm) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool)`

GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQosMonParams

`func (o *QosMonitoringInformationRm) SetReqQosMonParams(v []RequestedQosMonitoringParameter)`

SetReqQosMonParams sets ReqQosMonParams field to given value.

### HasReqQosMonParams

`func (o *QosMonitoringInformationRm) HasReqQosMonParams() bool`

HasReqQosMonParams returns a boolean if a field has been set.

### GetRepFreqs

`func (o *QosMonitoringInformationRm) GetRepFreqs() []ReportingFrequency`

GetRepFreqs returns the RepFreqs field if non-nil, zero value otherwise.

### GetRepFreqsOk

`func (o *QosMonitoringInformationRm) GetRepFreqsOk() (*[]ReportingFrequency, bool)`

GetRepFreqsOk returns a tuple with the RepFreqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepFreqs

`func (o *QosMonitoringInformationRm) SetRepFreqs(v []ReportingFrequency)`

SetRepFreqs sets RepFreqs field to given value.

### HasRepFreqs

`func (o *QosMonitoringInformationRm) HasRepFreqs() bool`

HasRepFreqs returns a boolean if a field has been set.

### GetRepThreshDl

`func (o *QosMonitoringInformationRm) GetRepThreshDl() int32`

GetRepThreshDl returns the RepThreshDl field if non-nil, zero value otherwise.

### GetRepThreshDlOk

`func (o *QosMonitoringInformationRm) GetRepThreshDlOk() (*int32, bool)`

GetRepThreshDlOk returns a tuple with the RepThreshDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshDl

`func (o *QosMonitoringInformationRm) SetRepThreshDl(v int32)`

SetRepThreshDl sets RepThreshDl field to given value.

### HasRepThreshDl

`func (o *QosMonitoringInformationRm) HasRepThreshDl() bool`

HasRepThreshDl returns a boolean if a field has been set.

### SetRepThreshDlNil

`func (o *QosMonitoringInformationRm) SetRepThreshDlNil(b bool)`

 SetRepThreshDlNil sets the value for RepThreshDl to be an explicit nil

### UnsetRepThreshDl
`func (o *QosMonitoringInformationRm) UnsetRepThreshDl()`

UnsetRepThreshDl ensures that no value is present for RepThreshDl, not even an explicit nil
### GetRepThreshUl

`func (o *QosMonitoringInformationRm) GetRepThreshUl() int32`

GetRepThreshUl returns the RepThreshUl field if non-nil, zero value otherwise.

### GetRepThreshUlOk

`func (o *QosMonitoringInformationRm) GetRepThreshUlOk() (*int32, bool)`

GetRepThreshUlOk returns a tuple with the RepThreshUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshUl

`func (o *QosMonitoringInformationRm) SetRepThreshUl(v int32)`

SetRepThreshUl sets RepThreshUl field to given value.

### HasRepThreshUl

`func (o *QosMonitoringInformationRm) HasRepThreshUl() bool`

HasRepThreshUl returns a boolean if a field has been set.

### SetRepThreshUlNil

`func (o *QosMonitoringInformationRm) SetRepThreshUlNil(b bool)`

 SetRepThreshUlNil sets the value for RepThreshUl to be an explicit nil

### UnsetRepThreshUl
`func (o *QosMonitoringInformationRm) UnsetRepThreshUl()`

UnsetRepThreshUl ensures that no value is present for RepThreshUl, not even an explicit nil
### GetRepThreshRp

`func (o *QosMonitoringInformationRm) GetRepThreshRp() int32`

GetRepThreshRp returns the RepThreshRp field if non-nil, zero value otherwise.

### GetRepThreshRpOk

`func (o *QosMonitoringInformationRm) GetRepThreshRpOk() (*int32, bool)`

GetRepThreshRpOk returns a tuple with the RepThreshRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshRp

`func (o *QosMonitoringInformationRm) SetRepThreshRp(v int32)`

SetRepThreshRp sets RepThreshRp field to given value.

### HasRepThreshRp

`func (o *QosMonitoringInformationRm) HasRepThreshRp() bool`

HasRepThreshRp returns a boolean if a field has been set.

### SetRepThreshRpNil

`func (o *QosMonitoringInformationRm) SetRepThreshRpNil(b bool)`

 SetRepThreshRpNil sets the value for RepThreshRp to be an explicit nil

### UnsetRepThreshRp
`func (o *QosMonitoringInformationRm) UnsetRepThreshRp()`

UnsetRepThreshRp ensures that no value is present for RepThreshRp, not even an explicit nil
### GetWaitTime

`func (o *QosMonitoringInformationRm) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *QosMonitoringInformationRm) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *QosMonitoringInformationRm) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *QosMonitoringInformationRm) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.

### SetWaitTimeNil

`func (o *QosMonitoringInformationRm) SetWaitTimeNil(b bool)`

 SetWaitTimeNil sets the value for WaitTime to be an explicit nil

### UnsetWaitTime
`func (o *QosMonitoringInformationRm) UnsetWaitTime()`

UnsetWaitTime ensures that no value is present for WaitTime, not even an explicit nil
### GetRepPeriod

`func (o *QosMonitoringInformationRm) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *QosMonitoringInformationRm) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *QosMonitoringInformationRm) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *QosMonitoringInformationRm) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### SetRepPeriodNil

`func (o *QosMonitoringInformationRm) SetRepPeriodNil(b bool)`

 SetRepPeriodNil sets the value for RepPeriod to be an explicit nil

### UnsetRepPeriod
`func (o *QosMonitoringInformationRm) UnsetRepPeriod()`

UnsetRepPeriod ensures that no value is present for RepPeriod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


