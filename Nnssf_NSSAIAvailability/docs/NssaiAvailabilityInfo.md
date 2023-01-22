# NssaiAvailabilityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedNssaiAvailabilityData** | [**[]SupportedNssaiAvailabilityData**](SupportedNssaiAvailabilityData.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**AmfSetId** | Pointer to **string** |  | [optional] 

## Methods

### NewNssaiAvailabilityInfo

`func NewNssaiAvailabilityInfo(supportedNssaiAvailabilityData []SupportedNssaiAvailabilityData, ) *NssaiAvailabilityInfo`

NewNssaiAvailabilityInfo instantiates a new NssaiAvailabilityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNssaiAvailabilityInfoWithDefaults

`func NewNssaiAvailabilityInfoWithDefaults() *NssaiAvailabilityInfo`

NewNssaiAvailabilityInfoWithDefaults instantiates a new NssaiAvailabilityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedNssaiAvailabilityData

`func (o *NssaiAvailabilityInfo) GetSupportedNssaiAvailabilityData() []SupportedNssaiAvailabilityData`

GetSupportedNssaiAvailabilityData returns the SupportedNssaiAvailabilityData field if non-nil, zero value otherwise.

### GetSupportedNssaiAvailabilityDataOk

`func (o *NssaiAvailabilityInfo) GetSupportedNssaiAvailabilityDataOk() (*[]SupportedNssaiAvailabilityData, bool)`

GetSupportedNssaiAvailabilityDataOk returns a tuple with the SupportedNssaiAvailabilityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedNssaiAvailabilityData

`func (o *NssaiAvailabilityInfo) SetSupportedNssaiAvailabilityData(v []SupportedNssaiAvailabilityData)`

SetSupportedNssaiAvailabilityData sets SupportedNssaiAvailabilityData field to given value.


### GetSupportedFeatures

`func (o *NssaiAvailabilityInfo) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NssaiAvailabilityInfo) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NssaiAvailabilityInfo) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NssaiAvailabilityInfo) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetAmfSetId

`func (o *NssaiAvailabilityInfo) GetAmfSetId() string`

GetAmfSetId returns the AmfSetId field if non-nil, zero value otherwise.

### GetAmfSetIdOk

`func (o *NssaiAvailabilityInfo) GetAmfSetIdOk() (*string, bool)`

GetAmfSetIdOk returns a tuple with the AmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSetId

`func (o *NssaiAvailabilityInfo) SetAmfSetId(v string)`

SetAmfSetId sets AmfSetId field to given value.

### HasAmfSetId

`func (o *NssaiAvailabilityInfo) HasAmfSetId() bool`

HasAmfSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


