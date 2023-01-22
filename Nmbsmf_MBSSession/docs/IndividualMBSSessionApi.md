# \IndividualMBSSessionApi

All URIs are relative to *https://example.com/nmbsmf-mbssession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Release**](IndividualMBSSessionApi.md#Release) | **Delete** /mbs-sessions/{mbsSessionRef} | Deletes an individual MBS session resource in the MB-SMF.
[**Update**](IndividualMBSSessionApi.md#Update) | **Patch** /mbs-sessions/{mbsSessionRef} | Updates an individual MBS session resource in the MB-SMF.



## Release

> Release(ctx, mbsSessionRef).Execute()

Deletes an individual MBS session resource in the MB-SMF.

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
    mbsSessionRef := "mbsSessionRef_example" // string | Unique ID of the MBS session to be released

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSSessionApi.Release(context.Background(), mbsSessionRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSSessionApi.Release``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsSessionRef** | **string** | Unique ID of the MBS session to be released | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseRequest struct via the builder pattern


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


## Update

> Update(ctx, mbsSessionRef).PatchItem(patchItem).Execute()

Updates an individual MBS session resource in the MB-SMF.

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
    mbsSessionRef := "mbsSessionRef_example" // string | Unique ID of the MBS session to be modified
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Data within the Update Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSSessionApi.Update(context.Background(), mbsSessionRef).PatchItem(patchItem).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSSessionApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsSessionRef** | **string** | Unique ID of the MBS session to be modified | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchItem** | [**[]PatchItem**](PatchItem.md) | Data within the Update Request | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

