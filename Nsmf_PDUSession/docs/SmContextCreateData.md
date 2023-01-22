# SmContextCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UnauthenticatedSupi** | Pointer to **bool** |  | [optional] [default to false]
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**PduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SelectedDnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**HplmnSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ServingNfId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ServiceName** | Pointer to [**ServiceName**](ServiceName.md) |  | [optional] 
**ServingNetwork** | [**PlmnIdNid**](PlmnIdNid.md) |  | 
**RequestType** | Pointer to [**RequestType**](RequestType.md) |  | [optional] 
**N1SmMsg** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**AnType** | [**AccessType**](AccessType.md) |  | 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**PresenceInLadn** | Pointer to [**PresenceState**](PresenceState.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**SmContextStatusUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**HSmfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**HSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SmfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AdditionalHsmfUri** | Pointer to **[]string** |  | [optional] 
**AdditionalHsmfId** | Pointer to **[]string** |  | [optional] 
**AdditionalSmfUri** | Pointer to **[]string** |  | [optional] 
**AdditionalSmfId** | Pointer to **[]string** |  | [optional] 
**OldPduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**PduSessionsActivateList** | Pointer to **[]int32** |  | [optional] 
**UeEpsPdnConnection** | Pointer to **string** | UE EPS PDN Connection container from SMF to AMF | [optional] 
**HoState** | Pointer to [**HoState**](HoState.md) |  | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**NrfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**SelMode** | Pointer to [**DnnSelectionMode**](DnnSelectionMode.md) |  | [optional] 
**BackupAmfInfo** | Pointer to [**[]BackupAmfInfo**](BackupAmfInfo.md) |  | [optional] 
**TraceData** | Pointer to [**NullableTraceData**](TraceData.md) |  | [optional] 
**UdmGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**HNwPubKeyId** | Pointer to **int32** |  | [optional] 
**EpsInterworkingInd** | Pointer to [**EpsInterworkingIndication**](EpsInterworkingIndication.md) |  | [optional] 
**IndirectForwardingFlag** | Pointer to **bool** |  | [optional] 
**DirectForwardingFlag** | Pointer to **bool** |  | [optional] 
**TargetId** | Pointer to [**NgRanTargetId**](NgRanTargetId.md) |  | [optional] 
**EpsBearerCtxStatus** | Pointer to **string** | EPS bearer context status | [optional] 
**CpCiotEnabled** | Pointer to **bool** |  | [optional] [default to false]
**CpOnlyInd** | Pointer to **bool** |  | [optional] [default to false]
**InvokeNef** | Pointer to **bool** |  | [optional] [default to false]
**MaRequestInd** | Pointer to **bool** |  | [optional] [default to false]
**MaNwUpgradeInd** | Pointer to **bool** |  | [optional] [default to false]
**N2SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoType** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**N2SmInfoExt1** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**N2SmInfoTypeExt1** | Pointer to [**N2SmInfoType**](N2SmInfoType.md) |  | [optional] 
**SmContextRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmContextSmfPlmnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**SmContextSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SmContextSmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SmContextSmfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**SmContextSmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**ExtendedNasSmTimerInd** | Pointer to **bool** |  | [optional] [default to false]
**DlDataWaitingInd** | Pointer to **bool** |  | [optional] [default to false]
**DdnFailureSubs** | Pointer to [**DdnFailureSubs**](DdnFailureSubs.md) |  | [optional] 
**SmfTransferInd** | Pointer to **bool** |  | [optional] [default to false]
**OldSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**OldSmContextRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**WAgfInfo** | Pointer to [**WAgfInfo**](WAgfInfo.md) |  | [optional] 
**TngfInfo** | Pointer to [**TngfInfo**](TngfInfo.md) |  | [optional] 
**TwifInfo** | Pointer to [**TwifInfo**](TwifInfo.md) |  | [optional] 
**RanUnchangedInd** | Pointer to **bool** |  | [optional] 
**SamePcfSelectionInd** | Pointer to **bool** |  | [optional] [default to false]
**TargetDnai** | Pointer to **string** | DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501. | [optional] 
**NrfManagementUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfDiscoveryUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfAccessTokenUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfOauth2Required** | Pointer to **map[string]bool** | Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \&quot;nnrf-nfm\&quot; or \&quot;nnrf-disc\&quot; | [optional] 
**SmfBindingInfo** | Pointer to **string** |  | [optional] 
**PvsInfo** | Pointer to [**[]ServerAddressingInfo**](ServerAddressingInfo.md) |  | [optional] 
**OnboardingInd** | Pointer to **bool** |  | [optional] [default to false]
**OldPduSessionRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmPolicyNotifyInd** | Pointer to **bool** |  | [optional] [default to false]
**PcfUeCallbackInfo** | Pointer to [**NullablePcfUeCallbackInfo**](PcfUeCallbackInfo.md) |  | [optional] 
**SatelliteBackhaulCat** | Pointer to [**SatelliteBackhaulCategory**](SatelliteBackhaulCategory.md) |  | [optional] 
**UpipSupported** | Pointer to **bool** |  | [optional] [default to false]
**UavAuthenticated** | Pointer to **bool** |  | [optional] 
**DisasterRoamingInd** | Pointer to **bool** |  | [optional] [default to false]
**AnchorSmfOauth2Required** | Pointer to **bool** |  | [optional] 
**SmContextSmfOauth2Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewSmContextCreateData

`func NewSmContextCreateData(servingNfId string, servingNetwork PlmnIdNid, anType AccessType, smContextStatusUri string, ) *SmContextCreateData`

NewSmContextCreateData instantiates a new SmContextCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextCreateDataWithDefaults

`func NewSmContextCreateDataWithDefaults() *SmContextCreateData`

NewSmContextCreateDataWithDefaults instantiates a new SmContextCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *SmContextCreateData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SmContextCreateData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SmContextCreateData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *SmContextCreateData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUnauthenticatedSupi

`func (o *SmContextCreateData) GetUnauthenticatedSupi() bool`

GetUnauthenticatedSupi returns the UnauthenticatedSupi field if non-nil, zero value otherwise.

### GetUnauthenticatedSupiOk

`func (o *SmContextCreateData) GetUnauthenticatedSupiOk() (*bool, bool)`

GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedSupi

`func (o *SmContextCreateData) SetUnauthenticatedSupi(v bool)`

SetUnauthenticatedSupi sets UnauthenticatedSupi field to given value.

### HasUnauthenticatedSupi

`func (o *SmContextCreateData) HasUnauthenticatedSupi() bool`

HasUnauthenticatedSupi returns a boolean if a field has been set.

### GetPei

`func (o *SmContextCreateData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *SmContextCreateData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *SmContextCreateData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *SmContextCreateData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetGpsi

`func (o *SmContextCreateData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmContextCreateData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmContextCreateData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *SmContextCreateData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPduSessionId

`func (o *SmContextCreateData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmContextCreateData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmContextCreateData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *SmContextCreateData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetDnn

`func (o *SmContextCreateData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmContextCreateData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmContextCreateData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *SmContextCreateData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSelectedDnn

`func (o *SmContextCreateData) GetSelectedDnn() string`

GetSelectedDnn returns the SelectedDnn field if non-nil, zero value otherwise.

### GetSelectedDnnOk

`func (o *SmContextCreateData) GetSelectedDnnOk() (*string, bool)`

GetSelectedDnnOk returns a tuple with the SelectedDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDnn

`func (o *SmContextCreateData) SetSelectedDnn(v string)`

SetSelectedDnn sets SelectedDnn field to given value.

### HasSelectedDnn

`func (o *SmContextCreateData) HasSelectedDnn() bool`

HasSelectedDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *SmContextCreateData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SmContextCreateData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SmContextCreateData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *SmContextCreateData) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetHplmnSnssai

`func (o *SmContextCreateData) GetHplmnSnssai() Snssai`

GetHplmnSnssai returns the HplmnSnssai field if non-nil, zero value otherwise.

### GetHplmnSnssaiOk

`func (o *SmContextCreateData) GetHplmnSnssaiOk() (*Snssai, bool)`

GetHplmnSnssaiOk returns a tuple with the HplmnSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHplmnSnssai

`func (o *SmContextCreateData) SetHplmnSnssai(v Snssai)`

SetHplmnSnssai sets HplmnSnssai field to given value.

### HasHplmnSnssai

`func (o *SmContextCreateData) HasHplmnSnssai() bool`

HasHplmnSnssai returns a boolean if a field has been set.

### GetServingNfId

`func (o *SmContextCreateData) GetServingNfId() string`

GetServingNfId returns the ServingNfId field if non-nil, zero value otherwise.

### GetServingNfIdOk

`func (o *SmContextCreateData) GetServingNfIdOk() (*string, bool)`

GetServingNfIdOk returns a tuple with the ServingNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNfId

`func (o *SmContextCreateData) SetServingNfId(v string)`

SetServingNfId sets ServingNfId field to given value.


### GetGuami

`func (o *SmContextCreateData) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *SmContextCreateData) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *SmContextCreateData) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *SmContextCreateData) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetServiceName

`func (o *SmContextCreateData) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SmContextCreateData) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SmContextCreateData) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *SmContextCreateData) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServingNetwork

`func (o *SmContextCreateData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *SmContextCreateData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *SmContextCreateData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.


### GetRequestType

`func (o *SmContextCreateData) GetRequestType() RequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *SmContextCreateData) GetRequestTypeOk() (*RequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *SmContextCreateData) SetRequestType(v RequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *SmContextCreateData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetN1SmMsg

`func (o *SmContextCreateData) GetN1SmMsg() RefToBinaryData`

GetN1SmMsg returns the N1SmMsg field if non-nil, zero value otherwise.

### GetN1SmMsgOk

`func (o *SmContextCreateData) GetN1SmMsgOk() (*RefToBinaryData, bool)`

GetN1SmMsgOk returns a tuple with the N1SmMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmMsg

`func (o *SmContextCreateData) SetN1SmMsg(v RefToBinaryData)`

SetN1SmMsg sets N1SmMsg field to given value.

### HasN1SmMsg

`func (o *SmContextCreateData) HasN1SmMsg() bool`

HasN1SmMsg returns a boolean if a field has been set.

### GetAnType

`func (o *SmContextCreateData) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *SmContextCreateData) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *SmContextCreateData) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.


### GetAdditionalAnType

`func (o *SmContextCreateData) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *SmContextCreateData) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *SmContextCreateData) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *SmContextCreateData) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.

### GetRatType

