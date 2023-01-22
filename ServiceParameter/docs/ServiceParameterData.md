# ServiceParameterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfServiceId** | Pointer to **string** | Identifies a service on behalf of which the AF is issuing the request. | [optional] 
**AppId** | Pointer to **string** | Identifies an application. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**AnyUeInd** | Pointer to **bool** | Identifies whether the AF request applies to any UE. This attribute shall set to \&quot;true\&quot; if applicable for any UE, otherwise, set to \&quot;false\&quot;.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**UeIpv4** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**UeIpv6** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**UeMac** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SubNotifEvents** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the AF to request the NEF to send a test notification as defined in clause 5.2.5.3 of 3GPP TS 29.122. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**ParamOverPc5** | Pointer to **string** | Represents configuration parameters for V2X communications over PC5 reference point.  | [optional] 
**ParamOverUu** | Pointer to **string** | Represents configuration parameters for V2X communications over Uu reference point.  | [optional] 
**ParamForProSeDd** | Pointer to **string** | Represents the service parameters for 5G ProSe direct discovery. | [optional] 
**ParamForProSeDc** | Pointer to **string** | Represents the service parameters for 5G ProSe direct communications. | [optional] 
**ParamForProSeU2NRelUe** | Pointer to **string** | Represents the service parameters for 5G ProSe UE-to-network relay UE. | [optional] 
**ParamForProSeRemUe** | Pointer to **string** | Represents the service parameters for 5G ProSe Remate UE. | [optional] 
**UrspGuidance** | Pointer to [**[]UrspRuleRequest**](UrspRuleRequest.md) | Contains the service parameter used to guide the URSP. | [optional] 
**MtcProviderId** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewServiceParameterData

`func NewServiceParameterData() *ServiceParameterData`

NewServiceParameterData instantiates a new ServiceParameterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceParameterDataWithDefaults

`func NewServiceParameterDataWithDefaults() *ServiceParameterData`

NewServiceParameterDataWithDefaults instantiates a new ServiceParameterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfServiceId

`func (o *ServiceParameterData) GetAfServiceId() string`

GetAfServiceId returns the AfServiceId field if non-nil, zero value otherwise.

### GetAfServiceIdOk

`func (o *ServiceParameterData) GetAfServiceIdOk() (*string, bool)`

GetAfServiceIdOk returns a tuple with the AfServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfServiceId

`func (o *ServiceParameterData) SetAfServiceId(v string)`

SetAfServiceId sets AfServiceId field to given value.

### HasAfServiceId

`func (o *ServiceParameterData) HasAfServiceId() bool`

HasAfServiceId returns a boolean if a field has been set.

### GetAppId

`func (o *ServiceParameterData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ServiceParameterData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ServiceParameterData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ServiceParameterData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnn

`func (o *ServiceParameterData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *ServiceParameterData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *ServiceParameterData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *ServiceParameterData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *ServiceParameterData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ServiceParameterData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ServiceParameterData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *ServiceParameterData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *ServiceParameterData) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *ServiceParameterData) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *ServiceParameterData) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *ServiceParameterData) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *ServiceParameterData) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *ServiceParameterData) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *ServiceParameterData) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *ServiceParameterData) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetGpsi

