# MatchInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**OpenMatchInfoForOpen** | Pointer to [**MatchInfoForOpen**](MatchInfoForOpen.md) |  | [optional] 
**RestrictedMatchInfo** | Pointer to [**MatchInfoForRestricted**](MatchInfoForRestricted.md) |  | [optional] 

## Methods

### NewMatchInformation

`func NewMatchInformation(discType DiscoveryType, ) *MatchInformation`

NewMatchInformation instantiates a new MatchInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchInformationWithDefaults

`func NewMatchInformationWithDefaults() *MatchInformation`

NewMatchInformationWithDefaults instantiates a new MatchInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *MatchInformation) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *MatchInformation) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *MatchInformation) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetOpenMatchInfoForOpen

`func (o *MatchInformation) GetOpenMatchInfoForOpen() MatchInfoForOpen`

GetOpenMatchInfoForOpen returns the OpenMatchInfoForOpen field if non-nil, zero value otherwise.

### GetOpenMatchInfoForOpenOk

`func (o *MatchInformation) GetOpenMatchInfoForOpenOk() (*MatchInfoForOpen, bool)`

GetOpenMatchInfoForOpenOk returns a tuple with the OpenMatchInfoForOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenMatchInfoForOpen

`func (o *MatchInformation) SetOpenMatchInfoForOpen(v MatchInfoForOpen)`

SetOpenMatchInfoForOpen sets OpenMatchInfoForOpen field to given value.

### HasOpenMatchInfoForOpen

`func (o *MatchInformation) HasOpenMatchInfoForOpen() bool`

HasOpenMatchInfoForOpen returns a boolean if a field has been set.

### GetRestrictedMatchInfo

`func (o *MatchInformation) GetRestrictedMatchInfo() MatchInfoForRestricted`

GetRestrictedMatchInfo returns the RestrictedMatchInfo field if non-nil, zero value otherwise.

### GetRestrictedMatchInfoOk

`func (o *MatchInformation) GetRestrictedMatchInfoOk() (*MatchInfoForRestricted, bool)`

GetRestrictedMatchInfoOk returns a tuple with the RestrictedMatchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedMatchInfo

`func (o *MatchInformation) SetRestrictedMatchInfo(v MatchInfoForRestricted)`

SetRestrictedMatchInfo sets RestrictedMatchInfo field to given value.

### HasRestrictedMatchInfo

`func (o *MatchInformation) HasRestrictedMatchInfo() bool`

HasRestrictedMatchInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


