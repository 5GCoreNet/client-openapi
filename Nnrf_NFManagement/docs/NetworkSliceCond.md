# NetworkSliceCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnssaiList** | [**[]Snssai**](Snssai.md) |  | 
**NsiList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNetworkSliceCond

`func NewNetworkSliceCond(snssaiList []Snssai, ) *NetworkSliceCond`

NewNetworkSliceCond instantiates a new NetworkSliceCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSliceCondWithDefaults

`func NewNetworkSliceCondWithDefaults() *NetworkSliceCond`

NewNetworkSliceCondWithDefaults instantiates a new NetworkSliceCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssaiList

`func (o *NetworkSliceCond) GetSnssaiList() []Snssai`

GetSnssaiList returns the SnssaiList field if non-nil, zero value otherwise.

### GetSnssaiListOk

`func (o *NetworkSliceCond) GetSnssaiListOk() (*[]Snssai, bool)`

GetSnssaiListOk returns a tuple with the SnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssaiList

`func (o *NetworkSliceCond) SetSnssaiList(v []Snssai)`

SetSnssaiList sets SnssaiList field to given value.


### GetNsiList

`func (o *NetworkSliceCond) GetNsiList() []string`

GetNsiList returns the NsiList field if non-nil, zero value otherwise.

### GetNsiListOk

`func (o *NetworkSliceCond) GetNsiListOk() (*[]string, bool)`

GetNsiListOk returns a tuple with the NsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiList

`func (o *NetworkSliceCond) SetNsiList(v []string)`

SetNsiList sets NsiList field to given value.

### HasNsiList

`func (o *NetworkSliceCond) HasNsiList() bool`

HasNsiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


