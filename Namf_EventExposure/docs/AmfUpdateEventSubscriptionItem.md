# AmfUpdateEventSubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** |  | 
**Path** | **string** |  | 
**Value** | Pointer to [**AmfEvent**](AmfEvent.md) |  | [optional] 
**PresenceInfo** | Pointer to [**PresenceInfo**](PresenceInfo.md) |  | [optional] 
**ExcludeSupiList** | Pointer to **[]string** |  | [optional] 
**ExcludeGpsiList** | Pointer to **[]string** |  | [optional] 
**IncludeSupiList** | Pointer to **[]string** |  | [optional] 
**IncludeGpsiList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAmfUpdateEventSubscriptionItem

`func NewAmfUpdateEventSubscriptionItem(op string, path string, ) *AmfUpdateEventSubscriptionItem`

NewAmfUpdateEventSubscriptionItem instantiates a new AmfUpdateEventSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfUpdateEventSubscriptionItemWithDefaults

`func NewAmfUpdateEventSubscriptionItemWithDefaults() *AmfUpdateEventSubscriptionItem`

NewAmfUpdateEventSubscriptionItemWithDefaults instantiates a new AmfUpdateEventSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *AmfUpdateEventSubscriptionItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *AmfUpdateEventSubscriptionItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *AmfUpdateEventSubscriptionItem) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *AmfUpdateEventSubscriptionItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AmfUpdateEventSubscriptionItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AmfUpdateEventSubscriptionItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *AmfUpdateEventSubscriptionItem) GetValue() AmfEvent`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AmfUpdateEventSubscriptionItem) GetValueOk() (*AmfEvent, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AmfUpdateEventSubscriptionItem) SetValue(v AmfEvent)`

SetValue sets Value field to given value.

### HasValue

`func (o *AmfUpdateEventSubscriptionItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetPresenceInfo

`func (o *AmfUpdateEventSubscriptionItem) GetPresenceInfo() PresenceInfo`

GetPresenceInfo returns the PresenceInfo field if non-nil, zero value otherwise.

### GetPresenceInfoOk

`func (o *AmfUpdateEventSubscriptionItem) GetPresenceInfoOk() (*PresenceInfo, bool)`

GetPresenceInfoOk returns a tuple with the PresenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInfo

`func (o *AmfUpdateEventSubscriptionItem) SetPresenceInfo(v PresenceInfo)`

SetPresenceInfo sets PresenceInfo field to given value.

### HasPresenceInfo

`func (o *AmfUpdateEventSubscriptionItem) HasPresenceInfo() bool`

HasPresenceInfo returns a boolean if a field has been set.

### GetExcludeSupiList

`func (o *AmfUpdateEventSubscriptionItem) GetExcludeSupiList() []string`

GetExcludeSupiList returns the ExcludeSupiList field if non-nil, zero value otherwise.

### GetExcludeSupiListOk

`func (o *AmfUpdateEventSubscriptionItem) GetExcludeSupiListOk() (*[]string, bool)`

GetExcludeSupiListOk returns a tuple with the ExcludeSupiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSupiList

`func (o *AmfUpdateEventSubscriptionItem) SetExcludeSupiList(v []string)`

SetExcludeSupiList sets ExcludeSupiList field to given value.

### HasExcludeSupiList

`func (o *AmfUpdateEventSubscriptionItem) HasExcludeSupiList() bool`

HasExcludeSupiList returns a boolean if a field has been set.

### GetExcludeGpsiList

`func (o *AmfUpdateEventSubscriptionItem) GetExcludeGpsiList() []string`

GetExcludeGpsiList returns the ExcludeGpsiList field if non-nil, zero value otherwise.

### GetExcludeGpsiListOk

`func (o *AmfUpdateEventSubscriptionItem) GetExcludeGpsiListOk() (*[]string, bool)`

GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGpsiList

`func (o *AmfUpdateEventSubscriptionItem) SetExcludeGpsiList(v []string)`

SetExcludeGpsiList sets ExcludeGpsiList field to given value.

### HasExcludeGpsiList

`func (o *AmfUpdateEventSubscriptionItem) HasExcludeGpsiList() bool`

HasExcludeGpsiList returns a boolean if a field has been set.

### GetIncludeSupiList

`func (o *AmfUpdateEventSubscriptionItem) GetIncludeSupiList() []string`

GetIncludeSupiList returns the IncludeSupiList field if non-nil, zero value otherwise.

### GetIncludeSupiListOk

`func (o *AmfUpdateEventSubscriptionItem) GetIncludeSupiListOk() (*[]string, bool)`

GetIncludeSupiListOk returns a tuple with the IncludeSupiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSupiList

`func (o *AmfUpdateEventSubscriptionItem) SetIncludeSupiList(v []string)`

SetIncludeSupiList sets IncludeSupiList field to given value.

### HasIncludeSupiList

`func (o *AmfUpdateEventSubscriptionItem) HasIncludeSupiList() bool`

HasIncludeSupiList returns a boolean if a field has been set.

### GetIncludeGpsiList

`func (o *AmfUpdateEventSubscriptionItem) GetIncludeGpsiList() []string`

GetIncludeGpsiList returns the IncludeGpsiList field if non-nil, zero value otherwise.

### GetIncludeGpsiListOk

`func (o *AmfUpdateEventSubscriptionItem) GetIncludeGpsiListOk() (*[]string, bool)`

GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGpsiList

`func (o *AmfUpdateEventSubscriptionItem) SetIncludeGpsiList(v []string)`

SetIncludeGpsiList sets IncludeGpsiList field to given value.

### HasIncludeGpsiList

`func (o *AmfUpdateEventSubscriptionItem) HasIncludeGpsiList() bool`

HasIncludeGpsiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


