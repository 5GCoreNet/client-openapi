# SliceAuthConfirmationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**EapMessage** | **NullableString** | contains an EAP packet | 

## Methods

### NewSliceAuthConfirmationData

`func NewSliceAuthConfirmationData(gpsi string, snssai Snssai, eapMessage NullableString, ) *SliceAuthConfirmationData`

NewSliceAuthConfirmationData instantiates a new SliceAuthConfirmationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceAuthConfirmationDataWithDefaults

`func NewSliceAuthConfirmationDataWithDefaults() *SliceAuthConfirmationData`

NewSliceAuthConfirmationDataWithDefaults instantiates a new SliceAuthConfirmationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *SliceAuthConfirmationData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SliceAuthConfirmationData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SliceAuthConfirmationData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetSnssai

`func (o *SliceAuthConfirmationData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SliceAuthConfirmationData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SliceAuthConfirmationData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetEapMessage

`func (o *SliceAuthConfirmationData) GetEapMessage() string`

GetEapMessage returns the EapMessage field if non-nil, zero value otherwise.

### GetEapMessageOk

`func (o *SliceAuthConfirmationData) GetEapMessageOk() (*string, bool)`

GetEapMessageOk returns a tuple with the EapMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapMessage

`func (o *SliceAuthConfirmationData) SetEapMessage(v string)`

SetEapMessage sets EapMessage field to given value.


### SetEapMessageNil

`func (o *SliceAuthConfirmationData) SetEapMessageNil(b bool)`

 SetEapMessageNil sets the value for EapMessage to be an explicit nil

### UnsetEapMessage
`func (o *SliceAuthConfirmationData) UnsetEapMessage()`

UnsetEapMessage ensures that no value is present for EapMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


