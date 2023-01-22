# UeAreaIndication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | Indicates country where UE is located | [optional] 
**InternationalAreaInd** | Pointer to **bool** | Indicates international area indication if UE is located in international area | [optional] [default to false]

## Methods

### NewUeAreaIndication

`func NewUeAreaIndication() *UeAreaIndication`

NewUeAreaIndication instantiates a new UeAreaIndication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeAreaIndicationWithDefaults

`func NewUeAreaIndicationWithDefaults() *UeAreaIndication`

NewUeAreaIndicationWithDefaults instantiates a new UeAreaIndication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *UeAreaIndication) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UeAreaIndication) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UeAreaIndication) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UeAreaIndication) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetInternationalAreaInd

`func (o *UeAreaIndication) GetInternationalAreaInd() bool`

GetInternationalAreaInd returns the InternationalAreaInd field if non-nil, zero value otherwise.

### GetInternationalAreaIndOk

`func (o *UeAreaIndication) GetInternationalAreaIndOk() (*bool, bool)`

GetInternationalAreaIndOk returns a tuple with the InternationalAreaInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternationalAreaInd

`func (o *UeAreaIndication) SetInternationalAreaInd(v bool)`

SetInternationalAreaInd sets InternationalAreaInd field to given value.

### HasInternationalAreaInd

`func (o *UeAreaIndication) HasInternationalAreaInd() bool`

HasInternationalAreaInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


