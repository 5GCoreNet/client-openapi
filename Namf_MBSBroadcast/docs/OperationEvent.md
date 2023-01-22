# OperationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpEventType** | [**OpEventType**](OpEventType.md) |  | 
**AmfId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NgranFailureEventList** | Pointer to [**[]NgranFailureEvent**](NgranFailureEvent.md) |  | [optional] 

## Methods

### NewOperationEvent

`func NewOperationEvent(opEventType OpEventType, ) *OperationEvent`

NewOperationEvent instantiates a new OperationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationEventWithDefaults

`func NewOperationEventWithDefaults() *OperationEvent`

NewOperationEventWithDefaults instantiates a new OperationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpEventType

`func (o *OperationEvent) GetOpEventType() OpEventType`

GetOpEventType returns the OpEventType field if non-nil, zero value otherwise.

### GetOpEventTypeOk

`func (o *OperationEvent) GetOpEventTypeOk() (*OpEventType, bool)`

GetOpEventTypeOk returns a tuple with the OpEventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpEventType

`func (o *OperationEvent) SetOpEventType(v OpEventType)`

SetOpEventType sets OpEventType field to given value.


### GetAmfId

`func (o *OperationEvent) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *OperationEvent) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *OperationEvent) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.

### HasAmfId

`func (o *OperationEvent) HasAmfId() bool`

HasAmfId returns a boolean if a field has been set.

### GetNgranFailureEventList

`func (o *OperationEvent) GetNgranFailureEventList() []NgranFailureEvent`

GetNgranFailureEventList returns the NgranFailureEventList field if non-nil, zero value otherwise.

### GetNgranFailureEventListOk

`func (o *OperationEvent) GetNgranFailureEventListOk() (*[]NgranFailureEvent, bool)`

GetNgranFailureEventListOk returns a tuple with the NgranFailureEventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNgranFailureEventList

`func (o *OperationEvent) SetNgranFailureEventList(v []NgranFailureEvent)`

SetNgranFailureEventList sets NgranFailureEventList field to given value.

### HasNgranFailureEventList

`func (o *OperationEvent) HasNgranFailureEventList() bool`

HasNgranFailureEventList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


