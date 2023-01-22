# ServiceAccessInformationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisioningSessionId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**ProvisioningSessionType** | [**ProvisioningSessionType**](ProvisioningSessionType.md) |  | 
**StreamingAccess** | Pointer to [**ServiceAccessInformationResourceStreamingAccess**](ServiceAccessInformationResourceStreamingAccess.md) |  | [optional] 
**ClientConsumptionReportingConfiguration** | Pointer to [**ServiceAccessInformationResourceClientConsumptionReportingConfiguration**](ServiceAccessInformationResourceClientConsumptionReportingConfiguration.md) |  | [optional] 
**DynamicPolicyInvocationConfiguration** | Pointer to [**ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration**](ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration.md) |  | [optional] 
**ClientMetricsReportingConfiguration** | Pointer to [**[]ServiceAccessInformationResourceClientMetricsReportingConfigurationInner**](ServiceAccessInformationResourceClientMetricsReportingConfigurationInner.md) |  | [optional] 
**NetworkAssistanceConfiguration** | Pointer to [**ServiceAccessInformationResourceNetworkAssistanceConfiguration**](ServiceAccessInformationResourceNetworkAssistanceConfiguration.md) |  | [optional] 
**ClientEdgeResourcesConfiguration** | Pointer to [**ServiceAccessInformationResourceClientEdgeResourcesConfiguration**](ServiceAccessInformationResourceClientEdgeResourcesConfiguration.md) |  | [optional] 

## Methods

### NewServiceAccessInformationResource

`func NewServiceAccessInformationResource(provisioningSessionId string, provisioningSessionType ProvisioningSessionType, ) *ServiceAccessInformationResource`

NewServiceAccessInformationResource instantiates a new ServiceAccessInformationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessInformationResourceWithDefaults

`func NewServiceAccessInformationResourceWithDefaults() *ServiceAccessInformationResource`

NewServiceAccessInformationResourceWithDefaults instantiates a new ServiceAccessInformationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisioningSessionId

`func (o *ServiceAccessInformationResource) GetProvisioningSessionId() string`

GetProvisioningSessionId returns the ProvisioningSessionId field if non-nil, zero value otherwise.

### GetProvisioningSessionIdOk

`func (o *ServiceAccessInformationResource) GetProvisioningSessionIdOk() (*string, bool)`

GetProvisioningSessionIdOk returns a tuple with the ProvisioningSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSessionId

`func (o *ServiceAccessInformationResource) SetProvisioningSessionId(v string)`

SetProvisioningSessionId sets ProvisioningSessionId field to given value.


### GetProvisioningSessionType

`func (o *ServiceAccessInformationResource) GetProvisioningSessionType() ProvisioningSessionType`

GetProvisioningSessionType returns the ProvisioningSessionType field if non-nil, zero value otherwise.

### GetProvisioningSessionTypeOk

`func (o *ServiceAccessInformationResource) GetProvisioningSessionTypeOk() (*ProvisioningSessionType, bool)`

GetProvisioningSessionTypeOk returns a tuple with the ProvisioningSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSessionType

`func (o *ServiceAccessInformationResource) SetProvisioningSessionType(v ProvisioningSessionType)`

SetProvisioningSessionType sets ProvisioningSessionType field to given value.


### GetStreamingAccess

`func (o *ServiceAccessInformationResource) GetStreamingAccess() ServiceAccessInformationResourceStreamingAccess`

GetStreamingAccess returns the StreamingAccess field if non-nil, zero value otherwise.

### GetStreamingAccessOk

`func (o *ServiceAccessInformationResource) GetStreamingAccessOk() (*ServiceAccessInformationResourceStreamingAccess, bool)`

GetStreamingAccessOk returns a tuple with the StreamingAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingAccess

`func (o *ServiceAccessInformationResource) SetStreamingAccess(v ServiceAccessInformationResourceStreamingAccess)`

SetStreamingAccess sets StreamingAccess field to given value.

### HasStreamingAccess

`func (o *ServiceAccessInformationResource) HasStreamingAccess() bool`

HasStreamingAccess returns a boolean if a field has been set.

### GetClientConsumptionReportingConfiguration

`func (o *ServiceAccessInformationResource) GetClientConsumptionReportingConfiguration() ServiceAccessInformationResourceClientConsumptionReportingConfiguration`

GetClientConsumptionReportingConfiguration returns the ClientConsumptionReportingConfiguration field if non-nil, zero value otherwise.

### GetClientConsumptionReportingConfigurationOk

`func (o *ServiceAccessInformationResource) GetClientConsumptionReportingConfigurationOk() (*ServiceAccessInformationResourceClientConsumptionReportingConfiguration, bool)`

GetClientConsumptionReportingConfigurationOk returns a tuple with the ClientConsumptionReportingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConsumptionReportingConfiguration

`func (o *ServiceAccessInformationResource) SetClientConsumptionReportingConfiguration(v ServiceAccessInformationResourceClientConsumptionReportingConfiguration)`

