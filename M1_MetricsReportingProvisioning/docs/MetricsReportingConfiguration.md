# MetricsReportingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricsReportingConfigurationId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**Scheme** | **string** | String providing an URI formatted according to RFC 3986. | 
**DataNetworkName** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**ReportingInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SamplePercentage** | Pointer to **float32** |  | [optional] 
**UrlFilters** | Pointer to **[]string** |  | [optional] 
**Metrics** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMetricsReportingConfiguration

`func NewMetricsReportingConfiguration(metricsReportingConfigurationId string, scheme string, ) *MetricsReportingConfiguration`

NewMetricsReportingConfiguration instantiates a new MetricsReportingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsReportingConfigurationWithDefaults

`func NewMetricsReportingConfigurationWithDefaults() *MetricsReportingConfiguration`

NewMetricsReportingConfigurationWithDefaults instantiates a new MetricsReportingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricsReportingConfigurationId

`func (o *MetricsReportingConfiguration) GetMetricsReportingConfigurationId() string`

GetMetricsReportingConfigurationId returns the MetricsReportingConfigurationId field if non-nil, zero value otherwise.

### GetMetricsReportingConfigurationIdOk

`func (o *MetricsReportingConfiguration) GetMetricsReportingConfigurationIdOk() (*string, bool)`

GetMetricsReportingConfigurationIdOk returns a tuple with the MetricsReportingConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsReportingConfigurationId

`func (o *MetricsReportingConfiguration) SetMetricsReportingConfigurationId(v string)`

SetMetricsReportingConfigurationId sets MetricsReportingConfigurationId field to given value.


### GetScheme

`func (o *MetricsReportingConfiguration) GetScheme() string`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *MetricsReportingConfiguration) GetSchemeOk() (*string, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *MetricsReportingConfiguration) SetScheme(v string)`

SetScheme sets Scheme field to given value.


### GetDataNetworkName

`func (o *MetricsReportingConfiguration) GetDataNetworkName() string`

GetDataNetworkName returns the DataNetworkName field if non-nil, zero value otherwise.

### GetDataNetworkNameOk

`func (o *MetricsReportingConfiguration) GetDataNetworkNameOk() (*string, bool)`

GetDataNetworkNameOk returns a tuple with the DataNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataNetworkName

`func (o *MetricsReportingConfiguration) SetDataNetworkName(v string)`

SetDataNetworkName sets DataNetworkName field to given value.

### HasDataNetworkName

`func (o *MetricsReportingConfiguration) HasDataNetworkName() bool`

HasDataNetworkName returns a boolean if a field has been set.

### GetReportingInterval

`func (o *MetricsReportingConfiguration) GetReportingInterval() int32`

GetReportingInterval returns the ReportingInterval field if non-nil, zero value otherwise.

### GetReportingIntervalOk

`func (o *MetricsReportingConfiguration) GetReportingIntervalOk() (*int32, bool)`

GetReportingIntervalOk returns a tuple with the ReportingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInterval

`func (o *MetricsReportingConfiguration) SetReportingInterval(v int32)`

SetReportingInterval sets ReportingInterval field to given value.

### HasReportingInterval

`func (o *MetricsReportingConfiguration) HasReportingInterval() bool`

HasReportingInterval returns a boolean if a field has been set.

### GetSamplePercentage

`func (o *MetricsReportingConfiguration) GetSamplePercentage() float32`

GetSamplePercentage returns the SamplePercentage field if non-nil, zero value otherwise.

### GetSamplePercentageOk

`func (o *MetricsReportingConfiguration) GetSamplePercentageOk() (*float32, bool)`

GetSamplePercentageOk returns a tuple with the SamplePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplePercentage

`func (o *MetricsReportingConfiguration) SetSamplePercentage(v float32)`

SetSamplePercentage sets SamplePercentage field to given value.

### HasSamplePercentage

`func (o *MetricsReportingConfiguration) HasSamplePercentage() bool`

HasSamplePercentage returns a boolean if a field has been set.

### GetUrlFilters

`func (o *MetricsReportingConfiguration) GetUrlFilters() []string`

GetUrlFilters returns the UrlFilters field if non-nil, zero value otherwise.

### GetUrlFiltersOk

`func (o *MetricsReportingConfiguration) GetUrlFiltersOk() (*[]string, bool)`

GetUrlFiltersOk returns a tuple with the UrlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFilters

`func (o *MetricsReportingConfiguration) SetUrlFilters(v []string)`

SetUrlFilters sets UrlFilters field to given value.

### HasUrlFilters

`func (o *MetricsReportingConfiguration) HasUrlFilters() bool`

HasUrlFilters returns a boolean if a field has been set.

### GetMetrics

`func (o *MetricsReportingConfiguration) GetMetrics() []string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricsReportingConfiguration) GetMetricsOk() (*[]string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricsReportingConfiguration) SetMetrics(v []string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricsReportingConfiguration) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


