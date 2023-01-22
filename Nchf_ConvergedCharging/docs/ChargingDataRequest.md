# ChargingDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberIdentifier** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**TenantIdentifier** | Pointer to **string** |  | [optional] 
**ChargingId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**MnSConsumerIdentifier** | Pointer to **string** |  | [optional] 
**NfConsumerIdentification** | [**NFIdentification**](NFIdentification.md) |  | 
**InvocationTimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**InvocationSequenceNumber** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**RetransmissionIndicator** | Pointer to **bool** |  | [optional] 
**OneTimeEvent** | Pointer to **bool** |  | [optional] 
**OneTimeEventType** | Pointer to [**OneTimeEventType**](OneTimeEventType.md) |  | [optional] 
**NotifyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ServiceSpecificationInfo** | Pointer to **string** |  | [optional] 
**MultipleUnitUsage** | Pointer to [**[]MultipleUnitUsage**](MultipleUnitUsage.md) |  | [optional] 
**Triggers** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**Easid** | Pointer to **string** |  | [optional] 
**Ednid** | Pointer to **string** |  | [optional] 
**EASProviderIdentifier** | Pointer to **string** |  | [optional] 
**PDUSessionChargingInformation** | Pointer to [**PDUSessionChargingInformation**](PDUSessionChargingInformation.md) |  | [optional] 
**RoamingQBCInformation** | Pointer to [**RoamingQBCInformation**](RoamingQBCInformation.md) |  | [optional] 
**SMSChargingInformation** | Pointer to [**SMSChargingInformation**](SMSChargingInformation.md) |  | [optional] 
**NEFChargingInformation** | Pointer to [**NEFChargingInformation**](NEFChargingInformation.md) |  | [optional] 
**RegistrationChargingInformation** | Pointer to [**RegistrationChargingInformation**](RegistrationChargingInformation.md) |  | [optional] 
**N2ConnectionChargingInformation** | Pointer to [**N2ConnectionChargingInformation**](N2ConnectionChargingInformation.md) |  | [optional] 
**LocationReportingChargingInformation** | Pointer to [**LocationReportingChargingInformation**](LocationReportingChargingInformation.md) |  | [optional] 
**NSPAChargingInformation** | Pointer to [**NSPAChargingInformation**](NSPAChargingInformation.md) |  | [optional] 
**NSMChargingInformation** | Pointer to [**NSMChargingInformation**](NSMChargingInformation.md) |  | [optional] 
**MMTelChargingInformation** | Pointer to [**MMTelChargingInformation**](MMTelChargingInformation.md) |  | [optional] 
**IMSChargingInformation** | Pointer to [**IMSChargingInformation**](IMSChargingInformation.md) |  | [optional] 
**EdgeInfrastructureUsageChargingInformation** | Pointer to [**EdgeInfrastructureUsageChargingInformation**](EdgeInfrastructureUsageChargingInformation.md) |  | [optional] 
**EASDeploymentChargingInformation** | Pointer to [**EASDeploymentChargingInformation**](EASDeploymentChargingInformation.md) |  | [optional] 
**DirectEdgeEnablingServiceChargingInformation** | Pointer to [**NEFChargingInformation**](NEFChargingInformation.md) |  | [optional] 
**ExposedEdgeEnablingServiceChargingInformation** | Pointer to [**NEFChargingInformation**](NEFChargingInformation.md) |  | [optional] 
**ProSeChargingInformation** | Pointer to [**ProseChargingInformation**](ProseChargingInformation.md) |  | [optional] 

## Methods

### NewChargingDataRequest

`func NewChargingDataRequest(nfConsumerIdentification NFIdentification, invocationTimeStamp time.Time, invocationSequenceNumber int32, ) *ChargingDataRequest`

NewChargingDataRequest instantiates a new ChargingDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingDataRequestWithDefaults

`func NewChargingDataRequestWithDefaults() *ChargingDataRequest`

NewChargingDataRequestWithDefaults instantiates a new ChargingDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberIdentifier

`func (o *ChargingDataRequest) GetSubscriberIdentifier() string`

GetSubscriberIdentifier returns the SubscriberIdentifier field if non-nil, zero value otherwise.

### GetSubscriberIdentifierOk

`func (o *ChargingDataRequest) GetSubscriberIdentifierOk() (*string, bool)`

GetSubscriberIdentifierOk returns a tuple with the SubscriberIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdentifier

