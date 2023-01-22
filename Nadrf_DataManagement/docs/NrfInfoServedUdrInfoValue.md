# NrfInfoServedUdrInfoValue

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

### NewNrfInfoServedUdrInfoValue

`func NewNrfInfoServedUdrInfoValue() *NrfInfoServedUdrInfoValue`

NewNrfInfoServedUdrInfoValue instantiates a new NrfInfoServedUdrInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedUdrInfoValueWithDefaults

`func NewNrfInfoServedUdrInfoValueWithDefaults() *NrfInfoServedUdrInfoValue`

NewNrfInfoServedUdrInfoValueWithDefaults instantiates a new NrfInfoServedUdrInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *NrfInfoServedUdrInfoValue) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NrfInfoServedUdrInfoValue) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NrfInfoServedUdrInfoValue) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NrfInfoServedUdrInfoValue) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetSupiRanges

`func (o *NrfInfoServedUdrInfoValue) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *NrfInfoServedUdrInfoValue) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *NrfInfoServedUdrInfoValue) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *NrfInfoServedUdrInfoValue) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *NrfInfoServedUdrInfoValue) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *NrfInfoServedUdrInfoValue) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *NrfInfoServedUdrInfoValue) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *NrfInfoServedUdrInfoValue) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdrInfoValue) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *NrfInfoServedUdrInfoValue) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdrInfoValue) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *NrfInfoServedUdrInfoValue) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetSupportedDataSets

`func (o *NrfInfoServedUdrInfoValue) GetSupportedDataSets() []DataSetId`

GetSupportedDataSets returns the SupportedDataSets field if non-nil, zero value otherwise.

### GetSupportedDataSetsOk

`func (o *NrfInfoServedUdrInfoValue) GetSupportedDataSetsOk() (*[]DataSetId, bool)`

GetSupportedDataSetsOk returns a tuple with the SupportedDataSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDataSets

`func (o *NrfInfoServedUdrInfoValue) SetSupportedDataSets(v []DataSetId)`

SetSupportedDataSets sets SupportedDataSets field to given value.

### HasSupportedDataSets

`func (o *NrfInfoServedUdrInfoValue) HasSupportedDataSets() bool`

HasSupportedDataSets returns a boolean if a field has been set.

### GetSharedDataIdRanges

`func (o *NrfInfoServedUdrInfoValue) GetSharedDataIdRanges() []SharedDataIdRange`

GetSharedDataIdRanges returns the SharedDataIdRanges field if non-nil, zero value otherwise.

### GetSharedDataIdRangesOk

`func (o *NrfInfoServedUdrInfoValue) GetSharedDataIdRangesOk() (*[]SharedDataIdRange, bool)`

GetSharedDataIdRangesOk returns a tuple with the SharedDataIdRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataIdRanges

`func (o *NrfInfoServedUdrInfoValue) SetSharedDataIdRanges(v []SharedDataIdRange)`

SetSharedDataIdRanges sets SharedDataIdRanges field to given value.

### HasSharedDataIdRanges

`func (o *NrfInfoServedUdrInfoValue) HasSharedDataIdRanges() bool`

HasSharedDataIdRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


