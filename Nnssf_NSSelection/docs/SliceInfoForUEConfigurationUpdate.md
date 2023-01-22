# SliceInfoForUEConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscribedNssai** | Pointer to [**[]SubscribedSnssai**](SubscribedSnssai.md) |  | [optional] 
**AllowedNssaiCurrentAccess** | Pointer to [**AllowedNssai**](AllowedNssai.md) |  | [optional] 
**AllowedNssaiOtherAccess** | Pointer to [**AllowedNssai**](AllowedNssai.md) |  | [optional] 
**DefaultConfiguredSnssaiInd** | Pointer to **bool** |  | [optional] 
**RequestedNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**MappingOfNssai** | Pointer to [**[]MappingOfSnssai**](MappingOfSnssai.md) |  | [optional] 
**UeSupNssrgInd** | Pointer to **bool** |  | [optional] 
**SuppressNssrgInd** | Pointer to **bool** |  | [optional] 
**RejectedNssaiRa** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**NsagSupported** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSliceInfoForUEConfigurationUpdate

`func NewSliceInfoForUEConfigurationUpdate() *SliceInfoForUEConfigurationUpdate`

NewSliceInfoForUEConfigurationUpdate instantiates a new SliceInfoForUEConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceInfoForUEConfigurationUpdateWithDefaults

`func NewSliceInfoForUEConfigurationUpdateWithDefaults() *SliceInfoForUEConfigurationUpdate`

NewSliceInfoForUEConfigurationUpdateWithDefaults instantiates a new SliceInfoForUEConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribedNssai

`func (o *SliceInfoForUEConfigurationUpdate) GetSubscribedNssai() []SubscribedSnssai`

GetSubscribedNssai returns the SubscribedNssai field if non-nil, zero value otherwise.

### GetSubscribedNssaiOk

`func (o *SliceInfoForUEConfigurationUpdate) GetSubscribedNssaiOk() (*[]SubscribedSnssai, bool)`

GetSubscribedNssaiOk returns a tuple with the SubscribedNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedNssai

`func (o *SliceInfoForUEConfigurationUpdate) SetSubscribedNssai(v []SubscribedSnssai)`

SetSubscribedNssai sets SubscribedNssai field to given value.

### HasSubscribedNssai

`func (o *SliceInfoForUEConfigurationUpdate) HasSubscribedNssai() bool`

HasSubscribedNssai returns a boolean if a field has been set.

### GetAllowedNssaiCurrentAccess

`func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiCurrentAccess() AllowedNssai`

GetAllowedNssaiCurrentAccess returns the AllowedNssaiCurrentAccess field if non-nil, zero value otherwise.

### GetAllowedNssaiCurrentAccessOk

`func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiCurrentAccessOk() (*AllowedNssai, bool)`

GetAllowedNssaiCurrentAccessOk returns a tuple with the AllowedNssaiCurrentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssaiCurrentAccess

`func (o *SliceInfoForUEConfigurationUpdate) SetAllowedNssaiCurrentAccess(v AllowedNssai)`

SetAllowedNssaiCurrentAccess sets AllowedNssaiCurrentAccess field to given value.

### HasAllowedNssaiCurrentAccess

`func (o *SliceInfoForUEConfigurationUpdate) HasAllowedNssaiCurrentAccess() bool`

HasAllowedNssaiCurrentAccess returns a boolean if a field has been set.

### GetAllowedNssaiOtherAccess

`func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiOtherAccess() AllowedNssai`

GetAllowedNssaiOtherAccess returns the AllowedNssaiOtherAccess field if non-nil, zero value otherwise.

### GetAllowedNssaiOtherAccessOk

`func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiOtherAccessOk() (*AllowedNssai, bool)`

GetAllowedNssaiOtherAccessOk returns a tuple with the AllowedNssaiOtherAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssaiOtherAccess

`func (o *SliceInfoForUEConfigurationUpdate) SetAllowedNssaiOtherAccess(v AllowedNssai)`

SetAllowedNssaiOtherAccess sets AllowedNssaiOtherAccess field to given value.

### HasAllowedNssaiOtherAccess

`func (o *SliceInfoForUEConfigurationUpdate) HasAllowedNssaiOtherAccess() bool`

HasAllowedNssaiOtherAccess returns a boolean if a field has been set.

### GetDefaultConfiguredSnssaiInd

`func (o *SliceInfoForUEConfigurationUpdate) GetDefaultConfiguredSnssaiInd() bool`

GetDefaultConfiguredSnssaiInd returns the DefaultConfiguredSnssaiInd field if non-nil, zero value otherwise.

### GetDefaultConfiguredSnssaiIndOk

`func (o *SliceInfoForUEConfigurationUpdate) GetDefaultConfiguredSnssaiIndOk() (*bool, bool)`

GetDefaultConfiguredSnssaiIndOk returns a tuple with the DefaultConfiguredSnssaiInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfiguredSnssaiInd

`func (o *SliceInfoForUEConfigurationUpdate) SetDefaultConfiguredSnssaiInd(v bool)`

SetDefaultConfiguredSnssaiInd sets DefaultConfiguredSnssaiInd field to given value.

### HasDefaultConfiguredSnssaiInd

