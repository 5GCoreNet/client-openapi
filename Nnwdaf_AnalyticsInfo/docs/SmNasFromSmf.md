# SmNasFromSmf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmNasType** | **string** |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**BackoffTimer** | **int32** | indicating a time in seconds. | 
**AppliedSmccType** | [**AppliedSmccType**](AppliedSmccType.md) |  | 

## Methods

### NewSmNasFromSmf

`func NewSmNasFromSmf(smNasType string, timeStamp time.Time, backoffTimer int32, appliedSmccType AppliedSmccType, ) *SmNasFromSmf`

NewSmNasFromSmf instantiates a new SmNasFromSmf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmNasFromSmfWithDefaults

`func NewSmNasFromSmfWithDefaults() *SmNasFromSmf`

NewSmNasFromSmfWithDefaults instantiates a new SmNasFromSmf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmNasType

`func (o *SmNasFromSmf) GetSmNasType() string`

GetSmNasType returns the SmNasType field if non-nil, zero value otherwise.

### GetSmNasTypeOk

`func (o *SmNasFromSmf) GetSmNasTypeOk() (*string, bool)`

GetSmNasTypeOk returns a tuple with the SmNasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmNasType

`func (o *SmNasFromSmf) SetSmNasType(v string)`

SetSmNasType sets SmNasType field to given value.


### GetTimeStamp

`func (o *SmNasFromSmf) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *SmNasFromSmf) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *SmNasFromSmf) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetBackoffTimer

`func (o *SmNasFromSmf) GetBackoffTimer() int32`

GetBackoffTimer returns the BackoffTimer field if non-nil, zero value otherwise.

### GetBackoffTimerOk

`func (o *SmNasFromSmf) GetBackoffTimerOk() (*int32, bool)`

GetBackoffTimerOk returns a tuple with the BackoffTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffTimer

`func (o *SmNasFromSmf) SetBackoffTimer(v int32)`

SetBackoffTimer sets BackoffTimer field to given value.


### GetAppliedSmccType

`func (o *SmNasFromSmf) GetAppliedSmccType() AppliedSmccType`

GetAppliedSmccType returns the AppliedSmccType field if non-nil, zero value otherwise.

### GetAppliedSmccTypeOk

`func (o *SmNasFromSmf) GetAppliedSmccTypeOk() (*AppliedSmccType, bool)`

GetAppliedSmccTypeOk returns a tuple with the AppliedSmccType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedSmccType

`func (o *SmNasFromSmf) SetAppliedSmccType(v AppliedSmccType)`

SetAppliedSmccType sets AppliedSmccType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


