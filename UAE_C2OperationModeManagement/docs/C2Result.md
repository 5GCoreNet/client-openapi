# C2Result

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**C2OpConfirmed** | **bool** |  | 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewC2Result

`func NewC2Result(c2OpConfirmed bool, ) *C2Result`

NewC2Result instantiates a new C2Result object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewC2ResultWithDefaults

`func NewC2ResultWithDefaults() *C2Result`

NewC2ResultWithDefaults instantiates a new C2Result object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetC2OpConfirmed

`func (o *C2Result) GetC2OpConfirmed() bool`

GetC2OpConfirmed returns the C2OpConfirmed field if non-nil, zero value otherwise.

### GetC2OpConfirmedOk

`func (o *C2Result) GetC2OpConfirmedOk() (*bool, bool)`

GetC2OpConfirmedOk returns a tuple with the C2OpConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2OpConfirmed

`func (o *C2Result) SetC2OpConfirmed(v bool)`

SetC2OpConfirmed sets C2OpConfirmed field to given value.


### GetSuppFeat

`func (o *C2Result) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *C2Result) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *C2Result) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *C2Result) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


