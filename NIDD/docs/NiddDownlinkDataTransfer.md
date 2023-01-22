# NiddDownlinkDataTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**Data** | **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | 
**ReliableDataService** | Pointer to **bool** | Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false). Default value is false.  | [optional] 
**RdsPort** | Pointer to [**RdsPort**](RdsPort.md) |  | [optional] 
**MaximumLatency** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**Priority** | Pointer to **int32** | It is used to indicate the priority of the non-IP data packet relative to other non-IP data packets. | [optional] 
**PdnEstablishmentOption** | Pointer to [**PdnEstablishmentOptions**](PdnEstablishmentOptions.md) |  | [optional] 
**DeliveryStatus** | Pointer to [**DeliveryStatus**](DeliveryStatus.md) |  | [optional] 
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 

## Methods

### NewNiddDownlinkDataTransfer

`func NewNiddDownlinkDataTransfer(data string, ) *NiddDownlinkDataTransfer`

NewNiddDownlinkDataTransfer instantiates a new NiddDownlinkDataTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddDownlinkDataTransferWithDefaults

`func NewNiddDownlinkDataTransferWithDefaults() *NiddDownlinkDataTransfer`

NewNiddDownlinkDataTransferWithDefaults instantiates a new NiddDownlinkDataTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *NiddDownlinkDataTransfer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NiddDownlinkDataTransfer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NiddDownlinkDataTransfer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NiddDownlinkDataTransfer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *NiddDownlinkDataTransfer) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *NiddDownlinkDataTransfer) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *NiddDownlinkDataTransfer) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *NiddDownlinkDataTransfer) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetMsisdn

`func (o *NiddDownlinkDataTransfer) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *NiddDownlinkDataTransfer) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *NiddDownlinkDataTransfer) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *NiddDownlinkDataTransfer) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetSelf

`func (o *NiddDownlinkDataTransfer) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NiddDownlinkDataTransfer) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NiddDownlinkDataTransfer) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *NiddDownlinkDataTransfer) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetData

`func (o *NiddDownlinkDataTransfer) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NiddDownlinkDataTransfer) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NiddDownlinkDataTransfer) SetData(v string)`

SetData sets Data field to given value.


### GetReliableDataService

`func (o *NiddDownlinkDataTransfer) GetReliableDataService() bool`

GetReliableDataService returns the ReliableDataService field if non-nil, zero value otherwise.

### GetReliableDataServiceOk

`func (o *NiddDownlinkDataTransfer) GetReliableDataServiceOk() (*bool, bool)`

GetReliableDataServiceOk returns a tuple with the ReliableDataService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableDataService

`func (o *NiddDownlinkDataTransfer) SetReliableDataService(v bool)`

SetReliableDataService sets ReliableDataService field to given value.

### HasReliableDataService

`func (o *NiddDownlinkDataTransfer) HasReliableDataService() bool`

HasReliableDataService returns a boolean if a field has been set.

### GetRdsPort

`func (o *NiddDownlinkDataTransfer) GetRdsPort() RdsPort`

GetRdsPort returns the RdsPort field if non-nil, zero value otherwise.

### GetRdsPortOk

`func (o *NiddDownlinkDataTransfer) GetRdsPortOk() (*RdsPort, bool)`

GetRdsPortOk returns a tuple with the RdsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPort

`func (o *NiddDownlinkDataTransfer) SetRdsPort(v RdsPort)`

SetRdsPort sets RdsPort field to given value.

### HasRdsPort

`func (o *NiddDownlinkDataTransfer) HasRdsPort() bool`

HasRdsPort returns a boolean if a field has been set.

### GetMaximumLatency

`func (o *NiddDownlinkDataTransfer) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *NiddDownlinkDataTransfer) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *NiddDownlinkDataTransfer) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *NiddDownlinkDataTransfer) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### GetPriority

`func (o *NiddDownlinkDataTransfer) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NiddDownlinkDataTransfer) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NiddDownlinkDataTransfer) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NiddDownlinkDataTransfer) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPdnEstablishmentOption

`func (o *NiddDownlinkDataTransfer) GetPdnEstablishmentOption() PdnEstablishmentOptions`

GetPdnEstablishmentOption returns the PdnEstablishmentOption field if non-nil, zero value otherwise.

### GetPdnEstablishmentOptionOk

`func (o *NiddDownlinkDataTransfer) GetPdnEstablishmentOptionOk() (*PdnEstablishmentOptions, bool)`

GetPdnEstablishmentOptionOk returns a tuple with the PdnEstablishmentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnEstablishmentOption

`func (o *NiddDownlinkDataTransfer) SetPdnEstablishmentOption(v PdnEstablishmentOptions)`

SetPdnEstablishmentOption sets PdnEstablishmentOption field to given value.

### HasPdnEstablishmentOption

`func (o *NiddDownlinkDataTransfer) HasPdnEstablishmentOption() bool`

HasPdnEstablishmentOption returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *NiddDownlinkDataTransfer) GetDeliveryStatus() DeliveryStatus`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *NiddDownlinkDataTransfer) GetDeliveryStatusOk() (*DeliveryStatus, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *NiddDownlinkDataTransfer) SetDeliveryStatus(v DeliveryStatus)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *NiddDownlinkDataTransfer) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetRequestedRetransmissionTime

`func (o *NiddDownlinkDataTransfer) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *NiddDownlinkDataTransfer) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *NiddDownlinkDataTransfer) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *NiddDownlinkDataTransfer) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


