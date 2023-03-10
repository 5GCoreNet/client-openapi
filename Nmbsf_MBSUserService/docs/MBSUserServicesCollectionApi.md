# \MBSUserServicesCollectionApi

All URIs are relative to *https://example.com/nmbsf-mbs-us/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSUserService**](MBSUserServicesCollectionApi.md#CreateMBSUserService) | **Post** /mbs-user-services | Request the creation of a new MBS User Service.
[**RetrieveMBSUserServices**](MBSUserServicesCollectionApi.md#RetrieveMBSUserServices) | **Get** /mbs-user-services | Retrieve all the active MBS User Service resources managed by the MBSF.



## CreateMBSUserService

> MBSUserService CreateMBSUserService(ctx).MBSUserService(mBSUserService).Execute()

Request the creation of a new MBS User Service.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsf_MBSUserService"
)

func main() {
    mBSUserService := *openapiclient.NewMBSUserService([]string{"ExtServiceIds_example"}, *openapiclient.NewMbsServiceType(), "ServClass_example", []openapiclient.ServiceAnnouncementMode{openapiclient.ServiceAnnouncementMode{ServiceAnnouncementModeOneOf: penapiclient.ServiceAnnouncementMode_oneOf("VIA_MBS_5")}}) // MBSUserService | Contains the parameters to request the creation of a new MBS User Service at the MBSF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserServicesCollectionApi.CreateMBSUserService(context.Background()).MBSUserService(mBSUserService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserServicesCollectionApi.CreateMBSUserService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSUserService`: MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `MBSUserServicesCollectionApi.CreateMBSUserService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSUserServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBSUserService** | [**MBSUserService**](MBSUserService.md) | Contains the parameters to request the creation of a new MBS User Service at the MBSF.  | 

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

Retrieve all the active MBS User Service resources managed by the MBSF.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsf_MBSUserService"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserServicesCollectionApi.RetrieveMBSUserServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserServicesCollectionApi.RetrieveMBSUserServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMBSUserServices`: []MBSUserService
    fmt.Fprintf(os.Stdout, "Response from `MBSUserServicesCollectionApi.RetrieveMBSUserServices`: %v\n", resp)
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

