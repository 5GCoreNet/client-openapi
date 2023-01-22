# ParameterProcessingInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A JSON pointer value that references an attribute within the notification object to which the processing instruction is applied.  | 
**Values** | **[]interface{}** | A list of values for the attribute identified by the name attribute. | 
**SumAttrs** | [**[]SummarizationAttribute**](SummarizationAttribute.md) | Attributes requested to be used in the summarized reports. | 
**AggrLevel** | Pointer to [**AggregationLevel**](AggregationLevel.md) |  | [optional] 
**Supis** | Pointer to **[]string** | Indicates the UEs for which processed reports are requested. | [optional] 
**Areas** | Pointer to [**[]NetworkAreaInfo**](NetworkAreaInfo.md) | Indicates the Areas of Interest for which processed reports are requested. | [optional] 

## Methods

### NewParameterProcessingInstruction

`func NewParameterProcessingInstruction(name string, values []interface{}, sumAttrs []SummarizationAttribute, ) *ParameterProcessingInstruction`

NewParameterProcessingInstruction instantiates a new ParameterProcessingInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterProcessingInstructionWithDefaults

`func NewParameterProcessingInstructionWithDefaults() *ParameterProcessingInstruction`

NewParameterProcessingInstructionWithDefaults instantiates a new ParameterProcessingInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParameterProcessingInstruction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParameterProcessingInstruction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParameterProcessingInstruction) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *ParameterProcessingInstruction) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ParameterProcessingInstruction) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ParameterProcessingInstruction) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetSumAttrs

`func (o *ParameterProcessingInstruction) GetSumAttrs() []SummarizationAttribute`

GetSumAttrs returns the SumAttrs field if non-nil, zero value otherwise.

### GetSumAttrsOk

`func (o *ParameterProcessingInstruction) GetSumAttrsOk() (*[]SummarizationAttribute, bool)`

GetSumAttrsOk returns a tuple with the SumAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSumAttrs

`func (o *ParameterProcessingInstruction) SetSumAttrs(v []SummarizationAttribute)`

SetSumAttrs sets SumAttrs field to given value.


### GetAggrLevel

`func (o *ParameterProcessingInstruction) GetAggrLevel() AggregationLevel`

GetAggrLevel returns the AggrLevel field if non-nil, zero value otherwise.

### GetAggrLevelOk

`func (o *ParameterProcessingInstruction) GetAggrLevelOk() (*AggregationLevel, bool)`

GetAggrLevelOk returns a tuple with the AggrLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrLevel

`func (o *ParameterProcessingInstruction) SetAggrLevel(v AggregationLevel)`

SetAggrLevel sets AggrLevel field to given value.

### HasAggrLevel

`func (o *ParameterProcessingInstruction) HasAggrLevel() bool`

HasAggrLevel returns a boolean if a field has been set.

### GetSupis

`func (o *ParameterProcessingInstruction) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *ParameterProcessingInstruction) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *ParameterProcessingInstruction) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *ParameterProcessingInstruction) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetAreas

`func (o *ParameterProcessingInstruction) GetAreas() []NetworkAreaInfo`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *ParameterProcessingInstruction) GetAreasOk() (*[]NetworkAreaInfo, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *ParameterProcessingInstruction) SetAreas(v []NetworkAreaInfo)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *ParameterProcessingInstruction) HasAreas() bool`

HasAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