`func (o *ServiceParameterData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *ServiceParameterData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *ServiceParameterData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *ServiceParameterData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetUeIpv4

`func (o *ServiceParameterData) GetUeIpv4() string`

GetUeIpv4 returns the UeIpv4 field if non-nil, zero value otherwise.

### GetUeIpv4Ok

`func (o *ServiceParameterData) GetUeIpv4Ok() (*string, bool)`

GetUeIpv4Ok returns a tuple with the UeIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4

`func (o *ServiceParameterData) SetUeIpv4(v string)`

SetUeIpv4 sets UeIpv4 field to given value.

### HasUeIpv4

`func (o *ServiceParameterData) HasUeIpv4() bool`

HasUeIpv4 returns a boolean if a field has been set.

### GetUeIpv6

`func (o *ServiceParameterData) GetUeIpv6() Ipv6Addr`

GetUeIpv6 returns the UeIpv6 field if non-nil, zero value otherwise.

### GetUeIpv6Ok

`func (o *ServiceParameterData) GetUeIpv6Ok() (*Ipv6Addr, bool)`

GetUeIpv6Ok returns a tuple with the UeIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6

`func (o *ServiceParameterData) SetUeIpv6(v Ipv6Addr)`

SetUeIpv6 sets UeIpv6 field to given value.

### HasUeIpv6

`func (o *ServiceParameterData) HasUeIpv6() bool`

HasUeIpv6 returns a boolean if a field has been set.

### GetUeMac

`func (o *ServiceParameterData) GetUeMac() string`

GetUeMac returns the UeMac field if non-nil, zero value otherwise.

### GetUeMacOk

`func (o *ServiceParameterData) GetUeMacOk() (*string, bool)`

GetUeMacOk returns a tuple with the UeMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMac

`func (o *ServiceParameterData) SetUeMac(v string)`

SetUeMac sets UeMac field to given value.

### HasUeMac

`func (o *ServiceParameterData) HasUeMac() bool`

HasUeMac returns a boolean if a field has been set.

### GetSelf

`func (o *ServiceParameterData) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ServiceParameterData) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ServiceParameterData) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ServiceParameterData) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSubNotifEvents

`func (o *ServiceParameterData) GetSubNotifEvents() []Event`

GetSubNotifEvents returns the SubNotifEvents field if non-nil, zero value otherwise.

### GetSubNotifEventsOk

`func (o *ServiceParameterData) GetSubNotifEventsOk() (*[]Event, bool)`

GetSubNotifEventsOk returns a tuple with the SubNotifEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNotifEvents

`func (o *ServiceParameterData) SetSubNotifEvents(v []Event)`

SetSubNotifEvents sets SubNotifEvents field to given value.

### HasSubNotifEvents

`func (o *ServiceParameterData) HasSubNotifEvents() bool`

HasSubNotifEvents returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *ServiceParameterData) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ServiceParameterData) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ServiceParameterData) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ServiceParameterData) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *ServiceParameterData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ServiceParameterData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ServiceParameterData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ServiceParameterData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ServiceParameterData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ServiceParameterData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ServiceParameterData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ServiceParameterData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetParamOverPc5

`func (o *ServiceParameterData) GetParamOverPc5() string`

GetParamOverPc5 returns the ParamOverPc5 field if non-nil, zero value otherwise.

### GetParamOverPc5Ok

`func (o *ServiceParameterData) GetParamOverPc5Ok() (*string, bool)`

GetParamOverPc5Ok returns a tuple with the ParamOverPc5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverPc5

`func (o *ServiceParameterData) SetParamOverPc5(v string)`

SetParamOverPc5 sets ParamOverPc5 field to given value.

### HasParamOverPc5

`func (o *ServiceParameterData) HasParamOverPc5() bool`

HasParamOverPc5 returns a boolean if a field has been set.

### GetParamOverUu

`func (o *ServiceParameterData) GetParamOverUu() string`

GetParamOverUu returns the ParamOverUu field if non-nil, zero value otherwise.

### GetParamOverUuOk

`func (o *ServiceParameterData) GetParamOverUuOk() (*string, bool)`

GetParamOverUuOk returns a tuple with the ParamOverUu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverUu

`func (o *ServiceParameterData) SetParamOverUu(v string)`

SetParamOverUu sets ParamOverUu field to given value.

### HasParamOverUu

`func (o *ServiceParameterData) HasParamOverUu() bool`

HasParamOverUu returns a boolean if a field has been set.

### GetParamForProSeDd

`func (o *ServiceParameterData) GetParamForProSeDd() string`

