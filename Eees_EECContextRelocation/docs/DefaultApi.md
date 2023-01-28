# \DefaultApi

All URIs are relative to *https://example.com/eees-eeccontextreloc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EecContextsGet**](DefaultApi.md#EecContextsGet) | **Get** /eec-contexts | 
[**EecContextsPost**](DefaultApi.md#EecContextsPost) | **Post** /eec-contexts | 



## EecContextsGet

> EECContext EecContextsGet(ctx).EesId(eesId).EecCntxId(eecCntxId).SessCntxs(sessCntxs).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EECContextRelocation"
)

func main() {
    eesId := "eesId_example" // string | Unique identifier of the requesting EES.
    eecCntxId := "eecCntxId_example" // string | Unique identifier of the EEC context.
    sessCntxs := map[string][]openapiclient.SessionContexts{ ... } // SessionContexts | List of service session context information being requested. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EecContextsGet(context.Background()).EesId(eesId).EecCntxId(eecCntxId).SessCntxs(sessCntxs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EecContextsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EecContextsGet`: EECContext
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EecContextsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEecContextsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eesId** | **string** | Unique identifier of the requesting EES. | 
 **eecCntxId** | **string** | Unique identifier of the EEC context. | 
 **sessCntxs** | [**SessionContexts**](SessionContexts.md) | List of service session context information being requested. | 

### Return type

[**EECContext**](EECContext.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EecContextsPost

> EECContextPushRes EecContextsPost(ctx).EECContextPush(eECContextPush).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Eees_EECContextRelocation"
)

func main() {
    eECContextPush := *openapiclient.NewEECContextPush("EesId_example", *openapiclient.NewEECContext("EecId_example", "CntxId_example")) // EECContextPush | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EecContextsPost(context.Background()).EECContextPush(eECContextPush).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EecContextsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EecContextsPost`: EECContextPushRes
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EecContextsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEecContextsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eECContextPush** | [**EECContextPush**](EECContextPush.md) |  | 

### Return type

[**EECContextPushRes**](EECContextPushRes.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

