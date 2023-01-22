# DownlinkMessageDeliveryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | Pointer to **string** | Represents the identifier of the V2X UE. | [optional] 
**GroupId** | Pointer to **string** | Represents the group ID for which a V2X message is addressed. | [optional] 
**Duration** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**GeoId** | Pointer to **string** | Represents a geographical area identifier. | [optional] 
**Payload** | **string** | string with format &#39;bytes&#39; as defined in OpenAPI | 

## Methods

### NewDownlinkMessageDeliveryData

`func NewDownlinkMessageDeliveryData(payload string, ) *DownlinkMessageDeliveryData`

NewDownlinkMessageDeliveryData instantiates a new DownlinkMessageDeliveryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownlinkMessageDeliveryDataWithDefaults

`func NewDownlinkMessageDeliveryDataWithDefaults() *DownlinkMessageDeliveryData`

NewDownlinkMessageDeliveryDataWithDefaults instantiates a new DownlinkMessageDeliveryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *DownlinkMessageDeliveryData) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *DownlinkMessageDeliveryData) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *DownlinkMessageDeliveryData) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *DownlinkMessageDeliveryData) HasUeId() bool`

HasUeId returns a boolean if a field has been set.

### GetGroupId

`func (o *DownlinkMessageDeliveryData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DownlinkMessageDeliveryData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DownlinkMessageDeliveryData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DownlinkMessageDeliveryData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetDuration

`func (o *DownlinkMessageDeliveryData) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DownlinkMessageDeliveryData) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DownlinkMessageDeliveryData) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DownlinkMessageDeliveryData) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGeoId

`func (o *DownlinkMessageDeliveryData) GetGeoId() string`

GetGeoId returns the GeoId field if non-nil, zero value otherwise.

### GetGeoIdOk

`func (o *DownlinkMessageDeliveryData) GetGeoIdOk() (*string, bool)`

GetGeoIdOk returns a tuple with the GeoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoId

`func (o *DownlinkMessageDeliveryData) SetGeoId(v string)`

SetGeoId sets GeoId field to given value.

### HasGeoId

`func (o *DownlinkMessageDeliveryData) HasGeoId() bool`

HasGeoId returns a boolean if a field has been set.

### GetPayload

`func (o *DownlinkMessageDeliveryData) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DownlinkMessageDeliveryData) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DownlinkMessageDeliveryData) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


