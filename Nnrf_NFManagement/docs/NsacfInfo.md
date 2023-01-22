# NsacfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NsacfCapability** | [**NsacfCapability**](NsacfCapability.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 

## Methods

### NewNsacfInfo

`func NewNsacfInfo(nsacfCapability NsacfCapability, ) *NsacfInfo`

NewNsacfInfo instantiates a new NsacfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsacfInfoWithDefaults

`func NewNsacfInfoWithDefaults() *NsacfInfo`

NewNsacfInfoWithDefaults instantiates a new NsacfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNsacfCapability

`func (o *NsacfInfo) GetNsacfCapability() NsacfCapability`

GetNsacfCapability returns the NsacfCapability field if non-nil, zero value otherwise.

### GetNsacfCapabilityOk

`func (o *NsacfInfo) GetNsacfCapabilityOk() (*NsacfCapability, bool)`

GetNsacfCapabilityOk returns a tuple with the NsacfCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacfCapability

`func (o *NsacfInfo) SetNsacfCapability(v NsacfCapability)`

SetNsacfCapability sets NsacfCapability field to given value.


### GetTaiList

`func (o *NsacfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NsacfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NsacfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NsacfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NsacfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NsacfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NsacfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NsacfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


