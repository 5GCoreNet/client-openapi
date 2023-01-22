# ExtAmfEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventList** | [**[]AmfEvent**](AmfEvent.md) |  | 
**EventNotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifyCorrelationId** | **string** |  | 
**NfId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**SubsChangeNotifyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**SubsChangeNotifyCorrelationId** | Pointer to **string** |  | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**GroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**ExcludeSupiList** | Pointer to **[]string** |  | [optional] 
**ExcludeGpsiList** | Pointer to **[]string** |  | [optional] 
**IncludeSupiList** | Pointer to **[]string** |  | [optional] 
**IncludeGpsiList** | Pointer to **[]string** |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**AnyUE** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**AmfEventMode**](AmfEventMode.md) |  | [optional] 
**SourceNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**BindingInfo** | Pointer to **[]string** |  | [optional] 
**SubscribingNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**EventSyncInd** | Pointer to **bool** |  | [optional] 
**NfConsumerInfo** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExtAmfEventSubscription

`func NewExtAmfEventSubscription(eventList []AmfEvent, eventNotifyUri string, notifyCorrelationId string, nfId string, ) *ExtAmfEventSubscription`

NewExtAmfEventSubscription instantiates a new ExtAmfEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtAmfEventSubscriptionWithDefaults

`func NewExtAmfEventSubscriptionWithDefaults() *ExtAmfEventSubscription`

NewExtAmfEventSubscriptionWithDefaults instantiates a new ExtAmfEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventList

`func (o *ExtAmfEventSubscription) GetEventList() []AmfEvent`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *ExtAmfEventSubscription) GetEventListOk() (*[]AmfEvent, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *ExtAmfEventSubscription) SetEventList(v []AmfEvent)`

SetEventList sets EventList field to given value.


### GetEventNotifyUri

`func (o *ExtAmfEventSubscription) GetEventNotifyUri() string`

GetEventNotifyUri returns the EventNotifyUri field if non-nil, zero value otherwise.

### GetEventNotifyUriOk

`func (o *ExtAmfEventSubscription) GetEventNotifyUriOk() (*string, bool)`

GetEventNotifyUriOk returns a tuple with the EventNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyUri

`func (o *ExtAmfEventSubscription) SetEventNotifyUri(v string)`

SetEventNotifyUri sets EventNotifyUri field to given value.


### GetNotifyCorrelationId

`func (o *ExtAmfEventSubscription) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *ExtAmfEventSubscription) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *ExtAmfEventSubscription) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.


### GetNfId

`func (o *ExtAmfEventSubscription) GetNfId() string`

GetNfId returns the NfId field if non-nil, zero value otherwise.

### GetNfIdOk

`func (o *ExtAmfEventSubscription) GetNfIdOk() (*string, bool)`

GetNfIdOk returns a tuple with the NfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfId

`func (o *ExtAmfEventSubscription) SetNfId(v string)`

SetNfId sets NfId field to given value.


### GetSubsChangeNotifyUri

`func (o *ExtAmfEventSubscription) GetSubsChangeNotifyUri() string`

GetSubsChangeNotifyUri returns the SubsChangeNotifyUri field if non-nil, zero value otherwise.

### GetSubsChangeNotifyUriOk

`func (o *ExtAmfEventSubscription) GetSubsChangeNotifyUriOk() (*string, bool)`

GetSubsChangeNotifyUriOk returns a tuple with the SubsChangeNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsChangeNotifyUri

`func (o *ExtAmfEventSubscription) SetSubsChangeNotifyUri(v string)`

SetSubsChangeNotifyUri sets SubsChangeNotifyUri field to given value.

### HasSubsChangeNotifyUri

`func (o *ExtAmfEventSubscription) HasSubsChangeNotifyUri() bool`

HasSubsChangeNotifyUri returns a boolean if a field has been set.

### GetSubsChangeNotifyCorrelationId

`func (o *ExtAmfEventSubscription) GetSubsChangeNotifyCorrelationId() string`

GetSubsChangeNotifyCorrelationId returns the SubsChangeNotifyCorrelationId field if non-nil, zero value otherwise.

### GetSubsChangeNotifyCorrelationIdOk

`func (o *ExtAmfEventSubscription) GetSubsChangeNotifyCorrelationIdOk() (*string, bool)`

GetSubsChangeNotifyCorrelationIdOk returns a tuple with the SubsChangeNotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsChangeNotifyCorrelationId

