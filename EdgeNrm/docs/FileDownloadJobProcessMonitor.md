# FileDownloadJobProcessMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ProgressPercentage** | Pointer to **int32** |  | [optional] 
**ProgressStateInfo** | Pointer to **string** |  | [optional] 
**ResultStateInfo** | Pointer to [**FileDownloadJobProcessMonitorResultStateInfo**](FileDownloadJobProcessMonitorResultStateInfo.md) |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**Timer** | Pointer to **int32** |  | [optional] 

## Methods

### NewFileDownloadJobProcessMonitor

`func NewFileDownloadJobProcessMonitor() *FileDownloadJobProcessMonitor`

NewFileDownloadJobProcessMonitor instantiates a new FileDownloadJobProcessMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDownloadJobProcessMonitorWithDefaults

`func NewFileDownloadJobProcessMonitorWithDefaults() *FileDownloadJobProcessMonitor`

NewFileDownloadJobProcessMonitorWithDefaults instantiates a new FileDownloadJobProcessMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *FileDownloadJobProcessMonitor) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *FileDownloadJobProcessMonitor) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *FileDownloadJobProcessMonitor) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *FileDownloadJobProcessMonitor) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStatus

`func (o *FileDownloadJobProcessMonitor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FileDownloadJobProcessMonitor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FileDownloadJobProcessMonitor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FileDownloadJobProcessMonitor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgressPercentage

`func (o *FileDownloadJobProcessMonitor) GetProgressPercentage() int32`

GetProgressPercentage returns the ProgressPercentage field if non-nil, zero value otherwise.

### GetProgressPercentageOk

`func (o *FileDownloadJobProcessMonitor) GetProgressPercentageOk() (*int32, bool)`

GetProgressPercentageOk returns a tuple with the ProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercentage

`func (o *FileDownloadJobProcessMonitor) SetProgressPercentage(v int32)`

SetProgressPercentage sets ProgressPercentage field to given value.

### HasProgressPercentage

`func (o *FileDownloadJobProcessMonitor) HasProgressPercentage() bool`

HasProgressPercentage returns a boolean if a field has been set.

### GetProgressStateInfo

`func (o *FileDownloadJobProcessMonitor) GetProgressStateInfo() string`

GetProgressStateInfo returns the ProgressStateInfo field if non-nil, zero value otherwise.

### GetProgressStateInfoOk

`func (o *FileDownloadJobProcessMonitor) GetProgressStateInfoOk() (*string, bool)`

GetProgressStateInfoOk returns a tuple with the ProgressStateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressStateInfo

`func (o *FileDownloadJobProcessMonitor) SetProgressStateInfo(v string)`

SetProgressStateInfo sets ProgressStateInfo field to given value.

### HasProgressStateInfo

`func (o *FileDownloadJobProcessMonitor) HasProgressStateInfo() bool`

HasProgressStateInfo returns a boolean if a field has been set.

### GetResultStateInfo

`func (o *FileDownloadJobProcessMonitor) GetResultStateInfo() FileDownloadJobProcessMonitorResultStateInfo`

GetResultStateInfo returns the ResultStateInfo field if non-nil, zero value otherwise.

### GetResultStateInfoOk

`func (o *FileDownloadJobProcessMonitor) GetResultStateInfoOk() (*FileDownloadJobProcessMonitorResultStateInfo, bool)`

GetResultStateInfoOk returns a tuple with the ResultStateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultStateInfo

`func (o *FileDownloadJobProcessMonitor) SetResultStateInfo(v FileDownloadJobProcessMonitorResultStateInfo)`

SetResultStateInfo sets ResultStateInfo field to given value.

### HasResultStateInfo

`func (o *FileDownloadJobProcessMonitor) HasResultStateInfo() bool`

HasResultStateInfo returns a boolean if a field has been set.

### GetStartTime

`func (o *FileDownloadJobProcessMonitor) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FileDownloadJobProcessMonitor) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FileDownloadJobProcessMonitor) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *FileDownloadJobProcessMonitor) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *FileDownloadJobProcessMonitor) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FileDownloadJobProcessMonitor) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FileDownloadJobProcessMonitor) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *FileDownloadJobProcessMonitor) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetTimer

`func (o *FileDownloadJobProcessMonitor) GetTimer() int32`

GetTimer returns the Timer field if non-nil, zero value otherwise.

### GetTimerOk

`func (o *FileDownloadJobProcessMonitor) GetTimerOk() (*int32, bool)`

GetTimerOk returns a tuple with the Timer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimer

`func (o *FileDownloadJobProcessMonitor) SetTimer(v int32)`

SetTimer sets Timer field to given value.

### HasTimer

`func (o *FileDownloadJobProcessMonitor) HasTimer() bool`

HasTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


