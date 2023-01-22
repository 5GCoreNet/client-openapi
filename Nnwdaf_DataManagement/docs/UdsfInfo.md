# UdsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**StorageIdRanges** | Pointer to [**map[string][]IdentityRange**](array.md) | A map (list of key-value pairs) where realmId serves as key and each value in the map is an array of IdentityRanges. Each IdentityRange is a range of storageIds.  | [optional] 

## Methods

### NewUdsfInfo

`func NewUdsfInfo() *UdsfInfo`

NewUdsfInfo instantiates a new UdsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdsfInfoWithDefaults

`func NewUdsfInfoWithDefaults() *UdsfInfo`

NewUdsfInfoWithDefaults instantiates a new UdsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UdsfInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UdsfInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UdsfInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UdsfInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *UdsfInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *UdsfInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *UdsfInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *UdsfInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetStorageIdRanges

`func (o *UdsfInfo) GetStorageIdRanges() map[string][]IdentityRange`

GetStorageIdRanges returns the StorageIdRanges field if non-nil, zero value otherwise.

### GetStorageIdRangesOk

`func (o *UdsfInfo) GetStorageIdRangesOk() (*map[string][]IdentityRange, bool)`

GetStorageIdRangesOk returns a tuple with the StorageIdRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageIdRanges

`func (o *UdsfInfo) SetStorageIdRanges(v map[string][]IdentityRange)`

SetStorageIdRanges sets StorageIdRanges field to given value.

### HasStorageIdRanges

`func (o *UdsfInfo) HasStorageIdRanges() bool`

HasStorageIdRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


