# ManagePortNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NiddConfiguration** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**ManagedPorts** | Pointer to [**[]ManagePort**](ManagePort.md) | Indicates the reserved RDS port configuration information. | [optional] 

## Methods

### NewManagePortNotification

`func NewManagePortNotification(niddConfiguration string, ) *ManagePortNotification`

NewManagePortNotification instantiates a new ManagePortNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagePortNotificationWithDefaults

`func NewManagePortNotificationWithDefaults() *ManagePortNotification`

NewManagePortNotificationWithDefaults instantiates a new ManagePortNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNiddConfiguration

`func (o *ManagePortNotification) GetNiddConfiguration() string`

GetNiddConfiguration returns the NiddConfiguration field if non-nil, zero value otherwise.

### GetNiddConfigurationOk

`func (o *ManagePortNotification) GetNiddConfigurationOk() (*string, bool)`

GetNiddConfigurationOk returns a tuple with the NiddConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddConfiguration

`func (o *ManagePortNotification) SetNiddConfiguration(v string)`

SetNiddConfiguration sets NiddConfiguration field to given value.


### GetExternalId

`func (o *ManagePortNotification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ManagePortNotification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ManagePortNotification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ManagePortNotification) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *ManagePortNotification) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *ManagePortNotification) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *ManagePortNotification) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *ManagePortNotification) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetManagedPorts

`func (o *ManagePortNotification) GetManagedPorts() []ManagePort`

GetManagedPorts returns the ManagedPorts field if non-nil, zero value otherwise.

### GetManagedPortsOk

`func (o *ManagePortNotification) GetManagedPortsOk() (*[]ManagePort, bool)`

GetManagedPortsOk returns a tuple with the ManagedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedPorts

`func (o *ManagePortNotification) SetManagedPorts(v []ManagePort)`

SetManagedPorts sets ManagedPorts field to given value.

### HasManagedPorts

`func (o *ManagePortNotification) HasManagedPorts() bool`

HasManagedPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


