# DataReportingProvisioningSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProvisioningSessionId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**AspId** | **string** | Contains an identity of an application service provider. | 
**ExternalApplicationId** | **string** | String providing an application identifier. | 
**InternalApplicationId** | Pointer to **string** | String providing an application identifier. | [optional] 
**EventId** | [**AfEvent**](AfEvent.md) |  | 
**DataReportingConfigurationIds** | **[]string** |  | 

## Methods

### NewDataReportingProvisioningSession

`func NewDataReportingProvisioningSession(provisioningSessionId string, aspId string, externalApplicationId string, eventId AfEvent, dataReportingConfigurationIds []string, ) *DataReportingProvisioningSession`

NewDataReportingProvisioningSession instantiates a new DataReportingProvisioningSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataReportingProvisioningSessionWithDefaults

`func NewDataReportingProvisioningSessionWithDefaults() *DataReportingProvisioningSession`

NewDataReportingProvisioningSessionWithDefaults instantiates a new DataReportingProvisioningSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvisioningSessionId

`func (o *DataReportingProvisioningSession) GetProvisioningSessionId() string`

GetProvisioningSessionId returns the ProvisioningSessionId field if non-nil, zero value otherwise.

### GetProvisioningSessionIdOk

`func (o *DataReportingProvisioningSession) GetProvisioningSessionIdOk() (*string, bool)`

GetProvisioningSessionIdOk returns a tuple with the ProvisioningSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSessionId

`func (o *DataReportingProvisioningSession) SetProvisioningSessionId(v string)`

SetProvisioningSessionId sets ProvisioningSessionId field to given value.


### GetAspId

`func (o *DataReportingProvisioningSession) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *DataReportingProvisioningSession) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *DataReportingProvisioningSession) SetAspId(v string)`

SetAspId sets AspId field to given value.


### GetExternalApplicationId

`func (o *DataReportingProvisioningSession) GetExternalApplicationId() string`

GetExternalApplicationId returns the ExternalApplicationId field if non-nil, zero value otherwise.

### GetExternalApplicationIdOk

`func (o *DataReportingProvisioningSession) GetExternalApplicationIdOk() (*string, bool)`

GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApplicationId

`func (o *DataReportingProvisioningSession) SetExternalApplicationId(v string)`

SetExternalApplicationId sets ExternalApplicationId field to given value.


### GetInternalApplicationId

`func (o *DataReportingProvisioningSession) GetInternalApplicationId() string`

GetInternalApplicationId returns the InternalApplicationId field if non-nil, zero value otherwise.

### GetInternalApplicationIdOk

`func (o *DataReportingProvisioningSession) GetInternalApplicationIdOk() (*string, bool)`

GetInternalApplicationIdOk returns a tuple with the InternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalApplicationId

`func (o *DataReportingProvisioningSession) SetInternalApplicationId(v string)`

SetInternalApplicationId sets InternalApplicationId field to given value.

### HasInternalApplicationId

`func (o *DataReportingProvisioningSession) HasInternalApplicationId() bool`

HasInternalApplicationId returns a boolean if a field has been set.

### GetEventId

`func (o *DataReportingProvisioningSession) GetEventId() AfEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *DataReportingProvisioningSession) GetEventIdOk() (*AfEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *DataReportingProvisioningSession) SetEventId(v AfEvent)`

SetEventId sets EventId field to given value.


### GetDataReportingConfigurationIds

`func (o *DataReportingProvisioningSession) GetDataReportingConfigurationIds() []string`

GetDataReportingConfigurationIds returns the DataReportingConfigurationIds field if non-nil, zero value otherwise.

### GetDataReportingConfigurationIdsOk

`func (o *DataReportingProvisioningSession) GetDataReportingConfigurationIdsOk() (*[]string, bool)`

GetDataReportingConfigurationIdsOk returns a tuple with the DataReportingConfigurationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReportingConfigurationIds

`func (o *DataReportingProvisioningSession) SetDataReportingConfigurationIds(v []string)`

SetDataReportingConfigurationIds sets DataReportingConfigurationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


