# BannedAuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BannedRpauid** | **string** | Contains the RPAUID. | 
**BannedPduid** | **string** | Contains the PDUID. | 
**RevocationResult** | Pointer to [**RevocationResult**](RevocationResult.md) |  | [optional] 

## Methods

### NewBannedAuthData

`func NewBannedAuthData(bannedRpauid string, bannedPduid string, ) *BannedAuthData`

NewBannedAuthData instantiates a new BannedAuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBannedAuthDataWithDefaults

`func NewBannedAuthDataWithDefaults() *BannedAuthData`

NewBannedAuthDataWithDefaults instantiates a new BannedAuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBannedRpauid

`func (o *BannedAuthData) GetBannedRpauid() string`

GetBannedRpauid returns the BannedRpauid field if non-nil, zero value otherwise.

### GetBannedRpauidOk

`func (o *BannedAuthData) GetBannedRpauidOk() (*string, bool)`

GetBannedRpauidOk returns a tuple with the BannedRpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedRpauid

`func (o *BannedAuthData) SetBannedRpauid(v string)`

SetBannedRpauid sets BannedRpauid field to given value.


### GetBannedPduid

`func (o *BannedAuthData) GetBannedPduid() string`

GetBannedPduid returns the BannedPduid field if non-nil, zero value otherwise.

### GetBannedPduidOk

`func (o *BannedAuthData) GetBannedPduidOk() (*string, bool)`

GetBannedPduidOk returns a tuple with the BannedPduid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannedPduid

`func (o *BannedAuthData) SetBannedPduid(v string)`

SetBannedPduid sets BannedPduid field to given value.


### GetRevocationResult

`func (o *BannedAuthData) GetRevocationResult() RevocationResult`

GetRevocationResult returns the RevocationResult field if non-nil, zero value otherwise.

### GetRevocationResultOk

`func (o *BannedAuthData) GetRevocationResultOk() (*RevocationResult, bool)`

GetRevocationResultOk returns a tuple with the RevocationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationResult

`func (o *BannedAuthData) SetRevocationResult(v RevocationResult)`

SetRevocationResult sets RevocationResult field to given value.

### HasRevocationResult

`func (o *BannedAuthData) HasRevocationResult() bool`

HasRevocationResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


