# ContextStatusSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfcInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**EventList** | [**[]ContextStatusEvent**](ContextStatusEvent.md) |  | 
**NotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewContextStatusSubscription

`func NewContextStatusSubscription(nfcInstanceId string, mbsSessionId MbsSessionId, eventList []ContextStatusEvent, notifyUri string, ) *ContextStatusSubscription`

NewContextStatusSubscription instantiates a new ContextStatusSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextStatusSubscriptionWithDefaults

`func NewContextStatusSubscriptionWithDefaults() *ContextStatusSubscription`

NewContextStatusSubscriptionWithDefaults instantiates a new ContextStatusSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfcInstanceId

`func (o *ContextStatusSubscription) GetNfcInstanceId() string`

GetNfcInstanceId returns the NfcInstanceId field if non-nil, zero value otherwise.

### GetNfcInstanceIdOk

`func (o *ContextStatusSubscription) GetNfcInstanceIdOk() (*string, bool)`

GetNfcInstanceIdOk returns a tuple with the NfcInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfcInstanceId

`func (o *ContextStatusSubscription) SetNfcInstanceId(v string)`

SetNfcInstanceId sets NfcInstanceId field to given value.


### GetMbsSessionId

`func (o *ContextStatusSubscription) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *ContextStatusSubscription) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *ContextStatusSubscription) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetEventList

`func (o *ContextStatusSubscription) GetEventList() []ContextStatusEvent`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *ContextStatusSubscription) GetEventListOk() (*[]ContextStatusEvent, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *ContextStatusSubscription) SetEventList(v []ContextStatusEvent)`

SetEventList sets EventList field to given value.


### GetNotifyUri

`func (o *ContextStatusSubscription) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *ContextStatusSubscription) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *ContextStatusSubscription) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.


### GetNotifyCorrelationId

`func (o *ContextStatusSubscription) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *ContextStatusSubscription) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *ContextStatusSubscription) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *ContextStatusSubscription) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ContextStatusSubscription) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ContextStatusSubscription) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ContextStatusSubscription) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ContextStatusSubscription) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


