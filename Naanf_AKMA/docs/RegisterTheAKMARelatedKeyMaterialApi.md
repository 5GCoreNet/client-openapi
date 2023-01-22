# \RegisterTheAKMARelatedKeyMaterialApi

All URIs are relative to *https://example.com/naanf-akma/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterAKMAKey**](RegisterTheAKMARelatedKeyMaterialApi.md#RegisterAKMAKey) | **Post** /register-anchorkey | Store AKMA related key material.



## RegisterAKMAKey

> AkmaKeyInfo RegisterAKMAKey(ctx).AkmaKeyInfo(akmaKeyInfo).Execute()

Store AKMA related key material.

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
    akmaKeyInfo := *openapiclient.NewAkmaKeyInfo("Supi_example", "AKId_example", "KAkma_example") // AkmaKeyInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RegisterTheAKMARelatedKeyMaterialApi.RegisterAKMAKey(context.Background()).AkmaKeyInfo(akmaKeyInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegisterTheAKMARelatedKeyMaterialApi.RegisterAKMAKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterAKMAKey`: AkmaKeyInfo
    fmt.Fprintf(os.Stdout, "Response from `RegisterTheAKMARelatedKeyMaterialApi.RegisterAKMAKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterAKMAKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **akmaKeyInfo** | [**AkmaKeyInfo**](AkmaKeyInfo.md) |  | 

### Return type

[**AkmaKeyInfo**](AkmaKeyInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

