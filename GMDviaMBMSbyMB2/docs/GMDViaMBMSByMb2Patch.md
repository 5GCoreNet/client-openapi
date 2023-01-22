# GMDViaMBMSByMb2Patch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**MbmsLocArea** | Pointer to [**MbmsLocArea**](MbmsLocArea.md) |  | [optional] 
**MessageDeliveryStartTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**GroupMessagePayload** | Pointer to **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewGMDViaMBMSByMb2Patch

`func NewGMDViaMBMSByMb2Patch() *GMDViaMBMSByMb2Patch`

NewGMDViaMBMSByMb2Patch instantiates a new GMDViaMBMSByMb2Patch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGMDViaMBMSByMb2PatchWithDefaults

`func NewGMDViaMBMSByMb2PatchWithDefaults() *GMDViaMBMSByMb2Patch`

NewGMDViaMBMSByMb2PatchWithDefaults instantiates a new GMDViaMBMSByMb2Patch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalGroupId

`func (o *GMDViaMBMSByMb2Patch) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *GMDViaMBMSByMb2Patch) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *GMDViaMBMSByMb2Patch) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *GMDViaMBMSByMb2Patch) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetMbmsLocArea

`func (o *GMDViaMBMSByMb2Patch) GetMbmsLocArea() MbmsLocArea`

GetMbmsLocArea returns the MbmsLocArea field if non-nil, zero value otherwise.

### GetMbmsLocAreaOk

`func (o *GMDViaMBMSByMb2Patch) GetMbmsLocAreaOk() (*MbmsLocArea, bool)`

GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsLocArea

`func (o *GMDViaMBMSByMb2Patch) SetMbmsLocArea(v MbmsLocArea)`

SetMbmsLocArea sets MbmsLocArea field to given value.

### HasMbmsLocArea

`func (o *GMDViaMBMSByMb2Patch) HasMbmsLocArea() bool`

HasMbmsLocArea returns a boolean if a field has been set.

### GetMessageDeliveryStartTime

`func (o *GMDViaMBMSByMb2Patch) GetMessageDeliveryStartTime() time.Time`

GetMessageDeliveryStartTime returns the MessageDeliveryStartTime field if non-nil, zero value otherwise.

### GetMessageDeliveryStartTimeOk

`func (o *GMDViaMBMSByMb2Patch) GetMessageDeliveryStartTimeOk() (*time.Time, bool)`

GetMessageDeliveryStartTimeOk returns a tuple with the MessageDeliveryStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageDeliveryStartTime

`func (o *GMDViaMBMSByMb2Patch) SetMessageDeliveryStartTime(v time.Time)`

SetMessageDeliveryStartTime sets MessageDeliveryStartTime field to given value.

### HasMessageDeliveryStartTime

`func (o *GMDViaMBMSByMb2Patch) HasMessageDeliveryStartTime() bool`

HasMessageDeliveryStartTime returns a boolean if a field has been set.

### GetGroupMessagePayload

`func (o *GMDViaMBMSByMb2Patch) GetGroupMessagePayload() string`

GetGroupMessagePayload returns the GroupMessagePayload field if non-nil, zero value otherwise.

### GetGroupMessagePayloadOk

`func (o *GMDViaMBMSByMb2Patch) GetGroupMessagePayloadOk() (*string, bool)`

GetGroupMessagePayloadOk returns a tuple with the GroupMessagePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMessagePayload

`func (o *GMDViaMBMSByMb2Patch) SetGroupMessagePayload(v string)`

SetGroupMessagePayload sets GroupMessagePayload field to given value.

### HasGroupMessagePayload

`func (o *GMDViaMBMSByMb2Patch) HasGroupMessagePayload() bool`

HasGroupMessagePayload returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *GMDViaMBMSByMb2Patch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *GMDViaMBMSByMb2Patch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *GMDViaMBMSByMb2Patch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *GMDViaMBMSByMb2Patch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


