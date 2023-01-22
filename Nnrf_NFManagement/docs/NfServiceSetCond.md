# NfServiceSetCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfServiceSetId** | **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | 
**NfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 

## Methods

### NewNfServiceSetCond

`func NewNfServiceSetCond(nfServiceSetId string, ) *NfServiceSetCond`

NewNfServiceSetCond instantiates a new NfServiceSetCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfServiceSetCondWithDefaults

`func NewNfServiceSetCondWithDefaults() *NfServiceSetCond`

NewNfServiceSetCondWithDefaults instantiates a new NfServiceSetCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfServiceSetId

`func (o *NfServiceSetCond) GetNfServiceSetId() string`

GetNfServiceSetId returns the NfServiceSetId field if non-nil, zero value otherwise.

### GetNfServiceSetIdOk

`func (o *NfServiceSetCond) GetNfServiceSetIdOk() (*string, bool)`

GetNfServiceSetIdOk returns a tuple with the NfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceSetId

`func (o *NfServiceSetCond) SetNfServiceSetId(v string)`

SetNfServiceSetId sets NfServiceSetId field to given value.


### GetNfSetId

`func (o *NfServiceSetCond) GetNfSetId() string`

GetNfSetId returns the NfSetId field if non-nil, zero value otherwise.

### GetNfSetIdOk

`func (o *NfServiceSetCond) GetNfSetIdOk() (*string, bool)`

GetNfSetIdOk returns a tuple with the NfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetId

`func (o *NfServiceSetCond) SetNfSetId(v string)`

SetNfSetId sets NfSetId field to given value.

### HasNfSetId

`func (o *NfServiceSetCond) HasNfSetId() bool`

HasNfSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


