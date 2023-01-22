# PduSessionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**UnauthenticatedSupi** | Pointer to **bool** |  | [optional] [default to false]
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**PduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**SelectedDnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**HplmnSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**VsmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**IsmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**ServingNetwork** | [**PlmnIdNid**](PlmnIdNid.md) |  | 
**RequestType** | Pointer to [**RequestType**](RequestType.md) |  | [optional] 
**EpsBearerId** | Pointer to **[]int32** |  | [optional] 
**PgwS8cFteid** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**VsmfPduSessionUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**IsmfPduSessionUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**VcnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**IcnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**N9ForwardingTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**AdditionalCnTunnelInfo** | Pointer to [**TunnelInfo**](TunnelInfo.md) |  | [optional] 
**AnType** | [**AccessType**](AccessType.md) |  | 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**AddUeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**N1SmInfoFromUe** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**UnknownN1SmInfo** | Pointer to [**RefToBinaryData**](RefToBinaryData.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**HPcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**HoPreparationIndication** | Pointer to **bool** |  | [optional] 
**SelMode** | Pointer to [**DnnSelectionMode**](DnnSelectionMode.md) |  | [optional] 
**AlwaysOnRequested** | Pointer to **bool** |  | [optional] [default to false]
**UdmGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 
**RoutingIndicator** | Pointer to **string** |  | [optional] 
**HNwPubKeyId** | Pointer to **int32** |  | [optional] 
**EpsInterworkingInd** | Pointer to [**EpsInterworkingIndication**](EpsInterworkingIndication.md) |  | [optional] 
**VSmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**ISmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RoamingChargingProfile** | Pointer to [**RoamingChargingProfile**](RoamingChargingProfile.md) |  | [optional] 
**ChargingId** | Pointer to **string** |  | [optional] 
**OldPduSessionId** | Pointer to **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | [optional] 
**EpsBearerCtxStatus** | Pointer to **string** | EPS bearer context status | [optional] 
**AmfNfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**Guami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**MaxIntegrityProtectedDataRateUl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**MaxIntegrityProtectedDataRateDl** | Pointer to [**MaxIntegrityProtectedDataRate**](MaxIntegrityProtectedDataRate.md) |  | [optional] 
**CpCiotEnabled** | Pointer to **bool** |  | [optional] [default to false]
**CpOnlyInd** | Pointer to **bool** |  | [optional] [default to false]
**InvokeNef** | Pointer to **bool** |  | [optional] [default to false]
**MaRequestInd** | Pointer to **bool** |  | [optional] [default to false]
**MaNwUpgradeInd** | Pointer to **bool** |  | [optional] [default to false]
**DnaiList** | Pointer to **[]string** |  | [optional] 
**PresenceInLadn** | Pointer to [**PresenceState**](PresenceState.md) |  | [optional] 
**SecondaryRatUsageInfo** | Pointer to [**[]SecondaryRatUsageInfo**](SecondaryRatUsageInfo.md) |  | [optional] 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**DlServingPlmnRateCtl** | Pointer to **int32** |  | [optional] 
**UpSecurityInfo** | Pointer to [**UpSecurityInfo**](UpSecurityInfo.md) |  | [optional] 
**VplmnQos** | Pointer to [**VplmnQos**](VplmnQos.md) |  | [optional] 
**OldSmContextRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**RedundantPduSessionInfo** | Pointer to [**RedundantPduSessionInformation**](RedundantPduSessionInformation.md) |  | [optional] 
**OldPduSessionRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmPolicyNotifyInd** | Pointer to **bool** |  | [optional] [default to false]
**PcfUeCallbackInfo** | Pointer to [**NullablePcfUeCallbackInfo**](PcfUeCallbackInfo.md) |  | [optional] 
**SatelliteBackhaulCat** | Pointer to [**SatelliteBackhaulCategory**](SatelliteBackhaulCategory.md) |  | [optional] 
**UpipSupported** | Pointer to **bool** |  | [optional] [default to false]
**UpCnxState** | Pointer to [**UpCnxState**](UpCnxState.md) |  | [optional] 
**DisasterRoamingInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPduSessionCreateData

`func NewPduSessionCreateData(dnn string, servingNetwork PlmnIdNid, anType AccessType, ) *PduSessionCreateData`

NewPduSessionCreateData instantiates a new PduSessionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionCreateDataWithDefaults

`func NewPduSessionCreateDataWithDefaults() *PduSessionCreateData`

NewPduSessionCreateDataWithDefaults instantiates a new PduSessionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *PduSessionCreateData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PduSessionCreateData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PduSessionCreateData) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PduSessionCreateData) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetUnauthenticatedSupi

`func (o *PduSessionCreateData) GetUnauthenticatedSupi() bool`

GetUnauthenticatedSupi returns the UnauthenticatedSupi field if non-nil, zero value otherwise.

### GetUnauthenticatedSupiOk

`func (o *PduSessionCreateData) GetUnauthenticatedSupiOk() (*bool, bool)`

GetUnauthenticatedSupiOk returns a tuple with the UnauthenticatedSupi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticatedSupi

`func (o *PduSessionCreateData) SetUnauthenticatedSupi(v bool)`

SetUnauthenticatedSupi sets UnauthenticatedSupi field to given value.

### HasUnauthenticatedSupi

`func (o *PduSessionCreateData) HasUnauthenticatedSupi() bool`

HasUnauthenticatedSupi returns a boolean if a field has been set.

### GetPei

