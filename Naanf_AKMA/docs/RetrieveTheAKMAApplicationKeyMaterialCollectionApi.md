# \RetrieveTheAKMAApplicationKeyMaterialCollectionApi

All URIs are relative to *https://example.com/naanf-akma/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAKMAAPPKeyMaterial**](RetrieveTheAKMAApplicationKeyMaterialCollectionApi.md#GetAKMAAPPKeyMaterial) | **Post** /retrieve-applicationkey | Request to retrieve AKMA Application Key information.



## GetAKMAAPPKeyMaterial

> AkmaAfKeyData GetAKMAAPPKeyMaterial(ctx).AkmaAfKeyRequest(akmaAfKeyRequest).Execute()

Request to retrieve AKMA Application Key information.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Naanf_AKMA"
)

func main() {
    akmaAfKeyRequest := *openapiclient.NewAkmaAfKeyRequest("AfId_example", "AKId_example") // AkmaAfKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrieveTheAKMAApplicationKeyMaterialCollectionApi.GetAKMAAPPKeyMaterial(context.Background()).AkmaAfKeyRequest(akmaAfKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrieveTheAKMAApplicationKeyMaterialCollectionApi.GetAKMAAPPKeyMaterial``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAKMAAPPKeyMaterial`: AkmaAfKeyData
    fmt.Fprintf(os.Stdout, "Response from `RetrieveTheAKMAApplicationKeyMaterialCollectionApi.GetAKMAAPPKeyMaterial`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAKMAAPPKeyMaterialRequest struct via the builder pattern


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

