# EventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]CAPIFEvent**](CAPIFEvent.md) | Subscribed events | 
**EventFilters** | Pointer to [**[]CAPIFEventFilter**](CAPIFEventFilter.md) | Subscribed event filters | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the CAPIF core function to send a test notification as defined in in clause 7.6. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewEventSubscription

`func NewEventSubscription(events []CAPIFEvent, notificationDestination string, ) *EventSubscription`

NewEventSubscription instantiates a new EventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionWithDefaults

`func NewEventSubscriptionWithDefaults() *EventSubscription`

NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventSubscription) GetEvents() []CAPIFEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventSubscription) GetEventsOk() (*[]CAPIFEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventSubscription) SetEvents(v []CAPIFEvent)`

SetEvents sets Events field to given value.


### GetEventFilters

`func (o *EventSubscription) GetEventFilters() []CAPIFEventFilter`

GetEventFilters returns the EventFilters field if non-nil, zero value otherwise.

### GetEventFiltersOk

`func (o *EventSubscription) GetEventFiltersOk() (*[]CAPIFEventFilter, bool)`

GetEventFiltersOk returns a tuple with the EventFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilters

`func (o *EventSubscription) SetEventFilters(v []CAPIFEventFilter)`

SetEventFilters sets EventFilters field to given value.

### HasEventFilters

`func (o *EventSubscription) HasEventFilters() bool`

HasEventFilters returns a boolean if a field has been set.

### GetEventReq

`func (o *EventSubscription) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *EventSubscription) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *EventSubscription) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *EventSubscription) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *EventSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *EventSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *EventSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *EventSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *EventSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *EventSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *EventSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *EventSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *EventSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *EventSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *EventSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EventSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EventSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EventSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EventSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


