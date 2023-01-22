# \CreationOfRoutingInfoApi

All URIs are relative to *https://example.com/nipsmgw-smservice/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RoutingInfo**](CreationOfRoutingInfoApi.md#RoutingInfo) | **Put** /mt-sm-infos/{gpsi} | Create the routing information for a given UE



## RoutingInfo

> CreatedRoutingData RoutingInfo(ctx, gpsi).CreateRoutingData(createRoutingData).Execute()

Create the routing information for a given UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    gpsi := "gpsi_example" // string | Generic Public Subscription Identifier (GPSI)
    createRoutingData := *openapiclient.NewCreateRoutingData("SmsfId_example") // CreateRoutingData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreationOfRoutingInfoApi.RoutingInfo(context.Background(), gpsi).CreateRoutingData(createRoutingData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreationOfRoutingInfoApi.RoutingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RoutingInfo`: CreatedRoutingData
    fmt.Fprintf(os.Stdout, "Response from `CreationOfRoutingInfoApi.RoutingInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gpsi** | **string** | Generic Public Subscription Identifier (GPSI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRoutingInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRoutingData** | [**CreateRoutingData**](CreateRoutingData.md) |  | 

### Return type

[**CreatedRoutingData**](CreatedRoutingData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

