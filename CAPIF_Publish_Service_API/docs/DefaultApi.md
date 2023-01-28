# \DefaultApi

All URIs are relative to *https://example.com/published-apis/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApfIdServiceApisGet**](DefaultApi.md#ApfIdServiceApisGet) | **Get** /{apfId}/service-apis | 
[**ApfIdServiceApisPost**](DefaultApi.md#ApfIdServiceApisPost) | **Post** /{apfId}/service-apis | 
[**ApfIdServiceApisServiceApiIdDelete**](DefaultApi.md#ApfIdServiceApisServiceApiIdDelete) | **Delete** /{apfId}/service-apis/{serviceApiId} | 
[**ApfIdServiceApisServiceApiIdGet**](DefaultApi.md#ApfIdServiceApisServiceApiIdGet) | **Get** /{apfId}/service-apis/{serviceApiId} | 
[**ApfIdServiceApisServiceApiIdPut**](DefaultApi.md#ApfIdServiceApisServiceApiIdPut) | **Put** /{apfId}/service-apis/{serviceApiId} | 



## ApfIdServiceApisGet

> []ServiceAPIDescription ApfIdServiceApisGet(ctx, apfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Publish_Service_API"
)

func main() {
    apfId := "apfId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApfIdServiceApisGet(context.Background(), apfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApfIdServiceApisGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApfIdServiceApisGet`: []ServiceAPIDescription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApfIdServiceApisGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApfIdServiceApisGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ServiceAPIDescription**](ServiceAPIDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApfIdServiceApisPost

> ServiceAPIDescription ApfIdServiceApisPost(ctx, apfId).ServiceAPIDescription(serviceAPIDescription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Publish_Service_API"
)

func main() {
    apfId := "apfId_example" // string | 
    serviceAPIDescription := *openapiclient.NewServiceAPIDescription("ApiName_example") // ServiceAPIDescription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApfIdServiceApisPost(context.Background(), apfId).ServiceAPIDescription(serviceAPIDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApfIdServiceApisPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApfIdServiceApisPost`: ServiceAPIDescription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApfIdServiceApisPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApfIdServiceApisPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceAPIDescription** | [**ServiceAPIDescription**](ServiceAPIDescription.md) |  | 

### Return type

[**ServiceAPIDescription**](ServiceAPIDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApfIdServiceApisServiceApiIdDelete

> ApfIdServiceApisServiceApiIdDelete(ctx, serviceApiId, apfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Publish_Service_API"
)

func main() {
    serviceApiId := "serviceApiId_example" // string | 
    apfId := "apfId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApfIdServiceApisServiceApiIdDelete(context.Background(), serviceApiId, apfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApfIdServiceApisServiceApiIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiId** | **string** |  | 
**apfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApfIdServiceApisServiceApiIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApfIdServiceApisServiceApiIdGet

> ServiceAPIDescription ApfIdServiceApisServiceApiIdGet(ctx, serviceApiId, apfId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Publish_Service_API"
)

func main() {
    serviceApiId := "serviceApiId_example" // string | 
    apfId := "apfId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApfIdServiceApisServiceApiIdGet(context.Background(), serviceApiId, apfId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApfIdServiceApisServiceApiIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApfIdServiceApisServiceApiIdGet`: ServiceAPIDescription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApfIdServiceApisServiceApiIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiId** | **string** |  | 
**apfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApfIdServiceApisServiceApiIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServiceAPIDescription**](ServiceAPIDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApfIdServiceApisServiceApiIdPut

> ServiceAPIDescription ApfIdServiceApisServiceApiIdPut(ctx, serviceApiId, apfId).ServiceAPIDescription(serviceAPIDescription).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Publish_Service_API"
)

func main() {
    serviceApiId := "serviceApiId_example" // string | 
    apfId := "apfId_example" // string | 
    serviceAPIDescription := *openapiclient.NewServiceAPIDescription("ApiName_example") // ServiceAPIDescription | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ApfIdServiceApisServiceApiIdPut(context.Background(), serviceApiId, apfId).ServiceAPIDescription(serviceAPIDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ApfIdServiceApisServiceApiIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApfIdServiceApisServiceApiIdPut`: ServiceAPIDescription
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ApfIdServiceApisServiceApiIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceApiId** | **string** |  | 
**apfId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApfIdServiceApisServiceApiIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceAPIDescription** | [**ServiceAPIDescription**](ServiceAPIDescription.md) |  | 

### Return type

[**ServiceAPIDescription**](ServiceAPIDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

