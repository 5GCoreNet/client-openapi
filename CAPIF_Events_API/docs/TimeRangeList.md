# TimeRangeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**StopTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewTimeRangeList

`func NewTimeRangeList() *TimeRangeList`

NewTimeRangeList instantiates a new TimeRangeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeRangeListWithDefaults

`func NewTimeRangeListWithDefaults() *TimeRangeList`

NewTimeRangeListWithDefaults instantiates a new TimeRangeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *TimeRangeList) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TimeRangeList) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TimeRangeList) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TimeRangeList) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopTime

`func (o *TimeRangeList) GetStopTime() time.Time`

GetStopTime returns the StopTime field if non-nil, zero value otherwise.

### GetStopTimeOk

`func (o *TimeRangeList) GetStopTimeOk() (*time.Time, bool)`

GetStopTimeOk returns a tuple with the StopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTime

`func (o *TimeRangeList) SetStopTime(v time.Time)`

SetStopTime sets StopTime field to given value.

### HasStopTime

`func (o *TimeRangeList) HasStopTime() bool`

HasStopTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


