# UnTrustAfInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | **string** |  | 
**SNssaiInfoList** | Pointer to [**[]SnssaiInfoItem**](SnssaiInfoItem.md) |  | [optional] 
**MappingInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewUnTrustAfInfo

`func NewUnTrustAfInfo(afId string, ) *UnTrustAfInfo`

NewUnTrustAfInfo instantiates a new UnTrustAfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnTrustAfInfoWithDefaults

`func NewUnTrustAfInfoWithDefaults() *UnTrustAfInfo`

NewUnTrustAfInfoWithDefaults instantiates a new UnTrustAfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *UnTrustAfInfo) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *UnTrustAfInfo) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *UnTrustAfInfo) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetSNssaiInfoList

`func (o *UnTrustAfInfo) GetSNssaiInfoList() []SnssaiInfoItem`

GetSNssaiInfoList returns the SNssaiInfoList field if non-nil, zero value otherwise.

### GetSNssaiInfoListOk

`func (o *UnTrustAfInfo) GetSNssaiInfoListOk() (*[]SnssaiInfoItem, bool)`

GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiInfoList

`func (o *UnTrustAfInfo) SetSNssaiInfoList(v []SnssaiInfoItem)`

SetSNssaiInfoList sets SNssaiInfoList field to given value.

### HasSNssaiInfoList

`func (o *UnTrustAfInfo) HasSNssaiInfoList() bool`

HasSNssaiInfoList returns a boolean if a field has been set.

### GetMappingInd

`func (o *UnTrustAfInfo) GetMappingInd() bool`

GetMappingInd returns the MappingInd field if non-nil, zero value otherwise.

### GetMappingIndOk

`func (o *UnTrustAfInfo) GetMappingIndOk() (*bool, bool)`

GetMappingIndOk returns a tuple with the MappingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingInd

`func (o *UnTrustAfInfo) SetMappingInd(v bool)`

SetMappingInd sets MappingInd field to given value.

### HasMappingInd

`func (o *UnTrustAfInfo) HasMappingInd() bool`

HasMappingInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


