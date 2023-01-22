# Guss

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BsfInfo** | Pointer to [**BsfInfo**](BsfInfo.md) |  | [optional] 
**UssList** | Pointer to [**[]UssListItem**](UssListItem.md) |  | [optional] 

## Methods

### NewGuss

`func NewGuss() *Guss`

NewGuss instantiates a new Guss object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGussWithDefaults

`func NewGussWithDefaults() *Guss`

NewGussWithDefaults instantiates a new Guss object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsfInfo

`func (o *Guss) GetBsfInfo() BsfInfo`

GetBsfInfo returns the BsfInfo field if non-nil, zero value otherwise.

### GetBsfInfoOk

`func (o *Guss) GetBsfInfoOk() (*BsfInfo, bool)`

GetBsfInfoOk returns a tuple with the BsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfo

`func (o *Guss) SetBsfInfo(v BsfInfo)`

SetBsfInfo sets BsfInfo field to given value.

### HasBsfInfo

`func (o *Guss) HasBsfInfo() bool`

HasBsfInfo returns a boolean if a field has been set.

### GetUssList

`func (o *Guss) GetUssList() []UssListItem`

GetUssList returns the UssList field if non-nil, zero value otherwise.

### GetUssListOk

`func (o *Guss) GetUssListOk() (*[]UssListItem, bool)`

GetUssListOk returns a tuple with the UssList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUssList

`func (o *Guss) SetUssList(v []UssListItem)`

SetUssList sets UssList field to given value.

### HasUssList

`func (o *Guss) HasUssList() bool`

HasUssList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


