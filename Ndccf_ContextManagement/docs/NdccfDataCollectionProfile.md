# NdccfDataCollectionProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnaSub** | Pointer to [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | [optional] 
**DataSub** | Pointer to [**DataNotification**](DataNotification.md) |  | [optional] 
**NwdafId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NwdafSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**AdrfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**AdrfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 

## Methods

### NewNdccfDataCollectionProfile

`func NewNdccfDataCollectionProfile() *NdccfDataCollectionProfile`

NewNdccfDataCollectionProfile instantiates a new NdccfDataCollectionProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNdccfDataCollectionProfileWithDefaults

`func NewNdccfDataCollectionProfileWithDefaults() *NdccfDataCollectionProfile`

NewNdccfDataCollectionProfileWithDefaults instantiates a new NdccfDataCollectionProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaSub

`func (o *NdccfDataCollectionProfile) GetAnaSub() NnwdafEventsSubscription`

GetAnaSub returns the AnaSub field if non-nil, zero value otherwise.

### GetAnaSubOk

`func (o *NdccfDataCollectionProfile) GetAnaSubOk() (*NnwdafEventsSubscription, bool)`

GetAnaSubOk returns a tuple with the AnaSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSub

`func (o *NdccfDataCollectionProfile) SetAnaSub(v NnwdafEventsSubscription)`

SetAnaSub sets AnaSub field to given value.

### HasAnaSub

`func (o *NdccfDataCollectionProfile) HasAnaSub() bool`

HasAnaSub returns a boolean if a field has been set.

### GetDataSub

`func (o *NdccfDataCollectionProfile) GetDataSub() DataNotification`

GetDataSub returns the DataSub field if non-nil, zero value otherwise.

### GetDataSubOk

`func (o *NdccfDataCollectionProfile) GetDataSubOk() (*DataNotification, bool)`

GetDataSubOk returns a tuple with the DataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSub

`func (o *NdccfDataCollectionProfile) SetDataSub(v DataNotification)`

SetDataSub sets DataSub field to given value.

### HasDataSub

`func (o *NdccfDataCollectionProfile) HasDataSub() bool`

HasDataSub returns a boolean if a field has been set.

### GetNwdafId

`func (o *NdccfDataCollectionProfile) GetNwdafId() string`

GetNwdafId returns the NwdafId field if non-nil, zero value otherwise.

### GetNwdafIdOk

`func (o *NdccfDataCollectionProfile) GetNwdafIdOk() (*string, bool)`

GetNwdafIdOk returns a tuple with the NwdafId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafId

`func (o *NdccfDataCollectionProfile) SetNwdafId(v string)`

SetNwdafId sets NwdafId field to given value.

### HasNwdafId

`func (o *NdccfDataCollectionProfile) HasNwdafId() bool`

HasNwdafId returns a boolean if a field has been set.

### GetNwdafSetId

`func (o *NdccfDataCollectionProfile) GetNwdafSetId() string`

GetNwdafSetId returns the NwdafSetId field if non-nil, zero value otherwise.

### GetNwdafSetIdOk

`func (o *NdccfDataCollectionProfile) GetNwdafSetIdOk() (*string, bool)`

GetNwdafSetIdOk returns a tuple with the NwdafSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafSetId

`func (o *NdccfDataCollectionProfile) SetNwdafSetId(v string)`

SetNwdafSetId sets NwdafSetId field to given value.

### HasNwdafSetId

`func (o *NdccfDataCollectionProfile) HasNwdafSetId() bool`

HasNwdafSetId returns a boolean if a field has been set.

### GetAdrfId

`func (o *NdccfDataCollectionProfile) GetAdrfId() string`

GetAdrfId returns the AdrfId field if non-nil, zero value otherwise.

### GetAdrfIdOk

`func (o *NdccfDataCollectionProfile) GetAdrfIdOk() (*string, bool)`

GetAdrfIdOk returns a tuple with the AdrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfId

`func (o *NdccfDataCollectionProfile) SetAdrfId(v string)`

SetAdrfId sets AdrfId field to given value.

### HasAdrfId

`func (o *NdccfDataCollectionProfile) HasAdrfId() bool`

HasAdrfId returns a boolean if a field has been set.

### GetAdrfSetId

`func (o *NdccfDataCollectionProfile) GetAdrfSetId() string`

GetAdrfSetId returns the AdrfSetId field if non-nil, zero value otherwise.

### GetAdrfSetIdOk

`func (o *NdccfDataCollectionProfile) GetAdrfSetIdOk() (*string, bool)`

GetAdrfSetIdOk returns a tuple with the AdrfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdrfSetId

`func (o *NdccfDataCollectionProfile) SetAdrfSetId(v string)`

SetAdrfSetId sets AdrfSetId field to given value.

### HasAdrfSetId

`func (o *NdccfDataCollectionProfile) HasAdrfSetId() bool`

HasAdrfSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


