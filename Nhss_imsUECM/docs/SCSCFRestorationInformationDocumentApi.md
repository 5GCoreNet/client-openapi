# \SCSCFRestorationInformationDocumentApi

All URIs are relative to *https://example.com/nhss-ims-uecm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteScscfRestorationInfo**](SCSCFRestorationInformationDocumentApi.md#DeleteScscfRestorationInfo) | **Delete** /{impu}/scscf-registration/scscf-restoration-info | Delete the S-CSCF restoration information of the UE
[**GetScscfRestorationInfo**](SCSCFRestorationInformationDocumentApi.md#GetScscfRestorationInfo) | **Get** /{impu}/scscf-registration/scscf-restoration-info | Retrieve the S-CSCF restoration information of the UE
[**UpdateScscfRestorationInfo**](SCSCFRestorationInformationDocumentApi.md#UpdateScscfRestorationInfo) | **Put** /{impu}/scscf-registration/scscf-restoration-info | Update the S-CSCF restoration information of the UE



## DeleteScscfRestorationInfo

> DeleteScscfRestorationInfo(ctx, impu).Execute()

Delete the S-CSCF restoration information of the UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsUECM"
)

func main() {
    impu := "impu_example" // string | Public identity of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCSCFRestorationInformationDocumentApi.DeleteScscfRestorationInfo(context.Background(), impu).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCSCFRestorationInformationDocumentApi.DeleteScscfRestorationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**impu** | **string** | Public identity of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScscfRestorationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScscfRestorationInfo

> ScscfRestorationInfoResponse GetScscfRestorationInfo(ctx, impu).Execute()

Retrieve the S-CSCF restoration information of the UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsUECM"
)

func main() {
    impu := "impu_example" // string | Public identity of the user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCSCFRestorationInformationDocumentApi.GetScscfRestorationInfo(context.Background(), impu).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCSCFRestorationInformationDocumentApi.GetScscfRestorationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScscfRestorationInfo`: ScscfRestorationInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `SCSCFRestorationInformationDocumentApi.GetScscfRestorationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**impu** | **string** | Public identity of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScscfRestorationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScscfRestorationInfoResponse**](ScscfRestorationInfoResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScscfRestorationInfo

> ScscfRestorationInfoResponse UpdateScscfRestorationInfo(ctx, impu).ScscfRestorationInfoRequest(scscfRestorationInfoRequest).Execute()

Update the S-CSCF restoration information of the UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsUECM"
)

func main() {
    impu := "impu_example" // string | Public identity of the user.
    scscfRestorationInfoRequest := *openapiclient.NewScscfRestorationInfoRequest() // ScscfRestorationInfoRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SCSCFRestorationInformationDocumentApi.UpdateScscfRestorationInfo(context.Background(), impu).ScscfRestorationInfoRequest(scscfRestorationInfoRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCSCFRestorationInformationDocumentApi.UpdateScscfRestorationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScscfRestorationInfo`: ScscfRestorationInfoResponse
    fmt.Fprintf(os.Stdout, "Response from `SCSCFRestorationInformationDocumentApi.UpdateScscfRestorationInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**impu** | **string** | Public identity of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScscfRestorationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scscfRestorationInfoRequest** | [**ScscfRestorationInfoRequest**](ScscfRestorationInfoRequest.md) |  | 

### Return type

[**ScscfRestorationInfoResponse**](ScscfRestorationInfoResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

