# \IndividualSMContextApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleaseSmContext**](IndividualSMContextApi.md#ReleaseSmContext) | **Post** /sm-contexts/{smContextRef}/release | Release SM Context
[**RetrieveSmContext**](IndividualSMContextApi.md#RetrieveSmContext) | **Post** /sm-contexts/{smContextRef}/retrieve | Retrieve SM Context
[**SendMoData**](IndividualSMContextApi.md#SendMoData) | **Post** /sm-contexts/{smContextRef}/send-mo-data | Send MO Data
[**UpdateSmContext**](IndividualSMContextApi.md#UpdateSmContext) | **Post** /sm-contexts/{smContextRef}/modify | Update SM Context



## ReleaseSmContext

> SmContextReleasedData ReleaseSmContext(ctx, smContextRef).SmContextReleaseData(smContextReleaseData).Execute()

Release SM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_PDUSession"
)

func main() {
    smContextRef := "smContextRef_example" // string | SM context reference
    smContextReleaseData := *openapiclient.NewSmContextReleaseData() // SmContextReleaseData | representation of the data to be sent to the SMF when releasing the SM context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextApi.ReleaseSmContext(context.Background(), smContextRef).SmContextReleaseData(smContextReleaseData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextApi.ReleaseSmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleaseSmContext`: SmContextReleasedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualSMContextApi.ReleaseSmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextRef** | **string** | SM context reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseSmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextReleaseData** | [**SmContextReleaseData**](SmContextReleaseData.md) | representation of the data to be sent to the SMF when releasing the SM context | 

### Return type

[**SmContextReleasedData**](SmContextReleasedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveSmContext

> SmContextRetrievedData RetrieveSmContext(ctx, smContextRef).SmContextRetrieveData(smContextRetrieveData).Execute()

Retrieve SM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_PDUSession"
)

func main() {
    smContextRef := "smContextRef_example" // string | SM context reference
    smContextRetrieveData := *openapiclient.NewSmContextRetrieveData() // SmContextRetrieveData | parameters used to retrieve the SM context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextApi.RetrieveSmContext(context.Background(), smContextRef).SmContextRetrieveData(smContextRetrieveData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextApi.RetrieveSmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveSmContext`: SmContextRetrievedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualSMContextApi.RetrieveSmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextRef** | **string** | SM context reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveSmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextRetrieveData** | [**SmContextRetrieveData**](SmContextRetrieveData.md) | parameters used to retrieve the SM context | 

### Return type

[**SmContextRetrievedData**](SmContextRetrievedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMoData

> SendMoData(ctx, smContextRef).JsonData(jsonData).BinaryMoData(binaryMoData).Execute()

Send MO Data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_PDUSession"
)

func main() {
    smContextRef := "smContextRef_example" // string | SM context reference
    jsonData := *openapiclient.NewSendMoDataReqData(*openapiclient.NewRefToBinaryData("ContentId_example")) // SendMoDataReqData |  (optional)
    binaryMoData := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextApi.SendMoData(context.Background(), smContextRef).JsonData(jsonData).BinaryMoData(binaryMoData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextApi.SendMoData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextRef** | **string** | SM context reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMoDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**SendMoDataReqData**](SendMoDataReqData.md) |  | 
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


## UpdateSmContext

> SmContextUpdatedData UpdateSmContext(ctx, smContextRef).SmContextUpdateData(smContextUpdateData).Execute()

Update SM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nsmf_PDUSession"
)

func main() {
    smContextRef := "smContextRef_example" // string | SM context reference
    smContextUpdateData := *openapiclient.NewSmContextUpdateData() // SmContextUpdateData | representation of the updates to apply to the SM context

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualSMContextApi.UpdateSmContext(context.Background(), smContextRef).SmContextUpdateData(smContextUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualSMContextApi.UpdateSmContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSmContext`: SmContextUpdatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualSMContextApi.UpdateSmContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smContextRef** | **string** | SM context reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSmContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextUpdateData** | [**SmContextUpdateData**](SmContextUpdateData.md) | representation of the updates to apply to the SM context | 

### Return type

[**SmContextUpdatedData**](SmContextUpdatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

