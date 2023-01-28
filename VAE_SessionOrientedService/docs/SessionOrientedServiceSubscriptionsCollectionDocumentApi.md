# \SessionOrientedServiceSubscriptionsCollectionDocumentApi

All URIs are relative to *https://example.com/vae-session-Oriented-service/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SessionOrientedServiceSubscriptionsCollectionDocumentApi.md#Create) | **Post** /subscriptions | VAE_SessionOrientedService resource create service Operation



## Create

> SessionOrientedData Create(ctx).SessionOrientedData(sessionOrientedData).Execute()

VAE_SessionOrientedService resource create service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/VAE_SessionOrientedService"
)

func main() {
    sessionOrientedData := *openapiclient.NewSessionOrientedData("UeId_example", "NotifUri_example", "ServiceId_example", "AppSerId_example") // SessionOrientedData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionOrientedServiceSubscriptionsCollectionDocumentApi.Create(context.Background()).SessionOrientedData(sessionOrientedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionOrientedServiceSubscriptionsCollectionDocumentApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: SessionOrientedData
    fmt.Fprintf(os.Stdout, "Response from `SessionOrientedServiceSubscriptionsCollectionDocumentApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionOrientedData** | [**SessionOrientedData**](SessionOrientedData.md) |  | 

### Return type

[**SessionOrientedData**](SessionOrientedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

