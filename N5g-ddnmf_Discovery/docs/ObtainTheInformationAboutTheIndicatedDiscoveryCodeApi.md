# \ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi

All URIs are relative to *https://example.com/n5g-ddnmf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MatchReport**](ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi.md#MatchReport) | **Post** /{ueId}/match-report | Obtain the information about the indicated discovery code from the 5G DDNMF



## MatchReport

> MatchReportRespData MatchReport(ctx, ueId).MatchReportReqData(matchReportReqData).Execute()

Obtain the information about the indicated discovery code from the 5G DDNMF

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
    ueId := "ueId_example" // string | Identifier of the UE
    matchReportReqData := *openapiclient.NewMatchReportReqData(*openapiclient.NewDiscoveryType()) // MatchReportReqData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi.MatchReport(context.Background(), ueId).MatchReportReqData(matchReportReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi.MatchReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MatchReport`: MatchReportRespData
    fmt.Fprintf(os.Stdout, "Response from `ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi.MatchReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiMatchReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matchReportReqData** | [**MatchReportReqData**](MatchReportReqData.md) |  | 

### Return type

[**MatchReportRespData**](MatchReportRespData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

