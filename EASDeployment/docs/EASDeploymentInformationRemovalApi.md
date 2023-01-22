# \EASDeploymentInformationRemovalApi

All URIs are relative to *https://example.com/3gpp-eas-deployment/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEDIs**](EASDeploymentInformationRemovalApi.md#DeleteEDIs) | **Post** /remove-edis | Remove EAS Deployment Information based on given criteria.



## DeleteEDIs

> DeleteEDIs(ctx).EdiDeleteCriteria(ediDeleteCriteria).Execute()

Remove EAS Deployment Information based on given criteria.

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
    ediDeleteCriteria := *openapiclient.NewEdiDeleteCriteria() // EdiDeleteCriteria | Criteria to be used for deleting EAS Deployment Information that match them.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EASDeploymentInformationRemovalApi.DeleteEDIs(context.Background()).EdiDeleteCriteria(ediDeleteCriteria).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EASDeploymentInformationRemovalApi.DeleteEDIs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEDIsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ediDeleteCriteria** | [**EdiDeleteCriteria**](EdiDeleteCriteria.md) | Criteria to be used for deleting EAS Deployment Information that match them. | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

