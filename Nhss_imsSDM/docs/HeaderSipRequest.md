# HeaderSipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | **string** |  | 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewHeaderSipRequest

`func NewHeaderSipRequest(header string, ) *HeaderSipRequest`

NewHeaderSipRequest instantiates a new HeaderSipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeaderSipRequestWithDefaults

`func NewHeaderSipRequestWithDefaults() *HeaderSipRequest`

NewHeaderSipRequestWithDefaults instantiates a new HeaderSipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *HeaderSipRequest) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *HeaderSipRequest) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *HeaderSipRequest) SetHeader(v string)`

SetHeader sets Header field to given value.


### GetContent

`func (o *HeaderSipRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HeaderSipRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HeaderSipRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *HeaderSipRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


