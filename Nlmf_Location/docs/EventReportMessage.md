# EventReportMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventClass** | [**EventClass**](EventClass.md) |  | 
**EventContent** | [**RefToBinaryData**](RefToBinaryData.md) |  | 

## Methods

### NewEventReportMessage

`func NewEventReportMessage(eventClass EventClass, eventContent RefToBinaryData, ) *EventReportMessage`

NewEventReportMessage instantiates a new EventReportMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventReportMessageWithDefaults

`func NewEventReportMessageWithDefaults() *EventReportMessage`

NewEventReportMessageWithDefaults instantiates a new EventReportMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventClass

`func (o *EventReportMessage) GetEventClass() EventClass`

GetEventClass returns the EventClass field if non-nil, zero value otherwise.

### GetEventClassOk

`func (o *EventReportMessage) GetEventClassOk() (*EventClass, bool)`

GetEventClassOk returns a tuple with the EventClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClass

`func (o *EventReportMessage) SetEventClass(v EventClass)`

SetEventClass sets EventClass field to given value.


### GetEventContent

`func (o *EventReportMessage) GetEventContent() RefToBinaryData`

GetEventContent returns the EventContent field if non-nil, zero value otherwise.

### GetEventContentOk

`func (o *EventReportMessage) GetEventContentOk() (*RefToBinaryData, bool)`

GetEventContentOk returns a tuple with the EventContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventContent

`func (o *EventReportMessage) SetEventContent(v RefToBinaryData)`

SetEventContent sets EventContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


