# FlatJweJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protected** | Pointer to **string** |  | [optional] 
**Unprotected** | Pointer to **map[string]interface{}** |  | [optional] 
**Header** | Pointer to **map[string]interface{}** |  | [optional] 
**EncryptedKey** | Pointer to **string** |  | [optional] 
**Aad** | Pointer to **string** |  | [optional] 
**Iv** | Pointer to **string** |  | [optional] 
**Ciphertext** | **string** |  | 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewFlatJweJson

`func NewFlatJweJson(ciphertext string, ) *FlatJweJson`

NewFlatJweJson instantiates a new FlatJweJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlatJweJsonWithDefaults

`func NewFlatJweJsonWithDefaults() *FlatJweJson`

NewFlatJweJsonWithDefaults instantiates a new FlatJweJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtected

`func (o *FlatJweJson) GetProtected() string`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *FlatJweJson) GetProtectedOk() (*string, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *FlatJweJson) SetProtected(v string)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *FlatJweJson) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetUnprotected

`func (o *FlatJweJson) GetUnprotected() map[string]interface{}`

GetUnprotected returns the Unprotected field if non-nil, zero value otherwise.

### GetUnprotectedOk

`func (o *FlatJweJson) GetUnprotectedOk() (*map[string]interface{}, bool)`

GetUnprotectedOk returns a tuple with the Unprotected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprotected

`func (o *FlatJweJson) SetUnprotected(v map[string]interface{})`

SetUnprotected sets Unprotected field to given value.

### HasUnprotected

`func (o *FlatJweJson) HasUnprotected() bool`

HasUnprotected returns a boolean if a field has been set.

### GetHeader

`func (o *FlatJweJson) GetHeader() map[string]interface{}`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FlatJweJson) GetHeaderOk() (*map[string]interface{}, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FlatJweJson) SetHeader(v map[string]interface{})`

SetHeader sets Header field to given value.

### HasHeader

`func (o *FlatJweJson) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetEncryptedKey

`func (o *FlatJweJson) GetEncryptedKey() string`

GetEncryptedKey returns the EncryptedKey field if non-nil, zero value otherwise.

### GetEncryptedKeyOk

`func (o *FlatJweJson) GetEncryptedKeyOk() (*string, bool)`

GetEncryptedKeyOk returns a tuple with the EncryptedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedKey

`func (o *FlatJweJson) SetEncryptedKey(v string)`

SetEncryptedKey sets EncryptedKey field to given value.

### HasEncryptedKey

`func (o *FlatJweJson) HasEncryptedKey() bool`

HasEncryptedKey returns a boolean if a field has been set.

### GetAad

`func (o *FlatJweJson) GetAad() string`

GetAad returns the Aad field if non-nil, zero value otherwise.

### GetAadOk

`func (o *FlatJweJson) GetAadOk() (*string, bool)`

GetAadOk returns a tuple with the Aad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAad

`func (o *FlatJweJson) SetAad(v string)`

SetAad sets Aad field to given value.

### HasAad

`func (o *FlatJweJson) HasAad() bool`

HasAad returns a boolean if a field has been set.

### GetIv

`func (o *FlatJweJson) GetIv() string`

GetIv returns the Iv field if non-nil, zero value otherwise.

### GetIvOk

`func (o *FlatJweJson) GetIvOk() (*string, bool)`

GetIvOk returns a tuple with the Iv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIv

`func (o *FlatJweJson) SetIv(v string)`

SetIv sets Iv field to given value.

### HasIv

`func (o *FlatJweJson) HasIv() bool`

HasIv returns a boolean if a field has been set.

### GetCiphertext

`func (o *FlatJweJson) GetCiphertext() string`

GetCiphertext returns the Ciphertext field if non-nil, zero value otherwise.

### GetCiphertextOk

`func (o *FlatJweJson) GetCiphertextOk() (*string, bool)`

GetCiphertextOk returns a tuple with the Ciphertext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphertext

`func (o *FlatJweJson) SetCiphertext(v string)`

SetCiphertext sets Ciphertext field to given value.


### GetTag

`func (o *FlatJweJson) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *FlatJweJson) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *FlatJweJson) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *FlatJweJson) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


