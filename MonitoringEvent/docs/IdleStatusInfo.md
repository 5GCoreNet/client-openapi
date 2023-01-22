# IdleStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**EdrxCycleLength** | Pointer to **float32** |  | [optional] 
**SuggestedNumberOfDlPackets** | Pointer to **int32** | Identifies the number of packets shall be buffered in the serving gateway. It shall be present if the idle status indication is requested by the SCS/AS with \&quot;idleStatusIndication\&quot; in the \&quot;monitoringEventSubscription\&quot; sets to \&quot;true\&quot;. | [optional] 
**IdleStatusTimestamp** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**PeriodicAUTimer** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 

## Methods

### NewIdleStatusInfo

`func NewIdleStatusInfo() *IdleStatusInfo`

NewIdleStatusInfo instantiates a new IdleStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdleStatusInfoWithDefaults

`func NewIdleStatusInfoWithDefaults() *IdleStatusInfo`

NewIdleStatusInfoWithDefaults instantiates a new IdleStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveTime

`func (o *IdleStatusInfo) GetActiveTime() int32`

GetActiveTime returns the ActiveTime field if non-nil, zero value otherwise.

### GetActiveTimeOk

`func (o *IdleStatusInfo) GetActiveTimeOk() (*int32, bool)`

GetActiveTimeOk returns a tuple with the ActiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTime

`func (o *IdleStatusInfo) SetActiveTime(v int32)`

SetActiveTime sets ActiveTime field to given value.

### HasActiveTime

`func (o *IdleStatusInfo) HasActiveTime() bool`

HasActiveTime returns a boolean if a field has been set.

### GetEdrxCycleLength

`func (o *IdleStatusInfo) GetEdrxCycleLength() float32`

GetEdrxCycleLength returns the EdrxCycleLength field if non-nil, zero value otherwise.

### GetEdrxCycleLengthOk

`func (o *IdleStatusInfo) GetEdrxCycleLengthOk() (*float32, bool)`

GetEdrxCycleLengthOk returns a tuple with the EdrxCycleLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdrxCycleLength

`func (o *IdleStatusInfo) SetEdrxCycleLength(v float32)`

SetEdrxCycleLength sets EdrxCycleLength field to given value.

### HasEdrxCycleLength

`func (o *IdleStatusInfo) HasEdrxCycleLength() bool`

HasEdrxCycleLength returns a boolean if a field has been set.

### GetSuggestedNumberOfDlPackets

`func (o *IdleStatusInfo) GetSuggestedNumberOfDlPackets() int32`

GetSuggestedNumberOfDlPackets returns the SuggestedNumberOfDlPackets field if non-nil, zero value otherwise.

### GetSuggestedNumberOfDlPacketsOk

`func (o *IdleStatusInfo) GetSuggestedNumberOfDlPacketsOk() (*int32, bool)`

GetSuggestedNumberOfDlPacketsOk returns a tuple with the SuggestedNumberOfDlPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedNumberOfDlPackets

`func (o *IdleStatusInfo) SetSuggestedNumberOfDlPackets(v int32)`

SetSuggestedNumberOfDlPackets sets SuggestedNumberOfDlPackets field to given value.

### HasSuggestedNumberOfDlPackets

`func (o *IdleStatusInfo) HasSuggestedNumberOfDlPackets() bool`

HasSuggestedNumberOfDlPackets returns a boolean if a field has been set.

### GetIdleStatusTimestamp

`func (o *IdleStatusInfo) GetIdleStatusTimestamp() time.Time`

GetIdleStatusTimestamp returns the IdleStatusTimestamp field if non-nil, zero value otherwise.

### GetIdleStatusTimestampOk

`func (o *IdleStatusInfo) GetIdleStatusTimestampOk() (*time.Time, bool)`

GetIdleStatusTimestampOk returns a tuple with the IdleStatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleStatusTimestamp

`func (o *IdleStatusInfo) SetIdleStatusTimestamp(v time.Time)`

SetIdleStatusTimestamp sets IdleStatusTimestamp field to given value.

### HasIdleStatusTimestamp

`func (o *IdleStatusInfo) HasIdleStatusTimestamp() bool`

HasIdleStatusTimestamp returns a boolean if a field has been set.

### GetPeriodicAUTimer

`func (o *IdleStatusInfo) GetPeriodicAUTimer() int32`

GetPeriodicAUTimer returns the PeriodicAUTimer field if non-nil, zero value otherwise.

### GetPeriodicAUTimerOk

`func (o *IdleStatusInfo) GetPeriodicAUTimerOk() (*int32, bool)`

GetPeriodicAUTimerOk returns a tuple with the PeriodicAUTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicAUTimer

`func (o *IdleStatusInfo) SetPeriodicAUTimer(v int32)`

SetPeriodicAUTimer sets PeriodicAUTimer field to given value.

### HasPeriodicAUTimer

`func (o *IdleStatusInfo) HasPeriodicAUTimer() bool`

HasPeriodicAUTimer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


