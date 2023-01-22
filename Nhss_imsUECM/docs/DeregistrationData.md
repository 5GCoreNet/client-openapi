# DeregistrationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeregReason** | [**DeregistrationReason**](DeregistrationReason.md) |  | 
**Impi** | **string** | IMS Private Identity of the UE | 
**AssociatedImpis** | Pointer to **[]string** |  | [optional] 
**EmergencyRegisteredIdentities** | Pointer to [**[]EmergencyRegisteredIdentity**](EmergencyRegisteredIdentity.md) |  | [optional] 

## Methods

### NewDeregistrationData

`func NewDeregistrationData(deregReason DeregistrationReason, impi string, ) *DeregistrationData`

NewDeregistrationData instantiates a new DeregistrationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeregistrationDataWithDefaults

`func NewDeregistrationDataWithDefaults() *DeregistrationData`

NewDeregistrationDataWithDefaults instantiates a new DeregistrationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeregReason

`func (o *DeregistrationData) GetDeregReason() DeregistrationReason`

GetDeregReason returns the DeregReason field if non-nil, zero value otherwise.

### GetDeregReasonOk

`func (o *DeregistrationData) GetDeregReasonOk() (*DeregistrationReason, bool)`

GetDeregReasonOk returns a tuple with the DeregReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeregReason

`func (o *DeregistrationData) SetDeregReason(v DeregistrationReason)`

SetDeregReason sets DeregReason field to given value.


### GetImpi

`func (o *DeregistrationData) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *DeregistrationData) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *DeregistrationData) SetImpi(v string)`

SetImpi sets Impi field to given value.


### GetAssociatedImpis

`func (o *DeregistrationData) GetAssociatedImpis() []string`

GetAssociatedImpis returns the AssociatedImpis field if non-nil, zero value otherwise.

### GetAssociatedImpisOk

`func (o *DeregistrationData) GetAssociatedImpisOk() (*[]string, bool)`

GetAssociatedImpisOk returns a tuple with the AssociatedImpis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedImpis

`func (o *DeregistrationData) SetAssociatedImpis(v []string)`

SetAssociatedImpis sets AssociatedImpis field to given value.

### HasAssociatedImpis

`func (o *DeregistrationData) HasAssociatedImpis() bool`

HasAssociatedImpis returns a boolean if a field has been set.

### GetEmergencyRegisteredIdentities

`func (o *DeregistrationData) GetEmergencyRegisteredIdentities() []EmergencyRegisteredIdentity`

GetEmergencyRegisteredIdentities returns the EmergencyRegisteredIdentities field if non-nil, zero value otherwise.

### GetEmergencyRegisteredIdentitiesOk

`func (o *DeregistrationData) GetEmergencyRegisteredIdentitiesOk() (*[]EmergencyRegisteredIdentity, bool)`

GetEmergencyRegisteredIdentitiesOk returns a tuple with the EmergencyRegisteredIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyRegisteredIdentities

`func (o *DeregistrationData) SetEmergencyRegisteredIdentities(v []EmergencyRegisteredIdentity)`

SetEmergencyRegisteredIdentities sets EmergencyRegisteredIdentities field to given value.

### HasEmergencyRegisteredIdentities

`func (o *DeregistrationData) HasEmergencyRegisteredIdentities() bool`

HasEmergencyRegisteredIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


