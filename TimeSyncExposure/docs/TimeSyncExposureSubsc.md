# TimeSyncExposureSubsc

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

### NewTimeSyncExposureSubsc

`func NewTimeSyncExposureSubsc(subsNotifId string, subsNotifUri string, ) *TimeSyncExposureSubsc`

NewTimeSyncExposureSubsc instantiates a new TimeSyncExposureSubsc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSyncExposureSubscWithDefaults

`func NewTimeSyncExposureSubscWithDefaults() *TimeSyncExposureSubsc`

NewTimeSyncExposureSubscWithDefaults instantiates a new TimeSyncExposureSubsc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExterGroupId

`func (o *TimeSyncExposureSubsc) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *TimeSyncExposureSubsc) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *TimeSyncExposureSubsc) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.

### HasExterGroupId

`func (o *TimeSyncExposureSubsc) HasExterGroupId() bool`

HasExterGroupId returns a boolean if a field has been set.

### GetGpsis

`func (o *TimeSyncExposureSubsc) GetGpsis() []string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *TimeSyncExposureSubsc) GetGpsisOk() (*[]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *TimeSyncExposureSubsc) SetGpsis(v []string)`

SetGpsis sets Gpsis field to given value.

### HasGpsis

`func (o *TimeSyncExposureSubsc) HasGpsis() bool`

HasGpsis returns a boolean if a field has been set.

### GetAnyUeInd

`func (o *TimeSyncExposureSubsc) GetAnyUeInd() bool`

GetAnyUeInd returns the AnyUeInd field if non-nil, zero value otherwise.

### GetAnyUeIndOk

`func (o *TimeSyncExposureSubsc) GetAnyUeIndOk() (*bool, bool)`

GetAnyUeIndOk returns a tuple with the AnyUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyUeInd

`func (o *TimeSyncExposureSubsc) SetAnyUeInd(v bool)`

SetAnyUeInd sets AnyUeInd field to given value.

### HasAnyUeInd

`func (o *TimeSyncExposureSubsc) HasAnyUeInd() bool`

HasAnyUeInd returns a boolean if a field has been set.

### GetAfServiceId

`func (o *TimeSyncExposureSubsc) GetAfServiceId() string`

GetAfServiceId returns the AfServiceId field if non-nil, zero value otherwise.

### GetAfServiceIdOk

`func (o *TimeSyncExposureSubsc) GetAfServiceIdOk() (*string, bool)`

GetAfServiceIdOk returns a tuple with the AfServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfServiceId

`func (o *TimeSyncExposureSubsc) SetAfServiceId(v string)`

SetAfServiceId sets AfServiceId field to given value.

### HasAfServiceId

`func (o *TimeSyncExposureSubsc) HasAfServiceId() bool`

HasAfServiceId returns a boolean if a field has been set.

### GetDnn

`func (o *TimeSyncExposureSubsc) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TimeSyncExposureSubsc) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TimeSyncExposureSubsc) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TimeSyncExposureSubsc) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *TimeSyncExposureSubsc) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TimeSyncExposureSubsc) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TimeSyncExposureSubsc) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TimeSyncExposureSubsc) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetSubsNotifId

`func (o *TimeSyncExposureSubsc) GetSubsNotifId() string`

GetSubsNotifId returns the SubsNotifId field if non-nil, zero value otherwise.

### GetSubsNotifIdOk

`func (o *TimeSyncExposureSubsc) GetSubsNotifIdOk() (*string, bool)`

GetSubsNotifIdOk returns a tuple with the SubsNotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifId

`func (o *TimeSyncExposureSubsc) SetSubsNotifId(v string)`

SetSubsNotifId sets SubsNotifId field to given value.


### GetSubsNotifUri

`func (o *TimeSyncExposureSubsc) GetSubsNotifUri() string`

GetSubsNotifUri returns the SubsNotifUri field if non-nil, zero value otherwise.

### GetSubsNotifUriOk

`func (o *TimeSyncExposureSubsc) GetSubsNotifUriOk() (*string, bool)`

GetSubsNotifUriOk returns a tuple with the SubsNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsNotifUri

`func (o *TimeSyncExposureSubsc) SetSubsNotifUri(v string)`

SetSubsNotifUri sets SubsNotifUri field to given value.


### GetSubscribedEvents

`func (o *TimeSyncExposureSubsc) GetSubscribedEvents() []SubscribedEvent`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *TimeSyncExposureSubsc) GetSubscribedEventsOk() (*[]SubscribedEvent, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *TimeSyncExposureSubsc) SetSubscribedEvents(v []SubscribedEvent)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *TimeSyncExposureSubsc) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### GetEventFilters

`func (o *TimeSyncExposureSubsc) GetEventFilters() []EventFilter`

GetEventFilters returns the EventFilters field if non-nil, zero value otherwise.

### GetEventFiltersOk

`func (o *TimeSyncExposureSubsc) GetEventFiltersOk() (*[]EventFilter, bool)`

GetEventFiltersOk returns a tuple with the EventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilters

`func (o *TimeSyncExposureSubsc) SetEventFilters(v []EventFilter)`

SetEventFilters sets EventFilters field to given value.

### HasEventFilters

`func (o *TimeSyncExposureSubsc) HasEventFilters() bool`

HasEventFilters returns a boolean if a field has been set.

### GetNotifMethod

`func (o *TimeSyncExposureSubsc) GetNotifMethod() NotificationMethod`

GetNotifMethod returns the NotifMethod field if non-nil, zero value otherwise.

### GetNotifMethodOk

`func (o *TimeSyncExposureSubsc) GetNotifMethodOk() (*NotificationMethod, bool)`

GetNotifMethodOk returns a tuple with the NotifMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifMethod

`func (o *TimeSyncExposureSubsc) SetNotifMethod(v NotificationMethod)`

SetNotifMethod sets NotifMethod field to given value.

### HasNotifMethod

`func (o *TimeSyncExposureSubsc) HasNotifMethod() bool`

HasNotifMethod returns a boolean if a field has been set.

### GetMaxReportNbr

`func (o *TimeSyncExposureSubsc) GetMaxReportNbr() int32`

GetMaxReportNbr returns the MaxReportNbr field if non-nil, zero value otherwise.

### GetMaxReportNbrOk

`func (o *TimeSyncExposureSubsc) GetMaxReportNbrOk() (*int32, bool)`

GetMaxReportNbrOk returns a tuple with the MaxReportNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReportNbr

`func (o *TimeSyncExposureSubsc) SetMaxReportNbr(v int32)`

SetMaxReportNbr sets MaxReportNbr field to given value.

### HasMaxReportNbr

`func (o *TimeSyncExposureSubsc) HasMaxReportNbr() bool`

HasMaxReportNbr returns a boolean if a field has been set.

### GetExpiry

`func (o *TimeSyncExposureSubsc) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *TimeSyncExposureSubsc) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *TimeSyncExposureSubsc) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *TimeSyncExposureSubsc) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *TimeSyncExposureSubsc) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *TimeSyncExposureSubsc) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *TimeSyncExposureSubsc) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *TimeSyncExposureSubsc) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *TimeSyncExposureSubsc) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *TimeSyncExposureSubsc) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *TimeSyncExposureSubsc) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *TimeSyncExposureSubsc) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *TimeSyncExposureSubsc) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *TimeSyncExposureSubsc) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *TimeSyncExposureSubsc) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *TimeSyncExposureSubsc) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *TimeSyncExposureSubsc) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *TimeSyncExposureSubsc) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *TimeSyncExposureSubsc) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *TimeSyncExposureSubsc) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


