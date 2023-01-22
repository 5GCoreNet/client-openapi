# SubNetworkSingleAllOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**ManagedElement** | Pointer to [**[]ManagedElementSingle**](ManagedElementSingle.md) |  | [optional] 
**NRFrequency** | Pointer to [**[]NRFrequencySingle**](NRFrequencySingle.md) |  | [optional] 
**ExternalGnbCuCpFunction** | Pointer to [**[]ExternalGnbCuCpFunctionSingle**](ExternalGnbCuCpFunctionSingle.md) |  | [optional] 
**ExternalENBFunction** | Pointer to [**[]ExternalENBFunctionSingle**](ExternalENBFunctionSingle.md) |  | [optional] 
**EUtranFrequency** | Pointer to [**[]EUtranFrequencySingle**](EUtranFrequencySingle.md) |  | [optional] 
**DESManagementFunction** | Pointer to [**DESManagementFunctionSingle**](DESManagementFunctionSingle.md) |  | [optional] 
**DRACHOptimizationFunction** | Pointer to [**DRACHOptimizationFunctionSingle**](DRACHOptimizationFunctionSingle.md) |  | [optional] 
**DMROFunction** | Pointer to [**DMROFunctionSingle**](DMROFunctionSingle.md) |  | [optional] 
**DLBOFunction** | Pointer to [**DLBOFunctionSingle**](DLBOFunctionSingle.md) |  | [optional] 
**DPCIConfigurationFunction** | Pointer to [**DPCIConfigurationFunctionSingle**](DPCIConfigurationFunctionSingle.md) |  | [optional] 
**CPCIConfigurationFunction** | Pointer to [**CPCIConfigurationFunctionSingle**](CPCIConfigurationFunctionSingle.md) |  | [optional] 
**CESManagementFunction** | Pointer to [**CESManagementFunctionSingle**](CESManagementFunctionSingle.md) |  | [optional] 
**Configurable5QISet** | Pointer to [**[]Configurable5QISetSingle**](Configurable5QISetSingle.md) |  | [optional] 
**RimRSGlobal** | Pointer to [**RimRSGlobalSingle**](RimRSGlobalSingle.md) |  | [optional] 
**Dynamic5QISet** | Pointer to [**[]Dynamic5QISetSingle**](Dynamic5QISetSingle.md) |  | [optional] 
**CCOFunction** | Pointer to [**CCOFunctionSingle**](CCOFunctionSingle.md) |  | [optional] 

## Methods

### NewSubNetworkSingleAllOf1

`func NewSubNetworkSingleAllOf1() *SubNetworkSingleAllOf1`

NewSubNetworkSingleAllOf1 instantiates a new SubNetworkSingleAllOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkSingleAllOf1WithDefaults

`func NewSubNetworkSingleAllOf1WithDefaults() *SubNetworkSingleAllOf1`

NewSubNetworkSingleAllOf1WithDefaults instantiates a new SubNetworkSingleAllOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubNetwork

