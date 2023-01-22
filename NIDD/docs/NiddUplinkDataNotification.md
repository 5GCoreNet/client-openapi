# NiddUplinkDataNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiddConfiguration** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**Data** | **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | 
**ReliableDataService** | Pointer to **bool** | Indicates whether the reliable data service acknowledgement is requested (true) or not (false).  | [optional] 
**RdsPort** | Pointer to [**RdsPort**](RdsPort.md) |  | [optional] 

## Methods

### NewNiddUplinkDataNotification

`func NewNiddUplinkDataNotification(niddConfiguration string, data string, ) *NiddUplinkDataNotification`

NewNiddUplinkDataNotification instantiates a new NiddUplinkDataNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddUplinkDataNotificationWithDefaults

`func NewNiddUplinkDataNotificationWithDefaults() *NiddUplinkDataNotification`

NewNiddUplinkDataNotificationWithDefaults instantiates a new NiddUplinkDataNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiddConfiguration

`func (o *NiddUplinkDataNotification) GetNiddConfiguration() string`

GetNiddConfiguration returns the NiddConfiguration field if non-nil, zero value otherwise.

### GetNiddConfigurationOk

`func (o *NiddUplinkDataNotification) GetNiddConfigurationOk() (*string, bool)`

GetNiddConfigurationOk returns a tuple with the NiddConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddConfiguration

`func (o *NiddUplinkDataNotification) SetNiddConfiguration(v string)`

SetNiddConfiguration sets NiddConfiguration field to given value.


### GetExternalId

`func (o *NiddUplinkDataNotification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NiddUplinkDataNotification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NiddUplinkDataNotification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NiddUplinkDataNotification) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *NiddUplinkDataNotification) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *NiddUplinkDataNotification) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *NiddUplinkDataNotification) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *NiddUplinkDataNotification) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetData

`func (o *NiddUplinkDataNotification) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NiddUplinkDataNotification) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NiddUplinkDataNotification) SetData(v string)`

SetData sets Data field to given value.


### GetReliableDataService

`func (o *NiddUplinkDataNotification) GetReliableDataService() bool`

GetReliableDataService returns the ReliableDataService field if non-nil, zero value otherwise.

### GetReliableDataServiceOk

`func (o *NiddUplinkDataNotification) GetReliableDataServiceOk() (*bool, bool)`

GetReliableDataServiceOk returns a tuple with the ReliableDataService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableDataService

`func (o *NiddUplinkDataNotification) SetReliableDataService(v bool)`

SetReliableDataService sets ReliableDataService field to given value.

### HasReliableDataService

`func (o *NiddUplinkDataNotification) HasReliableDataService() bool`

HasReliableDataService returns a boolean if a field has been set.

### GetRdsPort

`func (o *NiddUplinkDataNotification) GetRdsPort() RdsPort`

GetRdsPort returns the RdsPort field if non-nil, zero value otherwise.

### GetRdsPortOk

`func (o *NiddUplinkDataNotification) GetRdsPortOk() (*RdsPort, bool)`

GetRdsPortOk returns a tuple with the RdsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPort

`func (o *NiddUplinkDataNotification) SetRdsPort(v RdsPort)`

SetRdsPort sets RdsPort field to given value.

### HasRdsPort

`func (o *NiddUplinkDataNotification) HasRdsPort() bool`

HasRdsPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