GetParamForProSeDd returns the ParamForProSeDd field if non-nil, zero value otherwise.

### GetParamForProSeDdOk

`func (o *ServiceParameterData) GetParamForProSeDdOk() (*string, bool)`

GetParamForProSeDdOk returns a tuple with the ParamForProSeDd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeDd

`func (o *ServiceParameterData) SetParamForProSeDd(v string)`

SetParamForProSeDd sets ParamForProSeDd field to given value.

### HasParamForProSeDd

`func (o *ServiceParameterData) HasParamForProSeDd() bool`

HasParamForProSeDd returns a boolean if a field has been set.

### GetParamForProSeDc

`func (o *ServiceParameterData) GetParamForProSeDc() string`

GetParamForProSeDc returns the ParamForProSeDc field if non-nil, zero value otherwise.

### GetParamForProSeDcOk

`func (o *ServiceParameterData) GetParamForProSeDcOk() (*string, bool)`

GetParamForProSeDcOk returns a tuple with the ParamForProSeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeDc

`func (o *ServiceParameterData) SetParamForProSeDc(v string)`

SetParamForProSeDc sets ParamForProSeDc field to given value.

### HasParamForProSeDc

`func (o *ServiceParameterData) HasParamForProSeDc() bool`

HasParamForProSeDc returns a boolean if a field has been set.

### GetParamForProSeU2NRelUe

`func (o *ServiceParameterData) GetParamForProSeU2NRelUe() string`

GetParamForProSeU2NRelUe returns the ParamForProSeU2NRelUe field if non-nil, zero value otherwise.

### GetParamForProSeU2NRelUeOk

`func (o *ServiceParameterData) GetParamForProSeU2NRelUeOk() (*string, bool)`

GetParamForProSeU2NRelUeOk returns a tuple with the ParamForProSeU2NRelUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeU2NRelUe

`func (o *ServiceParameterData) SetParamForProSeU2NRelUe(v string)`

SetParamForProSeU2NRelUe sets ParamForProSeU2NRelUe field to given value.

### HasParamForProSeU2NRelUe

`func (o *ServiceParameterData) HasParamForProSeU2NRelUe() bool`

HasParamForProSeU2NRelUe returns a boolean if a field has been set.

### GetParamForProSeRemUe

`func (o *ServiceParameterData) GetParamForProSeRemUe() string`

GetParamForProSeRemUe returns the ParamForProSeRemUe field if non-nil, zero value otherwise.

### GetParamForProSeRemUeOk

`func (o *ServiceParameterData) GetParamForProSeRemUeOk() (*string, bool)`

GetParamForProSeRemUeOk returns a tuple with the ParamForProSeRemUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeRemUe

`func (o *ServiceParameterData) SetParamForProSeRemUe(v string)`

SetParamForProSeRemUe sets ParamForProSeRemUe field to given value.

### HasParamForProSeRemUe

`func (o *ServiceParameterData) HasParamForProSeRemUe() bool`

HasParamForProSeRemUe returns a boolean if a field has been set.

### GetUrspGuidance

`func (o *ServiceParameterData) GetUrspGuidance() []UrspRuleRequest`

GetUrspGuidance returns the UrspGuidance field if non-nil, zero value otherwise.

### GetUrspGuidanceOk

`func (o *ServiceParameterData) GetUrspGuidanceOk() (*[]UrspRuleRequest, bool)`

GetUrspGuidanceOk returns a tuple with the UrspGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrspGuidance

`func (o *ServiceParameterData) SetUrspGuidance(v []UrspRuleRequest)`

SetUrspGuidance sets UrspGuidance field to given value.

### HasUrspGuidance

`func (o *ServiceParameterData) HasUrspGuidance() bool`

HasUrspGuidance returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *ServiceParameterData) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *ServiceParameterData) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *ServiceParameterData) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *ServiceParameterData) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ServiceParameterData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ServiceParameterData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ServiceParameterData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ServiceParameterData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


