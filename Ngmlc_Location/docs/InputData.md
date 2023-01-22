# InputData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**ExtGroupId** | Pointer to **string** | String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.   | [optional] 
**IntGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExternalClientType** | [**ExternalClientType**](ExternalClientType.md) |  | 
**LocationQoS** | Pointer to [**LocationQoS**](LocationQoS.md) |  | [optional] 
**SupportedGADShapes** | Pointer to [**[]SupportedGADShapes**](SupportedGADShapes.md) |  | [optional] 
**ServiceIdentity** | Pointer to **string** | Contains the service identity | [optional] 
**ServiceCoverage** | Pointer to **[]string** |  | [optional] 
**LdrType** | Pointer to [**LdrType**](LdrType.md) |  | [optional] 
**PeriodicEventInfo** | Pointer to [**PeriodicEventInfo**](PeriodicEventInfo.md) |  | [optional] 
**AreaEventInfo** | Pointer to [**AreaEventInfoExt**](AreaEventInfoExt.md) |  | [optional] 
**MotionEventInfo** | Pointer to [**MotionEventInfo**](MotionEventInfo.md) |  | [optional] 
**LdrReference** | Pointer to **string** | LDR Reference. | [optional] 
**HgmlcCallBackUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**EventNotificationUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ExternalClientIdentification** | Pointer to **string** | Contains the external client identification | [optional] 
**AfId** | Pointer to **string** |  | [optional] 
**UePrivacyRequirements** | Pointer to [**UePrivacyRequirements**](UePrivacyRequirements.md) |  | [optional] 
**LcsServiceType** | Pointer to **int32** | LCS service type. | [optional] 
**VelocityRequested** | Pointer to [**VelocityRequested**](VelocityRequested.md) |  | [optional] 
**Priority** | Pointer to [**LcsPriority**](LcsPriority.md) |  | [optional] 
**LocationTypeRequested** | Pointer to [**LocationTypeRequested**](LocationTypeRequested.md) |  | [optional] 
**MaximumAgeOfLocationEstimate** | Pointer to **int32** | Indicates value of the age of the location estimate. | [optional] 
**AmfId** | Pointer to **string** | String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).   | [optional] 
**CodeWord** | Pointer to **string** | Contains the codeword | [optional] 
**ScheduledLocTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ReliableLocReq** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewInputData

`func NewInputData(externalClientType ExternalClientType, ) *InputData`

NewInputData instantiates a new InputData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDataWithDefaults

`func NewInputDataWithDefaults() *InputData`

NewInputDataWithDefaults instantiates a new InputData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *InputData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *InputData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *InputData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *InputData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *InputData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *InputData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *InputData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *InputData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetExtGroupId

`func (o *InputData) GetExtGroupId() string`

GetExtGroupId returns the ExtGroupId field if non-nil, zero value otherwise.

### GetExtGroupIdOk

`func (o *InputData) GetExtGroupIdOk() (*string, bool)`

GetExtGroupIdOk returns a tuple with the ExtGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtGroupId

`func (o *InputData) SetExtGroupId(v string)`

SetExtGroupId sets ExtGroupId field to given value.

### HasExtGroupId

`func (o *InputData) HasExtGroupId() bool`

HasExtGroupId returns a boolean if a field has been set.

### GetIntGroupId

`func (o *InputData) GetIntGroupId() string`

GetIntGroupId returns the IntGroupId field if non-nil, zero value otherwise.

### GetIntGroupIdOk

`func (o *InputData) GetIntGroupIdOk() (*string, bool)`

GetIntGroupIdOk returns a tuple with the IntGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGroupId

`func (o *InputData) SetIntGroupId(v string)`

SetIntGroupId sets IntGroupId field to given value.

### HasIntGroupId

`func (o *InputData) HasIntGroupId() bool`

HasIntGroupId returns a boolean if a field has been set.

### GetExternalClientType

`func (o *InputData) GetExternalClientType() ExternalClientType`

GetExternalClientType returns the ExternalClientType field if non-nil, zero value otherwise.

### GetExternalClientTypeOk

`func (o *InputData) GetExternalClientTypeOk() (*ExternalClientType, bool)`

GetExternalClientTypeOk returns a tuple with the ExternalClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClientType

`func (o *InputData) SetExternalClientType(v ExternalClientType)`

SetExternalClientType sets ExternalClientType field to given value.


### GetLocationQoS

`func (o *InputData) GetLocationQoS() LocationQoS`

GetLocationQoS returns the LocationQoS field if non-nil, zero value otherwise.

### GetLocationQoSOk

