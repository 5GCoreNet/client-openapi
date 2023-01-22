# EasDiscoverySubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Represents a unique identifier of the EEC. | 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**EasEventType** | [**EASDiscEventIDs**](EASDiscEventIDs.md) |  | 
**EasDiscoveryFilter** | Pointer to [**EasDiscoveryFilter**](EasDiscoveryFilter.md) |  | [optional] 
**EasDynInfoFilter** | Pointer to [**EasDynamicInfoFilter**](EasDynamicInfoFilter.md) |  | [optional] 
**EasSvcContinuity** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by Subscriber to request the ECS to send a test notification. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewEasDiscoverySubscription

`func NewEasDiscoverySubscription(eecId string, easEventType EASDiscEventIDs, ) *EasDiscoverySubscription`

NewEasDiscoverySubscription instantiates a new EasDiscoverySubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDiscoverySubscriptionWithDefaults

`func NewEasDiscoverySubscriptionWithDefaults() *EasDiscoverySubscription`

NewEasDiscoverySubscriptionWithDefaults instantiates a new EasDiscoverySubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *EasDiscoverySubscription) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *EasDiscoverySubscription) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *EasDiscoverySubscription) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetUeId

`func (o *EasDiscoverySubscription) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *EasDiscoverySubscription) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *EasDiscoverySubscription) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *EasDiscoverySubscription) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetEasEventType

`func (o *EasDiscoverySubscription) GetEasEventType() EASDiscEventIDs`

GetEasEventType returns the EasEventType field if non-nil, zero value otherwise.

### GetEasEventTypeOk

`func (o *EasDiscoverySubscription) GetEasEventTypeOk() (*EASDiscEventIDs, bool)`

GetEasEventTypeOk returns a tuple with the EasEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasEventType

`func (o *EasDiscoverySubscription) SetEasEventType(v EASDiscEventIDs)`

SetEasEventType sets EasEventType field to given value.


### GetEasDiscoveryFilter

`func (o *EasDiscoverySubscription) GetEasDiscoveryFilter() EasDiscoveryFilter`

GetEasDiscoveryFilter returns the EasDiscoveryFilter field if non-nil, zero value otherwise.

### GetEasDiscoveryFilterOk

`func (o *EasDiscoverySubscription) GetEasDiscoveryFilterOk() (*EasDiscoveryFilter, bool)`

GetEasDiscoveryFilterOk returns a tuple with the EasDiscoveryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDiscoveryFilter

`func (o *EasDiscoverySubscription) SetEasDiscoveryFilter(v EasDiscoveryFilter)`

SetEasDiscoveryFilter sets EasDiscoveryFilter field to given value.

### HasEasDiscoveryFilter

`func (o *EasDiscoverySubscription) HasEasDiscoveryFilter() bool`

HasEasDiscoveryFilter returns a boolean if a field has been set.

### GetEasDynInfoFilter

`func (o *EasDiscoverySubscription) GetEasDynInfoFilter() EasDynamicInfoFilter`

GetEasDynInfoFilter returns the EasDynInfoFilter field if non-nil, zero value otherwise.

### GetEasDynInfoFilterOk

`func (o *EasDiscoverySubscription) GetEasDynInfoFilterOk() (*EasDynamicInfoFilter, bool)`

GetEasDynInfoFilterOk returns a tuple with the EasDynInfoFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDynInfoFilter

`func (o *EasDiscoverySubscription) SetEasDynInfoFilter(v EasDynamicInfoFilter)`

SetEasDynInfoFilter sets EasDynInfoFilter field to given value.

### HasEasDynInfoFilter

`func (o *EasDiscoverySubscription) HasEasDynInfoFilter() bool`

HasEasDynInfoFilter returns a boolean if a field has been set.

### GetEasSvcContinuity

`func (o *EasDiscoverySubscription) GetEasSvcContinuity() []ACRScenario`

GetEasSvcContinuity returns the EasSvcContinuity field if non-nil, zero value otherwise.

### GetEasSvcContinuityOk

`func (o *EasDiscoverySubscription) GetEasSvcContinuityOk() (*[]ACRScenario, bool)`

GetEasSvcContinuityOk returns a tuple with the EasSvcContinuity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasSvcContinuity

`func (o *EasDiscoverySubscription) SetEasSvcContinuity(v []ACRScenario)`

SetEasSvcContinuity sets EasSvcContinuity field to given value.

### HasEasSvcContinuity

`func (o *EasDiscoverySubscription) HasEasSvcContinuity() bool`

HasEasSvcContinuity returns a boolean if a field has been set.

### GetExpTime

`func (o *EasDiscoverySubscription) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EasDiscoverySubscription) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EasDiscoverySubscription) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EasDiscoverySubscription) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *EasDiscoverySubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *EasDiscoverySubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *EasDiscoverySubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *EasDiscoverySubscription) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *EasDiscoverySubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *EasDiscoverySubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *EasDiscoverySubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *EasDiscoverySubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *EasDiscoverySubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *EasDiscoverySubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *EasDiscoverySubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *EasDiscoverySubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *EasDiscoverySubscription) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *EasDiscoverySubscription) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *EasDiscoverySubscription) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *EasDiscoverySubscription) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


