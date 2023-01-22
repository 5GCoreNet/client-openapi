# SliceInfoForRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscribedNssai** | Pointer to [**[]SubscribedSnssai**](SubscribedSnssai.md) |  | [optional] 
**AllowedNssaiCurrentAccess** | Pointer to [**AllowedNssai**](AllowedNssai.md) |  | [optional] 
**AllowedNssaiOtherAccess** | Pointer to [**AllowedNssai**](AllowedNssai.md) |  | [optional] 
**SNssaiForMapping** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RequestedNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**DefaultConfiguredSnssaiInd** | Pointer to **bool** |  | [optional] [default to false]
**MappingOfNssai** | Pointer to [**[]MappingOfSnssai**](MappingOfSnssai.md) |  | [optional] 
**RequestMapping** | Pointer to **bool** |  | [optional] 
**UeSupNssrgInd** | Pointer to **bool** |  | [optional] 
**SuppressNssrgInd** | Pointer to **bool** |  | [optional] 
**NsagSupported** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSliceInfoForRegistration

`func NewSliceInfoForRegistration() *SliceInfoForRegistration`

NewSliceInfoForRegistration instantiates a new SliceInfoForRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceInfoForRegistrationWithDefaults

`func NewSliceInfoForRegistrationWithDefaults() *SliceInfoForRegistration`

NewSliceInfoForRegistrationWithDefaults instantiates a new SliceInfoForRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscribedNssai

`func (o *SliceInfoForRegistration) GetSubscribedNssai() []SubscribedSnssai`

GetSubscribedNssai returns the SubscribedNssai field if non-nil, zero value otherwise.

### GetSubscribedNssaiOk

`func (o *SliceInfoForRegistration) GetSubscribedNssaiOk() (*[]SubscribedSnssai, bool)`

GetSubscribedNssaiOk returns a tuple with the SubscribedNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedNssai

`func (o *SliceInfoForRegistration) SetSubscribedNssai(v []SubscribedSnssai)`

SetSubscribedNssai sets SubscribedNssai field to given value.

### HasSubscribedNssai

`func (o *SliceInfoForRegistration) HasSubscribedNssai() bool`

HasSubscribedNssai returns a boolean if a field has been set.

### GetAllowedNssaiCurrentAccess

`func (o *SliceInfoForRegistration) GetAllowedNssaiCurrentAccess() AllowedNssai`

GetAllowedNssaiCurrentAccess returns the AllowedNssaiCurrentAccess field if non-nil, zero value otherwise.

### GetAllowedNssaiCurrentAccessOk

`func (o *SliceInfoForRegistration) GetAllowedNssaiCurrentAccessOk() (*AllowedNssai, bool)`

GetAllowedNssaiCurrentAccessOk returns a tuple with the AllowedNssaiCurrentAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssaiCurrentAccess

`func (o *SliceInfoForRegistration) SetAllowedNssaiCurrentAccess(v AllowedNssai)`

SetAllowedNssaiCurrentAccess sets AllowedNssaiCurrentAccess field to given value.

### HasAllowedNssaiCurrentAccess

`func (o *SliceInfoForRegistration) HasAllowedNssaiCurrentAccess() bool`

HasAllowedNssaiCurrentAccess returns a boolean if a field has been set.

### GetAllowedNssaiOtherAccess

`func (o *SliceInfoForRegistration) GetAllowedNssaiOtherAccess() AllowedNssai`

GetAllowedNssaiOtherAccess returns the AllowedNssaiOtherAccess field if non-nil, zero value otherwise.

### GetAllowedNssaiOtherAccessOk

`func (o *SliceInfoForRegistration) GetAllowedNssaiOtherAccessOk() (*AllowedNssai, bool)`

GetAllowedNssaiOtherAccessOk returns a tuple with the AllowedNssaiOtherAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssaiOtherAccess

`func (o *SliceInfoForRegistration) SetAllowedNssaiOtherAccess(v AllowedNssai)`

SetAllowedNssaiOtherAccess sets AllowedNssaiOtherAccess field to given value.

### HasAllowedNssaiOtherAccess

`func (o *SliceInfoForRegistration) HasAllowedNssaiOtherAccess() bool`

HasAllowedNssaiOtherAccess returns a boolean if a field has been set.

### GetSNssaiForMapping

`func (o *SliceInfoForRegistration) GetSNssaiForMapping() []Snssai`

GetSNssaiForMapping returns the SNssaiForMapping field if non-nil, zero value otherwise.

### GetSNssaiForMappingOk

`func (o *SliceInfoForRegistration) GetSNssaiForMappingOk() (*[]Snssai, bool)`

GetSNssaiForMappingOk returns a tuple with the SNssaiForMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiForMapping

`func (o *SliceInfoForRegistration) SetSNssaiForMapping(v []Snssai)`

SetSNssaiForMapping sets SNssaiForMapping field to given value.

### HasSNssaiForMapping

`func (o *SliceInfoForRegistration) HasSNssaiForMapping() bool`

HasSNssaiForMapping returns a boolean if a field has been set.

### GetRequestedNssai

`func (o *SliceInfoForRegistration) GetRequestedNssai() []Snssai`

GetRequestedNssai returns the RequestedNssai field if non-nil, zero value otherwise.

### GetRequestedNssaiOk

