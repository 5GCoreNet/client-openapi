# \MBSApplicationSessionContextsCollectionApi

All URIs are relative to *https://example.com/npcf-mbspolicyauth/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSAppSessionCtxt**](MBSApplicationSessionContextsCollectionApi.md#CreateMBSAppSessionCtxt) | **Post** / contexts | Request the creation of a new MBS Application Session Context.



## CreateMBSAppSessionCtxt

> MbsAppSessionCtxt CreateMBSAppSessionCtxt(ctx).MbsAppSessionCtxt(mbsAppSessionCtxt).Execute()

Request the creation of a new MBS Application Session Context.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_MBSPolicyAuthorization"
)

func main() {
    mbsAppSessionCtxt := *openapiclient.NewMbsAppSessionCtxt(*openapiclient.NewMbsSessionId()) // MbsAppSessionCtxt | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSApplicationSessionContextsCollectionApi.CreateMBSAppSessionCtxt(context.Background()).MbsAppSessionCtxt(mbsAppSessionCtxt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSApplicationSessionContextsCollectionApi.CreateMBSAppSessionCtxt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSAppSessionCtxt`: MbsAppSessionCtxt
    fmt.Fprintf(os.Stdout, "Response from `MBSApplicationSessionContextsCollectionApi.CreateMBSAppSessionCtxt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSAppSessionCtxtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mbsAppSessionCtxt** | [**MbsAppSessionCtxt**](MbsAppSessionCtxt.md) |  | 

### Return type

[**MbsAppSessionCtxt**](MbsAppSessionCtxt.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

