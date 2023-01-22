# GMDViaMBMSByxMBPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbmsLocArea** | Pointer to [**MbmsLocArea**](MbmsLocArea.md) |  | [optional] 
**MessageDeliveryStartTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**MessageDeliveryStopTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**GroupMessagePayload** | Pointer to **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewGMDViaMBMSByxMBPatch

`func NewGMDViaMBMSByxMBPatch() *GMDViaMBMSByxMBPatch`

NewGMDViaMBMSByxMBPatch instantiates a new GMDViaMBMSByxMBPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGMDViaMBMSByxMBPatchWithDefaults

`func NewGMDViaMBMSByxMBPatchWithDefaults() *GMDViaMBMSByxMBPatch`

NewGMDViaMBMSByxMBPatchWithDefaults instantiates a new GMDViaMBMSByxMBPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbmsLocArea

`func (o *GMDViaMBMSByxMBPatch) GetMbmsLocArea() MbmsLocArea`

GetMbmsLocArea returns the MbmsLocArea field if non-nil, zero value otherwise.

### GetMbmsLocAreaOk

`func (o *GMDViaMBMSByxMBPatch) GetMbmsLocAreaOk() (*MbmsLocArea, bool)`

GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsLocArea

`func (o *GMDViaMBMSByxMBPatch) SetMbmsLocArea(v MbmsLocArea)`

SetMbmsLocArea sets MbmsLocArea field to given value.

### HasMbmsLocArea

`func (o *GMDViaMBMSByxMBPatch) HasMbmsLocArea() bool`

HasMbmsLocArea returns a boolean if a field has been set.

### GetMessageDeliveryStartTime

`func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStartTime() time.Time`

GetMessageDeliveryStartTime returns the MessageDeliveryStartTime field if non-nil, zero value otherwise.

### GetMessageDeliveryStartTimeOk

`func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStartTimeOk() (*time.Time, bool)`

GetMessageDeliveryStartTimeOk returns a tuple with the MessageDeliveryStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDeliveryStartTime

`func (o *GMDViaMBMSByxMBPatch) SetMessageDeliveryStartTime(v time.Time)`

SetMessageDeliveryStartTime sets MessageDeliveryStartTime field to given value.

### HasMessageDeliveryStartTime

`func (o *GMDViaMBMSByxMBPatch) HasMessageDeliveryStartTime() bool`

HasMessageDeliveryStartTime returns a boolean if a field has been set.

### GetMessageDeliveryStopTime

`func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStopTime() time.Time`

GetMessageDeliveryStopTime returns the MessageDeliveryStopTime field if non-nil, zero value otherwise.

### GetMessageDeliveryStopTimeOk

`func (o *GMDViaMBMSByxMBPatch) GetMessageDeliveryStopTimeOk() (*time.Time, bool)`

GetMessageDeliveryStopTimeOk returns a tuple with the MessageDeliveryStopTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDeliveryStopTime

`func (o *GMDViaMBMSByxMBPatch) SetMessageDeliveryStopTime(v time.Time)`

SetMessageDeliveryStopTime sets MessageDeliveryStopTime field to given value.

### HasMessageDeliveryStopTime

`func (o *GMDViaMBMSByxMBPatch) HasMessageDeliveryStopTime() bool`

HasMessageDeliveryStopTime returns a boolean if a field has been set.

### GetGroupMessagePayload

`func (o *GMDViaMBMSByxMBPatch) GetGroupMessagePayload() string`

GetGroupMessagePayload returns the GroupMessagePayload field if non-nil, zero value otherwise.

### GetGroupMessagePayloadOk

`func (o *GMDViaMBMSByxMBPatch) GetGroupMessagePayloadOk() (*string, bool)`

GetGroupMessagePayloadOk returns a tuple with the GroupMessagePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMessagePayload

`func (o *GMDViaMBMSByxMBPatch) SetGroupMessagePayload(v string)`

SetGroupMessagePayload sets GroupMessagePayload field to given value.

### HasGroupMessagePayload

`func (o *GMDViaMBMSByxMBPatch) HasGroupMessagePayload() bool`

HasGroupMessagePayload returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *GMDViaMBMSByxMBPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *GMDViaMBMSByxMBPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *GMDViaMBMSByxMBPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *GMDViaMBMSByxMBPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


