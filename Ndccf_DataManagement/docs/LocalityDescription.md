# LocalityDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalityType** | [**LocalityType**](LocalityType.md) |  | 
**LocalityValue** | **string** |  | 
**AddlLocDescrItems** | Pointer to [**[]LocalityDescriptionItem**](LocalityDescriptionItem.md) |  | [optional] 

## Methods

### NewLocalityDescription

`func NewLocalityDescription(localityType LocalityType, localityValue string, ) *LocalityDescription`

NewLocalityDescription instantiates a new LocalityDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalityDescriptionWithDefaults

`func NewLocalityDescriptionWithDefaults() *LocalityDescription`

NewLocalityDescriptionWithDefaults instantiates a new LocalityDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalityType

`func (o *LocalityDescription) GetLocalityType() LocalityType`

GetLocalityType returns the LocalityType field if non-nil, zero value otherwise.

### GetLocalityTypeOk

`func (o *LocalityDescription) GetLocalityTypeOk() (*LocalityType, bool)`

GetLocalityTypeOk returns a tuple with the LocalityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityType

`func (o *LocalityDescription) SetLocalityType(v LocalityType)`

SetLocalityType sets LocalityType field to given value.


### GetLocalityValue

`func (o *LocalityDescription) GetLocalityValue() string`

GetLocalityValue returns the LocalityValue field if non-nil, zero value otherwise.

### GetLocalityValueOk

`func (o *LocalityDescription) GetLocalityValueOk() (*string, bool)`

GetLocalityValueOk returns a tuple with the LocalityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityValue

`func (o *LocalityDescription) SetLocalityValue(v string)`

SetLocalityValue sets LocalityValue field to given value.


### GetAddlLocDescrItems

`func (o *LocalityDescription) GetAddlLocDescrItems() []LocalityDescriptionItem`

GetAddlLocDescrItems returns the AddlLocDescrItems field if non-nil, zero value otherwise.

### GetAddlLocDescrItemsOk

`func (o *LocalityDescription) GetAddlLocDescrItemsOk() (*[]LocalityDescriptionItem, bool)`

GetAddlLocDescrItemsOk returns a tuple with the AddlLocDescrItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddlLocDescrItems

`func (o *LocalityDescription) SetAddlLocDescrItems(v []LocalityDescriptionItem)`

SetAddlLocDescrItems sets AddlLocDescrItems field to given value.

### HasAddlLocDescrItems

`func (o *LocalityDescription) HasAddlLocDescrItems() bool`

HasAddlLocDescrItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


