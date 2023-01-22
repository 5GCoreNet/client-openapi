# ProvisioningSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisioningSessionId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**ProvisioningSessionType** | [**ProvisioningSessionType**](ProvisioningSessionType.md) |  | 
**AspId** | Pointer to **string** | Contains an identity of an application service provider. | [optional] 
**ExternalApplicationId** | **string** | String providing an application identifier. | 
**ServerCertificateIds** | Pointer to **[]string** |  | [optional] 
**ContentPreparationTemplateIds** | Pointer to **[]string** |  | [optional] 
**MetricsReportingConfigurationIds** | Pointer to **[]string** |  | [optional] 
**PolicyTemplateIds** | Pointer to **[]string** |  | [optional] 
**EdgeResourcesConfigurationIds** | Pointer to **[]string** |  | [optional] 
**EventDataProcessingConfigurationIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProvisioningSession

`func NewProvisioningSession(provisioningSessionId string, provisioningSessionType ProvisioningSessionType, externalApplicationId string, ) *ProvisioningSession`

NewProvisioningSession instantiates a new ProvisioningSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSessionWithDefaults

`func NewProvisioningSessionWithDefaults() *ProvisioningSession`

NewProvisioningSessionWithDefaults instantiates a new ProvisioningSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisioningSessionId

`func (o *ProvisioningSession) GetProvisioningSessionId() string`

GetProvisioningSessionId returns the ProvisioningSessionId field if non-nil, zero value otherwise.

### GetProvisioningSessionIdOk

`func (o *ProvisioningSession) GetProvisioningSessionIdOk() (*string, bool)`

GetProvisioningSessionIdOk returns a tuple with the ProvisioningSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSessionId

`func (o *ProvisioningSession) SetProvisioningSessionId(v string)`

SetProvisioningSessionId sets ProvisioningSessionId field to given value.


### GetProvisioningSessionType

`func (o *ProvisioningSession) GetProvisioningSessionType() ProvisioningSessionType`

GetProvisioningSessionType returns the ProvisioningSessionType field if non-nil, zero value otherwise.

### GetProvisioningSessionTypeOk

`func (o *ProvisioningSession) GetProvisioningSessionTypeOk() (*ProvisioningSessionType, bool)`

GetProvisioningSessionTypeOk returns a tuple with the ProvisioningSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSessionType

`func (o *ProvisioningSession) SetProvisioningSessionType(v ProvisioningSessionType)`

SetProvisioningSessionType sets ProvisioningSessionType field to given value.


### GetAspId

`func (o *ProvisioningSession) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *ProvisioningSession) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *ProvisioningSession) SetAspId(v string)`

SetAspId sets AspId field to given value.

### HasAspId

`func (o *ProvisioningSession) HasAspId() bool`

HasAspId returns a boolean if a field has been set.

### GetExternalApplicationId

`func (o *ProvisioningSession) GetExternalApplicationId() string`

GetExternalApplicationId returns the ExternalApplicationId field if non-nil, zero value otherwise.

### GetExternalApplicationIdOk

`func (o *ProvisioningSession) GetExternalApplicationIdOk() (*string, bool)`

GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApplicationId

`func (o *ProvisioningSession) SetExternalApplicationId(v string)`

SetExternalApplicationId sets ExternalApplicationId field to given value.


### GetServerCertificateIds

`func (o *ProvisioningSession) GetServerCertificateIds() []string`

GetServerCertificateIds returns the ServerCertificateIds field if non-nil, zero value otherwise.

### GetServerCertificateIdsOk

`func (o *ProvisioningSession) GetServerCertificateIdsOk() (*[]string, bool)`

GetServerCertificateIdsOk returns a tuple with the ServerCertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertificateIds

`func (o *ProvisioningSession) SetServerCertificateIds(v []string)`

SetServerCertificateIds sets ServerCertificateIds field to given value.

### HasServerCertificateIds

`func (o *ProvisioningSession) HasServerCertificateIds() bool`

HasServerCertificateIds returns a boolean if a field has been set.

### GetContentPreparationTemplateIds

`func (o *ProvisioningSession) GetContentPreparationTemplateIds() []string`

GetContentPreparationTemplateIds returns the ContentPreparationTemplateIds field if non-nil, zero value otherwise.

