# RequestPosInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LcsClientType** | [**ExternalClientType**](ExternalClientType.md) |  | 
**LcsLocation** | [**LocationType**](LocationType.md) |  | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Priority** | Pointer to [**LcsPriority**](LcsPriority.md) |  | [optional] 
**LcsQoS** | Pointer to [**LocationQoS**](LocationQoS.md) |  | [optional] 
**VelocityRequested** | Pointer to [**VelocityRequested**](VelocityRequested.md) |  | [optional] 
**LcsSupportedGADShapes** | Pointer to [**SupportedGADShapes**](SupportedGADShapes.md) |  | [optional] 
**AdditionalLcsSuppGADShapes** | Pointer to [**[]SupportedGADShapes**](SupportedGADShapes.md) |  | [optional] 
**LocationNotificationUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**LcsServiceType** | Pointer to **int32** | LCS service type. | [optional] 
**LdrType** | Pointer to [**LdrType**](LdrType.md) |  | [optional] 
**HgmlcCallBackURI** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**LdrReference** | Pointer to **string** | LDR Reference. | [optional] 
**PeriodicEventInfo** | Pointer to [**PeriodicEventInfo**](PeriodicEventInfo.md) |  | [optional] 
**AreaEventInfo** | Pointer to [**AreaEventInfo**](AreaEventInfo.md) |  | [optional] 
**MotionEventInfo** | Pointer to [**MotionEventInfo**](MotionEventInfo.md) |  | [optional] 
**ExternalClientIdentification** | Pointer to **string** | Contains the external client identification | [optional] 
**AfID** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**CodeWord** | Pointer to **string** | Contains the codeword | [optional] 
**UePrivacyRequirements** | Pointer to [**UePrivacyRequirements**](UePrivacyRequirements.md) |  | [optional] 
**ScheduledLocTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ReliableLocReq** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRequestPosInfo

`func NewRequestPosInfo(lcsClientType ExternalClientType, lcsLocation LocationType, ) *RequestPosInfo`

NewRequestPosInfo instantiates a new RequestPosInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPosInfoWithDefaults

`func NewRequestPosInfoWithDefaults() *RequestPosInfo`

NewRequestPosInfoWithDefaults instantiates a new RequestPosInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLcsClientType

`func (o *RequestPosInfo) GetLcsClientType() ExternalClientType`

GetLcsClientType returns the LcsClientType field if non-nil, zero value otherwise.

### GetLcsClientTypeOk

`func (o *RequestPosInfo) GetLcsClientTypeOk() (*ExternalClientType, bool)`

GetLcsClientTypeOk returns a tuple with the LcsClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsClientType

`func (o *RequestPosInfo) SetLcsClientType(v ExternalClientType)`

SetLcsClientType sets LcsClientType field to given value.


### GetLcsLocation

`func (o *RequestPosInfo) GetLcsLocation() LocationType`

GetLcsLocation returns the LcsLocation field if non-nil, zero value otherwise.

### GetLcsLocationOk

`func (o *RequestPosInfo) GetLcsLocationOk() (*LocationType, bool)`

GetLcsLocationOk returns a tuple with the LcsLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsLocation

`func (o *RequestPosInfo) SetLcsLocation(v LocationType)`

SetLcsLocation sets LcsLocation field to given value.


### GetSupi

`func (o *RequestPosInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *RequestPosInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *RequestPosInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *RequestPosInfo) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGpsi

`func (o *RequestPosInfo) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *RequestPosInfo) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *RequestPosInfo) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *RequestPosInfo) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPriority

