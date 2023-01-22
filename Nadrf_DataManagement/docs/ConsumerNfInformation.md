# ConsumerNfInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 

## Methods

### NewConsumerNfInformation

`func NewConsumerNfInformation() *ConsumerNfInformation`

NewConsumerNfInformation instantiates a new ConsumerNfInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumerNfInformationWithDefaults

`func NewConsumerNfInformationWithDefaults() *ConsumerNfInformation`

NewConsumerNfInformationWithDefaults instantiates a new ConsumerNfInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfId

`func (o *ConsumerNfInformation) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *ConsumerNfInformation) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *ConsumerNfInformation) SetNfId(v string)`

SetNfId sets NfId field to given value.

### HasNfId

`func (o *ConsumerNfInformation) HasNfId() bool`

HasNfId returns a boolean if a field has been set.

### GetNfSetId

`func (o *ConsumerNfInformation) GetNfSetId() string`

GetNfSetId returns the NfSetId field if non-nil, zero value otherwise.

### GetNfSetIdOk

`func (o *ConsumerNfInformation) GetNfSetIdOk() (*string, bool)`

GetNfSetIdOk returns a tuple with the NfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetId

`func (o *ConsumerNfInformation) SetNfSetId(v string)`

SetNfSetId sets NfSetId field to given value.

### HasNfSetId

`func (o *ConsumerNfInformation) HasNfSetId() bool`

HasNfSetId returns a boolean if a field has been set.

### GetTaiList

`func (o *ConsumerNfInformation) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *ConsumerNfInformation) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *ConsumerNfInformation) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *ConsumerNfInformation) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