### GetContentPreparationTemplateIdsOk

`func (o *ProvisioningSession) GetContentPreparationTemplateIdsOk() (*[]string, bool)`

GetContentPreparationTemplateIdsOk returns a tuple with the ContentPreparationTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPreparationTemplateIds

`func (o *ProvisioningSession) SetContentPreparationTemplateIds(v []string)`

SetContentPreparationTemplateIds sets ContentPreparationTemplateIds field to given value.

### HasContentPreparationTemplateIds

`func (o *ProvisioningSession) HasContentPreparationTemplateIds() bool`

HasContentPreparationTemplateIds returns a boolean if a field has been set.

### GetMetricsReportingConfigurationIds

`func (o *ProvisioningSession) GetMetricsReportingConfigurationIds() []string`

GetMetricsReportingConfigurationIds returns the MetricsReportingConfigurationIds field if non-nil, zero value otherwise.

### GetMetricsReportingConfigurationIdsOk

`func (o *ProvisioningSession) GetMetricsReportingConfigurationIdsOk() (*[]string, bool)`

GetMetricsReportingConfigurationIdsOk returns a tuple with the MetricsReportingConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsReportingConfigurationIds

`func (o *ProvisioningSession) SetMetricsReportingConfigurationIds(v []string)`

SetMetricsReportingConfigurationIds sets MetricsReportingConfigurationIds field to given value.

### HasMetricsReportingConfigurationIds

`func (o *ProvisioningSession) HasMetricsReportingConfigurationIds() bool`

HasMetricsReportingConfigurationIds returns a boolean if a field has been set.

### GetPolicyTemplateIds

`func (o *ProvisioningSession) GetPolicyTemplateIds() []string`

GetPolicyTemplateIds returns the PolicyTemplateIds field if non-nil, zero value otherwise.

### GetPolicyTemplateIdsOk

`func (o *ProvisioningSession) GetPolicyTemplateIdsOk() (*[]string, bool)`

GetPolicyTemplateIdsOk returns a tuple with the PolicyTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTemplateIds

`func (o *ProvisioningSession) SetPolicyTemplateIds(v []string)`

SetPolicyTemplateIds sets PolicyTemplateIds field to given value.

### HasPolicyTemplateIds

`func (o *ProvisioningSession) HasPolicyTemplateIds() bool`

HasPolicyTemplateIds returns a boolean if a field has been set.

### GetEdgeResourcesConfigurationIds

`func (o *ProvisioningSession) GetEdgeResourcesConfigurationIds() []string`

GetEdgeResourcesConfigurationIds returns the EdgeResourcesConfigurationIds field if non-nil, zero value otherwise.

### GetEdgeResourcesConfigurationIdsOk

`func (o *ProvisioningSession) GetEdgeResourcesConfigurationIdsOk() (*[]string, bool)`

GetEdgeResourcesConfigurationIdsOk returns a tuple with the EdgeResourcesConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeResourcesConfigurationIds

`func (o *ProvisioningSession) SetEdgeResourcesConfigurationIds(v []string)`

SetEdgeResourcesConfigurationIds sets EdgeResourcesConfigurationIds field to given value.

### HasEdgeResourcesConfigurationIds

`func (o *ProvisioningSession) HasEdgeResourcesConfigurationIds() bool`

HasEdgeResourcesConfigurationIds returns a boolean if a field has been set.

### GetEventDataProcessingConfigurationIds

`func (o *ProvisioningSession) GetEventDataProcessingConfigurationIds() []string`

GetEventDataProcessingConfigurationIds returns the EventDataProcessingConfigurationIds field if non-nil, zero value otherwise.

### GetEventDataProcessingConfigurationIdsOk

`func (o *ProvisioningSession) GetEventDataProcessingConfigurationIdsOk() (*[]string, bool)`

GetEventDataProcessingConfigurationIdsOk returns a tuple with the EventDataProcessingConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataProcessingConfigurationIds

`func (o *ProvisioningSession) SetEventDataProcessingConfigurationIds(v []string)`

SetEventDataProcessingConfigurationIds sets EventDataProcessingConfigurationIds field to given value.

### HasEventDataProcessingConfigurationIds

`func (o *ProvisioningSession) HasEventDataProcessingConfigurationIds() bool`

HasEventDataProcessingConfigurationIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


