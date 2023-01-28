# \MFAFConfigurationCollectionApi

All URIs are relative to *https://example.com/nmfaf-3dadatamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMFAFConfiguration**](MFAFConfigurationCollectionApi.md#CreateMFAFConfiguration) | **Post** /configurations | Creates a new Individual MFAF Configuration resource.



## CreateMFAFConfiguration

> MfafConfiguration CreateMFAFConfiguration(ctx).MfafConfiguration(mfafConfiguration).Execute()

Creates a new Individual MFAF Configuration resource.

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
    mfafConfiguration := *openapiclient.NewMfafConfiguration([]openapiclient.MessageConfiguration{*openapiclient.NewMessageConfiguration("CorreId_example", "NotificationURI_example")}) // MfafConfiguration | Contains the information for the creation of a new Individual MFAF Configuration resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFAFConfigurationCollectionApi.CreateMFAFConfiguration(context.Background()).MfafConfiguration(mfafConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFAFConfigurationCollectionApi.CreateMFAFConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMFAFConfiguration`: MfafConfiguration
    fmt.Fprintf(os.Stdout, "Response from `MFAFConfigurationCollectionApi.CreateMFAFConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMFAFConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfafConfiguration** | [**MfafConfiguration**](MfafConfiguration.md) | Contains the information for the creation of a new Individual MFAF Configuration resource.  | 

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

