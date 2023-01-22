# UdrInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**SupportedDataSets** | Pointer to [**[]DataSetId**](DataSetId.md) |  | [optional] 
**SharedDataIdRanges** | Pointer to [**[]SharedDataIdRange**](SharedDataIdRange.md) |  | [optional] 

## Methods

### NewUdrInfo

`func NewUdrInfo() *UdrInfo`

NewUdrInfo instantiates a new UdrInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdrInfoWithDefaults

`func NewUdrInfoWithDefaults() *UdrInfo`

NewUdrInfoWithDefaults instantiates a new UdrInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UdrInfo) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UdrInfo) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UdrInfo) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UdrInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *UdrInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *UdrInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *UdrInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *UdrInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *UdrInfo) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *UdrInfo) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *UdrInfo) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *UdrInfo) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *UdrInfo) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *UdrInfo) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *UdrInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *UdrInfo) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetSupportedDataSets

`func (o *UdrInfo) GetSupportedDataSets() []DataSetId`

GetSupportedDataSets returns the SupportedDataSets field if non-nil, zero value otherwise.

### GetSupportedDataSetsOk

`func (o *UdrInfo) GetSupportedDataSetsOk() (*[]DataSetId, bool)`

GetSupportedDataSetsOk returns a tuple with the SupportedDataSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDataSets

`func (o *UdrInfo) SetSupportedDataSets(v []DataSetId)`

SetSupportedDataSets sets SupportedDataSets field to given value.

### HasSupportedDataSets

`func (o *UdrInfo) HasSupportedDataSets() bool`

HasSupportedDataSets returns a boolean if a field has been set.

### GetSharedDataIdRanges

`func (o *UdrInfo) GetSharedDataIdRanges() []SharedDataIdRange`

GetSharedDataIdRanges returns the SharedDataIdRanges field if non-nil, zero value otherwise.

### GetSharedDataIdRangesOk

`func (o *UdrInfo) GetSharedDataIdRangesOk() (*[]SharedDataIdRange, bool)`

GetSharedDataIdRangesOk returns a tuple with the SharedDataIdRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataIdRanges

`func (o *UdrInfo) SetSharedDataIdRanges(v []SharedDataIdRange)`

SetSharedDataIdRanges sets SharedDataIdRanges field to given value.

### HasSharedDataIdRanges

`func (o *UdrInfo) HasSharedDataIdRanges() bool`

HasSharedDataIdRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


