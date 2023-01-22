# ConnectionRequestType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Producer** | Pointer to [**ProducerIdType**](ProducerIdType.md) |  | [optional] 
**Streams** | Pointer to [**[]StreamInfoType**](StreamInfoType.md) |  | [optional] 

## Methods

### NewConnectionRequestType

`func NewConnectionRequestType() *ConnectionRequestType`

NewConnectionRequestType instantiates a new ConnectionRequestType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRequestTypeWithDefaults

`func NewConnectionRequestTypeWithDefaults() *ConnectionRequestType`

NewConnectionRequestTypeWithDefaults instantiates a new ConnectionRequestType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducer

`func (o *ConnectionRequestType) GetProducer() ProducerIdType`

GetProducer returns the Producer field if non-nil, zero value otherwise.

### GetProducerOk

`func (o *ConnectionRequestType) GetProducerOk() (*ProducerIdType, bool)`

GetProducerOk returns a tuple with the Producer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducer

`func (o *ConnectionRequestType) SetProducer(v ProducerIdType)`

SetProducer sets Producer field to given value.

### HasProducer

`func (o *ConnectionRequestType) HasProducer() bool`

HasProducer returns a boolean if a field has been set.

### GetStreams

`func (o *ConnectionRequestType) GetStreams() []StreamInfoType`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *ConnectionRequestType) GetStreamsOk() (*[]StreamInfoType, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *ConnectionRequestType) SetStreams(v []StreamInfoType)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *ConnectionRequestType) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


