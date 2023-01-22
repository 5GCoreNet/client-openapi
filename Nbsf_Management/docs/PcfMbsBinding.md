# PcfMbsBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**PcfFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PcfIpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) |  | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**BindLevel** | Pointer to [**BindingLevel**](BindingLevel.md) |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewPcfMbsBinding

`func NewPcfMbsBinding(mbsSessionId MbsSessionId, ) *PcfMbsBinding`

NewPcfMbsBinding instantiates a new PcfMbsBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPcfMbsBindingWithDefaults

`func NewPcfMbsBindingWithDefaults() *PcfMbsBinding`

NewPcfMbsBindingWithDefaults instantiates a new PcfMbsBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *PcfMbsBinding) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *PcfMbsBinding) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *PcfMbsBinding) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetPcfFqdn

`func (o *PcfMbsBinding) GetPcfFqdn() string`

GetPcfFqdn returns the PcfFqdn field if non-nil, zero value otherwise.

### GetPcfFqdnOk

`func (o *PcfMbsBinding) GetPcfFqdnOk() (*string, bool)`

GetPcfFqdnOk returns a tuple with the PcfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfFqdn

`func (o *PcfMbsBinding) SetPcfFqdn(v string)`

SetPcfFqdn sets PcfFqdn field to given value.

### HasPcfFqdn

`func (o *PcfMbsBinding) HasPcfFqdn() bool`

HasPcfFqdn returns a boolean if a field has been set.

### GetPcfIpEndPoints

`func (o *PcfMbsBinding) GetPcfIpEndPoints() []IpEndPoint`

GetPcfIpEndPoints returns the PcfIpEndPoints field if non-nil, zero value otherwise.

### GetPcfIpEndPointsOk

`func (o *PcfMbsBinding) GetPcfIpEndPointsOk() (*[]IpEndPoint, bool)`

GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfIpEndPoints

`func (o *PcfMbsBinding) SetPcfIpEndPoints(v []IpEndPoint)`

SetPcfIpEndPoints sets PcfIpEndPoints field to given value.

### HasPcfIpEndPoints

`func (o *PcfMbsBinding) HasPcfIpEndPoints() bool`

HasPcfIpEndPoints returns a boolean if a field has been set.

### GetPcfId

`func (o *PcfMbsBinding) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *PcfMbsBinding) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *PcfMbsBinding) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *PcfMbsBinding) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *PcfMbsBinding) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *PcfMbsBinding) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *PcfMbsBinding) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *PcfMbsBinding) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetBindLevel

`func (o *PcfMbsBinding) GetBindLevel() BindingLevel`

GetBindLevel returns the BindLevel field if non-nil, zero value otherwise.

### GetBindLevelOk

`func (o *PcfMbsBinding) GetBindLevelOk() (*BindingLevel, bool)`

GetBindLevelOk returns a tuple with the BindLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindLevel

`func (o *PcfMbsBinding) SetBindLevel(v BindingLevel)`

SetBindLevel sets BindLevel field to given value.

### HasBindLevel

`func (o *PcfMbsBinding) HasBindLevel() bool`

HasBindLevel returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *PcfMbsBinding) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *PcfMbsBinding) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *PcfMbsBinding) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *PcfMbsBinding) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetSuppFeat

`func (o *PcfMbsBinding) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *PcfMbsBinding) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *PcfMbsBinding) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *PcfMbsBinding) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


