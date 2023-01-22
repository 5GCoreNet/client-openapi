# UserPlaneEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**UserPlaneEvent**](UserPlaneEvent.md) |  | 
**AccumulatedUsage** | Pointer to [**AccumulatedUsage**](AccumulatedUsage.md) |  | [optional] 
**FlowIds** | Pointer to **[]int32** | Identifies the affected flows that were sent during event subscription. It might be omitted when the reported event applies to all the flows sent during the subscription.  | [optional] 
**AppliedQosRef** | Pointer to **string** | The currently applied QoS reference. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION. | [optional] 
**QosMonReports** | Pointer to [**[]QosMonitoringReport**](QosMonitoringReport.md) | Contains the QoS Monitoring Reporting information | [optional] 

## Methods

### NewUserPlaneEventReport

`func NewUserPlaneEventReport(event UserPlaneEvent, ) *UserPlaneEventReport`

NewUserPlaneEventReport instantiates a new UserPlaneEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPlaneEventReportWithDefaults

`func NewUserPlaneEventReportWithDefaults() *UserPlaneEventReport`

NewUserPlaneEventReportWithDefaults instantiates a new UserPlaneEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *UserPlaneEventReport) GetEvent() UserPlaneEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *UserPlaneEventReport) GetEventOk() (*UserPlaneEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *UserPlaneEventReport) SetEvent(v UserPlaneEvent)`

SetEvent sets Event field to given value.


### GetAccumulatedUsage

`func (o *UserPlaneEventReport) GetAccumulatedUsage() AccumulatedUsage`

GetAccumulatedUsage returns the AccumulatedUsage field if non-nil, zero value otherwise.

### GetAccumulatedUsageOk

`func (o *UserPlaneEventReport) GetAccumulatedUsageOk() (*AccumulatedUsage, bool)`

GetAccumulatedUsageOk returns a tuple with the AccumulatedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccumulatedUsage

`func (o *UserPlaneEventReport) SetAccumulatedUsage(v AccumulatedUsage)`

SetAccumulatedUsage sets AccumulatedUsage field to given value.

### HasAccumulatedUsage

`func (o *UserPlaneEventReport) HasAccumulatedUsage() bool`

HasAccumulatedUsage returns a boolean if a field has been set.

### GetFlowIds

`func (o *UserPlaneEventReport) GetFlowIds() []int32`

GetFlowIds returns the FlowIds field if non-nil, zero value otherwise.

### GetFlowIdsOk

`func (o *UserPlaneEventReport) GetFlowIdsOk() (*[]int32, bool)`

GetFlowIdsOk returns a tuple with the FlowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowIds

`func (o *UserPlaneEventReport) SetFlowIds(v []int32)`

SetFlowIds sets FlowIds field to given value.

### HasFlowIds

`func (o *UserPlaneEventReport) HasFlowIds() bool`

HasFlowIds returns a boolean if a field has been set.

### GetAppliedQosRef

`func (o *UserPlaneEventReport) GetAppliedQosRef() string`

GetAppliedQosRef returns the AppliedQosRef field if non-nil, zero value otherwise.

### GetAppliedQosRefOk

`func (o *UserPlaneEventReport) GetAppliedQosRefOk() (*string, bool)`

GetAppliedQosRefOk returns a tuple with the AppliedQosRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedQosRef

`func (o *UserPlaneEventReport) SetAppliedQosRef(v string)`

SetAppliedQosRef sets AppliedQosRef field to given value.

### HasAppliedQosRef

`func (o *UserPlaneEventReport) HasAppliedQosRef() bool`

HasAppliedQosRef returns a boolean if a field has been set.

### GetQosMonReports

`func (o *UserPlaneEventReport) GetQosMonReports() []QosMonitoringReport`

GetQosMonReports returns the QosMonReports field if non-nil, zero value otherwise.

### GetQosMonReportsOk

`func (o *UserPlaneEventReport) GetQosMonReportsOk() (*[]QosMonitoringReport, bool)`

GetQosMonReportsOk returns a tuple with the QosMonReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonReports

`func (o *UserPlaneEventReport) SetQosMonReports(v []QosMonitoringReport)`

SetQosMonReports sets QosMonReports field to given value.

### HasQosMonReports

`func (o *UserPlaneEventReport) HasQosMonReports() bool`

HasQosMonReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


