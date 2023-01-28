# \DefaultApi

All URIs are relative to *https://example.com/3GPPManagement/StreamingDataReportingMnS*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectionsConnectionIdGet**](DefaultApi.md#ConnectionsConnectionIdGet) | **Get** /connections/{connectionId} | Obtain information about a connection.
[**ConnectionsConnectionIdStreamsDelete**](DefaultApi.md#ConnectionsConnectionIdStreamsDelete) | **Delete** /connections/{connectionId}/streams | Remove reporting streams from an existing connection
[**ConnectionsConnectionIdStreamsGet**](DefaultApi.md#ConnectionsConnectionIdStreamsGet) | **Get** /connections/{connectionId}/streams | Obtain information about streams.
[**ConnectionsConnectionIdStreamsPost**](DefaultApi.md#ConnectionsConnectionIdStreamsPost) | **Post** /connections/{connectionId}/streams | Inform consumer about new reporting streams on an existing connection.
[**ConnectionsConnectionIdStreamsStreamIdGet**](DefaultApi.md#ConnectionsConnectionIdStreamsStreamIdGet) | **Get** /connections/{connectionId}/streams/{streamId} | Obtain information about stream
[**ConnectionsGet**](DefaultApi.md#ConnectionsGet) | **Get** /connections | Obtain information about connections.
[**ConnectionsPost**](DefaultApi.md#ConnectionsPost) | **Post** /connections | Inform consumer about reporting streams to be carried by the new connection and receive a new connection id.



## ConnectionsConnectionIdGet

> ConnectionInfoType ConnectionsConnectionIdGet(ctx, connectionId).Connection(connection).SecWebSocketExtensions(secWebSocketExtensions).SecWebSocketKey(secWebSocketKey).SecWebSocketProtocol(secWebSocketProtocol).SecWebSocketVersion(secWebSocketVersion).Execute()

Obtain information about a connection.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionId := "connectionId_example" // string | Indicate the ID (URI) of the connection for which the information is being retrieved
    connection := openapiclient.websocketHeaderConnection-Type("Upgrade") // WebsocketHeaderConnectionType |  (optional)
    secWebSocketExtensions := "secWebSocketExtensions_example" // string |  (optional)
    secWebSocketKey := "secWebSocketKey_example" // string |  (optional)
    secWebSocketProtocol := "secWebSocketProtocol_example" // string |  (optional)
    secWebSocketVersion := "secWebSocketVersion_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsConnectionIdGet(context.Background(), connectionId).Connection(connection).SecWebSocketExtensions(secWebSocketExtensions).SecWebSocketKey(secWebSocketKey).SecWebSocketProtocol(secWebSocketProtocol).SecWebSocketVersion(secWebSocketVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsConnectionIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsConnectionIdGet`: ConnectionInfoType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectionsConnectionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Indicate the ID (URI) of the connection for which the information is being retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsConnectionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connection** | [**WebsocketHeaderConnectionType**](WebsocketHeaderConnectionType.md) |  | 
 **secWebSocketExtensions** | **string** |  | 
 **secWebSocketKey** | **string** |  | 
 **secWebSocketProtocol** | **string** |  | 
 **secWebSocketVersion** | **string** |  | 

### Return type

[**ConnectionInfoType**](ConnectionInfoType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsConnectionIdStreamsDelete

> ConnectionsConnectionIdStreamsDelete(ctx, connectionId).StreamIds(streamIds).Execute()

Remove reporting streams from an existing connection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionId := "connectionId_example" // string | Indicate the ID (URI) of the connection for which the reporting stream information is being removed.
    streamIds := []string{"26F452550021"} // []string | The list of streamId for the stream(s) to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsConnectionIdStreamsDelete(context.Background(), connectionId).StreamIds(streamIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsConnectionIdStreamsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Indicate the ID (URI) of the connection for which the reporting stream information is being removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsConnectionIdStreamsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamIds** | **[]string** | The list of streamId for the stream(s) to be deleted. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsConnectionIdStreamsGet

> []StreamInfoWithReportersType ConnectionsConnectionIdStreamsGet(ctx, connectionId).StreamIds(streamIds).Execute()

Obtain information about streams.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionId := "connectionId_example" // string | Indicate the ID (URI) of the connection for which the information is being retrieved
    streamIds := []string{"26F452550021"} // []string | The list of streamId for which the stream information is to be retrieved.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsConnectionIdStreamsGet(context.Background(), connectionId).StreamIds(streamIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsConnectionIdStreamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsConnectionIdStreamsGet`: []StreamInfoWithReportersType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectionsConnectionIdStreamsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Indicate the ID (URI) of the connection for which the information is being retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsConnectionIdStreamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamIds** | **[]string** | The list of streamId for which the stream information is to be retrieved. | 

### Return type

[**[]StreamInfoWithReportersType**](StreamInfoWithReportersType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsConnectionIdStreamsPost

> []StreamInfoType ConnectionsConnectionIdStreamsPost(ctx, connectionId).StreamInfoType(streamInfoType).Execute()

Inform consumer about new reporting streams on an existing connection.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionId := "connectionId_example" // string | Indicate the ID (URI) of the connection for which the reporting stream information is being added.
    streamInfoType := []openapiclient.StreamInfoType{*openapiclient.NewStreamInfoType(openapiclient.streamType-Type("TRACE"), openapiclient.serializationFormat-Type("GPB"), openapiclient.streamInfo_Type_streamId{String: new(string)})} // []StreamInfoType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsConnectionIdStreamsPost(context.Background(), connectionId).StreamInfoType(streamInfoType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsConnectionIdStreamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsConnectionIdStreamsPost`: []StreamInfoType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectionsConnectionIdStreamsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Indicate the ID (URI) of the connection for which the reporting stream information is being added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsConnectionIdStreamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **streamInfoType** | [**[]StreamInfoType**](StreamInfoType.md) |  | 

### Return type

[**[]StreamInfoType**](StreamInfoType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsConnectionIdStreamsStreamIdGet

> StreamInfoWithReportersType ConnectionsConnectionIdStreamsStreamIdGet(ctx, connectionId, streamId).Execute()

Obtain information about stream



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionId := "connectionId_example" // string | Indicate the ID (URI) of the connection for which the information is being retrieved
    streamId := "streamId_example" // string | Indicate the ID of the reporting stream for which the information is being retrieved

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsConnectionIdStreamsStreamIdGet(context.Background(), connectionId, streamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsConnectionIdStreamsStreamIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsConnectionIdStreamsStreamIdGet`: StreamInfoWithReportersType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectionsConnectionIdStreamsStreamIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionId** | **string** | Indicate the ID (URI) of the connection for which the information is being retrieved | 
**streamId** | **string** | Indicate the ID of the reporting stream for which the information is being retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsConnectionIdStreamsStreamIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StreamInfoWithReportersType**](StreamInfoWithReportersType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsGet

> []ConnectionInfoType ConnectionsGet(ctx).ConnectionIdList(connectionIdList).Execute()

Obtain information about connections.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionIdList := []string{"Inner_example"} // []string | The list of connectionId for which the connection information is to be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsGet(context.Background()).ConnectionIdList(connectionIdList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectionsGet`: []ConnectionInfoType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ConnectionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionIdList** | **[]string** | The list of connectionId for which the connection information is to be returned. | 

### Return type

[**[]ConnectionInfoType**](ConnectionInfoType.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectionsPost

> ConnectionsPost(ctx).ConnectionRequestType(connectionRequestType).Execute()

Inform consumer about reporting streams to be carried by the new connection and receive a new connection id.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/StreamingDataMnS"
)

func main() {
    connectionRequestType := *openapiclient.NewConnectionRequestType() // ConnectionRequestType | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ConnectionsPost(context.Background()).ConnectionRequestType(connectionRequestType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ConnectionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionRequestType** | [**ConnectionRequestType**](ConnectionRequestType.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

