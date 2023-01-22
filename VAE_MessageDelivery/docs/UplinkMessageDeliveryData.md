# UplinkMessageDeliveryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**UeId** | **string** | Represents the identifier of the V2X UE. | 
**GeoId** | Pointer to **string** | Represents a geographical area identifier. | [optional] 
**Payload** | **string** | string with format &#39;bytes&#39; as defined in OpenAPI | 

## Methods

### NewUplinkMessageDeliveryData

`func NewUplinkMessageDeliveryData(resourceUri string, ueId string, payload string, ) *UplinkMessageDeliveryData`

NewUplinkMessageDeliveryData instantiates a new UplinkMessageDeliveryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplinkMessageDeliveryDataWithDefaults

`func NewUplinkMessageDeliveryDataWithDefaults() *UplinkMessageDeliveryData`

NewUplinkMessageDeliveryDataWithDefaults instantiates a new UplinkMessageDeliveryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *UplinkMessageDeliveryData) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *UplinkMessageDeliveryData) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *UplinkMessageDeliveryData) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetUeId

`func (o *UplinkMessageDeliveryData) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *UplinkMessageDeliveryData) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *UplinkMessageDeliveryData) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetGeoId

`func (o *UplinkMessageDeliveryData) GetGeoId() string`

GetGeoId returns the GeoId field if non-nil, zero value otherwise.

### GetGeoIdOk

`func (o *UplinkMessageDeliveryData) GetGeoIdOk() (*string, bool)`

GetGeoIdOk returns a tuple with the GeoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoId

`func (o *UplinkMessageDeliveryData) SetGeoId(v string)`

SetGeoId sets GeoId field to given value.

### HasGeoId

`func (o *UplinkMessageDeliveryData) HasGeoId() bool`

HasGeoId returns a boolean if a field has been set.

### GetPayload

`func (o *UplinkMessageDeliveryData) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *UplinkMessageDeliveryData) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *UplinkMessageDeliveryData) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


