# CollectiveBehaviourInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColAttrib** | [**[]PerUeAttribute**](PerUeAttribute.md) |  | 
**NoOfUes** | Pointer to **int32** | Total number of UEs that fulfil a collective within the area of interest. | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**ExtUeIds** | Pointer to **[]string** |  | [optional] 
**UeIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCollectiveBehaviourInfo

`func NewCollectiveBehaviourInfo(colAttrib []PerUeAttribute, ) *CollectiveBehaviourInfo`

NewCollectiveBehaviourInfo instantiates a new CollectiveBehaviourInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectiveBehaviourInfoWithDefaults

`func NewCollectiveBehaviourInfoWithDefaults() *CollectiveBehaviourInfo`

NewCollectiveBehaviourInfoWithDefaults instantiates a new CollectiveBehaviourInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColAttrib

`func (o *CollectiveBehaviourInfo) GetColAttrib() []PerUeAttribute`

GetColAttrib returns the ColAttrib field if non-nil, zero value otherwise.

### GetColAttribOk

`func (o *CollectiveBehaviourInfo) GetColAttribOk() (*[]PerUeAttribute, bool)`

GetColAttribOk returns a tuple with the ColAttrib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColAttrib

`func (o *CollectiveBehaviourInfo) SetColAttrib(v []PerUeAttribute)`

SetColAttrib sets ColAttrib field to given value.


### GetNoOfUes

`func (o *CollectiveBehaviourInfo) GetNoOfUes() int32`

GetNoOfUes returns the NoOfUes field if non-nil, zero value otherwise.

### GetNoOfUesOk

`func (o *CollectiveBehaviourInfo) GetNoOfUesOk() (*int32, bool)`

GetNoOfUesOk returns a tuple with the NoOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoOfUes

`func (o *CollectiveBehaviourInfo) SetNoOfUes(v int32)`

SetNoOfUes sets NoOfUes field to given value.

### HasNoOfUes

`func (o *CollectiveBehaviourInfo) HasNoOfUes() bool`

HasNoOfUes returns a boolean if a field has been set.

### GetAppIds

`func (o *CollectiveBehaviourInfo) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *CollectiveBehaviourInfo) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *CollectiveBehaviourInfo) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *CollectiveBehaviourInfo) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetExtUeIds

`func (o *CollectiveBehaviourInfo) GetExtUeIds() []string`

GetExtUeIds returns the ExtUeIds field if non-nil, zero value otherwise.

### GetExtUeIdsOk

`func (o *CollectiveBehaviourInfo) GetExtUeIdsOk() (*[]string, bool)`

GetExtUeIdsOk returns a tuple with the ExtUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtUeIds

`func (o *CollectiveBehaviourInfo) SetExtUeIds(v []string)`

SetExtUeIds sets ExtUeIds field to given value.

### HasExtUeIds

`func (o *CollectiveBehaviourInfo) HasExtUeIds() bool`

HasExtUeIds returns a boolean if a field has been set.

### GetUeIds

`func (o *CollectiveBehaviourInfo) GetUeIds() []string`

GetUeIds returns the UeIds field if non-nil, zero value otherwise.

### GetUeIdsOk

`func (o *CollectiveBehaviourInfo) GetUeIdsOk() (*[]string, bool)`

GetUeIdsOk returns a tuple with the UeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIds

`func (o *CollectiveBehaviourInfo) SetUeIds(v []string)`

SetUeIds sets UeIds field to given value.

### HasUeIds

`func (o *CollectiveBehaviourInfo) HasUeIds() bool`

HasUeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