`func (o *InputData) GetLocationQoSOk() (*LocationQoS, bool)`

GetLocationQoSOk returns a tuple with the LocationQoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationQoS

`func (o *InputData) SetLocationQoS(v LocationQoS)`

SetLocationQoS sets LocationQoS field to given value.

### HasLocationQoS

`func (o *InputData) HasLocationQoS() bool`

HasLocationQoS returns a boolean if a field has been set.

### GetSupportedGADShapes

`func (o *InputData) GetSupportedGADShapes() []SupportedGADShapes`

GetSupportedGADShapes returns the SupportedGADShapes field if non-nil, zero value otherwise.

### GetSupportedGADShapesOk

`func (o *InputData) GetSupportedGADShapesOk() (*[]SupportedGADShapes, bool)`

GetSupportedGADShapesOk returns a tuple with the SupportedGADShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedGADShapes

`func (o *InputData) SetSupportedGADShapes(v []SupportedGADShapes)`

SetSupportedGADShapes sets SupportedGADShapes field to given value.

### HasSupportedGADShapes

`func (o *InputData) HasSupportedGADShapes() bool`

HasSupportedGADShapes returns a boolean if a field has been set.

### GetServiceIdentity

`func (o *InputData) GetServiceIdentity() string`

GetServiceIdentity returns the ServiceIdentity field if non-nil, zero value otherwise.

### GetServiceIdentityOk

`func (o *InputData) GetServiceIdentityOk() (*string, bool)`

GetServiceIdentityOk returns a tuple with the ServiceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIdentity

`func (o *InputData) SetServiceIdentity(v string)`

SetServiceIdentity sets ServiceIdentity field to given value.

### HasServiceIdentity

`func (o *InputData) HasServiceIdentity() bool`

HasServiceIdentity returns a boolean if a field has been set.

### GetServiceCoverage

`func (o *InputData) GetServiceCoverage() []string`

GetServiceCoverage returns the ServiceCoverage field if non-nil, zero value otherwise.

### GetServiceCoverageOk

`func (o *InputData) GetServiceCoverageOk() (*[]string, bool)`

GetServiceCoverageOk returns a tuple with the ServiceCoverage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCoverage

`func (o *InputData) SetServiceCoverage(v []string)`

SetServiceCoverage sets ServiceCoverage field to given value.

### HasServiceCoverage

`func (o *InputData) HasServiceCoverage() bool`

HasServiceCoverage returns a boolean if a field has been set.

### GetLdrType

`func (o *InputData) GetLdrType() LdrType`

GetLdrType returns the LdrType field if non-nil, zero value otherwise.

### GetLdrTypeOk

`func (o *InputData) GetLdrTypeOk() (*LdrType, bool)`

GetLdrTypeOk returns a tuple with the LdrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrType

`func (o *InputData) SetLdrType(v LdrType)`

SetLdrType sets LdrType field to given value.

### HasLdrType

`func (o *InputData) HasLdrType() bool`

HasLdrType returns a boolean if a field has been set.

### GetPeriodicEventInfo

`func (o *InputData) GetPeriodicEventInfo() PeriodicEventInfo`

GetPeriodicEventInfo returns the PeriodicEventInfo field if non-nil, zero value otherwise.

### GetPeriodicEventInfoOk

`func (o *InputData) GetPeriodicEventInfoOk() (*PeriodicEventInfo, bool)`

GetPeriodicEventInfoOk returns a tuple with the PeriodicEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicEventInfo

`func (o *InputData) SetPeriodicEventInfo(v PeriodicEventInfo)`

SetPeriodicEventInfo sets PeriodicEventInfo field to given value.

### HasPeriodicEventInfo

`func (o *InputData) HasPeriodicEventInfo() bool`

HasPeriodicEventInfo returns a boolean if a field has been set.

### GetAreaEventInfo

`func (o *InputData) GetAreaEventInfo() AreaEventInfoExt`

GetAreaEventInfo returns the AreaEventInfo field if non-nil, zero value otherwise.

### GetAreaEventInfoOk

`func (o *InputData) GetAreaEventInfoOk() (*AreaEventInfoExt, bool)`

GetAreaEventInfoOk returns a tuple with the AreaEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaEventInfo

`func (o *InputData) SetAreaEventInfo(v AreaEventInfoExt)`

SetAreaEventInfo sets AreaEventInfo field to given value.

### HasAreaEventInfo

`func (o *InputData) HasAreaEventInfo() bool`

HasAreaEventInfo returns a boolean if a field has been set.

### GetMotionEventInfo

