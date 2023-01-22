# SearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidityPeriod** | Pointer to **int32** |  | [optional] 
**NfInstances** | [**[]NFProfile**](NFProfile.md) |  | 
**SearchId** | Pointer to **string** |  | [optional] 
**NumNfInstComplete** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**PreferredSearch** | Pointer to [**PreferredSearch**](PreferredSearch.md) |  | [optional] 
**NrfSupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**NfInstanceList** | Pointer to [**map[string]NfInstanceInfo**](NfInstanceInfo.md) | List of matching NF instances. The key of the map is the NF instance ID. | [optional] 
**AlteredPriorityInd** | Pointer to **bool** |  | [optional] 
**NoProfileMatchInfo** | Pointer to [**NoProfileMatchInfo**](NoProfileMatchInfo.md) |  | [optional] 

## Methods

### NewSearchResult

`func NewSearchResult(nfInstances []NFProfile, ) *SearchResult`

NewSearchResult instantiates a new SearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultWithDefaults

`func NewSearchResultWithDefaults() *SearchResult`

NewSearchResultWithDefaults instantiates a new SearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidityPeriod

`func (o *SearchResult) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *SearchResult) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *SearchResult) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.

### HasValidityPeriod

`func (o *SearchResult) HasValidityPeriod() bool`

HasValidityPeriod returns a boolean if a field has been set.

### GetNfInstances

`func (o *SearchResult) GetNfInstances() []NFProfile`

GetNfInstances returns the NfInstances field if non-nil, zero value otherwise.

### GetNfInstancesOk

`func (o *SearchResult) GetNfInstancesOk() (*[]NFProfile, bool)`

GetNfInstancesOk returns a tuple with the NfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstances

`func (o *SearchResult) SetNfInstances(v []NFProfile)`

SetNfInstances sets NfInstances field to given value.


### GetSearchId

`func (o *SearchResult) GetSearchId() string`

GetSearchId returns the SearchId field if non-nil, zero value otherwise.

### GetSearchIdOk

`func (o *SearchResult) GetSearchIdOk() (*string, bool)`

GetSearchIdOk returns a tuple with the SearchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchId

`func (o *SearchResult) SetSearchId(v string)`

SetSearchId sets SearchId field to given value.

### HasSearchId

`func (o *SearchResult) HasSearchId() bool`

HasSearchId returns a boolean if a field has been set.

### GetNumNfInstComplete

`func (o *SearchResult) GetNumNfInstComplete() int32`

GetNumNfInstComplete returns the NumNfInstComplete field if non-nil, zero value otherwise.

### GetNumNfInstCompleteOk

`func (o *SearchResult) GetNumNfInstCompleteOk() (*int32, bool)`

GetNumNfInstCompleteOk returns a tuple with the NumNfInstComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumNfInstComplete

`func (o *SearchResult) SetNumNfInstComplete(v int32)`

SetNumNfInstComplete sets NumNfInstComplete field to given value.

### HasNumNfInstComplete

`func (o *SearchResult) HasNumNfInstComplete() bool`

HasNumNfInstComplete returns a boolean if a field has been set.

### GetPreferredSearch

`func (o *SearchResult) GetPreferredSearch() PreferredSearch`

GetPreferredSearch returns the PreferredSearch field if non-nil, zero value otherwise.

### GetPreferredSearchOk

`func (o *SearchResult) GetPreferredSearchOk() (*PreferredSearch, bool)`

GetPreferredSearchOk returns a tuple with the PreferredSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSearch

`func (o *SearchResult) SetPreferredSearch(v PreferredSearch)`

SetPreferredSearch sets PreferredSearch field to given value.

### HasPreferredSearch

`func (o *SearchResult) HasPreferredSearch() bool`

HasPreferredSearch returns a boolean if a field has been set.

### GetNrfSupportedFeatures

`func (o *SearchResult) GetNrfSupportedFeatures() string`

GetNrfSupportedFeatures returns the NrfSupportedFeatures field if non-nil, zero value otherwise.

### GetNrfSupportedFeaturesOk

`func (o *SearchResult) GetNrfSupportedFeaturesOk() (*string, bool)`

GetNrfSupportedFeaturesOk returns a tuple with the NrfSupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfSupportedFeatures

`func (o *SearchResult) SetNrfSupportedFeatures(v string)`

SetNrfSupportedFeatures sets NrfSupportedFeatures field to given value.

### HasNrfSupportedFeatures

`func (o *SearchResult) HasNrfSupportedFeatures() bool`

HasNrfSupportedFeatures returns a boolean if a field has been set.

### GetNfInstanceList

`func (o *SearchResult) GetNfInstanceList() map[string]NfInstanceInfo`

GetNfInstanceList returns the NfInstanceList field if non-nil, zero value otherwise.

### GetNfInstanceListOk

`func (o *SearchResult) GetNfInstanceListOk() (*map[string]NfInstanceInfo, bool)`

GetNfInstanceListOk returns a tuple with the NfInstanceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceList

`func (o *SearchResult) SetNfInstanceList(v map[string]NfInstanceInfo)`

SetNfInstanceList sets NfInstanceList field to given value.

### HasNfInstanceList

`func (o *SearchResult) HasNfInstanceList() bool`

HasNfInstanceList returns a boolean if a field has been set.

### GetAlteredPriorityInd

`func (o *SearchResult) GetAlteredPriorityInd() bool`

GetAlteredPriorityInd returns the AlteredPriorityInd field if non-nil, zero value otherwise.

### GetAlteredPriorityIndOk

`func (o *SearchResult) GetAlteredPriorityIndOk() (*bool, bool)`

GetAlteredPriorityIndOk returns a tuple with the AlteredPriorityInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlteredPriorityInd

`func (o *SearchResult) SetAlteredPriorityInd(v bool)`

SetAlteredPriorityInd sets AlteredPriorityInd field to given value.

### HasAlteredPriorityInd

`func (o *SearchResult) HasAlteredPriorityInd() bool`

HasAlteredPriorityInd returns a boolean if a field has been set.

### GetNoProfileMatchInfo

`func (o *SearchResult) GetNoProfileMatchInfo() NoProfileMatchInfo`

GetNoProfileMatchInfo returns the NoProfileMatchInfo field if non-nil, zero value otherwise.

### GetNoProfileMatchInfoOk

`func (o *SearchResult) GetNoProfileMatchInfoOk() (*NoProfileMatchInfo, bool)`

GetNoProfileMatchInfoOk returns a tuple with the NoProfileMatchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoProfileMatchInfo

`func (o *SearchResult) SetNoProfileMatchInfo(v NoProfileMatchInfo)`

SetNoProfileMatchInfo sets NoProfileMatchInfo field to given value.

### HasNoProfileMatchInfo

`func (o *SearchResult) HasNoProfileMatchInfo() bool`

HasNoProfileMatchInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


