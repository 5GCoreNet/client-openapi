# \IndividualPDUSessionHSMFOrSMFApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleasePduSession**](IndividualPDUSessionHSMFOrSMFApi.md#ReleasePduSession) | **Post** /pdu-sessions/{pduSessionRef}/release | Release
[**RetrievePduSession**](IndividualPDUSessionHSMFOrSMFApi.md#RetrievePduSession) | **Post** /pdu-sessions/{pduSessionRef}/retrieve | Retrieve
[**TransferMoData**](IndividualPDUSessionHSMFOrSMFApi.md#TransferMoData) | **Post** /pdu-sessions/{pduSessionRef}/transfer-mo-data | Transfer MO Data
[**UpdatePduSession**](IndividualPDUSessionHSMFOrSMFApi.md#UpdatePduSession) | **Post** /pdu-sessions/{pduSessionRef}/modify | Update (initiated by V-SMF or I-SMF)



## ReleasePduSession

> ReleasedData ReleasePduSession(ctx, pduSessionRef).ReleaseData(releaseData).Execute()

Release

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
    pduSessionRef := "pduSessionRef_example" // string | PDU session reference
    releaseData := *openapiclient.NewReleaseData() // ReleaseData | data sent to H-SMF or SMF when releasing the PDU session (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPDUSessionHSMFOrSMFApi.ReleasePduSession(context.Background(), pduSessionRef).ReleaseData(releaseData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPDUSessionHSMFOrSMFApi.ReleasePduSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReleasePduSession`: ReleasedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualPDUSessionHSMFOrSMFApi.ReleasePduSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pduSessionRef** | **string** | PDU session reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleasePduSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseData** | [**ReleaseData**](ReleaseData.md) | data sent to H-SMF or SMF when releasing the PDU session | 

### Return type

[**ReleasedData**](ReleasedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrievePduSession

> RetrievedData RetrievePduSession(ctx, pduSessionRef).RetrieveData(retrieveData).Execute()

Retrieve

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
    pduSessionRef := "pduSessionRef_example" // string | PDU session reference
    retrieveData := *openapiclient.NewRetrieveData() // RetrieveData | representation of the payload of the Retrieve Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPDUSessionHSMFOrSMFApi.RetrievePduSession(context.Background(), pduSessionRef).RetrieveData(retrieveData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPDUSessionHSMFOrSMFApi.RetrievePduSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrievePduSession`: RetrievedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualPDUSessionHSMFOrSMFApi.RetrievePduSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pduSessionRef** | **string** | PDU session reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrievePduSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **retrieveData** | [**RetrieveData**](RetrieveData.md) | representation of the payload of the Retrieve Request | 

### Return type

[**RetrievedData**](RetrievedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferMoData

> TransferMoData(ctx, pduSessionRef).JsonData(jsonData).BinaryMoData(binaryMoData).Execute()

Transfer MO Data

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
    pduSessionRef := "pduSessionRef_example" // string | PDU session reference
    jsonData := *openapiclient.NewTransferMoDataReqData(*openapiclient.NewRefToBinaryData("ContentId_example")) // TransferMoDataReqData |  (optional)
    binaryMoData := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPDUSessionHSMFOrSMFApi.TransferMoData(context.Background(), pduSessionRef).JsonData(jsonData).BinaryMoData(binaryMoData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPDUSessionHSMFOrSMFApi.TransferMoData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pduSessionRef** | **string** | PDU session reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferMoDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**TransferMoDataReqData**](TransferMoDataReqData.md) |  | 
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


## UpdatePduSession

> HsmfUpdatedData UpdatePduSession(ctx, pduSessionRef).HsmfUpdateData(hsmfUpdateData).Execute()

Update (initiated by V-SMF or I-SMF)

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
    pduSessionRef := "pduSessionRef_example" // string | PDU session reference
    hsmfUpdateData := *openapiclient.NewHsmfUpdateData(*openapiclient.NewRequestIndication()) // HsmfUpdateData | representation of the updates to apply to the PDU session

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualPDUSessionHSMFOrSMFApi.UpdatePduSession(context.Background(), pduSessionRef).HsmfUpdateData(hsmfUpdateData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualPDUSessionHSMFOrSMFApi.UpdatePduSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePduSession`: HsmfUpdatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualPDUSessionHSMFOrSMFApi.UpdatePduSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pduSessionRef** | **string** | PDU session reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePduSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hsmfUpdateData** | [**HsmfUpdateData**](HsmfUpdateData.md) | representation of the updates to apply to the PDU session | 

### Return type

[**HsmfUpdatedData**](HsmfUpdatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

