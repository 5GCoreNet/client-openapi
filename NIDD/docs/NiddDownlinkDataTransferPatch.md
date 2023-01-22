# NiddDownlinkDataTransferPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | [optional] 
**ReliableDataService** | Pointer to **bool** | Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false).  | [optional] 
**RdsPort** | Pointer to [**RdsPort**](RdsPort.md) |  | [optional] 
**MaximumLatency** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**Priority** | Pointer to **int32** | It is used to indicate the priority of the non-IP data packet relative to other non-IP data packets. | [optional] 
**PdnEstablishmentOption** | Pointer to [**PdnEstablishmentOptions**](PdnEstablishmentOptions.md) |  | [optional] 

## Methods

### NewNiddDownlinkDataTransferPatch

`func NewNiddDownlinkDataTransferPatch() *NiddDownlinkDataTransferPatch`

NewNiddDownlinkDataTransferPatch instantiates a new NiddDownlinkDataTransferPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddDownlinkDataTransferPatchWithDefaults

`func NewNiddDownlinkDataTransferPatchWithDefaults() *NiddDownlinkDataTransferPatch`

NewNiddDownlinkDataTransferPatchWithDefaults instantiates a new NiddDownlinkDataTransferPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NiddDownlinkDataTransferPatch) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NiddDownlinkDataTransferPatch) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NiddDownlinkDataTransferPatch) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *NiddDownlinkDataTransferPatch) HasData() bool`

HasData returns a boolean if a field has been set.

### GetReliableDataService

`func (o *NiddDownlinkDataTransferPatch) GetReliableDataService() bool`

GetReliableDataService returns the ReliableDataService field if non-nil, zero value otherwise.

### GetReliableDataServiceOk

`func (o *NiddDownlinkDataTransferPatch) GetReliableDataServiceOk() (*bool, bool)`

GetReliableDataServiceOk returns a tuple with the ReliableDataService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableDataService

`func (o *NiddDownlinkDataTransferPatch) SetReliableDataService(v bool)`

SetReliableDataService sets ReliableDataService field to given value.

### HasReliableDataService

`func (o *NiddDownlinkDataTransferPatch) HasReliableDataService() bool`

HasReliableDataService returns a boolean if a field has been set.

### GetRdsPort

`func (o *NiddDownlinkDataTransferPatch) GetRdsPort() RdsPort`

GetRdsPort returns the RdsPort field if non-nil, zero value otherwise.

### GetRdsPortOk

`func (o *NiddDownlinkDataTransferPatch) GetRdsPortOk() (*RdsPort, bool)`

GetRdsPortOk returns a tuple with the RdsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPort

`func (o *NiddDownlinkDataTransferPatch) SetRdsPort(v RdsPort)`

SetRdsPort sets RdsPort field to given value.

### HasRdsPort

`func (o *NiddDownlinkDataTransferPatch) HasRdsPort() bool`

HasRdsPort returns a boolean if a field has been set.

### GetMaximumLatency

`func (o *NiddDownlinkDataTransferPatch) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *NiddDownlinkDataTransferPatch) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *NiddDownlinkDataTransferPatch) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *NiddDownlinkDataTransferPatch) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### GetPriority

`func (o *NiddDownlinkDataTransferPatch) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NiddDownlinkDataTransferPatch) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NiddDownlinkDataTransferPatch) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NiddDownlinkDataTransferPatch) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPdnEstablishmentOption

`func (o *NiddDownlinkDataTransferPatch) GetPdnEstablishmentOption() PdnEstablishmentOptions`

GetPdnEstablishmentOption returns the PdnEstablishmentOption field if non-nil, zero value otherwise.

### GetPdnEstablishmentOptionOk

`func (o *NiddDownlinkDataTransferPatch) GetPdnEstablishmentOptionOk() (*PdnEstablishmentOptions, bool)`

GetPdnEstablishmentOptionOk returns a tuple with the PdnEstablishmentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnEstablishmentOption

`func (o *NiddDownlinkDataTransferPatch) SetPdnEstablishmentOption(v PdnEstablishmentOptions)`

SetPdnEstablishmentOption sets PdnEstablishmentOption field to given value.

### HasPdnEstablishmentOption

`func (o *NiddDownlinkDataTransferPatch) HasPdnEstablishmentOption() bool`

HasPdnEstablishmentOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