`func (o *ChargingDataRequest) SetSubscriberIdentifier(v string)`

SetSubscriberIdentifier sets SubscriberIdentifier field to given value.

### HasSubscriberIdentifier

`func (o *ChargingDataRequest) HasSubscriberIdentifier() bool`

HasSubscriberIdentifier returns a boolean if a field has been set.

### GetTenantIdentifier

`func (o *ChargingDataRequest) GetTenantIdentifier() string`

GetTenantIdentifier returns the TenantIdentifier field if non-nil, zero value otherwise.

### GetTenantIdentifierOk

`func (o *ChargingDataRequest) GetTenantIdentifierOk() (*string, bool)`

GetTenantIdentifierOk returns a tuple with the TenantIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantIdentifier

`func (o *ChargingDataRequest) SetTenantIdentifier(v string)`

SetTenantIdentifier sets TenantIdentifier field to given value.

### HasTenantIdentifier

`func (o *ChargingDataRequest) HasTenantIdentifier() bool`

HasTenantIdentifier returns a boolean if a field has been set.

### GetChargingId

`func (o *ChargingDataRequest) GetChargingId() int32`

GetChargingId returns the ChargingId field if non-nil, zero value otherwise.

### GetChargingIdOk

`func (o *ChargingDataRequest) GetChargingIdOk() (*int32, bool)`

GetChargingIdOk returns a tuple with the ChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingId

`func (o *ChargingDataRequest) SetChargingId(v int32)`

SetChargingId sets ChargingId field to given value.

### HasChargingId

`func (o *ChargingDataRequest) HasChargingId() bool`

HasChargingId returns a boolean if a field has been set.

### GetMnSConsumerIdentifier

`func (o *ChargingDataRequest) GetMnSConsumerIdentifier() string`

GetMnSConsumerIdentifier returns the MnSConsumerIdentifier field if non-nil, zero value otherwise.

### GetMnSConsumerIdentifierOk

`func (o *ChargingDataRequest) GetMnSConsumerIdentifierOk() (*string, bool)`

GetMnSConsumerIdentifierOk returns a tuple with the MnSConsumerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnSConsumerIdentifier

`func (o *ChargingDataRequest) SetMnSConsumerIdentifier(v string)`

SetMnSConsumerIdentifier sets MnSConsumerIdentifier field to given value.

### HasMnSConsumerIdentifier

`func (o *ChargingDataRequest) HasMnSConsumerIdentifier() bool`

HasMnSConsumerIdentifier returns a boolean if a field has been set.

### GetNfConsumerIdentification

`func (o *ChargingDataRequest) GetNfConsumerIdentification() NFIdentification`

GetNfConsumerIdentification returns the NfConsumerIdentification field if non-nil, zero value otherwise.

### GetNfConsumerIdentificationOk

`func (o *ChargingDataRequest) GetNfConsumerIdentificationOk() (*NFIdentification, bool)`

GetNfConsumerIdentificationOk returns a tuple with the NfConsumerIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfConsumerIdentification

`func (o *ChargingDataRequest) SetNfConsumerIdentification(v NFIdentification)`

SetNfConsumerIdentification sets NfConsumerIdentification field to given value.


### GetInvocationTimeStamp

`func (o *ChargingDataRequest) GetInvocationTimeStamp() time.Time`

GetInvocationTimeStamp returns the InvocationTimeStamp field if non-nil, zero value otherwise.

### GetInvocationTimeStampOk

`func (o *ChargingDataRequest) GetInvocationTimeStampOk() (*time.Time, bool)`

GetInvocationTimeStampOk returns a tuple with the InvocationTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationTimeStamp

`func (o *ChargingDataRequest) SetInvocationTimeStamp(v time.Time)`

SetInvocationTimeStamp sets InvocationTimeStamp field to given value.


### GetInvocationSequenceNumber

`func (o *ChargingDataRequest) GetInvocationSequenceNumber() int32`

GetInvocationSequenceNumber returns the InvocationSequenceNumber field if non-nil, zero value otherwise.

### GetInvocationSequenceNumberOk

`func (o *ChargingDataRequest) GetInvocationSequenceNumberOk() (*int32, bool)`

GetInvocationSequenceNumberOk returns a tuple with the InvocationSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationSequenceNumber

`func (o *ChargingDataRequest) SetInvocationSequenceNumber(v int32)`

