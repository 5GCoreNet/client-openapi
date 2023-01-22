# NetworkSliceSubnetProviderCapabilitiesSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes**](NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes.md) |  | [optional] 

## Methods

### NewNetworkSliceSubnetProviderCapabilitiesSingle

`func NewNetworkSliceSubnetProviderCapabilitiesSingle(id NullableString, ) *NetworkSliceSubnetProviderCapabilitiesSingle`

NewNetworkSliceSubnetProviderCapabilitiesSingle instantiates a new NetworkSliceSubnetProviderCapabilitiesSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSliceSubnetProviderCapabilitiesSingleWithDefaults

`func NewNetworkSliceSubnetProviderCapabilitiesSingleWithDefaults() *NetworkSliceSubnetProviderCapabilitiesSingle`

NewNetworkSliceSubnetProviderCapabilitiesSingleWithDefaults instantiates a new NetworkSliceSubnetProviderCapabilitiesSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetAttributes() NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) GetAttributesOk() (*NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) SetAttributes(v NetworkSliceSubnetProviderCapabilitiesSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NetworkSliceSubnetProviderCapabilitiesSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


