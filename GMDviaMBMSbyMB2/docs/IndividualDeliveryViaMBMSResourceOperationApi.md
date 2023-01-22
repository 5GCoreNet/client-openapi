# \IndividualDeliveryViaMBMSResourceOperationApi

All URIs are relative to *https://example.com/3gpp-group-message-delivery-mb2/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndDeliveryViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#DeleteIndDeliveryViaMBMS) | **Delete** /{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId} | Deletes a delivery via MBMS resource for a given SCS/AS, a TMGI and a transcation Id.
[**FetchIndDeliveryViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#FetchIndDeliveryViaMBMS) | **Get** /{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId} | Read all group message delivery via MBMS resource for a given SCS/AS and a TMGI.
[**ModifyIndDeliveryViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#ModifyIndDeliveryViaMBMS) | **Patch** /{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId} | Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.
[**UpdateIndDeliveryViaMBMS**](IndividualDeliveryViaMBMSResourceOperationApi.md#UpdateIndDeliveryViaMBMS) | **Put** /{scsAsId}/tmgi-allocation/{tmgi}/delivery-via-mbms/{transactionId} | Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.



## DeleteIndDeliveryViaMBMS

> DeleteIndDeliveryViaMBMS(ctx, scsAsId, tmgi, transactionId).Execute()

Deletes a delivery via MBMS resource for a given SCS/AS, a TMGI and a transcation Id.

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
    tmgi := "tmgi_example" // string | TMGI
    transactionId := "transactionId_example" // string | Identifier of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.DeleteIndDeliveryViaMBMS(context.Background(), scsAsId, tmgi, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.DeleteIndDeliveryViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndDeliveryViaMBMSRequest struct via the builder pattern


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


## FetchIndDeliveryViaMBMS

> GMDViaMBMSByMb2 FetchIndDeliveryViaMBMS(ctx, scsAsId, tmgi, transactionId).Execute()

Read all group message delivery via MBMS resource for a given SCS/AS and a TMGI.

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
    tmgi := "tmgi_example" // string | TMGI
    transactionId := "transactionId_example" // string | Identifier of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.FetchIndDeliveryViaMBMS(context.Background(), scsAsId, tmgi, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.FetchIndDeliveryViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndDeliveryViaMBMS`: GMDViaMBMSByMb2
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeliveryViaMBMSResourceOperationApi.FetchIndDeliveryViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndDeliveryViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndDeliveryViaMBMS

> GMDViaMBMSByMb2 ModifyIndDeliveryViaMBMS(ctx, scsAsId, tmgi, transactionId).GMDViaMBMSByMb2Patch(gMDViaMBMSByMb2Patch).Execute()

Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.

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
    tmgi := "tmgi_example" // string | TMGI
    transactionId := "transactionId_example" // string | Identifier of transaction
    gMDViaMBMSByMb2Patch := *openapiclient.NewGMDViaMBMSByMb2Patch() // GMDViaMBMSByMb2Patch | representation of the GMD via MBMS by MB2 resource to be udpated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.ModifyIndDeliveryViaMBMS(context.Background(), scsAsId, tmgi, transactionId).GMDViaMBMSByMb2Patch(gMDViaMBMSByMb2Patch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.ModifyIndDeliveryViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndDeliveryViaMBMS`: GMDViaMBMSByMb2
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeliveryViaMBMSResourceOperationApi.ModifyIndDeliveryViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndDeliveryViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gMDViaMBMSByMb2Patch** | [**GMDViaMBMSByMb2Patch**](GMDViaMBMSByMb2Patch.md) | representation of the GMD via MBMS by MB2 resource to be udpated in the SCEF | 

### Return type

[**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndDeliveryViaMBMS

> GMDViaMBMSByMb2 UpdateIndDeliveryViaMBMS(ctx, scsAsId, tmgi, transactionId).GMDViaMBMSByMb2(gMDViaMBMSByMb2).Execute()

Updates a existing delivery via MBMS for a given SCS/AS, a TMGI and transaction Id.

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
    tmgi := "tmgi_example" // string | TMGI
    transactionId := "transactionId_example" // string | Identifier of transaction
    gMDViaMBMSByMb2 := *openapiclient.NewGMDViaMBMSByMb2("NotificationDestination_example") // GMDViaMBMSByMb2 | representation of the GMD via MBMS by MB2 resource to be udpated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualDeliveryViaMBMSResourceOperationApi.UpdateIndDeliveryViaMBMS(context.Background(), scsAsId, tmgi, transactionId).GMDViaMBMSByMb2(gMDViaMBMSByMb2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualDeliveryViaMBMSResourceOperationApi.UpdateIndDeliveryViaMBMS``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndDeliveryViaMBMS`: GMDViaMBMSByMb2
    fmt.Fprintf(os.Stdout, "Response from `IndividualDeliveryViaMBMSResourceOperationApi.UpdateIndDeliveryViaMBMS`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**tmgi** | **string** | TMGI | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndDeliveryViaMBMSRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **gMDViaMBMSByMb2** | [**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md) | representation of the GMD via MBMS by MB2 resource to be udpated in the SCEF | 

### Return type

[**GMDViaMBMSByMb2**](GMDViaMBMSByMb2.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

