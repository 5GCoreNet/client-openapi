# MulticastSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValGroupId** | **string** |  | 
**AnncMode** | [**ServiceAnnoucementMode**](ServiceAnnoucementMode.md) |  | 
**MultiQosReq** | **string** |  | 
**LocArea** | Pointer to [**MbmsLocArea**](MbmsLocArea.md) |  | [optional] 
**Duration** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Tmgi** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**LocalMbmsInfo** | Pointer to [**LocalMbmsInfo**](LocalMbmsInfo.md) |  | [optional] 
**LocalMbmsActInd** | Pointer to **bool** |  | [optional] 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**ReqTestNotif** | Pointer to **bool** |  | [optional] 
**WsNotifCfg** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**UpIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**UpIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**UpPortNum** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535. | [optional] 
**RadioFreqs** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewMulticastSubscription

`func NewMulticastSubscription(valGroupId string, anncMode ServiceAnnoucementMode, multiQosReq string, notifUri string, ) *MulticastSubscription`

NewMulticastSubscription instantiates a new MulticastSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMulticastSubscriptionWithDefaults

`func NewMulticastSubscriptionWithDefaults() *MulticastSubscription`

NewMulticastSubscriptionWithDefaults instantiates a new MulticastSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValGroupId

`func (o *MulticastSubscription) GetValGroupId() string`

GetValGroupId returns the ValGroupId field if non-nil, zero value otherwise.

### GetValGroupIdOk

`func (o *MulticastSubscription) GetValGroupIdOk() (*string, bool)`

GetValGroupIdOk returns a tuple with the ValGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValGroupId

`func (o *MulticastSubscription) SetValGroupId(v string)`

SetValGroupId sets ValGroupId field to given value.


### GetAnncMode

`func (o *MulticastSubscription) GetAnncMode() ServiceAnnoucementMode`

GetAnncMode returns the AnncMode field if non-nil, zero value otherwise.

### GetAnncModeOk

`func (o *MulticastSubscription) GetAnncModeOk() (*ServiceAnnoucementMode, bool)`

GetAnncModeOk returns a tuple with the AnncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnncMode

`func (o *MulticastSubscription) SetAnncMode(v ServiceAnnoucementMode)`

SetAnncMode sets AnncMode field to given value.


### GetMultiQosReq

`func (o *MulticastSubscription) GetMultiQosReq() string`

GetMultiQosReq returns the MultiQosReq field if non-nil, zero value otherwise.

### GetMultiQosReqOk

`func (o *MulticastSubscription) GetMultiQosReqOk() (*string, bool)`

GetMultiQosReqOk returns a tuple with the MultiQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiQosReq

`func (o *MulticastSubscription) SetMultiQosReq(v string)`

SetMultiQosReq sets MultiQosReq field to given value.


### GetLocArea

`func (o *MulticastSubscription) GetLocArea() MbmsLocArea`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *MulticastSubscription) GetLocAreaOk() (*MbmsLocArea, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *MulticastSubscription) SetLocArea(v MbmsLocArea)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *MulticastSubscription) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetDuration

`func (o *MulticastSubscription) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MulticastSubscription) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MulticastSubscription) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MulticastSubscription) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTmgi

`func (o *MulticastSubscription) GetTmgi() int32`

GetTmgi returns the Tmgi field if non-nil, zero value otherwise.

### GetTmgiOk

`func (o *MulticastSubscription) GetTmgiOk() (*int32, bool)`

GetTmgiOk returns a tuple with the Tmgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgi

`func (o *MulticastSubscription) SetTmgi(v int32)`

SetTmgi sets Tmgi field to given value.

### HasTmgi

`func (o *MulticastSubscription) HasTmgi() bool`

HasTmgi returns a boolean if a field has been set.

### GetLocalMbmsInfo

`func (o *MulticastSubscription) GetLocalMbmsInfo() LocalMbmsInfo`

GetLocalMbmsInfo returns the LocalMbmsInfo field if non-nil, zero value otherwise.

### GetLocalMbmsInfoOk

`func (o *MulticastSubscription) GetLocalMbmsInfoOk() (*LocalMbmsInfo, bool)`

GetLocalMbmsInfoOk returns a tuple with the LocalMbmsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMbmsInfo

`func (o *MulticastSubscription) SetLocalMbmsInfo(v LocalMbmsInfo)`

SetLocalMbmsInfo sets LocalMbmsInfo field to given value.

### HasLocalMbmsInfo

`func (o *MulticastSubscription) HasLocalMbmsInfo() bool`

HasLocalMbmsInfo returns a boolean if a field has been set.

### GetLocalMbmsActInd

`func (o *MulticastSubscription) GetLocalMbmsActInd() bool`

GetLocalMbmsActInd returns the LocalMbmsActInd field if non-nil, zero value otherwise.

### GetLocalMbmsActIndOk

`func (o *MulticastSubscription) GetLocalMbmsActIndOk() (*bool, bool)`

GetLocalMbmsActIndOk returns a tuple with the LocalMbmsActInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMbmsActInd

`func (o *MulticastSubscription) SetLocalMbmsActInd(v bool)`

SetLocalMbmsActInd sets LocalMbmsActInd field to given value.

### HasLocalMbmsActInd

`func (o *MulticastSubscription) HasLocalMbmsActInd() bool`

HasLocalMbmsActInd returns a boolean if a field has been set.

### GetNotifUri

`func (o *MulticastSubscription) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *MulticastSubscription) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *MulticastSubscription) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetReqTestNotif

