# NnwdafDataManagementSubsc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdrfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AdrfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**AnaSub** | Pointer to [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | [optional] 
**DataCollectPurposes** | Pointer to [**[]DataCollectionPurpose**](DataCollectionPurpose.md) | The purposes of data collection. This attribute may only be provided if the consumer has checked user consent.  | [optional] 
**DataSub** | Pointer to [**DataSubscription**](DataSubscription.md) |  | [optional] 
**FormatInstruct** | Pointer to [**FormattingInstruction**](FormattingInstruction.md) |  | [optional] 
**NotifCorrId** | **string** |  | 
**NotificURI** | **string** | String providing an URI formatted according to RFC 3986. | 
**ProcInstruct** | Pointer to [**ProcessingInstruction**](ProcessingInstruction.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**TargetNfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TargetNfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**TimePeriod** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 

## Methods

### NewNnwdafDataManagementSubsc

`func NewNnwdafDataManagementSubsc(notifCorrId string, notificURI string, ) *NnwdafDataManagementSubsc`

NewNnwdafDataManagementSubsc instantiates a new NnwdafDataManagementSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNnwdafDataManagementSubscWithDefaults

`func NewNnwdafDataManagementSubscWithDefaults() *NnwdafDataManagementSubsc`

NewNnwdafDataManagementSubscWithDefaults instantiates a new NnwdafDataManagementSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdrfId

`func (o *NnwdafDataManagementSubsc) GetAdrfId() string`

GetAdrfId returns the AdrfId field if non-nil, zero value otherwise.

### GetAdrfIdOk

`func (o *NnwdafDataManagementSubsc) GetAdrfIdOk() (*string, bool)`

GetAdrfIdOk returns a tuple with the AdrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfId

`func (o *NnwdafDataManagementSubsc) SetAdrfId(v string)`

SetAdrfId sets AdrfId field to given value.

### HasAdrfId

`func (o *NnwdafDataManagementSubsc) HasAdrfId() bool`

HasAdrfId returns a boolean if a field has been set.

### GetAdrfSetId

`func (o *NnwdafDataManagementSubsc) GetAdrfSetId() string`

GetAdrfSetId returns the AdrfSetId field if non-nil, zero value otherwise.

### GetAdrfSetIdOk

`func (o *NnwdafDataManagementSubsc) GetAdrfSetIdOk() (*string, bool)`

GetAdrfSetIdOk returns a tuple with the AdrfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfSetId

`func (o *NnwdafDataManagementSubsc) SetAdrfSetId(v string)`

SetAdrfSetId sets AdrfSetId field to given value.

### HasAdrfSetId

`func (o *NnwdafDataManagementSubsc) HasAdrfSetId() bool`

HasAdrfSetId returns a boolean if a field has been set.

### GetAnaSub

`func (o *NnwdafDataManagementSubsc) GetAnaSub() NnwdafEventsSubscription`

GetAnaSub returns the AnaSub field if non-nil, zero value otherwise.

### GetAnaSubOk

`func (o *NnwdafDataManagementSubsc) GetAnaSubOk() (*NnwdafEventsSubscription, bool)`

GetAnaSubOk returns a tuple with the AnaSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSub

`func (o *NnwdafDataManagementSubsc) SetAnaSub(v NnwdafEventsSubscription)`

SetAnaSub sets AnaSub field to given value.

### HasAnaSub

`func (o *NnwdafDataManagementSubsc) HasAnaSub() bool`

HasAnaSub returns a boolean if a field has been set.

### GetDataCollectPurposes

`func (o *NnwdafDataManagementSubsc) GetDataCollectPurposes() []DataCollectionPurpose`

GetDataCollectPurposes returns the DataCollectPurposes field if non-nil, zero value otherwise.

### GetDataCollectPurposesOk

`func (o *NnwdafDataManagementSubsc) GetDataCollectPurposesOk() (*[]DataCollectionPurpose, bool)`

GetDataCollectPurposesOk returns a tuple with the DataCollectPurposes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectPurposes

`func (o *NnwdafDataManagementSubsc) SetDataCollectPurposes(v []DataCollectionPurpose)`

SetDataCollectPurposes sets DataCollectPurposes field to given value.

### HasDataCollectPurposes

`func (o *NnwdafDataManagementSubsc) HasDataCollectPurposes() bool`

HasDataCollectPurposes returns a boolean if a field has been set.

### GetDataSub

`func (o *NnwdafDataManagementSubsc) GetDataSub() DataSubscription`

GetDataSub returns the DataSub field if non-nil, zero value otherwise.

### GetDataSubOk

`func (o *NnwdafDataManagementSubsc) GetDataSubOk() (*DataSubscription, bool)`

GetDataSubOk returns a tuple with the DataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSub

`func (o *NnwdafDataManagementSubsc) SetDataSub(v DataSubscription)`

SetDataSub sets DataSub field to given value.

### HasDataSub

`func (o *NnwdafDataManagementSubsc) HasDataSub() bool`

HasDataSub returns a boolean if a field has been set.

### GetFormatInstruct

`func (o *NnwdafDataManagementSubsc) GetFormatInstruct() FormattingInstruction`

GetFormatInstruct returns the FormatInstruct field if non-nil, zero value otherwise.

### GetFormatInstructOk

`func (o *NnwdafDataManagementSubsc) GetFormatInstructOk() (*FormattingInstruction, bool)`

GetFormatInstructOk returns a tuple with the FormatInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatInstruct

`func (o *NnwdafDataManagementSubsc) SetFormatInstruct(v FormattingInstruction)`

SetFormatInstruct sets FormatInstruct field to given value.

### HasFormatInstruct

`func (o *NnwdafDataManagementSubsc) HasFormatInstruct() bool`

HasFormatInstruct returns a boolean if a field has been set.

### GetNotifCorrId

`func (o *NnwdafDataManagementSubsc) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *NnwdafDataManagementSubsc) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *NnwdafDataManagementSubsc) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.


### GetNotificURI

`func (o *NnwdafDataManagementSubsc) GetNotificURI() string`

GetNotificURI returns the NotificURI field if non-nil, zero value otherwise.

### GetNotificURIOk

`func (o *NnwdafDataManagementSubsc) GetNotificURIOk() (*string, bool)`

GetNotificURIOk returns a tuple with the NotificURI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificURI

`func (o *NnwdafDataManagementSubsc) SetNotificURI(v string)`

SetNotificURI sets NotificURI field to given value.


### GetProcInstruct

`func (o *NnwdafDataManagementSubsc) GetProcInstruct() ProcessingInstruction`

GetProcInstruct returns the ProcInstruct field if non-nil, zero value otherwise.

### GetProcInstructOk

`func (o *NnwdafDataManagementSubsc) GetProcInstructOk() (*ProcessingInstruction, bool)`

GetProcInstructOk returns a tuple with the ProcInstruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcInstruct

`func (o *NnwdafDataManagementSubsc) SetProcInstruct(v ProcessingInstruction)`

SetProcInstruct sets ProcInstruct field to given value.

### HasProcInstruct

`func (o *NnwdafDataManagementSubsc) HasProcInstruct() bool`

HasProcInstruct returns a boolean if a field has been set.

### GetSuppFeat

`func (o *NnwdafDataManagementSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *NnwdafDataManagementSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *NnwdafDataManagementSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *NnwdafDataManagementSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetTargetNfId

`func (o *NnwdafDataManagementSubsc) GetTargetNfId() string`

GetTargetNfId returns the TargetNfId field if non-nil, zero value otherwise.

### GetTargetNfIdOk

`func (o *NnwdafDataManagementSubsc) GetTargetNfIdOk() (*string, bool)`

GetTargetNfIdOk returns a tuple with the TargetNfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfId

`func (o *NnwdafDataManagementSubsc) SetTargetNfId(v string)`

SetTargetNfId sets TargetNfId field to given value.

### HasTargetNfId

`func (o *NnwdafDataManagementSubsc) HasTargetNfId() bool`

HasTargetNfId returns a boolean if a field has been set.

### GetTargetNfSetId

`func (o *NnwdafDataManagementSubsc) GetTargetNfSetId() string`

GetTargetNfSetId returns the TargetNfSetId field if non-nil, zero value otherwise.

### GetTargetNfSetIdOk

`func (o *NnwdafDataManagementSubsc) GetTargetNfSetIdOk() (*string, bool)`

GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNfSetId

`func (o *NnwdafDataManagementSubsc) SetTargetNfSetId(v string)`

SetTargetNfSetId sets TargetNfSetId field to given value.

### HasTargetNfSetId

`func (o *NnwdafDataManagementSubsc) HasTargetNfSetId() bool`

HasTargetNfSetId returns a boolean if a field has been set.

### GetTimePeriod

`func (o *NnwdafDataManagementSubsc) GetTimePeriod() TimeWindow`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *NnwdafDataManagementSubsc) GetTimePeriodOk() (*TimeWindow, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *NnwdafDataManagementSubsc) SetTimePeriod(v TimeWindow)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *NnwdafDataManagementSubsc) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


