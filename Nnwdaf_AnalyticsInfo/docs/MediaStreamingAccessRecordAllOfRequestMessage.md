# MediaStreamingAccessRecordAllOfRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** |  | 
**Url** | **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | 
**ProtocolVersion** | **string** |  | 
**Range** | Pointer to **string** |  | [optional] 
**Size** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**BodySize** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**ContentType** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**UserIdentity** | Pointer to **string** |  | [optional] 
**Referer** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 

## Methods

### NewMediaStreamingAccessRecordAllOfRequestMessage

`func NewMediaStreamingAccessRecordAllOfRequestMessage(method string, url string, protocolVersion string, size int32, bodySize int32, ) *MediaStreamingAccessRecordAllOfRequestMessage`

NewMediaStreamingAccessRecordAllOfRequestMessage instantiates a new MediaStreamingAccessRecordAllOfRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaStreamingAccessRecordAllOfRequestMessageWithDefaults

`func NewMediaStreamingAccessRecordAllOfRequestMessageWithDefaults() *MediaStreamingAccessRecordAllOfRequestMessage`

NewMediaStreamingAccessRecordAllOfRequestMessageWithDefaults instantiates a new MediaStreamingAccessRecordAllOfRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetUrl

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetProtocolVersion

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetProtocolVersion() string`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetProtocolVersionOk() (*string, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetProtocolVersion(v string)`

SetProtocolVersion sets ProtocolVersion field to given value.


### GetRange

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSize

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetBodySize

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetBodySize() int32`

GetBodySize returns the BodySize field if non-nil, zero value otherwise.

### GetBodySizeOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetBodySizeOk() (*int32, bool)`

GetBodySizeOk returns a tuple with the BodySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodySize

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetBodySize(v int32)`

SetBodySize sets BodySize field to given value.


### GetContentType

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetUserAgent

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserIdentity

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserIdentity() string`

GetUserIdentity returns the UserIdentity field if non-nil, zero value otherwise.

### GetUserIdentityOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetUserIdentityOk() (*string, bool)`

GetUserIdentityOk returns a tuple with the UserIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdentity

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetUserIdentity(v string)`

SetUserIdentity sets UserIdentity field to given value.

### HasUserIdentity

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasUserIdentity() bool`

HasUserIdentity returns a boolean if a field has been set.

### GetReferer

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetReferer() string`

GetReferer returns the Referer field if non-nil, zero value otherwise.

### GetRefererOk

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) GetRefererOk() (*string, bool)`

GetRefererOk returns a tuple with the Referer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferer

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) SetReferer(v string)`

SetReferer sets Referer field to given value.

### HasReferer

`func (o *MediaStreamingAccessRecordAllOfRequestMessage) HasReferer() bool`

HasReferer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


