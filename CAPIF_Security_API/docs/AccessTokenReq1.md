# AccessTokenReq1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** |  | 
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**TargetNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**Scope** | **string** |  | 
**TargetNfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**RequesterPlmn** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**RequesterPlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**RequesterSnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RequesterFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**RequesterSnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**TargetPlmn** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**TargetSnpn** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**TargetSnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**TargetNsiList** | Pointer to **[]string** |  | [optional] 
**TargetNfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**TargetNfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**HnrfAccessTokenUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SourceNfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewAccessTokenReq1

`func NewAccessTokenReq1(grantType string, nfInstanceId string, scope string, ) *AccessTokenReq1`

NewAccessTokenReq1 instantiates a new AccessTokenReq1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenReq1WithDefaults

`func NewAccessTokenReq1WithDefaults() *AccessTokenReq1`

NewAccessTokenReq1WithDefaults instantiates a new AccessTokenReq1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *AccessTokenReq1) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *AccessTokenReq1) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *AccessTokenReq1) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetNfInstanceId

`func (o *AccessTokenReq1) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *AccessTokenReq1) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *AccessTokenReq1) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfType

`func (o *AccessTokenReq1) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *AccessTokenReq1) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *AccessTokenReq1) SetNfType(v NFType)`

SetNfType sets NfType field to given value.

### HasNfType

`func (o *AccessTokenReq1) HasNfType() bool`

HasNfType returns a boolean if a field has been set.

### GetTargetNfType

`func (o *AccessTokenReq1) GetTargetNfType() NFType`

GetTargetNfType returns the TargetNfType field if non-nil, zero value otherwise.

### GetTargetNfTypeOk

`func (o *AccessTokenReq1) GetTargetNfTypeOk() (*NFType, bool)`

GetTargetNfTypeOk returns a tuple with the TargetNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfType

`func (o *AccessTokenReq1) SetTargetNfType(v NFType)`

SetTargetNfType sets TargetNfType field to given value.

### HasTargetNfType

`func (o *AccessTokenReq1) HasTargetNfType() bool`

HasTargetNfType returns a boolean if a field has been set.

### GetScope

`func (o *AccessTokenReq1) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenReq1) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenReq1) SetScope(v string)`

SetScope sets Scope field to given value.


### GetTargetNfInstanceId

`func (o *AccessTokenReq1) GetTargetNfInstanceId() string`

GetTargetNfInstanceId returns the TargetNfInstanceId field if non-nil, zero value otherwise.

### GetTargetNfInstanceIdOk

`func (o *AccessTokenReq1) GetTargetNfInstanceIdOk() (*string, bool)`

GetTargetNfInstanceIdOk returns a tuple with the TargetNfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfInstanceId

`func (o *AccessTokenReq1) SetTargetNfInstanceId(v string)`

SetTargetNfInstanceId sets TargetNfInstanceId field to given value.

### HasTargetNfInstanceId

`func (o *AccessTokenReq1) HasTargetNfInstanceId() bool`

HasTargetNfInstanceId returns a boolean if a field has been set.

### GetRequesterPlmn

`func (o *AccessTokenReq1) GetRequesterPlmn() PlmnId`

GetRequesterPlmn returns the RequesterPlmn field if non-nil, zero value otherwise.

### GetRequesterPlmnOk

`func (o *AccessTokenReq1) GetRequesterPlmnOk() (*PlmnId, bool)`

GetRequesterPlmnOk returns a tuple with the RequesterPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPlmn

`func (o *AccessTokenReq1) SetRequesterPlmn(v PlmnId)`

SetRequesterPlmn sets RequesterPlmn field to given value.

### HasRequesterPlmn

`func (o *AccessTokenReq1) HasRequesterPlmn() bool`

HasRequesterPlmn returns a boolean if a field has been set.

### GetRequesterPlmnList

`func (o *AccessTokenReq1) GetRequesterPlmnList() []PlmnId`

GetRequesterPlmnList returns the RequesterPlmnList field if non-nil, zero value otherwise.

### GetRequesterPlmnListOk

`func (o *AccessTokenReq1) GetRequesterPlmnListOk() (*[]PlmnId, bool)`

GetRequesterPlmnListOk returns a tuple with the RequesterPlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterPlmnList

`func (o *AccessTokenReq1) SetRequesterPlmnList(v []PlmnId)`

SetRequesterPlmnList sets RequesterPlmnList field to given value.

### HasRequesterPlmnList

`func (o *AccessTokenReq1) HasRequesterPlmnList() bool`

HasRequesterPlmnList returns a boolean if a field has been set.

### GetRequesterSnssaiList

`func (o *AccessTokenReq1) GetRequesterSnssaiList() []Snssai`

GetRequesterSnssaiList returns the RequesterSnssaiList field if non-nil, zero value otherwise.

### GetRequesterSnssaiListOk

`func (o *AccessTokenReq1) GetRequesterSnssaiListOk() (*[]Snssai, bool)`

GetRequesterSnssaiListOk returns a tuple with the RequesterSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSnssaiList

`func (o *AccessTokenReq1) SetRequesterSnssaiList(v []Snssai)`

SetRequesterSnssaiList sets RequesterSnssaiList field to given value.

### HasRequesterSnssaiList

`func (o *AccessTokenReq1) HasRequesterSnssaiList() bool`

HasRequesterSnssaiList returns a boolean if a field has been set.

### GetRequesterFqdn

`func (o *AccessTokenReq1) GetRequesterFqdn() string`

GetRequesterFqdn returns the RequesterFqdn field if non-nil, zero value otherwise.

### GetRequesterFqdnOk

`func (o *AccessTokenReq1) GetRequesterFqdnOk() (*string, bool)`

GetRequesterFqdnOk returns a tuple with the RequesterFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFqdn

`func (o *AccessTokenReq1) SetRequesterFqdn(v string)`

SetRequesterFqdn sets RequesterFqdn field to given value.

### HasRequesterFqdn

`func (o *AccessTokenReq1) HasRequesterFqdn() bool`

HasRequesterFqdn returns a boolean if a field has been set.

### GetRequesterSnpnList

`func (o *AccessTokenReq1) GetRequesterSnpnList() []PlmnIdNid`

GetRequesterSnpnList returns the RequesterSnpnList field if non-nil, zero value otherwise.

### GetRequesterSnpnListOk

`func (o *AccessTokenReq1) GetRequesterSnpnListOk() (*[]PlmnIdNid, bool)`

GetRequesterSnpnListOk returns a tuple with the RequesterSnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterSnpnList

`func (o *AccessTokenReq1) SetRequesterSnpnList(v []PlmnIdNid)`

SetRequesterSnpnList sets RequesterSnpnList field to given value.

### HasRequesterSnpnList

`func (o *AccessTokenReq1) HasRequesterSnpnList() bool`

HasRequesterSnpnList returns a boolean if a field has been set.

### GetTargetPlmn

`func (o *AccessTokenReq1) GetTargetPlmn() PlmnId`

GetTargetPlmn returns the TargetPlmn field if non-nil, zero value otherwise.

### GetTargetPlmnOk

`func (o *AccessTokenReq1) GetTargetPlmnOk() (*PlmnId, bool)`

GetTargetPlmnOk returns a tuple with the TargetPlmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlmn

`func (o *AccessTokenReq1) SetTargetPlmn(v PlmnId)`

SetTargetPlmn sets TargetPlmn field to given value.

### HasTargetPlmn

`func (o *AccessTokenReq1) HasTargetPlmn() bool`

HasTargetPlmn returns a boolean if a field has been set.

### GetTargetSnpn

`func (o *AccessTokenReq1) GetTargetSnpn() PlmnIdNid`

GetTargetSnpn returns the TargetSnpn field if non-nil, zero value otherwise.

### GetTargetSnpnOk

`func (o *AccessTokenReq1) GetTargetSnpnOk() (*PlmnIdNid, bool)`

GetTargetSnpnOk returns a tuple with the TargetSnpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSnpn

`func (o *AccessTokenReq1) SetTargetSnpn(v PlmnIdNid)`

SetTargetSnpn sets TargetSnpn field to given value.

### HasTargetSnpn

`func (o *AccessTokenReq1) HasTargetSnpn() bool`

HasTargetSnpn returns a boolean if a field has been set.

### GetTargetSnssaiList

`func (o *AccessTokenReq1) GetTargetSnssaiList() []Snssai`

GetTargetSnssaiList returns the TargetSnssaiList field if non-nil, zero value otherwise.

### GetTargetSnssaiListOk

`func (o *AccessTokenReq1) GetTargetSnssaiListOk() (*[]Snssai, bool)`

GetTargetSnssaiListOk returns a tuple with the TargetSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSnssaiList

`func (o *AccessTokenReq1) SetTargetSnssaiList(v []Snssai)`

SetTargetSnssaiList sets TargetSnssaiList field to given value.

### HasTargetSnssaiList

`func (o *AccessTokenReq1) HasTargetSnssaiList() bool`

HasTargetSnssaiList returns a boolean if a field has been set.

### GetTargetNsiList

`func (o *AccessTokenReq1) GetTargetNsiList() []string`

GetTargetNsiList returns the TargetNsiList field if non-nil, zero value otherwise.

### GetTargetNsiListOk

`func (o *AccessTokenReq1) GetTargetNsiListOk() (*[]string, bool)`

GetTargetNsiListOk returns a tuple with the TargetNsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNsiList

`func (o *AccessTokenReq1) SetTargetNsiList(v []string)`

SetTargetNsiList sets TargetNsiList field to given value.

### HasTargetNsiList

`func (o *AccessTokenReq1) HasTargetNsiList() bool`

HasTargetNsiList returns a boolean if a field has been set.

### GetTargetNfSetId

`func (o *AccessTokenReq1) GetTargetNfSetId() string`

GetTargetNfSetId returns the TargetNfSetId field if non-nil, zero value otherwise.

### GetTargetNfSetIdOk

`func (o *AccessTokenReq1) GetTargetNfSetIdOk() (*string, bool)`

GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfSetId

`func (o *AccessTokenReq1) SetTargetNfSetId(v string)`

SetTargetNfSetId sets TargetNfSetId field to given value.

### HasTargetNfSetId

`func (o *AccessTokenReq1) HasTargetNfSetId() bool`

HasTargetNfSetId returns a boolean if a field has been set.

### GetTargetNfServiceSetId

`func (o *AccessTokenReq1) GetTargetNfServiceSetId() string`

GetTargetNfServiceSetId returns the TargetNfServiceSetId field if non-nil, zero value otherwise.

### GetTargetNfServiceSetIdOk

`func (o *AccessTokenReq1) GetTargetNfServiceSetIdOk() (*string, bool)`

GetTargetNfServiceSetIdOk returns a tuple with the TargetNfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfServiceSetId

`func (o *AccessTokenReq1) SetTargetNfServiceSetId(v string)`

SetTargetNfServiceSetId sets TargetNfServiceSetId field to given value.

### HasTargetNfServiceSetId

`func (o *AccessTokenReq1) HasTargetNfServiceSetId() bool`

HasTargetNfServiceSetId returns a boolean if a field has been set.

### GetHnrfAccessTokenUri

`func (o *AccessTokenReq1) GetHnrfAccessTokenUri() string`

GetHnrfAccessTokenUri returns the HnrfAccessTokenUri field if non-nil, zero value otherwise.

### GetHnrfAccessTokenUriOk

`func (o *AccessTokenReq1) GetHnrfAccessTokenUriOk() (*string, bool)`

GetHnrfAccessTokenUriOk returns a tuple with the HnrfAccessTokenUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHnrfAccessTokenUri

`func (o *AccessTokenReq1) SetHnrfAccessTokenUri(v string)`

SetHnrfAccessTokenUri sets HnrfAccessTokenUri field to given value.

### HasHnrfAccessTokenUri

`func (o *AccessTokenReq1) HasHnrfAccessTokenUri() bool`

HasHnrfAccessTokenUri returns a boolean if a field has been set.

### GetSourceNfInstanceId

`func (o *AccessTokenReq1) GetSourceNfInstanceId() string`

GetSourceNfInstanceId returns the SourceNfInstanceId field if non-nil, zero value otherwise.

### GetSourceNfInstanceIdOk

`func (o *AccessTokenReq1) GetSourceNfInstanceIdOk() (*string, bool)`

GetSourceNfInstanceIdOk returns a tuple with the SourceNfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNfInstanceId

`func (o *AccessTokenReq1) SetSourceNfInstanceId(v string)`

SetSourceNfInstanceId sets SourceNfInstanceId field to given value.

### HasSourceNfInstanceId

`func (o *AccessTokenReq1) HasSourceNfInstanceId() bool`

HasSourceNfInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