SetInvocationSequenceNumber sets InvocationSequenceNumber field to given value.


### GetRetransmissionIndicator

`func (o *ChargingDataRequest) GetRetransmissionIndicator() bool`

GetRetransmissionIndicator returns the RetransmissionIndicator field if non-nil, zero value otherwise.

### GetRetransmissionIndicatorOk

`func (o *ChargingDataRequest) GetRetransmissionIndicatorOk() (*bool, bool)`

GetRetransmissionIndicatorOk returns a tuple with the RetransmissionIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmissionIndicator

`func (o *ChargingDataRequest) SetRetransmissionIndicator(v bool)`

SetRetransmissionIndicator sets RetransmissionIndicator field to given value.

### HasRetransmissionIndicator

`func (o *ChargingDataRequest) HasRetransmissionIndicator() bool`

HasRetransmissionIndicator returns a boolean if a field has been set.

### GetOneTimeEvent

`func (o *ChargingDataRequest) GetOneTimeEvent() bool`

GetOneTimeEvent returns the OneTimeEvent field if non-nil, zero value otherwise.

### GetOneTimeEventOk

`func (o *ChargingDataRequest) GetOneTimeEventOk() (*bool, bool)`

GetOneTimeEventOk returns a tuple with the OneTimeEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeEvent

`func (o *ChargingDataRequest) SetOneTimeEvent(v bool)`

SetOneTimeEvent sets OneTimeEvent field to given value.

### HasOneTimeEvent

`func (o *ChargingDataRequest) HasOneTimeEvent() bool`

HasOneTimeEvent returns a boolean if a field has been set.

### GetOneTimeEventType

`func (o *ChargingDataRequest) GetOneTimeEventType() OneTimeEventType`

GetOneTimeEventType returns the OneTimeEventType field if non-nil, zero value otherwise.

### GetOneTimeEventTypeOk

`func (o *ChargingDataRequest) GetOneTimeEventTypeOk() (*OneTimeEventType, bool)`

GetOneTimeEventTypeOk returns a tuple with the OneTimeEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeEventType

`func (o *ChargingDataRequest) SetOneTimeEventType(v OneTimeEventType)`

SetOneTimeEventType sets OneTimeEventType field to given value.

### HasOneTimeEventType

`func (o *ChargingDataRequest) HasOneTimeEventType() bool`

HasOneTimeEventType returns a boolean if a field has been set.

### GetNotifyUri

`func (o *ChargingDataRequest) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *ChargingDataRequest) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *ChargingDataRequest) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.

### HasNotifyUri

`func (o *ChargingDataRequest) HasNotifyUri() bool`

HasNotifyUri returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *ChargingDataRequest) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *ChargingDataRequest) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *ChargingDataRequest) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *ChargingDataRequest) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetServiceSpecificationInfo

`func (o *ChargingDataRequest) GetServiceSpecificationInfo() string`

GetServiceSpecificationInfo returns the ServiceSpecificationInfo field if non-nil, zero value otherwise.

### GetServiceSpecificationInfoOk

`func (o *ChargingDataRequest) GetServiceSpecificationInfoOk() (*string, bool)`

GetServiceSpecificationInfoOk returns a tuple with the ServiceSpecificationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpecificationInfo

`func (o *ChargingDataRequest) SetServiceSpecificationInfo(v string)`

SetServiceSpecificationInfo sets ServiceSpecificationInfo field to given value.

### HasServiceSpecificationInfo

`func (o *ChargingDataRequest) HasServiceSpecificationInfo() bool`

HasServiceSpecificationInfo returns a boolean if a field has been set.

### GetMultipleUnitUsage

`func (o *ChargingDataRequest) GetMultipleUnitUsage() []MultipleUnitUsage`

GetMultipleUnitUsage returns the MultipleUnitUsage field if non-nil, zero value otherwise.

### GetMultipleUnitUsageOk

`func (o *ChargingDataRequest) GetMultipleUnitUsageOk() (*[]MultipleUnitUsage, bool)`

GetMultipleUnitUsageOk returns a tuple with the MultipleUnitUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleUnitUsage

`func (o *ChargingDataRequest) SetMultipleUnitUsage(v []MultipleUnitUsage)`

SetMultipleUnitUsage sets MultipleUnitUsage field to given value.

### HasMultipleUnitUsage

`func (o *ChargingDataRequest) HasMultipleUnitUsage() bool`

HasMultipleUnitUsage returns a boolean if a field has been set.

### GetTriggers

`func (o *ChargingDataRequest) GetTriggers() []Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ChargingDataRequest) GetTriggersOk() (*[]Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ChargingDataRequest) SetTriggers(v []Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ChargingDataRequest) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetEasid

