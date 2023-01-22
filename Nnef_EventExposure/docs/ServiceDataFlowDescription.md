# ServiceDataFlowDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowDescription** | Pointer to [**IpPacketFilterSet**](IpPacketFilterSet.md) |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceDataFlowDescription

`func NewServiceDataFlowDescription() *ServiceDataFlowDescription`

NewServiceDataFlowDescription instantiates a new ServiceDataFlowDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceDataFlowDescriptionWithDefaults

`func NewServiceDataFlowDescriptionWithDefaults() *ServiceDataFlowDescription`

NewServiceDataFlowDescriptionWithDefaults instantiates a new ServiceDataFlowDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowDescription

`func (o *ServiceDataFlowDescription) GetFlowDescription() IpPacketFilterSet`

GetFlowDescription returns the FlowDescription field if non-nil, zero value otherwise.

### GetFlowDescriptionOk

`func (o *ServiceDataFlowDescription) GetFlowDescriptionOk() (*IpPacketFilterSet, bool)`

GetFlowDescriptionOk returns a tuple with the FlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDescription

`func (o *ServiceDataFlowDescription) SetFlowDescription(v IpPacketFilterSet)`

SetFlowDescription sets FlowDescription field to given value.

### HasFlowDescription

`func (o *ServiceDataFlowDescription) HasFlowDescription() bool`

HasFlowDescription returns a boolean if a field has been set.

### GetDomainName

`func (o *ServiceDataFlowDescription) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ServiceDataFlowDescription) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ServiceDataFlowDescription) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ServiceDataFlowDescription) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


