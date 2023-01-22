# ProcessingInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | [**DccfEvent**](DccfEvent.md) |  | 
**ProcInterval** | **int32** | indicating a time in seconds. | 
**ParamProcInstructs** | Pointer to [**[]ParameterProcessingInstruction**](ParameterProcessingInstruction.md) | List of event parameter names, and for each event parameter name, respective event parameter values and sets of the attributes to be used in the summarized reports.  | [optional] 

## Methods

### NewProcessingInstruction

`func NewProcessingInstruction(eventId DccfEvent, procInterval int32, ) *ProcessingInstruction`

NewProcessingInstruction instantiates a new ProcessingInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingInstructionWithDefaults

`func NewProcessingInstructionWithDefaults() *ProcessingInstruction`

NewProcessingInstructionWithDefaults instantiates a new ProcessingInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ProcessingInstruction) GetEventId() DccfEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ProcessingInstruction) GetEventIdOk() (*DccfEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ProcessingInstruction) SetEventId(v DccfEvent)`

SetEventId sets EventId field to given value.


### GetProcInterval

`func (o *ProcessingInstruction) GetProcInterval() int32`

GetProcInterval returns the ProcInterval field if non-nil, zero value otherwise.

### GetProcIntervalOk

`func (o *ProcessingInstruction) GetProcIntervalOk() (*int32, bool)`

GetProcIntervalOk returns a tuple with the ProcInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcInterval

`func (o *ProcessingInstruction) SetProcInterval(v int32)`

SetProcInterval sets ProcInterval field to given value.


### GetParamProcInstructs

`func (o *ProcessingInstruction) GetParamProcInstructs() []ParameterProcessingInstruction`

GetParamProcInstructs returns the ParamProcInstructs field if non-nil, zero value otherwise.

### GetParamProcInstructsOk

`func (o *ProcessingInstruction) GetParamProcInstructsOk() (*[]ParameterProcessingInstruction, bool)`

GetParamProcInstructsOk returns a tuple with the ParamProcInstructs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamProcInstructs

`func (o *ProcessingInstruction) SetParamProcInstructs(v []ParameterProcessingInstruction)`

SetParamProcInstructs sets ParamProcInstructs field to given value.

### HasParamProcInstructs

`func (o *ProcessingInstruction) HasParamProcInstructs() bool`

HasParamProcInstructs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


