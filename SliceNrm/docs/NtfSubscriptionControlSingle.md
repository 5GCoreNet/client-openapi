# NtfSubscriptionControlSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**NtfSubscriptionControlSingleAllOfAttributes**](NtfSubscriptionControlSingleAllOfAttributes.md) |  | [optional] 
**HeartbeatControl** | Pointer to [**HeartbeatControlSingle**](HeartbeatControlSingle.md) |  | [optional] 

## Methods

### NewNtfSubscriptionControlSingle

`func NewNtfSubscriptionControlSingle(id NullableString, ) *NtfSubscriptionControlSingle`

NewNtfSubscriptionControlSingle instantiates a new NtfSubscriptionControlSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNtfSubscriptionControlSingleWithDefaults

`func NewNtfSubscriptionControlSingleWithDefaults() *NtfSubscriptionControlSingle`

NewNtfSubscriptionControlSingleWithDefaults instantiates a new NtfSubscriptionControlSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NtfSubscriptionControlSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NtfSubscriptionControlSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NtfSubscriptionControlSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *NtfSubscriptionControlSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NtfSubscriptionControlSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *NtfSubscriptionControlSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *NtfSubscriptionControlSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *NtfSubscriptionControlSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *NtfSubscriptionControlSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *NtfSubscriptionControlSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *NtfSubscriptionControlSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *NtfSubscriptionControlSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *NtfSubscriptionControlSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *NtfSubscriptionControlSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *NtfSubscriptionControlSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *NtfSubscriptionControlSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *NtfSubscriptionControlSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *NtfSubscriptionControlSingle) GetAttributes() NtfSubscriptionControlSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NtfSubscriptionControlSingle) GetAttributesOk() (*NtfSubscriptionControlSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NtfSubscriptionControlSingle) SetAttributes(v NtfSubscriptionControlSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NtfSubscriptionControlSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetHeartbeatControl

`func (o *NtfSubscriptionControlSingle) GetHeartbeatControl() HeartbeatControlSingle`

GetHeartbeatControl returns the HeartbeatControl field if non-nil, zero value otherwise.

### GetHeartbeatControlOk

`func (o *NtfSubscriptionControlSingle) GetHeartbeatControlOk() (*HeartbeatControlSingle, bool)`

GetHeartbeatControlOk returns a tuple with the HeartbeatControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatControl

`func (o *NtfSubscriptionControlSingle) SetHeartbeatControl(v HeartbeatControlSingle)`

SetHeartbeatControl sets HeartbeatControl field to given value.

### HasHeartbeatControl

`func (o *NtfSubscriptionControlSingle) HasHeartbeatControl() bool`

HasHeartbeatControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


