# PduSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**SmContextRef** | **string** | String providing an URI formatted according to RFC 3986. | 
**SNssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**SelectedDnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**AccessType** | [**AccessType**](AccessType.md) |  | 
**AdditionalAccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**AllocatedEbiList** | Pointer to [**[]EbiArpMapping**](EbiArpMapping.md) |  | [optional] 
**HsmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**HsmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**HsmfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**SmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**VsmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**VsmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**VsmfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**VsmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**IsmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**IsmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**IsmfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**IsmfBinding** | Pointer to [**SbiBindingLevel**](SbiBindingLevel.md) |  | [optional] 
**NsInstance** | Pointer to **string** | Contains the Identifier of the selected Network Slice instance | [optional] 
**SmfServiceInstanceId** | Pointer to **string** |  | [optional] 
**MaPduSession** | Pointer to **bool** |  | [optional] [default to false]
**CnAssistedRanPara** | Pointer to [**CnAssistedRanPara**](CnAssistedRanPara.md) |  | [optional] 
**NrfManagementUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfDiscoveryUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**NrfAccessTokenUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SmfBindingInfo** | Pointer to **string** |  | [optional] 
**VsmfBindingInfo** | Pointer to **string** |  | [optional] 
**IsmfBindingInfo** | Pointer to **string** |  | [optional] 
**AdditionalSnssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**InterPlmnApiRoot** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**PgwFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PgwIpAddr** | Pointer to [**IpAddress**](IpAddress.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**AnchorSmfSupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AnchorSmfOauth2Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewPduSessionContext

`func NewPduSessionContext(pduSessionId int32, smContextRef string, sNssai Snssai, dnn string, accessType AccessType, ) *PduSessionContext`

NewPduSessionContext instantiates a new PduSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduSessionContextWithDefaults

`func NewPduSessionContextWithDefaults() *PduSessionContext`

NewPduSessionContextWithDefaults instantiates a new PduSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPduSessionId

`func (o *PduSessionContext) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *PduSessionContext) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *PduSessionContext) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetSmContextRef

`func (o *PduSessionContext) GetSmContextRef() string`

GetSmContextRef returns the SmContextRef field if non-nil, zero value otherwise.

### GetSmContextRefOk

`func (o *PduSessionContext) GetSmContextRefOk() (*string, bool)`

GetSmContextRefOk returns a tuple with the SmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextRef

`func (o *PduSessionContext) SetSmContextRef(v string)`

SetSmContextRef sets SmContextRef field to given value.


### GetSNssai

`func (o *PduSessionContext) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *PduSessionContext) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *PduSessionContext) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.


### GetDnn

`func (o *PduSessionContext) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *PduSessionContext) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *PduSessionContext) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSelectedDnn

`func (o *PduSessionContext) GetSelectedDnn() string`

GetSelectedDnn returns the SelectedDnn field if non-nil, zero value otherwise.

### GetSelectedDnnOk

`func (o *PduSessionContext) GetSelectedDnnOk() (*string, bool)`

GetSelectedDnnOk returns a tuple with the SelectedDnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedDnn

`func (o *PduSessionContext) SetSelectedDnn(v string)`

SetSelectedDnn sets SelectedDnn field to given value.

### HasSelectedDnn

`func (o *PduSessionContext) HasSelectedDnn() bool`

HasSelectedDnn returns a boolean if a field has been set.

### GetAccessType

`func (o *PduSessionContext) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PduSessionContext) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PduSessionContext) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.


### GetAdditionalAccessType

`func (o *PduSessionContext) GetAdditionalAccessType() AccessType`

GetAdditionalAccessType returns the AdditionalAccessType field if non-nil, zero value otherwise.

### GetAdditionalAccessTypeOk

`func (o *PduSessionContext) GetAdditionalAccessTypeOk() (*AccessType, bool)`

GetAdditionalAccessTypeOk returns a tuple with the AdditionalAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAccessType

`func (o *PduSessionContext) SetAdditionalAccessType(v AccessType)`

SetAdditionalAccessType sets AdditionalAccessType field to given value.

### HasAdditionalAccessType

`func (o *PduSessionContext) HasAdditionalAccessType() bool`

HasAdditionalAccessType returns a boolean if a field has been set.

### GetAllocatedEbiList

