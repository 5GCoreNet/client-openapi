# \AuthenticationStatusDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthenticationStatus**](AuthenticationStatusDocumentApi.md#CreateAuthenticationStatus) | **Put** /subscription-data/{ueId}/authentication-data/authentication-status | To store the Authentication Status data of a UE



## CreateAuthenticationStatus

> CreateAuthenticationStatus(ctx, ueId).AuthEvent(authEvent).Execute()

To store the Authentication Status data of a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    ueId := "ueId_example" // string | UE id
    authEvent := *openapiclient.NewAuthEvent("NfInstanceId_example", false, time.Now(), *openapiclient.NewAuthType(), "ServingNetworkName_example") // AuthEvent | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationStatusDocumentApi.CreateAuthenticationStatus(context.Background(), ueId).AuthEvent(authEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationStatusDocumentApi.CreateAuthenticationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthenticationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authEvent** | [**AuthEvent**](AuthEvent.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

