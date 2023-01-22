# EventReportingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventReportCounter** | Pointer to **int32** | Number of event reports received from the target UE. | [optional] 
**EventReportDuration** | Pointer to **int32** | Duration of event reporting. | [optional] 

## Methods

### NewEventReportingStatus

`func NewEventReportingStatus() *EventReportingStatus`

NewEventReportingStatus instantiates a new EventReportingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventReportingStatusWithDefaults

`func NewEventReportingStatusWithDefaults() *EventReportingStatus`

NewEventReportingStatusWithDefaults instantiates a new EventReportingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventReportCounter

`func (o *EventReportingStatus) GetEventReportCounter() int32`

GetEventReportCounter returns the EventReportCounter field if non-nil, zero value otherwise.

### GetEventReportCounterOk

`func (o *EventReportingStatus) GetEventReportCounterOk() (*int32, bool)`

GetEventReportCounterOk returns a tuple with the EventReportCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReportCounter

`func (o *EventReportingStatus) SetEventReportCounter(v int32)`

SetEventReportCounter sets EventReportCounter field to given value.

### HasEventReportCounter

`func (o *EventReportingStatus) HasEventReportCounter() bool`

HasEventReportCounter returns a boolean if a field has been set.

### GetEventReportDuration

`func (o *EventReportingStatus) GetEventReportDuration() int32`

GetEventReportDuration returns the EventReportDuration field if non-nil, zero value otherwise.

### GetEventReportDurationOk

`func (o *EventReportingStatus) GetEventReportDurationOk() (*int32, bool)`

GetEventReportDurationOk returns a tuple with the EventReportDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReportDuration

`func (o *EventReportingStatus) SetEventReportDuration(v int32)`

SetEventReportDuration sets EventReportDuration field to given value.

### HasEventReportDuration

`func (o *EventReportingStatus) HasEventReportDuration() bool`

HasEventReportDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