`func (o *InputData) GetMotionEventInfo() MotionEventInfo`

GetMotionEventInfo returns the MotionEventInfo field if non-nil, zero value otherwise.

### GetMotionEventInfoOk

`func (o *InputData) GetMotionEventInfoOk() (*MotionEventInfo, bool)`

GetMotionEventInfoOk returns a tuple with the MotionEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionEventInfo

`func (o *InputData) SetMotionEventInfo(v MotionEventInfo)`

SetMotionEventInfo sets MotionEventInfo field to given value.

### HasMotionEventInfo

`func (o *InputData) HasMotionEventInfo() bool`

HasMotionEventInfo returns a boolean if a field has been set.

### GetLdrReference

`func (o *InputData) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *InputData) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *InputData) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.

### HasLdrReference

`func (o *InputData) HasLdrReference() bool`

HasLdrReference returns a boolean if a field has been set.

### GetHgmlcCallBackUri

`func (o *InputData) GetHgmlcCallBackUri() string`

GetHgmlcCallBackUri returns the HgmlcCallBackUri field if non-nil, zero value otherwise.

### GetHgmlcCallBackUriOk

`func (o *InputData) GetHgmlcCallBackUriOk() (*string, bool)`

GetHgmlcCallBackUriOk returns a tuple with the HgmlcCallBackUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackUri

`func (o *InputData) SetHgmlcCallBackUri(v string)`

SetHgmlcCallBackUri sets HgmlcCallBackUri field to given value.

### HasHgmlcCallBackUri

`func (o *InputData) HasHgmlcCallBackUri() bool`

HasHgmlcCallBackUri returns a boolean if a field has been set.

### GetEventNotificationUri

`func (o *InputData) GetEventNotificationUri() string`

GetEventNotificationUri returns the EventNotificationUri field if non-nil, zero value otherwise.

### GetEventNotificationUriOk

`func (o *InputData) GetEventNotificationUriOk() (*string, bool)`

GetEventNotificationUriOk returns a tuple with the EventNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotificationUri

`func (o *InputData) SetEventNotificationUri(v string)`

SetEventNotificationUri sets EventNotificationUri field to given value.

### HasEventNotificationUri

`func (o *InputData) HasEventNotificationUri() bool`

HasEventNotificationUri returns a boolean if a field has been set.

### GetExternalClientIdentification

`func (o *InputData) GetExternalClientIdentification() string`

GetExternalClientIdentification returns the ExternalClientIdentification field if non-nil, zero value otherwise.

### GetExternalClientIdentificationOk

`func (o *InputData) GetExternalClientIdentificationOk() (*string, bool)`

GetExternalClientIdentificationOk returns a tuple with the ExternalClientIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClientIdentification

`func (o *InputData) SetExternalClientIdentification(v string)`

SetExternalClientIdentification sets ExternalClientIdentification field to given value.

### HasExternalClientIdentification

`func (o *InputData) HasExternalClientIdentification() bool`

HasExternalClientIdentification returns a boolean if a field has been set.

### GetAfId

`func (o *InputData) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *InputData) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *InputData) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *InputData) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetUePrivacyRequirements

`func (o *InputData) GetUePrivacyRequirements() UePrivacyRequirements`

GetUePrivacyRequirements returns the UePrivacyRequirements field if non-nil, zero value otherwise.

### GetUePrivacyRequirementsOk

`func (o *InputData) GetUePrivacyRequirementsOk() (*UePrivacyRequirements, bool)`

GetUePrivacyRequirementsOk returns a tuple with the UePrivacyRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePrivacyRequirements

`func (o *InputData) SetUePrivacyRequirements(v UePrivacyRequirements)`

SetUePrivacyRequirements sets UePrivacyRequirements field to given value.

### HasUePrivacyRequirements

`func (o *InputData) HasUePrivacyRequirements() bool`

HasUePrivacyRequirements returns a boolean if a field has been set.

### GetLcsServiceType

`func (o *InputData) GetLcsServiceType() int32`

GetLcsServiceType returns the LcsServiceType field if non-nil, zero value otherwise.

### GetLcsServiceTypeOk

`func (o *InputData) GetLcsServiceTypeOk() (*int32, bool)`

GetLcsServiceTypeOk returns a tuple with the LcsServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsServiceType

`func (o *InputData) SetLcsServiceType(v int32)`

SetLcsServiceType sets LcsServiceType field to given value.

### HasLcsServiceType

`func (o *InputData) HasLcsServiceType() bool`

HasLcsServiceType returns a boolean if a field has been set.

### GetVelocityRequested

