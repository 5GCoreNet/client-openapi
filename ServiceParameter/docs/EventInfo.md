# EventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCause** | Pointer to [**Failure**](Failure.md) |  | [optional] 

## Methods

### NewEventInfo

`func NewEventInfo() *EventInfo`

NewEventInfo instantiates a new EventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventInfoWithDefaults

`func NewEventInfoWithDefaults() *EventInfo`

NewEventInfoWithDefaults instantiates a new EventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCause

`func (o *EventInfo) GetFailureCause() Failure`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *EventInfo) GetFailureCauseOk() (*Failure, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *EventInfo) SetFailureCause(v Failure)`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *EventInfo) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


