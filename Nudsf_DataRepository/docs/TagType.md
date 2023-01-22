# TagType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** |  | 
**KeyType** | [**KeyType**](KeyType.md) |  | 
**Sort** | Pointer to **bool** |  | [optional] [default to false]
**Presence** | Pointer to **bool** |  | [optional] 

## Methods

### NewTagType

`func NewTagType(tagName string, keyType KeyType, ) *TagType`

NewTagType instantiates a new TagType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagTypeWithDefaults

`func NewTagTypeWithDefaults() *TagType`

NewTagTypeWithDefaults instantiates a new TagType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *TagType) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *TagType) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *TagType) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetKeyType

`func (o *TagType) GetKeyType() KeyType`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *TagType) GetKeyTypeOk() (*KeyType, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *TagType) SetKeyType(v KeyType)`

SetKeyType sets KeyType field to given value.


### GetSort

`func (o *TagType) GetSort() bool`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TagType) GetSortOk() (*bool, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TagType) SetSort(v bool)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TagType) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetPresence

`func (o *TagType) GetPresence() bool`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *TagType) GetPresenceOk() (*bool, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *TagType) SetPresence(v bool)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *TagType) HasPresence() bool`

HasPresence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


