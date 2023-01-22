# UeACResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcuFailureList** | Pointer to [**map[string][]AcuFailureItem**](array.md) | A map (list of key-value pairs) where the key of the map shall be UE&#39;s supi | [optional] 

## Methods

### NewUeACResponseData

`func NewUeACResponseData() *UeACResponseData`

NewUeACResponseData instantiates a new UeACResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeACResponseDataWithDefaults

`func NewUeACResponseDataWithDefaults() *UeACResponseData`

NewUeACResponseDataWithDefaults instantiates a new UeACResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcuFailureList

`func (o *UeACResponseData) GetAcuFailureList() map[string][]AcuFailureItem`

GetAcuFailureList returns the AcuFailureList field if non-nil, zero value otherwise.

### GetAcuFailureListOk

`func (o *UeACResponseData) GetAcuFailureListOk() (*map[string][]AcuFailureItem, bool)`

GetAcuFailureListOk returns a tuple with the AcuFailureList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcuFailureList

`func (o *UeACResponseData) SetAcuFailureList(v map[string][]AcuFailureItem)`

SetAcuFailureList sets AcuFailureList field to given value.

### HasAcuFailureList

`func (o *UeACResponseData) HasAcuFailureList() bool`

HasAcuFailureList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


