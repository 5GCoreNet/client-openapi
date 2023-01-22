# ScheduleOfDayType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to [**DayOfWeekType**](DayOfWeekType.md) |  | [optional] 
**IntervalsOfDay** | Pointer to [**[]TimeIntervalType**](TimeIntervalType.md) |  | [optional] 

## Methods

### NewScheduleOfDayType

`func NewScheduleOfDayType() *ScheduleOfDayType`

NewScheduleOfDayType instantiates a new ScheduleOfDayType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleOfDayTypeWithDefaults

`func NewScheduleOfDayTypeWithDefaults() *ScheduleOfDayType`

NewScheduleOfDayTypeWithDefaults instantiates a new ScheduleOfDayType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *ScheduleOfDayType) GetDayOfWeek() DayOfWeekType`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *ScheduleOfDayType) GetDayOfWeekOk() (*DayOfWeekType, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *ScheduleOfDayType) SetDayOfWeek(v DayOfWeekType)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *ScheduleOfDayType) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetIntervalsOfDay

`func (o *ScheduleOfDayType) GetIntervalsOfDay() []TimeIntervalType`

GetIntervalsOfDay returns the IntervalsOfDay field if non-nil, zero value otherwise.

### GetIntervalsOfDayOk

`func (o *ScheduleOfDayType) GetIntervalsOfDayOk() (*[]TimeIntervalType, bool)`

GetIntervalsOfDayOk returns a tuple with the IntervalsOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalsOfDay

`func (o *ScheduleOfDayType) SetIntervalsOfDay(v []TimeIntervalType)`

SetIntervalsOfDay sets IntervalsOfDay field to given value.

### HasIntervalsOfDay

`func (o *ScheduleOfDayType) HasIntervalsOfDay() bool`

HasIntervalsOfDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