`func (o *ChargingDataRequest) GetEasid() string`

GetEasid returns the Easid field if non-nil, zero value otherwise.

### GetEasidOk

`func (o *ChargingDataRequest) GetEasidOk() (*string, bool)`

GetEasidOk returns a tuple with the Easid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasid

`func (o *ChargingDataRequest) SetEasid(v string)`

SetEasid sets Easid field to given value.

### HasEasid

`func (o *ChargingDataRequest) HasEasid() bool`

HasEasid returns a boolean if a field has been set.

### GetEdnid

`func (o *ChargingDataRequest) GetEdnid() string`

GetEdnid returns the Ednid field if non-nil, zero value otherwise.

### GetEdnidOk

`func (o *ChargingDataRequest) GetEdnidOk() (*string, bool)`

GetEdnidOk returns a tuple with the Ednid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnid

`func (o *ChargingDataRequest) SetEdnid(v string)`

SetEdnid sets Ednid field to given value.

### HasEdnid

`func (o *ChargingDataRequest) HasEdnid() bool`

HasEdnid returns a boolean if a field has been set.

### GetEASProviderIdentifier

`func (o *ChargingDataRequest) GetEASProviderIdentifier() string`

GetEASProviderIdentifier returns the EASProviderIdentifier field if non-nil, zero value otherwise.

### GetEASProviderIdentifierOk

`func (o *ChargingDataRequest) GetEASProviderIdentifierOk() (*string, bool)`

GetEASProviderIdentifierOk returns a tuple with the EASProviderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASProviderIdentifier

`func (o *ChargingDataRequest) SetEASProviderIdentifier(v string)`

SetEASProviderIdentifier sets EASProviderIdentifier field to given value.

### HasEASProviderIdentifier

`func (o *ChargingDataRequest) HasEASProviderIdentifier() bool`

HasEASProviderIdentifier returns a boolean if a field has been set.

### GetPDUSessionChargingInformation

`func (o *ChargingDataRequest) GetPDUSessionChargingInformation() PDUSessionChargingInformation`

GetPDUSessionChargingInformation returns the PDUSessionChargingInformation field if non-nil, zero value otherwise.

### GetPDUSessionChargingInformationOk

`func (o *ChargingDataRequest) GetPDUSessionChargingInformationOk() (*PDUSessionChargingInformation, bool)`

GetPDUSessionChargingInformationOk returns a tuple with the PDUSessionChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPDUSessionChargingInformation

`func (o *ChargingDataRequest) SetPDUSessionChargingInformation(v PDUSessionChargingInformation)`

SetPDUSessionChargingInformation sets PDUSessionChargingInformation field to given value.

### HasPDUSessionChargingInformation

`func (o *ChargingDataRequest) HasPDUSessionChargingInformation() bool`

HasPDUSessionChargingInformation returns a boolean if a field has been set.

### GetRoamingQBCInformation

`func (o *ChargingDataRequest) GetRoamingQBCInformation() RoamingQBCInformation`

GetRoamingQBCInformation returns the RoamingQBCInformation field if non-nil, zero value otherwise.

### GetRoamingQBCInformationOk

`func (o *ChargingDataRequest) GetRoamingQBCInformationOk() (*RoamingQBCInformation, bool)`

GetRoamingQBCInformationOk returns a tuple with the RoamingQBCInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingQBCInformation

`func (o *ChargingDataRequest) SetRoamingQBCInformation(v RoamingQBCInformation)`

SetRoamingQBCInformation sets RoamingQBCInformation field to given value.

### HasRoamingQBCInformation

`func (o *ChargingDataRequest) HasRoamingQBCInformation() bool`

HasRoamingQBCInformation returns a boolean if a field has been set.

### GetSMSChargingInformation

`func (o *ChargingDataRequest) GetSMSChargingInformation() SMSChargingInformation`

GetSMSChargingInformation returns the SMSChargingInformation field if non-nil, zero value otherwise.

### GetSMSChargingInformationOk

