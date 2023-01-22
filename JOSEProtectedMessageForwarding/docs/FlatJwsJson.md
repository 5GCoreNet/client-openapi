# FlatJwsJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | **string** |  | 
**Protected** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **map[string]interface{}** |  | [optional] 
**Signature** | **string** |  | 

## Methods

### NewFlatJwsJson

`func NewFlatJwsJson(payload string, signature string, ) *FlatJwsJson`

NewFlatJwsJson instantiates a new FlatJwsJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatJwsJsonWithDefaults

`func NewFlatJwsJsonWithDefaults() *FlatJwsJson`

NewFlatJwsJsonWithDefaults instantiates a new FlatJwsJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *FlatJwsJson) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *FlatJwsJson) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *FlatJwsJson) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetProtected

`func (o *FlatJwsJson) GetProtected() string`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *FlatJwsJson) GetProtectedOk() (*string, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *FlatJwsJson) SetProtected(v string)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *FlatJwsJson) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetHeader

`func (o *FlatJwsJson) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FlatJwsJson) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FlatJwsJson) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *FlatJwsJson) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetSignature

`func (o *FlatJwsJson) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *FlatJwsJson) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *FlatJwsJson) SetSignature(v string)`

SetSignature sets Signature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


