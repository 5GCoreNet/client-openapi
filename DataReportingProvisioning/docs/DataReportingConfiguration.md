# DataReportingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataReportingConfigurationId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**DataCollectionClientType** | [**DataCollectionClientType**](DataCollectionClientType.md) |  | 
**AuthorizationURL** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 
**DataAccessProfiles** | [**[]DataAccessProfile**](DataAccessProfile.md) |  | 

## Methods

### NewDataReportingConfiguration

`func NewDataReportingConfiguration(dataReportingConfigurationId string, dataCollectionClientType DataCollectionClientType, dataAccessProfiles []DataAccessProfile, ) *DataReportingConfiguration`

NewDataReportingConfiguration instantiates a new DataReportingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataReportingConfigurationWithDefaults

`func NewDataReportingConfigurationWithDefaults() *DataReportingConfiguration`

NewDataReportingConfigurationWithDefaults instantiates a new DataReportingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataReportingConfigurationId

`func (o *DataReportingConfiguration) GetDataReportingConfigurationId() string`

GetDataReportingConfigurationId returns the DataReportingConfigurationId field if non-nil, zero value otherwise.

### GetDataReportingConfigurationIdOk

`func (o *DataReportingConfiguration) GetDataReportingConfigurationIdOk() (*string, bool)`

GetDataReportingConfigurationIdOk returns a tuple with the DataReportingConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataReportingConfigurationId

`func (o *DataReportingConfiguration) SetDataReportingConfigurationId(v string)`

SetDataReportingConfigurationId sets DataReportingConfigurationId field to given value.


### GetDataCollectionClientType

`func (o *DataReportingConfiguration) GetDataCollectionClientType() DataCollectionClientType`

GetDataCollectionClientType returns the DataCollectionClientType field if non-nil, zero value otherwise.

### GetDataCollectionClientTypeOk

`func (o *DataReportingConfiguration) GetDataCollectionClientTypeOk() (*DataCollectionClientType, bool)`

GetDataCollectionClientTypeOk returns a tuple with the DataCollectionClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataCollectionClientType

`func (o *DataReportingConfiguration) SetDataCollectionClientType(v DataCollectionClientType)`

SetDataCollectionClientType sets DataCollectionClientType field to given value.


### GetAuthorizationURL

`func (o *DataReportingConfiguration) GetAuthorizationURL() string`

GetAuthorizationURL returns the AuthorizationURL field if non-nil, zero value otherwise.

### GetAuthorizationURLOk

`func (o *DataReportingConfiguration) GetAuthorizationURLOk() (*string, bool)`

GetAuthorizationURLOk returns a tuple with the AuthorizationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationURL

`func (o *DataReportingConfiguration) SetAuthorizationURL(v string)`

SetAuthorizationURL sets AuthorizationURL field to given value.

### HasAuthorizationURL

`func (o *DataReportingConfiguration) HasAuthorizationURL() bool`

HasAuthorizationURL returns a boolean if a field has been set.

### GetDataAccessProfiles

`func (o *DataReportingConfiguration) GetDataAccessProfiles() []DataAccessProfile`

GetDataAccessProfiles returns the DataAccessProfiles field if non-nil, zero value otherwise.

### GetDataAccessProfilesOk

`func (o *DataReportingConfiguration) GetDataAccessProfilesOk() (*[]DataAccessProfile, bool)`

GetDataAccessProfilesOk returns a tuple with the DataAccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessProfiles

`func (o *DataReportingConfiguration) SetDataAccessProfiles(v []DataAccessProfile)`

SetDataAccessProfiles sets DataAccessProfiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


