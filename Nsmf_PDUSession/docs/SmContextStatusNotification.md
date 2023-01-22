# SmContextStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusInfo** | [**StatusInfo**](StatusInfo.md) |  | 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ApnRateStatus** | Pointer to [**ApnRateStatus**](ApnRateStatus.md) |  | [optional] 
**DdnFailureStatus** | Pointer to **bool** |  | [optional] [default to false]
**NotifyCorrelationIdsForddnFailure** | Pointer to **[]string** |  | [optional] 
**NewIntermediateSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NewSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NewSmfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**OldSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**OldSmContextRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AltAnchorSmfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AltAnchorSmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TargetDnaiInfo** | Pointer to [**TargetDnaiInfo**](TargetDnaiInfo.md) |  | [optional] 
**OldPduSessionRef** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**InterPlmnApiRoot** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewSmContextStatusNotification

`func NewSmContextStatusNotification(statusInfo StatusInfo, ) *SmContextStatusNotification`

NewSmContextStatusNotification instantiates a new SmContextStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextStatusNotificationWithDefaults

`func NewSmContextStatusNotificationWithDefaults() *SmContextStatusNotification`

NewSmContextStatusNotificationWithDefaults instantiates a new SmContextStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusInfo

`func (o *SmContextStatusNotification) GetStatusInfo() StatusInfo`

GetStatusInfo returns the StatusInfo field if non-nil, zero value otherwise.

### GetStatusInfoOk

`func (o *SmContextStatusNotification) GetStatusInfoOk() (*StatusInfo, bool)`

GetStatusInfoOk returns a tuple with the StatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusInfo

`func (o *SmContextStatusNotification) SetStatusInfo(v StatusInfo)`

SetStatusInfo sets StatusInfo field to given value.


### GetSmallDataRateStatus

