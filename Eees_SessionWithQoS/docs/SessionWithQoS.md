# SessionWithQoS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**EasId** | **string** | Identifier of an EAS. | 
**UeIpv4Addr** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166. | [optional] 
**UeIpv6Addr** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] 
**IpDomain** | Pointer to **string** |  | [optional] 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**IntGrpId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExtGrpId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**IpFlows** | **[]string** | Contains the flow description for the Uplink and/or Downlink IP flows. | 
**QosReference** | Pointer to **string** | Identifies a pre-defined QoS information. | [optional] 
**AltQosReference** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority.  | [optional] 
**Events** | Pointer to [**[]UserPlaneEvent**](UserPlaneEvent.md) | Indicates the events subscribed by the EAS. | [optional] 
**SponsorInformation** | Pointer to [**SponsorInformation**](SponsorInformation.md) |  | [optional] 
**QosMonInfo** | Pointer to [**QosMonitoringInformation**](QosMonitoringInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**MaxbrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the EES to send a test notification as defined in 3GPP TS 29.122 [6]. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSessionWithQoS

`func NewSessionWithQoS(easId string, ipFlows []string, ) *SessionWithQoS`

NewSessionWithQoS instantiates a new SessionWithQoS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithQoSWithDefaults

`func NewSessionWithQoSWithDefaults() *SessionWithQoS`

NewSessionWithQoSWithDefaults instantiates a new SessionWithQoS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *SessionWithQoS) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *SessionWithQoS) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *SessionWithQoS) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *SessionWithQoS) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetEasId

`func (o *SessionWithQoS) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *SessionWithQoS) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *SessionWithQoS) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetUeIpv4Addr

`func (o *SessionWithQoS) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *SessionWithQoS) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *SessionWithQoS) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *SessionWithQoS) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetUeIpv6Addr

`func (o *SessionWithQoS) GetUeIpv6Addr() string`

GetUeIpv6Addr returns the UeIpv6Addr field if non-nil, zero value otherwise.

### GetUeIpv6AddrOk

`func (o *SessionWithQoS) GetUeIpv6AddrOk() (*string, bool)`

GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Addr

`func (o *SessionWithQoS) SetUeIpv6Addr(v string)`

SetUeIpv6Addr sets UeIpv6Addr field to given value.

### HasUeIpv6Addr

`func (o *SessionWithQoS) HasUeIpv6Addr() bool`

HasUeIpv6Addr returns a boolean if a field has been set.

### GetIpDomain

`func (o *SessionWithQoS) GetIpDomain() string`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *SessionWithQoS) GetIpDomainOk() (*string, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *SessionWithQoS) SetIpDomain(v string)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *SessionWithQoS) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetUeId

`func (o *SessionWithQoS) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *SessionWithQoS) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *SessionWithQoS) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *SessionWithQoS) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetIntGrpId

`func (o *SessionWithQoS) GetIntGrpId() string`

GetIntGrpId returns the IntGrpId field if non-nil, zero value otherwise.

### GetIntGrpIdOk

`func (o *SessionWithQoS) GetIntGrpIdOk() (*string, bool)`

GetIntGrpIdOk returns a tuple with the IntGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGrpId

`func (o *SessionWithQoS) SetIntGrpId(v string)`

SetIntGrpId sets IntGrpId field to given value.

### HasIntGrpId

`func (o *SessionWithQoS) HasIntGrpId() bool`

HasIntGrpId returns a boolean if a field has been set.

### GetExtGrpId

`func (o *SessionWithQoS) GetExtGrpId() string`

GetExtGrpId returns the ExtGrpId field if non-nil, zero value otherwise.

### GetExtGrpIdOk

`func (o *SessionWithQoS) GetExtGrpIdOk() (*string, bool)`

GetExtGrpIdOk returns a tuple with the ExtGrpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGrpId

`func (o *SessionWithQoS) SetExtGrpId(v string)`

SetExtGrpId sets ExtGrpId field to given value.

### HasExtGrpId

`func (o *SessionWithQoS) HasExtGrpId() bool`

HasExtGrpId returns a boolean if a field has been set.

### GetIpFlows

`func (o *SessionWithQoS) GetIpFlows() []string`

GetIpFlows returns the IpFlows field if non-nil, zero value otherwise.

### GetIpFlowsOk

`func (o *SessionWithQoS) GetIpFlowsOk() (*[]string, bool)`

GetIpFlowsOk returns a tuple with the IpFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFlows

`func (o *SessionWithQoS) SetIpFlows(v []string)`

SetIpFlows sets IpFlows field to given value.


### GetQosReference

`func (o *SessionWithQoS) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *SessionWithQoS) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *SessionWithQoS) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *SessionWithQoS) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetAltQosReference

`func (o *SessionWithQoS) GetAltQosReference() []string`

GetAltQosReference returns the AltQosReference field if non-nil, zero value otherwise.

### GetAltQosReferenceOk

`func (o *SessionWithQoS) GetAltQosReferenceOk() (*[]string, bool)`

GetAltQosReferenceOk returns a tuple with the AltQosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReference

`func (o *SessionWithQoS) SetAltQosReference(v []string)`

SetAltQosReference sets AltQosReference field to given value.

### HasAltQosReference

`func (o *SessionWithQoS) HasAltQosReference() bool`

HasAltQosReference returns a boolean if a field has been set.

### GetEvents

`func (o *SessionWithQoS) GetEvents() []UserPlaneEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SessionWithQoS) GetEventsOk() (*[]UserPlaneEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SessionWithQoS) SetEvents(v []UserPlaneEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SessionWithQoS) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSponsorInformation

`func (o *SessionWithQoS) GetSponsorInformation() SponsorInformation`

GetSponsorInformation returns the SponsorInformation field if non-nil, zero value otherwise.

### GetSponsorInformationOk

`func (o *SessionWithQoS) GetSponsorInformationOk() (*SponsorInformation, bool)`

GetSponsorInformationOk returns a tuple with the SponsorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInformation

`func (o *SessionWithQoS) SetSponsorInformation(v SponsorInformation)`

SetSponsorInformation sets SponsorInformation field to given value.

### HasSponsorInformation

`func (o *SessionWithQoS) HasSponsorInformation() bool`

HasSponsorInformation returns a boolean if a field has been set.

### GetQosMonInfo

`func (o *SessionWithQoS) GetQosMonInfo() QosMonitoringInformation`

GetQosMonInfo returns the QosMonInfo field if non-nil, zero value otherwise.

### GetQosMonInfoOk

`func (o *SessionWithQoS) GetQosMonInfoOk() (*QosMonitoringInformation, bool)`

GetQosMonInfoOk returns a tuple with the QosMonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonInfo

`func (o *SessionWithQoS) SetQosMonInfo(v QosMonitoringInformation)`

SetQosMonInfo sets QosMonInfo field to given value.

### HasQosMonInfo

`func (o *SessionWithQoS) HasQosMonInfo() bool`

HasQosMonInfo returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *SessionWithQoS) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *SessionWithQoS) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *SessionWithQoS) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *SessionWithQoS) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetDnn

`func (o *SessionWithQoS) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SessionWithQoS) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SessionWithQoS) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SessionWithQoS) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *SessionWithQoS) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SessionWithQoS) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SessionWithQoS) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *SessionWithQoS) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetMaxbrUl

`func (o *SessionWithQoS) GetMaxbrUl() string`

GetMaxbrUl returns the MaxbrUl field if non-nil, zero value otherwise.

### GetMaxbrUlOk

`func (o *SessionWithQoS) GetMaxbrUlOk() (*string, bool)`

GetMaxbrUlOk returns a tuple with the MaxbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrUl

`func (o *SessionWithQoS) SetMaxbrUl(v string)`

SetMaxbrUl sets MaxbrUl field to given value.

### HasMaxbrUl

`func (o *SessionWithQoS) HasMaxbrUl() bool`

HasMaxbrUl returns a boolean if a field has been set.

### GetMaxbrDl

`func (o *SessionWithQoS) GetMaxbrDl() string`

GetMaxbrDl returns the MaxbrDl field if non-nil, zero value otherwise.

### GetMaxbrDlOk

`func (o *SessionWithQoS) GetMaxbrDlOk() (*string, bool)`

GetMaxbrDlOk returns a tuple with the MaxbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrDl

`func (o *SessionWithQoS) SetMaxbrDl(v string)`

SetMaxbrDl sets MaxbrDl field to given value.

### HasMaxbrDl

`func (o *SessionWithQoS) HasMaxbrDl() bool`

HasMaxbrDl returns a boolean if a field has been set.

### GetDisUeNotif

`func (o *SessionWithQoS) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *SessionWithQoS) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *SessionWithQoS) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *SessionWithQoS) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *SessionWithQoS) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *SessionWithQoS) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *SessionWithQoS) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *SessionWithQoS) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *SessionWithQoS) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *SessionWithQoS) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *SessionWithQoS) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *SessionWithQoS) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SessionWithQoS) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SessionWithQoS) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SessionWithQoS) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SessionWithQoS) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


