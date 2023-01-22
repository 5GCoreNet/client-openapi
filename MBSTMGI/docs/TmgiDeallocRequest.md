# TmgiDeallocRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfId** | **string** |  | 
**Tmgis** | [**[]Tmgi**](Tmgi.md) |  | 

## Methods

### NewTmgiDeallocRequest

`func NewTmgiDeallocRequest(afId string, tmgis []Tmgi, ) *TmgiDeallocRequest`

NewTmgiDeallocRequest instantiates a new TmgiDeallocRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTmgiDeallocRequestWithDefaults

`func NewTmgiDeallocRequestWithDefaults() *TmgiDeallocRequest`

NewTmgiDeallocRequestWithDefaults instantiates a new TmgiDeallocRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfId

`func (o *TmgiDeallocRequest) GetAfId() string`

GetAfId returns the AfId field if non-nil, zero value otherwise.

### GetAfIdOk

`func (o *TmgiDeallocRequest) GetAfIdOk() (*string, bool)`

GetAfIdOk returns a tuple with the AfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfId

`func (o *TmgiDeallocRequest) SetAfId(v string)`

SetAfId sets AfId field to given value.


### GetTmgis

`func (o *TmgiDeallocRequest) GetTmgis() []Tmgi`

GetTmgis returns the Tmgis field if non-nil, zero value otherwise.

### GetTmgisOk

`func (o *TmgiDeallocRequest) GetTmgisOk() (*[]Tmgi, bool)`

GetTmgisOk returns a tuple with the Tmgis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgis

`func (o *TmgiDeallocRequest) SetTmgis(v []Tmgi)`

SetTmgis sets Tmgis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


