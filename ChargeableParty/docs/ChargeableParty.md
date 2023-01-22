# ChargeableParty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**NotificationDestination** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**ExterAppId** | Pointer to **string** | Identifies the external Application Identifier. | [optional] 
**Ipv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**Ipv6Addr** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] 
**MacAddr** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**FlowInfo** | Pointer to [**[]FlowInfo**](FlowInfo.md) | Describes the application flows. | [optional] 
**EthFlowInfo** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Identifies Ethernet packet flows. | [optional] 
**SponsorInformation** | [**SponsorInformation**](SponsorInformation.md) |  | 
**SponsoringEnabled** | **bool** | Indicates whether the sponsoring data connectivity is enabled (true) or not (false).  | 
**ReferenceId** | Pointer to **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | [optional] 
**ServAuthInfo** | Pointer to [**ServAuthInfo**](ServAuthInfo.md) |  | [optional] 
**UsageThreshold** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**Events** | Pointer to [**[]Event**](Event.md) | Represents the list of event(s) to which the SCS/AS requests to subscribe to. | [optional] 

## Methods

### NewChargeableParty

`func NewChargeableParty(notificationDestination string, sponsorInformation SponsorInformation, sponsoringEnabled bool, ) *ChargeableParty`

NewChargeableParty instantiates a new ChargeableParty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargeablePartyWithDefaults

`func NewChargeablePartyWithDefaults() *ChargeableParty`

NewChargeablePartyWithDefaults instantiates a new ChargeableParty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *ChargeableParty) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ChargeableParty) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ChargeableParty) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *ChargeableParty) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ChargeableParty) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ChargeableParty) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ChargeableParty) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ChargeableParty) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetDnn

`func (o *ChargeableParty) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *ChargeableParty) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *ChargeableParty) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *ChargeableParty) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *ChargeableParty) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *ChargeableParty) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *ChargeableParty) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *ChargeableParty) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *ChargeableParty) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ChargeableParty) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ChargeableParty) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *ChargeableParty) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ChargeableParty) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ChargeableParty) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ChargeableParty) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ChargeableParty) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ChargeableParty) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ChargeableParty) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ChargeableParty) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetExterAppId

`func (o *ChargeableParty) GetExterAppId() string`

GetExterAppId returns the ExterAppId field if non-nil, zero value otherwise.

### GetExterAppIdOk

`func (o *ChargeableParty) GetExterAppIdOk() (*string, bool)`

GetExterAppIdOk returns a tuple with the ExterAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterAppId

`func (o *ChargeableParty) SetExterAppId(v string)`

SetExterAppId sets ExterAppId field to given value.

### HasExterAppId

`func (o *ChargeableParty) HasExterAppId() bool`

HasExterAppId returns a boolean if a field has been set.

### GetIpv4Addr

`func (o *ChargeableParty) GetIpv4Addr() string`

GetIpv4Addr returns the Ipv4Addr field if non-nil, zero value otherwise.

### GetIpv4AddrOk

`func (o *ChargeableParty) GetIpv4AddrOk() (*string, bool)`

GetIpv4AddrOk returns a tuple with the Ipv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addr

`func (o *ChargeableParty) SetIpv4Addr(v string)`

SetIpv4Addr sets Ipv4Addr field to given value.

### HasIpv4Addr

`func (o *ChargeableParty) HasIpv4Addr() bool`

HasIpv4Addr returns a boolean if a field has been set.

### GetIpDomain

`func (o *ChargeableParty) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *ChargeableParty) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *ChargeableParty) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *ChargeableParty) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetIpv6Addr

`func (o *ChargeableParty) GetIpv6Addr() string`

GetIpv6Addr returns the Ipv6Addr field if non-nil, zero value otherwise.

### GetIpv6AddrOk

`func (o *ChargeableParty) GetIpv6AddrOk() (*string, bool)`

GetIpv6AddrOk returns a tuple with the Ipv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addr

`func (o *ChargeableParty) SetIpv6Addr(v string)`

SetIpv6Addr sets Ipv6Addr field to given value.

### HasIpv6Addr

`func (o *ChargeableParty) HasIpv6Addr() bool`

HasIpv6Addr returns a boolean if a field has been set.

### GetMacAddr

`func (o *ChargeableParty) GetMacAddr() string`

GetMacAddr returns the MacAddr field if non-nil, zero value otherwise.

### GetMacAddrOk

`func (o *ChargeableParty) GetMacAddrOk() (*string, bool)`