`func (o *SliceInfoForRegistration) GetRequestedNssaiOk() (*[]Snssai, bool)`

GetRequestedNssaiOk returns a tuple with the RequestedNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedNssai

`func (o *SliceInfoForRegistration) SetRequestedNssai(v []Snssai)`

SetRequestedNssai sets RequestedNssai field to given value.

### HasRequestedNssai

`func (o *SliceInfoForRegistration) HasRequestedNssai() bool`

HasRequestedNssai returns a boolean if a field has been set.

### GetDefaultConfiguredSnssaiInd

`func (o *SliceInfoForRegistration) GetDefaultConfiguredSnssaiInd() bool`

GetDefaultConfiguredSnssaiInd returns the DefaultConfiguredSnssaiInd field if non-nil, zero value otherwise.

### GetDefaultConfiguredSnssaiIndOk

`func (o *SliceInfoForRegistration) GetDefaultConfiguredSnssaiIndOk() (*bool, bool)`

GetDefaultConfiguredSnssaiIndOk returns a tuple with the DefaultConfiguredSnssaiInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfiguredSnssaiInd

`func (o *SliceInfoForRegistration) SetDefaultConfiguredSnssaiInd(v bool)`

SetDefaultConfiguredSnssaiInd sets DefaultConfiguredSnssaiInd field to given value.

### HasDefaultConfiguredSnssaiInd

`func (o *SliceInfoForRegistration) HasDefaultConfiguredSnssaiInd() bool`

HasDefaultConfiguredSnssaiInd returns a boolean if a field has been set.

### GetMappingOfNssai

`func (o *SliceInfoForRegistration) GetMappingOfNssai() []MappingOfSnssai`

GetMappingOfNssai returns the MappingOfNssai field if non-nil, zero value otherwise.

### GetMappingOfNssaiOk

`func (o *SliceInfoForRegistration) GetMappingOfNssaiOk() (*[]MappingOfSnssai, bool)`

GetMappingOfNssaiOk returns a tuple with the MappingOfNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingOfNssai

`func (o *SliceInfoForRegistration) SetMappingOfNssai(v []MappingOfSnssai)`

SetMappingOfNssai sets MappingOfNssai field to given value.

### HasMappingOfNssai

`func (o *SliceInfoForRegistration) HasMappingOfNssai() bool`

HasMappingOfNssai returns a boolean if a field has been set.

### GetRequestMapping

`func (o *SliceInfoForRegistration) GetRequestMapping() bool`

GetRequestMapping returns the RequestMapping field if non-nil, zero value otherwise.

### GetRequestMappingOk

`func (o *SliceInfoForRegistration) GetRequestMappingOk() (*bool, bool)`

GetRequestMappingOk returns a tuple with the RequestMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMapping

`func (o *SliceInfoForRegistration) SetRequestMapping(v bool)`

SetRequestMapping sets RequestMapping field to given value.

### HasRequestMapping

`func (o *SliceInfoForRegistration) HasRequestMapping() bool`

HasRequestMapping returns a boolean if a field has been set.

### GetUeSupNssrgInd

`func (o *SliceInfoForRegistration) GetUeSupNssrgInd() bool`

GetUeSupNssrgInd returns the UeSupNssrgInd field if non-nil, zero value otherwise.

### GetUeSupNssrgIndOk

`func (o *SliceInfoForRegistration) GetUeSupNssrgIndOk() (*bool, bool)`

GetUeSupNssrgIndOk returns a tuple with the UeSupNssrgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSupNssrgInd

`func (o *SliceInfoForRegistration) SetUeSupNssrgInd(v bool)`

SetUeSupNssrgInd sets UeSupNssrgInd field to given value.

### HasUeSupNssrgInd

`func (o *SliceInfoForRegistration) HasUeSupNssrgInd() bool`

HasUeSupNssrgInd returns a boolean if a field has been set.

### GetSuppressNssrgInd

`func (o *SliceInfoForRegistration) GetSuppressNssrgInd() bool`

GetSuppressNssrgInd returns the SuppressNssrgInd field if non-nil, zero value otherwise.

### GetSuppressNssrgIndOk

`func (o *SliceInfoForRegistration) GetSuppressNssrgIndOk() (*bool, bool)`

GetSuppressNssrgIndOk returns a tuple with the SuppressNssrgInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressNssrgInd

`func (o *SliceInfoForRegistration) SetSuppressNssrgInd(v bool)`

SetSuppressNssrgInd sets SuppressNssrgInd field to given value.

### HasSuppressNssrgInd

`func (o *SliceInfoForRegistration) HasSuppressNssrgInd() bool`

HasSuppressNssrgInd returns a boolean if a field has been set.

### GetNsagSupported

`func (o *SliceInfoForRegistration) GetNsagSupported() bool`

GetNsagSupported returns the NsagSupported field if non-nil, zero value otherwise.

### GetNsagSupportedOk

`func (o *SliceInfoForRegistration) GetNsagSupportedOk() (*bool, bool)`

GetNsagSupportedOk returns a tuple with the NsagSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsagSupported

`func (o *SliceInfoForRegistration) SetNsagSupported(v bool)`

SetNsagSupported sets NsagSupported field to given value.

### HasNsagSupported

`func (o *SliceInfoForRegistration) HasNsagSupported() bool`

HasNsagSupported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


