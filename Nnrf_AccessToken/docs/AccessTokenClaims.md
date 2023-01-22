# AccessTokenClaims

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iss** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**Sub** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**Aud** | [**AccessTokenClaimsAud**](AccessTokenClaimsAud.md) |  | 
**Scope** | **string** |  | 
**Exp** | **int32** |  | 
**ConsumerPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ConsumerSnpnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**ProducerPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**ProducerSnpnId** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**ProducerSnssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**ProducerNsiList** | Pointer to **[]string** |  | [optional] 
**ProducerNfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**ProducerNfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 
**SourceNfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewAccessTokenClaims

`func NewAccessTokenClaims(iss string, sub string, aud AccessTokenClaimsAud, scope string, exp int32, ) *AccessTokenClaims`

NewAccessTokenClaims instantiates a new AccessTokenClaims object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenClaimsWithDefaults

`func NewAccessTokenClaimsWithDefaults() *AccessTokenClaims`

NewAccessTokenClaimsWithDefaults instantiates a new AccessTokenClaims object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIss

`func (o *AccessTokenClaims) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *AccessTokenClaims) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *AccessTokenClaims) SetIss(v string)`

SetIss sets Iss field to given value.


### GetSub

`func (o *AccessTokenClaims) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *AccessTokenClaims) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *AccessTokenClaims) SetSub(v string)`

SetSub sets Sub field to given value.


### GetAud

