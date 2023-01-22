# NetworkStatusReportingSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**NotificationDestination** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**LocationArea** | [**LocationArea**](LocationArea.md) |  | 
**TimeDuration** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**ThresholdValues** | Pointer to **[]int32** | Identifies a list of congestion level(s) with exact value that the SCS/AS requests to be informed of when reached. | [optional] 
**ThresholdTypes** | Pointer to [**[]CongestionType**](CongestionType.md) | Identifies a list of congestion level(s) with abstracted value that the SCS/AS requests to be informed of when reached. | [optional] 

## Methods

### NewNetworkStatusReportingSubscription

`func NewNetworkStatusReportingSubscription(notificationDestination string, locationArea LocationArea, ) *NetworkStatusReportingSubscription`

NewNetworkStatusReportingSubscription instantiates a new NetworkStatusReportingSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStatusReportingSubscriptionWithDefaults

`func NewNetworkStatusReportingSubscriptionWithDefaults() *NetworkStatusReportingSubscription`

NewNetworkStatusReportingSubscriptionWithDefaults instantiates a new NetworkStatusReportingSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NetworkStatusReportingSubscription) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NetworkStatusReportingSubscription) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NetworkStatusReportingSubscription) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *NetworkStatusReportingSubscription) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NetworkStatusReportingSubscription) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NetworkStatusReportingSubscription) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NetworkStatusReportingSubscription) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NetworkStatusReportingSubscription) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *NetworkStatusReportingSubscription) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *NetworkStatusReportingSubscription) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *NetworkStatusReportingSubscription) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *NetworkStatusReportingSubscription) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *NetworkStatusReportingSubscription) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *NetworkStatusReportingSubscription) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *NetworkStatusReportingSubscription) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *NetworkStatusReportingSubscription) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *NetworkStatusReportingSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *NetworkStatusReportingSubscription) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *NetworkStatusReportingSubscription) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetLocationArea

`func (o *NetworkStatusReportingSubscription) GetLocationArea() LocationArea`

GetLocationArea returns the LocationArea field if non-nil, zero value otherwise.

### GetLocationAreaOk

`func (o *NetworkStatusReportingSubscription) GetLocationAreaOk() (*LocationArea, bool)`

GetLocationAreaOk returns a tuple with the LocationArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea

`func (o *NetworkStatusReportingSubscription) SetLocationArea(v LocationArea)`

SetLocationArea sets LocationArea field to given value.


### GetTimeDuration

`func (o *NetworkStatusReportingSubscription) GetTimeDuration() time.Time`

GetTimeDuration returns the TimeDuration field if non-nil, zero value otherwise.

### GetTimeDurationOk

`func (o *NetworkStatusReportingSubscription) GetTimeDurationOk() (*time.Time, bool)`

GetTimeDurationOk returns a tuple with the TimeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDuration

`func (o *NetworkStatusReportingSubscription) SetTimeDuration(v time.Time)`

SetTimeDuration sets TimeDuration field to given value.

### HasTimeDuration

`func (o *NetworkStatusReportingSubscription) HasTimeDuration() bool`

HasTimeDuration returns a boolean if a field has been set.

### GetThresholdValues

`func (o *NetworkStatusReportingSubscription) GetThresholdValues() []int32`

GetThresholdValues returns the ThresholdValues field if non-nil, zero value otherwise.

### GetThresholdValuesOk

`func (o *NetworkStatusReportingSubscription) GetThresholdValuesOk() (*[]int32, bool)`

GetThresholdValuesOk returns a tuple with the ThresholdValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValues

`func (o *NetworkStatusReportingSubscription) SetThresholdValues(v []int32)`

SetThresholdValues sets ThresholdValues field to given value.

### HasThresholdValues

`func (o *NetworkStatusReportingSubscription) HasThresholdValues() bool`

HasThresholdValues returns a boolean if a field has been set.

### GetThresholdTypes

`func (o *NetworkStatusReportingSubscription) GetThresholdTypes() []CongestionType`

GetThresholdTypes returns the ThresholdTypes field if non-nil, zero value otherwise.

### GetThresholdTypesOk

`func (o *NetworkStatusReportingSubscription) GetThresholdTypesOk() (*[]CongestionType, bool)`

GetThresholdTypesOk returns a tuple with the ThresholdTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdTypes

`func (o *NetworkStatusReportingSubscription) SetThresholdTypes(v []CongestionType)`

SetThresholdTypes sets ThresholdTypes field to given value.

### HasThresholdTypes

`func (o *NetworkStatusReportingSubscription) HasThresholdTypes() bool`

HasThresholdTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


