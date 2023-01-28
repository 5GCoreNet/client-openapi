# \MBSSessionsCollectionApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSSession**](MBSSessionsCollectionApi.md#CreateMBSSession) | **Post** /mbs-sessions | Request the creation of a new MBS Session.



## CreateMBSSession

> MbsSessionCreateRsp CreateMBSSession(ctx).MbsSessionCreateReq(mbsSessionCreateReq).Execute()

Request the creation of a new MBS Session.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsSessionCreateReq := *openapiclient.NewMbsSessionCreateReq("AfId_example", *openapiclient.NewMbsSession(*openapiclient.NewMbsServiceType())) // MbsSessionCreateReq | Representation of the new MBS session to be created at the NEF.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSSessionsCollectionApi.CreateMBSSession(context.Background()).MbsSessionCreateReq(mbsSessionCreateReq).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSSessionsCollectionApi.CreateMBSSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSSession`: MbsSessionCreateRsp
    fmt.Fprintf(os.Stdout, "Response from `MBSSessionsCollectionApi.CreateMBSSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mbsSessionCreateReq** | [**MbsSessionCreateReq**](MbsSessionCreateReq.md) | Representation of the new MBS session to be created at the NEF. | 

### Return type

[**MbsSessionCreateRsp**](MbsSessionCreateRsp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

