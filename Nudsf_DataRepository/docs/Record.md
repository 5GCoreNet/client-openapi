# Record

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**RecordMeta**](RecordMeta.md) |  | 
**Blocks** | Pointer to **[]interface{}** | list of opaque Block&#39;s in this Record | [optional] 

## Methods

### NewRecord

`func NewRecord(meta RecordMeta, ) *Record`

NewRecord instantiates a new Record object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordWithDefaults

`func NewRecordWithDefaults() *Record`

NewRecordWithDefaults instantiates a new Record object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Record) GetMeta() RecordMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Record) GetMetaOk() (*RecordMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Record) SetMeta(v RecordMeta)`

SetMeta sets Meta field to given value.


### GetBlocks

`func (o *Record) GetBlocks() []interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *Record) GetBlocksOk() (*[]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *Record) SetBlocks(v []interface{})`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *Record) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