`func (o *SliceInfoForUEConfigurationUpdate) HasDefaultConfiguredSnssaiInd() bool`

HasDefaultConfiguredSnssaiInd returns a boolean if a field has been set.

### GetRequestedNssai

`func (o *SliceInfoForUEConfigurationUpdate) GetRequestedNssai() []Snssai`

GetRequestedNssai returns the RequestedNssai field if non-nil, zero value otherwise.

### GetRequestedNssaiOk

`func (o *SliceInfoForUEConfigurationUpdate) GetRequestedNssaiOk() (*[]Snssai, bool)`

GetRequestedNssaiOk returns a tuple with the RequestedNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedNssai

`func (o *SliceInfoForUEConfigurationUpdate) SetRequestedNssai(v []Snssai)`

SetRequestedNssai sets RequestedNssai field to given value.

### HasRequestedNssai

`func (o *SliceInfoForUEConfigurationUpdate) HasRequestedNssai() bool`

HasRequestedNssai returns a boolean if a field has been set.

### GetMappingOfNssai

`func (o *SliceInfoForUEConfigurationUpdate) GetMappingOfNssai() []MappingOfSnssai`

GetMappingOfNssai returns the MappingOfNssai field if non-nil, zero value otherwise.

### GetMappingOfNssaiOk

`func (o *SliceInfoForUEConfigurationUpdate) GetMappingOfNssaiOk() (*[]MappingOfSnssai, bool)`

GetMappingOfNssaiOk returns a tuple with the MappingOfNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingOfNssai

`func (o *SliceInfoForUEConfigurationUpdate) SetMappingOfNssai(v []MappingOfSnssai)`

SetMappingOfNssai sets MappingOfNssai field to given value.

### HasMappingOfNssai

`func (o *SliceInfoForUEConfigurationUpdate) HasMappingOfNssai() bool`

HasMappingOfNssai returns a boolean if a field has been set.

### GetUeSupNssrgInd

`func (o *SliceInfoForUEConfigurationUpdate) GetUeSupNssrgInd() bool`

GetUeSupNssrgInd returns the UeSupNssrgInd field if non-nil, zero value otherwise.

### GetUeSupNssrgIndOk

`func (o *SliceInfoForUEConfigurationUpdate) GetUeSupNssrgIndOk() (*bool, bool)`

GetUeSupNssrgIndOk returns a tuple with the UeSupNssrgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSupNssrgInd

`func (o *SliceInfoForUEConfigurationUpdate) SetUeSupNssrgInd(v bool)`

SetUeSupNssrgInd sets UeSupNssrgInd field to given value.

### HasUeSupNssrgInd

`func (o *SliceInfoForUEConfigurationUpdate) HasUeSupNssrgInd() bool`

HasUeSupNssrgInd returns a boolean if a field has been set.

### GetSuppressNssrgInd

`func (o *SliceInfoForUEConfigurationUpdate) GetSuppressNssrgInd() bool`

GetSuppressNssrgInd returns the SuppressNssrgInd field if non-nil, zero value otherwise.

### GetSuppressNssrgIndOk

`func (o *SliceInfoForUEConfigurationUpdate) GetSuppressNssrgIndOk() (*bool, bool)`

GetSuppressNssrgIndOk returns a tuple with the SuppressNssrgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressNssrgInd

`func (o *SliceInfoForUEConfigurationUpdate) SetSuppressNssrgInd(v bool)`

SetSuppressNssrgInd sets SuppressNssrgInd field to given value.

### HasSuppressNssrgInd

`func (o *SliceInfoForUEConfigurationUpdate) HasSuppressNssrgInd() bool`

HasSuppressNssrgInd returns a boolean if a field has been set.

### GetRejectedNssaiRa

`func (o *SliceInfoForUEConfigurationUpdate) GetRejectedNssaiRa() []Snssai`

GetRejectedNssaiRa returns the RejectedNssaiRa field if non-nil, zero value otherwise.

### GetRejectedNssaiRaOk

`func (o *SliceInfoForUEConfigurationUpdate) GetRejectedNssaiRaOk() (*[]Snssai, bool)`

GetRejectedNssaiRaOk returns a tuple with the RejectedNssaiRa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedNssaiRa

`func (o *SliceInfoForUEConfigurationUpdate) SetRejectedNssaiRa(v []Snssai)`

SetRejectedNssaiRa sets RejectedNssaiRa field to given value.

### HasRejectedNssaiRa

`func (o *SliceInfoForUEConfigurationUpdate) HasRejectedNssaiRa() bool`

HasRejectedNssaiRa returns a boolean if a field has been set.

### GetNsagSupported

`func (o *SliceInfoForUEConfigurationUpdate) GetNsagSupported() bool`

GetNsagSupported returns the NsagSupported field if non-nil, zero value otherwise.

### GetNsagSupportedOk

`func (o *SliceInfoForUEConfigurationUpdate) GetNsagSupportedOk() (*bool, bool)`

GetNsagSupportedOk returns a tuple with the NsagSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsagSupported

`func (o *SliceInfoForUEConfigurationUpdate) SetNsagSupported(v bool)`

SetNsagSupported sets NsagSupported field to given value.

### HasNsagSupported

`func (o *SliceInfoForUEConfigurationUpdate) HasNsagSupported() bool`

HasNsagSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


