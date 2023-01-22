# SharedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedDataId** | **string** | Identifies globally and uniquely a piece of subscription data shared by multiple UEs; the value shall start with the HPLMN id (MCC/MNC) followed by a hyphen followed by a local Id as allocated by the home network operator  | 
**SharedImsIfcData** | Pointer to [**Ifcs**](Ifcs.md) |  | [optional] 

## Methods

### NewSharedData

`func NewSharedData(sharedDataId string, ) *SharedData`

NewSharedData instantiates a new SharedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedDataWithDefaults

`func NewSharedDataWithDefaults() *SharedData`

NewSharedDataWithDefaults instantiates a new SharedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedDataId

`func (o *SharedData) GetSharedDataId() string`

GetSharedDataId returns the SharedDataId field if non-nil, zero value otherwise.

### GetSharedDataIdOk

`func (o *SharedData) GetSharedDataIdOk() (*string, bool)`

GetSharedDataIdOk returns a tuple with the SharedDataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedDataId

`func (o *SharedData) SetSharedDataId(v string)`

SetSharedDataId sets SharedDataId field to given value.


### GetSharedImsIfcData

`func (o *SharedData) GetSharedImsIfcData() Ifcs`

GetSharedImsIfcData returns the SharedImsIfcData field if non-nil, zero value otherwise.

### GetSharedImsIfcDataOk

`func (o *SharedData) GetSharedImsIfcDataOk() (*Ifcs, bool)`

GetSharedImsIfcDataOk returns a tuple with the SharedImsIfcData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedImsIfcData

`func (o *SharedData) SetSharedImsIfcData(v Ifcs)`

SetSharedImsIfcData sets SharedImsIfcData field to given value.

### HasSharedImsIfcData

`func (o *SharedData) HasSharedImsIfcData() bool`

HasSharedImsIfcData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