`func (o *RequestPosInfo) GetPriority() LcsPriority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RequestPosInfo) GetPriorityOk() (*LcsPriority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RequestPosInfo) SetPriority(v LcsPriority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RequestPosInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetLcsQoS

`func (o *RequestPosInfo) GetLcsQoS() LocationQoS`

GetLcsQoS returns the LcsQoS field if non-nil, zero value otherwise.

### GetLcsQoSOk

`func (o *RequestPosInfo) GetLcsQoSOk() (*LocationQoS, bool)`

GetLcsQoSOk returns a tuple with the LcsQoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsQoS

`func (o *RequestPosInfo) SetLcsQoS(v LocationQoS)`

SetLcsQoS sets LcsQoS field to given value.

### HasLcsQoS

`func (o *RequestPosInfo) HasLcsQoS() bool`

HasLcsQoS returns a boolean if a field has been set.

### GetVelocityRequested

`func (o *RequestPosInfo) GetVelocityRequested() VelocityRequested`

GetVelocityRequested returns the VelocityRequested field if non-nil, zero value otherwise.

### GetVelocityRequestedOk

`func (o *RequestPosInfo) GetVelocityRequestedOk() (*VelocityRequested, bool)`

GetVelocityRequestedOk returns a tuple with the VelocityRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocityRequested

`func (o *RequestPosInfo) SetVelocityRequested(v VelocityRequested)`

SetVelocityRequested sets VelocityRequested field to given value.

### HasVelocityRequested

`func (o *RequestPosInfo) HasVelocityRequested() bool`

HasVelocityRequested returns a boolean if a field has been set.

### GetLcsSupportedGADShapes

`func (o *RequestPosInfo) GetLcsSupportedGADShapes() SupportedGADShapes`

GetLcsSupportedGADShapes returns the LcsSupportedGADShapes field if non-nil, zero value otherwise.

### GetLcsSupportedGADShapesOk

`func (o *RequestPosInfo) GetLcsSupportedGADShapesOk() (*SupportedGADShapes, bool)`

GetLcsSupportedGADShapesOk returns a tuple with the LcsSupportedGADShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsSupportedGADShapes

`func (o *RequestPosInfo) SetLcsSupportedGADShapes(v SupportedGADShapes)`

SetLcsSupportedGADShapes sets LcsSupportedGADShapes field to given value.

### HasLcsSupportedGADShapes

`func (o *RequestPosInfo) HasLcsSupportedGADShapes() bool`

HasLcsSupportedGADShapes returns a boolean if a field has been set.

### GetAdditionalLcsSuppGADShapes

`func (o *RequestPosInfo) GetAdditionalLcsSuppGADShapes() []SupportedGADShapes`

GetAdditionalLcsSuppGADShapes returns the AdditionalLcsSuppGADShapes field if non-nil, zero value otherwise.

### GetAdditionalLcsSuppGADShapesOk

`func (o *RequestPosInfo) GetAdditionalLcsSuppGADShapesOk() (*[]SupportedGADShapes, bool)`

GetAdditionalLcsSuppGADShapesOk returns a tuple with the AdditionalLcsSuppGADShapes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalLcsSuppGADShapes

`func (o *RequestPosInfo) SetAdditionalLcsSuppGADShapes(v []SupportedGADShapes)`

SetAdditionalLcsSuppGADShapes sets AdditionalLcsSuppGADShapes field to given value.

### HasAdditionalLcsSuppGADShapes

`func (o *RequestPosInfo) HasAdditionalLcsSuppGADShapes() bool`

HasAdditionalLcsSuppGADShapes returns a boolean if a field has been set.

### GetLocationNotificationUri

`func (o *RequestPosInfo) GetLocationNotificationUri() string`

GetLocationNotificationUri returns the LocationNotificationUri field if non-nil, zero value otherwise.

### GetLocationNotificationUriOk

`func (o *RequestPosInfo) GetLocationNotificationUriOk() (*string, bool)`

GetLocationNotificationUriOk returns a tuple with the LocationNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNotificationUri

`func (o *RequestPosInfo) SetLocationNotificationUri(v string)`

SetLocationNotificationUri sets LocationNotificationUri field to given value.

### HasLocationNotificationUri

`func (o *RequestPosInfo) HasLocationNotificationUri() bool`

HasLocationNotificationUri returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *RequestPosInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *RequestPosInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *RequestPosInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *RequestPosInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetOldGuami

`func (o *RequestPosInfo) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *RequestPosInfo) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *RequestPosInfo) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *RequestPosInfo) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.

### GetPei

