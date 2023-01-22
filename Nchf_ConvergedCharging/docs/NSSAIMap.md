# NSSAIMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServingSnssai** | [**Snssai**](Snssai.md) |  | 
**HomeSnssai** | [**Snssai**](Snssai.md) |  | 

## Methods

### NewNSSAIMap

`func NewNSSAIMap(servingSnssai Snssai, homeSnssai Snssai, ) *NSSAIMap`

NewNSSAIMap instantiates a new NSSAIMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSSAIMapWithDefaults

`func NewNSSAIMapWithDefaults() *NSSAIMap`

NewNSSAIMapWithDefaults instantiates a new NSSAIMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServingSnssai

`func (o *NSSAIMap) GetServingSnssai() Snssai`

GetServingSnssai returns the ServingSnssai field if non-nil, zero value otherwise.

### GetServingSnssaiOk

`func (o *NSSAIMap) GetServingSnssaiOk() (*Snssai, bool)`

GetServingSnssaiOk returns a tuple with the ServingSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingSnssai

`func (o *NSSAIMap) SetServingSnssai(v Snssai)`

SetServingSnssai sets ServingSnssai field to given value.


### GetHomeSnssai

`func (o *NSSAIMap) GetHomeSnssai() Snssai`

GetHomeSnssai returns the HomeSnssai field if non-nil, zero value otherwise.

### GetHomeSnssaiOk

`func (o *NSSAIMap) GetHomeSnssaiOk() (*Snssai, bool)`

GetHomeSnssaiOk returns a tuple with the HomeSnssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeSnssai

`func (o *NSSAIMap) SetHomeSnssai(v Snssai)`

SetHomeSnssai sets HomeSnssai field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


