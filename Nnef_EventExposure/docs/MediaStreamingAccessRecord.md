# MediaStreamingAccessRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**MediaStreamHandlerEndpointAddress** | [**EndpointAddress**](EndpointAddress.md) |  | 
**ApplicationServerEndpointAddress** | [**EndpointAddress**](EndpointAddress.md) |  | 
**SessionIdentifier** | Pointer to **string** |  | [optional] 
**RequestMessage** | [**MediaStreamingAccessRecordAllOfRequestMessage**](MediaStreamingAccessRecordAllOfRequestMessage.md) |  | 
**CacheStatus** | Pointer to [**CacheStatus**](CacheStatus.md) |  | [optional] 
**ResponseMessage** | [**MediaStreamingAccessRecordAllOfResponseMessage**](MediaStreamingAccessRecordAllOfResponseMessage.md) |  | 
**ProcessingLatency** | **float32** | string with format &#39;float&#39; as defined in OpenAPI. | 
**ConnectionMetrics** | Pointer to [**MediaStreamingAccessRecordAllOfConnectionMetrics**](MediaStreamingAccessRecordAllOfConnectionMetrics.md) |  | [optional] 

## Methods

### NewMediaStreamingAccessRecord

`func NewMediaStreamingAccessRecord(timestamp time.Time, mediaStreamHandlerEndpointAddress EndpointAddress, applicationServerEndpointAddress EndpointAddress, requestMessage MediaStreamingAccessRecordAllOfRequestMessage, responseMessage MediaStreamingAccessRecordAllOfResponseMessage, processingLatency float32, ) *MediaStreamingAccessRecord`

NewMediaStreamingAccessRecord instantiates a new MediaStreamingAccessRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaStreamingAccessRecordWithDefaults

`func NewMediaStreamingAccessRecordWithDefaults() *MediaStreamingAccessRecord`

NewMediaStreamingAccessRecordWithDefaults instantiates a new MediaStreamingAccessRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *MediaStreamingAccessRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MediaStreamingAccessRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MediaStreamingAccessRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetMediaStreamHandlerEndpointAddress

`func (o *MediaStreamingAccessRecord) GetMediaStreamHandlerEndpointAddress() EndpointAddress`

GetMediaStreamHandlerEndpointAddress returns the MediaStreamHandlerEndpointAddress field if non-nil, zero value otherwise.

### GetMediaStreamHandlerEndpointAddressOk

`func (o *MediaStreamingAccessRecord) GetMediaStreamHandlerEndpointAddressOk() (*EndpointAddress, bool)`

GetMediaStreamHandlerEndpointAddressOk returns a tuple with the MediaStreamHandlerEndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreamHandlerEndpointAddress

`func (o *MediaStreamingAccessRecord) SetMediaStreamHandlerEndpointAddress(v EndpointAddress)`

SetMediaStreamHandlerEndpointAddress sets MediaStreamHandlerEndpointAddress field to given value.


### GetApplicationServerEndpointAddress

`func (o *MediaStreamingAccessRecord) GetApplicationServerEndpointAddress() EndpointAddress`

GetApplicationServerEndpointAddress returns the ApplicationServerEndpointAddress field if non-nil, zero value otherwise.

### GetApplicationServerEndpointAddressOk

`func (o *MediaStreamingAccessRecord) GetApplicationServerEndpointAddressOk() (*EndpointAddress, bool)`

GetApplicationServerEndpointAddressOk returns a tuple with the ApplicationServerEndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationServerEndpointAddress

`func (o *MediaStreamingAccessRecord) SetApplicationServerEndpointAddress(v EndpointAddress)`

SetApplicationServerEndpointAddress sets ApplicationServerEndpointAddress field to given value.


### GetSessionIdentifier

`func (o *MediaStreamingAccessRecord) GetSessionIdentifier() string`

GetSessionIdentifier returns the SessionIdentifier field if non-nil, zero value otherwise.

### GetSessionIdentifierOk

`func (o *MediaStreamingAccessRecord) GetSessionIdentifierOk() (*string, bool)`

GetSessionIdentifierOk returns a tuple with the SessionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdentifier

`func (o *MediaStreamingAccessRecord) SetSessionIdentifier(v string)`

SetSessionIdentifier sets SessionIdentifier field to given value.

### HasSessionIdentifier

`func (o *MediaStreamingAccessRecord) HasSessionIdentifier() bool`

HasSessionIdentifier returns a boolean if a field has been set.

### GetRequestMessage

`func (o *MediaStreamingAccessRecord) GetRequestMessage() MediaStreamingAccessRecordAllOfRequestMessage`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *MediaStreamingAccessRecord) GetRequestMessageOk() (*MediaStreamingAccessRecordAllOfRequestMessage, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *MediaStreamingAccessRecord) SetRequestMessage(v MediaStreamingAccessRecordAllOfRequestMessage)`

SetRequestMessage sets RequestMessage field to given value.


### GetCacheStatus

`func (o *MediaStreamingAccessRecord) GetCacheStatus() CacheStatus`

GetCacheStatus returns the CacheStatus field if non-nil, zero value otherwise.

### GetCacheStatusOk

`func (o *MediaStreamingAccessRecord) GetCacheStatusOk() (*CacheStatus, bool)`

GetCacheStatusOk returns a tuple with the CacheStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheStatus

`func (o *MediaStreamingAccessRecord) SetCacheStatus(v CacheStatus)`

SetCacheStatus sets CacheStatus field to given value.

### HasCacheStatus

`func (o *MediaStreamingAccessRecord) HasCacheStatus() bool`

HasCacheStatus returns a boolean if a field has been set.

### GetResponseMessage

`func (o *MediaStreamingAccessRecord) GetResponseMessage() MediaStreamingAccessRecordAllOfResponseMessage`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *MediaStreamingAccessRecord) GetResponseMessageOk() (*MediaStreamingAccessRecordAllOfResponseMessage, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *MediaStreamingAccessRecord) SetResponseMessage(v MediaStreamingAccessRecordAllOfResponseMessage)`

SetResponseMessage sets ResponseMessage field to given value.


### GetProcessingLatency

`func (o *MediaStreamingAccessRecord) GetProcessingLatency() float32`

GetProcessingLatency returns the ProcessingLatency field if non-nil, zero value otherwise.

### GetProcessingLatencyOk

`func (o *MediaStreamingAccessRecord) GetProcessingLatencyOk() (*float32, bool)`

GetProcessingLatencyOk returns a tuple with the ProcessingLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingLatency

`func (o *MediaStreamingAccessRecord) SetProcessingLatency(v float32)`

SetProcessingLatency sets ProcessingLatency field to given value.


### GetConnectionMetrics

`func (o *MediaStreamingAccessRecord) GetConnectionMetrics() MediaStreamingAccessRecordAllOfConnectionMetrics`

GetConnectionMetrics returns the ConnectionMetrics field if non-nil, zero value otherwise.

### GetConnectionMetricsOk

`func (o *MediaStreamingAccessRecord) GetConnectionMetricsOk() (*MediaStreamingAccessRecordAllOfConnectionMetrics, bool)`

GetConnectionMetricsOk returns a tuple with the ConnectionMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionMetrics

`func (o *MediaStreamingAccessRecord) SetConnectionMetrics(v MediaStreamingAccessRecordAllOfConnectionMetrics)`

SetConnectionMetrics sets ConnectionMetrics field to given value.

### HasConnectionMetrics

`func (o *MediaStreamingAccessRecord) HasConnectionMetrics() bool`

HasConnectionMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


