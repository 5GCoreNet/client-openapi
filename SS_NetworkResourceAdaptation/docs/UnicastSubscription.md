# UnicastSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValTgtUe** | [**ValTargetUe**](ValTargetUe.md) |  | 
**UniQosReq** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**ReqTestNotif** | Pointer to **bool** |  | [optional] 
**WsNotifCfg** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewUnicastSubscription

`func NewUnicastSubscription(valTgtUe ValTargetUe, notifUri string, ) *UnicastSubscription`

NewUnicastSubscription instantiates a new UnicastSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnicastSubscriptionWithDefaults

`func NewUnicastSubscriptionWithDefaults() *UnicastSubscription`

NewUnicastSubscriptionWithDefaults instantiates a new UnicastSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValTgtUe

`func (o *UnicastSubscription) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *UnicastSubscription) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *UnicastSubscription) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.


### GetUniQosReq

`func (o *UnicastSubscription) GetUniQosReq() string`

GetUniQosReq returns the UniQosReq field if non-nil, zero value otherwise.

### GetUniQosReqOk

`func (o *UnicastSubscription) GetUniQosReqOk() (*string, bool)`

GetUniQosReqOk returns a tuple with the UniQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniQosReq

`func (o *UnicastSubscription) SetUniQosReq(v string)`

SetUniQosReq sets UniQosReq field to given value.

### HasUniQosReq

`func (o *UnicastSubscription) HasUniQosReq() bool`

HasUniQosReq returns a boolean if a field has been set.

### GetDuration

`func (o *UnicastSubscription) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *UnicastSubscription) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *UnicastSubscription) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *UnicastSubscription) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetNotifUri

`func (o *UnicastSubscription) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *UnicastSubscription) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *UnicastSubscription) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetReqTestNotif

`func (o *UnicastSubscription) GetReqTestNotif() bool`

GetReqTestNotif returns the ReqTestNotif field if non-nil, zero value otherwise.

### GetReqTestNotifOk

`func (o *UnicastSubscription) GetReqTestNotifOk() (*bool, bool)`

GetReqTestNotifOk returns a tuple with the ReqTestNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqTestNotif

`func (o *UnicastSubscription) SetReqTestNotif(v bool)`

SetReqTestNotif sets ReqTestNotif field to given value.

### HasReqTestNotif

`func (o *UnicastSubscription) HasReqTestNotif() bool`

HasReqTestNotif returns a boolean if a field has been set.

### GetWsNotifCfg

`func (o *UnicastSubscription) GetWsNotifCfg() WebsockNotifConfig`

GetWsNotifCfg returns the WsNotifCfg field if non-nil, zero value otherwise.

### GetWsNotifCfgOk

`func (o *UnicastSubscription) GetWsNotifCfgOk() (*WebsockNotifConfig, bool)`

GetWsNotifCfgOk returns a tuple with the WsNotifCfg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWsNotifCfg

`func (o *UnicastSubscription) SetWsNotifCfg(v WebsockNotifConfig)`

SetWsNotifCfg sets WsNotifCfg field to given value.

### HasWsNotifCfg

`func (o *UnicastSubscription) HasWsNotifCfg() bool`

HasWsNotifCfg returns a boolean if a field has been set.

### GetSuppFeat

`func (o *UnicastSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *UnicastSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *UnicastSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *UnicastSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


