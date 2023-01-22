# DynamicPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DynamicPolicyId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**PolicyTemplateId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**ServiceDataFlowDescriptions** | [**[]ServiceDataFlowDescription**](ServiceDataFlowDescription.md) |  | 
**ProvisioningSessionId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**QosSpecification** | Pointer to [**M5QoSSpecification**](M5QoSSpecification.md) |  | [optional] 
**EnforcementMethod** | Pointer to **string** |  | [optional] 
**EnforcementBitRate** | Pointer to **int32** |  | [optional] 

## Methods

### NewDynamicPolicy

`func NewDynamicPolicy(dynamicPolicyId string, policyTemplateId string, serviceDataFlowDescriptions []ServiceDataFlowDescription, provisioningSessionId string, ) *DynamicPolicy`

NewDynamicPolicy instantiates a new DynamicPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicPolicyWithDefaults

`func NewDynamicPolicyWithDefaults() *DynamicPolicy`

NewDynamicPolicyWithDefaults instantiates a new DynamicPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamicPolicyId

`func (o *DynamicPolicy) GetDynamicPolicyId() string`

GetDynamicPolicyId returns the DynamicPolicyId field if non-nil, zero value otherwise.

### GetDynamicPolicyIdOk

`func (o *DynamicPolicy) GetDynamicPolicyIdOk() (*string, bool)`

GetDynamicPolicyIdOk returns a tuple with the DynamicPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPolicyId

`func (o *DynamicPolicy) SetDynamicPolicyId(v string)`

SetDynamicPolicyId sets DynamicPolicyId field to given value.


### GetPolicyTemplateId

`func (o *DynamicPolicy) GetPolicyTemplateId() string`

GetPolicyTemplateId returns the PolicyTemplateId field if non-nil, zero value otherwise.

### GetPolicyTemplateIdOk

`func (o *DynamicPolicy) GetPolicyTemplateIdOk() (*string, bool)`

GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTemplateId

`func (o *DynamicPolicy) SetPolicyTemplateId(v string)`

SetPolicyTemplateId sets PolicyTemplateId field to given value.


### GetServiceDataFlowDescriptions

`func (o *DynamicPolicy) GetServiceDataFlowDescriptions() []ServiceDataFlowDescription`

GetServiceDataFlowDescriptions returns the ServiceDataFlowDescriptions field if non-nil, zero value otherwise.

### GetServiceDataFlowDescriptionsOk

`func (o *DynamicPolicy) GetServiceDataFlowDescriptionsOk() (*[]ServiceDataFlowDescription, bool)`

GetServiceDataFlowDescriptionsOk returns a tuple with the ServiceDataFlowDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataFlowDescriptions

`func (o *DynamicPolicy) SetServiceDataFlowDescriptions(v []ServiceDataFlowDescription)`

SetServiceDataFlowDescriptions sets ServiceDataFlowDescriptions field to given value.


### GetProvisioningSessionId

`func (o *DynamicPolicy) GetProvisioningSessionId() string`

GetProvisioningSessionId returns the ProvisioningSessionId field if non-nil, zero value otherwise.

### GetProvisioningSessionIdOk

`func (o *DynamicPolicy) GetProvisioningSessionIdOk() (*string, bool)`

GetProvisioningSessionIdOk returns a tuple with the ProvisioningSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSessionId

`func (o *DynamicPolicy) SetProvisioningSessionId(v string)`

SetProvisioningSessionId sets ProvisioningSessionId field to given value.


### GetQosSpecification

`func (o *DynamicPolicy) GetQosSpecification() M5QoSSpecification`

GetQosSpecification returns the QosSpecification field if non-nil, zero value otherwise.

### GetQosSpecificationOk

`func (o *DynamicPolicy) GetQosSpecificationOk() (*M5QoSSpecification, bool)`

GetQosSpecificationOk returns a tuple with the QosSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosSpecification

`func (o *DynamicPolicy) SetQosSpecification(v M5QoSSpecification)`

SetQosSpecification sets QosSpecification field to given value.

### HasQosSpecification

`func (o *DynamicPolicy) HasQosSpecification() bool`

HasQosSpecification returns a boolean if a field has been set.

### GetEnforcementMethod

`func (o *DynamicPolicy) GetEnforcementMethod() string`

GetEnforcementMethod returns the EnforcementMethod field if non-nil, zero value otherwise.

### GetEnforcementMethodOk

`func (o *DynamicPolicy) GetEnforcementMethodOk() (*string, bool)`

GetEnforcementMethodOk returns a tuple with the EnforcementMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementMethod

`func (o *DynamicPolicy) SetEnforcementMethod(v string)`

SetEnforcementMethod sets EnforcementMethod field to given value.

### HasEnforcementMethod

`func (o *DynamicPolicy) HasEnforcementMethod() bool`

HasEnforcementMethod returns a boolean if a field has been set.

### GetEnforcementBitRate

`func (o *DynamicPolicy) GetEnforcementBitRate() int32`

GetEnforcementBitRate returns the EnforcementBitRate field if non-nil, zero value otherwise.

### GetEnforcementBitRateOk

`func (o *DynamicPolicy) GetEnforcementBitRateOk() (*int32, bool)`

GetEnforcementBitRateOk returns a tuple with the EnforcementBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementBitRate

`func (o *DynamicPolicy) SetEnforcementBitRate(v int32)`

SetEnforcementBitRate sets EnforcementBitRate field to given value.

### HasEnforcementBitRate

`func (o *DynamicPolicy) HasEnforcementBitRate() bool`

HasEnforcementBitRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


