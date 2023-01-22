# ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAddresses** | **[]string** | A set of application endpoint addresses. | 
**ValidPolicyTemplateIds** | **[]string** |  | 
**SdfMethods** | [**[]SdfMethod**](SdfMethod.md) |  | 
**ExternalReferences** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceAccessInformationResourceDynamicPolicyInvocationConfiguration

`func NewServiceAccessInformationResourceDynamicPolicyInvocationConfiguration(serverAddresses []string, validPolicyTemplateIds []string, sdfMethods []SdfMethod, ) *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration`

NewServiceAccessInformationResourceDynamicPolicyInvocationConfiguration instantiates a new ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessInformationResourceDynamicPolicyInvocationConfigurationWithDefaults

`func NewServiceAccessInformationResourceDynamicPolicyInvocationConfigurationWithDefaults() *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration`

NewServiceAccessInformationResourceDynamicPolicyInvocationConfigurationWithDefaults instantiates a new ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerAddresses

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetServerAddresses() []string`

GetServerAddresses returns the ServerAddresses field if non-nil, zero value otherwise.

### GetServerAddressesOk

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetServerAddressesOk() (*[]string, bool)`

GetServerAddressesOk returns a tuple with the ServerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddresses

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetServerAddresses(v []string)`

SetServerAddresses sets ServerAddresses field to given value.


### GetValidPolicyTemplateIds

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetValidPolicyTemplateIds() []string`

GetValidPolicyTemplateIds returns the ValidPolicyTemplateIds field if non-nil, zero value otherwise.

### GetValidPolicyTemplateIdsOk

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetValidPolicyTemplateIdsOk() (*[]string, bool)`

GetValidPolicyTemplateIdsOk returns a tuple with the ValidPolicyTemplateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPolicyTemplateIds

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetValidPolicyTemplateIds(v []string)`

SetValidPolicyTemplateIds sets ValidPolicyTemplateIds field to given value.


### GetSdfMethods

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetSdfMethods() []SdfMethod`

GetSdfMethods returns the SdfMethods field if non-nil, zero value otherwise.

### GetSdfMethodsOk

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetSdfMethodsOk() (*[]SdfMethod, bool)`

GetSdfMethodsOk returns a tuple with the SdfMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdfMethods

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetSdfMethods(v []SdfMethod)`

SetSdfMethods sets SdfMethods field to given value.


### GetExternalReferences

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetExternalReferences() []string`

GetExternalReferences returns the ExternalReferences field if non-nil, zero value otherwise.

### GetExternalReferencesOk

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetExternalReferencesOk() (*[]string, bool)`

GetExternalReferencesOk returns a tuple with the ExternalReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReferences

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetExternalReferences(v []string)`

SetExternalReferences sets ExternalReferences field to given value.

### HasExternalReferences

`func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) HasExternalReferences() bool`

HasExternalReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


