# ResourcesGenericNrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**Attributes** | Pointer to [**FileSingleAllOfAttributes**](FileSingleAllOfAttributes.md) |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**MnsAgent** | Pointer to [**[]MnsAgentSingle**](MnsAgentSingle.md) |  | [optional] 
**Files** | Pointer to [**[]FilesSingle**](FilesSingle.md) |  | [optional] 
**HeartbeatControl** | Pointer to [**HeartbeatControlSingle**](HeartbeatControlSingle.md) |  | [optional] 
**MnsInfo** | Pointer to [**[]MnsInfoSingle**](MnsInfoSingle.md) |  | [optional] 
**MnsLabel** | Pointer to **string** |  | [optional] 
**MnsType** | Pointer to **string** |  | [optional] 
**MnsVersion** | Pointer to **string** |  | [optional] 
**MnsAddress** | Pointer to **string** |  | [optional] 
**MnsScope** | Pointer to **[]string** | List of the managed object instances that can be accessed using the MnS. If a complete SubNetwork can be accessed using the MnS, this attribute may contain the DN of the SubNetwork instead of the DNs of the individual managed entities within the SubNetwork. | [optional] 

## Methods

### NewResourcesGenericNrm

`func NewResourcesGenericNrm(id NullableString, ) *ResourcesGenericNrm`

NewResourcesGenericNrm instantiates a new ResourcesGenericNrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesGenericNrmWithDefaults

`func NewResourcesGenericNrmWithDefaults() *ResourcesGenericNrm`

NewResourcesGenericNrmWithDefaults instantiates a new ResourcesGenericNrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcesGenericNrm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesGenericNrm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesGenericNrm) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ResourcesGenericNrm) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResourcesGenericNrm) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ResourcesGenericNrm) GetAttributes() FileSingleAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ResourcesGenericNrm) GetAttributesOk() (*FileSingleAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ResourcesGenericNrm) SetAttributes(v FileSingleAllOfAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ResourcesGenericNrm) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ResourcesGenericNrm) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ResourcesGenericNrm) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ResourcesGenericNrm) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ResourcesGenericNrm) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetObjectClass

`func (o *ResourcesGenericNrm) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ResourcesGenericNrm) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ResourcesGenericNrm) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ResourcesGenericNrm) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ResourcesGenericNrm) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ResourcesGenericNrm) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ResourcesGenericNrm) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ResourcesGenericNrm) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetMnsAgent

`func (o *ResourcesGenericNrm) GetMnsAgent() []MnsAgentSingle`

GetMnsAgent returns the MnsAgent field if non-nil, zero value otherwise.

### GetMnsAgentOk

`func (o *ResourcesGenericNrm) GetMnsAgentOk() (*[]MnsAgentSingle, bool)`

GetMnsAgentOk returns a tuple with the MnsAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAgent

`func (o *ResourcesGenericNrm) SetMnsAgent(v []MnsAgentSingle)`

SetMnsAgent sets MnsAgent field to given value.

### HasMnsAgent

`func (o *ResourcesGenericNrm) HasMnsAgent() bool`

HasMnsAgent returns a boolean if a field has been set.

### GetFiles

`func (o *ResourcesGenericNrm) GetFiles() []FilesSingle`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ResourcesGenericNrm) GetFilesOk() (*[]FilesSingle, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ResourcesGenericNrm) SetFiles(v []FilesSingle)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ResourcesGenericNrm) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetHeartbeatControl

`func (o *ResourcesGenericNrm) GetHeartbeatControl() HeartbeatControlSingle`

GetHeartbeatControl returns the HeartbeatControl field if non-nil, zero value otherwise.

### GetHeartbeatControlOk

`func (o *ResourcesGenericNrm) GetHeartbeatControlOk() (*HeartbeatControlSingle, bool)`

GetHeartbeatControlOk returns a tuple with the HeartbeatControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatControl

`func (o *ResourcesGenericNrm) SetHeartbeatControl(v HeartbeatControlSingle)`

SetHeartbeatControl sets HeartbeatControl field to given value.

### HasHeartbeatControl

`func (o *ResourcesGenericNrm) HasHeartbeatControl() bool`

HasHeartbeatControl returns a boolean if a field has been set.

### GetMnsInfo

`func (o *ResourcesGenericNrm) GetMnsInfo() []MnsInfoSingle`

GetMnsInfo returns the MnsInfo field if non-nil, zero value otherwise.

### GetMnsInfoOk

`func (o *ResourcesGenericNrm) GetMnsInfoOk() (*[]MnsInfoSingle, bool)`

GetMnsInfoOk returns a tuple with the MnsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsInfo

`func (o *ResourcesGenericNrm) SetMnsInfo(v []MnsInfoSingle)`

SetMnsInfo sets MnsInfo field to given value.

### HasMnsInfo

`func (o *ResourcesGenericNrm) HasMnsInfo() bool`

HasMnsInfo returns a boolean if a field has been set.

### GetMnsLabel

`func (o *ResourcesGenericNrm) GetMnsLabel() string`

GetMnsLabel returns the MnsLabel field if non-nil, zero value otherwise.

### GetMnsLabelOk

`func (o *ResourcesGenericNrm) GetMnsLabelOk() (*string, bool)`

GetMnsLabelOk returns a tuple with the MnsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsLabel

`func (o *ResourcesGenericNrm) SetMnsLabel(v string)`

SetMnsLabel sets MnsLabel field to given value.

### HasMnsLabel

`func (o *ResourcesGenericNrm) HasMnsLabel() bool`

HasMnsLabel returns a boolean if a field has been set.

### GetMnsType

`func (o *ResourcesGenericNrm) GetMnsType() string`

GetMnsType returns the MnsType field if non-nil, zero value otherwise.

### GetMnsTypeOk

`func (o *ResourcesGenericNrm) GetMnsTypeOk() (*string, bool)`

GetMnsTypeOk returns a tuple with the MnsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsType

`func (o *ResourcesGenericNrm) SetMnsType(v string)`

SetMnsType sets MnsType field to given value.

### HasMnsType

`func (o *ResourcesGenericNrm) HasMnsType() bool`

HasMnsType returns a boolean if a field has been set.

### GetMnsVersion

`func (o *ResourcesGenericNrm) GetMnsVersion() string`

GetMnsVersion returns the MnsVersion field if non-nil, zero value otherwise.

### GetMnsVersionOk

`func (o *ResourcesGenericNrm) GetMnsVersionOk() (*string, bool)`

GetMnsVersionOk returns a tuple with the MnsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsVersion

`func (o *ResourcesGenericNrm) SetMnsVersion(v string)`

SetMnsVersion sets MnsVersion field to given value.

### HasMnsVersion

`func (o *ResourcesGenericNrm) HasMnsVersion() bool`

HasMnsVersion returns a boolean if a field has been set.

### GetMnsAddress

`func (o *ResourcesGenericNrm) GetMnsAddress() string`

GetMnsAddress returns the MnsAddress field if non-nil, zero value otherwise.

### GetMnsAddressOk

`func (o *ResourcesGenericNrm) GetMnsAddressOk() (*string, bool)`

GetMnsAddressOk returns a tuple with the MnsAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsAddress

`func (o *ResourcesGenericNrm) SetMnsAddress(v string)`

SetMnsAddress sets MnsAddress field to given value.

### HasMnsAddress

`func (o *ResourcesGenericNrm) HasMnsAddress() bool`

HasMnsAddress returns a boolean if a field has been set.

### GetMnsScope

`func (o *ResourcesGenericNrm) GetMnsScope() []string`

GetMnsScope returns the MnsScope field if non-nil, zero value otherwise.

### GetMnsScopeOk

`func (o *ResourcesGenericNrm) GetMnsScopeOk() (*[]string, bool)`

GetMnsScopeOk returns a tuple with the MnsScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnsScope

`func (o *ResourcesGenericNrm) SetMnsScope(v []string)`

SetMnsScope sets MnsScope field to given value.

### HasMnsScope

`func (o *ResourcesGenericNrm) HasMnsScope() bool`

HasMnsScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