`func (o *PduSessionContext) GetAllocatedEbiList() []EbiArpMapping`

GetAllocatedEbiList returns the AllocatedEbiList field if non-nil, zero value otherwise.

### GetAllocatedEbiListOk

`func (o *PduSessionContext) GetAllocatedEbiListOk() (*[]EbiArpMapping, bool)`

GetAllocatedEbiListOk returns a tuple with the AllocatedEbiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedEbiList

`func (o *PduSessionContext) SetAllocatedEbiList(v []EbiArpMapping)`

SetAllocatedEbiList sets AllocatedEbiList field to given value.

### HasAllocatedEbiList

`func (o *PduSessionContext) HasAllocatedEbiList() bool`

HasAllocatedEbiList returns a boolean if a field has been set.

### GetHsmfId

`func (o *PduSessionContext) GetHsmfId() string`

GetHsmfId returns the HsmfId field if non-nil, zero value otherwise.

### GetHsmfIdOk

`func (o *PduSessionContext) GetHsmfIdOk() (*string, bool)`

GetHsmfIdOk returns a tuple with the HsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfId

`func (o *PduSessionContext) SetHsmfId(v string)`

SetHsmfId sets HsmfId field to given value.

### HasHsmfId

`func (o *PduSessionContext) HasHsmfId() bool`

HasHsmfId returns a boolean if a field has been set.

### GetHsmfSetId

`func (o *PduSessionContext) GetHsmfSetId() string`

GetHsmfSetId returns the HsmfSetId field if non-nil, zero value otherwise.

### GetHsmfSetIdOk

`func (o *PduSessionContext) GetHsmfSetIdOk() (*string, bool)`

GetHsmfSetIdOk returns a tuple with the HsmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfSetId

`func (o *PduSessionContext) SetHsmfSetId(v string)`

SetHsmfSetId sets HsmfSetId field to given value.

### HasHsmfSetId

`func (o *PduSessionContext) HasHsmfSetId() bool`

HasHsmfSetId returns a boolean if a field has been set.

### GetHsmfServiceSetId

`func (o *PduSessionContext) GetHsmfServiceSetId() string`

GetHsmfServiceSetId returns the HsmfServiceSetId field if non-nil, zero value otherwise.

### GetHsmfServiceSetIdOk

`func (o *PduSessionContext) GetHsmfServiceSetIdOk() (*string, bool)`

GetHsmfServiceSetIdOk returns a tuple with the HsmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsmfServiceSetId

`func (o *PduSessionContext) SetHsmfServiceSetId(v string)`

SetHsmfServiceSetId sets HsmfServiceSetId field to given value.

### HasHsmfServiceSetId

`func (o *PduSessionContext) HasHsmfServiceSetId() bool`

HasHsmfServiceSetId returns a boolean if a field has been set.

### GetSmfBinding

`func (o *PduSessionContext) GetSmfBinding() SbiBindingLevel`

GetSmfBinding returns the SmfBinding field if non-nil, zero value otherwise.

### GetSmfBindingOk

`func (o *PduSessionContext) GetSmfBindingOk() (*SbiBindingLevel, bool)`

GetSmfBindingOk returns a tuple with the SmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfBinding

`func (o *PduSessionContext) SetSmfBinding(v SbiBindingLevel)`

SetSmfBinding sets SmfBinding field to given value.

### HasSmfBinding

`func (o *PduSessionContext) HasSmfBinding() bool`

HasSmfBinding returns a boolean if a field has been set.

### GetVsmfId

`func (o *PduSessionContext) GetVsmfId() string`

GetVsmfId returns the VsmfId field if non-nil, zero value otherwise.

### GetVsmfIdOk

`func (o *PduSessionContext) GetVsmfIdOk() (*string, bool)`

GetVsmfIdOk returns a tuple with the VsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfId

`func (o *PduSessionContext) SetVsmfId(v string)`

SetVsmfId sets VsmfId field to given value.

### HasVsmfId

`func (o *PduSessionContext) HasVsmfId() bool`

HasVsmfId returns a boolean if a field has been set.

### GetVsmfSetId

`func (o *PduSessionContext) GetVsmfSetId() string`

GetVsmfSetId returns the VsmfSetId field if non-nil, zero value otherwise.

