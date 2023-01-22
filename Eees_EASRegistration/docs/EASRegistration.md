# EASRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasProf** | [**EASProfile**](EASProfile.md) |  | 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewEASRegistration

`func NewEASRegistration(easProf EASProfile, ) *EASRegistration`

NewEASRegistration instantiates a new EASRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEASRegistrationWithDefaults

`func NewEASRegistrationWithDefaults() *EASRegistration`

NewEASRegistrationWithDefaults instantiates a new EASRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasProf

`func (o *EASRegistration) GetEasProf() EASProfile`

GetEasProf returns the EasProf field if non-nil, zero value otherwise.

### GetEasProfOk

`func (o *EASRegistration) GetEasProfOk() (*EASProfile, bool)`

GetEasProfOk returns a tuple with the EasProf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasProf

`func (o *EASRegistration) SetEasProf(v EASProfile)`

SetEasProf sets EasProf field to given value.


### GetExpTime

`func (o *EASRegistration) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *EASRegistration) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *EASRegistration) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *EASRegistration) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetSuppFeat

`func (o *EASRegistration) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *EASRegistration) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *EASRegistration) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *EASRegistration) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


