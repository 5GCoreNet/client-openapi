# CCOPilotPollutionParametersSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**CCOParametersAttrAllOfAttributes**](CCOParametersAttrAllOfAttributes.md) |  | [optional] 

## Methods

### NewCCOPilotPollutionParametersSingle

`func NewCCOPilotPollutionParametersSingle(id NullableString, ) *CCOPilotPollutionParametersSingle`

NewCCOPilotPollutionParametersSingle instantiates a new CCOPilotPollutionParametersSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCOPilotPollutionParametersSingleWithDefaults

`func NewCCOPilotPollutionParametersSingleWithDefaults() *CCOPilotPollutionParametersSingle`

NewCCOPilotPollutionParametersSingleWithDefaults instantiates a new CCOPilotPollutionParametersSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CCOPilotPollutionParametersSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CCOPilotPollutionParametersSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CCOPilotPollutionParametersSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CCOPilotPollutionParametersSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CCOPilotPollutionParametersSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *CCOPilotPollutionParametersSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *CCOPilotPollutionParametersSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *CCOPilotPollutionParametersSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *CCOPilotPollutionParametersSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *CCOPilotPollutionParametersSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *CCOPilotPollutionParametersSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *CCOPilotPollutionParametersSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *CCOPilotPollutionParametersSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *CCOPilotPollutionParametersSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *CCOPilotPollutionParametersSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *CCOPilotPollutionParametersSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *CCOPilotPollutionParametersSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *CCOPilotPollutionParametersSingle) GetAttributes() CCOParametersAttrAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CCOPilotPollutionParametersSingle) GetAttributesOk() (*CCOParametersAttrAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CCOPilotPollutionParametersSingle) SetAttributes(v CCOParametersAttrAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CCOPilotPollutionParametersSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


