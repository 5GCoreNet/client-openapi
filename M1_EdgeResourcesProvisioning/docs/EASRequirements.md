# EASRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasProviderIds** | **[]string** |  | 
**EasType** | **string** |  | 
**EasFeatures** | **[]string** |  | 
**ServiceKpi** | Pointer to [**EASServiceKPI**](EASServiceKPI.md) |  | [optional] 
**ServiceArea** | Pointer to [**GeographicalServiceArea**](GeographicalServiceArea.md) |  | [optional] 
**ServiceAvailabilitySchedule** | [**[]ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | 
**ServiceContinuityScenarios** | [**[]ACRScenario**](ACRScenario.md) |  | 
**ServiceContinuitySupport** | Pointer to [**[]ACRScenario**](ACRScenario.md) |  | [optional] 

## Methods

### NewEASRequirements

`func NewEASRequirements(easProviderIds []string, easType string, easFeatures []string, serviceAvailabilitySchedule []ScheduledCommunicationTime, serviceContinuityScenarios []ACRScenario, ) *EASRequirements`

NewEASRequirements instantiates a new EASRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEASRequirementsWithDefaults

`func NewEASRequirementsWithDefaults() *EASRequirements`

NewEASRequirementsWithDefaults instantiates a new EASRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasProviderIds

`func (o *EASRequirements) GetEasProviderIds() []string`

GetEasProviderIds returns the EasProviderIds field if non-nil, zero value otherwise.

### GetEasProviderIdsOk

`func (o *EASRequirements) GetEasProviderIdsOk() (*[]string, bool)`

GetEasProviderIdsOk returns a tuple with the EasProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasProviderIds

`func (o *EASRequirements) SetEasProviderIds(v []string)`

SetEasProviderIds sets EasProviderIds field to given value.


### GetEasType

`func (o *EASRequirements) GetEasType() string`

GetEasType returns the EasType field if non-nil, zero value otherwise.

### GetEasTypeOk

`func (o *EASRequirements) GetEasTypeOk() (*string, bool)`

GetEasTypeOk returns a tuple with the EasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasType

`func (o *EASRequirements) SetEasType(v string)`

SetEasType sets EasType field to given value.


### GetEasFeatures

`func (o *EASRequirements) GetEasFeatures() []string`

GetEasFeatures returns the EasFeatures field if non-nil, zero value otherwise.

### GetEasFeaturesOk

`func (o *EASRequirements) GetEasFeaturesOk() (*[]string, bool)`

GetEasFeaturesOk returns a tuple with the EasFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasFeatures

`func (o *EASRequirements) SetEasFeatures(v []string)`

SetEasFeatures sets EasFeatures field to given value.


### GetServiceKpi

`func (o *EASRequirements) GetServiceKpi() EASServiceKPI`

GetServiceKpi returns the ServiceKpi field if non-nil, zero value otherwise.

### GetServiceKpiOk

`func (o *EASRequirements) GetServiceKpiOk() (*EASServiceKPI, bool)`

GetServiceKpiOk returns a tuple with the ServiceKpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKpi

`func (o *EASRequirements) SetServiceKpi(v EASServiceKPI)`

SetServiceKpi sets ServiceKpi field to given value.

### HasServiceKpi

`func (o *EASRequirements) HasServiceKpi() bool`

HasServiceKpi returns a boolean if a field has been set.

### GetServiceArea

`func (o *EASRequirements) GetServiceArea() GeographicalServiceArea`

GetServiceArea returns the ServiceArea field if non-nil, zero value otherwise.

### GetServiceAreaOk

`func (o *EASRequirements) GetServiceAreaOk() (*GeographicalServiceArea, bool)`

GetServiceAreaOk returns a tuple with the ServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceArea

`func (o *EASRequirements) SetServiceArea(v GeographicalServiceArea)`

SetServiceArea sets ServiceArea field to given value.

### HasServiceArea

`func (o *EASRequirements) HasServiceArea() bool`

HasServiceArea returns a boolean if a field has been set.

### GetServiceAvailabilitySchedule

`func (o *EASRequirements) GetServiceAvailabilitySchedule() []ScheduledCommunicationTime`

GetServiceAvailabilitySchedule returns the ServiceAvailabilitySchedule field if non-nil, zero value otherwise.

### GetServiceAvailabilityScheduleOk

`func (o *EASRequirements) GetServiceAvailabilityScheduleOk() (*[]ScheduledCommunicationTime, bool)`

GetServiceAvailabilityScheduleOk returns a tuple with the ServiceAvailabilitySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAvailabilitySchedule

`func (o *EASRequirements) SetServiceAvailabilitySchedule(v []ScheduledCommunicationTime)`

SetServiceAvailabilitySchedule sets ServiceAvailabilitySchedule field to given value.


### GetServiceContinuityScenarios

`func (o *EASRequirements) GetServiceContinuityScenarios() []ACRScenario`

GetServiceContinuityScenarios returns the ServiceContinuityScenarios field if non-nil, zero value otherwise.

### GetServiceContinuityScenariosOk

`func (o *EASRequirements) GetServiceContinuityScenariosOk() (*[]ACRScenario, bool)`

GetServiceContinuityScenariosOk returns a tuple with the ServiceContinuityScenarios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContinuityScenarios

`func (o *EASRequirements) SetServiceContinuityScenarios(v []ACRScenario)`

SetServiceContinuityScenarios sets ServiceContinuityScenarios field to given value.


### GetServiceContinuitySupport

`func (o *EASRequirements) GetServiceContinuitySupport() []ACRScenario`

GetServiceContinuitySupport returns the ServiceContinuitySupport field if non-nil, zero value otherwise.

### GetServiceContinuitySupportOk

`func (o *EASRequirements) GetServiceContinuitySupportOk() (*[]ACRScenario, bool)`

GetServiceContinuitySupportOk returns a tuple with the ServiceContinuitySupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceContinuitySupport

`func (o *EASRequirements) SetServiceContinuitySupport(v []ACRScenario)`

SetServiceContinuitySupport sets ServiceContinuitySupport field to given value.

### HasServiceContinuitySupport

`func (o *EASRequirements) HasServiceContinuitySupport() bool`

HasServiceContinuitySupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