`func (o *SmContextCreateData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *SmContextCreateData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *SmContextCreateData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *SmContextCreateData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetPresenceInLadn

`func (o *SmContextCreateData) GetPresenceInLadn() PresenceState`

GetPresenceInLadn returns the PresenceInLadn field if non-nil, zero value otherwise.

### GetPresenceInLadnOk

`func (o *SmContextCreateData) GetPresenceInLadnOk() (*PresenceState, bool)`

GetPresenceInLadnOk returns a tuple with the PresenceInLadn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInLadn

`func (o *SmContextCreateData) SetPresenceInLadn(v PresenceState)`

SetPresenceInLadn sets PresenceInLadn field to given value.

### HasPresenceInLadn

`func (o *SmContextCreateData) HasPresenceInLadn() bool`

HasPresenceInLadn returns a boolean if a field has been set.

### GetUeLocation

`func (o *SmContextCreateData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *SmContextCreateData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *SmContextCreateData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *SmContextCreateData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmContextCreateData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmContextCreateData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmContextCreateData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmContextCreateData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *SmContextCreateData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *SmContextCreateData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *SmContextCreateData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *SmContextCreateData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetSmContextStatusUri

`func (o *SmContextCreateData) GetSmContextStatusUri() string`

GetSmContextStatusUri returns the SmContextStatusUri field if non-nil, zero value otherwise.

### GetSmContextStatusUriOk

`func (o *SmContextCreateData) GetSmContextStatusUriOk() (*string, bool)`

GetSmContextStatusUriOk returns a tuple with the SmContextStatusUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextStatusUri

`func (o *SmContextCreateData) SetSmContextStatusUri(v string)`

SetSmContextStatusUri sets SmContextStatusUri field to given value.


### GetHSmfUri

`func (o *SmContextCreateData) GetHSmfUri() string`

GetHSmfUri returns the HSmfUri field if non-nil, zero value otherwise.

### GetHSmfUriOk

`func (o *SmContextCreateData) GetHSmfUriOk() (*string, bool)`

GetHSmfUriOk returns a tuple with the HSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfUri

`func (o *SmContextCreateData) SetHSmfUri(v string)`

SetHSmfUri sets HSmfUri field to given value.

### HasHSmfUri

`func (o *SmContextCreateData) HasHSmfUri() bool`

HasHSmfUri returns a boolean if a field has been set.

### GetHSmfId

`func (o *SmContextCreateData) GetHSmfId() string`

GetHSmfId returns the HSmfId field if non-nil, zero value otherwise.

### GetHSmfIdOk

`func (o *SmContextCreateData) GetHSmfIdOk() (*string, bool)`

GetHSmfIdOk returns a tuple with the HSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHSmfId

`func (o *SmContextCreateData) SetHSmfId(v string)`

SetHSmfId sets HSmfId field to given value.

### HasHSmfId

`func (o *SmContextCreateData) HasHSmfId() bool`

HasHSmfId returns a boolean if a field has been set.

### GetSmfUri

`func (o *SmContextCreateData) GetSmfUri() string`

GetSmfUri returns the SmfUri field if non-nil, zero value otherwise.

### GetSmfUriOk

`func (o *SmContextCreateData) GetSmfUriOk() (*string, bool)`

GetSmfUriOk returns a tuple with the SmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfUri

`func (o *SmContextCreateData) SetSmfUri(v string)`

SetSmfUri sets SmfUri field to given value.

### HasSmfUri

`func (o *SmContextCreateData) HasSmfUri() bool`

HasSmfUri returns a boolean if a field has been set.

### GetSmfId

`func (o *SmContextCreateData) GetSmfId() string`

GetSmfId returns the SmfId field if non-nil, zero value otherwise.

### GetSmfIdOk

`func (o *SmContextCreateData) GetSmfIdOk() (*string, bool)`

GetSmfIdOk returns a tuple with the SmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfId

`func (o *SmContextCreateData) SetSmfId(v string)`

SetSmfId sets SmfId field to given value.

### HasSmfId

`func (o *SmContextCreateData) HasSmfId() bool`

HasSmfId returns a boolean if a field has been set.

### GetAdditionalHsmfUri

`func (o *SmContextCreateData) GetAdditionalHsmfUri() []string`

GetAdditionalHsmfUri returns the AdditionalHsmfUri field if non-nil, zero value otherwise.

### GetAdditionalHsmfUriOk

`func (o *SmContextCreateData) GetAdditionalHsmfUriOk() (*[]string, bool)`

GetAdditionalHsmfUriOk returns a tuple with the AdditionalHsmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHsmfUri

`func (o *SmContextCreateData) SetAdditionalHsmfUri(v []string)`

SetAdditionalHsmfUri sets AdditionalHsmfUri field to given value.

### HasAdditionalHsmfUri

`func (o *SmContextCreateData) HasAdditionalHsmfUri() bool`

HasAdditionalHsmfUri returns a boolean if a field has been set.

### GetAdditionalHsmfId

`func (o *SmContextCreateData) GetAdditionalHsmfId() []string`

GetAdditionalHsmfId returns the AdditionalHsmfId field if non-nil, zero value otherwise.

### GetAdditionalHsmfIdOk

`func (o *SmContextCreateData) GetAdditionalHsmfIdOk() (*[]string, bool)`

GetAdditionalHsmfIdOk returns a tuple with the AdditionalHsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHsmfId

`func (o *SmContextCreateData) SetAdditionalHsmfId(v []string)`

SetAdditionalHsmfId sets AdditionalHsmfId field to given value.

### HasAdditionalHsmfId

`func (o *SmContextCreateData) HasAdditionalHsmfId() bool`

HasAdditionalHsmfId returns a boolean if a field has been set.

### GetAdditionalSmfUri

`func (o *SmContextCreateData) GetAdditionalSmfUri() []string`

GetAdditionalSmfUri returns the AdditionalSmfUri field if non-nil, zero value otherwise.

### GetAdditionalSmfUriOk

`func (o *SmContextCreateData) GetAdditionalSmfUriOk() (*[]string, bool)`

GetAdditionalSmfUriOk returns a tuple with the AdditionalSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSmfUri

`func (o *SmContextCreateData) SetAdditionalSmfUri(v []string)`

SetAdditionalSmfUri sets AdditionalSmfUri field to given value.

### HasAdditionalSmfUri

`func (o *SmContextCreateData) HasAdditionalSmfUri() bool`

HasAdditionalSmfUri returns a boolean if a field has been set.

### GetAdditionalSmfId

`func (o *SmContextCreateData) GetAdditionalSmfId() []string`

GetAdditionalSmfId returns the AdditionalSmfId field if non-nil, zero value otherwise.

### GetAdditionalSmfIdOk

`func (o *SmContextCreateData) GetAdditionalSmfIdOk() (*[]string, bool)`

GetAdditionalSmfIdOk returns a tuple with the AdditionalSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSmfId

`func (o *SmContextCreateData) SetAdditionalSmfId(v []string)`

SetAdditionalSmfId sets AdditionalSmfId field to given value.

### HasAdditionalSmfId

`func (o *SmContextCreateData) HasAdditionalSmfId() bool`

HasAdditionalSmfId returns a boolean if a field has been set.

### GetOldPduSessionId

`func (o *SmContextCreateData) GetOldPduSessionId() int32`

GetOldPduSessionId returns the OldPduSessionId field if non-nil, zero value otherwise.

### GetOldPduSessionIdOk

`func (o *SmContextCreateData) GetOldPduSessionIdOk() (*int32, bool)`

GetOldPduSessionIdOk returns a tuple with the OldPduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionId

`func (o *SmContextCreateData) SetOldPduSessionId(v int32)`

SetOldPduSessionId sets OldPduSessionId field to given value.

### HasOldPduSessionId

`func (o *SmContextCreateData) HasOldPduSessionId() bool`

HasOldPduSessionId returns a boolean if a field has been set.

### GetPduSessionsActivateList

`func (o *SmContextCreateData) GetPduSessionsActivateList() []int32`

GetPduSessionsActivateList returns the PduSessionsActivateList field if non-nil, zero value otherwise.

### GetPduSessionsActivateListOk

`func (o *SmContextCreateData) GetPduSessionsActivateListOk() (*[]int32, bool)`

GetPduSessionsActivateListOk returns a tuple with the PduSessionsActivateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionsActivateList

`func (o *SmContextCreateData) SetPduSessionsActivateList(v []int32)`

SetPduSessionsActivateList sets PduSessionsActivateList field to given value.

### HasPduSessionsActivateList

`func (o *SmContextCreateData) HasPduSessionsActivateList() bool`

HasPduSessionsActivateList returns a boolean if a field has been set.

### GetUeEpsPdnConnection

`func (o *SmContextCreateData) GetUeEpsPdnConnection() string`

GetUeEpsPdnConnection returns the UeEpsPdnConnection field if non-nil, zero value otherwise.

### GetUeEpsPdnConnectionOk

`func (o *SmContextCreateData) GetUeEpsPdnConnectionOk() (*string, bool)`

GetUeEpsPdnConnectionOk returns a tuple with the UeEpsPdnConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeEpsPdnConnection

`func (o *SmContextCreateData) SetUeEpsPdnConnection(v string)`

SetUeEpsPdnConnection sets UeEpsPdnConnection field to given value.

### HasUeEpsPdnConnection

`func (o *SmContextCreateData) HasUeEpsPdnConnection() bool`

HasUeEpsPdnConnection returns a boolean if a field has been set.

### GetHoState

`func (o *SmContextCreateData) GetHoState() HoState`

GetHoState returns the HoState field if non-nil, zero value otherwise.

### GetHoStateOk

`func (o *SmContextCreateData) GetHoStateOk() (*HoState, bool)`

GetHoStateOk returns a tuple with the HoState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoState

`func (o *SmContextCreateData) SetHoState(v HoState)`

SetHoState sets HoState field to given value.

### HasHoState

`func (o *SmContextCreateData) HasHoState() bool`

HasHoState returns a boolean if a field has been set.

### GetPcfId

`func (o *SmContextCreateData) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *SmContextCreateData) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *SmContextCreateData) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *SmContextCreateData) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfGroupId

`func (o *SmContextCreateData) GetPcfGroupId() string`

GetPcfGroupId returns the PcfGroupId field if non-nil, zero value otherwise.

### GetPcfGroupIdOk

`func (o *SmContextCreateData) GetPcfGroupIdOk() (*string, bool)`

GetPcfGroupIdOk returns a tuple with the PcfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfGroupId

`func (o *SmContextCreateData) SetPcfGroupId(v string)`

SetPcfGroupId sets PcfGroupId field to given value.

### HasPcfGroupId

`func (o *SmContextCreateData) HasPcfGroupId() bool`

HasPcfGroupId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *SmContextCreateData) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *SmContextCreateData) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *SmContextCreateData) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *SmContextCreateData) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetNrfUri

`func (o *SmContextCreateData) GetNrfUri() string`

GetNrfUri returns the NrfUri field if non-nil, zero value otherwise.

### GetNrfUriOk

`func (o *SmContextCreateData) GetNrfUriOk() (*string, bool)`

GetNrfUriOk returns a tuple with the NrfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfUri

`func (o *SmContextCreateData) SetNrfUri(v string)`

SetNrfUri sets NrfUri field to given value.

### HasNrfUri

`func (o *SmContextCreateData) HasNrfUri() bool`

HasNrfUri returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetSelMode

`func (o *SmContextCreateData) GetSelMode() DnnSelectionMode`

GetSelMode returns the SelMode field if non-nil, zero value otherwise.

### GetSelModeOk

`func (o *SmContextCreateData) GetSelModeOk() (*DnnSelectionMode, bool)`

GetSelModeOk returns a tuple with the SelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelMode

`func (o *SmContextCreateData) SetSelMode(v DnnSelectionMode)`

SetSelMode sets SelMode field to given value.

### HasSelMode

`func (o *SmContextCreateData) HasSelMode() bool`

HasSelMode returns a boolean if a field has been set.

### GetBackupAmfInfo

`func (o *SmContextCreateData) GetBackupAmfInfo() []BackupAmfInfo`

GetBackupAmfInfo returns the BackupAmfInfo field if non-nil, zero value otherwise.

### GetBackupAmfInfoOk

`func (o *SmContextCreateData) GetBackupAmfInfoOk() (*[]BackupAmfInfo, bool)`

GetBackupAmfInfoOk returns a tuple with the BackupAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupAmfInfo

`func (o *SmContextCreateData) SetBackupAmfInfo(v []BackupAmfInfo)`

SetBackupAmfInfo sets BackupAmfInfo field to given value.

### HasBackupAmfInfo

`func (o *SmContextCreateData) HasBackupAmfInfo() bool`

HasBackupAmfInfo returns a boolean if a field has been set.

### GetTraceData

`func (o *SmContextCreateData) GetTraceData() TraceData`

GetTraceData returns the TraceData field if non-nil, zero value otherwise.

### GetTraceDataOk

`func (o *SmContextCreateData) GetTraceDataOk() (*TraceData, bool)`

GetTraceDataOk returns a tuple with the TraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceData

`func (o *SmContextCreateData) SetTraceData(v TraceData)`

SetTraceData sets TraceData field to given value.

### HasTraceData

`func (o *SmContextCreateData) HasTraceData() bool`

HasTraceData returns a boolean if a field has been set.

### SetTraceDataNil

`func (o *SmContextCreateData) SetTraceDataNil(b bool)`

 SetTraceDataNil sets the value for TraceData to be an explicit nil

### UnsetTraceData
`func (o *SmContextCreateData) UnsetTraceData()`

UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
### GetUdmGroupId

`func (o *SmContextCreateData) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *SmContextCreateData) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *SmContextCreateData) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *SmContextCreateData) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *SmContextCreateData) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *SmContextCreateData) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *SmContextCreateData) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *SmContextCreateData) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetHNwPubKeyId

