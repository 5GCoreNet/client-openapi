# SupportedSnssai

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SNssai** | [**ExtSnssai**](ExtSnssai.md) |  | 
**RestrictionInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewSupportedSnssai

`func NewSupportedSnssai(sNssai ExtSnssai, ) *SupportedSnssai`

NewSupportedSnssai instantiates a new SupportedSnssai object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedSnssaiWithDefaults

`func NewSupportedSnssaiWithDefaults() *SupportedSnssai`

NewSupportedSnssaiWithDefaults instantiates a new SupportedSnssai object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSNssai

`func (o *SupportedSnssai) GetSNssai() ExtSnssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *SupportedSnssai) GetSNssaiOk() (*ExtSnssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *SupportedSnssai) SetSNssai(v ExtSnssai)`

SetSNssai sets SNssai field to given value.


### GetRestrictionInd

`func (o *SupportedSnssai) GetRestrictionInd() bool`

GetRestrictionInd returns the RestrictionInd field if non-nil, zero value otherwise.

### GetRestrictionIndOk

`func (o *SupportedSnssai) GetRestrictionIndOk() (*bool, bool)`

GetRestrictionIndOk returns a tuple with the RestrictionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictionInd

`func (o *SupportedSnssai) SetRestrictionInd(v bool)`

SetRestrictionInd sets RestrictionInd field to given value.

### HasRestrictionInd

`func (o *SupportedSnssai) HasRestrictionInd() bool`

HasRestrictionInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


