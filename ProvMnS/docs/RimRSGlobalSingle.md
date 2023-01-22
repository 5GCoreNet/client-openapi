# RimRSGlobalSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**RimRSGlobalSingleAllOfAttributes**](RimRSGlobalSingleAllOfAttributes.md) |  | [optional] 
**RimRSSet** | Pointer to [**[]RimRSSetSingle**](RimRSSetSingle.md) |  | [optional] 

## Methods

### NewRimRSGlobalSingle

`func NewRimRSGlobalSingle(id NullableString, ) *RimRSGlobalSingle`

NewRimRSGlobalSingle instantiates a new RimRSGlobalSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRimRSGlobalSingleWithDefaults

`func NewRimRSGlobalSingleWithDefaults() *RimRSGlobalSingle`

NewRimRSGlobalSingleWithDefaults instantiates a new RimRSGlobalSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RimRSGlobalSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RimRSGlobalSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RimRSGlobalSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *RimRSGlobalSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RimRSGlobalSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *RimRSGlobalSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *RimRSGlobalSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *RimRSGlobalSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *RimRSGlobalSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *RimRSGlobalSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *RimRSGlobalSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *RimRSGlobalSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *RimRSGlobalSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *RimRSGlobalSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *RimRSGlobalSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *RimRSGlobalSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *RimRSGlobalSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *RimRSGlobalSingle) GetAttributes() RimRSGlobalSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RimRSGlobalSingle) GetAttributesOk() (*RimRSGlobalSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RimRSGlobalSingle) SetAttributes(v RimRSGlobalSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RimRSGlobalSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRimRSSet

`func (o *RimRSGlobalSingle) GetRimRSSet() []RimRSSetSingle`

GetRimRSSet returns the RimRSSet field if non-nil, zero value otherwise.

### GetRimRSSetOk

`func (o *RimRSGlobalSingle) GetRimRSSetOk() (*[]RimRSSetSingle, bool)`

GetRimRSSetOk returns a tuple with the RimRSSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRimRSSet

`func (o *RimRSGlobalSingle) SetRimRSSet(v []RimRSSetSingle)`

SetRimRSSet sets RimRSSet field to given value.

### HasRimRSSet

`func (o *RimRSGlobalSingle) HasRimRSSet() bool`

HasRimRSSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


