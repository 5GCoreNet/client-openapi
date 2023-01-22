# UAVAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**ServiceLevelId** | Pointer to **string** |  | [optional] 
**AuthMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**AuthResult** | Pointer to [**AuthResult**](AuthResult.md) |  | [optional] 
**NotifyCorrId** | Pointer to **string** |  | [optional] 

## Methods

### NewUAVAuthResponse

`func NewUAVAuthResponse(gpsi string, ) *UAVAuthResponse`

NewUAVAuthResponse instantiates a new UAVAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUAVAuthResponseWithDefaults

`func NewUAVAuthResponseWithDefaults() *UAVAuthResponse`

NewUAVAuthResponseWithDefaults instantiates a new UAVAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *UAVAuthResponse) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UAVAuthResponse) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UAVAuthResponse) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetServiceLevelId

`func (o *UAVAuthResponse) GetServiceLevelId() string`

GetServiceLevelId returns the ServiceLevelId field if non-nil, zero value otherwise.

### GetServiceLevelIdOk

`func (o *UAVAuthResponse) GetServiceLevelIdOk() (*string, bool)`

GetServiceLevelIdOk returns a tuple with the ServiceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelId

`func (o *UAVAuthResponse) SetServiceLevelId(v string)`

SetServiceLevelId sets ServiceLevelId field to given value.

### HasServiceLevelId

`func (o *UAVAuthResponse) HasServiceLevelId() bool`

HasServiceLevelId returns a boolean if a field has been set.

### GetAuthMsg

`func (o *UAVAuthResponse) GetAuthMsg() RefToBinaryData`

GetAuthMsg returns the AuthMsg field if non-nil, zero value otherwise.

### GetAuthMsgOk

`func (o *UAVAuthResponse) GetAuthMsgOk() (*RefToBinaryData, bool)`

GetAuthMsgOk returns a tuple with the AuthMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMsg

`func (o *UAVAuthResponse) SetAuthMsg(v RefToBinaryData)`

SetAuthMsg sets AuthMsg field to given value.

### HasAuthMsg

`func (o *UAVAuthResponse) HasAuthMsg() bool`

HasAuthMsg returns a boolean if a field has been set.

### GetAuthResult

`func (o *UAVAuthResponse) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *UAVAuthResponse) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *UAVAuthResponse) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *UAVAuthResponse) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.

### GetNotifyCorrId

`func (o *UAVAuthResponse) GetNotifyCorrId() string`

GetNotifyCorrId returns the NotifyCorrId field if non-nil, zero value otherwise.

### GetNotifyCorrIdOk

`func (o *UAVAuthResponse) GetNotifyCorrIdOk() (*string, bool)`

GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrId

`func (o *UAVAuthResponse) SetNotifyCorrId(v string)`

SetNotifyCorrId sets NotifyCorrId field to given value.

### HasNotifyCorrId

`func (o *UAVAuthResponse) HasNotifyCorrId() bool`

HasNotifyCorrId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


