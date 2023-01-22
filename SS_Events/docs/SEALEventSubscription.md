# SEALEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriberId** | **string** | String identifying the subscriber of the event. | 
**EventSubs** | [**[]EventSubscription**](EventSubscription.md) | Subscribed events. | 
**EventReq** | [**ReportingInformation**](ReportingInformation.md) |  | 
**NotificationDestination** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the SEAL server to send a test notification. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**EventDetails** | Pointer to [**[]SEALEventDetail**](SEALEventDetail.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSEALEventSubscription

`func NewSEALEventSubscription(subscriberId string, eventSubs []EventSubscription, eventReq ReportingInformation, notificationDestination string, ) *SEALEventSubscription`

NewSEALEventSubscription instantiates a new SEALEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSEALEventSubscriptionWithDefaults

`func NewSEALEventSubscriptionWithDefaults() *SEALEventSubscription`

NewSEALEventSubscriptionWithDefaults instantiates a new SEALEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriberId

`func (o *SEALEventSubscription) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *SEALEventSubscription) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *SEALEventSubscription) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.


### GetEventSubs

`func (o *SEALEventSubscription) GetEventSubs() []EventSubscription`

GetEventSubs returns the EventSubs field if non-nil, zero value otherwise.

### GetEventSubsOk

`func (o *SEALEventSubscription) GetEventSubsOk() (*[]EventSubscription, bool)`

GetEventSubsOk returns a tuple with the EventSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubs

`func (o *SEALEventSubscription) SetEventSubs(v []EventSubscription)`

SetEventSubs sets EventSubs field to given value.


### GetEventReq

`func (o *SEALEventSubscription) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *SEALEventSubscription) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *SEALEventSubscription) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.


### GetNotificationDestination

`func (o *SEALEventSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *SEALEventSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *SEALEventSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *SEALEventSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *SEALEventSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *SEALEventSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *SEALEventSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *SEALEventSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *SEALEventSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *SEALEventSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *SEALEventSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetEventDetails

`func (o *SEALEventSubscription) GetEventDetails() []SEALEventDetail`

GetEventDetails returns the EventDetails field if non-nil, zero value otherwise.

### GetEventDetailsOk

`func (o *SEALEventSubscription) GetEventDetailsOk() (*[]SEALEventDetail, bool)`

GetEventDetailsOk returns a tuple with the EventDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDetails

`func (o *SEALEventSubscription) SetEventDetails(v []SEALEventDetail)`

SetEventDetails sets EventDetails field to given value.

### HasEventDetails

`func (o *SEALEventSubscription) HasEventDetails() bool`

HasEventDetails returns a boolean if a field has been set.

### GetSuppFeat

`func (o *SEALEventSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *SEALEventSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *SEALEventSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *SEALEventSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


