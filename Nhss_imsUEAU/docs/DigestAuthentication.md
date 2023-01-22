# DigestAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DigestRealm** | **string** |  | 
**DigestAlgorithm** | [**SipDigestAlgorithm**](SipDigestAlgorithm.md) |  | 
**DigestQop** | [**SipDigestQop**](SipDigestQop.md) |  | 
**Ha1** | **string** |  | 

## Methods

### NewDigestAuthentication

`func NewDigestAuthentication(digestRealm string, digestAlgorithm SipDigestAlgorithm, digestQop SipDigestQop, ha1 string, ) *DigestAuthentication`

NewDigestAuthentication instantiates a new DigestAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDigestAuthenticationWithDefaults

`func NewDigestAuthenticationWithDefaults() *DigestAuthentication`

NewDigestAuthenticationWithDefaults instantiates a new DigestAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigestRealm

`func (o *DigestAuthentication) GetDigestRealm() string`

GetDigestRealm returns the DigestRealm field if non-nil, zero value otherwise.

### GetDigestRealmOk

`func (o *DigestAuthentication) GetDigestRealmOk() (*string, bool)`

GetDigestRealmOk returns a tuple with the DigestRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestRealm

`func (o *DigestAuthentication) SetDigestRealm(v string)`

SetDigestRealm sets DigestRealm field to given value.


### GetDigestAlgorithm

`func (o *DigestAuthentication) GetDigestAlgorithm() SipDigestAlgorithm`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *DigestAuthentication) GetDigestAlgorithmOk() (*SipDigestAlgorithm, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *DigestAuthentication) SetDigestAlgorithm(v SipDigestAlgorithm)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.


### GetDigestQop

`func (o *DigestAuthentication) GetDigestQop() SipDigestQop`

GetDigestQop returns the DigestQop field if non-nil, zero value otherwise.

### GetDigestQopOk

`func (o *DigestAuthentication) GetDigestQopOk() (*SipDigestQop, bool)`

GetDigestQopOk returns a tuple with the DigestQop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestQop

`func (o *DigestAuthentication) SetDigestQop(v SipDigestQop)`

SetDigestQop sets DigestQop field to given value.


### GetHa1

`func (o *DigestAuthentication) GetHa1() string`

GetHa1 returns the Ha1 field if non-nil, zero value otherwise.

### GetHa1Ok

`func (o *DigestAuthentication) GetHa1Ok() (*string, bool)`

GetHa1Ok returns a tuple with the Ha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa1

`func (o *DigestAuthentication) SetHa1(v string)`

SetHa1 sets Ha1 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


