# SliceAuthInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**EapIdRsp** | **NullableString** | contains an EAP packet | 
**AmfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**ReauthNotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**RevocNotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewSliceAuthInfo

`func NewSliceAuthInfo(gpsi string, snssai Snssai, eapIdRsp NullableString, ) *SliceAuthInfo`

NewSliceAuthInfo instantiates a new SliceAuthInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceAuthInfoWithDefaults

`func NewSliceAuthInfoWithDefaults() *SliceAuthInfo`

NewSliceAuthInfoWithDefaults instantiates a new SliceAuthInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *SliceAuthInfo) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SliceAuthInfo) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SliceAuthInfo) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetSnssai

`func (o *SliceAuthInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SliceAuthInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SliceAuthInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetEapIdRsp

`func (o *SliceAuthInfo) GetEapIdRsp() string`

GetEapIdRsp returns the EapIdRsp field if non-nil, zero value otherwise.

### GetEapIdRspOk

`func (o *SliceAuthInfo) GetEapIdRspOk() (*string, bool)`

GetEapIdRspOk returns a tuple with the EapIdRsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapIdRsp

`func (o *SliceAuthInfo) SetEapIdRsp(v string)`

SetEapIdRsp sets EapIdRsp field to given value.


### SetEapIdRspNil

`func (o *SliceAuthInfo) SetEapIdRspNil(b bool)`

 SetEapIdRspNil sets the value for EapIdRsp to be an explicit nil

### UnsetEapIdRsp
`func (o *SliceAuthInfo) UnsetEapIdRsp()`

UnsetEapIdRsp ensures that no value is present for EapIdRsp, not even an explicit nil
### GetAmfInstanceId

`func (o *SliceAuthInfo) GetAmfInstanceId() string`

GetAmfInstanceId returns the AmfInstanceId field if non-nil, zero value otherwise.

### GetAmfInstanceIdOk

`func (o *SliceAuthInfo) GetAmfInstanceIdOk() (*string, bool)`

GetAmfInstanceIdOk returns a tuple with the AmfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInstanceId

`func (o *SliceAuthInfo) SetAmfInstanceId(v string)`

SetAmfInstanceId sets AmfInstanceId field to given value.

### HasAmfInstanceId

`func (o *SliceAuthInfo) HasAmfInstanceId() bool`

HasAmfInstanceId returns a boolean if a field has been set.

### GetReauthNotifUri

`func (o *SliceAuthInfo) GetReauthNotifUri() string`

GetReauthNotifUri returns the ReauthNotifUri field if non-nil, zero value otherwise.

### GetReauthNotifUriOk

`func (o *SliceAuthInfo) GetReauthNotifUriOk() (*string, bool)`

GetReauthNotifUriOk returns a tuple with the ReauthNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthNotifUri

`func (o *SliceAuthInfo) SetReauthNotifUri(v string)`

SetReauthNotifUri sets ReauthNotifUri field to given value.

### HasReauthNotifUri

`func (o *SliceAuthInfo) HasReauthNotifUri() bool`

HasReauthNotifUri returns a boolean if a field has been set.

### GetRevocNotifUri

`func (o *SliceAuthInfo) GetRevocNotifUri() string`

GetRevocNotifUri returns the RevocNotifUri field if non-nil, zero value otherwise.

### GetRevocNotifUriOk

`func (o *SliceAuthInfo) GetRevocNotifUriOk() (*string, bool)`

GetRevocNotifUriOk returns a tuple with the RevocNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocNotifUri

`func (o *SliceAuthInfo) SetRevocNotifUri(v string)`

SetRevocNotifUri sets RevocNotifUri field to given value.

### HasRevocNotifUri

`func (o *SliceAuthInfo) HasRevocNotifUri() bool`

HasRevocNotifUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