`func (o *SmContextCreateData) GetHNwPubKeyId() int32`

GetHNwPubKeyId returns the HNwPubKeyId field if non-nil, zero value otherwise.

### GetHNwPubKeyIdOk

`func (o *SmContextCreateData) GetHNwPubKeyIdOk() (*int32, bool)`

GetHNwPubKeyIdOk returns a tuple with the HNwPubKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHNwPubKeyId

`func (o *SmContextCreateData) SetHNwPubKeyId(v int32)`

SetHNwPubKeyId sets HNwPubKeyId field to given value.

### HasHNwPubKeyId

`func (o *SmContextCreateData) HasHNwPubKeyId() bool`

HasHNwPubKeyId returns a boolean if a field has been set.

### GetEpsInterworkingInd

`func (o *SmContextCreateData) GetEpsInterworkingInd() EpsInterworkingIndication`

GetEpsInterworkingInd returns the EpsInterworkingInd field if non-nil, zero value otherwise.

### GetEpsInterworkingIndOk

`func (o *SmContextCreateData) GetEpsInterworkingIndOk() (*EpsInterworkingIndication, bool)`

GetEpsInterworkingIndOk returns a tuple with the EpsInterworkingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInd

`func (o *SmContextCreateData) SetEpsInterworkingInd(v EpsInterworkingIndication)`

SetEpsInterworkingInd sets EpsInterworkingInd field to given value.

### HasEpsInterworkingInd

`func (o *SmContextCreateData) HasEpsInterworkingInd() bool`

HasEpsInterworkingInd returns a boolean if a field has been set.

### GetIndirectForwardingFlag

`func (o *SmContextCreateData) GetIndirectForwardingFlag() bool`

GetIndirectForwardingFlag returns the IndirectForwardingFlag field if non-nil, zero value otherwise.

### GetIndirectForwardingFlagOk

`func (o *SmContextCreateData) GetIndirectForwardingFlagOk() (*bool, bool)`

GetIndirectForwardingFlagOk returns a tuple with the IndirectForwardingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectForwardingFlag

`func (o *SmContextCreateData) SetIndirectForwardingFlag(v bool)`

SetIndirectForwardingFlag sets IndirectForwardingFlag field to given value.

### HasIndirectForwardingFlag

`func (o *SmContextCreateData) HasIndirectForwardingFlag() bool`

HasIndirectForwardingFlag returns a boolean if a field has been set.

### GetDirectForwardingFlag

`func (o *SmContextCreateData) GetDirectForwardingFlag() bool`

GetDirectForwardingFlag returns the DirectForwardingFlag field if non-nil, zero value otherwise.

### GetDirectForwardingFlagOk

`func (o *SmContextCreateData) GetDirectForwardingFlagOk() (*bool, bool)`

GetDirectForwardingFlagOk returns a tuple with the DirectForwardingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectForwardingFlag

`func (o *SmContextCreateData) SetDirectForwardingFlag(v bool)`

SetDirectForwardingFlag sets DirectForwardingFlag field to given value.

### HasDirectForwardingFlag

`func (o *SmContextCreateData) HasDirectForwardingFlag() bool`

HasDirectForwardingFlag returns a boolean if a field has been set.

### GetTargetId

`func (o *SmContextCreateData) GetTargetId() NgRanTargetId`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *SmContextCreateData) GetTargetIdOk() (*NgRanTargetId, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *SmContextCreateData) SetTargetId(v NgRanTargetId)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *SmContextCreateData) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetEpsBearerCtxStatus

`func (o *SmContextCreateData) GetEpsBearerCtxStatus() string`

GetEpsBearerCtxStatus returns the EpsBearerCtxStatus field if non-nil, zero value otherwise.

### GetEpsBearerCtxStatusOk

`func (o *SmContextCreateData) GetEpsBearerCtxStatusOk() (*string, bool)`

GetEpsBearerCtxStatusOk returns a tuple with the EpsBearerCtxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerCtxStatus

`func (o *SmContextCreateData) SetEpsBearerCtxStatus(v string)`

SetEpsBearerCtxStatus sets EpsBearerCtxStatus field to given value.

### HasEpsBearerCtxStatus

`func (o *SmContextCreateData) HasEpsBearerCtxStatus() bool`

HasEpsBearerCtxStatus returns a boolean if a field has been set.

### GetCpCiotEnabled

`func (o *SmContextCreateData) GetCpCiotEnabled() bool`

GetCpCiotEnabled returns the CpCiotEnabled field if non-nil, zero value otherwise.

### GetCpCiotEnabledOk

`func (o *SmContextCreateData) GetCpCiotEnabledOk() (*bool, bool)`

GetCpCiotEnabledOk returns a tuple with the CpCiotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpCiotEnabled

`func (o *SmContextCreateData) SetCpCiotEnabled(v bool)`

SetCpCiotEnabled sets CpCiotEnabled field to given value.

### HasCpCiotEnabled

`func (o *SmContextCreateData) HasCpCiotEnabled() bool`

HasCpCiotEnabled returns a boolean if a field has been set.

### GetCpOnlyInd

`func (o *SmContextCreateData) GetCpOnlyInd() bool`

GetCpOnlyInd returns the CpOnlyInd field if non-nil, zero value otherwise.

### GetCpOnlyIndOk

`func (o *SmContextCreateData) GetCpOnlyIndOk() (*bool, bool)`

GetCpOnlyIndOk returns a tuple with the CpOnlyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpOnlyInd

`func (o *SmContextCreateData) SetCpOnlyInd(v bool)`

SetCpOnlyInd sets CpOnlyInd field to given value.

### HasCpOnlyInd

`func (o *SmContextCreateData) HasCpOnlyInd() bool`

HasCpOnlyInd returns a boolean if a field has been set.

### GetInvokeNef

`func (o *SmContextCreateData) GetInvokeNef() bool`

GetInvokeNef returns the InvokeNef field if non-nil, zero value otherwise.

### GetInvokeNefOk

`func (o *SmContextCreateData) GetInvokeNefOk() (*bool, bool)`

GetInvokeNefOk returns a tuple with the InvokeNef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeNef

`func (o *SmContextCreateData) SetInvokeNef(v bool)`

SetInvokeNef sets InvokeNef field to given value.

### HasInvokeNef

`func (o *SmContextCreateData) HasInvokeNef() bool`

HasInvokeNef returns a boolean if a field has been set.

### GetMaRequestInd

`func (o *SmContextCreateData) GetMaRequestInd() bool`

GetMaRequestInd returns the MaRequestInd field if non-nil, zero value otherwise.

### GetMaRequestIndOk

`func (o *SmContextCreateData) GetMaRequestIndOk() (*bool, bool)`

