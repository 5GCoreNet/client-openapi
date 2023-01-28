# \ProvidingTheReceptionStatusOfTheAcknowledgementOfSteeringOfRoamingInformationReceptionByTheUEApi

All URIs are relative to *https://example.com/nsoraf-sor/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SorAckInfo**](ProvidingTheReceptionStatusOfTheAcknowledgementOfSteeringOfRoamingInformationReceptionByTheUEApi.md#SorAckInfo) | **Put** /{supi}/sor-information/sor-ack | SoR Acknowledgment Reception Notification



## SorAckInfo

> SorAckInfo(ctx, supi).SorAckInfo(sorAckInfo).Execute()

SoR Acknowledgment Reception Notification

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsoraf_SOR"
)

func main() {
    supi := "supi_example" // string | Identifier of the UE
    sorAckInfo := *openapiclient.NewSorAckInfo(*openapiclient.NewSorAckStatus(), time.Now()) // SorAckInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvidingTheReceptionStatusOfTheAcknowledgementOfSteeringOfRoamingInformationReceptionByTheUEApi.SorAckInfo(context.Background(), supi).SorAckInfo(sorAckInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvidingTheReceptionStatusOfTheAcknowledgementOfSteeringOfRoamingInformationReceptionByTheUEApi.SorAckInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**supi** | **string** | Identifier of the UE | 

### Other Parameters

Other parameters are passed through a pointer to a apiSorAckInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sorAckInfo** | [**SorAckInfo**](SorAckInfo.md) |  | 

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

