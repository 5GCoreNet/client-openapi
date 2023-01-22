# AmInfluSub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfTransId** | **string** |  | 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**AnyUeInd** | Pointer to **bool** | Identifies whether the AF request applies to any UE. This attribute shall set to \&quot;true\&quot; if applicable for any UE, otherwise, set to \&quot;false\&quot;.  | [optional] 
**DnnSnssaiInfos** | Pointer to [**[]DnnSnssaiInformation**](DnnSnssaiInformation.md) | Each of the element identifies a (DNN, S-NSSAI) combination. | [optional] 
**AfAppIds** | Pointer to **[]string** | Each of the element identifies an application. | [optional] 
**HighThruInd** | Pointer to **bool** |  | [optional] 
**GeoAreas** | Pointer to [**[]GeographicalArea**](GeographicalArea.md) | Identifies geographic areas of the user where the request is applicable. | [optional] 
**PolicyDuration** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SubscribedEvents** | Pointer to [**[]AmInfluEvent**](AmInfluEvent.md) | Indicates one or more AM influence related events. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the AF to request the NEF to send a test notification as defined in clause 5.2.5.3 of 3GPP TS 29.122. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAmInfluSub

`func NewAmInfluSub(afTransId string, ) *AmInfluSub`

NewAmInfluSub instantiates a new AmInfluSub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmInfluSubWithDefaults

`func NewAmInfluSubWithDefaults() *AmInfluSub`

NewAmInfluSubWithDefaults instantiates a new AmInfluSub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfTransId

`func (o *AmInfluSub) GetAfTransId() string`

GetAfTransId returns the AfTransId field if non-nil, zero value otherwise.

### GetAfTransIdOk

`func (o *AmInfluSub) GetAfTransIdOk() (*string, bool)`

GetAfTransIdOk returns a tuple with the AfTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfTransId

`func (o *AmInfluSub) SetAfTransId(v string)`

SetAfTransId sets AfTransId field to given value.


### GetGpsi

`func (o *AmInfluSub) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AmInfluSub) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AmInfluSub) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AmInfluSub) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *AmInfluSub) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *AmInfluSub) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *AmInfluSub) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *AmInfluSub) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *AmInfluSub) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *AmInfluSub) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *AmInfluSub) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *AmInfluSub) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetDnnSnssaiInfos

`func (o *AmInfluSub) GetDnnSnssaiInfos() []DnnSnssaiInformation`

GetDnnSnssaiInfos returns the DnnSnssaiInfos field if non-nil, zero value otherwise.

### GetDnnSnssaiInfosOk

`func (o *AmInfluSub) GetDnnSnssaiInfosOk() (*[]DnnSnssaiInformation, bool)`

GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssaiInfos

`func (o *AmInfluSub) SetDnnSnssaiInfos(v []DnnSnssaiInformation)`

SetDnnSnssaiInfos sets DnnSnssaiInfos field to given value.

### HasDnnSnssaiInfos

`func (o *AmInfluSub) HasDnnSnssaiInfos() bool`

HasDnnSnssaiInfos returns a boolean if a field has been set.

### GetAfAppIds

`func (o *AmInfluSub) GetAfAppIds() []string`

GetAfAppIds returns the AfAppIds field if non-nil, zero value otherwise.

### GetAfAppIdsOk

`func (o *AmInfluSub) GetAfAppIdsOk() (*[]string, bool)`

GetAfAppIdsOk returns a tuple with the AfAppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppIds

`func (o *AmInfluSub) SetAfAppIds(v []string)`

SetAfAppIds sets AfAppIds field to given value.

### HasAfAppIds

`func (o *AmInfluSub) HasAfAppIds() bool`

HasAfAppIds returns a boolean if a field has been set.

### GetHighThruInd

`func (o *AmInfluSub) GetHighThruInd() bool`

GetHighThruInd returns the HighThruInd field if non-nil, zero value otherwise.

### GetHighThruIndOk

`func (o *AmInfluSub) GetHighThruIndOk() (*bool, bool)`

GetHighThruIndOk returns a tuple with the HighThruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThruInd

`func (o *AmInfluSub) SetHighThruInd(v bool)`

SetHighThruInd sets HighThruInd field to given value.

### HasHighThruInd

`func (o *AmInfluSub) HasHighThruInd() bool`

HasHighThruInd returns a boolean if a field has been set.

### GetGeoAreas

`func (o *AmInfluSub) GetGeoAreas() []GeographicalArea`

GetGeoAreas returns the GeoAreas field if non-nil, zero value otherwise.

### GetGeoAreasOk

`func (o *AmInfluSub) GetGeoAreasOk() (*[]GeographicalArea, bool)`

GetGeoAreasOk returns a tuple with the GeoAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoAreas

`func (o *AmInfluSub) SetGeoAreas(v []GeographicalArea)`

SetGeoAreas sets GeoAreas field to given value.

### HasGeoAreas

`func (o *AmInfluSub) HasGeoAreas() bool`

HasGeoAreas returns a boolean if a field has been set.

### GetPolicyDuration

`func (o *AmInfluSub) GetPolicyDuration() int32`

GetPolicyDuration returns the PolicyDuration field if non-nil, zero value otherwise.

### GetPolicyDurationOk

`func (o *AmInfluSub) GetPolicyDurationOk() (*int32, bool)`

GetPolicyDurationOk returns a tuple with the PolicyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDuration

`func (o *AmInfluSub) SetPolicyDuration(v int32)`

SetPolicyDuration sets PolicyDuration field to given value.

### HasPolicyDuration

`func (o *AmInfluSub) HasPolicyDuration() bool`

HasPolicyDuration returns a boolean if a field has been set.

### GetSelf

`func (o *AmInfluSub) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AmInfluSub) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AmInfluSub) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AmInfluSub) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSubscribedEvents

`func (o *AmInfluSub) GetSubscribedEvents() []AmInfluEvent`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *AmInfluSub) GetSubscribedEventsOk() (*[]AmInfluEvent, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *AmInfluSub) SetSubscribedEvents(v []AmInfluEvent)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *AmInfluSub) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *AmInfluSub) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *AmInfluSub) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *AmInfluSub) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *AmInfluSub) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *AmInfluSub) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *AmInfluSub) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *AmInfluSub) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *AmInfluSub) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *AmInfluSub) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *AmInfluSub) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *AmInfluSub) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *AmInfluSub) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AmInfluSub) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AmInfluSub) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AmInfluSub) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AmInfluSub) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