`func (o *ChargingDataRequest) GetSMSChargingInformationOk() (*SMSChargingInformation, bool)`

GetSMSChargingInformationOk returns a tuple with the SMSChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMSChargingInformation

`func (o *ChargingDataRequest) SetSMSChargingInformation(v SMSChargingInformation)`

SetSMSChargingInformation sets SMSChargingInformation field to given value.

### HasSMSChargingInformation

`func (o *ChargingDataRequest) HasSMSChargingInformation() bool`

HasSMSChargingInformation returns a boolean if a field has been set.

### GetNEFChargingInformation

`func (o *ChargingDataRequest) GetNEFChargingInformation() NEFChargingInformation`

GetNEFChargingInformation returns the NEFChargingInformation field if non-nil, zero value otherwise.

### GetNEFChargingInformationOk

`func (o *ChargingDataRequest) GetNEFChargingInformationOk() (*NEFChargingInformation, bool)`

GetNEFChargingInformationOk returns a tuple with the NEFChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEFChargingInformation

`func (o *ChargingDataRequest) SetNEFChargingInformation(v NEFChargingInformation)`

SetNEFChargingInformation sets NEFChargingInformation field to given value.

### HasNEFChargingInformation

`func (o *ChargingDataRequest) HasNEFChargingInformation() bool`

HasNEFChargingInformation returns a boolean if a field has been set.

### GetRegistrationChargingInformation

`func (o *ChargingDataRequest) GetRegistrationChargingInformation() RegistrationChargingInformation`

GetRegistrationChargingInformation returns the RegistrationChargingInformation field if non-nil, zero value otherwise.

### GetRegistrationChargingInformationOk

`func (o *ChargingDataRequest) GetRegistrationChargingInformationOk() (*RegistrationChargingInformation, bool)`

GetRegistrationChargingInformationOk returns a tuple with the RegistrationChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationChargingInformation

`func (o *ChargingDataRequest) SetRegistrationChargingInformation(v RegistrationChargingInformation)`

SetRegistrationChargingInformation sets RegistrationChargingInformation field to given value.

### HasRegistrationChargingInformation

`func (o *ChargingDataRequest) HasRegistrationChargingInformation() bool`

HasRegistrationChargingInformation returns a boolean if a field has been set.

### GetN2ConnectionChargingInformation

`func (o *ChargingDataRequest) GetN2ConnectionChargingInformation() N2ConnectionChargingInformation`

GetN2ConnectionChargingInformation returns the N2ConnectionChargingInformation field if non-nil, zero value otherwise.

### GetN2ConnectionChargingInformationOk

`func (o *ChargingDataRequest) GetN2ConnectionChargingInformationOk() (*N2ConnectionChargingInformation, bool)`

GetN2ConnectionChargingInformationOk returns a tuple with the N2ConnectionChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2ConnectionChargingInformation

`func (o *ChargingDataRequest) SetN2ConnectionChargingInformation(v N2ConnectionChargingInformation)`

SetN2ConnectionChargingInformation sets N2ConnectionChargingInformation field to given value.

### HasN2ConnectionChargingInformation

`func (o *ChargingDataRequest) HasN2ConnectionChargingInformation() bool`

HasN2ConnectionChargingInformation returns a boolean if a field has been set.

### GetLocationReportingChargingInformation

`func (o *ChargingDataRequest) GetLocationReportingChargingInformation() LocationReportingChargingInformation`

GetLocationReportingChargingInformation returns the LocationReportingChargingInformation field if non-nil, zero value otherwise.

### GetLocationReportingChargingInformationOk

`func (o *ChargingDataRequest) GetLocationReportingChargingInformationOk() (*LocationReportingChargingInformation, bool)`

GetLocationReportingChargingInformationOk returns a tuple with the LocationReportingChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReportingChargingInformation

`func (o *ChargingDataRequest) SetLocationReportingChargingInformation(v LocationReportingChargingInformation)`

SetLocationReportingChargingInformation sets LocationReportingChargingInformation field to given value.

### HasLocationReportingChargingInformation

`func (o *ChargingDataRequest) HasLocationReportingChargingInformation() bool`

HasLocationReportingChargingInformation returns a boolean if a field has been set.

### GetNSPAChargingInformation

`func (o *ChargingDataRequest) GetNSPAChargingInformation() NSPAChargingInformation`

