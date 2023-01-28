# \IndividualUeContextDocumentApi

All URIs are relative to *https://example.com/namf-comm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRelocateUEContext**](IndividualUeContextDocumentApi.md#CancelRelocateUEContext) | **Post** /ue-contexts/{ueContextId}/cancel-relocate | Namf_Communication CancelRelocateUEContext service Operation
[**CreateUEContext**](IndividualUeContextDocumentApi.md#CreateUEContext) | **Put** /ue-contexts/{ueContextId} | Namf_Communication CreateUEContext service Operation
[**EBIAssignment**](IndividualUeContextDocumentApi.md#EBIAssignment) | **Post** /ue-contexts/{ueContextId}/assign-ebi | Namf_Communication EBI Assignment service Operation
[**RegistrationStatusUpdate**](IndividualUeContextDocumentApi.md#RegistrationStatusUpdate) | **Post** /ue-contexts/{ueContextId}/transfer-update | Namf_Communication RegistrationStatusUpdate service Operation
[**ReleaseUEContext**](IndividualUeContextDocumentApi.md#ReleaseUEContext) | **Post** /ue-contexts/{ueContextId}/release | Namf_Communication ReleaseUEContext service Operation
[**RelocateUEContext**](IndividualUeContextDocumentApi.md#RelocateUEContext) | **Post** /ue-contexts/{ueContextId}/relocate | Namf_Communication RelocateUEContext service Operation
[**UEContextTransfer**](IndividualUeContextDocumentApi.md#UEContextTransfer) | **Post** /ue-contexts/{ueContextId}/transfer | Namf_Communication UEContextTransfer service Operation



## CancelRelocateUEContext

> CancelRelocateUEContext(ctx, ueContextId).JsonData(jsonData).BinaryDataGtpcMessage(binaryDataGtpcMessage).Execute()

Namf_Communication CancelRelocateUEContext service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    jsonData := *openapiclient.NewUeContextCancelRelocateData(*openapiclient.NewRefToBinaryData("ContentId_example")) // UeContextCancelRelocateData |  (optional)
    binaryDataGtpcMessage := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.CancelRelocateUEContext(context.Background(), ueContextId).JsonData(jsonData).BinaryDataGtpcMessage(binaryDataGtpcMessage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.CancelRelocateUEContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRelocateUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**UeContextCancelRelocateData**](UeContextCancelRelocateData.md) |  | 
 **binaryDataGtpcMessage** | ***os.File** |  | 

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


## CreateUEContext

> UeContextCreatedData CreateUEContext(ctx, ueContextId).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).BinaryDataN2InformationExt1(binaryDataN2InformationExt1).BinaryDataN2InformationExt2(binaryDataN2InformationExt2).BinaryDataN2InformationExt3(binaryDataN2InformationExt3).BinaryDataN2InformationExt4(binaryDataN2InformationExt4).BinaryDataN2InformationExt5(binaryDataN2InformationExt5).BinaryDataN2InformationExt6(binaryDataN2InformationExt6).BinaryDataN2InformationExt7(binaryDataN2InformationExt7).BinaryDataN2InformationExt8(binaryDataN2InformationExt8).BinaryDataN2InformationExt9(binaryDataN2InformationExt9).BinaryDataN2InformationExt10(binaryDataN2InformationExt10).BinaryDataN2InformationExt11(binaryDataN2InformationExt11).BinaryDataN2InformationExt12(binaryDataN2InformationExt12).BinaryDataN2InformationExt13(binaryDataN2InformationExt13).BinaryDataN2InformationExt14(binaryDataN2InformationExt14).BinaryDataN2InformationExt15(binaryDataN2InformationExt15).BinaryDataN2InformationExt16(binaryDataN2InformationExt16).BinaryDataN2InformationExt17(binaryDataN2InformationExt17).Execute()

Namf_Communication CreateUEContext service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    jsonData := *openapiclient.NewUeContextCreateData(*openapiclient.NewUeContext(), *openapiclient.NewNgRanTargetId(openapiclient.GlobalRanNodeId{Interface{}: new(interface{})}, *openapiclient.NewTai(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example"), "Tac_example")), *openapiclient.NewN2InfoContent(*openapiclient.NewRefToBinaryData("ContentId_example")), []openapiclient.N2SmInformation{*openapiclient.NewN2SmInformation(int32(123))}) // UeContextCreateData |  (optional)
    binaryDataN2Information := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt1 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt2 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt3 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt4 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt5 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt6 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt7 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt8 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt9 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt10 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt11 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt12 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt13 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt14 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt15 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt16 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt17 := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.CreateUEContext(context.Background(), ueContextId).JsonData(jsonData).BinaryDataN2Information(binaryDataN2Information).BinaryDataN2InformationExt1(binaryDataN2InformationExt1).BinaryDataN2InformationExt2(binaryDataN2InformationExt2).BinaryDataN2InformationExt3(binaryDataN2InformationExt3).BinaryDataN2InformationExt4(binaryDataN2InformationExt4).BinaryDataN2InformationExt5(binaryDataN2InformationExt5).BinaryDataN2InformationExt6(binaryDataN2InformationExt6).BinaryDataN2InformationExt7(binaryDataN2InformationExt7).BinaryDataN2InformationExt8(binaryDataN2InformationExt8).BinaryDataN2InformationExt9(binaryDataN2InformationExt9).BinaryDataN2InformationExt10(binaryDataN2InformationExt10).BinaryDataN2InformationExt11(binaryDataN2InformationExt11).BinaryDataN2InformationExt12(binaryDataN2InformationExt12).BinaryDataN2InformationExt13(binaryDataN2InformationExt13).BinaryDataN2InformationExt14(binaryDataN2InformationExt14).BinaryDataN2InformationExt15(binaryDataN2InformationExt15).BinaryDataN2InformationExt16(binaryDataN2InformationExt16).BinaryDataN2InformationExt17(binaryDataN2InformationExt17).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.CreateUEContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUEContext`: UeContextCreatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.CreateUEContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**UeContextCreateData**](UeContextCreateData.md) |  | 
 **binaryDataN2Information** | ***os.File** |  | 
 **binaryDataN2InformationExt1** | ***os.File** |  | 
 **binaryDataN2InformationExt2** | ***os.File** |  | 
 **binaryDataN2InformationExt3** | ***os.File** |  | 
 **binaryDataN2InformationExt4** | ***os.File** |  | 
 **binaryDataN2InformationExt5** | ***os.File** |  | 
 **binaryDataN2InformationExt6** | ***os.File** |  | 
 **binaryDataN2InformationExt7** | ***os.File** |  | 
 **binaryDataN2InformationExt8** | ***os.File** |  | 
 **binaryDataN2InformationExt9** | ***os.File** |  | 
 **binaryDataN2InformationExt10** | ***os.File** |  | 
 **binaryDataN2InformationExt11** | ***os.File** |  | 
 **binaryDataN2InformationExt12** | ***os.File** |  | 
 **binaryDataN2InformationExt13** | ***os.File** |  | 
 **binaryDataN2InformationExt14** | ***os.File** |  | 
 **binaryDataN2InformationExt15** | ***os.File** |  | 
 **binaryDataN2InformationExt16** | ***os.File** |  | 
 **binaryDataN2InformationExt17** | ***os.File** |  | 

### Return type

[**UeContextCreatedData**](UeContextCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EBIAssignment

> AssignedEbiData EBIAssignment(ctx, ueContextId).AssignEbiData(assignEbiData).Execute()

Namf_Communication EBI Assignment service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    assignEbiData := *openapiclient.NewAssignEbiData(int32(123)) // AssignEbiData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.EBIAssignment(context.Background(), ueContextId).AssignEbiData(assignEbiData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.EBIAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EBIAssignment`: AssignedEbiData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.EBIAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiEBIAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignEbiData** | [**AssignEbiData**](AssignEbiData.md) |  | 

### Return type

[**AssignedEbiData**](AssignedEbiData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegistrationStatusUpdate

> UeRegStatusUpdateRspData RegistrationStatusUpdate(ctx, ueContextId).UeRegStatusUpdateReqData(ueRegStatusUpdateReqData).Execute()

Namf_Communication RegistrationStatusUpdate service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    ueRegStatusUpdateReqData := *openapiclient.NewUeRegStatusUpdateReqData(*openapiclient.NewUeContextTransferStatus()) // UeRegStatusUpdateReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.RegistrationStatusUpdate(context.Background(), ueContextId).UeRegStatusUpdateReqData(ueRegStatusUpdateReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.RegistrationStatusUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegistrationStatusUpdate`: UeRegStatusUpdateRspData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.RegistrationStatusUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegistrationStatusUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueRegStatusUpdateReqData** | [**UeRegStatusUpdateReqData**](UeRegStatusUpdateReqData.md) |  | 

### Return type

[**UeRegStatusUpdateRspData**](UeRegStatusUpdateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseUEContext

> ReleaseUEContext(ctx, ueContextId).UEContextRelease(uEContextRelease).Execute()

Namf_Communication ReleaseUEContext service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    uEContextRelease := *openapiclient.NewUEContextRelease(*openapiclient.NewNgApCause(int32(123), int32(123))) // UEContextRelease | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.ReleaseUEContext(context.Background(), ueContextId).UEContextRelease(uEContextRelease).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.ReleaseUEContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uEContextRelease** | [**UEContextRelease**](UEContextRelease.md) |  | 

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


## RelocateUEContext

> UeContextRelocatedData RelocateUEContext(ctx, ueContextId).JsonData(jsonData).BinaryDataForwardRelocationRequest(binaryDataForwardRelocationRequest).BinaryDataN2Information(binaryDataN2Information).BinaryDataN2InformationExt1(binaryDataN2InformationExt1).BinaryDataN2InformationExt2(binaryDataN2InformationExt2).BinaryDataN2InformationExt3(binaryDataN2InformationExt3).BinaryDataN2InformationExt4(binaryDataN2InformationExt4).BinaryDataN2InformationExt5(binaryDataN2InformationExt5).BinaryDataN2InformationExt6(binaryDataN2InformationExt6).BinaryDataN2InformationExt7(binaryDataN2InformationExt7).BinaryDataN2InformationExt8(binaryDataN2InformationExt8).BinaryDataN2InformationExt9(binaryDataN2InformationExt9).BinaryDataN2InformationExt10(binaryDataN2InformationExt10).BinaryDataN2InformationExt11(binaryDataN2InformationExt11).BinaryDataN2InformationExt12(binaryDataN2InformationExt12).BinaryDataN2InformationExt13(binaryDataN2InformationExt13).BinaryDataN2InformationExt14(binaryDataN2InformationExt14).BinaryDataN2InformationExt15(binaryDataN2InformationExt15).BinaryDataN2InformationExt16(binaryDataN2InformationExt16).Execute()

Namf_Communication RelocateUEContext service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    jsonData := *openapiclient.NewUeContextRelocateData(*openapiclient.NewUeContext(), *openapiclient.NewNgRanTargetId(openapiclient.GlobalRanNodeId{Interface{}: new(interface{})}, *openapiclient.NewTai(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example"), "Tac_example")), *openapiclient.NewN2InfoContent(*openapiclient.NewRefToBinaryData("ContentId_example")), *openapiclient.NewRefToBinaryData("ContentId_example")) // UeContextRelocateData |  (optional)
    binaryDataForwardRelocationRequest := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2Information := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt1 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt2 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt3 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt4 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt5 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt6 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt7 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt8 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt9 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt10 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt11 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt12 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt13 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt14 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt15 := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2InformationExt16 := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.RelocateUEContext(context.Background(), ueContextId).JsonData(jsonData).BinaryDataForwardRelocationRequest(binaryDataForwardRelocationRequest).BinaryDataN2Information(binaryDataN2Information).BinaryDataN2InformationExt1(binaryDataN2InformationExt1).BinaryDataN2InformationExt2(binaryDataN2InformationExt2).BinaryDataN2InformationExt3(binaryDataN2InformationExt3).BinaryDataN2InformationExt4(binaryDataN2InformationExt4).BinaryDataN2InformationExt5(binaryDataN2InformationExt5).BinaryDataN2InformationExt6(binaryDataN2InformationExt6).BinaryDataN2InformationExt7(binaryDataN2InformationExt7).BinaryDataN2InformationExt8(binaryDataN2InformationExt8).BinaryDataN2InformationExt9(binaryDataN2InformationExt9).BinaryDataN2InformationExt10(binaryDataN2InformationExt10).BinaryDataN2InformationExt11(binaryDataN2InformationExt11).BinaryDataN2InformationExt12(binaryDataN2InformationExt12).BinaryDataN2InformationExt13(binaryDataN2InformationExt13).BinaryDataN2InformationExt14(binaryDataN2InformationExt14).BinaryDataN2InformationExt15(binaryDataN2InformationExt15).BinaryDataN2InformationExt16(binaryDataN2InformationExt16).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.RelocateUEContext``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RelocateUEContext`: UeContextRelocatedData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.RelocateUEContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRelocateUEContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jsonData** | [**UeContextRelocateData**](UeContextRelocateData.md) |  | 
 **binaryDataForwardRelocationRequest** | ***os.File** |  | 
 **binaryDataN2Information** | ***os.File** |  | 
 **binaryDataN2InformationExt1** | ***os.File** |  | 
 **binaryDataN2InformationExt2** | ***os.File** |  | 
 **binaryDataN2InformationExt3** | ***os.File** |  | 
 **binaryDataN2InformationExt4** | ***os.File** |  | 
 **binaryDataN2InformationExt5** | ***os.File** |  | 
 **binaryDataN2InformationExt6** | ***os.File** |  | 
 **binaryDataN2InformationExt7** | ***os.File** |  | 
 **binaryDataN2InformationExt8** | ***os.File** |  | 
 **binaryDataN2InformationExt9** | ***os.File** |  | 
 **binaryDataN2InformationExt10** | ***os.File** |  | 
 **binaryDataN2InformationExt11** | ***os.File** |  | 
 **binaryDataN2InformationExt12** | ***os.File** |  | 
 **binaryDataN2InformationExt13** | ***os.File** |  | 
 **binaryDataN2InformationExt14** | ***os.File** |  | 
 **binaryDataN2InformationExt15** | ***os.File** |  | 
 **binaryDataN2InformationExt16** | ***os.File** |  | 

### Return type

[**UeContextRelocatedData**](UeContextRelocatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UEContextTransfer

> UeContextTransferRspData UEContextTransfer(ctx, ueContextId).UeContextTransferReqData(ueContextTransferReqData).Execute()

Namf_Communication UEContextTransfer service Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Namf_Communication"
)

func main() {
    ueContextId := "ueContextId_example" // string | UE Context Identifier
    ueContextTransferReqData := *openapiclient.NewUeContextTransferReqData(*openapiclient.NewTransferReason(), openapiclient.AccessType("3GPP_ACCESS")) // UeContextTransferReqData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualUeContextDocumentApi.UEContextTransfer(context.Background(), ueContextId).UeContextTransferReqData(ueContextTransferReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualUeContextDocumentApi.UEContextTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UEContextTransfer`: UeContextTransferRspData
    fmt.Fprintf(os.Stdout, "Response from `IndividualUeContextDocumentApi.UEContextTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ueContextId** | **string** | UE Context Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiUEContextTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ueContextTransferReqData** | [**UeContextTransferReqData**](UeContextTransferReqData.md) |  | 

### Return type

[**UeContextTransferRspData**](UeContextTransferRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

