# GMDViaMBMSByxMB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**NotificationDestination** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**MbmsLocArea** | Pointer to [**MbmsLocArea**](MbmsLocArea.md) |  | [optional] 
**MessageDeliveryStartTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**MessageDeliveryStopTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**GroupMessagePayload** | Pointer to **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | [optional] 
**ScefMessageDeliveryIPv4** | Pointer to **string** | string identifying a Ipv4 address formatted in the \&quot;dotted decimal\&quot; notation as defined in IETF RFC 1166, with \&quot;readOnly&#x3D;true\&quot; property. | [optional] [readonly] 
**ScefMessageDeliveryIPv6** | Pointer to **string** | string identifying a Ipv6 address formatted according to clause 4 in IETF RFC 5952, with \&quot;readOnly&#x3D;true\&quot; property. The mixed Ipv4 Ipv6 notation according to clause 5 of IETF RFC 5952 shall not be used. | [optional] [readonly] 
**ScefMessageDeliveryPort** | Pointer to **int32** | Unsigned integer with valid values between 0 and 65535, with \&quot;readOnly&#x3D;true\&quot; property. | [optional] [readonly] 

## Methods

### NewGMDViaMBMSByxMB

`func NewGMDViaMBMSByxMB(notificationDestination string, ) *GMDViaMBMSByxMB`

NewGMDViaMBMSByxMB instantiates a new GMDViaMBMSByxMB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGMDViaMBMSByxMBWithDefaults

`func NewGMDViaMBMSByxMBWithDefaults() *GMDViaMBMSByxMB`

NewGMDViaMBMSByxMBWithDefaults instantiates a new GMDViaMBMSByxMB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *GMDViaMBMSByxMB) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GMDViaMBMSByxMB) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GMDViaMBMSByxMB) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *GMDViaMBMSByxMB) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *GMDViaMBMSByxMB) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *GMDViaMBMSByxMB) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *GMDViaMBMSByxMB) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *GMDViaMBMSByxMB) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *GMDViaMBMSByxMB) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *GMDViaMBMSByxMB) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *GMDViaMBMSByxMB) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *GMDViaMBMSByxMB) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *GMDViaMBMSByxMB) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *GMDViaMBMSByxMB) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *GMDViaMBMSByxMB) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetMbmsLocArea

`func (o *GMDViaMBMSByxMB) GetMbmsLocArea() MbmsLocArea`

GetMbmsLocArea returns the MbmsLocArea field if non-nil, zero value otherwise.

### GetMbmsLocAreaOk

`func (o *GMDViaMBMSByxMB) GetMbmsLocAreaOk() (*MbmsLocArea, bool)`

GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsLocArea

`func (o *GMDViaMBMSByxMB) SetMbmsLocArea(v MbmsLocArea)`

SetMbmsLocArea sets MbmsLocArea field to given value.

### HasMbmsLocArea

`func (o *GMDViaMBMSByxMB) HasMbmsLocArea() bool`

HasMbmsLocArea returns a boolean if a field has been set.

### GetMessageDeliveryStartTime

`func (o *GMDViaMBMSByxMB) GetMessageDeliveryStartTime() time.Time`

GetMessageDeliveryStartTime returns the MessageDeliveryStartTime field if non-nil, zero value otherwise.

### GetMessageDeliveryStartTimeOk

`func (o *GMDViaMBMSByxMB) GetMessageDeliveryStartTimeOk() (*time.Time, bool)`

GetMessageDeliveryStartTimeOk returns a tuple with the MessageDeliveryStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDeliveryStartTime

`func (o *GMDViaMBMSByxMB) SetMessageDeliveryStartTime(v time.Time)`

SetMessageDeliveryStartTime sets MessageDeliveryStartTime field to given value.

### HasMessageDeliveryStartTime

`func (o *GMDViaMBMSByxMB) HasMessageDeliveryStartTime() bool`

HasMessageDeliveryStartTime returns a boolean if a field has been set.

### GetMessageDeliveryStopTime

`func (o *GMDViaMBMSByxMB) GetMessageDeliveryStopTime() time.Time`

GetMessageDeliveryStopTime returns the MessageDeliveryStopTime field if non-nil, zero value otherwise.

### GetMessageDeliveryStopTimeOk

`func (o *GMDViaMBMSByxMB) GetMessageDeliveryStopTimeOk() (*time.Time, bool)`

