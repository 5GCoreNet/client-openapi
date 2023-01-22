# \MBSUserServicesApi

All URIs are relative to *https://example.com/3gpp-mbs-us/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSUserService**](MBSUserServicesApi.md#CreateMBSUserService) | **Post** /mbs-user-services | Request the creation of a new Individual MBS User Service resource.
[**RetrieveMBSUserServices**](MBSUserServicesApi.md#RetrieveMBSUserServices) | **Get** /mbs-user-services | Retrieve all the active MBS User Service resources managed by the NEF.



## CreateMBSUserService

> MBSUserService CreateMBSUserService(ctx).MBSUserService(mBSUserService).Execute()

Request the creation of a new Individual MBS User Service resource.

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
    mBSUserService := *openapiclient.NewMBSUserService([]string{"ExtServiceIds_example"}, *openapiclient.NewMbsServiceType(), "ServClass_example", []openapiclient.ServiceAnnouncementMode{openapiclient.ServiceAnnouncementMode{ServiceAnnouncementModeOneOf: penapiclient.ServiceAnnouncementMode_oneOf("VIA_MBS_5")}}) // MBSUserService | Contains the parameters to request the creation of a new MBS User Service at the NEF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserServicesApi.CreateMBSUserService(context.Background()).MBSUserService(mBSUserService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserServicesApi.CreateMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `MBSUserServicesApi.CreateMBSUserService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSUserServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBSUserService** | [**MBSUserService**](MBSUserService.md) | Contains the parameters to request the creation of a new MBS User Service at the NEF.  | 

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


## RetrieveMBSUserServices

> []MBSUserService RetrieveMBSUserServices(ctx).Execute()

Retrieve all the active MBS User Service resources managed by the NEF.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserServicesApi.RetrieveMBSUserServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserServicesApi.RetrieveMBSUserServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMBSUserServices`: []MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `MBSUserServicesApi.RetrieveMBSUserServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMBSUserServicesRequest struct via the builder pattern


### Return type

[**[]MBSUserService**](MBSUserService.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

