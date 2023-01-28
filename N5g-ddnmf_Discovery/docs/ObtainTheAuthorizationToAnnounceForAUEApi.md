# \ObtainTheAuthorizationToAnnounceForAUEApi

All URIs are relative to *https://example.com/n5g-ddnmf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObtainAnnounceAuth**](ObtainTheAuthorizationToAnnounceForAUEApi.md#ObtainAnnounceAuth) | **Put** /{ueId}/announce-authorize/{discEntryId} | Obtain the authorization to announce for a UE



## ObtainAnnounceAuth

> ObtainAnnounceAuth(ctx, ueId, discEntryId).AnnounceAuthReqData(announceAuthReqData).Execute()

Obtain the authorization to announce for a UE

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/N5g-ddnmf_Discovery"
)

func main() {
    ueId := "ueId_example" // string | Identifier of the UE
    discEntryId := "discEntryId_example" // string | Discovery Entry Id
    announceAuthReqData := *openapiclient.NewAnnounceAuthReqData(*openapiclient.NewDiscoveryType()) // AnnounceAuthReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObtainTheAuthorizationToAnnounceForAUEApi.ObtainAnnounceAuth(context.Background(), ueId, discEntryId).AnnounceAuthReqData(announceAuthReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObtainTheAuthorizationToAnnounceForAUEApi.ObtainAnnounceAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 
**discEntryId** | **string** | Discovery Entry Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiObtainAnnounceAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **announceAuthReqData** | [**AnnounceAuthReqData**](AnnounceAuthReqData.md) |  | 

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

