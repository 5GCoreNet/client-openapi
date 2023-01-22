# DataReportingSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | [optional] 
**ValidUntil** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ExternalApplicationId** | **string** | String providing an application identifier. | 
**SupportedDomains** | [**[]DataDomain**](DataDomain.md) |  | 
**ReportingConditions** | Pointer to [**DataReportingSessionReportingConditions**](DataReportingSessionReportingConditions.md) |  | [optional] 

## Methods

### NewDataReportingSession

`func NewDataReportingSession(externalApplicationId string, supportedDomains []DataDomain, ) *DataReportingSession`

NewDataReportingSession instantiates a new DataReportingSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataReportingSessionWithDefaults

`func NewDataReportingSessionWithDefaults() *DataReportingSession`

NewDataReportingSessionWithDefaults instantiates a new DataReportingSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *DataReportingSession) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DataReportingSession) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DataReportingSession) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DataReportingSession) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetValidUntil

`func (o *DataReportingSession) GetValidUntil() time.Time`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *DataReportingSession) GetValidUntilOk() (*time.Time, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *DataReportingSession) SetValidUntil(v time.Time)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *DataReportingSession) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetExternalApplicationId

`func (o *DataReportingSession) GetExternalApplicationId() string`

GetExternalApplicationId returns the ExternalApplicationId field if non-nil, zero value otherwise.

### GetExternalApplicationIdOk

`func (o *DataReportingSession) GetExternalApplicationIdOk() (*string, bool)`

GetExternalApplicationIdOk returns a tuple with the ExternalApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalApplicationId

`func (o *DataReportingSession) SetExternalApplicationId(v string)`

SetExternalApplicationId sets ExternalApplicationId field to given value.


### GetSupportedDomains

`func (o *DataReportingSession) GetSupportedDomains() []DataDomain`

GetSupportedDomains returns the SupportedDomains field if non-nil, zero value otherwise.

### GetSupportedDomainsOk

`func (o *DataReportingSession) GetSupportedDomainsOk() (*[]DataDomain, bool)`

GetSupportedDomainsOk returns a tuple with the SupportedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDomains

`func (o *DataReportingSession) SetSupportedDomains(v []DataDomain)`

SetSupportedDomains sets SupportedDomains field to given value.


### GetReportingConditions

`func (o *DataReportingSession) GetReportingConditions() DataReportingSessionReportingConditions`

GetReportingConditions returns the ReportingConditions field if non-nil, zero value otherwise.

### GetReportingConditionsOk

`func (o *DataReportingSession) GetReportingConditionsOk() (*DataReportingSessionReportingConditions, bool)`

GetReportingConditionsOk returns a tuple with the ReportingConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingConditions

`func (o *DataReportingSession) SetReportingConditions(v DataReportingSessionReportingConditions)`

SetReportingConditions sets ReportingConditions field to given value.

### HasReportingConditions

`func (o *DataReportingSession) HasReportingConditions() bool`

HasReportingConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


