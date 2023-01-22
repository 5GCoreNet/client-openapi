# \PSUserStateInfoRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPsUserStateInfo**](PSUserStateInfoRetrievalApi.md#GetPsUserStateInfo) | **Get** /{imsUeId}/access-data/ps-domain/user-state | Retrieve the user state information in PS domain



## GetPsUserStateInfo

> PsUserState GetPsUserStateInfo(ctx, imsUeId).RequestedNodes(requestedNodes).PrivateIdentity(privateIdentity).Execute()

Retrieve the user state information in PS domain

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
    requestedNodes := []openapiclient.RequestedNode{*openapiclient.NewRequestedNode()} // []RequestedNode | Indicates the serving node(s) for which the request is applicable. (optional)
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PSUserStateInfoRetrievalApi.GetPsUserStateInfo(context.Background(), imsUeId).RequestedNodes(requestedNodes).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PSUserStateInfoRetrievalApi.GetPsUserStateInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPsUserStateInfo`: PsUserState
    fmt.Fprintf(os.Stdout, "Response from `PSUserStateInfoRetrievalApi.GetPsUserStateInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPsUserStateInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestedNodes** | [**[]RequestedNode**](RequestedNode.md) | Indicates the serving node(s) for which the request is applicable. | 
 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**PsUserState**](PsUserState.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