GetMaRequestIndOk returns a tuple with the MaRequestInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaRequestInd

`func (o *SmContextCreateData) SetMaRequestInd(v bool)`

SetMaRequestInd sets MaRequestInd field to given value.

### HasMaRequestInd

`func (o *SmContextCreateData) HasMaRequestInd() bool`

HasMaRequestInd returns a boolean if a field has been set.

### GetMaNwUpgradeInd

`func (o *SmContextCreateData) GetMaNwUpgradeInd() bool`

GetMaNwUpgradeInd returns the MaNwUpgradeInd field if non-nil, zero value otherwise.

### GetMaNwUpgradeIndOk

`func (o *SmContextCreateData) GetMaNwUpgradeIndOk() (*bool, bool)`

GetMaNwUpgradeIndOk returns a tuple with the MaNwUpgradeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaNwUpgradeInd

`func (o *SmContextCreateData) SetMaNwUpgradeInd(v bool)`

SetMaNwUpgradeInd sets MaNwUpgradeInd field to given value.

### HasMaNwUpgradeInd

`func (o *SmContextCreateData) HasMaNwUpgradeInd() bool`

HasMaNwUpgradeInd returns a boolean if a field has been set.

### GetN2SmInfo

`func (o *SmContextCreateData) GetN2SmInfo() RefToBinaryData`

GetN2SmInfo returns the N2SmInfo field if non-nil, zero value otherwise.

### GetN2SmInfoOk

`func (o *SmContextCreateData) GetN2SmInfoOk() (*RefToBinaryData, bool)`

GetN2SmInfoOk returns a tuple with the N2SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfo

`func (o *SmContextCreateData) SetN2SmInfo(v RefToBinaryData)`

SetN2SmInfo sets N2SmInfo field to given value.

### HasN2SmInfo

`func (o *SmContextCreateData) HasN2SmInfo() bool`

HasN2SmInfo returns a boolean if a field has been set.

### GetN2SmInfoType

`func (o *SmContextCreateData) GetN2SmInfoType() N2SmInfoType`

GetN2SmInfoType returns the N2SmInfoType field if non-nil, zero value otherwise.

### GetN2SmInfoTypeOk

`func (o *SmContextCreateData) GetN2SmInfoTypeOk() (*N2SmInfoType, bool)`

GetN2SmInfoTypeOk returns a tuple with the N2SmInfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoType

`func (o *SmContextCreateData) SetN2SmInfoType(v N2SmInfoType)`

SetN2SmInfoType sets N2SmInfoType field to given value.

### HasN2SmInfoType

`func (o *SmContextCreateData) HasN2SmInfoType() bool`

HasN2SmInfoType returns a boolean if a field has been set.

### GetN2SmInfoExt1

`func (o *SmContextCreateData) GetN2SmInfoExt1() RefToBinaryData`

GetN2SmInfoExt1 returns the N2SmInfoExt1 field if non-nil, zero value otherwise.

### GetN2SmInfoExt1Ok

`func (o *SmContextCreateData) GetN2SmInfoExt1Ok() (*RefToBinaryData, bool)`

GetN2SmInfoExt1Ok returns a tuple with the N2SmInfoExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoExt1

`func (o *SmContextCreateData) SetN2SmInfoExt1(v RefToBinaryData)`

SetN2SmInfoExt1 sets N2SmInfoExt1 field to given value.

### HasN2SmInfoExt1

`func (o *SmContextCreateData) HasN2SmInfoExt1() bool`

HasN2SmInfoExt1 returns a boolean if a field has been set.

### GetN2SmInfoTypeExt1

`func (o *SmContextCreateData) GetN2SmInfoTypeExt1() N2SmInfoType`

GetN2SmInfoTypeExt1 returns the N2SmInfoTypeExt1 field if non-nil, zero value otherwise.

### GetN2SmInfoTypeExt1Ok

`func (o *SmContextCreateData) GetN2SmInfoTypeExt1Ok() (*N2SmInfoType, bool)`

GetN2SmInfoTypeExt1Ok returns a tuple with the N2SmInfoTypeExt1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2SmInfoTypeExt1

`func (o *SmContextCreateData) SetN2SmInfoTypeExt1(v N2SmInfoType)`

SetN2SmInfoTypeExt1 sets N2SmInfoTypeExt1 field to given value.

### HasN2SmInfoTypeExt1

`func (o *SmContextCreateData) HasN2SmInfoTypeExt1() bool`

HasN2SmInfoTypeExt1 returns a boolean if a field has been set.

### GetSmContextRef

`func (o *SmContextCreateData) GetSmContextRef() string`

GetSmContextRef returns the SmContextRef field if non-nil, zero value otherwise.

### GetSmContextRefOk

`func (o *SmContextCreateData) GetSmContextRefOk() (*string, bool)`

GetSmContextRefOk returns a tuple with the SmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextRef

`func (o *SmContextCreateData) SetSmContextRef(v string)`

SetSmContextRef sets SmContextRef field to given value.

### HasSmContextRef

`func (o *SmContextCreateData) HasSmContextRef() bool`

HasSmContextRef returns a boolean if a field has been set.

### GetSmContextSmfPlmnId

`func (o *SmContextCreateData) GetSmContextSmfPlmnId() PlmnIdNid`

GetSmContextSmfPlmnId returns the SmContextSmfPlmnId field if non-nil, zero value otherwise.

### GetSmContextSmfPlmnIdOk

`func (o *SmContextCreateData) GetSmContextSmfPlmnIdOk() (*PlmnIdNid, bool)`

GetSmContextSmfPlmnIdOk returns a tuple with the SmContextSmfPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfPlmnId

`func (o *SmContextCreateData) SetSmContextSmfPlmnId(v PlmnIdNid)`

SetSmContextSmfPlmnId sets SmContextSmfPlmnId field to given value.

### HasSmContextSmfPlmnId

`func (o *SmContextCreateData) HasSmContextSmfPlmnId() bool`

HasSmContextSmfPlmnId returns a boolean if a field has been set.

### GetSmContextSmfId

`func (o *SmContextCreateData) GetSmContextSmfId() string`

GetSmContextSmfId returns the SmContextSmfId field if non-nil, zero value otherwise.

### GetSmContextSmfIdOk

`func (o *SmContextCreateData) GetSmContextSmfIdOk() (*string, bool)`

GetSmContextSmfIdOk returns a tuple with the SmContextSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfId

`func (o *SmContextCreateData) SetSmContextSmfId(v string)`

SetSmContextSmfId sets SmContextSmfId field to given value.

### HasSmContextSmfId

`func (o *SmContextCreateData) HasSmContextSmfId() bool`

HasSmContextSmfId returns a boolean if a field has been set.

### GetSmContextSmfSetId

`func (o *SmContextCreateData) GetSmContextSmfSetId() string`

GetSmContextSmfSetId returns the SmContextSmfSetId field if non-nil, zero value otherwise.

### GetSmContextSmfSetIdOk

`func (o *SmContextCreateData) GetSmContextSmfSetIdOk() (*string, bool)`

GetSmContextSmfSetIdOk returns a tuple with the SmContextSmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfSetId

`func (o *SmContextCreateData) SetSmContextSmfSetId(v string)`

SetSmContextSmfSetId sets SmContextSmfSetId field to given value.

### HasSmContextSmfSetId

`func (o *SmContextCreateData) HasSmContextSmfSetId() bool`

HasSmContextSmfSetId returns a boolean if a field has been set.

### GetSmContextSmfServiceSetId

`func (o *SmContextCreateData) GetSmContextSmfServiceSetId() string`

GetSmContextSmfServiceSetId returns the SmContextSmfServiceSetId field if non-nil, zero value otherwise.

### GetSmContextSmfServiceSetIdOk

`func (o *SmContextCreateData) GetSmContextSmfServiceSetIdOk() (*string, bool)`

