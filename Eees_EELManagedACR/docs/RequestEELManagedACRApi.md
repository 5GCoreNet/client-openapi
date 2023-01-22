# \RequestEELManagedACRApi

All URIs are relative to *https://example.com/eees-eel-acr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestEELManagedACR**](RequestEELManagedACRApi.md#RequestEELManagedACR) | **Post** /request-eelacr | Request the EES (e.g. S-EES) to handle all the operations of an ACR.



## RequestEELManagedACR

> EELACRResp RequestEELManagedACR(ctx).EELACRReq(eELACRReq).Execute()

Request the EES (e.g. S-EES) to handle all the operations of an ACR.

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
    eELACRReq := *openapiclient.NewEELACRReq("UeId_example", []openapiclient.EasCharacteristics{*openapiclient.NewEasCharacteristics()}) // EELACRReq | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RequestEELManagedACRApi.RequestEELManagedACR(context.Background()).EELACRReq(eELACRReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RequestEELManagedACRApi.RequestEELManagedACR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestEELManagedACR`: EELACRResp
    fmt.Fprintf(os.Stdout, "Response from `RequestEELManagedACRApi.RequestEELManagedACR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestEELManagedACRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eELACRReq** | [**EELACRReq**](EELACRReq.md) |  | 

### Return type

[**EELACRResp**](EELACRResp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