### GetVsmfSetIdOk

`func (o *PduSessionContext) GetVsmfSetIdOk() (*string, bool)`

GetVsmfSetIdOk returns a tuple with the VsmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfSetId

`func (o *PduSessionContext) SetVsmfSetId(v string)`

SetVsmfSetId sets VsmfSetId field to given value.

### HasVsmfSetId

`func (o *PduSessionContext) HasVsmfSetId() bool`

HasVsmfSetId returns a boolean if a field has been set.

### GetVsmfServiceSetId

`func (o *PduSessionContext) GetVsmfServiceSetId() string`

GetVsmfServiceSetId returns the VsmfServiceSetId field if non-nil, zero value otherwise.

### GetVsmfServiceSetIdOk

`func (o *PduSessionContext) GetVsmfServiceSetIdOk() (*string, bool)`

GetVsmfServiceSetIdOk returns a tuple with the VsmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfServiceSetId

`func (o *PduSessionContext) SetVsmfServiceSetId(v string)`

SetVsmfServiceSetId sets VsmfServiceSetId field to given value.

### HasVsmfServiceSetId

`func (o *PduSessionContext) HasVsmfServiceSetId() bool`

HasVsmfServiceSetId returns a boolean if a field has been set.

### GetVsmfBinding

`func (o *PduSessionContext) GetVsmfBinding() SbiBindingLevel`

GetVsmfBinding returns the VsmfBinding field if non-nil, zero value otherwise.

### GetVsmfBindingOk

`func (o *PduSessionContext) GetVsmfBindingOk() (*SbiBindingLevel, bool)`

GetVsmfBindingOk returns a tuple with the VsmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfBinding

`func (o *PduSessionContext) SetVsmfBinding(v SbiBindingLevel)`

SetVsmfBinding sets VsmfBinding field to given value.

### HasVsmfBinding

`func (o *PduSessionContext) HasVsmfBinding() bool`

HasVsmfBinding returns a boolean if a field has been set.

### GetIsmfId

`func (o *PduSessionContext) GetIsmfId() string`

GetIsmfId returns the IsmfId field if non-nil, zero value otherwise.

### GetIsmfIdOk

`func (o *PduSessionContext) GetIsmfIdOk() (*string, bool)`

GetIsmfIdOk returns a tuple with the IsmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfId

`func (o *PduSessionContext) SetIsmfId(v string)`

SetIsmfId sets IsmfId field to given value.

### HasIsmfId

`func (o *PduSessionContext) HasIsmfId() bool`

HasIsmfId returns a boolean if a field has been set.

### GetIsmfSetId

`func (o *PduSessionContext) GetIsmfSetId() string`

GetIsmfSetId returns the IsmfSetId field if non-nil, zero value otherwise.

### GetIsmfSetIdOk

`func (o *PduSessionContext) GetIsmfSetIdOk() (*string, bool)`

GetIsmfSetIdOk returns a tuple with the IsmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfSetId

`func (o *PduSessionContext) SetIsmfSetId(v string)`

SetIsmfSetId sets IsmfSetId field to given value.

### HasIsmfSetId

`func (o *PduSessionContext) HasIsmfSetId() bool`

HasIsmfSetId returns a boolean if a field has been set.

### GetIsmfServiceSetId

`func (o *PduSessionContext) GetIsmfServiceSetId() string`

GetIsmfServiceSetId returns the IsmfServiceSetId field if non-nil, zero value otherwise.

### GetIsmfServiceSetIdOk

`func (o *PduSessionContext) GetIsmfServiceSetIdOk() (*string, bool)`

GetIsmfServiceSetIdOk returns a tuple with the IsmfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfServiceSetId

`func (o *PduSessionContext) SetIsmfServiceSetId(v string)`

SetIsmfServiceSetId sets IsmfServiceSetId field to given value.

### HasIsmfServiceSetId

`func (o *PduSessionContext) HasIsmfServiceSetId() bool`

HasIsmfServiceSetId returns a boolean if a field has been set.

### GetIsmfBinding

`func (o *PduSessionContext) GetIsmfBinding() SbiBindingLevel`

GetIsmfBinding returns the IsmfBinding field if non-nil, zero value otherwise.

### GetIsmfBindingOk

`func (o *PduSessionContext) GetIsmfBindingOk() (*SbiBindingLevel, bool)`

