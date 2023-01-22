# NiddConfigurationStatusNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiddConfiguration** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**Status** | [**NiddStatus**](NiddStatus.md) |  | 
**RdsCapIndication** | Pointer to **bool** | It indicates whether the network capability for the reliable data service is enabled or not. | [optional] 
**RdsPort** | Pointer to [**RdsPort**](RdsPort.md) |  | [optional] 

## Methods

### NewNiddConfigurationStatusNotification

`func NewNiddConfigurationStatusNotification(niddConfiguration string, status NiddStatus, ) *NiddConfigurationStatusNotification`

NewNiddConfigurationStatusNotification instantiates a new NiddConfigurationStatusNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddConfigurationStatusNotificationWithDefaults

`func NewNiddConfigurationStatusNotificationWithDefaults() *NiddConfigurationStatusNotification`

NewNiddConfigurationStatusNotificationWithDefaults instantiates a new NiddConfigurationStatusNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiddConfiguration

`func (o *NiddConfigurationStatusNotification) GetNiddConfiguration() string`

GetNiddConfiguration returns the NiddConfiguration field if non-nil, zero value otherwise.

### GetNiddConfigurationOk

`func (o *NiddConfigurationStatusNotification) GetNiddConfigurationOk() (*string, bool)`

GetNiddConfigurationOk returns a tuple with the NiddConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddConfiguration

`func (o *NiddConfigurationStatusNotification) SetNiddConfiguration(v string)`

SetNiddConfiguration sets NiddConfiguration field to given value.


### GetExternalId

`func (o *NiddConfigurationStatusNotification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NiddConfigurationStatusNotification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NiddConfigurationStatusNotification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NiddConfigurationStatusNotification) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *NiddConfigurationStatusNotification) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *NiddConfigurationStatusNotification) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *NiddConfigurationStatusNotification) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *NiddConfigurationStatusNotification) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetStatus

`func (o *NiddConfigurationStatusNotification) GetStatus() NiddStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NiddConfigurationStatusNotification) GetStatusOk() (*NiddStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NiddConfigurationStatusNotification) SetStatus(v NiddStatus)`

SetStatus sets Status field to given value.


### GetRdsCapIndication

`func (o *NiddConfigurationStatusNotification) GetRdsCapIndication() bool`

GetRdsCapIndication returns the RdsCapIndication field if non-nil, zero value otherwise.

### GetRdsCapIndicationOk

`func (o *NiddConfigurationStatusNotification) GetRdsCapIndicationOk() (*bool, bool)`

GetRdsCapIndicationOk returns a tuple with the RdsCapIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsCapIndication

`func (o *NiddConfigurationStatusNotification) SetRdsCapIndication(v bool)`

SetRdsCapIndication sets RdsCapIndication field to given value.

### HasRdsCapIndication

`func (o *NiddConfigurationStatusNotification) HasRdsCapIndication() bool`

HasRdsCapIndication returns a boolean if a field has been set.

### GetRdsPort

`func (o *NiddConfigurationStatusNotification) GetRdsPort() RdsPort`

GetRdsPort returns the RdsPort field if non-nil, zero value otherwise.

### GetRdsPortOk

`func (o *NiddConfigurationStatusNotification) GetRdsPortOk() (*RdsPort, bool)`

GetRdsPortOk returns a tuple with the RdsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPort

`func (o *NiddConfigurationStatusNotification) SetRdsPort(v RdsPort)`

SetRdsPort sets RdsPort field to given value.

### HasRdsPort

`func (o *NiddConfigurationStatusNotification) HasRdsPort() bool`

HasRdsPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


