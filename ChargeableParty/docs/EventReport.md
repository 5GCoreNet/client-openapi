# EventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**Event**](Event.md) |  | 
**AccumulatedUsage** | Pointer to [**AccumulatedUsage**](AccumulatedUsage.md) |  | [optional] 
**FlowIds** | Pointer to **[]int32** | Identifies the IP flows that were sent during event subscription | [optional] 

## Methods

### NewEventReport

`func NewEventReport(event Event, ) *EventReport`

NewEventReport instantiates a new EventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventReportWithDefaults

`func NewEventReportWithDefaults() *EventReport`

NewEventReportWithDefaults instantiates a new EventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventReport) GetEvent() Event`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventReport) GetEventOk() (*Event, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventReport) SetEvent(v Event)`

SetEvent sets Event field to given value.


### GetAccumulatedUsage

`func (o *EventReport) GetAccumulatedUsage() AccumulatedUsage`

GetAccumulatedUsage returns the AccumulatedUsage field if non-nil, zero value otherwise.

### GetAccumulatedUsageOk

`func (o *EventReport) GetAccumulatedUsageOk() (*AccumulatedUsage, bool)`

GetAccumulatedUsageOk returns a tuple with the AccumulatedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedUsage

`func (o *EventReport) SetAccumulatedUsage(v AccumulatedUsage)`

SetAccumulatedUsage sets AccumulatedUsage field to given value.

### HasAccumulatedUsage

`func (o *EventReport) HasAccumulatedUsage() bool`

HasAccumulatedUsage returns a boolean if a field has been set.

### GetFlowIds

`func (o *EventReport) GetFlowIds() []int32`

GetFlowIds returns the FlowIds field if non-nil, zero value otherwise.

### GetFlowIdsOk

`func (o *EventReport) GetFlowIdsOk() (*[]int32, bool)`

GetFlowIdsOk returns a tuple with the FlowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowIds

`func (o *EventReport) SetFlowIds(v []int32)`

SetFlowIds sets FlowIds field to given value.

### HasFlowIds

`func (o *EventReport) HasFlowIds() bool`

HasFlowIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


