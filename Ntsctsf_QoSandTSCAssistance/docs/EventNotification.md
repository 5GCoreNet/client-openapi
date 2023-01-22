# EventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**TscEvent**](TscEvent.md) |  | 
**FlowIds** | Pointer to **[]int32** | Identifies the IP flows that were sent during event subscription. | [optional] 
**QosMonReports** | Pointer to [**[]QosMonitoringReport**](QosMonitoringReport.md) |  | [optional] 
**UsgRep** | Pointer to [**AccumulatedUsage**](AccumulatedUsage.md) |  | [optional] 
**AppliedQosRef** | Pointer to **string** | The currently applied alternative QoS requirement referring to an alternative QoS reference or a requested alternative QoS parameter set. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION.  | [optional] 

## Methods

### NewEventNotification

`func NewEventNotification(event TscEvent, ) *EventNotification`

NewEventNotification instantiates a new EventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventNotificationWithDefaults

`func NewEventNotificationWithDefaults() *EventNotification`

NewEventNotificationWithDefaults instantiates a new EventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *EventNotification) GetEvent() TscEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventNotification) GetEventOk() (*TscEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventNotification) SetEvent(v TscEvent)`

SetEvent sets Event field to given value.


### GetFlowIds

`func (o *EventNotification) GetFlowIds() []int32`

GetFlowIds returns the FlowIds field if non-nil, zero value otherwise.

### GetFlowIdsOk

`func (o *EventNotification) GetFlowIdsOk() (*[]int32, bool)`

GetFlowIdsOk returns a tuple with the FlowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowIds

`func (o *EventNotification) SetFlowIds(v []int32)`

SetFlowIds sets FlowIds field to given value.

### HasFlowIds

`func (o *EventNotification) HasFlowIds() bool`

HasFlowIds returns a boolean if a field has been set.

### GetQosMonReports

`func (o *EventNotification) GetQosMonReports() []QosMonitoringReport`

GetQosMonReports returns the QosMonReports field if non-nil, zero value otherwise.

### GetQosMonReportsOk

`func (o *EventNotification) GetQosMonReportsOk() (*[]QosMonitoringReport, bool)`

GetQosMonReportsOk returns a tuple with the QosMonReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonReports

`func (o *EventNotification) SetQosMonReports(v []QosMonitoringReport)`

SetQosMonReports sets QosMonReports field to given value.

### HasQosMonReports

`func (o *EventNotification) HasQosMonReports() bool`

HasQosMonReports returns a boolean if a field has been set.

### GetUsgRep

`func (o *EventNotification) GetUsgRep() AccumulatedUsage`

GetUsgRep returns the UsgRep field if non-nil, zero value otherwise.

### GetUsgRepOk

`func (o *EventNotification) GetUsgRepOk() (*AccumulatedUsage, bool)`

GetUsgRepOk returns a tuple with the UsgRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsgRep

`func (o *EventNotification) SetUsgRep(v AccumulatedUsage)`

SetUsgRep sets UsgRep field to given value.

### HasUsgRep

`func (o *EventNotification) HasUsgRep() bool`

HasUsgRep returns a boolean if a field has been set.

### GetAppliedQosRef

`func (o *EventNotification) GetAppliedQosRef() string`

GetAppliedQosRef returns the AppliedQosRef field if non-nil, zero value otherwise.

### GetAppliedQosRefOk

`func (o *EventNotification) GetAppliedQosRefOk() (*string, bool)`

GetAppliedQosRefOk returns a tuple with the AppliedQosRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedQosRef

`func (o *EventNotification) SetAppliedQosRef(v string)`

SetAppliedQosRef sets AppliedQosRef field to given value.

### HasAppliedQosRef

`func (o *EventNotification) HasAppliedQosRef() bool`

HasAppliedQosRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


