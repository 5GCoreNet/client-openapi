# ScpDomainRoutingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScpDomainList** | [**map[string]ScpDomainConnectivity**](ScpDomainConnectivity.md) | This IE shall contain a map of SCP domain interconnection information, where the key of the map is a SCP domain. The value of each entry shall be the interconnectivity information of the the SCP domain indicated by the key. An empty map indicates that there is no SCP domain currently registered in the NRF.  | 

## Methods

### NewScpDomainRoutingInformation

`func NewScpDomainRoutingInformation(scpDomainList map[string]ScpDomainConnectivity, ) *ScpDomainRoutingInformation`

NewScpDomainRoutingInformation instantiates a new ScpDomainRoutingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScpDomainRoutingInformationWithDefaults

`func NewScpDomainRoutingInformationWithDefaults() *ScpDomainRoutingInformation`

NewScpDomainRoutingInformationWithDefaults instantiates a new ScpDomainRoutingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScpDomainList

`func (o *ScpDomainRoutingInformation) GetScpDomainList() map[string]ScpDomainConnectivity`

GetScpDomainList returns the ScpDomainList field if non-nil, zero value otherwise.

### GetScpDomainListOk

`func (o *ScpDomainRoutingInformation) GetScpDomainListOk() (*map[string]ScpDomainConnectivity, bool)`

GetScpDomainListOk returns a tuple with the ScpDomainList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomainList

`func (o *ScpDomainRoutingInformation) SetScpDomainList(v map[string]ScpDomainConnectivity)`

SetScpDomainList sets ScpDomainList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


