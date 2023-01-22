# HttpPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IePath** | **string** |  | 
**IeValueLocation** | [**IeLocation**](IeLocation.md) |  | 
**Value** | **map[string]interface{}** |  | 

## Methods

### NewHttpPayload

`func NewHttpPayload(iePath string, ieValueLocation IeLocation, value map[string]interface{}, ) *HttpPayload`

NewHttpPayload instantiates a new HttpPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpPayloadWithDefaults

`func NewHttpPayloadWithDefaults() *HttpPayload`

NewHttpPayloadWithDefaults instantiates a new HttpPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIePath

`func (o *HttpPayload) GetIePath() string`

GetIePath returns the IePath field if non-nil, zero value otherwise.

### GetIePathOk

`func (o *HttpPayload) GetIePathOk() (*string, bool)`

GetIePathOk returns a tuple with the IePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIePath

`func (o *HttpPayload) SetIePath(v string)`

SetIePath sets IePath field to given value.


### GetIeValueLocation

`func (o *HttpPayload) GetIeValueLocation() IeLocation`

GetIeValueLocation returns the IeValueLocation field if non-nil, zero value otherwise.

### GetIeValueLocationOk

`func (o *HttpPayload) GetIeValueLocationOk() (*IeLocation, bool)`

GetIeValueLocationOk returns a tuple with the IeValueLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIeValueLocation

`func (o *HttpPayload) SetIeValueLocation(v IeLocation)`

SetIeValueLocation sets IeValueLocation field to given value.


### GetValue

`func (o *HttpPayload) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HttpPayload) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HttpPayload) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


