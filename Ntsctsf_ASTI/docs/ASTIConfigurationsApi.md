# \ASTIConfigurationsApi

All URIs are relative to *https://example.com/ntsctsf-asti/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestStatusof5GAccessStratumTimeDistribution**](ASTIConfigurationsApi.md#RequestStatusof5GAccessStratumTimeDistribution) | **Post** /configurations/retrieve | Request the status of the 5G access stratum time distribution for a list of UEs.



## RequestStatusof5GAccessStratumTimeDistribution

> StatusResponseData RequestStatusof5GAccessStratumTimeDistribution(ctx).StatusRequestData(statusRequestData).Execute()

Request the status of the 5G access stratum time distribution for a list of UEs.

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
    statusRequestData := openapiclient.StatusRequestData{Interface{}: new(interface{})} // StatusRequestData | Contains the information for the status of the 5G access stratum time distribution. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ASTIConfigurationsApi.RequestStatusof5GAccessStratumTimeDistribution(context.Background()).StatusRequestData(statusRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ASTIConfigurationsApi.RequestStatusof5GAccessStratumTimeDistribution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestStatusof5GAccessStratumTimeDistribution`: StatusResponseData
    fmt.Fprintf(os.Stdout, "Response from `ASTIConfigurationsApi.RequestStatusof5GAccessStratumTimeDistribution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestStatusof5GAccessStratumTimeDistributionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusRequestData** | [**StatusRequestData**](StatusRequestData.md) | Contains the information for the status of the 5G access stratum time distribution.  | 

### Return type

[**StatusResponseData**](StatusResponseData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