`func (o *ExtAmfEventSubscription) SetSubsChangeNotifyCorrelationId(v string)`

SetSubsChangeNotifyCorrelationId sets SubsChangeNotifyCorrelationId field to given value.

### HasSubsChangeNotifyCorrelationId

`func (o *ExtAmfEventSubscription) HasSubsChangeNotifyCorrelationId() bool`

HasSubsChangeNotifyCorrelationId returns a boolean if a field has been set.

### GetSupi

`func (o *ExtAmfEventSubscription) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ExtAmfEventSubscription) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ExtAmfEventSubscription) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *ExtAmfEventSubscription) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetGroupId

`func (o *ExtAmfEventSubscription) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ExtAmfEventSubscription) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ExtAmfEventSubscription) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ExtAmfEventSubscription) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetExcludeSupiList

`func (o *ExtAmfEventSubscription) GetExcludeSupiList() []string`

GetExcludeSupiList returns the ExcludeSupiList field if non-nil, zero value otherwise.

### GetExcludeSupiListOk

`func (o *ExtAmfEventSubscription) GetExcludeSupiListOk() (*[]string, bool)`

GetExcludeSupiListOk returns a tuple with the ExcludeSupiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSupiList

`func (o *ExtAmfEventSubscription) SetExcludeSupiList(v []string)`

SetExcludeSupiList sets ExcludeSupiList field to given value.

### HasExcludeSupiList

`func (o *ExtAmfEventSubscription) HasExcludeSupiList() bool`

HasExcludeSupiList returns a boolean if a field has been set.

### GetExcludeGpsiList

`func (o *ExtAmfEventSubscription) GetExcludeGpsiList() []string`

GetExcludeGpsiList returns the ExcludeGpsiList field if non-nil, zero value otherwise.

### GetExcludeGpsiListOk

`func (o *ExtAmfEventSubscription) GetExcludeGpsiListOk() (*[]string, bool)`

GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGpsiList

`func (o *ExtAmfEventSubscription) SetExcludeGpsiList(v []string)`

SetExcludeGpsiList sets ExcludeGpsiList field to given value.

### HasExcludeGpsiList

`func (o *ExtAmfEventSubscription) HasExcludeGpsiList() bool`

HasExcludeGpsiList returns a boolean if a field has been set.

### GetIncludeSupiList

`func (o *ExtAmfEventSubscription) GetIncludeSupiList() []string`

GetIncludeSupiList returns the IncludeSupiList field if non-nil, zero value otherwise.

### GetIncludeSupiListOk

`func (o *ExtAmfEventSubscription) GetIncludeSupiListOk() (*[]string, bool)`

GetIncludeSupiListOk returns a tuple with the IncludeSupiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSupiList

`func (o *ExtAmfEventSubscription) SetIncludeSupiList(v []string)`

SetIncludeSupiList sets IncludeSupiList field to given value.

### HasIncludeSupiList

`func (o *ExtAmfEventSubscription) HasIncludeSupiList() bool`

HasIncludeSupiList returns a boolean if a field has been set.

### GetIncludeGpsiList

`func (o *ExtAmfEventSubscription) GetIncludeGpsiList() []string`

GetIncludeGpsiList returns the IncludeGpsiList field if non-nil, zero value otherwise.

### GetIncludeGpsiListOk

`func (o *ExtAmfEventSubscription) GetIncludeGpsiListOk() (*[]string, bool)`

GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGpsiList

`func (o *ExtAmfEventSubscription) SetIncludeGpsiList(v []string)`

SetIncludeGpsiList sets IncludeGpsiList field to given value.

### HasIncludeGpsiList

`func (o *ExtAmfEventSubscription) HasIncludeGpsiList() bool`

HasIncludeGpsiList returns a boolean if a field has been set.

### GetGpsi

`func (o *ExtAmfEventSubscription) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *ExtAmfEventSubscription) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *ExtAmfEventSubscription) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *ExtAmfEventSubscription) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPei

`func (o *ExtAmfEventSubscription) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *ExtAmfEventSubscription) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *ExtAmfEventSubscription) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *ExtAmfEventSubscription) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetAnyUE

`func (o *ExtAmfEventSubscription) GetAnyUE() bool`

GetAnyUE returns the AnyUE field if non-nil, zero value otherwise.

### GetAnyUEOk

