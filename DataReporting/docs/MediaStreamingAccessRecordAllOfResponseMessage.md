# MediaStreamingAccessRecordAllOfResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**Size** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**BodySize** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**ContentType** | Pointer to **string** |  | [optional] 

## Methods

### NewMediaStreamingAccessRecordAllOfResponseMessage

`func NewMediaStreamingAccessRecordAllOfResponseMessage(responseCode int32, size int32, bodySize int32, ) *MediaStreamingAccessRecordAllOfResponseMessage`

NewMediaStreamingAccessRecordAllOfResponseMessage instantiates a new MediaStreamingAccessRecordAllOfResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaStreamingAccessRecordAllOfResponseMessageWithDefaults

`func NewMediaStreamingAccessRecordAllOfResponseMessageWithDefaults() *MediaStreamingAccessRecordAllOfResponseMessage`

NewMediaStreamingAccessRecordAllOfResponseMessageWithDefaults instantiates a new MediaStreamingAccessRecordAllOfResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetResponseCode() int32`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetResponseCodeOk() (*int32, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetResponseCode(v int32)`

SetResponseCode sets ResponseCode field to given value.


### GetSize

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetBodySize

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetBodySize() int32`

GetBodySize returns the BodySize field if non-nil, zero value otherwise.

### GetBodySizeOk

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetBodySizeOk() (*int32, bool)`

GetBodySizeOk returns a tuple with the BodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySize

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetBodySize(v int32)`

SetBodySize sets BodySize field to given value.


### GetContentType

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MediaStreamingAccessRecordAllOfResponseMessage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


