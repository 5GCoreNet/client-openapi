# AuthorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationData** | [**[]UserIdentifier**](UserIdentifier.md) |  | 
**AllowedDnnList** | Pointer to [**[]AccessAndMobilitySubscriptionDataSubscribedDnnListInner**](AccessAndMobilitySubscriptionDataSubscribedDnnListInner.md) |  | [optional] 
**AllowedSnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**AllowedMtcProviders** | Pointer to [**[]MtcProvider**](MtcProvider.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewAuthorizationData

`func NewAuthorizationData(authorizationData []UserIdentifier, ) *AuthorizationData`

NewAuthorizationData instantiates a new AuthorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDataWithDefaults

`func NewAuthorizationDataWithDefaults() *AuthorizationData`

NewAuthorizationDataWithDefaults instantiates a new AuthorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationData

`func (o *AuthorizationData) GetAuthorizationData() []UserIdentifier`

GetAuthorizationData returns the AuthorizationData field if non-nil, zero value otherwise.

### GetAuthorizationDataOk

`func (o *AuthorizationData) GetAuthorizationDataOk() (*[]UserIdentifier, bool)`

GetAuthorizationDataOk returns a tuple with the AuthorizationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationData

`func (o *AuthorizationData) SetAuthorizationData(v []UserIdentifier)`

SetAuthorizationData sets AuthorizationData field to given value.


### GetAllowedDnnList

`func (o *AuthorizationData) GetAllowedDnnList() []AccessAndMobilitySubscriptionDataSubscribedDnnListInner`

GetAllowedDnnList returns the AllowedDnnList field if non-nil, zero value otherwise.

### GetAllowedDnnListOk

`func (o *AuthorizationData) GetAllowedDnnListOk() (*[]AccessAndMobilitySubscriptionDataSubscribedDnnListInner, bool)`

GetAllowedDnnListOk returns a tuple with the AllowedDnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDnnList

`func (o *AuthorizationData) SetAllowedDnnList(v []AccessAndMobilitySubscriptionDataSubscribedDnnListInner)`

SetAllowedDnnList sets AllowedDnnList field to given value.

### HasAllowedDnnList

`func (o *AuthorizationData) HasAllowedDnnList() bool`

HasAllowedDnnList returns a boolean if a field has been set.

### GetAllowedSnssaiList

`func (o *AuthorizationData) GetAllowedSnssaiList() []Snssai`

GetAllowedSnssaiList returns the AllowedSnssaiList field if non-nil, zero value otherwise.

### GetAllowedSnssaiListOk

`func (o *AuthorizationData) GetAllowedSnssaiListOk() (*[]Snssai, bool)`

GetAllowedSnssaiListOk returns a tuple with the AllowedSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnssaiList

`func (o *AuthorizationData) SetAllowedSnssaiList(v []Snssai)`

SetAllowedSnssaiList sets AllowedSnssaiList field to given value.

### HasAllowedSnssaiList

`func (o *AuthorizationData) HasAllowedSnssaiList() bool`

HasAllowedSnssaiList returns a boolean if a field has been set.

### GetAllowedMtcProviders

`func (o *AuthorizationData) GetAllowedMtcProviders() []MtcProvider`

GetAllowedMtcProviders returns the AllowedMtcProviders field if non-nil, zero value otherwise.

### GetAllowedMtcProvidersOk

`func (o *AuthorizationData) GetAllowedMtcProvidersOk() (*[]MtcProvider, bool)`

GetAllowedMtcProvidersOk returns a tuple with the AllowedMtcProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMtcProviders

`func (o *AuthorizationData) SetAllowedMtcProviders(v []MtcProvider)`

SetAllowedMtcProviders sets AllowedMtcProviders field to given value.

### HasAllowedMtcProviders

`func (o *AuthorizationData) HasAllowedMtcProviders() bool`

HasAllowedMtcProviders returns a boolean if a field has been set.

### GetValidityTime

`func (o *AuthorizationData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AuthorizationData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AuthorizationData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *AuthorizationData) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