GetSmContextSmfServiceSetIdOk returns a tuple with the SmContextSmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfServiceSetId

`func (o *SmContextCreateData) SetSmContextSmfServiceSetId(v string)`

SetSmContextSmfServiceSetId sets SmContextSmfServiceSetId field to given value.

### HasSmContextSmfServiceSetId

`func (o *SmContextCreateData) HasSmContextSmfServiceSetId() bool`

HasSmContextSmfServiceSetId returns a boolean if a field has been set.

### GetSmContextSmfBinding

`func (o *SmContextCreateData) GetSmContextSmfBinding() SbiBindingLevel`

GetSmContextSmfBinding returns the SmContextSmfBinding field if non-nil, zero value otherwise.

### GetSmContextSmfBindingOk

`func (o *SmContextCreateData) GetSmContextSmfBindingOk() (*SbiBindingLevel, bool)`

GetSmContextSmfBindingOk returns a tuple with the SmContextSmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfBinding

`func (o *SmContextCreateData) SetSmContextSmfBinding(v SbiBindingLevel)`

SetSmContextSmfBinding sets SmContextSmfBinding field to given value.

### HasSmContextSmfBinding

`func (o *SmContextCreateData) HasSmContextSmfBinding() bool`

HasSmContextSmfBinding returns a boolean if a field has been set.

### GetUpCnxState

`func (o *SmContextCreateData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *SmContextCreateData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *SmContextCreateData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *SmContextCreateData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.

### GetSmallDataRateStatus

`func (o *SmContextCreateData) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextCreateData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextCreateData) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextCreateData) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *SmContextCreateData) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *SmContextCreateData) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *SmContextCreateData) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *SmContextCreateData) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetExtendedNasSmTimerInd

`func (o *SmContextCreateData) GetExtendedNasSmTimerInd() bool`

GetExtendedNasSmTimerInd returns the ExtendedNasSmTimerInd field if non-nil, zero value otherwise.

### GetExtendedNasSmTimerIndOk

`func (o *SmContextCreateData) GetExtendedNasSmTimerIndOk() (*bool, bool)`

GetExtendedNasSmTimerIndOk returns a tuple with the ExtendedNasSmTimerInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedNasSmTimerInd

`func (o *SmContextCreateData) SetExtendedNasSmTimerInd(v bool)`

SetExtendedNasSmTimerInd sets ExtendedNasSmTimerInd field to given value.

### HasExtendedNasSmTimerInd

`func (o *SmContextCreateData) HasExtendedNasSmTimerInd() bool`

HasExtendedNasSmTimerInd returns a boolean if a field has been set.

### GetDlDataWaitingInd

`func (o *SmContextCreateData) GetDlDataWaitingInd() bool`

GetDlDataWaitingInd returns the DlDataWaitingInd field if non-nil, zero value otherwise.

### GetDlDataWaitingIndOk

`func (o *SmContextCreateData) GetDlDataWaitingIndOk() (*bool, bool)`

GetDlDataWaitingIndOk returns a tuple with the DlDataWaitingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlDataWaitingInd

`func (o *SmContextCreateData) SetDlDataWaitingInd(v bool)`

SetDlDataWaitingInd sets DlDataWaitingInd field to given value.

### HasDlDataWaitingInd

`func (o *SmContextCreateData) HasDlDataWaitingInd() bool`

HasDlDataWaitingInd returns a boolean if a field has been set.

### GetDdnFailureSubs

`func (o *SmContextCreateData) GetDdnFailureSubs() DdnFailureSubs`

GetDdnFailureSubs returns the DdnFailureSubs field if non-nil, zero value otherwise.

### GetDdnFailureSubsOk

`func (o *SmContextCreateData) GetDdnFailureSubsOk() (*DdnFailureSubs, bool)`

GetDdnFailureSubsOk returns a tuple with the DdnFailureSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnFailureSubs

`func (o *SmContextCreateData) SetDdnFailureSubs(v DdnFailureSubs)`

SetDdnFailureSubs sets DdnFailureSubs field to given value.

### HasDdnFailureSubs

`func (o *SmContextCreateData) HasDdnFailureSubs() bool`

HasDdnFailureSubs returns a boolean if a field has been set.

### GetSmfTransferInd

`func (o *SmContextCreateData) GetSmfTransferInd() bool`

GetSmfTransferInd returns the SmfTransferInd field if non-nil, zero value otherwise.

### GetSmfTransferIndOk

`func (o *SmContextCreateData) GetSmfTransferIndOk() (*bool, bool)`

GetSmfTransferIndOk returns a tuple with the SmfTransferInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfTransferInd

`func (o *SmContextCreateData) SetSmfTransferInd(v bool)`

SetSmfTransferInd sets SmfTransferInd field to given value.

### HasSmfTransferInd

`func (o *SmContextCreateData) HasSmfTransferInd() bool`

HasSmfTransferInd returns a boolean if a field has been set.

### GetOldSmfId

`func (o *SmContextCreateData) GetOldSmfId() string`

GetOldSmfId returns the OldSmfId field if non-nil, zero value otherwise.

### GetOldSmfIdOk

`func (o *SmContextCreateData) GetOldSmfIdOk() (*string, bool)`

GetOldSmfIdOk returns a tuple with the OldSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmfId

`func (o *SmContextCreateData) SetOldSmfId(v string)`

SetOldSmfId sets OldSmfId field to given value.

### HasOldSmfId

`func (o *SmContextCreateData) HasOldSmfId() bool`

HasOldSmfId returns a boolean if a field has been set.

### GetOldSmContextRef

`func (o *SmContextCreateData) GetOldSmContextRef() string`

GetOldSmContextRef returns the OldSmContextRef field if non-nil, zero value otherwise.

### GetOldSmContextRefOk

`func (o *SmContextCreateData) GetOldSmContextRefOk() (*string, bool)`

GetOldSmContextRefOk returns a tuple with the OldSmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmContextRef

`func (o *SmContextCreateData) SetOldSmContextRef(v string)`

SetOldSmContextRef sets OldSmContextRef field to given value.

### HasOldSmContextRef

`func (o *SmContextCreateData) HasOldSmContextRef() bool`

HasOldSmContextRef returns a boolean if a field has been set.

### GetWAgfInfo

`func (o *SmContextCreateData) GetWAgfInfo() WAgfInfo`

GetWAgfInfo returns the WAgfInfo field if non-nil, zero value otherwise.

### GetWAgfInfoOk

`func (o *SmContextCreateData) GetWAgfInfoOk() (*WAgfInfo, bool)`

GetWAgfInfoOk returns a tuple with the WAgfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWAgfInfo

`func (o *SmContextCreateData) SetWAgfInfo(v WAgfInfo)`

SetWAgfInfo sets WAgfInfo field to given value.

### HasWAgfInfo

`func (o *SmContextCreateData) HasWAgfInfo() bool`

HasWAgfInfo returns a boolean if a field has been set.

### GetTngfInfo

`func (o *SmContextCreateData) GetTngfInfo() TngfInfo`

GetTngfInfo returns the TngfInfo field if non-nil, zero value otherwise.

### GetTngfInfoOk

`func (o *SmContextCreateData) GetTngfInfoOk() (*TngfInfo, bool)`

GetTngfInfoOk returns a tuple with the TngfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTngfInfo

`func (o *SmContextCreateData) SetTngfInfo(v TngfInfo)`

SetTngfInfo sets TngfInfo field to given value.

### HasTngfInfo

`func (o *SmContextCreateData) HasTngfInfo() bool`

HasTngfInfo returns a boolean if a field has been set.

### GetTwifInfo

`func (o *SmContextCreateData) GetTwifInfo() TwifInfo`

GetTwifInfo returns the TwifInfo field if non-nil, zero value otherwise.

### GetTwifInfoOk

