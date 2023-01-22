# EventFilter

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

### NewEventFilter

`func NewEventFilter() *EventFilter`

NewEventFilter instantiates a new EventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventFilterWithDefaults

`func NewEventFilterWithDefaults() *EventFilter`

NewEventFilterWithDefaults instantiates a new EventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsis

`func (o *EventFilter) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *EventFilter) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *EventFilter) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *EventFilter) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetSupis

`func (o *EventFilter) GetSupis() []string`

GetSupis returns the Supis field if non-nil, zero value otherwise.

### GetSupisOk

`func (o *EventFilter) GetSupisOk() (*[]string, bool)`

GetSupisOk returns a tuple with the Supis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupis

`func (o *EventFilter) SetSupis(v []string)`

SetSupis sets Supis field to given value.

### HasSupis

`func (o *EventFilter) HasSupis() bool`

HasSupis returns a boolean if a field has been set.

### GetExterGroupIds

`func (o *EventFilter) GetExterGroupIds() []string`

GetExterGroupIds returns the ExterGroupIds field if non-nil, zero value otherwise.

### GetExterGroupIdsOk

`func (o *EventFilter) GetExterGroupIdsOk() (*[]string, bool)`

GetExterGroupIdsOk returns a tuple with the ExterGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupIds

`func (o *EventFilter) SetExterGroupIds(v []string)`

SetExterGroupIds sets ExterGroupIds field to given value.

### HasExterGroupIds

`func (o *EventFilter) HasExterGroupIds() bool`

HasExterGroupIds returns a boolean if a field has been set.

### GetInterGroupIds

`func (o *EventFilter) GetInterGroupIds() []string`

GetInterGroupIds returns the InterGroupIds field if non-nil, zero value otherwise.

### GetInterGroupIdsOk

`func (o *EventFilter) GetInterGroupIdsOk() (*[]string, bool)`

GetInterGroupIdsOk returns a tuple with the InterGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupIds

`func (o *EventFilter) SetInterGroupIds(v []string)`

SetInterGroupIds sets InterGroupIds field to given value.

### HasInterGroupIds

`func (o *EventFilter) HasInterGroupIds() bool`

HasInterGroupIds returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *EventFilter) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *EventFilter) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *EventFilter) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *EventFilter) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetAppIds

`func (o *EventFilter) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *EventFilter) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *EventFilter) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *EventFilter) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetLocArea

`func (o *EventFilter) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *EventFilter) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *EventFilter) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *EventFilter) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetCollAttrs

`func (o *EventFilter) GetCollAttrs() []CollectiveBehaviourFilter`

GetCollAttrs returns the CollAttrs field if non-nil, zero value otherwise.

### GetCollAttrsOk

`func (o *EventFilter) GetCollAttrsOk() (*[]CollectiveBehaviourFilter, bool)`

GetCollAttrsOk returns a tuple with the CollAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollAttrs

`func (o *EventFilter) SetCollAttrs(v []CollectiveBehaviourFilter)`

SetCollAttrs sets CollAttrs field to given value.

### HasCollAttrs

`func (o *EventFilter) HasCollAttrs() bool`

HasCollAttrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


