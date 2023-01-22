# TrafficDescriptorComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDescs** | Pointer to [**map[string]AppDescriptor1**](AppDescriptor1.md) | Describes the operation systems and the corresponding applications for each operation systems. The key of map is osId. | [optional] 
**FlowDescs** | Pointer to **[]string** | Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic. The content of the string has the same encoding as the IPFilterRule AVP value as defined in IETF RFC 6733. | [optional] 
**DomainDescs** | Pointer to **[]string** | FQDN(s) or a regular expression which are used as a domain name matching criteria. | [optional] 
**EthFlowDescs** | Pointer to [**[]EthFlowDescription**](EthFlowDescription.md) | Descriptor(s) for destination information of non-IP traffic in which only ethernet flow description is defined. | [optional] 
**Dnns** | Pointer to **[]string** | This is matched against the DNN information provided by the application. | [optional] 
**ConnCaps** | Pointer to [**[]ConnectionCapabilities**](ConnectionCapabilities.md) | This is matched against the information provided by a UE application when it requests a network connection with certain capabilities. | [optional] 

## Methods

### NewTrafficDescriptorComponents

`func NewTrafficDescriptorComponents() *TrafficDescriptorComponents`

NewTrafficDescriptorComponents instantiates a new TrafficDescriptorComponents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficDescriptorComponentsWithDefaults

`func NewTrafficDescriptorComponentsWithDefaults() *TrafficDescriptorComponents`

NewTrafficDescriptorComponentsWithDefaults instantiates a new TrafficDescriptorComponents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDescs

`func (o *TrafficDescriptorComponents) GetAppDescs() map[string]AppDescriptor1`

GetAppDescs returns the AppDescs field if non-nil, zero value otherwise.

### GetAppDescsOk

`func (o *TrafficDescriptorComponents) GetAppDescsOk() (*map[string]AppDescriptor1, bool)`

GetAppDescsOk returns a tuple with the AppDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescs

`func (o *TrafficDescriptorComponents) SetAppDescs(v map[string]AppDescriptor1)`

SetAppDescs sets AppDescs field to given value.

### HasAppDescs

`func (o *TrafficDescriptorComponents) HasAppDescs() bool`

HasAppDescs returns a boolean if a field has been set.

### GetFlowDescs

`func (o *TrafficDescriptorComponents) GetFlowDescs() []string`

GetFlowDescs returns the FlowDescs field if non-nil, zero value otherwise.

### GetFlowDescsOk

`func (o *TrafficDescriptorComponents) GetFlowDescsOk() (*[]string, bool)`

GetFlowDescsOk returns a tuple with the FlowDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescs

`func (o *TrafficDescriptorComponents) SetFlowDescs(v []string)`

SetFlowDescs sets FlowDescs field to given value.

### HasFlowDescs

`func (o *TrafficDescriptorComponents) HasFlowDescs() bool`

HasFlowDescs returns a boolean if a field has been set.

### GetDomainDescs

`func (o *TrafficDescriptorComponents) GetDomainDescs() []string`

GetDomainDescs returns the DomainDescs field if non-nil, zero value otherwise.

### GetDomainDescsOk

`func (o *TrafficDescriptorComponents) GetDomainDescsOk() (*[]string, bool)`

GetDomainDescsOk returns a tuple with the DomainDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescs

`func (o *TrafficDescriptorComponents) SetDomainDescs(v []string)`

SetDomainDescs sets DomainDescs field to given value.

### HasDomainDescs

`func (o *TrafficDescriptorComponents) HasDomainDescs() bool`

HasDomainDescs returns a boolean if a field has been set.

### GetEthFlowDescs

`func (o *TrafficDescriptorComponents) GetEthFlowDescs() []EthFlowDescription`

GetEthFlowDescs returns the EthFlowDescs field if non-nil, zero value otherwise.

### GetEthFlowDescsOk

`func (o *TrafficDescriptorComponents) GetEthFlowDescsOk() (*[]EthFlowDescription, bool)`

GetEthFlowDescsOk returns a tuple with the EthFlowDescs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthFlowDescs

`func (o *TrafficDescriptorComponents) SetEthFlowDescs(v []EthFlowDescription)`

SetEthFlowDescs sets EthFlowDescs field to given value.

### HasEthFlowDescs

`func (o *TrafficDescriptorComponents) HasEthFlowDescs() bool`

HasEthFlowDescs returns a boolean if a field has been set.

### GetDnns

`func (o *TrafficDescriptorComponents) GetDnns() []string`

GetDnns returns the Dnns field if non-nil, zero value otherwise.

### GetDnnsOk

`func (o *TrafficDescriptorComponents) GetDnnsOk() (*[]string, bool)`

GetDnnsOk returns a tuple with the Dnns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnns

`func (o *TrafficDescriptorComponents) SetDnns(v []string)`

SetDnns sets Dnns field to given value.

### HasDnns

`func (o *TrafficDescriptorComponents) HasDnns() bool`

HasDnns returns a boolean if a field has been set.

### GetConnCaps

`func (o *TrafficDescriptorComponents) GetConnCaps() []ConnectionCapabilities`

GetConnCaps returns the ConnCaps field if non-nil, zero value otherwise.

### GetConnCapsOk

`func (o *TrafficDescriptorComponents) GetConnCapsOk() (*[]ConnectionCapabilities, bool)`

GetConnCapsOk returns a tuple with the ConnCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnCaps

`func (o *TrafficDescriptorComponents) SetConnCaps(v []ConnectionCapabilities)`

SetConnCaps sets ConnCaps field to given value.

### HasConnCaps

`func (o *TrafficDescriptorComponents) HasConnCaps() bool`

HasConnCaps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


