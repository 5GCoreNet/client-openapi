# RacsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**RacsConfigs** | [**map[string]RacsConfiguration**](RacsConfiguration.md) | Identifies the configuration related to manufacturer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key. The response shall include successfully provisioned RACS data.  | 
**RacsReports** | Pointer to [**map[string]RacsFailureReport**](RacsFailureReport.md) | Contains the RACS IDs for which the RACS data are not provisioned successfully. The failure reason is also included. Any string value can be used as a key of the map.  | [optional] [readonly] 

## Methods

### NewRacsData

`func NewRacsData(racsConfigs map[string]RacsConfiguration, ) *RacsData`

NewRacsData instantiates a new RacsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacsDataWithDefaults

`func NewRacsDataWithDefaults() *RacsData`

NewRacsDataWithDefaults instantiates a new RacsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuppFeat

`func (o *RacsData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *RacsData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *RacsData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *RacsData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetRacsConfigs

`func (o *RacsData) GetRacsConfigs() map[string]RacsConfiguration`

GetRacsConfigs returns the RacsConfigs field if non-nil, zero value otherwise.

### GetRacsConfigsOk

`func (o *RacsData) GetRacsConfigsOk() (*map[string]RacsConfiguration, bool)`

GetRacsConfigsOk returns a tuple with the RacsConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsConfigs

`func (o *RacsData) SetRacsConfigs(v map[string]RacsConfiguration)`

SetRacsConfigs sets RacsConfigs field to given value.


### GetRacsReports

`func (o *RacsData) GetRacsReports() map[string]RacsFailureReport`

GetRacsReports returns the RacsReports field if non-nil, zero value otherwise.

### GetRacsReportsOk

`func (o *RacsData) GetRacsReportsOk() (*map[string]RacsFailureReport, bool)`

GetRacsReportsOk returns a tuple with the RacsReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacsReports

`func (o *RacsData) SetRacsReports(v map[string]RacsFailureReport)`

SetRacsReports sets RacsReports field to given value.

### HasRacsReports

`func (o *RacsData) HasRacsReports() bool`

HasRacsReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


