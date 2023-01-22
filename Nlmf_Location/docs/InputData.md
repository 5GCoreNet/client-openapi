# InputData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalClientType** | Pointer to [**ExternalClientType**](ExternalClientType.md) |  | [optional] 
**CorrelationID** | Pointer to **string** | LCS Correlation ID. | [optional] 
**AmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**LocationQoS** | Pointer to [**LocationQoS**](LocationQoS.md) |  | [optional] 
**SupportedGADShapes** | Pointer to [**[]SupportedGADShapes**](SupportedGADShapes.md) |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**EcgiOnSecondNode** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**NcgiOnSecondNode** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**Priority** | Pointer to [**LcsPriority**](LcsPriority.md) |  | [optional] 
**VelocityRequested** | Pointer to [**VelocityRequested**](VelocityRequested.md) |  | [optional] 
**UeLcsCap** | Pointer to [**UeLcsCapability**](UeLcsCapability.md) |  | [optional] 
**LcsServiceType** | Pointer to **int32** | LCS service type. | [optional] 
**LdrType** | Pointer to [**LdrType**](LdrType.md) |  | [optional] 
**HgmlcCallBackURI** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**VgmlcAddress** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**LdrReference** | Pointer to **string** | LDR Reference. | [optional] 
**PeriodicEventInfo** | Pointer to [**PeriodicEventInfo**](PeriodicEventInfo.md) |  | [optional] 
**AreaEventInfo** | Pointer to [**AreaEventInfo**](AreaEventInfo.md) |  | [optional] 
**MotionEventInfo** | Pointer to [**MotionEventInfo**](MotionEventInfo.md) |  | [optional] 
**ReportingAccessTypes** | Pointer to [**[]ReportingAccessType**](ReportingAccessType.md) |  | [optional] 
**UeConnectivityStates** | Pointer to [**UeConnectivityState**](UeConnectivityState.md) |  | [optional] 
**UeLocationServiceInd** | Pointer to [**UeLocationServiceInd**](UeLocationServiceInd.md) |  | [optional] 
**MoAssistanceDataTypes** | Pointer to [**LcsBroadcastAssistanceTypesData**](LcsBroadcastAssistanceTypesData.md) |  | [optional] 
**LppMessage** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**LppMessageExt** | Pointer to [**[]RefToBinaryData**](RefToBinaryData.md) | Indicates the lpp message extension. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**UePositioningCap** | Pointer to **string** | Positioning capabilities supported by the UE. A string encoding the \&quot;ProvideCapabilities-r9-IEs\&quot; IE as specified in clause 6.3 of 3GPP TS 37.355 (start from octet 1). | [optional] 
**TnapId** | Pointer to [**TnapId**](TnapId.md) |  | [optional] 
**TwapId** | Pointer to [**TwapId**](TwapId.md) |  | [optional] 
**UeCountryDetInd** | Pointer to **bool** |  | [optional] 
**ScheduledLocTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ReliableLocReq** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewInputData

`func NewInputData() *InputData`

NewInputData instantiates a new InputData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputDataWithDefaults

`func NewInputDataWithDefaults() *InputData`

NewInputDataWithDefaults instantiates a new InputData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasExternalClientType

`func (o *InputData) HasExternalClientType() bool`

HasExternalClientType returns a boolean if a field has been set.

### GetCorrelationID

`func (o *InputData) GetCorrelationID() string`

GetCorrelationID returns the CorrelationID field if non-nil, zero value otherwise.

### GetCorrelationIDOk

`func (o *InputData) GetCorrelationIDOk() (*string, bool)`

GetCorrelationIDOk returns a tuple with the CorrelationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationID

`func (o *InputData) SetCorrelationID(v string)`

SetCorrelationID sets CorrelationID field to given value.

### HasCorrelationID

`func (o *InputData) HasCorrelationID() bool`

HasCorrelationID returns a boolean if a field has been set.

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

### GetPei

`func (o *InputData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *InputData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *InputData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *InputData) HasPei() bool`

HasPei returns a boolean if a field has been set.

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

### GetEcgi

`func (o *InputData) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *InputData) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *InputData) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *InputData) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetEcgiOnSecondNode

`func (o *InputData) GetEcgiOnSecondNode() Ecgi`

GetEcgiOnSecondNode returns the EcgiOnSecondNode field if non-nil, zero value otherwise.

### GetEcgiOnSecondNodeOk

`func (o *InputData) GetEcgiOnSecondNodeOk() (*Ecgi, bool)`

GetEcgiOnSecondNodeOk returns a tuple with the EcgiOnSecondNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgiOnSecondNode

`func (o *InputData) SetEcgiOnSecondNode(v Ecgi)`

SetEcgiOnSecondNode sets EcgiOnSecondNode field to given value.

### HasEcgiOnSecondNode

`func (o *InputData) HasEcgiOnSecondNode() bool`

HasEcgiOnSecondNode returns a boolean if a field has been set.

### GetNcgi

`func (o *InputData) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *InputData) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *InputData) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *InputData) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetNcgiOnSecondNode

`func (o *InputData) GetNcgiOnSecondNode() Ncgi`