`func (o *MulticastSubscription) GetReqTestNotif() bool`

GetReqTestNotif returns the ReqTestNotif field if non-nil, zero value otherwise.

### GetReqTestNotifOk

`func (o *MulticastSubscription) GetReqTestNotifOk() (*bool, bool)`

GetReqTestNotifOk returns a tuple with the ReqTestNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTestNotif

`func (o *MulticastSubscription) SetReqTestNotif(v bool)`

SetReqTestNotif sets ReqTestNotif field to given value.

### HasReqTestNotif

`func (o *MulticastSubscription) HasReqTestNotif() bool`

HasReqTestNotif returns a boolean if a field has been set.

### GetWsNotifCfg

`func (o *MulticastSubscription) GetWsNotifCfg() WebsockNotifConfig`

GetWsNotifCfg returns the WsNotifCfg field if non-nil, zero value otherwise.

### GetWsNotifCfgOk

`func (o *MulticastSubscription) GetWsNotifCfgOk() (*WebsockNotifConfig, bool)`

GetWsNotifCfgOk returns a tuple with the WsNotifCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsNotifCfg

`func (o *MulticastSubscription) SetWsNotifCfg(v WebsockNotifConfig)`

SetWsNotifCfg sets WsNotifCfg field to given value.

### HasWsNotifCfg

`func (o *MulticastSubscription) HasWsNotifCfg() bool`

HasWsNotifCfg returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MulticastSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MulticastSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MulticastSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MulticastSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetUpIpv4Addr

`func (o *MulticastSubscription) GetUpIpv4Addr() string`

GetUpIpv4Addr returns the UpIpv4Addr field if non-nil, zero value otherwise.

### GetUpIpv4AddrOk

`func (o *MulticastSubscription) GetUpIpv4AddrOk() (*string, bool)`

GetUpIpv4AddrOk returns a tuple with the UpIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpIpv4Addr

`func (o *MulticastSubscription) SetUpIpv4Addr(v string)`

SetUpIpv4Addr sets UpIpv4Addr field to given value.

### HasUpIpv4Addr

`func (o *MulticastSubscription) HasUpIpv4Addr() bool`

HasUpIpv4Addr returns a boolean if a field has been set.

### GetUpIpv6Addr

`func (o *MulticastSubscription) GetUpIpv6Addr() Ipv6Addr`

GetUpIpv6Addr returns the UpIpv6Addr field if non-nil, zero value otherwise.

### GetUpIpv6AddrOk

`func (o *MulticastSubscription) GetUpIpv6AddrOk() (*Ipv6Addr, bool)`

GetUpIpv6AddrOk returns a tuple with the UpIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpIpv6Addr

`func (o *MulticastSubscription) SetUpIpv6Addr(v Ipv6Addr)`

SetUpIpv6Addr sets UpIpv6Addr field to given value.

### HasUpIpv6Addr

`func (o *MulticastSubscription) HasUpIpv6Addr() bool`

HasUpIpv6Addr returns a boolean if a field has been set.

### GetUpPortNum

`func (o *MulticastSubscription) GetUpPortNum() int32`

GetUpPortNum returns the UpPortNum field if non-nil, zero value otherwise.

### GetUpPortNumOk

`func (o *MulticastSubscription) GetUpPortNumOk() (*int32, bool)`

GetUpPortNumOk returns a tuple with the UpPortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpPortNum

`func (o *MulticastSubscription) SetUpPortNum(v int32)`

SetUpPortNum sets UpPortNum field to given value.

### HasUpPortNum

`func (o *MulticastSubscription) HasUpPortNum() bool`

HasUpPortNum returns a boolean if a field has been set.

### GetRadioFreqs

`func (o *MulticastSubscription) GetRadioFreqs() []int32`

GetRadioFreqs returns the RadioFreqs field if non-nil, zero value otherwise.

### GetRadioFreqsOk

`func (o *MulticastSubscription) GetRadioFreqsOk() (*[]int32, bool)`

GetRadioFreqsOk returns a tuple with the RadioFreqs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioFreqs

`func (o *MulticastSubscription) SetRadioFreqs(v []int32)`

SetRadioFreqs sets RadioFreqs field to given value.

### HasRadioFreqs

`func (o *MulticastSubscription) HasRadioFreqs() bool`

HasRadioFreqs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