`func (o *PduSessionCreateData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *PduSessionCreateData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *PduSessionCreateData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *PduSessionCreateData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetPduSessionId

`func (o *PduSessionCreateData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *PduSessionCreateData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *PduSessionCreateData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.

### HasPduSessionId

`func (o *PduSessionCreateData) HasPduSessionId() bool`

HasPduSessionId returns a boolean if a field has been set.

### GetDnn

`func (o *PduSessionCreateData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionCreateData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionCreateData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSelectedDnn

`func (o *PduSessionCreateData) GetSelectedDnn() string`

GetSelectedDnn returns the SelectedDnn field if non-nil, zero value otherwise.

### GetSelectedDnnOk

`func (o *PduSessionCreateData) GetSelectedDnnOk() (*string, bool)`

GetSelectedDnnOk returns a tuple with the SelectedDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDnn

`func (o *PduSessionCreateData) SetSelectedDnn(v string)`

SetSelectedDnn sets SelectedDnn field to given value.

### HasSelectedDnn

`func (o *PduSessionCreateData) HasSelectedDnn() bool`

HasSelectedDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *PduSessionCreateData) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *PduSessionCreateData) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *PduSessionCreateData) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *PduSessionCreateData) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetHplmnSnssai

`func (o *PduSessionCreateData) GetHplmnSnssai() Snssai`

GetHplmnSnssai returns the HplmnSnssai field if non-nil, zero value otherwise.

### GetHplmnSnssaiOk

`func (o *PduSessionCreateData) GetHplmnSnssaiOk() (*Snssai, bool)`

GetHplmnSnssaiOk returns a tuple with the HplmnSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHplmnSnssai

`func (o *PduSessionCreateData) SetHplmnSnssai(v Snssai)`

SetHplmnSnssai sets HplmnSnssai field to given value.

### HasHplmnSnssai

`func (o *PduSessionCreateData) HasHplmnSnssai() bool`

HasHplmnSnssai returns a boolean if a field has been set.

### GetVsmfId

`func (o *PduSessionCreateData) GetVsmfId() string`

GetVsmfId returns the VsmfId field if non-nil, zero value otherwise.

### GetVsmfIdOk

`func (o *PduSessionCreateData) GetVsmfIdOk() (*string, bool)`

GetVsmfIdOk returns a tuple with the VsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfId

`func (o *PduSessionCreateData) SetVsmfId(v string)`

SetVsmfId sets VsmfId field to given value.

### HasVsmfId

`func (o *PduSessionCreateData) HasVsmfId() bool`

HasVsmfId returns a boolean if a field has been set.

### GetIsmfId

`func (o *PduSessionCreateData) GetIsmfId() string`

GetIsmfId returns the IsmfId field if non-nil, zero value otherwise.

### GetIsmfIdOk

`func (o *PduSessionCreateData) GetIsmfIdOk() (*string, bool)`

GetIsmfIdOk returns a tuple with the IsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfId

`func (o *PduSessionCreateData) SetIsmfId(v string)`

SetIsmfId sets IsmfId field to given value.

### HasIsmfId

`func (o *PduSessionCreateData) HasIsmfId() bool`

HasIsmfId returns a boolean if a field has been set.

### GetServingNetwork

`func (o *PduSessionCreateData) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *PduSessionCreateData) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *PduSessionCreateData) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.


### GetRequestType

`func (o *PduSessionCreateData) GetRequestType() RequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *PduSessionCreateData) GetRequestTypeOk() (*RequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *PduSessionCreateData) SetRequestType(v RequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *PduSessionCreateData) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetEpsBearerId

`func (o *PduSessionCreateData) GetEpsBearerId() []int32`

GetEpsBearerId returns the EpsBearerId field if non-nil, zero value otherwise.

### GetEpsBearerIdOk

`func (o *PduSessionCreateData) GetEpsBearerIdOk() (*[]int32, bool)`

GetEpsBearerIdOk returns a tuple with the EpsBearerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerId

`func (o *PduSessionCreateData) SetEpsBearerId(v []int32)`

SetEpsBearerId sets EpsBearerId field to given value.

### HasEpsBearerId

`func (o *PduSessionCreateData) HasEpsBearerId() bool`

HasEpsBearerId returns a boolean if a field has been set.

### GetPgwS8cFteid

`func (o *PduSessionCreateData) GetPgwS8cFteid() string`

GetPgwS8cFteid returns the PgwS8cFteid field if non-nil, zero value otherwise.

### GetPgwS8cFteidOk

`func (o *PduSessionCreateData) GetPgwS8cFteidOk() (*string, bool)`

GetPgwS8cFteidOk returns a tuple with the PgwS8cFteid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwS8cFteid

`func (o *PduSessionCreateData) SetPgwS8cFteid(v string)`

SetPgwS8cFteid sets PgwS8cFteid field to given value.

### HasPgwS8cFteid

`func (o *PduSessionCreateData) HasPgwS8cFteid() bool`

HasPgwS8cFteid returns a boolean if a field has been set.

### GetVsmfPduSessionUri

`func (o *PduSessionCreateData) GetVsmfPduSessionUri() string`

GetVsmfPduSessionUri returns the VsmfPduSessionUri field if non-nil, zero value otherwise.

### GetVsmfPduSessionUriOk

`func (o *PduSessionCreateData) GetVsmfPduSessionUriOk() (*string, bool)`

GetVsmfPduSessionUriOk returns a tuple with the VsmfPduSessionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfPduSessionUri

`func (o *PduSessionCreateData) SetVsmfPduSessionUri(v string)`

SetVsmfPduSessionUri sets VsmfPduSessionUri field to given value.

### HasVsmfPduSessionUri

`func (o *PduSessionCreateData) HasVsmfPduSessionUri() bool`

HasVsmfPduSessionUri returns a boolean if a field has been set.

### GetIsmfPduSessionUri

`func (o *PduSessionCreateData) GetIsmfPduSessionUri() string`

GetIsmfPduSessionUri returns the IsmfPduSessionUri field if non-nil, zero value otherwise.

### GetIsmfPduSessionUriOk

`func (o *PduSessionCreateData) GetIsmfPduSessionUriOk() (*string, bool)`

GetIsmfPduSessionUriOk returns a tuple with the IsmfPduSessionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfPduSessionUri

`func (o *PduSessionCreateData) SetIsmfPduSessionUri(v string)`

SetIsmfPduSessionUri sets IsmfPduSessionUri field to given value.

### HasIsmfPduSessionUri

`func (o *PduSessionCreateData) HasIsmfPduSessionUri() bool`

HasIsmfPduSessionUri returns a boolean if a field has been set.

### GetVcnTunnelInfo

`func (o *PduSessionCreateData) GetVcnTunnelInfo() TunnelInfo`

GetVcnTunnelInfo returns the VcnTunnelInfo field if non-nil, zero value otherwise.

### GetVcnTunnelInfoOk

`func (o *PduSessionCreateData) GetVcnTunnelInfoOk() (*TunnelInfo, bool)`

GetVcnTunnelInfoOk returns a tuple with the VcnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcnTunnelInfo

`func (o *PduSessionCreateData) SetVcnTunnelInfo(v TunnelInfo)`

SetVcnTunnelInfo sets VcnTunnelInfo field to given value.

### HasVcnTunnelInfo

`func (o *PduSessionCreateData) HasVcnTunnelInfo() bool`

HasVcnTunnelInfo returns a boolean if a field has been set.

### GetIcnTunnelInfo

`func (o *PduSessionCreateData) GetIcnTunnelInfo() TunnelInfo`

GetIcnTunnelInfo returns the IcnTunnelInfo field if non-nil, zero value otherwise.

### GetIcnTunnelInfoOk

`func (o *PduSessionCreateData) GetIcnTunnelInfoOk() (*TunnelInfo, bool)`

GetIcnTunnelInfoOk returns a tuple with the IcnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcnTunnelInfo

`func (o *PduSessionCreateData) SetIcnTunnelInfo(v TunnelInfo)`

SetIcnTunnelInfo sets IcnTunnelInfo field to given value.

### HasIcnTunnelInfo

`func (o *PduSessionCreateData) HasIcnTunnelInfo() bool`

HasIcnTunnelInfo returns a boolean if a field has been set.

### GetN9ForwardingTunnelInfo

`func (o *PduSessionCreateData) GetN9ForwardingTunnelInfo() TunnelInfo`

GetN9ForwardingTunnelInfo returns the N9ForwardingTunnelInfo field if non-nil, zero value otherwise.

### GetN9ForwardingTunnelInfoOk

`func (o *PduSessionCreateData) GetN9ForwardingTunnelInfoOk() (*TunnelInfo, bool)`

GetN9ForwardingTunnelInfoOk returns a tuple with the N9ForwardingTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN9ForwardingTunnelInfo

`func (o *PduSessionCreateData) SetN9ForwardingTunnelInfo(v TunnelInfo)`

SetN9ForwardingTunnelInfo sets N9ForwardingTunnelInfo field to given value.

### HasN9ForwardingTunnelInfo

`func (o *PduSessionCreateData) HasN9ForwardingTunnelInfo() bool`

HasN9ForwardingTunnelInfo returns a boolean if a field has been set.

### GetAdditionalCnTunnelInfo

`func (o *PduSessionCreateData) GetAdditionalCnTunnelInfo() TunnelInfo`

GetAdditionalCnTunnelInfo returns the AdditionalCnTunnelInfo field if non-nil, zero value otherwise.

### GetAdditionalCnTunnelInfoOk

`func (o *PduSessionCreateData) GetAdditionalCnTunnelInfoOk() (*TunnelInfo, bool)`

GetAdditionalCnTunnelInfoOk returns a tuple with the AdditionalCnTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCnTunnelInfo

`func (o *PduSessionCreateData) SetAdditionalCnTunnelInfo(v TunnelInfo)`

SetAdditionalCnTunnelInfo sets AdditionalCnTunnelInfo field to given value.

### HasAdditionalCnTunnelInfo

`func (o *PduSessionCreateData) HasAdditionalCnTunnelInfo() bool`

HasAdditionalCnTunnelInfo returns a boolean if a field has been set.

### GetAnType

`func (o *PduSessionCreateData) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *PduSessionCreateData) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *PduSessionCreateData) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.


