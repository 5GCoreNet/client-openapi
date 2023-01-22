# UriList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksValueSchema**](LinksValueSchema.md) | List of the URI of NF instances. It has two members whose names are item and self. The item attribute contains an array of URIs.  | [optional] 
**TotalItemCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewUriList

`func NewUriList() *UriList`

NewUriList instantiates a new UriList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUriListWithDefaults

`func NewUriListWithDefaults() *UriList`

NewUriListWithDefaults instantiates a new UriList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UriList) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UriList) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UriList) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UriList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTotalItemCount

`func (o *UriList) GetTotalItemCount() int32`

GetTotalItemCount returns the TotalItemCount field if non-nil, zero value otherwise.

### GetTotalItemCountOk

`func (o *UriList) GetTotalItemCountOk() (*int32, bool)`

GetTotalItemCountOk returns a tuple with the TotalItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItemCount

`func (o *UriList) SetTotalItemCount(v int32)`

SetTotalItemCount sets TotalItemCount field to given value.

### HasTotalItemCount

`func (o *UriList) HasTotalItemCount() bool`

HasTotalItemCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