GetMessageDeliveryStopTimeOk returns a tuple with the MessageDeliveryStopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDeliveryStopTime

`func (o *GMDViaMBMSByxMB) SetMessageDeliveryStopTime(v time.Time)`

SetMessageDeliveryStopTime sets MessageDeliveryStopTime field to given value.

### HasMessageDeliveryStopTime

`func (o *GMDViaMBMSByxMB) HasMessageDeliveryStopTime() bool`

HasMessageDeliveryStopTime returns a boolean if a field has been set.

### GetGroupMessagePayload

`func (o *GMDViaMBMSByxMB) GetGroupMessagePayload() string`

GetGroupMessagePayload returns the GroupMessagePayload field if non-nil, zero value otherwise.

### GetGroupMessagePayloadOk

`func (o *GMDViaMBMSByxMB) GetGroupMessagePayloadOk() (*string, bool)`

GetGroupMessagePayloadOk returns a tuple with the GroupMessagePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMessagePayload

`func (o *GMDViaMBMSByxMB) SetGroupMessagePayload(v string)`

SetGroupMessagePayload sets GroupMessagePayload field to given value.

### HasGroupMessagePayload

`func (o *GMDViaMBMSByxMB) HasGroupMessagePayload() bool`

HasGroupMessagePayload returns a boolean if a field has been set.

### GetScefMessageDeliveryIPv4

`func (o *GMDViaMBMSByxMB) GetScefMessageDeliveryIPv4() string`

GetScefMessageDeliveryIPv4 returns the ScefMessageDeliveryIPv4 field if non-nil, zero value otherwise.

### GetScefMessageDeliveryIPv4Ok

`func (o *GMDViaMBMSByxMB) GetScefMessageDeliveryIPv4Ok() (*string, bool)`

GetScefMessageDeliveryIPv4Ok returns a tuple with the ScefMessageDeliveryIPv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefMessageDeliveryIPv4

`func (o *GMDViaMBMSByxMB) SetScefMessageDeliveryIPv4(v string)`

SetScefMessageDeliveryIPv4 sets ScefMessageDeliveryIPv4 field to given value.

### HasScefMessageDeliveryIPv4

`func (o *GMDViaMBMSByxMB) HasScefMessageDeliveryIPv4() bool`

HasScefMessageDeliveryIPv4 returns a boolean if a field has been set.

### GetScefMessageDeliveryIPv6

`func (o *GMDViaMBMSByxMB) GetScefMessageDeliveryIPv6() string`

GetScefMessageDeliveryIPv6 returns the ScefMessageDeliveryIPv6 field if non-nil, zero value otherwise.

### GetScefMessageDeliveryIPv6Ok

`func (o *GMDViaMBMSByxMB) GetScefMessageDeliveryIPv6Ok() (*string, bool)`

GetScefMessageDeliveryIPv6Ok returns a tuple with the ScefMessageDeliveryIPv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefMessageDeliveryIPv6

`func (o *GMDViaMBMSByxMB) SetScefMessageDeliveryIPv6(v string)`

SetScefMessageDeliveryIPv6 sets ScefMessageDeliveryIPv6 field to given value.

### HasScefMessageDeliveryIPv6

`func (o *GMDViaMBMSByxMB) HasScefMessageDeliveryIPv6() bool`

HasScefMessageDeliveryIPv6 returns a boolean if a field has been set.

### GetScefMessageDeliveryPort

`func (o *GMDViaMBMSByxMB) GetScefMessageDeliveryPort() int32`

GetScefMessageDeliveryPort returns the ScefMessageDeliveryPort field if non-nil, zero value otherwise.

### GetScefMessageDeliveryPortOk

`func (o *GMDViaMBMSByxMB) GetScefMessageDeliveryPortOk() (*int32, bool)`

GetScefMessageDeliveryPortOk returns a tuple with the ScefMessageDeliveryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScefMessageDeliveryPort

`func (o *GMDViaMBMSByxMB) SetScefMessageDeliveryPort(v int32)`

SetScefMessageDeliveryPort sets ScefMessageDeliveryPort field to given value.

### HasScefMessageDeliveryPort

`func (o *GMDViaMBMSByxMB) HasScefMessageDeliveryPort() bool`

HasScefMessageDeliveryPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


