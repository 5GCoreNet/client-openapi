# DataReportingConfigurationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationURL** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 
**DataAccessProfiles** | Pointer to [**[]DataAccessProfile**](DataAccessProfile.md) |  | [optional] 

## Methods

### NewDataReportingConfigurationPatch

`func NewDataReportingConfigurationPatch() *DataReportingConfigurationPatch`

NewDataReportingConfigurationPatch instantiates a new DataReportingConfigurationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataReportingConfigurationPatchWithDefaults

`func NewDataReportingConfigurationPatchWithDefaults() *DataReportingConfigurationPatch`

NewDataReportingConfigurationPatchWithDefaults instantiates a new DataReportingConfigurationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationURL

`func (o *DataReportingConfigurationPatch) GetAuthorizationURL() string`

GetAuthorizationURL returns the AuthorizationURL field if non-nil, zero value otherwise.

### GetAuthorizationURLOk

`func (o *DataReportingConfigurationPatch) GetAuthorizationURLOk() (*string, bool)`

GetAuthorizationURLOk returns a tuple with the AuthorizationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationURL

`func (o *DataReportingConfigurationPatch) SetAuthorizationURL(v string)`

SetAuthorizationURL sets AuthorizationURL field to given value.

### HasAuthorizationURL

`func (o *DataReportingConfigurationPatch) HasAuthorizationURL() bool`

HasAuthorizationURL returns a boolean if a field has been set.

### GetDataAccessProfiles

`func (o *DataReportingConfigurationPatch) GetDataAccessProfiles() []DataAccessProfile`

GetDataAccessProfiles returns the DataAccessProfiles field if non-nil, zero value otherwise.

### GetDataAccessProfilesOk

`func (o *DataReportingConfigurationPatch) GetDataAccessProfilesOk() (*[]DataAccessProfile, bool)`

GetDataAccessProfilesOk returns a tuple with the DataAccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessProfiles

`func (o *DataReportingConfigurationPatch) SetDataAccessProfiles(v []DataAccessProfile)`

SetDataAccessProfiles sets DataAccessProfiles field to given value.

### HasDataAccessProfiles

`func (o *DataReportingConfigurationPatch) HasDataAccessProfiles() bool`

HasDataAccessProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


