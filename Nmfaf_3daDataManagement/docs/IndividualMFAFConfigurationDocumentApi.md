# \IndividualMFAFConfigurationDocumentApi

All URIs are relative to *https://example.com/nmfaf-3dadatamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMFAFConfiguration**](IndividualMFAFConfigurationDocumentApi.md#DeleteMFAFConfiguration) | **Delete** /configurations/{transRefId} | Deletes an existing Individual MFAF Configuration resource.
[**UpdateMFAFConfiguration**](IndividualMFAFConfigurationDocumentApi.md#UpdateMFAFConfiguration) | **Put** /configurations/{transRefId} | Updates an existing Individual MFAF Configuration resource.



## DeleteMFAFConfiguration

> DeleteMFAFConfiguration(ctx, transRefId).Execute()

Deletes an existing Individual MFAF Configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmfaf_3daDataManagement"
)

func main() {
    transRefId := "transRefId_example" // string | Unique identifier of the individual MFAF Configurations resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMFAFConfigurationDocumentApi.DeleteMFAFConfiguration(context.Background(), transRefId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMFAFConfigurationDocumentApi.DeleteMFAFConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transRefId** | **string** | Unique identifier of the individual MFAF Configurations resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMFAFConfigurationRequest struct via the builder pattern


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


## UpdateMFAFConfiguration

> MfafConfiguration UpdateMFAFConfiguration(ctx, transRefId).MfafConfiguration(mfafConfiguration).Execute()

Updates an existing Individual MFAF Configuration resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmfaf_3daDataManagement"
)

func main() {
    transRefId := "transRefId_example" // string | Unique identifier of the individual MFAF Configurations resource.
    mfafConfiguration := *openapiclient.NewMfafConfiguration([]openapiclient.MessageConfiguration{*openapiclient.NewMessageConfiguration("CorreId_example", "NotificationURI_example")}) // MfafConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMFAFConfigurationDocumentApi.UpdateMFAFConfiguration(context.Background(), transRefId).MfafConfiguration(mfafConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMFAFConfigurationDocumentApi.UpdateMFAFConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMFAFConfiguration`: MfafConfiguration
    fmt.Fprintf(os.Stdout, "Response from `IndividualMFAFConfigurationDocumentApi.UpdateMFAFConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transRefId** | **string** | Unique identifier of the individual MFAF Configurations resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMFAFConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mfafConfiguration** | [**MfafConfiguration**](MfafConfiguration.md) |  | 

### Return type

[**MfafConfiguration**](MfafConfiguration.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

