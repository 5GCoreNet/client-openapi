# MbsSessPolCtrlData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var5qis** | Pointer to **[]int32** |  | [optional] 
**MaxMbsArpLevel** | Pointer to **NullableInt32** | nullable true shall not be used for this attribute. Unsigned integer indicating the ARP Priority Level (see clause 5.7.2.2 of 3GPP TS 23.501, within the range 1 to 15.Values are ordered in decreasing order of priority, i.e. with 1 as the highest priority and 15 as the lowest priority.   | [optional] 
**MaxMbsSessionAmbr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxGbr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsSessPolCtrlData

`func NewMbsSessPolCtrlData() *MbsSessPolCtrlData`

NewMbsSessPolCtrlData instantiates a new MbsSessPolCtrlData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessPolCtrlDataWithDefaults

`func NewMbsSessPolCtrlDataWithDefaults() *MbsSessPolCtrlData`

NewMbsSessPolCtrlDataWithDefaults instantiates a new MbsSessPolCtrlData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar5qis

`func (o *MbsSessPolCtrlData) GetVar5qis() []int32`

GetVar5qis returns the Var5qis field if non-nil, zero value otherwise.

### GetVar5qisOk

`func (o *MbsSessPolCtrlData) GetVar5qisOk() (*[]int32, bool)`

GetVar5qisOk returns a tuple with the Var5qis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qis

`func (o *MbsSessPolCtrlData) SetVar5qis(v []int32)`

SetVar5qis sets Var5qis field to given value.

### HasVar5qis

`func (o *MbsSessPolCtrlData) HasVar5qis() bool`

HasVar5qis returns a boolean if a field has been set.

### GetMaxMbsArpLevel

`func (o *MbsSessPolCtrlData) GetMaxMbsArpLevel() int32`

GetMaxMbsArpLevel returns the MaxMbsArpLevel field if non-nil, zero value otherwise.

### GetMaxMbsArpLevelOk

`func (o *MbsSessPolCtrlData) GetMaxMbsArpLevelOk() (*int32, bool)`

GetMaxMbsArpLevelOk returns a tuple with the MaxMbsArpLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMbsArpLevel

`func (o *MbsSessPolCtrlData) SetMaxMbsArpLevel(v int32)`

SetMaxMbsArpLevel sets MaxMbsArpLevel field to given value.

### HasMaxMbsArpLevel

`func (o *MbsSessPolCtrlData) HasMaxMbsArpLevel() bool`

HasMaxMbsArpLevel returns a boolean if a field has been set.

### SetMaxMbsArpLevelNil

`func (o *MbsSessPolCtrlData) SetMaxMbsArpLevelNil(b bool)`

 SetMaxMbsArpLevelNil sets the value for MaxMbsArpLevel to be an explicit nil

### UnsetMaxMbsArpLevel
`func (o *MbsSessPolCtrlData) UnsetMaxMbsArpLevel()`

UnsetMaxMbsArpLevel ensures that no value is present for MaxMbsArpLevel, not even an explicit nil
### GetMaxMbsSessionAmbr

`func (o *MbsSessPolCtrlData) GetMaxMbsSessionAmbr() string`

GetMaxMbsSessionAmbr returns the MaxMbsSessionAmbr field if non-nil, zero value otherwise.

### GetMaxMbsSessionAmbrOk

`func (o *MbsSessPolCtrlData) GetMaxMbsSessionAmbrOk() (*string, bool)`

GetMaxMbsSessionAmbrOk returns a tuple with the MaxMbsSessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMbsSessionAmbr

`func (o *MbsSessPolCtrlData) SetMaxMbsSessionAmbr(v string)`

SetMaxMbsSessionAmbr sets MaxMbsSessionAmbr field to given value.

### HasMaxMbsSessionAmbr

`func (o *MbsSessPolCtrlData) HasMaxMbsSessionAmbr() bool`

HasMaxMbsSessionAmbr returns a boolean if a field has been set.

### GetMaxGbr

`func (o *MbsSessPolCtrlData) GetMaxGbr() string`

GetMaxGbr returns the MaxGbr field if non-nil, zero value otherwise.

### GetMaxGbrOk

`func (o *MbsSessPolCtrlData) GetMaxGbrOk() (*string, bool)`

GetMaxGbrOk returns a tuple with the MaxGbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGbr

`func (o *MbsSessPolCtrlData) SetMaxGbr(v string)`

SetMaxGbr sets MaxGbr field to given value.

### HasMaxGbr

`func (o *MbsSessPolCtrlData) HasMaxGbr() bool`

HasMaxGbr returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MbsSessPolCtrlData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MbsSessPolCtrlData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MbsSessPolCtrlData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MbsSessPolCtrlData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


