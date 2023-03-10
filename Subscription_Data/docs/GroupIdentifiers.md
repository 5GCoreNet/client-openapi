# GroupIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtGroupId** | Pointer to **string** |  | [optional] 
**IntGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**UeIdList** | Pointer to [**[]UeId**](UeId.md) |  | [optional] 
**AllowedAfIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupIdentifiers

`func NewGroupIdentifiers() *GroupIdentifiers`

NewGroupIdentifiers instantiates a new GroupIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupIdentifiersWithDefaults

`func NewGroupIdentifiersWithDefaults() *GroupIdentifiers`

NewGroupIdentifiersWithDefaults instantiates a new GroupIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtGroupId

`func (o *GroupIdentifiers) GetExtGroupId() string`

GetExtGroupId returns the ExtGroupId field if non-nil, zero value otherwise.

### GetExtGroupIdOk

`func (o *GroupIdentifiers) GetExtGroupIdOk() (*string, bool)`

GetExtGroupIdOk returns a tuple with the ExtGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGroupId

`func (o *GroupIdentifiers) SetExtGroupId(v string)`

SetExtGroupId sets ExtGroupId field to given value.

### HasExtGroupId

`func (o *GroupIdentifiers) HasExtGroupId() bool`

HasExtGroupId returns a boolean if a field has been set.

### GetIntGroupId

`func (o *GroupIdentifiers) GetIntGroupId() string`

GetIntGroupId returns the IntGroupId field if non-nil, zero value otherwise.

### GetIntGroupIdOk

`func (o *GroupIdentifiers) GetIntGroupIdOk() (*string, bool)`

GetIntGroupIdOk returns a tuple with the IntGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGroupId

`func (o *GroupIdentifiers) SetIntGroupId(v string)`

SetIntGroupId sets IntGroupId field to given value.

### HasIntGroupId

`func (o *GroupIdentifiers) HasIntGroupId() bool`

HasIntGroupId returns a boolean if a field has been set.

### GetUeIdList

`func (o *GroupIdentifiers) GetUeIdList() []UeId`

GetUeIdList returns the UeIdList field if non-nil, zero value otherwise.

### GetUeIdListOk

`func (o *GroupIdentifiers) GetUeIdListOk() (*[]UeId, bool)`

GetUeIdListOk returns a tuple with the UeIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIdList

`func (o *GroupIdentifiers) SetUeIdList(v []UeId)`

SetUeIdList sets UeIdList field to given value.

### HasUeIdList

`func (o *GroupIdentifiers) HasUeIdList() bool`

HasUeIdList returns a boolean if a field has been set.

### GetAllowedAfIds

`func (o *GroupIdentifiers) GetAllowedAfIds() []string`

GetAllowedAfIds returns the AllowedAfIds field if non-nil, zero value otherwise.

### GetAllowedAfIdsOk

`func (o *GroupIdentifiers) GetAllowedAfIdsOk() (*[]string, bool)`

GetAllowedAfIdsOk returns a tuple with the AllowedAfIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAfIds

`func (o *GroupIdentifiers) SetAllowedAfIds(v []string)`

SetAllowedAfIds sets AllowedAfIds field to given value.

### HasAllowedAfIds

`func (o *GroupIdentifiers) HasAllowedAfIds() bool`

HasAllowedAfIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


