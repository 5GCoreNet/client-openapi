# AssuranceClosedControlLoopSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**AssuranceClosedControlLoopSingleAllOfAttributes**](AssuranceClosedControlLoopSingleAllOfAttributes.md) |  | [optional] 
**AssuranceGoal** | Pointer to [**[]AssuranceGoalSingle**](AssuranceGoalSingle.md) |  | [optional] 

## Methods

### NewAssuranceClosedControlLoopSingle

`func NewAssuranceClosedControlLoopSingle(id NullableString, ) *AssuranceClosedControlLoopSingle`

NewAssuranceClosedControlLoopSingle instantiates a new AssuranceClosedControlLoopSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceClosedControlLoopSingleWithDefaults

`func NewAssuranceClosedControlLoopSingleWithDefaults() *AssuranceClosedControlLoopSingle`

NewAssuranceClosedControlLoopSingleWithDefaults instantiates a new AssuranceClosedControlLoopSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssuranceClosedControlLoopSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssuranceClosedControlLoopSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssuranceClosedControlLoopSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AssuranceClosedControlLoopSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AssuranceClosedControlLoopSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *AssuranceClosedControlLoopSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AssuranceClosedControlLoopSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AssuranceClosedControlLoopSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *AssuranceClosedControlLoopSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *AssuranceClosedControlLoopSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *AssuranceClosedControlLoopSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *AssuranceClosedControlLoopSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *AssuranceClosedControlLoopSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *AssuranceClosedControlLoopSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *AssuranceClosedControlLoopSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *AssuranceClosedControlLoopSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *AssuranceClosedControlLoopSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *AssuranceClosedControlLoopSingle) GetAttributes() AssuranceClosedControlLoopSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AssuranceClosedControlLoopSingle) GetAttributesOk() (*AssuranceClosedControlLoopSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AssuranceClosedControlLoopSingle) SetAttributes(v AssuranceClosedControlLoopSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AssuranceClosedControlLoopSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetAssuranceGoal

`func (o *AssuranceClosedControlLoopSingle) GetAssuranceGoal() []AssuranceGoalSingle`

GetAssuranceGoal returns the AssuranceGoal field if non-nil, zero value otherwise.

### GetAssuranceGoalOk

`func (o *AssuranceClosedControlLoopSingle) GetAssuranceGoalOk() (*[]AssuranceGoalSingle, bool)`

GetAssuranceGoalOk returns a tuple with the AssuranceGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssuranceGoal

`func (o *AssuranceClosedControlLoopSingle) SetAssuranceGoal(v []AssuranceGoalSingle)`

SetAssuranceGoal sets AssuranceGoal field to given value.

### HasAssuranceGoal

`func (o *AssuranceClosedControlLoopSingle) HasAssuranceGoal() bool`

HasAssuranceGoal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


