# \DefaultApi

All URIs are relative to *https://example.com/3gpp-akma/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveAKMAAppKey**](DefaultApi.md#RetrieveAKMAAppKey) | **Post** /retrieve | Retrieve AKMA Application Key Information.



## RetrieveAKMAAppKey

> AkmaAfKeyData RetrieveAKMAAppKey(ctx).AkmaAfKeyRequest(akmaAfKeyRequest).Execute()

Retrieve AKMA Application Key Information.

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
    akmaAfKeyRequest := *openapiclient.NewAkmaAfKeyRequest("AfId_example", "AKId_example") // AkmaAfKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RetrieveAKMAAppKey(context.Background()).AkmaAfKeyRequest(akmaAfKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RetrieveAKMAAppKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveAKMAAppKey`: AkmaAfKeyData
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RetrieveAKMAAppKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveAKMAAppKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **akmaAfKeyRequest** | [**AkmaAfKeyRequest**](AkmaAfKeyRequest.md) |  | 

### Return type

[**AkmaAfKeyData**](AkmaAfKeyData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

