# DistSessionSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfcInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**EventList** | [**[]DistSessionEventType**](DistSessionEventType.md) |  | 
**NotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**DistSessionSubscUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewDistSessionSubscription

`func NewDistSessionSubscription(eventList []DistSessionEventType, notifyUri string, ) *DistSessionSubscription`

NewDistSessionSubscription instantiates a new DistSessionSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistSessionSubscriptionWithDefaults

`func NewDistSessionSubscriptionWithDefaults() *DistSessionSubscription`

NewDistSessionSubscriptionWithDefaults instantiates a new DistSessionSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfcInstanceId

`func (o *DistSessionSubscription) GetNfcInstanceId() string`

GetNfcInstanceId returns the NfcInstanceId field if non-nil, zero value otherwise.

### GetNfcInstanceIdOk

`func (o *DistSessionSubscription) GetNfcInstanceIdOk() (*string, bool)`

GetNfcInstanceIdOk returns a tuple with the NfcInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfcInstanceId

`func (o *DistSessionSubscription) SetNfcInstanceId(v string)`

SetNfcInstanceId sets NfcInstanceId field to given value.

### HasNfcInstanceId

`func (o *DistSessionSubscription) HasNfcInstanceId() bool`

HasNfcInstanceId returns a boolean if a field has been set.

### GetEventList

`func (o *DistSessionSubscription) GetEventList() []DistSessionEventType`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *DistSessionSubscription) GetEventListOk() (*[]DistSessionEventType, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *DistSessionSubscription) SetEventList(v []DistSessionEventType)`

SetEventList sets EventList field to given value.


### GetNotifyUri

`func (o *DistSessionSubscription) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *DistSessionSubscription) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *DistSessionSubscription) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.


### GetNotifyCorrelationId

`func (o *DistSessionSubscription) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *DistSessionSubscription) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *DistSessionSubscription) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *DistSessionSubscription) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetExpiryTime

`func (o *DistSessionSubscription) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *DistSessionSubscription) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *DistSessionSubscription) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *DistSessionSubscription) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetDistSessionSubscUri

`func (o *DistSessionSubscription) GetDistSessionSubscUri() string`

GetDistSessionSubscUri returns the DistSessionSubscUri field if non-nil, zero value otherwise.

### GetDistSessionSubscUriOk

`func (o *DistSessionSubscription) GetDistSessionSubscUriOk() (*string, bool)`

GetDistSessionSubscUriOk returns a tuple with the DistSessionSubscUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistSessionSubscUri

`func (o *DistSessionSubscription) SetDistSessionSubscUri(v string)`

SetDistSessionSubscUri sets DistSessionSubscUri field to given value.

### HasDistSessionSubscUri

`func (o *DistSessionSubscription) HasDistSessionSubscUri() bool`

HasDistSessionSubscUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


