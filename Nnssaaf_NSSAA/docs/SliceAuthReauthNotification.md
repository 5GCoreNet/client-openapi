# SliceAuthReauthNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifType** | [**SliceAuthNotificationType**](SliceAuthNotificationType.md) |  | 
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 

## Methods

### NewSliceAuthReauthNotification

`func NewSliceAuthReauthNotification(notifType SliceAuthNotificationType, gpsi string, snssai Snssai, ) *SliceAuthReauthNotification`

NewSliceAuthReauthNotification instantiates a new SliceAuthReauthNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceAuthReauthNotificationWithDefaults

`func NewSliceAuthReauthNotificationWithDefaults() *SliceAuthReauthNotification`

NewSliceAuthReauthNotificationWithDefaults instantiates a new SliceAuthReauthNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifType

`func (o *SliceAuthReauthNotification) GetNotifType() SliceAuthNotificationType`

GetNotifType returns the NotifType field if non-nil, zero value otherwise.

### GetNotifTypeOk

`func (o *SliceAuthReauthNotification) GetNotifTypeOk() (*SliceAuthNotificationType, bool)`

GetNotifTypeOk returns a tuple with the NotifType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifType

`func (o *SliceAuthReauthNotification) SetNotifType(v SliceAuthNotificationType)`

SetNotifType sets NotifType field to given value.


### GetGpsi

`func (o *SliceAuthReauthNotification) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SliceAuthReauthNotification) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SliceAuthReauthNotification) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetSnssai

`func (o *SliceAuthReauthNotification) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SliceAuthReauthNotification) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SliceAuthReauthNotification) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetSupi

`func (o *SliceAuthReauthNotification) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SliceAuthReauthNotification) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SliceAuthReauthNotification) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *SliceAuthReauthNotification) HasSupi() bool`

HasSupi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


