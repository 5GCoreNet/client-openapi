# PrivateIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateIdentity** | **string** | String containing a Private User Identity or a Private Service Identity | 
**PrivateIdentityType** | [**PrivateIdentityType**](PrivateIdentityType.md) |  | 

## Methods

### NewPrivateIdentity

`func NewPrivateIdentity(privateIdentity string, privateIdentityType PrivateIdentityType, ) *PrivateIdentity`

NewPrivateIdentity instantiates a new PrivateIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateIdentityWithDefaults

`func NewPrivateIdentityWithDefaults() *PrivateIdentity`

NewPrivateIdentityWithDefaults instantiates a new PrivateIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateIdentity

`func (o *PrivateIdentity) GetPrivateIdentity() string`

GetPrivateIdentity returns the PrivateIdentity field if non-nil, zero value otherwise.

### GetPrivateIdentityOk

`func (o *PrivateIdentity) GetPrivateIdentityOk() (*string, bool)`

GetPrivateIdentityOk returns a tuple with the PrivateIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIdentity

`func (o *PrivateIdentity) SetPrivateIdentity(v string)`

SetPrivateIdentity sets PrivateIdentity field to given value.


### GetPrivateIdentityType

`func (o *PrivateIdentity) GetPrivateIdentityType() PrivateIdentityType`

GetPrivateIdentityType returns the PrivateIdentityType field if non-nil, zero value otherwise.

### GetPrivateIdentityTypeOk

`func (o *PrivateIdentity) GetPrivateIdentityTypeOk() (*PrivateIdentityType, bool)`

GetPrivateIdentityTypeOk returns a tuple with the PrivateIdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIdentityType

`func (o *PrivateIdentity) SetPrivateIdentityType(v PrivateIdentityType)`

SetPrivateIdentityType sets PrivateIdentityType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


