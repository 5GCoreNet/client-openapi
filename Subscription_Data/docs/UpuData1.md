# UpuData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecPacket** | Pointer to **string** | Contains a secure packet. | [optional] 
**DefaultConfNssai** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**RoutingId** | Pointer to **string** | Represents a routing indicator. | [optional] 

## Methods

### NewUpuData1

`func NewUpuData1() *UpuData1`

NewUpuData1 instantiates a new UpuData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpuData1WithDefaults

`func NewUpuData1WithDefaults() *UpuData1`

NewUpuData1WithDefaults instantiates a new UpuData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecPacket

`func (o *UpuData1) GetSecPacket() string`

GetSecPacket returns the SecPacket field if non-nil, zero value otherwise.

### GetSecPacketOk

`func (o *UpuData1) GetSecPacketOk() (*string, bool)`

GetSecPacketOk returns a tuple with the SecPacket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecPacket

`func (o *UpuData1) SetSecPacket(v string)`

SetSecPacket sets SecPacket field to given value.

### HasSecPacket

`func (o *UpuData1) HasSecPacket() bool`

HasSecPacket returns a boolean if a field has been set.

### GetDefaultConfNssai

`func (o *UpuData1) GetDefaultConfNssai() []Snssai`

GetDefaultConfNssai returns the DefaultConfNssai field if non-nil, zero value otherwise.

### GetDefaultConfNssaiOk

`func (o *UpuData1) GetDefaultConfNssaiOk() (*[]Snssai, bool)`

GetDefaultConfNssaiOk returns a tuple with the DefaultConfNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConfNssai

`func (o *UpuData1) SetDefaultConfNssai(v []Snssai)`

SetDefaultConfNssai sets DefaultConfNssai field to given value.

### HasDefaultConfNssai

`func (o *UpuData1) HasDefaultConfNssai() bool`

HasDefaultConfNssai returns a boolean if a field has been set.

### GetRoutingId

`func (o *UpuData1) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### GetRoutingIdOk

`func (o *UpuData1) GetRoutingIdOk() (*string, bool)`

GetRoutingIdOk returns a tuple with the RoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingId

`func (o *UpuData1) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### HasRoutingId

`func (o *UpuData1) HasRoutingId() bool`

HasRoutingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


