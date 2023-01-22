# TrustAfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiInfoList** | Pointer to [**[]SnssaiInfoItem**](SnssaiInfoItem.md) |  | [optional] 
**AfEvents** | Pointer to [**[]AfEvent**](AfEvent.md) |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**InternalGroupId** | Pointer to **[]string** |  | [optional] 
**MappingInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTrustAfInfo

`func NewTrustAfInfo() *TrustAfInfo`

NewTrustAfInfo instantiates a new TrustAfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrustAfInfoWithDefaults

`func NewTrustAfInfoWithDefaults() *TrustAfInfo`

NewTrustAfInfoWithDefaults instantiates a new TrustAfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiInfoList

`func (o *TrustAfInfo) GetSNssaiInfoList() []SnssaiInfoItem`

GetSNssaiInfoList returns the SNssaiInfoList field if non-nil, zero value otherwise.

### GetSNssaiInfoListOk

`func (o *TrustAfInfo) GetSNssaiInfoListOk() (*[]SnssaiInfoItem, bool)`

GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiInfoList

`func (o *TrustAfInfo) SetSNssaiInfoList(v []SnssaiInfoItem)`

SetSNssaiInfoList sets SNssaiInfoList field to given value.

### HasSNssaiInfoList

`func (o *TrustAfInfo) HasSNssaiInfoList() bool`

HasSNssaiInfoList returns a boolean if a field has been set.

### GetAfEvents

`func (o *TrustAfInfo) GetAfEvents() []AfEvent`

GetAfEvents returns the AfEvents field if non-nil, zero value otherwise.

### GetAfEventsOk

`func (o *TrustAfInfo) GetAfEventsOk() (*[]AfEvent, bool)`

GetAfEventsOk returns a tuple with the AfEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEvents

`func (o *TrustAfInfo) SetAfEvents(v []AfEvent)`

SetAfEvents sets AfEvents field to given value.

### HasAfEvents

`func (o *TrustAfInfo) HasAfEvents() bool`

HasAfEvents returns a boolean if a field has been set.

### GetAppIds

`func (o *TrustAfInfo) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *TrustAfInfo) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *TrustAfInfo) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *TrustAfInfo) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetInternalGroupId

`func (o *TrustAfInfo) GetInternalGroupId() []string`

GetInternalGroupId returns the InternalGroupId field if non-nil, zero value otherwise.

### GetInternalGroupIdOk

`func (o *TrustAfInfo) GetInternalGroupIdOk() (*[]string, bool)`

GetInternalGroupIdOk returns a tuple with the InternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupId

`func (o *TrustAfInfo) SetInternalGroupId(v []string)`

SetInternalGroupId sets InternalGroupId field to given value.

### HasInternalGroupId

`func (o *TrustAfInfo) HasInternalGroupId() bool`

HasInternalGroupId returns a boolean if a field has been set.

### GetMappingInd

`func (o *TrustAfInfo) GetMappingInd() bool`

GetMappingInd returns the MappingInd field if non-nil, zero value otherwise.

### GetMappingIndOk

`func (o *TrustAfInfo) GetMappingIndOk() (*bool, bool)`

GetMappingIndOk returns a tuple with the MappingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingInd

`func (o *TrustAfInfo) SetMappingInd(v bool)`

SetMappingInd sets MappingInd field to given value.

### HasMappingInd

`func (o *TrustAfInfo) HasMappingInd() bool`

HasMappingInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