`func (o *InputData) GetVelocityRequested() VelocityRequested`

GetVelocityRequested returns the VelocityRequested field if non-nil, zero value otherwise.

### GetVelocityRequestedOk

`func (o *InputData) GetVelocityRequestedOk() (*VelocityRequested, bool)`

GetVelocityRequestedOk returns a tuple with the VelocityRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityRequested

`func (o *InputData) SetVelocityRequested(v VelocityRequested)`

SetVelocityRequested sets VelocityRequested field to given value.

### HasVelocityRequested

`func (o *InputData) HasVelocityRequested() bool`

HasVelocityRequested returns a boolean if a field has been set.

### GetPriority

`func (o *InputData) GetPriority() LcsPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InputData) GetPriorityOk() (*LcsPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InputData) SetPriority(v LcsPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InputData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetLocationTypeRequested

`func (o *InputData) GetLocationTypeRequested() LocationTypeRequested`

GetLocationTypeRequested returns the LocationTypeRequested field if non-nil, zero value otherwise.

### GetLocationTypeRequestedOk

`func (o *InputData) GetLocationTypeRequestedOk() (*LocationTypeRequested, bool)`

GetLocationTypeRequestedOk returns a tuple with the LocationTypeRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationTypeRequested

`func (o *InputData) SetLocationTypeRequested(v LocationTypeRequested)`

SetLocationTypeRequested sets LocationTypeRequested field to given value.

### HasLocationTypeRequested

`func (o *InputData) HasLocationTypeRequested() bool`

HasLocationTypeRequested returns a boolean if a field has been set.

### GetMaximumAgeOfLocationEstimate

`func (o *InputData) GetMaximumAgeOfLocationEstimate() int32`

GetMaximumAgeOfLocationEstimate returns the MaximumAgeOfLocationEstimate field if non-nil, zero value otherwise.

### GetMaximumAgeOfLocationEstimateOk

`func (o *InputData) GetMaximumAgeOfLocationEstimateOk() (*int32, bool)`

GetMaximumAgeOfLocationEstimateOk returns a tuple with the MaximumAgeOfLocationEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumAgeOfLocationEstimate

`func (o *InputData) SetMaximumAgeOfLocationEstimate(v int32)`

SetMaximumAgeOfLocationEstimate sets MaximumAgeOfLocationEstimate field to given value.

### HasMaximumAgeOfLocationEstimate

`func (o *InputData) HasMaximumAgeOfLocationEstimate() bool`

HasMaximumAgeOfLocationEstimate returns a boolean if a field has been set.

### GetAmfId

`func (o *InputData) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *InputData) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *InputData) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.

### HasAmfId

`func (o *InputData) HasAmfId() bool`

HasAmfId returns a boolean if a field has been set.

### GetCodeWord

`func (o *InputData) GetCodeWord() string`

GetCodeWord returns the CodeWord field if non-nil, zero value otherwise.

### GetCodeWordOk

`func (o *InputData) GetCodeWordOk() (*string, bool)`

GetCodeWordOk returns a tuple with the CodeWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeWord

`func (o *InputData) SetCodeWord(v string)`

SetCodeWord sets CodeWord field to given value.

### HasCodeWord

`func (o *InputData) HasCodeWord() bool`

HasCodeWord returns a boolean if a field has been set.

### GetScheduledLocTime

`func (o *InputData) GetScheduledLocTime() time.Time`

GetScheduledLocTime returns the ScheduledLocTime field if non-nil, zero value otherwise.

### GetScheduledLocTimeOk

`func (o *InputData) GetScheduledLocTimeOk() (*time.Time, bool)`

GetScheduledLocTimeOk returns a tuple with the ScheduledLocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledLocTime

`func (o *InputData) SetScheduledLocTime(v time.Time)`

SetScheduledLocTime sets ScheduledLocTime field to given value.

### HasScheduledLocTime

`func (o *InputData) HasScheduledLocTime() bool`

HasScheduledLocTime returns a boolean if a field has been set.

### GetReliableLocReq

`func (o *InputData) GetReliableLocReq() bool`

GetReliableLocReq returns the ReliableLocReq field if non-nil, zero value otherwise.

### GetReliableLocReqOk

`func (o *InputData) GetReliableLocReqOk() (*bool, bool)`

GetReliableLocReqOk returns a tuple with the ReliableLocReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableLocReq

`func (o *InputData) SetReliableLocReq(v bool)`

SetReliableLocReq sets ReliableLocReq field to given value.

### HasReliableLocReq

`func (o *InputData) HasReliableLocReq() bool`

HasReliableLocReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


