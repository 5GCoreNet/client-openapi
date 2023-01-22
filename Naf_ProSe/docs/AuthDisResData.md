# AuthDisResData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResponseType** | [**AuthResponseType**](AuthResponseType.md) |  | 
**ProseAppCodeSuffixPool** | Pointer to [**ProseApplicationCodeSuffixPool**](ProseApplicationCodeSuffixPool.md) |  | [optional] 
**Pduids** | Pointer to **[]string** |  | [optional] 
**RestrictedCodeSuffixPool** | Pointer to [**[]RestrictedCodeSuffixPool**](RestrictedCodeSuffixPool.md) |  | [optional] 
**ProseAppMasks** | Pointer to **[]string** |  | [optional] 
**ProSeRestrictedMasks** | Pointer to **[]string** |  | [optional] 
**ResAppLevelContainer** | Pointer to **string** | Contains the Application Level Container. | [optional] 
**TargetDataSet** | Pointer to [**[]TargetData**](TargetData.md) |  | [optional] 
**TargetPduid** | Pointer to **string** | Contains the PDUID. | [optional] 
**MetaData** | Pointer to **string** | Contains the metadata. | [optional] 

## Methods

### NewAuthDisResData

`func NewAuthDisResData(authResponseType AuthResponseType, ) *AuthDisResData`

NewAuthDisResData instantiates a new AuthDisResData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthDisResDataWithDefaults

`func NewAuthDisResDataWithDefaults() *AuthDisResData`

NewAuthDisResDataWithDefaults instantiates a new AuthDisResData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResponseType

`func (o *AuthDisResData) GetAuthResponseType() AuthResponseType`

GetAuthResponseType returns the AuthResponseType field if non-nil, zero value otherwise.

### GetAuthResponseTypeOk

`func (o *AuthDisResData) GetAuthResponseTypeOk() (*AuthResponseType, bool)`

GetAuthResponseTypeOk returns a tuple with the AuthResponseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResponseType

`func (o *AuthDisResData) SetAuthResponseType(v AuthResponseType)`

SetAuthResponseType sets AuthResponseType field to given value.


### GetProseAppCodeSuffixPool

`func (o *AuthDisResData) GetProseAppCodeSuffixPool() ProseApplicationCodeSuffixPool`

GetProseAppCodeSuffixPool returns the ProseAppCodeSuffixPool field if non-nil, zero value otherwise.

### GetProseAppCodeSuffixPoolOk

`func (o *AuthDisResData) GetProseAppCodeSuffixPoolOk() (*ProseApplicationCodeSuffixPool, bool)`

GetProseAppCodeSuffixPoolOk returns a tuple with the ProseAppCodeSuffixPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCodeSuffixPool

`func (o *AuthDisResData) SetProseAppCodeSuffixPool(v ProseApplicationCodeSuffixPool)`

SetProseAppCodeSuffixPool sets ProseAppCodeSuffixPool field to given value.

### HasProseAppCodeSuffixPool

`func (o *AuthDisResData) HasProseAppCodeSuffixPool() bool`

HasProseAppCodeSuffixPool returns a boolean if a field has been set.

### GetPduids

`func (o *AuthDisResData) GetPduids() []string`

GetPduids returns the Pduids field if non-nil, zero value otherwise.

### GetPduidsOk

`func (o *AuthDisResData) GetPduidsOk() (*[]string, bool)`

GetPduidsOk returns a tuple with the Pduids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduids

`func (o *AuthDisResData) SetPduids(v []string)`

SetPduids sets Pduids field to given value.

### HasPduids

`func (o *AuthDisResData) HasPduids() bool`

HasPduids returns a boolean if a field has been set.

### GetRestrictedCodeSuffixPool

`func (o *AuthDisResData) GetRestrictedCodeSuffixPool() []RestrictedCodeSuffixPool`

GetRestrictedCodeSuffixPool returns the RestrictedCodeSuffixPool field if non-nil, zero value otherwise.

### GetRestrictedCodeSuffixPoolOk

`func (o *AuthDisResData) GetRestrictedCodeSuffixPoolOk() (*[]RestrictedCodeSuffixPool, bool)`

GetRestrictedCodeSuffixPoolOk returns a tuple with the RestrictedCodeSuffixPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCodeSuffixPool

`func (o *AuthDisResData) SetRestrictedCodeSuffixPool(v []RestrictedCodeSuffixPool)`

SetRestrictedCodeSuffixPool sets RestrictedCodeSuffixPool field to given value.

### HasRestrictedCodeSuffixPool

`func (o *AuthDisResData) HasRestrictedCodeSuffixPool() bool`

HasRestrictedCodeSuffixPool returns a boolean if a field has been set.

### GetProseAppMasks

`func (o *AuthDisResData) GetProseAppMasks() []string`

GetProseAppMasks returns the ProseAppMasks field if non-nil, zero value otherwise.

### GetProseAppMasksOk

`func (o *AuthDisResData) GetProseAppMasksOk() (*[]string, bool)`

GetProseAppMasksOk returns a tuple with the ProseAppMasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppMasks

`func (o *AuthDisResData) SetProseAppMasks(v []string)`

SetProseAppMasks sets ProseAppMasks field to given value.

### HasProseAppMasks

`func (o *AuthDisResData) HasProseAppMasks() bool`

HasProseAppMasks returns a boolean if a field has been set.

### GetProSeRestrictedMasks

`func (o *AuthDisResData) GetProSeRestrictedMasks() []string`

GetProSeRestrictedMasks returns the ProSeRestrictedMasks field if non-nil, zero value otherwise.

### GetProSeRestrictedMasksOk

`func (o *AuthDisResData) GetProSeRestrictedMasksOk() (*[]string, bool)`

GetProSeRestrictedMasksOk returns a tuple with the ProSeRestrictedMasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProSeRestrictedMasks

`func (o *AuthDisResData) SetProSeRestrictedMasks(v []string)`

SetProSeRestrictedMasks sets ProSeRestrictedMasks field to given value.

### HasProSeRestrictedMasks

`func (o *AuthDisResData) HasProSeRestrictedMasks() bool`

HasProSeRestrictedMasks returns a boolean if a field has been set.

### GetResAppLevelContainer

`func (o *AuthDisResData) GetResAppLevelContainer() string`

GetResAppLevelContainer returns the ResAppLevelContainer field if non-nil, zero value otherwise.

### GetResAppLevelContainerOk

`func (o *AuthDisResData) GetResAppLevelContainerOk() (*string, bool)`

GetResAppLevelContainerOk returns a tuple with the ResAppLevelContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResAppLevelContainer

`func (o *AuthDisResData) SetResAppLevelContainer(v string)`

SetResAppLevelContainer sets ResAppLevelContainer field to given value.

### HasResAppLevelContainer

`func (o *AuthDisResData) HasResAppLevelContainer() bool`

HasResAppLevelContainer returns a boolean if a field has been set.

### GetTargetDataSet

`func (o *AuthDisResData) GetTargetDataSet() []TargetData`

GetTargetDataSet returns the TargetDataSet field if non-nil, zero value otherwise.

### GetTargetDataSetOk

`func (o *AuthDisResData) GetTargetDataSetOk() (*[]TargetData, bool)`

GetTargetDataSetOk returns a tuple with the TargetDataSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDataSet

`func (o *AuthDisResData) SetTargetDataSet(v []TargetData)`

SetTargetDataSet sets TargetDataSet field to given value.

### HasTargetDataSet

`func (o *AuthDisResData) HasTargetDataSet() bool`

HasTargetDataSet returns a boolean if a field has been set.

### GetTargetPduid

`func (o *AuthDisResData) GetTargetPduid() string`

GetTargetPduid returns the TargetPduid field if non-nil, zero value otherwise.

### GetTargetPduidOk

`func (o *AuthDisResData) GetTargetPduidOk() (*string, bool)`

GetTargetPduidOk returns a tuple with the TargetPduid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPduid

`func (o *AuthDisResData) SetTargetPduid(v string)`

SetTargetPduid sets TargetPduid field to given value.

### HasTargetPduid

`func (o *AuthDisResData) HasTargetPduid() bool`

HasTargetPduid returns a boolean if a field has been set.

### GetMetaData

`func (o *AuthDisResData) GetMetaData() string`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AuthDisResData) GetMetaDataOk() (*string, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AuthDisResData) SetMetaData(v string)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *AuthDisResData) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


