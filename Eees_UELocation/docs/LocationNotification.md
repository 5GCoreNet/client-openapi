# LocationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubId** | **string** | Identifier of the location information subscription for which the location information notification is related to.  | 
**LocEvs** | [**[]LocationEvent**](LocationEvent.md) | List of notifications with location information. | 

## Methods

### NewLocationNotification

`func NewLocationNotification(subId string, locEvs []LocationEvent, ) *LocationNotification`

NewLocationNotification instantiates a new LocationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationNotificationWithDefaults

`func NewLocationNotificationWithDefaults() *LocationNotification`

NewLocationNotificationWithDefaults instantiates a new LocationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubId

`func (o *LocationNotification) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *LocationNotification) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *LocationNotification) SetSubId(v string)`

SetSubId sets SubId field to given value.


### GetLocEvs

`func (o *LocationNotification) GetLocEvs() []LocationEvent`

GetLocEvs returns the LocEvs field if non-nil, zero value otherwise.

### GetLocEvsOk

`func (o *LocationNotification) GetLocEvsOk() (*[]LocationEvent, bool)`

GetLocEvsOk returns a tuple with the LocEvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocEvs

`func (o *LocationNotification) SetLocEvs(v []LocationEvent)`

SetLocEvs sets LocEvs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


