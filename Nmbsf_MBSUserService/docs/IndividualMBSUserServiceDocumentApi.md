# \IndividualMBSUserServiceDocumentApi

All URIs are relative to *https://example.com/nmbsf-mbs-us/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSUserService**](IndividualMBSUserServiceDocumentApi.md#DeleteIndMBSUserService) | **Delete** /mbs-user-services/{mbsUserServId} | Request the deletion of an existing Individual MBS User Service resource.
[**ModifyIndMBSUserService**](IndividualMBSUserServiceDocumentApi.md#ModifyIndMBSUserService) | **Patch** /mbs-user-services/{mbsUserServId} | Request the modification of an existing Individual MBS User Service resource.
[**RetrieveIndMBSUserService**](IndividualMBSUserServiceDocumentApi.md#RetrieveIndMBSUserService) | **Get** /mbs-user-services/{mbsUserServId} | Retrieve an existing Individual MBS User Service resource.
[**UpdateIndMBSUserService**](IndividualMBSUserServiceDocumentApi.md#UpdateIndMBSUserService) | **Put** /mbs-user-services/{mbsUserServId} | Request the update of an existing Individual MBS User Service resource.



## DeleteIndMBSUserService

> DeleteIndMBSUserService(ctx, mbsUserServId).Execute()

Request the deletion of an existing Individual MBS User Service resource.

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
    mbsUserServId := "mbsUserServId_example" // string | Identifier of the Individual MBS User Service resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserServiceDocumentApi.DeleteIndMBSUserService(context.Background(), mbsUserServId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceDocumentApi.DeleteIndMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMBSUserServiceRequest struct via the builder pattern


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


## ModifyIndMBSUserService

> MBSUserService ModifyIndMBSUserService(ctx, mbsUserServId).MBSUserServicePatch(mBSUserServicePatch).Execute()

Request the modification of an existing Individual MBS User Service resource.

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
    mbsUserServId := "mbsUserServId_example" // string | Identifier of the Individual MBS User Service resource.
    mBSUserServicePatch := *openapiclient.NewMBSUserServicePatch() // MBSUserServicePatch | Contains the parameters to request the modification of the Individual MBS User Service resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserServiceDocumentApi.ModifyIndMBSUserService(context.Background(), mbsUserServId).MBSUserServicePatch(mBSUserServicePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceDocumentApi.ModifyIndMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserServiceDocumentApi.ModifyIndMBSUserService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndMBSUserServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mBSUserServicePatch** | [**MBSUserServicePatch**](MBSUserServicePatch.md) | Contains the parameters to request the modification of the Individual MBS User Service resource.  | 

### Return type

[**MBSUserService**](MBSUserService.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndMBSUserService

> MBSUserService RetrieveIndMBSUserService(ctx, mbsUserServId).Execute()

Retrieve an existing Individual MBS User Service resource.

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
    mbsUserServId := "mbsUserServId_example" // string | Identifier of the Individual MBS User Service resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserServiceDocumentApi.RetrieveIndMBSUserService(context.Background(), mbsUserServId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceDocumentApi.RetrieveIndMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserServiceDocumentApi.RetrieveIndMBSUserService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndMBSUserServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MBSUserService**](MBSUserService.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndMBSUserService

> MBSUserService UpdateIndMBSUserService(ctx, mbsUserServId).MBSUserService(mBSUserService).Execute()

Request the update of an existing Individual MBS User Service resource.

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
    mbsUserServId := "mbsUserServId_example" // string | Identifier of the Individual MBS User Service resource.
    mBSUserService := *openapiclient.NewMBSUserService([]string{"ExtServiceIds_example"}, *openapiclient.NewMbsServiceType(), "ServClass_example", []openapiclient.ServiceAnnouncementMode{openapiclient.ServiceAnnouncementMode{ServiceAnnouncementModeOneOf: penapiclient.ServiceAnnouncementMode_oneOf("VIA_MBS_5")}}) // MBSUserService | Contains the updated representation of the Individual MBS User Service resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserServiceDocumentApi.UpdateIndMBSUserService(context.Background(), mbsUserServId).MBSUserService(mBSUserService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceDocumentApi.UpdateIndMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserServiceDocumentApi.UpdateIndMBSUserService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndMBSUserServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mBSUserService** | [**MBSUserService**](MBSUserService.md) | Contains the updated representation of the Individual MBS User Service resource.  | 

### Return type

[**MBSUserService**](MBSUserService.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

