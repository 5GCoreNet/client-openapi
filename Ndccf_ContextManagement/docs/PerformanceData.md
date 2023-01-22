# PerformanceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pdb** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 
**Plr** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 
**ThrputUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ThrputDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewPerformanceData

`func NewPerformanceData() *PerformanceData`

NewPerformanceData instantiates a new PerformanceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerformanceDataWithDefaults

`func NewPerformanceDataWithDefaults() *PerformanceData`

NewPerformanceDataWithDefaults instantiates a new PerformanceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPdb

`func (o *PerformanceData) GetPdb() int32`

GetPdb returns the Pdb field if non-nil, zero value otherwise.

### GetPdbOk

`func (o *PerformanceData) GetPdbOk() (*int32, bool)`

GetPdbOk returns a tuple with the Pdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdb

`func (o *PerformanceData) SetPdb(v int32)`

SetPdb sets Pdb field to given value.

### HasPdb

`func (o *PerformanceData) HasPdb() bool`

HasPdb returns a boolean if a field has been set.

### GetPlr

`func (o *PerformanceData) GetPlr() int32`

GetPlr returns the Plr field if non-nil, zero value otherwise.

### GetPlrOk

`func (o *PerformanceData) GetPlrOk() (*int32, bool)`

GetPlrOk returns a tuple with the Plr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlr

`func (o *PerformanceData) SetPlr(v int32)`

SetPlr sets Plr field to given value.

### HasPlr

`func (o *PerformanceData) HasPlr() bool`

HasPlr returns a boolean if a field has been set.

### GetThrputUl

`func (o *PerformanceData) GetThrputUl() string`

GetThrputUl returns the ThrputUl field if non-nil, zero value otherwise.

### GetThrputUlOk

`func (o *PerformanceData) GetThrputUlOk() (*string, bool)`

GetThrputUlOk returns a tuple with the ThrputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrputUl

`func (o *PerformanceData) SetThrputUl(v string)`

SetThrputUl sets ThrputUl field to given value.

### HasThrputUl

`func (o *PerformanceData) HasThrputUl() bool`

HasThrputUl returns a boolean if a field has been set.

### GetThrputDl

`func (o *PerformanceData) GetThrputDl() string`

GetThrputDl returns the ThrputDl field if non-nil, zero value otherwise.

### GetThrputDlOk

`func (o *PerformanceData) GetThrputDlOk() (*string, bool)`

GetThrputDlOk returns a tuple with the ThrputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrputDl

`func (o *PerformanceData) SetThrputDl(v string)`

SetThrputDl sets ThrputDl field to given value.

### HasThrputDl

`func (o *PerformanceData) HasThrputDl() bool`

HasThrputDl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