GetNSPAChargingInformation returns the NSPAChargingInformation field if non-nil, zero value otherwise.

### GetNSPAChargingInformationOk

`func (o *ChargingDataRequest) GetNSPAChargingInformationOk() (*NSPAChargingInformation, bool)`

GetNSPAChargingInformationOk returns a tuple with the NSPAChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSPAChargingInformation

`func (o *ChargingDataRequest) SetNSPAChargingInformation(v NSPAChargingInformation)`

SetNSPAChargingInformation sets NSPAChargingInformation field to given value.

### HasNSPAChargingInformation

`func (o *ChargingDataRequest) HasNSPAChargingInformation() bool`

HasNSPAChargingInformation returns a boolean if a field has been set.

### GetNSMChargingInformation

`func (o *ChargingDataRequest) GetNSMChargingInformation() NSMChargingInformation`

GetNSMChargingInformation returns the NSMChargingInformation field if non-nil, zero value otherwise.

### GetNSMChargingInformationOk

`func (o *ChargingDataRequest) GetNSMChargingInformationOk() (*NSMChargingInformation, bool)`

GetNSMChargingInformationOk returns a tuple with the NSMChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSMChargingInformation

`func (o *ChargingDataRequest) SetNSMChargingInformation(v NSMChargingInformation)`

SetNSMChargingInformation sets NSMChargingInformation field to given value.

### HasNSMChargingInformation

`func (o *ChargingDataRequest) HasNSMChargingInformation() bool`

HasNSMChargingInformation returns a boolean if a field has been set.

### GetMMTelChargingInformation

`func (o *ChargingDataRequest) GetMMTelChargingInformation() MMTelChargingInformation`

GetMMTelChargingInformation returns the MMTelChargingInformation field if non-nil, zero value otherwise.

### GetMMTelChargingInformationOk

`func (o *ChargingDataRequest) GetMMTelChargingInformationOk() (*MMTelChargingInformation, bool)`

GetMMTelChargingInformationOk returns a tuple with the MMTelChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMMTelChargingInformation

`func (o *ChargingDataRequest) SetMMTelChargingInformation(v MMTelChargingInformation)`

SetMMTelChargingInformation sets MMTelChargingInformation field to given value.

### HasMMTelChargingInformation

`func (o *ChargingDataRequest) HasMMTelChargingInformation() bool`

HasMMTelChargingInformation returns a boolean if a field has been set.

### GetIMSChargingInformation

`func (o *ChargingDataRequest) GetIMSChargingInformation() IMSChargingInformation`

GetIMSChargingInformation returns the IMSChargingInformation field if non-nil, zero value otherwise.

### GetIMSChargingInformationOk

`func (o *ChargingDataRequest) GetIMSChargingInformationOk() (*IMSChargingInformation, bool)`

GetIMSChargingInformationOk returns a tuple with the IMSChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIMSChargingInformation

`func (o *ChargingDataRequest) SetIMSChargingInformation(v IMSChargingInformation)`

SetIMSChargingInformation sets IMSChargingInformation field to given value.

### HasIMSChargingInformation

`func (o *ChargingDataRequest) HasIMSChargingInformation() bool`

HasIMSChargingInformation returns a boolean if a field has been set.

### GetEdgeInfrastructureUsageChargingInformation

`func (o *ChargingDataRequest) GetEdgeInfrastructureUsageChargingInformation() EdgeInfrastructureUsageChargingInformation`

GetEdgeInfrastructureUsageChargingInformation returns the EdgeInfrastructureUsageChargingInformation field if non-nil, zero value otherwise.

### GetEdgeInfrastructureUsageChargingInformationOk

`func (o *ChargingDataRequest) GetEdgeInfrastructureUsageChargingInformationOk() (*EdgeInfrastructureUsageChargingInformation, bool)`

GetEdgeInfrastructureUsageChargingInformationOk returns a tuple with the EdgeInfrastructureUsageChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeInfrastructureUsageChargingInformation

`func (o *ChargingDataRequest) SetEdgeInfrastructureUsageChargingInformation(v EdgeInfrastructureUsageChargingInformation)`

SetEdgeInfrastructureUsageChargingInformation sets EdgeInfrastructureUsageChargingInformation field to given value.

### HasEdgeInfrastructureUsageChargingInformation

`func (o *ChargingDataRequest) HasEdgeInfrastructureUsageChargingInformation() bool`

