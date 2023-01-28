# \IndividualMBSDistributionSessionApi

All URIs are relative to *https://example.com/nmbstf-distsession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Destroy**](IndividualMBSDistributionSessionApi.md#Destroy) | **Delete** /dist-sessions/{distSessionRef} | Deletes an individual MBS distribution session resource in the MBSTF.
[**Retrieve**](IndividualMBSDistributionSessionApi.md#Retrieve) | **Get** /dist-sessions/{distSessionRef} | Retrieves an individual MBS distribution session resource in the MBSTF.
[**Update**](IndividualMBSDistributionSessionApi.md#Update) | **Patch** /dist-sessions/{distSessionRef} | Updates an individual MBS distribution session resource in the MBSTF.



## Destroy

> Destroy(ctx, distSessionRef).Execute()

Deletes an individual MBS distribution session resource in the MBSTF.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbstf_DistSession"
)

func main() {
    distSessionRef := "distSessionRef_example" // string | Unique ID of the MBS distribution session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSDistributionSessionApi.Destroy(context.Background(), distSessionRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSDistributionSessionApi.Destroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distSessionRef** | **string** | Unique ID of the MBS distribution session | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyRequest struct via the builder pattern


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


## Retrieve

> DistSession Retrieve(ctx, distSessionRef).Execute()

Retrieves an individual MBS distribution session resource in the MBSTF.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbstf_DistSession"
)

func main() {
    distSessionRef := "distSessionRef_example" // string | Unique ID of the MBS distribution session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSDistributionSessionApi.Retrieve(context.Background(), distSessionRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSDistributionSessionApi.Retrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Retrieve`: DistSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSDistributionSessionApi.Retrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distSessionRef** | **string** | Unique ID of the MBS distribution session | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DistSession**](DistSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> DistSession Update(ctx, distSessionRef).PatchItem(patchItem).Execute()

Updates an individual MBS distribution session resource in the MBSTF.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbstf_DistSession"
)

func main() {
    distSessionRef := "distSessionRef_example" // string | Unique ID of the MBS distribution session
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Data within the Update Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSDistributionSessionApi.Update(context.Background(), distSessionRef).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSDistributionSessionApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: DistSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSDistributionSessionApi.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distSessionRef** | **string** | Unique ID of the MBS distribution session | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) | Data within the Update Request | 

### Return type

[**DistSession**](DistSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

