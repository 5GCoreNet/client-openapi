# \MBSSessionsCollectionApi

All URIs are relative to *https://example.com/nmbsmf-mbssession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContextUpdate**](MBSSessionsCollectionApi.md#ContextUpdate) | **Post** /mbs-sessions/contexts/update | ContextUpdate
[**Create**](MBSSessionsCollectionApi.md#Create) | **Post** /mbs-sessions | Create



## ContextUpdate

> ContextUpdateRspData ContextUpdate(ctx).ContextUpdateReqData(contextUpdateReqData).Execute()

ContextUpdate

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsmf_MBSSession"
)

func main() {
    contextUpdateReqData := *openapiclient.NewContextUpdateReqData("NfcInstanceId_example", *openapiclient.NewMbsSessionId()) // ContextUpdateReqData | Data within the ContextUpdate Request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSSessionsCollectionApi.ContextUpdate(context.Background()).ContextUpdateReqData(contextUpdateReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSSessionsCollectionApi.ContextUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContextUpdate`: ContextUpdateRspData
    fmt.Fprintf(os.Stdout, "Response from `MBSSessionsCollectionApi.ContextUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContextUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextUpdateReqData** | [**ContextUpdateReqData**](ContextUpdateReqData.md) | Data within the ContextUpdate Request | 

### Return type

[**ContextUpdateRspData**](ContextUpdateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json, multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> CreateRspData Create(ctx).CreateReqData(createReqData).Execute()

Create

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsmf_MBSSession"
)

func main() {
    createReqData := *openapiclient.NewCreateReqData(*openapiclient.NewExtMbsSession()) // CreateReqData | Representation of the MBS session to be created in the MB-SMF Creates an individual MBS session resource in the MB-SMF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSSessionsCollectionApi.Create(context.Background()).CreateReqData(createReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSSessionsCollectionApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: CreateRspData
    fmt.Fprintf(os.Stdout, "Response from `MBSSessionsCollectionApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReqData** | [**CreateReqData**](CreateReqData.md) | Representation of the MBS session to be created in the MB-SMF Creates an individual MBS session resource in the MB-SMF.  | 

### Return type

[**CreateRspData**](CreateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