GetIsmfBindingOk returns a tuple with the IsmfBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfBinding

`func (o *PduSessionContext) SetIsmfBinding(v SbiBindingLevel)`

SetIsmfBinding sets IsmfBinding field to given value.

### HasIsmfBinding

`func (o *PduSessionContext) HasIsmfBinding() bool`

HasIsmfBinding returns a boolean if a field has been set.

### GetNsInstance

`func (o *PduSessionContext) GetNsInstance() string`

GetNsInstance returns the NsInstance field if non-nil, zero value otherwise.

### GetNsInstanceOk

`func (o *PduSessionContext) GetNsInstanceOk() (*string, bool)`

GetNsInstanceOk returns a tuple with the NsInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsInstance

`func (o *PduSessionContext) SetNsInstance(v string)`

SetNsInstance sets NsInstance field to given value.

### HasNsInstance

`func (o *PduSessionContext) HasNsInstance() bool`

HasNsInstance returns a boolean if a field has been set.

### GetSmfServiceInstanceId

`func (o *PduSessionContext) GetSmfServiceInstanceId() string`

GetSmfServiceInstanceId returns the SmfServiceInstanceId field if non-nil, zero value otherwise.

### GetSmfServiceInstanceIdOk

`func (o *PduSessionContext) GetSmfServiceInstanceIdOk() (*string, bool)`

GetSmfServiceInstanceIdOk returns a tuple with the SmfServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfServiceInstanceId

`func (o *PduSessionContext) SetSmfServiceInstanceId(v string)`

SetSmfServiceInstanceId sets SmfServiceInstanceId field to given value.

### HasSmfServiceInstanceId

`func (o *PduSessionContext) HasSmfServiceInstanceId() bool`

HasSmfServiceInstanceId returns a boolean if a field has been set.

### GetMaPduSession

`func (o *PduSessionContext) GetMaPduSession() bool`

GetMaPduSession returns the MaPduSession field if non-nil, zero value otherwise.

### GetMaPduSessionOk

`func (o *PduSessionContext) GetMaPduSessionOk() (*bool, bool)`

GetMaPduSessionOk returns a tuple with the MaPduSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaPduSession

`func (o *PduSessionContext) SetMaPduSession(v bool)`

SetMaPduSession sets MaPduSession field to given value.

### HasMaPduSession

`func (o *PduSessionContext) HasMaPduSession() bool`

HasMaPduSession returns a boolean if a field has been set.

### GetCnAssistedRanPara

`func (o *PduSessionContext) GetCnAssistedRanPara() CnAssistedRanPara`

GetCnAssistedRanPara returns the CnAssistedRanPara field if non-nil, zero value otherwise.

### GetCnAssistedRanParaOk

`func (o *PduSessionContext) GetCnAssistedRanParaOk() (*CnAssistedRanPara, bool)`

GetCnAssistedRanParaOk returns a tuple with the CnAssistedRanPara field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnAssistedRanPara

`func (o *PduSessionContext) SetCnAssistedRanPara(v CnAssistedRanPara)`

SetCnAssistedRanPara sets CnAssistedRanPara field to given value.

### HasCnAssistedRanPara

`func (o *PduSessionContext) HasCnAssistedRanPara() bool`

HasCnAssistedRanPara returns a boolean if a field has been set.

### GetNrfManagementUri

`func (o *PduSessionContext) GetNrfManagementUri() string`

GetNrfManagementUri returns the NrfManagementUri field if non-nil, zero value otherwise.

### GetNrfManagementUriOk

`func (o *PduSessionContext) GetNrfManagementUriOk() (*string, bool)`

GetNrfManagementUriOk returns a tuple with the NrfManagementUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfManagementUri

`func (o *PduSessionContext) SetNrfManagementUri(v string)`

SetNrfManagementUri sets NrfManagementUri field to given value.

### HasNrfManagementUri

`func (o *PduSessionContext) HasNrfManagementUri() bool`

HasNrfManagementUri returns a boolean if a field has been set.

### GetNrfDiscoveryUri

`func (o *PduSessionContext) GetNrfDiscoveryUri() string`

GetNrfDiscoveryUri returns the NrfDiscoveryUri field if non-nil, zero value otherwise.

### GetNrfDiscoveryUriOk

