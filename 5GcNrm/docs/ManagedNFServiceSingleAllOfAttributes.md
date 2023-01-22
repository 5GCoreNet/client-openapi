# ManagedNFServiceSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserLabel** | Pointer to **string** |  | [optional] 
**NFServiceType** | Pointer to [**NFServiceType**](NFServiceType.md) |  | [optional] 
**SAP** | Pointer to [**SAP**](SAP.md) |  | [optional] 
**Operations** | Pointer to [**[]Operation**](Operation.md) |  | [optional] 
**AdministrativeState** | Pointer to [**AdministrativeState**](AdministrativeState.md) |  | [optional] 
**OperationalState** | Pointer to [**OperationalState**](OperationalState.md) |  | [optional] 
**UsageState** | Pointer to [**UsageState**](UsageState.md) |  | [optional] 
**RegistrationState** | Pointer to [**RegistrationState**](RegistrationState.md) |  | [optional] 

## Methods

### NewManagedNFServiceSingleAllOfAttributes

`func NewManagedNFServiceSingleAllOfAttributes() *ManagedNFServiceSingleAllOfAttributes`

NewManagedNFServiceSingleAllOfAttributes instantiates a new ManagedNFServiceSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedNFServiceSingleAllOfAttributesWithDefaults

`func NewManagedNFServiceSingleAllOfAttributesWithDefaults() *ManagedNFServiceSingleAllOfAttributes`

NewManagedNFServiceSingleAllOfAttributesWithDefaults instantiates a new ManagedNFServiceSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserLabel

`func (o *ManagedNFServiceSingleAllOfAttributes) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ManagedNFServiceSingleAllOfAttributes) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ManagedNFServiceSingleAllOfAttributes) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetNFServiceType

`func (o *ManagedNFServiceSingleAllOfAttributes) GetNFServiceType() NFServiceType`

GetNFServiceType returns the NFServiceType field if non-nil, zero value otherwise.

### GetNFServiceTypeOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetNFServiceTypeOk() (*NFServiceType, bool)`

GetNFServiceTypeOk returns a tuple with the NFServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFServiceType

`func (o *ManagedNFServiceSingleAllOfAttributes) SetNFServiceType(v NFServiceType)`

SetNFServiceType sets NFServiceType field to given value.

### HasNFServiceType

`func (o *ManagedNFServiceSingleAllOfAttributes) HasNFServiceType() bool`

HasNFServiceType returns a boolean if a field has been set.

### GetSAP

`func (o *ManagedNFServiceSingleAllOfAttributes) GetSAP() SAP`

GetSAP returns the SAP field if non-nil, zero value otherwise.

### GetSAPOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetSAPOk() (*SAP, bool)`

GetSAPOk returns a tuple with the SAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAP

`func (o *ManagedNFServiceSingleAllOfAttributes) SetSAP(v SAP)`

SetSAP sets SAP field to given value.

### HasSAP

`func (o *ManagedNFServiceSingleAllOfAttributes) HasSAP() bool`

HasSAP returns a boolean if a field has been set.

### GetOperations

`func (o *ManagedNFServiceSingleAllOfAttributes) GetOperations() []Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetOperationsOk() (*[]Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ManagedNFServiceSingleAllOfAttributes) SetOperations(v []Operation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ManagedNFServiceSingleAllOfAttributes) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetAdministrativeState

`func (o *ManagedNFServiceSingleAllOfAttributes) GetAdministrativeState() AdministrativeState`

GetAdministrativeState returns the AdministrativeState field if non-nil, zero value otherwise.

### GetAdministrativeStateOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool)`

GetAdministrativeStateOk returns a tuple with the AdministrativeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeState

`func (o *ManagedNFServiceSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState)`

SetAdministrativeState sets AdministrativeState field to given value.

### HasAdministrativeState

`func (o *ManagedNFServiceSingleAllOfAttributes) HasAdministrativeState() bool`

HasAdministrativeState returns a boolean if a field has been set.

### GetOperationalState

`func (o *ManagedNFServiceSingleAllOfAttributes) GetOperationalState() OperationalState`

GetOperationalState returns the OperationalState field if non-nil, zero value otherwise.

### GetOperationalStateOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool)`

GetOperationalStateOk returns a tuple with the OperationalState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalState

`func (o *ManagedNFServiceSingleAllOfAttributes) SetOperationalState(v OperationalState)`

SetOperationalState sets OperationalState field to given value.

### HasOperationalState

`func (o *ManagedNFServiceSingleAllOfAttributes) HasOperationalState() bool`

HasOperationalState returns a boolean if a field has been set.

### GetUsageState

`func (o *ManagedNFServiceSingleAllOfAttributes) GetUsageState() UsageState`

GetUsageState returns the UsageState field if non-nil, zero value otherwise.

### GetUsageStateOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetUsageStateOk() (*UsageState, bool)`

GetUsageStateOk returns a tuple with the UsageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageState

`func (o *ManagedNFServiceSingleAllOfAttributes) SetUsageState(v UsageState)`

SetUsageState sets UsageState field to given value.

### HasUsageState

`func (o *ManagedNFServiceSingleAllOfAttributes) HasUsageState() bool`

HasUsageState returns a boolean if a field has been set.

### GetRegistrationState

`func (o *ManagedNFServiceSingleAllOfAttributes) GetRegistrationState() RegistrationState`

GetRegistrationState returns the RegistrationState field if non-nil, zero value otherwise.

### GetRegistrationStateOk

`func (o *ManagedNFServiceSingleAllOfAttributes) GetRegistrationStateOk() (*RegistrationState, bool)`

GetRegistrationStateOk returns a tuple with the RegistrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationState

`func (o *ManagedNFServiceSingleAllOfAttributes) SetRegistrationState(v RegistrationState)`

SetRegistrationState sets RegistrationState field to given value.

### HasRegistrationState

`func (o *ManagedNFServiceSingleAllOfAttributes) HasRegistrationState() bool`

HasRegistrationState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


