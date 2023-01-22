# ImsServiceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIdentifierList** | [**[]PublicIdentifier**](PublicIdentifier.md) |  | 
**Ifcs** | Pointer to [**Ifcs**](Ifcs.md) |  | [optional] 
**CnServiceAuthorization** | Pointer to [**CoreNetworkServiceAuthorization**](CoreNetworkServiceAuthorization.md) |  | [optional] 

## Methods

### NewImsServiceProfile

`func NewImsServiceProfile(publicIdentifierList []PublicIdentifier, ) *ImsServiceProfile`

NewImsServiceProfile instantiates a new ImsServiceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImsServiceProfileWithDefaults

`func NewImsServiceProfileWithDefaults() *ImsServiceProfile`

NewImsServiceProfileWithDefaults instantiates a new ImsServiceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIdentifierList

`func (o *ImsServiceProfile) GetPublicIdentifierList() []PublicIdentifier`

GetPublicIdentifierList returns the PublicIdentifierList field if non-nil, zero value otherwise.

### GetPublicIdentifierListOk

`func (o *ImsServiceProfile) GetPublicIdentifierListOk() (*[]PublicIdentifier, bool)`

GetPublicIdentifierListOk returns a tuple with the PublicIdentifierList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIdentifierList

`func (o *ImsServiceProfile) SetPublicIdentifierList(v []PublicIdentifier)`

SetPublicIdentifierList sets PublicIdentifierList field to given value.


### GetIfcs

`func (o *ImsServiceProfile) GetIfcs() Ifcs`

GetIfcs returns the Ifcs field if non-nil, zero value otherwise.

### GetIfcsOk

`func (o *ImsServiceProfile) GetIfcsOk() (*Ifcs, bool)`

GetIfcsOk returns a tuple with the Ifcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfcs

`func (o *ImsServiceProfile) SetIfcs(v Ifcs)`

SetIfcs sets Ifcs field to given value.

### HasIfcs

`func (o *ImsServiceProfile) HasIfcs() bool`

HasIfcs returns a boolean if a field has been set.

### GetCnServiceAuthorization

`func (o *ImsServiceProfile) GetCnServiceAuthorization() CoreNetworkServiceAuthorization`

GetCnServiceAuthorization returns the CnServiceAuthorization field if non-nil, zero value otherwise.

### GetCnServiceAuthorizationOk

`func (o *ImsServiceProfile) GetCnServiceAuthorizationOk() (*CoreNetworkServiceAuthorization, bool)`

GetCnServiceAuthorizationOk returns a tuple with the CnServiceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnServiceAuthorization

`func (o *ImsServiceProfile) SetCnServiceAuthorization(v CoreNetworkServiceAuthorization)`

SetCnServiceAuthorization sets CnServiceAuthorization field to given value.

### HasCnServiceAuthorization

`func (o *ImsServiceProfile) HasCnServiceAuthorization() bool`

HasCnServiceAuthorization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