`func (o *PduSessionContext) GetNrfDiscoveryUriOk() (*string, bool)`

GetNrfDiscoveryUriOk returns a tuple with the NrfDiscoveryUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfDiscoveryUri

`func (o *PduSessionContext) SetNrfDiscoveryUri(v string)`

SetNrfDiscoveryUri sets NrfDiscoveryUri field to given value.

### HasNrfDiscoveryUri

`func (o *PduSessionContext) HasNrfDiscoveryUri() bool`

HasNrfDiscoveryUri returns a boolean if a field has been set.

### GetNrfAccessTokenUri

`func (o *PduSessionContext) GetNrfAccessTokenUri() string`

GetNrfAccessTokenUri returns the NrfAccessTokenUri field if non-nil, zero value otherwise.

### GetNrfAccessTokenUriOk

`func (o *PduSessionContext) GetNrfAccessTokenUriOk() (*string, bool)`

GetNrfAccessTokenUriOk returns a tuple with the NrfAccessTokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfAccessTokenUri

`func (o *PduSessionContext) SetNrfAccessTokenUri(v string)`

SetNrfAccessTokenUri sets NrfAccessTokenUri field to given value.

### HasNrfAccessTokenUri

`func (o *PduSessionContext) HasNrfAccessTokenUri() bool`

HasNrfAccessTokenUri returns a boolean if a field has been set.

### GetSmfBindingInfo

`func (o *PduSessionContext) GetSmfBindingInfo() string`

GetSmfBindingInfo returns the SmfBindingInfo field if non-nil, zero value otherwise.

### GetSmfBindingInfoOk

`func (o *PduSessionContext) GetSmfBindingInfoOk() (*string, bool)`

GetSmfBindingInfoOk returns a tuple with the SmfBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfBindingInfo

`func (o *PduSessionContext) SetSmfBindingInfo(v string)`

SetSmfBindingInfo sets SmfBindingInfo field to given value.

### HasSmfBindingInfo

`func (o *PduSessionContext) HasSmfBindingInfo() bool`

HasSmfBindingInfo returns a boolean if a field has been set.

### GetVsmfBindingInfo

`func (o *PduSessionContext) GetVsmfBindingInfo() string`

GetVsmfBindingInfo returns the VsmfBindingInfo field if non-nil, zero value otherwise.

### GetVsmfBindingInfoOk

`func (o *PduSessionContext) GetVsmfBindingInfoOk() (*string, bool)`

GetVsmfBindingInfoOk returns a tuple with the VsmfBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsmfBindingInfo

`func (o *PduSessionContext) SetVsmfBindingInfo(v string)`

SetVsmfBindingInfo sets VsmfBindingInfo field to given value.

### HasVsmfBindingInfo

`func (o *PduSessionContext) HasVsmfBindingInfo() bool`

HasVsmfBindingInfo returns a boolean if a field has been set.

### GetIsmfBindingInfo

`func (o *PduSessionContext) GetIsmfBindingInfo() string`

GetIsmfBindingInfo returns the IsmfBindingInfo field if non-nil, zero value otherwise.

### GetIsmfBindingInfoOk

`func (o *PduSessionContext) GetIsmfBindingInfoOk() (*string, bool)`

GetIsmfBindingInfoOk returns a tuple with the IsmfBindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsmfBindingInfo

`func (o *PduSessionContext) SetIsmfBindingInfo(v string)`

SetIsmfBindingInfo sets IsmfBindingInfo field to given value.

### HasIsmfBindingInfo

`func (o *PduSessionContext) HasIsmfBindingInfo() bool`

HasIsmfBindingInfo returns a boolean if a field has been set.

### GetAdditionalSnssai

`func (o *PduSessionContext) GetAdditionalSnssai() Snssai`

GetAdditionalSnssai returns the AdditionalSnssai field if non-nil, zero value otherwise.

### GetAdditionalSnssaiOk

`func (o *PduSessionContext) GetAdditionalSnssaiOk() (*Snssai, bool)`

GetAdditionalSnssaiOk returns a tuple with the AdditionalSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSnssai

`func (o *PduSessionContext) SetAdditionalSnssai(v Snssai)`

SetAdditionalSnssai sets AdditionalSnssai field to given value.

### HasAdditionalSnssai

