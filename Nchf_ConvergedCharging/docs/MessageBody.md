# MessageBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | 
**ContentLength** | **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | 
**ContentDisposition** | Pointer to **string** |  | [optional] 
**Originator** | Pointer to [**OriginatorPartyType**](OriginatorPartyType.md) |  | [optional] 

## Methods

### NewMessageBody

`func NewMessageBody(contentType string, contentLength int32, ) *MessageBody`

NewMessageBody instantiates a new MessageBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageBodyWithDefaults

`func NewMessageBodyWithDefaults() *MessageBody`

NewMessageBodyWithDefaults instantiates a new MessageBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *MessageBody) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MessageBody) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MessageBody) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetContentLength

`func (o *MessageBody) GetContentLength() int32`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *MessageBody) GetContentLengthOk() (*int32, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *MessageBody) SetContentLength(v int32)`

SetContentLength sets ContentLength field to given value.


### GetContentDisposition

`func (o *MessageBody) GetContentDisposition() string`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *MessageBody) GetContentDispositionOk() (*string, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *MessageBody) SetContentDisposition(v string)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *MessageBody) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetOriginator

`func (o *MessageBody) GetOriginator() OriginatorPartyType`

GetOriginator returns the Originator field if non-nil, zero value otherwise.

### GetOriginatorOk

`func (o *MessageBody) GetOriginatorOk() (*OriginatorPartyType, bool)`

GetOriginatorOk returns a tuple with the Originator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginator

`func (o *MessageBody) SetOriginator(v OriginatorPartyType)`

SetOriginator sets Originator field to given value.

### HasOriginator

`func (o *MessageBody) HasOriginator() bool`

HasOriginator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