`func (o *RequestPosInfo) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *RequestPosInfo) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *RequestPosInfo) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *RequestPosInfo) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetLcsServiceType

`func (o *RequestPosInfo) GetLcsServiceType() int32`

GetLcsServiceType returns the LcsServiceType field if non-nil, zero value otherwise.

### GetLcsServiceTypeOk

`func (o *RequestPosInfo) GetLcsServiceTypeOk() (*int32, bool)`

GetLcsServiceTypeOk returns a tuple with the LcsServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsServiceType

`func (o *RequestPosInfo) SetLcsServiceType(v int32)`

SetLcsServiceType sets LcsServiceType field to given value.

### HasLcsServiceType

`func (o *RequestPosInfo) HasLcsServiceType() bool`

HasLcsServiceType returns a boolean if a field has been set.

### GetLdrType

`func (o *RequestPosInfo) GetLdrType() LdrType`

GetLdrType returns the LdrType field if non-nil, zero value otherwise.

### GetLdrTypeOk

`func (o *RequestPosInfo) GetLdrTypeOk() (*LdrType, bool)`

GetLdrTypeOk returns a tuple with the LdrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrType

`func (o *RequestPosInfo) SetLdrType(v LdrType)`

SetLdrType sets LdrType field to given value.

### HasLdrType

`func (o *RequestPosInfo) HasLdrType() bool`

HasLdrType returns a boolean if a field has been set.

### GetHgmlcCallBackURI

`func (o *RequestPosInfo) GetHgmlcCallBackURI() string`

GetHgmlcCallBackURI returns the HgmlcCallBackURI field if non-nil, zero value otherwise.

### GetHgmlcCallBackURIOk

`func (o *RequestPosInfo) GetHgmlcCallBackURIOk() (*string, bool)`

GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackURI

`func (o *RequestPosInfo) SetHgmlcCallBackURI(v string)`

SetHgmlcCallBackURI sets HgmlcCallBackURI field to given value.

### HasHgmlcCallBackURI

`func (o *RequestPosInfo) HasHgmlcCallBackURI() bool`

HasHgmlcCallBackURI returns a boolean if a field has been set.

### GetLdrReference

`func (o *RequestPosInfo) GetLdrReference() string`

GetLdrReference returns the LdrReference field if non-nil, zero value otherwise.

### GetLdrReferenceOk

`func (o *RequestPosInfo) GetLdrReferenceOk() (*string, bool)`

GetLdrReferenceOk returns a tuple with the LdrReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrReference

`func (o *RequestPosInfo) SetLdrReference(v string)`

SetLdrReference sets LdrReference field to given value.

### HasLdrReference

`func (o *RequestPosInfo) HasLdrReference() bool`

HasLdrReference returns a boolean if a field has been set.

### GetPeriodicEventInfo

`func (o *RequestPosInfo) GetPeriodicEventInfo() PeriodicEventInfo`

GetPeriodicEventInfo returns the PeriodicEventInfo field if non-nil, zero value otherwise.

### GetPeriodicEventInfoOk

`func (o *RequestPosInfo) GetPeriodicEventInfoOk() (*PeriodicEventInfo, bool)`

GetPeriodicEventInfoOk returns a tuple with the PeriodicEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicEventInfo

`func (o *RequestPosInfo) SetPeriodicEventInfo(v PeriodicEventInfo)`

SetPeriodicEventInfo sets PeriodicEventInfo field to given value.

### HasPeriodicEventInfo

`func (o *RequestPosInfo) HasPeriodicEventInfo() bool`

HasPeriodicEventInfo returns a boolean if a field has been set.

### GetAreaEventInfo

`func (o *RequestPosInfo) GetAreaEventInfo() AreaEventInfo`

GetAreaEventInfo returns the AreaEventInfo field if non-nil, zero value otherwise.

### GetAreaEventInfoOk

`func (o *RequestPosInfo) GetAreaEventInfoOk() (*AreaEventInfo, bool)`

GetAreaEventInfoOk returns a tuple with the AreaEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaEventInfo

`func (o *RequestPosInfo) SetAreaEventInfo(v AreaEventInfo)`

SetAreaEventInfo sets AreaEventInfo field to given value.

### HasAreaEventInfo

`func (o *RequestPosInfo) HasAreaEventInfo() bool`

HasAreaEventInfo returns a boolean if a field has been set.

### GetMotionEventInfo

`func (o *RequestPosInfo) GetMotionEventInfo() MotionEventInfo`

GetMotionEventInfo returns the MotionEventInfo field if non-nil, zero value otherwise.

### GetMotionEventInfoOk

`func (o *RequestPosInfo) GetMotionEventInfoOk() (*MotionEventInfo, bool)`

GetMotionEventInfoOk returns a tuple with the MotionEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotionEventInfo

`func (o *RequestPosInfo) SetMotionEventInfo(v MotionEventInfo)`

SetMotionEventInfo sets MotionEventInfo field to given value.

### HasMotionEventInfo

`func (o *RequestPosInfo) HasMotionEventInfo() bool`

HasMotionEventInfo returns a boolean if a field has been set.

### GetExternalClientIdentification

`func (o *RequestPosInfo) GetExternalClientIdentification() string`

GetExternalClientIdentification returns the ExternalClientIdentification field if non-nil, zero value otherwise.

### GetExternalClientIdentificationOk

`func (o *RequestPosInfo) GetExternalClientIdentificationOk() (*string, bool)`

GetExternalClientIdentificationOk returns a tuple with the ExternalClientIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalClientIdentification

`func (o *RequestPosInfo) SetExternalClientIdentification(v string)`

SetExternalClientIdentification sets ExternalClientIdentification field to given value.

### HasExternalClientIdentification

`func (o *RequestPosInfo) HasExternalClientIdentification() bool`

HasExternalClientIdentification returns a boolean if a field has been set.

### GetAfID

`func (o *RequestPosInfo) GetAfID() string`

GetAfID returns the AfID field if non-nil, zero value otherwise.

### GetAfIDOk

`func (o *RequestPosInfo) GetAfIDOk() (*string, bool)`

GetAfIDOk returns a tuple with the AfID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfID

`func (o *RequestPosInfo) SetAfID(v string)`

SetAfID sets AfID field to given value.

### HasAfID

`func (o *RequestPosInfo) HasAfID() bool`

HasAfID returns a boolean if a field has been set.

### GetCodeWord

`func (o *RequestPosInfo) GetCodeWord() string`

GetCodeWord returns the CodeWord field if non-nil, zero value otherwise.

### GetCodeWordOk

`func (o *RequestPosInfo) GetCodeWordOk() (*string, bool)`

GetCodeWordOk returns a tuple with the CodeWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeWord

`func (o *RequestPosInfo) SetCodeWord(v string)`

SetCodeWord sets CodeWord field to given value.

### HasCodeWord

`func (o *RequestPosInfo) HasCodeWord() bool`

HasCodeWord returns a boolean if a field has been set.

### GetUePrivacyRequirements

`func (o *RequestPosInfo) GetUePrivacyRequirements() UePrivacyRequirements`

GetUePrivacyRequirements returns the UePrivacyRequirements field if non-nil, zero value otherwise.

### GetUePrivacyRequirementsOk

`func (o *RequestPosInfo) GetUePrivacyRequirementsOk() (*UePrivacyRequirements, bool)`

GetUePrivacyRequirementsOk returns a tuple with the UePrivacyRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePrivacyRequirements

`func (o *RequestPosInfo) SetUePrivacyRequirements(v UePrivacyRequirements)`

SetUePrivacyRequirements sets UePrivacyRequirements field to given value.

### HasUePrivacyRequirements

`func (o *RequestPosInfo) HasUePrivacyRequirements() bool`

HasUePrivacyRequirements returns a boolean if a field has been set.

### GetScheduledLocTime

`func (o *RequestPosInfo) GetScheduledLocTime() time.Time`

GetScheduledLocTime returns the ScheduledLocTime field if non-nil, zero value otherwise.

### GetScheduledLocTimeOk

`func (o *RequestPosInfo) GetScheduledLocTimeOk() (*time.Time, bool)`

GetScheduledLocTimeOk returns a tuple with the ScheduledLocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledLocTime

`func (o *RequestPosInfo) SetScheduledLocTime(v time.Time)`

SetScheduledLocTime sets ScheduledLocTime field to given value.

### HasScheduledLocTime

`func (o *RequestPosInfo) HasScheduledLocTime() bool`

HasScheduledLocTime returns a boolean if a field has been set.

### GetReliableLocReq

`func (o *RequestPosInfo) GetReliableLocReq() bool`

GetReliableLocReq returns the ReliableLocReq field if non-nil, zero value otherwise.

### GetReliableLocReqOk

`func (o *RequestPosInfo) GetReliableLocReqOk() (*bool, bool)`

GetReliableLocReqOk returns a tuple with the ReliableLocReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableLocReq

`func (o *RequestPosInfo) SetReliableLocReq(v bool)`

SetReliableLocReq sets ReliableLocReq field to given value.

### HasReliableLocReq

`func (o *RequestPosInfo) HasReliableLocReq() bool`

HasReliableLocReq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


