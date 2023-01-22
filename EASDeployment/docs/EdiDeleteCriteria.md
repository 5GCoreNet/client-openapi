# EdiDeleteCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | Pointer to **string** | Represents an AF identifier. | [optional] 
**DnnSnssai** | Pointer to [**DnnSnssaiInformation**](DnnSnssaiInformation.md) |  | [optional] 

## Methods

### NewEdiDeleteCriteria

`func NewEdiDeleteCriteria() *EdiDeleteCriteria`

NewEdiDeleteCriteria instantiates a new EdiDeleteCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdiDeleteCriteriaWithDefaults

`func NewEdiDeleteCriteriaWithDefaults() *EdiDeleteCriteria`

NewEdiDeleteCriteriaWithDefaults instantiates a new EdiDeleteCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *EdiDeleteCriteria) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *EdiDeleteCriteria) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *EdiDeleteCriteria) SetAfId(v string)`

SetAfId sets AfId field to given value.

### HasAfId

`func (o *EdiDeleteCriteria) HasAfId() bool`

HasAfId returns a boolean if a field has been set.

### GetDnnSnssai

`func (o *EdiDeleteCriteria) GetDnnSnssai() DnnSnssaiInformation`

GetDnnSnssai returns the DnnSnssai field if non-nil, zero value otherwise.

### GetDnnSnssaiOk

`func (o *EdiDeleteCriteria) GetDnnSnssaiOk() (*DnnSnssaiInformation, bool)`

GetDnnSnssaiOk returns a tuple with the DnnSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssai

`func (o *EdiDeleteCriteria) SetDnnSnssai(v DnnSnssaiInformation)`

SetDnnSnssai sets DnnSnssai field to given value.

### HasDnnSnssai

`func (o *EdiDeleteCriteria) HasDnnSnssai() bool`

HasDnnSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


