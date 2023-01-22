# MetaSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemaId** | **string** | Represents the Identifier of a Meta schema. | 
**MetaTags** | [**[]TagType**](TagType.md) |  | 

## Methods

### NewMetaSchema

`func NewMetaSchema(schemaId string, metaTags []TagType, ) *MetaSchema`

NewMetaSchema instantiates a new MetaSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaSchemaWithDefaults

`func NewMetaSchemaWithDefaults() *MetaSchema`

NewMetaSchemaWithDefaults instantiates a new MetaSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemaId

`func (o *MetaSchema) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *MetaSchema) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *MetaSchema) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.


### GetMetaTags

`func (o *MetaSchema) GetMetaTags() []TagType`

GetMetaTags returns the MetaTags field if non-nil, zero value otherwise.

### GetMetaTagsOk

`func (o *MetaSchema) GetMetaTagsOk() (*[]TagType, bool)`

GetMetaTagsOk returns a tuple with the MetaTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTags

`func (o *MetaSchema) SetMetaTags(v []TagType)`

SetMetaTags sets MetaTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


