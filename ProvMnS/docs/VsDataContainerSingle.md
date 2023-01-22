# VsDataContainerSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**VsDataContainerSingleAttributes**](VsDataContainerSingleAttributes.md) |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 

## Methods

### NewVsDataContainerSingle

`func NewVsDataContainerSingle() *VsDataContainerSingle`

NewVsDataContainerSingle instantiates a new VsDataContainerSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVsDataContainerSingleWithDefaults

`func NewVsDataContainerSingleWithDefaults() *VsDataContainerSingle`

NewVsDataContainerSingleWithDefaults instantiates a new VsDataContainerSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VsDataContainerSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VsDataContainerSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VsDataContainerSingle) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VsDataContainerSingle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *VsDataContainerSingle) GetAttributes() VsDataContainerSingleAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VsDataContainerSingle) GetAttributesOk() (*VsDataContainerSingleAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VsDataContainerSingle) SetAttributes(v VsDataContainerSingleAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *VsDataContainerSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *VsDataContainerSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *VsDataContainerSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *VsDataContainerSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *VsDataContainerSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


