# NadrfDataStoreSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnaSub** | Pointer to [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | [optional] 
**DataSub** | Pointer to [**DataSubscription**](DataSubscription.md) |  | [optional] 
**TargetNfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TargetNfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**FormatInstruct** | Pointer to [**FormattingInstruction**](FormattingInstruction.md) |  | [optional] 
**ProcInstruct** | Pointer to [**ProcessingInstruction**](ProcessingInstruction.md) |  | [optional] 

## Methods

### NewNadrfDataStoreSubscription

`func NewNadrfDataStoreSubscription() *NadrfDataStoreSubscription`

NewNadrfDataStoreSubscription instantiates a new NadrfDataStoreSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNadrfDataStoreSubscriptionWithDefaults

`func NewNadrfDataStoreSubscriptionWithDefaults() *NadrfDataStoreSubscription`

NewNadrfDataStoreSubscriptionWithDefaults instantiates a new NadrfDataStoreSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaSub

`func (o *NadrfDataStoreSubscription) GetAnaSub() NnwdafEventsSubscription`

GetAnaSub returns the AnaSub field if non-nil, zero value otherwise.

### GetAnaSubOk

`func (o *NadrfDataStoreSubscription) GetAnaSubOk() (*NnwdafEventsSubscription, bool)`

GetAnaSubOk returns a tuple with the AnaSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSub

`func (o *NadrfDataStoreSubscription) SetAnaSub(v NnwdafEventsSubscription)`

SetAnaSub sets AnaSub field to given value.

### HasAnaSub

`func (o *NadrfDataStoreSubscription) HasAnaSub() bool`

HasAnaSub returns a boolean if a field has been set.

### GetDataSub

`func (o *NadrfDataStoreSubscription) GetDataSub() DataSubscription`

GetDataSub returns the DataSub field if non-nil, zero value otherwise.

### GetDataSubOk

`func (o *NadrfDataStoreSubscription) GetDataSubOk() (*DataSubscription, bool)`

GetDataSubOk returns a tuple with the DataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSub

`func (o *NadrfDataStoreSubscription) SetDataSub(v DataSubscription)`

SetDataSub sets DataSub field to given value.

### HasDataSub

`func (o *NadrfDataStoreSubscription) HasDataSub() bool`

HasDataSub returns a boolean if a field has been set.

### GetTargetNfId

`func (o *NadrfDataStoreSubscription) GetTargetNfId() string`

GetTargetNfId returns the TargetNfId field if non-nil, zero value otherwise.

### GetTargetNfIdOk

`func (o *NadrfDataStoreSubscription) GetTargetNfIdOk() (*string, bool)`

GetTargetNfIdOk returns a tuple with the TargetNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfId

`func (o *NadrfDataStoreSubscription) SetTargetNfId(v string)`

SetTargetNfId sets TargetNfId field to given value.

### HasTargetNfId

`func (o *NadrfDataStoreSubscription) HasTargetNfId() bool`

HasTargetNfId returns a boolean if a field has been set.

### GetTargetNfSetId

`func (o *NadrfDataStoreSubscription) GetTargetNfSetId() string`

GetTargetNfSetId returns the TargetNfSetId field if non-nil, zero value otherwise.

### GetTargetNfSetIdOk

`func (o *NadrfDataStoreSubscription) GetTargetNfSetIdOk() (*string, bool)`

GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfSetId

`func (o *NadrfDataStoreSubscription) SetTargetNfSetId(v string)`

SetTargetNfSetId sets TargetNfSetId field to given value.

### HasTargetNfSetId

`func (o *NadrfDataStoreSubscription) HasTargetNfSetId() bool`

HasTargetNfSetId returns a boolean if a field has been set.

### GetFormatInstruct

`func (o *NadrfDataStoreSubscription) GetFormatInstruct() FormattingInstruction`

GetFormatInstruct returns the FormatInstruct field if non-nil, zero value otherwise.

### GetFormatInstructOk

`func (o *NadrfDataStoreSubscription) GetFormatInstructOk() (*FormattingInstruction, bool)`

GetFormatInstructOk returns a tuple with the FormatInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatInstruct

`func (o *NadrfDataStoreSubscription) SetFormatInstruct(v FormattingInstruction)`

SetFormatInstruct sets FormatInstruct field to given value.

### HasFormatInstruct

`func (o *NadrfDataStoreSubscription) HasFormatInstruct() bool`

HasFormatInstruct returns a boolean if a field has been set.

### GetProcInstruct

`func (o *NadrfDataStoreSubscription) GetProcInstruct() ProcessingInstruction`

GetProcInstruct returns the ProcInstruct field if non-nil, zero value otherwise.

### GetProcInstructOk

`func (o *NadrfDataStoreSubscription) GetProcInstructOk() (*ProcessingInstruction, bool)`

GetProcInstructOk returns a tuple with the ProcInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcInstruct

`func (o *NadrfDataStoreSubscription) SetProcInstruct(v ProcessingInstruction)`

SetProcInstruct sets ProcInstruct field to given value.

### HasProcInstruct

`func (o *NadrfDataStoreSubscription) HasProcInstruct() bool`

HasProcInstruct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