`func (o *SmContextCreateData) GetTwifInfoOk() (*TwifInfo, bool)`

GetTwifInfoOk returns a tuple with the TwifInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwifInfo

`func (o *SmContextCreateData) SetTwifInfo(v TwifInfo)`

SetTwifInfo sets TwifInfo field to given value.

### HasTwifInfo

`func (o *SmContextCreateData) HasTwifInfo() bool`

HasTwifInfo returns a boolean if a field has been set.

### GetRanUnchangedInd

`func (o *SmContextCreateData) GetRanUnchangedInd() bool`

GetRanUnchangedInd returns the RanUnchangedInd field if non-nil, zero value otherwise.

### GetRanUnchangedIndOk

`func (o *SmContextCreateData) GetRanUnchangedIndOk() (*bool, bool)`

GetRanUnchangedIndOk returns a tuple with the RanUnchangedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUnchangedInd

`func (o *SmContextCreateData) SetRanUnchangedInd(v bool)`

SetRanUnchangedInd sets RanUnchangedInd field to given value.

### HasRanUnchangedInd

`func (o *SmContextCreateData) HasRanUnchangedInd() bool`

HasRanUnchangedInd returns a boolean if a field has been set.

### GetSamePcfSelectionInd

`func (o *SmContextCreateData) GetSamePcfSelectionInd() bool`

GetSamePcfSelectionInd returns the SamePcfSelectionInd field if non-nil, zero value otherwise.

### GetSamePcfSelectionIndOk

`func (o *SmContextCreateData) GetSamePcfSelectionIndOk() (*bool, bool)`

GetSamePcfSelectionIndOk returns a tuple with the SamePcfSelectionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePcfSelectionInd

`func (o *SmContextCreateData) SetSamePcfSelectionInd(v bool)`

SetSamePcfSelectionInd sets SamePcfSelectionInd field to given value.

### HasSamePcfSelectionInd

`func (o *SmContextCreateData) HasSamePcfSelectionInd() bool`

HasSamePcfSelectionInd returns a boolean if a field has been set.

### GetTargetDnai

`func (o *SmContextCreateData) GetTargetDnai() string`

GetTargetDnai returns the TargetDnai field if non-nil, zero value otherwise.

### GetTargetDnaiOk

`func (o *SmContextCreateData) GetTargetDnaiOk() (*string, bool)`

GetTargetDnaiOk returns a tuple with the TargetDnai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnai

`func (o *SmContextCreateData) SetTargetDnai(v string)`

SetTargetDnai sets TargetDnai field to given value.

### HasTargetDnai

`func (o *SmContextCreateData) HasTargetDnai() bool`

HasTargetDnai returns a boolean if a field has been set.

### GetNrfManagementUri

`func (o *SmContextCreateData) GetNrfManagementUri() string`

GetNrfManagementUri returns the NrfManagementUri field if non-nil, zero value otherwise.

### GetNrfManagementUriOk

`func (o *SmContextCreateData) GetNrfManagementUriOk() (*string, bool)`

GetNrfManagementUriOk returns a tuple with the NrfManagementUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfManagementUri

`func (o *SmContextCreateData) SetNrfManagementUri(v string)`

SetNrfManagementUri sets NrfManagementUri field to given value.

### HasNrfManagementUri

`func (o *SmContextCreateData) HasNrfManagementUri() bool`

HasNrfManagementUri returns a boolean if a field has been set.

### GetNrfDiscoveryUri

`func (o *SmContextCreateData) GetNrfDiscoveryUri() string`

GetNrfDiscoveryUri returns the NrfDiscoveryUri field if non-nil, zero value otherwise.

### GetNrfDiscoveryUriOk

`func (o *SmContextCreateData) GetNrfDiscoveryUriOk() (*string, bool)`

GetNrfDiscoveryUriOk returns a tuple with the NrfDiscoveryUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfDiscoveryUri

`func (o *SmContextCreateData) SetNrfDiscoveryUri(v string)`

SetNrfDiscoveryUri sets NrfDiscoveryUri field to given value.

### HasNrfDiscoveryUri

`func (o *SmContextCreateData) HasNrfDiscoveryUri() bool`

HasNrfDiscoveryUri returns a boolean if a field has been set.

### GetNrfAccessTokenUri

`func (o *SmContextCreateData) GetNrfAccessTokenUri() string`

GetNrfAccessTokenUri returns the NrfAccessTokenUri field if non-nil, zero value otherwise.

### GetNrfAccessTokenUriOk

`func (o *SmContextCreateData) GetNrfAccessTokenUriOk() (*string, bool)`

GetNrfAccessTokenUriOk returns a tuple with the NrfAccessTokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfAccessTokenUri

`func (o *SmContextCreateData) SetNrfAccessTokenUri(v string)`

SetNrfAccessTokenUri sets NrfAccessTokenUri field to given value.

### HasNrfAccessTokenUri

`func (o *SmContextCreateData) HasNrfAccessTokenUri() bool`

HasNrfAccessTokenUri returns a boolean if a field has been set.

### GetNrfOauth2Required

`func (o *SmContextCreateData) GetNrfOauth2Required() map[string]bool`

GetNrfOauth2Required returns the NrfOauth2Required field if non-nil, zero value otherwise.

### GetNrfOauth2RequiredOk

`func (o *SmContextCreateData) GetNrfOauth2RequiredOk() (*map[string]bool, bool)`

GetNrfOauth2RequiredOk returns a tuple with the NrfOauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfOauth2Required

`func (o *SmContextCreateData) SetNrfOauth2Required(v map[string]bool)`

SetNrfOauth2Required sets NrfOauth2Required field to given value.

### HasNrfOauth2Required

`func (o *SmContextCreateData) HasNrfOauth2Required() bool`

HasNrfOauth2Required returns a boolean if a field has been set.

### GetSmfBindingInfo

`func (o *SmContextCreateData) GetSmfBindingInfo() string`

GetSmfBindingInfo returns the SmfBindingInfo field if non-nil, zero value otherwise.

### GetSmfBindingInfoOk

`func (o *SmContextCreateData) GetSmfBindingInfoOk() (*string, bool)`

GetSmfBindingInfoOk returns a tuple with the SmfBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfBindingInfo

`func (o *SmContextCreateData) SetSmfBindingInfo(v string)`

SetSmfBindingInfo sets SmfBindingInfo field to given value.

### HasSmfBindingInfo

`func (o *SmContextCreateData) HasSmfBindingInfo() bool`

HasSmfBindingInfo returns a boolean if a field has been set.

### GetPvsInfo

`func (o *SmContextCreateData) GetPvsInfo() []ServerAddressingInfo`

GetPvsInfo returns the PvsInfo field if non-nil, zero value otherwise.

### GetPvsInfoOk

`func (o *SmContextCreateData) GetPvsInfoOk() (*[]ServerAddressingInfo, bool)`

GetPvsInfoOk returns a tuple with the PvsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsInfo

`func (o *SmContextCreateData) SetPvsInfo(v []ServerAddressingInfo)`

SetPvsInfo sets PvsInfo field to given value.

### HasPvsInfo

`func (o *SmContextCreateData) HasPvsInfo() bool`

HasPvsInfo returns a boolean if a field has been set.

### GetOnboardingInd

`func (o *SmContextCreateData) GetOnboardingInd() bool`

GetOnboardingInd returns the OnboardingInd field if non-nil, zero value otherwise.

### GetOnboardingIndOk

`func (o *SmContextCreateData) GetOnboardingIndOk() (*bool, bool)`

GetOnboardingIndOk returns a tuple with the OnboardingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingInd

`func (o *SmContextCreateData) SetOnboardingInd(v bool)`

SetOnboardingInd sets OnboardingInd field to given value.

### HasOnboardingInd

`func (o *SmContextCreateData) HasOnboardingInd() bool`

