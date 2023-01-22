# AcrMgntEventsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**EasId** | **string** | Identifier of an EAS. | 
**EventSubscs** | [**[]AcrMgntEventSubsc**](AcrMgntEventSubsc.md) | The subscribed ACR management events. | 
**EvtReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**EventReports** | Pointer to [**[]AcrMgntEventReport**](AcrMgntEventReport.md) | The ACR management event report(s). | [optional] 
**AvailabilityInfo** | Pointer to [**AvailabilityNotif**](AvailabilityNotif.md) |  | [optional] 
**FailEventReports** | Pointer to [**[]FailureAcrMgntEventInfo**](FailureAcrMgntEventInfo.md) | Failure event reports. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the EAS to request the EES to send a test notification. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewAcrMgntEventsSubscription

`func NewAcrMgntEventsSubscription(easId string, eventSubscs []AcrMgntEventSubsc, notificationDestination string, ) *AcrMgntEventsSubscription`

NewAcrMgntEventsSubscription instantiates a new AcrMgntEventsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcrMgntEventsSubscriptionWithDefaults

`func NewAcrMgntEventsSubscriptionWithDefaults() *AcrMgntEventsSubscription`

NewAcrMgntEventsSubscriptionWithDefaults instantiates a new AcrMgntEventsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *AcrMgntEventsSubscription) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *AcrMgntEventsSubscription) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *AcrMgntEventsSubscription) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *AcrMgntEventsSubscription) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetEasId

`func (o *AcrMgntEventsSubscription) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *AcrMgntEventsSubscription) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *AcrMgntEventsSubscription) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetEventSubscs

`func (o *AcrMgntEventsSubscription) GetEventSubscs() []AcrMgntEventSubsc`

GetEventSubscs returns the EventSubscs field if non-nil, zero value otherwise.

### GetEventSubscsOk

`func (o *AcrMgntEventsSubscription) GetEventSubscsOk() (*[]AcrMgntEventSubsc, bool)`

GetEventSubscsOk returns a tuple with the EventSubscs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubscs

`func (o *AcrMgntEventsSubscription) SetEventSubscs(v []AcrMgntEventSubsc)`

SetEventSubscs sets EventSubscs field to given value.


### GetEvtReq

`func (o *AcrMgntEventsSubscription) GetEvtReq() ReportingInformation`

GetEvtReq returns the EvtReq field if non-nil, zero value otherwise.

### GetEvtReqOk

`func (o *AcrMgntEventsSubscription) GetEvtReqOk() (*ReportingInformation, bool)`

GetEvtReqOk returns a tuple with the EvtReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvtReq

`func (o *AcrMgntEventsSubscription) SetEvtReq(v ReportingInformation)`

SetEvtReq sets EvtReq field to given value.

### HasEvtReq

`func (o *AcrMgntEventsSubscription) HasEvtReq() bool`

HasEvtReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *AcrMgntEventsSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *AcrMgntEventsSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *AcrMgntEventsSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetEventReports

`func (o *AcrMgntEventsSubscription) GetEventReports() []AcrMgntEventReport`

GetEventReports returns the EventReports field if non-nil, zero value otherwise.

### GetEventReportsOk

`func (o *AcrMgntEventsSubscription) GetEventReportsOk() (*[]AcrMgntEventReport, bool)`

GetEventReportsOk returns a tuple with the EventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReports

`func (o *AcrMgntEventsSubscription) SetEventReports(v []AcrMgntEventReport)`

SetEventReports sets EventReports field to given value.

### HasEventReports

`func (o *AcrMgntEventsSubscription) HasEventReports() bool`

HasEventReports returns a boolean if a field has been set.

### GetAvailabilityInfo

`func (o *AcrMgntEventsSubscription) GetAvailabilityInfo() AvailabilityNotif`

GetAvailabilityInfo returns the AvailabilityInfo field if non-nil, zero value otherwise.

### GetAvailabilityInfoOk

`func (o *AcrMgntEventsSubscription) GetAvailabilityInfoOk() (*AvailabilityNotif, bool)`

GetAvailabilityInfoOk returns a tuple with the AvailabilityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityInfo

`func (o *AcrMgntEventsSubscription) SetAvailabilityInfo(v AvailabilityNotif)`

SetAvailabilityInfo sets AvailabilityInfo field to given value.

### HasAvailabilityInfo

`func (o *AcrMgntEventsSubscription) HasAvailabilityInfo() bool`

HasAvailabilityInfo returns a boolean if a field has been set.

### GetFailEventReports

`func (o *AcrMgntEventsSubscription) GetFailEventReports() []FailureAcrMgntEventInfo`

GetFailEventReports returns the FailEventReports field if non-nil, zero value otherwise.

### GetFailEventReportsOk

`func (o *AcrMgntEventsSubscription) GetFailEventReportsOk() (*[]FailureAcrMgntEventInfo, bool)`

GetFailEventReportsOk returns a tuple with the FailEventReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailEventReports

`func (o *AcrMgntEventsSubscription) SetFailEventReports(v []FailureAcrMgntEventInfo)`

SetFailEventReports sets FailEventReports field to given value.

### HasFailEventReports

`func (o *AcrMgntEventsSubscription) HasFailEventReports() bool`

HasFailEventReports returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *AcrMgntEventsSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *AcrMgntEventsSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *AcrMgntEventsSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *AcrMgntEventsSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *AcrMgntEventsSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *AcrMgntEventsSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *AcrMgntEventsSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *AcrMgntEventsSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *AcrMgntEventsSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *AcrMgntEventsSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *AcrMgntEventsSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *AcrMgntEventsSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


