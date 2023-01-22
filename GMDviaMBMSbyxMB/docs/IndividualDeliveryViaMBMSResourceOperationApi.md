# \IndividualDeliveryViaMBMSResourceOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-xmb/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndGMDViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#DeleteIndGMDViaMBMS) | **Delete** /{scsAsId}/services/{serviceId}/delivery-via-mbms/{transactionId} | Deletes a delivery via MBMS resource for a given SCS/AS, a service Id and a transcation Id.
[**FetchIndGMDViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#FetchIndGMDViaMBMS) | **Get** /{scsAsId}/services/{serviceId}/delivery-via-mbms/{transactionId} | Read all group message delivery via MBMS resource for a given SCS/AS and a service Id.
[**ModifyIndGMDViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#ModifyIndGMDViaMBMS) | **Patch** /{scsAsId}/services/{serviceId}/delivery-via-mbms/{transactionId} | Updates an existing delivery via MBMS for a given SCS/AS, a service Id and transaction Id.
[**UpdateIndGMDViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#UpdateIndGMDViaMBMS) | **Put** /{scsAsId}/services/{serviceId}/delivery-via-mbms/{transactionId} | Updates an existing delivery via MBMS for a given SCS/AS, a service Id and transaction Id.



## DeleteIndGMDViaMBMS

> DeleteIndGMDViaMBMS(ctx, scsAsId, serviceId, transactionId).Execute()

Deletes a delivery via MBMS resource for a given SCS/AS, a service Id and a transcation Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id
    transactionId := "transactionId_example" // string | Identifier of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.DeleteIndGMDViaMBMS(context.Background(), scsAsId, serviceId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.DeleteIndGMDViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndGMDViaMBMSRequest struct via the builder pattern


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


## FetchIndGMDViaMBMS

> GMDViaMBMSByxMB FetchIndGMDViaMBMS(ctx, scsAsId, serviceId, transactionId).Execute()

Read all group message delivery via MBMS resource for a given SCS/AS and a service Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id
    transactionId := "transactionId_example" // string | Identifier of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.FetchIndGMDViaMBMS(context.Background(), scsAsId, serviceId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.FetchIndGMDViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndGMDViaMBMS`: GMDViaMBMSByxMB
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeliveryViaMBMSResourceOperationApi.FetchIndGMDViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndGMDViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GMDViaMBMSByxMB**](GMDViaMBMSByxMB.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndGMDViaMBMS

> GMDViaMBMSByxMB ModifyIndGMDViaMBMS(ctx, scsAsId, serviceId, transactionId).GMDViaMBMSByxMBPatch(gMDViaMBMSByxMBPatch).Execute()

Updates an existing delivery via MBMS for a given SCS/AS, a service Id and transaction Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id
    transactionId := "transactionId_example" // string | Identifier of transaction
    gMDViaMBMSByxMBPatch := *openapiclient.NewGMDViaMBMSByxMBPatch() // GMDViaMBMSByxMBPatch | representation of the GMD via MBMS by xMB resource to be udpated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.ModifyIndGMDViaMBMS(context.Background(), scsAsId, serviceId, transactionId).GMDViaMBMSByxMBPatch(gMDViaMBMSByxMBPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.ModifyIndGMDViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndGMDViaMBMS`: GMDViaMBMSByxMB
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeliveryViaMBMSResourceOperationApi.ModifyIndGMDViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndGMDViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gMDViaMBMSByxMBPatch** | [**GMDViaMBMSByxMBPatch**](GMDViaMBMSByxMBPatch.md) | representation of the GMD via MBMS by xMB resource to be udpated in the SCEF | 

### Return type

[**GMDViaMBMSByxMB**](GMDViaMBMSByxMB.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndGMDViaMBMS

> GMDViaMBMSByxMB UpdateIndGMDViaMBMS(ctx, scsAsId, serviceId, transactionId).GMDViaMBMSByxMB(gMDViaMBMSByxMB).Execute()

Updates an existing delivery via MBMS for a given SCS/AS, a service Id and transaction Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    serviceId := "serviceId_example" // string | Service Id
    transactionId := "transactionId_example" // string | Identifier of transaction
    gMDViaMBMSByxMB := *openapiclient.NewGMDViaMBMSByxMB("NotificationDestination_example") // GMDViaMBMSByxMB | representation of the GMD via MBMS by xMB resource to be udpated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.UpdateIndGMDViaMBMS(context.Background(), scsAsId, serviceId, transactionId).GMDViaMBMSByxMB(gMDViaMBMSByxMB).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.UpdateIndGMDViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndGMDViaMBMS`: GMDViaMBMSByxMB
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeliveryViaMBMSResourceOperationApi.UpdateIndGMDViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**serviceId** | **string** | Service Id | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndGMDViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gMDViaMBMSByxMB** | [**GMDViaMBMSByxMB**](GMDViaMBMSByxMB.md) | representation of the GMD via MBMS by xMB resource to be udpated in the SCEF | 

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

