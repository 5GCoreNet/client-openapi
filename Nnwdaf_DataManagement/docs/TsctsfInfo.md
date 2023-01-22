# TsctsfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssaiInfoList** | Pointer to [**map[string]SnssaiTsctsfInfoItem**](SnssaiTsctsfInfoItem.md) | A map (list of key-value pairs) where a valid JSON string serves as key | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**InternalGroupIdentifiersRanges** | Pointer to [**[]InternalGroupIdRange**](InternalGroupIdRange.md) |  | [optional] 

## Methods

### NewTsctsfInfo

`func NewTsctsfInfo() *TsctsfInfo`

NewTsctsfInfo instantiates a new TsctsfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTsctsfInfoWithDefaults

`func NewTsctsfInfoWithDefaults() *TsctsfInfo`

NewTsctsfInfoWithDefaults instantiates a new TsctsfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssaiInfoList

`func (o *TsctsfInfo) GetSNssaiInfoList() map[string]SnssaiTsctsfInfoItem`

GetSNssaiInfoList returns the SNssaiInfoList field if non-nil, zero value otherwise.

### GetSNssaiInfoListOk

`func (o *TsctsfInfo) GetSNssaiInfoListOk() (*map[string]SnssaiTsctsfInfoItem, bool)`

GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiInfoList

`func (o *TsctsfInfo) SetSNssaiInfoList(v map[string]SnssaiTsctsfInfoItem)`

SetSNssaiInfoList sets SNssaiInfoList field to given value.

### HasSNssaiInfoList

`func (o *TsctsfInfo) HasSNssaiInfoList() bool`

HasSNssaiInfoList returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *TsctsfInfo) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *TsctsfInfo) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *TsctsfInfo) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *TsctsfInfo) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetSupiRanges

`func (o *TsctsfInfo) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *TsctsfInfo) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *TsctsfInfo) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *TsctsfInfo) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *TsctsfInfo) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *TsctsfInfo) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *TsctsfInfo) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *TsctsfInfo) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetInternalGroupIdentifiersRanges

`func (o *TsctsfInfo) GetInternalGroupIdentifiersRanges() []InternalGroupIdRange`

GetInternalGroupIdentifiersRanges returns the InternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetInternalGroupIdentifiersRangesOk

`func (o *TsctsfInfo) GetInternalGroupIdentifiersRangesOk() (*[]InternalGroupIdRange, bool)`

GetInternalGroupIdentifiersRangesOk returns a tuple with the InternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupIdentifiersRanges

`func (o *TsctsfInfo) SetInternalGroupIdentifiersRanges(v []InternalGroupIdRange)`

SetInternalGroupIdentifiersRanges sets InternalGroupIdentifiersRanges field to given value.

### HasInternalGroupIdentifiersRanges

`func (o *TsctsfInfo) HasInternalGroupIdentifiersRanges() bool`

HasInternalGroupIdentifiersRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


