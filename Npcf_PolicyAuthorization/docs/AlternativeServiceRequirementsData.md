# AlternativeServiceRequirementsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltQosParamSetRef** | **string** | Reference to this alternative QoS related parameter set. | 
**GbrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**GbrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**Pdb** | Pointer to **int32** | Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.  | [optional] 

## Methods

### NewAlternativeServiceRequirementsData

`func NewAlternativeServiceRequirementsData(altQosParamSetRef string, ) *AlternativeServiceRequirementsData`

NewAlternativeServiceRequirementsData instantiates a new AlternativeServiceRequirementsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlternativeServiceRequirementsDataWithDefaults

`func NewAlternativeServiceRequirementsDataWithDefaults() *AlternativeServiceRequirementsData`

NewAlternativeServiceRequirementsDataWithDefaults instantiates a new AlternativeServiceRequirementsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltQosParamSetRef

`func (o *AlternativeServiceRequirementsData) GetAltQosParamSetRef() string`

GetAltQosParamSetRef returns the AltQosParamSetRef field if non-nil, zero value otherwise.

### GetAltQosParamSetRefOk

`func (o *AlternativeServiceRequirementsData) GetAltQosParamSetRefOk() (*string, bool)`

GetAltQosParamSetRefOk returns a tuple with the AltQosParamSetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosParamSetRef

`func (o *AlternativeServiceRequirementsData) SetAltQosParamSetRef(v string)`

SetAltQosParamSetRef sets AltQosParamSetRef field to given value.


### GetGbrUl

`func (o *AlternativeServiceRequirementsData) GetGbrUl() string`

GetGbrUl returns the GbrUl field if non-nil, zero value otherwise.

### GetGbrUlOk

`func (o *AlternativeServiceRequirementsData) GetGbrUlOk() (*string, bool)`

GetGbrUlOk returns a tuple with the GbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrUl

`func (o *AlternativeServiceRequirementsData) SetGbrUl(v string)`

SetGbrUl sets GbrUl field to given value.

### HasGbrUl

`func (o *AlternativeServiceRequirementsData) HasGbrUl() bool`

HasGbrUl returns a boolean if a field has been set.

### GetGbrDl

`func (o *AlternativeServiceRequirementsData) GetGbrDl() string`

GetGbrDl returns the GbrDl field if non-nil, zero value otherwise.

### GetGbrDlOk

`func (o *AlternativeServiceRequirementsData) GetGbrDlOk() (*string, bool)`

GetGbrDlOk returns a tuple with the GbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrDl

`func (o *AlternativeServiceRequirementsData) SetGbrDl(v string)`

SetGbrDl sets GbrDl field to given value.

### HasGbrDl

`func (o *AlternativeServiceRequirementsData) HasGbrDl() bool`

HasGbrDl returns a boolean if a field has been set.

### GetPdb

`func (o *AlternativeServiceRequirementsData) GetPdb() int32`

GetPdb returns the Pdb field if non-nil, zero value otherwise.

### GetPdbOk

`func (o *AlternativeServiceRequirementsData) GetPdbOk() (*int32, bool)`

GetPdbOk returns a tuple with the Pdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdb

`func (o *AlternativeServiceRequirementsData) SetPdb(v int32)`

SetPdb sets Pdb field to given value.

### HasPdb

`func (o *AlternativeServiceRequirementsData) HasPdb() bool`

HasPdb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


