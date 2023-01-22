# MbsServiceArea

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NcgiList** | Pointer to [**[]NcgiTai**](NcgiTai.md) | List of NR cell Ids | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) | List of tracking area Ids | [optional] 

## Methods

### NewMbsServiceArea

`func NewMbsServiceArea() *MbsServiceArea`

NewMbsServiceArea instantiates a new MbsServiceArea object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsServiceAreaWithDefaults

`func NewMbsServiceAreaWithDefaults() *MbsServiceArea`

NewMbsServiceAreaWithDefaults instantiates a new MbsServiceArea object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNcgiList

`func (o *MbsServiceArea) GetNcgiList() []NcgiTai`

GetNcgiList returns the NcgiList field if non-nil, zero value otherwise.

### GetNcgiListOk

`func (o *MbsServiceArea) GetNcgiListOk() (*[]NcgiTai, bool)`

GetNcgiListOk returns a tuple with the NcgiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgiList

`func (o *MbsServiceArea) SetNcgiList(v []NcgiTai)`

SetNcgiList sets NcgiList field to given value.

### HasNcgiList

`func (o *MbsServiceArea) HasNcgiList() bool`

HasNcgiList returns a boolean if a field has been set.

### GetTaiList

`func (o *MbsServiceArea) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *MbsServiceArea) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *MbsServiceArea) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *MbsServiceArea) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


