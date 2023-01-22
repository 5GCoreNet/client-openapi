# RdsDownlinkDataDeliveryFailureAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**SupportedUeFormats** | Pointer to [**[]SerializationFormat**](SerializationFormat.md) | Indicates the serialization format(s) that are supported by the UE on the associated RDS port. | [optional] 

## Methods

### NewRdsDownlinkDataDeliveryFailureAllOf

`func NewRdsDownlinkDataDeliveryFailureAllOf() *RdsDownlinkDataDeliveryFailureAllOf`

NewRdsDownlinkDataDeliveryFailureAllOf instantiates a new RdsDownlinkDataDeliveryFailureAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsDownlinkDataDeliveryFailureAllOfWithDefaults

`func NewRdsDownlinkDataDeliveryFailureAllOfWithDefaults() *RdsDownlinkDataDeliveryFailureAllOf`

NewRdsDownlinkDataDeliveryFailureAllOfWithDefaults instantiates a new RdsDownlinkDataDeliveryFailureAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedRetransmissionTime

`func (o *RdsDownlinkDataDeliveryFailureAllOf) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *RdsDownlinkDataDeliveryFailureAllOf) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *RdsDownlinkDataDeliveryFailureAllOf) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *RdsDownlinkDataDeliveryFailureAllOf) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.

### GetSupportedUeFormats

`func (o *RdsDownlinkDataDeliveryFailureAllOf) GetSupportedUeFormats() []SerializationFormat`

GetSupportedUeFormats returns the SupportedUeFormats field if non-nil, zero value otherwise.

### GetSupportedUeFormatsOk

`func (o *RdsDownlinkDataDeliveryFailureAllOf) GetSupportedUeFormatsOk() (*[]SerializationFormat, bool)`

GetSupportedUeFormatsOk returns a tuple with the SupportedUeFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedUeFormats

`func (o *RdsDownlinkDataDeliveryFailureAllOf) SetSupportedUeFormats(v []SerializationFormat)`

SetSupportedUeFormats sets SupportedUeFormats field to given value.

### HasSupportedUeFormats

`func (o *RdsDownlinkDataDeliveryFailureAllOf) HasSupportedUeFormats() bool`

HasSupportedUeFormats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


