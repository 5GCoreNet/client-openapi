# QosMonitoringInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReqQosMonParams** | [**[]RequestedQosMonitoringParameter**](RequestedQosMonitoringParameter.md) |  | 
**RepFreqs** | [**[]ReportingFrequency**](ReportingFrequency.md) |  | 
**RepThreshDl** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RepThreshUl** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**RepThreshRp** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**WaitTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewQosMonitoringInformation

`func NewQosMonitoringInformation(reqQosMonParams []RequestedQosMonitoringParameter, repFreqs []ReportingFrequency, ) *QosMonitoringInformation`

NewQosMonitoringInformation instantiates a new QosMonitoringInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosMonitoringInformationWithDefaults

`func NewQosMonitoringInformationWithDefaults() *QosMonitoringInformation`

NewQosMonitoringInformationWithDefaults instantiates a new QosMonitoringInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReqQosMonParams

`func (o *QosMonitoringInformation) GetReqQosMonParams() []RequestedQosMonitoringParameter`

GetReqQosMonParams returns the ReqQosMonParams field if non-nil, zero value otherwise.

### GetReqQosMonParamsOk

`func (o *QosMonitoringInformation) GetReqQosMonParamsOk() (*[]RequestedQosMonitoringParameter, bool)`

GetReqQosMonParamsOk returns a tuple with the ReqQosMonParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqQosMonParams

`func (o *QosMonitoringInformation) SetReqQosMonParams(v []RequestedQosMonitoringParameter)`

SetReqQosMonParams sets ReqQosMonParams field to given value.


### GetRepFreqs

`func (o *QosMonitoringInformation) GetRepFreqs() []ReportingFrequency`

GetRepFreqs returns the RepFreqs field if non-nil, zero value otherwise.

### GetRepFreqsOk

`func (o *QosMonitoringInformation) GetRepFreqsOk() (*[]ReportingFrequency, bool)`

GetRepFreqsOk returns a tuple with the RepFreqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepFreqs

`func (o *QosMonitoringInformation) SetRepFreqs(v []ReportingFrequency)`

SetRepFreqs sets RepFreqs field to given value.


### GetRepThreshDl

`func (o *QosMonitoringInformation) GetRepThreshDl() int32`

GetRepThreshDl returns the RepThreshDl field if non-nil, zero value otherwise.

### GetRepThreshDlOk

`func (o *QosMonitoringInformation) GetRepThreshDlOk() (*int32, bool)`

GetRepThreshDlOk returns a tuple with the RepThreshDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshDl

`func (o *QosMonitoringInformation) SetRepThreshDl(v int32)`

SetRepThreshDl sets RepThreshDl field to given value.

### HasRepThreshDl

`func (o *QosMonitoringInformation) HasRepThreshDl() bool`

HasRepThreshDl returns a boolean if a field has been set.

### GetRepThreshUl

`func (o *QosMonitoringInformation) GetRepThreshUl() int32`

GetRepThreshUl returns the RepThreshUl field if non-nil, zero value otherwise.

### GetRepThreshUlOk

`func (o *QosMonitoringInformation) GetRepThreshUlOk() (*int32, bool)`

GetRepThreshUlOk returns a tuple with the RepThreshUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshUl

`func (o *QosMonitoringInformation) SetRepThreshUl(v int32)`

SetRepThreshUl sets RepThreshUl field to given value.

### HasRepThreshUl

`func (o *QosMonitoringInformation) HasRepThreshUl() bool`

HasRepThreshUl returns a boolean if a field has been set.

### GetRepThreshRp

`func (o *QosMonitoringInformation) GetRepThreshRp() int32`

GetRepThreshRp returns the RepThreshRp field if non-nil, zero value otherwise.

### GetRepThreshRpOk

`func (o *QosMonitoringInformation) GetRepThreshRpOk() (*int32, bool)`

GetRepThreshRpOk returns a tuple with the RepThreshRp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepThreshRp

`func (o *QosMonitoringInformation) SetRepThreshRp(v int32)`

SetRepThreshRp sets RepThreshRp field to given value.

### HasRepThreshRp

`func (o *QosMonitoringInformation) HasRepThreshRp() bool`

HasRepThreshRp returns a boolean if a field has been set.

### GetWaitTime

`func (o *QosMonitoringInformation) GetWaitTime() int32`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *QosMonitoringInformation) GetWaitTimeOk() (*int32, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *QosMonitoringInformation) SetWaitTime(v int32)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *QosMonitoringInformation) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.

### GetRepPeriod

`func (o *QosMonitoringInformation) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *QosMonitoringInformation) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *QosMonitoringInformation) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *QosMonitoringInformation) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