GetNcgiOnSecondNode returns the NcgiOnSecondNode field if non-nil, zero value otherwise.

### GetNcgiOnSecondNodeOk

`func (o *InputData) GetNcgiOnSecondNodeOk() (*Ncgi, bool)`

GetNcgiOnSecondNodeOk returns a tuple with the NcgiOnSecondNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgiOnSecondNode

`func (o *InputData) SetNcgiOnSecondNode(v Ncgi)`

SetNcgiOnSecondNode sets NcgiOnSecondNode field to given value.

### HasNcgiOnSecondNode

`func (o *InputData) HasNcgiOnSecondNode() bool`

HasNcgiOnSecondNode returns a boolean if a field has been set.

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

### GetUeLcsCap

`func (o *InputData) GetUeLcsCap() UeLcsCapability`

GetUeLcsCap returns the UeLcsCap field if non-nil, zero value otherwise.

### GetUeLcsCapOk

`func (o *InputData) GetUeLcsCapOk() (*UeLcsCapability, bool)`

GetUeLcsCapOk returns a tuple with the UeLcsCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLcsCap

`func (o *InputData) SetUeLcsCap(v UeLcsCapability)`

SetUeLcsCap sets UeLcsCap field to given value.

### HasUeLcsCap

`func (o *InputData) HasUeLcsCap() bool`

HasUeLcsCap returns a boolean if a field has been set.

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

### GetHgmlcCallBackURI

`func (o *InputData) GetHgmlcCallBackURI() string`

GetHgmlcCallBackURI returns the HgmlcCallBackURI field if non-nil, zero value otherwise.

### GetHgmlcCallBackURIOk

`func (o *InputData) GetHgmlcCallBackURIOk() (*string, bool)`

GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgmlcCallBackURI

`func (o *InputData) SetHgmlcCallBackURI(v string)`

SetHgmlcCallBackURI sets HgmlcCallBackURI field to given value.

### HasHgmlcCallBackURI

`func (o *InputData) HasHgmlcCallBackURI() bool`

HasHgmlcCallBackURI returns a boolean if a field has been set.

### GetVgmlcAddress

`func (o *InputData) GetVgmlcAddress() string`

GetVgmlcAddress returns the VgmlcAddress field if non-nil, zero value otherwise.

### GetVgmlcAddressOk

`func (o *InputData) GetVgmlcAddressOk() (*string, bool)`

GetVgmlcAddressOk returns a tuple with the VgmlcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgmlcAddress

`func (o *InputData) SetVgmlcAddress(v string)`

SetVgmlcAddress sets VgmlcAddress field to given value.

### HasVgmlcAddress

`func (o *InputData) HasVgmlcAddress() bool`

HasVgmlcAddress returns a boolean if a field has been set.

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

`func (o *InputData) GetAreaEventInfo() AreaEventInfo`

GetAreaEventInfo returns the AreaEventInfo field if non-nil, zero value otherwise.

### GetAreaEventInfoOk

`func (o *InputData) GetAreaEventInfoOk() (*AreaEventInfo, bool)`

GetAreaEventInfoOk returns a tuple with the AreaEventInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaEventInfo

`func (o *InputData) SetAreaEventInfo(v AreaEventInfo)`

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

### GetReportingAccessTypes

`func (o *InputData) GetReportingAccessTypes() []ReportingAccessType`

GetReportingAccessTypes returns the ReportingAccessTypes field if non-nil, zero value otherwise.

### GetReportingAccessTypesOk

`func (o *InputData) GetReportingAccessTypesOk() (*[]ReportingAccessType, bool)`

GetReportingAccessTypesOk returns a tuple with the ReportingAccessTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingAccessTypes

`func (o *InputData) SetReportingAccessTypes(v []ReportingAccessType)`

SetReportingAccessTypes sets ReportingAccessTypes field to given value.

### HasReportingAccessTypes

`func (o *InputData) HasReportingAccessTypes() bool`

HasReportingAccessTypes returns a boolean if a field has been set.

### GetUeConnectivityStates

`func (o *InputData) GetUeConnectivityStates() UeConnectivityState`

GetUeConnectivityStates returns the UeConnectivityStates field if non-nil, zero value otherwise.

### GetUeConnectivityStatesOk

`func (o *InputData) GetUeConnectivityStatesOk() (*UeConnectivityState, bool)`

GetUeConnectivityStatesOk returns a tuple with the UeConnectivityStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeConnectivityStates

`func (o *InputData) SetUeConnectivityStates(v UeConnectivityState)`

SetUeConnectivityStates sets UeConnectivityStates field to given value.

### HasUeConnectivityStates

`func (o *InputData) HasUeConnectivityStates() bool`

HasUeConnectivityStates returns a boolean if a field has been set.

### GetUeLocationServiceInd

`func (o *InputData) GetUeLocationServiceInd() UeLocationServiceInd`

GetUeLocationServiceInd returns the UeLocationServiceInd field if non-nil, zero value otherwise.

### GetUeLocationServiceIndOk

`func (o *InputData) GetUeLocationServiceIndOk() (*UeLocationServiceInd, bool)`

