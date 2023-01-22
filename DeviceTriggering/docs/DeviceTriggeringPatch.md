# DeviceTriggeringPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidityPeriod** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**Priority** | Pointer to [**Priority**](Priority.md) |  | [optional] 
**ApplicationPortId** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535. | [optional] 
**AppSrcPortId** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535. | [optional] 
**TriggerPayload** | Pointer to **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 

## Methods

### NewDeviceTriggeringPatch

`func NewDeviceTriggeringPatch() *DeviceTriggeringPatch`

NewDeviceTriggeringPatch instantiates a new DeviceTriggeringPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceTriggeringPatchWithDefaults

`func NewDeviceTriggeringPatchWithDefaults() *DeviceTriggeringPatch`

NewDeviceTriggeringPatchWithDefaults instantiates a new DeviceTriggeringPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidityPeriod

`func (o *DeviceTriggeringPatch) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *DeviceTriggeringPatch) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *DeviceTriggeringPatch) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *DeviceTriggeringPatch) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### GetPriority

`func (o *DeviceTriggeringPatch) GetPriority() Priority`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DeviceTriggeringPatch) GetPriorityOk() (*Priority, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DeviceTriggeringPatch) SetPriority(v Priority)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DeviceTriggeringPatch) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetApplicationPortId

`func (o *DeviceTriggeringPatch) GetApplicationPortId() int32`

GetApplicationPortId returns the ApplicationPortId field if non-nil, zero value otherwise.

### GetApplicationPortIdOk

`func (o *DeviceTriggeringPatch) GetApplicationPortIdOk() (*int32, bool)`

GetApplicationPortIdOk returns a tuple with the ApplicationPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPortId

`func (o *DeviceTriggeringPatch) SetApplicationPortId(v int32)`

SetApplicationPortId sets ApplicationPortId field to given value.

### HasApplicationPortId

`func (o *DeviceTriggeringPatch) HasApplicationPortId() bool`

HasApplicationPortId returns a boolean if a field has been set.

### GetAppSrcPortId

`func (o *DeviceTriggeringPatch) GetAppSrcPortId() int32`

GetAppSrcPortId returns the AppSrcPortId field if non-nil, zero value otherwise.

### GetAppSrcPortIdOk

`func (o *DeviceTriggeringPatch) GetAppSrcPortIdOk() (*int32, bool)`

GetAppSrcPortIdOk returns a tuple with the AppSrcPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSrcPortId

`func (o *DeviceTriggeringPatch) SetAppSrcPortId(v int32)`

SetAppSrcPortId sets AppSrcPortId field to given value.

### HasAppSrcPortId

`func (o *DeviceTriggeringPatch) HasAppSrcPortId() bool`

HasAppSrcPortId returns a boolean if a field has been set.

### GetTriggerPayload

`func (o *DeviceTriggeringPatch) GetTriggerPayload() string`

GetTriggerPayload returns the TriggerPayload field if non-nil, zero value otherwise.

### GetTriggerPayloadOk

`func (o *DeviceTriggeringPatch) GetTriggerPayloadOk() (*string, bool)`

GetTriggerPayloadOk returns a tuple with the TriggerPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPayload

`func (o *DeviceTriggeringPatch) SetTriggerPayload(v string)`

SetTriggerPayload sets TriggerPayload field to given value.

### HasTriggerPayload

`func (o *DeviceTriggeringPatch) HasTriggerPayload() bool`

HasTriggerPayload returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *DeviceTriggeringPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *DeviceTriggeringPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *DeviceTriggeringPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *DeviceTriggeringPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRequestTestNotification

`func (o *DeviceTriggeringPatch) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *DeviceTriggeringPatch) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *DeviceTriggeringPatch) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *DeviceTriggeringPatch) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *DeviceTriggeringPatch) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *DeviceTriggeringPatch) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *DeviceTriggeringPatch) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *DeviceTriggeringPatch) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


