# BsfNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifCorreId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**PcfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**PcfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**BindLevel** | Pointer to [**BindingLevel**](BindingLevel.md) |  | [optional] 
**EventNotifs** | [**[]BsfEventNotification**](BsfEventNotification.md) | Notifications about Individual Events. | 

## Methods

### NewBsfNotification

`func NewBsfNotification(notifCorreId string, eventNotifs []BsfEventNotification, ) *BsfNotification`

NewBsfNotification instantiates a new BsfNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBsfNotificationWithDefaults

`func NewBsfNotificationWithDefaults() *BsfNotification`

NewBsfNotificationWithDefaults instantiates a new BsfNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifCorreId

`func (o *BsfNotification) GetNotifCorreId() string`

GetNotifCorreId returns the NotifCorreId field if non-nil, zero value otherwise.

### GetNotifCorreIdOk

`func (o *BsfNotification) GetNotifCorreIdOk() (*string, bool)`

GetNotifCorreIdOk returns a tuple with the NotifCorreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorreId

`func (o *BsfNotification) SetNotifCorreId(v string)`

SetNotifCorreId sets NotifCorreId field to given value.


### GetPcfId

`func (o *BsfNotification) GetPcfId() string`

GetPcfId returns the PcfId field if non-nil, zero value otherwise.

### GetPcfIdOk

`func (o *BsfNotification) GetPcfIdOk() (*string, bool)`

GetPcfIdOk returns a tuple with the PcfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfId

`func (o *BsfNotification) SetPcfId(v string)`

SetPcfId sets PcfId field to given value.

### HasPcfId

`func (o *BsfNotification) HasPcfId() bool`

HasPcfId returns a boolean if a field has been set.

### GetPcfSetId

`func (o *BsfNotification) GetPcfSetId() string`

GetPcfSetId returns the PcfSetId field if non-nil, zero value otherwise.

### GetPcfSetIdOk

`func (o *BsfNotification) GetPcfSetIdOk() (*string, bool)`

GetPcfSetIdOk returns a tuple with the PcfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfSetId

`func (o *BsfNotification) SetPcfSetId(v string)`

SetPcfSetId sets PcfSetId field to given value.

### HasPcfSetId

`func (o *BsfNotification) HasPcfSetId() bool`

HasPcfSetId returns a boolean if a field has been set.

### GetBindLevel

`func (o *BsfNotification) GetBindLevel() BindingLevel`

GetBindLevel returns the BindLevel field if non-nil, zero value otherwise.

### GetBindLevelOk

`func (o *BsfNotification) GetBindLevelOk() (*BindingLevel, bool)`

GetBindLevelOk returns a tuple with the BindLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindLevel

`func (o *BsfNotification) SetBindLevel(v BindingLevel)`

SetBindLevel sets BindLevel field to given value.

### HasBindLevel

`func (o *BsfNotification) HasBindLevel() bool`

HasBindLevel returns a boolean if a field has been set.

### GetEventNotifs

`func (o *BsfNotification) GetEventNotifs() []BsfEventNotification`

GetEventNotifs returns the EventNotifs field if non-nil, zero value otherwise.

### GetEventNotifsOk

`func (o *BsfNotification) GetEventNotifsOk() (*[]BsfEventNotification, bool)`

GetEventNotifsOk returns a tuple with the EventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifs

`func (o *BsfNotification) SetEventNotifs(v []BsfEventNotification)`

SetEventNotifs sets EventNotifs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


