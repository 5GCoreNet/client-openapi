# AppReqNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Result** | [**ReservationResult**](ReservationResult.md) |  | 

## Methods

### NewAppReqNotification

`func NewAppReqNotification(resourceUri string, result ReservationResult, ) *AppReqNotification`

NewAppReqNotification instantiates a new AppReqNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppReqNotificationWithDefaults

`func NewAppReqNotificationWithDefaults() *AppReqNotification`

NewAppReqNotificationWithDefaults instantiates a new AppReqNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *AppReqNotification) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *AppReqNotification) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *AppReqNotification) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetResult

`func (o *AppReqNotification) GetResult() ReservationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *AppReqNotification) GetResultOk() (*ReservationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *AppReqNotification) SetResult(v ReservationResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


