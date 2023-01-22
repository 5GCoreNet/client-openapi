# AnnounceAuthReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**OpenDiscData** | Pointer to [**AnnounceDiscDataForOpen**](AnnounceDiscDataForOpen.md) |  | [optional] 
**RestrictedDiscData** | Pointer to [**AnnounceDiscDataForRestricted**](AnnounceDiscDataForRestricted.md) |  | [optional] 

## Methods

### NewAnnounceAuthReqData

`func NewAnnounceAuthReqData(discType DiscoveryType, ) *AnnounceAuthReqData`

NewAnnounceAuthReqData instantiates a new AnnounceAuthReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnounceAuthReqDataWithDefaults

`func NewAnnounceAuthReqDataWithDefaults() *AnnounceAuthReqData`

NewAnnounceAuthReqDataWithDefaults instantiates a new AnnounceAuthReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *AnnounceAuthReqData) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *AnnounceAuthReqData) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *AnnounceAuthReqData) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetOpenDiscData

`func (o *AnnounceAuthReqData) GetOpenDiscData() AnnounceDiscDataForOpen`

GetOpenDiscData returns the OpenDiscData field if non-nil, zero value otherwise.

### GetOpenDiscDataOk

`func (o *AnnounceAuthReqData) GetOpenDiscDataOk() (*AnnounceDiscDataForOpen, bool)`

GetOpenDiscDataOk returns a tuple with the OpenDiscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenDiscData

`func (o *AnnounceAuthReqData) SetOpenDiscData(v AnnounceDiscDataForOpen)`

SetOpenDiscData sets OpenDiscData field to given value.

### HasOpenDiscData

`func (o *AnnounceAuthReqData) HasOpenDiscData() bool`

HasOpenDiscData returns a boolean if a field has been set.

### GetRestrictedDiscData

`func (o *AnnounceAuthReqData) GetRestrictedDiscData() AnnounceDiscDataForRestricted`

GetRestrictedDiscData returns the RestrictedDiscData field if non-nil, zero value otherwise.

### GetRestrictedDiscDataOk

`func (o *AnnounceAuthReqData) GetRestrictedDiscDataOk() (*AnnounceDiscDataForRestricted, bool)`

GetRestrictedDiscDataOk returns a tuple with the RestrictedDiscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDiscData

`func (o *AnnounceAuthReqData) SetRestrictedDiscData(v AnnounceDiscDataForRestricted)`

SetRestrictedDiscData sets RestrictedDiscData field to given value.

### HasRestrictedDiscData

`func (o *AnnounceAuthReqData) HasRestrictedDiscData() bool`

HasRestrictedDiscData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