GetMacAddrOk returns a tuple with the MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr

`func (o *ChargeableParty) SetMacAddr(v string)`

SetMacAddr sets MacAddr field to given value.

### HasMacAddr

`func (o *ChargeableParty) HasMacAddr() bool`

HasMacAddr returns a boolean if a field has been set.

### GetFlowInfo

`func (o *ChargeableParty) GetFlowInfo() []FlowInfo`

GetFlowInfo returns the FlowInfo field if non-nil, zero value otherwise.

### GetFlowInfoOk

`func (o *ChargeableParty) GetFlowInfoOk() (*[]FlowInfo, bool)`

GetFlowInfoOk returns a tuple with the FlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowInfo

`func (o *ChargeableParty) SetFlowInfo(v []FlowInfo)`

SetFlowInfo sets FlowInfo field to given value.

### HasFlowInfo

`func (o *ChargeableParty) HasFlowInfo() bool`

HasFlowInfo returns a boolean if a field has been set.

### GetEthFlowInfo

`func (o *ChargeableParty) GetEthFlowInfo() []EthFlowDescription`

GetEthFlowInfo returns the EthFlowInfo field if non-nil, zero value otherwise.

### GetEthFlowInfoOk

`func (o *ChargeableParty) GetEthFlowInfoOk() (*[]EthFlowDescription, bool)`

GetEthFlowInfoOk returns a tuple with the EthFlowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowInfo

`func (o *ChargeableParty) SetEthFlowInfo(v []EthFlowDescription)`

SetEthFlowInfo sets EthFlowInfo field to given value.

### HasEthFlowInfo

`func (o *ChargeableParty) HasEthFlowInfo() bool`

HasEthFlowInfo returns a boolean if a field has been set.

### GetSponsorInformation

`func (o *ChargeableParty) GetSponsorInformation() SponsorInformation`

GetSponsorInformation returns the SponsorInformation field if non-nil, zero value otherwise.

### GetSponsorInformationOk

`func (o *ChargeableParty) GetSponsorInformationOk() (*SponsorInformation, bool)`

GetSponsorInformationOk returns a tuple with the SponsorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInformation

`func (o *ChargeableParty) SetSponsorInformation(v SponsorInformation)`

SetSponsorInformation sets SponsorInformation field to given value.


### GetSponsoringEnabled

`func (o *ChargeableParty) GetSponsoringEnabled() bool`

GetSponsoringEnabled returns the SponsoringEnabled field if non-nil, zero value otherwise.

### GetSponsoringEnabledOk

`func (o *ChargeableParty) GetSponsoringEnabledOk() (*bool, bool)`

GetSponsoringEnabledOk returns a tuple with the SponsoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsoringEnabled

`func (o *ChargeableParty) SetSponsoringEnabled(v bool)`

SetSponsoringEnabled sets SponsoringEnabled field to given value.


### GetReferenceId

`func (o *ChargeableParty) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ChargeableParty) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ChargeableParty) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ChargeableParty) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetServAuthInfo

`func (o *ChargeableParty) GetServAuthInfo() ServAuthInfo`

GetServAuthInfo returns the ServAuthInfo field if non-nil, zero value otherwise.

### GetServAuthInfoOk

`func (o *ChargeableParty) GetServAuthInfoOk() (*ServAuthInfo, bool)`

GetServAuthInfoOk returns a tuple with the ServAuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAuthInfo

`func (o *ChargeableParty) SetServAuthInfo(v ServAuthInfo)`

SetServAuthInfo sets ServAuthInfo field to given value.

### HasServAuthInfo

`func (o *ChargeableParty) HasServAuthInfo() bool`

HasServAuthInfo returns a boolean if a field has been set.

### GetUsageThreshold

`func (o *ChargeableParty) GetUsageThreshold() UsageThreshold`

GetUsageThreshold returns the UsageThreshold field if non-nil, zero value otherwise.

### GetUsageThresholdOk

`func (o *ChargeableParty) GetUsageThresholdOk() (*UsageThreshold, bool)`

GetUsageThresholdOk returns a tuple with the UsageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageThreshold

`func (o *ChargeableParty) SetUsageThreshold(v UsageThreshold)`

SetUsageThreshold sets UsageThreshold field to given value.

### HasUsageThreshold

`func (o *ChargeableParty) HasUsageThreshold() bool`

HasUsageThreshold returns a boolean if a field has been set.

### GetEvents

`func (o *ChargeableParty) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ChargeableParty) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ChargeableParty) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ChargeableParty) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


