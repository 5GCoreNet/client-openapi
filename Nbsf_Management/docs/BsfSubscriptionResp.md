# BsfSubscriptionResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]BsfEvent**](BsfEvent.md) | Contain te subscribed events. | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifCorreId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**SnssaiDnnPairs** | Pointer to [**SnssaiDnnPair**](SnssaiDnnPair.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**BindLevel** | Pointer to [**BindingLevel**](BindingLevel.md) |  | [optional] 
**EventNotifs** | [**[]BsfEventNotification**](BsfEventNotification.md) | Notifications about Individual Events. | 

## Methods

### NewBsfSubscriptionResp

`func NewBsfSubscriptionResp(events []BsfEvent, notifUri string, notifCorreId string, supi string, eventNotifs []BsfEventNotification, ) *BsfSubscriptionResp`

NewBsfSubscriptionResp instantiates a new BsfSubscriptionResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsfSubscriptionRespWithDefaults

`func NewBsfSubscriptionRespWithDefaults() *BsfSubscriptionResp`

NewBsfSubscriptionRespWithDefaults instantiates a new BsfSubscriptionResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *BsfSubscriptionResp) GetEvents() []BsfEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *BsfSubscriptionResp) GetEventsOk() (*[]BsfEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *BsfSubscriptionResp) SetEvents(v []BsfEvent)`

SetEvents sets Events field to given value.


### GetNotifUri

`func (o *BsfSubscriptionResp) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *BsfSubscriptionResp) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *BsfSubscriptionResp) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetNotifCorreId

`func (o *BsfSubscriptionResp) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *BsfSubscriptionResp) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *BsfSubscriptionResp) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.


### GetSupi

`func (o *BsfSubscriptionResp) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *BsfSubscriptionResp) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *BsfSubscriptionResp) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetGpsi

`func (o *BsfSubscriptionResp) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *BsfSubscriptionResp) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *BsfSubscriptionResp) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *BsfSubscriptionResp) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSnssaiDnnPairs

`func (o *BsfSubscriptionResp) GetSnssaiDnnPairs() SnssaiDnnPair`

GetSnssaiDnnPairs returns the SnssaiDnnPairs field if non-nil, zero value otherwise.

### GetSnssaiDnnPairsOk

`func (o *BsfSubscriptionResp) GetSnssaiDnnPairsOk() (*SnssaiDnnPair, bool)`

GetSnssaiDnnPairsOk returns a tuple with the SnssaiDnnPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiDnnPairs

`func (o *BsfSubscriptionResp) SetSnssaiDnnPairs(v SnssaiDnnPair)`

SetSnssaiDnnPairs sets SnssaiDnnPairs field to given value.

### HasSnssaiDnnPairs

`func (o *BsfSubscriptionResp) HasSnssaiDnnPairs() bool`

HasSnssaiDnnPairs returns a boolean if a field has been set.

### GetSuppFeat

`func (o *BsfSubscriptionResp) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *BsfSubscriptionResp) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *BsfSubscriptionResp) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *BsfSubscriptionResp) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetPcfId

`func (o *BsfSubscriptionResp) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *BsfSubscriptionResp) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *BsfSubscriptionResp) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *BsfSubscriptionResp) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *BsfSubscriptionResp) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *BsfSubscriptionResp) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *BsfSubscriptionResp) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *BsfSubscriptionResp) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetBindLevel

`func (o *BsfSubscriptionResp) GetBindLevel() BindingLevel`

GetBindLevel returns the BindLevel field if non-nil, zero value otherwise.

### GetBindLevelOk

`func (o *BsfSubscriptionResp) GetBindLevelOk() (*BindingLevel, bool)`

GetBindLevelOk returns a tuple with the BindLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindLevel

`func (o *BsfSubscriptionResp) SetBindLevel(v BindingLevel)`

SetBindLevel sets BindLevel field to given value.

### HasBindLevel

`func (o *BsfSubscriptionResp) HasBindLevel() bool`

HasBindLevel returns a boolean if a field has been set.

### GetEventNotifs

`func (o *BsfSubscriptionResp) GetEventNotifs() []BsfEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *BsfSubscriptionResp) GetEventNotifsOk() (*[]BsfEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *BsfSubscriptionResp) SetEventNotifs(v []BsfEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


