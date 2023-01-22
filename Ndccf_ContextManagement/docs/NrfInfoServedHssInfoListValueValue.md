# NrfInfoServedHssInfoListValueValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**ImsiRanges** | Pointer to [**[]ImsiRange**](ImsiRange.md) |  | [optional] 
**ImsPrivateIdentityRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ImsPublicIdentityRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**MsisdnRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ExternalGroupIdentifiersRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**HssDiameterAddress** | Pointer to [**NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 
**AdditionalDiamAddresses** | Pointer to [**[]NetworkNodeDiameterAddress**](NetworkNodeDiameterAddress.md) |  | [optional] 

## Methods

### NewNrfInfoServedHssInfoListValueValue

`func NewNrfInfoServedHssInfoListValueValue() *NrfInfoServedHssInfoListValueValue`

NewNrfInfoServedHssInfoListValueValue instantiates a new NrfInfoServedHssInfoListValueValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedHssInfoListValueValueWithDefaults

`func NewNrfInfoServedHssInfoListValueValueWithDefaults() *NrfInfoServedHssInfoListValueValue`

NewNrfInfoServedHssInfoListValueValueWithDefaults instantiates a new NrfInfoServedHssInfoListValueValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *NrfInfoServedHssInfoListValueValue) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *NrfInfoServedHssInfoListValueValue) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *NrfInfoServedHssInfoListValueValue) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *NrfInfoServedHssInfoListValueValue) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetImsiRanges

`func (o *NrfInfoServedHssInfoListValueValue) GetImsiRanges() []ImsiRange`

GetImsiRanges returns the ImsiRanges field if non-nil, zero value otherwise.

### GetImsiRangesOk

`func (o *NrfInfoServedHssInfoListValueValue) GetImsiRangesOk() (*[]ImsiRange, bool)`

GetImsiRangesOk returns a tuple with the ImsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsiRanges

`func (o *NrfInfoServedHssInfoListValueValue) SetImsiRanges(v []ImsiRange)`

SetImsiRanges sets ImsiRanges field to given value.

### HasImsiRanges

`func (o *NrfInfoServedHssInfoListValueValue) HasImsiRanges() bool`

HasImsiRanges returns a boolean if a field has been set.

### GetImsPrivateIdentityRanges

`func (o *NrfInfoServedHssInfoListValueValue) GetImsPrivateIdentityRanges() []IdentityRange`

GetImsPrivateIdentityRanges returns the ImsPrivateIdentityRanges field if non-nil, zero value otherwise.

### GetImsPrivateIdentityRangesOk

`func (o *NrfInfoServedHssInfoListValueValue) GetImsPrivateIdentityRangesOk() (*[]IdentityRange, bool)`

GetImsPrivateIdentityRangesOk returns a tuple with the ImsPrivateIdentityRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsPrivateIdentityRanges

`func (o *NrfInfoServedHssInfoListValueValue) SetImsPrivateIdentityRanges(v []IdentityRange)`

SetImsPrivateIdentityRanges sets ImsPrivateIdentityRanges field to given value.

### HasImsPrivateIdentityRanges

`func (o *NrfInfoServedHssInfoListValueValue) HasImsPrivateIdentityRanges() bool`

HasImsPrivateIdentityRanges returns a boolean if a field has been set.

### GetImsPublicIdentityRanges

`func (o *NrfInfoServedHssInfoListValueValue) GetImsPublicIdentityRanges() []IdentityRange`

GetImsPublicIdentityRanges returns the ImsPublicIdentityRanges field if non-nil, zero value otherwise.

### GetImsPublicIdentityRangesOk

`func (o *NrfInfoServedHssInfoListValueValue) GetImsPublicIdentityRangesOk() (*[]IdentityRange, bool)`

GetImsPublicIdentityRangesOk returns a tuple with the ImsPublicIdentityRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsPublicIdentityRanges

`func (o *NrfInfoServedHssInfoListValueValue) SetImsPublicIdentityRanges(v []IdentityRange)`

SetImsPublicIdentityRanges sets ImsPublicIdentityRanges field to given value.

