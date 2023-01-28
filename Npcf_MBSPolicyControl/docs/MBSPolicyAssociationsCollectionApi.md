# \MBSPolicyAssociationsCollectionApi

All URIs are relative to *https://example.com/npcf-mbspolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSPolicy**](MBSPolicyAssociationsCollectionApi.md#CreateMBSPolicy) | **Post** /mbs-policies | Request the creation of a new MBS Policy Association.



## CreateMBSPolicy

> MbsPolicyDecision CreateMBSPolicy(ctx).MbsPolicyCtxtData(mbsPolicyCtxtData).Execute()

Request the creation of a new MBS Policy Association.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_MBSPolicyControl"
)

func main() {
    mbsPolicyCtxtData := *openapiclient.NewMbsPolicyCtxtData(*openapiclient.NewMbsSessionId()) // MbsPolicyCtxtData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSPolicyAssociationsCollectionApi.CreateMBSPolicy(context.Background()).MbsPolicyCtxtData(mbsPolicyCtxtData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSPolicyAssociationsCollectionApi.CreateMBSPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSPolicy`: MbsPolicyDecision
    fmt.Fprintf(os.Stdout, "Response from `MBSPolicyAssociationsCollectionApi.CreateMBSPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mbsPolicyCtxtData** | [**MbsPolicyCtxtData**](MbsPolicyCtxtData.md) |  | 

### Return type

[**MbsPolicyDecision**](MbsPolicyDecision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