`func (o *SmContextStatusNotification) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextStatusNotification) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextStatusNotification) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextStatusNotification) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetApnRateStatus

`func (o *SmContextStatusNotification) GetApnRateStatus() ApnRateStatus`

GetApnRateStatus returns the ApnRateStatus field if non-nil, zero value otherwise.

### GetApnRateStatusOk

`func (o *SmContextStatusNotification) GetApnRateStatusOk() (*ApnRateStatus, bool)`

GetApnRateStatusOk returns a tuple with the ApnRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnRateStatus

`func (o *SmContextStatusNotification) SetApnRateStatus(v ApnRateStatus)`

SetApnRateStatus sets ApnRateStatus field to given value.

### HasApnRateStatus

`func (o *SmContextStatusNotification) HasApnRateStatus() bool`

HasApnRateStatus returns a boolean if a field has been set.

### GetDdnFailureStatus

`func (o *SmContextStatusNotification) GetDdnFailureStatus() bool`

GetDdnFailureStatus returns the DdnFailureStatus field if non-nil, zero value otherwise.

### GetDdnFailureStatusOk

`func (o *SmContextStatusNotification) GetDdnFailureStatusOk() (*bool, bool)`

GetDdnFailureStatusOk returns a tuple with the DdnFailureStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnFailureStatus

`func (o *SmContextStatusNotification) SetDdnFailureStatus(v bool)`

SetDdnFailureStatus sets DdnFailureStatus field to given value.

### HasDdnFailureStatus

`func (o *SmContextStatusNotification) HasDdnFailureStatus() bool`

HasDdnFailureStatus returns a boolean if a field has been set.

### GetNotifyCorrelationIdsForddnFailure

`func (o *SmContextStatusNotification) GetNotifyCorrelationIdsForddnFailure() []string`

GetNotifyCorrelationIdsForddnFailure returns the NotifyCorrelationIdsForddnFailure field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdsForddnFailureOk

`func (o *SmContextStatusNotification) GetNotifyCorrelationIdsForddnFailureOk() (*[]string, bool)`

GetNotifyCorrelationIdsForddnFailureOk returns a tuple with the NotifyCorrelationIdsForddnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationIdsForddnFailure

`func (o *SmContextStatusNotification) SetNotifyCorrelationIdsForddnFailure(v []string)`

SetNotifyCorrelationIdsForddnFailure sets NotifyCorrelationIdsForddnFailure field to given value.

### HasNotifyCorrelationIdsForddnFailure

`func (o *SmContextStatusNotification) HasNotifyCorrelationIdsForddnFailure() bool`

HasNotifyCorrelationIdsForddnFailure returns a boolean if a field has been set.

### GetNewIntermediateSmfId

`func (o *SmContextStatusNotification) GetNewIntermediateSmfId() string`

GetNewIntermediateSmfId returns the NewIntermediateSmfId field if non-nil, zero value otherwise.

### GetNewIntermediateSmfIdOk

`func (o *SmContextStatusNotification) GetNewIntermediateSmfIdOk() (*string, bool)`

GetNewIntermediateSmfIdOk returns a tuple with the NewIntermediateSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIntermediateSmfId

`func (o *SmContextStatusNotification) SetNewIntermediateSmfId(v string)`

SetNewIntermediateSmfId sets NewIntermediateSmfId field to given value.

### HasNewIntermediateSmfId

`func (o *SmContextStatusNotification) HasNewIntermediateSmfId() bool`

HasNewIntermediateSmfId returns a boolean if a field has been set.

### GetNewSmfId

`func (o *SmContextStatusNotification) GetNewSmfId() string`

GetNewSmfId returns the NewSmfId field if non-nil, zero value otherwise.

### GetNewSmfIdOk

`func (o *SmContextStatusNotification) GetNewSmfIdOk() (*string, bool)`

GetNewSmfIdOk returns a tuple with the NewSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSmfId

`func (o *SmContextStatusNotification) SetNewSmfId(v string)`

SetNewSmfId sets NewSmfId field to given value.

### HasNewSmfId

`func (o *SmContextStatusNotification) HasNewSmfId() bool`

HasNewSmfId returns a boolean if a field has been set.

### GetNewSmfSetId

`func (o *SmContextStatusNotification) GetNewSmfSetId() string`

GetNewSmfSetId returns the NewSmfSetId field if non-nil, zero value otherwise.

### GetNewSmfSetIdOk

`func (o *SmContextStatusNotification) GetNewSmfSetIdOk() (*string, bool)`

GetNewSmfSetIdOk returns a tuple with the NewSmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSmfSetId

`func (o *SmContextStatusNotification) SetNewSmfSetId(v string)`

SetNewSmfSetId sets NewSmfSetId field to given value.

### HasNewSmfSetId

`func (o *SmContextStatusNotification) HasNewSmfSetId() bool`

HasNewSmfSetId returns a boolean if a field has been set.

### GetOldSmfId

`func (o *SmContextStatusNotification) GetOldSmfId() string`

GetOldSmfId returns the OldSmfId field if non-nil, zero value otherwise.

### GetOldSmfIdOk

`func (o *SmContextStatusNotification) GetOldSmfIdOk() (*string, bool)`

GetOldSmfIdOk returns a tuple with the OldSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmfId

`func (o *SmContextStatusNotification) SetOldSmfId(v string)`

SetOldSmfId sets OldSmfId field to given value.

### HasOldSmfId

`func (o *SmContextStatusNotification) HasOldSmfId() bool`

HasOldSmfId returns a boolean if a field has been set.

### GetOldSmContextRef

`func (o *SmContextStatusNotification) GetOldSmContextRef() string`

GetOldSmContextRef returns the OldSmContextRef field if non-nil, zero value otherwise.

### GetOldSmContextRefOk

`func (o *SmContextStatusNotification) GetOldSmContextRefOk() (*string, bool)`

GetOldSmContextRefOk returns a tuple with the OldSmContextRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSmContextRef

`func (o *SmContextStatusNotification) SetOldSmContextRef(v string)`

SetOldSmContextRef sets OldSmContextRef field to given value.

### HasOldSmContextRef

`func (o *SmContextStatusNotification) HasOldSmContextRef() bool`

HasOldSmContextRef returns a boolean if a field has been set.

### GetAltAnchorSmfUri

`func (o *SmContextStatusNotification) GetAltAnchorSmfUri() string`

GetAltAnchorSmfUri returns the AltAnchorSmfUri field if non-nil, zero value otherwise.

### GetAltAnchorSmfUriOk

`func (o *SmContextStatusNotification) GetAltAnchorSmfUriOk() (*string, bool)`

GetAltAnchorSmfUriOk returns a tuple with the AltAnchorSmfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAnchorSmfUri

`func (o *SmContextStatusNotification) SetAltAnchorSmfUri(v string)`

SetAltAnchorSmfUri sets AltAnchorSmfUri field to given value.

### HasAltAnchorSmfUri

`func (o *SmContextStatusNotification) HasAltAnchorSmfUri() bool`

HasAltAnchorSmfUri returns a boolean if a field has been set.

### GetAltAnchorSmfId

`func (o *SmContextStatusNotification) GetAltAnchorSmfId() string`

GetAltAnchorSmfId returns the AltAnchorSmfId field if non-nil, zero value otherwise.

### GetAltAnchorSmfIdOk

`func (o *SmContextStatusNotification) GetAltAnchorSmfIdOk() (*string, bool)`

GetAltAnchorSmfIdOk returns a tuple with the AltAnchorSmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltAnchorSmfId

`func (o *SmContextStatusNotification) SetAltAnchorSmfId(v string)`

SetAltAnchorSmfId sets AltAnchorSmfId field to given value.

### HasAltAnchorSmfId

`func (o *SmContextStatusNotification) HasAltAnchorSmfId() bool`

HasAltAnchorSmfId returns a boolean if a field has been set.

### GetTargetDnaiInfo

`func (o *SmContextStatusNotification) GetTargetDnaiInfo() TargetDnaiInfo`

GetTargetDnaiInfo returns the TargetDnaiInfo field if non-nil, zero value otherwise.

### GetTargetDnaiInfoOk

`func (o *SmContextStatusNotification) GetTargetDnaiInfoOk() (*TargetDnaiInfo, bool)`

GetTargetDnaiInfoOk returns a tuple with the TargetDnaiInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDnaiInfo

`func (o *SmContextStatusNotification) SetTargetDnaiInfo(v TargetDnaiInfo)`

SetTargetDnaiInfo sets TargetDnaiInfo field to given value.

### HasTargetDnaiInfo

`func (o *SmContextStatusNotification) HasTargetDnaiInfo() bool`

HasTargetDnaiInfo returns a boolean if a field has been set.

### GetOldPduSessionRef

`func (o *SmContextStatusNotification) GetOldPduSessionRef() string`

GetOldPduSessionRef returns the OldPduSessionRef field if non-nil, zero value otherwise.

### GetOldPduSessionRefOk

`func (o *SmContextStatusNotification) GetOldPduSessionRefOk() (*string, bool)`

GetOldPduSessionRefOk returns a tuple with the OldPduSessionRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPduSessionRef

`func (o *SmContextStatusNotification) SetOldPduSessionRef(v string)`

SetOldPduSessionRef sets OldPduSessionRef field to given value.

### HasOldPduSessionRef

`func (o *SmContextStatusNotification) HasOldPduSessionRef() bool`

HasOldPduSessionRef returns a boolean if a field has been set.

### GetInterPlmnApiRoot

`func (o *SmContextStatusNotification) GetInterPlmnApiRoot() string`

GetInterPlmnApiRoot returns the InterPlmnApiRoot field if non-nil, zero value otherwise.

### GetInterPlmnApiRootOk

`func (o *SmContextStatusNotification) GetInterPlmnApiRootOk() (*string, bool)`

GetInterPlmnApiRootOk returns a tuple with the InterPlmnApiRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnApiRoot

`func (o *SmContextStatusNotification) SetInterPlmnApiRoot(v string)`

SetInterPlmnApiRoot sets InterPlmnApiRoot field to given value.

### HasInterPlmnApiRoot

`func (o *SmContextStatusNotification) HasInterPlmnApiRoot() bool`

HasInterPlmnApiRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


