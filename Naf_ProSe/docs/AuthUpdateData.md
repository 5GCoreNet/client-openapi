# AuthUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetRpauid** | **string** | Contains the RPAUID. | 
**BannedAuthData** | [**[]BannedAuthData**](BannedAuthData.md) |  | 

## Methods

### NewAuthUpdateData

`func NewAuthUpdateData(targetRpauid string, bannedAuthData []BannedAuthData, ) *AuthUpdateData`

NewAuthUpdateData instantiates a new AuthUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUpdateDataWithDefaults

`func NewAuthUpdateDataWithDefaults() *AuthUpdateData`

NewAuthUpdateDataWithDefaults instantiates a new AuthUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetRpauid

`func (o *AuthUpdateData) GetTargetRpauid() string`

GetTargetRpauid returns the TargetRpauid field if non-nil, zero value otherwise.

### GetTargetRpauidOk

`func (o *AuthUpdateData) GetTargetRpauidOk() (*string, bool)`

GetTargetRpauidOk returns a tuple with the TargetRpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRpauid

`func (o *AuthUpdateData) SetTargetRpauid(v string)`

SetTargetRpauid sets TargetRpauid field to given value.


### GetBannedAuthData

`func (o *AuthUpdateData) GetBannedAuthData() []BannedAuthData`

GetBannedAuthData returns the BannedAuthData field if non-nil, zero value otherwise.

### GetBannedAuthDataOk

`func (o *AuthUpdateData) GetBannedAuthDataOk() (*[]BannedAuthData, bool)`

GetBannedAuthDataOk returns a tuple with the BannedAuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedAuthData

`func (o *AuthUpdateData) SetBannedAuthData(v []BannedAuthData)`

SetBannedAuthData sets BannedAuthData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


