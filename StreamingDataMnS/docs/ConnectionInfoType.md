# ConnectionInfoType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to **string** | Resource URI | [optional] 
**Producer** | Pointer to [**ProducerIdType**](ProducerIdType.md) |  | [optional] 
**Streams** | Pointer to **[]string** |  | [optional] 

## Methods

### NewConnectionInfoType

`func NewConnectionInfoType() *ConnectionInfoType`

NewConnectionInfoType instantiates a new ConnectionInfoType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionInfoTypeWithDefaults

`func NewConnectionInfoTypeWithDefaults() *ConnectionInfoType`

NewConnectionInfoTypeWithDefaults instantiates a new ConnectionInfoType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *ConnectionInfoType) GetConnection() string`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ConnectionInfoType) GetConnectionOk() (*string, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ConnectionInfoType) SetConnection(v string)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ConnectionInfoType) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetProducer

`func (o *ConnectionInfoType) GetProducer() ProducerIdType`

GetProducer returns the Producer field if non-nil, zero value otherwise.

### GetProducerOk

`func (o *ConnectionInfoType) GetProducerOk() (*ProducerIdType, bool)`

GetProducerOk returns a tuple with the Producer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducer

`func (o *ConnectionInfoType) SetProducer(v ProducerIdType)`

SetProducer sets Producer field to given value.

### HasProducer

`func (o *ConnectionInfoType) HasProducer() bool`

HasProducer returns a boolean if a field has been set.

### GetStreams

`func (o *ConnectionInfoType) GetStreams() []string`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *ConnectionInfoType) GetStreamsOk() (*[]string, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *ConnectionInfoType) SetStreams(v []string)`

SetStreams sets Streams field to given value.

### HasStreams

`func (o *ConnectionInfoType) HasStreams() bool`

HasStreams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


