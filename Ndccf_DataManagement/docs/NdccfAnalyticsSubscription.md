# NdccfAnalyticsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnaSub** | [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | 
**AnaNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**AnaNotifCorrId** | **string** | Notification correlation identifier. | 
**FormatInstruct** | Pointer to [**FormattingInstruction**](FormattingInstruction.md) |  | [optional] 
**ProcInstructs** | Pointer to [**[]ProcessingInstruction**](ProcessingInstruction.md) | Processing instructions to be used for sending event notifications. | [optional] 
**TargetNfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TargetNfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**AdrfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**ArdfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**TimePeriod** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**DataCollectPurposes** | Pointer to [**[]DataCollectionPurpose**](DataCollectionPurpose.md) | The purposes of data collection. This attribute may only be provided if the consumer has checked user consent.  | [optional] 

## Methods

### NewNdccfAnalyticsSubscription

`func NewNdccfAnalyticsSubscription(anaSub NnwdafEventsSubscription, anaNotifUri string, anaNotifCorrId string, ) *NdccfAnalyticsSubscription`

NewNdccfAnalyticsSubscription instantiates a new NdccfAnalyticsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNdccfAnalyticsSubscriptionWithDefaults

`func NewNdccfAnalyticsSubscriptionWithDefaults() *NdccfAnalyticsSubscription`

NewNdccfAnalyticsSubscriptionWithDefaults instantiates a new NdccfAnalyticsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaSub

`func (o *NdccfAnalyticsSubscription) GetAnaSub() NnwdafEventsSubscription`

GetAnaSub returns the AnaSub field if non-nil, zero value otherwise.

### GetAnaSubOk

`func (o *NdccfAnalyticsSubscription) GetAnaSubOk() (*NnwdafEventsSubscription, bool)`

GetAnaSubOk returns a tuple with the AnaSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSub

`func (o *NdccfAnalyticsSubscription) SetAnaSub(v NnwdafEventsSubscription)`

SetAnaSub sets AnaSub field to given value.


### GetAnaNotifUri

`func (o *NdccfAnalyticsSubscription) GetAnaNotifUri() string`

GetAnaNotifUri returns the AnaNotifUri field if non-nil, zero value otherwise.

### GetAnaNotifUriOk

`func (o *NdccfAnalyticsSubscription) GetAnaNotifUriOk() (*string, bool)`

GetAnaNotifUriOk returns a tuple with the AnaNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifUri

`func (o *NdccfAnalyticsSubscription) SetAnaNotifUri(v string)`

SetAnaNotifUri sets AnaNotifUri field to given value.


### GetAnaNotifCorrId

`func (o *NdccfAnalyticsSubscription) GetAnaNotifCorrId() string`

GetAnaNotifCorrId returns the AnaNotifCorrId field if non-nil, zero value otherwise.

### GetAnaNotifCorrIdOk

`func (o *NdccfAnalyticsSubscription) GetAnaNotifCorrIdOk() (*string, bool)`

GetAnaNotifCorrIdOk returns a tuple with the AnaNotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaNotifCorrId

`func (o *NdccfAnalyticsSubscription) SetAnaNotifCorrId(v string)`

SetAnaNotifCorrId sets AnaNotifCorrId field to given value.


### GetFormatInstruct

`func (o *NdccfAnalyticsSubscription) GetFormatInstruct() FormattingInstruction`

GetFormatInstruct returns the FormatInstruct field if non-nil, zero value otherwise.

### GetFormatInstructOk

`func (o *NdccfAnalyticsSubscription) GetFormatInstructOk() (*FormattingInstruction, bool)`

GetFormatInstructOk returns a tuple with the FormatInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatInstruct

`func (o *NdccfAnalyticsSubscription) SetFormatInstruct(v FormattingInstruction)`

SetFormatInstruct sets FormatInstruct field to given value.

### HasFormatInstruct

`func (o *NdccfAnalyticsSubscription) HasFormatInstruct() bool`

HasFormatInstruct returns a boolean if a field has been set.

### GetProcInstructs

`func (o *NdccfAnalyticsSubscription) GetProcInstructs() []ProcessingInstruction`

GetProcInstructs returns the ProcInstructs field if non-nil, zero value otherwise.

### GetProcInstructsOk

`func (o *NdccfAnalyticsSubscription) GetProcInstructsOk() (*[]ProcessingInstruction, bool)`

GetProcInstructsOk returns a tuple with the ProcInstructs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcInstructs

