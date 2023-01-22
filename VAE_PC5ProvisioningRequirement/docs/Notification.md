# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Result** | [**Result**](Result.md) |  | 

## Methods

### NewNotification

`func NewNotification(resourceUri string, result Result, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceUri

`func (o *Notification) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *Notification) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *Notification) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.


### GetResult

`func (o *Notification) GetResult() Result`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Notification) GetResultOk() (*Result, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Notification) SetResult(v Result)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


