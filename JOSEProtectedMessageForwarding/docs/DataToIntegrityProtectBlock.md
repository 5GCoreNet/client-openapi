# DataToIntegrityProtectBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | Pointer to [**MetaData**](MetaData.md) |  | [optional] 
**RequestLine** | Pointer to [**RequestLine**](RequestLine.md) |  | [optional] 
**StatusLine** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to [**[]HttpHeader**](HttpHeader.md) |  | [optional] 
**Payload** | Pointer to [**[]HttpPayload**](HttpPayload.md) |  | [optional] 

## Methods

### NewDataToIntegrityProtectBlock

`func NewDataToIntegrityProtectBlock() *DataToIntegrityProtectBlock`

NewDataToIntegrityProtectBlock instantiates a new DataToIntegrityProtectBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataToIntegrityProtectBlockWithDefaults

`func NewDataToIntegrityProtectBlockWithDefaults() *DataToIntegrityProtectBlock`

NewDataToIntegrityProtectBlockWithDefaults instantiates a new DataToIntegrityProtectBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *DataToIntegrityProtectBlock) GetMetaData() MetaData`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *DataToIntegrityProtectBlock) GetMetaDataOk() (*MetaData, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *DataToIntegrityProtectBlock) SetMetaData(v MetaData)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *DataToIntegrityProtectBlock) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetRequestLine

`func (o *DataToIntegrityProtectBlock) GetRequestLine() RequestLine`

GetRequestLine returns the RequestLine field if non-nil, zero value otherwise.

### GetRequestLineOk

`func (o *DataToIntegrityProtectBlock) GetRequestLineOk() (*RequestLine, bool)`

GetRequestLineOk returns a tuple with the RequestLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLine

`func (o *DataToIntegrityProtectBlock) SetRequestLine(v RequestLine)`

SetRequestLine sets RequestLine field to given value.

### HasRequestLine

`func (o *DataToIntegrityProtectBlock) HasRequestLine() bool`

HasRequestLine returns a boolean if a field has been set.

### GetStatusLine

`func (o *DataToIntegrityProtectBlock) GetStatusLine() string`

GetStatusLine returns the StatusLine field if non-nil, zero value otherwise.

### GetStatusLineOk

`func (o *DataToIntegrityProtectBlock) GetStatusLineOk() (*string, bool)`

GetStatusLineOk returns a tuple with the StatusLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusLine

`func (o *DataToIntegrityProtectBlock) SetStatusLine(v string)`

SetStatusLine sets StatusLine field to given value.

### HasStatusLine

`func (o *DataToIntegrityProtectBlock) HasStatusLine() bool`

HasStatusLine returns a boolean if a field has been set.

### GetHeaders

`func (o *DataToIntegrityProtectBlock) GetHeaders() []HttpHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DataToIntegrityProtectBlock) GetHeadersOk() (*[]HttpHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DataToIntegrityProtectBlock) SetHeaders(v []HttpHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DataToIntegrityProtectBlock) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetPayload

`func (o *DataToIntegrityProtectBlock) GetPayload() []HttpPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DataToIntegrityProtectBlock) GetPayloadOk() (*[]HttpPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DataToIntegrityProtectBlock) SetPayload(v []HttpPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DataToIntegrityProtectBlock) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


