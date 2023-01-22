# NefCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** |  | 
**AfEvents** | Pointer to [**[]AfEvent**](AfEvent.md) |  | [optional] 
**SnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**PfdData** | Pointer to [**PfdData**](PfdData.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ServedFqdnList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNefCond

`func NewNefCond(conditionType string, ) *NefCond`

NewNefCond instantiates a new NefCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefCondWithDefaults

`func NewNefCondWithDefaults() *NefCond`

NewNefCondWithDefaults instantiates a new NefCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *NefCond) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *NefCond) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *NefCond) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetAfEvents

`func (o *NefCond) GetAfEvents() []AfEvent`

GetAfEvents returns the AfEvents field if non-nil, zero value otherwise.

### GetAfEventsOk

`func (o *NefCond) GetAfEventsOk() (*[]AfEvent, bool)`

GetAfEventsOk returns a tuple with the AfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEvents

`func (o *NefCond) SetAfEvents(v []AfEvent)`

SetAfEvents sets AfEvents field to given value.

### HasAfEvents

`func (o *NefCond) HasAfEvents() bool`

HasAfEvents returns a boolean if a field has been set.

### GetSnssaiList

`func (o *NefCond) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *NefCond) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *NefCond) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.

### HasSnssaiList

`func (o *NefCond) HasSnssaiList() bool`

HasSnssaiList returns a boolean if a field has been set.

### GetPfdData

`func (o *NefCond) GetPfdData() PfdData`

GetPfdData returns the PfdData field if non-nil, zero value otherwise.

### GetPfdDataOk

`func (o *NefCond) GetPfdDataOk() (*PfdData, bool)`

GetPfdDataOk returns a tuple with the PfdData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdData

`func (o *NefCond) SetPfdData(v PfdData)`

SetPfdData sets PfdData field to given value.

### HasPfdData

`func (o *NefCond) HasPfdData() bool`

HasPfdData returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NefCond) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NefCond) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NefCond) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NefCond) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *NefCond) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *NefCond) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *NefCond) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *NefCond) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetServedFqdnList

`func (o *NefCond) GetServedFqdnList() []string`

GetServedFqdnList returns the ServedFqdnList field if non-nil, zero value otherwise.

### GetServedFqdnListOk

`func (o *NefCond) GetServedFqdnListOk() (*[]string, bool)`

GetServedFqdnListOk returns a tuple with the ServedFqdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedFqdnList

`func (o *NefCond) SetServedFqdnList(v []string)`

SetServedFqdnList sets ServedFqdnList field to given value.

### HasServedFqdnList

`func (o *NefCond) HasServedFqdnList() bool`

HasServedFqdnList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


