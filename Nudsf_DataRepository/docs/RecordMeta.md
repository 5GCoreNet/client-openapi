# RecordMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**CallbackReference** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**Tags** | Pointer to **map[string][]string** | A dictionary of {\&quot;tagName\&quot;: [ \&quot;tagValue\&quot;, ...] }. A tag name can be used to retrieve a Record. The tagValue are unique. | [optional] 

## Methods

### NewRecordMeta

`func NewRecordMeta() *RecordMeta`

NewRecordMeta instantiates a new RecordMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordMetaWithDefaults

`func NewRecordMetaWithDefaults() *RecordMeta`

NewRecordMetaWithDefaults instantiates a new RecordMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *RecordMeta) GetTtl() time.Time`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordMeta) GetTtlOk() (*time.Time, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordMeta) SetTtl(v time.Time)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordMeta) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetCallbackReference

`func (o *RecordMeta) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *RecordMeta) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *RecordMeta) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.

### HasCallbackReference

`func (o *RecordMeta) HasCallbackReference() bool`

HasCallbackReference returns a boolean if a field has been set.

### GetTags

`func (o *RecordMeta) GetTags() map[string][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RecordMeta) GetTagsOk() (*map[string][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RecordMeta) SetTags(v map[string][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RecordMeta) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


