# NsiInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NrfId** | **string** | String providing an URI formatted according to RFC 3986. | 
**NsiId** | Pointer to **string** | Contains the Identifier of the selected Network Slice instance | [optional] 
**NrfNfMgtUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfAccessTokenUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfOauth2Required** | Pointer to **map[string]bool** | Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \&quot;nnrf-nfm\&quot; or \&quot;nnrf-disc\&quot;  | [optional] 

## Methods

### NewNsiInformation

`func NewNsiInformation(nrfId string, ) *NsiInformation`

NewNsiInformation instantiates a new NsiInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsiInformationWithDefaults

`func NewNsiInformationWithDefaults() *NsiInformation`

NewNsiInformationWithDefaults instantiates a new NsiInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNrfId

`func (o *NsiInformation) GetNrfId() string`

GetNrfId returns the NrfId field if non-nil, zero value otherwise.

### GetNrfIdOk

`func (o *NsiInformation) GetNrfIdOk() (*string, bool)`

GetNrfIdOk returns a tuple with the NrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfId

`func (o *NsiInformation) SetNrfId(v string)`

SetNrfId sets NrfId field to given value.


### GetNsiId

`func (o *NsiInformation) GetNsiId() string`

GetNsiId returns the NsiId field if non-nil, zero value otherwise.

### GetNsiIdOk

`func (o *NsiInformation) GetNsiIdOk() (*string, bool)`

GetNsiIdOk returns a tuple with the NsiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiId

`func (o *NsiInformation) SetNsiId(v string)`

SetNsiId sets NsiId field to given value.

### HasNsiId

`func (o *NsiInformation) HasNsiId() bool`

HasNsiId returns a boolean if a field has been set.

### GetNrfNfMgtUri

`func (o *NsiInformation) GetNrfNfMgtUri() string`

GetNrfNfMgtUri returns the NrfNfMgtUri field if non-nil, zero value otherwise.

### GetNrfNfMgtUriOk

`func (o *NsiInformation) GetNrfNfMgtUriOk() (*string, bool)`

GetNrfNfMgtUriOk returns a tuple with the NrfNfMgtUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfNfMgtUri

`func (o *NsiInformation) SetNrfNfMgtUri(v string)`

SetNrfNfMgtUri sets NrfNfMgtUri field to given value.

### HasNrfNfMgtUri

`func (o *NsiInformation) HasNrfNfMgtUri() bool`

HasNrfNfMgtUri returns a boolean if a field has been set.

### GetNrfAccessTokenUri

`func (o *NsiInformation) GetNrfAccessTokenUri() string`

GetNrfAccessTokenUri returns the NrfAccessTokenUri field if non-nil, zero value otherwise.

### GetNrfAccessTokenUriOk

`func (o *NsiInformation) GetNrfAccessTokenUriOk() (*string, bool)`

GetNrfAccessTokenUriOk returns a tuple with the NrfAccessTokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfAccessTokenUri

`func (o *NsiInformation) SetNrfAccessTokenUri(v string)`

SetNrfAccessTokenUri sets NrfAccessTokenUri field to given value.

### HasNrfAccessTokenUri

`func (o *NsiInformation) HasNrfAccessTokenUri() bool`

HasNrfAccessTokenUri returns a boolean if a field has been set.

### GetNrfOauth2Required

`func (o *NsiInformation) GetNrfOauth2Required() map[string]bool`

GetNrfOauth2Required returns the NrfOauth2Required field if non-nil, zero value otherwise.

### GetNrfOauth2RequiredOk

`func (o *NsiInformation) GetNrfOauth2RequiredOk() (*map[string]bool, bool)`

GetNrfOauth2RequiredOk returns a tuple with the NrfOauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfOauth2Required

`func (o *NsiInformation) SetNrfOauth2Required(v map[string]bool)`

SetNrfOauth2Required sets NrfOauth2Required field to given value.

### HasNrfOauth2Required

`func (o *NsiInformation) HasNrfOauth2Required() bool`

HasNrfOauth2Required returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


