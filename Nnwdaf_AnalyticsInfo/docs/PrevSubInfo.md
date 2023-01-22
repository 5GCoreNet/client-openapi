# PrevSubInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProducerId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**ProducerSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SubscriptionId** | **string** | The identifier of a subscription. | 
**NfAnaEvents** | Pointer to [**[]NwdafEvent**](NwdafEvent.md) |  | [optional] 
**UeAnaEvents** | Pointer to [**[]UeAnalyticsContextDescriptor**](UeAnalyticsContextDescriptor.md) |  | [optional] 

## Methods

### NewPrevSubInfo

`func NewPrevSubInfo(subscriptionId string, ) *PrevSubInfo`

NewPrevSubInfo instantiates a new PrevSubInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrevSubInfoWithDefaults

`func NewPrevSubInfoWithDefaults() *PrevSubInfo`

NewPrevSubInfoWithDefaults instantiates a new PrevSubInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducerId

`func (o *PrevSubInfo) GetProducerId() string`

GetProducerId returns the ProducerId field if non-nil, zero value otherwise.

### GetProducerIdOk

`func (o *PrevSubInfo) GetProducerIdOk() (*string, bool)`

GetProducerIdOk returns a tuple with the ProducerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerId

`func (o *PrevSubInfo) SetProducerId(v string)`

SetProducerId sets ProducerId field to given value.

### HasProducerId

`func (o *PrevSubInfo) HasProducerId() bool`

HasProducerId returns a boolean if a field has been set.

### GetProducerSetId

`func (o *PrevSubInfo) GetProducerSetId() string`

GetProducerSetId returns the ProducerSetId field if non-nil, zero value otherwise.

### GetProducerSetIdOk

`func (o *PrevSubInfo) GetProducerSetIdOk() (*string, bool)`

GetProducerSetIdOk returns a tuple with the ProducerSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerSetId

`func (o *PrevSubInfo) SetProducerSetId(v string)`

SetProducerSetId sets ProducerSetId field to given value.

### HasProducerSetId

`func (o *PrevSubInfo) HasProducerSetId() bool`

HasProducerSetId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *PrevSubInfo) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *PrevSubInfo) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *PrevSubInfo) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetNfAnaEvents

`func (o *PrevSubInfo) GetNfAnaEvents() []NwdafEvent`

GetNfAnaEvents returns the NfAnaEvents field if non-nil, zero value otherwise.

### GetNfAnaEventsOk

`func (o *PrevSubInfo) GetNfAnaEventsOk() (*[]NwdafEvent, bool)`

GetNfAnaEventsOk returns a tuple with the NfAnaEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfAnaEvents

`func (o *PrevSubInfo) SetNfAnaEvents(v []NwdafEvent)`

SetNfAnaEvents sets NfAnaEvents field to given value.

### HasNfAnaEvents

`func (o *PrevSubInfo) HasNfAnaEvents() bool`

HasNfAnaEvents returns a boolean if a field has been set.

### GetUeAnaEvents

`func (o *PrevSubInfo) GetUeAnaEvents() []UeAnalyticsContextDescriptor`

GetUeAnaEvents returns the UeAnaEvents field if non-nil, zero value otherwise.

### GetUeAnaEventsOk

`func (o *PrevSubInfo) GetUeAnaEventsOk() (*[]UeAnalyticsContextDescriptor, bool)`

GetUeAnaEventsOk returns a tuple with the UeAnaEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeAnaEvents

`func (o *PrevSubInfo) SetUeAnaEvents(v []UeAnalyticsContextDescriptor)`

SetUeAnaEvents sets UeAnaEvents field to given value.

### HasUeAnaEvents

`func (o *PrevSubInfo) HasUeAnaEvents() bool`

HasUeAnaEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


