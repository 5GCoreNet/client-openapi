# \GenerateAuthVectorApi

All URIs are relative to *https://example.com/nhss-ueau/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateAV**](GenerateAuthVectorApi.md#GenerateAV) | **Post** /generate-av | Generate authentication vector for the UE



## GenerateAV

> AvGenerationResponse GenerateAV(ctx).AvGenerationRequest(avGenerationRequest).Execute()

Generate authentication vector for the UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_UEAU"
)

func main() {
    avGenerationRequest := *openapiclient.NewAvGenerationRequest("Imsi_example", *openapiclient.NewAuthType(), "ServingNetworkName_example") // AvGenerationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GenerateAuthVectorApi.GenerateAV(context.Background()).AvGenerationRequest(avGenerationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenerateAuthVectorApi.GenerateAV``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAV`: AvGenerationResponse
    fmt.Fprintf(os.Stdout, "Response from `GenerateAuthVectorApi.GenerateAV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avGenerationRequest** | [**AvGenerationRequest**](AvGenerationRequest.md) |  | 

### Return type

[**AvGenerationResponse**](AvGenerationResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

