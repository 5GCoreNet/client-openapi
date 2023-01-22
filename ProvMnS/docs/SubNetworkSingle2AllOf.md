# SubNetworkSingle2AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**NetworkSlice** | Pointer to [**[]NetworkSliceSingle**](NetworkSliceSingle.md) |  | [optional] 
**NetworkSliceSubnet** | Pointer to [**[]NetworkSliceSubnetSingle**](NetworkSliceSubnetSingle.md) |  | [optional] 
**EPTransport** | Pointer to [**[]EPTransportSingle**](EPTransportSingle.md) |  | [optional] 
**NetworkSliceSubnetProviderCapabilities** | Pointer to [**[]NetworkSliceSubnetProviderCapabilitiesSingle**](NetworkSliceSubnetProviderCapabilitiesSingle.md) |  | [optional] 
**FeasibilityCheckAndReservationJob** | Pointer to [**[]FeasibilityCheckAndReservationJobSingle**](FeasibilityCheckAndReservationJobSingle.md) |  | [optional] 

## Methods

### NewSubNetworkSingle2AllOf

`func NewSubNetworkSingle2AllOf() *SubNetworkSingle2AllOf`

NewSubNetworkSingle2AllOf instantiates a new SubNetworkSingle2AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubNetworkSingle2AllOfWithDefaults

`func NewSubNetworkSingle2AllOfWithDefaults() *SubNetworkSingle2AllOf`

NewSubNetworkSingle2AllOfWithDefaults instantiates a new SubNetworkSingle2AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubNetwork

`func (o *SubNetworkSingle2AllOf) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *SubNetworkSingle2AllOf) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *SubNetworkSingle2AllOf) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *SubNetworkSingle2AllOf) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetNetworkSlice

`func (o *SubNetworkSingle2AllOf) GetNetworkSlice() []NetworkSliceSingle`

GetNetworkSlice returns the NetworkSlice field if non-nil, zero value otherwise.

### GetNetworkSliceOk

`func (o *SubNetworkSingle2AllOf) GetNetworkSliceOk() (*[]NetworkSliceSingle, bool)`

GetNetworkSliceOk returns a tuple with the NetworkSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSlice

`func (o *SubNetworkSingle2AllOf) SetNetworkSlice(v []NetworkSliceSingle)`

SetNetworkSlice sets NetworkSlice field to given value.

### HasNetworkSlice

`func (o *SubNetworkSingle2AllOf) HasNetworkSlice() bool`

HasNetworkSlice returns a boolean if a field has been set.

### GetNetworkSliceSubnet

`func (o *SubNetworkSingle2AllOf) GetNetworkSliceSubnet() []NetworkSliceSubnetSingle`

GetNetworkSliceSubnet returns the NetworkSliceSubnet field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetOk

`func (o *SubNetworkSingle2AllOf) GetNetworkSliceSubnetOk() (*[]NetworkSliceSubnetSingle, bool)`

GetNetworkSliceSubnetOk returns a tuple with the NetworkSliceSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnet

`func (o *SubNetworkSingle2AllOf) SetNetworkSliceSubnet(v []NetworkSliceSubnetSingle)`

SetNetworkSliceSubnet sets NetworkSliceSubnet field to given value.

### HasNetworkSliceSubnet

`func (o *SubNetworkSingle2AllOf) HasNetworkSliceSubnet() bool`

HasNetworkSliceSubnet returns a boolean if a field has been set.

### GetEPTransport

`func (o *SubNetworkSingle2AllOf) GetEPTransport() []EPTransportSingle`

GetEPTransport returns the EPTransport field if non-nil, zero value otherwise.

### GetEPTransportOk

`func (o *SubNetworkSingle2AllOf) GetEPTransportOk() (*[]EPTransportSingle, bool)`

GetEPTransportOk returns a tuple with the EPTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPTransport

`func (o *SubNetworkSingle2AllOf) SetEPTransport(v []EPTransportSingle)`

SetEPTransport sets EPTransport field to given value.

### HasEPTransport

`func (o *SubNetworkSingle2AllOf) HasEPTransport() bool`

HasEPTransport returns a boolean if a field has been set.

### GetNetworkSliceSubnetProviderCapabilities

`func (o *SubNetworkSingle2AllOf) GetNetworkSliceSubnetProviderCapabilities() []NetworkSliceSubnetProviderCapabilitiesSingle`

GetNetworkSliceSubnetProviderCapabilities returns the NetworkSliceSubnetProviderCapabilities field if non-nil, zero value otherwise.

### GetNetworkSliceSubnetProviderCapabilitiesOk

`func (o *SubNetworkSingle2AllOf) GetNetworkSliceSubnetProviderCapabilitiesOk() (*[]NetworkSliceSubnetProviderCapabilitiesSingle, bool)`

GetNetworkSliceSubnetProviderCapabilitiesOk returns a tuple with the NetworkSliceSubnetProviderCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSliceSubnetProviderCapabilities

`func (o *SubNetworkSingle2AllOf) SetNetworkSliceSubnetProviderCapabilities(v []NetworkSliceSubnetProviderCapabilitiesSingle)`

SetNetworkSliceSubnetProviderCapabilities sets NetworkSliceSubnetProviderCapabilities field to given value.

### HasNetworkSliceSubnetProviderCapabilities

`func (o *SubNetworkSingle2AllOf) HasNetworkSliceSubnetProviderCapabilities() bool`

HasNetworkSliceSubnetProviderCapabilities returns a boolean if a field has been set.

### GetFeasibilityCheckAndReservationJob

`func (o *SubNetworkSingle2AllOf) GetFeasibilityCheckAndReservationJob() []FeasibilityCheckAndReservationJobSingle`

GetFeasibilityCheckAndReservationJob returns the FeasibilityCheckAndReservationJob field if non-nil, zero value otherwise.

### GetFeasibilityCheckAndReservationJobOk

`func (o *SubNetworkSingle2AllOf) GetFeasibilityCheckAndReservationJobOk() (*[]FeasibilityCheckAndReservationJobSingle, bool)`

GetFeasibilityCheckAndReservationJobOk returns a tuple with the FeasibilityCheckAndReservationJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeasibilityCheckAndReservationJob

`func (o *SubNetworkSingle2AllOf) SetFeasibilityCheckAndReservationJob(v []FeasibilityCheckAndReservationJobSingle)`

SetFeasibilityCheckAndReservationJob sets FeasibilityCheckAndReservationJob field to given value.

### HasFeasibilityCheckAndReservationJob

`func (o *SubNetworkSingle2AllOf) HasFeasibilityCheckAndReservationJob() bool`

HasFeasibilityCheckAndReservationJob returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


