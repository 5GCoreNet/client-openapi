# ScheduleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduleOption** | Pointer to [**ScheduleOptionType**](ScheduleOptionType.md) |  | [optional] 
**DailySchedule** | Pointer to [**[]TimeIntervalType**](TimeIntervalType.md) |  | [optional] 
**WeeklySchedule** | Pointer to [**[]ScheduleOfDayType**](ScheduleOfDayType.md) |  | [optional] 

## Methods

### NewScheduleType

`func NewScheduleType() *ScheduleType`

NewScheduleType instantiates a new ScheduleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTypeWithDefaults

`func NewScheduleTypeWithDefaults() *ScheduleType`

NewScheduleTypeWithDefaults instantiates a new ScheduleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduleOption

`func (o *ScheduleType) GetScheduleOption() ScheduleOptionType`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *ScheduleType) GetScheduleOptionOk() (*ScheduleOptionType, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *ScheduleType) SetScheduleOption(v ScheduleOptionType)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *ScheduleType) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetDailySchedule

`func (o *ScheduleType) GetDailySchedule() []TimeIntervalType`

GetDailySchedule returns the DailySchedule field if non-nil, zero value otherwise.

### GetDailyScheduleOk

`func (o *ScheduleType) GetDailyScheduleOk() (*[]TimeIntervalType, bool)`

GetDailyScheduleOk returns a tuple with the DailySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailySchedule

`func (o *ScheduleType) SetDailySchedule(v []TimeIntervalType)`

SetDailySchedule sets DailySchedule field to given value.

### HasDailySchedule

`func (o *ScheduleType) HasDailySchedule() bool`

HasDailySchedule returns a boolean if a field has been set.

### GetWeeklySchedule

`func (o *ScheduleType) GetWeeklySchedule() []ScheduleOfDayType`

GetWeeklySchedule returns the WeeklySchedule field if non-nil, zero value otherwise.

### GetWeeklyScheduleOk

`func (o *ScheduleType) GetWeeklyScheduleOk() (*[]ScheduleOfDayType, bool)`

GetWeeklyScheduleOk returns a tuple with the WeeklySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeeklySchedule

`func (o *ScheduleType) SetWeeklySchedule(v []ScheduleOfDayType)`

SetWeeklySchedule sets WeeklySchedule field to given value.

### HasWeeklySchedule

`func (o *ScheduleType) HasWeeklySchedule() bool`

HasWeeklySchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


