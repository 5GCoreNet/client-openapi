# MLModelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MlFileAddrs** | Pointer to [**[]MLModelAddr**](MLModelAddr.md) |  | [optional] 
**ModelProvId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**ModelProvSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 

## Methods

### NewMLModelInfo

`func NewMLModelInfo() *MLModelInfo`

NewMLModelInfo instantiates a new MLModelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLModelInfoWithDefaults

`func NewMLModelInfoWithDefaults() *MLModelInfo`

NewMLModelInfoWithDefaults instantiates a new MLModelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMlFileAddrs

`func (o *MLModelInfo) GetMlFileAddrs() []MLModelAddr`

GetMlFileAddrs returns the MlFileAddrs field if non-nil, zero value otherwise.

### GetMlFileAddrsOk

`func (o *MLModelInfo) GetMlFileAddrsOk() (*[]MLModelAddr, bool)`

GetMlFileAddrsOk returns a tuple with the MlFileAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlFileAddrs

`func (o *MLModelInfo) SetMlFileAddrs(v []MLModelAddr)`

SetMlFileAddrs sets MlFileAddrs field to given value.

### HasMlFileAddrs

`func (o *MLModelInfo) HasMlFileAddrs() bool`

HasMlFileAddrs returns a boolean if a field has been set.

### GetModelProvId

`func (o *MLModelInfo) GetModelProvId() string`

GetModelProvId returns the ModelProvId field if non-nil, zero value otherwise.

### GetModelProvIdOk

`func (o *MLModelInfo) GetModelProvIdOk() (*string, bool)`

GetModelProvIdOk returns a tuple with the ModelProvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvId

`func (o *MLModelInfo) SetModelProvId(v string)`

SetModelProvId sets ModelProvId field to given value.

### HasModelProvId

`func (o *MLModelInfo) HasModelProvId() bool`

HasModelProvId returns a boolean if a field has been set.

### GetModelProvSetId

`func (o *MLModelInfo) GetModelProvSetId() string`

GetModelProvSetId returns the ModelProvSetId field if non-nil, zero value otherwise.

### GetModelProvSetIdOk

`func (o *MLModelInfo) GetModelProvSetIdOk() (*string, bool)`

GetModelProvSetIdOk returns a tuple with the ModelProvSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvSetId

`func (o *MLModelInfo) SetModelProvSetId(v string)`

SetModelProvSetId sets ModelProvSetId field to given value.

### HasModelProvSetId

`func (o *MLModelInfo) HasModelProvSetId() bool`

HasModelProvSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


