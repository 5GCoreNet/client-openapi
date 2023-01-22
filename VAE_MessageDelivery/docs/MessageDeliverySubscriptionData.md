# MessageDeliverySubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSerId** | **string** | Represents the V2X application specific server identifier. | 
**ServiceId** | **string** | Represents the V2X service ID to which a V2X message belongs. | 
**GeoId** | Pointer to **string** | Represents a geographical area identifier. | [optional] 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the NF service consumer to request the VAE server to send a test notification as defined in clause 6.1.5.3. Set to false or omitted otherwise.  | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMessageDeliverySubscriptionData

`func NewMessageDeliverySubscriptionData(appSerId string, serviceId string, notifUri string, ) *MessageDeliverySubscriptionData`

NewMessageDeliverySubscriptionData instantiates a new MessageDeliverySubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDeliverySubscriptionDataWithDefaults

`func NewMessageDeliverySubscriptionDataWithDefaults() *MessageDeliverySubscriptionData`

NewMessageDeliverySubscriptionDataWithDefaults instantiates a new MessageDeliverySubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSerId

`func (o *MessageDeliverySubscriptionData) GetAppSerId() string`

GetAppSerId returns the AppSerId field if non-nil, zero value otherwise.

### GetAppSerIdOk

`func (o *MessageDeliverySubscriptionData) GetAppSerIdOk() (*string, bool)`

GetAppSerIdOk returns a tuple with the AppSerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSerId

`func (o *MessageDeliverySubscriptionData) SetAppSerId(v string)`

SetAppSerId sets AppSerId field to given value.


### GetServiceId

`func (o *MessageDeliverySubscriptionData) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MessageDeliverySubscriptionData) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MessageDeliverySubscriptionData) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetGeoId

`func (o *MessageDeliverySubscriptionData) GetGeoId() string`

GetGeoId returns the GeoId field if non-nil, zero value otherwise.

### GetGeoIdOk

`func (o *MessageDeliverySubscriptionData) GetGeoIdOk() (*string, bool)`

GetGeoIdOk returns a tuple with the GeoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoId

`func (o *MessageDeliverySubscriptionData) SetGeoId(v string)`

SetGeoId sets GeoId field to given value.

### HasGeoId

`func (o *MessageDeliverySubscriptionData) HasGeoId() bool`

HasGeoId returns a boolean if a field has been set.

### GetNotifUri

`func (o *MessageDeliverySubscriptionData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *MessageDeliverySubscriptionData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *MessageDeliverySubscriptionData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.


### GetRequestTestNotification

`func (o *MessageDeliverySubscriptionData) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *MessageDeliverySubscriptionData) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *MessageDeliverySubscriptionData) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *MessageDeliverySubscriptionData) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *MessageDeliverySubscriptionData) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *MessageDeliverySubscriptionData) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *MessageDeliverySubscriptionData) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *MessageDeliverySubscriptionData) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MessageDeliverySubscriptionData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MessageDeliverySubscriptionData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MessageDeliverySubscriptionData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MessageDeliverySubscriptionData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


