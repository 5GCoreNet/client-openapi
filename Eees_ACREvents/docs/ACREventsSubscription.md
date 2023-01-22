# ACREventsSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Represents a unique identifier of the EEC. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EasIds** | **[]string** | The list of application identifiers of the EASs. | 
**AcIds** | Pointer to **[]string** | List of AC identities | [optional] 
**EventIds** | [**ACREventIDs**](ACREventIDs.md) |  | 
**NotificationDestination** | **string** | string providing an URI formatted according to IETF RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the ECS to send a test notification. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewACREventsSubscription

`func NewACREventsSubscription(eecId string, easIds []string, eventIds ACREventIDs, notificationDestination string, ) *ACREventsSubscription`

NewACREventsSubscription instantiates a new ACREventsSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACREventsSubscriptionWithDefaults

`func NewACREventsSubscriptionWithDefaults() *ACREventsSubscription`

NewACREventsSubscriptionWithDefaults instantiates a new ACREventsSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *ACREventsSubscription) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *ACREventsSubscription) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *ACREventsSubscription) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetUeId

`func (o *ACREventsSubscription) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ACREventsSubscription) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ACREventsSubscription) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ACREventsSubscription) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetExpTime

`func (o *ACREventsSubscription) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ACREventsSubscription) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ACREventsSubscription) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ACREventsSubscription) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEasIds

`func (o *ACREventsSubscription) GetEasIds() []string`

GetEasIds returns the EasIds field if non-nil, zero value otherwise.

### GetEasIdsOk

`func (o *ACREventsSubscription) GetEasIdsOk() (*[]string, bool)`

GetEasIdsOk returns a tuple with the EasIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasIds

`func (o *ACREventsSubscription) SetEasIds(v []string)`

SetEasIds sets EasIds field to given value.


### GetAcIds

`func (o *ACREventsSubscription) GetAcIds() []string`

GetAcIds returns the AcIds field if non-nil, zero value otherwise.

### GetAcIdsOk

`func (o *ACREventsSubscription) GetAcIdsOk() (*[]string, bool)`

GetAcIdsOk returns a tuple with the AcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcIds

`func (o *ACREventsSubscription) SetAcIds(v []string)`

SetAcIds sets AcIds field to given value.

### HasAcIds

`func (o *ACREventsSubscription) HasAcIds() bool`

HasAcIds returns a boolean if a field has been set.

### GetEventIds

`func (o *ACREventsSubscription) GetEventIds() ACREventIDs`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *ACREventsSubscription) GetEventIdsOk() (*ACREventIDs, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *ACREventsSubscription) SetEventIds(v ACREventIDs)`

SetEventIds sets EventIds field to given value.


### GetNotificationDestination

`func (o *ACREventsSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ACREventsSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ACREventsSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *ACREventsSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *ACREventsSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *ACREventsSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *ACREventsSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *ACREventsSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *ACREventsSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *ACREventsSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *ACREventsSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *ACREventsSubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *ACREventsSubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *ACREventsSubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *ACREventsSubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


