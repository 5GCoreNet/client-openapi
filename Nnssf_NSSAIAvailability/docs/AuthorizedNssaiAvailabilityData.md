# AuthorizedNssaiAvailabilityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | [**Tai**](Tai.md) |  | 
**SupportedSnssaiList** | [**[]ExtSnssai**](ExtSnssai.md) |  | 
**RestrictedSnssaiList** | Pointer to [**[]RestrictedSnssai**](RestrictedSnssai.md) |  | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**NsagInfos** | Pointer to [**[]NsagInfo**](NsagInfo.md) |  | [optional] 

## Methods

### NewAuthorizedNssaiAvailabilityData

`func NewAuthorizedNssaiAvailabilityData(tai Tai, supportedSnssaiList []ExtSnssai, ) *AuthorizedNssaiAvailabilityData`

NewAuthorizedNssaiAvailabilityData instantiates a new AuthorizedNssaiAvailabilityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizedNssaiAvailabilityDataWithDefaults

`func NewAuthorizedNssaiAvailabilityDataWithDefaults() *AuthorizedNssaiAvailabilityData`

NewAuthorizedNssaiAvailabilityDataWithDefaults instantiates a new AuthorizedNssaiAvailabilityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *AuthorizedNssaiAvailabilityData) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *AuthorizedNssaiAvailabilityData) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *AuthorizedNssaiAvailabilityData) SetTai(v Tai)`

SetTai sets Tai field to given value.


### GetSupportedSnssaiList

`func (o *AuthorizedNssaiAvailabilityData) GetSupportedSnssaiList() []ExtSnssai`

GetSupportedSnssaiList returns the SupportedSnssaiList field if non-nil, zero value otherwise.

### GetSupportedSnssaiListOk

`func (o *AuthorizedNssaiAvailabilityData) GetSupportedSnssaiListOk() (*[]ExtSnssai, bool)`

GetSupportedSnssaiListOk returns a tuple with the SupportedSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSnssaiList

`func (o *AuthorizedNssaiAvailabilityData) SetSupportedSnssaiList(v []ExtSnssai)`

SetSupportedSnssaiList sets SupportedSnssaiList field to given value.


### GetRestrictedSnssaiList

`func (o *AuthorizedNssaiAvailabilityData) GetRestrictedSnssaiList() []RestrictedSnssai`

GetRestrictedSnssaiList returns the RestrictedSnssaiList field if non-nil, zero value otherwise.

### GetRestrictedSnssaiListOk

`func (o *AuthorizedNssaiAvailabilityData) GetRestrictedSnssaiListOk() (*[]RestrictedSnssai, bool)`

GetRestrictedSnssaiListOk returns a tuple with the RestrictedSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedSnssaiList

`func (o *AuthorizedNssaiAvailabilityData) SetRestrictedSnssaiList(v []RestrictedSnssai)`

SetRestrictedSnssaiList sets RestrictedSnssaiList field to given value.

### HasRestrictedSnssaiList

`func (o *AuthorizedNssaiAvailabilityData) HasRestrictedSnssaiList() bool`

HasRestrictedSnssaiList returns a boolean if a field has been set.

### GetTaiList

`func (o *AuthorizedNssaiAvailabilityData) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *AuthorizedNssaiAvailabilityData) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *AuthorizedNssaiAvailabilityData) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *AuthorizedNssaiAvailabilityData) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *AuthorizedNssaiAvailabilityData) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *AuthorizedNssaiAvailabilityData) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *AuthorizedNssaiAvailabilityData) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *AuthorizedNssaiAvailabilityData) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetNsagInfos

`func (o *AuthorizedNssaiAvailabilityData) GetNsagInfos() []NsagInfo`

GetNsagInfos returns the NsagInfos field if non-nil, zero value otherwise.

### GetNsagInfosOk

`func (o *AuthorizedNssaiAvailabilityData) GetNsagInfosOk() (*[]NsagInfo, bool)`

GetNsagInfosOk returns a tuple with the NsagInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsagInfos

`func (o *AuthorizedNssaiAvailabilityData) SetNsagInfos(v []NsagInfo)`

SetNsagInfos sets NsagInfos field to given value.

### HasNsagInfos

`func (o *AuthorizedNssaiAvailabilityData) HasNsagInfos() bool`

HasNsagInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