GetUeLocationServiceIndOk returns a tuple with the UeLocationServiceInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocationServiceInd

`func (o *InputData) SetUeLocationServiceInd(v UeLocationServiceInd)`

SetUeLocationServiceInd sets UeLocationServiceInd field to given value.

### HasUeLocationServiceInd

`func (o *InputData) HasUeLocationServiceInd() bool`

HasUeLocationServiceInd returns a boolean if a field has been set.

### GetMoAssistanceDataTypes

`func (o *InputData) GetMoAssistanceDataTypes() LcsBroadcastAssistanceTypesData`

GetMoAssistanceDataTypes returns the MoAssistanceDataTypes field if non-nil, zero value otherwise.

### GetMoAssistanceDataTypesOk

`func (o *InputData) GetMoAssistanceDataTypesOk() (*LcsBroadcastAssistanceTypesData, bool)`

GetMoAssistanceDataTypesOk returns a tuple with the MoAssistanceDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoAssistanceDataTypes

`func (o *InputData) SetMoAssistanceDataTypes(v LcsBroadcastAssistanceTypesData)`

SetMoAssistanceDataTypes sets MoAssistanceDataTypes field to given value.

### HasMoAssistanceDataTypes

`func (o *InputData) HasMoAssistanceDataTypes() bool`

HasMoAssistanceDataTypes returns a boolean if a field has been set.

### GetLppMessage

`func (o *InputData) GetLppMessage() RefToBinaryData`

GetLppMessage returns the LppMessage field if non-nil, zero value otherwise.

### GetLppMessageOk

`func (o *InputData) GetLppMessageOk() (*RefToBinaryData, bool)`

GetLppMessageOk returns a tuple with the LppMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLppMessage

`func (o *InputData) SetLppMessage(v RefToBinaryData)`

SetLppMessage sets LppMessage field to given value.

### HasLppMessage

`func (o *InputData) HasLppMessage() bool`

HasLppMessage returns a boolean if a field has been set.

### GetLppMessageExt

`func (o *InputData) GetLppMessageExt() []RefToBinaryData`

GetLppMessageExt returns the LppMessageExt field if non-nil, zero value otherwise.

### GetLppMessageExtOk

`func (o *InputData) GetLppMessageExtOk() (*[]RefToBinaryData, bool)`

GetLppMessageExtOk returns a tuple with the LppMessageExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLppMessageExt

`func (o *InputData) SetLppMessageExt(v []RefToBinaryData)`

SetLppMessageExt sets LppMessageExt field to given value.

### HasLppMessageExt

`func (o *InputData) HasLppMessageExt() bool`

HasLppMessageExt returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *InputData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *InputData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *InputData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *InputData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetUePositioningCap

`func (o *InputData) GetUePositioningCap() string`

GetUePositioningCap returns the UePositioningCap field if non-nil, zero value otherwise.

### GetUePositioningCapOk

`func (o *InputData) GetUePositioningCapOk() (*string, bool)`

GetUePositioningCapOk returns a tuple with the UePositioningCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUePositioningCap

`func (o *InputData) SetUePositioningCap(v string)`

SetUePositioningCap sets UePositioningCap field to given value.

### HasUePositioningCap

`func (o *InputData) HasUePositioningCap() bool`

HasUePositioningCap returns a boolean if a field has been set.

### GetTnapId

`func (o *InputData) GetTnapId() TnapId`

GetTnapId returns the TnapId field if non-nil, zero value otherwise.

### GetTnapIdOk

`func (o *InputData) GetTnapIdOk() (*TnapId, bool)`

GetTnapIdOk returns a tuple with the TnapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTnapId

`func (o *InputData) SetTnapId(v TnapId)`

SetTnapId sets TnapId field to given value.

### HasTnapId

`func (o *InputData) HasTnapId() bool`

HasTnapId returns a boolean if a field has been set.

### GetTwapId

`func (o *InputData) GetTwapId() TwapId`

GetTwapId returns the TwapId field if non-nil, zero value otherwise.

### GetTwapIdOk

`func (o *InputData) GetTwapIdOk() (*TwapId, bool)`

GetTwapIdOk returns a tuple with the TwapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwapId

`func (o *InputData) SetTwapId(v TwapId)`

SetTwapId sets TwapId field to given value.

### HasTwapId

`func (o *InputData) HasTwapId() bool`

HasTwapId returns a boolean if a field has been set.

### GetUeCountryDetInd

`func (o *InputData) GetUeCountryDetInd() bool`

GetUeCountryDetInd returns the UeCountryDetInd field if non-nil, zero value otherwise.

### GetUeCountryDetIndOk

`func (o *InputData) GetUeCountryDetIndOk() (*bool, bool)`

GetUeCountryDetIndOk returns a tuple with the UeCountryDetInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCountryDetInd

`func (o *InputData) SetUeCountryDetInd(v bool)`

SetUeCountryDetInd sets UeCountryDetInd field to given value.

### HasUeCountryDetInd

`func (o *InputData) HasUeCountryDetInd() bool`

HasUeCountryDetInd returns a boolean if a field has been set.

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


