# FiveQICharacteristicsSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**FiveQIValue** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**PriorityLevel** | Pointer to **int32** |  | [optional] 
**PacketDelayBudget** | Pointer to **int32** |  | [optional] 
**PacketErrorRate** | Pointer to [**PacketErrorRate**](PacketErrorRate.md) |  | [optional] 
**AveragingWindow** | Pointer to **int32** |  | [optional] 
**MaximumDataBurstVolume** | Pointer to **int32** |  | [optional] 

## Methods

### NewFiveQICharacteristicsSingle

`func NewFiveQICharacteristicsSingle(id NullableString, ) *FiveQICharacteristicsSingle`

NewFiveQICharacteristicsSingle instantiates a new FiveQICharacteristicsSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiveQICharacteristicsSingleWithDefaults

`func NewFiveQICharacteristicsSingleWithDefaults() *FiveQICharacteristicsSingle`

NewFiveQICharacteristicsSingleWithDefaults instantiates a new FiveQICharacteristicsSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FiveQICharacteristicsSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiveQICharacteristicsSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiveQICharacteristicsSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *FiveQICharacteristicsSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *FiveQICharacteristicsSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *FiveQICharacteristicsSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *FiveQICharacteristicsSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *FiveQICharacteristicsSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *FiveQICharacteristicsSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *FiveQICharacteristicsSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *FiveQICharacteristicsSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *FiveQICharacteristicsSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *FiveQICharacteristicsSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *FiveQICharacteristicsSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *FiveQICharacteristicsSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *FiveQICharacteristicsSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *FiveQICharacteristicsSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetFiveQIValue

`func (o *FiveQICharacteristicsSingle) GetFiveQIValue() int32`

GetFiveQIValue returns the FiveQIValue field if non-nil, zero value otherwise.

### GetFiveQIValueOk

`func (o *FiveQICharacteristicsSingle) GetFiveQIValueOk() (*int32, bool)`

GetFiveQIValueOk returns a tuple with the FiveQIValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveQIValue

`func (o *FiveQICharacteristicsSingle) SetFiveQIValue(v int32)`

SetFiveQIValue sets FiveQIValue field to given value.

### HasFiveQIValue

`func (o *FiveQICharacteristicsSingle) HasFiveQIValue() bool`

HasFiveQIValue returns a boolean if a field has been set.

### GetResourceType

`func (o *FiveQICharacteristicsSingle) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *FiveQICharacteristicsSingle) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *FiveQICharacteristicsSingle) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *FiveQICharacteristicsSingle) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetPriorityLevel

`func (o *FiveQICharacteristicsSingle) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *FiveQICharacteristicsSingle) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *FiveQICharacteristicsSingle) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.

### HasPriorityLevel

`func (o *FiveQICharacteristicsSingle) HasPriorityLevel() bool`

HasPriorityLevel returns a boolean if a field has been set.

### GetPacketDelayBudget

`func (o *FiveQICharacteristicsSingle) GetPacketDelayBudget() int32`

GetPacketDelayBudget returns the PacketDelayBudget field if non-nil, zero value otherwise.

### GetPacketDelayBudgetOk

`func (o *FiveQICharacteristicsSingle) GetPacketDelayBudgetOk() (*int32, bool)`

GetPacketDelayBudgetOk returns a tuple with the PacketDelayBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketDelayBudget

`func (o *FiveQICharacteristicsSingle) SetPacketDelayBudget(v int32)`

SetPacketDelayBudget sets PacketDelayBudget field to given value.

### HasPacketDelayBudget

`func (o *FiveQICharacteristicsSingle) HasPacketDelayBudget() bool`

HasPacketDelayBudget returns a boolean if a field has been set.

### GetPacketErrorRate

`func (o *FiveQICharacteristicsSingle) GetPacketErrorRate() PacketErrorRate`

GetPacketErrorRate returns the PacketErrorRate field if non-nil, zero value otherwise.

### GetPacketErrorRateOk

`func (o *FiveQICharacteristicsSingle) GetPacketErrorRateOk() (*PacketErrorRate, bool)`

GetPacketErrorRateOk returns a tuple with the PacketErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketErrorRate

`func (o *FiveQICharacteristicsSingle) SetPacketErrorRate(v PacketErrorRate)`

SetPacketErrorRate sets PacketErrorRate field to given value.

### HasPacketErrorRate

`func (o *FiveQICharacteristicsSingle) HasPacketErrorRate() bool`

HasPacketErrorRate returns a boolean if a field has been set.

### GetAveragingWindow

`func (o *FiveQICharacteristicsSingle) GetAveragingWindow() int32`

GetAveragingWindow returns the AveragingWindow field if non-nil, zero value otherwise.

### GetAveragingWindowOk

`func (o *FiveQICharacteristicsSingle) GetAveragingWindowOk() (*int32, bool)`

GetAveragingWindowOk returns a tuple with the AveragingWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragingWindow

`func (o *FiveQICharacteristicsSingle) SetAveragingWindow(v int32)`

SetAveragingWindow sets AveragingWindow field to given value.

### HasAveragingWindow

`func (o *FiveQICharacteristicsSingle) HasAveragingWindow() bool`

HasAveragingWindow returns a boolean if a field has been set.

### GetMaximumDataBurstVolume

`func (o *FiveQICharacteristicsSingle) GetMaximumDataBurstVolume() int32`

GetMaximumDataBurstVolume returns the MaximumDataBurstVolume field if non-nil, zero value otherwise.

### GetMaximumDataBurstVolumeOk

`func (o *FiveQICharacteristicsSingle) GetMaximumDataBurstVolumeOk() (*int32, bool)`

GetMaximumDataBurstVolumeOk returns a tuple with the MaximumDataBurstVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDataBurstVolume

`func (o *FiveQICharacteristicsSingle) SetMaximumDataBurstVolume(v int32)`

SetMaximumDataBurstVolume sets MaximumDataBurstVolume field to given value.

### HasMaximumDataBurstVolume

`func (o *FiveQICharacteristicsSingle) HasMaximumDataBurstVolume() bool`

HasMaximumDataBurstVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


