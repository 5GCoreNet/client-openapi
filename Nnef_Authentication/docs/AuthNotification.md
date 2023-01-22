# AuthNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**ServiceLevelId** | **string** |  | 
**NotifyCorrId** | **string** |  | 
**AuthMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**NotifType** | [**NotifType**](NotifType.md) |  | 

## Methods

### NewAuthNotification

`func NewAuthNotification(gpsi string, serviceLevelId string, notifyCorrId string, notifType NotifType, ) *AuthNotification`

NewAuthNotification instantiates a new AuthNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthNotificationWithDefaults

`func NewAuthNotificationWithDefaults() *AuthNotification`

NewAuthNotificationWithDefaults instantiates a new AuthNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *AuthNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AuthNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AuthNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetServiceLevelId

`func (o *AuthNotification) GetServiceLevelId() string`

GetServiceLevelId returns the ServiceLevelId field if non-nil, zero value otherwise.

### GetServiceLevelIdOk

`func (o *AuthNotification) GetServiceLevelIdOk() (*string, bool)`

GetServiceLevelIdOk returns a tuple with the ServiceLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLevelId

`func (o *AuthNotification) SetServiceLevelId(v string)`

SetServiceLevelId sets ServiceLevelId field to given value.


### GetNotifyCorrId

`func (o *AuthNotification) GetNotifyCorrId() string`

GetNotifyCorrId returns the NotifyCorrId field if non-nil, zero value otherwise.

### GetNotifyCorrIdOk

`func (o *AuthNotification) GetNotifyCorrIdOk() (*string, bool)`

GetNotifyCorrIdOk returns a tuple with the NotifyCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrId

`func (o *AuthNotification) SetNotifyCorrId(v string)`

SetNotifyCorrId sets NotifyCorrId field to given value.


### GetAuthMsg

`func (o *AuthNotification) GetAuthMsg() RefToBinaryData`

GetAuthMsg returns the AuthMsg field if non-nil, zero value otherwise.

### GetAuthMsgOk

`func (o *AuthNotification) GetAuthMsgOk() (*RefToBinaryData, bool)`

GetAuthMsgOk returns a tuple with the AuthMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMsg

`func (o *AuthNotification) SetAuthMsg(v RefToBinaryData)`

SetAuthMsg sets AuthMsg field to given value.

### HasAuthMsg

`func (o *AuthNotification) HasAuthMsg() bool`

HasAuthMsg returns a boolean if a field has been set.

### GetNotifType

`func (o *AuthNotification) GetNotifType() NotifType`

GetNotifType returns the NotifType field if non-nil, zero value otherwise.

### GetNotifTypeOk

`func (o *AuthNotification) GetNotifTypeOk() (*NotifType, bool)`

GetNotifTypeOk returns a tuple with the NotifType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifType

`func (o *AuthNotification) SetNotifType(v NotifType)`

SetNotifType sets NotifType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


