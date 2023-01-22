# CommonBeamformingFunctionSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
**Beam** | Pointer to [**[]BeamSingle**](BeamSingle.md) |  | [optional] 

## Methods

### NewCommonBeamformingFunctionSingle

`func NewCommonBeamformingFunctionSingle(id NullableString, ) *CommonBeamformingFunctionSingle`

NewCommonBeamformingFunctionSingle instantiates a new CommonBeamformingFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonBeamformingFunctionSingleWithDefaults

`func NewCommonBeamformingFunctionSingleWithDefaults() *CommonBeamformingFunctionSingle`

NewCommonBeamformingFunctionSingleWithDefaults instantiates a new CommonBeamformingFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommonBeamformingFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommonBeamformingFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommonBeamformingFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *CommonBeamformingFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CommonBeamformingFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *CommonBeamformingFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *CommonBeamformingFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *CommonBeamformingFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *CommonBeamformingFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *CommonBeamformingFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *CommonBeamformingFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *CommonBeamformingFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *CommonBeamformingFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *CommonBeamformingFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *CommonBeamformingFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *CommonBeamformingFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *CommonBeamformingFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *CommonBeamformingFunctionSingle) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CommonBeamformingFunctionSingle) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CommonBeamformingFunctionSingle) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CommonBeamformingFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBeam

`func (o *CommonBeamformingFunctionSingle) GetBeam() []BeamSingle`

GetBeam returns the Beam field if non-nil, zero value otherwise.

### GetBeamOk

`func (o *CommonBeamformingFunctionSingle) GetBeamOk() (*[]BeamSingle, bool)`

GetBeamOk returns a tuple with the Beam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeam

`func (o *CommonBeamformingFunctionSingle) SetBeam(v []BeamSingle)`

SetBeam sets Beam field to given value.

### HasBeam

`func (o *CommonBeamformingFunctionSingle) HasBeam() bool`

HasBeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


