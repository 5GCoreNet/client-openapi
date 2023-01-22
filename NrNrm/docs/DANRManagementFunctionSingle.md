# DANRManagementFunctionSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**DANRManagementFunctionSingleAllOfAttributes**](DANRManagementFunctionSingleAllOfAttributes.md) |  | [optional] 

## Methods

### NewDANRManagementFunctionSingle

`func NewDANRManagementFunctionSingle(id NullableString, ) *DANRManagementFunctionSingle`

NewDANRManagementFunctionSingle instantiates a new DANRManagementFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDANRManagementFunctionSingleWithDefaults

`func NewDANRManagementFunctionSingleWithDefaults() *DANRManagementFunctionSingle`

NewDANRManagementFunctionSingleWithDefaults instantiates a new DANRManagementFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DANRManagementFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DANRManagementFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DANRManagementFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *DANRManagementFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DANRManagementFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *DANRManagementFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *DANRManagementFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *DANRManagementFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *DANRManagementFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *DANRManagementFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *DANRManagementFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *DANRManagementFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *DANRManagementFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *DANRManagementFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *DANRManagementFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *DANRManagementFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *DANRManagementFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *DANRManagementFunctionSingle) GetAttributes() DANRManagementFunctionSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *DANRManagementFunctionSingle) GetAttributesOk() (*DANRManagementFunctionSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *DANRManagementFunctionSingle) SetAttributes(v DANRManagementFunctionSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *DANRManagementFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