`func (o *SubNetworkSingleAllOf1) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *SubNetworkSingleAllOf1) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *SubNetworkSingleAllOf1) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *SubNetworkSingleAllOf1) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetManagedElement

`func (o *SubNetworkSingleAllOf1) GetManagedElement() []ManagedElementSingle`

GetManagedElement returns the ManagedElement field if non-nil, zero value otherwise.

### GetManagedElementOk

`func (o *SubNetworkSingleAllOf1) GetManagedElementOk() (*[]ManagedElementSingle, bool)`

GetManagedElementOk returns a tuple with the ManagedElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElement

`func (o *SubNetworkSingleAllOf1) SetManagedElement(v []ManagedElementSingle)`

SetManagedElement sets ManagedElement field to given value.

### HasManagedElement

`func (o *SubNetworkSingleAllOf1) HasManagedElement() bool`

HasManagedElement returns a boolean if a field has been set.

### GetNRFrequency

`func (o *SubNetworkSingleAllOf1) GetNRFrequency() []NRFrequencySingle`

GetNRFrequency returns the NRFrequency field if non-nil, zero value otherwise.

### GetNRFrequencyOk

`func (o *SubNetworkSingleAllOf1) GetNRFrequencyOk() (*[]NRFrequencySingle, bool)`

GetNRFrequencyOk returns a tuple with the NRFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNRFrequency

`func (o *SubNetworkSingleAllOf1) SetNRFrequency(v []NRFrequencySingle)`

SetNRFrequency sets NRFrequency field to given value.

### HasNRFrequency

`func (o *SubNetworkSingleAllOf1) HasNRFrequency() bool`

HasNRFrequency returns a boolean if a field has been set.

### GetExternalGnbCuCpFunction

`func (o *SubNetworkSingleAllOf1) GetExternalGnbCuCpFunction() []ExternalGnbCuCpFunctionSingle`

GetExternalGnbCuCpFunction returns the ExternalGnbCuCpFunction field if non-nil, zero value otherwise.

### GetExternalGnbCuCpFunctionOk

`func (o *SubNetworkSingleAllOf1) GetExternalGnbCuCpFunctionOk() (*[]ExternalGnbCuCpFunctionSingle, bool)`

GetExternalGnbCuCpFunctionOk returns a tuple with the ExternalGnbCuCpFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGnbCuCpFunction

`func (o *SubNetworkSingleAllOf1) SetExternalGnbCuCpFunction(v []ExternalGnbCuCpFunctionSingle)`

SetExternalGnbCuCpFunction sets ExternalGnbCuCpFunction field to given value.

### HasExternalGnbCuCpFunction

`func (o *SubNetworkSingleAllOf1) HasExternalGnbCuCpFunction() bool`

HasExternalGnbCuCpFunction returns a boolean if a field has been set.

### GetExternalENBFunction

`func (o *SubNetworkSingleAllOf1) GetExternalENBFunction() []ExternalENBFunctionSingle`

GetExternalENBFunction returns the ExternalENBFunction field if non-nil, zero value otherwise.

### GetExternalENBFunctionOk

`func (o *SubNetworkSingleAllOf1) GetExternalENBFunctionOk() (*[]ExternalENBFunctionSingle, bool)`

GetExternalENBFunctionOk returns a tuple with the ExternalENBFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalENBFunction

`func (o *SubNetworkSingleAllOf1) SetExternalENBFunction(v []ExternalENBFunctionSingle)`

SetExternalENBFunction sets ExternalENBFunction field to given value.

### HasExternalENBFunction

`func (o *SubNetworkSingleAllOf1) HasExternalENBFunction() bool`

HasExternalENBFunction returns a boolean if a field has been set.

### GetEUtranFrequency

`func (o *SubNetworkSingleAllOf1) GetEUtranFrequency() []EUtranFrequencySingle`

GetEUtranFrequency returns the EUtranFrequency field if non-nil, zero value otherwise.

### GetEUtranFrequencyOk

`func (o *SubNetworkSingleAllOf1) GetEUtranFrequencyOk() (*[]EUtranFrequencySingle, bool)`

GetEUtranFrequencyOk returns a tuple with the EUtranFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranFrequency

`func (o *SubNetworkSingleAllOf1) SetEUtranFrequency(v []EUtranFrequencySingle)`

SetEUtranFrequency sets EUtranFrequency field to given value.

### HasEUtranFrequency

`func (o *SubNetworkSingleAllOf1) HasEUtranFrequency() bool`

HasEUtranFrequency returns a boolean if a field has been set.

### GetDESManagementFunction

`func (o *SubNetworkSingleAllOf1) GetDESManagementFunction() DESManagementFunctionSingle`

GetDESManagementFunction returns the DESManagementFunction field if non-nil, zero value otherwise.

### GetDESManagementFunctionOk

`func (o *SubNetworkSingleAllOf1) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool)`

GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDESManagementFunction

`func (o *SubNetworkSingleAllOf1) SetDESManagementFunction(v DESManagementFunctionSingle)`

SetDESManagementFunction sets DESManagementFunction field to given value.

### HasDESManagementFunction

`func (o *SubNetworkSingleAllOf1) HasDESManagementFunction() bool`

HasDESManagementFunction returns a boolean if a field has been set.

### GetDRACHOptimizationFunction

`func (o *SubNetworkSingleAllOf1) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle`

GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field if non-nil, zero value otherwise.

### GetDRACHOptimizationFunctionOk

`func (o *SubNetworkSingleAllOf1) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool)`

GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDRACHOptimizationFunction

`func (o *SubNetworkSingleAllOf1) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle)`

SetDRACHOptimizationFunction sets DRACHOptimizationFunction field to given value.

### HasDRACHOptimizationFunction

`func (o *SubNetworkSingleAllOf1) HasDRACHOptimizationFunction() bool`

HasDRACHOptimizationFunction returns a boolean if a field has been set.

### GetDMROFunction

`func (o *SubNetworkSingleAllOf1) GetDMROFunction() DMROFunctionSingle`

GetDMROFunction returns the DMROFunction field if non-nil, zero value otherwise.

### GetDMROFunctionOk

`func (o *SubNetworkSingleAllOf1) GetDMROFunctionOk() (*DMROFunctionSingle, bool)`

GetDMROFunctionOk returns a tuple with the DMROFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDMROFunction

`func (o *SubNetworkSingleAllOf1) SetDMROFunction(v DMROFunctionSingle)`

SetDMROFunction sets DMROFunction field to given value.

### HasDMROFunction

`func (o *SubNetworkSingleAllOf1) HasDMROFunction() bool`

HasDMROFunction returns a boolean if a field has been set.

### GetDLBOFunction

`func (o *SubNetworkSingleAllOf1) GetDLBOFunction() DLBOFunctionSingle`

GetDLBOFunction returns the DLBOFunction field if non-nil, zero value otherwise.

### GetDLBOFunctionOk

`func (o *SubNetworkSingleAllOf1) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool)`

GetDLBOFunctionOk returns a tuple with the DLBOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDLBOFunction

`func (o *SubNetworkSingleAllOf1) SetDLBOFunction(v DLBOFunctionSingle)`

SetDLBOFunction sets DLBOFunction field to given value.

### HasDLBOFunction

`func (o *SubNetworkSingleAllOf1) HasDLBOFunction() bool`

HasDLBOFunction returns a boolean if a field has been set.

### GetDPCIConfigurationFunction

`func (o *SubNetworkSingleAllOf1) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle`

GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetDPCIConfigurationFunctionOk

`func (o *SubNetworkSingleAllOf1) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool)`

GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPCIConfigurationFunction

`func (o *SubNetworkSingleAllOf1) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle)`

SetDPCIConfigurationFunction sets DPCIConfigurationFunction field to given value.

### HasDPCIConfigurationFunction

`func (o *SubNetworkSingleAllOf1) HasDPCIConfigurationFunction() bool`

HasDPCIConfigurationFunction returns a boolean if a field has been set.

### GetCPCIConfigurationFunction

`func (o *SubNetworkSingleAllOf1) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle`

GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field if non-nil, zero value otherwise.

### GetCPCIConfigurationFunctionOk

`func (o *SubNetworkSingleAllOf1) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool)`

GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCPCIConfigurationFunction

`func (o *SubNetworkSingleAllOf1) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle)`

SetCPCIConfigurationFunction sets CPCIConfigurationFunction field to given value.

### HasCPCIConfigurationFunction

`func (o *SubNetworkSingleAllOf1) HasCPCIConfigurationFunction() bool`

HasCPCIConfigurationFunction returns a boolean if a field has been set.

### GetCESManagementFunction

`func (o *SubNetworkSingleAllOf1) GetCESManagementFunction() CESManagementFunctionSingle`

GetCESManagementFunction returns the CESManagementFunction field if non-nil, zero value otherwise.

### GetCESManagementFunctionOk

`func (o *SubNetworkSingleAllOf1) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool)`

GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCESManagementFunction

`func (o *SubNetworkSingleAllOf1) SetCESManagementFunction(v CESManagementFunctionSingle)`

SetCESManagementFunction sets CESManagementFunction field to given value.

### HasCESManagementFunction

`func (o *SubNetworkSingleAllOf1) HasCESManagementFunction() bool`

HasCESManagementFunction returns a boolean if a field has been set.

### GetConfigurable5QISet

`func (o *SubNetworkSingleAllOf1) GetConfigurable5QISet() []Configurable5QISetSingle`

GetConfigurable5QISet returns the Configurable5QISet field if non-nil, zero value otherwise.

### GetConfigurable5QISetOk

`func (o *SubNetworkSingleAllOf1) GetConfigurable5QISetOk() (*[]Configurable5QISetSingle, bool)`

GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable5QISet

`func (o *SubNetworkSingleAllOf1) SetConfigurable5QISet(v []Configurable5QISetSingle)`

SetConfigurable5QISet sets Configurable5QISet field to given value.

### HasConfigurable5QISet

`func (o *SubNetworkSingleAllOf1) HasConfigurable5QISet() bool`

HasConfigurable5QISet returns a boolean if a field has been set.

### GetRimRSGlobal

`func (o *SubNetworkSingleAllOf1) GetRimRSGlobal() RimRSGlobalSingle`

GetRimRSGlobal returns the RimRSGlobal field if non-nil, zero value otherwise.

### GetRimRSGlobalOk

`func (o *SubNetworkSingleAllOf1) GetRimRSGlobalOk() (*RimRSGlobalSingle, bool)`

GetRimRSGlobalOk returns a tuple with the RimRSGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRimRSGlobal

`func (o *SubNetworkSingleAllOf1) SetRimRSGlobal(v RimRSGlobalSingle)`

SetRimRSGlobal sets RimRSGlobal field to given value.

### HasRimRSGlobal

`func (o *SubNetworkSingleAllOf1) HasRimRSGlobal() bool`

HasRimRSGlobal returns a boolean if a field has been set.

### GetDynamic5QISet

`func (o *SubNetworkSingleAllOf1) GetDynamic5QISet() []Dynamic5QISetSingle`

GetDynamic5QISet returns the Dynamic5QISet field if non-nil, zero value otherwise.

### GetDynamic5QISetOk

`func (o *SubNetworkSingleAllOf1) GetDynamic5QISetOk() (*[]Dynamic5QISetSingle, bool)`

GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamic5QISet

`func (o *SubNetworkSingleAllOf1) SetDynamic5QISet(v []Dynamic5QISetSingle)`

SetDynamic5QISet sets Dynamic5QISet field to given value.

### HasDynamic5QISet

`func (o *SubNetworkSingleAllOf1) HasDynamic5QISet() bool`

HasDynamic5QISet returns a boolean if a field has been set.

### GetCCOFunction

`func (o *SubNetworkSingleAllOf1) GetCCOFunction() CCOFunctionSingle`

GetCCOFunction returns the CCOFunction field if non-nil, zero value otherwise.

### GetCCOFunctionOk

`func (o *SubNetworkSingleAllOf1) GetCCOFunctionOk() (*CCOFunctionSingle, bool)`

GetCCOFunctionOk returns a tuple with the CCOFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCCOFunction

`func (o *SubNetworkSingleAllOf1) SetCCOFunction(v CCOFunctionSingle)`

SetCCOFunction sets CCOFunction field to given value.

### HasCCOFunction

`func (o *SubNetworkSingleAllOf1) HasCCOFunction() bool`

HasCCOFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