HasEdgeInfrastructureUsageChargingInformation returns a boolean if a field has been set.

### GetEASDeploymentChargingInformation

`func (o *ChargingDataRequest) GetEASDeploymentChargingInformation() EASDeploymentChargingInformation`

GetEASDeploymentChargingInformation returns the EASDeploymentChargingInformation field if non-nil, zero value otherwise.

### GetEASDeploymentChargingInformationOk

`func (o *ChargingDataRequest) GetEASDeploymentChargingInformationOk() (*EASDeploymentChargingInformation, bool)`

GetEASDeploymentChargingInformationOk returns a tuple with the EASDeploymentChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASDeploymentChargingInformation

`func (o *ChargingDataRequest) SetEASDeploymentChargingInformation(v EASDeploymentChargingInformation)`

SetEASDeploymentChargingInformation sets EASDeploymentChargingInformation field to given value.

### HasEASDeploymentChargingInformation

`func (o *ChargingDataRequest) HasEASDeploymentChargingInformation() bool`

HasEASDeploymentChargingInformation returns a boolean if a field has been set.

### GetDirectEdgeEnablingServiceChargingInformation

`func (o *ChargingDataRequest) GetDirectEdgeEnablingServiceChargingInformation() NEFChargingInformation`

GetDirectEdgeEnablingServiceChargingInformation returns the DirectEdgeEnablingServiceChargingInformation field if non-nil, zero value otherwise.

### GetDirectEdgeEnablingServiceChargingInformationOk

`func (o *ChargingDataRequest) GetDirectEdgeEnablingServiceChargingInformationOk() (*NEFChargingInformation, bool)`

GetDirectEdgeEnablingServiceChargingInformationOk returns a tuple with the DirectEdgeEnablingServiceChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectEdgeEnablingServiceChargingInformation

`func (o *ChargingDataRequest) SetDirectEdgeEnablingServiceChargingInformation(v NEFChargingInformation)`

SetDirectEdgeEnablingServiceChargingInformation sets DirectEdgeEnablingServiceChargingInformation field to given value.

### HasDirectEdgeEnablingServiceChargingInformation

`func (o *ChargingDataRequest) HasDirectEdgeEnablingServiceChargingInformation() bool`

HasDirectEdgeEnablingServiceChargingInformation returns a boolean if a field has been set.

### GetExposedEdgeEnablingServiceChargingInformation

`func (o *ChargingDataRequest) GetExposedEdgeEnablingServiceChargingInformation() NEFChargingInformation`

GetExposedEdgeEnablingServiceChargingInformation returns the ExposedEdgeEnablingServiceChargingInformation field if non-nil, zero value otherwise.

### GetExposedEdgeEnablingServiceChargingInformationOk

`func (o *ChargingDataRequest) GetExposedEdgeEnablingServiceChargingInformationOk() (*NEFChargingInformation, bool)`

GetExposedEdgeEnablingServiceChargingInformationOk returns a tuple with the ExposedEdgeEnablingServiceChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedEdgeEnablingServiceChargingInformation

`func (o *ChargingDataRequest) SetExposedEdgeEnablingServiceChargingInformation(v NEFChargingInformation)`

SetExposedEdgeEnablingServiceChargingInformation sets ExposedEdgeEnablingServiceChargingInformation field to given value.

### HasExposedEdgeEnablingServiceChargingInformation

`func (o *ChargingDataRequest) HasExposedEdgeEnablingServiceChargingInformation() bool`

HasExposedEdgeEnablingServiceChargingInformation returns a boolean if a field has been set.

### GetProSeChargingInformation

`func (o *ChargingDataRequest) GetProSeChargingInformation() ProseChargingInformation`

GetProSeChargingInformation returns the ProSeChargingInformation field if non-nil, zero value otherwise.

### GetProSeChargingInformationOk

`func (o *ChargingDataRequest) GetProSeChargingInformationOk() (*ProseChargingInformation, bool)`

GetProSeChargingInformationOk returns a tuple with the ProSeChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProSeChargingInformation

`func (o *ChargingDataRequest) SetProSeChargingInformation(v ProseChargingInformation)`

SetProSeChargingInformation sets ProSeChargingInformation field to given value.

### HasProSeChargingInformation

`func (o *ChargingDataRequest) HasProSeChargingInformation() bool`

HasProSeChargingInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


