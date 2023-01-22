# HdMapDynamicInfoNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NearbyUeInfo** | [**[]NearbyUeInfo**](NearbyUeInfo.md) | Contains the informaiotn of nearby UEs. | 

## Methods

### NewHdMapDynamicInfoNotification

`func NewHdMapDynamicInfoNotification(resourceUri string, nearbyUeInfo []NearbyUeInfo, ) *HdMapDynamicInfoNotification`

NewHdMapDynamicInfoNotification instantiates a new HdMapDynamicInfoNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHdMapDynamicInfoNotificationWithDefaults

`func NewHdMapDynamicInfoNotificationWithDefaults() *HdMapDynamicInfoNotification`

NewHdMapDynamicInfoNotificationWithDefaults instantiates a new HdMapDynamicInfoNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *HdMapDynamicInfoNotification) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *HdMapDynamicInfoNotification) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *HdMapDynamicInfoNotification) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetNearbyUeInfo

`func (o *HdMapDynamicInfoNotification) GetNearbyUeInfo() []NearbyUeInfo`

GetNearbyUeInfo returns the NearbyUeInfo field if non-nil, zero value otherwise.

### GetNearbyUeInfoOk

`func (o *HdMapDynamicInfoNotification) GetNearbyUeInfoOk() (*[]NearbyUeInfo, bool)`

GetNearbyUeInfoOk returns a tuple with the NearbyUeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNearbyUeInfo

`func (o *HdMapDynamicInfoNotification) SetNearbyUeInfo(v []NearbyUeInfo)`

SetNearbyUeInfo sets NearbyUeInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


