# \CSUserStateInfoRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCsUserStateInfo**](CSUserStateInfoRetrievalApi.md#GetCsUserStateInfo) | **Get** /{imsUeId}/access-data/cs-domain/user-state | Retrieve the user state information in CS domain



## GetCsUserStateInfo

> CsUserState GetCsUserStateInfo(ctx, imsUeId).PrivateIdentity(privateIdentity).Execute()

Retrieve the user state information in CS domain

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSUserStateInfoRetrievalApi.GetCsUserStateInfo(context.Background(), imsUeId).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSUserStateInfoRetrievalApi.GetCsUserStateInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCsUserStateInfo`: CsUserState
    fmt.Fprintf(os.Stdout, "Response from `CSUserStateInfoRetrievalApi.GetCsUserStateInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCsUserStateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**CsUserState**](CsUserState.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

