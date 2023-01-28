# \UpdateTheRoamingInformationOfTheEPCDomainDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateRoamingInformation**](UpdateTheRoamingInformationOfTheEPCDomainDocumentApi.md#UpdateRoamingInformation) | **Put** /subscription-data/{ueId}/context-data/roaming-information | Update the Roaming Information of the EPC domain



## UpdateRoamingInformation

> RoamingInfoUpdate UpdateRoamingInformation(ctx, ueId).RoamingInfoUpdate(roamingInfoUpdate).Execute()

Update the Roaming Information of the EPC domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Subscription_Data"
)

func main() {
    ueId := "ueId_example" // string | UE id
    roamingInfoUpdate := *openapiclient.NewRoamingInfoUpdate(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")) // RoamingInfoUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpdateTheRoamingInformationOfTheEPCDomainDocumentApi.UpdateRoamingInformation(context.Background(), ueId).RoamingInfoUpdate(roamingInfoUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpdateTheRoamingInformationOfTheEPCDomainDocumentApi.UpdateRoamingInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRoamingInformation`: RoamingInfoUpdate
    fmt.Fprintf(os.Stdout, "Response from `UpdateTheRoamingInformationOfTheEPCDomainDocumentApi.UpdateRoamingInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoamingInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roamingInfoUpdate** | [**RoamingInfoUpdate**](RoamingInfoUpdate.md) |  | 

### Return type

[**RoamingInfoUpdate**](RoamingInfoUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