### HasImsPublicIdentityRanges

`func (o *NrfInfoServedHssInfoListValueValue) HasImsPublicIdentityRanges() bool`

HasImsPublicIdentityRanges returns a boolean if a field has been set.

### GetMsisdnRanges

`func (o *NrfInfoServedHssInfoListValueValue) GetMsisdnRanges() []IdentityRange`

GetMsisdnRanges returns the MsisdnRanges field if non-nil, zero value otherwise.

### GetMsisdnRangesOk

`func (o *NrfInfoServedHssInfoListValueValue) GetMsisdnRangesOk() (*[]IdentityRange, bool)`

GetMsisdnRangesOk returns a tuple with the MsisdnRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdnRanges

`func (o *NrfInfoServedHssInfoListValueValue) SetMsisdnRanges(v []IdentityRange)`

SetMsisdnRanges sets MsisdnRanges field to given value.

### HasMsisdnRanges

`func (o *NrfInfoServedHssInfoListValueValue) HasMsisdnRanges() bool`

HasMsisdnRanges returns a boolean if a field has been set.

### GetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedHssInfoListValueValue) GetExternalGroupIdentifiersRanges() []IdentityRange`

GetExternalGroupIdentifiersRanges returns the ExternalGroupIdentifiersRanges field if non-nil, zero value otherwise.

### GetExternalGroupIdentifiersRangesOk

`func (o *NrfInfoServedHssInfoListValueValue) GetExternalGroupIdentifiersRangesOk() (*[]IdentityRange, bool)`

GetExternalGroupIdentifiersRangesOk returns a tuple with the ExternalGroupIdentifiersRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupIdentifiersRanges

`func (o *NrfInfoServedHssInfoListValueValue) SetExternalGroupIdentifiersRanges(v []IdentityRange)`

SetExternalGroupIdentifiersRanges sets ExternalGroupIdentifiersRanges field to given value.

### HasExternalGroupIdentifiersRanges

`func (o *NrfInfoServedHssInfoListValueValue) HasExternalGroupIdentifiersRanges() bool`

HasExternalGroupIdentifiersRanges returns a boolean if a field has been set.

### GetHssDiameterAddress

`func (o *NrfInfoServedHssInfoListValueValue) GetHssDiameterAddress() NetworkNodeDiameterAddress`

GetHssDiameterAddress returns the HssDiameterAddress field if non-nil, zero value otherwise.

### GetHssDiameterAddressOk

`func (o *NrfInfoServedHssInfoListValueValue) GetHssDiameterAddressOk() (*NetworkNodeDiameterAddress, bool)`

GetHssDiameterAddressOk returns a tuple with the HssDiameterAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssDiameterAddress

`func (o *NrfInfoServedHssInfoListValueValue) SetHssDiameterAddress(v NetworkNodeDiameterAddress)`

SetHssDiameterAddress sets HssDiameterAddress field to given value.

### HasHssDiameterAddress

`func (o *NrfInfoServedHssInfoListValueValue) HasHssDiameterAddress() bool`

HasHssDiameterAddress returns a boolean if a field has been set.

### GetAdditionalDiamAddresses

`func (o *NrfInfoServedHssInfoListValueValue) GetAdditionalDiamAddresses() []NetworkNodeDiameterAddress`

GetAdditionalDiamAddresses returns the AdditionalDiamAddresses field if non-nil, zero value otherwise.

### GetAdditionalDiamAddressesOk

`func (o *NrfInfoServedHssInfoListValueValue) GetAdditionalDiamAddressesOk() (*[]NetworkNodeDiameterAddress, bool)`

GetAdditionalDiamAddressesOk returns a tuple with the AdditionalDiamAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDiamAddresses

`func (o *NrfInfoServedHssInfoListValueValue) SetAdditionalDiamAddresses(v []NetworkNodeDiameterAddress)`

SetAdditionalDiamAddresses sets AdditionalDiamAddresses field to given value.

### HasAdditionalDiamAddresses

`func (o *NrfInfoServedHssInfoListValueValue) HasAdditionalDiamAddresses() bool`

HasAdditionalDiamAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


