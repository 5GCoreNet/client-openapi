# \InitiateACRApi

All URIs are relative to *https://example.com/eees-appctxtreloc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Initiate**](InitiateACRApi.md#Initiate) | **Post** /initiate | Request the initiation of ACR.



## Initiate

> Initiate(ctx).AcrInitReq(acrInitReq).Execute()

Request the initiation of ACR.

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
    acrInitReq := *openapiclient.NewAcrInitReq("RequestorId_example", openapiclient.EndPoint{Interface{}: new(interface{})}, false) // AcrInitReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InitiateACRApi.Initiate(context.Background()).AcrInitReq(acrInitReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InitiateACRApi.Initiate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acrInitReq** | [**AcrInitReq**](AcrInitReq.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

