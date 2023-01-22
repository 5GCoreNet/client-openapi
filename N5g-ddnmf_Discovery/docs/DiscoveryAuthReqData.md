# DiscoveryAuthReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**RestrictedDiscData** | Pointer to [**DiscDataForRestricted**](DiscDataForRestricted.md) |  | [optional] 

## Methods

### NewDiscoveryAuthReqData

`func NewDiscoveryAuthReqData(discType DiscoveryType, ) *DiscoveryAuthReqData`

NewDiscoveryAuthReqData instantiates a new DiscoveryAuthReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryAuthReqDataWithDefaults

`func NewDiscoveryAuthReqDataWithDefaults() *DiscoveryAuthReqData`

NewDiscoveryAuthReqDataWithDefaults instantiates a new DiscoveryAuthReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *DiscoveryAuthReqData) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *DiscoveryAuthReqData) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *DiscoveryAuthReqData) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetRestrictedDiscData

`func (o *DiscoveryAuthReqData) GetRestrictedDiscData() DiscDataForRestricted`

GetRestrictedDiscData returns the RestrictedDiscData field if non-nil, zero value otherwise.

### GetRestrictedDiscDataOk

`func (o *DiscoveryAuthReqData) GetRestrictedDiscDataOk() (*DiscDataForRestricted, bool)`

GetRestrictedDiscDataOk returns a tuple with the RestrictedDiscData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDiscData

`func (o *DiscoveryAuthReqData) SetRestrictedDiscData(v DiscDataForRestricted)`

SetRestrictedDiscData sets RestrictedDiscData field to given value.

### HasRestrictedDiscData

`func (o *DiscoveryAuthReqData) HasRestrictedDiscData() bool`

HasRestrictedDiscData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


