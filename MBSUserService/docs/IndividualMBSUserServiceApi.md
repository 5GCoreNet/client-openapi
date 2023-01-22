# \IndividualMBSUserServiceApi

All URIs are relative to *https://example.com/3gpp-mbs-us/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndivMBSUserService**](IndividualMBSUserServiceApi.md#DeleteIndivMBSUserService) | **Delete** /mbs-user-services/{mbsUserServId} | Deletes an existing Individual MBS User Service resource.
[**ModifyIndivMBSUserService**](IndividualMBSUserServiceApi.md#ModifyIndivMBSUserService) | **Patch** /mbs-user-services/{mbsUserServId} | Request the modification of an existing Individual MBS User Service resource.
[**RetrieveIndivMBSUserService**](IndividualMBSUserServiceApi.md#RetrieveIndivMBSUserService) | **Get** /mbs-user-services/{mbsUserServId} | Retrieve an existing Individual MBS User Service resource.
[**UpdateIndivMBSUserService**](IndividualMBSUserServiceApi.md#UpdateIndivMBSUserService) | **Put** /mbs-user-services/{mbsUserServId} | Request the update of an existing Individual MBS User Service resource.



## DeleteIndivMBSUserService

> DeleteIndivMBSUserService(ctx, mbsUserServId).Execute()

Deletes an existing Individual MBS User Service resource.

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
    resp, r, err := apiClient.IndividualMBSUserServiceApi.DeleteIndivMBSUserService(context.Background(), mbsUserServId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceApi.DeleteIndivMBSUserService``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndivMBSUserServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndivMBSUserService

> MBSUserService ModifyIndivMBSUserService(ctx, mbsUserServId).MBSUserServicePatch(mBSUserServicePatch).Execute()

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
    resp, r, err := apiClient.IndividualMBSUserServiceApi.ModifyIndivMBSUserService(context.Background(), mbsUserServId).MBSUserServicePatch(mBSUserServicePatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceApi.ModifyIndivMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndivMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserServiceApi.ModifyIndivMBSUserService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndivMBSUserServiceRequest struct via the builder pattern


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


## RetrieveIndivMBSUserService

> MBSUserService RetrieveIndivMBSUserService(ctx, mbsUserServId).Execute()

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
    resp, r, err := apiClient.IndividualMBSUserServiceApi.RetrieveIndivMBSUserService(context.Background(), mbsUserServId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceApi.RetrieveIndivMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndivMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserServiceApi.RetrieveIndivMBSUserService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndivMBSUserServiceRequest struct via the builder pattern


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


## UpdateIndivMBSUserService

> MBSUserService UpdateIndivMBSUserService(ctx, mbsUserServId).MBSUserService(mBSUserService).Execute()

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
    resp, r, err := apiClient.IndividualMBSUserServiceApi.UpdateIndivMBSUserService(context.Background(), mbsUserServId).MBSUserService(mBSUserService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserServiceApi.UpdateIndivMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndivMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserServiceApi.UpdateIndivMBSUserService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsUserServId** | **string** | Identifier of the Individual MBS User Service resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndivMBSUserServiceRequest struct via the builder pattern


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

