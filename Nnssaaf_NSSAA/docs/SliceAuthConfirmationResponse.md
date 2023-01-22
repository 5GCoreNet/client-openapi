# SliceAuthConfirmationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**EapMessage** | **NullableString** | contains an EAP packet | 
**AuthResult** | Pointer to [**AuthStatus**](AuthStatus.md) |  | [optional] 

## Methods

### NewSliceAuthConfirmationResponse

`func NewSliceAuthConfirmationResponse(gpsi string, snssai Snssai, eapMessage NullableString, ) *SliceAuthConfirmationResponse`

NewSliceAuthConfirmationResponse instantiates a new SliceAuthConfirmationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceAuthConfirmationResponseWithDefaults

`func NewSliceAuthConfirmationResponseWithDefaults() *SliceAuthConfirmationResponse`

NewSliceAuthConfirmationResponseWithDefaults instantiates a new SliceAuthConfirmationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *SliceAuthConfirmationResponse) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SliceAuthConfirmationResponse) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SliceAuthConfirmationResponse) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetSnssai

`func (o *SliceAuthConfirmationResponse) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SliceAuthConfirmationResponse) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SliceAuthConfirmationResponse) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetEapMessage

`func (o *SliceAuthConfirmationResponse) GetEapMessage() string`

GetEapMessage returns the EapMessage field if non-nil, zero value otherwise.

### GetEapMessageOk

`func (o *SliceAuthConfirmationResponse) GetEapMessageOk() (*string, bool)`

GetEapMessageOk returns a tuple with the EapMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapMessage

`func (o *SliceAuthConfirmationResponse) SetEapMessage(v string)`

SetEapMessage sets EapMessage field to given value.


### SetEapMessageNil

`func (o *SliceAuthConfirmationResponse) SetEapMessageNil(b bool)`

 SetEapMessageNil sets the value for EapMessage to be an explicit nil

### UnsetEapMessage
`func (o *SliceAuthConfirmationResponse) UnsetEapMessage()`

UnsetEapMessage ensures that no value is present for EapMessage, not even an explicit nil
### GetAuthResult

`func (o *SliceAuthConfirmationResponse) GetAuthResult() AuthStatus`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *SliceAuthConfirmationResponse) GetAuthResultOk() (*AuthStatus, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *SliceAuthConfirmationResponse) SetAuthResult(v AuthStatus)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *SliceAuthConfirmationResponse) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


