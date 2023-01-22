# \IndividualSMContextDocumentApi

All URIs are relative to *https://example.com/nnef-smcontext/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](IndividualSMContextDocumentApi.md#Delete) | **Post** /sm-contexts/{smContextId}/release | Delete SM Context
[**Deliver**](IndividualSMContextDocumentApi.md#Deliver) | **Post** /sm-contexts/{smContextId}/deliver | Deliver Uplink MO Data
[**Update**](IndividualSMContextDocumentApi.md#Update) | **Post** /sm-contexts/{smContextId}/update | Update SM Context



## Delete

> SmContextReleasedData Delete(ctx, smContextId).SmContextReleaseData(smContextReleaseData).Execute()

Delete SM Context

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
    smContextId := "smContextId_example" // string | SM Context Resource ID
    smContextReleaseData := *openapiclient.NewSmContextReleaseData(*openapiclient.NewReleaseCause()) // SmContextReleaseData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextDocumentApi.Delete(context.Background(), smContextId).SmContextReleaseData(smContextReleaseData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextDocumentApi.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: SmContextReleasedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualSMContextDocumentApi.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextId** | **string** | SM Context Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextReleaseData** | [**SmContextReleaseData**](SmContextReleaseData.md) |  | 

### Return type

[**SmContextReleasedData**](SmContextReleasedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Deliver

> Deliver(ctx, smContextId).JsonData(jsonData).BinaryMoData(binaryMoData).Execute()

Deliver Uplink MO Data

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
    smContextId := "smContextId_example" // string | SM Context Resource ID
    jsonData := *openapiclient.NewDeliverReqData(*openapiclient.NewRefToBinaryData("ContentId_example")) // DeliverReqData |  (optional)
    binaryMoData := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextDocumentApi.Deliver(context.Background(), smContextId).JsonData(jsonData).BinaryMoData(binaryMoData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextDocumentApi.Deliver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextId** | **string** | SM Context Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeliverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**DeliverReqData**](DeliverReqData.md) |  | 
 **binaryMoData** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, smContextId).SmContextUpdateData(smContextUpdateData).Execute()

Update SM Context

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
    smContextId := "smContextId_example" // string | SM Context Resource ID
    smContextUpdateData := *openapiclient.NewSmContextUpdateData() // SmContextUpdateData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextDocumentApi.Update(context.Background(), smContextId).SmContextUpdateData(smContextUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextDocumentApi.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextId** | **string** | SM Context Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextUpdateData** | [**SmContextUpdateData**](SmContextUpdateData.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

