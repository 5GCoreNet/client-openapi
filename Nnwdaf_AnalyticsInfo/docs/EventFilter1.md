# EventFilter1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsis** | Pointer to **[]string** |  | [optional] 
**Supis** | Pointer to **[]string** |  | [optional] 
**ExterGroupIds** | Pointer to **[]string** |  | [optional] 
**InterGroupIds** | Pointer to **[]string** |  | [optional] 
**AnyUeInd** | Pointer to **bool** |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**LocArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**CollAttrs** | Pointer to [**[]CollectiveBehaviourFilter**](CollectiveBehaviourFilter.md) |  | [optional] 

## Methods

### NewEventFilter1

`func NewEventFilter1() *EventFilter1`

NewEventFilter1 instantiates a new EventFilter1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFilter1WithDefaults

`func NewEventFilter1WithDefaults() *EventFilter1`

NewEventFilter1WithDefaults instantiates a new EventFilter1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsis

`func (o *EventFilter1) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *EventFilter1) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *EventFilter1) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *EventFilter1) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetSupis

`func (o *EventFilter1) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *EventFilter1) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *EventFilter1) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *EventFilter1) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetExterGroupIds

`func (o *EventFilter1) GetExterGroupIds() []string`

GetExterGroupIds returns the ExterGroupIds field if non-nil, zero value otherwise.

### GetExterGroupIdsOk

`func (o *EventFilter1) GetExterGroupIdsOk() (*[]string, bool)`

GetExterGroupIdsOk returns a tuple with the ExterGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupIds

`func (o *EventFilter1) SetExterGroupIds(v []string)`

SetExterGroupIds sets ExterGroupIds field to given value.

### HasExterGroupIds

`func (o *EventFilter1) HasExterGroupIds() bool`

HasExterGroupIds returns a boolean if a field has been set.

### GetInterGroupIds

`func (o *EventFilter1) GetInterGroupIds() []string`

GetInterGroupIds returns the InterGroupIds field if non-nil, zero value otherwise.

### GetInterGroupIdsOk

`func (o *EventFilter1) GetInterGroupIdsOk() (*[]string, bool)`

GetInterGroupIdsOk returns a tuple with the InterGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupIds

`func (o *EventFilter1) SetInterGroupIds(v []string)`

SetInterGroupIds sets InterGroupIds field to given value.

### HasInterGroupIds

`func (o *EventFilter1) HasInterGroupIds() bool`

HasInterGroupIds returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *EventFilter1) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *EventFilter1) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *EventFilter1) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *EventFilter1) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetAppIds

`func (o *EventFilter1) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventFilter1) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventFilter1) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventFilter1) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetLocArea

`func (o *EventFilter1) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *EventFilter1) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *EventFilter1) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *EventFilter1) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetCollAttrs

`func (o *EventFilter1) GetCollAttrs() []CollectiveBehaviourFilter`

GetCollAttrs returns the CollAttrs field if non-nil, zero value otherwise.

### GetCollAttrsOk

`func (o *EventFilter1) GetCollAttrsOk() (*[]CollectiveBehaviourFilter, bool)`

GetCollAttrsOk returns a tuple with the CollAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollAttrs

`func (o *EventFilter1) SetCollAttrs(v []CollectiveBehaviourFilter)`

SetCollAttrs sets CollAttrs field to given value.

### HasCollAttrs

`func (o *EventFilter1) HasCollAttrs() bool`

HasCollAttrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


