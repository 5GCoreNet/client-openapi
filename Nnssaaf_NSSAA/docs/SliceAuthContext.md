# SliceAuthContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**AuthCtxId** | **string** | contains the resource ID of slice authentication context | 
**EapMessage** | **NullableString** | contains an EAP packet | 

## Methods

### NewSliceAuthContext

`func NewSliceAuthContext(gpsi string, snssai Snssai, authCtxId string, eapMessage NullableString, ) *SliceAuthContext`

NewSliceAuthContext instantiates a new SliceAuthContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceAuthContextWithDefaults

`func NewSliceAuthContextWithDefaults() *SliceAuthContext`

NewSliceAuthContextWithDefaults instantiates a new SliceAuthContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *SliceAuthContext) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SliceAuthContext) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SliceAuthContext) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetSnssai

`func (o *SliceAuthContext) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SliceAuthContext) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SliceAuthContext) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetAuthCtxId

`func (o *SliceAuthContext) GetAuthCtxId() string`

GetAuthCtxId returns the AuthCtxId field if non-nil, zero value otherwise.

### GetAuthCtxIdOk

`func (o *SliceAuthContext) GetAuthCtxIdOk() (*string, bool)`

GetAuthCtxIdOk returns a tuple with the AuthCtxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCtxId

`func (o *SliceAuthContext) SetAuthCtxId(v string)`

SetAuthCtxId sets AuthCtxId field to given value.


### GetEapMessage

`func (o *SliceAuthContext) GetEapMessage() string`

GetEapMessage returns the EapMessage field if non-nil, zero value otherwise.

### GetEapMessageOk

`func (o *SliceAuthContext) GetEapMessageOk() (*string, bool)`

GetEapMessageOk returns a tuple with the EapMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapMessage

`func (o *SliceAuthContext) SetEapMessage(v string)`

SetEapMessage sets EapMessage field to given value.


### SetEapMessageNil

`func (o *SliceAuthContext) SetEapMessageNil(b bool)`

 SetEapMessageNil sets the value for EapMessage to be an explicit nil

### UnsetEapMessage
`func (o *SliceAuthContext) UnsetEapMessage()`

UnsetEapMessage ensures that no value is present for EapMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