`func (o *NdccfAnalyticsSubscription) SetProcInstructs(v []ProcessingInstruction)`

SetProcInstructs sets ProcInstructs field to given value.

### HasProcInstructs

`func (o *NdccfAnalyticsSubscription) HasProcInstructs() bool`

HasProcInstructs returns a boolean if a field has been set.

### GetTargetNfId

`func (o *NdccfAnalyticsSubscription) GetTargetNfId() string`

GetTargetNfId returns the TargetNfId field if non-nil, zero value otherwise.

### GetTargetNfIdOk

`func (o *NdccfAnalyticsSubscription) GetTargetNfIdOk() (*string, bool)`

GetTargetNfIdOk returns a tuple with the TargetNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfId

`func (o *NdccfAnalyticsSubscription) SetTargetNfId(v string)`

SetTargetNfId sets TargetNfId field to given value.

### HasTargetNfId

`func (o *NdccfAnalyticsSubscription) HasTargetNfId() bool`

HasTargetNfId returns a boolean if a field has been set.

### GetTargetNfSetId

`func (o *NdccfAnalyticsSubscription) GetTargetNfSetId() string`

GetTargetNfSetId returns the TargetNfSetId field if non-nil, zero value otherwise.

### GetTargetNfSetIdOk

`func (o *NdccfAnalyticsSubscription) GetTargetNfSetIdOk() (*string, bool)`

GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfSetId

`func (o *NdccfAnalyticsSubscription) SetTargetNfSetId(v string)`

SetTargetNfSetId sets TargetNfSetId field to given value.

### HasTargetNfSetId

`func (o *NdccfAnalyticsSubscription) HasTargetNfSetId() bool`

HasTargetNfSetId returns a boolean if a field has been set.

### GetAdrfId

`func (o *NdccfAnalyticsSubscription) GetAdrfId() string`

GetAdrfId returns the AdrfId field if non-nil, zero value otherwise.

### GetAdrfIdOk

`func (o *NdccfAnalyticsSubscription) GetAdrfIdOk() (*string, bool)`

GetAdrfIdOk returns a tuple with the AdrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfId

`func (o *NdccfAnalyticsSubscription) SetAdrfId(v string)`

SetAdrfId sets AdrfId field to given value.

### HasAdrfId

`func (o *NdccfAnalyticsSubscription) HasAdrfId() bool`

HasAdrfId returns a boolean if a field has been set.

### GetArdfSetId

`func (o *NdccfAnalyticsSubscription) GetArdfSetId() string`

GetArdfSetId returns the ArdfSetId field if non-nil, zero value otherwise.

### GetArdfSetIdOk

`func (o *NdccfAnalyticsSubscription) GetArdfSetIdOk() (*string, bool)`

GetArdfSetIdOk returns a tuple with the ArdfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArdfSetId

`func (o *NdccfAnalyticsSubscription) SetArdfSetId(v string)`

SetArdfSetId sets ArdfSetId field to given value.

### HasArdfSetId

`func (o *NdccfAnalyticsSubscription) HasArdfSetId() bool`

HasArdfSetId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *NdccfAnalyticsSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *NdccfAnalyticsSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *NdccfAnalyticsSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *NdccfAnalyticsSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetTimePeriod

`func (o *NdccfAnalyticsSubscription) GetTimePeriod() TimeWindow`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *NdccfAnalyticsSubscription) GetTimePeriodOk() (*TimeWindow, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *NdccfAnalyticsSubscription) SetTimePeriod(v TimeWindow)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *NdccfAnalyticsSubscription) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetDataCollectPurposes

`func (o *NdccfAnalyticsSubscription) GetDataCollectPurposes() []DataCollectionPurpose`

GetDataCollectPurposes returns the DataCollectPurposes field if non-nil, zero value otherwise.

### GetDataCollectPurposesOk

`func (o *NdccfAnalyticsSubscription) GetDataCollectPurposesOk() (*[]DataCollectionPurpose, bool)`

GetDataCollectPurposesOk returns a tuple with the DataCollectPurposes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectPurposes

`func (o *NdccfAnalyticsSubscription) SetDataCollectPurposes(v []DataCollectionPurpose)`

SetDataCollectPurposes sets DataCollectPurposes field to given value.

### HasDataCollectPurposes

`func (o *NdccfAnalyticsSubscription) HasDataCollectPurposes() bool`

HasDataCollectPurposes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