SetClientConsumptionReportingConfiguration sets ClientConsumptionReportingConfiguration field to given value.

### HasClientConsumptionReportingConfiguration

`func (o *ServiceAccessInformationResource) HasClientConsumptionReportingConfiguration() bool`

HasClientConsumptionReportingConfiguration returns a boolean if a field has been set.

### GetDynamicPolicyInvocationConfiguration

`func (o *ServiceAccessInformationResource) GetDynamicPolicyInvocationConfiguration() ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration`

GetDynamicPolicyInvocationConfiguration returns the DynamicPolicyInvocationConfiguration field if non-nil, zero value otherwise.

### GetDynamicPolicyInvocationConfigurationOk

`func (o *ServiceAccessInformationResource) GetDynamicPolicyInvocationConfigurationOk() (*ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration, bool)`

GetDynamicPolicyInvocationConfigurationOk returns a tuple with the DynamicPolicyInvocationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicPolicyInvocationConfiguration

`func (o *ServiceAccessInformationResource) SetDynamicPolicyInvocationConfiguration(v ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration)`

SetDynamicPolicyInvocationConfiguration sets DynamicPolicyInvocationConfiguration field to given value.

### HasDynamicPolicyInvocationConfiguration

`func (o *ServiceAccessInformationResource) HasDynamicPolicyInvocationConfiguration() bool`

HasDynamicPolicyInvocationConfiguration returns a boolean if a field has been set.

### GetClientMetricsReportingConfiguration

`func (o *ServiceAccessInformationResource) GetClientMetricsReportingConfiguration() []ServiceAccessInformationResourceClientMetricsReportingConfigurationInner`

GetClientMetricsReportingConfiguration returns the ClientMetricsReportingConfiguration field if non-nil, zero value otherwise.

### GetClientMetricsReportingConfigurationOk

`func (o *ServiceAccessInformationResource) GetClientMetricsReportingConfigurationOk() (*[]ServiceAccessInformationResourceClientMetricsReportingConfigurationInner, bool)`

GetClientMetricsReportingConfigurationOk returns a tuple with the ClientMetricsReportingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMetricsReportingConfiguration

`func (o *ServiceAccessInformationResource) SetClientMetricsReportingConfiguration(v []ServiceAccessInformationResourceClientMetricsReportingConfigurationInner)`

SetClientMetricsReportingConfiguration sets ClientMetricsReportingConfiguration field to given value.

### HasClientMetricsReportingConfiguration

`func (o *ServiceAccessInformationResource) HasClientMetricsReportingConfiguration() bool`

HasClientMetricsReportingConfiguration returns a boolean if a field has been set.

### GetNetworkAssistanceConfiguration

`func (o *ServiceAccessInformationResource) GetNetworkAssistanceConfiguration() ServiceAccessInformationResourceNetworkAssistanceConfiguration`

GetNetworkAssistanceConfiguration returns the NetworkAssistanceConfiguration field if non-nil, zero value otherwise.

### GetNetworkAssistanceConfigurationOk

`func (o *ServiceAccessInformationResource) GetNetworkAssistanceConfigurationOk() (*ServiceAccessInformationResourceNetworkAssistanceConfiguration, bool)`

GetNetworkAssistanceConfigurationOk returns a tuple with the NetworkAssistanceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAssistanceConfiguration

`func (o *ServiceAccessInformationResource) SetNetworkAssistanceConfiguration(v ServiceAccessInformationResourceNetworkAssistanceConfiguration)`

SetNetworkAssistanceConfiguration sets NetworkAssistanceConfiguration field to given value.

### HasNetworkAssistanceConfiguration

`func (o *ServiceAccessInformationResource) HasNetworkAssistanceConfiguration() bool`

HasNetworkAssistanceConfiguration returns a boolean if a field has been set.

### GetClientEdgeResourcesConfiguration

`func (o *ServiceAccessInformationResource) GetClientEdgeResourcesConfiguration() ServiceAccessInformationResourceClientEdgeResourcesConfiguration`

GetClientEdgeResourcesConfiguration returns the ClientEdgeResourcesConfiguration field if non-nil, zero value otherwise.

### GetClientEdgeResourcesConfigurationOk

`func (o *ServiceAccessInformationResource) GetClientEdgeResourcesConfigurationOk() (*ServiceAccessInformationResourceClientEdgeResourcesConfiguration, bool)`

GetClientEdgeResourcesConfigurationOk returns a tuple with the ClientEdgeResourcesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEdgeResourcesConfiguration

`func (o *ServiceAccessInformationResource) SetClientEdgeResourcesConfiguration(v ServiceAccessInformationResourceClientEdgeResourcesConfiguration)`

SetClientEdgeResourcesConfiguration sets ClientEdgeResourcesConfiguration field to given value.

### HasClientEdgeResourcesConfiguration

`func (o *ServiceAccessInformationResource) HasClientEdgeResourcesConfiguration() bool`

HasClientEdgeResourcesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


