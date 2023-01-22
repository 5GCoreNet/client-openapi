# ServiceAccessInformationResourceClientConsumptionReportingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportingInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**ServerAddresses** | **[]string** | A set of application endpoint addresses. | 
**LocationReporting** | **bool** |  | 
**AccessReporting** | Pointer to **bool** |  | [optional] 
**SamplePercentage** | **float32** |  | 

## Methods

### NewServiceAccessInformationResourceClientConsumptionReportingConfiguration

`func NewServiceAccessInformationResourceClientConsumptionReportingConfiguration(serverAddresses []string, locationReporting bool, samplePercentage float32, ) *ServiceAccessInformationResourceClientConsumptionReportingConfiguration`

NewServiceAccessInformationResourceClientConsumptionReportingConfiguration instantiates a new ServiceAccessInformationResourceClientConsumptionReportingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessInformationResourceClientConsumptionReportingConfigurationWithDefaults

`func NewServiceAccessInformationResourceClientConsumptionReportingConfigurationWithDefaults() *ServiceAccessInformationResourceClientConsumptionReportingConfiguration`

NewServiceAccessInformationResourceClientConsumptionReportingConfigurationWithDefaults instantiates a new ServiceAccessInformationResourceClientConsumptionReportingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportingInterval

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetReportingInterval() int32`

GetReportingInterval returns the ReportingInterval field if non-nil, zero value otherwise.

### GetReportingIntervalOk

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetReportingIntervalOk() (*int32, bool)`

GetReportingIntervalOk returns a tuple with the ReportingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInterval

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetReportingInterval(v int32)`

SetReportingInterval sets ReportingInterval field to given value.

### HasReportingInterval

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) HasReportingInterval() bool`

HasReportingInterval returns a boolean if a field has been set.

### GetServerAddresses

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetServerAddresses() []string`

GetServerAddresses returns the ServerAddresses field if non-nil, zero value otherwise.

### GetServerAddressesOk

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetServerAddressesOk() (*[]string, bool)`

GetServerAddressesOk returns a tuple with the ServerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddresses

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetServerAddresses(v []string)`

SetServerAddresses sets ServerAddresses field to given value.


### GetLocationReporting

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetLocationReporting() bool`

GetLocationReporting returns the LocationReporting field if non-nil, zero value otherwise.

### GetLocationReportingOk

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetLocationReportingOk() (*bool, bool)`

GetLocationReportingOk returns a tuple with the LocationReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationReporting

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetLocationReporting(v bool)`

SetLocationReporting sets LocationReporting field to given value.


### GetAccessReporting

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetAccessReporting() bool`

GetAccessReporting returns the AccessReporting field if non-nil, zero value otherwise.

### GetAccessReportingOk

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetAccessReportingOk() (*bool, bool)`

GetAccessReportingOk returns a tuple with the AccessReporting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessReporting

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetAccessReporting(v bool)`

SetAccessReporting sets AccessReporting field to given value.

### HasAccessReporting

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) HasAccessReporting() bool`

HasAccessReporting returns a boolean if a field has been set.

### GetSamplePercentage

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetSamplePercentage() float32`

GetSamplePercentage returns the SamplePercentage field if non-nil, zero value otherwise.

### GetSamplePercentageOk

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) GetSamplePercentageOk() (*float32, bool)`

GetSamplePercentageOk returns a tuple with the SamplePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplePercentage

`func (o *ServiceAccessInformationResourceClientConsumptionReportingConfiguration) SetSamplePercentage(v float32)`

SetSamplePercentage sets SamplePercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


