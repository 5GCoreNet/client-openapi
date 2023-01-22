# ReauthRevokeNotify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**ServiceLevelId** | **string** |  | 
**NotifyCorrId** | Pointer to **string** |  | [optional] 
**AuthMsg** | Pointer to **string** |  | [optional] 
**NotifyType** | [**NotifyType**](NotifyType.md) |  | 
**IpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 

## Methods

### NewReauthRevokeNotify

`func NewReauthRevokeNotify(gpsi string, serviceLevelId string, notifyType NotifyType, ) *ReauthRevokeNotify`

NewReauthRevokeNotify instantiates a new ReauthRevokeNotify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReauthRevokeNotifyWithDefaults

`func NewReauthRevokeNotifyWithDefaults() *ReauthRevokeNotify`

NewReauthRevokeNotifyWithDefaults instantiates a new ReauthRevokeNotify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *ReauthRevokeNotify) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *ReauthRevokeNotify) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *ReauthRevokeNotify) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetServiceLevelId

`func (o *ReauthRevokeNotify) GetServiceLevelId() string`

GetServiceLevelId returns the ServiceLevelId field if non-nil, zero value otherwise.

### GetServiceLevelIdOk

`func (o *ReauthRevokeNotify) GetServiceLevelIdOk() (*string, bool)`

GetServiceLevelIdOk returns a tuple with the ServiceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelId

`func (o *ReauthRevokeNotify) SetServiceLevelId(v string)`

SetServiceLevelId sets ServiceLevelId field to given value.


### GetNotifyCorrId

`func (o *ReauthRevokeNotify) GetNotifyCorrId() string`

GetNotifyCorrId returns the NotifyCorrId field if non-nil, zero value otherwise.

### GetNotifyCorrIdOk

`func (o *ReauthRevokeNotify) GetNotifyCorrIdOk() (*string, bool)`

GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrId

`func (o *ReauthRevokeNotify) SetNotifyCorrId(v string)`

SetNotifyCorrId sets NotifyCorrId field to given value.

### HasNotifyCorrId

`func (o *ReauthRevokeNotify) HasNotifyCorrId() bool`

HasNotifyCorrId returns a boolean if a field has been set.

### GetAuthMsg

`func (o *ReauthRevokeNotify) GetAuthMsg() string`

GetAuthMsg returns the AuthMsg field if non-nil, zero value otherwise.

### GetAuthMsgOk

`func (o *ReauthRevokeNotify) GetAuthMsgOk() (*string, bool)`

GetAuthMsgOk returns a tuple with the AuthMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMsg

`func (o *ReauthRevokeNotify) SetAuthMsg(v string)`

SetAuthMsg sets AuthMsg field to given value.

### HasAuthMsg

`func (o *ReauthRevokeNotify) HasAuthMsg() bool`

HasAuthMsg returns a boolean if a field has been set.

### GetNotifyType

`func (o *ReauthRevokeNotify) GetNotifyType() NotifyType`

GetNotifyType returns the NotifyType field if non-nil, zero value otherwise.

### GetNotifyTypeOk

`func (o *ReauthRevokeNotify) GetNotifyTypeOk() (*NotifyType, bool)`

GetNotifyTypeOk returns a tuple with the NotifyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyType

`func (o *ReauthRevokeNotify) SetNotifyType(v NotifyType)`

SetNotifyType sets NotifyType field to given value.


### GetIpAddr

`func (o *ReauthRevokeNotify) GetIpAddr() IpAddr`

GetIpAddr returns the IpAddr field if non-nil, zero value otherwise.

### GetIpAddrOk

`func (o *ReauthRevokeNotify) GetIpAddrOk() (*IpAddr, bool)`

GetIpAddrOk returns a tuple with the IpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddr

`func (o *ReauthRevokeNotify) SetIpAddr(v IpAddr)`

SetIpAddr sets IpAddr field to given value.

### HasIpAddr

`func (o *ReauthRevokeNotify) HasIpAddr() bool`

HasIpAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


