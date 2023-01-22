# IntendedN32Purpose

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsagePurpose** | [**N32Purpose**](N32Purpose.md) |  | 
**AdditionalInfo** | Pointer to **string** |  | [optional] 
**Cause** | Pointer to **string** |  | [optional] 

## Methods

### NewIntendedN32Purpose

`func NewIntendedN32Purpose(usagePurpose N32Purpose, ) *IntendedN32Purpose`

NewIntendedN32Purpose instantiates a new IntendedN32Purpose object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntendedN32PurposeWithDefaults

`func NewIntendedN32PurposeWithDefaults() *IntendedN32Purpose`

NewIntendedN32PurposeWithDefaults instantiates a new IntendedN32Purpose object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsagePurpose

`func (o *IntendedN32Purpose) GetUsagePurpose() N32Purpose`

GetUsagePurpose returns the UsagePurpose field if non-nil, zero value otherwise.

### GetUsagePurposeOk

`func (o *IntendedN32Purpose) GetUsagePurposeOk() (*N32Purpose, bool)`

GetUsagePurposeOk returns a tuple with the UsagePurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsagePurpose

`func (o *IntendedN32Purpose) SetUsagePurpose(v N32Purpose)`

SetUsagePurpose sets UsagePurpose field to given value.


### GetAdditionalInfo

`func (o *IntendedN32Purpose) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *IntendedN32Purpose) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *IntendedN32Purpose) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *IntendedN32Purpose) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetCause

`func (o *IntendedN32Purpose) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *IntendedN32Purpose) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *IntendedN32Purpose) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *IntendedN32Purpose) HasCause() bool`

HasCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


