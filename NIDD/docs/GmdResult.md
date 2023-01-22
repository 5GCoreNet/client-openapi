# GmdResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**DeliveryStatus** | [**DeliveryStatus**](DeliveryStatus.md) |  | 
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewGmdResult

`func NewGmdResult(deliveryStatus DeliveryStatus, ) *GmdResult`

NewGmdResult instantiates a new GmdResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGmdResultWithDefaults

`func NewGmdResultWithDefaults() *GmdResult`

NewGmdResultWithDefaults instantiates a new GmdResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *GmdResult) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GmdResult) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GmdResult) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GmdResult) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *GmdResult) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *GmdResult) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *GmdResult) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *GmdResult) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *GmdResult) GetDeliveryStatus() DeliveryStatus`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *GmdResult) GetDeliveryStatusOk() (*DeliveryStatus, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *GmdResult) SetDeliveryStatus(v DeliveryStatus)`

SetDeliveryStatus sets DeliveryStatus field to given value.


### GetRequestedRetransmissionTime

`func (o *GmdResult) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *GmdResult) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *GmdResult) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *GmdResult) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


