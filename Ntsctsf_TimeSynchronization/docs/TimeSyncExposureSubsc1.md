# TimeSyncExposureSubsc1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExterGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**Gpsis** | Pointer to **[]string** | Contains a list of UE for which the time synchronization capabilities is requested.  | [optional] 
**AnyUeInd** | Pointer to **bool** | Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \&quot;false\&quot; is used, if not present.  | [optional] 
**AfServiceId** | Pointer to **string** | Identifies a service on behalf of which the AF is issuing the request. | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**SubsNotifId** | **string** | Notification Correlation ID assigned by the NF service consumer. | 
**SubsNotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**SubscribedEvents** | Pointer to [**[]SubscribedEvent**](SubscribedEvent.md) | Subscribed events | [optional] 
**EventFilters** | Pointer to [**[]EventFilter**](EventFilter.md) | Contains the filter conditions to match for notifying the event(s) of time synchronization capabilities for a list of UE(s).  | [optional] 
**NotifMethod** | Pointer to [**NotificationMethod**](NotificationMethod.md) |  | [optional] 
**MaxReportNbr** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3 of 3GPP TS 29.122. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewTimeSyncExposureSubsc1

`func NewTimeSyncExposureSubsc1(subsNotifId string, subsNotifUri string, ) *TimeSyncExposureSubsc1`

NewTimeSyncExposureSubsc1 instantiates a new TimeSyncExposureSubsc1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncExposureSubsc1WithDefaults

`func NewTimeSyncExposureSubsc1WithDefaults() *TimeSyncExposureSubsc1`

NewTimeSyncExposureSubsc1WithDefaults instantiates a new TimeSyncExposureSubsc1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExterGroupId

`func (o *TimeSyncExposureSubsc1) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *TimeSyncExposureSubsc1) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *TimeSyncExposureSubsc1) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.

### HasExterGroupId

`func (o *TimeSyncExposureSubsc1) HasExterGroupId() bool`

HasExterGroupId returns a boolean if a field has been set.

### GetGpsis

`func (o *TimeSyncExposureSubsc1) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *TimeSyncExposureSubsc1) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *TimeSyncExposureSubsc1) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *TimeSyncExposureSubsc1) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *TimeSyncExposureSubsc1) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *TimeSyncExposureSubsc1) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *TimeSyncExposureSubsc1) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *TimeSyncExposureSubsc1) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetAfServiceId

`func (o *TimeSyncExposureSubsc1) GetAfServiceId() string`

GetAfServiceId returns the AfServiceId field if non-nil, zero value otherwise.

### GetAfServiceIdOk

`func (o *TimeSyncExposureSubsc1) GetAfServiceIdOk() (*string, bool)`

GetAfServiceIdOk returns a tuple with the AfServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfServiceId

`func (o *TimeSyncExposureSubsc1) SetAfServiceId(v string)`

SetAfServiceId sets AfServiceId field to given value.

### HasAfServiceId

`func (o *TimeSyncExposureSubsc1) HasAfServiceId() bool`

HasAfServiceId returns a boolean if a field has been set.

### GetDnn

`func (o *TimeSyncExposureSubsc1) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TimeSyncExposureSubsc1) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TimeSyncExposureSubsc1) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TimeSyncExposureSubsc1) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *TimeSyncExposureSubsc1) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TimeSyncExposureSubsc1) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TimeSyncExposureSubsc1) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TimeSyncExposureSubsc1) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetSubsNotifId

`func (o *TimeSyncExposureSubsc1) GetSubsNotifId() string`

GetSubsNotifId returns the SubsNotifId field if non-nil, zero value otherwise.

### GetSubsNotifIdOk

`func (o *TimeSyncExposureSubsc1) GetSubsNotifIdOk() (*string, bool)`

GetSubsNotifIdOk returns a tuple with the SubsNotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifId

`func (o *TimeSyncExposureSubsc1) SetSubsNotifId(v string)`

SetSubsNotifId sets SubsNotifId field to given value.


### GetSubsNotifUri

`func (o *TimeSyncExposureSubsc1) GetSubsNotifUri() string`

GetSubsNotifUri returns the SubsNotifUri field if non-nil, zero value otherwise.

### GetSubsNotifUriOk

`func (o *TimeSyncExposureSubsc1) GetSubsNotifUriOk() (*string, bool)`

GetSubsNotifUriOk returns a tuple with the SubsNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifUri

`func (o *TimeSyncExposureSubsc1) SetSubsNotifUri(v string)`

SetSubsNotifUri sets SubsNotifUri field to given value.


### GetSubscribedEvents

`func (o *TimeSyncExposureSubsc1) GetSubscribedEvents() []SubscribedEvent`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *TimeSyncExposureSubsc1) GetSubscribedEventsOk() (*[]SubscribedEvent, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *TimeSyncExposureSubsc1) SetSubscribedEvents(v []SubscribedEvent)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *TimeSyncExposureSubsc1) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetEventFilters

`func (o *TimeSyncExposureSubsc1) GetEventFilters() []EventFilter`

GetEventFilters returns the EventFilters field if non-nil, zero value otherwise.

### GetEventFiltersOk

`func (o *TimeSyncExposureSubsc1) GetEventFiltersOk() (*[]EventFilter, bool)`

GetEventFiltersOk returns a tuple with the EventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilters

`func (o *TimeSyncExposureSubsc1) SetEventFilters(v []EventFilter)`

SetEventFilters sets EventFilters field to given value.

### HasEventFilters

`func (o *TimeSyncExposureSubsc1) HasEventFilters() bool`

HasEventFilters returns a boolean if a field has been set.

### GetNotifMethod

`func (o *TimeSyncExposureSubsc1) GetNotifMethod() NotificationMethod`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *TimeSyncExposureSubsc1) GetNotifMethodOk() (*NotificationMethod, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *TimeSyncExposureSubsc1) SetNotifMethod(v NotificationMethod)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *TimeSyncExposureSubsc1) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetMaxReportNbr

`func (o *TimeSyncExposureSubsc1) GetMaxReportNbr() int32`

GetMaxReportNbr returns the MaxReportNbr field if non-nil, zero value otherwise.

### GetMaxReportNbrOk

`func (o *TimeSyncExposureSubsc1) GetMaxReportNbrOk() (*int32, bool)`

GetMaxReportNbrOk returns a tuple with the MaxReportNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReportNbr

`func (o *TimeSyncExposureSubsc1) SetMaxReportNbr(v int32)`

SetMaxReportNbr sets MaxReportNbr field to given value.

### HasMaxReportNbr

`func (o *TimeSyncExposureSubsc1) HasMaxReportNbr() bool`

HasMaxReportNbr returns a boolean if a field has been set.

### GetExpiry

`func (o *TimeSyncExposureSubsc1) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *TimeSyncExposureSubsc1) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *TimeSyncExposureSubsc1) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *TimeSyncExposureSubsc1) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *TimeSyncExposureSubsc1) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *TimeSyncExposureSubsc1) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *TimeSyncExposureSubsc1) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *TimeSyncExposureSubsc1) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *TimeSyncExposureSubsc1) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *TimeSyncExposureSubsc1) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *TimeSyncExposureSubsc1) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *TimeSyncExposureSubsc1) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *TimeSyncExposureSubsc1) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *TimeSyncExposureSubsc1) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *TimeSyncExposureSubsc1) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *TimeSyncExposureSubsc1) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *TimeSyncExposureSubsc1) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *TimeSyncExposureSubsc1) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *TimeSyncExposureSubsc1) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *TimeSyncExposureSubsc1) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