`func (o *PduSessionContext) HasAdditionalSnssai() bool`

HasAdditionalSnssai returns a boolean if a field has been set.

### GetInterPlmnApiRoot

`func (o *PduSessionContext) GetInterPlmnApiRoot() string`

GetInterPlmnApiRoot returns the InterPlmnApiRoot field if non-nil, zero value otherwise.

### GetInterPlmnApiRootOk

`func (o *PduSessionContext) GetInterPlmnApiRootOk() (*string, bool)`

GetInterPlmnApiRootOk returns a tuple with the InterPlmnApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnApiRoot

`func (o *PduSessionContext) SetInterPlmnApiRoot(v string)`

SetInterPlmnApiRoot sets InterPlmnApiRoot field to given value.

### HasInterPlmnApiRoot

`func (o *PduSessionContext) HasInterPlmnApiRoot() bool`

HasInterPlmnApiRoot returns a boolean if a field has been set.

### GetPgwFqdn

`func (o *PduSessionContext) GetPgwFqdn() string`

GetPgwFqdn returns the PgwFqdn field if non-nil, zero value otherwise.

### GetPgwFqdnOk

`func (o *PduSessionContext) GetPgwFqdnOk() (*string, bool)`

GetPgwFqdnOk returns a tuple with the PgwFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwFqdn

`func (o *PduSessionContext) SetPgwFqdn(v string)`

SetPgwFqdn sets PgwFqdn field to given value.

### HasPgwFqdn

`func (o *PduSessionContext) HasPgwFqdn() bool`

HasPgwFqdn returns a boolean if a field has been set.

### GetPgwIpAddr

`func (o *PduSessionContext) GetPgwIpAddr() IpAddress`

GetPgwIpAddr returns the PgwIpAddr field if non-nil, zero value otherwise.

### GetPgwIpAddrOk

`func (o *PduSessionContext) GetPgwIpAddrOk() (*IpAddress, bool)`

GetPgwIpAddrOk returns a tuple with the PgwIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgwIpAddr

`func (o *PduSessionContext) SetPgwIpAddr(v IpAddress)`

SetPgwIpAddr sets PgwIpAddr field to given value.

### HasPgwIpAddr

`func (o *PduSessionContext) HasPgwIpAddr() bool`

HasPgwIpAddr returns a boolean if a field has been set.

### GetPlmnId

`func (o *PduSessionContext) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PduSessionContext) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PduSessionContext) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *PduSessionContext) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetAnchorSmfSupportedFeatures

`func (o *PduSessionContext) GetAnchorSmfSupportedFeatures() string`

GetAnchorSmfSupportedFeatures returns the AnchorSmfSupportedFeatures field if non-nil, zero value otherwise.

### GetAnchorSmfSupportedFeaturesOk

`func (o *PduSessionContext) GetAnchorSmfSupportedFeaturesOk() (*string, bool)`

GetAnchorSmfSupportedFeaturesOk returns a tuple with the AnchorSmfSupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorSmfSupportedFeatures

`func (o *PduSessionContext) SetAnchorSmfSupportedFeatures(v string)`

SetAnchorSmfSupportedFeatures sets AnchorSmfSupportedFeatures field to given value.

### HasAnchorSmfSupportedFeatures

`func (o *PduSessionContext) HasAnchorSmfSupportedFeatures() bool`

HasAnchorSmfSupportedFeatures returns a boolean if a field has been set.

### GetAnchorSmfOauth2Required

`func (o *PduSessionContext) GetAnchorSmfOauth2Required() bool`

GetAnchorSmfOauth2Required returns the AnchorSmfOauth2Required field if non-nil, zero value otherwise.

### GetAnchorSmfOauth2RequiredOk

`func (o *PduSessionContext) GetAnchorSmfOauth2RequiredOk() (*bool, bool)`

GetAnchorSmfOauth2RequiredOk returns a tuple with the AnchorSmfOauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchorSmfOauth2Required

`func (o *PduSessionContext) SetAnchorSmfOauth2Required(v bool)`

SetAnchorSmfOauth2Required sets AnchorSmfOauth2Required field to given value.

### HasAnchorSmfOauth2Required

`func (o *PduSessionContext) HasAnchorSmfOauth2Required() bool`

HasAnchorSmfOauth2Required returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


