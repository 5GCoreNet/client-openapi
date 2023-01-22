# \DCCFDataCollectionProfilesCollectionApi

All URIs are relative to *https://example.com/ndccf-contextmanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDCCFDataCollectionProfile**](DCCFDataCollectionProfilesCollectionApi.md#CreateDCCFDataCollectionProfile) | **Post** /data-collection-profiles | Creates a new Individual DCCF Data Collection Profile resource.



## CreateDCCFDataCollectionProfile

> NdccfDataCollectionProfile CreateDCCFDataCollectionProfile(ctx).NdccfDataCollectionProfile(ndccfDataCollectionProfile).Execute()

Creates a new Individual DCCF Data Collection Profile resource.

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
    ndccfDataCollectionProfile := *openapiclient.NewNdccfDataCollectionProfile() // NdccfDataCollectionProfile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DCCFDataCollectionProfilesCollectionApi.CreateDCCFDataCollectionProfile(context.Background()).NdccfDataCollectionProfile(ndccfDataCollectionProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DCCFDataCollectionProfilesCollectionApi.CreateDCCFDataCollectionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDCCFDataCollectionProfile`: NdccfDataCollectionProfile
    fmt.Fprintf(os.Stdout, "Response from `DCCFDataCollectionProfilesCollectionApi.CreateDCCFDataCollectionProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDCCFDataCollectionProfileRequest struct via the builder pattern


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

