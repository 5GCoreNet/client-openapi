# SipAuthenticationInfoResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impi** | **string** | IMS Private Identity of the UE | 
**Var3gAkaAvs** | Pointer to [**[]Model3GAkaAv**](Model3GAkaAv.md) |  | [optional] 
**DigestAuth** | Pointer to [**DigestAuthentication**](DigestAuthentication.md) |  | [optional] 
**LineIdentifierList** | Pointer to **[]string** |  | [optional] 
**IpAddress** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 

## Methods

### NewSipAuthenticationInfoResult

`func NewSipAuthenticationInfoResult(impi string, ) *SipAuthenticationInfoResult`

NewSipAuthenticationInfoResult instantiates a new SipAuthenticationInfoResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipAuthenticationInfoResultWithDefaults

`func NewSipAuthenticationInfoResultWithDefaults() *SipAuthenticationInfoResult`

NewSipAuthenticationInfoResultWithDefaults instantiates a new SipAuthenticationInfoResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpi

`func (o *SipAuthenticationInfoResult) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *SipAuthenticationInfoResult) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *SipAuthenticationInfoResult) SetImpi(v string)`

SetImpi sets Impi field to given value.


### GetVar3gAkaAvs

`func (o *SipAuthenticationInfoResult) GetVar3gAkaAvs() []Model3GAkaAv`

GetVar3gAkaAvs returns the Var3gAkaAvs field if non-nil, zero value otherwise.

### GetVar3gAkaAvsOk

`func (o *SipAuthenticationInfoResult) GetVar3gAkaAvsOk() (*[]Model3GAkaAv, bool)`

GetVar3gAkaAvsOk returns a tuple with the Var3gAkaAvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3gAkaAvs

`func (o *SipAuthenticationInfoResult) SetVar3gAkaAvs(v []Model3GAkaAv)`

SetVar3gAkaAvs sets Var3gAkaAvs field to given value.

### HasVar3gAkaAvs

`func (o *SipAuthenticationInfoResult) HasVar3gAkaAvs() bool`

HasVar3gAkaAvs returns a boolean if a field has been set.

### GetDigestAuth

`func (o *SipAuthenticationInfoResult) GetDigestAuth() DigestAuthentication`

GetDigestAuth returns the DigestAuth field if non-nil, zero value otherwise.

### GetDigestAuthOk

`func (o *SipAuthenticationInfoResult) GetDigestAuthOk() (*DigestAuthentication, bool)`

GetDigestAuthOk returns a tuple with the DigestAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAuth

`func (o *SipAuthenticationInfoResult) SetDigestAuth(v DigestAuthentication)`

SetDigestAuth sets DigestAuth field to given value.

### HasDigestAuth

`func (o *SipAuthenticationInfoResult) HasDigestAuth() bool`

HasDigestAuth returns a boolean if a field has been set.

### GetLineIdentifierList

`func (o *SipAuthenticationInfoResult) GetLineIdentifierList() []string`

GetLineIdentifierList returns the LineIdentifierList field if non-nil, zero value otherwise.

### GetLineIdentifierListOk

`func (o *SipAuthenticationInfoResult) GetLineIdentifierListOk() (*[]string, bool)`

GetLineIdentifierListOk returns a tuple with the LineIdentifierList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineIdentifierList

`func (o *SipAuthenticationInfoResult) SetLineIdentifierList(v []string)`

SetLineIdentifierList sets LineIdentifierList field to given value.

### HasLineIdentifierList

`func (o *SipAuthenticationInfoResult) HasLineIdentifierList() bool`

HasLineIdentifierList returns a boolean if a field has been set.

### GetIpAddress

`func (o *SipAuthenticationInfoResult) GetIpAddress() IpAddr`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *SipAuthenticationInfoResult) GetIpAddressOk() (*IpAddr, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *SipAuthenticationInfoResult) SetIpAddress(v IpAddr)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *SipAuthenticationInfoResult) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


