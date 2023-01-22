# MbsSessionSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | Pointer to [**MbsSessionId**](MbsSessionId.md) |  | [optional] 
**AreaSessionId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | [optional] 
**EventList** | [**[]MbsSessionEvent**](MbsSessionEvent.md) |  | 
**NotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotifyCorrelationId** | Pointer to **string** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NfcInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**MbsSessionSubscUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewMbsSessionSubscription

`func NewMbsSessionSubscription(eventList []MbsSessionEvent, notifyUri string, ) *MbsSessionSubscription`

NewMbsSessionSubscription instantiates a new MbsSessionSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionSubscriptionWithDefaults

`func NewMbsSessionSubscriptionWithDefaults() *MbsSessionSubscription`

NewMbsSessionSubscriptionWithDefaults instantiates a new MbsSessionSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MbsSessionSubscription) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MbsSessionSubscription) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MbsSessionSubscription) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.

### HasMbsSessionId

`func (o *MbsSessionSubscription) HasMbsSessionId() bool`

HasMbsSessionId returns a boolean if a field has been set.

### GetAreaSessionId

`func (o *MbsSessionSubscription) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *MbsSessionSubscription) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *MbsSessionSubscription) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.

### HasAreaSessionId

`func (o *MbsSessionSubscription) HasAreaSessionId() bool`

HasAreaSessionId returns a boolean if a field has been set.

### GetEventList

`func (o *MbsSessionSubscription) GetEventList() []MbsSessionEvent`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *MbsSessionSubscription) GetEventListOk() (*[]MbsSessionEvent, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *MbsSessionSubscription) SetEventList(v []MbsSessionEvent)`

SetEventList sets EventList field to given value.


### GetNotifyUri

`func (o *MbsSessionSubscription) GetNotifyUri() string`

GetNotifyUri returns the NotifyUri field if non-nil, zero value otherwise.

### GetNotifyUriOk

`func (o *MbsSessionSubscription) GetNotifyUriOk() (*string, bool)`

GetNotifyUriOk returns a tuple with the NotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyUri

`func (o *MbsSessionSubscription) SetNotifyUri(v string)`

SetNotifyUri sets NotifyUri field to given value.


### GetNotifyCorrelationId

`func (o *MbsSessionSubscription) GetNotifyCorrelationId() string`

GetNotifyCorrelationId returns the NotifyCorrelationId field if non-nil, zero value otherwise.

### GetNotifyCorrelationIdOk

`func (o *MbsSessionSubscription) GetNotifyCorrelationIdOk() (*string, bool)`

GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyCorrelationId

`func (o *MbsSessionSubscription) SetNotifyCorrelationId(v string)`

SetNotifyCorrelationId sets NotifyCorrelationId field to given value.

### HasNotifyCorrelationId

`func (o *MbsSessionSubscription) HasNotifyCorrelationId() bool`

HasNotifyCorrelationId returns a boolean if a field has been set.

### GetExpiryTime

`func (o *MbsSessionSubscription) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *MbsSessionSubscription) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *MbsSessionSubscription) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *MbsSessionSubscription) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetNfcInstanceId

`func (o *MbsSessionSubscription) GetNfcInstanceId() string`

GetNfcInstanceId returns the NfcInstanceId field if non-nil, zero value otherwise.

### GetNfcInstanceIdOk

`func (o *MbsSessionSubscription) GetNfcInstanceIdOk() (*string, bool)`

GetNfcInstanceIdOk returns a tuple with the NfcInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfcInstanceId

`func (o *MbsSessionSubscription) SetNfcInstanceId(v string)`

SetNfcInstanceId sets NfcInstanceId field to given value.

### HasNfcInstanceId

`func (o *MbsSessionSubscription) HasNfcInstanceId() bool`

HasNfcInstanceId returns a boolean if a field has been set.

### GetMbsSessionSubscUri

`func (o *MbsSessionSubscription) GetMbsSessionSubscUri() string`

GetMbsSessionSubscUri returns the MbsSessionSubscUri field if non-nil, zero value otherwise.

### GetMbsSessionSubscUriOk

`func (o *MbsSessionSubscription) GetMbsSessionSubscUriOk() (*string, bool)`

GetMbsSessionSubscUriOk returns a tuple with the MbsSessionSubscUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionSubscUri

`func (o *MbsSessionSubscription) SetMbsSessionSubscUri(v string)`

SetMbsSessionSubscUri sets MbsSessionSubscUri field to given value.

### HasMbsSessionSubscUri

`func (o *MbsSessionSubscription) HasMbsSessionSubscUri() bool`

HasMbsSessionSubscUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


