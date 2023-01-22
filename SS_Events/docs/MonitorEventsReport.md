# MonitorEventsReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgtUe** | [**ValTargetUe**](ValTargetUe.md) |  | 
**Evnts** | [**[]MonitorEvents**](MonitorEvents.md) | List of monitoring and analytics events related to VAL UE. | 

## Methods

### NewMonitorEventsReport

`func NewMonitorEventsReport(tgtUe ValTargetUe, evnts []MonitorEvents, ) *MonitorEventsReport`

NewMonitorEventsReport instantiates a new MonitorEventsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorEventsReportWithDefaults

`func NewMonitorEventsReportWithDefaults() *MonitorEventsReport`

NewMonitorEventsReportWithDefaults instantiates a new MonitorEventsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgtUe

`func (o *MonitorEventsReport) GetTgtUe() ValTargetUe`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *MonitorEventsReport) GetTgtUeOk() (*ValTargetUe, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *MonitorEventsReport) SetTgtUe(v ValTargetUe)`

SetTgtUe sets TgtUe field to given value.


### GetEvnts

`func (o *MonitorEventsReport) GetEvnts() []MonitorEvents`

GetEvnts returns the Evnts field if non-nil, zero value otherwise.

### GetEvntsOk

`func (o *MonitorEventsReport) GetEvntsOk() (*[]MonitorEvents, bool)`

GetEvntsOk returns a tuple with the Evnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvnts

`func (o *MonitorEventsReport) SetEvnts(v []MonitorEvents)`

SetEvnts sets Evnts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


