# NotifSummaryReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | [**DccfEvent**](DccfEvent.md) |  | 
**ProcInterval** | **int32** | indicating a time in seconds. | 
**EventReports** | [**[]EventParamReport**](EventParamReport.md) | List of event parameter reports. | 

## Methods

### NewNotifSummaryReport

`func NewNotifSummaryReport(eventId DccfEvent, procInterval int32, eventReports []EventParamReport, ) *NotifSummaryReport`

NewNotifSummaryReport instantiates a new NotifSummaryReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifSummaryReportWithDefaults

`func NewNotifSummaryReportWithDefaults() *NotifSummaryReport`

NewNotifSummaryReportWithDefaults instantiates a new NotifSummaryReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *NotifSummaryReport) GetEventId() DccfEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *NotifSummaryReport) GetEventIdOk() (*DccfEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *NotifSummaryReport) SetEventId(v DccfEvent)`

SetEventId sets EventId field to given value.


### GetProcInterval

`func (o *NotifSummaryReport) GetProcInterval() int32`

GetProcInterval returns the ProcInterval field if non-nil, zero value otherwise.

### GetProcIntervalOk

`func (o *NotifSummaryReport) GetProcIntervalOk() (*int32, bool)`

GetProcIntervalOk returns a tuple with the ProcInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcInterval

`func (o *NotifSummaryReport) SetProcInterval(v int32)`

SetProcInterval sets ProcInterval field to given value.


### GetEventReports

`func (o *NotifSummaryReport) GetEventReports() []EventParamReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *NotifSummaryReport) GetEventReportsOk() (*[]EventParamReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *NotifSummaryReport) SetEventReports(v []EventParamReport)`

SetEventReports sets EventReports field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


