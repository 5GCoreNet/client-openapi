# \PCSCFRestorationIndicationApi

All URIs are relative to *https://example.com/npcf-policyauthorization/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PcscfRestoration**](PCSCFRestorationIndicationApi.md#PcscfRestoration) | **Post** /app-sessions/pcscf-restoration | Indicates P-CSCF restoration and does not create an Individual Application Session Context



## PcscfRestoration

> PcscfRestoration(ctx).PcscfRestorationRequestData(pcscfRestorationRequestData).Execute()

Indicates P-CSCF restoration and does not create an Individual Application Session Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_PolicyAuthorization"
)

func main() {
    pcscfRestorationRequestData := openapiclient.PcscfRestorationRequestData{Interface{}: new(interface{})} // PcscfRestorationRequestData | PCSCF Restoration Indication.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PCSCFRestorationIndicationApi.PcscfRestoration(context.Background()).PcscfRestorationRequestData(pcscfRestorationRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PCSCFRestorationIndicationApi.PcscfRestoration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPcscfRestorationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pcscfRestorationRequestData** | [**PcscfRestorationRequestData**](PcscfRestorationRequestData.md) | PCSCF Restoration Indication. | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