`func (o *AccessTokenClaims) GetAud() AccessTokenClaimsAud`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *AccessTokenClaims) GetAudOk() (*AccessTokenClaimsAud, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *AccessTokenClaims) SetAud(v AccessTokenClaimsAud)`

SetAud sets Aud field to given value.


### GetScope

`func (o *AccessTokenClaims) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenClaims) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenClaims) SetScope(v string)`

SetScope sets Scope field to given value.


### GetExp

`func (o *AccessTokenClaims) GetExp() int32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *AccessTokenClaims) GetExpOk() (*int32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *AccessTokenClaims) SetExp(v int32)`

SetExp sets Exp field to given value.


### GetConsumerPlmnId

`func (o *AccessTokenClaims) GetConsumerPlmnId() PlmnId`

GetConsumerPlmnId returns the ConsumerPlmnId field if non-nil, zero value otherwise.

### GetConsumerPlmnIdOk

`func (o *AccessTokenClaims) GetConsumerPlmnIdOk() (*PlmnId, bool)`

GetConsumerPlmnIdOk returns a tuple with the ConsumerPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerPlmnId

`func (o *AccessTokenClaims) SetConsumerPlmnId(v PlmnId)`

SetConsumerPlmnId sets ConsumerPlmnId field to given value.

### HasConsumerPlmnId

`func (o *AccessTokenClaims) HasConsumerPlmnId() bool`

HasConsumerPlmnId returns a boolean if a field has been set.

### GetConsumerSnpnId

`func (o *AccessTokenClaims) GetConsumerSnpnId() PlmnIdNid`

GetConsumerSnpnId returns the ConsumerSnpnId field if non-nil, zero value otherwise.

### GetConsumerSnpnIdOk

`func (o *AccessTokenClaims) GetConsumerSnpnIdOk() (*PlmnIdNid, bool)`

GetConsumerSnpnIdOk returns a tuple with the ConsumerSnpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerSnpnId

`func (o *AccessTokenClaims) SetConsumerSnpnId(v PlmnIdNid)`

SetConsumerSnpnId sets ConsumerSnpnId field to given value.

### HasConsumerSnpnId

`func (o *AccessTokenClaims) HasConsumerSnpnId() bool`

HasConsumerSnpnId returns a boolean if a field has been set.

### GetProducerPlmnId

`func (o *AccessTokenClaims) GetProducerPlmnId() PlmnId`

GetProducerPlmnId returns the ProducerPlmnId field if non-nil, zero value otherwise.

### GetProducerPlmnIdOk

`func (o *AccessTokenClaims) GetProducerPlmnIdOk() (*PlmnId, bool)`

GetProducerPlmnIdOk returns a tuple with the ProducerPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerPlmnId

`func (o *AccessTokenClaims) SetProducerPlmnId(v PlmnId)`

SetProducerPlmnId sets ProducerPlmnId field to given value.

### HasProducerPlmnId

`func (o *AccessTokenClaims) HasProducerPlmnId() bool`

HasProducerPlmnId returns a boolean if a field has been set.

### GetProducerSnpnId

`func (o *AccessTokenClaims) GetProducerSnpnId() PlmnIdNid`

GetProducerSnpnId returns the ProducerSnpnId field if non-nil, zero value otherwise.

### GetProducerSnpnIdOk

`func (o *AccessTokenClaims) GetProducerSnpnIdOk() (*PlmnIdNid, bool)`

GetProducerSnpnIdOk returns a tuple with the ProducerSnpnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerSnpnId

`func (o *AccessTokenClaims) SetProducerSnpnId(v PlmnIdNid)`

SetProducerSnpnId sets ProducerSnpnId field to given value.

### HasProducerSnpnId

`func (o *AccessTokenClaims) HasProducerSnpnId() bool`

HasProducerSnpnId returns a boolean if a field has been set.

### GetProducerSnssaiList

`func (o *AccessTokenClaims) GetProducerSnssaiList() []Snssai`

GetProducerSnssaiList returns the ProducerSnssaiList field if non-nil, zero value otherwise.

### GetProducerSnssaiListOk

`func (o *AccessTokenClaims) GetProducerSnssaiListOk() (*[]Snssai, bool)`

GetProducerSnssaiListOk returns a tuple with the ProducerSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerSnssaiList

`func (o *AccessTokenClaims) SetProducerSnssaiList(v []Snssai)`

SetProducerSnssaiList sets ProducerSnssaiList field to given value.

### HasProducerSnssaiList

`func (o *AccessTokenClaims) HasProducerSnssaiList() bool`

HasProducerSnssaiList returns a boolean if a field has been set.

### GetProducerNsiList

`func (o *AccessTokenClaims) GetProducerNsiList() []string`

GetProducerNsiList returns the ProducerNsiList field if non-nil, zero value otherwise.

### GetProducerNsiListOk

`func (o *AccessTokenClaims) GetProducerNsiListOk() (*[]string, bool)`

GetProducerNsiListOk returns a tuple with the ProducerNsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerNsiList

`func (o *AccessTokenClaims) SetProducerNsiList(v []string)`

SetProducerNsiList sets ProducerNsiList field to given value.

### HasProducerNsiList

`func (o *AccessTokenClaims) HasProducerNsiList() bool`

HasProducerNsiList returns a boolean if a field has been set.

### GetProducerNfSetId

`func (o *AccessTokenClaims) GetProducerNfSetId() string`

GetProducerNfSetId returns the ProducerNfSetId field if non-nil, zero value otherwise.

### GetProducerNfSetIdOk

`func (o *AccessTokenClaims) GetProducerNfSetIdOk() (*string, bool)`

GetProducerNfSetIdOk returns a tuple with the ProducerNfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerNfSetId

`func (o *AccessTokenClaims) SetProducerNfSetId(v string)`

SetProducerNfSetId sets ProducerNfSetId field to given value.

### HasProducerNfSetId

`func (o *AccessTokenClaims) HasProducerNfSetId() bool`

HasProducerNfSetId returns a boolean if a field has been set.

### GetProducerNfServiceSetId

`func (o *AccessTokenClaims) GetProducerNfServiceSetId() string`

GetProducerNfServiceSetId returns the ProducerNfServiceSetId field if non-nil, zero value otherwise.

### GetProducerNfServiceSetIdOk

`func (o *AccessTokenClaims) GetProducerNfServiceSetIdOk() (*string, bool)`

GetProducerNfServiceSetIdOk returns a tuple with the ProducerNfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerNfServiceSetId

`func (o *AccessTokenClaims) SetProducerNfServiceSetId(v string)`

SetProducerNfServiceSetId sets ProducerNfServiceSetId field to given value.

### HasProducerNfServiceSetId

`func (o *AccessTokenClaims) HasProducerNfServiceSetId() bool`

HasProducerNfServiceSetId returns a boolean if a field has been set.

### GetSourceNfInstanceId

`func (o *AccessTokenClaims) GetSourceNfInstanceId() string`

GetSourceNfInstanceId returns the SourceNfInstanceId field if non-nil, zero value otherwise.

### GetSourceNfInstanceIdOk

`func (o *AccessTokenClaims) GetSourceNfInstanceIdOk() (*string, bool)`

GetSourceNfInstanceIdOk returns a tuple with the SourceNfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNfInstanceId

`func (o *AccessTokenClaims) SetSourceNfInstanceId(v string)`

SetSourceNfInstanceId sets SourceNfInstanceId field to given value.

### HasSourceNfInstanceId

`func (o *AccessTokenClaims) HasSourceNfInstanceId() bool`

HasSourceNfInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


