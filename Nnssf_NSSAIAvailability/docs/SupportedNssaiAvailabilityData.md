# SupportedNssaiAvailabilityData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | [**Tai**](Tai.md) |  | 
**SupportedSnssaiList** | [**[]ExtSnssai**](ExtSnssai.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**NsagInfos** | Pointer to [**[]NsagInfo**](NsagInfo.md) |  | [optional] 

## Methods

### NewSupportedNssaiAvailabilityData

`func NewSupportedNssaiAvailabilityData(tai Tai, supportedSnssaiList []ExtSnssai, ) *SupportedNssaiAvailabilityData`

NewSupportedNssaiAvailabilityData instantiates a new SupportedNssaiAvailabilityData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedNssaiAvailabilityDataWithDefaults

`func NewSupportedNssaiAvailabilityDataWithDefaults() *SupportedNssaiAvailabilityData`

NewSupportedNssaiAvailabilityDataWithDefaults instantiates a new SupportedNssaiAvailabilityData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *SupportedNssaiAvailabilityData) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *SupportedNssaiAvailabilityData) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *SupportedNssaiAvailabilityData) SetTai(v Tai)`

SetTai sets Tai field to given value.


### GetSupportedSnssaiList

`func (o *SupportedNssaiAvailabilityData) GetSupportedSnssaiList() []ExtSnssai`

GetSupportedSnssaiList returns the SupportedSnssaiList field if non-nil, zero value otherwise.

### GetSupportedSnssaiListOk

`func (o *SupportedNssaiAvailabilityData) GetSupportedSnssaiListOk() (*[]ExtSnssai, bool)`

GetSupportedSnssaiListOk returns a tuple with the SupportedSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSnssaiList

`func (o *SupportedNssaiAvailabilityData) SetSupportedSnssaiList(v []ExtSnssai)`

SetSupportedSnssaiList sets SupportedSnssaiList field to given value.


### GetTaiList

`func (o *SupportedNssaiAvailabilityData) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *SupportedNssaiAvailabilityData) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *SupportedNssaiAvailabilityData) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *SupportedNssaiAvailabilityData) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *SupportedNssaiAvailabilityData) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *SupportedNssaiAvailabilityData) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *SupportedNssaiAvailabilityData) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *SupportedNssaiAvailabilityData) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetNsagInfos

`func (o *SupportedNssaiAvailabilityData) GetNsagInfos() []NsagInfo`

GetNsagInfos returns the NsagInfos field if non-nil, zero value otherwise.

### GetNsagInfosOk

`func (o *SupportedNssaiAvailabilityData) GetNsagInfosOk() (*[]NsagInfo, bool)`

GetNsagInfosOk returns a tuple with the NsagInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsagInfos

`func (o *SupportedNssaiAvailabilityData) SetNsagInfos(v []NsagInfo)`

SetNsagInfos sets NsagInfos field to given value.

### HasNsagInfos

`func (o *SupportedNssaiAvailabilityData) HasNsagInfos() bool`

HasNsagInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