### GetAdditionalAnType

`func (o *PduSessionCreateData) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *PduSessionCreateData) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *PduSessionCreateData) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *PduSessionCreateData) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.

### GetRatType

`func (o *PduSessionCreateData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *PduSessionCreateData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *PduSessionCreateData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *PduSessionCreateData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetUeLocation

`func (o *PduSessionCreateData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *PduSessionCreateData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *PduSessionCreateData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *PduSessionCreateData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *PduSessionCreateData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *PduSessionCreateData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *PduSessionCreateData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *PduSessionCreateData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.

### GetAddUeLocation

`func (o *PduSessionCreateData) GetAddUeLocation() UserLocation`

GetAddUeLocation returns the AddUeLocation field if non-nil, zero value otherwise.

### GetAddUeLocationOk

`func (o *PduSessionCreateData) GetAddUeLocationOk() (*UserLocation, bool)`

GetAddUeLocationOk returns a tuple with the AddUeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddUeLocation

`func (o *PduSessionCreateData) SetAddUeLocation(v UserLocation)`

SetAddUeLocation sets AddUeLocation field to given value.

### HasAddUeLocation

`func (o *PduSessionCreateData) HasAddUeLocation() bool`

HasAddUeLocation returns a boolean if a field has been set.

### GetGpsi

`func (o *PduSessionCreateData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PduSessionCreateData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PduSessionCreateData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *PduSessionCreateData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetN1SmInfoFromUe

`func (o *PduSessionCreateData) GetN1SmInfoFromUe() RefToBinaryData`

GetN1SmInfoFromUe returns the N1SmInfoFromUe field if non-nil, zero value otherwise.

### GetN1SmInfoFromUeOk

`func (o *PduSessionCreateData) GetN1SmInfoFromUeOk() (*RefToBinaryData, bool)`

GetN1SmInfoFromUeOk returns a tuple with the N1SmInfoFromUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN1SmInfoFromUe

`func (o *PduSessionCreateData) SetN1SmInfoFromUe(v RefToBinaryData)`

SetN1SmInfoFromUe sets N1SmInfoFromUe field to given value.

### HasN1SmInfoFromUe

`func (o *PduSessionCreateData) HasN1SmInfoFromUe() bool`

HasN1SmInfoFromUe returns a boolean if a field has been set.

### GetUnknownN1SmInfo

`func (o *PduSessionCreateData) GetUnknownN1SmInfo() RefToBinaryData`

GetUnknownN1SmInfo returns the UnknownN1SmInfo field if non-nil, zero value otherwise.

### GetUnknownN1SmInfoOk

`func (o *PduSessionCreateData) GetUnknownN1SmInfoOk() (*RefToBinaryData, bool)`

GetUnknownN1SmInfoOk returns a tuple with the UnknownN1SmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownN1SmInfo

`func (o *PduSessionCreateData) SetUnknownN1SmInfo(v RefToBinaryData)`

SetUnknownN1SmInfo sets UnknownN1SmInfo field to given value.

### HasUnknownN1SmInfo

`func (o *PduSessionCreateData) HasUnknownN1SmInfo() bool`

HasUnknownN1SmInfo returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *PduSessionCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *PduSessionCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *PduSessionCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *PduSessionCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetHPcfId

`func (o *PduSessionCreateData) GetHPcfId() string`

GetHPcfId returns the HPcfId field if non-nil, zero value otherwise.

### GetHPcfIdOk

`func (o *PduSessionCreateData) GetHPcfIdOk() (*string, bool)`

GetHPcfIdOk returns a tuple with the HPcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHPcfId

`func (o *PduSessionCreateData) SetHPcfId(v string)`

SetHPcfId sets HPcfId field to given value.

### HasHPcfId

`func (o *PduSessionCreateData) HasHPcfId() bool`

HasHPcfId returns a boolean if a field has been set.

### GetPcfId

`func (o *PduSessionCreateData) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PduSessionCreateData) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PduSessionCreateData) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PduSessionCreateData) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfGroupId

`func (o *PduSessionCreateData) GetPcfGroupId() string`

GetPcfGroupId returns the PcfGroupId field if non-nil, zero value otherwise.

### GetPcfGroupIdOk

`func (o *PduSessionCreateData) GetPcfGroupIdOk() (*string, bool)`

GetPcfGroupIdOk returns a tuple with the PcfGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfGroupId

`func (o *PduSessionCreateData) SetPcfGroupId(v string)`

SetPcfGroupId sets PcfGroupId field to given value.

### HasPcfGroupId

`func (o *PduSessionCreateData) HasPcfGroupId() bool`

HasPcfGroupId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *PduSessionCreateData) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *PduSessionCreateData) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *PduSessionCreateData) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *PduSessionCreateData) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetHoPreparationIndication

`func (o *PduSessionCreateData) GetHoPreparationIndication() bool`

GetHoPreparationIndication returns the HoPreparationIndication field if non-nil, zero value otherwise.

### GetHoPreparationIndicationOk

`func (o *PduSessionCreateData) GetHoPreparationIndicationOk() (*bool, bool)`

GetHoPreparationIndicationOk returns a tuple with the HoPreparationIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoPreparationIndication

`func (o *PduSessionCreateData) SetHoPreparationIndication(v bool)`

SetHoPreparationIndication sets HoPreparationIndication field to given value.

### HasHoPreparationIndication

`func (o *PduSessionCreateData) HasHoPreparationIndication() bool`

HasHoPreparationIndication returns a boolean if a field has been set.

### GetSelMode

`func (o *PduSessionCreateData) GetSelMode() DnnSelectionMode`

GetSelMode returns the SelMode field if non-nil, zero value otherwise.

### GetSelModeOk

`func (o *PduSessionCreateData) GetSelModeOk() (*DnnSelectionMode, bool)`

GetSelModeOk returns a tuple with the SelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelMode

`func (o *PduSessionCreateData) SetSelMode(v DnnSelectionMode)`

SetSelMode sets SelMode field to given value.

### HasSelMode

`func (o *PduSessionCreateData) HasSelMode() bool`

HasSelMode returns a boolean if a field has been set.

### GetAlwaysOnRequested

`func (o *PduSessionCreateData) GetAlwaysOnRequested() bool`

GetAlwaysOnRequested returns the AlwaysOnRequested field if non-nil, zero value otherwise.

### GetAlwaysOnRequestedOk

`func (o *PduSessionCreateData) GetAlwaysOnRequestedOk() (*bool, bool)`

GetAlwaysOnRequestedOk returns a tuple with the AlwaysOnRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysOnRequested

`func (o *PduSessionCreateData) SetAlwaysOnRequested(v bool)`

SetAlwaysOnRequested sets AlwaysOnRequested field to given value.

### HasAlwaysOnRequested

`func (o *PduSessionCreateData) HasAlwaysOnRequested() bool`

HasAlwaysOnRequested returns a boolean if a field has been set.

### GetUdmGroupId

`func (o *PduSessionCreateData) GetUdmGroupId() string`

GetUdmGroupId returns the UdmGroupId field if non-nil, zero value otherwise.

### GetUdmGroupIdOk

`func (o *PduSessionCreateData) GetUdmGroupIdOk() (*string, bool)`

GetUdmGroupIdOk returns a tuple with the UdmGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmGroupId

`func (o *PduSessionCreateData) SetUdmGroupId(v string)`

SetUdmGroupId sets UdmGroupId field to given value.

### HasUdmGroupId

`func (o *PduSessionCreateData) HasUdmGroupId() bool`

HasUdmGroupId returns a boolean if a field has been set.

### GetRoutingIndicator

`func (o *PduSessionCreateData) GetRoutingIndicator() string`

GetRoutingIndicator returns the RoutingIndicator field if non-nil, zero value otherwise.

### GetRoutingIndicatorOk

`func (o *PduSessionCreateData) GetRoutingIndicatorOk() (*string, bool)`

GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingIndicator

`func (o *PduSessionCreateData) SetRoutingIndicator(v string)`

SetRoutingIndicator sets RoutingIndicator field to given value.

### HasRoutingIndicator

`func (o *PduSessionCreateData) HasRoutingIndicator() bool`

HasRoutingIndicator returns a boolean if a field has been set.

### GetHNwPubKeyId

`func (o *PduSessionCreateData) GetHNwPubKeyId() int32`

GetHNwPubKeyId returns the HNwPubKeyId field if non-nil, zero value otherwise.

### GetHNwPubKeyIdOk

`func (o *PduSessionCreateData) GetHNwPubKeyIdOk() (*int32, bool)`

GetHNwPubKeyIdOk returns a tuple with the HNwPubKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHNwPubKeyId

`func (o *PduSessionCreateData) SetHNwPubKeyId(v int32)`

SetHNwPubKeyId sets HNwPubKeyId field to given value.

### HasHNwPubKeyId

`func (o *PduSessionCreateData) HasHNwPubKeyId() bool`

HasHNwPubKeyId returns a boolean if a field has been set.

### GetEpsInterworkingInd

`func (o *PduSessionCreateData) GetEpsInterworkingInd() EpsInterworkingIndication`

GetEpsInterworkingInd returns the EpsInterworkingInd field if non-nil, zero value otherwise.

### GetEpsInterworkingIndOk

`func (o *PduSessionCreateData) GetEpsInterworkingIndOk() (*EpsInterworkingIndication, bool)`

GetEpsInterworkingIndOk returns a tuple with the EpsInterworkingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsInterworkingInd

`func (o *PduSessionCreateData) SetEpsInterworkingInd(v EpsInterworkingIndication)`

SetEpsInterworkingInd sets EpsInterworkingInd field to given value.

### HasEpsInterworkingInd

`func (o *PduSessionCreateData) HasEpsInterworkingInd() bool`

HasEpsInterworkingInd returns a boolean if a field has been set.

### GetVSmfServiceInstanceId

`func (o *PduSessionCreateData) GetVSmfServiceInstanceId() string`

GetVSmfServiceInstanceId returns the VSmfServiceInstanceId field if non-nil, zero value otherwise.

### GetVSmfServiceInstanceIdOk

`func (o *PduSessionCreateData) GetVSmfServiceInstanceIdOk() (*string, bool)`

GetVSmfServiceInstanceIdOk returns a tuple with the VSmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSmfServiceInstanceId

`func (o *PduSessionCreateData) SetVSmfServiceInstanceId(v string)`

SetVSmfServiceInstanceId sets VSmfServiceInstanceId field to given value.

### HasVSmfServiceInstanceId

`func (o *PduSessionCreateData) HasVSmfServiceInstanceId() bool`

HasVSmfServiceInstanceId returns a boolean if a field has been set.

### GetISmfServiceInstanceId

`func (o *PduSessionCreateData) GetISmfServiceInstanceId() string`

GetISmfServiceInstanceId returns the ISmfServiceInstanceId field if non-nil, zero value otherwise.

### GetISmfServiceInstanceIdOk

`func (o *PduSessionCreateData) GetISmfServiceInstanceIdOk() (*string, bool)`

GetISmfServiceInstanceIdOk returns a tuple with the ISmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISmfServiceInstanceId

`func (o *PduSessionCreateData) SetISmfServiceInstanceId(v string)`

SetISmfServiceInstanceId sets ISmfServiceInstanceId field to given value.

### HasISmfServiceInstanceId

`func (o *PduSessionCreateData) HasISmfServiceInstanceId() bool`

HasISmfServiceInstanceId returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *PduSessionCreateData) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *PduSessionCreateData) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *PduSessionCreateData) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *PduSessionCreateData) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetRoamingChargingProfile

`func (o *PduSessionCreateData) GetRoamingChargingProfile() RoamingChargingProfile`

GetRoamingChargingProfile returns the RoamingChargingProfile field if non-nil, zero value otherwise.

### GetRoamingChargingProfileOk

`func (o *PduSessionCreateData) GetRoamingChargingProfileOk() (*RoamingChargingProfile, bool)`

GetRoamingChargingProfileOk returns a tuple with the RoamingChargingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamingChargingProfile

`func (o *PduSessionCreateData) SetRoamingChargingProfile(v RoamingChargingProfile)`

SetRoamingChargingProfile sets RoamingChargingProfile field to given value.

### HasRoamingChargingProfile

`func (o *PduSessionCreateData) HasRoamingChargingProfile() bool`

HasRoamingChargingProfile returns a boolean if a field has been set.

### GetChargingId

`func (o *PduSessionCreateData) GetChargingId() string`

GetChargingId returns the ChargingId field if non-nil, zero value otherwise.

### GetChargingIdOk

`func (o *PduSessionCreateData) GetChargingIdOk() (*string, bool)`

GetChargingIdOk returns a tuple with the ChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingId

`func (o *PduSessionCreateData) SetChargingId(v string)`

SetChargingId sets ChargingId field to given value.

### HasChargingId

`func (o *PduSessionCreateData) HasChargingId() bool`

HasChargingId returns a boolean if a field has been set.

### GetOldPduSessionId

`func (o *PduSessionCreateData) GetOldPduSessionId() int32`

GetOldPduSessionId returns the OldPduSessionId field if non-nil, zero value otherwise.

### GetOldPduSessionIdOk

`func (o *PduSessionCreateData) GetOldPduSessionIdOk() (*int32, bool)`

GetOldPduSessionIdOk returns a tuple with the OldPduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionId

`func (o *PduSessionCreateData) SetOldPduSessionId(v int32)`

SetOldPduSessionId sets OldPduSessionId field to given value.

### HasOldPduSessionId

`func (o *PduSessionCreateData) HasOldPduSessionId() bool`

HasOldPduSessionId returns a boolean if a field has been set.

### GetEpsBearerCtxStatus

`func (o *PduSessionCreateData) GetEpsBearerCtxStatus() string`

GetEpsBearerCtxStatus returns the EpsBearerCtxStatus field if non-nil, zero value otherwise.

### GetEpsBearerCtxStatusOk

`func (o *PduSessionCreateData) GetEpsBearerCtxStatusOk() (*string, bool)`

GetEpsBearerCtxStatusOk returns a tuple with the EpsBearerCtxStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsBearerCtxStatus

`func (o *PduSessionCreateData) SetEpsBearerCtxStatus(v string)`

SetEpsBearerCtxStatus sets EpsBearerCtxStatus field to given value.

### HasEpsBearerCtxStatus

`func (o *PduSessionCreateData) HasEpsBearerCtxStatus() bool`

HasEpsBearerCtxStatus returns a boolean if a field has been set.

### GetAmfNfId

`func (o *PduSessionCreateData) GetAmfNfId() string`

GetAmfNfId returns the AmfNfId field if non-nil, zero value otherwise.

### GetAmfNfIdOk

`func (o *PduSessionCreateData) GetAmfNfIdOk() (*string, bool)`

GetAmfNfIdOk returns a tuple with the AmfNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfNfId

`func (o *PduSessionCreateData) SetAmfNfId(v string)`

SetAmfNfId sets AmfNfId field to given value.

### HasAmfNfId

`func (o *PduSessionCreateData) HasAmfNfId() bool`

HasAmfNfId returns a boolean if a field has been set.

### GetGuami

`func (o *PduSessionCreateData) GetGuami() Guami`

GetGuami returns the Guami field if non-nil, zero value otherwise.

### GetGuamiOk

`func (o *PduSessionCreateData) GetGuamiOk() (*Guami, bool)`

GetGuamiOk returns a tuple with the Guami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuami

`func (o *PduSessionCreateData) SetGuami(v Guami)`

SetGuami sets Guami field to given value.

### HasGuami

`func (o *PduSessionCreateData) HasGuami() bool`

HasGuami returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateUl

`func (o *PduSessionCreateData) GetMaxIntegrityProtectedDataRateUl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateUl returns the MaxIntegrityProtectedDataRateUl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateUlOk

`func (o *PduSessionCreateData) GetMaxIntegrityProtectedDataRateUlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateUlOk returns a tuple with the MaxIntegrityProtectedDataRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateUl

`func (o *PduSessionCreateData) SetMaxIntegrityProtectedDataRateUl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateUl sets MaxIntegrityProtectedDataRateUl field to given value.

### HasMaxIntegrityProtectedDataRateUl

`func (o *PduSessionCreateData) HasMaxIntegrityProtectedDataRateUl() bool`

HasMaxIntegrityProtectedDataRateUl returns a boolean if a field has been set.

### GetMaxIntegrityProtectedDataRateDl

`func (o *PduSessionCreateData) GetMaxIntegrityProtectedDataRateDl() MaxIntegrityProtectedDataRate`

GetMaxIntegrityProtectedDataRateDl returns the MaxIntegrityProtectedDataRateDl field if non-nil, zero value otherwise.

### GetMaxIntegrityProtectedDataRateDlOk

`func (o *PduSessionCreateData) GetMaxIntegrityProtectedDataRateDlOk() (*MaxIntegrityProtectedDataRate, bool)`

GetMaxIntegrityProtectedDataRateDlOk returns a tuple with the MaxIntegrityProtectedDataRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIntegrityProtectedDataRateDl

`func (o *PduSessionCreateData) SetMaxIntegrityProtectedDataRateDl(v MaxIntegrityProtectedDataRate)`

SetMaxIntegrityProtectedDataRateDl sets MaxIntegrityProtectedDataRateDl field to given value.

### HasMaxIntegrityProtectedDataRateDl

`func (o *PduSessionCreateData) HasMaxIntegrityProtectedDataRateDl() bool`

HasMaxIntegrityProtectedDataRateDl returns a boolean if a field has been set.

### GetCpCiotEnabled

`func (o *PduSessionCreateData) GetCpCiotEnabled() bool`

GetCpCiotEnabled returns the CpCiotEnabled field if non-nil, zero value otherwise.

### GetCpCiotEnabledOk

`func (o *PduSessionCreateData) GetCpCiotEnabledOk() (*bool, bool)`

GetCpCiotEnabledOk returns a tuple with the CpCiotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpCiotEnabled

`func (o *PduSessionCreateData) SetCpCiotEnabled(v bool)`

SetCpCiotEnabled sets CpCiotEnabled field to given value.

### HasCpCiotEnabled

`func (o *PduSessionCreateData) HasCpCiotEnabled() bool`

HasCpCiotEnabled returns a boolean if a field has been set.

### GetCpOnlyInd

`func (o *PduSessionCreateData) GetCpOnlyInd() bool`

GetCpOnlyInd returns the CpOnlyInd field if non-nil, zero value otherwise.

### GetCpOnlyIndOk

`func (o *PduSessionCreateData) GetCpOnlyIndOk() (*bool, bool)`

GetCpOnlyIndOk returns a tuple with the CpOnlyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpOnlyInd

`func (o *PduSessionCreateData) SetCpOnlyInd(v bool)`

SetCpOnlyInd sets CpOnlyInd field to given value.

### HasCpOnlyInd

`func (o *PduSessionCreateData) HasCpOnlyInd() bool`

HasCpOnlyInd returns a boolean if a field has been set.

### GetInvokeNef

`func (o *PduSessionCreateData) GetInvokeNef() bool`

GetInvokeNef returns the InvokeNef field if non-nil, zero value otherwise.

### GetInvokeNefOk

`func (o *PduSessionCreateData) GetInvokeNefOk() (*bool, bool)`

GetInvokeNefOk returns a tuple with the InvokeNef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeNef

`func (o *PduSessionCreateData) SetInvokeNef(v bool)`

SetInvokeNef sets InvokeNef field to given value.

### HasInvokeNef

`func (o *PduSessionCreateData) HasInvokeNef() bool`

HasInvokeNef returns a boolean if a field has been set.

### GetMaRequestInd

`func (o *PduSessionCreateData) GetMaRequestInd() bool`

GetMaRequestInd returns the MaRequestInd field if non-nil, zero value otherwise.

### GetMaRequestIndOk

`func (o *PduSessionCreateData) GetMaRequestIndOk() (*bool, bool)`

GetMaRequestIndOk returns a tuple with the MaRequestInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaRequestInd

`func (o *PduSessionCreateData) SetMaRequestInd(v bool)`

SetMaRequestInd sets MaRequestInd field to given value.

### HasMaRequestInd

`func (o *PduSessionCreateData) HasMaRequestInd() bool`

HasMaRequestInd returns a boolean if a field has been set.

### GetMaNwUpgradeInd

`func (o *PduSessionCreateData) GetMaNwUpgradeInd() bool`

GetMaNwUpgradeInd returns the MaNwUpgradeInd field if non-nil, zero value otherwise.

### GetMaNwUpgradeIndOk

`func (o *PduSessionCreateData) GetMaNwUpgradeIndOk() (*bool, bool)`

GetMaNwUpgradeIndOk returns a tuple with the MaNwUpgradeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaNwUpgradeInd

`func (o *PduSessionCreateData) SetMaNwUpgradeInd(v bool)`

SetMaNwUpgradeInd sets MaNwUpgradeInd field to given value.

### HasMaNwUpgradeInd

`func (o *PduSessionCreateData) HasMaNwUpgradeInd() bool`

HasMaNwUpgradeInd returns a boolean if a field has been set.

### GetDnaiList

`func (o *PduSessionCreateData) GetDnaiList() []string`

GetDnaiList returns the DnaiList field if non-nil, zero value otherwise.

### GetDnaiListOk

`func (o *PduSessionCreateData) GetDnaiListOk() (*[]string, bool)`

GetDnaiListOk returns a tuple with the DnaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiList

`func (o *PduSessionCreateData) SetDnaiList(v []string)`

SetDnaiList sets DnaiList field to given value.

### HasDnaiList

`func (o *PduSessionCreateData) HasDnaiList() bool`

HasDnaiList returns a boolean if a field has been set.

### GetPresenceInLadn

`func (o *PduSessionCreateData) GetPresenceInLadn() PresenceState`

GetPresenceInLadn returns the PresenceInLadn field if non-nil, zero value otherwise.

### GetPresenceInLadnOk

`func (o *PduSessionCreateData) GetPresenceInLadnOk() (*PresenceState, bool)`

GetPresenceInLadnOk returns a tuple with the PresenceInLadn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInLadn

`func (o *PduSessionCreateData) SetPresenceInLadn(v PresenceState)`

SetPresenceInLadn sets PresenceInLadn field to given value.

### HasPresenceInLadn

`func (o *PduSessionCreateData) HasPresenceInLadn() bool`

HasPresenceInLadn returns a boolean if a field has been set.

### GetSecondaryRatUsageInfo

`func (o *PduSessionCreateData) GetSecondaryRatUsageInfo() []SecondaryRatUsageInfo`

GetSecondaryRatUsageInfo returns the SecondaryRatUsageInfo field if non-nil, zero value otherwise.

### GetSecondaryRatUsageInfoOk

`func (o *PduSessionCreateData) GetSecondaryRatUsageInfoOk() (*[]SecondaryRatUsageInfo, bool)`

GetSecondaryRatUsageInfoOk returns a tuple with the SecondaryRatUsageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRatUsageInfo

`func (o *PduSessionCreateData) SetSecondaryRatUsageInfo(v []SecondaryRatUsageInfo)`

SetSecondaryRatUsageInfo sets SecondaryRatUsageInfo field to given value.

### HasSecondaryRatUsageInfo

`func (o *PduSessionCreateData) HasSecondaryRatUsageInfo() bool`

HasSecondaryRatUsageInfo returns a boolean if a field has been set.

### GetSmallDataRateStatus

`func (o *PduSessionCreateData) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *PduSessionCreateData) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *PduSessionCreateData) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *PduSessionCreateData) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *PduSessionCreateData) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *PduSessionCreateData) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *PduSessionCreateData) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *PduSessionCreateData) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetDlServingPlmnRateCtl

`func (o *PduSessionCreateData) GetDlServingPlmnRateCtl() int32`

GetDlServingPlmnRateCtl returns the DlServingPlmnRateCtl field if non-nil, zero value otherwise.

### GetDlServingPlmnRateCtlOk

`func (o *PduSessionCreateData) GetDlServingPlmnRateCtlOk() (*int32, bool)`

GetDlServingPlmnRateCtlOk returns a tuple with the DlServingPlmnRateCtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlServingPlmnRateCtl

`func (o *PduSessionCreateData) SetDlServingPlmnRateCtl(v int32)`

SetDlServingPlmnRateCtl sets DlServingPlmnRateCtl field to given value.

### HasDlServingPlmnRateCtl

`func (o *PduSessionCreateData) HasDlServingPlmnRateCtl() bool`

HasDlServingPlmnRateCtl returns a boolean if a field has been set.

### GetUpSecurityInfo

`func (o *PduSessionCreateData) GetUpSecurityInfo() UpSecurityInfo`

GetUpSecurityInfo returns the UpSecurityInfo field if non-nil, zero value otherwise.

### GetUpSecurityInfoOk

`func (o *PduSessionCreateData) GetUpSecurityInfoOk() (*UpSecurityInfo, bool)`

GetUpSecurityInfoOk returns a tuple with the UpSecurityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpSecurityInfo

`func (o *PduSessionCreateData) SetUpSecurityInfo(v UpSecurityInfo)`

SetUpSecurityInfo sets UpSecurityInfo field to given value.

### HasUpSecurityInfo

`func (o *PduSessionCreateData) HasUpSecurityInfo() bool`

HasUpSecurityInfo returns a boolean if a field has been set.

### GetVplmnQos

`func (o *PduSessionCreateData) GetVplmnQos() VplmnQos`

GetVplmnQos returns the VplmnQos field if non-nil, zero value otherwise.

### GetVplmnQosOk

`func (o *PduSessionCreateData) GetVplmnQosOk() (*VplmnQos, bool)`

GetVplmnQosOk returns a tuple with the VplmnQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVplmnQos

`func (o *PduSessionCreateData) SetVplmnQos(v VplmnQos)`

SetVplmnQos sets VplmnQos field to given value.

### HasVplmnQos

`func (o *PduSessionCreateData) HasVplmnQos() bool`

HasVplmnQos returns a boolean if a field has been set.

### GetOldSmContextRef

`func (o *PduSessionCreateData) GetOldSmContextRef() string`

GetOldSmContextRef returns the OldSmContextRef field if non-nil, zero value otherwise.

### GetOldSmContextRefOk

`func (o *PduSessionCreateData) GetOldSmContextRefOk() (*string, bool)`

GetOldSmContextRefOk returns a tuple with the OldSmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmContextRef

`func (o *PduSessionCreateData) SetOldSmContextRef(v string)`

SetOldSmContextRef sets OldSmContextRef field to given value.

### HasOldSmContextRef

`func (o *PduSessionCreateData) HasOldSmContextRef() bool`

HasOldSmContextRef returns a boolean if a field has been set.

### GetRedundantPduSessionInfo

`func (o *PduSessionCreateData) GetRedundantPduSessionInfo() RedundantPduSessionInformation`

GetRedundantPduSessionInfo returns the RedundantPduSessionInfo field if non-nil, zero value otherwise.

### GetRedundantPduSessionInfoOk

`func (o *PduSessionCreateData) GetRedundantPduSessionInfoOk() (*RedundantPduSessionInformation, bool)`

GetRedundantPduSessionInfoOk returns a tuple with the RedundantPduSessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundantPduSessionInfo

`func (o *PduSessionCreateData) SetRedundantPduSessionInfo(v RedundantPduSessionInformation)`

SetRedundantPduSessionInfo sets RedundantPduSessionInfo field to given value.

### HasRedundantPduSessionInfo

`func (o *PduSessionCreateData) HasRedundantPduSessionInfo() bool`

HasRedundantPduSessionInfo returns a boolean if a field has been set.

### GetOldPduSessionRef

`func (o *PduSessionCreateData) GetOldPduSessionRef() string`

GetOldPduSessionRef returns the OldPduSessionRef field if non-nil, zero value otherwise.

### GetOldPduSessionRefOk

`func (o *PduSessionCreateData) GetOldPduSessionRefOk() (*string, bool)`

GetOldPduSessionRefOk returns a tuple with the OldPduSessionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionRef

`func (o *PduSessionCreateData) SetOldPduSessionRef(v string)`

SetOldPduSessionRef sets OldPduSessionRef field to given value.

### HasOldPduSessionRef

`func (o *PduSessionCreateData) HasOldPduSessionRef() bool`

HasOldPduSessionRef returns a boolean if a field has been set.

### GetSmPolicyNotifyInd

`func (o *PduSessionCreateData) GetSmPolicyNotifyInd() bool`

GetSmPolicyNotifyInd returns the SmPolicyNotifyInd field if non-nil, zero value otherwise.

### GetSmPolicyNotifyIndOk

`func (o *PduSessionCreateData) GetSmPolicyNotifyIndOk() (*bool, bool)`

GetSmPolicyNotifyIndOk returns a tuple with the SmPolicyNotifyInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmPolicyNotifyInd

`func (o *PduSessionCreateData) SetSmPolicyNotifyInd(v bool)`

SetSmPolicyNotifyInd sets SmPolicyNotifyInd field to given value.

### HasSmPolicyNotifyInd

`func (o *PduSessionCreateData) HasSmPolicyNotifyInd() bool`

HasSmPolicyNotifyInd returns a boolean if a field has been set.

### GetPcfUeCallbackInfo

`func (o *PduSessionCreateData) GetPcfUeCallbackInfo() PcfUeCallbackInfo`

GetPcfUeCallbackInfo returns the PcfUeCallbackInfo field if non-nil, zero value otherwise.

### GetPcfUeCallbackInfoOk

`func (o *PduSessionCreateData) GetPcfUeCallbackInfoOk() (*PcfUeCallbackInfo, bool)`

GetPcfUeCallbackInfoOk returns a tuple with the PcfUeCallbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfUeCallbackInfo

`func (o *PduSessionCreateData) SetPcfUeCallbackInfo(v PcfUeCallbackInfo)`

SetPcfUeCallbackInfo sets PcfUeCallbackInfo field to given value.

### HasPcfUeCallbackInfo

`func (o *PduSessionCreateData) HasPcfUeCallbackInfo() bool`

HasPcfUeCallbackInfo returns a boolean if a field has been set.

### SetPcfUeCallbackInfoNil

`func (o *PduSessionCreateData) SetPcfUeCallbackInfoNil(b bool)`

 SetPcfUeCallbackInfoNil sets the value for PcfUeCallbackInfo to be an explicit nil

### UnsetPcfUeCallbackInfo
`func (o *PduSessionCreateData) UnsetPcfUeCallbackInfo()`

UnsetPcfUeCallbackInfo ensures that no value is present for PcfUeCallbackInfo, not even an explicit nil
### GetSatelliteBackhaulCat

`func (o *PduSessionCreateData) GetSatelliteBackhaulCat() SatelliteBackhaulCategory`

GetSatelliteBackhaulCat returns the SatelliteBackhaulCat field if non-nil, zero value otherwise.

### GetSatelliteBackhaulCatOk

`func (o *PduSessionCreateData) GetSatelliteBackhaulCatOk() (*SatelliteBackhaulCategory, bool)`

GetSatelliteBackhaulCatOk returns a tuple with the SatelliteBackhaulCat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSatelliteBackhaulCat

`func (o *PduSessionCreateData) SetSatelliteBackhaulCat(v SatelliteBackhaulCategory)`

SetSatelliteBackhaulCat sets SatelliteBackhaulCat field to given value.

### HasSatelliteBackhaulCat

`func (o *PduSessionCreateData) HasSatelliteBackhaulCat() bool`

HasSatelliteBackhaulCat returns a boolean if a field has been set.

### GetUpipSupported

`func (o *PduSessionCreateData) GetUpipSupported() bool`

GetUpipSupported returns the UpipSupported field if non-nil, zero value otherwise.

### GetUpipSupportedOk

`func (o *PduSessionCreateData) GetUpipSupportedOk() (*bool, bool)`

GetUpipSupportedOk returns a tuple with the UpipSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpipSupported

`func (o *PduSessionCreateData) SetUpipSupported(v bool)`

SetUpipSupported sets UpipSupported field to given value.

### HasUpipSupported

`func (o *PduSessionCreateData) HasUpipSupported() bool`

HasUpipSupported returns a boolean if a field has been set.

### GetUpCnxState

`func (o *PduSessionCreateData) GetUpCnxState() UpCnxState`

GetUpCnxState returns the UpCnxState field if non-nil, zero value otherwise.

### GetUpCnxStateOk

`func (o *PduSessionCreateData) GetUpCnxStateOk() (*UpCnxState, bool)`

GetUpCnxStateOk returns a tuple with the UpCnxState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpCnxState

`func (o *PduSessionCreateData) SetUpCnxState(v UpCnxState)`

SetUpCnxState sets UpCnxState field to given value.

### HasUpCnxState

`func (o *PduSessionCreateData) HasUpCnxState() bool`

HasUpCnxState returns a boolean if a field has been set.

### GetDisasterRoamingInd

`func (o *PduSessionCreateData) GetDisasterRoamingInd() bool`

GetDisasterRoamingInd returns the DisasterRoamingInd field if non-nil, zero value otherwise.

### GetDisasterRoamingIndOk

`func (o *PduSessionCreateData) GetDisasterRoamingIndOk() (*bool, bool)`

GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisasterRoamingInd

`func (o *PduSessionCreateData) SetDisasterRoamingInd(v bool)`

SetDisasterRoamingInd sets DisasterRoamingInd field to given value.

### HasDisasterRoamingInd

`func (o *PduSessionCreateData) HasDisasterRoamingInd() bool`

HasDisasterRoamingInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


