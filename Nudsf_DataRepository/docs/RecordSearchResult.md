# RecordSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**References** | Pointer to **[]string** |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MatchingRecords** | Pointer to [**map[string]Record**](Record.md) | A map (list of key-value pairs where recordId serves as key) of Records | [optional] 

## Methods

### NewRecordSearchResult

`func NewRecordSearchResult(count int32, ) *RecordSearchResult`

NewRecordSearchResult instantiates a new RecordSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSearchResultWithDefaults

`func NewRecordSearchResultWithDefaults() *RecordSearchResult`

NewRecordSearchResultWithDefaults instantiates a new RecordSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RecordSearchResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecordSearchResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecordSearchResult) SetCount(v int32)`

SetCount sets Count field to given value.


### GetReferences

`func (o *RecordSearchResult) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *RecordSearchResult) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *RecordSearchResult) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *RecordSearchResult) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *RecordSearchResult) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *RecordSearchResult) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *RecordSearchResult) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *RecordSearchResult) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMatchingRecords

`func (o *RecordSearchResult) GetMatchingRecords() map[string]Record`

GetMatchingRecords returns the MatchingRecords field if non-nil, zero value otherwise.

### GetMatchingRecordsOk

`func (o *RecordSearchResult) GetMatchingRecordsOk() (*map[string]Record, bool)`

GetMatchingRecordsOk returns a tuple with the MatchingRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRecords

`func (o *RecordSearchResult) SetMatchingRecords(v map[string]Record)`

SetMatchingRecords sets MatchingRecords field to given value.

### HasMatchingRecords

`func (o *RecordSearchResult) HasMatchingRecords() bool`

HasMatchingRecords returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