`func (o *ExtAmfEventSubscription) GetAnyUEOk() (*bool, bool)`

GetAnyUEOk returns a tuple with the AnyUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUE

`func (o *ExtAmfEventSubscription) SetAnyUE(v bool)`

SetAnyUE sets AnyUE field to given value.

### HasAnyUE

`func (o *ExtAmfEventSubscription) HasAnyUE() bool`

HasAnyUE returns a boolean if a field has been set.

### GetOptions

`func (o *ExtAmfEventSubscription) GetOptions() AmfEventMode`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ExtAmfEventSubscription) GetOptionsOk() (*AmfEventMode, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ExtAmfEventSubscription) SetOptions(v AmfEventMode)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ExtAmfEventSubscription) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetSourceNfType

`func (o *ExtAmfEventSubscription) GetSourceNfType() NFType`

GetSourceNfType returns the SourceNfType field if non-nil, zero value otherwise.

### GetSourceNfTypeOk

`func (o *ExtAmfEventSubscription) GetSourceNfTypeOk() (*NFType, bool)`

GetSourceNfTypeOk returns a tuple with the SourceNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNfType

`func (o *ExtAmfEventSubscription) SetSourceNfType(v NFType)`

SetSourceNfType sets SourceNfType field to given value.

### HasSourceNfType

`func (o *ExtAmfEventSubscription) HasSourceNfType() bool`

HasSourceNfType returns a boolean if a field has been set.

### GetBindingInfo

`func (o *ExtAmfEventSubscription) GetBindingInfo() []string`

GetBindingInfo returns the BindingInfo field if non-nil, zero value otherwise.

### GetBindingInfoOk

`func (o *ExtAmfEventSubscription) GetBindingInfoOk() (*[]string, bool)`

GetBindingInfoOk returns a tuple with the BindingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingInfo

`func (o *ExtAmfEventSubscription) SetBindingInfo(v []string)`

SetBindingInfo sets BindingInfo field to given value.

### HasBindingInfo

`func (o *ExtAmfEventSubscription) HasBindingInfo() bool`

HasBindingInfo returns a boolean if a field has been set.

### GetSubscribingNfType

`func (o *ExtAmfEventSubscription) GetSubscribingNfType() NFType`

GetSubscribingNfType returns the SubscribingNfType field if non-nil, zero value otherwise.

### GetSubscribingNfTypeOk

`func (o *ExtAmfEventSubscription) GetSubscribingNfTypeOk() (*NFType, bool)`

GetSubscribingNfTypeOk returns a tuple with the SubscribingNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribingNfType

`func (o *ExtAmfEventSubscription) SetSubscribingNfType(v NFType)`

SetSubscribingNfType sets SubscribingNfType field to given value.

### HasSubscribingNfType

`func (o *ExtAmfEventSubscription) HasSubscribingNfType() bool`

HasSubscribingNfType returns a boolean if a field has been set.

### GetEventSyncInd

`func (o *ExtAmfEventSubscription) GetEventSyncInd() bool`

GetEventSyncInd returns the EventSyncInd field if non-nil, zero value otherwise.

### GetEventSyncIndOk

`func (o *ExtAmfEventSubscription) GetEventSyncIndOk() (*bool, bool)`

GetEventSyncIndOk returns a tuple with the EventSyncInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSyncInd

`func (o *ExtAmfEventSubscription) SetEventSyncInd(v bool)`

SetEventSyncInd sets EventSyncInd field to given value.

### HasEventSyncInd

`func (o *ExtAmfEventSubscription) HasEventSyncInd() bool`

HasEventSyncInd returns a boolean if a field has been set.

### GetNfConsumerInfo

`func (o *ExtAmfEventSubscription) GetNfConsumerInfo() []string`

GetNfConsumerInfo returns the NfConsumerInfo field if non-nil, zero value otherwise.

### GetNfConsumerInfoOk

`func (o *ExtAmfEventSubscription) GetNfConsumerInfoOk() (*[]string, bool)`

GetNfConsumerInfoOk returns a tuple with the NfConsumerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfConsumerInfo

`func (o *ExtAmfEventSubscription) SetNfConsumerInfo(v []string)`

SetNfConsumerInfo sets NfConsumerInfo field to given value.

### HasNfConsumerInfo

`func (o *ExtAmfEventSubscription) HasNfConsumerInfo() bool`

HasNfConsumerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


