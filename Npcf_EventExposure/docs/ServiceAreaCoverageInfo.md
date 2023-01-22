# ServiceAreaCoverageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TacList** | **[]string** | Indicates a list of Tracking Areas where the service is allowed. | 
**ServingNetwork** | Pointer to [**PlmnIdNid**](PlmnIdNid.md) |  | [optional] 

## Methods

### NewServiceAreaCoverageInfo

`func NewServiceAreaCoverageInfo(tacList []string, ) *ServiceAreaCoverageInfo`

NewServiceAreaCoverageInfo instantiates a new ServiceAreaCoverageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAreaCoverageInfoWithDefaults

`func NewServiceAreaCoverageInfoWithDefaults() *ServiceAreaCoverageInfo`

NewServiceAreaCoverageInfoWithDefaults instantiates a new ServiceAreaCoverageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTacList

`func (o *ServiceAreaCoverageInfo) GetTacList() []string`

GetTacList returns the TacList field if non-nil, zero value otherwise.

### GetTacListOk

`func (o *ServiceAreaCoverageInfo) GetTacListOk() (*[]string, bool)`

GetTacListOk returns a tuple with the TacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacList

`func (o *ServiceAreaCoverageInfo) SetTacList(v []string)`

SetTacList sets TacList field to given value.


### GetServingNetwork

`func (o *ServiceAreaCoverageInfo) GetServingNetwork() PlmnIdNid`

GetServingNetwork returns the ServingNetwork field if non-nil, zero value otherwise.

### GetServingNetworkOk

`func (o *ServiceAreaCoverageInfo) GetServingNetworkOk() (*PlmnIdNid, bool)`

GetServingNetworkOk returns a tuple with the ServingNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingNetwork

`func (o *ServiceAreaCoverageInfo) SetServingNetwork(v PlmnIdNid)`

SetServingNetwork sets ServingNetwork field to given value.

### HasServingNetwork

`func (o *ServiceAreaCoverageInfo) HasServingNetwork() bool`

HasServingNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


