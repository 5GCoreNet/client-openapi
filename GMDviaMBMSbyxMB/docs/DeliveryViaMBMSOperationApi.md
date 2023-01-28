# \DeliveryViaMBMSOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-xmb/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGMDViaMBMS**](DeliveryViaMBMSOperationApi.md#CreateGMDViaMBMS) | **Post** /{scsAsId}/services/{serviceId}/delivery-via-mbms | Creates a new delivery via MBMS for a given SCS/AS and a service Id.
[**FetchAllGMDViaMBMS**](DeliveryViaMBMSOperationApi.md#FetchAllGMDViaMBMS) | **Get** /{scsAsId}/services/{serviceId}/delivery-via-mbms | Read all group message delivery via MBMS resource for a given SCS/AS and a service id.



## CreateGMDViaMBMS

> GMDViaMBMSByxMB CreateGMDViaMBMS(ctx, scsAsId, serviceId).GMDViaMBMSByxMB(gMDViaMBMSByxMB).Execute()

Creates a new delivery via MBMS for a given SCS/AS and a service Id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/GMDviaMBMSbyxMB"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id
    gMDViaMBMSByxMB := *openapiclient.NewGMDViaMBMSByxMB("NotificationDestination_example") // GMDViaMBMSByxMB | representation of the GMD via MBMS by xMB resource to be Created in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryViaMBMSOperationApi.CreateGMDViaMBMS(context.Background(), scsAsId, serviceId).GMDViaMBMSByxMB(gMDViaMBMSByxMB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryViaMBMSOperationApi.CreateGMDViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGMDViaMBMS`: GMDViaMBMSByxMB
    fmt.Fprintf(os.Stdout, "Response from `DeliveryViaMBMSOperationApi.CreateGMDViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGMDViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gMDViaMBMSByxMB** | [**GMDViaMBMSByxMB**](GMDViaMBMSByxMB.md) | representation of the GMD via MBMS by xMB resource to be Created in the SCEF | 

### Return type

[**GMDViaMBMSByxMB**](GMDViaMBMSByxMB.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllGMDViaMBMS

> []GMDViaMBMSByxMB FetchAllGMDViaMBMS(ctx, scsAsId, serviceId).Execute()

Read all group message delivery via MBMS resource for a given SCS/AS and a service id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/GMDviaMBMSbyxMB"
)

func main() {
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeliveryViaMBMSOperationApi.FetchAllGMDViaMBMS(context.Background(), scsAsId, serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeliveryViaMBMSOperationApi.FetchAllGMDViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllGMDViaMBMS`: []GMDViaMBMSByxMB
    fmt.Fprintf(os.Stdout, "Response from `DeliveryViaMBMSOperationApi.FetchAllGMDViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllGMDViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GMDViaMBMSByxMB**](GMDViaMBMSByxMB.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