HasOnboardingInd returns a boolean if a field has been set.

### GetOldPduSessionRef

`func (o *SmContextCreateData) GetOldPduSessionRef() string`

GetOldPduSessionRef returns the OldPduSessionRef field if non-nil, zero value otherwise.

### GetOldPduSessionRefOk

`func (o *SmContextCreateData) GetOldPduSessionRefOk() (*string, bool)`

GetOldPduSessionRefOk returns a tuple with the OldPduSessionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionRef

`func (o *SmContextCreateData) SetOldPduSessionRef(v string)`

SetOldPduSessionRef sets OldPduSessionRef field to given value.

### HasOldPduSessionRef

`func (o *SmContextCreateData) HasOldPduSessionRef() bool`

HasOldPduSessionRef returns a boolean if a field has been set.

### GetSmPolicyNotifyInd

`func (o *SmContextCreateData) GetSmPolicyNotifyInd() bool`

GetSmPolicyNotifyInd returns the SmPolicyNotifyInd field if non-nil, zero value otherwise.

### GetSmPolicyNotifyIndOk

`func (o *SmContextCreateData) GetSmPolicyNotifyIndOk() (*bool, bool)`

GetSmPolicyNotifyIndOk returns a tuple with the SmPolicyNotifyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyNotifyInd

`func (o *SmContextCreateData) SetSmPolicyNotifyInd(v bool)`

SetSmPolicyNotifyInd sets SmPolicyNotifyInd field to given value.

### HasSmPolicyNotifyInd

`func (o *SmContextCreateData) HasSmPolicyNotifyInd() bool`

HasSmPolicyNotifyInd returns a boolean if a field has been set.

### GetPcfUeCallbackInfo

`func (o *SmContextCreateData) GetPcfUeCallbackInfo() PcfUeCallbackInfo`

GetPcfUeCallbackInfo returns the PcfUeCallbackInfo field if non-nil, zero value otherwise.

### GetPcfUeCallbackInfoOk

`func (o *SmContextCreateData) GetPcfUeCallbackInfoOk() (*PcfUeCallbackInfo, bool)`

GetPcfUeCallbackInfoOk returns a tuple with the PcfUeCallbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeCallbackInfo

`func (o *SmContextCreateData) SetPcfUeCallbackInfo(v PcfUeCallbackInfo)`

SetPcfUeCallbackInfo sets PcfUeCallbackInfo field to given value.

### HasPcfUeCallbackInfo

`func (o *SmContextCreateData) HasPcfUeCallbackInfo() bool`

HasPcfUeCallbackInfo returns a boolean if a field has been set.

### SetPcfUeCallbackInfoNil

`func (o *SmContextCreateData) SetPcfUeCallbackInfoNil(b bool)`

 SetPcfUeCallbackInfoNil sets the value for PcfUeCallbackInfo to be an explicit nil

### UnsetPcfUeCallbackInfo
`func (o *SmContextCreateData) UnsetPcfUeCallbackInfo()`

UnsetPcfUeCallbackInfo ensures that no value is present for PcfUeCallbackInfo, not even an explicit nil
### GetSatelliteBackhaulCat

`func (o *SmContextCreateData) GetSatelliteBackhaulCat() SatelliteBackhaulCategory`

GetSatelliteBackhaulCat returns the SatelliteBackhaulCat field if non-nil, zero value otherwise.

### GetSatelliteBackhaulCatOk

`func (o *SmContextCreateData) GetSatelliteBackhaulCatOk() (*SatelliteBackhaulCategory, bool)`

GetSatelliteBackhaulCatOk returns a tuple with the SatelliteBackhaulCat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteBackhaulCat

`func (o *SmContextCreateData) SetSatelliteBackhaulCat(v SatelliteBackhaulCategory)`

SetSatelliteBackhaulCat sets SatelliteBackhaulCat field to given value.

### HasSatelliteBackhaulCat

`func (o *SmContextCreateData) HasSatelliteBackhaulCat() bool`

HasSatelliteBackhaulCat returns a boolean if a field has been set.

### GetUpipSupported

`func (o *SmContextCreateData) GetUpipSupported() bool`

GetUpipSupported returns the UpipSupported field if non-nil, zero value otherwise.

### GetUpipSupportedOk

`func (o *SmContextCreateData) GetUpipSupportedOk() (*bool, bool)`

GetUpipSupportedOk returns a tuple with the UpipSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpipSupported

`func (o *SmContextCreateData) SetUpipSupported(v bool)`

SetUpipSupported sets UpipSupported field to given value.

### HasUpipSupported

`func (o *SmContextCreateData) HasUpipSupported() bool`

HasUpipSupported returns a boolean if a field has been set.

### GetUavAuthenticated

`func (o *SmContextCreateData) GetUavAuthenticated() bool`

GetUavAuthenticated returns the UavAuthenticated field if non-nil, zero value otherwise.

### GetUavAuthenticatedOk

`func (o *SmContextCreateData) GetUavAuthenticatedOk() (*bool, bool)`

GetUavAuthenticatedOk returns a tuple with the UavAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUavAuthenticated

`func (o *SmContextCreateData) SetUavAuthenticated(v bool)`

SetUavAuthenticated sets UavAuthenticated field to given value.

### HasUavAuthenticated

`func (o *SmContextCreateData) HasUavAuthenticated() bool`

HasUavAuthenticated returns a boolean if a field has been set.

### GetDisasterRoamingInd

`func (o *SmContextCreateData) GetDisasterRoamingInd() bool`

GetDisasterRoamingInd returns the DisasterRoamingInd field if non-nil, zero value otherwise.

### GetDisasterRoamingIndOk

`func (o *SmContextCreateData) GetDisasterRoamingIndOk() (*bool, bool)`

GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRoamingInd

`func (o *SmContextCreateData) SetDisasterRoamingInd(v bool)`

SetDisasterRoamingInd sets DisasterRoamingInd field to given value.

### HasDisasterRoamingInd

`func (o *SmContextCreateData) HasDisasterRoamingInd() bool`

HasDisasterRoamingInd returns a boolean if a field has been set.

### GetAnchorSmfOauth2Required

`func (o *SmContextCreateData) GetAnchorSmfOauth2Required() bool`

GetAnchorSmfOauth2Required returns the AnchorSmfOauth2Required field if non-nil, zero value otherwise.

### GetAnchorSmfOauth2RequiredOk

`func (o *SmContextCreateData) GetAnchorSmfOauth2RequiredOk() (*bool, bool)`

GetAnchorSmfOauth2RequiredOk returns a tuple with the AnchorSmfOauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorSmfOauth2Required

`func (o *SmContextCreateData) SetAnchorSmfOauth2Required(v bool)`

SetAnchorSmfOauth2Required sets AnchorSmfOauth2Required field to given value.

### HasAnchorSmfOauth2Required

`func (o *SmContextCreateData) HasAnchorSmfOauth2Required() bool`

HasAnchorSmfOauth2Required returns a boolean if a field has been set.

### GetSmContextSmfOauth2Required

`func (o *SmContextCreateData) GetSmContextSmfOauth2Required() bool`

GetSmContextSmfOauth2Required returns the SmContextSmfOauth2Required field if non-nil, zero value otherwise.

### GetSmContextSmfOauth2RequiredOk

`func (o *SmContextCreateData) GetSmContextSmfOauth2RequiredOk() (*bool, bool)`

GetSmContextSmfOauth2RequiredOk returns a tuple with the SmContextSmfOauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextSmfOauth2Required

`func (o *SmContextCreateData) SetSmContextSmfOauth2Required(v bool)`

SetSmContextSmfOauth2Required sets SmContextSmfOauth2Required field to given value.

### HasSmContextSmfOauth2Required

`func (o *SmContextCreateData) HasSmContextSmfOauth2Required() bool`

HasSmContextSmfOauth2Required returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


