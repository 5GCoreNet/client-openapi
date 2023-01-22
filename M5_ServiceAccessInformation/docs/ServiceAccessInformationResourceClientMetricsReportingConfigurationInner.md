# ServiceAccessInformationResourceClientMetricsReportingConfigurationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAddresses** | **[]string** | A set of application endpoint addresses. | 
**Scheme** | **string** | String providing an URI formatted according to RFC 3986. | 
**DataNetworkName** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**ReportingInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SamplePercentage** | **float32** |  | 
**UrlFilters** | **[]string** |  | 
**Metrics** | **[]string** |  | 

## Methods

### NewServiceAccessInformationResourceClientMetricsReportingConfigurationInner

`func NewServiceAccessInformationResourceClientMetricsReportingConfigurationInner(serverAddresses []string, scheme string, samplePercentage float32, urlFilters []string, metrics []string, ) *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner`

NewServiceAccessInformationResourceClientMetricsReportingConfigurationInner instantiates a new ServiceAccessInformationResourceClientMetricsReportingConfigurationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccessInformationResourceClientMetricsReportingConfigurationInnerWithDefaults

`func NewServiceAccessInformationResourceClientMetricsReportingConfigurationInnerWithDefaults() *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner`

NewServiceAccessInformationResourceClientMetricsReportingConfigurationInnerWithDefaults instantiates a new ServiceAccessInformationResourceClientMetricsReportingConfigurationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerAddresses

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetServerAddresses() []string`

GetServerAddresses returns the ServerAddresses field if non-nil, zero value otherwise.

### GetServerAddressesOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetServerAddressesOk() (*[]string, bool)`

GetServerAddressesOk returns a tuple with the ServerAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddresses

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetServerAddresses(v []string)`

SetServerAddresses sets ServerAddresses field to given value.


### GetScheme

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetDataNetworkName

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetDataNetworkName() string`

GetDataNetworkName returns the DataNetworkName field if non-nil, zero value otherwise.

### GetDataNetworkNameOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetDataNetworkNameOk() (*string, bool)`

GetDataNetworkNameOk returns a tuple with the DataNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNetworkName

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetDataNetworkName(v string)`

SetDataNetworkName sets DataNetworkName field to given value.

### HasDataNetworkName

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) HasDataNetworkName() bool`

HasDataNetworkName returns a boolean if a field has been set.

### GetReportingInterval

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetReportingInterval() int32`

GetReportingInterval returns the ReportingInterval field if non-nil, zero value otherwise.

### GetReportingIntervalOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetReportingIntervalOk() (*int32, bool)`

GetReportingIntervalOk returns a tuple with the ReportingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInterval

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetReportingInterval(v int32)`

SetReportingInterval sets ReportingInterval field to given value.

### HasReportingInterval

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) HasReportingInterval() bool`

HasReportingInterval returns a boolean if a field has been set.

### GetSamplePercentage

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetSamplePercentage() float32`

GetSamplePercentage returns the SamplePercentage field if non-nil, zero value otherwise.

### GetSamplePercentageOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetSamplePercentageOk() (*float32, bool)`

GetSamplePercentageOk returns a tuple with the SamplePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplePercentage

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetSamplePercentage(v float32)`

SetSamplePercentage sets SamplePercentage field to given value.


### GetUrlFilters

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetUrlFilters() []string`

GetUrlFilters returns the UrlFilters field if non-nil, zero value otherwise.

### GetUrlFiltersOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetUrlFiltersOk() (*[]string, bool)`

GetUrlFiltersOk returns a tuple with the UrlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFilters

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetUrlFilters(v []string)`

SetUrlFilters sets UrlFilters field to given value.


### GetMetrics

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


