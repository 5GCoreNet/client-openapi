# DeviceTriggering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ValidityPeriod** | **int32** | Unsigned integer identifying a period of time in units of seconds. | 
**Priority** | [**Priority**](Priority.md) |  | 
**ApplicationPortId** | **int32** | Unsigned integer with valid values between 0 and 65535. | 
**AppSrcPortId** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535. | [optional] 
**TriggerPayload** | **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | 
**NotificationDestination** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**DeliveryResult** | Pointer to [**DeliveryResult**](DeliveryResult.md) |  | [optional] 

## Methods

### NewDeviceTriggering

`func NewDeviceTriggering(validityPeriod int32, priority Priority, applicationPortId int32, triggerPayload string, notificationDestination string, ) *DeviceTriggering`

NewDeviceTriggering instantiates a new DeviceTriggering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTriggeringWithDefaults

`func NewDeviceTriggeringWithDefaults() *DeviceTriggering`

NewDeviceTriggeringWithDefaults instantiates a new DeviceTriggering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *DeviceTriggering) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *DeviceTriggering) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *DeviceTriggering) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *DeviceTriggering) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetExternalId

`func (o *DeviceTriggering) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DeviceTriggering) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DeviceTriggering) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *DeviceTriggering) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *DeviceTriggering) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *DeviceTriggering) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *DeviceTriggering) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *DeviceTriggering) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *DeviceTriggering) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *DeviceTriggering) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *DeviceTriggering) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *DeviceTriggering) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetValidityPeriod

`func (o *DeviceTriggering) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *DeviceTriggering) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *DeviceTriggering) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.


### GetPriority

`func (o *DeviceTriggering) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeviceTriggering) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeviceTriggering) SetPriority(v Priority)`

SetPriority sets Priority field to given value.


### GetApplicationPortId

`func (o *DeviceTriggering) GetApplicationPortId() int32`

GetApplicationPortId returns the ApplicationPortId field if non-nil, zero value otherwise.

### GetApplicationPortIdOk

`func (o *DeviceTriggering) GetApplicationPortIdOk() (*int32, bool)`

GetApplicationPortIdOk returns a tuple with the ApplicationPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPortId

`func (o *DeviceTriggering) SetApplicationPortId(v int32)`

SetApplicationPortId sets ApplicationPortId field to given value.


### GetAppSrcPortId

`func (o *DeviceTriggering) GetAppSrcPortId() int32`

GetAppSrcPortId returns the AppSrcPortId field if non-nil, zero value otherwise.

### GetAppSrcPortIdOk

`func (o *DeviceTriggering) GetAppSrcPortIdOk() (*int32, bool)`

GetAppSrcPortIdOk returns a tuple with the AppSrcPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSrcPortId

`func (o *DeviceTriggering) SetAppSrcPortId(v int32)`

SetAppSrcPortId sets AppSrcPortId field to given value.

### HasAppSrcPortId

`func (o *DeviceTriggering) HasAppSrcPortId() bool`

HasAppSrcPortId returns a boolean if a field has been set.

### GetTriggerPayload

`func (o *DeviceTriggering) GetTriggerPayload() string`

GetTriggerPayload returns the TriggerPayload field if non-nil, zero value otherwise.

### GetTriggerPayloadOk

`func (o *DeviceTriggering) GetTriggerPayloadOk() (*string, bool)`

GetTriggerPayloadOk returns a tuple with the TriggerPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPayload

`func (o *DeviceTriggering) SetTriggerPayload(v string)`

SetTriggerPayload sets TriggerPayload field to given value.


### GetNotificationDestination

`func (o *DeviceTriggering) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *DeviceTriggering) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *DeviceTriggering) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *DeviceTriggering) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *DeviceTriggering) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *DeviceTriggering) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *DeviceTriggering) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *DeviceTriggering) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *DeviceTriggering) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *DeviceTriggering) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *DeviceTriggering) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetDeliveryResult

`func (o *DeviceTriggering) GetDeliveryResult() DeliveryResult`

GetDeliveryResult returns the DeliveryResult field if non-nil, zero value otherwise.

### GetDeliveryResultOk

`func (o *DeviceTriggering) GetDeliveryResultOk() (*DeliveryResult, bool)`

GetDeliveryResultOk returns a tuple with the DeliveryResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryResult

`func (o *DeviceTriggering) SetDeliveryResult(v DeliveryResult)`

SetDeliveryResult sets DeliveryResult field to given value.

### HasDeliveryResult

`func (o *DeviceTriggering) HasDeliveryResult() bool`

HasDeliveryResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


