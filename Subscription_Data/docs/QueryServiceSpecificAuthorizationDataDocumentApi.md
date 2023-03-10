# \QueryServiceSpecificAuthorizationDataDocumentApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSSAuData**](QueryServiceSpecificAuthorizationDataDocumentApi.md#GetSSAuData) | **Get** /subscription-data/{ueId}/service-specific-authorization-data/{serviceType} | Retrieve ServiceSpecific Authorization Data



## GetSSAuData

> AuthorizationData GetSSAuData(ctx, ueId, serviceType).SingleNssai(singleNssai).Dnn(dnn).MtcProviderInformation(mtcProviderInformation).AfId(afId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()

Retrieve ServiceSpecific Authorization Data

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
    ueId := "ueId_example" // string | UE ID
    serviceType := *openapiclient.NewServiceType() // ServiceType | Service Type
    singleNssai := map[string][]openapiclient.VarSnssai{ ... } // VarSnssai | single NSSAI
    dnn := "dnn_example" // string | DNN
    mtcProviderInformation := "mtcProviderInformation_example" // string | MTC Provider Information (optional)
    afId := "afId_example" // string | Application Function Identifier (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryServiceSpecificAuthorizationDataDocumentApi.GetSSAuData(context.Background(), ueId, serviceType).SingleNssai(singleNssai).Dnn(dnn).MtcProviderInformation(mtcProviderInformation).AfId(afId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryServiceSpecificAuthorizationDataDocumentApi.GetSSAuData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSAuData`: AuthorizationData
    fmt.Fprintf(os.Stdout, "Response from `QueryServiceSpecificAuthorizationDataDocumentApi.GetSSAuData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueId** | **string** | UE ID | 
**serviceType** | [**ServiceType**](.md) | Service Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSAuDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **singleNssai** | [**VarSnssai**](VarSnssai.md) | single NSSAI | 
 **dnn** | **string** | DNN | 
 **mtcProviderInformation** | **string** | MTC Provider Information | 
 **afId** | **string** | Application Function Identifier | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 

### Return type

[**AuthorizationData**](AuthorizationData.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

