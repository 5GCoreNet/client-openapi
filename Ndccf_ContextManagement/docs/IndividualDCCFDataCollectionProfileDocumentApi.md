# \IndividualDCCFDataCollectionProfileDocumentApi

All URIs are relative to *https://example.com/ndccf-contextmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDCCFDataCollectionProfile**](IndividualDCCFDataCollectionProfileDocumentApi.md#DeleteDCCFDataCollectionProfile) | **Delete** /data-collection-profiles/{profileId} | Deletes an existing Individual DCCF Data Subscription resource.
[**UpdateDCCFDataCollectionProfile**](IndividualDCCFDataCollectionProfileDocumentApi.md#UpdateDCCFDataCollectionProfile) | **Put** /data-collection-profiles/{profileId} | Updates an existing Individual DCCF Data Collection Profile resource.



## DeleteDCCFDataCollectionProfile

> DeleteDCCFDataCollectionProfile(ctx, profileId).Execute()

Deletes an existing Individual DCCF Data Subscription resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndccf_ContextManagement"
)

func main() {
    profileId := "profileId_example" // string | String identifying a data collection profile at the Ndccf_ContextManagement Service. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDCCFDataCollectionProfileDocumentApi.DeleteDCCFDataCollectionProfile(context.Background(), profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDCCFDataCollectionProfileDocumentApi.DeleteDCCFDataCollectionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | String identifying a data collection profile at the Ndccf_ContextManagement Service.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDCCFDataCollectionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDCCFDataCollectionProfile

> NdccfDataCollectionProfile UpdateDCCFDataCollectionProfile(ctx, profileId).NdccfDataCollectionProfile(ndccfDataCollectionProfile).Execute()

Updates an existing Individual DCCF Data Collection Profile resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Ndccf_ContextManagement"
)

func main() {
    profileId := "profileId_example" // string | String identifying a data collection profile at the Ndccf_ContextManagement Service. 
    ndccfDataCollectionProfile := *openapiclient.NewNdccfDataCollectionProfile() // NdccfDataCollectionProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDCCFDataCollectionProfileDocumentApi.UpdateDCCFDataCollectionProfile(context.Background(), profileId).NdccfDataCollectionProfile(ndccfDataCollectionProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDCCFDataCollectionProfileDocumentApi.UpdateDCCFDataCollectionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDCCFDataCollectionProfile`: NdccfDataCollectionProfile
    fmt.Fprintf(os.Stdout, "Response from `IndividualDCCFDataCollectionProfileDocumentApi.UpdateDCCFDataCollectionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | String identifying a data collection profile at the Ndccf_ContextManagement Service.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDCCFDataCollectionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ndccfDataCollectionProfile** | [**NdccfDataCollectionProfile**](NdccfDataCollectionProfile.md) |  | 

### Return type

[**NdccfDataCollectionProfile**](NdccfDataCollectionProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

