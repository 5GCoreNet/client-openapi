# AnalyticsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NwdafId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NwdafSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**NwdafSubscriptionList** | [**[]NwdafSubscription**](NwdafSubscription.md) |  | 

## Methods

### NewAnalyticsSubscription

`func NewAnalyticsSubscription(nwdafSubscriptionList []NwdafSubscription, ) *AnalyticsSubscription`

NewAnalyticsSubscription instantiates a new AnalyticsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsSubscriptionWithDefaults

`func NewAnalyticsSubscriptionWithDefaults() *AnalyticsSubscription`

NewAnalyticsSubscriptionWithDefaults instantiates a new AnalyticsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNwdafId

`func (o *AnalyticsSubscription) GetNwdafId() string`

GetNwdafId returns the NwdafId field if non-nil, zero value otherwise.

### GetNwdafIdOk

`func (o *AnalyticsSubscription) GetNwdafIdOk() (*string, bool)`

GetNwdafIdOk returns a tuple with the NwdafId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafId

`func (o *AnalyticsSubscription) SetNwdafId(v string)`

SetNwdafId sets NwdafId field to given value.

### HasNwdafId

`func (o *AnalyticsSubscription) HasNwdafId() bool`

HasNwdafId returns a boolean if a field has been set.

### GetNwdafSetId

`func (o *AnalyticsSubscription) GetNwdafSetId() string`

GetNwdafSetId returns the NwdafSetId field if non-nil, zero value otherwise.

### GetNwdafSetIdOk

`func (o *AnalyticsSubscription) GetNwdafSetIdOk() (*string, bool)`

GetNwdafSetIdOk returns a tuple with the NwdafSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafSetId

`func (o *AnalyticsSubscription) SetNwdafSetId(v string)`

SetNwdafSetId sets NwdafSetId field to given value.

### HasNwdafSetId

`func (o *AnalyticsSubscription) HasNwdafSetId() bool`

HasNwdafSetId returns a boolean if a field has been set.

### GetNwdafSubscriptionList

`func (o *AnalyticsSubscription) GetNwdafSubscriptionList() []NwdafSubscription`

GetNwdafSubscriptionList returns the NwdafSubscriptionList field if non-nil, zero value otherwise.

### GetNwdafSubscriptionListOk

`func (o *AnalyticsSubscription) GetNwdafSubscriptionListOk() (*[]NwdafSubscription, bool)`

GetNwdafSubscriptionListOk returns a tuple with the NwdafSubscriptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafSubscriptionList

`func (o *AnalyticsSubscription) SetNwdafSubscriptionList(v []NwdafSubscription)`

SetNwdafSubscriptionList sets NwdafSubscriptionList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


