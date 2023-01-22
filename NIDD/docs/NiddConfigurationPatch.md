# NiddConfigurationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **NullableTime** | string with format \&quot;date-time\&quot; as defined in OpenAPI with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**ReliableDataService** | Pointer to **NullableBool** | Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false).  | [optional] 
**RdsPorts** | Pointer to [**[]RdsPort**](RdsPort.md) | Indicates the static port configuration that is used for reliable data transfer between specific applications using RDS (as defined in clause 5.2.4 and 5.2.5 of 3GPP TS 24.250). | [optional] 
**PdnEstablishmentOption** | Pointer to [**PdnEstablishmentOptionsRm**](PdnEstablishmentOptionsRm.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewNiddConfigurationPatch

`func NewNiddConfigurationPatch() *NiddConfigurationPatch`

NewNiddConfigurationPatch instantiates a new NiddConfigurationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddConfigurationPatchWithDefaults

`func NewNiddConfigurationPatchWithDefaults() *NiddConfigurationPatch`

NewNiddConfigurationPatchWithDefaults instantiates a new NiddConfigurationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *NiddConfigurationPatch) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NiddConfigurationPatch) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NiddConfigurationPatch) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NiddConfigurationPatch) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *NiddConfigurationPatch) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *NiddConfigurationPatch) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetReliableDataService

`func (o *NiddConfigurationPatch) GetReliableDataService() bool`

GetReliableDataService returns the ReliableDataService field if non-nil, zero value otherwise.

### GetReliableDataServiceOk

`func (o *NiddConfigurationPatch) GetReliableDataServiceOk() (*bool, bool)`

GetReliableDataServiceOk returns a tuple with the ReliableDataService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableDataService

`func (o *NiddConfigurationPatch) SetReliableDataService(v bool)`

SetReliableDataService sets ReliableDataService field to given value.

### HasReliableDataService

`func (o *NiddConfigurationPatch) HasReliableDataService() bool`

HasReliableDataService returns a boolean if a field has been set.

### SetReliableDataServiceNil

`func (o *NiddConfigurationPatch) SetReliableDataServiceNil(b bool)`

 SetReliableDataServiceNil sets the value for ReliableDataService to be an explicit nil

### UnsetReliableDataService
`func (o *NiddConfigurationPatch) UnsetReliableDataService()`

UnsetReliableDataService ensures that no value is present for ReliableDataService, not even an explicit nil
### GetRdsPorts

`func (o *NiddConfigurationPatch) GetRdsPorts() []RdsPort`

GetRdsPorts returns the RdsPorts field if non-nil, zero value otherwise.

### GetRdsPortsOk

`func (o *NiddConfigurationPatch) GetRdsPortsOk() (*[]RdsPort, bool)`

GetRdsPortsOk returns a tuple with the RdsPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPorts

`func (o *NiddConfigurationPatch) SetRdsPorts(v []RdsPort)`

SetRdsPorts sets RdsPorts field to given value.

### HasRdsPorts

`func (o *NiddConfigurationPatch) HasRdsPorts() bool`

HasRdsPorts returns a boolean if a field has been set.

### GetPdnEstablishmentOption

`func (o *NiddConfigurationPatch) GetPdnEstablishmentOption() PdnEstablishmentOptionsRm`

GetPdnEstablishmentOption returns the PdnEstablishmentOption field if non-nil, zero value otherwise.

### GetPdnEstablishmentOptionOk

`func (o *NiddConfigurationPatch) GetPdnEstablishmentOptionOk() (*PdnEstablishmentOptionsRm, bool)`

GetPdnEstablishmentOptionOk returns a tuple with the PdnEstablishmentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnEstablishmentOption

`func (o *NiddConfigurationPatch) SetPdnEstablishmentOption(v PdnEstablishmentOptionsRm)`

SetPdnEstablishmentOption sets PdnEstablishmentOption field to given value.

### HasPdnEstablishmentOption

`func (o *NiddConfigurationPatch) HasPdnEstablishmentOption() bool`

HasPdnEstablishmentOption returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *NiddConfigurationPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *NiddConfigurationPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *NiddConfigurationPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *NiddConfigurationPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


