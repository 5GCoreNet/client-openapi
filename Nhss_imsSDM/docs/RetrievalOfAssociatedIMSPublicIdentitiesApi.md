# \RetrievalOfAssociatedIMSPublicIdentitiesApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImsAssocIds**](RetrievalOfAssociatedIMSPublicIdentitiesApi.md#GetImsAssocIds) | **Get** /{imsUeId}/identities/ims-associated-identities | Retrieve the associated identities to the IMS public identity included in the service request 



## GetImsAssocIds

> ImsAssociatedIdentities GetImsAssocIds(ctx, imsUeId).Execute()

Retrieve the associated identities to the IMS public identity included in the service request 

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
    imsUeId := "imsUeId_example" // string | IMS Public Identity

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RetrievalOfAssociatedIMSPublicIdentitiesApi.GetImsAssocIds(context.Background(), imsUeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RetrievalOfAssociatedIMSPublicIdentitiesApi.GetImsAssocIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImsAssocIds`: ImsAssociatedIdentities
    fmt.Fprintf(os.Stdout, "Response from `RetrievalOfAssociatedIMSPublicIdentitiesApi.GetImsAssocIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImsAssocIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImsAssociatedIdentities**](ImsAssociatedIdentities.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

