# PublicIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImsPublicId** | **string** | String containing an IMS Public Identity in SIP URI format or TEL URI format | 
**IdentityType** | [**IdentityType**](IdentityType.md) |  | 
**IrsIsDefault** | Pointer to **bool** |  | [optional] 
**AliasGroupId** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicIdentity

`func NewPublicIdentity(imsPublicId string, identityType IdentityType, ) *PublicIdentity`

NewPublicIdentity instantiates a new PublicIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIdentityWithDefaults

`func NewPublicIdentityWithDefaults() *PublicIdentity`

NewPublicIdentityWithDefaults instantiates a new PublicIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsPublicId

`func (o *PublicIdentity) GetImsPublicId() string`

GetImsPublicId returns the ImsPublicId field if non-nil, zero value otherwise.

### GetImsPublicIdOk

`func (o *PublicIdentity) GetImsPublicIdOk() (*string, bool)`

GetImsPublicIdOk returns a tuple with the ImsPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsPublicId

`func (o *PublicIdentity) SetImsPublicId(v string)`

SetImsPublicId sets ImsPublicId field to given value.


### GetIdentityType

`func (o *PublicIdentity) GetIdentityType() IdentityType`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *PublicIdentity) GetIdentityTypeOk() (*IdentityType, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *PublicIdentity) SetIdentityType(v IdentityType)`

SetIdentityType sets IdentityType field to given value.


### GetIrsIsDefault

`func (o *PublicIdentity) GetIrsIsDefault() bool`

GetIrsIsDefault returns the IrsIsDefault field if non-nil, zero value otherwise.

### GetIrsIsDefaultOk

`func (o *PublicIdentity) GetIrsIsDefaultOk() (*bool, bool)`

GetIrsIsDefaultOk returns a tuple with the IrsIsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrsIsDefault

`func (o *PublicIdentity) SetIrsIsDefault(v bool)`

SetIrsIsDefault sets IrsIsDefault field to given value.

### HasIrsIsDefault

`func (o *PublicIdentity) HasIrsIsDefault() bool`

HasIrsIsDefault returns a boolean if a field has been set.

### GetAliasGroupId

`func (o *PublicIdentity) GetAliasGroupId() string`

GetAliasGroupId returns the AliasGroupId field if non-nil, zero value otherwise.

### GetAliasGroupIdOk

`func (o *PublicIdentity) GetAliasGroupIdOk() (*string, bool)`

GetAliasGroupIdOk returns a tuple with the AliasGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasGroupId

`func (o *PublicIdentity) SetAliasGroupId(v string)`

SetAliasGroupId sets AliasGroupId field to given value.

### HasAliasGroupId

`func (o *PublicIdentity) HasAliasGroupId() bool`

HasAliasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


