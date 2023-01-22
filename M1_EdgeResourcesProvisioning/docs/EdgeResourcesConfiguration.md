# EdgeResourcesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgeResourcesConfigurationId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**EdgeManagementMode** | [**EdgeManagementMode**](EdgeManagementMode.md) |  | 
**EligibilityCriteria** | Pointer to [**EdgeProcessingEligibilityCriteria**](EdgeProcessingEligibilityCriteria.md) |  | [optional] 
**EasRequirements** | [**EASRequirements**](EASRequirements.md) |  | 
**EasRelocationRequirements** | Pointer to [**M1EASRelocationRequirements**](M1EASRelocationRequirements.md) |  | [optional] 

## Methods

### NewEdgeResourcesConfiguration

`func NewEdgeResourcesConfiguration(edgeResourcesConfigurationId string, edgeManagementMode EdgeManagementMode, easRequirements EASRequirements, ) *EdgeResourcesConfiguration`

NewEdgeResourcesConfiguration instantiates a new EdgeResourcesConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeResourcesConfigurationWithDefaults

`func NewEdgeResourcesConfigurationWithDefaults() *EdgeResourcesConfiguration`

NewEdgeResourcesConfigurationWithDefaults instantiates a new EdgeResourcesConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdgeResourcesConfigurationId

`func (o *EdgeResourcesConfiguration) GetEdgeResourcesConfigurationId() string`

GetEdgeResourcesConfigurationId returns the EdgeResourcesConfigurationId field if non-nil, zero value otherwise.

### GetEdgeResourcesConfigurationIdOk

`func (o *EdgeResourcesConfiguration) GetEdgeResourcesConfigurationIdOk() (*string, bool)`

GetEdgeResourcesConfigurationIdOk returns a tuple with the EdgeResourcesConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeResourcesConfigurationId

`func (o *EdgeResourcesConfiguration) SetEdgeResourcesConfigurationId(v string)`

SetEdgeResourcesConfigurationId sets EdgeResourcesConfigurationId field to given value.


### GetEdgeManagementMode

`func (o *EdgeResourcesConfiguration) GetEdgeManagementMode() EdgeManagementMode`

GetEdgeManagementMode returns the EdgeManagementMode field if non-nil, zero value otherwise.

### GetEdgeManagementModeOk

`func (o *EdgeResourcesConfiguration) GetEdgeManagementModeOk() (*EdgeManagementMode, bool)`

GetEdgeManagementModeOk returns a tuple with the EdgeManagementMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdgeManagementMode

`func (o *EdgeResourcesConfiguration) SetEdgeManagementMode(v EdgeManagementMode)`

SetEdgeManagementMode sets EdgeManagementMode field to given value.


### GetEligibilityCriteria

`func (o *EdgeResourcesConfiguration) GetEligibilityCriteria() EdgeProcessingEligibilityCriteria`

GetEligibilityCriteria returns the EligibilityCriteria field if non-nil, zero value otherwise.

### GetEligibilityCriteriaOk

`func (o *EdgeResourcesConfiguration) GetEligibilityCriteriaOk() (*EdgeProcessingEligibilityCriteria, bool)`

GetEligibilityCriteriaOk returns a tuple with the EligibilityCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityCriteria

`func (o *EdgeResourcesConfiguration) SetEligibilityCriteria(v EdgeProcessingEligibilityCriteria)`

SetEligibilityCriteria sets EligibilityCriteria field to given value.

### HasEligibilityCriteria

`func (o *EdgeResourcesConfiguration) HasEligibilityCriteria() bool`

HasEligibilityCriteria returns a boolean if a field has been set.

### GetEasRequirements

`func (o *EdgeResourcesConfiguration) GetEasRequirements() EASRequirements`

GetEasRequirements returns the EasRequirements field if non-nil, zero value otherwise.

### GetEasRequirementsOk

`func (o *EdgeResourcesConfiguration) GetEasRequirementsOk() (*EASRequirements, bool)`

GetEasRequirementsOk returns a tuple with the EasRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasRequirements

`func (o *EdgeResourcesConfiguration) SetEasRequirements(v EASRequirements)`

SetEasRequirements sets EasRequirements field to given value.


### GetEasRelocationRequirements

`func (o *EdgeResourcesConfiguration) GetEasRelocationRequirements() M1EASRelocationRequirements`

GetEasRelocationRequirements returns the EasRelocationRequirements field if non-nil, zero value otherwise.

### GetEasRelocationRequirementsOk

`func (o *EdgeResourcesConfiguration) GetEasRelocationRequirementsOk() (*M1EASRelocationRequirements, bool)`

GetEasRelocationRequirementsOk returns a tuple with the EasRelocationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasRelocationRequirements

`func (o *EdgeResourcesConfiguration) SetEasRelocationRequirements(v M1EASRelocationRequirements)`

SetEasRelocationRequirements sets EasRelocationRequirements field to given value.

### HasEasRelocationRequirements

`func (o *EdgeResourcesConfiguration) HasEasRelocationRequirements() bool`

HasEasRelocationRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


