# EncodedHttpHeaderValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncBlockIndex** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 

## Methods

### NewEncodedHttpHeaderValue

`func NewEncodedHttpHeaderValue(encBlockIndex int32, ) *EncodedHttpHeaderValue`

NewEncodedHttpHeaderValue instantiates a new EncodedHttpHeaderValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodedHttpHeaderValueWithDefaults

`func NewEncodedHttpHeaderValueWithDefaults() *EncodedHttpHeaderValue`

NewEncodedHttpHeaderValueWithDefaults instantiates a new EncodedHttpHeaderValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncBlockIndex

`func (o *EncodedHttpHeaderValue) GetEncBlockIndex() int32`

GetEncBlockIndex returns the EncBlockIndex field if non-nil, zero value otherwise.

### GetEncBlockIndexOk

`func (o *EncodedHttpHeaderValue) GetEncBlockIndexOk() (*int32, bool)`

GetEncBlockIndexOk returns a tuple with the EncBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncBlockIndex

`func (o *EncodedHttpHeaderValue) SetEncBlockIndex(v int32)`

SetEncBlockIndex sets EncBlockIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


