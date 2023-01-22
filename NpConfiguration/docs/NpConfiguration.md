# NpConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MtcProviderId** | Pointer to **string** | Identifies the MTC Service Provider and/or MTC Application. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**MaximumLatency** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**MaximumResponseTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**SuggestedNumberOfDlPackets** | Pointer to **int32** | This parameter may be included to identify the number of packets that the serving gateway shall buffer in case that the UE is not reachable. | [optional] 
**GroupReportingGuardTime** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**ValidityTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**UeIpAddr** | Pointer to [**IpAddr**](IpAddr.md) |  | [optional] 
**UeMacAddr** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 

## Methods

### NewNpConfiguration

`func NewNpConfiguration() *NpConfiguration`

NewNpConfiguration instantiates a new NpConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNpConfigurationWithDefaults

`func NewNpConfigurationWithDefaults() *NpConfiguration`

NewNpConfigurationWithDefaults instantiates a new NpConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NpConfiguration) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NpConfiguration) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NpConfiguration) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *NpConfiguration) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NpConfiguration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NpConfiguration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NpConfiguration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NpConfiguration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *NpConfiguration) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *NpConfiguration) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *NpConfiguration) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *NpConfiguration) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetDnn

`func (o *NpConfiguration) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *NpConfiguration) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *NpConfiguration) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *NpConfiguration) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetExternalId

`func (o *NpConfiguration) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NpConfiguration) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NpConfiguration) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NpConfiguration) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *NpConfiguration) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *NpConfiguration) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *NpConfiguration) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *NpConfiguration) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *NpConfiguration) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *NpConfiguration) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *NpConfiguration) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *NpConfiguration) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetMaximumLatency

`func (o *NpConfiguration) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *NpConfiguration) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *NpConfiguration) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *NpConfiguration) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### GetMaximumResponseTime

`func (o *NpConfiguration) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *NpConfiguration) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *NpConfiguration) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *NpConfiguration) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### GetSuggestedNumberOfDlPackets

`func (o *NpConfiguration) GetSuggestedNumberOfDlPackets() int32`

GetSuggestedNumberOfDlPackets returns the SuggestedNumberOfDlPackets field if non-nil, zero value otherwise.

### GetSuggestedNumberOfDlPacketsOk

`func (o *NpConfiguration) GetSuggestedNumberOfDlPacketsOk() (*int32, bool)`

GetSuggestedNumberOfDlPacketsOk returns a tuple with the SuggestedNumberOfDlPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedNumberOfDlPackets

`func (o *NpConfiguration) SetSuggestedNumberOfDlPackets(v int32)`

SetSuggestedNumberOfDlPackets sets SuggestedNumberOfDlPackets field to given value.

### HasSuggestedNumberOfDlPackets

`func (o *NpConfiguration) HasSuggestedNumberOfDlPackets() bool`

HasSuggestedNumberOfDlPackets returns a boolean if a field has been set.

### GetGroupReportingGuardTime

`func (o *NpConfiguration) GetGroupReportingGuardTime() int32`

GetGroupReportingGuardTime returns the GroupReportingGuardTime field if non-nil, zero value otherwise.

### GetGroupReportingGuardTimeOk

`func (o *NpConfiguration) GetGroupReportingGuardTimeOk() (*int32, bool)`

GetGroupReportingGuardTimeOk returns a tuple with the GroupReportingGuardTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupReportingGuardTime

`func (o *NpConfiguration) SetGroupReportingGuardTime(v int32)`

SetGroupReportingGuardTime sets GroupReportingGuardTime field to given value.

### HasGroupReportingGuardTime

`func (o *NpConfiguration) HasGroupReportingGuardTime() bool`

HasGroupReportingGuardTime returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *NpConfiguration) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *NpConfiguration) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *NpConfiguration) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *NpConfiguration) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *NpConfiguration) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *NpConfiguration) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *NpConfiguration) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *NpConfiguration) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *NpConfiguration) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *NpConfiguration) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *NpConfiguration) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *NpConfiguration) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetValidityTime

`func (o *NpConfiguration) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *NpConfiguration) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *NpConfiguration) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *NpConfiguration) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetSnssai

`func (o *NpConfiguration) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NpConfiguration) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NpConfiguration) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *NpConfiguration) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetUeIpAddr

`func (o *NpConfiguration) GetUeIpAddr() IpAddr`

GetUeIpAddr returns the UeIpAddr field if non-nil, zero value otherwise.

### GetUeIpAddrOk

`func (o *NpConfiguration) GetUeIpAddrOk() (*IpAddr, bool)`

GetUeIpAddrOk returns a tuple with the UeIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpAddr

`func (o *NpConfiguration) SetUeIpAddr(v IpAddr)`

SetUeIpAddr sets UeIpAddr field to given value.

### HasUeIpAddr

`func (o *NpConfiguration) HasUeIpAddr() bool`

HasUeIpAddr returns a boolean if a field has been set.

### GetUeMacAddr

`func (o *NpConfiguration) GetUeMacAddr() string`

GetUeMacAddr returns the UeMacAddr field if non-nil, zero value otherwise.

### GetUeMacAddrOk

`func (o *NpConfiguration) GetUeMacAddrOk() (*string, bool)`

GetUeMacAddrOk returns a tuple with the UeMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMacAddr

`func (o *NpConfiguration) SetUeMacAddr(v string)`

SetUeMacAddr sets UeMacAddr field to given value.

### HasUeMacAddr

`func (o *NpConfiguration) HasUeMacAddr() bool`

HasUeMacAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


