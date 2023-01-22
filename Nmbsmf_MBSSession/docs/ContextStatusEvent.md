# ContextStatusEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**ContextStatusEventType**](ContextStatusEventType.md) |  | 
**ImmediateReportInd** | Pointer to **bool** |  | [optional] [default to false]
**ReportingMode** | Pointer to [**ReportingMode**](ReportingMode.md) |  | [optional] 

## Methods

### NewContextStatusEvent

`func NewContextStatusEvent(eventType ContextStatusEventType, ) *ContextStatusEvent`

NewContextStatusEvent instantiates a new ContextStatusEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextStatusEventWithDefaults

`func NewContextStatusEventWithDefaults() *ContextStatusEvent`

NewContextStatusEventWithDefaults instantiates a new ContextStatusEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ContextStatusEvent) GetEventType() ContextStatusEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ContextStatusEvent) GetEventTypeOk() (*ContextStatusEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ContextStatusEvent) SetEventType(v ContextStatusEventType)`

SetEventType sets EventType field to given value.


### GetImmediateReportInd

`func (o *ContextStatusEvent) GetImmediateReportInd() bool`

GetImmediateReportInd returns the ImmediateReportInd field if non-nil, zero value otherwise.

### GetImmediateReportIndOk

`func (o *ContextStatusEvent) GetImmediateReportIndOk() (*bool, bool)`

GetImmediateReportIndOk returns a tuple with the ImmediateReportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateReportInd

`func (o *ContextStatusEvent) SetImmediateReportInd(v bool)`

SetImmediateReportInd sets ImmediateReportInd field to given value.

### HasImmediateReportInd

`func (o *ContextStatusEvent) HasImmediateReportInd() bool`

HasImmediateReportInd returns a boolean if a field has been set.

### GetReportingMode

`func (o *ContextStatusEvent) GetReportingMode() ReportingMode`

GetReportingMode returns the ReportingMode field if non-nil, zero value otherwise.

### GetReportingModeOk

`func (o *ContextStatusEvent) GetReportingModeOk() (*ReportingMode, bool)`

GetReportingModeOk returns a tuple with the ReportingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingMode

`func (o *ContextStatusEvent) SetReportingMode(v ReportingMode)`

SetReportingMode sets ReportingMode field to given value.

### HasReportingMode

`func (o *ContextStatusEvent) HasReportingMode() bool`

HasReportingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


