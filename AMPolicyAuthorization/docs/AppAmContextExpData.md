# AppAmContextExpData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**EvSubscs** | Pointer to [**AmEventsSubscData**](AmEventsSubscData.md) |  | [optional] 
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**HighThruInd** | Pointer to **bool** |  | [optional] 
**CovReqs** | Pointer to [**[]GeographicalArea**](GeographicalArea.md) |  | [optional] 
**PolicyDuration** | Pointer to **NullableInt32** | Unsigned integer identifying a period of time in units of seconds with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the AF to request the NEF to send a test notification as defined in clause 5.2.5.3 of 3GPP TS 29.122. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 

## Methods

### NewAppAmContextExpData

`func NewAppAmContextExpData(gpsi string, ) *AppAmContextExpData`

NewAppAmContextExpData instantiates a new AppAmContextExpData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAmContextExpDataWithDefaults

`func NewAppAmContextExpDataWithDefaults() *AppAmContextExpData`

NewAppAmContextExpDataWithDefaults instantiates a new AppAmContextExpData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AppAmContextExpData) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AppAmContextExpData) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AppAmContextExpData) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AppAmContextExpData) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetEvSubscs

`func (o *AppAmContextExpData) GetEvSubscs() AmEventsSubscData`

GetEvSubscs returns the EvSubscs field if non-nil, zero value otherwise.

### GetEvSubscsOk

`func (o *AppAmContextExpData) GetEvSubscsOk() (*AmEventsSubscData, bool)`

GetEvSubscsOk returns a tuple with the EvSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubscs

`func (o *AppAmContextExpData) SetEvSubscs(v AmEventsSubscData)`

SetEvSubscs sets EvSubscs field to given value.

### HasEvSubscs

`func (o *AppAmContextExpData) HasEvSubscs() bool`

HasEvSubscs returns a boolean if a field has been set.

### GetGpsi

`func (o *AppAmContextExpData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AppAmContextExpData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AppAmContextExpData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetHighThruInd

`func (o *AppAmContextExpData) GetHighThruInd() bool`

GetHighThruInd returns the HighThruInd field if non-nil, zero value otherwise.

### GetHighThruIndOk

`func (o *AppAmContextExpData) GetHighThruIndOk() (*bool, bool)`

GetHighThruIndOk returns a tuple with the HighThruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThruInd

`func (o *AppAmContextExpData) SetHighThruInd(v bool)`

SetHighThruInd sets HighThruInd field to given value.

### HasHighThruInd

`func (o *AppAmContextExpData) HasHighThruInd() bool`

HasHighThruInd returns a boolean if a field has been set.

### GetCovReqs

`func (o *AppAmContextExpData) GetCovReqs() []GeographicalArea`

GetCovReqs returns the CovReqs field if non-nil, zero value otherwise.

### GetCovReqsOk

`func (o *AppAmContextExpData) GetCovReqsOk() (*[]GeographicalArea, bool)`

GetCovReqsOk returns a tuple with the CovReqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovReqs

`func (o *AppAmContextExpData) SetCovReqs(v []GeographicalArea)`

SetCovReqs sets CovReqs field to given value.

### HasCovReqs

`func (o *AppAmContextExpData) HasCovReqs() bool`

HasCovReqs returns a boolean if a field has been set.

### SetCovReqsNil

`func (o *AppAmContextExpData) SetCovReqsNil(b bool)`

 SetCovReqsNil sets the value for CovReqs to be an explicit nil

### UnsetCovReqs
`func (o *AppAmContextExpData) UnsetCovReqs()`

UnsetCovReqs ensures that no value is present for CovReqs, not even an explicit nil
### GetPolicyDuration

`func (o *AppAmContextExpData) GetPolicyDuration() int32`

GetPolicyDuration returns the PolicyDuration field if non-nil, zero value otherwise.

### GetPolicyDurationOk

`func (o *AppAmContextExpData) GetPolicyDurationOk() (*int32, bool)`

GetPolicyDurationOk returns a tuple with the PolicyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDuration

`func (o *AppAmContextExpData) SetPolicyDuration(v int32)`

SetPolicyDuration sets PolicyDuration field to given value.

### HasPolicyDuration

`func (o *AppAmContextExpData) HasPolicyDuration() bool`

HasPolicyDuration returns a boolean if a field has been set.

### SetPolicyDurationNil

`func (o *AppAmContextExpData) SetPolicyDurationNil(b bool)`

 SetPolicyDurationNil sets the value for PolicyDuration to be an explicit nil

### UnsetPolicyDuration
`func (o *AppAmContextExpData) UnsetPolicyDuration()`

UnsetPolicyDuration ensures that no value is present for PolicyDuration, not even an explicit nil
### GetSuppFeat

`func (o *AppAmContextExpData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AppAmContextExpData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AppAmContextExpData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AppAmContextExpData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *AppAmContextExpData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *AppAmContextExpData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *AppAmContextExpData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *AppAmContextExpData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *AppAmContextExpData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *AppAmContextExpData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *AppAmContextExpData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *AppAmContextExpData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


