# \ParameterUpdateInTheUESMSContextInSMSFApi

All URIs are relative to *https://example.com/nsmsf-sms/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SMSServiceParameterUpdate**](ParameterUpdateInTheUESMSContextInSMSFApi.md#SMSServiceParameterUpdate) | **Patch** /ue-contexts/{supi} | Update a parameter in the UE SMS Context in SMSF



## SMSServiceParameterUpdate

> SMSServiceParameterUpdate200Response SMSServiceParameterUpdate(ctx, supi).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()

Update a parameter in the UE SMS Context in SMSF

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
    supi := "supi_example" // string | Subscriber Permanent Identifier (SUPI)
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ParameterUpdateInTheUESMSContextInSMSFApi.SMSServiceParameterUpdate(context.Background(), supi).PatchItem(patchItem).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ParameterUpdateInTheUESMSContextInSMSFApi.SMSServiceParameterUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SMSServiceParameterUpdate`: SMSServiceParameterUpdate200Response
    fmt.Fprintf(os.Stdout, "Response from `ParameterUpdateInTheUESMSContextInSMSFApi.SMSServiceParameterUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Subscriber Permanent Identifier (SUPI) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSMSServiceParameterUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**SMSServiceParameterUpdate200Response**](SMSServiceParameterUpdate200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

